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
	EnableValidate  *bool
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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": dara.String("csas.aliyuncs.com"),
		"public":      dara.String("csas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("csas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Appends associated terminal devices to a static device label in batches.
//
// @param request - AddDeviceGroupMatchDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDeviceGroupMatchDevicesResponse
func (client *Client) AddDeviceGroupMatchDevicesWithOptions(request *AddDeviceGroupMatchDevicesRequest, runtime *dara.RuntimeOptions) (_result *AddDeviceGroupMatchDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DevTags) {
		bodyFlat["DevTags"] = request.DevTags
	}

	if !dara.IsNil(request.DeviceGroupId) {
		body["DeviceGroupId"] = request.DeviceGroupId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDeviceGroupMatchDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDeviceGroupMatchDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Appends associated terminal devices to a static device label in batches.
//
// @param request - AddDeviceGroupMatchDevicesRequest
//
// @return AddDeviceGroupMatchDevicesResponse
func (client *Client) AddDeviceGroupMatchDevices(request *AddDeviceGroupMatchDevicesRequest) (_result *AddDeviceGroupMatchDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDeviceGroupMatchDevicesResponse{}
	_body, _err := client.AddDeviceGroupMatchDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Appends entries in batches to the virus scan blacklists and whitelists for a specified operating system without overwriting existing entries. Quotas are calculated independently for each combination of matching dimension and list type. Each combination allows a maximum of 10,000 whitelist entries and 1,000 blacklist entries. If the quota is exceeded after appending, the entire batch fails.
//
// @param request - AddVirusScanAdditionalListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVirusScanAdditionalListsResponse
func (client *Client) AddVirusScanAdditionalListsWithOptions(request *AddVirusScanAdditionalListsRequest, runtime *dara.RuntimeOptions) (_result *AddVirusScanAdditionalListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalLists) {
		bodyFlat["AdditionalLists"] = request.AdditionalLists
	}

	if !dara.IsNil(request.DevType) {
		body["DevType"] = request.DevType
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVirusScanAdditionalLists"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVirusScanAdditionalListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Appends entries in batches to the virus scan blacklists and whitelists for a specified operating system without overwriting existing entries. Quotas are calculated independently for each combination of matching dimension and list type. Each combination allows a maximum of 10,000 whitelist entries and 1,000 blacklist entries. If the quota is exceeded after appending, the entire batch fails.
//
// @param request - AddVirusScanAdditionalListsRequest
//
// @return AddVirusScanAdditionalListsResponse
func (client *Client) AddVirusScanAdditionalLists(request *AddVirusScanAdditionalListsRequest) (_result *AddVirusScanAdditionalListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddVirusScanAdditionalListsResponse{}
	_body, _err := client.AddVirusScanAdditionalListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches the private access applications of a Connector under the current Alibaba Cloud account.
//
// @param tmpReq - AttachApplication2ConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachApplication2ConnectorResponse
func (client *Client) AttachApplication2ConnectorWithOptions(tmpReq *AttachApplication2ConnectorRequest, runtime *dara.RuntimeOptions) (_result *AttachApplication2ConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AttachApplication2ConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationIds) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, dara.String("ApplicationIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIdsShrink) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachApplication2Connector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachApplication2ConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches the private access applications of a Connector under the current Alibaba Cloud account.
//
// @param request - AttachApplication2ConnectorRequest
//
// @return AttachApplication2ConnectorResponse
func (client *Client) AttachApplication2Connector(request *AttachApplication2ConnectorRequest) (_result *AttachApplication2ConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachApplication2ConnectorResponse{}
	_body, _err := client.AttachApplication2ConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a business policy to a specified approval process.
//
// @param request - AttachPolicy2ApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachPolicy2ApprovalProcessResponse
func (client *Client) AttachPolicy2ApprovalProcessWithOptions(request *AttachPolicy2ApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *AttachPolicy2ApprovalProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.ProcessId) {
		body["ProcessId"] = request.ProcessId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AttachPolicy2ApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AttachPolicy2ApprovalProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a business policy to a specified approval process.
//
// @param request - AttachPolicy2ApprovalProcessRequest
//
// @return AttachPolicy2ApprovalProcessResponse
func (client *Client) AttachPolicy2ApprovalProcess(request *AttachPolicy2ApprovalProcessRequest) (_result *AttachPolicy2ApprovalProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AttachPolicy2ApprovalProcessResponse{}
	_body, _err := client.AttachPolicy2ApprovalProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Inserts domain name entries into a domain name list in batches.
//
// Description:
//
// Appends domain name entries in batches to a specified domain name list (`ListId`). Domain names must be second-level or higher domain names. Wildcard domain names (`*.example.com`) are supported, but overly broad patterns such as `*.com` or `*.com.cn` are prohibited.
//
// @param request - BatchCreateDomainItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateDomainItemsResponse
func (client *Client) BatchCreateDomainItemsWithOptions(request *BatchCreateDomainItemsRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateDomainItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DomainItems) {
		bodyFlat["DomainItems"] = request.DomainItems
	}

	if !dara.IsNil(request.ListId) {
		body["ListId"] = request.ListId
	}

	if !dara.IsNil(request.ListType) {
		body["ListType"] = request.ListType
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateDomainItems"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateDomainItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Inserts domain name entries into a domain name list in batches.
//
// Description:
//
// Appends domain name entries in batches to a specified domain name list (`ListId`). Domain names must be second-level or higher domain names. Wildcard domain names (`*.example.com`) are supported, but overly broad patterns such as `*.com` or `*.com.cn` are prohibited.
//
// @param request - BatchCreateDomainItemsRequest
//
// @return BatchCreateDomainItemsResponse
func (client *Client) BatchCreateDomainItems(request *BatchCreateDomainItemsRequest) (_result *BatchCreateDomainItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateDomainItemsResponse{}
	_body, _err := client.BatchCreateDomainItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch deletes domain name entries from a domain name list.
//
// Description:
//
// Batch deletes domain name entries from a specified domain name list by entry IDs (`ItemIds`, obtained from the `ItemId` field returned by ListDomainItems).
//
// @param request - BatchDeleteDomainItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDomainItemsResponse
func (client *Client) BatchDeleteDomainItemsWithOptions(request *BatchDeleteDomainItemsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDomainItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ItemIds) {
		bodyFlat["ItemIds"] = request.ItemIds
	}

	if !dara.IsNil(request.ListId) {
		body["ListId"] = request.ListId
	}

	if !dara.IsNil(request.ListType) {
		body["ListType"] = request.ListType
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDomainItems"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDomainItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch deletes domain name entries from a domain name list.
//
// Description:
//
// Batch deletes domain name entries from a specified domain name list by entry IDs (`ItemIds`, obtained from the `ItemId` field returned by ListDomainItems).
//
// @param request - BatchDeleteDomainItemsRequest
//
// @return BatchDeleteDomainItemsResponse
func (client *Client) BatchDeleteDomainItems(request *BatchDeleteDomainItemsRequest) (_result *BatchDeleteDomainItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDomainItemsResponse{}
	_body, _err := client.BatchDeleteDomainItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes internal-facing applications in batches.
//
// Description:
//
// Applications that are referenced by office network recognition or policies cannot be deleted. References:
//
// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists internal-facing access applications in batches.
//
// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): Lists internal-facing access policies in batches.
//
// @param request - BatchDeletePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeletePrivateAccessApplicationResponse
func (client *Client) BatchDeletePrivateAccessApplicationWithOptions(request *BatchDeletePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *BatchDeletePrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeletePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeletePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes internal-facing applications in batches.
//
// Description:
//
// Applications that are referenced by office network recognition or policies cannot be deleted. References:
//
// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists internal-facing access applications in batches.
//
// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): Lists internal-facing access policies in batches.
//
// @param request - BatchDeletePrivateAccessApplicationRequest
//
// @return BatchDeletePrivateAccessApplicationResponse
func (client *Client) BatchDeletePrivateAccessApplication(request *BatchDeletePrivateAccessApplicationRequest) (_result *BatchDeletePrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeletePrivateAccessApplicationResponse{}
	_body, _err := client.BatchDeletePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes internal network access policies in batches.
//
// @param request - BatchDeletePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeletePrivateAccessPolicyResponse
func (client *Client) BatchDeletePrivateAccessPolicyWithOptions(request *BatchDeletePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *BatchDeletePrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIds) {
		bodyFlat["PolicyIds"] = request.PolicyIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeletePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeletePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes internal network access policies in batches.
//
// @param request - BatchDeletePrivateAccessPolicyRequest
//
// @return BatchDeletePrivateAccessPolicyResponse
func (client *Client) BatchDeletePrivateAccessPolicy(request *BatchDeletePrivateAccessPolicyRequest) (_result *BatchDeletePrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeletePrivateAccessPolicyResponse{}
	_body, _err := client.BatchDeletePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels multiple virus scan tasks that have not yet expired in a batch. After cancellation, terminals no longer pull and execute the tasks. Scans already running on terminals are not interrupted.
//
// @param request - CancelVirusScanTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelVirusScanTasksResponse
func (client *Client) CancelVirusScanTasksWithOptions(request *CancelVirusScanTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelVirusScanTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		bodyFlat["TaskIds"] = request.TaskIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelVirusScanTasks"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelVirusScanTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels multiple virus scan tasks that have not yet expired in a batch. After cancellation, terminals no longer pull and execute the tasks. Scans already running on terminals are not interrupted.
//
// @param request - CancelVirusScanTasksRequest
//
// @return CancelVirusScanTasksResponse
func (client *Client) CancelVirusScanTasks(request *CancelVirusScanTasksRequest) (_result *CancelVirusScanTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelVirusScanTasksResponse{}
	_body, _err := client.CancelVirusScanTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels multiple vulnerability scanning tasks that have not yet expired in a batch.
//
// @param request - CancelVulScanTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelVulScanTasksResponse
func (client *Client) CancelVulScanTasksWithOptions(request *CancelVulScanTasksRequest, runtime *dara.RuntimeOptions) (_result *CancelVulScanTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		bodyFlat["TaskIds"] = request.TaskIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelVulScanTasks"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelVulScanTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels multiple vulnerability scanning tasks that have not yet expired in a batch.
//
// @param request - CancelVulScanTasksRequest
//
// @return CancelVulScanTasksResponse
func (client *Client) CancelVulScanTasks(request *CancelVulScanTasksRequest) (_result *CancelVulScanTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelVulScanTasksResponse{}
	_body, _err := client.CancelVulScanTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an approval flow under the current Alibaba Cloud account.
//
// @param tmpReq - CreateApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApprovalProcessResponse
func (client *Client) CreateApprovalProcessWithOptions(tmpReq *CreateApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *CreateApprovalProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateApprovalProcessShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MatchSchemas) {
		request.MatchSchemasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchSchemas, dara.String("MatchSchemas"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchSchemasShrink) {
		body["MatchSchemas"] = request.MatchSchemasShrink
	}

	if !dara.IsNil(request.ProcessName) {
		body["ProcessName"] = request.ProcessName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ProcessNodes) {
		bodyFlat["ProcessNodes"] = request.ProcessNodes
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApprovalProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an approval flow under the current Alibaba Cloud account.
//
// @param request - CreateApprovalProcessRequest
//
// @return CreateApprovalProcessResponse
func (client *Client) CreateApprovalProcess(request *CreateApprovalProcessRequest) (_result *CreateApprovalProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateApprovalProcessResponse{}
	_body, _err := client.CreateApprovalProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a custom identity source user for your Alibaba Cloud account.
//
// @param request - CreateClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientUserResponse
func (client *Client) CreateClientUserWithOptions(request *CreateClientUserRequest, runtime *dara.RuntimeOptions) (_result *CreateClientUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	if !dara.IsNil(request.MobileNumber) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a custom identity source user for your Alibaba Cloud account.
//
// @param request - CreateClientUserRequest
//
// @return CreateClientUserResponse
func (client *Client) CreateClientUser(request *CreateClientUserRequest) (_result *CreateClientUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateClientUserResponse{}
	_body, _err := client.CreateClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a connector.
//
// @param request - CreateConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnectorResponse
func (client *Client) CreateConnectorWithOptions(request *CreateConnectorRequest, runtime *dara.RuntimeOptions) (_result *CreateConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bandwidth) {
		body["Bandwidth"] = request.Bandwidth
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Region) {
		body["Region"] = request.Region
	}

	if !dara.IsNil(request.SwitchStatus) {
		body["SwitchStatus"] = request.SwitchStatus
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConnector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a connector.
//
// @param request - CreateConnectorRequest
//
// @return CreateConnectorResponse
func (client *Client) CreateConnector(request *CreateConnectorRequest) (_result *CreateConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConnectorResponse{}
	_body, _err := client.CreateConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a device label.
//
// @param tmpReq - CreateDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDeviceGroupResponse
func (client *Client) CreateDeviceGroupWithOptions(tmpReq *CreateDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDeviceGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DynamicRule) {
		request.DynamicRuleShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DynamicRule, dara.String("DynamicRule"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicOperator) {
		body["DynamicOperator"] = request.DynamicOperator
	}

	if !dara.IsNil(request.DynamicRuleShrink) {
		body["DynamicRule"] = request.DynamicRuleShrink
	}

	if !dara.IsNil(request.GroupType) {
		body["GroupType"] = request.GroupType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDeviceGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a device label.
//
// @param request - CreateDeviceGroupRequest
//
// @return CreateDeviceGroupResponse
func (client *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (_result *CreateDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDeviceGroupResponse{}
	_body, _err := client.CreateDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a domain name list.
//
// Description:
//
// Creates a domain name list of a specified type (blacklist or whitelist) under the current tenant and returns the ListId of the new list. A maximum of 100 lists can be created for each list type per tenant.
//
// @param request - CreateDomainMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainMetaResponse
func (client *Client) CreateDomainMetaWithOptions(request *CreateDomainMetaRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ListType) {
		body["ListType"] = request.ListType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomainMeta"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a domain name list.
//
// Description:
//
// Creates a domain name list of a specified type (blacklist or whitelist) under the current tenant and returns the ListId of the new list. A maximum of 100 lists can be created for each list type per tenant.
//
// @param request - CreateDomainMetaRequest
//
// @return CreateDomainMetaResponse
func (client *Client) CreateDomainMeta(request *CreateDomainMetaRequest) (_result *CreateDomainMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainMetaResponse{}
	_body, _err := client.CreateDomainMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create dynamic routes for the current Alibaba Cloud account.
//
// Description:
//
// By default, you can create a maximum of 100 dynamic routes.
//
// @param request - CreateDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDynamicRouteResponse
func (client *Client) CreateDynamicRouteWithOptions(request *CreateDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *CreateDynamicRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicRouteType) {
		body["DynamicRouteType"] = request.DynamicRouteType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextHop) {
		body["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionIds) {
		bodyFlat["RegionIds"] = request.RegionIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create dynamic routes for the current Alibaba Cloud account.
//
// Description:
//
// By default, you can create a maximum of 100 dynamic routes.
//
// @param request - CreateDynamicRouteRequest
//
// @return CreateDynamicRouteResponse
func (client *Client) CreateDynamicRoute(request *CreateDynamicRouteRequest) (_result *CreateDynamicRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDynamicRouteResponse{}
	_body, _err := client.CreateDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an enterprise acceleration policy.
//
// @param request - CreateEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseAcceleratePolicyResponse
func (client *Client) CreateEnterpriseAcceleratePolicyWithOptions(request *CreateEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseAcceleratePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationType) {
		body["AccelerationType"] = request.AccelerationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ShowInClient) {
		body["ShowInClient"] = request.ShowInClient
	}

	if !dara.IsNil(request.UpstreamHost) {
		body["UpstreamHost"] = request.UpstreamHost
	}

	if !dara.IsNil(request.UpstreamPort) {
		body["UpstreamPort"] = request.UpstreamPort
	}

	if !dara.IsNil(request.UpstreamType) {
		body["UpstreamType"] = request.UpstreamType
	}

	if !dara.IsNil(request.UserAttributeGroup) {
		body["UserAttributeGroup"] = request.UserAttributeGroup
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an enterprise acceleration policy.
//
// @param request - CreateEnterpriseAcceleratePolicyRequest
//
// @return CreateEnterpriseAcceleratePolicyResponse
func (client *Client) CreateEnterpriseAcceleratePolicy(request *CreateEnterpriseAcceleratePolicyRequest) (_result *CreateEnterpriseAcceleratePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CreateEnterpriseAcceleratePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an enterprise acceleration address.
//
// @param request - CreateEnterpriseAccelerateTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEnterpriseAccelerateTargetResponse
func (client *Client) CreateEnterpriseAccelerateTargetWithOptions(request *CreateEnterpriseAccelerateTargetRequest, runtime *dara.RuntimeOptions) (_result *CreateEnterpriseAccelerateTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Target) {
		bodyFlat["Target"] = request.Target
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEnterpriseAccelerateTarget"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEnterpriseAccelerateTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an enterprise acceleration address.
//
// @param request - CreateEnterpriseAccelerateTargetRequest
//
// @return CreateEnterpriseAccelerateTargetResponse
func (client *Client) CreateEnterpriseAccelerateTarget(request *CreateEnterpriseAccelerateTargetRequest) (_result *CreateEnterpriseAccelerateTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEnterpriseAccelerateTargetResponse{}
	_body, _err := client.CreateEnterpriseAccelerateTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a traffic forwarding rule.
//
// @param request - CreateForwardStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateForwardStrategyResponse
func (client *Client) CreateForwardStrategyWithOptions(request *CreateForwardStrategyRequest, runtime *dara.RuntimeOptions) (_result *CreateForwardStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationId) {
		body["DestinationId"] = request.DestinationId
	}

	if !dara.IsNil(request.DestinationType) {
		body["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateForwardStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateForwardStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a traffic forwarding rule.
//
// @param request - CreateForwardStrategyRequest
//
// @return CreateForwardStrategyResponse
func (client *Client) CreateForwardStrategy(request *CreateForwardStrategyRequest) (_result *CreateForwardStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateForwardStrategyResponse{}
	_body, _err := client.CreateForwardStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a department for a custom identity source in the current Alibaba Cloud account.
//
// @param request - CreateIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdpDepartmentResponse
func (client *Client) CreateIdpDepartmentWithOptions(request *CreateIdpDepartmentRequest, runtime *dara.RuntimeOptions) (_result *CreateIdpDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentName) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIdpDepartment"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIdpDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a department for a custom identity source in the current Alibaba Cloud account.
//
// @param request - CreateIdpDepartmentRequest
//
// @return CreateIdpDepartmentResponse
func (client *Client) CreateIdpDepartment(request *CreateIdpDepartmentRequest) (_result *CreateIdpDepartmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateIdpDepartmentResponse{}
	_body, _err := client.CreateIdpDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a diagnostic task for internal network access.
//
// @param tmpReq - CreatePADiagnosisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePADiagnosisTaskResponse
func (client *Client) CreatePADiagnosisTaskWithOptions(tmpReq *CreatePADiagnosisTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatePADiagnosisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePADiagnosisTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UdpExtraConfigs) {
		request.UdpExtraConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UdpExtraConfigs, dara.String("UdpExtraConfigs"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DevTag) {
		body["DevTag"] = request.DevTag
	}

	if !dara.IsNil(request.DiagnoseType) {
		body["DiagnoseType"] = request.DiagnoseType
	}

	if !dara.IsNil(request.Host) {
		body["Host"] = request.Host
	}

	if !dara.IsNil(request.PopId) {
		body["PopId"] = request.PopId
	}

	if !dara.IsNil(request.PopMode) {
		body["PopMode"] = request.PopMode
	}

	if !dara.IsNil(request.Port) {
		body["Port"] = request.Port
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.UdpExtraConfigsShrink) {
		body["UdpExtraConfigs"] = request.UdpExtraConfigsShrink
	}

	if !dara.IsNil(request.UserGroupId) {
		body["UserGroupId"] = request.UserGroupId
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePADiagnosisTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePADiagnosisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a diagnostic task for internal network access.
//
// @param request - CreatePADiagnosisTaskRequest
//
// @return CreatePADiagnosisTaskResponse
func (client *Client) CreatePADiagnosisTask(request *CreatePADiagnosisTaskRequest) (_result *CreatePADiagnosisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePADiagnosisTaskResponse{}
	_body, _err := client.CreatePADiagnosisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an internal-facing access application under the current Alibaba Cloud account.
//
// Description:
//
// You can create up to 500 internal-facing access applications by default.
//
// @param tmpReq - CreatePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessApplicationResponse
func (client *Client) CreatePrivateAccessApplicationWithOptions(tmpReq *CreatePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePrivateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.L7Config) {
		request.L7ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.L7Config, dara.String("L7Config"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UnauthorizedAccessConfig) {
		request.UnauthorizedAccessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UnauthorizedAccessConfig, dara.String("UnauthorizedAccessConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AddressGroups) {
		bodyFlat["AddressGroups"] = request.AddressGroups
	}

	if !dara.IsNil(request.Addresses) {
		bodyFlat["Addresses"] = request.Addresses
	}

	if !dara.IsNil(request.BrowserAccessStatus) {
		body["BrowserAccessStatus"] = request.BrowserAccessStatus
	}

	if !dara.IsNil(request.ConfigMode) {
		body["ConfigMode"] = request.ConfigMode
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.L7ConfigShrink) {
		body["L7Config"] = request.L7ConfigShrink
	}

	if !dara.IsNil(request.L7ProxyDomainAutomaticPrefix) {
		body["L7ProxyDomainAutomaticPrefix"] = request.L7ProxyDomainAutomaticPrefix
	}

	if !dara.IsNil(request.L7ProxyDomainCustom) {
		body["L7ProxyDomainCustom"] = request.L7ProxyDomainCustom
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PortRanges) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.UnauthorizedAccessConfigShrink) {
		body["UnauthorizedAccessConfig"] = request.UnauthorizedAccessConfigShrink
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an internal-facing access application under the current Alibaba Cloud account.
//
// Description:
//
// You can create up to 500 internal-facing access applications by default.
//
// @param request - CreatePrivateAccessApplicationRequest
//
// @return CreatePrivateAccessApplicationResponse
func (client *Client) CreatePrivateAccessApplication(request *CreatePrivateAccessApplicationRequest) (_result *CreatePrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePrivateAccessApplicationResponse{}
	_body, _err := client.CreatePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Private Access Policy
//
// Description:
//
// By default, up to 500 private access policies can be created.
//
// @param request - CreatePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessPolicyResponse
func (client *Client) CreatePrivateAccessPolicyWithOptions(request *CreatePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.CustomUserAttributes) {
		bodyFlat["CustomUserAttributes"] = request.CustomUserAttributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceAttributeAction) {
		body["DeviceAttributeAction"] = request.DeviceAttributeAction
	}

	if !dara.IsNil(request.DeviceAttributeId) {
		body["DeviceAttributeId"] = request.DeviceAttributeId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PolicyAction) {
		body["PolicyAction"] = request.PolicyAction
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.TriggerTemplateId) {
		body["TriggerTemplateId"] = request.TriggerTemplateId
	}

	if !dara.IsNil(request.TrustedProcessGroupIds) {
		bodyFlat["TrustedProcessGroupIds"] = request.TrustedProcessGroupIds
	}

	if !dara.IsNil(request.TrustedProcessStatus) {
		body["TrustedProcessStatus"] = request.TrustedProcessStatus
	}

	if !dara.IsNil(request.TrustedSoftwareIds) {
		bodyFlat["TrustedSoftwareIds"] = request.TrustedSoftwareIds
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserGroupMode) {
		body["UserGroupMode"] = request.UserGroupMode
	}

	if !dara.IsNil(request.ValidFrom) {
		body["ValidFrom"] = request.ValidFrom
	}

	if !dara.IsNil(request.ValidTimeStatus) {
		body["ValidTimeStatus"] = request.ValidTimeStatus
	}

	if !dara.IsNil(request.ValidUntil) {
		body["ValidUntil"] = request.ValidUntil
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Private Access Policy
//
// Description:
//
// By default, up to 500 private access policies can be created.
//
// @param request - CreatePrivateAccessPolicyRequest
//
// @return CreatePrivateAccessPolicyResponse
func (client *Client) CreatePrivateAccessPolicy(request *CreatePrivateAccessPolicyRequest) (_result *CreatePrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePrivateAccessPolicyResponse{}
	_body, _err := client.CreatePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an internal-facing access tag under the current Alibaba Cloud account.
//
// Description:
//
// You can create up to 500 internal-facing access tags by default.
//
// @param request - CreatePrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessTagResponse
func (client *Client) CreatePrivateAccessTagWithOptions(request *CreatePrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *CreatePrivateAccessTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an internal-facing access tag under the current Alibaba Cloud account.
//
// Description:
//
// You can create up to 500 internal-facing access tags by default.
//
// @param request - CreatePrivateAccessTagRequest
//
// @return CreatePrivateAccessTagResponse
func (client *Client) CreatePrivateAccessTag(request *CreatePrivateAccessTagRequest) (_result *CreatePrivateAccessTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePrivateAccessTagResponse{}
	_body, _err := client.CreatePrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a software ban policy.
//
// @param request - CreateProhibitedPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProhibitedPolicyResponse
func (client *Client) CreateProhibitedPolicyWithOptions(request *CreateProhibitedPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateProhibitedPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowReport) {
		body["AllowReport"] = request.AllowReport
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ForceKill) {
		body["ForceKill"] = request.ForceKill
	}

	if !dara.IsNil(request.MainButtonTextCh) {
		body["MainButtonTextCh"] = request.MainButtonTextCh
	}

	if !dara.IsNil(request.MainButtonTextEn) {
		body["MainButtonTextEn"] = request.MainButtonTextEn
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MinorButtonTextCh) {
		body["MinorButtonTextCh"] = request.MinorButtonTextCh
	}

	if !dara.IsNil(request.MinorButtonTextEn) {
		body["MinorButtonTextEn"] = request.MinorButtonTextEn
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.PromptCh) {
		body["PromptCh"] = request.PromptCh
	}

	if !dara.IsNil(request.PromptEn) {
		body["PromptEn"] = request.PromptEn
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.SoftwareIds) {
		bodyFlat["SoftwareIds"] = request.SoftwareIds
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.TitleCh) {
		body["TitleCh"] = request.TitleCh
	}

	if !dara.IsNil(request.TitleEn) {
		body["TitleEn"] = request.TitleEn
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProhibitedPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProhibitedPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a software ban policy.
//
// @param request - CreateProhibitedPolicyRequest
//
// @return CreateProhibitedPolicyResponse
func (client *Client) CreateProhibitedPolicy(request *CreateProhibitedPolicyRequest) (_result *CreateProhibitedPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProhibitedPolicyResponse{}
	_body, _err := client.CreateProhibitedPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom disabled software entry.
//
// @param request - CreateProhibitedSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProhibitedSoftwareResponse
func (client *Client) CreateProhibitedSoftwareWithOptions(request *CreateProhibitedSoftwareRequest, runtime *dara.RuntimeOptions) (_result *CreateProhibitedSoftwareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.LinuxProcesses) {
		bodyFlat["LinuxProcesses"] = request.LinuxProcesses
	}

	if !dara.IsNil(request.MacOSProcesses) {
		bodyFlat["MacOSProcesses"] = request.MacOSProcesses
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.WindowsProcesses) {
		bodyFlat["WindowsProcesses"] = request.WindowsProcesses
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProhibitedSoftware"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProhibitedSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom disabled software entry.
//
// @param request - CreateProhibitedSoftwareRequest
//
// @return CreateProhibitedSoftwareResponse
func (client *Client) CreateProhibitedSoftware(request *CreateProhibitedSoftwareRequest) (_result *CreateProhibitedSoftwareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProhibitedSoftwareResponse{}
	_body, _err := client.CreateProhibitedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom disabled software tag.
//
// @param request - CreateProhibitedTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProhibitedTagResponse
func (client *Client) CreateProhibitedTagWithOptions(request *CreateProhibitedTagRequest, runtime *dara.RuntimeOptions) (_result *CreateProhibitedTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProhibitedTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProhibitedTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom disabled software tag.
//
// @param request - CreateProhibitedTagRequest
//
// @return CreateProhibitedTagResponse
func (client *Client) CreateProhibitedTag(request *CreateProhibitedTagRequest) (_result *CreateProhibitedTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProhibitedTagResponse{}
	_body, _err := client.CreateProhibitedTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a device registration policy for your Alibaba Cloud account.
//
// @param tmpReq - CreateRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRegistrationPolicyResponse
func (client *Client) CreateRegistrationPolicyWithOptions(tmpReq *CreateRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateRegistrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateRegistrationPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CompanyLimitCount) {
		request.CompanyLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompanyLimitCount, dara.String("CompanyLimitCount"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PersonalLimitCount) {
		request.PersonalLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonalLimitCount, dara.String("PersonalLimitCount"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CompanyLimitCountShrink) {
		body["CompanyLimitCount"] = request.CompanyLimitCountShrink
	}

	if !dara.IsNil(request.CompanyLimitType) {
		body["CompanyLimitType"] = request.CompanyLimitType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PersonalLimitCountShrink) {
		body["PersonalLimitCount"] = request.PersonalLimitCountShrink
	}

	if !dara.IsNil(request.PersonalLimitType) {
		body["PersonalLimitType"] = request.PersonalLimitType
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a device registration policy for your Alibaba Cloud account.
//
// @param request - CreateRegistrationPolicyRequest
//
// @return CreateRegistrationPolicyResponse
func (client *Client) CreateRegistrationPolicy(request *CreateRegistrationPolicyRequest) (_result *CreateRegistrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRegistrationPolicyResponse{}
	_body, _err := client.CreateRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a software distribution task.
//
// Description:
//
// - After a task is created, its initial status is disabled.
//
// - MatchMode determines how to specify the matching target parameters: when set to UserGroupNormal, you must pass only UserGroupIds. When set to DeviceGroupNormal, you must pass only DeviceGroupIds. When set to DevTagNormal, you must pass only DevTags. Requests that contain parameters not matching the MatchMode value are rejected.
//
// - SupportOs supports only a single operating system value.
//
// @param request - CreateSoftwarelibDistributeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSoftwarelibDistributeTaskResponse
func (client *Client) CreateSoftwarelibDistributeTaskWithOptions(request *CreateSoftwarelibDistributeTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateSoftwarelibDistributeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DevTags) {
		bodyFlat["DevTags"] = request.DevTags
	}

	if !dara.IsNil(request.DeviceGroupIds) {
		bodyFlat["DeviceGroupIds"] = request.DeviceGroupIds
	}

	if !dara.IsNil(request.ExecuteMode) {
		body["ExecuteMode"] = request.ExecuteMode
	}

	if !dara.IsNil(request.ExecuteParameters) {
		body["ExecuteParameters"] = request.ExecuteParameters
	}

	if !dara.IsNil(request.ExecutePeriod) {
		body["ExecutePeriod"] = request.ExecutePeriod
	}

	if !dara.IsNil(request.ExpireMode) {
		body["ExpireMode"] = request.ExpireMode
	}

	if !dara.IsNil(request.GmtExpired) {
		body["GmtExpired"] = request.GmtExpired
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RetryTimes) {
		body["RetryTimes"] = request.RetryTimes
	}

	if !dara.IsNil(request.RunAsAccount) {
		body["RunAsAccount"] = request.RunAsAccount
	}

	if !dara.IsNil(request.SoftwareId) {
		body["SoftwareId"] = request.SoftwareId
	}

	if !dara.IsNil(request.SoftwareName) {
		body["SoftwareName"] = request.SoftwareName
	}

	if !dara.IsNil(request.SupportOs) {
		body["SupportOs"] = request.SupportOs
	}

	if !dara.IsNil(request.TaskType) {
		body["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.VersionId) {
		body["VersionId"] = request.VersionId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSoftwarelibDistributeTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSoftwarelibDistributeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a software distribution task.
//
// Description:
//
// - After a task is created, its initial status is disabled.
//
// - MatchMode determines how to specify the matching target parameters: when set to UserGroupNormal, you must pass only UserGroupIds. When set to DeviceGroupNormal, you must pass only DeviceGroupIds. When set to DevTagNormal, you must pass only DevTags. Requests that contain parameters not matching the MatchMode value are rejected.
//
// - SupportOs supports only a single operating system value.
//
// @param request - CreateSoftwarelibDistributeTaskRequest
//
// @return CreateSoftwarelibDistributeTaskResponse
func (client *Client) CreateSoftwarelibDistributeTask(request *CreateSoftwarelibDistributeTaskRequest) (_result *CreateSoftwarelibDistributeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSoftwarelibDistributeTaskResponse{}
	_body, _err := client.CreateSoftwarelibDistributeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a software version.
//
// Description:
//
// - Within the same software, the combination of operating system and version number must be unique. If a duplicate is created, a ResourceDuplicated error is returned.
//
// - A newly created version has an initial publish status of unpublished.
//
// - A newly created version has the highest priority. The priorities of other versions under the same software are shifted down accordingly.
//
// @param request - CreateSoftwarelibVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSoftwarelibVersionResponse
func (client *Client) CreateSoftwarelibVersionWithOptions(request *CreateSoftwarelibVersionRequest, runtime *dara.RuntimeOptions) (_result *CreateSoftwarelibVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Md5) {
		body["Md5"] = request.Md5
	}

	if !dara.IsNil(request.Os) {
		body["Os"] = request.Os
	}

	if !dara.IsNil(request.PublisherType) {
		body["PublisherType"] = request.PublisherType
	}

	if !dara.IsNil(request.SoftwareId) {
		body["SoftwareId"] = request.SoftwareId
	}

	if !dara.IsNil(request.SoftwareName) {
		body["SoftwareName"] = request.SoftwareName
	}

	if !dara.IsNil(request.SoftwarePkgName) {
		body["SoftwarePkgName"] = request.SoftwarePkgName
	}

	if !dara.IsNil(request.SoftwarePkgSize) {
		body["SoftwarePkgSize"] = request.SoftwarePkgSize
	}

	if !dara.IsNil(request.SoftwareUrl) {
		body["SoftwareUrl"] = request.SoftwareUrl
	}

	if !dara.IsNil(request.SoftwareVersion) {
		body["SoftwareVersion"] = request.SoftwareVersion
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSoftwarelibVersion"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSoftwarelibVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a software version.
//
// Description:
//
// - Within the same software, the combination of operating system and version number must be unique. If a duplicate is created, a ResourceDuplicated error is returned.
//
// - A newly created version has an initial publish status of unpublished.
//
// - A newly created version has the highest priority. The priorities of other versions under the same software are shifted down accordingly.
//
// @param request - CreateSoftwarelibVersionRequest
//
// @return CreateSoftwarelibVersionResponse
func (client *Client) CreateSoftwarelibVersion(request *CreateSoftwarelibVersionRequest) (_result *CreateSoftwarelibVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSoftwarelibVersionResponse{}
	_body, _err := client.CreateSoftwarelibVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user group for your Alibaba Cloud account.
//
// Description:
//
// You can create up to 500 user groups.
//
// @param request - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		bodyFlat["Attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user group for your Alibaba Cloud account.
//
// Description:
//
// You can create up to 500 user groups.
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduled virus scan policy that automatically sends scan tasks to user terminal devices within the effective scope based on the configured cycle.
//
// @param request - CreateVirusScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirusScanScheduledStrategyResponse
func (client *Client) CreateVirusScanScheduledStrategyWithOptions(request *CreateVirusScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *CreateVirusScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HighRiskOperation) {
		body["HighRiskOperation"] = request.HighRiskOperation
	}

	if !dara.IsNil(request.LowRiskOperation) {
		body["LowRiskOperation"] = request.LowRiskOperation
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MaxCpuUsage) {
		body["MaxCpuUsage"] = request.MaxCpuUsage
	}

	if !dara.IsNil(request.MidRiskOperation) {
		body["MidRiskOperation"] = request.MidRiskOperation
	}

	if !dara.IsNil(request.PerformanceMode) {
		body["PerformanceMode"] = request.PerformanceMode
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ScanBeginTime) {
		body["ScanBeginTime"] = request.ScanBeginTime
	}

	if !dara.IsNil(request.ScanEndTime) {
		body["ScanEndTime"] = request.ScanEndTime
	}

	if !dara.IsNil(request.ScanFrequency) {
		body["ScanFrequency"] = request.ScanFrequency
	}

	if !dara.IsNil(request.ScanInterval) {
		body["ScanInterval"] = request.ScanInterval
	}

	if !dara.IsNil(request.ScanMode) {
		body["ScanMode"] = request.ScanMode
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ScanPath) {
		bodyFlat["ScanPath"] = request.ScanPath
	}

	if !dara.IsNil(request.ScanTargets) {
		bodyFlat["ScanTargets"] = request.ScanTargets
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategyDescription) {
		body["StrategyDescription"] = request.StrategyDescription
	}

	if !dara.IsNil(request.StrategyName) {
		body["StrategyName"] = request.StrategyName
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirusScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirusScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled virus scan policy that automatically sends scan tasks to user terminal devices within the effective scope based on the configured cycle.
//
// @param request - CreateVirusScanScheduledStrategyRequest
//
// @return CreateVirusScanScheduledStrategyResponse
func (client *Client) CreateVirusScanScheduledStrategy(request *CreateVirusScanScheduledStrategyRequest) (_result *CreateVirusScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVirusScanScheduledStrategyResponse{}
	_body, _err := client.CreateVirusScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instant virus scan task and delivers it to user endpoint devices within the effective scope. The task takes effect immediately after creation. A maximum of 10 tasks can be created per Alibaba Cloud account per minute.
//
// @param request - CreateVirusScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVirusScanTaskResponse
func (client *Client) CreateVirusScanTaskWithOptions(request *CreateVirusScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVirusScanTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.HighRiskOperation) {
		body["HighRiskOperation"] = request.HighRiskOperation
	}

	if !dara.IsNil(request.LowRiskOperation) {
		body["LowRiskOperation"] = request.LowRiskOperation
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MaxCpuUsage) {
		body["MaxCpuUsage"] = request.MaxCpuUsage
	}

	if !dara.IsNil(request.MidRiskOperation) {
		body["MidRiskOperation"] = request.MidRiskOperation
	}

	if !dara.IsNil(request.PerformanceMode) {
		body["PerformanceMode"] = request.PerformanceMode
	}

	if !dara.IsNil(request.ScanMode) {
		body["ScanMode"] = request.ScanMode
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ScanPath) {
		bodyFlat["ScanPath"] = request.ScanPath
	}

	if !dara.IsNil(request.ScanTargets) {
		bodyFlat["ScanTargets"] = request.ScanTargets
	}

	if !dara.IsNil(request.TaskDescription) {
		body["TaskDescription"] = request.TaskDescription
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVirusScanTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVirusScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instant virus scan task and delivers it to user endpoint devices within the effective scope. The task takes effect immediately after creation. A maximum of 10 tasks can be created per Alibaba Cloud account per minute.
//
// @param request - CreateVirusScanTaskRequest
//
// @return CreateVirusScanTaskResponse
func (client *Client) CreateVirusScanTask(request *CreateVirusScanTaskRequest) (_result *CreateVirusScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVirusScanTaskResponse{}
	_body, _err := client.CreateVirusScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scheduled vulnerability scanning policy that automatically sends vulnerability scanning tasks to user endpoint devices within the effective scope based on the configured cycle.
//
// @param request - CreateVulScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVulScanScheduledStrategyResponse
func (client *Client) CreateVulScanScheduledStrategyWithOptions(request *CreateVulScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *CreateVulScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ScanBeginTime) {
		body["ScanBeginTime"] = request.ScanBeginTime
	}

	if !dara.IsNil(request.ScanEndTime) {
		body["ScanEndTime"] = request.ScanEndTime
	}

	if !dara.IsNil(request.ScanFrequency) {
		body["ScanFrequency"] = request.ScanFrequency
	}

	if !dara.IsNil(request.ScanInterval) {
		body["ScanInterval"] = request.ScanInterval
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategyDescription) {
		body["StrategyDescription"] = request.StrategyDescription
	}

	if !dara.IsNil(request.StrategyName) {
		body["StrategyName"] = request.StrategyName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVulScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVulScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scheduled vulnerability scanning policy that automatically sends vulnerability scanning tasks to user endpoint devices within the effective scope based on the configured cycle.
//
// @param request - CreateVulScanScheduledStrategyRequest
//
// @return CreateVulScanScheduledStrategyResponse
func (client *Client) CreateVulScanScheduledStrategy(request *CreateVulScanScheduledStrategyRequest) (_result *CreateVulScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVulScanScheduledStrategyResponse{}
	_body, _err := client.CreateVulScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an instant vulnerability scanning task and delivers it to user endpoint devices within the effective scope.
//
// @param request - CreateVulScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVulScanTaskResponse
func (client *Client) CreateVulScanTaskWithOptions(request *CreateVulScanTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVulScanTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndTimestamp) {
		body["EndTimestamp"] = request.EndTimestamp
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.TaskDescription) {
		body["TaskDescription"] = request.TaskDescription
	}

	if !dara.IsNil(request.TaskName) {
		body["TaskName"] = request.TaskName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVulScanTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVulScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an instant vulnerability scanning task and delivers it to user endpoint devices within the effective scope.
//
// @param request - CreateVulScanTaskRequest
//
// @return CreateVulScanTaskResponse
func (client *Client) CreateVulScanTask(request *CreateVulScanTaskRequest) (_result *CreateVulScanTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVulScanTaskResponse{}
	_body, _err := client.CreateVulScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a vulnerability fix task that delivers the patch for a specified vulnerability to user endpoint devices and performs the installation.
//
// @param tmpReq - CreateVulnerabilityFixTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVulnerabilityFixTaskResponse
func (client *Client) CreateVulnerabilityFixTaskWithOptions(tmpReq *CreateVulnerabilityFixTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateVulnerabilityFixTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateVulnerabilityFixTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WuyingVulFixConfig) {
		request.WuyingVulFixConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WuyingVulFixConfig, dara.String("WuyingVulFixConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DevTags) {
		bodyFlat["DevTags"] = request.DevTags
	}

	if !dara.IsNil(request.FixMode) {
		body["FixMode"] = request.FixMode
	}

	if !dara.IsNil(request.MaxDownloadSpeed) {
		body["MaxDownloadSpeed"] = request.MaxDownloadSpeed
	}

	if !dara.IsNil(request.UpdateId) {
		body["UpdateId"] = request.UpdateId
	}

	if !dara.IsNil(request.WuyingVulFixConfigShrink) {
		body["WuyingVulFixConfig"] = request.WuyingVulFixConfigShrink
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVulnerabilityFixTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVulnerabilityFixTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a vulnerability fix task that delivers the patch for a specified vulnerability to user endpoint devices and performs the installation.
//
// @param request - CreateVulnerabilityFixTaskRequest
//
// @return CreateVulnerabilityFixTaskResponse
func (client *Client) CreateVulnerabilityFixTask(request *CreateVulnerabilityFixTaskRequest) (_result *CreateVulnerabilityFixTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVulnerabilityFixTaskResponse{}
	_body, _err := client.CreateVulnerabilityFixTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the invisible watermark transparent background image for web watermarks, screen watermarks, and App watermarks.
//
// @param tmpReq - CreateWmBaseImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmBaseImageResponse
func (client *Client) CreateWmBaseImageWithOptions(tmpReq *CreateWmBaseImageRequest, runtime *dara.RuntimeOptions) (_result *CreateWmBaseImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWmBaseImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageControl) {
		request.ImageControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageControl, dara.String("ImageControl"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["comment"] = request.Comment
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Height) {
		body["Height"] = request.Height
	}

	if !dara.IsNil(request.ImageControlShrink) {
		body["ImageControl"] = request.ImageControlShrink
	}

	if !dara.IsNil(request.Opacity) {
		body["Opacity"] = request.Opacity
	}

	if !dara.IsNil(request.Scale) {
		body["Scale"] = request.Scale
	}

	if !dara.IsNil(request.Width) {
		body["Width"] = request.Width
	}

	if !dara.IsNil(request.WmInfoBytesB64) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmInfoUint) {
		body["WmInfoUint"] = request.WmInfoUint
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmBaseImage"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmBaseImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the invisible watermark transparent background image for web watermarks, screen watermarks, and App watermarks.
//
// @param request - CreateWmBaseImageRequest
//
// @return CreateWmBaseImageResponse
func (client *Client) CreateWmBaseImage(request *CreateWmBaseImageRequest) (_result *CreateWmBaseImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWmBaseImageResponse{}
	_body, _err := client.CreateWmBaseImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a digital watermarking embedding task.
//
// Description:
//
// You can create a maximum of 500 user groups by default.
//
// @param tmpReq - CreateWmEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmEmbedTaskResponse
func (client *Client) CreateWmEmbedTaskWithOptions(tmpReq *CreateWmEmbedTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateWmEmbedTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWmEmbedTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AudioControl) {
		request.AudioControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AudioControl, dara.String("AudioControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CsvControl) {
		request.CsvControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CsvControl, dara.String("CsvControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DocumentControl) {
		request.DocumentControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentControl, dara.String("DocumentControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ImageControl) {
		request.ImageControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageControl, dara.String("ImageControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoControl) {
		request.VideoControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoControl, dara.String("VideoControl"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AudioControlShrink) {
		body["AudioControl"] = request.AudioControlShrink
	}

	if !dara.IsNil(request.CsvControlShrink) {
		body["CsvControl"] = request.CsvControlShrink
	}

	if !dara.IsNil(request.DocumentControlShrink) {
		body["DocumentControl"] = request.DocumentControlShrink
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Filename) {
		body["Filename"] = request.Filename
	}

	if !dara.IsNil(request.ImageControlShrink) {
		body["ImageControl"] = request.ImageControlShrink
	}

	if !dara.IsNil(request.ImageEmbedJpegQuality) {
		body["ImageEmbedJpegQuality"] = request.ImageEmbedJpegQuality
	}

	if !dara.IsNil(request.ImageEmbedLevel) {
		body["ImageEmbedLevel"] = request.ImageEmbedLevel
	}

	if !dara.IsNil(request.InvisibleEnable) {
		body["InvisibleEnable"] = request.InvisibleEnable
	}

	if !dara.IsNil(request.VideoBitrate) {
		body["VideoBitrate"] = request.VideoBitrate
	}

	if !dara.IsNil(request.VideoControlShrink) {
		body["VideoControl"] = request.VideoControlShrink
	}

	if !dara.IsNil(request.VideoIsLong) {
		body["VideoIsLong"] = request.VideoIsLong
	}

	if !dara.IsNil(request.WmInfoBytesB64) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmInfoUint) {
		body["WmInfoUint"] = request.WmInfoUint
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmEmbedTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmEmbedTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a digital watermarking embedding task.
//
// Description:
//
// You can create a maximum of 500 user groups by default.
//
// @param request - CreateWmEmbedTaskRequest
//
// @return CreateWmEmbedTaskResponse
func (client *Client) CreateWmEmbedTask(request *CreateWmEmbedTaskRequest) (_result *CreateWmEmbedTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWmEmbedTaskResponse{}
	_body, _err := client.CreateWmEmbedTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a digital watermarking extraction task.
//
// @param tmpReq - CreateWmExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmExtractTaskResponse
func (client *Client) CreateWmExtractTaskWithOptions(tmpReq *CreateWmExtractTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateWmExtractTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateWmExtractTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CsvControl) {
		request.CsvControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CsvControl, dara.String("CsvControl"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ImageExtractParamsOpenApi) {
		request.ImageExtractParamsOpenApiShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageExtractParamsOpenApi, dara.String("ImageExtractParamsOpenApi"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CsvControlShrink) {
		query["CsvControl"] = request.CsvControlShrink
	}

	if !dara.IsNil(request.ImageExtractParamsOpenApiShrink) {
		query["ImageExtractParamsOpenApi"] = request.ImageExtractParamsOpenApiShrink
	}

	if !dara.IsNil(request.IsClientEmbed) {
		query["IsClientEmbed"] = request.IsClientEmbed
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIsCapture) {
		body["DocumentIsCapture"] = request.DocumentIsCapture
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Filename) {
		body["Filename"] = request.Filename
	}

	if !dara.IsNil(request.VideoIsLong) {
		body["VideoIsLong"] = request.VideoIsLong
	}

	if !dara.IsNil(request.VideoSpeed) {
		body["VideoSpeed"] = request.VideoSpeed
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmExtractTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmExtractTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a digital watermarking extraction task.
//
// @param request - CreateWmExtractTaskRequest
//
// @return CreateWmExtractTaskResponse
func (client *Client) CreateWmExtractTask(request *CreateWmExtractTaskRequest) (_result *CreateWmExtractTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWmExtractTaskResponse{}
	_body, _err := client.CreateWmExtractTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mapping from string-format watermark information to digital-format watermark information.
//
// @param request - CreateWmInfoMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmInfoMappingResponse
func (client *Client) CreateWmInfoMappingWithOptions(request *CreateWmInfoMappingRequest, runtime *dara.RuntimeOptions) (_result *CreateWmInfoMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WmInfoBytesB64) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !dara.IsNil(request.WmInfoSize) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !dara.IsNil(request.WmType) {
		body["WmType"] = request.WmType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWmInfoMapping"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWmInfoMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mapping from string-format watermark information to digital-format watermark information.
//
// @param request - CreateWmInfoMappingRequest
//
// @return CreateWmInfoMappingResponse
func (client *Client) CreateWmInfoMapping(request *CreateWmInfoMappingRequest) (_result *CreateWmInfoMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWmInfoMappingResponse{}
	_body, _err := client.CreateWmInfoMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes approval processes in batches from your Alibaba Cloud account.
//
// @param request - DeleteApprovalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApprovalProcessesResponse
func (client *Client) DeleteApprovalProcessesWithOptions(request *DeleteApprovalProcessesRequest, runtime *dara.RuntimeOptions) (_result *DeleteApprovalProcessesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ProcessIds) {
		bodyFlat["ProcessIds"] = request.ProcessIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApprovalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApprovalProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes approval processes in batches from your Alibaba Cloud account.
//
// @param request - DeleteApprovalProcessesRequest
//
// @return DeleteApprovalProcessesResponse
func (client *Client) DeleteApprovalProcesses(request *DeleteApprovalProcessesRequest) (_result *DeleteApprovalProcessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteApprovalProcessesResponse{}
	_body, _err := client.DeleteApprovalProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a specified user from your Alibaba Cloud account\\"s custom identity source.
//
// @param request - DeleteClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientUserResponse
func (client *Client) DeleteClientUserWithOptions(request *DeleteClientUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a specified user from your Alibaba Cloud account\\"s custom identity source.
//
// @param request - DeleteClientUserRequest
//
// @return DeleteClientUserResponse
func (client *Client) DeleteClientUser(request *DeleteClientUserRequest) (_result *DeleteClientUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientUserResponse{}
	_body, _err := client.DeleteClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a connector.
//
// @param request - DeleteConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectorResponse
func (client *Client) DeleteConnectorWithOptions(request *DeleteConnectorRequest, runtime *dara.RuntimeOptions) (_result *DeleteConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a connector.
//
// @param request - DeleteConnectorRequest
//
// @return DeleteConnectorResponse
func (client *Client) DeleteConnector(request *DeleteConnectorRequest) (_result *DeleteConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConnectorResponse{}
	_body, _err := client.DeleteConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a ConnectorClient under the current Alibaba Cloud account.
//
// @param request - DeleteConnectorClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectorClientResponse
func (client *Client) DeleteConnectorClientWithOptions(request *DeleteConnectorClientRequest, runtime *dara.RuntimeOptions) (_result *DeleteConnectorClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.DevTag) {
		body["DevTag"] = request.DevTag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnectorClient"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectorClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a ConnectorClient under the current Alibaba Cloud account.
//
// @param request - DeleteConnectorClientRequest
//
// @return DeleteConnectorClientResponse
func (client *Client) DeleteConnectorClient(request *DeleteConnectorClientRequest) (_result *DeleteConnectorClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConnectorClientResponse{}
	_body, _err := client.DeleteConnectorClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes instance tags in batches.
//
// @param request - DeleteDeviceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDeviceGroupsResponse
func (client *Client) DeleteDeviceGroupsWithOptions(request *DeleteDeviceGroupsRequest, runtime *dara.RuntimeOptions) (_result *DeleteDeviceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupIds) {
		bodyFlat["DeviceGroupIds"] = request.DeviceGroupIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDeviceGroups"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDeviceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes instance tags in batches.
//
// @param request - DeleteDeviceGroupsRequest
//
// @return DeleteDeviceGroupsResponse
func (client *Client) DeleteDeviceGroups(request *DeleteDeviceGroupsRequest) (_result *DeleteDeviceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDeviceGroupsResponse{}
	_body, _err := client.DeleteDeviceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes detection records of a specified vulnerability from specified user endpoint devices in batches.
//
// @param request - DeleteDevicesVulnerabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDevicesVulnerabilityResponse
func (client *Client) DeleteDevicesVulnerabilityWithOptions(request *DeleteDevicesVulnerabilityRequest, runtime *dara.RuntimeOptions) (_result *DeleteDevicesVulnerabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DevTags) {
		bodyFlat["DevTags"] = request.DevTags
	}

	if !dara.IsNil(request.UpdateId) {
		body["UpdateId"] = request.UpdateId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDevicesVulnerability"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDevicesVulnerabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes detection records of a specified vulnerability from specified user endpoint devices in batches.
//
// @param request - DeleteDevicesVulnerabilityRequest
//
// @return DeleteDevicesVulnerabilityResponse
func (client *Client) DeleteDevicesVulnerability(request *DeleteDevicesVulnerabilityRequest) (_result *DeleteDevicesVulnerabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDevicesVulnerabilityResponse{}
	_body, _err := client.DeleteDevicesVulnerabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a domain name list.
//
// Description:
//
// Deletes a specified domain name list under the current tenant. Before deletion, the system checks whether any domain name policy references the list. If a reference exists, the deletion is rejected.
//
// @param request - DeleteDomainMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainMetaResponse
func (client *Client) DeleteDomainMetaWithOptions(request *DeleteDomainMetaRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ListId) {
		body["ListId"] = request.ListId
	}

	if !dara.IsNil(request.ListType) {
		body["ListType"] = request.ListType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDomainMeta"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a domain name list.
//
// Description:
//
// Deletes a specified domain name list under the current tenant. Before deletion, the system checks whether any domain name policy references the list. If a reference exists, the deletion is rejected.
//
// @param request - DeleteDomainMetaRequest
//
// @return DeleteDomainMetaResponse
func (client *Client) DeleteDomainMeta(request *DeleteDomainMetaRequest) (_result *DeleteDomainMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainMetaResponse{}
	_body, _err := client.DeleteDomainMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a dynamic route from your current Alibaba Cloud account.
//
// @param request - DeleteDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicRouteResponse
func (client *Client) DeleteDynamicRouteWithOptions(request *DeleteDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *DeleteDynamicRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DynamicRouteId) {
		query["DynamicRouteId"] = request.DynamicRouteId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a dynamic route from your current Alibaba Cloud account.
//
// @param request - DeleteDynamicRouteRequest
//
// @return DeleteDynamicRouteResponse
func (client *Client) DeleteDynamicRoute(request *DeleteDynamicRouteRequest) (_result *DeleteDynamicRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDynamicRouteResponse{}
	_body, _err := client.DeleteDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an enterprise acceleration policy.
//
// @param request - DeleteEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseAcceleratePolicyResponse
func (client *Client) DeleteEnterpriseAcceleratePolicyWithOptions(request *DeleteEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseAcceleratePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an enterprise acceleration policy.
//
// @param request - DeleteEnterpriseAcceleratePolicyRequest
//
// @return DeleteEnterpriseAcceleratePolicyResponse
func (client *Client) DeleteEnterpriseAcceleratePolicy(request *DeleteEnterpriseAcceleratePolicyRequest) (_result *DeleteEnterpriseAcceleratePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.DeleteEnterpriseAcceleratePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes enterprise acceleration addresses.
//
// @param request - DeleteEnterpriseAccelerateTargetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEnterpriseAccelerateTargetResponse
func (client *Client) DeleteEnterpriseAccelerateTargetWithOptions(request *DeleteEnterpriseAccelerateTargetRequest, runtime *dara.RuntimeOptions) (_result *DeleteEnterpriseAccelerateTargetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Target) {
		bodyFlat["Target"] = request.Target
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEnterpriseAccelerateTarget"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEnterpriseAccelerateTargetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes enterprise acceleration addresses.
//
// @param request - DeleteEnterpriseAccelerateTargetRequest
//
// @return DeleteEnterpriseAccelerateTargetResponse
func (client *Client) DeleteEnterpriseAccelerateTarget(request *DeleteEnterpriseAccelerateTargetRequest) (_result *DeleteEnterpriseAccelerateTargetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEnterpriseAccelerateTargetResponse{}
	_body, _err := client.DeleteEnterpriseAccelerateTargetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule.
//
// @param request - DeleteForwardStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteForwardStrategyResponse
func (client *Client) DeleteForwardStrategyWithOptions(request *DeleteForwardStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteForwardStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ForwardId) {
		body["ForwardId"] = request.ForwardId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteForwardStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteForwardStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a forwarding rule.
//
// @param request - DeleteForwardStrategyRequest
//
// @return DeleteForwardStrategyResponse
func (client *Client) DeleteForwardStrategy(request *DeleteForwardStrategyRequest) (_result *DeleteForwardStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteForwardStrategyResponse{}
	_body, _err := client.DeleteForwardStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a department from a custom identity provider in your Alibaba Cloud account.
//
// @param request - DeleteIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdpDepartmentResponse
func (client *Client) DeleteIdpDepartmentWithOptions(request *DeleteIdpDepartmentRequest, runtime *dara.RuntimeOptions) (_result *DeleteIdpDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIdpDepartment"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIdpDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a department from a custom identity provider in your Alibaba Cloud account.
//
// @param request - DeleteIdpDepartmentRequest
//
// @return DeleteIdpDepartmentResponse
func (client *Client) DeleteIdpDepartment(request *DeleteIdpDepartmentRequest) (_result *DeleteIdpDepartmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIdpDepartmentResponse{}
	_body, _err := client.DeleteIdpDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user\\"s One-Time Password (OTP) configuration.
//
// @param request - DeleteOtpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOtpConfigResponse
func (client *Client) DeleteOtpConfigWithOptions(request *DeleteOtpConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteOtpConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOtpConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOtpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user\\"s One-Time Password (OTP) configuration.
//
// @param request - DeleteOtpConfigRequest
//
// @return DeleteOtpConfigResponse
func (client *Client) DeleteOtpConfig(request *DeleteOtpConfigRequest) (_result *DeleteOtpConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteOtpConfigResponse{}
	_body, _err := client.DeleteOtpConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an internal network access application from your Alibaba Cloud account.
//
// Description:
//
// You cannot delete an application if it is referenced by an office zone or a policy. For more information, see:
//
// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists internal network access applications.
//
// - [ListPrivateAccessPolicies](~~ListPrivateAccessPolices~~): Lists internal network access policies.
//
// @param request - DeletePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessApplicationResponse
func (client *Client) DeletePrivateAccessApplicationWithOptions(request *DeletePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an internal network access application from your Alibaba Cloud account.
//
// Description:
//
// You cannot delete an application if it is referenced by an office zone or a policy. For more information, see:
//
// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists internal network access applications.
//
// - [ListPrivateAccessPolicies](~~ListPrivateAccessPolices~~): Lists internal network access policies.
//
// @param request - DeletePrivateAccessApplicationRequest
//
// @return DeletePrivateAccessApplicationResponse
func (client *Client) DeletePrivateAccessApplication(request *DeletePrivateAccessApplicationRequest) (_result *DeletePrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateAccessApplicationResponse{}
	_body, _err := client.DeletePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a private network access policy for the current Alibaba Cloud account.
//
// @param request - DeletePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessPolicyResponse
func (client *Client) DeletePrivateAccessPolicyWithOptions(request *DeletePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a private network access policy for the current Alibaba Cloud account.
//
// @param request - DeletePrivateAccessPolicyRequest
//
// @return DeletePrivateAccessPolicyResponse
func (client *Client) DeletePrivateAccessPolicy(request *DeletePrivateAccessPolicyRequest) (_result *DeletePrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateAccessPolicyResponse{}
	_body, _err := client.DeletePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an internal access tag from the current Alibaba Cloud account.
//
// Description:
//
// Deletion is not allowed when the tag is referenced by applications, office networks, or policies. References:
//
// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists internal access applications.
//
// - [ListPrivateAccessTags](~~ListPrivateAccessTags~~): Lists internal access tags.
//
// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): Lists internal access policies.
//
// @param request - DeletePrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessTagResponse
func (client *Client) DeletePrivateAccessTagWithOptions(request *DeletePrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *DeletePrivateAccessTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TagId) {
		body["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an internal access tag from the current Alibaba Cloud account.
//
// Description:
//
// Deletion is not allowed when the tag is referenced by applications, office networks, or policies. References:
//
// - [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): Lists internal access applications.
//
// - [ListPrivateAccessTags](~~ListPrivateAccessTags~~): Lists internal access tags.
//
// - [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): Lists internal access policies.
//
// @param request - DeletePrivateAccessTagRequest
//
// @return DeletePrivateAccessTagResponse
func (client *Client) DeletePrivateAccessTag(request *DeletePrivateAccessTagRequest) (_result *DeletePrivateAccessTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePrivateAccessTagResponse{}
	_body, _err := client.DeletePrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes software prohibition policies in batches.
//
// @param request - DeleteProhibitedPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProhibitedPoliciesResponse
func (client *Client) DeleteProhibitedPoliciesWithOptions(request *DeleteProhibitedPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteProhibitedPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIds) {
		bodyFlat["PolicyIds"] = request.PolicyIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProhibitedPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProhibitedPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes software prohibition policies in batches.
//
// @param request - DeleteProhibitedPoliciesRequest
//
// @return DeleteProhibitedPoliciesResponse
func (client *Client) DeleteProhibitedPolicies(request *DeleteProhibitedPoliciesRequest) (_result *DeleteProhibitedPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProhibitedPoliciesResponse{}
	_body, _err := client.DeleteProhibitedPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes custom prohibited software in batches.
//
// @param request - DeleteProhibitedSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProhibitedSoftwareResponse
func (client *Client) DeleteProhibitedSoftwareWithOptions(request *DeleteProhibitedSoftwareRequest, runtime *dara.RuntimeOptions) (_result *DeleteProhibitedSoftwareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.SoftwareIds) {
		bodyFlat["SoftwareIds"] = request.SoftwareIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProhibitedSoftware"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProhibitedSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom prohibited software in batches.
//
// @param request - DeleteProhibitedSoftwareRequest
//
// @return DeleteProhibitedSoftwareResponse
func (client *Client) DeleteProhibitedSoftware(request *DeleteProhibitedSoftwareRequest) (_result *DeleteProhibitedSoftwareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProhibitedSoftwareResponse{}
	_body, _err := client.DeleteProhibitedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes custom prohibited software labels in batches.
//
// @param request - DeleteProhibitedTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProhibitedTagsResponse
func (client *Client) DeleteProhibitedTagsWithOptions(request *DeleteProhibitedTagsRequest, runtime *dara.RuntimeOptions) (_result *DeleteProhibitedTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProhibitedTags"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProhibitedTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom prohibited software labels in batches.
//
// @param request - DeleteProhibitedTagsRequest
//
// @return DeleteProhibitedTagsResponse
func (client *Client) DeleteProhibitedTags(request *DeleteProhibitedTagsRequest) (_result *DeleteProhibitedTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProhibitedTagsResponse{}
	_body, _err := client.DeleteProhibitedTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch delete device registration policies under your Alibaba Cloud account.
//
// @param request - DeleteRegistrationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistrationPoliciesResponse
func (client *Client) DeleteRegistrationPoliciesWithOptions(request *DeleteRegistrationPoliciesRequest, runtime *dara.RuntimeOptions) (_result *DeleteRegistrationPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.PolicyIds) {
		bodyFlat["PolicyIds"] = request.PolicyIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRegistrationPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRegistrationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch delete device registration policies under your Alibaba Cloud account.
//
// @param request - DeleteRegistrationPoliciesRequest
//
// @return DeleteRegistrationPoliciesResponse
func (client *Client) DeleteRegistrationPolicies(request *DeleteRegistrationPoliciesRequest) (_result *DeleteRegistrationPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRegistrationPoliciesResponse{}
	_body, _err := client.DeleteRegistrationPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes user endpoint devices in batches.
//
// Description:
//
// You can delete up to 100 devices at a time. Each device must be in a non-online status. If some device IDs in the specified collection do not meet the status requirement, only the devices that meet the requirement are deleted, and the operation still returns a success response.
//
// @param request - DeleteUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserDevicesResponse
func (client *Client) DeleteUserDevicesWithOptions(request *DeleteUserDevicesRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes user endpoint devices in batches.
//
// Description:
//
// You can delete up to 100 devices at a time. Each device must be in a non-online status. If some device IDs in the specified collection do not meet the status requirement, only the devices that meet the requirement are deleted, and the operation still returns a success response.
//
// @param request - DeleteUserDevicesRequest
//
// @return DeleteUserDevicesResponse
func (client *Client) DeleteUserDevices(request *DeleteUserDevicesRequest) (_result *DeleteUserDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserDevicesResponse{}
	_body, _err := client.DeleteUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a user group in your Alibaba Cloud account.
//
// Description:
//
// For more information, see:
//
// - [ListPolicesForUserGroup](~~ListPolicesForUserGroup~~): Query policies attached to a user group.
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupId) {
		body["UserGroupId"] = request.UserGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a user group in your Alibaba Cloud account.
//
// Description:
//
// For more information, see:
//
// - [ListPolicesForUserGroup](~~ListPolicesForUserGroup~~): Query policies attached to a user group.
//
// @param request - DeleteUserGroupRequest
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a virus file record that failed to be handled. Only records with a handling action of Fail can be deleted. This operation does not delete the actual file on the user\\"s endpoint device.
//
// @param request - DeleteVirusFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirusFileResponse
func (client *Client) DeleteVirusFileWithOptions(request *DeleteVirusFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirusFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DevTag) {
		body["DevTag"] = request.DevTag
	}

	if !dara.IsNil(request.FileMd5) {
		body["FileMd5"] = request.FileMd5
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirusFile"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirusFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a virus file record that failed to be handled. Only records with a handling action of Fail can be deleted. This operation does not delete the actual file on the user\\"s endpoint device.
//
// @param request - DeleteVirusFileRequest
//
// @return DeleteVirusFileResponse
func (client *Client) DeleteVirusFile(request *DeleteVirusFileRequest) (_result *DeleteVirusFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVirusFileResponse{}
	_body, _err := client.DeleteVirusFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes virus scheduled scan policies in batches. After deletion, no new scan tasks are triggered, but scan tasks that have already been dispatched are not affected. If any policy ID does not belong to the current Alibaba Cloud account, the entire deletion fails.
//
// @param request - DeleteVirusScanScheduledStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVirusScanScheduledStrategiesResponse
func (client *Client) DeleteVirusScanScheduledStrategiesWithOptions(request *DeleteVirusScanScheduledStrategiesRequest, runtime *dara.RuntimeOptions) (_result *DeleteVirusScanScheduledStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.StrategyIds) {
		bodyFlat["StrategyIds"] = request.StrategyIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVirusScanScheduledStrategies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVirusScanScheduledStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes virus scheduled scan policies in batches. After deletion, no new scan tasks are triggered, but scan tasks that have already been dispatched are not affected. If any policy ID does not belong to the current Alibaba Cloud account, the entire deletion fails.
//
// @param request - DeleteVirusScanScheduledStrategiesRequest
//
// @return DeleteVirusScanScheduledStrategiesResponse
func (client *Client) DeleteVirusScanScheduledStrategies(request *DeleteVirusScanScheduledStrategiesRequest) (_result *DeleteVirusScanScheduledStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVirusScanScheduledStrategiesResponse{}
	_body, _err := client.DeleteVirusScanScheduledStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specified scheduled vulnerability scanning policy.
//
// @param request - DeleteVulScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVulScanScheduledStrategyResponse
func (client *Client) DeleteVulScanScheduledStrategyWithOptions(request *DeleteVulScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *DeleteVulScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.StrategyId) {
		body["StrategyId"] = request.StrategyId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVulScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVulScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specified scheduled vulnerability scanning policy.
//
// @param request - DeleteVulScanScheduledStrategyRequest
//
// @return DeleteVulScanScheduledStrategyResponse
func (client *Client) DeleteVulScanScheduledStrategy(request *DeleteVulScanScheduledStrategyRequest) (_result *DeleteVulScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVulScanScheduledStrategyResponse{}
	_body, _err := client.DeleteVulScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DetachApplication2Connector is deprecated, please use csas::2023-01-20::ModifyForwardStrategy instead.
//
// Summary:
//
// Detaches private network access applications from a Connector in your Alibaba Cloud account.
//
// @param tmpReq - DetachApplication2ConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachApplication2ConnectorResponse
func (client *Client) DetachApplication2ConnectorWithOptions(tmpReq *DetachApplication2ConnectorRequest, runtime *dara.RuntimeOptions) (_result *DetachApplication2ConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetachApplication2ConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApplicationIds) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, dara.String("ApplicationIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIdsShrink) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachApplication2Connector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachApplication2ConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DetachApplication2Connector is deprecated, please use csas::2023-01-20::ModifyForwardStrategy instead.
//
// Summary:
//
// Detaches private network access applications from a Connector in your Alibaba Cloud account.
//
// @param request - DetachApplication2ConnectorRequest
//
// @return DetachApplication2ConnectorResponse
// Deprecated
func (client *Client) DetachApplication2Connector(request *DetachApplication2ConnectorRequest) (_result *DetachApplication2ConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachApplication2ConnectorResponse{}
	_body, _err := client.DetachApplication2ConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detach a business policy from an approval process.
//
// @param request - DetachPolicy2ApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachPolicy2ApprovalProcessResponse
func (client *Client) DetachPolicy2ApprovalProcessWithOptions(request *DetachPolicy2ApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *DetachPolicy2ApprovalProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.ProcessId) {
		body["ProcessId"] = request.ProcessId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetachPolicy2ApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetachPolicy2ApprovalProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detach a business policy from an approval process.
//
// @param request - DetachPolicy2ApprovalProcessRequest
//
// @return DetachPolicy2ApprovalProcessResponse
func (client *Client) DetachPolicy2ApprovalProcess(request *DetachPolicy2ApprovalProcessRequest) (_result *DetachPolicy2ApprovalProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetachPolicy2ApprovalProcessResponse{}
	_body, _err := client.DetachPolicy2ApprovalProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables an enterprise acceleration policy.
//
// @param request - DisableEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableEnterpriseAcceleratePolicyResponse
func (client *Client) DisableEnterpriseAcceleratePolicyWithOptions(request *DisableEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *DisableEnterpriseAcceleratePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables an enterprise acceleration policy.
//
// @param request - DisableEnterpriseAcceleratePolicyRequest
//
// @return DisableEnterpriseAcceleratePolicyResponse
func (client *Client) DisableEnterpriseAcceleratePolicy(request *DisableEnterpriseAcceleratePolicyRequest) (_result *DisableEnterpriseAcceleratePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.DisableEnterpriseAcceleratePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables an enterprise acceleration policy.
//
// @param request - EnableEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableEnterpriseAcceleratePolicyResponse
func (client *Client) EnableEnterpriseAcceleratePolicyWithOptions(request *EnableEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *EnableEnterpriseAcceleratePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables an enterprise acceleration policy.
//
// @param request - EnableEnterpriseAcceleratePolicyRequest
//
// @return EnableEnterpriseAcceleratePolicyResponse
func (client *Client) EnableEnterpriseAcceleratePolicy(request *EnableEnterpriseAcceleratePolicyRequest) (_result *EnableEnterpriseAcceleratePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.EnableEnterpriseAcceleratePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Export a list of user terminal devices to an Excel file.
//
// @param request - ExportUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportUserDevicesResponse
func (client *Client) ExportUserDevicesWithOptions(request *ExportUserDevicesRequest, runtime *dara.RuntimeOptions) (_result *ExportUserDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AppStatuses) {
		bodyFlat["AppStatuses"] = request.AppStatuses
	}

	if !dara.IsNil(request.Department) {
		body["Department"] = request.Department
	}

	if !dara.IsNil(request.DeviceBelong) {
		body["DeviceBelong"] = request.DeviceBelong
	}

	if !dara.IsNil(request.DeviceStatuses) {
		bodyFlat["DeviceStatuses"] = request.DeviceStatuses
	}

	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	if !dara.IsNil(request.DeviceTypes) {
		bodyFlat["DeviceTypes"] = request.DeviceTypes
	}

	if !dara.IsNil(request.DlpStatuses) {
		bodyFlat["DlpStatuses"] = request.DlpStatuses
	}

	if !dara.IsNil(request.Hostname) {
		body["Hostname"] = request.Hostname
	}

	if !dara.IsNil(request.IaStatuses) {
		bodyFlat["IaStatuses"] = request.IaStatuses
	}

	if !dara.IsNil(request.Mac) {
		body["Mac"] = request.Mac
	}

	if !dara.IsNil(request.NacStatuses) {
		bodyFlat["NacStatuses"] = request.NacStatuses
	}

	if !dara.IsNil(request.PaStatuses) {
		bodyFlat["PaStatuses"] = request.PaStatuses
	}

	if !dara.IsNil(request.SaseUserId) {
		body["SaseUserId"] = request.SaseUserId
	}

	if !dara.IsNil(request.SharingStatus) {
		body["SharingStatus"] = request.SharingStatus
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportUserDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Export a list of user terminal devices to an Excel file.
//
// @param request - ExportUserDevicesRequest
//
// @return ExportUserDevicesResponse
func (client *Client) ExportUserDevices(request *ExportUserDevicesRequest) (_result *ExportUserDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportUserDevicesResponse{}
	_body, _err := client.ExportUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the identity provider configuration enabled for your Alibaba Cloud account.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActiveIdpConfigResponse
func (client *Client) GetActiveIdpConfigWithOptions(runtime *dara.RuntimeOptions) (_result *GetActiveIdpConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetActiveIdpConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetActiveIdpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the identity provider configuration enabled for your Alibaba Cloud account.
//
// @return GetActiveIdpConfigResponse
func (client *Client) GetActiveIdpConfig() (_result *GetActiveIdpConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetActiveIdpConfigResponse{}
	_body, _err := client.GetActiveIdpConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the real-time antivirus defense policy of the current Alibaba Cloud account.
//
// @param request - GetAntiVirusRealTimeDefenceStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAntiVirusRealTimeDefenceStrategyResponse
func (client *Client) GetAntiVirusRealTimeDefenceStrategyWithOptions(request *GetAntiVirusRealTimeDefenceStrategyRequest, runtime *dara.RuntimeOptions) (_result *GetAntiVirusRealTimeDefenceStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetAntiVirusRealTimeDefenceStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAntiVirusRealTimeDefenceStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the real-time antivirus defense policy of the current Alibaba Cloud account.
//
// @param request - GetAntiVirusRealTimeDefenceStrategyRequest
//
// @return GetAntiVirusRealTimeDefenceStrategyResponse
func (client *Client) GetAntiVirusRealTimeDefenceStrategy(request *GetAntiVirusRealTimeDefenceStrategyRequest) (_result *GetAntiVirusRealTimeDefenceStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAntiVirusRealTimeDefenceStrategyResponse{}
	_body, _err := client.GetAntiVirusRealTimeDefenceStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an approval instance under the current Alibaba Cloud account.
//
// @param request - GetApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalResponse
func (client *Client) GetApprovalWithOptions(request *GetApprovalRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApproval"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an approval instance under the current Alibaba Cloud account.
//
// @param request - GetApprovalRequest
//
// @return GetApprovalResponse
func (client *Client) GetApproval(request *GetApprovalRequest) (_result *GetApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApprovalResponse{}
	_body, _err := client.GetApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an approval process under the current Alibaba Cloud account.
//
// @param request - GetApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalProcessResponse
func (client *Client) GetApprovalProcessWithOptions(request *GetApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an approval process under the current Alibaba Cloud account.
//
// @param request - GetApprovalProcessRequest
//
// @return GetApprovalProcessResponse
func (client *Client) GetApprovalProcess(request *GetApprovalProcessRequest) (_result *GetApprovalProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApprovalProcessResponse{}
	_body, _err := client.GetApprovalProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an approval template for your Alibaba Cloud account.
//
// @param request - GetApprovalSchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalSchemaResponse
func (client *Client) GetApprovalSchemaWithOptions(request *GetApprovalSchemaRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalSchemaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApprovalSchema"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalSchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an approval template for your Alibaba Cloud account.
//
// @param request - GetApprovalSchemaRequest
//
// @return GetApprovalSchemaResponse
func (client *Client) GetApprovalSchema(request *GetApprovalSchemaRequest) (_result *GetApprovalSchemaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApprovalSchemaResponse{}
	_body, _err := client.GetApprovalSchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the auto-start and anti-uninstall policy under the current Alibaba Cloud account.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBootAndAntiUninstallPolicyResponse
func (client *Client) GetBootAndAntiUninstallPolicyWithOptions(runtime *dara.RuntimeOptions) (_result *GetBootAndAntiUninstallPolicyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetBootAndAntiUninstallPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBootAndAntiUninstallPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the auto-start and anti-uninstall policy under the current Alibaba Cloud account.
//
// @return GetBootAndAntiUninstallPolicyResponse
func (client *Client) GetBootAndAntiUninstallPolicy() (_result *GetBootAndAntiUninstallPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBootAndAntiUninstallPolicyResponse{}
	_body, _err := client.GetBootAndAntiUninstallPolicyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user from a custom identity source in your Alibaba Cloud account.
//
// @param request - GetClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientUserResponse
func (client *Client) GetClientUserWithOptions(request *GetClientUserRequest, runtime *dara.RuntimeOptions) (_result *GetClientUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a user from a custom identity source in your Alibaba Cloud account.
//
// @param request - GetClientUserRequest
//
// @return GetClientUserResponse
func (client *Client) GetClientUser(request *GetClientUserRequest) (_result *GetClientUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClientUserResponse{}
	_body, _err := client.GetClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a connector.
//
// @param request - GetConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectorResponse
func (client *Client) GetConnectorWithOptions(request *GetConnectorRequest, runtime *dara.RuntimeOptions) (_result *GetConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a connector.
//
// @param request - GetConnectorRequest
//
// @return GetConnectorResponse
func (client *Client) GetConnector(request *GetConnectorRequest) (_result *GetConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConnectorResponse{}
	_body, _err := client.GetConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a ConnectorClient.
//
// @param request - GetConnectorClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConnectorClientResponse
func (client *Client) GetConnectorClientWithOptions(request *GetConnectorClientRequest, runtime *dara.RuntimeOptions) (_result *GetConnectorClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetConnectorClient"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConnectorClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a ConnectorClient.
//
// @param request - GetConnectorClientRequest
//
// @return GetConnectorClientResponse
func (client *Client) GetConnectorClient(request *GetConnectorClientRequest) (_result *GetConnectorClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetConnectorClientResponse{}
	_body, _err := client.GetConnectorClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified device label.
//
// @param request - GetDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceGroupResponse
func (client *Client) GetDeviceGroupWithOptions(request *GetDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified device label.
//
// @param request - GetDeviceGroupRequest
//
// @return GetDeviceGroupResponse
func (client *Client) GetDeviceGroup(request *GetDeviceGroupRequest) (_result *GetDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceGroupResponse{}
	_body, _err := client.GetDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the online time distribution of a specified terminal device on a specified date, aggregated by minute.
//
// @param request - GetDeviceOnlineHeatmapRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceOnlineHeatmapResponse
func (client *Client) GetDeviceOnlineHeatmapWithOptions(request *GetDeviceOnlineHeatmapRequest, runtime *dara.RuntimeOptions) (_result *GetDeviceOnlineHeatmapResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.DevTag) {
		query["DevTag"] = request.DevTag
	}

	if !dara.IsNil(request.SaseUserId) {
		query["SaseUserId"] = request.SaseUserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceOnlineHeatmap"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceOnlineHeatmapResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the online time distribution of a specified terminal device on a specified date, aggregated by minute.
//
// @param request - GetDeviceOnlineHeatmapRequest
//
// @return GetDeviceOnlineHeatmapResponse
func (client *Client) GetDeviceOnlineHeatmap(request *GetDeviceOnlineHeatmapRequest) (_result *GetDeviceOnlineHeatmapResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeviceOnlineHeatmapResponse{}
	_body, _err := client.GetDeviceOnlineHeatmapWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves details about a dynamic route in your Alibaba Cloud account.
//
// @param request - GetDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDynamicRouteResponse
func (client *Client) GetDynamicRouteWithOptions(request *GetDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *GetDynamicRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details about a dynamic route in your Alibaba Cloud account.
//
// @param request - GetDynamicRouteRequest
//
// @return GetDynamicRouteResponse
func (client *Client) GetDynamicRoute(request *GetDynamicRouteRequest) (_result *GetDynamicRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDynamicRouteResponse{}
	_body, _err := client.GetDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a forwarding rule.
//
// Description:
//
// Creates a domain name list of a specified type (blacklist/whitelist) under the current tenant and returns the ListId of the new list. You can create up to 100 lists of each type per tenant.
//
// @param request - GetForwardStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetForwardStrategyResponse
func (client *Client) GetForwardStrategyWithOptions(request *GetForwardStrategyRequest, runtime *dara.RuntimeOptions) (_result *GetForwardStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetForwardStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetForwardStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a forwarding rule.
//
// Description:
//
// Creates a domain name list of a specified type (blacklist/whitelist) under the current tenant and returns the ListId of the new list. You can create up to 100 lists of each type per tenant.
//
// @param request - GetForwardStrategyRequest
//
// @return GetForwardStrategyResponse
func (client *Client) GetForwardStrategy(request *GetForwardStrategyRequest) (_result *GetForwardStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetForwardStrategyResponse{}
	_body, _err := client.GetForwardStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve the configuration details of a specified identity provider for your Alibaba Cloud account.
//
// @param request - GetIdpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdpConfigResponse
func (client *Client) GetIdpConfigWithOptions(request *GetIdpConfigRequest, runtime *dara.RuntimeOptions) (_result *GetIdpConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIdpConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIdpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve the configuration details of a specified identity provider for your Alibaba Cloud account.
//
// @param request - GetIdpConfigRequest
//
// @return GetIdpConfigResponse
func (client *Client) GetIdpConfig(request *GetIdpConfigRequest) (_result *GetIdpConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIdpConfigResponse{}
	_body, _err := client.GetIdpConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the phone number whitelist for visitor admission SMS logon.
//
// Description:
//
// Retrieves all phone numbers in the whitelist.
//
// @param request - GetNacPortalSmsPhoneWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNacPortalSmsPhoneWhitelistResponse
func (client *Client) GetNacPortalSmsPhoneWhitelistWithOptions(request *GetNacPortalSmsPhoneWhitelistRequest, runtime *dara.RuntimeOptions) (_result *GetNacPortalSmsPhoneWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetNacPortalSmsPhoneWhitelist"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNacPortalSmsPhoneWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the phone number whitelist for visitor admission SMS logon.
//
// Description:
//
// Retrieves all phone numbers in the whitelist.
//
// @param request - GetNacPortalSmsPhoneWhitelistRequest
//
// @return GetNacPortalSmsPhoneWhitelistResponse
func (client *Client) GetNacPortalSmsPhoneWhitelist(request *GetNacPortalSmsPhoneWhitelistRequest) (_result *GetNacPortalSmsPhoneWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNacPortalSmsPhoneWhitelistResponse{}
	_body, _err := client.GetNacPortalSmsPhoneWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a private access diagnostic task.
//
// @param request - GetPADiagnosisTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPADiagnosisTaskResponse
func (client *Client) GetPADiagnosisTaskWithOptions(request *GetPADiagnosisTaskRequest, runtime *dara.RuntimeOptions) (_result *GetPADiagnosisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPADiagnosisTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPADiagnosisTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a private access diagnostic task.
//
// @param request - GetPADiagnosisTaskRequest
//
// @return GetPADiagnosisTaskResponse
func (client *Client) GetPADiagnosisTask(request *GetPADiagnosisTaskRequest) (_result *GetPADiagnosisTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPADiagnosisTaskResponse{}
	_body, _err := client.GetPADiagnosisTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an internal-facing access application under the current Alibaba Cloud account.
//
// @param request - GetPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateAccessApplicationResponse
func (client *Client) GetPrivateAccessApplicationWithOptions(request *GetPrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetPrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an internal-facing access application under the current Alibaba Cloud account.
//
// @param request - GetPrivateAccessApplicationRequest
//
// @return GetPrivateAccessApplicationResponse
func (client *Client) GetPrivateAccessApplication(request *GetPrivateAccessApplicationRequest) (_result *GetPrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPrivateAccessApplicationResponse{}
	_body, _err := client.GetPrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Intranet Access Policy Details
//
// @param request - GetPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateAccessPolicyResponse
func (client *Client) GetPrivateAccessPolicyWithOptions(request *GetPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetPrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Intranet Access Policy Details
//
// @param request - GetPrivateAccessPolicyRequest
//
// @return GetPrivateAccessPolicyResponse
func (client *Client) GetPrivateAccessPolicy(request *GetPrivateAccessPolicyRequest) (_result *GetPrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPrivateAccessPolicyResponse{}
	_body, _err := client.GetPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified software prohibition policy.
//
// @param request - GetProhibitedPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProhibitedPolicyResponse
func (client *Client) GetProhibitedPolicyWithOptions(request *GetProhibitedPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetProhibitedPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProhibitedPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProhibitedPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified software prohibition policy.
//
// @param request - GetProhibitedPolicyRequest
//
// @return GetProhibitedPolicyResponse
func (client *Client) GetProhibitedPolicy(request *GetProhibitedPolicyRequest) (_result *GetProhibitedPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProhibitedPolicyResponse{}
	_body, _err := client.GetProhibitedPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified prohibited software.
//
// @param tmpReq - GetProhibitedSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProhibitedSoftwareResponse
func (client *Client) GetProhibitedSoftwareWithOptions(tmpReq *GetProhibitedSoftwareRequest, runtime *dara.RuntimeOptions) (_result *GetProhibitedSoftwareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetProhibitedSoftwareShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SoftwareId) {
		request.SoftwareIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SoftwareId, dara.String("SoftwareId"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProhibitedSoftware"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProhibitedSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified prohibited software.
//
// @param request - GetProhibitedSoftwareRequest
//
// @return GetProhibitedSoftwareResponse
func (client *Client) GetProhibitedSoftware(request *GetProhibitedSoftwareRequest) (_result *GetProhibitedSoftwareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProhibitedSoftwareResponse{}
	_body, _err := client.GetProhibitedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a device registration policy within the current Alibaba Cloud account.
//
// @param request - GetRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistrationPolicyResponse
func (client *Client) GetRegistrationPolicyWithOptions(request *GetRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetRegistrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a device registration policy within the current Alibaba Cloud account.
//
// @param request - GetRegistrationPolicyRequest
//
// @return GetRegistrationPolicyResponse
func (client *Client) GetRegistrationPolicy(request *GetRegistrationPolicyRequest) (_result *GetRegistrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRegistrationPolicyResponse{}
	_body, _err := client.GetRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a user endpoint device under the current Alibaba Cloud account.
//
// @param request - GetUserDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDeviceResponse
func (client *Client) GetUserDeviceWithOptions(request *GetUserDeviceRequest, runtime *dara.RuntimeOptions) (_result *GetUserDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserDevice"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a user endpoint device under the current Alibaba Cloud account.
//
// @param request - GetUserDeviceRequest
//
// @return GetUserDeviceResponse
func (client *Client) GetUserDevice(request *GetUserDeviceRequest) (_result *GetUserDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserDeviceResponse{}
	_body, _err := client.GetUserDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the workload usage trends of a specified endpoint device.
//
// @param request - GetUserDeviceWorkloadTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDeviceWorkloadTrendResponse
func (client *Client) GetUserDeviceWorkloadTrendWithOptions(request *GetUserDeviceWorkloadTrendRequest, runtime *dara.RuntimeOptions) (_result *GetUserDeviceWorkloadTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTag) {
		query["DeviceTag"] = request.DeviceTag
	}

	if !dara.IsNil(request.From) {
		query["From"] = request.From
	}

	if !dara.IsNil(request.To) {
		query["To"] = request.To
	}

	if !dara.IsNil(request.WorkloadType) {
		query["WorkloadType"] = request.WorkloadType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserDeviceWorkloadTrend"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserDeviceWorkloadTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the workload usage trends of a specified endpoint device.
//
// @param request - GetUserDeviceWorkloadTrendRequest
//
// @return GetUserDeviceWorkloadTrendResponse
func (client *Client) GetUserDeviceWorkloadTrend(request *GetUserDeviceWorkloadTrendRequest) (_result *GetUserDeviceWorkloadTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserDeviceWorkloadTrendResponse{}
	_body, _err := client.GetUserDeviceWorkloadTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a user group in the current Alibaba Cloud account.
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *dara.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a user group in the current Alibaba Cloud account.
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global anti-virus configuration of the current Alibaba Cloud account, including the virus file upload switch and upload limits. If the current Alibaba Cloud account does not have its own configuration record, the default configurations are returned.
//
// @param request - GetVirusScanGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVirusScanGlobalConfigResponse
func (client *Client) GetVirusScanGlobalConfigWithOptions(request *GetVirusScanGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *GetVirusScanGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetVirusScanGlobalConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVirusScanGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global anti-virus configuration of the current Alibaba Cloud account, including the virus file upload switch and upload limits. If the current Alibaba Cloud account does not have its own configuration record, the default configurations are returned.
//
// @param request - GetVirusScanGlobalConfigRequest
//
// @return GetVirusScanGlobalConfigResponse
func (client *Client) GetVirusScanGlobalConfig(request *GetVirusScanGlobalConfigRequest) (_result *GetVirusScanGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVirusScanGlobalConfigResponse{}
	_body, _err := client.GetVirusScanGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified scheduled virus scan policy.
//
// @param request - GetVirusScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVirusScanScheduledStrategyResponse
func (client *Client) GetVirusScanScheduledStrategyWithOptions(request *GetVirusScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *GetVirusScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVirusScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVirusScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified scheduled virus scan policy.
//
// @param request - GetVirusScanScheduledStrategyRequest
//
// @return GetVirusScanScheduledStrategyResponse
func (client *Client) GetVirusScanScheduledStrategy(request *GetVirusScanScheduledStrategyRequest) (_result *GetVirusScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVirusScanScheduledStrategyResponse{}
	_body, _err := client.GetVirusScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the global configuration of vulnerability scanning for the current Alibaba Cloud account.
//
// @param request - GetVulScanGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulScanGlobalConfigResponse
func (client *Client) GetVulScanGlobalConfigWithOptions(request *GetVulScanGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *GetVulScanGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetVulScanGlobalConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVulScanGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the global configuration of vulnerability scanning for the current Alibaba Cloud account.
//
// @param request - GetVulScanGlobalConfigRequest
//
// @return GetVulScanGlobalConfigResponse
func (client *Client) GetVulScanGlobalConfig(request *GetVulScanGlobalConfigRequest) (_result *GetVulScanGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVulScanGlobalConfigResponse{}
	_body, _err := client.GetVulScanGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the complete configuration of a specified vulnerability scheduled scan policy.
//
// @param request - GetVulScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulScanScheduledStrategyResponse
func (client *Client) GetVulScanScheduledStrategyWithOptions(request *GetVulScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *GetVulScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVulScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVulScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the complete configuration of a specified vulnerability scheduled scan policy.
//
// @param request - GetVulScanScheduledStrategyRequest
//
// @return GetVulScanScheduledStrategyResponse
func (client *Client) GetVulScanScheduledStrategy(request *GetVulScanScheduledStrategyRequest) (_result *GetVulScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVulScanScheduledStrategyResponse{}
	_body, _err := client.GetVulScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified vulnerability.
//
// @param request - GetVulnerabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVulnerabilityResponse
func (client *Client) GetVulnerabilityWithOptions(request *GetVulnerabilityRequest, runtime *dara.RuntimeOptions) (_result *GetVulnerabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVulnerability"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVulnerabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified vulnerability.
//
// @param request - GetVulnerabilityRequest
//
// @return GetVulnerabilityResponse
func (client *Client) GetVulnerability(request *GetVulnerabilityRequest) (_result *GetVulnerabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVulnerabilityResponse{}
	_body, _err := client.GetVulnerabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Use the job ID obtained from creating a watermark embedding job to query the embedding job result.
//
// @param request - GetWmEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWmEmbedTaskResponse
func (client *Client) GetWmEmbedTaskWithOptions(request *GetWmEmbedTaskRequest, runtime *dara.RuntimeOptions) (_result *GetWmEmbedTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWmEmbedTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWmEmbedTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use the job ID obtained from creating a watermark embedding job to query the embedding job result.
//
// @param request - GetWmEmbedTaskRequest
//
// @return GetWmEmbedTaskResponse
func (client *Client) GetWmEmbedTask(request *GetWmEmbedTaskRequest) (_result *GetWmEmbedTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWmEmbedTaskResponse{}
	_body, _err := client.GetWmEmbedTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of a watermark extraction task using the task ID that is returned when you create the task.
//
// @param request - GetWmExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWmExtractTaskResponse
func (client *Client) GetWmExtractTaskWithOptions(request *GetWmExtractTaskRequest, runtime *dara.RuntimeOptions) (_result *GetWmExtractTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWmExtractTask"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWmExtractTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of a watermark extraction task using the task ID that is returned when you create the task.
//
// @param request - GetWmExtractTaskRequest
//
// @return GetWmExtractTaskResponse
func (client *Client) GetWmExtractTask(request *GetWmExtractTaskRequest) (_result *GetWmExtractTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWmExtractTaskResponse{}
	_body, _err := client.GetWmExtractTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports acceleration addresses in batches.
//
// @param request - ImportEnterpriseAccelerateTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportEnterpriseAccelerateTargetsResponse
func (client *Client) ImportEnterpriseAccelerateTargetsWithOptions(request *ImportEnterpriseAccelerateTargetsRequest, runtime *dara.RuntimeOptions) (_result *ImportEnterpriseAccelerateTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	if !dara.IsNil(request.FileUrl) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportEnterpriseAccelerateTargets"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportEnterpriseAccelerateTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports acceleration addresses in batches.
//
// @param request - ImportEnterpriseAccelerateTargetsRequest
//
// @return ImportEnterpriseAccelerateTargetsResponse
func (client *Client) ImportEnterpriseAccelerateTargets(request *ImportEnterpriseAccelerateTargetsRequest) (_result *ImportEnterpriseAccelerateTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportEnterpriseAccelerateTargetsResponse{}
	_body, _err := client.ImportEnterpriseAccelerateTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the applications associated with one or more private access policies.
//
// @param request - ListApplicationsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForPrivateAccessPolicyResponse
func (client *Client) ListApplicationsForPrivateAccessPolicyWithOptions(request *ListApplicationsForPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForPrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the applications associated with one or more private access policies.
//
// @param request - ListApplicationsForPrivateAccessPolicyRequest
//
// @return ListApplicationsForPrivateAccessPolicyResponse
func (client *Client) ListApplicationsForPrivateAccessPolicy(request *ListApplicationsForPrivateAccessPolicyRequest) (_result *ListApplicationsForPrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsForPrivateAccessPolicyResponse{}
	_body, _err := client.ListApplicationsForPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch-query applications associated with private network access tags within your Alibaba Cloud account.
//
// @param request - ListApplicationsForPrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForPrivateAccessTagResponse
func (client *Client) ListApplicationsForPrivateAccessTagWithOptions(request *ListApplicationsForPrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *ListApplicationsForPrivateAccessTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApplicationsForPrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApplicationsForPrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch-query applications associated with private network access tags within your Alibaba Cloud account.
//
// @param request - ListApplicationsForPrivateAccessTagRequest
//
// @return ListApplicationsForPrivateAccessTagResponse
func (client *Client) ListApplicationsForPrivateAccessTag(request *ListApplicationsForPrivateAccessTagRequest) (_result *ListApplicationsForPrivateAccessTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApplicationsForPrivateAccessTagResponse{}
	_body, _err := client.ListApplicationsForPrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of approval processes under the current Alibaba Cloud account.
//
// @param request - ListApprovalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalProcessesResponse
func (client *Client) ListApprovalProcessesWithOptions(request *ListApprovalProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalProcessesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of approval processes under the current Alibaba Cloud account.
//
// @param request - ListApprovalProcessesRequest
//
// @return ListApprovalProcessesResponse
func (client *Client) ListApprovalProcesses(request *ListApprovalProcessesRequest) (_result *ListApprovalProcessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApprovalProcessesResponse{}
	_body, _err := client.ListApprovalProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the approval flows associated with approval rendering templates.
//
// @param request - ListApprovalProcessesForApprovalSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalProcessesForApprovalSchemasResponse
func (client *Client) ListApprovalProcessesForApprovalSchemasWithOptions(request *ListApprovalProcessesForApprovalSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalProcessesForApprovalSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalProcessesForApprovalSchemas"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalProcessesForApprovalSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the approval flows associated with approval rendering templates.
//
// @param request - ListApprovalProcessesForApprovalSchemasRequest
//
// @return ListApprovalProcessesForApprovalSchemasResponse
func (client *Client) ListApprovalProcessesForApprovalSchemas(request *ListApprovalProcessesForApprovalSchemasRequest) (_result *ListApprovalProcessesForApprovalSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApprovalProcessesForApprovalSchemasResponse{}
	_body, _err := client.ListApprovalProcessesForApprovalSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the approval templates for your Alibaba Cloud account.
//
// @param request - ListApprovalSchemasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalSchemasResponse
func (client *Client) ListApprovalSchemasWithOptions(request *ListApprovalSchemasRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalSchemasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalSchemas"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the approval templates for your Alibaba Cloud account.
//
// @param request - ListApprovalSchemasRequest
//
// @return ListApprovalSchemasResponse
func (client *Client) ListApprovalSchemas(request *ListApprovalSchemasRequest) (_result *ListApprovalSchemasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApprovalSchemasResponse{}
	_body, _err := client.ListApprovalSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the rendering templates associated with approval processes.
//
// @param request - ListApprovalSchemasForApprovalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalSchemasForApprovalProcessesResponse
func (client *Client) ListApprovalSchemasForApprovalProcessesWithOptions(request *ListApprovalSchemasForApprovalProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalSchemasForApprovalProcessesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovalSchemasForApprovalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalSchemasForApprovalProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the rendering templates associated with approval processes.
//
// @param request - ListApprovalSchemasForApprovalProcessesRequest
//
// @return ListApprovalSchemasForApprovalProcessesResponse
func (client *Client) ListApprovalSchemasForApprovalProcesses(request *ListApprovalSchemasForApprovalProcessesRequest) (_result *ListApprovalSchemasForApprovalProcessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApprovalSchemasForApprovalProcessesResponse{}
	_body, _err := client.ListApprovalSchemasForApprovalProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of approval instances under the current Alibaba Cloud account.
//
// @param request - ListApprovalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApprovalsResponse
func (client *Client) ListApprovalsWithOptions(request *ListApprovalsRequest, runtime *dara.RuntimeOptions) (_result *ListApprovalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApprovals"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApprovalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of approval instances under the current Alibaba Cloud account.
//
// @param request - ListApprovalsRequest
//
// @return ListApprovalsResponse
func (client *Client) ListApprovals(request *ListApprovalsRequest) (_result *ListApprovalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListApprovalsResponse{}
	_body, _err := client.ListApprovalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query user information from custom identity sources in your Alibaba Cloud account.
//
// @param request - ListClientUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientUsersResponse
func (client *Client) ListClientUsersWithOptions(request *ListClientUsersRequest, runtime *dara.RuntimeOptions) (_result *ListClientUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClientUsers"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClientUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query user information from custom identity sources in your Alibaba Cloud account.
//
// @param request - ListClientUsersRequest
//
// @return ListClientUsersResponse
func (client *Client) ListClientUsers(request *ListClientUsersRequest) (_result *ListClientUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClientUsersResponse{}
	_body, _err := client.ListClientUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query connectors in batches.
//
// @param request - ListConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectorsResponse
func (client *Client) ListConnectorsWithOptions(request *ListConnectorsRequest, runtime *dara.RuntimeOptions) (_result *ListConnectorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnectors"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query connectors in batches.
//
// @param request - ListConnectorsRequest
//
// @return ListConnectorsResponse
func (client *Client) ListConnectors(request *ListConnectorsRequest) (_result *ListConnectorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConnectorsResponse{}
	_body, _err := client.ListConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of device groups under the current Alibaba Cloud account by using paging.
//
// @param request - ListDeviceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeviceGroupsResponse
func (client *Client) ListDeviceGroupsWithOptions(request *ListDeviceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListDeviceGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DeviceGroupIds) {
		query["DeviceGroupIds"] = request.DeviceGroupIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeviceGroups"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeviceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of device groups under the current Alibaba Cloud account by using paging.
//
// @param request - ListDeviceGroupsRequest
//
// @return ListDeviceGroupsResponse
func (client *Client) ListDeviceGroups(request *ListDeviceGroupsRequest) (_result *ListDeviceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeviceGroupsResponse{}
	_body, _err := client.ListDeviceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user endpoint devices affected by a specified vulnerability and their remediation status by paging.
//
// @param request - ListDevicesForVulnerabilityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDevicesForVulnerabilityResponse
func (client *Client) ListDevicesForVulnerabilityWithOptions(request *ListDevicesForVulnerabilityRequest, runtime *dara.RuntimeOptions) (_result *ListDevicesForVulnerabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDevicesForVulnerability"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDevicesForVulnerabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user endpoint devices affected by a specified vulnerability and their remediation status by paging.
//
// @param request - ListDevicesForVulnerabilityRequest
//
// @return ListDevicesForVulnerabilityResponse
func (client *Client) ListDevicesForVulnerability(request *ListDevicesForVulnerabilityRequest) (_result *ListDevicesForVulnerabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDevicesForVulnerabilityResponse{}
	_body, _err := client.ListDevicesForVulnerabilityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries domain name entries in a domain name list by paging.
//
// Description:
//
// Queries the details of domain name entries in a specified domain name list by paging. Use this operation together with ListDomainMetas: first obtain the `ListId`, and then use this operation to perform paging through the domain names in the list.
//
// @param request - ListDomainItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainItemsResponse
func (client *Client) ListDomainItemsWithOptions(request *ListDomainItemsRequest, runtime *dara.RuntimeOptions) (_result *ListDomainItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.ItemValue) {
		query["ItemValue"] = request.ItemValue
	}

	if !dara.IsNil(request.ListId) {
		query["ListId"] = request.ListId
	}

	if !dara.IsNil(request.ListType) {
		query["ListType"] = request.ListType
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomainItems"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries domain name entries in a domain name list by paging.
//
// Description:
//
// Queries the details of domain name entries in a specified domain name list by paging. Use this operation together with ListDomainMetas: first obtain the `ListId`, and then use this operation to perform paging through the domain names in the list.
//
// @param request - ListDomainItemsRequest
//
// @return ListDomainItemsResponse
func (client *Client) ListDomainItems(request *ListDomainItemsRequest) (_result *ListDomainItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainItemsResponse{}
	_body, _err := client.ListDomainItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of domain name lists.
//
// Description:
//
// Performs a paged query on the metadata of domain name lists (the header information of domain name blacklists/whitelists, excluding the specific domain name entries within the lists) for the current tenant with paging. You can filter by list type (blacklist/whitelist), perform fuzzy search by name, and specify whether to include system built-in default template lists in the results. Each record includes the number of domain name entries in the list.
//
// @param request - ListDomainMetasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainMetasResponse
func (client *Client) ListDomainMetasWithOptions(request *ListDomainMetasRequest, runtime *dara.RuntimeOptions) (_result *ListDomainMetasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DefaultTemplate) {
		query["DefaultTemplate"] = request.DefaultTemplate
	}

	if !dara.IsNil(request.ListType) {
		query["ListType"] = request.ListType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDomainMetas"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDomainMetasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of domain name lists.
//
// Description:
//
// Performs a paged query on the metadata of domain name lists (the header information of domain name blacklists/whitelists, excluding the specific domain name entries within the lists) for the current tenant with paging. You can filter by list type (blacklist/whitelist), perform fuzzy search by name, and specify whether to include system built-in default template lists in the results. Each record includes the number of domain name entries in the list.
//
// @param request - ListDomainMetasRequest
//
// @return ListDomainMetasResponse
func (client *Client) ListDomainMetas(request *ListDomainMetasRequest) (_result *ListDomainMetasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDomainMetasResponse{}
	_body, _err := client.ListDomainMetasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Query Dynamic Policy Disposal Processes
//
// @param request - ListDynamicDisposalProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicDisposalProcessesResponse
func (client *Client) ListDynamicDisposalProcessesWithOptions(request *ListDynamicDisposalProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicDisposalProcessesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicDisposalProcesses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicDisposalProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Query Dynamic Policy Disposal Processes
//
// @param request - ListDynamicDisposalProcessesRequest
//
// @return ListDynamicDisposalProcessesResponse
func (client *Client) ListDynamicDisposalProcesses(request *ListDynamicDisposalProcessesRequest) (_result *ListDynamicDisposalProcessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDynamicDisposalProcessesResponse{}
	_body, _err := client.ListDynamicDisposalProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the regions that are supported by the access points of SASE POP clusters for your Alibaba Cloud account.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicRouteRegionsResponse
func (client *Client) ListDynamicRouteRegionsWithOptions(runtime *dara.RuntimeOptions) (_result *ListDynamicRouteRegionsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicRouteRegions"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicRouteRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the regions that are supported by the access points of SASE POP clusters for your Alibaba Cloud account.
//
// @return ListDynamicRouteRegionsResponse
func (client *Client) ListDynamicRouteRegions() (_result *ListDynamicRouteRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDynamicRouteRegionsResponse{}
	_body, _err := client.ListDynamicRouteRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about all dynamic routes for the current Alibaba Cloud account.
//
// @param request - ListDynamicRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicRoutesResponse
func (client *Client) ListDynamicRoutesWithOptions(request *ListDynamicRoutesRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicRoutesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicRoutes"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about all dynamic routes for the current Alibaba Cloud account.
//
// @param request - ListDynamicRoutesRequest
//
// @return ListDynamicRoutesResponse
func (client *Client) ListDynamicRoutes(request *ListDynamicRoutesRequest) (_result *ListDynamicRoutesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDynamicRoutesResponse{}
	_body, _err := client.ListDynamicRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries audit logs for enterprise acceleration.
//
// @param request - ListEnterpriseAccelerateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseAccelerateLogsResponse
func (client *Client) ListEnterpriseAccelerateLogsWithOptions(request *ListEnterpriseAccelerateLogsRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseAccelerateLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseAccelerateLogs"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseAccelerateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries audit logs for enterprise acceleration.
//
// @param request - ListEnterpriseAccelerateLogsRequest
//
// @return ListEnterpriseAccelerateLogsResponse
func (client *Client) ListEnterpriseAccelerateLogs(request *ListEnterpriseAccelerateLogsRequest) (_result *ListEnterpriseAccelerateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnterpriseAccelerateLogsResponse{}
	_body, _err := client.ListEnterpriseAccelerateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query acceleration policies
//
// @param request - ListEnterpriseAcceleratePoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseAcceleratePoliciesResponse
func (client *Client) ListEnterpriseAcceleratePoliciesWithOptions(request *ListEnterpriseAcceleratePoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseAcceleratePoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseAcceleratePolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseAcceleratePoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query acceleration policies
//
// @param request - ListEnterpriseAcceleratePoliciesRequest
//
// @return ListEnterpriseAcceleratePoliciesResponse
func (client *Client) ListEnterpriseAcceleratePolicies(request *ListEnterpriseAcceleratePoliciesRequest) (_result *ListEnterpriseAcceleratePoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnterpriseAcceleratePoliciesResponse{}
	_body, _err := client.ListEnterpriseAcceleratePoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of enterprise acceleration targets.
//
// @param request - ListEnterpriseAccelerateTargetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnterpriseAccelerateTargetsResponse
func (client *Client) ListEnterpriseAccelerateTargetsWithOptions(request *ListEnterpriseAccelerateTargetsRequest, runtime *dara.RuntimeOptions) (_result *ListEnterpriseAccelerateTargetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnterpriseAccelerateTargets"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnterpriseAccelerateTargetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of enterprise acceleration targets.
//
// @param request - ListEnterpriseAccelerateTargetsRequest
//
// @return ListEnterpriseAccelerateTargetsResponse
func (client *Client) ListEnterpriseAccelerateTargets(request *ListEnterpriseAccelerateTargetsRequest) (_result *ListEnterpriseAccelerateTargetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnterpriseAccelerateTargetsResponse{}
	_body, _err := client.ListEnterpriseAccelerateTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists excess end-user device registration applications in the current Alibaba Cloud account.
//
// @param request - ListExcessiveDeviceRegistrationApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExcessiveDeviceRegistrationApplicationsResponse
func (client *Client) ListExcessiveDeviceRegistrationApplicationsWithOptions(request *ListExcessiveDeviceRegistrationApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListExcessiveDeviceRegistrationApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExcessiveDeviceRegistrationApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExcessiveDeviceRegistrationApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists excess end-user device registration applications in the current Alibaba Cloud account.
//
// @param request - ListExcessiveDeviceRegistrationApplicationsRequest
//
// @return ListExcessiveDeviceRegistrationApplicationsResponse
func (client *Client) ListExcessiveDeviceRegistrationApplications(request *ListExcessiveDeviceRegistrationApplicationsRequest) (_result *ListExcessiveDeviceRegistrationApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExcessiveDeviceRegistrationApplicationsResponse{}
	_body, _err := client.ListExcessiveDeviceRegistrationApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries forwarding rules in batches.
//
// @param request - ListForwardStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListForwardStrategiesResponse
func (client *Client) ListForwardStrategiesWithOptions(request *ListForwardStrategiesRequest, runtime *dara.RuntimeOptions) (_result *ListForwardStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListForwardStrategies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListForwardStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries forwarding rules in batches.
//
// @param request - ListForwardStrategiesRequest
//
// @return ListForwardStrategiesResponse
func (client *Client) ListForwardStrategies(request *ListForwardStrategiesRequest) (_result *ListForwardStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListForwardStrategiesResponse{}
	_body, _err := client.ListForwardStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the bindings of forwarding rules in batches.
//
// @param request - ListForwardStrategyBindingItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListForwardStrategyBindingItemsResponse
func (client *Client) ListForwardStrategyBindingItemsWithOptions(request *ListForwardStrategyBindingItemsRequest, runtime *dara.RuntimeOptions) (_result *ListForwardStrategyBindingItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListForwardStrategyBindingItems"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListForwardStrategyBindingItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the bindings of forwarding rules in batches.
//
// @param request - ListForwardStrategyBindingItemsRequest
//
// @return ListForwardStrategyBindingItemsResponse
func (client *Client) ListForwardStrategyBindingItems(request *ListForwardStrategyBindingItemsRequest) (_result *ListForwardStrategyBindingItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListForwardStrategyBindingItemsResponse{}
	_body, _err := client.ListForwardStrategyBindingItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists identity provider configurations for the current Alibaba Cloud account.
//
// @param request - ListIdpConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdpConfigsResponse
func (client *Client) ListIdpConfigsWithOptions(request *ListIdpConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListIdpConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdpConfigs"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdpConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists identity provider configurations for the current Alibaba Cloud account.
//
// @param request - ListIdpConfigsRequest
//
// @return ListIdpConfigsResponse
func (client *Client) ListIdpConfigs(request *ListIdpConfigsRequest) (_result *ListIdpConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIdpConfigsResponse{}
	_body, _err := client.ListIdpConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves department information for a custom identity provider (IdP) associated with your Alibaba Cloud account.
//
// @param request - ListIdpDepartmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdpDepartmentsResponse
func (client *Client) ListIdpDepartmentsWithOptions(request *ListIdpDepartmentsRequest, runtime *dara.RuntimeOptions) (_result *ListIdpDepartmentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdpDepartments"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdpDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves department information for a custom identity provider (IdP) associated with your Alibaba Cloud account.
//
// @param request - ListIdpDepartmentsRequest
//
// @return ListIdpDepartmentsResponse
func (client *Client) ListIdpDepartments(request *ListIdpDepartmentsRequest) (_result *ListIdpDepartmentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIdpDepartmentsResponse{}
	_body, _err := client.ListIdpDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists onboarded users.
//
// @param request - ListNacUserCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNacUserCertResponse
func (client *Client) ListNacUserCertWithOptions(request *ListNacUserCertRequest, runtime *dara.RuntimeOptions) (_result *ListNacUserCertResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Department) {
		query["Department"] = request.Department
	}

	if !dara.IsNil(request.DeviceType) {
		query["DeviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNacUserCert"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNacUserCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists onboarded users.
//
// @param request - ListNacUserCertRequest
//
// @return ListNacUserCertResponse
func (client *Client) ListNacUserCert(request *ListNacUserCertRequest) (_result *ListNacUserCertResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNacUserCertResponse{}
	_body, _err := client.ListNacUserCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries administrator operation audit logs in batches.
//
// Description:
//
// - StartTime and EndTime are UNIX timestamps in seconds. StartTime must be earlier than EndTime. The interval between them cannot exceed 30 days, and StartTime cannot be more than 31 days before the current time.
//
// - If OperationStatus is not specified, only successful operation records are returned.
//
// - Results are sorted by operation time in descending order.
//
// - The return values of OperationFunc, OperationPage, and OperationType are localized based on the request language.
//
// - The values of filter parameters cannot contain single quotation marks (\\"), double quotation marks ("), or backslashes (\\\\). Otherwise, an InvalidParameter error is returned.
//
// @param request - ListOperationAuditLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOperationAuditLogsResponse
func (client *Client) ListOperationAuditLogsWithOptions(request *ListOperationAuditLogsRequest, runtime *dara.RuntimeOptions) (_result *ListOperationAuditLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EventType) {
		query["EventType"] = request.EventType
	}

	if !dara.IsNil(request.OperationFunc) {
		query["OperationFunc"] = request.OperationFunc
	}

	if !dara.IsNil(request.OperationStatus) {
		query["OperationStatus"] = request.OperationStatus
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OperatorId) {
		query["OperatorId"] = request.OperatorId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOperationAuditLogs"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOperationAuditLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries administrator operation audit logs in batches.
//
// Description:
//
// - StartTime and EndTime are UNIX timestamps in seconds. StartTime must be earlier than EndTime. The interval between them cannot exceed 30 days, and StartTime cannot be more than 31 days before the current time.
//
// - If OperationStatus is not specified, only successful operation records are returned.
//
// - Results are sorted by operation time in descending order.
//
// - The return values of OperationFunc, OperationPage, and OperationType are localized based on the request language.
//
// - The values of filter parameters cannot contain single quotation marks (\\"), double quotation marks ("), or backslashes (\\\\). Otherwise, an InvalidParameter error is returned.
//
// @param request - ListOperationAuditLogsRequest
//
// @return ListOperationAuditLogsResponse
func (client *Client) ListOperationAuditLogs(request *ListOperationAuditLogsRequest) (_result *ListOperationAuditLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOperationAuditLogsResponse{}
	_body, _err := client.ListOperationAuditLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries policies for private access applications in your Alibaba Cloud account in batches.
//
// @param request - ListPolicesForPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForPrivateAccessApplicationResponse
func (client *Client) ListPolicesForPrivateAccessApplicationWithOptions(request *ListPolicesForPrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListPolicesForPrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicesForPrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicesForPrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries policies for private access applications in your Alibaba Cloud account in batches.
//
// @param request - ListPolicesForPrivateAccessApplicationRequest
//
// @return ListPolicesForPrivateAccessApplicationResponse
func (client *Client) ListPolicesForPrivateAccessApplication(request *ListPolicesForPrivateAccessApplicationRequest) (_result *ListPolicesForPrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicesForPrivateAccessApplicationResponse{}
	_body, _err := client.ListPolicesForPrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch query policies for private network access tags in your Alibaba Cloud account.
//
// @param request - ListPolicesForPrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForPrivateAccessTagResponse
func (client *Client) ListPolicesForPrivateAccessTagWithOptions(request *ListPolicesForPrivateAccessTagRequest, runtime *dara.RuntimeOptions) (_result *ListPolicesForPrivateAccessTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicesForPrivateAccessTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicesForPrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch query policies for private network access tags in your Alibaba Cloud account.
//
// @param request - ListPolicesForPrivateAccessTagRequest
//
// @return ListPolicesForPrivateAccessTagResponse
func (client *Client) ListPolicesForPrivateAccessTag(request *ListPolicesForPrivateAccessTagRequest) (_result *ListPolicesForPrivateAccessTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicesForPrivateAccessTagResponse{}
	_body, _err := client.ListPolicesForPrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries policies for multiple user groups within your Alibaba Cloud account.
//
// @param request - ListPolicesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForUserGroupResponse
func (client *Client) ListPolicesForUserGroupWithOptions(request *ListPolicesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListPolicesForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPolicesForUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPolicesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries policies for multiple user groups within your Alibaba Cloud account.
//
// @param request - ListPolicesForUserGroupRequest
//
// @return ListPolicesForUserGroupResponse
func (client *Client) ListPolicesForUserGroup(request *ListPolicesForUserGroupRequest) (_result *ListPolicesForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPolicesForUserGroupResponse{}
	_body, _err := client.ListPolicesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves bandwidth usage statistics for Secure Access Service Edge (SASE) points of presence (POPs).
//
// @param request - ListPopTrafficStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPopTrafficStatisticsResponse
func (client *Client) ListPopTrafficStatisticsWithOptions(request *ListPopTrafficStatisticsRequest, runtime *dara.RuntimeOptions) (_result *ListPopTrafficStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPopTrafficStatistics"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPopTrafficStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves bandwidth usage statistics for Secure Access Service Edge (SASE) points of presence (POPs).
//
// @param request - ListPopTrafficStatisticsRequest
//
// @return ListPopTrafficStatisticsResponse
func (client *Client) ListPopTrafficStatistics(request *ListPopTrafficStatisticsRequest) (_result *ListPopTrafficStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPopTrafficStatisticsResponse{}
	_body, _err := client.ListPopTrafficStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Layer 7 switches of internal-facing applications in batches.
//
// @param request - ListPrivateAccessApplicationL7SwitchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationL7SwitchesResponse
func (client *Client) ListPrivateAccessApplicationL7SwitchesWithOptions(request *ListPrivateAccessApplicationL7SwitchesRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessApplicationL7SwitchesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessApplicationL7Switches"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessApplicationL7SwitchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Layer 7 switches of internal-facing applications in batches.
//
// @param request - ListPrivateAccessApplicationL7SwitchesRequest
//
// @return ListPrivateAccessApplicationL7SwitchesResponse
func (client *Client) ListPrivateAccessApplicationL7Switches(request *ListPrivateAccessApplicationL7SwitchesRequest) (_result *ListPrivateAccessApplicationL7SwitchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateAccessApplicationL7SwitchesResponse{}
	_body, _err := client.ListPrivateAccessApplicationL7SwitchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about all internal-facing access applications under the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationsResponse
func (client *Client) ListPrivateAccessApplicationsWithOptions(request *ListPrivateAccessApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessModes) {
		query["AccessModes"] = request.AccessModes
	}

	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.ApplicationIds) {
		query["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ConnectorId) {
		query["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about all internal-facing access applications under the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessApplicationsRequest
//
// @return ListPrivateAccessApplicationsResponse
func (client *Client) ListPrivateAccessApplications(request *ListPrivateAccessApplicationsRequest) (_result *ListPrivateAccessApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateAccessApplicationsResponse{}
	_body, _err := client.ListPrivateAccessApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List private access applications for dynamic routes in your Alibaba Cloud account.
//
// @param request - ListPrivateAccessApplicationsForDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationsForDynamicRouteResponse
func (client *Client) ListPrivateAccessApplicationsForDynamicRouteWithOptions(request *ListPrivateAccessApplicationsForDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessApplicationsForDynamicRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessApplicationsForDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessApplicationsForDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List private access applications for dynamic routes in your Alibaba Cloud account.
//
// @param request - ListPrivateAccessApplicationsForDynamicRouteRequest
//
// @return ListPrivateAccessApplicationsForDynamicRouteResponse
func (client *Client) ListPrivateAccessApplicationsForDynamicRoute(request *ListPrivateAccessApplicationsForDynamicRouteRequest) (_result *ListPrivateAccessApplicationsForDynamicRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateAccessApplicationsForDynamicRouteResponse{}
	_body, _err := client.ListPrivateAccessApplicationsForDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all private access policies under the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessPolicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessPolicesResponse
func (client *Client) ListPrivateAccessPolicesWithOptions(request *ListPrivateAccessPolicesRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessPolicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessPolices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessPolicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all private access policies under the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessPolicesRequest
//
// @return ListPrivateAccessPolicesResponse
func (client *Client) ListPrivateAccessPolices(request *ListPrivateAccessPolicesRequest) (_result *ListPrivateAccessPolicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateAccessPolicesResponse{}
	_body, _err := client.ListPrivateAccessPolicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all internal access tags within the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessTagsResponse
func (client *Client) ListPrivateAccessTagsWithOptions(request *ListPrivateAccessTagsRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessTags"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all internal access tags within the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsRequest
//
// @return ListPrivateAccessTagsResponse
func (client *Client) ListPrivateAccessTags(request *ListPrivateAccessTagsRequest) (_result *ListPrivateAccessTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateAccessTagsResponse{}
	_body, _err := client.ListPrivateAccessTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the tags for dynamic routes in your Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsForDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessTagsForDynamicRouteResponse
func (client *Client) ListPrivateAccessTagsForDynamicRouteWithOptions(request *ListPrivateAccessTagsForDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *ListPrivateAccessTagsForDynamicRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrivateAccessTagsForDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPrivateAccessTagsForDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the tags for dynamic routes in your Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsForDynamicRouteRequest
//
// @return ListPrivateAccessTagsForDynamicRouteResponse
func (client *Client) ListPrivateAccessTagsForDynamicRoute(request *ListPrivateAccessTagsForDynamicRouteRequest) (_result *ListPrivateAccessTagsForDynamicRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPrivateAccessTagsForDynamicRouteResponse{}
	_body, _err := client.ListPrivateAccessTagsForDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of software prohibition policies under the current Alibaba Cloud account by paging.
//
// @param tmpReq - ListProhibitedPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProhibitedPoliciesResponse
func (client *Client) ListProhibitedPoliciesWithOptions(tmpReq *ListProhibitedPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListProhibitedPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProhibitedPoliciesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SoftwareId) {
		request.SoftwareIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SoftwareId, dara.String("SoftwareId"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProhibitedPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProhibitedPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of software prohibition policies under the current Alibaba Cloud account by paging.
//
// @param request - ListProhibitedPoliciesRequest
//
// @return ListProhibitedPoliciesResponse
func (client *Client) ListProhibitedPolicies(request *ListProhibitedPoliciesRequest) (_result *ListProhibitedPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProhibitedPoliciesResponse{}
	_body, _err := client.ListProhibitedPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of prohibited software under the current Alibaba Cloud account by using paging.
//
// @param tmpReq - ListProhibitedSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProhibitedSoftwareResponse
func (client *Client) ListProhibitedSoftwareWithOptions(tmpReq *ListProhibitedSoftwareRequest, runtime *dara.RuntimeOptions) (_result *ListProhibitedSoftwareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProhibitedSoftwareShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TagId) {
		request.TagIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagId, dara.String("TagId"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProhibitedSoftware"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProhibitedSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of prohibited software under the current Alibaba Cloud account by using paging.
//
// @param request - ListProhibitedSoftwareRequest
//
// @return ListProhibitedSoftwareResponse
func (client *Client) ListProhibitedSoftware(request *ListProhibitedSoftwareRequest) (_result *ListProhibitedSoftwareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProhibitedSoftwareResponse{}
	_body, _err := client.ListProhibitedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of prohibited software tags under the current Alibaba Cloud account by paging.
//
// @param tmpReq - ListProhibitedTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProhibitedTagsResponse
func (client *Client) ListProhibitedTagsWithOptions(tmpReq *ListProhibitedTagsRequest, runtime *dara.RuntimeOptions) (_result *ListProhibitedTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListProhibitedTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SoftwareId) {
		request.SoftwareIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SoftwareId, dara.String("SoftwareId"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.SoftwareIdShrink) {
		query["SoftwareId"] = request.SoftwareIdShrink
	}

	if !dara.IsNil(request.TagIds) {
		query["TagIds"] = request.TagIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProhibitedTags"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProhibitedTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of prohibited software tags under the current Alibaba Cloud account by paging.
//
// @param request - ListProhibitedTagsRequest
//
// @return ListProhibitedTagsResponse
func (client *Client) ListProhibitedTags(request *ListProhibitedTagsRequest) (_result *ListProhibitedTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProhibitedTagsResponse{}
	_body, _err := client.ListProhibitedTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of device registration policies for your Alibaba Cloud account.
//
// @param request - ListRegistrationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistrationPoliciesResponse
func (client *Client) ListRegistrationPoliciesWithOptions(request *ListRegistrationPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListRegistrationPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistrationPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistrationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of device registration policies for your Alibaba Cloud account.
//
// @param request - ListRegistrationPoliciesRequest
//
// @return ListRegistrationPoliciesResponse
func (client *Client) ListRegistrationPolicies(request *ListRegistrationPoliciesRequest) (_result *ListRegistrationPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegistrationPoliciesResponse{}
	_body, _err := client.ListRegistrationPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the device registration policies that are associated with user groups in your Alibaba Cloud account.
//
// @param request - ListRegistrationPoliciesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistrationPoliciesForUserGroupResponse
func (client *Client) ListRegistrationPoliciesForUserGroupWithOptions(request *ListRegistrationPoliciesForUserGroupRequest, runtime *dara.RuntimeOptions) (_result *ListRegistrationPoliciesForUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegistrationPoliciesForUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegistrationPoliciesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the device registration policies that are associated with user groups in your Alibaba Cloud account.
//
// @param request - ListRegistrationPoliciesForUserGroupRequest
//
// @return ListRegistrationPoliciesForUserGroupResponse
func (client *Client) ListRegistrationPoliciesForUserGroup(request *ListRegistrationPoliciesForUserGroupRequest) (_result *ListRegistrationPoliciesForUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegistrationPoliciesForUserGroupResponse{}
	_body, _err := client.ListRegistrationPoliciesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of risk events under the current Alibaba Cloud account.
//
// Description:
//
// ## Operation description
//
// - This operation is used for paging query of risk events that meet specified conditional criteria.
//
// - `CurrentPage` and `PageSize` are required parameters that specify the current page number and the number of entries per page.
//
// - You can set parameters such as `RiskId`, `RiskScene`, and `RiskCategory` to perform exact or fuzzy queries for specific risk events.
//
// - The `Status` and `StatusList` parameters cannot be used at the same time. They are used to filter risk events by disposition status.
//
// - Fuzzy match queries are supported by settings `PolicyName` and `Username`.
//
// - The response includes the total number of risk events that meet the query conditions and their details.
//
// @param request - ListRiskItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRiskItemsResponse
func (client *Client) ListRiskItemsWithOptions(request *ListRiskItemsRequest, runtime *dara.RuntimeOptions) (_result *ListRiskItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyName) {
		query["PolicyName"] = request.PolicyName
	}

	if !dara.IsNil(request.RiskCategory) {
		query["RiskCategory"] = request.RiskCategory
	}

	if !dara.IsNil(request.RiskId) {
		query["RiskId"] = request.RiskId
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.RiskScene) {
		query["RiskScene"] = request.RiskScene
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRiskItems"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRiskItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of risk events under the current Alibaba Cloud account.
//
// Description:
//
// ## Operation description
//
// - This operation is used for paging query of risk events that meet specified conditional criteria.
//
// - `CurrentPage` and `PageSize` are required parameters that specify the current page number and the number of entries per page.
//
// - You can set parameters such as `RiskId`, `RiskScene`, and `RiskCategory` to perform exact or fuzzy queries for specific risk events.
//
// - The `Status` and `StatusList` parameters cannot be used at the same time. They are used to filter risk events by disposition status.
//
// - Fuzzy match queries are supported by settings `PolicyName` and `Username`.
//
// - The response includes the total number of risk events that meet the query conditions and their details.
//
// @param request - ListRiskItemsRequest
//
// @return ListRiskItemsResponse
func (client *Client) ListRiskItems(request *ListRiskItemsRequest) (_result *ListRiskItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRiskItemsResponse{}
	_body, _err := client.ListRiskItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of user labels.
//
// @param tmpReq - ListSaseUserTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSaseUserTagsResponse
func (client *Client) ListSaseUserTagsWithOptions(tmpReq *ListSaseUserTagsRequest, runtime *dara.RuntimeOptions) (_result *ListSaseUserTagsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSaseUserTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TagIds) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, dara.String("TagIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TagIdsShrink) {
		body["TagIds"] = request.TagIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSaseUserTags"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSaseUserTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of user labels.
//
// @param request - ListSaseUserTagsRequest
//
// @return ListSaseUserTagsResponse
func (client *Client) ListSaseUserTags(request *ListSaseUserTagsRequest) (_result *ListSaseUserTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSaseUserTagsResponse{}
	_body, _err := client.ListSaseUserTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of software installed on user endpoint devices under the current Alibaba Cloud account.
//
// @param request - ListSoftwareForUserDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSoftwareForUserDeviceResponse
func (client *Client) ListSoftwareForUserDeviceWithOptions(request *ListSoftwareForUserDeviceRequest, runtime *dara.RuntimeOptions) (_result *ListSoftwareForUserDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSoftwareForUserDevice"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSoftwareForUserDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of software installed on user endpoint devices under the current Alibaba Cloud account.
//
// @param request - ListSoftwareForUserDeviceRequest
//
// @return ListSoftwareForUserDeviceResponse
func (client *Client) ListSoftwareForUserDevice(request *ListSoftwareForUserDeviceRequest) (_result *ListSoftwareForUserDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSoftwareForUserDeviceResponse{}
	_body, _err := client.ListSoftwareForUserDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries software in the software library in batches.
//
// Description:
//
// - Use CurrentPage and PageSize for pagination. NextToken and MaxResults do not take effect.
//
// - SoftwareName supports fuzzy match.
//
// - The Versions field (software version list) is not returned in the response. To query software versions, call [ListSoftwarelibVersion](~~ListSoftwarelibVersion~~).
//
// @param request - ListSoftwarelibSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSoftwarelibSoftwareResponse
func (client *Client) ListSoftwarelibSoftwareWithOptions(request *ListSoftwarelibSoftwareRequest, runtime *dara.RuntimeOptions) (_result *ListSoftwarelibSoftwareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClassifyId) {
		query["ClassifyId"] = request.ClassifyId
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Os) {
		query["Os"] = request.Os
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SoftwareName) {
		query["SoftwareName"] = request.SoftwareName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSoftwarelibSoftware"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSoftwarelibSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries software in the software library in batches.
//
// Description:
//
// - Use CurrentPage and PageSize for pagination. NextToken and MaxResults do not take effect.
//
// - SoftwareName supports fuzzy match.
//
// - The Versions field (software version list) is not returned in the response. To query software versions, call [ListSoftwarelibVersion](~~ListSoftwarelibVersion~~).
//
// @param request - ListSoftwarelibSoftwareRequest
//
// @return ListSoftwarelibSoftwareResponse
func (client *Client) ListSoftwarelibSoftware(request *ListSoftwarelibSoftwareRequest) (_result *ListSoftwarelibSoftwareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSoftwarelibSoftwareResponse{}
	_body, _err := client.ListSoftwarelibSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch query tags for private network access applications under the current Alibaba Cloud account.
//
// @param request - ListTagsForPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsForPrivateAccessApplicationResponse
func (client *Client) ListTagsForPrivateAccessApplicationWithOptions(request *ListTagsForPrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *ListTagsForPrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagsForPrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsForPrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch query tags for private network access applications under the current Alibaba Cloud account.
//
// @param request - ListTagsForPrivateAccessApplicationRequest
//
// @return ListTagsForPrivateAccessApplicationResponse
func (client *Client) ListTagsForPrivateAccessApplication(request *ListTagsForPrivateAccessApplicationRequest) (_result *ListTagsForPrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagsForPrivateAccessApplicationResponse{}
	_body, _err := client.ListTagsForPrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags of internal network access policies in your Alibaba Cloud account.
//
// @param request - ListTagsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsForPrivateAccessPolicyResponse
func (client *Client) ListTagsForPrivateAccessPolicyWithOptions(request *ListTagsForPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListTagsForPrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagsForPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags of internal network access policies in your Alibaba Cloud account.
//
// @param request - ListTagsForPrivateAccessPolicyRequest
//
// @return ListTagsForPrivateAccessPolicyResponse
func (client *Client) ListTagsForPrivateAccessPolicy(request *ListTagsForPrivateAccessPolicyRequest) (_result *ListTagsForPrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagsForPrivateAccessPolicyResponse{}
	_body, _err := client.ListTagsForPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of uninstall applications under the current Alibaba Cloud account in batches.
//
// @param request - ListUninstallApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUninstallApplicationsResponse
func (client *Client) ListUninstallApplicationsWithOptions(request *ListUninstallApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListUninstallApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUninstallApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUninstallApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of uninstall applications under the current Alibaba Cloud account in batches.
//
// @param request - ListUninstallApplicationsRequest
//
// @return ListUninstallApplicationsResponse
func (client *Client) ListUninstallApplications(request *ListUninstallApplicationsRequest) (_result *ListUninstallApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUninstallApplicationsResponse{}
	_body, _err := client.ListUninstallApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the application permissions of the logged-in user in the current Alibaba Cloud account.
//
// @param request - ListUserApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserApplicationsResponse
func (client *Client) ListUserApplicationsWithOptions(request *ListUserApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListUserApplicationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserApplications"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the application permissions of the logged-in user in the current Alibaba Cloud account.
//
// @param request - ListUserApplicationsRequest
//
// @return ListUserApplicationsResponse
func (client *Client) ListUserApplications(request *ListUserApplicationsRequest) (_result *ListUserApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserApplicationsResponse{}
	_body, _err := client.ListUserApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of user endpoint devices under the current Alibaba Cloud account.
//
// @param request - ListUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDevicesResponse
func (client *Client) ListUserDevicesWithOptions(request *ListUserDevicesRequest, runtime *dara.RuntimeOptions) (_result *ListUserDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppStatuses) {
		query["AppStatuses"] = request.AppStatuses
	}

	if !dara.IsNil(request.AppVersions) {
		query["AppVersions"] = request.AppVersions
	}

	if !dara.IsNil(request.AutoLoginStatuses) {
		query["AutoLoginStatuses"] = request.AutoLoginStatuses
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Department) {
		query["Department"] = request.Department
	}

	if !dara.IsNil(request.DeviceBelong) {
		query["DeviceBelong"] = request.DeviceBelong
	}

	if !dara.IsNil(request.DeviceGroupId) {
		query["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.DeviceStatuses) {
		query["DeviceStatuses"] = request.DeviceStatuses
	}

	if !dara.IsNil(request.DeviceTags) {
		query["DeviceTags"] = request.DeviceTags
	}

	if !dara.IsNil(request.DeviceTypes) {
		query["DeviceTypes"] = request.DeviceTypes
	}

	if !dara.IsNil(request.DlpStatuses) {
		query["DlpStatuses"] = request.DlpStatuses
	}

	if !dara.IsNil(request.Hostname) {
		query["Hostname"] = request.Hostname
	}

	if !dara.IsNil(request.IaStatuses) {
		query["IaStatuses"] = request.IaStatuses
	}

	if !dara.IsNil(request.InnerIp) {
		query["InnerIp"] = request.InnerIp
	}

	if !dara.IsNil(request.Mac) {
		query["Mac"] = request.Mac
	}

	if !dara.IsNil(request.NacStatuses) {
		query["NacStatuses"] = request.NacStatuses
	}

	if !dara.IsNil(request.PaStatuses) {
		query["PaStatuses"] = request.PaStatuses
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SaseUserId) {
		query["SaseUserId"] = request.SaseUserId
	}

	if !dara.IsNil(request.SharingStatus) {
		query["SharingStatus"] = request.SharingStatus
	}

	if !dara.IsNil(request.SnBios) {
		query["SnBios"] = request.SnBios
	}

	if !dara.IsNil(request.SnSystem) {
		query["SnSystem"] = request.SnSystem
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.Workshop) {
		query["Workshop"] = request.Workshop
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of user endpoint devices under the current Alibaba Cloud account.
//
// @param request - ListUserDevicesRequest
//
// @return ListUserDevicesResponse
func (client *Client) ListUserDevices(request *ListUserDevicesRequest) (_result *ListUserDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserDevicesResponse{}
	_body, _err := client.ListUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about all user groups in your Alibaba Cloud account.
//
// @param request - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroups"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about all user groups in your Alibaba Cloud account.
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve user groups for private network access policies in batches under your Alibaba Cloud account.
//
// @param request - ListUserGroupsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForPrivateAccessPolicyResponse
func (client *Client) ListUserGroupsForPrivateAccessPolicyWithOptions(request *ListUserGroupsForPrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsForPrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupsForPrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve user groups for private network access policies in batches under your Alibaba Cloud account.
//
// @param request - ListUserGroupsForPrivateAccessPolicyRequest
//
// @return ListUserGroupsForPrivateAccessPolicyResponse
func (client *Client) ListUserGroupsForPrivateAccessPolicy(request *ListUserGroupsForPrivateAccessPolicyRequest) (_result *ListUserGroupsForPrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsForPrivateAccessPolicyResponse{}
	_body, _err := client.ListUserGroupsForPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the user groups associated with device registration policies in your Alibaba Cloud account.
//
// @param request - ListUserGroupsForRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForRegistrationPolicyResponse
func (client *Client) ListUserGroupsForRegistrationPolicyWithOptions(request *ListUserGroupsForRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *ListUserGroupsForRegistrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserGroupsForRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserGroupsForRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the user groups associated with device registration policies in your Alibaba Cloud account.
//
// @param request - ListUserGroupsForRegistrationPolicyRequest
//
// @return ListUserGroupsForRegistrationPolicyResponse
func (client *Client) ListUserGroupsForRegistrationPolicy(request *ListUserGroupsForRegistrationPolicyRequest) (_result *ListUserGroupsForRegistrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserGroupsForRegistrationPolicyResponse{}
	_body, _err := client.ListUserGroupsForRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of user zero trust policies.
//
// @param request - ListUserPrivateAccessPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPrivateAccessPoliciesResponse
func (client *Client) ListUserPrivateAccessPoliciesWithOptions(request *ListUserPrivateAccessPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListUserPrivateAccessPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPrivateAccessPolicies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPrivateAccessPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of user zero trust policies.
//
// @param request - ListUserPrivateAccessPoliciesRequest
//
// @return ListUserPrivateAccessPoliciesResponse
func (client *Client) ListUserPrivateAccessPolicies(request *ListUserPrivateAccessPoliciesRequest) (_result *ListUserPrivateAccessPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPrivateAccessPoliciesResponse{}
	_body, _err := client.ListUserPrivateAccessPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the users for the current Alibaba Cloud account.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the users for the current Alibaba Cloud account.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virus files detected under the current Alibaba Cloud account and their disposition status with paging. Supports filtering by virus type, risk level, user terminal device, user, and discovery time.
//
// @param request - ListVirusFileStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusFileStatusesResponse
func (client *Client) ListVirusFileStatusesWithOptions(request *ListVirusFileStatusesRequest, runtime *dara.RuntimeOptions) (_result *ListVirusFileStatusesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusFileStatuses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusFileStatusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virus files detected under the current Alibaba Cloud account and their disposition status with paging. Supports filtering by virus type, risk level, user terminal device, user, and discovery time.
//
// @param request - ListVirusFileStatusesRequest
//
// @return ListVirusFileStatusesResponse
func (client *Client) ListVirusFileStatuses(request *ListVirusFileStatusesRequest) (_result *ListVirusFileStatusesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusFileStatusesResponse{}
	_body, _err := client.ListVirusFileStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询病毒扫描额外名单
//
// @param request - ListVirusScanAdditionalListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanAdditionalListsResponse
func (client *Client) ListVirusScanAdditionalListsWithOptions(request *ListVirusScanAdditionalListsRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanAdditionalListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanAdditionalLists"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanAdditionalListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询病毒扫描额外名单
//
// @param request - ListVirusScanAdditionalListsRequest
//
// @return ListVirusScanAdditionalListsResponse
func (client *Client) ListVirusScanAdditionalLists(request *ListVirusScanAdditionalListsRequest) (_result *ListVirusScanAdditionalListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusScanAdditionalListsResponse{}
	_body, _err := client.ListVirusScanAdditionalListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries virus scheduled scan policies under the current Alibaba Cloud account with paging.
//
// @param request - ListVirusScanScheduledStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanScheduledStrategiesResponse
func (client *Client) ListVirusScanScheduledStrategiesWithOptions(request *ListVirusScanScheduledStrategiesRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanScheduledStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanScheduledStrategies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanScheduledStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries virus scheduled scan policies under the current Alibaba Cloud account with paging.
//
// @param request - ListVirusScanScheduledStrategiesRequest
//
// @return ListVirusScanScheduledStrategiesResponse
func (client *Client) ListVirusScanScheduledStrategies(request *ListVirusScanScheduledStrategiesRequest) (_result *ListVirusScanScheduledStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusScanScheduledStrategiesResponse{}
	_body, _err := client.ListVirusScanScheduledStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution progress of specified virus scan tasks on user terminal devices in batches.
//
// @param request - ListVirusScanTaskStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanTaskStatusesResponse
func (client *Client) ListVirusScanTaskStatusesWithOptions(request *ListVirusScanTaskStatusesRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanTaskStatusesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanTaskStatuses"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanTaskStatusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution progress of specified virus scan tasks on user terminal devices in batches.
//
// @param request - ListVirusScanTaskStatusesRequest
//
// @return ListVirusScanTaskStatusesResponse
func (client *Client) ListVirusScanTaskStatuses(request *ListVirusScanTaskStatusesRequest) (_result *ListVirusScanTaskStatusesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusScanTaskStatusesResponse{}
	_body, _err := client.ListVirusScanTaskStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the detection result statistics of specified virus scan tasks in batches.
//
// @param request - ListVirusScanTaskSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanTaskSummaryResponse
func (client *Client) ListVirusScanTaskSummaryWithOptions(request *ListVirusScanTaskSummaryRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanTaskSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanTaskSummary"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanTaskSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the detection result statistics of specified virus scan tasks in batches.
//
// @param request - ListVirusScanTaskSummaryRequest
//
// @return ListVirusScanTaskSummaryResponse
func (client *Client) ListVirusScanTaskSummary(request *ListVirusScanTaskSummaryRequest) (_result *ListVirusScanTaskSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusScanTaskSummaryResponse{}
	_body, _err := client.ListVirusScanTaskSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries instant virus scan tasks under the current Alibaba Cloud account by paging.
//
// @param request - ListVirusScanTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVirusScanTasksResponse
func (client *Client) ListVirusScanTasksWithOptions(request *ListVirusScanTasksRequest, runtime *dara.RuntimeOptions) (_result *ListVirusScanTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVirusScanTasks"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVirusScanTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries instant virus scan tasks under the current Alibaba Cloud account by paging.
//
// @param request - ListVirusScanTasksRequest
//
// @return ListVirusScanTasksResponse
func (client *Client) ListVirusScanTasks(request *ListVirusScanTasksRequest) (_result *ListVirusScanTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVirusScanTasksResponse{}
	_body, _err := client.ListVirusScanTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries scheduled vulnerability scan policies under the current Alibaba Cloud account by paging.
//
// @param request - ListVulScanScheduledStrategiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVulScanScheduledStrategiesResponse
func (client *Client) ListVulScanScheduledStrategiesWithOptions(request *ListVulScanScheduledStrategiesRequest, runtime *dara.RuntimeOptions) (_result *ListVulScanScheduledStrategiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVulScanScheduledStrategies"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVulScanScheduledStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scheduled vulnerability scan policies under the current Alibaba Cloud account by paging.
//
// @param request - ListVulScanScheduledStrategiesRequest
//
// @return ListVulScanScheduledStrategiesResponse
func (client *Client) ListVulScanScheduledStrategies(request *ListVulScanScheduledStrategiesRequest) (_result *ListVulScanScheduledStrategiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVulScanScheduledStrategiesResponse{}
	_body, _err := client.ListVulScanScheduledStrategiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries vulnerability scanning tasks under the current Alibaba Cloud account by paged query.
//
// @param request - ListVulScanTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVulScanTasksResponse
func (client *Client) ListVulScanTasksWithOptions(request *ListVulScanTasksRequest, runtime *dara.RuntimeOptions) (_result *ListVulScanTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVulScanTasks"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVulScanTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries vulnerability scanning tasks under the current Alibaba Cloud account by paged query.
//
// @param request - ListVulScanTasksRequest
//
// @return ListVulScanTasksResponse
func (client *Client) ListVulScanTasks(request *ListVulScanTasksRequest) (_result *ListVulScanTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVulScanTasksResponse{}
	_body, _err := client.ListVulScanTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries vulnerabilities detected by scans under the current Alibaba Cloud account by using paged query with paging.
//
// @param request - ListVulnerabilitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVulnerabilitiesResponse
func (client *Client) ListVulnerabilitiesWithOptions(request *ListVulnerabilitiesRequest, runtime *dara.RuntimeOptions) (_result *ListVulnerabilitiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVulnerabilities"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVulnerabilitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries vulnerabilities detected by scans under the current Alibaba Cloud account by using paged query with paging.
//
// @param request - ListVulnerabilitiesRequest
//
// @return ListVulnerabilitiesResponse
func (client *Client) ListVulnerabilities(request *ListVulnerabilitiesRequest) (_result *ListVulnerabilitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVulnerabilitiesResponse{}
	_body, _err := client.ListVulnerabilitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Look up an existing watermark information mapping to retrieve the corresponding string-formatted watermark information from numeric-formatted watermark data.
//
// @param request - LookupWmInfoMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupWmInfoMappingResponse
func (client *Client) LookupWmInfoMappingWithOptions(request *LookupWmInfoMappingRequest, runtime *dara.RuntimeOptions) (_result *LookupWmInfoMappingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LookupWmInfoMapping"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LookupWmInfoMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Look up an existing watermark information mapping to retrieve the corresponding string-formatted watermark information from numeric-formatted watermark data.
//
// @param request - LookupWmInfoMappingRequest
//
// @return LookupWmInfoMappingResponse
func (client *Client) LookupWmInfoMapping(request *LookupWmInfoMappingRequest) (_result *LookupWmInfoMappingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LookupWmInfoMappingResponse{}
	_body, _err := client.LookupWmInfoMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an enterprise acceleration policy.
//
// @param request - ModifyEnterpriseAcceleratePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyEnterpriseAcceleratePolicyResponse
func (client *Client) ModifyEnterpriseAcceleratePolicyWithOptions(request *ModifyEnterpriseAcceleratePolicyRequest, runtime *dara.RuntimeOptions) (_result *ModifyEnterpriseAcceleratePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccelerationType) {
		body["AccelerationType"] = request.AccelerationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EapId) {
		body["EapId"] = request.EapId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OnTls) {
		body["OnTls"] = request.OnTls
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ShowInClient) {
		body["ShowInClient"] = request.ShowInClient
	}

	if !dara.IsNil(request.UpstreamHost) {
		body["UpstreamHost"] = request.UpstreamHost
	}

	if !dara.IsNil(request.UpstreamPort) {
		body["UpstreamPort"] = request.UpstreamPort
	}

	if !dara.IsNil(request.UpstreamType) {
		body["UpstreamType"] = request.UpstreamType
	}

	if !dara.IsNil(request.UserAttributeGroup) {
		body["UserAttributeGroup"] = request.UserAttributeGroup
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyEnterpriseAcceleratePolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an enterprise acceleration policy.
//
// @param request - ModifyEnterpriseAcceleratePolicyRequest
//
// @return ModifyEnterpriseAcceleratePolicyResponse
func (client *Client) ModifyEnterpriseAcceleratePolicy(request *ModifyEnterpriseAcceleratePolicyRequest) (_result *ModifyEnterpriseAcceleratePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyEnterpriseAcceleratePolicyResponse{}
	_body, _err := client.ModifyEnterpriseAcceleratePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a forwarding rule.
//
// @param request - ModifyForwardStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyForwardStrategyResponse
func (client *Client) ModifyForwardStrategyWithOptions(request *ModifyForwardStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyForwardStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationId) {
		body["DestinationId"] = request.DestinationId
	}

	if !dara.IsNil(request.DestinationType) {
		body["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.ForwardId) {
		body["ForwardId"] = request.ForwardId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyForwardStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyForwardStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a forwarding rule.
//
// @param request - ModifyForwardStrategyRequest
//
// @return ModifyForwardStrategyResponse
func (client *Client) ModifyForwardStrategy(request *ModifyForwardStrategyRequest) (_result *ModifyForwardStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyForwardStrategyResponse{}
	_body, _err := client.ModifyForwardStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the binding items of a forwarding rule.
//
// @param request - ModifyForwardStrategyBindingItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyForwardStrategyBindingItemsResponse
func (client *Client) ModifyForwardStrategyBindingItemsWithOptions(request *ModifyForwardStrategyBindingItemsRequest, runtime *dara.RuntimeOptions) (_result *ModifyForwardStrategyBindingItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ItemIds) {
		query["ItemIds"] = request.ItemIds
	}

	if !dara.IsNil(request.MatchMode) {
		query["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.ModifyType) {
		query["ModifyType"] = request.ModifyType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ForwardId) {
		body["ForwardId"] = request.ForwardId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyForwardStrategyBindingItems"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyForwardStrategyBindingItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the binding items of a forwarding rule.
//
// @param request - ModifyForwardStrategyBindingItemsRequest
//
// @return ModifyForwardStrategyBindingItemsResponse
func (client *Client) ModifyForwardStrategyBindingItems(request *ModifyForwardStrategyBindingItemsRequest) (_result *ModifyForwardStrategyBindingItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyForwardStrategyBindingItemsResponse{}
	_body, _err := client.ModifyForwardStrategyBindingItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes associated terminal devices from a static device label in batches.
//
// @param request - RemoveDeviceGroupMatchDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDeviceGroupMatchDevicesResponse
func (client *Client) RemoveDeviceGroupMatchDevicesWithOptions(request *RemoveDeviceGroupMatchDevicesRequest, runtime *dara.RuntimeOptions) (_result *RemoveDeviceGroupMatchDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DevTags) {
		bodyFlat["DevTags"] = request.DevTags
	}

	if !dara.IsNil(request.DeviceGroupId) {
		body["DeviceGroupId"] = request.DeviceGroupId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDeviceGroupMatchDevices"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDeviceGroupMatchDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes associated terminal devices from a static device label in batches.
//
// @param request - RemoveDeviceGroupMatchDevicesRequest
//
// @return RemoveDeviceGroupMatchDevicesResponse
func (client *Client) RemoveDeviceGroupMatchDevices(request *RemoveDeviceGroupMatchDevicesRequest) (_result *RemoveDeviceGroupMatchDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDeviceGroupMatchDevicesResponse{}
	_body, _err := client.RemoveDeviceGroupMatchDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes virus scan blacklists and whitelists entries in batch by entry IDs. The entire removal operation is failed if any of the specified entry IDs do not belong to the current Alibaba Cloud account.
//
// @param request - RemoveVirusScanAdditionalListsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveVirusScanAdditionalListsResponse
func (client *Client) RemoveVirusScanAdditionalListsWithOptions(request *RemoveVirusScanAdditionalListsRequest, runtime *dara.RuntimeOptions) (_result *RemoveVirusScanAdditionalListsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ListIds) {
		bodyFlat["ListIds"] = request.ListIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveVirusScanAdditionalLists"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveVirusScanAdditionalListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes virus scan blacklists and whitelists entries in batch by entry IDs. The entire removal operation is failed if any of the specified entry IDs do not belong to the current Alibaba Cloud account.
//
// @param request - RemoveVirusScanAdditionalListsRequest
//
// @return RemoveVirusScanAdditionalListsResponse
func (client *Client) RemoveVirusScanAdditionalLists(request *RemoveVirusScanAdditionalListsRequest) (_result *RemoveVirusScanAdditionalListsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveVirusScanAdditionalListsResponse{}
	_body, _err := client.RemoveVirusScanAdditionalListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes a user device session.
//
// @param request - RevokeUserDeviceSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeUserDeviceSessionResponse
func (client *Client) RevokeUserDeviceSessionWithOptions(request *RevokeUserDeviceSessionRequest, runtime *dara.RuntimeOptions) (_result *RevokeUserDeviceSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DevTag) {
		body["DevTag"] = request.DevTag
	}

	if !dara.IsNil(request.SaseUserId) {
		body["SaseUserId"] = request.SaseUserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeUserDeviceSession"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeUserDeviceSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes a user device session.
//
// @param request - RevokeUserDeviceSessionRequest
//
// @return RevokeUserDeviceSessionResponse
func (client *Client) RevokeUserDeviceSession(request *RevokeUserDeviceSessionRequest) (_result *RevokeUserDeviceSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeUserDeviceSessionResponse{}
	_body, _err := client.RevokeUserDeviceSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI RevokeUserSession is deprecated
//
// Summary:
//
// Revokes a user logon session.
//
// @param request - RevokeUserSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeUserSessionResponse
func (client *Client) RevokeUserSessionWithOptions(request *RevokeUserSessionRequest, runtime *dara.RuntimeOptions) (_result *RevokeUserSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalIds) {
		query["ExternalIds"] = request.ExternalIds
	}

	if !dara.IsNil(request.IdpId) {
		query["IdpId"] = request.IdpId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeUserSession"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeUserSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI RevokeUserSession is deprecated
//
// Summary:
//
// Revokes a user logon session.
//
// @param request - RevokeUserSessionRequest
//
// @return RevokeUserSessionResponse
// Deprecated
func (client *Client) RevokeUserSession(request *RevokeUserSessionRequest) (_result *RevokeUserSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeUserSessionResponse{}
	_body, _err := client.RevokeUserSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the real-time anti-virus defense policy for the current Alibaba Cloud account. The first call creates the policy, and subsequent calls update it. The complete updated configuration is returned. When configuring for the first time, Status, MatchMode, HighRiskOperation, MidRiskOperation, LowRiskOperation, and ScanTargets are all required. ScanTargets and Whitelist are full replacements. The collection you pass in replaces the existing configuration. When MatchMode is set to UserGroupNormal, you must pass in the complete UserGroupIds on every call. When Status is not set to Disabled, the system validates the endpoint hardening license count. The call fails if the count exceeds the purchased licenses.
//
// @param request - UpdateAntiVirusRealTimeDefenceStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAntiVirusRealTimeDefenceStrategyResponse
func (client *Client) UpdateAntiVirusRealTimeDefenceStrategyWithOptions(request *UpdateAntiVirusRealTimeDefenceStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateAntiVirusRealTimeDefenceStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HighRiskOperation) {
		body["HighRiskOperation"] = request.HighRiskOperation
	}

	if !dara.IsNil(request.LowRiskOperation) {
		body["LowRiskOperation"] = request.LowRiskOperation
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MaxCpuUsage) {
		body["MaxCpuUsage"] = request.MaxCpuUsage
	}

	if !dara.IsNil(request.MidRiskOperation) {
		body["MidRiskOperation"] = request.MidRiskOperation
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ScanTargets) {
		bodyFlat["ScanTargets"] = request.ScanTargets
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAntiVirusRealTimeDefenceStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAntiVirusRealTimeDefenceStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the real-time anti-virus defense policy for the current Alibaba Cloud account. The first call creates the policy, and subsequent calls update it. The complete updated configuration is returned. When configuring for the first time, Status, MatchMode, HighRiskOperation, MidRiskOperation, LowRiskOperation, and ScanTargets are all required. ScanTargets and Whitelist are full replacements. The collection you pass in replaces the existing configuration. When MatchMode is set to UserGroupNormal, you must pass in the complete UserGroupIds on every call. When Status is not set to Disabled, the system validates the endpoint hardening license count. The call fails if the count exceeds the purchased licenses.
//
// @param request - UpdateAntiVirusRealTimeDefenceStrategyRequest
//
// @return UpdateAntiVirusRealTimeDefenceStrategyResponse
func (client *Client) UpdateAntiVirusRealTimeDefenceStrategy(request *UpdateAntiVirusRealTimeDefenceStrategyRequest) (_result *UpdateAntiVirusRealTimeDefenceStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAntiVirusRealTimeDefenceStrategyResponse{}
	_body, _err := client.UpdateAntiVirusRealTimeDefenceStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an approval process under the current Alibaba Cloud account.
//
// @param tmpReq - UpdateApprovalProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApprovalProcessResponse
func (client *Client) UpdateApprovalProcessWithOptions(tmpReq *UpdateApprovalProcessRequest, runtime *dara.RuntimeOptions) (_result *UpdateApprovalProcessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateApprovalProcessShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MatchSchemaConfigs) {
		request.MatchSchemaConfigsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchSchemaConfigs, dara.String("MatchSchemaConfigs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.MatchSchemas) {
		request.MatchSchemasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MatchSchemas, dara.String("MatchSchemas"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalType) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !dara.IsNil(request.EventLabel) {
		query["EventLabel"] = request.EventLabel
	}

	if !dara.IsNil(request.ExternalConfig) {
		query["ExternalConfig"] = request.ExternalConfig
	}

	if !dara.IsNil(request.MatchSchemaConfigsShrink) {
		query["MatchSchemaConfigs"] = request.MatchSchemaConfigsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchSchemasShrink) {
		body["MatchSchemas"] = request.MatchSchemasShrink
	}

	if !dara.IsNil(request.ProcessId) {
		body["ProcessId"] = request.ProcessId
	}

	if !dara.IsNil(request.ProcessName) {
		body["ProcessName"] = request.ProcessName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ProcessNodes) {
		bodyFlat["ProcessNodes"] = request.ProcessNodes
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApprovalProcess"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApprovalProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an approval process under the current Alibaba Cloud account.
//
// @param request - UpdateApprovalProcessRequest
//
// @return UpdateApprovalProcessResponse
func (client *Client) UpdateApprovalProcess(request *UpdateApprovalProcessRequest) (_result *UpdateApprovalProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApprovalProcessResponse{}
	_body, _err := client.UpdateApprovalProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the instance status of an approval under the current Alibaba Cloud account.
//
// @param request - UpdateApprovalStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApprovalStatusResponse
func (client *Client) UpdateApprovalStatusWithOptions(request *UpdateApprovalStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateApprovalStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalId) {
		query["ApprovalId"] = request.ApprovalId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApprovalStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApprovalStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the instance status of an approval under the current Alibaba Cloud account.
//
// @param request - UpdateApprovalStatusRequest
//
// @return UpdateApprovalStatusResponse
func (client *Client) UpdateApprovalStatus(request *UpdateApprovalStatusRequest) (_result *UpdateApprovalStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateApprovalStatusResponse{}
	_body, _err := client.UpdateApprovalStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the auto-start and anti-uninstall policy for your Alibaba Cloud account.
//
// @param tmpReq - UpdateBootAndAntiUninstallPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBootAndAntiUninstallPolicyResponse
func (client *Client) UpdateBootAndAntiUninstallPolicyWithOptions(tmpReq *UpdateBootAndAntiUninstallPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateBootAndAntiUninstallPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateBootAndAntiUninstallPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.BlockContent) {
		request.BlockContentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.BlockContent, dara.String("BlockContent"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowReport) {
		body["AllowReport"] = request.AllowReport
	}

	if !dara.IsNil(request.BlockContentShrink) {
		body["BlockContent"] = request.BlockContentShrink
	}

	if !dara.IsNil(request.IsAntiUninstall) {
		body["IsAntiUninstall"] = request.IsAntiUninstall
	}

	if !dara.IsNil(request.IsBoot) {
		body["IsBoot"] = request.IsBoot
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.WhitelistUsers) {
		bodyFlat["WhitelistUsers"] = request.WhitelistUsers
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBootAndAntiUninstallPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBootAndAntiUninstallPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the auto-start and anti-uninstall policy for your Alibaba Cloud account.
//
// @param request - UpdateBootAndAntiUninstallPolicyRequest
//
// @return UpdateBootAndAntiUninstallPolicyResponse
func (client *Client) UpdateBootAndAntiUninstallPolicy(request *UpdateBootAndAntiUninstallPolicyRequest) (_result *UpdateBootAndAntiUninstallPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBootAndAntiUninstallPolicyResponse{}
	_body, _err := client.UpdateBootAndAntiUninstallPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update user information for a custom identity provider in your Alibaba Cloud account.
//
// @param request - UpdateClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserResponse
func (client *Client) UpdateClientUserWithOptions(request *UpdateClientUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MobileNumber) {
		query["MobileNumber"] = request.MobileNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientUser"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update user information for a custom identity provider in your Alibaba Cloud account.
//
// @param request - UpdateClientUserRequest
//
// @return UpdateClientUserResponse
func (client *Client) UpdateClientUser(request *UpdateClientUserRequest) (_result *UpdateClientUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClientUserResponse{}
	_body, _err := client.UpdateClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the password for a specified user that belongs to a custom identity source.
//
// @param request - UpdateClientUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserPasswordResponse
func (client *Client) UpdateClientUserPasswordWithOptions(request *UpdateClientUserPasswordRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientUserPasswordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientUserPassword"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the password for a specified user that belongs to a custom identity source.
//
// @param request - UpdateClientUserPasswordRequest
//
// @return UpdateClientUserPasswordResponse
func (client *Client) UpdateClientUserPassword(request *UpdateClientUserPasswordRequest) (_result *UpdateClientUserPasswordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClientUserPasswordResponse{}
	_body, _err := client.UpdateClientUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the status of a specified user from a custom identity source for your Alibaba Cloud account.
//
// @param request - UpdateClientUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserStatusResponse
func (client *Client) UpdateClientUserStatusWithOptions(request *UpdateClientUserStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateClientUserStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClientUserStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClientUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the status of a specified user from a custom identity source for your Alibaba Cloud account.
//
// @param request - UpdateClientUserStatusRequest
//
// @return UpdateClientUserStatusResponse
func (client *Client) UpdateClientUserStatus(request *UpdateClientUserStatusRequest) (_result *UpdateClientUserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClientUserStatusResponse{}
	_body, _err := client.UpdateClientUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a Connector instance under the current Alibaba Cloud account.
//
// @param request - UpdateConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectorResponse
func (client *Client) UpdateConnectorWithOptions(request *UpdateConnectorRequest, runtime *dara.RuntimeOptions) (_result *UpdateConnectorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccelerateStatus) {
		body["AccelerateStatus"] = request.AccelerateStatus
	}

	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SwitchStatus) {
		body["SwitchStatus"] = request.SwitchStatus
	}

	if !dara.IsNil(request.VipCidr) {
		body["VipCidr"] = request.VipCidr
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnector"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a Connector instance under the current Alibaba Cloud account.
//
// @param request - UpdateConnectorRequest
//
// @return UpdateConnectorResponse
func (client *Client) UpdateConnector(request *UpdateConnectorRequest) (_result *UpdateConnectorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConnectorResponse{}
	_body, _err := client.UpdateConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a ConnectorClient under the current Alibaba Cloud account.
//
// @param request - UpdateConnectorClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectorClientResponse
func (client *Client) UpdateConnectorClientWithOptions(request *UpdateConnectorClientRequest, runtime *dara.RuntimeOptions) (_result *UpdateConnectorClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConnectorId) {
		body["ConnectorId"] = request.ConnectorId
	}

	if !dara.IsNil(request.DevTag) {
		body["DevTag"] = request.DevTag
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnectorClient"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectorClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a ConnectorClient under the current Alibaba Cloud account.
//
// @param request - UpdateConnectorClientRequest
//
// @return UpdateConnectorClientResponse
func (client *Client) UpdateConnectorClient(request *UpdateConnectorClientRequest) (_result *UpdateConnectorClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConnectorClientResponse{}
	_body, _err := client.UpdateConnectorClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a device label.
//
// @param request - UpdateDeviceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDeviceGroupResponse
func (client *Client) UpdateDeviceGroupWithOptions(request *UpdateDeviceGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateDeviceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceGroupId) {
		body["DeviceGroupId"] = request.DeviceGroupId
	}

	if !dara.IsNil(request.DynamicOperator) {
		body["DynamicOperator"] = request.DynamicOperator
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDeviceGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDeviceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a device label.
//
// @param request - UpdateDeviceGroupRequest
//
// @return UpdateDeviceGroupResponse
func (client *Client) UpdateDeviceGroup(request *UpdateDeviceGroupRequest) (_result *UpdateDeviceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDeviceGroupResponse{}
	_body, _err := client.UpdateDeviceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of a domain name list.
//
// @param request - UpdateDomainMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainMetaResponse
func (client *Client) UpdateDomainMetaWithOptions(request *UpdateDomainMetaRequest, runtime *dara.RuntimeOptions) (_result *UpdateDomainMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ListId) {
		body["ListId"] = request.ListId
	}

	if !dara.IsNil(request.ListType) {
		body["ListType"] = request.ListType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDomainMeta"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDomainMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of a domain name list.
//
// @param request - UpdateDomainMetaRequest
//
// @return UpdateDomainMetaResponse
func (client *Client) UpdateDomainMeta(request *UpdateDomainMetaRequest) (_result *UpdateDomainMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDomainMetaResponse{}
	_body, _err := client.UpdateDomainMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a dynamic route in your Alibaba Cloud account.
//
// @param request - UpdateDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDynamicRouteResponse
func (client *Client) UpdateDynamicRouteWithOptions(request *UpdateDynamicRouteRequest, runtime *dara.RuntimeOptions) (_result *UpdateDynamicRouteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DynamicRouteId) {
		body["DynamicRouteId"] = request.DynamicRouteId
	}

	if !dara.IsNil(request.DynamicRouteType) {
		body["DynamicRouteType"] = request.DynamicRouteType
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NextHop) {
		body["NextHop"] = request.NextHop
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.RegionIds) {
		bodyFlat["RegionIds"] = request.RegionIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDynamicRoute"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a dynamic route in your Alibaba Cloud account.
//
// @param request - UpdateDynamicRouteRequest
//
// @return UpdateDynamicRouteResponse
func (client *Client) UpdateDynamicRoute(request *UpdateDynamicRouteRequest) (_result *UpdateDynamicRouteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDynamicRouteResponse{}
	_body, _err := client.UpdateDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the status of device registration applications that exceed your Alibaba Cloud account\\"s quota.
//
// @param request - UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatusWithOptions(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateExcessiveDeviceRegistrationApplicationsStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateExcessiveDeviceRegistrationApplicationsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the status of device registration applications that exceed your Alibaba Cloud account\\"s quota.
//
// @param request - UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
//
// @return UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatus(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) (_result *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateExcessiveDeviceRegistrationApplicationsStatusResponse{}
	_body, _err := client.UpdateExcessiveDeviceRegistrationApplicationsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a department from a custom identity provider for the current Alibaba Cloud account.
//
// @param request - UpdateIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdpDepartmentResponse
func (client *Client) UpdateIdpDepartmentWithOptions(request *UpdateIdpDepartmentRequest, runtime *dara.RuntimeOptions) (_result *UpdateIdpDepartmentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DepartmentId) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !dara.IsNil(request.DepartmentName) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !dara.IsNil(request.IdpConfigId) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIdpDepartment"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIdpDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a department from a custom identity provider for the current Alibaba Cloud account.
//
// @param request - UpdateIdpDepartmentRequest
//
// @return UpdateIdpDepartmentResponse
func (client *Client) UpdateIdpDepartment(request *UpdateIdpDepartmentRequest) (_result *UpdateIdpDepartmentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIdpDepartmentResponse{}
	_body, _err := client.UpdateIdpDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the phone number whitelist for visitor access SMS logon.
//
// Description:
//
// - A maximum of 1024 phone numbers are supported.
//
// - Duplicate phone numbers are not allowed. Phone numbers in invalid formats are rejected. Only Chinese mainland phone numbers are supported.
//
// - You must update all phone numbers at once. Incremental updates are not supported.
//
// @param request - UpdateNacPortalSmsPhoneWhitelistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacPortalSmsPhoneWhitelistResponse
func (client *Client) UpdateNacPortalSmsPhoneWhitelistWithOptions(request *UpdateNacPortalSmsPhoneWhitelistRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacPortalSmsPhoneWhitelistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Phones) {
		query["Phones"] = request.Phones
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacPortalSmsPhoneWhitelist"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacPortalSmsPhoneWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the phone number whitelist for visitor access SMS logon.
//
// Description:
//
// - A maximum of 1024 phone numbers are supported.
//
// - Duplicate phone numbers are not allowed. Phone numbers in invalid formats are rejected. Only Chinese mainland phone numbers are supported.
//
// - You must update all phone numbers at once. Incremental updates are not supported.
//
// @param request - UpdateNacPortalSmsPhoneWhitelistRequest
//
// @return UpdateNacPortalSmsPhoneWhitelistResponse
func (client *Client) UpdateNacPortalSmsPhoneWhitelist(request *UpdateNacPortalSmsPhoneWhitelistRequest) (_result *UpdateNacPortalSmsPhoneWhitelistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacPortalSmsPhoneWhitelistResponse{}
	_body, _err := client.UpdateNacPortalSmsPhoneWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the network access certificate status for users in your Alibaba Cloud account.
//
// @param request - UpdateNacUserCertStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacUserCertStatusResponse
func (client *Client) UpdateNacUserCertStatusWithOptions(request *UpdateNacUserCertStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateNacUserCertStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.IdList) {
		bodyFlat["IdList"] = request.IdList
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNacUserCertStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNacUserCertStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the network access certificate status for users in your Alibaba Cloud account.
//
// @param request - UpdateNacUserCertStatusRequest
//
// @return UpdateNacUserCertStatusResponse
func (client *Client) UpdateNacUserCertStatus(request *UpdateNacUserCertStatusRequest) (_result *UpdateNacUserCertStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNacUserCertStatusResponse{}
	_body, _err := client.UpdateNacUserCertStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a private access application under the current Alibaba Cloud account.
//
// @param tmpReq - UpdatePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessApplicationResponse
func (client *Client) UpdatePrivateAccessApplicationWithOptions(tmpReq *UpdatePrivateAccessApplicationRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateAccessApplicationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdatePrivateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.L7Config) {
		request.L7ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.L7Config, dara.String("L7Config"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UnauthorizedAccessConfig) {
		request.UnauthorizedAccessConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UnauthorizedAccessConfig, dara.String("UnauthorizedAccessConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AddressGroups) {
		bodyFlat["AddressGroups"] = request.AddressGroups
	}

	if !dara.IsNil(request.Addresses) {
		bodyFlat["Addresses"] = request.Addresses
	}

	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ConfigMode) {
		body["ConfigMode"] = request.ConfigMode
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.L7ConfigShrink) {
		body["L7Config"] = request.L7ConfigShrink
	}

	if !dara.IsNil(request.L7ProxyDomainAutomaticPrefix) {
		body["L7ProxyDomainAutomaticPrefix"] = request.L7ProxyDomainAutomaticPrefix
	}

	if !dara.IsNil(request.L7ProxyDomainCustom) {
		body["L7ProxyDomainCustom"] = request.L7ProxyDomainCustom
	}

	if !dara.IsNil(request.L7ProxyDomainPrivate) {
		body["L7ProxyDomainPrivate"] = request.L7ProxyDomainPrivate
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PortRanges) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.UnauthorizedAccessConfigShrink) {
		body["UnauthorizedAccessConfig"] = request.UnauthorizedAccessConfigShrink
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateAccessApplication"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a private access application under the current Alibaba Cloud account.
//
// @param request - UpdatePrivateAccessApplicationRequest
//
// @return UpdatePrivateAccessApplicationResponse
func (client *Client) UpdatePrivateAccessApplication(request *UpdatePrivateAccessApplicationRequest) (_result *UpdatePrivateAccessApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePrivateAccessApplicationResponse{}
	_body, _err := client.UpdatePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the Layer 7 access switch for an internal-facing application.
//
// @param request - UpdatePrivateAccessApplicationL7SwitchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessApplicationL7SwitchResponse
func (client *Client) UpdatePrivateAccessApplicationL7SwitchWithOptions(request *UpdatePrivateAccessApplicationL7SwitchRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateAccessApplicationL7SwitchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.DevTagMarkStatus) {
		body["DevTagMarkStatus"] = request.DevTagMarkStatus
	}

	if !dara.IsNil(request.DownloadAuditStatus) {
		body["DownloadAuditStatus"] = request.DownloadAuditStatus
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.PortRanges) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !dara.IsNil(request.SrcIpMarkStatus) {
		body["SrcIpMarkStatus"] = request.SrcIpMarkStatus
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeoutSec) {
		body["TimeoutSec"] = request.TimeoutSec
	}

	if !dara.IsNil(request.UserMarkStatus) {
		body["UserMarkStatus"] = request.UserMarkStatus
	}

	if !dara.IsNil(request.ZeroTrustStatus) {
		body["ZeroTrustStatus"] = request.ZeroTrustStatus
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateAccessApplicationL7Switch"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateAccessApplicationL7SwitchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the Layer 7 access switch for an internal-facing application.
//
// @param request - UpdatePrivateAccessApplicationL7SwitchRequest
//
// @return UpdatePrivateAccessApplicationL7SwitchResponse
func (client *Client) UpdatePrivateAccessApplicationL7Switch(request *UpdatePrivateAccessApplicationL7SwitchRequest) (_result *UpdatePrivateAccessApplicationL7SwitchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePrivateAccessApplicationL7SwitchResponse{}
	_body, _err := client.UpdatePrivateAccessApplicationL7SwitchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update an internal network access policy for your Alibaba Cloud account.
//
// @param request - UpdatePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessPolicyResponse
func (client *Client) UpdatePrivateAccessPolicyWithOptions(request *UpdatePrivateAccessPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdatePrivateAccessPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.ApplicationType) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !dara.IsNil(request.CustomUserAttributes) {
		bodyFlat["CustomUserAttributes"] = request.CustomUserAttributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceAttributeAction) {
		body["DeviceAttributeAction"] = request.DeviceAttributeAction
	}

	if !dara.IsNil(request.DeviceAttributeId) {
		body["DeviceAttributeId"] = request.DeviceAttributeId
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PolicyAction) {
		body["PolicyAction"] = request.PolicyAction
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.TriggerTemplateId) {
		body["TriggerTemplateId"] = request.TriggerTemplateId
	}

	if !dara.IsNil(request.TrustedProcessGroupIds) {
		bodyFlat["TrustedProcessGroupIds"] = request.TrustedProcessGroupIds
	}

	if !dara.IsNil(request.TrustedProcessStatus) {
		body["TrustedProcessStatus"] = request.TrustedProcessStatus
	}

	if !dara.IsNil(request.TrustedSoftwareIds) {
		bodyFlat["TrustedSoftwareIds"] = request.TrustedSoftwareIds
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.UserGroupMode) {
		body["UserGroupMode"] = request.UserGroupMode
	}

	if !dara.IsNil(request.ValidFrom) {
		body["ValidFrom"] = request.ValidFrom
	}

	if !dara.IsNil(request.ValidTimeStatus) {
		body["ValidTimeStatus"] = request.ValidTimeStatus
	}

	if !dara.IsNil(request.ValidUntil) {
		body["ValidUntil"] = request.ValidUntil
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePrivateAccessPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update an internal network access policy for your Alibaba Cloud account.
//
// @param request - UpdatePrivateAccessPolicyRequest
//
// @return UpdatePrivateAccessPolicyResponse
func (client *Client) UpdatePrivateAccessPolicy(request *UpdatePrivateAccessPolicyRequest) (_result *UpdatePrivateAccessPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePrivateAccessPolicyResponse{}
	_body, _err := client.UpdatePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a software prohibition policy.
//
// @param request - UpdateProhibitedPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProhibitedPolicyResponse
func (client *Client) UpdateProhibitedPolicyWithOptions(request *UpdateProhibitedPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateProhibitedPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowReport) {
		body["AllowReport"] = request.AllowReport
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ForceKill) {
		body["ForceKill"] = request.ForceKill
	}

	if !dara.IsNil(request.MainButtonTextCh) {
		body["MainButtonTextCh"] = request.MainButtonTextCh
	}

	if !dara.IsNil(request.MainButtonTextEn) {
		body["MainButtonTextEn"] = request.MainButtonTextEn
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MinorButtonTextCh) {
		body["MinorButtonTextCh"] = request.MinorButtonTextCh
	}

	if !dara.IsNil(request.MinorButtonTextEn) {
		body["MinorButtonTextEn"] = request.MinorButtonTextEn
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicyType) {
		body["PolicyType"] = request.PolicyType
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.PromptCh) {
		body["PromptCh"] = request.PromptCh
	}

	if !dara.IsNil(request.PromptEn) {
		body["PromptEn"] = request.PromptEn
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.SoftwareIds) {
		bodyFlat["SoftwareIds"] = request.SoftwareIds
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.TitleCh) {
		body["TitleCh"] = request.TitleCh
	}

	if !dara.IsNil(request.TitleEn) {
		body["TitleEn"] = request.TitleEn
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProhibitedPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProhibitedPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a software prohibition policy.
//
// @param request - UpdateProhibitedPolicyRequest
//
// @return UpdateProhibitedPolicyResponse
func (client *Client) UpdateProhibitedPolicy(request *UpdateProhibitedPolicyRequest) (_result *UpdateProhibitedPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProhibitedPolicyResponse{}
	_body, _err := client.UpdateProhibitedPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom prohibited software entry.
//
// @param request - UpdateProhibitedSoftwareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProhibitedSoftwareResponse
func (client *Client) UpdateProhibitedSoftwareWithOptions(request *UpdateProhibitedSoftwareRequest, runtime *dara.RuntimeOptions) (_result *UpdateProhibitedSoftwareResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.LinuxProcesses) {
		bodyFlat["LinuxProcesses"] = request.LinuxProcesses
	}

	if !dara.IsNil(request.MacOSProcesses) {
		bodyFlat["MacOSProcesses"] = request.MacOSProcesses
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SoftwareId) {
		body["SoftwareId"] = request.SoftwareId
	}

	if !dara.IsNil(request.TagIds) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !dara.IsNil(request.WindowsProcesses) {
		bodyFlat["WindowsProcesses"] = request.WindowsProcesses
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProhibitedSoftware"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProhibitedSoftwareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom prohibited software entry.
//
// @param request - UpdateProhibitedSoftwareRequest
//
// @return UpdateProhibitedSoftwareResponse
func (client *Client) UpdateProhibitedSoftware(request *UpdateProhibitedSoftwareRequest) (_result *UpdateProhibitedSoftwareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProhibitedSoftwareResponse{}
	_body, _err := client.UpdateProhibitedSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom prohibited software tag.
//
// @param request - UpdateProhibitedTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProhibitedTagResponse
func (client *Client) UpdateProhibitedTagWithOptions(request *UpdateProhibitedTagRequest, runtime *dara.RuntimeOptions) (_result *UpdateProhibitedTagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.TagId) {
		body["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProhibitedTag"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProhibitedTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom prohibited software tag.
//
// @param request - UpdateProhibitedTagRequest
//
// @return UpdateProhibitedTagResponse
func (client *Client) UpdateProhibitedTag(request *UpdateProhibitedTagRequest) (_result *UpdateProhibitedTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProhibitedTagResponse{}
	_body, _err := client.UpdateProhibitedTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a device registration policy for your Alibaba Cloud account.
//
// @param tmpReq - UpdateRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRegistrationPolicyResponse
func (client *Client) UpdateRegistrationPolicyWithOptions(tmpReq *UpdateRegistrationPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateRegistrationPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateRegistrationPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CompanyLimitCount) {
		request.CompanyLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompanyLimitCount, dara.String("CompanyLimitCount"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.PersonalLimitCount) {
		request.PersonalLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonalLimitCount, dara.String("PersonalLimitCount"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CompanyLimitCountShrink) {
		body["CompanyLimitCount"] = request.CompanyLimitCountShrink
	}

	if !dara.IsNil(request.CompanyLimitType) {
		body["CompanyLimitType"] = request.CompanyLimitType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PersonalLimitCountShrink) {
		body["PersonalLimitCount"] = request.PersonalLimitCountShrink
	}

	if !dara.IsNil(request.PersonalLimitType) {
		body["PersonalLimitType"] = request.PersonalLimitType
	}

	if !dara.IsNil(request.PolicyId) {
		body["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRegistrationPolicy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a device registration policy for your Alibaba Cloud account.
//
// @param request - UpdateRegistrationPolicyRequest
//
// @return UpdateRegistrationPolicyResponse
func (client *Client) UpdateRegistrationPolicy(request *UpdateRegistrationPolicyRequest) (_result *UpdateRegistrationPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRegistrationPolicyResponse{}
	_body, _err := client.UpdateRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the current processing status and conclusion of a specified risk event.
//
// Description:
//
// ## Request description
//
// - This operation allows you to update the processing status of a specific risk event under your Alibaba Cloud account.
//
// - When `Status` is set to `Processed`, you must provide the `RiskConfirm` parameter to specify the manually confirmed risk conclusion.
//
// - If `Status` is `Unprocess` or `Processing`, do not include the `RiskConfirm` parameter.
//
// - The `RiskScene` parameter is optional. If not provided, the system automatically populates it based on `RiskId`.
//
// - The `RiskConfirmDesc` field provides additional explanation or remarks for the processing decision. The length must be 1 to 128 characters.
//
// @param request - UpdateRiskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRiskStatusResponse
func (client *Client) UpdateRiskStatusWithOptions(request *UpdateRiskStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateRiskStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RiskConfirm) {
		body["RiskConfirm"] = request.RiskConfirm
	}

	if !dara.IsNil(request.RiskConfirmDesc) {
		body["RiskConfirmDesc"] = request.RiskConfirmDesc
	}

	if !dara.IsNil(request.RiskId) {
		body["RiskId"] = request.RiskId
	}

	if !dara.IsNil(request.RiskScene) {
		body["RiskScene"] = request.RiskScene
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRiskStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRiskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the current processing status and conclusion of a specified risk event.
//
// Description:
//
// ## Request description
//
// - This operation allows you to update the processing status of a specific risk event under your Alibaba Cloud account.
//
// - When `Status` is set to `Processed`, you must provide the `RiskConfirm` parameter to specify the manually confirmed risk conclusion.
//
// - If `Status` is `Unprocess` or `Processing`, do not include the `RiskConfirm` parameter.
//
// - The `RiskScene` parameter is optional. If not provided, the system automatically populates it based on `RiskId`.
//
// - The `RiskConfirmDesc` field provides additional explanation or remarks for the processing decision. The length must be 1 to 128 characters.
//
// @param request - UpdateRiskStatusRequest
//
// @return UpdateRiskStatusResponse
func (client *Client) UpdateRiskStatus(request *UpdateRiskStatusRequest) (_result *UpdateRiskStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRiskStatusResponse{}
	_body, _err := client.UpdateRiskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of uninstall applications in batches under the current Alibaba Cloud account.
//
// @param request - UpdateUninstallApplicationsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUninstallApplicationsStatusResponse
func (client *Client) UpdateUninstallApplicationsStatusWithOptions(request *UpdateUninstallApplicationsStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUninstallApplicationsStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationIds) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUninstallApplicationsStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUninstallApplicationsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of uninstall applications in batches under the current Alibaba Cloud account.
//
// @param request - UpdateUninstallApplicationsStatusRequest
//
// @return UpdateUninstallApplicationsStatusResponse
func (client *Client) UpdateUninstallApplicationsStatus(request *UpdateUninstallApplicationsStatusRequest) (_result *UpdateUninstallApplicationsStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUninstallApplicationsStatusResponse{}
	_body, _err := client.UpdateUninstallApplicationsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the sharing status of devices for multiple enterprise users.
//
// @param request - UpdateUserDevicesSharingStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDevicesSharingStatusResponse
func (client *Client) UpdateUserDevicesSharingStatusWithOptions(request *UpdateUserDevicesSharingStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDevicesSharingStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	if !dara.IsNil(request.SharingStatus) {
		body["SharingStatus"] = request.SharingStatus
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserDevicesSharingStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDevicesSharingStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the sharing status of devices for multiple enterprise users.
//
// @param request - UpdateUserDevicesSharingStatusRequest
//
// @return UpdateUserDevicesSharingStatusResponse
func (client *Client) UpdateUserDevicesSharingStatus(request *UpdateUserDevicesSharingStatusRequest) (_result *UpdateUserDevicesSharingStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserDevicesSharingStatusResponse{}
	_body, _err := client.UpdateUserDevicesSharingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update the status of endpoint devices for your Alibaba Cloud account.
//
// @param request - UpdateUserDevicesStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDevicesStatusResponse
func (client *Client) UpdateUserDevicesStatusWithOptions(request *UpdateUserDevicesStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserDevicesStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceAction) {
		body["DeviceAction"] = request.DeviceAction
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTags) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserDevicesStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserDevicesStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update the status of endpoint devices for your Alibaba Cloud account.
//
// @param request - UpdateUserDevicesStatusRequest
//
// @return UpdateUserDevicesStatusResponse
func (client *Client) UpdateUserDevicesStatus(request *UpdateUserDevicesStatusRequest) (_result *UpdateUserDevicesStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserDevicesStatusResponse{}
	_body, _err := client.UpdateUserDevicesStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a user group under the current Alibaba Cloud account.
//
// @param request - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Attributes) {
		bodyFlat["Attributes"] = request.Attributes
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ModifyType) {
		body["ModifyType"] = request.ModifyType
	}

	if !dara.IsNil(request.UserGroupId) {
		body["UserGroupId"] = request.UserGroupId
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUserGroup"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a user group under the current Alibaba Cloud account.
//
// @param request - UpdateUserGroupRequest
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of users in your Alibaba Cloud account.
//
// @param request - UpdateUsersStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUsersStatusResponse
func (client *Client) UpdateUsersStatusWithOptions(request *UpdateUsersStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateUsersStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SaseUserIds) {
		query["SaseUserIds"] = request.SaseUserIds
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUsersStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUsersStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of users in your Alibaba Cloud account.
//
// @param request - UpdateUsersStatusRequest
//
// @return UpdateUsersStatusResponse
func (client *Client) UpdateUsersStatus(request *UpdateUsersStatusRequest) (_result *UpdateUsersStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUsersStatusResponse{}
	_body, _err := client.UpdateUsersStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Quarantines or trusts a virus file on a specified user terminal device. DevTag, FilePath, and FileMd5 together identify a virus file record. The call fails if the record does not exist. Quarantine is an asynchronous operation. After the server creates a disposal task, the user terminal device pulls and executes it. The same virus file record can only be disposed of once within one minute.
//
// @param request - UpdateVirusFileStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVirusFileStatusResponse
func (client *Client) UpdateVirusFileStatusWithOptions(request *UpdateVirusFileStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateVirusFileStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DevTag) {
		body["DevTag"] = request.DevTag
	}

	if !dara.IsNil(request.FileMd5) {
		body["FileMd5"] = request.FileMd5
	}

	if !dara.IsNil(request.FilePath) {
		body["FilePath"] = request.FilePath
	}

	if !dara.IsNil(request.Operation) {
		body["Operation"] = request.Operation
	}

	if !dara.IsNil(request.VirusType) {
		body["VirusType"] = request.VirusType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVirusFileStatus"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVirusFileStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Quarantines or trusts a virus file on a specified user terminal device. DevTag, FilePath, and FileMd5 together identify a virus file record. The call fails if the record does not exist. Quarantine is an asynchronous operation. After the server creates a disposal task, the user terminal device pulls and executes it. The same virus file record can only be disposed of once within one minute.
//
// @param request - UpdateVirusFileStatusRequest
//
// @return UpdateVirusFileStatusResponse
func (client *Client) UpdateVirusFileStatus(request *UpdateVirusFileStatusRequest) (_result *UpdateVirusFileStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVirusFileStatusResponse{}
	_body, _err := client.UpdateVirusFileStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the anti-virus global configuration for the current Alibaba Cloud account. The four configuration items are treated as a whole and are entirely overwritten with each call. Therefore, pass in the complete configuration with each call: set VirusFileUpload to false, UploadFileSuffixBlacklist to empty, and UploadFileMaxSize and UploadFileMaxSpeed to 0 (no limit). After VirusFileUpload is changed, the virus file upload module switch is synchronously updated, which affects whether cloud-based STS tokens are issued to user terminal devices.
//
// @param request - UpdateVirusScanGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVirusScanGlobalConfigResponse
func (client *Client) UpdateVirusScanGlobalConfigWithOptions(request *UpdateVirusScanGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateVirusScanGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.UploadFileMaxSize) {
		body["UploadFileMaxSize"] = request.UploadFileMaxSize
	}

	if !dara.IsNil(request.UploadFileMaxSpeed) {
		body["UploadFileMaxSpeed"] = request.UploadFileMaxSpeed
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UploadFileSuffixBlacklist) {
		bodyFlat["UploadFileSuffixBlacklist"] = request.UploadFileSuffixBlacklist
	}

	if !dara.IsNil(request.VirusFileUpload) {
		body["VirusFileUpload"] = request.VirusFileUpload
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVirusScanGlobalConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVirusScanGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the anti-virus global configuration for the current Alibaba Cloud account. The four configuration items are treated as a whole and are entirely overwritten with each call. Therefore, pass in the complete configuration with each call: set VirusFileUpload to false, UploadFileSuffixBlacklist to empty, and UploadFileMaxSize and UploadFileMaxSpeed to 0 (no limit). After VirusFileUpload is changed, the virus file upload module switch is synchronously updated, which affects whether cloud-based STS tokens are issued to user terminal devices.
//
// @param request - UpdateVirusScanGlobalConfigRequest
//
// @return UpdateVirusScanGlobalConfigResponse
func (client *Client) UpdateVirusScanGlobalConfig(request *UpdateVirusScanGlobalConfigRequest) (_result *UpdateVirusScanGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVirusScanGlobalConfigResponse{}
	_body, _err := client.UpdateVirusScanGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified scheduled virus scan policy. The Whitelist parameter performs a full overwrite, meaning the provided list replaces the existing exception user list of the policy.
//
// @param request - UpdateVirusScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVirusScanScheduledStrategyResponse
func (client *Client) UpdateVirusScanScheduledStrategyWithOptions(request *UpdateVirusScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateVirusScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HighRiskOperation) {
		body["HighRiskOperation"] = request.HighRiskOperation
	}

	if !dara.IsNil(request.LowRiskOperation) {
		body["LowRiskOperation"] = request.LowRiskOperation
	}

	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.MaxCpuUsage) {
		body["MaxCpuUsage"] = request.MaxCpuUsage
	}

	if !dara.IsNil(request.MidRiskOperation) {
		body["MidRiskOperation"] = request.MidRiskOperation
	}

	if !dara.IsNil(request.PerformanceMode) {
		body["PerformanceMode"] = request.PerformanceMode
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ScanBeginTime) {
		body["ScanBeginTime"] = request.ScanBeginTime
	}

	if !dara.IsNil(request.ScanEndTime) {
		body["ScanEndTime"] = request.ScanEndTime
	}

	if !dara.IsNil(request.ScanFrequency) {
		body["ScanFrequency"] = request.ScanFrequency
	}

	if !dara.IsNil(request.ScanInterval) {
		body["ScanInterval"] = request.ScanInterval
	}

	if !dara.IsNil(request.ScanMode) {
		body["ScanMode"] = request.ScanMode
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.ScanPath) {
		bodyFlat["ScanPath"] = request.ScanPath
	}

	if !dara.IsNil(request.ScanTargets) {
		bodyFlat["ScanTargets"] = request.ScanTargets
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategyDescription) {
		body["StrategyDescription"] = request.StrategyDescription
	}

	if !dara.IsNil(request.StrategyId) {
		body["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.StrategyName) {
		body["StrategyName"] = request.StrategyName
	}

	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVirusScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVirusScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified scheduled virus scan policy. The Whitelist parameter performs a full overwrite, meaning the provided list replaces the existing exception user list of the policy.
//
// @param request - UpdateVirusScanScheduledStrategyRequest
//
// @return UpdateVirusScanScheduledStrategyResponse
func (client *Client) UpdateVirusScanScheduledStrategy(request *UpdateVirusScanScheduledStrategyRequest) (_result *UpdateVirusScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVirusScanScheduledStrategyResponse{}
	_body, _err := client.UpdateVirusScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the global vulnerability scanning configuration for the current Alibaba Cloud account and returns the complete updated configuration.
//
// @param tmpReq - UpdateVulScanGlobalConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVulScanGlobalConfigResponse
func (client *Client) UpdateVulScanGlobalConfigWithOptions(tmpReq *UpdateVulScanGlobalConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateVulScanGlobalConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateVulScanGlobalConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WuyingVulFixConfig) {
		request.WuyingVulFixConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WuyingVulFixConfig, dara.String("WuyingVulFixConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxDownloadSpeed) {
		body["MaxDownloadSpeed"] = request.MaxDownloadSpeed
	}

	if !dara.IsNil(request.WuyingVulFixConfigShrink) {
		body["WuyingVulFixConfig"] = request.WuyingVulFixConfigShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVulScanGlobalConfig"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVulScanGlobalConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the global vulnerability scanning configuration for the current Alibaba Cloud account and returns the complete updated configuration.
//
// @param request - UpdateVulScanGlobalConfigRequest
//
// @return UpdateVulScanGlobalConfigResponse
func (client *Client) UpdateVulScanGlobalConfig(request *UpdateVulScanGlobalConfigRequest) (_result *UpdateVulScanGlobalConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVulScanGlobalConfigResponse{}
	_body, _err := client.UpdateVulScanGlobalConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified vulnerability scheduled scan policy and returns the complete updated configuration.
//
// @param request - UpdateVulScanScheduledStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVulScanScheduledStrategyResponse
func (client *Client) UpdateVulScanScheduledStrategyWithOptions(request *UpdateVulScanScheduledStrategyRequest, runtime *dara.RuntimeOptions) (_result *UpdateVulScanScheduledStrategyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MatchMode) {
		body["MatchMode"] = request.MatchMode
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ScanBeginTime) {
		body["ScanBeginTime"] = request.ScanBeginTime
	}

	if !dara.IsNil(request.ScanEndTime) {
		body["ScanEndTime"] = request.ScanEndTime
	}

	if !dara.IsNil(request.ScanFrequency) {
		body["ScanFrequency"] = request.ScanFrequency
	}

	if !dara.IsNil(request.ScanInterval) {
		body["ScanInterval"] = request.ScanInterval
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.StrategyDescription) {
		body["StrategyDescription"] = request.StrategyDescription
	}

	if !dara.IsNil(request.StrategyId) {
		body["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.StrategyName) {
		body["StrategyName"] = request.StrategyName
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupIds) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !dara.IsNil(request.Whitelist) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVulScanScheduledStrategy"),
		Version:     dara.String("2023-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVulScanScheduledStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the configuration of a specified vulnerability scheduled scan policy and returns the complete updated configuration.
//
// @param request - UpdateVulScanScheduledStrategyRequest
//
// @return UpdateVulScanScheduledStrategyResponse
func (client *Client) UpdateVulScanScheduledStrategy(request *UpdateVulScanScheduledStrategyRequest) (_result *UpdateVulScanScheduledStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateVulScanScheduledStrategyResponse{}
	_body, _err := client.UpdateVulScanScheduledStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
