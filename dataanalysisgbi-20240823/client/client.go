// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RunDataAnalysisRequest struct {
	// example:
	//
	// true
	GenerateSqlOnly *bool `json:"generateSqlOnly,omitempty" xml:"generateSqlOnly,omitempty"`
	// This parameter is required.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// sessionID
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// STANDARD_MIX
	SpecificationType *string `json:"specificationType,omitempty" xml:"specificationType,omitempty"`
}

func (s RunDataAnalysisRequest) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisRequest) SetGenerateSqlOnly(v bool) *RunDataAnalysisRequest {
	s.GenerateSqlOnly = &v
	return s
}

func (s *RunDataAnalysisRequest) SetQuery(v string) *RunDataAnalysisRequest {
	s.Query = &v
	return s
}

func (s *RunDataAnalysisRequest) SetSessionId(v string) *RunDataAnalysisRequest {
	s.SessionId = &v
	return s
}

func (s *RunDataAnalysisRequest) SetSpecificationType(v string) *RunDataAnalysisRequest {
	s.SpecificationType = &v
	return s
}

type RunDataAnalysisResponseBody struct {
	Code           *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data           *RunDataAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int64                           `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                          `json:"message,omitempty" xml:"message,omitempty"`
}

func (s RunDataAnalysisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBody) SetCode(v string) *RunDataAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *RunDataAnalysisResponseBody) SetData(v *RunDataAnalysisResponseBodyData) *RunDataAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *RunDataAnalysisResponseBody) SetHttpStatusCode(v int64) *RunDataAnalysisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RunDataAnalysisResponseBody) SetMessage(v string) *RunDataAnalysisResponseBody {
	s.Message = &v
	return s
}

type RunDataAnalysisResponseBodyData struct {
	// example:
	//
	// Access was denied, message: No such namespace namespaces/tech-scp-chain7.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// rewrite
	Event          *string `json:"event,omitempty" xml:"event,omitempty"`
	Evidence       *string `json:"evidence,omitempty" xml:"evidence,omitempty"`
	HttpStatusCode *int64  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// DA2578F7-88A5-5D6E-9305-33E724E97D60
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Rewrite   *string   `json:"rewrite,omitempty" xml:"rewrite,omitempty"`
	Selector  []*string `json:"selector,omitempty" xml:"selector,omitempty" type:"Repeated"`
	// example:
	//
	// sessionid1
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// example:
	//
	// select p.product_id, p.product_name, sum(o.quantity) as total_sales from products p join orders o on p.product_id = o.product_id where o.order_date between \\"2022-10-22\\" and \\"2024-10-22\\" group by p.product_id, p.product_name having total_sales > 5
	Sql     *string                                 `json:"sql,omitempty" xml:"sql,omitempty"`
	SqlData *RunDataAnalysisResponseBodyDataSqlData `json:"sqlData,omitempty" xml:"sqlData,omitempty" type:"Struct"`
	// example:
	//
	// Can not issue data manipulation statements with executeQuery()
	SqlError      *string                                       `json:"sqlError,omitempty" xml:"sqlError,omitempty"`
	Visualization *RunDataAnalysisResponseBodyDataVisualization `json:"visualization,omitempty" xml:"visualization,omitempty" type:"Struct"`
}

func (s RunDataAnalysisResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyData) SetErrorMessage(v string) *RunDataAnalysisResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetEvent(v string) *RunDataAnalysisResponseBodyData {
	s.Event = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetEvidence(v string) *RunDataAnalysisResponseBodyData {
	s.Evidence = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetHttpStatusCode(v int64) *RunDataAnalysisResponseBodyData {
	s.HttpStatusCode = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetRequestId(v string) *RunDataAnalysisResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetRewrite(v string) *RunDataAnalysisResponseBodyData {
	s.Rewrite = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSelector(v []*string) *RunDataAnalysisResponseBodyData {
	s.Selector = v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSessionId(v string) *RunDataAnalysisResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSql(v string) *RunDataAnalysisResponseBodyData {
	s.Sql = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSqlData(v *RunDataAnalysisResponseBodyDataSqlData) *RunDataAnalysisResponseBodyData {
	s.SqlData = v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetSqlError(v string) *RunDataAnalysisResponseBodyData {
	s.SqlError = &v
	return s
}

func (s *RunDataAnalysisResponseBodyData) SetVisualization(v *RunDataAnalysisResponseBodyDataVisualization) *RunDataAnalysisResponseBodyData {
	s.Visualization = v
	return s
}

type RunDataAnalysisResponseBodyDataSqlData struct {
	Column []*string                `json:"column,omitempty" xml:"column,omitempty" type:"Repeated"`
	Data   []map[string]interface{} `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s RunDataAnalysisResponseBodyDataSqlData) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyDataSqlData) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyDataSqlData) SetColumn(v []*string) *RunDataAnalysisResponseBodyDataSqlData {
	s.Column = v
	return s
}

func (s *RunDataAnalysisResponseBodyDataSqlData) SetData(v []map[string]interface{}) *RunDataAnalysisResponseBodyDataSqlData {
	s.Data = v
	return s
}

type RunDataAnalysisResponseBodyDataVisualization struct {
	Data *RunDataAnalysisResponseBodyDataVisualizationData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Text *string                                           `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunDataAnalysisResponseBodyDataVisualization) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyDataVisualization) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyDataVisualization) SetData(v *RunDataAnalysisResponseBodyDataVisualizationData) *RunDataAnalysisResponseBodyDataVisualization {
	s.Data = v
	return s
}

func (s *RunDataAnalysisResponseBodyDataVisualization) SetText(v string) *RunDataAnalysisResponseBodyDataVisualization {
	s.Text = &v
	return s
}

type RunDataAnalysisResponseBodyDataVisualizationData struct {
	// example:
	//
	// bar
	PlotType *string   `json:"plotType,omitempty" xml:"plotType,omitempty"`
	XAxis    []*string `json:"xAxis,omitempty" xml:"xAxis,omitempty" type:"Repeated"`
	YAxis    []*string `json:"yAxis,omitempty" xml:"yAxis,omitempty" type:"Repeated"`
}

func (s RunDataAnalysisResponseBodyDataVisualizationData) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponseBodyDataVisualizationData) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponseBodyDataVisualizationData) SetPlotType(v string) *RunDataAnalysisResponseBodyDataVisualizationData {
	s.PlotType = &v
	return s
}

func (s *RunDataAnalysisResponseBodyDataVisualizationData) SetXAxis(v []*string) *RunDataAnalysisResponseBodyDataVisualizationData {
	s.XAxis = v
	return s
}

func (s *RunDataAnalysisResponseBodyDataVisualizationData) SetYAxis(v []*string) *RunDataAnalysisResponseBodyDataVisualizationData {
	s.YAxis = v
	return s
}

type RunDataAnalysisResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunDataAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDataAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s RunDataAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunDataAnalysisResponse) SetHeaders(v map[string]*string) *RunDataAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunDataAnalysisResponse) SetStatusCode(v int32) *RunDataAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDataAnalysisResponse) SetBody(v *RunDataAnalysisResponseBody) *RunDataAnalysisResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dataanalysisgbi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 运行数据分析
//
// @param request - RunDataAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDataAnalysisResponse
func (client *Client) RunDataAnalysisWithOptions(workspaceId *string, request *RunDataAnalysisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunDataAnalysisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GenerateSqlOnly)) {
		body["generateSqlOnly"] = request.GenerateSqlOnly
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["sessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationType)) {
		body["specificationType"] = request.SpecificationType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunDataAnalysis"),
		Version:     tea.String("2024-08-23"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/gbi/runDataAnalysis"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunDataAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行数据分析
//
// @param request - RunDataAnalysisRequest
//
// @return RunDataAnalysisResponse
func (client *Client) RunDataAnalysis(workspaceId *string, request *RunDataAnalysisRequest) (_result *RunDataAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunDataAnalysisResponse{}
	_body, _err := client.RunDataAnalysisWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
