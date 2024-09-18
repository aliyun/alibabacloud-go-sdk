// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AbolishDeploymentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1606087c-9ac4-43f0-83a8-0b5ced21XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s AbolishDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentRequest) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentRequest) SetId(v string) *AbolishDeploymentRequest {
	s.Id = &v
	return s
}

func (s *AbolishDeploymentRequest) SetProjectId(v string) *AbolishDeploymentRequest {
	s.ProjectId = &v
	return s
}

type AbolishDeploymentResponseBody struct {
	// example:
	//
	// 55D786C9-DD57-524D-884C-C5239278XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbolishDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentResponseBody) SetRequestId(v string) *AbolishDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbolishDeploymentResponseBody) SetSuccess(v bool) *AbolishDeploymentResponseBody {
	s.Success = &v
	return s
}

type AbolishDeploymentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s AbolishDeploymentResponse) GoString() string {
	return s.String()
}

func (s *AbolishDeploymentResponse) SetHeaders(v map[string]*string) *AbolishDeploymentResponse {
	s.Headers = v
	return s
}

func (s *AbolishDeploymentResponse) SetStatusCode(v int32) *AbolishDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishDeploymentResponse) SetBody(v *AbolishDeploymentResponseBody) *AbolishDeploymentResponse {
	s.Body = v
	return s
}

type CreateDeploymentRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ObjectIds []*string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty" type:"Repeated"`
	// 项目Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentRequest) SetDescription(v string) *CreateDeploymentRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentRequest) SetObjectIds(v []*string) *CreateDeploymentRequest {
	s.ObjectIds = v
	return s
}

func (s *CreateDeploymentRequest) SetProjectId(v string) *CreateDeploymentRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeploymentRequest) SetType(v string) *CreateDeploymentRequest {
	s.Type = &v
	return s
}

type CreateDeploymentShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	ObjectIdsShrink *string `json:"ObjectIds,omitempty" xml:"ObjectIds,omitempty"`
	// 项目Id
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Online
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDeploymentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDeploymentShrinkRequest) SetDescription(v string) *CreateDeploymentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetObjectIdsShrink(v string) *CreateDeploymentShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetProjectId(v string) *CreateDeploymentShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDeploymentShrinkRequest) SetType(v string) *CreateDeploymentShrinkRequest {
	s.Type = &v
	return s
}

type CreateDeploymentResponseBody struct {
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponseBody) SetId(v string) *CreateDeploymentResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDeploymentResponseBody) SetRequestId(v string) *CreateDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type CreateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponse) SetHeaders(v map[string]*string) *CreateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentResponse) SetStatusCode(v int32) *CreateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentResponse) SetBody(v *CreateDeploymentResponseBody) *CreateDeploymentResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetProjectId(v string) *CreateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFunctionRequest) SetSpec(v string) *CreateFunctionRequest {
	s.Spec = &v
	return s
}

type CreateFunctionResponseBody struct {
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// AE49C88D-5BEE-5ADD-8B8C-C4BBC0D7XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetId(v string) *CreateFunctionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

type CreateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetStatusCode(v int32) *CreateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateNodeRequest struct {
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeRequest) SetContainerId(v string) *CreateNodeRequest {
	s.ContainerId = &v
	return s
}

func (s *CreateNodeRequest) SetProjectId(v string) *CreateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateNodeRequest) SetScene(v string) *CreateNodeRequest {
	s.Scene = &v
	return s
}

func (s *CreateNodeRequest) SetSpec(v string) *CreateNodeRequest {
	s.Spec = &v
	return s
}

type CreateNodeResponseBody struct {
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeResponseBody) SetId(v string) *CreateNodeResponseBody {
	s.Id = &v
	return s
}

func (s *CreateNodeResponseBody) SetRequestId(v string) *CreateNodeResponseBody {
	s.RequestId = &v
	return s
}

type CreateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeResponse) SetHeaders(v map[string]*string) *CreateNodeResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeResponse) SetStatusCode(v int32) *CreateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeResponse) SetBody(v *CreateNodeResponseBody) *CreateNodeResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// 资源文件的项目id
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetProjectId(v string) *CreateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateResourceRequest) SetSpec(v string) *CreateResourceRequest {
	s.Spec = &v
	return s
}

type CreateResourceResponseBody struct {
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// A5B97987-66EA-5563-9599-A2752292XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetId(v string) *CreateResourceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type CreateWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionRequest) SetProjectId(v string) *CreateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateWorkflowDefinitionRequest) SetSpec(v string) *CreateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type CreateWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0EF298E5-0940-5AC7-9CB0-65025070XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponseBody) SetId(v string) *CreateWorkflowDefinitionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateWorkflowDefinitionResponseBody) SetRequestId(v string) *CreateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

type CreateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *CreateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *CreateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetStatusCode(v int32) *CreateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWorkflowDefinitionResponse) SetBody(v *CreateWorkflowDefinitionResponseBody) *CreateWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type DeleteFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) SetId(v string) *DeleteFunctionRequest {
	s.Id = &v
	return s
}

func (s *DeleteFunctionRequest) SetProjectId(v string) *DeleteFunctionRequest {
	s.ProjectId = &v
	return s
}

type DeleteFunctionResponseBody struct {
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetSuccess(v bool) *DeleteFunctionResponseBody {
	s.Success = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetStatusCode(v int32) *DeleteFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResponse) SetBody(v *DeleteFunctionResponseBody) *DeleteFunctionResponse {
	s.Body = v
	return s
}

type DeleteNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeRequest) SetId(v string) *DeleteNodeRequest {
	s.Id = &v
	return s
}

func (s *DeleteNodeRequest) SetProjectId(v string) *DeleteNodeRequest {
	s.ProjectId = &v
	return s
}

type DeleteNodeResponseBody struct {
	// example:
	//
	// A1E54497-5122-505E-91C6-BAC14980XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponseBody) SetRequestId(v string) *DeleteNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeResponseBody) SetSuccess(v bool) *DeleteNodeResponseBody {
	s.Success = &v
	return s
}

type DeleteNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeResponse) SetHeaders(v map[string]*string) *DeleteNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeResponse) SetStatusCode(v int32) *DeleteNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeResponse) SetBody(v *DeleteNodeResponseBody) *DeleteNodeResponse {
	s.Body = v
	return s
}

type DeleteResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) SetId(v string) *DeleteResourceRequest {
	s.Id = &v
	return s
}

func (s *DeleteResourceRequest) SetProjectId(v string) *DeleteResourceRequest {
	s.ProjectId = &v
	return s
}

type DeleteResourceResponseBody struct {
	// example:
	//
	// 88198F19-A36B-52A9-AE44-4518A688XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceResponseBody) SetSuccess(v bool) *DeleteResourceResponseBody {
	s.Success = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type DeleteWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s DeleteWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionRequest) SetId(v string) *DeleteWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *DeleteWorkflowDefinitionRequest) SetProjectId(v string) *DeleteWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type DeleteWorkflowDefinitionResponseBody struct {
	// example:
	//
	// B17730C0-D959-548A-AE23-E754177CXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponseBody) SetRequestId(v string) *DeleteWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponseBody) SetSuccess(v bool) *DeleteWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type DeleteWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *DeleteWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetStatusCode(v int32) *DeleteWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWorkflowDefinitionResponse) SetBody(v *DeleteWorkflowDefinitionResponseBody) *DeleteWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type ExecDeploymentStageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ExecDeploymentStageRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageRequest) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageRequest) SetCode(v string) *ExecDeploymentStageRequest {
	s.Code = &v
	return s
}

func (s *ExecDeploymentStageRequest) SetId(v string) *ExecDeploymentStageRequest {
	s.Id = &v
	return s
}

func (s *ExecDeploymentStageRequest) SetProjectId(v string) *ExecDeploymentStageRequest {
	s.ProjectId = &v
	return s
}

type ExecDeploymentStageResponseBody struct {
	// example:
	//
	// AFBB799F-8578-51C5-A766-E922EDB8XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecDeploymentStageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageResponseBody) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageResponseBody) SetRequestId(v string) *ExecDeploymentStageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecDeploymentStageResponseBody) SetSuccess(v bool) *ExecDeploymentStageResponseBody {
	s.Success = &v
	return s
}

type ExecDeploymentStageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecDeploymentStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecDeploymentStageResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecDeploymentStageResponse) GoString() string {
	return s.String()
}

func (s *ExecDeploymentStageResponse) SetHeaders(v map[string]*string) *ExecDeploymentStageResponse {
	s.Headers = v
	return s
}

func (s *ExecDeploymentStageResponse) SetStatusCode(v int32) *ExecDeploymentStageResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecDeploymentStageResponse) SetBody(v *ExecDeploymentStageResponseBody) *ExecDeploymentStageResponse {
	s.Body = v
	return s
}

type GetDeploymentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetDeploymentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentRequest) SetId(v string) *GetDeploymentRequest {
	s.Id = &v
	return s
}

func (s *GetDeploymentRequest) SetProjectId(v string) *GetDeploymentRequest {
	s.ProjectId = &v
	return s
}

type GetDeploymentResponseBody struct {
	Pipeline *GetDeploymentResponseBodyPipeline `json:"Pipeline,omitempty" xml:"Pipeline,omitempty" type:"Struct"`
	// example:
	//
	// 08468352-032C-5262-AEDC-68C9FA05XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDeploymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBody) SetPipeline(v *GetDeploymentResponseBodyPipeline) *GetDeploymentResponseBody {
	s.Pipeline = v
	return s
}

func (s *GetDeploymentResponseBody) SetRequestId(v string) *GetDeploymentResponseBody {
	s.RequestId = &v
	return s
}

type GetDeploymentResponseBodyPipeline struct {
	// 发布包创建时间戳
	//
	// example:
	//
	// 1724984066000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 创建人
	//
	// example:
	//
	// 137946317766XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 发布流程Id
	//
	// example:
	//
	// a7ef0634-20ec-4a7c-a214-54020f91XXXX
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1724984066000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 56160
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 步骤详情
	Stages []*GetDeploymentResponseBodyPipelineStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// 发布流程状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeploymentResponseBodyPipeline) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBodyPipeline) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyPipeline) SetCreateTime(v int64) *GetDeploymentResponseBodyPipeline {
	s.CreateTime = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetCreator(v string) *GetDeploymentResponseBodyPipeline {
	s.Creator = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetId(v string) *GetDeploymentResponseBodyPipeline {
	s.Id = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetMessage(v string) *GetDeploymentResponseBodyPipeline {
	s.Message = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetModifyTime(v int64) *GetDeploymentResponseBodyPipeline {
	s.ModifyTime = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetProjectId(v string) *GetDeploymentResponseBodyPipeline {
	s.ProjectId = &v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetStages(v []*GetDeploymentResponseBodyPipelineStages) *GetDeploymentResponseBodyPipeline {
	s.Stages = v
	return s
}

func (s *GetDeploymentResponseBodyPipeline) SetStatus(v string) *GetDeploymentResponseBodyPipeline {
	s.Status = &v
	return s
}

type GetDeploymentResponseBodyPipelineStages struct {
	// 阶段代号
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 阶段描述
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	Detail      map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// 阶段信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 阶段名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 阶段状态
	//
	// example:
	//
	// INIT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 步骤
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// 阶段类型
	//
	// example:
	//
	// BUILD
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDeploymentResponseBodyPipelineStages) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponseBodyPipelineStages) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponseBodyPipelineStages) SetCode(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Code = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetDescription(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Description = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetDetail(v map[string]interface{}) *GetDeploymentResponseBodyPipelineStages {
	s.Detail = v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetMessage(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Message = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetName(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Name = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetStatus(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Status = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetStep(v int32) *GetDeploymentResponseBodyPipelineStages {
	s.Step = &v
	return s
}

func (s *GetDeploymentResponseBodyPipelineStages) SetType(v string) *GetDeploymentResponseBodyPipelineStages {
	s.Type = &v
	return s
}

type GetDeploymentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentResponse) SetHeaders(v map[string]*string) *GetDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentResponse) SetStatusCode(v int32) *GetDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentResponse) SetBody(v *GetDeploymentResponseBody) *GetDeploymentResponse {
	s.Body = v
	return s
}

type GetFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) SetId(v string) *GetFunctionRequest {
	s.Id = &v
	return s
}

func (s *GetFunctionRequest) SetProjectId(v string) *GetFunctionRequest {
	s.ProjectId = &v
	return s
}

type GetFunctionResponseBody struct {
	Function *GetFunctionResponseBodyFunction `json:"Function,omitempty" xml:"Function,omitempty" type:"Struct"`
	// example:
	//
	// 6CF95929-6D12-5A88-8CC3-4B2F4C2EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBody) SetFunction(v *GetFunctionResponseBodyFunction) *GetFunctionResponseBody {
	s.Function = v
	return s
}

func (s *GetFunctionResponseBody) SetRequestId(v string) *GetFunctionResponseBody {
	s.RequestId = &v
	return s
}

type GetFunctionResponseBodyFunction struct {
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1724506661000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetFunctionResponseBodyFunction) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBodyFunction) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBodyFunction) SetCreateTime(v int64) *GetFunctionResponseBodyFunction {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetId(v string) *GetFunctionResponseBodyFunction {
	s.Id = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetModifyTime(v int64) *GetFunctionResponseBodyFunction {
	s.ModifyTime = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetName(v string) *GetFunctionResponseBodyFunction {
	s.Name = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetOwner(v string) *GetFunctionResponseBodyFunction {
	s.Owner = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetProjectId(v string) *GetFunctionResponseBodyFunction {
	s.ProjectId = &v
	return s
}

func (s *GetFunctionResponseBodyFunction) SetSpec(v string) *GetFunctionResponseBodyFunction {
	s.Spec = &v
	return s
}

type GetFunctionResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetStatusCode(v int32) *GetFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResponse) SetBody(v *GetFunctionResponseBody) *GetFunctionResponse {
	s.Body = v
	return s
}

type GetNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNodeRequest) GoString() string {
	return s.String()
}

func (s *GetNodeRequest) SetId(v string) *GetNodeRequest {
	s.Id = &v
	return s
}

func (s *GetNodeRequest) SetProjectId(v string) *GetNodeRequest {
	s.ProjectId = &v
	return s
}

type GetNodeResponseBody struct {
	Node *GetNodeResponseBodyNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBody) SetNode(v *GetNodeResponseBodyNode) *GetNodeResponseBody {
	s.Node = v
	return s
}

func (s *GetNodeResponseBody) SetRequestId(v string) *GetNodeResponseBody {
	s.RequestId = &v
	return s
}

type GetNodeResponseBodyNode struct {
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1700539206000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 196596664824XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetNodeResponseBodyNode) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponseBodyNode) GoString() string {
	return s.String()
}

func (s *GetNodeResponseBodyNode) SetCreateTime(v int64) *GetNodeResponseBodyNode {
	s.CreateTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetId(v string) *GetNodeResponseBodyNode {
	s.Id = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetModifyTime(v int64) *GetNodeResponseBodyNode {
	s.ModifyTime = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetName(v string) *GetNodeResponseBodyNode {
	s.Name = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetOwner(v string) *GetNodeResponseBodyNode {
	s.Owner = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetProjectId(v string) *GetNodeResponseBodyNode {
	s.ProjectId = &v
	return s
}

func (s *GetNodeResponseBodyNode) SetSpec(v string) *GetNodeResponseBodyNode {
	s.Spec = &v
	return s
}

type GetNodeResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNodeResponse) GoString() string {
	return s.String()
}

func (s *GetNodeResponse) SetHeaders(v map[string]*string) *GetNodeResponse {
	s.Headers = v
	return s
}

func (s *GetNodeResponse) SetStatusCode(v int32) *GetNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeResponse) SetBody(v *GetNodeResponseBody) *GetNodeResponse {
	s.Body = v
	return s
}

type GetResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) SetId(v string) *GetResourceRequest {
	s.Id = &v
	return s
}

func (s *GetResourceRequest) SetProjectId(v string) *GetResourceRequest {
	s.ProjectId = &v
	return s
}

type GetResourceResponseBody struct {
	// example:
	//
	// E871F6C0-2EFF-5790-A00D-C57543EEXXXX
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resource  *GetResourceResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetResource(v *GetResourceResponseBodyResource) *GetResourceResponseBody {
	s.Resource = v
	return s
}

type GetResourceResponseBodyResource struct {
	// example:
	//
	// 1700539206000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1700539206000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetResourceResponseBodyResource) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyResource) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyResource) SetCreateTime(v int64) *GetResourceResponseBodyResource {
	s.CreateTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetId(v string) *GetResourceResponseBodyResource {
	s.Id = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetModifyTime(v int64) *GetResourceResponseBodyResource {
	s.ModifyTime = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetName(v string) *GetResourceResponseBodyResource {
	s.Name = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetOwner(v string) *GetResourceResponseBodyResource {
	s.Owner = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetProjectId(v string) *GetResourceResponseBodyResource {
	s.ProjectId = &v
	return s
}

func (s *GetResourceResponseBodyResource) SetSpec(v string) *GetResourceResponseBodyResource {
	s.Spec = &v
	return s
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type GetWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionRequest) SetId(v string) *GetWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionRequest) SetProjectId(v string) *GetWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type GetWorkflowDefinitionResponseBody struct {
	// example:
	//
	// F2BDD628-8A21-5BD1-B930-1A2D5989XXXX
	RequestId          *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WorkflowDefinition *GetWorkflowDefinitionResponseBodyWorkflowDefinition `json:"WorkflowDefinition,omitempty" xml:"WorkflowDefinition,omitempty" type:"Struct"`
}

func (s GetWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBody) SetRequestId(v string) *GetWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBody) SetWorkflowDefinition(v *GetWorkflowDefinitionResponseBodyWorkflowDefinition) *GetWorkflowDefinitionResponseBody {
	s.WorkflowDefinition = v
	return s
}

type GetWorkflowDefinitionResponseBodyWorkflowDefinition struct {
	// example:
	//
	// 1708481905000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1708481905000
	ModifyTime *int64  `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 307XXX
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Spec      *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponseBodyWorkflowDefinition) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetCreateTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.CreateTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Id = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetModifyTime(v int64) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ModifyTime = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetName(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Name = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetOwner(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Owner = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetProjectId(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.ProjectId = &v
	return s
}

func (s *GetWorkflowDefinitionResponseBodyWorkflowDefinition) SetSpec(v string) *GetWorkflowDefinitionResponseBodyWorkflowDefinition {
	s.Spec = &v
	return s
}

type GetWorkflowDefinitionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *GetWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *GetWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetStatusCode(v int32) *GetWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkflowDefinitionResponse) SetBody(v *GetWorkflowDefinitionResponseBody) *GetWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type ListDeploymentsRequest struct {
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentsRequest) SetCreator(v string) *ListDeploymentsRequest {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageNumber(v int32) *ListDeploymentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsRequest) SetPageSize(v int32) *ListDeploymentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsRequest) SetProjectId(v string) *ListDeploymentsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsRequest) SetStatus(v string) *ListDeploymentsRequest {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBody struct {
	PagingInfo *ListDeploymentsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDeploymentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBody) SetPagingInfo(v *ListDeploymentsResponseBodyPagingInfo) *ListDeploymentsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDeploymentsResponseBody) SetRequestId(v string) *ListDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfo struct {
	Deployments []*ListDeploymentsResponseBodyPagingInfoDeployments `json:"Deployments,omitempty" xml:"Deployments,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2524
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetDeployments(v []*ListDeploymentsResponseBodyPagingInfoDeployments) *ListDeploymentsResponseBodyPagingInfo {
	s.Deployments = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetPageNumber(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetPageSize(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfo) SetTotalCount(v string) *ListDeploymentsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfoDeployments struct {
	// 发布包创建时间戳
	//
	// example:
	//
	// 1702736654000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 创建人
	//
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 发布流程Id
	//
	// example:
	//
	// ddf354a5-03df-48fc-94c1-cc973f79XXXX
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1702736654000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 项目Id
	//
	// example:
	//
	// 44683
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 步骤详情
	Stages []*ListDeploymentsResponseBodyPagingInfoDeploymentsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// 发布流程状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfoDeployments) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfoDeployments) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetCreateTime(v int64) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.CreateTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetCreator(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Creator = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetId(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Id = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetMessage(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Message = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetModifyTime(v int64) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.ModifyTime = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetProjectId(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.ProjectId = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetStages(v []*ListDeploymentsResponseBodyPagingInfoDeploymentsStages) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Stages = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeployments) SetStatus(v string) *ListDeploymentsResponseBodyPagingInfoDeployments {
	s.Status = &v
	return s
}

type ListDeploymentsResponseBodyPagingInfoDeploymentsStages struct {
	// 阶段代号
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 阶段描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 阶段详细信息
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// 阶段信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 阶段名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 阶段状态
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 步骤
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// 阶段类型
	//
	// example:
	//
	// CHECK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDeploymentsResponseBodyPagingInfoDeploymentsStages) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponseBodyPagingInfoDeploymentsStages) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetCode(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Code = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetDescription(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Description = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetDetail(v map[string]interface{}) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Detail = v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetMessage(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Message = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetName(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Name = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetStatus(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Status = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetStep(v int32) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Step = &v
	return s
}

func (s *ListDeploymentsResponseBodyPagingInfoDeploymentsStages) SetType(v string) *ListDeploymentsResponseBodyPagingInfoDeploymentsStages {
	s.Type = &v
	return s
}

type ListDeploymentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeploymentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeploymentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeploymentsResponse) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponse) SetHeaders(v map[string]*string) *ListDeploymentsResponse {
	s.Headers = v
	return s
}

func (s *ListDeploymentsResponse) SetStatusCode(v int32) *ListDeploymentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeploymentsResponse) SetBody(v *ListDeploymentsResponseBody) *ListDeploymentsResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetOwner(v string) *ListFunctionsRequest {
	s.Owner = &v
	return s
}

func (s *ListFunctionsRequest) SetPageNumber(v int32) *ListFunctionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsRequest) SetPageSize(v int32) *ListFunctionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsRequest) SetProjectId(v string) *ListFunctionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsRequest) SetType(v string) *ListFunctionsRequest {
	s.Type = &v
	return s
}

type ListFunctionsResponseBody struct {
	PagingInfo *ListFunctionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 89FB2BF0-EB00-5D03-9C34-05931001XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetPagingInfo(v *ListFunctionsResponseBodyPagingInfo) *ListFunctionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

type ListFunctionsResponseBodyPagingInfo struct {
	Functions []*ListFunctionsResponseBodyPagingInfoFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfo) SetFunctions(v []*ListFunctionsResponseBodyPagingInfoFunctions) *ListFunctionsResponseBodyPagingInfo {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetPageSize(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListFunctionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctions struct {
	// ARM集群资源文件列表
	//
	// example:
	//
	// xxx.jar,yyy.jar
	ArmResource *string `json:"ArmResource,omitempty" xml:"ArmResource,omitempty"`
	// 函数实现类名
	//
	// example:
	//
	// com.demo.Main
	ClassName *string `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	// 命令描述
	//
	// example:
	//
	// testUdf(xx,yy)
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// 1655953028000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 函数注册到的数据源信息
	DataSource *ListFunctionsResponseBodyPagingInfoFunctionsDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 数据库名，可选
	//
	// example:
	//
	// odps_first
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// 对funciotn的描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 嵌套函数代码内容
	//
	// example:
	//
	// print(\\"hello,world!\\")
	EmbeddedCode *string `json:"EmbeddedCode,omitempty" xml:"EmbeddedCode,omitempty"`
	// 嵌套代码类型
	//
	// example:
	//
	// Python2
	EmbeddedCodeType *string `json:"EmbeddedCodeType,omitempty" xml:"EmbeddedCodeType,omitempty"`
	// 嵌套资源类型
	//
	// example:
	//
	// File
	EmbeddedResourceType *string `json:"EmbeddedResourceType,omitempty" xml:"EmbeddedResourceType,omitempty"`
	// 示例说明
	ExampleDescription *string `json:"ExampleDescription,omitempty" xml:"ExampleDescription,omitempty"`
	// 函数的实现代码
	//
	// example:
	//
	// xxx.jar,yyy.jar
	FileResource *string `json:"FileResource,omitempty" xml:"FileResource,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 580667964888595XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1655953028000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 代表资源名称的资源属性字段
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 函数责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 命令描述
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// 项目Id
	//
	// example:
	//
	// 307XXX
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 返回值说明
	ReturnValueDescription *string `json:"ReturnValueDescription,omitempty" xml:"ReturnValueDescription,omitempty"`
	// 运行时资源组信息
	RuntimeResource *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// 工作流的脚本信息
	Script *ListFunctionsResponseBodyPagingInfoFunctionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 函数类型
	//
	// example:
	//
	// MATH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctions) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetArmResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ArmResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetClassName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ClassName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCommandDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CommandDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetCreateTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDataSource(v *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DataSource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDatabaseName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.DatabaseName = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Description = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCode(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCode = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedCodeType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedCodeType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetEmbeddedResourceType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.EmbeddedResourceType = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetExampleDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ExampleDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetFileResource(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.FileResource = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetModifyTime(v int64) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ModifyTime = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetOwner(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Owner = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetParameterDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ParameterDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetProjectId(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ProjectId = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetReturnValueDescription(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.ReturnValueDescription = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetRuntimeResource(v *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.RuntimeResource = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetScript(v *ListFunctionsResponseBodyPagingInfoFunctionsScript) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Script = v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctions) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctions {
	s.Type = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsDataSource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetName(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Name = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsDataSource) SetType(v string) *ListFunctionsResponseBodyPagingInfoFunctionsDataSource {
	s.Type = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource struct {
	// 运行时资源组Id
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource) SetResourceGroupId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsScript struct {
	// 脚本的id
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScript) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetId(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Id = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetPath(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Path = &v
	return s
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScript) SetRuntime(v *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) *ListFunctionsResponseBodyPagingInfoFunctionsScript {
	s.Runtime = v
	return s
}

type ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_FUNCTION
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime) SetCommand(v string) *ListFunctionsResponseBodyPagingInfoFunctionsScriptRuntime {
	s.Command = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListNodeDependenciesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListNodeDependenciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesRequest) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesRequest) SetId(v string) *ListNodeDependenciesRequest {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageNumber(v int32) *ListNodeDependenciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetPageSize(v int32) *ListNodeDependenciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesRequest) SetProjectId(v string) *ListNodeDependenciesRequest {
	s.ProjectId = &v
	return s
}

type ListNodeDependenciesResponseBody struct {
	PagingInfo *ListNodeDependenciesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 204EAF68-CCE3-5112-8DA0-E7A60F02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeDependenciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBody) SetPagingInfo(v *ListNodeDependenciesResponseBodyPagingInfo) *ListNodeDependenciesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodeDependenciesResponseBody) SetRequestId(v string) *ListNodeDependenciesResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfo struct {
	Nodes []*ListNodeDependenciesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 90
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetNodes(v []*ListNodeDependenciesResponseBodyPagingInfoNodes) *ListNodeDependenciesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetPageSize(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodeDependenciesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodes struct {
	// 节点的创建时间
	//
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 数据源信息
	DataSource *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 723932906364267XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 节点输入
	Inputs *ListNodeDependenciesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// 属性修改时间
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 节点名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 节点输出
	Outputs *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// 节点的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 65133
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// 资源组信息
	RuntimeResource *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// 工作流的脚本信息
	Script *ListNodeDependenciesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 调度策略
	Strategy *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// 标签信息
	Tags []*ListNodeDependenciesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 调度任务Id
	//
	// example:
	//
	// 580667964888595XXXX
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 触发器信息
	Trigger *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetInputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetProjectId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetScript(v *ListNodeDependenciesResponseBodyPagingInfoNodesScript) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTags(v []*ListNodeDependenciesResponseBodyPagingInfoNodesTags) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTaskId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) *ListNodeDependenciesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 543218872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 463497880880954XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 463497880880954XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource struct {
	// 资源组id
	//
	// example:
	//
	// S_res_group_XXXX_XXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScript struct {
	// 脚本的id
	//
	// example:
	//
	// 853573334108680XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) *ListNodeDependenciesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesStrategy struct {
	// 生成实例的模式
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// 重试时间间隔
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// 允许重跑的模式
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// 重试次数
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// 超时时间
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodeDependenciesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTags struct {
	// 标签键
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

type ListNodeDependenciesResponseBodyPagingInfoNodesTrigger struct {
	// 触发器的cron表达式
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// 结束时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 触发器id
	//
	// example:
	//
	// 543680677872062XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 开始时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// 触发器类型
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetId(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodeDependenciesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

type ListNodeDependenciesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeDependenciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeDependenciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeDependenciesResponse) GoString() string {
	return s.String()
}

func (s *ListNodeDependenciesResponse) SetHeaders(v map[string]*string) *ListNodeDependenciesResponse {
	s.Headers = v
	return s
}

func (s *ListNodeDependenciesResponse) SetStatusCode(v int32) *ListNodeDependenciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeDependenciesResponse) SetBody(v *ListNodeDependenciesResponseBody) *ListNodeDependenciesResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	// example:
	//
	// 860438872620113XXXX
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// example:
	//
	// Normal
	Rerurrence *string `json:"Rerurrence,omitempty" xml:"Rerurrence,omitempty"`
	// example:
	//
	// DATAWORKS_PROJECT
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetContainerId(v string) *ListNodesRequest {
	s.ContainerId = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetProjectId(v string) *ListNodesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListNodesRequest) SetRerunMode(v string) *ListNodesRequest {
	s.RerunMode = &v
	return s
}

func (s *ListNodesRequest) SetRerurrence(v string) *ListNodesRequest {
	s.Rerurrence = &v
	return s
}

func (s *ListNodesRequest) SetScene(v string) *ListNodesRequest {
	s.Scene = &v
	return s
}

type ListNodesResponseBody struct {
	PagingInfo *ListNodesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2197B9C4-39CE-55EA-8EEA-FDBAE52DXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetPagingInfo(v *ListNodesResponseBodyPagingInfo) *ListNodesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListNodesResponseBodyPagingInfo struct {
	Nodes []*ListNodesResponseBodyPagingInfoNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 42
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfo) SetNodes(v []*ListNodesResponseBodyPagingInfoNodes) *ListNodesResponseBodyPagingInfo {
	s.Nodes = v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageNumber(v string) *ListNodesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetPageSize(v string) *ListNodesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfo) SetTotalCount(v string) *ListNodesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodes struct {
	// 节点的创建时间
	//
	// example:
	//
	// 1722910655000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 数据源信息
	DataSource *ListNodesResponseBodyPagingInfoNodesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 860438872620113XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 节点输入
	Inputs *ListNodesResponseBodyPagingInfoNodesInputs `json:"Inputs,omitempty" xml:"Inputs,omitempty" type:"Struct"`
	// 属性修改时间
	//
	// example:
	//
	// 1722910655000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 节点名
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 节点输出
	Outputs *ListNodesResponseBodyPagingInfoNodesOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Struct"`
	// 节点的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 33233
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Normal
	Recurrence *string `json:"Recurrence,omitempty" xml:"Recurrence,omitempty"`
	// 资源组信息
	RuntimeResource *ListNodesResponseBodyPagingInfoNodesRuntimeResource `json:"RuntimeResource,omitempty" xml:"RuntimeResource,omitempty" type:"Struct"`
	// 工作流的脚本信息
	Script *ListNodesResponseBodyPagingInfoNodesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 调度策略
	Strategy *ListNodesResponseBodyPagingInfoNodesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
	// 标签信息
	Tags []*ListNodesResponseBodyPagingInfoNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// 调度任务Id
	//
	// example:
	//
	// 88888888888
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 触发器信息
	Trigger *ListNodesResponseBodyPagingInfoNodesTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetCreateTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDataSource(v *ListNodesResponseBodyPagingInfoNodesDataSource) *ListNodesResponseBodyPagingInfoNodes {
	s.DataSource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetDescription(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetInputs(v *ListNodesResponseBodyPagingInfoNodesInputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Inputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetModifyTime(v int64) *ListNodesResponseBodyPagingInfoNodes {
	s.ModifyTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetName(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOutputs(v *ListNodesResponseBodyPagingInfoNodesOutputs) *ListNodesResponseBodyPagingInfoNodes {
	s.Outputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetOwner(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Owner = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetProjectId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.ProjectId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRecurrence(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.Recurrence = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetRuntimeResource(v *ListNodesResponseBodyPagingInfoNodesRuntimeResource) *ListNodesResponseBodyPagingInfoNodes {
	s.RuntimeResource = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetScript(v *ListNodesResponseBodyPagingInfoNodesScript) *ListNodesResponseBodyPagingInfoNodes {
	s.Script = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetStrategy(v *ListNodesResponseBodyPagingInfoNodesStrategy) *ListNodesResponseBodyPagingInfoNodes {
	s.Strategy = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTags(v []*ListNodesResponseBodyPagingInfoNodesTags) *ListNodesResponseBodyPagingInfoNodes {
	s.Tags = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTaskId(v string) *ListNodesResponseBodyPagingInfoNodes {
	s.TaskId = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodes) SetTrigger(v *ListNodesResponseBodyPagingInfoNodesTrigger) *ListNodesResponseBodyPagingInfoNodes {
	s.Trigger = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesDataSource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetName(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesDataSource) SetType(v string) *ListNodesResponseBodyPagingInfoNodesDataSource {
	s.Type = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodesResponseBodyPagingInfoNodesInputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodesResponseBodyPagingInfoNodesInputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesInputsTables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesInputsVariables) *ListNodesResponseBodyPagingInfoNodesInputs {
	s.Variables = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 623731286945488XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesInputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesInputsTables {
	s.Guid = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 543211286945488XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// input
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 222
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetId(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariables {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesInputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 623731286945488XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesInputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputs struct {
	// 节点输出列表
	NodeOutputs []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs `json:"NodeOutputs,omitempty" xml:"NodeOutputs,omitempty" type:"Repeated"`
	// 表列表
	Tables []*ListNodesResponseBodyPagingInfoNodesOutputsTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// 变量列表
	Variables []*ListNodesResponseBodyPagingInfoNodesOutputsVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetNodeOutputs(v []*ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.NodeOutputs = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetTables(v []*ListNodesResponseBodyPagingInfoNodesOutputsTables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Tables = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputs) SetVariables(v []*ListNodesResponseBodyPagingInfoNodesOutputsVariables) *ListNodesResponseBodyPagingInfoNodesOutputs {
	s.Variables = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs) SetData(v string) *ListNodesResponseBodyPagingInfoNodesOutputsNodeOutputs {
	s.Data = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsTables struct {
	// 表id
	//
	// example:
	//
	// odps.autotest.test_output_table_1
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsTables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsTables) SetGuid(v string) *ListNodesResponseBodyPagingInfoNodesOutputsTables {
	s.Guid = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariables struct {
	// 制品类型
	//
	// example:
	//
	// Variable
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// 变量id
	//
	// example:
	//
	// 623731286945488XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 变量名
	//
	// example:
	//
	// output
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 变量所属节点
	Node *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Struct"`
	// 范围
	//
	// example:
	//
	// NodeParameter
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// 类型
	//
	// example:
	//
	// Constant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 变量值
	//
	// example:
	//
	// 111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariables) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetArtifactType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.ArtifactType = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetId(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetName(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetNode(v *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Node = v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetScope(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Scope = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetType(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariables) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariables {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode struct {
	// 节点输出
	//
	// example:
	//
	// 860438872620113XXXX
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode) SetOutput(v string) *ListNodesResponseBodyPagingInfoNodesOutputsVariablesNode {
	s.Output = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesRuntimeResource struct {
	// 资源组id
	//
	// example:
	//
	// S_resgrop_xxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesRuntimeResource) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesRuntimeResource) SetResourceGroupId(v string) *ListNodesResponseBodyPagingInfoNodesRuntimeResource {
	s.ResourceGroupId = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesScript struct {
	// 脚本的id
	//
	// example:
	//
	// 853573334108680XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 脚本路径
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListNodesResponseBodyPagingInfoNodesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListNodesResponseBodyPagingInfoNodesScript) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScript) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetId(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetPath(v string) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Path = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesScript) SetRuntime(v *ListNodesResponseBodyPagingInfoNodesScriptRuntime) *ListNodesResponseBodyPagingInfoNodesScript {
	s.Runtime = v
	return s
}

type ListNodesResponseBodyPagingInfoNodesScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_SQL
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesScriptRuntime) SetCommand(v string) *ListNodesResponseBodyPagingInfoNodesScriptRuntime {
	s.Command = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesStrategy struct {
	// 生成实例的模式
	//
	// example:
	//
	// T+1
	InstanceMode *string `json:"InstanceMode,omitempty" xml:"InstanceMode,omitempty"`
	// 重试时间间隔
	//
	// example:
	//
	// 180000
	RerunInterval *int32 `json:"RerunInterval,omitempty" xml:"RerunInterval,omitempty"`
	// 允许重跑的模式
	//
	// example:
	//
	// Allowed
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// 重试次数
	//
	// example:
	//
	// 3
	RerunTimes *int32 `json:"RerunTimes,omitempty" xml:"RerunTimes,omitempty"`
	// 超时时间
	//
	// example:
	//
	// 0
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesStrategy) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetInstanceMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.InstanceMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunInterval(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunInterval = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunMode(v string) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunMode = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetRerunTimes(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.RerunTimes = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesStrategy) SetTimeout(v int32) *ListNodesResponseBodyPagingInfoNodesStrategy {
	s.Timeout = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesTags struct {
	// 标签键
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTags) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetKey(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Key = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTags) SetValue(v string) *ListNodesResponseBodyPagingInfoNodesTags {
	s.Value = &v
	return s
}

type ListNodesResponseBodyPagingInfoNodesTrigger struct {
	// 触发器的cron表达式
	//
	// example:
	//
	// 00 00 00 	- 	- ?
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// 结束时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 9999-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 触发器id
	//
	// example:
	//
	// 543680677872062XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 开始时间，格式为yyyy-MM-dd HH:mm:ss
	//
	// example:
	//
	// 1970-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时区
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// 触发器类型
	//
	// example:
	//
	// Scheduler
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyPagingInfoNodesTrigger) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetCron(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Cron = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetEndTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.EndTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetId(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetStartTime(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.StartTime = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetTimezone(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Timezone = &v
	return s
}

func (s *ListNodesResponseBodyPagingInfoNodesTrigger) SetType(v string) *ListNodesResponseBodyPagingInfoNodesTrigger {
	s.Type = &v
	return s
}

type ListNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetStatusCode(v int32) *ListNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// python
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetOwner(v string) *ListResourcesRequest {
	s.Owner = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProjectId(v string) *ListResourcesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesRequest) SetType(v string) *ListResourcesRequest {
	s.Type = &v
	return s
}

type ListResourcesResponseBody struct {
	PagingInfo *ListResourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPagingInfo(v *ListResourcesResponseBodyPagingInfo) *ListResourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListResourcesResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Resources []*ListResourcesResponseBodyPagingInfoResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 131
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageNumber(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetPageSize(v int32) *ListResourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetResources(v []*ListResourcesResponseBodyPagingInfoResources) *ListResourcesResponseBodyPagingInfo {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfo) SetTotalCount(v int32) *ListResourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResources struct {
	// example:
	//
	// 1724505917000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 函数注册到的数据源信息
	DataSource *ListResourcesResponseBodyPagingInfoResourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 资源文件的最近修改时间
	//
	// example:
	//
	// 1724505917000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// math.py
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 资源文件的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 资源文件的项目id
	//
	// example:
	//
	// 344247
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 工作流的脚本信息
	Script *ListResourcesResponseBodyPagingInfoResourcesScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 文件目标存储路径
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// 文件资源来源存储类型
	//
	// example:
	//
	// local
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// 文件来源路径
	//
	// example:
	//
	// XXX/unknown/ide/1/XXX/20240820200851_963a9da676de44ef8d06a6576a8c4d6a.py
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// 文件目标存储类型
	//
	// example:
	//
	// oss
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// 资源类型
	//
	// example:
	//
	// jar
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResources) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetCreateTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetDataSource(v *ListResourcesResponseBodyPagingInfoResourcesDataSource) *ListResourcesResponseBodyPagingInfoResources {
	s.DataSource = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetId(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetModifyTime(v int64) *ListResourcesResponseBodyPagingInfoResources {
	s.ModifyTime = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetName(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetOwner(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetProjectId(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.ProjectId = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetScript(v *ListResourcesResponseBodyPagingInfoResourcesScript) *ListResourcesResponseBodyPagingInfoResources {
	s.Script = v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourcePath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourcePath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetSourceType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.SourceType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetPath(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetPath = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetTargetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.TargetType = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResources) SetType(v string) *ListResourcesResponseBodyPagingInfoResources {
	s.Type = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesDataSource struct {
	// 数据源名称
	//
	// example:
	//
	// odps_first
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 数据源类型
	//
	// example:
	//
	// odps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetName(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesDataSource) SetType(v string) *ListResourcesResponseBodyPagingInfoResourcesDataSource {
	s.Type = &v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesScript struct {
	// 工作流脚本的id
	//
	// example:
	//
	// 123348864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 工作流的脚本路径
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScript) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetId(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Id = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetPath(v string) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Path = &v
	return s
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScript) SetRuntime(v *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) *ListResourcesResponseBodyPagingInfoResourcesScript {
	s.Runtime = v
	return s
}

type ListResourcesResponseBodyPagingInfoResourcesScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// ODPS_PYTHON
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime) SetCommand(v string) *ListResourcesResponseBodyPagingInfoResourcesScriptRuntime {
	s.Command = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListWorkflowDefinitionsRequest struct {
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsRequest) SetOwner(v string) *ListWorkflowDefinitionsRequest {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageNumber(v int32) *ListWorkflowDefinitionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetPageSize(v int32) *ListWorkflowDefinitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetProjectId(v string) *ListWorkflowDefinitionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsRequest) SetType(v string) *ListWorkflowDefinitionsRequest {
	s.Type = &v
	return s
}

type ListWorkflowDefinitionsResponseBody struct {
	PagingInfo *ListWorkflowDefinitionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 8C3ED0C5-ABAB-55E1-854B-DAC02B11XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBody) SetPagingInfo(v *ListWorkflowDefinitionsResponseBodyPagingInfo) *ListWorkflowDefinitionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBody) SetRequestId(v string) *ListWorkflowDefinitionsResponseBody {
	s.RequestId = &v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 227
	TotalCount          *int32                                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkflowDefinitions []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions `json:"WorkflowDefinitions,omitempty" xml:"WorkflowDefinitions,omitempty" type:"Repeated"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetPageSize(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfo) SetWorkflowDefinitions(v []*ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) *ListWorkflowDefinitionsResponseBodyPagingInfo {
	s.WorkflowDefinitions = v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions struct {
	// 工作流的创建时间
	//
	// example:
	//
	// 1698057323000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 工作流的描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 工作流定义的唯一ID
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 工作流的最近修改时间
	//
	// example:
	//
	// 1698057323000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// 工作流的名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 工作流的责任人
	//
	// example:
	//
	// 110755000425XXXX
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// 工作流定义的所属项目空间
	//
	// This parameter is required.
	//
	// example:
	//
	// 4710
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 工作流的脚本信息
	Script *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript `json:"Script,omitempty" xml:"Script,omitempty" type:"Struct"`
	// 工作流类型，可选值：CycleWorkflow、ManualWorkflow，分别表示周期工作流和手动工作流
	//
	// example:
	//
	// CycleWorkflow
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetCreateTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.CreateTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetDescription(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Description = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetModifyTime(v int64) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ModifyTime = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetName(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Name = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetOwner(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Owner = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetProjectId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.ProjectId = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetScript(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Script = v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions) SetType(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitions {
	s.Type = &v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript struct {
	// 工作流脚本的id
	//
	// example:
	//
	// 698002781368644XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 工作流的脚本路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 脚本的运行时信息
	Runtime *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime `json:"Runtime,omitempty" xml:"Runtime,omitempty" type:"Struct"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetId(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Id = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetPath(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Path = &v
	return s
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript) SetRuntime(v *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScript {
	s.Runtime = v
	return s
}

type ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime struct {
	// 脚本所属类型
	//
	// example:
	//
	// WORKFLOW
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime) SetCommand(v string) *ListWorkflowDefinitionsResponseBodyPagingInfoWorkflowDefinitionsScriptRuntime {
	s.Command = &v
	return s
}

type ListWorkflowDefinitionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkflowDefinitionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkflowDefinitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkflowDefinitionsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkflowDefinitionsResponse) SetHeaders(v map[string]*string) *ListWorkflowDefinitionsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetStatusCode(v int32) *ListWorkflowDefinitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkflowDefinitionsResponse) SetBody(v *ListWorkflowDefinitionsResponseBody) *ListWorkflowDefinitionsResponse {
	s.Body = v
	return s
}

type MoveFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionRequest) GoString() string {
	return s.String()
}

func (s *MoveFunctionRequest) SetId(v string) *MoveFunctionRequest {
	s.Id = &v
	return s
}

func (s *MoveFunctionRequest) SetPath(v string) *MoveFunctionRequest {
	s.Path = &v
	return s
}

func (s *MoveFunctionRequest) SetProjectId(v string) *MoveFunctionRequest {
	s.ProjectId = &v
	return s
}

type MoveFunctionResponseBody struct {
	// example:
	//
	// 48C0E2F0-52BA-5888-BDFA-28F1B9AFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponseBody) SetRequestId(v string) *MoveFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveFunctionResponseBody) SetSuccess(v bool) *MoveFunctionResponseBody {
	s.Success = &v
	return s
}

type MoveFunctionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFunctionResponse) GoString() string {
	return s.String()
}

func (s *MoveFunctionResponse) SetHeaders(v map[string]*string) *MoveFunctionResponse {
	s.Headers = v
	return s
}

func (s *MoveFunctionResponse) SetStatusCode(v int32) *MoveFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveFunctionResponse) SetBody(v *MoveFunctionResponseBody) *MoveFunctionResponse {
	s.Body = v
	return s
}

type MoveNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeRequest) GoString() string {
	return s.String()
}

func (s *MoveNodeRequest) SetId(v string) *MoveNodeRequest {
	s.Id = &v
	return s
}

func (s *MoveNodeRequest) SetPath(v string) *MoveNodeRequest {
	s.Path = &v
	return s
}

func (s *MoveNodeRequest) SetProjectId(v string) *MoveNodeRequest {
	s.ProjectId = &v
	return s
}

type MoveNodeResponseBody struct {
	// example:
	//
	// C99E2BE6-9DEA-5C2E-8F51-1DDCFEADXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeResponseBody) GoString() string {
	return s.String()
}

func (s *MoveNodeResponseBody) SetRequestId(v string) *MoveNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveNodeResponseBody) SetSuccess(v bool) *MoveNodeResponseBody {
	s.Success = &v
	return s
}

type MoveNodeResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveNodeResponse) GoString() string {
	return s.String()
}

func (s *MoveNodeResponse) SetHeaders(v map[string]*string) *MoveNodeResponse {
	s.Headers = v
	return s
}

func (s *MoveNodeResponse) SetStatusCode(v int32) *MoveNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveNodeResponse) SetBody(v *MoveNodeResponseBody) *MoveNodeResponse {
	s.Body = v
	return s
}

type MoveResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceRequest) SetId(v string) *MoveResourceRequest {
	s.Id = &v
	return s
}

func (s *MoveResourceRequest) SetPath(v string) *MoveResourceRequest {
	s.Path = &v
	return s
}

func (s *MoveResourceRequest) SetProjectId(v string) *MoveResourceRequest {
	s.ProjectId = &v
	return s
}

type MoveResourceResponseBody struct {
	// example:
	//
	// F332BED4-DD73-5972-A9C2-642BA6CFXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceResponseBody) SetRequestId(v string) *MoveResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveResourceResponseBody) SetSuccess(v bool) *MoveResourceResponseBody {
	s.Success = &v
	return s
}

type MoveResourceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceResponse) SetHeaders(v map[string]*string) *MoveResourceResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceResponse) SetStatusCode(v int32) *MoveResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceResponse) SetBody(v *MoveResourceResponseBody) *MoveResourceResponse {
	s.Body = v
	return s
}

type MoveWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/demo
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s MoveWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionRequest) SetId(v string) *MoveWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetPath(v string) *MoveWorkflowDefinitionRequest {
	s.Path = &v
	return s
}

func (s *MoveWorkflowDefinitionRequest) SetProjectId(v string) *MoveWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type MoveWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 05ADAF4F-7709-5FB1-B606-3513483FXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MoveWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponseBody) SetRequestId(v string) *MoveWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveWorkflowDefinitionResponseBody) SetSuccess(v bool) *MoveWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type MoveWorkflowDefinitionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *MoveWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *MoveWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetStatusCode(v int32) *MoveWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveWorkflowDefinitionResponse) SetBody(v *MoveWorkflowDefinitionResponseBody) *MoveWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type RenameFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10002
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionRequest) GoString() string {
	return s.String()
}

func (s *RenameFunctionRequest) SetId(v string) *RenameFunctionRequest {
	s.Id = &v
	return s
}

func (s *RenameFunctionRequest) SetName(v string) *RenameFunctionRequest {
	s.Name = &v
	return s
}

func (s *RenameFunctionRequest) SetProjectId(v string) *RenameFunctionRequest {
	s.ProjectId = &v
	return s
}

type RenameFunctionResponseBody struct {
	// example:
	//
	// 1ED4C97F-BA2A-57C5-BA7C-8853627EXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponseBody) SetRequestId(v string) *RenameFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameFunctionResponseBody) SetSuccess(v string) *RenameFunctionResponseBody {
	s.Success = &v
	return s
}

type RenameFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameFunctionResponse) GoString() string {
	return s.String()
}

func (s *RenameFunctionResponse) SetHeaders(v map[string]*string) *RenameFunctionResponse {
	s.Headers = v
	return s
}

func (s *RenameFunctionResponse) SetStatusCode(v int32) *RenameFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameFunctionResponse) SetBody(v *RenameFunctionResponseBody) *RenameFunctionResponse {
	s.Body = v
	return s
}

type RenameNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeRequest) GoString() string {
	return s.String()
}

func (s *RenameNodeRequest) SetId(v string) *RenameNodeRequest {
	s.Id = &v
	return s
}

func (s *RenameNodeRequest) SetName(v string) *RenameNodeRequest {
	s.Name = &v
	return s
}

func (s *RenameNodeRequest) SetProjectId(v string) *RenameNodeRequest {
	s.ProjectId = &v
	return s
}

type RenameNodeResponseBody struct {
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeResponseBody) GoString() string {
	return s.String()
}

func (s *RenameNodeResponseBody) SetRequestId(v string) *RenameNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameNodeResponseBody) SetSuccess(v bool) *RenameNodeResponseBody {
	s.Success = &v
	return s
}

type RenameNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameNodeResponse) GoString() string {
	return s.String()
}

func (s *RenameNodeResponse) SetHeaders(v map[string]*string) *RenameNodeResponse {
	s.Headers = v
	return s
}

func (s *RenameNodeResponse) SetStatusCode(v int32) *RenameNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameNodeResponse) SetBody(v *RenameNodeResponseBody) *RenameNodeResponse {
	s.Body = v
	return s
}

type RenameResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceRequest) GoString() string {
	return s.String()
}

func (s *RenameResourceRequest) SetId(v string) *RenameResourceRequest {
	s.Id = &v
	return s
}

func (s *RenameResourceRequest) SetName(v string) *RenameResourceRequest {
	s.Name = &v
	return s
}

func (s *RenameResourceRequest) SetProjectId(v string) *RenameResourceRequest {
	s.ProjectId = &v
	return s
}

type RenameResourceResponseBody struct {
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA8XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameResourceResponseBody) SetRequestId(v string) *RenameResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameResourceResponseBody) SetSuccess(v bool) *RenameResourceResponseBody {
	s.Success = &v
	return s
}

type RenameResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameResourceResponse) GoString() string {
	return s.String()
}

func (s *RenameResourceResponse) SetHeaders(v map[string]*string) *RenameResourceResponse {
	s.Headers = v
	return s
}

func (s *RenameResourceResponse) SetStatusCode(v int32) *RenameResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameResourceResponse) SetBody(v *RenameResourceResponseBody) *RenameResourceResponse {
	s.Body = v
	return s
}

type RenameWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s RenameWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionRequest) SetId(v string) *RenameWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetName(v string) *RenameWorkflowDefinitionRequest {
	s.Name = &v
	return s
}

func (s *RenameWorkflowDefinitionRequest) SetProjectId(v string) *RenameWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

type RenameWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 975BD43D-C421-595C-A29C-565A8AD5XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenameWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponseBody) SetRequestId(v string) *RenameWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameWorkflowDefinitionResponseBody) SetSuccess(v bool) *RenameWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type RenameWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenameWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenameWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *RenameWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *RenameWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetStatusCode(v int32) *RenameWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameWorkflowDefinitionResponse) SetBody(v *RenameWorkflowDefinitionResponseBody) *RenameWorkflowDefinitionResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 463497880880954XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetId(v string) *UpdateFunctionRequest {
	s.Id = &v
	return s
}

func (s *UpdateFunctionRequest) SetProjectId(v string) *UpdateFunctionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFunctionRequest) SetSpec(v string) *UpdateFunctionRequest {
	s.Spec = &v
	return s
}

type UpdateFunctionResponseBody struct {
	// example:
	//
	// 12123960-CB2C-5086-868E-C6C1D024XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSuccess(v bool) *UpdateFunctionResponseBody {
	s.Success = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetStatusCode(v int32) *UpdateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdateNodeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeRequest) SetId(v string) *UpdateNodeRequest {
	s.Id = &v
	return s
}

func (s *UpdateNodeRequest) SetProjectId(v string) *UpdateNodeRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateNodeRequest) SetSpec(v string) *UpdateNodeRequest {
	s.Spec = &v
	return s
}

type UpdateNodeResponseBody struct {
	// example:
	//
	// 99EBE7CF-69C0-5089-BE3E-79563C31XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponseBody) SetRequestId(v string) *UpdateNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeResponseBody) SetSuccess(v bool) *UpdateNodeResponseBody {
	s.Success = &v
	return s
}

type UpdateNodeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeResponse) SetHeaders(v map[string]*string) *UpdateNodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeResponse) SetStatusCode(v int32) *UpdateNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeResponse) SetBody(v *UpdateNodeResponseBody) *UpdateNodeResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 543217824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetId(v string) *UpdateResourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateResourceRequest) SetProjectId(v string) *UpdateResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateResourceRequest) SetSpec(v string) *UpdateResourceRequest {
	s.Spec = &v
	return s
}

type UpdateResourceResponseBody struct {
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA81XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetSuccess(v bool) *UpdateResourceResponseBody {
	s.Success = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
	s.Body = v
	return s
}

type UpdateWorkflowDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 652567824470354XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s UpdateWorkflowDefinitionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionRequest) SetId(v string) *UpdateWorkflowDefinitionRequest {
	s.Id = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetProjectId(v string) *UpdateWorkflowDefinitionRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateWorkflowDefinitionRequest) SetSpec(v string) *UpdateWorkflowDefinitionRequest {
	s.Spec = &v
	return s
}

type UpdateWorkflowDefinitionResponseBody struct {
	// example:
	//
	// 20BF7E80-668A-5620-8AD8-879B8FEAXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWorkflowDefinitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponseBody) SetRequestId(v string) *UpdateWorkflowDefinitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponseBody) SetSuccess(v bool) *UpdateWorkflowDefinitionResponseBody {
	s.Success = &v
	return s
}

type UpdateWorkflowDefinitionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWorkflowDefinitionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkflowDefinitionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWorkflowDefinitionResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowDefinitionResponse) SetHeaders(v map[string]*string) *UpdateWorkflowDefinitionResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetStatusCode(v int32) *UpdateWorkflowDefinitionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkflowDefinitionResponse) SetBody(v *UpdateWorkflowDefinitionResponseBody) *UpdateWorkflowDefinitionResponse {
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
		"ap-northeast-1":        tea.String("dataworks.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("dataworks.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        tea.String("dataworks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("dataworks.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("dataworks.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("dataworks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            tea.String("dataworks.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            tea.String("dataworks.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou":           tea.String("dataworks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           tea.String("dataworks.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          tea.String("dataworks.aliyuncs.com"),
		"cn-qingdao":            tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai":           tea.String("dataworks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("dataworks.cn-shenzhen.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("dataworks.aliyuncs.com"),
		"eu-central-1":          tea.String("dataworks.eu-central-1.aliyuncs.com"),
		"eu-west-1":             tea.String("dataworks.eu-west-1.aliyuncs.com"),
		"me-east-1":             tea.String("dataworks.me-east-1.aliyuncs.com"),
		"us-east-1":             tea.String("dataworks.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("dataworks.us-west-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("dataworks.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("dataworks.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("dataworks.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataworks-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 终止发布流程
//
// @param request - AbolishDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishDeploymentResponse
func (client *Client) AbolishDeploymentWithOptions(request *AbolishDeploymentRequest, runtime *util.RuntimeOptions) (_result *AbolishDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AbolishDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbolishDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 终止发布流程
//
// @param request - AbolishDeploymentRequest
//
// @return AbolishDeploymentResponse
func (client *Client) AbolishDeployment(request *AbolishDeploymentRequest) (_result *AbolishDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbolishDeploymentResponse{}
	_body, _err := client.AbolishDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发布流程
//
// @param tmpReq - CreateDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeploymentWithOptions(tmpReq *CreateDeploymentRequest, runtime *util.RuntimeOptions) (_result *CreateDeploymentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDeploymentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ObjectIds)) {
		request.ObjectIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ObjectIds, tea.String("ObjectIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectIdsShrink)) {
		body["ObjectIds"] = request.ObjectIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建发布流程
//
// @param request - CreateDeploymentRequest
//
// @return CreateDeploymentResponse
func (client *Client) CreateDeployment(request *CreateDeploymentRequest) (_result *CreateDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDeploymentResponse{}
	_body, _err := client.CreateDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建udf函数
//
// @param request - CreateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResponse
func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建udf函数
//
// @param request - CreateFunctionRequest
//
// @return CreateFunctionResponse
func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建节点
//
// @param request - CreateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeResponse
func (client *Client) CreateNodeWithOptions(request *CreateNodeRequest, runtime *util.RuntimeOptions) (_result *CreateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContainerId)) {
		body["ContainerId"] = request.ContainerId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Scene)) {
		body["Scene"] = request.Scene
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建节点
//
// @param request - CreateNodeRequest
//
// @return CreateNodeResponse
func (client *Client) CreateNode(request *CreateNodeRequest) (_result *CreateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeResponse{}
	_body, _err := client.CreateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源文件
//
// @param request - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源文件
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinitionWithOptions(request *CreateWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *CreateWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建工作流
//
// @param request - CreateWorkflowDefinitionRequest
//
// @return CreateWorkflowDefinitionResponse
func (client *Client) CreateWorkflowDefinition(request *CreateWorkflowDefinitionRequest) (_result *CreateWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWorkflowDefinitionResponse{}
	_body, _err := client.CreateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除udf函数
//
// @param request - DeleteFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除udf函数
//
// @param request - DeleteFunctionRequest
//
// @return DeleteFunctionResponse
func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除节点
//
// @param request - DeleteNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeResponse
func (client *Client) DeleteNodeWithOptions(request *DeleteNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除节点
//
// @param request - DeleteNodeRequest
//
// @return DeleteNodeResponse
func (client *Client) DeleteNode(request *DeleteNodeRequest) (_result *DeleteNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeResponse{}
	_body, _err := client.DeleteNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除资源文件
//
// @param request - DeleteResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceResponse
func (client *Client) DeleteResourceWithOptions(request *DeleteResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除资源文件
//
// @param request - DeleteResourceRequest
//
// @return DeleteResourceResponse
func (client *Client) DeleteResource(request *DeleteResourceRequest) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinitionWithOptions(request *DeleteWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除工作流
//
// @param request - DeleteWorkflowDefinitionRequest
//
// @return DeleteWorkflowDefinitionResponse
func (client *Client) DeleteWorkflowDefinition(request *DeleteWorkflowDefinitionRequest) (_result *DeleteWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteWorkflowDefinitionResponse{}
	_body, _err := client.DeleteWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 执行Deployment一个阶段
//
// @param request - ExecDeploymentStageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecDeploymentStageResponse
func (client *Client) ExecDeploymentStageWithOptions(request *ExecDeploymentStageRequest, runtime *util.RuntimeOptions) (_result *ExecDeploymentStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["Code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecDeploymentStage"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecDeploymentStageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行Deployment一个阶段
//
// @param request - ExecDeploymentStageRequest
//
// @return ExecDeploymentStageResponse
func (client *Client) ExecDeploymentStage(request *ExecDeploymentStageRequest) (_result *ExecDeploymentStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecDeploymentStageResponse{}
	_body, _err := client.ExecDeploymentStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发布流程详情
//
// @param request - GetDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithOptions(request *GetDeploymentRequest, runtime *util.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployment"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发布流程详情
//
// @param request - GetDeploymentRequest
//
// @return GetDeploymentResponse
func (client *Client) GetDeployment(request *GetDeploymentRequest) (_result *GetDeploymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeploymentResponse{}
	_body, _err := client.GetDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取函数信息
//
// @param request - GetFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResponse
func (client *Client) GetFunctionWithOptions(request *GetFunctionRequest, runtime *util.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取函数信息
//
// @param request - GetFunctionRequest
//
// @return GetFunctionResponse
func (client *Client) GetFunction(request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithOptions(request *GetNodeRequest, runtime *util.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetNodeRequest
//
// @return GetNodeResponse
func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取资源文件
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取资源文件
//
// @param request - GetResourceRequest
//
// @return GetResourceResponse
func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - GetWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinitionWithOptions(request *GetWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *GetWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取工作流详情
//
// @param request - GetWorkflowDefinitionRequest
//
// @return GetWorkflowDefinitionResponse
func (client *Client) GetWorkflowDefinition(request *GetWorkflowDefinitionRequest) (_result *GetWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkflowDefinitionResponse{}
	_body, _err := client.GetWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取发布流程
//
// @param request - ListDeploymentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithOptions(request *ListDeploymentsRequest, runtime *util.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDeployments"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取发布流程
//
// @param request - ListDeploymentsRequest
//
// @return ListDeploymentsResponse
func (client *Client) ListDeployments(request *ListDeploymentsRequest) (_result *ListDeploymentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.ListDeploymentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取udf函数列表
//
// @param request - ListFunctionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionsResponse
func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取udf函数列表
//
// @param request - ListFunctionsRequest
//
// @return ListFunctionsResponse
func (client *Client) ListFunctions(request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点依赖列表
//
// @param request - ListNodeDependenciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependenciesWithOptions(request *ListNodeDependenciesRequest, runtime *util.RuntimeOptions) (_result *ListNodeDependenciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeDependencies"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点依赖列表
//
// @param request - ListNodeDependenciesRequest
//
// @return ListNodeDependenciesResponse
func (client *Client) ListNodeDependencies(request *ListNodeDependenciesRequest) (_result *ListNodeDependenciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeDependenciesResponse{}
	_body, _err := client.ListNodeDependenciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取节点列表
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodes"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取节点列表
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页获取资源文件
//
// @param request - ListResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourcesResponse
func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页获取资源文件
//
// @param request - ListResourcesRequest
//
// @return ListResourcesResponse
func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取workflowDefinition的列表
//
// @param request - ListWorkflowDefinitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitionsWithOptions(request *ListWorkflowDefinitionsRequest, runtime *util.RuntimeOptions) (_result *ListWorkflowDefinitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkflowDefinitions"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取workflowDefinition的列表
//
// @param request - ListWorkflowDefinitionsRequest
//
// @return ListWorkflowDefinitionsResponse
func (client *Client) ListWorkflowDefinitions(request *ListWorkflowDefinitionsRequest) (_result *ListWorkflowDefinitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkflowDefinitionsResponse{}
	_body, _err := client.ListWorkflowDefinitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动funciton的路径
//
// @param request - MoveFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFunctionResponse
func (client *Client) MoveFunctionWithOptions(request *MoveFunctionRequest, runtime *util.RuntimeOptions) (_result *MoveFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动funciton的路径
//
// @param request - MoveFunctionRequest
//
// @return MoveFunctionResponse
func (client *Client) MoveFunction(request *MoveFunctionRequest) (_result *MoveFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveFunctionResponse{}
	_body, _err := client.MoveFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动节点路径
//
// @param request - MoveNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveNodeResponse
func (client *Client) MoveNodeWithOptions(request *MoveNodeRequest, runtime *util.RuntimeOptions) (_result *MoveNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动节点路径
//
// @param request - MoveNodeRequest
//
// @return MoveNodeResponse
func (client *Client) MoveNode(request *MoveNodeRequest) (_result *MoveNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveNodeResponse{}
	_body, _err := client.MoveNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动资源文件路径
//
// @param request - MoveResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceResponse
func (client *Client) MoveResourceWithOptions(request *MoveResourceRequest, runtime *util.RuntimeOptions) (_result *MoveResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动资源文件路径
//
// @param request - MoveResourceRequest
//
// @return MoveResourceResponse
func (client *Client) MoveResource(request *MoveResourceRequest) (_result *MoveResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceResponse{}
	_body, _err := client.MoveResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动工作流的路径
//
// @param request - MoveWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinitionWithOptions(request *MoveWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *MoveWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动工作流的路径
//
// @param request - MoveWorkflowDefinitionRequest
//
// @return MoveWorkflowDefinitionResponse
func (client *Client) MoveWorkflowDefinition(request *MoveWorkflowDefinitionRequest) (_result *MoveWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveWorkflowDefinitionResponse{}
	_body, _err := client.MoveWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对function重命名
//
// @param request - RenameFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameFunctionResponse
func (client *Client) RenameFunctionWithOptions(request *RenameFunctionRequest, runtime *util.RuntimeOptions) (_result *RenameFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对function重命名
//
// @param request - RenameFunctionRequest
//
// @return RenameFunctionResponse
func (client *Client) RenameFunction(request *RenameFunctionRequest) (_result *RenameFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameFunctionResponse{}
	_body, _err := client.RenameFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名节点
//
// @param request - RenameNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameNodeResponse
func (client *Client) RenameNodeWithOptions(request *RenameNodeRequest, runtime *util.RuntimeOptions) (_result *RenameNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名节点
//
// @param request - RenameNodeRequest
//
// @return RenameNodeResponse
func (client *Client) RenameNode(request *RenameNodeRequest) (_result *RenameNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameNodeResponse{}
	_body, _err := client.RenameNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对资源文件重命名
//
// @param request - RenameResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameResourceResponse
func (client *Client) RenameResourceWithOptions(request *RenameResourceRequest, runtime *util.RuntimeOptions) (_result *RenameResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对资源文件重命名
//
// @param request - RenameResourceRequest
//
// @return RenameResourceResponse
func (client *Client) RenameResource(request *RenameResourceRequest) (_result *RenameResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameResourceResponse{}
	_body, _err := client.RenameResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重命名工作流
//
// @param request - RenameWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinitionWithOptions(request *RenameWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *RenameWorkflowDefinitionResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重命名工作流
//
// @param request - RenameWorkflowDefinitionRequest
//
// @return RenameWorkflowDefinitionResponse
func (client *Client) RenameWorkflowDefinition(request *RenameWorkflowDefinitionRequest) (_result *RenameWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameWorkflowDefinitionResponse{}
	_body, _err := client.RenameWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新udf函数
//
// @param request - UpdateFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新udf函数
//
// @param request - UpdateFunctionRequest
//
// @return UpdateFunctionResponse
func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新节点
//
// @param request - UpdateNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeResponse
func (client *Client) UpdateNodeWithOptions(request *UpdateNodeRequest, runtime *util.RuntimeOptions) (_result *UpdateNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNode"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新节点
//
// @param request - UpdateNodeRequest
//
// @return UpdateNodeResponse
func (client *Client) UpdateNode(request *UpdateNodeRequest) (_result *UpdateNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodeResponse{}
	_body, _err := client.UpdateNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新资源文件
//
// @param request - UpdateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourceResponse
func (client *Client) UpdateResourceWithOptions(request *UpdateResourceRequest, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新资源文件
//
// @param request - UpdateResourceRequest
//
// @return UpdateResourceResponse
func (client *Client) UpdateResource(request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinitionWithOptions(request *UpdateWorkflowDefinitionRequest, runtime *util.RuntimeOptions) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		body["Spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateWorkflowDefinition"),
		Version:     tea.String("2024-05-18"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新工作流
//
// @param request - UpdateWorkflowDefinitionRequest
//
// @return UpdateWorkflowDefinitionResponse
func (client *Client) UpdateWorkflowDefinition(request *UpdateWorkflowDefinitionRequest) (_result *UpdateWorkflowDefinitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateWorkflowDefinitionResponse{}
	_body, _err := client.UpdateWorkflowDefinitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
