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

type ErrorResponse struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s ErrorResponse) GoString() string {
	return s.String()
}

func (s *ErrorResponse) SetCode(v string) *ErrorResponse {
	s.Code = &v
	return s
}

func (s *ErrorResponse) SetMessage(v string) *ErrorResponse {
	s.Message = &v
	return s
}

func (s *ErrorResponse) SetRequestId(v string) *ErrorResponse {
	s.RequestId = &v
	return s
}

type CreateGraphRequest struct {
	Body      *string `json:"body,omitempty" xml:"body,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s CreateGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGraphRequest) GoString() string {
	return s.String()
}

func (s *CreateGraphRequest) SetBody(v string) *CreateGraphRequest {
	s.Body = &v
	return s
}

func (s *CreateGraphRequest) SetNamespace(v string) *CreateGraphRequest {
	s.Namespace = &v
	return s
}

type CreateGraphResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGraphResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGraphResponseBody) SetCode(v string) *CreateGraphResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGraphResponseBody) SetMessage(v string) *CreateGraphResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGraphResponseBody) SetRequestId(v string) *CreateGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGraphResponseBody) SetResult(v map[string]interface{}) *CreateGraphResponseBody {
	s.Result = v
	return s
}

type CreateGraphResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGraphResponse) GoString() string {
	return s.String()
}

func (s *CreateGraphResponse) SetHeaders(v map[string]*string) *CreateGraphResponse {
	s.Headers = v
	return s
}

func (s *CreateGraphResponse) SetStatusCode(v int32) *CreateGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGraphResponse) SetBody(v *CreateGraphResponseBody) *CreateGraphResponse {
	s.Body = v
	return s
}

type CreateGraphSchemaRequest struct {
	Body      *string `json:"body,omitempty" xml:"body,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s CreateGraphSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGraphSchemaRequest) GoString() string {
	return s.String()
}

func (s *CreateGraphSchemaRequest) SetBody(v string) *CreateGraphSchemaRequest {
	s.Body = &v
	return s
}

func (s *CreateGraphSchemaRequest) SetNamespace(v string) *CreateGraphSchemaRequest {
	s.Namespace = &v
	return s
}

type CreateGraphSchemaResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateGraphSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGraphSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGraphSchemaResponseBody) SetCode(v string) *CreateGraphSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGraphSchemaResponseBody) SetMessage(v string) *CreateGraphSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGraphSchemaResponseBody) SetRequestId(v string) *CreateGraphSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGraphSchemaResponseBody) SetResult(v map[string]interface{}) *CreateGraphSchemaResponseBody {
	s.Result = v
	return s
}

type CreateGraphSchemaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateGraphSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGraphSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGraphSchemaResponse) GoString() string {
	return s.String()
}

func (s *CreateGraphSchemaResponse) SetHeaders(v map[string]*string) *CreateGraphSchemaResponse {
	s.Headers = v
	return s
}

func (s *CreateGraphSchemaResponse) SetStatusCode(v int32) *CreateGraphSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGraphSchemaResponse) SetBody(v *CreateGraphSchemaResponseBody) *CreateGraphSchemaResponse {
	s.Body = v
	return s
}

type DeleteGraphRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s DeleteGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGraphRequest) GoString() string {
	return s.String()
}

func (s *DeleteGraphRequest) SetNamespace(v string) *DeleteGraphRequest {
	s.Namespace = &v
	return s
}

type DeleteGraphResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGraphResponseBody) SetCode(v string) *DeleteGraphResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGraphResponseBody) SetMessage(v string) *DeleteGraphResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGraphResponseBody) SetRequestId(v string) *DeleteGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGraphResponseBody) SetResult(v map[string]interface{}) *DeleteGraphResponseBody {
	s.Result = v
	return s
}

type DeleteGraphResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGraphResponse) GoString() string {
	return s.String()
}

func (s *DeleteGraphResponse) SetHeaders(v map[string]*string) *DeleteGraphResponse {
	s.Headers = v
	return s
}

func (s *DeleteGraphResponse) SetStatusCode(v int32) *DeleteGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGraphResponse) SetBody(v *DeleteGraphResponseBody) *DeleteGraphResponse {
	s.Body = v
	return s
}

type GetGraphRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s GetGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGraphRequest) GoString() string {
	return s.String()
}

func (s *GetGraphRequest) SetNamespace(v string) *GetGraphRequest {
	s.Namespace = &v
	return s
}

type GetGraphResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGraphResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphResponseBody) SetCode(v string) *GetGraphResponseBody {
	s.Code = &v
	return s
}

func (s *GetGraphResponseBody) SetMessage(v string) *GetGraphResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphResponseBody) SetRequestId(v string) *GetGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphResponseBody) SetResult(v map[string]interface{}) *GetGraphResponseBody {
	s.Result = v
	return s
}

type GetGraphResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGraphResponse) GoString() string {
	return s.String()
}

func (s *GetGraphResponse) SetHeaders(v map[string]*string) *GetGraphResponse {
	s.Headers = v
	return s
}

func (s *GetGraphResponse) SetStatusCode(v int32) *GetGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGraphResponse) SetBody(v *GetGraphResponseBody) *GetGraphResponse {
	s.Body = v
	return s
}

type GetGraphSchemaRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s GetGraphSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGraphSchemaRequest) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaRequest) SetNamespace(v string) *GetGraphSchemaRequest {
	s.Namespace = &v
	return s
}

type GetGraphSchemaResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetGraphSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGraphSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaResponseBody) SetCode(v string) *GetGraphSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetMessage(v string) *GetGraphSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetRequestId(v string) *GetGraphSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGraphSchemaResponseBody) SetResult(v map[string]interface{}) *GetGraphSchemaResponseBody {
	s.Result = v
	return s
}

type GetGraphSchemaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetGraphSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGraphSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGraphSchemaResponse) GoString() string {
	return s.String()
}

func (s *GetGraphSchemaResponse) SetHeaders(v map[string]*string) *GetGraphSchemaResponse {
	s.Headers = v
	return s
}

func (s *GetGraphSchemaResponse) SetStatusCode(v int32) *GetGraphSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGraphSchemaResponse) SetBody(v *GetGraphSchemaResponseBody) *GetGraphSchemaResponse {
	s.Body = v
	return s
}

type GetIgraphLabelBackFlowResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetIgraphLabelBackFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphLabelBackFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetIgraphLabelBackFlowResponseBody) SetCode(v string) *GetIgraphLabelBackFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetIgraphLabelBackFlowResponseBody) SetMessage(v string) *GetIgraphLabelBackFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetIgraphLabelBackFlowResponseBody) SetRequestId(v string) *GetIgraphLabelBackFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIgraphLabelBackFlowResponseBody) SetResult(v map[string]interface{}) *GetIgraphLabelBackFlowResponseBody {
	s.Result = v
	return s
}

type GetIgraphLabelBackFlowResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIgraphLabelBackFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIgraphLabelBackFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphLabelBackFlowResponse) GoString() string {
	return s.String()
}

func (s *GetIgraphLabelBackFlowResponse) SetHeaders(v map[string]*string) *GetIgraphLabelBackFlowResponse {
	s.Headers = v
	return s
}

func (s *GetIgraphLabelBackFlowResponse) SetStatusCode(v int32) *GetIgraphLabelBackFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIgraphLabelBackFlowResponse) SetBody(v *GetIgraphLabelBackFlowResponseBody) *GetIgraphLabelBackFlowResponse {
	s.Body = v
	return s
}

type GetIgraphLabelLastBackflowResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetIgraphLabelLastBackflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphLabelLastBackflowResponseBody) GoString() string {
	return s.String()
}

func (s *GetIgraphLabelLastBackflowResponseBody) SetCode(v string) *GetIgraphLabelLastBackflowResponseBody {
	s.Code = &v
	return s
}

func (s *GetIgraphLabelLastBackflowResponseBody) SetMessage(v string) *GetIgraphLabelLastBackflowResponseBody {
	s.Message = &v
	return s
}

func (s *GetIgraphLabelLastBackflowResponseBody) SetRequestId(v string) *GetIgraphLabelLastBackflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIgraphLabelLastBackflowResponseBody) SetResult(v map[string]interface{}) *GetIgraphLabelLastBackflowResponseBody {
	s.Result = v
	return s
}

type GetIgraphLabelLastBackflowResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIgraphLabelLastBackflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIgraphLabelLastBackflowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphLabelLastBackflowResponse) GoString() string {
	return s.String()
}

func (s *GetIgraphLabelLastBackflowResponse) SetHeaders(v map[string]*string) *GetIgraphLabelLastBackflowResponse {
	s.Headers = v
	return s
}

func (s *GetIgraphLabelLastBackflowResponse) SetStatusCode(v int32) *GetIgraphLabelLastBackflowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIgraphLabelLastBackflowResponse) SetBody(v *GetIgraphLabelLastBackflowResponseBody) *GetIgraphLabelLastBackflowResponse {
	s.Body = v
	return s
}

type GetIgraphTableDetailResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetIgraphTableDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetIgraphTableDetailResponseBody) SetCode(v string) *GetIgraphTableDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetIgraphTableDetailResponseBody) SetMessage(v string) *GetIgraphTableDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetIgraphTableDetailResponseBody) SetRequestId(v string) *GetIgraphTableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIgraphTableDetailResponseBody) SetResult(v map[string]interface{}) *GetIgraphTableDetailResponseBody {
	s.Result = v
	return s
}

type GetIgraphTableDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIgraphTableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIgraphTableDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphTableDetailResponse) GoString() string {
	return s.String()
}

func (s *GetIgraphTableDetailResponse) SetHeaders(v map[string]*string) *GetIgraphTableDetailResponse {
	s.Headers = v
	return s
}

func (s *GetIgraphTableDetailResponse) SetStatusCode(v int32) *GetIgraphTableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIgraphTableDetailResponse) SetBody(v *GetIgraphTableDetailResponseBody) *GetIgraphTableDetailResponse {
	s.Body = v
	return s
}

type GetIgraphTablesBackFlowRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s GetIgraphTablesBackFlowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphTablesBackFlowRequest) GoString() string {
	return s.String()
}

func (s *GetIgraphTablesBackFlowRequest) SetNamespace(v string) *GetIgraphTablesBackFlowRequest {
	s.Namespace = &v
	return s
}

type GetIgraphTablesBackFlowResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetIgraphTablesBackFlowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphTablesBackFlowResponseBody) GoString() string {
	return s.String()
}

func (s *GetIgraphTablesBackFlowResponseBody) SetCode(v string) *GetIgraphTablesBackFlowResponseBody {
	s.Code = &v
	return s
}

func (s *GetIgraphTablesBackFlowResponseBody) SetMessage(v string) *GetIgraphTablesBackFlowResponseBody {
	s.Message = &v
	return s
}

func (s *GetIgraphTablesBackFlowResponseBody) SetRequestId(v string) *GetIgraphTablesBackFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIgraphTablesBackFlowResponseBody) SetResult(v map[string]interface{}) *GetIgraphTablesBackFlowResponseBody {
	s.Result = v
	return s
}

type GetIgraphTablesBackFlowResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetIgraphTablesBackFlowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIgraphTablesBackFlowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIgraphTablesBackFlowResponse) GoString() string {
	return s.String()
}

func (s *GetIgraphTablesBackFlowResponse) SetHeaders(v map[string]*string) *GetIgraphTablesBackFlowResponse {
	s.Headers = v
	return s
}

func (s *GetIgraphTablesBackFlowResponse) SetStatusCode(v int32) *GetIgraphTablesBackFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIgraphTablesBackFlowResponse) SetBody(v *GetIgraphTablesBackFlowResponseBody) *GetIgraphTablesBackFlowResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetNamespace(v string) *GetInstanceRequest {
	s.Namespace = &v
	return s
}

type GetInstanceResponseBody struct {
	Code      *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                        `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GetInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetMessage(v string) *GetInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResult(v *GetInstanceResponseBodyResult) *GetInstanceResponseBody {
	s.Result = v
	return s
}

type GetInstanceResponseBodyResult struct {
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s GetInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyResult) SetResourceGroupId(v string) *GetInstanceResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetInstanceDigestRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s GetInstanceDigestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDigestRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDigestRequest) SetNamespace(v string) *GetInstanceDigestRequest {
	s.Namespace = &v
	return s
}

type GetInstanceDigestResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetInstanceDigestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDigestResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceDigestResponseBody) SetCode(v string) *GetInstanceDigestResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceDigestResponseBody) SetMessage(v string) *GetInstanceDigestResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceDigestResponseBody) SetRequestId(v string) *GetInstanceDigestResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceDigestResponseBody) SetResult(v map[string]interface{}) *GetInstanceDigestResponseBody {
	s.Result = v
	return s
}

type GetInstanceDigestResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetInstanceDigestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetInstanceDigestResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceDigestResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceDigestResponse) SetHeaders(v map[string]*string) *GetInstanceDigestResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceDigestResponse) SetStatusCode(v int32) *GetInstanceDigestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceDigestResponse) SetBody(v *GetInstanceDigestResponseBody) *GetInstanceDigestResponse {
	s.Body = v
	return s
}

type GetTableDetailResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetTableDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDetailResponseBody) SetCode(v string) *GetTableDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableDetailResponseBody) SetMessage(v string) *GetTableDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableDetailResponseBody) SetRequestId(v string) *GetTableDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDetailResponseBody) SetResult(v map[string]interface{}) *GetTableDetailResponseBody {
	s.Result = v
	return s
}

type GetTableDetailResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTableDetailResponse) SetHeaders(v map[string]*string) *GetTableDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTableDetailResponse) SetStatusCode(v int32) *GetTableDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableDetailResponse) SetBody(v *GetTableDetailResponseBody) *GetTableDetailResponse {
	s.Body = v
	return s
}

type GetTableLastBackflowRequest struct {
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
}

func (s GetTableLastBackflowRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableLastBackflowRequest) GoString() string {
	return s.String()
}

func (s *GetTableLastBackflowRequest) SetPartition(v string) *GetTableLastBackflowRequest {
	s.Partition = &v
	return s
}

type GetTableLastBackflowResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetTableLastBackflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableLastBackflowResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableLastBackflowResponseBody) SetCode(v string) *GetTableLastBackflowResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableLastBackflowResponseBody) SetMessage(v string) *GetTableLastBackflowResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableLastBackflowResponseBody) SetRequestId(v string) *GetTableLastBackflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableLastBackflowResponseBody) SetResult(v map[string]interface{}) *GetTableLastBackflowResponseBody {
	s.Result = v
	return s
}

type GetTableLastBackflowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableLastBackflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableLastBackflowResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableLastBackflowResponse) GoString() string {
	return s.String()
}

func (s *GetTableLastBackflowResponse) SetHeaders(v map[string]*string) *GetTableLastBackflowResponse {
	s.Headers = v
	return s
}

func (s *GetTableLastBackflowResponse) SetStatusCode(v int32) *GetTableLastBackflowResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableLastBackflowResponse) SetBody(v *GetTableLastBackflowResponseBody) *GetTableLastBackflowResponse {
	s.Body = v
	return s
}

type ListDemoGraphSchemasResponseBody struct {
	Code         *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message      *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId    *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ReturnTimeMs *string                `json:"returnTimeMs,omitempty" xml:"returnTimeMs,omitempty"`
}

func (s ListDemoGraphSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDemoGraphSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListDemoGraphSchemasResponseBody) SetCode(v string) *ListDemoGraphSchemasResponseBody {
	s.Code = &v
	return s
}

func (s *ListDemoGraphSchemasResponseBody) SetMessage(v string) *ListDemoGraphSchemasResponseBody {
	s.Message = &v
	return s
}

func (s *ListDemoGraphSchemasResponseBody) SetRequestId(v string) *ListDemoGraphSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDemoGraphSchemasResponseBody) SetResult(v map[string]interface{}) *ListDemoGraphSchemasResponseBody {
	s.Result = v
	return s
}

func (s *ListDemoGraphSchemasResponseBody) SetReturnTimeMs(v string) *ListDemoGraphSchemasResponseBody {
	s.ReturnTimeMs = &v
	return s
}

type ListDemoGraphSchemasResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDemoGraphSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDemoGraphSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDemoGraphSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListDemoGraphSchemasResponse) SetHeaders(v map[string]*string) *ListDemoGraphSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListDemoGraphSchemasResponse) SetStatusCode(v int32) *ListDemoGraphSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDemoGraphSchemasResponse) SetBody(v *ListDemoGraphSchemasResponseBody) *ListDemoGraphSchemasResponse {
	s.Body = v
	return s
}

type ListGraphSchemasRequest struct {
	Namespace    *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	PageLimit    *string `json:"pageLimit,omitempty" xml:"pageLimit,omitempty"`
	PageStart    *string `json:"pageStart,omitempty" xml:"pageStart,omitempty"`
	ReturnSpec   *string `json:"returnSpec,omitempty" xml:"returnSpec,omitempty"`
	SchemaStatus *string `json:"schemaStatus,omitempty" xml:"schemaStatus,omitempty"`
}

func (s ListGraphSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGraphSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasRequest) SetNamespace(v string) *ListGraphSchemasRequest {
	s.Namespace = &v
	return s
}

func (s *ListGraphSchemasRequest) SetPageLimit(v string) *ListGraphSchemasRequest {
	s.PageLimit = &v
	return s
}

func (s *ListGraphSchemasRequest) SetPageStart(v string) *ListGraphSchemasRequest {
	s.PageStart = &v
	return s
}

func (s *ListGraphSchemasRequest) SetReturnSpec(v string) *ListGraphSchemasRequest {
	s.ReturnSpec = &v
	return s
}

func (s *ListGraphSchemasRequest) SetSchemaStatus(v string) *ListGraphSchemasRequest {
	s.SchemaStatus = &v
	return s
}

type ListGraphSchemasResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListGraphSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGraphSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasResponseBody) SetCode(v string) *ListGraphSchemasResponseBody {
	s.Code = &v
	return s
}

func (s *ListGraphSchemasResponseBody) SetMessage(v string) *ListGraphSchemasResponseBody {
	s.Message = &v
	return s
}

func (s *ListGraphSchemasResponseBody) SetRequestId(v string) *ListGraphSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGraphSchemasResponseBody) SetResult(v map[string]interface{}) *ListGraphSchemasResponseBody {
	s.Result = v
	return s
}

type ListGraphSchemasResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGraphSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListGraphSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGraphSchemasResponse) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasResponse) SetHeaders(v map[string]*string) *ListGraphSchemasResponse {
	s.Headers = v
	return s
}

func (s *ListGraphSchemasResponse) SetStatusCode(v int32) *ListGraphSchemasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGraphSchemasResponse) SetBody(v *ListGraphSchemasResponseBody) *ListGraphSchemasResponse {
	s.Body = v
	return s
}

type ListIgraphInstancesRequest struct {
	PageContinue    *string                           `json:"pageContinue,omitempty" xml:"pageContinue,omitempty"`
	PageLimit       *string                           `json:"pageLimit,omitempty" xml:"pageLimit,omitempty"`
	ResourceGroupId *string                           `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Tags            []*ListIgraphInstancesRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListIgraphInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIgraphInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListIgraphInstancesRequest) SetPageContinue(v string) *ListIgraphInstancesRequest {
	s.PageContinue = &v
	return s
}

func (s *ListIgraphInstancesRequest) SetPageLimit(v string) *ListIgraphInstancesRequest {
	s.PageLimit = &v
	return s
}

func (s *ListIgraphInstancesRequest) SetResourceGroupId(v string) *ListIgraphInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIgraphInstancesRequest) SetTags(v []*ListIgraphInstancesRequestTags) *ListIgraphInstancesRequest {
	s.Tags = v
	return s
}

type ListIgraphInstancesRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListIgraphInstancesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListIgraphInstancesRequestTags) GoString() string {
	return s.String()
}

func (s *ListIgraphInstancesRequestTags) SetKey(v string) *ListIgraphInstancesRequestTags {
	s.Key = &v
	return s
}

func (s *ListIgraphInstancesRequestTags) SetValue(v string) *ListIgraphInstancesRequestTags {
	s.Value = &v
	return s
}

type ListIgraphInstancesShrinkRequest struct {
	PageContinue    *string `json:"pageContinue,omitempty" xml:"pageContinue,omitempty"`
	PageLimit       *string `json:"pageLimit,omitempty" xml:"pageLimit,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	TagsShrink      *string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s ListIgraphInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIgraphInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListIgraphInstancesShrinkRequest) SetPageContinue(v string) *ListIgraphInstancesShrinkRequest {
	s.PageContinue = &v
	return s
}

func (s *ListIgraphInstancesShrinkRequest) SetPageLimit(v string) *ListIgraphInstancesShrinkRequest {
	s.PageLimit = &v
	return s
}

func (s *ListIgraphInstancesShrinkRequest) SetResourceGroupId(v string) *ListIgraphInstancesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIgraphInstancesShrinkRequest) SetTagsShrink(v string) *ListIgraphInstancesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListIgraphInstancesResponseBody struct {
	Code      *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ListIgraphInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListIgraphInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIgraphInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIgraphInstancesResponseBody) SetCode(v string) *ListIgraphInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIgraphInstancesResponseBody) SetMessage(v string) *ListIgraphInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIgraphInstancesResponseBody) SetRequestId(v string) *ListIgraphInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIgraphInstancesResponseBody) SetResult(v []*ListIgraphInstancesResponseBodyResult) *ListIgraphInstancesResponseBody {
	s.Result = v
	return s
}

type ListIgraphInstancesResponseBodyResult struct {
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
}

func (s ListIgraphInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListIgraphInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListIgraphInstancesResponseBodyResult) SetResourceGroupId(v string) *ListIgraphInstancesResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

type ListIgraphInstancesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListIgraphInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListIgraphInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIgraphInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListIgraphInstancesResponse) SetHeaders(v map[string]*string) *ListIgraphInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListIgraphInstancesResponse) SetStatusCode(v int32) *ListIgraphInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIgraphInstancesResponse) SetBody(v *ListIgraphInstancesResponseBody) *ListIgraphInstancesResponse {
	s.Body = v
	return s
}

type ListInstanceGraphRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListInstanceGraphRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceGraphRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceGraphRequest) SetNamespace(v string) *ListInstanceGraphRequest {
	s.Namespace = &v
	return s
}

type ListInstanceGraphResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListInstanceGraphResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceGraphResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceGraphResponseBody) SetCode(v string) *ListInstanceGraphResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceGraphResponseBody) SetMessage(v string) *ListInstanceGraphResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstanceGraphResponseBody) SetRequestId(v string) *ListInstanceGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceGraphResponseBody) SetResult(v map[string]interface{}) *ListInstanceGraphResponseBody {
	s.Result = v
	return s
}

type ListInstanceGraphResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceGraphResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceGraphResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceGraphResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceGraphResponse) SetHeaders(v map[string]*string) *ListInstanceGraphResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceGraphResponse) SetStatusCode(v int32) *ListInstanceGraphResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceGraphResponse) SetBody(v *ListInstanceGraphResponseBody) *ListInstanceGraphResponse {
	s.Body = v
	return s
}

type PublishGraphSchemaRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s PublishGraphSchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishGraphSchemaRequest) GoString() string {
	return s.String()
}

func (s *PublishGraphSchemaRequest) SetNamespace(v string) *PublishGraphSchemaRequest {
	s.Namespace = &v
	return s
}

type PublishGraphSchemaResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PublishGraphSchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PublishGraphSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *PublishGraphSchemaResponseBody) SetCode(v string) *PublishGraphSchemaResponseBody {
	s.Code = &v
	return s
}

func (s *PublishGraphSchemaResponseBody) SetMessage(v string) *PublishGraphSchemaResponseBody {
	s.Message = &v
	return s
}

func (s *PublishGraphSchemaResponseBody) SetRequestId(v string) *PublishGraphSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishGraphSchemaResponseBody) SetResult(v map[string]interface{}) *PublishGraphSchemaResponseBody {
	s.Result = v
	return s
}

type PublishGraphSchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PublishGraphSchemaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PublishGraphSchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishGraphSchemaResponse) GoString() string {
	return s.String()
}

func (s *PublishGraphSchemaResponse) SetHeaders(v map[string]*string) *PublishGraphSchemaResponse {
	s.Headers = v
	return s
}

func (s *PublishGraphSchemaResponse) SetStatusCode(v int32) *PublishGraphSchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishGraphSchemaResponse) SetBody(v *PublishGraphSchemaResponseBody) *PublishGraphSchemaResponse {
	s.Body = v
	return s
}

type SearchIgraphDemoResponseBody struct {
	Code         *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message      *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId    *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result       map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
	ReturnTimeMs *int64                 `json:"returnTimeMs,omitempty" xml:"returnTimeMs,omitempty"`
}

func (s SearchIgraphDemoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchIgraphDemoResponseBody) GoString() string {
	return s.String()
}

func (s *SearchIgraphDemoResponseBody) SetCode(v string) *SearchIgraphDemoResponseBody {
	s.Code = &v
	return s
}

func (s *SearchIgraphDemoResponseBody) SetMessage(v string) *SearchIgraphDemoResponseBody {
	s.Message = &v
	return s
}

func (s *SearchIgraphDemoResponseBody) SetRequestId(v string) *SearchIgraphDemoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchIgraphDemoResponseBody) SetResult(v map[string]interface{}) *SearchIgraphDemoResponseBody {
	s.Result = v
	return s
}

func (s *SearchIgraphDemoResponseBody) SetReturnTimeMs(v int64) *SearchIgraphDemoResponseBody {
	s.ReturnTimeMs = &v
	return s
}

type SearchIgraphDemoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchIgraphDemoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchIgraphDemoResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchIgraphDemoResponse) GoString() string {
	return s.String()
}

func (s *SearchIgraphDemoResponse) SetHeaders(v map[string]*string) *SearchIgraphDemoResponse {
	s.Headers = v
	return s
}

func (s *SearchIgraphDemoResponse) SetStatusCode(v int32) *SearchIgraphDemoResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchIgraphDemoResponse) SetBody(v *SearchIgraphDemoResponseBody) *SearchIgraphDemoResponse {
	s.Body = v
	return s
}

type TriggerLabelBackflowRequest struct {
	OdpsPartition *string `json:"odpsPartition,omitempty" xml:"odpsPartition,omitempty"`
	Timestamp     *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s TriggerLabelBackflowRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerLabelBackflowRequest) GoString() string {
	return s.String()
}

func (s *TriggerLabelBackflowRequest) SetOdpsPartition(v string) *TriggerLabelBackflowRequest {
	s.OdpsPartition = &v
	return s
}

func (s *TriggerLabelBackflowRequest) SetTimestamp(v string) *TriggerLabelBackflowRequest {
	s.Timestamp = &v
	return s
}

type TriggerLabelBackflowResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s TriggerLabelBackflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerLabelBackflowResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerLabelBackflowResponseBody) SetCode(v string) *TriggerLabelBackflowResponseBody {
	s.Code = &v
	return s
}

func (s *TriggerLabelBackflowResponseBody) SetMessage(v string) *TriggerLabelBackflowResponseBody {
	s.Message = &v
	return s
}

func (s *TriggerLabelBackflowResponseBody) SetRequestId(v string) *TriggerLabelBackflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerLabelBackflowResponseBody) SetResult(v map[string]interface{}) *TriggerLabelBackflowResponseBody {
	s.Result = v
	return s
}

type TriggerLabelBackflowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TriggerLabelBackflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TriggerLabelBackflowResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerLabelBackflowResponse) GoString() string {
	return s.String()
}

func (s *TriggerLabelBackflowResponse) SetHeaders(v map[string]*string) *TriggerLabelBackflowResponse {
	s.Headers = v
	return s
}

func (s *TriggerLabelBackflowResponse) SetStatusCode(v int32) *TriggerLabelBackflowResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerLabelBackflowResponse) SetBody(v *TriggerLabelBackflowResponseBody) *TriggerLabelBackflowResponse {
	s.Body = v
	return s
}

type UpdateGraphDescriptionRequest struct {
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s UpdateGraphDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGraphDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateGraphDescriptionRequest) SetNamespace(v string) *UpdateGraphDescriptionRequest {
	s.Namespace = &v
	return s
}

type UpdateGraphDescriptionResponseBody struct {
	Code      *string                `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateGraphDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGraphDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGraphDescriptionResponseBody) SetCode(v string) *UpdateGraphDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGraphDescriptionResponseBody) SetMessage(v string) *UpdateGraphDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGraphDescriptionResponseBody) SetRequestId(v string) *UpdateGraphDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGraphDescriptionResponseBody) SetResult(v map[string]interface{}) *UpdateGraphDescriptionResponseBody {
	s.Result = v
	return s
}

type UpdateGraphDescriptionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateGraphDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateGraphDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGraphDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateGraphDescriptionResponse) SetHeaders(v map[string]*string) *UpdateGraphDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateGraphDescriptionResponse) SetStatusCode(v int32) *UpdateGraphDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGraphDescriptionResponse) SetBody(v *UpdateGraphDescriptionResponseBody) *UpdateGraphDescriptionResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("igraph"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateGraphWithOptions(instanceId *string, graphName *string, request *CreateGraphRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGraph"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName))),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGraph(instanceId *string, graphName *string, request *CreateGraphRequest) (_result *CreateGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGraphResponse{}
	_body, _err := client.CreateGraphWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGraphSchemaWithOptions(instanceId *string, graphName *string, request *CreateGraphSchemaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGraphSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGraphSchema"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/schemas"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGraphSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGraphSchema(instanceId *string, graphName *string, request *CreateGraphSchemaRequest) (_result *CreateGraphSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGraphSchemaResponse{}
	_body, _err := client.CreateGraphSchemaWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGraphWithOptions(instanceId *string, graphName *string, request *DeleteGraphRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGraph"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGraph(instanceId *string, graphName *string, request *DeleteGraphRequest) (_result *DeleteGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGraphResponse{}
	_body, _err := client.DeleteGraphWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGraphWithOptions(instanceId *string, graphName *string, request *GetGraphRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGraph"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGraph(instanceId *string, graphName *string, request *GetGraphRequest) (_result *GetGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGraphResponse{}
	_body, _err := client.GetGraphWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGraphSchemaWithOptions(instanceId *string, graphName *string, graphSchemaName *string, request *GetGraphSchemaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGraphSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGraphSchema"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/schemas/" + tea.StringValue(openapiutil.GetEncodeParam(graphSchemaName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGraphSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGraphSchema(instanceId *string, graphName *string, graphSchemaName *string, request *GetGraphSchemaRequest) (_result *GetGraphSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGraphSchemaResponse{}
	_body, _err := client.GetGraphSchemaWithOptions(instanceId, graphName, graphSchemaName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIgraphLabelBackFlowWithOptions(graphName *string, instanceId *string, labelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIgraphLabelBackFlowResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIgraphLabelBackFlow"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/label/" + tea.StringValue(openapiutil.GetEncodeParam(labelName)) + "/backflow"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIgraphLabelBackFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIgraphLabelBackFlow(graphName *string, instanceId *string, labelName *string) (_result *GetIgraphLabelBackFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIgraphLabelBackFlowResponse{}
	_body, _err := client.GetIgraphLabelBackFlowWithOptions(graphName, instanceId, labelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIgraphLabelLastBackflowWithOptions(instanceId *string, graphName *string, labelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIgraphLabelLastBackflowResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIgraphLabelLastBackflow"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/label/" + tea.StringValue(openapiutil.GetEncodeParam(labelName)) + "/backflow/last"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIgraphLabelLastBackflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIgraphLabelLastBackflow(instanceId *string, graphName *string, labelName *string) (_result *GetIgraphLabelLastBackflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIgraphLabelLastBackflowResponse{}
	_body, _err := client.GetIgraphLabelLastBackflowWithOptions(instanceId, graphName, labelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIgraphTableDetailWithOptions(instanceId *string, graphName *string, tableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIgraphTableDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetIgraphTableDetail"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName)) + "/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIgraphTableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIgraphTableDetail(instanceId *string, graphName *string, tableName *string) (_result *GetIgraphTableDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIgraphTableDetailResponse{}
	_body, _err := client.GetIgraphTableDetailWithOptions(instanceId, graphName, tableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIgraphTablesBackFlowWithOptions(instanceId *string, graphName *string, request *GetIgraphTablesBackFlowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIgraphTablesBackFlowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIgraphTablesBackFlow"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/backflows"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIgraphTablesBackFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIgraphTablesBackFlow(instanceId *string, graphName *string, request *GetIgraphTablesBackFlowRequest) (_result *GetIgraphTablesBackFlowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIgraphTablesBackFlowResponse{}
	_body, _err := client.GetIgraphTablesBackFlowWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(instanceId *string, request *GetInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(instanceId *string, request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceDigestWithOptions(instanceId *string, request *GetInstanceDigestRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceDigestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceDigest"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/digest"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceDigestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstanceDigest(instanceId *string, request *GetInstanceDigestRequest) (_result *GetInstanceDigestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceDigestResponse{}
	_body, _err := client.GetInstanceDigestWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableDetailWithOptions(instanceId *string, tableName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableDetailResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableDetail"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/tables/" + tea.StringValue(openapiutil.GetEncodeParam(tableName)) + "/detail"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableDetail(instanceId *string, tableName *string) (_result *GetTableDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableDetailResponse{}
	_body, _err := client.GetTableDetailWithOptions(instanceId, tableName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableLastBackflowWithOptions(instanceId *string, tableName *string, request *GetTableLastBackflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableLastBackflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Partition)) {
		query["partition"] = request.Partition
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableLastBackflow"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/table/" + tea.StringValue(openapiutil.GetEncodeParam(tableName)) + "/backflow/last"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableLastBackflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableLastBackflow(instanceId *string, tableName *string, request *GetTableLastBackflowRequest) (_result *GetTableLastBackflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableLastBackflowResponse{}
	_body, _err := client.GetTableLastBackflowWithOptions(instanceId, tableName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDemoGraphSchemasWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDemoGraphSchemasResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListDemoGraphSchemas"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/demo/schemas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDemoGraphSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDemoGraphSchemas() (_result *ListDemoGraphSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDemoGraphSchemasResponse{}
	_body, _err := client.ListDemoGraphSchemasWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGraphSchemasWithOptions(instanceId *string, graphName *string, request *ListGraphSchemasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGraphSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.PageLimit)) {
		query["pageLimit"] = request.PageLimit
	}

	if !tea.BoolValue(util.IsUnset(request.PageStart)) {
		query["pageStart"] = request.PageStart
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnSpec)) {
		query["returnSpec"] = request.ReturnSpec
	}

	if !tea.BoolValue(util.IsUnset(request.SchemaStatus)) {
		query["schemaStatus"] = request.SchemaStatus
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGraphSchemas"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/schemas"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGraphSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGraphSchemas(instanceId *string, graphName *string, request *ListGraphSchemasRequest) (_result *ListGraphSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGraphSchemasResponse{}
	_body, _err := client.ListGraphSchemasWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIgraphInstancesWithOptions(tmpReq *ListIgraphInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIgraphInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListIgraphInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageContinue)) {
		query["pageContinue"] = request.PageContinue
	}

	if !tea.BoolValue(util.IsUnset(request.PageLimit)) {
		query["pageLimit"] = request.PageLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIgraphInstances"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIgraphInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIgraphInstances(request *ListIgraphInstancesRequest) (_result *ListIgraphInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIgraphInstancesResponse{}
	_body, _err := client.ListIgraphInstancesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceGraphWithOptions(instanceId *string, request *ListInstanceGraphRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstanceGraphResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceGraph"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceGraph(instanceId *string, request *ListInstanceGraphRequest) (_result *ListInstanceGraphResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstanceGraphResponse{}
	_body, _err := client.ListInstanceGraphWithOptions(instanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PublishGraphSchemaWithOptions(instanceId *string, graphName *string, graphSchemaName *string, request *PublishGraphSchemaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishGraphSchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PublishGraphSchema"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/schemas/" + tea.StringValue(openapiutil.GetEncodeParam(graphSchemaName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PublishGraphSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PublishGraphSchema(instanceId *string, graphName *string, graphSchemaName *string, request *PublishGraphSchemaRequest) (_result *PublishGraphSchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishGraphSchemaResponse{}
	_body, _err := client.PublishGraphSchemaWithOptions(instanceId, graphName, graphSchemaName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchIgraphDemoWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchIgraphDemoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("SearchIgraphDemo"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/tool/actions/search_demo"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchIgraphDemoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchIgraphDemo() (_result *SearchIgraphDemoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchIgraphDemoResponse{}
	_body, _err := client.SearchIgraphDemoWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TriggerLabelBackflowWithOptions(instanceId *string, graphName *string, labelName *string, request *TriggerLabelBackflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TriggerLabelBackflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OdpsPartition)) {
		query["odpsPartition"] = request.OdpsPartition
	}

	if !tea.BoolValue(util.IsUnset(request.Timestamp)) {
		query["timestamp"] = request.Timestamp
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerLabelBackflow"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/label/" + tea.StringValue(openapiutil.GetEncodeParam(labelName)) + "/backflow"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerLabelBackflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TriggerLabelBackflow(instanceId *string, graphName *string, labelName *string, request *TriggerLabelBackflowRequest) (_result *TriggerLabelBackflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TriggerLabelBackflowResponse{}
	_body, _err := client.TriggerLabelBackflowWithOptions(instanceId, graphName, labelName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGraphDescriptionWithOptions(instanceId *string, graphName *string, request *UpdateGraphDescriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGraphDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["namespace"] = request.Namespace
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGraphDescription"),
		Version:     tea.String("2021-06-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/igraph/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceId)) + "/graphs/" + tea.StringValue(openapiutil.GetEncodeParam(graphName)) + "/description"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGraphDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGraphDescription(instanceId *string, graphName *string, request *UpdateGraphDescriptionRequest) (_result *UpdateGraphDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGraphDescriptionResponse{}
	_body, _err := client.UpdateGraphDescriptionWithOptions(instanceId, graphName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
