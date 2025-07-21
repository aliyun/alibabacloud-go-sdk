// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.SignatureAlgorithm = dara.String("v2")
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("wyota"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备激活
//
// @param request - ActivateDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateDeviceResponse
func (client *Client) ActivateDeviceWithOptions(request *ActivateDeviceRequest, runtime *dara.RuntimeOptions) (_result *ActivateDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateDevice"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备激活
//
// @param request - ActivateDeviceRequest
//
// @return ActivateDeviceResponse
func (client *Client) ActivateDevice(request *ActivateDeviceRequest) (_result *ActivateDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ActivateDeviceResponse{}
	_body, _err := client.ActivateDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过序列号添加设备
//
// @param request - AddDeviceFromSNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceFromSNResponse
func (client *Client) AddDeviceFromSNWithOptions(request *AddDeviceFromSNRequest, runtime *dara.RuntimeOptions) (_result *AddDeviceFromSNResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.CustomProperty) {
		body["CustomProperty"] = request.CustomProperty
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.LabelContents) {
		body["LabelContents"] = request.LabelContents
	}

	if !dara.IsNil(request.SecureNetworkType) {
		body["SecureNetworkType"] = request.SecureNetworkType
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDeviceFromSN"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDeviceFromSNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过序列号添加设备
//
// @param request - AddDeviceFromSNRequest
//
// @return AddDeviceFromSNResponse
func (client *Client) AddDeviceFromSN(request *AddDeviceFromSNRequest) (_result *AddDeviceFromSNResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDeviceFromSNResponse{}
	_body, _err := client.AddDeviceFromSNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增设备座位和标签
//
// @param request - AddDeviceSeatsAndLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceSeatsAndLabelsResponse
func (client *Client) AddDeviceSeatsAndLabelsWithOptions(request *AddDeviceSeatsAndLabelsRequest, runtime *dara.RuntimeOptions) (_result *AddDeviceSeatsAndLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IsUnique) {
		body["IsUnique"] = request.IsUnique
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.LabelList) {
		body["LabelList"] = request.LabelList
	}

	if !dara.IsNil(request.SeatName) {
		body["SeatName"] = request.SeatName
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDeviceSeatsAndLabels"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDeviceSeatsAndLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增设备座位和标签
//
// @param request - AddDeviceSeatsAndLabelsRequest
//
// @return AddDeviceSeatsAndLabelsResponse
func (client *Client) AddDeviceSeatsAndLabels(request *AddDeviceSeatsAndLabelsRequest) (_result *AddDeviceSeatsAndLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDeviceSeatsAndLabelsResponse{}
	_body, _err := client.AddDeviceSeatsAndLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过CSV文件添加设备
//
// @param request - AddDevicesFromCSVRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDevicesFromCSVResponse
func (client *Client) AddDevicesFromCSVWithOptions(request *AddDevicesFromCSVRequest, runtime *dara.RuntimeOptions) (_result *AddDevicesFromCSVResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.SeatCol) {
		body["SeatCol"] = request.SeatCol
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.SiteName) {
		body["SiteName"] = request.SiteName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDevicesFromCSV"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDevicesFromCSVResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过CSV文件添加设备
//
// @param request - AddDevicesFromCSVRequest
//
// @return AddDevicesFromCSVResponse
func (client *Client) AddDevicesFromCSV(request *AddDevicesFromCSVRequest) (_result *AddDevicesFromCSVResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDevicesFromCSVResponse{}
	_body, _err := client.AddDevicesFromCSVWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加标签
//
// @param request - AddLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLabelsResponse
func (client *Client) AddLabelsWithOptions(request *AddLabelsRequest, runtime *dara.RuntimeOptions) (_result *AddLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelContents) {
		body["LabelContents"] = request.LabelContents
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLabels"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加标签
//
// @param request - AddLabelsRequest
//
// @return AddLabelsResponse
func (client *Client) AddLabels(request *AddLabelsRequest) (_result *AddLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddLabelsResponse{}
	_body, _err := client.AddLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 新增或更新设备工位
//
// @param request - AddOrUpdateDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddOrUpdateDeviceSeatsResponse
func (client *Client) AddOrUpdateDeviceSeatsWithOptions(request *AddOrUpdateDeviceSeatsRequest, runtime *dara.RuntimeOptions) (_result *AddOrUpdateDeviceSeatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.UserCustomId) {
		body["UserCustomId"] = request.UserCustomId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddOrUpdateDeviceSeats"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddOrUpdateDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 新增或更新设备工位
//
// @param request - AddOrUpdateDeviceSeatsRequest
//
// @return AddOrUpdateDeviceSeatsResponse
func (client *Client) AddOrUpdateDeviceSeats(request *AddOrUpdateDeviceSeatsRequest) (_result *AddOrUpdateDeviceSeatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddOrUpdateDeviceSeatsResponse{}
	_body, _err := client.AddOrUpdateDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalResponse
func (client *Client) AddTerminalWithOptions(request *AddTerminalRequest, runtime *dara.RuntimeOptions) (_result *AddTerminalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTerminal"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalRequest
//
// @return AddTerminalResponse
func (client *Client) AddTerminal(request *AddTerminalRequest) (_result *AddTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTerminalResponse{}
	_body, _err := client.AddTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalsResponse
func (client *Client) AddTerminalsWithOptions(request *AddTerminalsRequest, runtime *dara.RuntimeOptions) (_result *AddTerminalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AddTerminalParams) {
		bodyFlat["AddTerminalParams"] = request.AddTerminalParams
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTerminals"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalsRequest
//
// @return AddTerminalsResponse
func (client *Client) AddTerminals(request *AddTerminalsRequest) (_result *AddTerminalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTerminalsResponse{}
	_body, _err := client.AddTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备绑定终端用户
//
// @param request - AttachEndUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachEndUsersResponse
func (client *Client) AttachEndUsersWithOptions(request *AttachEndUsersRequest, runtime *dara.RuntimeOptions) (_result *AttachEndUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachEndUsers"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachEndUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备绑定终端用户
//
// @param request - AttachEndUsersRequest
//
// @return AttachEndUsersResponse
func (client *Client) AttachEndUsers(request *AttachEndUsersRequest) (_result *AttachEndUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachEndUsersResponse{}
	_body, _err := client.AttachEndUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - AttachLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLabelResponse
func (client *Client) AttachLabelWithOptions(request *AttachLabelRequest, runtime *dara.RuntimeOptions) (_result *AttachLabelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachLabel"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - AttachLabelRequest
//
// @return AttachLabelResponse
func (client *Client) AttachLabel(request *AttachLabelRequest) (_result *AttachLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachLabelResponse{}
	_body, _err := client.AttachLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量绑定标签
//
// @param request - AttachLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachLabelsResponse
func (client *Client) AttachLabelsWithOptions(request *AttachLabelsRequest, runtime *dara.RuntimeOptions) (_result *AttachLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelIds) {
		body["LabelIds"] = request.LabelIds
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.SerialNoList) {
		body["SerialNoList"] = request.SerialNoList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachLabels"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量绑定标签
//
// @param request - AttachLabelsRequest
//
// @return AttachLabelsResponse
func (client *Client) AttachLabels(request *AttachLabelsRequest) (_result *AttachLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachLabelsResponse{}
	_body, _err := client.AttachLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindAccountLessLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUserWithOptions(request *BindAccountLessLoginUserRequest, runtime *dara.RuntimeOptions) (_result *BindAccountLessLoginUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAccountLessLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAccountLessLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindAccountLessLoginUserRequest
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUser(request *BindAccountLessLoginUserRequest) (_result *BindAccountLessLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAccountLessLoginUserResponse{}
	_body, _err := client.BindAccountLessLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindPasswordFreeLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPasswordFreeLoginUserResponse
func (client *Client) BindPasswordFreeLoginUserWithOptions(request *BindPasswordFreeLoginUserRequest, runtime *dara.RuntimeOptions) (_result *BindPasswordFreeLoginUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPasswordFreeLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPasswordFreeLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindPasswordFreeLoginUserRequest
//
// @return BindPasswordFreeLoginUserResponse
func (client *Client) BindPasswordFreeLoginUser(request *BindPasswordFreeLoginUserRequest) (_result *BindPasswordFreeLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPasswordFreeLoginUserResponse{}
	_body, _err := client.BindPasswordFreeLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查uuid有效性
//
// @param request - CheckUuidValidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUuidValidResponse
func (client *Client) CheckUuidValidWithOptions(request *CheckUuidValidRequest, runtime *dara.RuntimeOptions) (_result *CheckUuidValidResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bluetooth) {
		body["Bluetooth"] = request.Bluetooth
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.ChipId) {
		body["ChipId"] = request.ChipId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.CustomId) {
		body["CustomId"] = request.CustomId
	}

	if !dara.IsNil(request.EtherMac) {
		body["EtherMac"] = request.EtherMac
	}

	if !dara.IsNil(request.HostOsInfo) {
		body["HostOsInfo"] = request.HostOsInfo
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.Wlan) {
		body["Wlan"] = request.Wlan
	}

	if !dara.IsNil(request.WosAppVersion) {
		body["WosAppVersion"] = request.WosAppVersion
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckUuidValid"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckUuidValidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查uuid有效性
//
// @param request - CheckUuidValidRequest
//
// @return CheckUuidValidResponse
func (client *Client) CheckUuidValid(request *CheckUuidValidRequest) (_result *CheckUuidValidResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckUuidValidResponse{}
	_body, _err := client.CheckUuidValidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateAppOtaTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppOtaTaskResponse
func (client *Client) CreateAppOtaTaskWithOptions(request *CreateAppOtaTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAppOtaTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersionUid) {
		query["AppVersionUid"] = request.AppVersionUid
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ClientIdList) {
		query["ClientIdList"] = request.ClientIdList
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ForceUpgrade) {
		query["ForceUpgrade"] = request.ForceUpgrade
	}

	if !dara.IsNil(request.Label) {
		query["Label"] = request.Label
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.Regions) {
		query["Regions"] = request.Regions
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.TenantIdList) {
		query["TenantIdList"] = request.TenantIdList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppOtaTask"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppOtaTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务
//
// @param request - CreateAppOtaTaskRequest
//
// @return CreateAppOtaTaskResponse
func (client *Client) CreateAppOtaTask(request *CreateAppOtaTaskRequest) (_result *CreateAppOtaTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppOtaTaskResponse{}
	_body, _err := client.CreateAppOtaTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建版本
//
// @param request - CreateAppOtaVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppOtaVersionResponse
func (client *Client) CreateAppOtaVersionWithOptions(request *CreateAppOtaVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateAppOtaVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Arch) {
		query["Arch"] = request.Arch
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.DownloadUrl) {
		query["DownloadUrl"] = request.DownloadUrl
	}

	if !dara.IsNil(request.Md5) {
		query["Md5"] = request.Md5
	}

	if !dara.IsNil(request.Os) {
		query["Os"] = request.Os
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.OtaType) {
		query["OtaType"] = request.OtaType
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.RelationVersionUids) {
		query["RelationVersionUids"] = request.RelationVersionUids
	}

	if !dara.IsNil(request.ReleaseNote) {
		query["ReleaseNote"] = request.ReleaseNote
	}

	if !dara.IsNil(request.ReleaseNoteEn) {
		query["ReleaseNoteEn"] = request.ReleaseNoteEn
	}

	if !dara.IsNil(request.ReleaseNoteJp) {
		query["ReleaseNoteJp"] = request.ReleaseNoteJp
	}

	if !dara.IsNil(request.Size) {
		query["Size"] = request.Size
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.SnapshotRegionId) {
		query["SnapshotRegionId"] = request.SnapshotRegionId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VersionType) {
		query["VersionType"] = request.VersionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAppOtaVersion"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAppOtaVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建版本
//
// @param request - CreateAppOtaVersionRequest
//
// @return CreateAppOtaVersionResponse
func (client *Client) CreateAppOtaVersion(request *CreateAppOtaVersionRequest) (_result *CreateAppOtaVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAppOtaVersionResponse{}
	_body, _err := client.CreateAppOtaVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除版本
//
// @param request - DeleteAppOtaVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppOtaVersionsResponse
func (client *Client) DeleteAppOtaVersionsWithOptions(request *DeleteAppOtaVersionsRequest, runtime *dara.RuntimeOptions) (_result *DeleteAppOtaVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VersionUidList) {
		query["VersionUidList"] = request.VersionUidList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAppOtaVersions"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAppOtaVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除版本
//
// @param request - DeleteAppOtaVersionsRequest
//
// @return DeleteAppOtaVersionsResponse
func (client *Client) DeleteAppOtaVersions(request *DeleteAppOtaVersionsRequest) (_result *DeleteAppOtaVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAppOtaVersionsResponse{}
	_body, _err := client.DeleteAppOtaVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除设备
//
// @param request - DeleteDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDevicesResponse
func (client *Client) DeleteDevicesWithOptions(request *DeleteDevicesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Uuids) {
		query["Uuids"] = request.Uuids
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.SerialNos) {
		body["SerialNos"] = request.SerialNos
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDevices"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除设备
//
// @param request - DeleteDevicesRequest
//
// @return DeleteDevicesResponse
func (client *Client) DeleteDevices(request *DeleteDevicesRequest) (_result *DeleteDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDevicesResponse{}
	_body, _err := client.DeleteDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLabelResponse
func (client *Client) DeleteLabelWithOptions(request *DeleteLabelRequest, runtime *dara.RuntimeOptions) (_result *DeleteLabelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Force) {
		body["Force"] = request.Force
	}

	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLabel"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除标签
//
// @param request - DeleteLabelRequest
//
// @return DeleteLabelResponse
func (client *Client) DeleteLabel(request *DeleteLabelRequest) (_result *DeleteLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLabelResponse{}
	_body, _err := client.DeleteLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询版本
//
// @param request - DescribeAppOtaVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppOtaVersionResponse
func (client *Client) DescribeAppOtaVersionWithOptions(request *DescribeAppOtaVersionRequest, runtime *dara.RuntimeOptions) (_result *DescribeAppOtaVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.NullChannel) {
		query["NullChannel"] = request.NullChannel
	}

	if !dara.IsNil(request.OtaType) {
		query["OtaType"] = request.OtaType
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.VersionUid) {
		query["VersionUid"] = request.VersionUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAppOtaVersion"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAppOtaVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询版本
//
// @param request - DescribeAppOtaVersionRequest
//
// @return DescribeAppOtaVersionResponse
func (client *Client) DescribeAppOtaVersion(request *DescribeAppOtaVersionRequest) (_result *DescribeAppOtaVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAppOtaVersionResponse{}
	_body, _err := client.DescribeAppOtaVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备座位
//
// @param request - DescribeDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceSeatsResponse
func (client *Client) DescribeDeviceSeatsWithOptions(request *DescribeDeviceSeatsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceSeatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.SerialNoList) {
		body["SerialNoList"] = request.SerialNoList
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceSeats"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备座位
//
// @param request - DescribeDeviceSeatsRequest
//
// @return DescribeDeviceSeatsResponse
func (client *Client) DescribeDeviceSeats(request *DescribeDeviceSeatsRequest) (_result *DescribeDeviceSeatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceSeatsResponse{}
	_body, _err := client.DescribeDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询版本信息
//
// @param request - DescribeDeviceVersionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceVersionDetailResponse
func (client *Client) DescribeDeviceVersionDetailWithOptions(request *DescribeDeviceVersionDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceVersionDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.NetworkType) {
		body["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.VersionName) {
		body["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceVersionDetail"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceVersionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询版本信息
//
// @param request - DescribeDeviceVersionDetailRequest
//
// @return DescribeDeviceVersionDetailResponse
func (client *Client) DescribeDeviceVersionDetail(request *DescribeDeviceVersionDetailRequest) (_result *DescribeDeviceVersionDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDeviceVersionDetailResponse{}
	_body, _err := client.DescribeDeviceVersionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备标签数量
//
// @param request - DescribeSnLabelCountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSnLabelCountsResponse
func (client *Client) DescribeSnLabelCountsWithOptions(request *DescribeSnLabelCountsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSnLabelCountsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelList) {
		body["LabelList"] = request.LabelList
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneName) {
		body["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSnLabelCounts"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSnLabelCountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备标签数量
//
// @param request - DescribeSnLabelCountsRequest
//
// @return DescribeSnLabelCountsResponse
func (client *Client) DescribeSnLabelCounts(request *DescribeSnLabelCountsRequest) (_result *DescribeSnLabelCountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSnLabelCountsResponse{}
	_body, _err := client.DescribeSnLabelCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工作区域
//
// @param request - DescribeWorkZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeWorkZonesResponse
func (client *Client) DescribeWorkZonesWithOptions(request *DescribeWorkZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeWorkZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ZoneIdList) {
		body["ZoneIdList"] = request.ZoneIdList
	}

	if !dara.IsNil(request.ZoneNameList) {
		body["ZoneNameList"] = request.ZoneNameList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeWorkZones"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeWorkZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工作区域
//
// @param request - DescribeWorkZonesRequest
//
// @return DescribeWorkZonesResponse
func (client *Client) DescribeWorkZones(request *DescribeWorkZonesRequest) (_result *DescribeWorkZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeWorkZonesResponse{}
	_body, _err := client.DescribeWorkZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备解绑终端用户
//
// @param request - DetachEndUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachEndUsersResponse
func (client *Client) DetachEndUsersWithOptions(request *DetachEndUsersRequest, runtime *dara.RuntimeOptions) (_result *DetachEndUsersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserIds) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachEndUsers"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachEndUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备解绑终端用户
//
// @param request - DetachEndUsersRequest
//
// @return DetachEndUsersResponse
func (client *Client) DetachEndUsers(request *DetachEndUsersRequest) (_result *DetachEndUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachEndUsersResponse{}
	_body, _err := client.DetachEndUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - DetachLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachLabelResponse
func (client *Client) DetachLabelWithOptions(request *DetachLabelRequest, runtime *dara.RuntimeOptions) (_result *DetachLabelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachLabel"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备绑定标签
//
// @param request - DetachLabelRequest
//
// @return DetachLabelResponse
func (client *Client) DetachLabel(request *DetachLabelRequest) (_result *DetachLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachLabelResponse{}
	_body, _err := client.DetachLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量解绑标签
//
// @param request - DetachLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachLabelsResponse
func (client *Client) DetachLabelsWithOptions(request *DetachLabelsRequest, runtime *dara.RuntimeOptions) (_result *DetachLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelIds) {
		body["LabelIds"] = request.LabelIds
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.SerialNoList) {
		body["SerialNoList"] = request.SerialNoList
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachLabels"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量解绑标签
//
// @param request - DetachLabelsRequest
//
// @return DetachLabelsResponse
func (client *Client) DetachLabels(request *DetachLabelsRequest) (_result *DetachLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachLabelsResponse{}
	_body, _err := client.DetachLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户上传的文件
//
// @param request - GenerateOssUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateOssUrlResponse
func (client *Client) GenerateOssUrlWithOptions(request *GenerateOssUrlRequest, runtime *dara.RuntimeOptions) (_result *GenerateOssUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ObjectNameList) {
		body["ObjectNameList"] = request.ObjectNameList
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateOssUrl"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateOssUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户上传的文件
//
// @param request - GenerateOssUrlRequest
//
// @return GenerateOssUrlResponse
func (client *Client) GenerateOssUrl(request *GenerateOssUrlRequest) (_result *GenerateOssUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateOssUrlResponse{}
	_body, _err := client.GenerateOssUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取应用最新版本
//
// @param request - GetAppOtaLatestVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppOtaLatestVersionResponse
func (client *Client) GetAppOtaLatestVersionWithOptions(request *GetAppOtaLatestVersionRequest, runtime *dara.RuntimeOptions) (_result *GetAppOtaLatestVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseVersion) {
		query["BaseVersion"] = request.BaseVersion
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientUid) {
		query["ClientUid"] = request.ClientUid
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppOtaLatestVersion"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppOtaLatestVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用最新版本
//
// @param request - GetAppOtaLatestVersionRequest
//
// @return GetAppOtaLatestVersionResponse
func (client *Client) GetAppOtaLatestVersion(request *GetAppOtaLatestVersionRequest) (_result *GetAppOtaLatestVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAppOtaLatestVersionResponse{}
	_body, _err := client.GetAppOtaLatestVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备配置
//
// @param request - GetDeviceConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceConfigsResponse
func (client *Client) GetDeviceConfigsWithOptions(request *GetDeviceConfigsRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		body["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.NetworkType) {
		body["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.UrclVersion) {
		body["UrclVersion"] = request.UrclVersion
	}

	if !dara.IsNil(request.UserCustomId) {
		body["UserCustomId"] = request.UserCustomId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceConfigs"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备配置
//
// @param request - GetDeviceConfigsRequest
//
// @return GetDeviceConfigsResponse
func (client *Client) GetDeviceConfigs(request *GetDeviceConfigsRequest) (_result *GetDeviceConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceConfigsResponse{}
	_body, _err := client.GetDeviceConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取是否开启自动升级状态
//
// @param request - GetDeviceOtaAutoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaAutoStatusResponse
func (client *Client) GetDeviceOtaAutoStatusWithOptions(request *GetDeviceOtaAutoStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceOtaAutoStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceOtaAutoStatus"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceOtaAutoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取是否开启自动升级状态
//
// @param request - GetDeviceOtaAutoStatusRequest
//
// @return GetDeviceOtaAutoStatusResponse
func (client *Client) GetDeviceOtaAutoStatus(request *GetDeviceOtaAutoStatusRequest) (_result *GetDeviceOtaAutoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceOtaAutoStatusResponse{}
	_body, _err := client.GetDeviceOtaAutoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaInfoResponse
func (client *Client) GetDeviceOtaInfoWithOptions(request *GetDeviceOtaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceOtaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseVersion) {
		body["BaseVersion"] = request.BaseVersion
	}

	if !dara.IsNil(request.Channel) {
		body["Channel"] = request.Channel
	}

	if !dara.IsNil(request.DeviceId) {
		body["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.NetworkType) {
		body["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.OsVersion) {
		body["OsVersion"] = request.OsVersion
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TargetVersionType) {
		body["TargetVersionType"] = request.TargetVersionType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceOtaInfo"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoRequest
//
// @return GetDeviceOtaInfoResponse
func (client *Client) GetDeviceOtaInfo(request *GetDeviceOtaInfoRequest) (_result *GetDeviceOtaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceOtaInfoResponse{}
	_body, _err := client.GetDeviceOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaInfoTestResponse
func (client *Client) GetDeviceOtaInfoTestWithOptions(request *GetDeviceOtaInfoTestRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceOtaInfoTestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseVersion) {
		body["BaseVersion"] = request.BaseVersion
	}

	if !dara.IsNil(request.DeviceId) {
		body["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceOtaInfoTest"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceOtaInfoTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备升级信息
//
// @param request - GetDeviceOtaInfoTestRequest
//
// @return GetDeviceOtaInfoTestResponse
func (client *Client) GetDeviceOtaInfoTest(request *GetDeviceOtaInfoTestRequest) (_result *GetDeviceOtaInfoTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceOtaInfoTestResponse{}
	_body, _err := client.GetDeviceOtaInfoTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户任务版本信息
//
// @param request - GetDeviceOtaTaskVersionInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOtaTaskVersionInfoResponse
func (client *Client) GetDeviceOtaTaskVersionInfoWithOptions(request *GetDeviceOtaTaskVersionInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceOtaTaskVersionInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceOtaTaskVersionInfo"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceOtaTaskVersionInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户任务版本信息
//
// @param request - GetDeviceOtaTaskVersionInfoRequest
//
// @return GetDeviceOtaTaskVersionInfoResponse
func (client *Client) GetDeviceOtaTaskVersionInfo(request *GetDeviceOtaTaskVersionInfoRequest) (_result *GetDeviceOtaTaskVersionInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceOtaTaskVersionInfoResponse{}
	_body, _err := client.GetDeviceOtaTaskVersionInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获得设备升级详情
//
// @param request - GetDeviceUpgradeStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceUpgradeStatusResponse
func (client *Client) GetDeviceUpgradeStatusWithOptions(request *GetDeviceUpgradeStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceUpgradeStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppVersion) {
		query["AppVersion"] = request.AppVersion
	}

	if !dara.IsNil(request.ClientUid) {
		query["ClientUid"] = request.ClientUid
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.TaskUid) {
		query["TaskUid"] = request.TaskUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceUpgradeStatus"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceUpgradeStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获得设备升级详情
//
// @param request - GetDeviceUpgradeStatusRequest
//
// @return GetDeviceUpgradeStatusResponse
func (client *Client) GetDeviceUpgradeStatus(request *GetDeviceUpgradeStatusRequest) (_result *GetDeviceUpgradeStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceUpgradeStatusResponse{}
	_body, _err := client.GetDeviceUpgradeStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导出设备工位列表
//
// @param request - GetExportDeviceInfoOssUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExportDeviceInfoOssUrlResponse
func (client *Client) GetExportDeviceInfoOssUrlWithOptions(request *GetExportDeviceInfoOssUrlRequest, runtime *dara.RuntimeOptions) (_result *GetExportDeviceInfoOssUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	if !dara.IsNil(request.ZoneName) {
		body["ZoneName"] = request.ZoneName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExportDeviceInfoOssUrl"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExportDeviceInfoOssUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导出设备工位列表
//
// @param request - GetExportDeviceInfoOssUrlRequest
//
// @return GetExportDeviceInfoOssUrlResponse
func (client *Client) GetExportDeviceInfoOssUrl(request *GetExportDeviceInfoOssUrlRequest) (_result *GetExportDeviceInfoOssUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExportDeviceInfoOssUrlResponse{}
	_body, _err := client.GetExportDeviceInfoOssUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询OSS配置信息
//
// @param request - GetFbOssConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFbOssConfigResponse
func (client *Client) GetFbOssConfigWithOptions(request *GetFbOssConfigRequest, runtime *dara.RuntimeOptions) (_result *GetFbOssConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AreaSite) {
		body["AreaSite"] = request.AreaSite
	}

	if !dara.IsNil(request.DirPrefix) {
		body["DirPrefix"] = request.DirPrefix
	}

	if !dara.IsNil(request.IsDedicatedLine) {
		body["IsDedicatedLine"] = request.IsDedicatedLine
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFbOssConfig"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFbOssConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询OSS配置信息
//
// @param request - GetFbOssConfigRequest
//
// @return GetFbOssConfigResponse
func (client *Client) GetFbOssConfig(request *GetFbOssConfigRequest) (_result *GetFbOssConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFbOssConfigResponse{}
	_body, _err := client.GetFbOssConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取OSS配置
//
// @param request - GetOssConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssConfigResponse
func (client *Client) GetOssConfigWithOptions(request *GetOssConfigRequest, runtime *dara.RuntimeOptions) (_result *GetOssConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssConfig"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取OSS配置
//
// @param request - GetOssConfigRequest
//
// @return GetOssConfigResponse
func (client *Client) GetOssConfig(request *GetOssConfigRequest) (_result *GetOssConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOssConfigResponse{}
	_body, _err := client.GetOssConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取版本下载地址
//
// @param request - GetVersionDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVersionDownloadUrlResponse
func (client *Client) GetVersionDownloadUrlWithOptions(request *GetVersionDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetVersionDownloadUrlResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVersionDownloadUrl"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVersionDownloadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取版本下载地址
//
// @param request - GetVersionDownloadUrlRequest
//
// @return GetVersionDownloadUrlResponse
func (client *Client) GetVersionDownloadUrl(request *GetVersionDownloadUrlRequest) (_result *GetVersionDownloadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVersionDownloadUrlResponse{}
	_body, _err := client.GetVersionDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户已绑定的可信设备列表
//
// @param request - ListBoundDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBoundDevicesResponse
func (client *Client) ListBoundDevicesWithOptions(request *ListBoundDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListBoundDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		body["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DirectoryId) {
		body["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	if !dara.IsNil(request.LastLoginUser) {
		body["LastLoginUser"] = request.LastLoginUser
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.UserType) {
		body["UserType"] = request.UserType
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBoundDevices"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBoundDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户已绑定的可信设备列表
//
// @param request - ListBoundDevicesRequest
//
// @return ListBoundDevicesResponse
func (client *Client) ListBoundDevices(request *ListBoundDevicesRequest) (_result *ListBoundDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBoundDevicesResponse{}
	_body, _err := client.ListBoundDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户ota任务
//
// @param request - ListDeviceOtaTaskByTenantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceOtaTaskByTenantResponse
func (client *Client) ListDeviceOtaTaskByTenantWithOptions(request *ListDeviceOtaTaskByTenantRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceOtaTaskByTenantResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceOtaTaskByTenant"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceOtaTaskByTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户ota任务
//
// @param request - ListDeviceOtaTaskByTenantRequest
//
// @return ListDeviceOtaTaskByTenantResponse
func (client *Client) ListDeviceOtaTaskByTenant(request *ListDeviceOtaTaskByTenantRequest) (_result *ListDeviceOtaTaskByTenantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeviceOtaTaskByTenantResponse{}
	_body, _err := client.ListDeviceOtaTaskByTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备座位列表
//
// @param request - ListDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceSeatsResponse
func (client *Client) ListDeviceSeatsWithOptions(request *ListDeviceSeatsRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceSeatsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.SeatNo) {
		body["SeatNo"] = request.SeatNo
	}

	if !dara.IsNil(request.SerialNoList) {
		body["SerialNoList"] = request.SerialNoList
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.ZoneId) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceSeats"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备座位列表
//
// @param request - ListDeviceSeatsRequest
//
// @return ListDeviceSeatsResponse
func (client *Client) ListDeviceSeats(request *ListDeviceSeatsRequest) (_result *ListDeviceSeatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeviceSeatsResponse{}
	_body, _err := client.ListDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取设备列表
//
// @param request - ListDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDevicesResponse
func (client *Client) ListDevicesWithOptions(request *ListDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DeviceIpV4) {
		query["DeviceIpV4"] = request.DeviceIpV4
	}

	if !dara.IsNil(request.DeviceName) {
		query["DeviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.DeviceOS) {
		query["DeviceOS"] = request.DeviceOS
	}

	if !dara.IsNil(request.DevicePlatform) {
		query["DevicePlatform"] = request.DevicePlatform
	}

	if !dara.IsNil(request.LastLoginUser) {
		query["LastLoginUser"] = request.LastLoginUser
	}

	if !dara.IsNil(request.LocationInfo) {
		query["LocationInfo"] = request.LocationInfo
	}

	if !dara.IsNil(request.UserType) {
		query["UserType"] = request.UserType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.DeviceGroupId) {
		body["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDevices"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取设备列表
//
// @param request - ListDevicesRequest
//
// @return ListDevicesResponse
func (client *Client) ListDevices(request *ListDevicesRequest) (_result *ListDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDevicesResponse{}
	_body, _err := client.ListDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户问题标签
//
// @param request - ListFbIssueLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFbIssueLabelsResponse
func (client *Client) ListFbIssueLabelsWithOptions(runtime *dara.RuntimeOptions) (_result *ListFbIssueLabelsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListFbIssueLabels"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFbIssueLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户问题标签
//
// @return ListFbIssueLabelsResponse
func (client *Client) ListFbIssueLabels() (_result *ListFbIssueLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFbIssueLabelsResponse{}
	_body, _err := client.ListFbIssueLabelsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据语言类型和调用方查询标签列表
//
// @param request - ListFbIssueLabelsByLCRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFbIssueLabelsByLCResponse
func (client *Client) ListFbIssueLabelsByLCWithOptions(request *ListFbIssueLabelsByLCRequest, runtime *dara.RuntimeOptions) (_result *ListFbIssueLabelsByLCResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Caller) {
		body["Caller"] = request.Caller
	}

	if !dara.IsNil(request.LanguageType) {
		body["LanguageType"] = request.LanguageType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFbIssueLabelsByLC"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFbIssueLabelsByLCResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据语言类型和调用方查询标签列表
//
// @param request - ListFbIssueLabelsByLCRequest
//
// @return ListFbIssueLabelsByLCResponse
func (client *Client) ListFbIssueLabelsByLC(request *ListFbIssueLabelsByLCRequest) (_result *ListFbIssueLabelsByLCResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFbIssueLabelsByLCResponse{}
	_body, _err := client.ListFbIssueLabelsByLCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取标签列表
//
// @param request - ListLabelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLabelsResponse
func (client *Client) ListLabelsWithOptions(request *ListLabelsRequest, runtime *dara.RuntimeOptions) (_result *ListLabelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLabels"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLabelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取标签列表
//
// @param request - ListLabelsRequest
//
// @return ListLabelsResponse
func (client *Client) ListLabels(request *ListLabelsRequest) (_result *ListLabelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLabelsResponse{}
	_body, _err := client.ListLabelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取租户升级设备信息
//
// @param request - ListTenantDeviceOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTenantDeviceOtaInfoResponse
func (client *Client) ListTenantDeviceOtaInfoWithOptions(request *ListTenantDeviceOtaInfoRequest, runtime *dara.RuntimeOptions) (_result *ListTenantDeviceOtaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTenantDeviceOtaInfo"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTenantDeviceOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取租户升级设备信息
//
// @param request - ListTenantDeviceOtaInfoRequest
//
// @return ListTenantDeviceOtaInfoResponse
func (client *Client) ListTenantDeviceOtaInfo(request *ListTenantDeviceOtaInfoRequest) (_result *ListTenantDeviceOtaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTenantDeviceOtaInfoResponse{}
	_body, _err := client.ListTenantDeviceOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询终端列表
//
// @param request - ListTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalResponse
func (client *Client) ListTerminalWithOptions(request *ListTerminalRequest, runtime *dara.RuntimeOptions) (_result *ListTerminalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	if !dara.IsNil(request.Ipv4) {
		body["Ipv4"] = request.Ipv4
	}

	if !dara.IsNil(request.LocationInfo) {
		body["LocationInfo"] = request.LocationInfo
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchKeyword) {
		body["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTerminal"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询终端列表
//
// @param request - ListTerminalRequest
//
// @return ListTerminalResponse
func (client *Client) ListTerminal(request *ListTerminalRequest) (_result *ListTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTerminalResponse{}
	_body, _err := client.ListTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询终端基本信息
//
// @param request - ListTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalsResponse
func (client *Client) ListTerminalsWithOptions(request *ListTerminalsRequest, runtime *dara.RuntimeOptions) (_result *ListTerminalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PasswordFreeLoginUser) {
		body["PasswordFreeLoginUser"] = request.PasswordFreeLoginUser
	}

	if !dara.IsNil(request.SearchKeyword) {
		body["SearchKeyword"] = request.SearchKeyword
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.SerialNumbers) {
		bodyFlat["SerialNumbers"] = request.SerialNumbers
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !dara.IsNil(request.Uuids) {
		bodyFlat["Uuids"] = request.Uuids
	}

	if !dara.IsNil(request.WithBindUser) {
		body["WithBindUser"] = request.WithBindUser
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTerminals"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询终端基本信息
//
// @param request - ListTerminalsRequest
//
// @return ListTerminalsResponse
func (client *Client) ListTerminals(request *ListTerminalsRequest) (_result *ListTerminalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTerminalsResponse{}
	_body, _err := client.ListTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询可信设备列表
//
// @param request - ListTrustDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTrustDevicesResponse
func (client *Client) ListTrustDevicesWithOptions(request *ListTrustDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListTrustDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.UserCustomId) {
		body["UserCustomId"] = request.UserCustomId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTrustDevices"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTrustDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可信设备列表
//
// @param request - ListTrustDevicesRequest
//
// @return ListTrustDevicesResponse
func (client *Client) ListTrustDevices(request *ListTrustDevicesRequest) (_result *ListTrustDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTrustDevicesResponse{}
	_body, _err := client.ListTrustDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户未绑定的可信设备列表
//
// @param request - ListUnbindDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUnbindDevicesResponse
func (client *Client) ListUnbindDevicesWithOptions(request *ListUnbindDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListUnbindDevicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdDomain) {
		body["AdDomain"] = request.AdDomain
	}

	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.DirectoryId) {
		body["DirectoryId"] = request.DirectoryId
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	if !dara.IsNil(request.LastLoginUser) {
		body["LastLoginUser"] = request.LastLoginUser
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.UserType) {
		body["UserType"] = request.UserType
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUnbindDevices"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUnbindDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户未绑定的可信设备列表
//
// @param request - ListUnbindDevicesRequest
//
// @return ListUnbindDevicesResponse
func (client *Client) ListUnbindDevices(request *ListUnbindDevicesRequest) (_result *ListUnbindDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUnbindDevicesResponse{}
	_body, _err := client.ListUnbindDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询问题反馈列表
//
// @param request - ListUserFbAcIssuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserFbAcIssuesResponse
func (client *Client) ListUserFbAcIssuesWithOptions(request *ListUserFbAcIssuesRequest, runtime *dara.RuntimeOptions) (_result *ListUserFbAcIssuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		body["Account"] = request.Account
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.ErrorMessage) {
		body["ErrorMessage"] = request.ErrorMessage
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.IssueId) {
		body["IssueId"] = request.IssueId
	}

	if !dara.IsNil(request.Label) {
		body["Label"] = request.Label
	}

	if !dara.IsNil(request.ReservedA) {
		body["ReservedA"] = request.ReservedA
	}

	if !dara.IsNil(request.ReservedB) {
		body["ReservedB"] = request.ReservedB
	}

	if !dara.IsNil(request.UserEmail) {
		body["UserEmail"] = request.UserEmail
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserFbAcIssues"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserFbAcIssuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询问题反馈列表
//
// @param request - ListUserFbAcIssuesRequest
//
// @return ListUserFbAcIssuesResponse
func (client *Client) ListUserFbAcIssues(request *ListUserFbAcIssuesRequest) (_result *ListUserFbAcIssuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserFbAcIssuesResponse{}
	_body, _err := client.ListUserFbAcIssuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户反馈问题列表
//
// @param request - ListUserFbIssuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserFbIssuesResponse
func (client *Client) ListUserFbIssuesWithOptions(request *ListUserFbIssuesRequest, runtime *dara.RuntimeOptions) (_result *ListUserFbIssuesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientModel) {
		body["ClientModel"] = request.ClientModel
	}

	if !dara.IsNil(request.ClientSn) {
		body["ClientSn"] = request.ClientSn
	}

	if !dara.IsNil(request.CustomerId) {
		body["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DesktopId) {
		body["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.ErrorCode) {
		body["ErrorCode"] = request.ErrorCode
	}

	if !dara.IsNil(request.ErrorMsg) {
		body["ErrorMsg"] = request.ErrorMsg
	}

	if !dara.IsNil(request.FbType) {
		body["FbType"] = request.FbType
	}

	if !dara.IsNil(request.IssueId) {
		body["IssueId"] = request.IssueId
	}

	if !dara.IsNil(request.IssueLabel) {
		body["IssueLabel"] = request.IssueLabel
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UserEmail) {
		body["UserEmail"] = request.UserEmail
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.WasRead) {
		body["WasRead"] = request.WasRead
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserFbIssues"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserFbIssuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户反馈问题列表
//
// @param request - ListUserFbIssuesRequest
//
// @return ListUserFbIssuesResponse
func (client *Client) ListUserFbIssues(request *ListUserFbIssuesRequest) (_result *ListUserFbIssuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserFbIssuesResponse{}
	_body, _err := client.ListUserFbIssuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改设备安全入网类型
//
// @param request - ModifyDevicesSecureNetworkTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDevicesSecureNetworkTypeResponse
func (client *Client) ModifyDevicesSecureNetworkTypeWithOptions(request *ModifyDevicesSecureNetworkTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifyDevicesSecureNetworkTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllDevices) {
		body["AllDevices"] = request.AllDevices
	}

	if !dara.IsNil(request.SecureNetworkType) {
		body["SecureNetworkType"] = request.SecureNetworkType
	}

	if !dara.IsNil(request.SerialNos) {
		body["SerialNos"] = request.SerialNos
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDevicesSecureNetworkType"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDevicesSecureNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改设备安全入网类型
//
// @param request - ModifyDevicesSecureNetworkTypeRequest
//
// @return ModifyDevicesSecureNetworkTypeResponse
func (client *Client) ModifyDevicesSecureNetworkType(request *ModifyDevicesSecureNetworkTypeRequest) (_result *ModifyDevicesSecureNetworkTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDevicesSecureNetworkTypeResponse{}
	_body, _err := client.ModifyDevicesSecureNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 匿名api修改安全入网配置
//
// @param request - ModifySecureNetworkTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecureNetworkTypeResponse
func (client *Client) ModifySecureNetworkTypeWithOptions(request *ModifySecureNetworkTypeRequest, runtime *dara.RuntimeOptions) (_result *ModifySecureNetworkTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SecureNetworkType) {
		body["SecureNetworkType"] = request.SecureNetworkType
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifySecureNetworkType"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifySecureNetworkTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 匿名api修改安全入网配置
//
// @param request - ModifySecureNetworkTypeRequest
//
// @return ModifySecureNetworkTypeResponse
func (client *Client) ModifySecureNetworkType(request *ModifySecureNetworkTypeRequest) (_result *ModifySecureNetworkTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifySecureNetworkTypeResponse{}
	_body, _err := client.ModifySecureNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - RegisterDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDeviceResponse
func (client *Client) RegisterDeviceWithOptions(request *RegisterDeviceRequest, runtime *dara.RuntimeOptions) (_result *RegisterDeviceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bluetooth) {
		body["Bluetooth"] = request.Bluetooth
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.ChipId) {
		body["ChipId"] = request.ChipId
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.Cpu) {
		body["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.CustomId) {
		body["CustomId"] = request.CustomId
	}

	if !dara.IsNil(request.EtherMac) {
		body["EtherMac"] = request.EtherMac
	}

	if !dara.IsNil(request.Memory) {
		body["Memory"] = request.Memory
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.Storage) {
		body["Storage"] = request.Storage
	}

	if !dara.IsNil(request.Token) {
		body["Token"] = request.Token
	}

	if !dara.IsNil(request.Wlan) {
		body["Wlan"] = request.Wlan
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterDevice"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - RegisterDeviceRequest
//
// @return RegisterDeviceResponse
func (client *Client) RegisterDevice(request *RegisterDeviceRequest) (_result *RegisterDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterDeviceResponse{}
	_body, _err := client.RegisterDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报升级信息
//
// @param request - ReportAppOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportAppOtaInfoResponse
func (client *Client) ReportAppOtaInfoWithOptions(request *ReportAppOtaInfoRequest, runtime *dara.RuntimeOptions) (_result *ReportAppOtaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseVersion) {
		query["BaseVersion"] = request.BaseVersion
	}

	if !dara.IsNil(request.ClientType) {
		query["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientUid) {
		query["ClientUid"] = request.ClientUid
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.OsType) {
		query["OsType"] = request.OsType
	}

	if !dara.IsNil(request.Project) {
		query["Project"] = request.Project
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetVersion) {
		query["TargetVersion"] = request.TargetVersion
	}

	if !dara.IsNil(request.TaskUid) {
		query["TaskUid"] = request.TaskUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportAppOtaInfo"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportAppOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报升级信息
//
// @param request - ReportAppOtaInfoRequest
//
// @return ReportAppOtaInfoResponse
func (client *Client) ReportAppOtaInfo(request *ReportAppOtaInfoRequest) (_result *ReportAppOtaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportAppOtaInfoResponse{}
	_body, _err := client.ReportAppOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级信息上报
//
// @param request - ReportDeviceOtaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportDeviceOtaInfoResponse
func (client *Client) ReportDeviceOtaInfoWithOptions(request *ReportDeviceOtaInfoRequest, runtime *dara.RuntimeOptions) (_result *ReportDeviceOtaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaseVersion) {
		body["BaseVersion"] = request.BaseVersion
	}

	if !dara.IsNil(request.DeviceId) {
		body["DeviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.Note) {
		body["Note"] = request.Note
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TargetVersion) {
		body["TargetVersion"] = request.TargetVersion
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportDeviceOtaInfo"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportDeviceOtaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级信息上报
//
// @param request - ReportDeviceOtaInfoRequest
//
// @return ReportDeviceOtaInfoResponse
func (client *Client) ReportDeviceOtaInfo(request *ReportDeviceOtaInfoRequest) (_result *ReportDeviceOtaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportDeviceOtaInfoResponse{}
	_body, _err := client.ReportDeviceOtaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用中心用户问题反馈
//
// @param tmpReq - ReportUserFbAcIssueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportUserFbAcIssueResponse
func (client *Client) ReportUserFbAcIssueWithOptions(tmpReq *ReportUserFbAcIssueRequest, runtime *dara.RuntimeOptions) (_result *ReportUserFbAcIssueResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ReportUserFbAcIssueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileList) {
		request.FileListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileList, dara.String("FileList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		body["Account"] = request.Account
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.ErrorMsg) {
		body["ErrorMsg"] = request.ErrorMsg
	}

	if !dara.IsNil(request.FileListShrink) {
		body["FileList"] = request.FileListShrink
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Labels) {
		body["Labels"] = request.Labels
	}

	if !dara.IsNil(request.ReservedA) {
		body["ReservedA"] = request.ReservedA
	}

	if !dara.IsNil(request.ReservedB) {
		body["ReservedB"] = request.ReservedB
	}

	if !dara.IsNil(request.UserEmail) {
		body["UserEmail"] = request.UserEmail
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportUserFbAcIssue"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportUserFbAcIssueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用中心用户问题反馈
//
// @param request - ReportUserFbAcIssueRequest
//
// @return ReportUserFbAcIssueResponse
func (client *Client) ReportUserFbAcIssue(request *ReportUserFbAcIssueRequest) (_result *ReportUserFbAcIssueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportUserFbAcIssueResponse{}
	_body, _err := client.ReportUserFbAcIssueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上报用户反馈问题
//
// @param tmpReq - ReportUserFbIssueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportUserFbIssueResponse
func (client *Client) ReportUserFbIssueWithOptions(tmpReq *ReportUserFbIssueRequest, runtime *dara.RuntimeOptions) (_result *ReportUserFbIssueResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ReportUserFbIssueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileList) {
		request.FileListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileList, dara.String("FileList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.ClientAppVersion) {
		body["ClientAppVersion"] = request.ClientAppVersion
	}

	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientModel) {
		body["ClientModel"] = request.ClientModel
	}

	if !dara.IsNil(request.ClientOsName) {
		body["ClientOsName"] = request.ClientOsName
	}

	if !dara.IsNil(request.ClientSn) {
		body["ClientSn"] = request.ClientSn
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.CustomerId) {
		body["CustomerId"] = request.CustomerId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DesktopId) {
		body["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopType) {
		body["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.ErrorCode) {
		body["ErrorCode"] = request.ErrorCode
	}

	if !dara.IsNil(request.ErrorMsg) {
		body["ErrorMsg"] = request.ErrorMsg
	}

	if !dara.IsNil(request.FbType) {
		body["FbType"] = request.FbType
	}

	if !dara.IsNil(request.FileListShrink) {
		body["FileList"] = request.FileListShrink
	}

	if !dara.IsNil(request.IsSubstituteReport) {
		body["IsSubstituteReport"] = request.IsSubstituteReport
	}

	if !dara.IsNil(request.IssueLabel) {
		body["IssueLabel"] = request.IssueLabel
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OccurTime) {
		body["OccurTime"] = request.OccurTime
	}

	if !dara.IsNil(request.ReservedA) {
		body["ReservedA"] = request.ReservedA
	}

	if !dara.IsNil(request.ReservedB) {
		body["ReservedB"] = request.ReservedB
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TelNo) {
		body["TelNo"] = request.TelNo
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	if !dara.IsNil(request.UserEmail) {
		body["UserEmail"] = request.UserEmail
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	if !dara.IsNil(request.UserName) {
		body["UserName"] = request.UserName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WyId) {
		body["WyId"] = request.WyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportUserFbIssue"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReportUserFbIssueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上报用户反馈问题
//
// @param request - ReportUserFbIssueRequest
//
// @return ReportUserFbIssueResponse
func (client *Client) ReportUserFbIssue(request *ReportUserFbIssueRequest) (_result *ReportUserFbIssueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReportUserFbIssueResponse{}
	_body, _err := client.ReportUserFbIssueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向终端发送运维命令
//
// @param request - SendOpsMessageToTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminalsWithOptions(request *SendOpsMessageToTerminalsRequest, runtime *dara.RuntimeOptions) (_result *SendOpsMessageToTerminalsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Delay) {
		query["Delay"] = request.Delay
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Msg) {
		body["Msg"] = request.Msg
	}

	if !dara.IsNil(request.OpsAction) {
		body["OpsAction"] = request.OpsAction
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Uuids) {
		bodyFlat["Uuids"] = request.Uuids
	}

	if !dara.IsNil(request.WaitForAck) {
		body["WaitForAck"] = request.WaitForAck
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendOpsMessageToTerminals"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendOpsMessageToTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向终端发送运维命令
//
// @param request - SendOpsMessageToTerminalsRequest
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminals(request *SendOpsMessageToTerminalsRequest) (_result *SendOpsMessageToTerminalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendOpsMessageToTerminalsResponse{}
	_body, _err := client.SendOpsMessageToTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置租户ota自动开启/关闭
//
// @param request - SetDeviceOtaAutoStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeviceOtaAutoStatusResponse
func (client *Client) SetDeviceOtaAutoStatusWithOptions(request *SetDeviceOtaAutoStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDeviceOtaAutoStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoUpdate) {
		body["AutoUpdate"] = request.AutoUpdate
	}

	if !dara.IsNil(request.AutoUpdateTimeSchedule) {
		body["AutoUpdateTimeSchedule"] = request.AutoUpdateTimeSchedule
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ForceUpgrade) {
		body["ForceUpgrade"] = request.ForceUpgrade
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDeviceOtaAutoStatus"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDeviceOtaAutoStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置租户ota自动开启/关闭
//
// @param request - SetDeviceOtaAutoStatusRequest
//
// @return SetDeviceOtaAutoStatusResponse
func (client *Client) SetDeviceOtaAutoStatus(request *SetDeviceOtaAutoStatusRequest) (_result *SetDeviceOtaAutoStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDeviceOtaAutoStatusResponse{}
	_body, _err := client.SetDeviceOtaAutoStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 租户设置设备ota任务的状态
//
// @param request - SetDeviceOtaTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeviceOtaTaskStatusResponse
func (client *Client) SetDeviceOtaTaskStatusWithOptions(request *SetDeviceOtaTaskStatusRequest, runtime *dara.RuntimeOptions) (_result *SetDeviceOtaTaskStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OperationStatus) {
		body["OperationStatus"] = request.OperationStatus
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDeviceOtaTaskStatus"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDeviceOtaTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 租户设置设备ota任务的状态
//
// @param request - SetDeviceOtaTaskStatusRequest
//
// @return SetDeviceOtaTaskStatusResponse
func (client *Client) SetDeviceOtaTaskStatus(request *SetDeviceOtaTaskStatusRequest) (_result *SetDeviceOtaTaskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDeviceOtaTaskStatusResponse{}
	_body, _err := client.SetDeviceOtaTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑免账号登录用户
//
// @param request - UnbindAccountLessLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUserWithOptions(request *UnbindAccountLessLoginUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindAccountLessLoginUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAccountLessLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAccountLessLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑免账号登录用户
//
// @param request - UnbindAccountLessLoginUserRequest
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUser(request *UnbindAccountLessLoginUserRequest) (_result *UnbindAccountLessLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindAccountLessLoginUserResponse{}
	_body, _err := client.UnbindAccountLessLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑设备座位
//
// @param tmpReq - UnbindDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDeviceSeatsResponse
func (client *Client) UnbindDeviceSeatsWithOptions(tmpReq *UnbindDeviceSeatsRequest, runtime *dara.RuntimeOptions) (_result *UnbindDeviceSeatsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UnbindDeviceSeatsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SerialNoList) {
		request.SerialNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SerialNoList, dara.String("SerialNoList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SerialNoListShrink) {
		body["SerialNoList"] = request.SerialNoListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDeviceSeats"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDeviceSeatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑设备座位
//
// @param request - UnbindDeviceSeatsRequest
//
// @return UnbindDeviceSeatsResponse
func (client *Client) UnbindDeviceSeats(request *UnbindDeviceSeatsRequest) (_result *UnbindDeviceSeatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindDeviceSeatsResponse{}
	_body, _err := client.UnbindDeviceSeatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑免密登录用户
//
// @param request - UnbindPasswordFreeLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPasswordFreeLoginUserResponse
func (client *Client) UnbindPasswordFreeLoginUserWithOptions(request *UnbindPasswordFreeLoginUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindPasswordFreeLoginUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPasswordFreeLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPasswordFreeLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑免密登录用户
//
// @param request - UnbindPasswordFreeLoginUserRequest
//
// @return UnbindPasswordFreeLoginUserResponse
func (client *Client) UnbindPasswordFreeLoginUser(request *UnbindPasswordFreeLoginUserRequest) (_result *UnbindPasswordFreeLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindPasswordFreeLoginUserResponse{}
	_body, _err := client.UnbindPasswordFreeLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新设备别名
//
// @param request - UpdateAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAliasResponse
func (client *Client) UpdateAliasWithOptions(request *UpdateAliasRequest, runtime *dara.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAlias"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新设备别名
//
// @param request - UpdateAliasRequest
//
// @return UpdateAliasResponse
func (client *Client) UpdateAlias(request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新设备绑定的终端用户
//
// @param request - UpdateDeviceBindedEndUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeviceBindedEndUserResponse
func (client *Client) UpdateDeviceBindedEndUserWithOptions(request *UpdateDeviceBindedEndUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeviceBindedEndUserResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.SourceAdEndUsers) {
		body["SourceAdEndUsers"] = request.SourceAdEndUsers
	}

	if !dara.IsNil(request.SourceEndUserIds) {
		body["SourceEndUserIds"] = request.SourceEndUserIds
	}

	if !dara.IsNil(request.TargetAdEndUsers) {
		body["TargetAdEndUsers"] = request.TargetAdEndUsers
	}

	if !dara.IsNil(request.TargetEndUserIds) {
		body["TargetEndUserIds"] = request.TargetEndUserIds
	}

	if !dara.IsNil(request.UserType) {
		body["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeviceBindedEndUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeviceBindedEndUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新设备绑定的终端用户
//
// @param request - UpdateDeviceBindedEndUserRequest
//
// @return UpdateDeviceBindedEndUserResponse
func (client *Client) UpdateDeviceBindedEndUser(request *UpdateDeviceBindedEndUserRequest) (_result *UpdateDeviceBindedEndUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDeviceBindedEndUserResponse{}
	_body, _err := client.UpdateDeviceBindedEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改标签
//
// @param request - UpdateLabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLabelResponse
func (client *Client) UpdateLabelWithOptions(request *UpdateLabelRequest, runtime *dara.RuntimeOptions) (_result *UpdateLabelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LabelContent) {
		body["LabelContent"] = request.LabelContent
	}

	if !dara.IsNil(request.LabelId) {
		body["LabelId"] = request.LabelId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLabel"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改标签
//
// @param request - UpdateLabelRequest
//
// @return UpdateLabelResponse
func (client *Client) UpdateLabel(request *UpdateLabelRequest) (_result *UpdateLabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLabelResponse{}
	_body, _err := client.UpdateLabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改终端策略
//
// @param request - UpdateTerminalPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTerminalPolicyResponse
func (client *Client) UpdateTerminalPolicyWithOptions(request *UpdateTerminalPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateTerminalPolicyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowManualLockScreen) {
		body["AllowManualLockScreen"] = request.AllowManualLockScreen
	}

	if !dara.IsNil(request.BackgroundModeTitle) {
		body["BackgroundModeTitle"] = request.BackgroundModeTitle
	}

	if !dara.IsNil(request.CustomScreenCastRes) {
		body["CustomScreenCastRes"] = request.CustomScreenCastRes
	}

	if !dara.IsNil(request.DisplayLayout) {
		body["DisplayLayout"] = request.DisplayLayout
	}

	if !dara.IsNil(request.DisplayResolution) {
		body["DisplayResolution"] = request.DisplayResolution
	}

	if !dara.IsNil(request.DisplayScaleRatio) {
		body["DisplayScaleRatio"] = request.DisplayScaleRatio
	}

	if !dara.IsNil(request.EnableAutoLockScreen) {
		body["EnableAutoLockScreen"] = request.EnableAutoLockScreen
	}

	if !dara.IsNil(request.EnableAutoLogin) {
		body["EnableAutoLogin"] = request.EnableAutoLogin
	}

	if !dara.IsNil(request.EnableBackgroundMode) {
		body["EnableBackgroundMode"] = request.EnableBackgroundMode
	}

	if !dara.IsNil(request.EnableBluetooth) {
		body["EnableBluetooth"] = request.EnableBluetooth
	}

	if !dara.IsNil(request.EnableControlPanel) {
		body["EnableControlPanel"] = request.EnableControlPanel
	}

	if !dara.IsNil(request.EnableImmersiveMode) {
		body["EnableImmersiveMode"] = request.EnableImmersiveMode
	}

	if !dara.IsNil(request.EnableLockScreenHotKey) {
		body["EnableLockScreenHotKey"] = request.EnableLockScreenHotKey
	}

	if !dara.IsNil(request.EnableModifyPassword) {
		body["EnableModifyPassword"] = request.EnableModifyPassword
	}

	if !dara.IsNil(request.EnableScanLogin) {
		body["EnableScanLogin"] = request.EnableScanLogin
	}

	if !dara.IsNil(request.EnableScheduledReboot) {
		body["EnableScheduledReboot"] = request.EnableScheduledReboot
	}

	if !dara.IsNil(request.EnableScheduledShutdown) {
		body["EnableScheduledShutdown"] = request.EnableScheduledShutdown
	}

	if !dara.IsNil(request.EnableSmsLogin) {
		body["EnableSmsLogin"] = request.EnableSmsLogin
	}

	if !dara.IsNil(request.EnableSwitchPersonal) {
		body["EnableSwitchPersonal"] = request.EnableSwitchPersonal
	}

	if !dara.IsNil(request.EnableWlan) {
		body["EnableWlan"] = request.EnableWlan
	}

	if !dara.IsNil(request.FollowCloudReboot) {
		body["FollowCloudReboot"] = request.FollowCloudReboot
	}

	if !dara.IsNil(request.FollowCloudShutdown) {
		body["FollowCloudShutdown"] = request.FollowCloudShutdown
	}

	if !dara.IsNil(request.FollowTerminalReboot) {
		body["FollowTerminalReboot"] = request.FollowTerminalReboot
	}

	if !dara.IsNil(request.FollowTerminalShutdown) {
		body["FollowTerminalShutdown"] = request.FollowTerminalShutdown
	}

	if !dara.IsNil(request.ForceSetPinCode) {
		body["ForceSetPinCode"] = request.ForceSetPinCode
	}

	if !dara.IsNil(request.IdleTimeout) {
		body["IdleTimeout"] = request.IdleTimeout
	}

	if !dara.IsNil(request.IdleTimeoutAction) {
		body["IdleTimeoutAction"] = request.IdleTimeoutAction
	}

	if !dara.IsNil(request.LockScreenPasswordRequired) {
		body["LockScreenPasswordRequired"] = request.LockScreenPasswordRequired
	}

	if !dara.IsNil(request.LockScreenTimeout) {
		body["LockScreenTimeout"] = request.LockScreenTimeout
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PowerButtonDefine) {
		body["PowerButtonDefine"] = request.PowerButtonDefine
	}

	if !dara.IsNil(request.PowerButtonDefineForAs) {
		body["PowerButtonDefineForAs"] = request.PowerButtonDefineForAs
	}

	if !dara.IsNil(request.PowerButtonDefineForNs) {
		body["PowerButtonDefineForNs"] = request.PowerButtonDefineForNs
	}

	if !dara.IsNil(request.PowerOnBehavior) {
		body["PowerOnBehavior"] = request.PowerOnBehavior
	}

	if !dara.IsNil(request.RunningMode) {
		body["RunningMode"] = request.RunningMode
	}

	if !dara.IsNil(request.ScheduledReboot) {
		body["ScheduledReboot"] = request.ScheduledReboot
	}

	if !dara.IsNil(request.ScheduledShutdown) {
		body["ScheduledShutdown"] = request.ScheduledShutdown
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ScreenCastResPaths) {
		bodyFlat["ScreenCastResPaths"] = request.ScreenCastResPaths
	}

	if !dara.IsNil(request.SettingLock) {
		body["SettingLock"] = request.SettingLock
	}

	if !dara.IsNil(request.TerminalPolicyId) {
		body["TerminalPolicyId"] = request.TerminalPolicyId
	}

	if !dara.IsNil(request.UnlockMethod) {
		body["UnlockMethod"] = request.UnlockMethod
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTerminalPolicy"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTerminalPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改终端策略
//
// @param request - UpdateTerminalPolicyRequest
//
// @return UpdateTerminalPolicyResponse
func (client *Client) UpdateTerminalPolicy(request *UpdateTerminalPolicyRequest) (_result *UpdateTerminalPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTerminalPolicyResponse{}
	_body, _err := client.UpdateTerminalPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
