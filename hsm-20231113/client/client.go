// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfigClusterSubnetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigClusterSubnetRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetRequest) SetClusterId(v string) *ConfigClusterSubnetRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) SetRegionId(v string) *ConfigClusterSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigClusterSubnetRequest) SetVSwitchIds(v []*string) *ConfigClusterSubnetRequest {
	s.VSwitchIds = v
	return s
}

func (s *ConfigClusterSubnetRequest) SetVpcId(v string) *ConfigClusterSubnetRequest {
	s.VpcId = &v
	return s
}

type ConfigClusterSubnetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsg****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-7xvkh90cw39p0****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ConfigClusterSubnetShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetShrinkRequest) SetClusterId(v string) *ConfigClusterSubnetShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetRegionId(v string) *ConfigClusterSubnetShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetVSwitchIdsShrink(v string) *ConfigClusterSubnetShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *ConfigClusterSubnetShrinkRequest) SetVpcId(v string) *ConfigClusterSubnetShrinkRequest {
	s.VpcId = &v
	return s
}

type ConfigClusterSubnetResponseBody struct {
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterSubnetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetResponseBody) SetRequestId(v string) *ConfigClusterSubnetResponseBody {
	s.RequestId = &v
	return s
}

type ConfigClusterSubnetResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterSubnetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterSubnetResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigClusterSubnetResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetResponse) SetHeaders(v map[string]*string) *ConfigClusterSubnetResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterSubnetResponse) SetStatusCode(v int32) *ConfigClusterSubnetResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterSubnetResponse) SetBody(v *ConfigClusterSubnetResponseBody) *ConfigClusterSubnetResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("hsm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 配置集群子网
//
// @param tmpReq - ConfigClusterSubnetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConfigClusterSubnetResponse
func (client *Client) ConfigClusterSubnetWithOptions(tmpReq *ConfigClusterSubnetRequest, runtime *util.RuntimeOptions) (_result *ConfigClusterSubnetResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ConfigClusterSubnetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.VSwitchIds)) {
		request.VSwitchIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VSwitchIds, tea.String("VSwitchIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdsShrink)) {
		body["VSwitchIds"] = request.VSwitchIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigClusterSubnet"),
		Version:     tea.String("2023-11-13"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigClusterSubnetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 配置集群子网
//
// @param request - ConfigClusterSubnetRequest
//
// @return ConfigClusterSubnetResponse
func (client *Client) ConfigClusterSubnet(request *ConfigClusterSubnetRequest) (_result *ConfigClusterSubnetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigClusterSubnetResponse{}
	_body, _err := client.ConfigClusterSubnetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
