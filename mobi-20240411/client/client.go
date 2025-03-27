// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateAppFromTemplateRequest struct {
	// example:
	//
	// [{"type":"bailianapp","name":"rag","actualParameter":"11"}]
	ActualParameters   *string `json:"ActualParameters,omitempty" xml:"ActualParameters,omitempty"`
	ConnectionsContent *string `json:"ConnectionsContent,omitempty" xml:"ConnectionsContent,omitempty"`
	DatabasesContent   *string `json:"DatabasesContent,omitempty" xml:"DatabasesContent,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	From               *string `json:"From,omitempty" xml:"From,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4c892729-9950-4353-8146-33542b869e01
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// Web
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	VariablesContent *string `json:"VariablesContent,omitempty" xml:"VariablesContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1731664463*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAppFromTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateRequest) SetActualParameters(v string) *CreateAppFromTemplateRequest {
	s.ActualParameters = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetConnectionsContent(v string) *CreateAppFromTemplateRequest {
	s.ConnectionsContent = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetDatabasesContent(v string) *CreateAppFromTemplateRequest {
	s.DatabasesContent = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetDescription(v string) *CreateAppFromTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetFrom(v string) *CreateAppFromTemplateRequest {
	s.From = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetIcon(v string) *CreateAppFromTemplateRequest {
	s.Icon = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetName(v string) *CreateAppFromTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetTemplateId(v string) *CreateAppFromTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetType(v string) *CreateAppFromTemplateRequest {
	s.Type = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetVariablesContent(v string) *CreateAppFromTemplateRequest {
	s.VariablesContent = &v
	return s
}

func (s *CreateAppFromTemplateRequest) SetWorkspaceId(v string) *CreateAppFromTemplateRequest {
	s.WorkspaceId = &v
	return s
}

type CreateAppFromTemplateResponseBody struct {
	Data *CreateAppFromTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 36F552F7-E61E-556A-9957-8284318D1B9C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppFromTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppFromTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateResponseBody) SetData(v *CreateAppFromTemplateResponseBodyData) *CreateAppFromTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateAppFromTemplateResponseBody) SetRequestId(v string) *CreateAppFromTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppFromTemplateResponseBodyData struct {
	// example:
	//
	// 172050620*****
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-03-26T10:22Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-03-26T10:22Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// -1
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1731664463*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateAppFromTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateAppFromTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateResponseBodyData) SetAppId(v string) *CreateAppFromTemplateResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetDescription(v string) *CreateAppFromTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetGmtCreate(v string) *CreateAppFromTemplateResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetGmtModified(v string) *CreateAppFromTemplateResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetIcon(v string) *CreateAppFromTemplateResponseBodyData {
	s.Icon = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetName(v string) *CreateAppFromTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateAppFromTemplateResponseBodyData) SetWorkspaceId(v string) *CreateAppFromTemplateResponseBodyData {
	s.WorkspaceId = &v
	return s
}

type CreateAppFromTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppFromTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppFromTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppFromTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppFromTemplateResponse) SetHeaders(v map[string]*string) *CreateAppFromTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppFromTemplateResponse) SetStatusCode(v int32) *CreateAppFromTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppFromTemplateResponse) SetBody(v *CreateAppFromTemplateResponseBody) *CreateAppFromTemplateResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("mobi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 通过模板创建应用
//
// @param request - CreateAppFromTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppFromTemplateResponse
func (client *Client) CreateAppFromTemplateWithOptions(request *CreateAppFromTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateAppFromTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActualParameters)) {
		query["ActualParameters"] = request.ActualParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectionsContent)) {
		query["ConnectionsContent"] = request.ConnectionsContent
	}

	if !tea.BoolValue(util.IsUnset(request.DatabasesContent)) {
		query["DatabasesContent"] = request.DatabasesContent
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.From)) {
		query["From"] = request.From
	}

	if !tea.BoolValue(util.IsUnset(request.Icon)) {
		query["Icon"] = request.Icon
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.VariablesContent)) {
		query["VariablesContent"] = request.VariablesContent
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppFromTemplate"),
		Version:     tea.String("2024-04-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAppFromTemplateResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAppFromTemplateResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 通过模板创建应用
//
// @param request - CreateAppFromTemplateRequest
//
// @return CreateAppFromTemplateResponse
func (client *Client) CreateAppFromTemplate(request *CreateAppFromTemplateRequest) (_result *CreateAppFromTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppFromTemplateResponse{}
	_body, _err := client.CreateAppFromTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
