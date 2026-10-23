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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("csasedr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Queries security alert types.
//
// @param request - DescribeNsasSuspEventTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNsasSuspEventTypeResponse
func (client *Client) DescribeNsasSuspEventTypeWithOptions(request *DescribeNsasSuspEventTypeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNsasSuspEventTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContainerFieldName) {
		query["ContainerFieldName"] = request.ContainerFieldName
	}

	if !dara.IsNil(request.ContainerFieldValue) {
		query["ContainerFieldValue"] = request.ContainerFieldValue
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.Uuids) {
		query["Uuids"] = request.Uuids
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNsasSuspEventType"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNsasSuspEventTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries security alert types.
//
// @param request - DescribeNsasSuspEventTypeRequest
//
// @return DescribeNsasSuspEventTypeResponse
func (client *Client) DescribeNsasSuspEventType(request *DescribeNsasSuspEventTypeRequest) (_result *DescribeNsasSuspEventTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNsasSuspEventTypeResponse{}
	_body, _err := client.DescribeNsasSuspEventTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the handling operations, operability, and whitelist rule configurations for a specified endpoint alert in its current state.
//
// Description:
//
// This operation queries the handling operations that can be displayed for a specified endpoint alert. When initiating a subsequent response, you must select the operation items returned in this response where UserCanOperate is true, and submit parameters as required by the response operation. The query result does not guarantee that the subsequent response will succeed.
//
// The actual returned results are affected by feature switches, canary releases, dynamic allowlists, alert templates, alert statuses, and filter conditions. The list may be empty. The examples are used to illustrate field formats or rule combinations and do not represent that a specific alert will definitely return these values. MappingMarkFields and its subfields are generic compatibility structures. The current query pipeline does not populate their elements, and you should not infer that this operation supports the corresponding capabilities based on them.
//
// @param request - DescribeSecurityEventOperationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityEventOperationsResponse
func (client *Client) DescribeSecurityEventOperationsWithOptions(request *DescribeSecurityEventOperationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SecurityEventId) {
		query["SecurityEventId"] = request.SecurityEventId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSecurityEventOperations"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the handling operations, operability, and whitelist rule configurations for a specified endpoint alert in its current state.
//
// Description:
//
// This operation queries the handling operations that can be displayed for a specified endpoint alert. When initiating a subsequent response, you must select the operation items returned in this response where UserCanOperate is true, and submit parameters as required by the response operation. The query result does not guarantee that the subsequent response will succeed.
//
// The actual returned results are affected by feature switches, canary releases, dynamic allowlists, alert templates, alert statuses, and filter conditions. The list may be empty. The examples are used to illustrate field formats or rule combinations and do not represent that a specific alert will definitely return these values. MappingMarkFields and its subfields are generic compatibility structures. The current query pipeline does not populate their elements, and you should not infer that this operation supports the corresponding capabilities based on them.
//
// @param request - DescribeSecurityEventOperationsRequest
//
// @return DescribeSecurityEventOperationsResponse
func (client *Client) DescribeSecurityEventOperations(request *DescribeSecurityEventOperationsRequest) (_result *DescribeSecurityEventOperationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSecurityEventOperationsResponse{}
	_body, _err := client.DescribeSecurityEventOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries endpoint security alerts and supports filtering by alert, endpoint, and processing status.
//
// @param request - DescribeSuspEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSuspEventsResponse
func (client *Client) DescribeSuspEventsWithOptions(request *DescribeSuspEventsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSuspEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlarmEventName) {
		query["AlarmEventName"] = request.AlarmEventName
	}

	if !dara.IsNil(request.AlarmUniqueInfo) {
		query["AlarmUniqueInfo"] = request.AlarmUniqueInfo
	}

	if !dara.IsNil(request.AssetsTypeList) {
		query["AssetsTypeList"] = request.AssetsTypeList
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ContainerFieldName) {
		query["ContainerFieldName"] = request.ContainerFieldName
	}

	if !dara.IsNil(request.ContainerFieldValue) {
		query["ContainerFieldValue"] = request.ContainerFieldValue
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Dealed) {
		query["Dealed"] = request.Dealed
	}

	if !dara.IsNil(request.Department) {
		query["Department"] = request.Department
	}

	if !dara.IsNil(request.EventNames) {
		query["EventNames"] = request.EventNames
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.HostName) {
		query["HostName"] = request.HostName
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.InnerIp) {
		query["InnerIp"] = request.InnerIp
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Levels) {
		query["Levels"] = request.Levels
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OperateErrorCodeList) {
		query["OperateErrorCodeList"] = request.OperateErrorCodeList
	}

	if !dara.IsNil(request.OperateTimeEnd) {
		query["OperateTimeEnd"] = request.OperateTimeEnd
	}

	if !dara.IsNil(request.OperateTimeStart) {
		query["OperateTimeStart"] = request.OperateTimeStart
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentEventTypes) {
		query["ParentEventTypes"] = request.ParentEventTypes
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SortColumn) {
		query["SortColumn"] = request.SortColumn
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	if !dara.IsNil(request.SrcIp) {
		query["SrcIp"] = request.SrcIp
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StrictMode) {
		query["StrictMode"] = request.StrictMode
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	if !dara.IsNil(request.TimeEnd) {
		query["TimeEnd"] = request.TimeEnd
	}

	if !dara.IsNil(request.TimeStart) {
		query["TimeStart"] = request.TimeStart
	}

	if !dara.IsNil(request.UniqueInfo) {
		query["UniqueInfo"] = request.UniqueInfo
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	if !dara.IsNil(request.Uuids) {
		query["Uuids"] = request.Uuids
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TacticId) {
		body["TacticId"] = request.TacticId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSuspEvents"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries endpoint security alerts and supports filtering by alert, endpoint, and processing status.
//
// @param request - DescribeSuspEventsRequest
//
// @return DescribeSuspEventsResponse
func (client *Client) DescribeSuspEvents(request *DescribeSuspEventsRequest) (_result *DescribeSuspEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeSuspEventsResponse{}
	_body, _err := client.DescribeSuspEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Handles security alert events.
//
// @param request - HandleSecurityEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return HandleSecurityEventsResponse
func (client *Client) HandleSecurityEventsWithOptions(request *HandleSecurityEventsRequest, runtime *dara.RuntimeOptions) (_result *HandleSecurityEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MarkBatch) {
		query["MarkBatch"] = request.MarkBatch
	}

	if !dara.IsNil(request.MarkMissParam) {
		query["MarkMissParam"] = request.MarkMissParam
	}

	if !dara.IsNil(request.OperationCode) {
		query["OperationCode"] = request.OperationCode
	}

	if !dara.IsNil(request.OperationParams) {
		query["OperationParams"] = request.OperationParams
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	if !dara.IsNil(request.SecurityEventIds) {
		query["SecurityEventIds"] = request.SecurityEventIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("HandleSecurityEvents"),
		Version:     dara.String("2024-11-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Handles security alert events.
//
// @param request - HandleSecurityEventsRequest
//
// @return HandleSecurityEventsResponse
func (client *Client) HandleSecurityEvents(request *HandleSecurityEventsRequest) (_result *HandleSecurityEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &HandleSecurityEventsResponse{}
	_body, _err := client.HandleSecurityEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
