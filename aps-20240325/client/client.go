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

type GetFxCustomerTypeRequest struct {
	Uid *int64 `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetFxCustomerTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFxCustomerTypeRequest) GoString() string {
	return s.String()
}

func (s *GetFxCustomerTypeRequest) SetUid(v int64) *GetFxCustomerTypeRequest {
	s.Uid = &v
	return s
}

type GetFxCustomerTypeResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetFxCustomerTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFxCustomerTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFxCustomerTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetFxCustomerTypeResponseBody) SetCode(v string) *GetFxCustomerTypeResponseBody {
	s.Code = &v
	return s
}

func (s *GetFxCustomerTypeResponseBody) SetData(v *GetFxCustomerTypeResponseBodyData) *GetFxCustomerTypeResponseBody {
	s.Data = v
	return s
}

func (s *GetFxCustomerTypeResponseBody) SetMessage(v string) *GetFxCustomerTypeResponseBody {
	s.Message = &v
	return s
}

func (s *GetFxCustomerTypeResponseBody) SetRequestId(v string) *GetFxCustomerTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFxCustomerTypeResponseBody) SetSuccess(v bool) *GetFxCustomerTypeResponseBody {
	s.Success = &v
	return s
}

type GetFxCustomerTypeResponseBodyData struct {
	CustomerType *int32 `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	IsSub        *int32 `json:"IsSub,omitempty" xml:"IsSub,omitempty"`
	ParentId     *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s GetFxCustomerTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetFxCustomerTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFxCustomerTypeResponseBodyData) SetCustomerType(v int32) *GetFxCustomerTypeResponseBodyData {
	s.CustomerType = &v
	return s
}

func (s *GetFxCustomerTypeResponseBodyData) SetIsSub(v int32) *GetFxCustomerTypeResponseBodyData {
	s.IsSub = &v
	return s
}

func (s *GetFxCustomerTypeResponseBodyData) SetParentId(v int64) *GetFxCustomerTypeResponseBodyData {
	s.ParentId = &v
	return s
}

type GetFxCustomerTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFxCustomerTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFxCustomerTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFxCustomerTypeResponse) GoString() string {
	return s.String()
}

func (s *GetFxCustomerTypeResponse) SetHeaders(v map[string]*string) *GetFxCustomerTypeResponse {
	s.Headers = v
	return s
}

func (s *GetFxCustomerTypeResponse) SetStatusCode(v int32) *GetFxCustomerTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFxCustomerTypeResponse) SetBody(v *GetFxCustomerTypeResponseBody) *GetFxCustomerTypeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("aps"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetFxCustomerTypeWithOptions(request *GetFxCustomerTypeRequest, runtime *util.RuntimeOptions) (_result *GetFxCustomerTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFxCustomerType"),
		Version:     tea.String("2024-03-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFxCustomerTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFxCustomerType(request *GetFxCustomerTypeRequest) (_result *GetFxCustomerTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFxCustomerTypeResponse{}
	_body, _err := client.GetFxCustomerTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
