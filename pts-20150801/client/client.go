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

type CreateTransactionRequest struct {
	ScriptId        *int32  `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	TransactionName *string `json:"TransactionName,omitempty" xml:"TransactionName,omitempty"`
}

func (s CreateTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionRequest) GoString() string {
	return s.String()
}

func (s *CreateTransactionRequest) SetScriptId(v int32) *CreateTransactionRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateTransactionRequest) SetTransactionName(v string) *CreateTransactionRequest {
	s.TransactionName = &v
	return s
}

type CreateTransactionResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TransactionId *int32  `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CreateTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransactionResponseBody) SetRequestId(v string) *CreateTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransactionResponseBody) SetTransactionId(v int32) *CreateTransactionResponseBody {
	s.TransactionId = &v
	return s
}

type CreateTransactionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionResponse) GoString() string {
	return s.String()
}

func (s *CreateTransactionResponse) SetHeaders(v map[string]*string) *CreateTransactionResponse {
	s.Headers = v
	return s
}

func (s *CreateTransactionResponse) SetStatusCode(v int32) *CreateTransactionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransactionResponse) SetBody(v *CreateTransactionResponseBody) *CreateTransactionResponse {
	s.Body = v
	return s
}

type GetKeySecretResponseBody struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Secret    *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
}

func (s GetKeySecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetKeySecretResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeySecretResponseBody) SetKey(v string) *GetKeySecretResponseBody {
	s.Key = &v
	return s
}

func (s *GetKeySecretResponseBody) SetRequestId(v string) *GetKeySecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeySecretResponseBody) SetSecret(v string) *GetKeySecretResponseBody {
	s.Secret = &v
	return s
}

type GetKeySecretResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKeySecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKeySecretResponse) String() string {
	return tea.Prettify(s)
}

func (s GetKeySecretResponse) GoString() string {
	return s.String()
}

func (s *GetKeySecretResponse) SetHeaders(v map[string]*string) *GetKeySecretResponse {
	s.Headers = v
	return s
}

func (s *GetKeySecretResponse) SetStatusCode(v int32) *GetKeySecretResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKeySecretResponse) SetBody(v *GetKeySecretResponseBody) *GetKeySecretResponse {
	s.Body = v
	return s
}

type GetScriptRequest struct {
	ScriptId *int32  `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	Tfsname  *string `json:"Tfsname,omitempty" xml:"Tfsname,omitempty"`
}

func (s GetScriptRequest) String() string {
	return tea.Prettify(s)
}

func (s GetScriptRequest) GoString() string {
	return s.String()
}

func (s *GetScriptRequest) SetScriptId(v int32) *GetScriptRequest {
	s.ScriptId = &v
	return s
}

func (s *GetScriptRequest) SetTfsname(v string) *GetScriptRequest {
	s.Tfsname = &v
	return s
}

type GetScriptResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Script    *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s GetScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptResponseBody) SetRequestId(v string) *GetScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptResponseBody) SetScript(v string) *GetScriptResponseBody {
	s.Script = &v
	return s
}

type GetScriptResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScriptResponse) GoString() string {
	return s.String()
}

func (s *GetScriptResponse) SetHeaders(v map[string]*string) *GetScriptResponse {
	s.Headers = v
	return s
}

func (s *GetScriptResponse) SetStatusCode(v int32) *GetScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScriptResponse) SetBody(v *GetScriptResponseBody) *GetScriptResponse {
	s.Body = v
	return s
}

type GetTasksRequest struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTasksRequest) GoString() string {
	return s.String()
}

func (s *GetTasksRequest) SetStatus(v string) *GetTasksRequest {
	s.Status = &v
	return s
}

type GetTasksResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     *string `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s GetTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTasksResponseBody) GoString() string {
	return s.String()
}

func (s *GetTasksResponseBody) SetRequestId(v string) *GetTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTasksResponseBody) SetTasks(v string) *GetTasksResponseBody {
	s.Tasks = &v
	return s
}

type GetTasksResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTasksResponse) GoString() string {
	return s.String()
}

func (s *GetTasksResponse) SetHeaders(v map[string]*string) *GetTasksResponse {
	s.Headers = v
	return s
}

func (s *GetTasksResponse) SetStatusCode(v int32) *GetTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTasksResponse) SetBody(v *GetTasksResponseBody) *GetTasksResponse {
	s.Body = v
	return s
}

type ReportLogSampleRequest struct {
	LogSample  *string `json:"LogSample,omitempty" xml:"LogSample,omitempty"`
	ScenarioId *int32  `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	Wskey      *string `json:"Wskey,omitempty" xml:"Wskey,omitempty"`
}

func (s ReportLogSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportLogSampleRequest) GoString() string {
	return s.String()
}

func (s *ReportLogSampleRequest) SetLogSample(v string) *ReportLogSampleRequest {
	s.LogSample = &v
	return s
}

func (s *ReportLogSampleRequest) SetScenarioId(v int32) *ReportLogSampleRequest {
	s.ScenarioId = &v
	return s
}

func (s *ReportLogSampleRequest) SetWskey(v string) *ReportLogSampleRequest {
	s.Wskey = &v
	return s
}

type ReportLogSampleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportLogSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportLogSampleResponseBody) GoString() string {
	return s.String()
}

func (s *ReportLogSampleResponseBody) SetRequestId(v string) *ReportLogSampleResponseBody {
	s.RequestId = &v
	return s
}

type ReportLogSampleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportLogSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportLogSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportLogSampleResponse) GoString() string {
	return s.String()
}

func (s *ReportLogSampleResponse) SetHeaders(v map[string]*string) *ReportLogSampleResponse {
	s.Headers = v
	return s
}

func (s *ReportLogSampleResponse) SetStatusCode(v int32) *ReportLogSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportLogSampleResponse) SetBody(v *ReportLogSampleResponseBody) *ReportLogSampleResponse {
	s.Body = v
	return s
}

type ReportTestSampleRequest struct {
	TestSample *string `json:"TestSample,omitempty" xml:"TestSample,omitempty"`
}

func (s ReportTestSampleRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTestSampleRequest) GoString() string {
	return s.String()
}

func (s *ReportTestSampleRequest) SetTestSample(v string) *ReportTestSampleRequest {
	s.TestSample = &v
	return s
}

type ReportTestSampleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportTestSampleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportTestSampleResponseBody) GoString() string {
	return s.String()
}

func (s *ReportTestSampleResponseBody) SetRequestId(v string) *ReportTestSampleResponseBody {
	s.RequestId = &v
	return s
}

type ReportTestSampleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportTestSampleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportTestSampleResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportTestSampleResponse) GoString() string {
	return s.String()
}

func (s *ReportTestSampleResponse) SetHeaders(v map[string]*string) *ReportTestSampleResponse {
	s.Headers = v
	return s
}

func (s *ReportTestSampleResponse) SetStatusCode(v int32) *ReportTestSampleResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportTestSampleResponse) SetBody(v *ReportTestSampleResponseBody) *ReportTestSampleResponse {
	s.Body = v
	return s
}

type ReportVuserRequest struct {
	GmtCreated *int64  `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	ScenarioId *int32  `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	Vuser      *int32  `json:"Vuser,omitempty" xml:"Vuser,omitempty"`
	Wskey      *string `json:"Wskey,omitempty" xml:"Wskey,omitempty"`
}

func (s ReportVuserRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportVuserRequest) GoString() string {
	return s.String()
}

func (s *ReportVuserRequest) SetGmtCreated(v int64) *ReportVuserRequest {
	s.GmtCreated = &v
	return s
}

func (s *ReportVuserRequest) SetScenarioId(v int32) *ReportVuserRequest {
	s.ScenarioId = &v
	return s
}

func (s *ReportVuserRequest) SetVuser(v int32) *ReportVuserRequest {
	s.Vuser = &v
	return s
}

func (s *ReportVuserRequest) SetWskey(v string) *ReportVuserRequest {
	s.Wskey = &v
	return s
}

type ReportVuserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportVuserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportVuserResponseBody) GoString() string {
	return s.String()
}

func (s *ReportVuserResponseBody) SetRequestId(v string) *ReportVuserResponseBody {
	s.RequestId = &v
	return s
}

type ReportVuserResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportVuserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportVuserResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportVuserResponse) GoString() string {
	return s.String()
}

func (s *ReportVuserResponse) SetHeaders(v map[string]*string) *ReportVuserResponse {
	s.Headers = v
	return s
}

func (s *ReportVuserResponse) SetStatusCode(v int32) *ReportVuserResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportVuserResponse) SetBody(v *ReportVuserResponseBody) *ReportVuserResponse {
	s.Body = v
	return s
}

type SendWangWangRequest struct {
	Msg   *string                `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Title *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	To    map[string]interface{} `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendWangWangRequest) String() string {
	return tea.Prettify(s)
}

func (s SendWangWangRequest) GoString() string {
	return s.String()
}

func (s *SendWangWangRequest) SetMsg(v string) *SendWangWangRequest {
	s.Msg = &v
	return s
}

func (s *SendWangWangRequest) SetTitle(v string) *SendWangWangRequest {
	s.Title = &v
	return s
}

func (s *SendWangWangRequest) SetTo(v map[string]interface{}) *SendWangWangRequest {
	s.To = v
	return s
}

type SendWangWangShrinkRequest struct {
	Msg      *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
	ToShrink *string `json:"To,omitempty" xml:"To,omitempty"`
}

func (s SendWangWangShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendWangWangShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendWangWangShrinkRequest) SetMsg(v string) *SendWangWangShrinkRequest {
	s.Msg = &v
	return s
}

func (s *SendWangWangShrinkRequest) SetTitle(v string) *SendWangWangShrinkRequest {
	s.Title = &v
	return s
}

func (s *SendWangWangShrinkRequest) SetToShrink(v string) *SendWangWangShrinkRequest {
	s.ToShrink = &v
	return s
}

type SendWangWangResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendWangWangResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendWangWangResponseBody) GoString() string {
	return s.String()
}

func (s *SendWangWangResponseBody) SetRequestId(v string) *SendWangWangResponseBody {
	s.RequestId = &v
	return s
}

type SendWangWangResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendWangWangResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendWangWangResponse) String() string {
	return tea.Prettify(s)
}

func (s SendWangWangResponse) GoString() string {
	return s.String()
}

func (s *SendWangWangResponse) SetHeaders(v map[string]*string) *SendWangWangResponse {
	s.Headers = v
	return s
}

func (s *SendWangWangResponse) SetStatusCode(v int32) *SendWangWangResponse {
	s.StatusCode = &v
	return s
}

func (s *SendWangWangResponse) SetBody(v *SendWangWangResponseBody) *SendWangWangResponse {
	s.Body = v
	return s
}

type SetScenarioStatusRequest struct {
	NodeIp     *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	ScenarioId *int32  `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Wskey      *string `json:"Wskey,omitempty" xml:"Wskey,omitempty"`
}

func (s SetScenarioStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetScenarioStatusRequest) GoString() string {
	return s.String()
}

func (s *SetScenarioStatusRequest) SetNodeIp(v string) *SetScenarioStatusRequest {
	s.NodeIp = &v
	return s
}

func (s *SetScenarioStatusRequest) SetScenarioId(v int32) *SetScenarioStatusRequest {
	s.ScenarioId = &v
	return s
}

func (s *SetScenarioStatusRequest) SetStatus(v int32) *SetScenarioStatusRequest {
	s.Status = &v
	return s
}

func (s *SetScenarioStatusRequest) SetWskey(v string) *SetScenarioStatusRequest {
	s.Wskey = &v
	return s
}

type SetScenarioStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetScenarioStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetScenarioStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetScenarioStatusResponseBody) SetRequestId(v string) *SetScenarioStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetScenarioStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetScenarioStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetScenarioStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetScenarioStatusResponse) GoString() string {
	return s.String()
}

func (s *SetScenarioStatusResponse) SetHeaders(v map[string]*string) *SetScenarioStatusResponse {
	s.Headers = v
	return s
}

func (s *SetScenarioStatusResponse) SetStatusCode(v int32) *SetScenarioStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetScenarioStatusResponse) SetBody(v *SetScenarioStatusResponseBody) *SetScenarioStatusResponse {
	s.Body = v
	return s
}

type SetTaskStatusRequest struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Wskey  *string `json:"Wskey,omitempty" xml:"Wskey,omitempty"`
}

func (s SetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *SetTaskStatusRequest) SetStatus(v string) *SetTaskStatusRequest {
	s.Status = &v
	return s
}

func (s *SetTaskStatusRequest) SetWskey(v string) *SetTaskStatusRequest {
	s.Wskey = &v
	return s
}

type SetTaskStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetTaskStatusResponseBody) SetRequestId(v string) *SetTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *SetTaskStatusResponse) SetHeaders(v map[string]*string) *SetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *SetTaskStatusResponse) SetStatusCode(v int32) *SetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTaskStatusResponse) SetBody(v *SetTaskStatusResponseBody) *SetTaskStatusResponse {
	s.Body = v
	return s
}

type StopTaskRequest struct {
	Msg    *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	TaskId *int32  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StopTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTaskRequest) SetMsg(v string) *StopTaskRequest {
	s.Msg = &v
	return s
}

func (s *StopTaskRequest) SetTaskId(v int32) *StopTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StopTaskRequest) SetType(v string) *StopTaskRequest {
	s.Type = &v
	return s
}

type StopTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopTaskResponseBody) SetRequestId(v string) *StopTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopTaskResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTaskResponse) GoString() string {
	return s.String()
}

func (s *StopTaskResponse) SetHeaders(v map[string]*string) *StopTaskResponse {
	s.Headers = v
	return s
}

func (s *StopTaskResponse) SetStatusCode(v int32) *StopTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTaskResponse) SetBody(v *StopTaskResponseBody) *StopTaskResponse {
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
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
 * @deprecated
 *
 * @param request CreateTransactionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateTransactionResponse
 */
// Deprecated
func (client *Client) CreateTransactionWithOptions(request *CreateTransactionRequest, runtime *util.RuntimeOptions) (_result *CreateTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScriptId)) {
		query["ScriptId"] = request.ScriptId
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionName)) {
		query["TransactionName"] = request.TransactionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTransaction"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTransactionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request CreateTransactionRequest
 * @return CreateTransactionResponse
 */
// Deprecated
func (client *Client) CreateTransaction(request *CreateTransactionRequest) (_result *CreateTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransactionResponse{}
	_body, _err := client.CreateTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetKeySecretRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetKeySecretResponse
 */
// Deprecated
func (client *Client) GetKeySecretWithOptions(runtime *util.RuntimeOptions) (_result *GetKeySecretResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetKeySecret"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetKeySecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @return GetKeySecretResponse
 */
// Deprecated
func (client *Client) GetKeySecret() (_result *GetKeySecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetKeySecretResponse{}
	_body, _err := client.GetKeySecretWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetScriptRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetScriptResponse
 */
// Deprecated
func (client *Client) GetScriptWithOptions(request *GetScriptRequest, runtime *util.RuntimeOptions) (_result *GetScriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ScriptId)) {
		query["ScriptId"] = request.ScriptId
	}

	if !tea.BoolValue(util.IsUnset(request.Tfsname)) {
		query["Tfsname"] = request.Tfsname
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetScript"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetScriptRequest
 * @return GetScriptResponse
 */
// Deprecated
func (client *Client) GetScript(request *GetScriptRequest) (_result *GetScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetScriptResponse{}
	_body, _err := client.GetScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetTasksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTasksResponse
 */
// Deprecated
func (client *Client) GetTasksWithOptions(request *GetTasksRequest, runtime *util.RuntimeOptions) (_result *GetTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTasks"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request GetTasksRequest
 * @return GetTasksResponse
 */
// Deprecated
func (client *Client) GetTasks(request *GetTasksRequest) (_result *GetTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTasksResponse{}
	_body, _err := client.GetTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ReportLogSampleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReportLogSampleResponse
 */
// Deprecated
func (client *Client) ReportLogSampleWithOptions(request *ReportLogSampleRequest, runtime *util.RuntimeOptions) (_result *ReportLogSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogSample)) {
		query["LogSample"] = request.LogSample
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioId)) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !tea.BoolValue(util.IsUnset(request.Wskey)) {
		query["Wskey"] = request.Wskey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportLogSample"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportLogSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ReportLogSampleRequest
 * @return ReportLogSampleResponse
 */
// Deprecated
func (client *Client) ReportLogSample(request *ReportLogSampleRequest) (_result *ReportLogSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportLogSampleResponse{}
	_body, _err := client.ReportLogSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ReportTestSampleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReportTestSampleResponse
 */
// Deprecated
func (client *Client) ReportTestSampleWithOptions(request *ReportTestSampleRequest, runtime *util.RuntimeOptions) (_result *ReportTestSampleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TestSample)) {
		query["TestSample"] = request.TestSample
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportTestSample"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportTestSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ReportTestSampleRequest
 * @return ReportTestSampleResponse
 */
// Deprecated
func (client *Client) ReportTestSample(request *ReportTestSampleRequest) (_result *ReportTestSampleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportTestSampleResponse{}
	_body, _err := client.ReportTestSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ReportVuserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReportVuserResponse
 */
// Deprecated
func (client *Client) ReportVuserWithOptions(request *ReportVuserRequest, runtime *util.RuntimeOptions) (_result *ReportVuserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GmtCreated)) {
		query["GmtCreated"] = request.GmtCreated
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioId)) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !tea.BoolValue(util.IsUnset(request.Vuser)) {
		query["Vuser"] = request.Vuser
	}

	if !tea.BoolValue(util.IsUnset(request.Wskey)) {
		query["Wskey"] = request.Wskey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReportVuser"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportVuserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request ReportVuserRequest
 * @return ReportVuserResponse
 */
// Deprecated
func (client *Client) ReportVuser(request *ReportVuserRequest) (_result *ReportVuserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReportVuserResponse{}
	_body, _err := client.ReportVuserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param tmpReq SendWangWangRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SendWangWangResponse
 */
// Deprecated
func (client *Client) SendWangWangWithOptions(tmpReq *SendWangWangRequest, runtime *util.RuntimeOptions) (_result *SendWangWangResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendWangWangShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.To)) {
		request.ToShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.To, tea.String("To"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		query["Msg"] = request.Msg
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.ToShrink)) {
		query["To"] = request.ToShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendWangWang"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendWangWangResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request SendWangWangRequest
 * @return SendWangWangResponse
 */
// Deprecated
func (client *Client) SendWangWang(request *SendWangWangRequest) (_result *SendWangWangResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendWangWangResponse{}
	_body, _err := client.SendWangWangWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request SetScenarioStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetScenarioStatusResponse
 */
// Deprecated
func (client *Client) SetScenarioStatusWithOptions(request *SetScenarioStatusRequest, runtime *util.RuntimeOptions) (_result *SetScenarioStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		query["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.ScenarioId)) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Wskey)) {
		query["Wskey"] = request.Wskey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetScenarioStatus"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetScenarioStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request SetScenarioStatusRequest
 * @return SetScenarioStatusResponse
 */
// Deprecated
func (client *Client) SetScenarioStatus(request *SetScenarioStatusRequest) (_result *SetScenarioStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetScenarioStatusResponse{}
	_body, _err := client.SetScenarioStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request SetTaskStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetTaskStatusResponse
 */
// Deprecated
func (client *Client) SetTaskStatusWithOptions(request *SetTaskStatusRequest, runtime *util.RuntimeOptions) (_result *SetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Wskey)) {
		query["Wskey"] = request.Wskey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTaskStatus"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request SetTaskStatusRequest
 * @return SetTaskStatusResponse
 */
// Deprecated
func (client *Client) SetTaskStatus(request *SetTaskStatusRequest) (_result *SetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTaskStatusResponse{}
	_body, _err := client.SetTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request StopTaskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopTaskResponse
 */
// Deprecated
func (client *Client) StopTaskWithOptions(request *StopTaskRequest, runtime *util.RuntimeOptions) (_result *StopTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Msg)) {
		query["Msg"] = request.Msg
	}

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
		Action:      tea.String("StopTask"),
		Version:     tea.String("2015-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated
 *
 * @param request StopTaskRequest
 * @return StopTaskResponse
 */
// Deprecated
func (client *Client) StopTask(request *StopTaskRequest) (_result *StopTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopTaskResponse{}
	_body, _err := client.StopTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
