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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("nis"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Initiates a task for analyzing network reachability.
//
// Description:
//
// You can call this operation to initiate a task for analyzing network reachability by specifying only the information about the source and destination. You do not need to create a network path for reachability analysis. The analysis result is not recorded in the system. If you want to record the path parameters and analysis result in the Network Intelligence Service (NIS) console, we recommend that you call the **createNetworkReachableAnalysis*	- operation.
//
// @param request - CreateAndAnalyzeNetworkPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAndAnalyzeNetworkPathResponse
func (client *Client) CreateAndAnalyzeNetworkPathWithOptions(request *CreateAndAnalyzeNetworkPathRequest, runtime *dara.RuntimeOptions) (_result *CreateAndAnalyzeNetworkPathResponse, _err error) {
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
		Action:      dara.String("CreateAndAnalyzeNetworkPath"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAndAnalyzeNetworkPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates a task for analyzing network reachability.
//
// Description:
//
// You can call this operation to initiate a task for analyzing network reachability by specifying only the information about the source and destination. You do not need to create a network path for reachability analysis. The analysis result is not recorded in the system. If you want to record the path parameters and analysis result in the Network Intelligence Service (NIS) console, we recommend that you call the **createNetworkReachableAnalysis*	- operation.
//
// @param request - CreateAndAnalyzeNetworkPathRequest
//
// @return CreateAndAnalyzeNetworkPathResponse
func (client *Client) CreateAndAnalyzeNetworkPath(request *CreateAndAnalyzeNetworkPathRequest) (_result *CreateAndAnalyzeNetworkPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAndAnalyzeNetworkPathResponse{}
	_body, _err := client.CreateAndAnalyzeNetworkPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a network path in the cloud for reachability analysis.
//
// Description:
//
//	  You can call the **CreateNetworkPath*	- operation to create network paths in multiple networking scenarios and between multiple resources. After a path is created, the path parameters are saved for repeated analysis.
//
//		- You can create up to 100 network paths within one Alibaba Cloud account.
//
// @param request - CreateNetworkPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkPathResponse
func (client *Client) CreateNetworkPathWithOptions(request *CreateNetworkPathRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkPathDescription) {
		query["NetworkPathDescription"] = request.NetworkPathDescription
	}

	if !dara.IsNil(request.NetworkPathName) {
		query["NetworkPathName"] = request.NetworkPathName
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceIpAddress) {
		query["SourceIpAddress"] = request.SourceIpAddress
	}

	if !dara.IsNil(request.SourcePort) {
		query["SourcePort"] = request.SourcePort
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetIpAddress) {
		query["TargetIpAddress"] = request.TargetIpAddress
	}

	if !dara.IsNil(request.TargetPort) {
		query["TargetPort"] = request.TargetPort
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkPath"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a network path in the cloud for reachability analysis.
//
// Description:
//
//	  You can call the **CreateNetworkPath*	- operation to create network paths in multiple networking scenarios and between multiple resources. After a path is created, the path parameters are saved for repeated analysis.
//
//		- You can create up to 100 network paths within one Alibaba Cloud account.
//
// @param request - CreateNetworkPathRequest
//
// @return CreateNetworkPathResponse
func (client *Client) CreateNetworkPath(request *CreateNetworkPathRequest) (_result *CreateNetworkPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkPathResponse{}
	_body, _err := client.CreateNetworkPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task for analyzing network reachability.
//
// Description:
//
//	  The **CreateNetworkReachableAnalysis*	- operation is used to create a task for analyzing the reachability of the network path that is created by calling the **CreateNetworkPath*	- operation and record the analysis results.
//
//		- The **CreateNetworkReachableAnalysis*	- operation can be called to repeatedly analyze the reachability of a network path.
//
//		- You can create up to 1,000 reachability analysis records within one Alibaba Cloud account.
//
// @param request - CreateNetworkReachableAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetworkReachableAnalysisResponse
func (client *Client) CreateNetworkReachableAnalysisWithOptions(request *CreateNetworkReachableAnalysisRequest, runtime *dara.RuntimeOptions) (_result *CreateNetworkReachableAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkPathId) {
		query["NetworkPathId"] = request.NetworkPathId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetworkReachableAnalysis"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateNetworkReachableAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task for analyzing network reachability.
//
// Description:
//
//	  The **CreateNetworkReachableAnalysis*	- operation is used to create a task for analyzing the reachability of the network path that is created by calling the **CreateNetworkPath*	- operation and record the analysis results.
//
//		- The **CreateNetworkReachableAnalysis*	- operation can be called to repeatedly analyze the reachability of a network path.
//
//		- You can create up to 1,000 reachability analysis records within one Alibaba Cloud account.
//
// @param request - CreateNetworkReachableAnalysisRequest
//
// @return CreateNetworkReachableAnalysisResponse
func (client *Client) CreateNetworkReachableAnalysis(request *CreateNetworkReachableAnalysisRequest) (_result *CreateNetworkReachableAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateNetworkReachableAnalysisResponse{}
	_body, _err := client.CreateNetworkReachableAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a network path.
//
// @param tmpReq - DeleteNetworkPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkPathResponse
func (client *Client) DeleteNetworkPathWithOptions(tmpReq *DeleteNetworkPathRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkPathResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteNetworkPathShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetworkPathIds) {
		request.NetworkPathIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkPathIds, dara.String("NetworkPathIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkPathIdsShrink) {
		query["NetworkPathIds"] = request.NetworkPathIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkPath"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a network path.
//
// @param request - DeleteNetworkPathRequest
//
// @return DeleteNetworkPathResponse
func (client *Client) DeleteNetworkPath(request *DeleteNetworkPathRequest) (_result *DeleteNetworkPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkPathResponse{}
	_body, _err := client.DeleteNetworkPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task for analyzing network reachability.
//
// @param tmpReq - DeleteNetworkReachableAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNetworkReachableAnalysisResponse
func (client *Client) DeleteNetworkReachableAnalysisWithOptions(tmpReq *DeleteNetworkReachableAnalysisRequest, runtime *dara.RuntimeOptions) (_result *DeleteNetworkReachableAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteNetworkReachableAnalysisShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NetworkReachableAnalysisIds) {
		request.NetworkReachableAnalysisIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NetworkReachableAnalysisIds, dara.String("NetworkReachableAnalysisIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkReachableAnalysisIdsShrink) {
		query["NetworkReachableAnalysisIds"] = request.NetworkReachableAnalysisIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNetworkReachableAnalysis"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNetworkReachableAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task for analyzing network reachability.
//
// @param request - DeleteNetworkReachableAnalysisRequest
//
// @return DeleteNetworkReachableAnalysisResponse
func (client *Client) DeleteNetworkReachableAnalysis(request *DeleteNetworkReachableAnalysisRequest) (_result *DeleteNetworkReachableAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNetworkReachableAnalysisResponse{}
	_body, _err := client.DeleteNetworkReachableAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除报告
//
// @param request - DeleteNisInspectionReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNisInspectionReportResponse
func (client *Client) DeleteNisInspectionReportWithOptions(request *DeleteNisInspectionReportRequest, runtime *dara.RuntimeOptions) (_result *DeleteNisInspectionReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionReportId) {
		query["InspectionReportId"] = request.InspectionReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNisInspectionReport"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNisInspectionReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除报告
//
// @param request - DeleteNisInspectionReportRequest
//
// @return DeleteNisInspectionReportResponse
func (client *Client) DeleteNisInspectionReport(request *DeleteNisInspectionReportRequest) (_result *DeleteNisInspectionReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNisInspectionReportResponse{}
	_body, _err := client.DeleteNisInspectionReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除巡检任务
//
// @param request - DeleteNisInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNisInspectionTaskResponse
func (client *Client) DeleteNisInspectionTaskWithOptions(request *DeleteNisInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteNisInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionTaskId) {
		query["InspectionTaskId"] = request.InspectionTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNisInspectionTask"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteNisInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除巡检任务
//
// @param request - DeleteNisInspectionTaskRequest
//
// @return DeleteNisInspectionTaskResponse
func (client *Client) DeleteNisInspectionTask(request *DeleteNisInspectionTaskRequest) (_result *DeleteNisInspectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteNisInspectionTaskResponse{}
	_body, _err := client.DeleteNisInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 受影响资源列表
//
// @param request - DescribeNisInspectionRecommendationResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNisInspectionRecommendationResourcesResponse
func (client *Client) DescribeNisInspectionRecommendationResourcesWithOptions(request *DescribeNisInspectionRecommendationResourcesRequest, runtime *dara.RuntimeOptions) (_result *DescribeNisInspectionRecommendationResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionReportId) {
		query["InspectionReportId"] = request.InspectionReportId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RecommendationCode) {
		query["RecommendationCode"] = request.RecommendationCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNisInspectionRecommendationResources"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNisInspectionRecommendationResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 受影响资源列表
//
// @param request - DescribeNisInspectionRecommendationResourcesRequest
//
// @return DescribeNisInspectionRecommendationResourcesResponse
func (client *Client) DescribeNisInspectionRecommendationResources(request *DescribeNisInspectionRecommendationResourcesRequest) (_result *DescribeNisInspectionRecommendationResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNisInspectionRecommendationResourcesResponse{}
	_body, _err := client.DescribeNisInspectionRecommendationResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 报告巡检项列表
//
// @param tmpReq - DescribeNisInspectionReportCheckItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNisInspectionReportCheckItemsResponse
func (client *Client) DescribeNisInspectionReportCheckItemsWithOptions(tmpReq *DescribeNisInspectionReportCheckItemsRequest, runtime *dara.RuntimeOptions) (_result *DescribeNisInspectionReportCheckItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribeNisInspectionReportCheckItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ResourceType) {
		request.ResourceTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceType, dara.String("ResourceType"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RiskLevel) {
		request.RiskLevelShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RiskLevel, dara.String("RiskLevel"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryCode) {
		query["CategoryCode"] = request.CategoryCode
	}

	if !dara.IsNil(request.InspectionReportId) {
		query["InspectionReportId"] = request.InspectionReportId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceTypeShrink) {
		query["ResourceType"] = request.ResourceTypeShrink
	}

	if !dara.IsNil(request.RiskLevelShrink) {
		query["RiskLevel"] = request.RiskLevelShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNisInspectionReportCheckItems"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNisInspectionReportCheckItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 报告巡检项列表
//
// @param request - DescribeNisInspectionReportCheckItemsRequest
//
// @return DescribeNisInspectionReportCheckItemsResponse
func (client *Client) DescribeNisInspectionReportCheckItems(request *DescribeNisInspectionReportCheckItemsRequest) (_result *DescribeNisInspectionReportCheckItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNisInspectionReportCheckItemsResponse{}
	_body, _err := client.DescribeNisInspectionReportCheckItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询报告状态
//
// @param request - DescribeNisInspectionReportStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNisInspectionReportStatusResponse
func (client *Client) DescribeNisInspectionReportStatusWithOptions(request *DescribeNisInspectionReportStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeNisInspectionReportStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionReportId) {
		query["InspectionReportId"] = request.InspectionReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNisInspectionReportStatus"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNisInspectionReportStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询报告状态
//
// @param request - DescribeNisInspectionReportStatusRequest
//
// @return DescribeNisInspectionReportStatusResponse
func (client *Client) DescribeNisInspectionReportStatus(request *DescribeNisInspectionReportStatusRequest) (_result *DescribeNisInspectionReportStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNisInspectionReportStatusResponse{}
	_body, _err := client.DescribeNisInspectionReportStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 报告总结信息
//
// @param request - DescribeNisInspectionReportSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNisInspectionReportSummaryResponse
func (client *Client) DescribeNisInspectionReportSummaryWithOptions(request *DescribeNisInspectionReportSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeNisInspectionReportSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionReportId) {
		query["InspectionReportId"] = request.InspectionReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNisInspectionReportSummary"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNisInspectionReportSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 报告总结信息
//
// @param request - DescribeNisInspectionReportSummaryRequest
//
// @return DescribeNisInspectionReportSummaryResponse
func (client *Client) DescribeNisInspectionReportSummary(request *DescribeNisInspectionReportSummaryRequest) (_result *DescribeNisInspectionReportSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNisInspectionReportSummaryResponse{}
	_body, _err := client.DescribeNisInspectionReportSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询巡检任务
//
// @param request - DescribeNisInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNisInspectionTaskResponse
func (client *Client) DescribeNisInspectionTaskWithOptions(request *DescribeNisInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeNisInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionTaskId) {
		query["InspectionTaskId"] = request.InspectionTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNisInspectionTask"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNisInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询巡检任务
//
// @param request - DescribeNisInspectionTaskRequest
//
// @return DescribeNisInspectionTaskResponse
func (client *Client) DescribeNisInspectionTask(request *DescribeNisInspectionTaskRequest) (_result *DescribeNisInspectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNisInspectionTaskResponse{}
	_body, _err := client.DescribeNisInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询流量分析topN异步任务结果
//
// @param request - DescribeNisTrafficRankingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNisTrafficRankingResponse
func (client *Client) DescribeNisTrafficRankingWithOptions(request *DescribeNisTrafficRankingRequest, runtime *dara.RuntimeOptions) (_result *DescribeNisTrafficRankingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NisTrafficRankingId) {
		query["NisTrafficRankingId"] = request.NisTrafficRankingId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNisTrafficRanking"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNisTrafficRankingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流量分析topN异步任务结果
//
// @param request - DescribeNisTrafficRankingRequest
//
// @return DescribeNisTrafficRankingResponse
func (client *Client) DescribeNisTrafficRanking(request *DescribeNisTrafficRankingRequest) (_result *DescribeNisTrafficRankingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNisTrafficRankingResponse{}
	_body, _err := client.DescribeNisTrafficRankingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetInternetTuple is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the rankings of Internet traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Internet traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param tmpReq - GetInternetTupleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInternetTupleResponse
func (client *Client) GetInternetTupleWithOptions(tmpReq *GetInternetTupleRequest, runtime *dara.RuntimeOptions) (_result *GetInternetTupleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetInternetTupleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CloudIpList) {
		request.CloudIpListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CloudIpList, dara.String("CloudIpList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InstanceList) {
		request.InstanceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceList, dara.String("InstanceList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CloudIp) {
		query["CloudIp"] = request.CloudIp
	}

	if !dara.IsNil(request.CloudIpListShrink) {
		query["CloudIpList"] = request.CloudIpListShrink
	}

	if !dara.IsNil(request.CloudIsp) {
		query["CloudIsp"] = request.CloudIsp
	}

	if !dara.IsNil(request.CloudPort) {
		query["CloudPort"] = request.CloudPort
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceListShrink) {
		query["InstanceList"] = request.InstanceListShrink
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OtherCity) {
		query["OtherCity"] = request.OtherCity
	}

	if !dara.IsNil(request.OtherCountry) {
		query["OtherCountry"] = request.OtherCountry
	}

	if !dara.IsNil(request.OtherIp) {
		query["OtherIp"] = request.OtherIp
	}

	if !dara.IsNil(request.OtherIsp) {
		query["OtherIsp"] = request.OtherIsp
	}

	if !dara.IsNil(request.OtherPort) {
		query["OtherPort"] = request.OtherPort
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.TupleType) {
		query["TupleType"] = request.TupleType
	}

	if !dara.IsNil(request.UseMultiAccount) {
		query["UseMultiAccount"] = request.UseMultiAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInternetTuple"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInternetTupleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetInternetTuple is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the rankings of Internet traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Internet traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param request - GetInternetTupleRequest
//
// @return GetInternetTupleResponse
// Deprecated
func (client *Client) GetInternetTuple(request *GetInternetTupleRequest) (_result *GetInternetTupleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInternetTupleResponse{}
	_body, _err := client.GetInternetTupleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetNatTopN is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the real-time SNAT performance ranking of a NAT gateway.
//
// @param request - GetNatTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNatTopNResponse
func (client *Client) GetNatTopNWithOptions(request *GetNatTopNRequest, runtime *dara.RuntimeOptions) (_result *GetNatTopNResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Ip) {
		query["Ip"] = request.Ip
	}

	if !dara.IsNil(request.NatGatewayId) {
		query["NatGatewayId"] = request.NatGatewayId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNatTopN"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNatTopNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetNatTopN is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the real-time SNAT performance ranking of a NAT gateway.
//
// @param request - GetNatTopNRequest
//
// @return GetNatTopNResponse
// Deprecated
func (client *Client) GetNatTopN(request *GetNatTopNRequest) (_result *GetNatTopNResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNatTopNResponse{}
	_body, _err := client.GetNatTopNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the results of network reachability analysis.
//
// Description:
//
// *GetNetworkReachableAnalysis*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can query the state of the task for analyzing network reachability.
//
//   - The **init*	- state indicates that the task is in progress.
//
//   - The **finish*	- state indicates that the task is complete. In this state, you can obtain the analysis result.
//
// @param request - GetNetworkReachableAnalysisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNetworkReachableAnalysisResponse
func (client *Client) GetNetworkReachableAnalysisWithOptions(request *GetNetworkReachableAnalysisRequest, runtime *dara.RuntimeOptions) (_result *GetNetworkReachableAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NetworkReachableAnalysisId) {
		query["NetworkReachableAnalysisId"] = request.NetworkReachableAnalysisId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNetworkReachableAnalysis"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNetworkReachableAnalysisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the results of network reachability analysis.
//
// Description:
//
// *GetNetworkReachableAnalysis*	- is an asynchronous operation. After a request is sent, the system returns a request ID and runs the task in the background. You can query the state of the task for analyzing network reachability.
//
//   - The **init*	- state indicates that the task is in progress.
//
//   - The **finish*	- state indicates that the task is complete. In this state, you can obtain the analysis result.
//
// @param request - GetNetworkReachableAnalysisRequest
//
// @return GetNetworkReachableAnalysisResponse
func (client *Client) GetNetworkReachableAnalysis(request *GetNetworkReachableAnalysisRequest) (_result *GetNetworkReachableAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNetworkReachableAnalysisResponse{}
	_body, _err := client.GetNetworkReachableAnalysisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云网络指标趋势
//
// @param tmpReq - GetNisNetworkMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNisNetworkMetricsResponse
func (client *Client) GetNisNetworkMetricsWithOptions(tmpReq *GetNisNetworkMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetNisNetworkMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNisNetworkMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Dimensions) {
		request.DimensionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Dimensions, dara.String("Dimensions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.DimensionsShrink) {
		query["Dimensions"] = request.DimensionsShrink
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ScanBy) {
		query["ScanBy"] = request.ScanBy
	}

	if !dara.IsNil(request.StepMinutes) {
		query["StepMinutes"] = request.StepMinutes
	}

	if !dara.IsNil(request.UseCrossAccount) {
		query["UseCrossAccount"] = request.UseCrossAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNisNetworkMetrics"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNisNetworkMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云网络指标趋势
//
// @param request - GetNisNetworkMetricsRequest
//
// @return GetNisNetworkMetricsResponse
func (client *Client) GetNisNetworkMetrics(request *GetNisNetworkMetricsRequest) (_result *GetNisNetworkMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNisNetworkMetricsResponse{}
	_body, _err := client.GetNisNetworkMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云网络指标排名
//
// @param tmpReq - GetNisNetworkRankingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNisNetworkRankingResponse
func (client *Client) GetNisNetworkRankingWithOptions(tmpReq *GetNisNetworkRankingRequest, runtime *dara.RuntimeOptions) (_result *GetNisNetworkRankingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNisNetworkRankingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.UseCrossAccount) {
		query["UseCrossAccount"] = request.UseCrossAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNisNetworkRanking"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNisNetworkRankingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云网络指标排名
//
// @param request - GetNisNetworkRankingRequest
//
// @return GetNisNetworkRankingResponse
func (client *Client) GetNisNetworkRanking(request *GetNisNetworkRankingRequest) (_result *GetNisNetworkRankingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNisNetworkRankingResponse{}
	_body, _err := client.GetNisNetworkRankingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetTransitRouterFlowTopN is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the rankings of inter-region traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Inter-region traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param tmpReq - GetTransitRouterFlowTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTransitRouterFlowTopNResponse
func (client *Client) GetTransitRouterFlowTopNWithOptions(tmpReq *GetTransitRouterFlowTopNRequest, runtime *dara.RuntimeOptions) (_result *GetTransitRouterFlowTopNResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTransitRouterFlowTopNShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("AccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.BandwithPackageId) {
		query["BandwithPackageId"] = request.BandwithPackageId
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OtherIp) {
		query["OtherIp"] = request.OtherIp
	}

	if !dara.IsNil(request.OtherPort) {
		query["OtherPort"] = request.OtherPort
	}

	if !dara.IsNil(request.OtherRegion) {
		query["OtherRegion"] = request.OtherRegion
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.ThisIp) {
		query["ThisIp"] = request.ThisIp
	}

	if !dara.IsNil(request.ThisPort) {
		query["ThisPort"] = request.ThisPort
	}

	if !dara.IsNil(request.ThisRegion) {
		query["ThisRegion"] = request.ThisRegion
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.UseMultiAccount) {
		query["UseMultiAccount"] = request.UseMultiAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTransitRouterFlowTopN"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTransitRouterFlowTopNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetTransitRouterFlowTopN is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the rankings of inter-region traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Inter-region traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param request - GetTransitRouterFlowTopNRequest
//
// @return GetTransitRouterFlowTopNResponse
// Deprecated
func (client *Client) GetTransitRouterFlowTopN(request *GetTransitRouterFlowTopNRequest) (_result *GetTransitRouterFlowTopNResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTransitRouterFlowTopNResponse{}
	_body, _err := client.GetTransitRouterFlowTopNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetVbrFlowTopN is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the rankings of hybrid cloud traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Hybrid cloud traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param tmpReq - GetVbrFlowTopNRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVbrFlowTopNResponse
func (client *Client) GetVbrFlowTopNWithOptions(tmpReq *GetVbrFlowTopNRequest, runtime *dara.RuntimeOptions) (_result *GetVbrFlowTopNResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetVbrFlowTopNShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("AccountIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.AttachmentId) {
		query["AttachmentId"] = request.AttachmentId
	}

	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.CenId) {
		query["CenId"] = request.CenId
	}

	if !dara.IsNil(request.CloudIp) {
		query["CloudIp"] = request.CloudIp
	}

	if !dara.IsNil(request.CloudPort) {
		query["CloudPort"] = request.CloudPort
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.GroupBy) {
		query["GroupBy"] = request.GroupBy
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.OtherIp) {
		query["OtherIp"] = request.OtherIp
	}

	if !dara.IsNil(request.OtherPort) {
		query["OtherPort"] = request.OtherPort
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.UseMultiAccount) {
		query["UseMultiAccount"] = request.UseMultiAccount
	}

	if !dara.IsNil(request.VirtualBorderRouterId) {
		query["VirtualBorderRouterId"] = request.VirtualBorderRouterId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVbrFlowTopN"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVbrFlowTopNResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetVbrFlowTopN is deprecated, please use nis::2021-12-16::GetNisNetworkRanking instead.
//
// Summary:
//
// Queries the rankings of hybrid cloud traffic data in the form of 1-tuple, 2-tuple, or 5-tuple. Hybrid cloud traffic data can be ranked by metrics such as traffic volumes and the number of packets.
//
// @param request - GetVbrFlowTopNRequest
//
// @return GetVbrFlowTopNResponse
// Deprecated
func (client *Client) GetVbrFlowTopN(request *GetVbrFlowTopNRequest) (_result *GetVbrFlowTopNResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVbrFlowTopNResponse{}
	_body, _err := client.GetVbrFlowTopNWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 巡检资源类型列表
//
// @param request - ListNisInspectionResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNisInspectionResourceTypeResponse
func (client *Client) ListNisInspectionResourceTypeWithOptions(runtime *dara.RuntimeOptions) (_result *ListNisInspectionResourceTypeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListNisInspectionResourceType"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNisInspectionResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 巡检资源类型列表
//
// @return ListNisInspectionResourceTypeResponse
func (client *Client) ListNisInspectionResourceType() (_result *ListNisInspectionResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNisInspectionResourceTypeResponse{}
	_body, _err := client.ListNisInspectionResourceTypeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询巡检报告列表
//
// @param request - ListNisInspectionTaskReportsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNisInspectionTaskReportsResponse
func (client *Client) ListNisInspectionTaskReportsWithOptions(request *ListNisInspectionTaskReportsRequest, runtime *dara.RuntimeOptions) (_result *ListNisInspectionTaskReportsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionTaskId) {
		query["InspectionTaskId"] = request.InspectionTaskId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNisInspectionTaskReports"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNisInspectionTaskReportsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询巡检报告列表
//
// @param request - ListNisInspectionTaskReportsRequest
//
// @return ListNisInspectionTaskReportsResponse
func (client *Client) ListNisInspectionTaskReports(request *ListNisInspectionTaskReportsRequest) (_result *ListNisInspectionTaskReportsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNisInspectionTaskReportsResponse{}
	_body, _err := client.ListNisInspectionTaskReportsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 巡检任务列表
//
// @param request - ListNisInspectionTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNisInspectionTasksResponse
func (client *Client) ListNisInspectionTasksWithOptions(request *ListNisInspectionTasksRequest, runtime *dara.RuntimeOptions) (_result *ListNisInspectionTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionName) {
		query["InspectionName"] = request.InspectionName
	}

	if !dara.IsNil(request.InspectionProject) {
		query["InspectionProject"] = request.InspectionProject
	}

	if !dara.IsNil(request.InspectionTaskId) {
		query["InspectionTaskId"] = request.InspectionTaskId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNisInspectionTasks"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNisInspectionTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 巡检任务列表
//
// @param request - ListNisInspectionTasksRequest
//
// @return ListNisInspectionTasksResponse
func (client *Client) ListNisInspectionTasks(request *ListNisInspectionTasksRequest) (_result *ListNisInspectionTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNisInspectionTasksResponse{}
	_body, _err := client.ListNisInspectionTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 请补充描述开启任务
//
// @param request - StartNisInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartNisInspectionTaskResponse
func (client *Client) StartNisInspectionTaskWithOptions(request *StartNisInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *StartNisInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionTaskId) {
		query["InspectionTaskId"] = request.InspectionTaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartNisInspectionTask"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartNisInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 请补充描述开启任务
//
// @param request - StartNisInspectionTaskRequest
//
// @return StartNisInspectionTaskResponse
func (client *Client) StartNisInspectionTask(request *StartNisInspectionTaskRequest) (_result *StartNisInspectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartNisInspectionTaskResponse{}
	_body, _err := client.StartNisInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取云网络流量指标排名
//
// @param tmpReq - StartNisTrafficRankingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartNisTrafficRankingResponse
func (client *Client) StartNisTrafficRankingWithOptions(tmpReq *StartNisTrafficRankingRequest, runtime *dara.RuntimeOptions) (_result *StartNisTrafficRankingResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &StartNisTrafficRankingShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.GroupBy) {
		request.GroupByShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GroupBy, dara.String("GroupBy"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FilterShrink) {
		query["Filter"] = request.FilterShrink
	}

	if !dara.IsNil(request.GroupByShrink) {
		query["GroupBy"] = request.GroupByShrink
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.Sort) {
		query["Sort"] = request.Sort
	}

	if !dara.IsNil(request.StorageInterval) {
		query["StorageInterval"] = request.StorageInterval
	}

	if !dara.IsNil(request.TopN) {
		query["TopN"] = request.TopN
	}

	if !dara.IsNil(request.TrafficAnalyzerId) {
		query["TrafficAnalyzerId"] = request.TrafficAnalyzerId
	}

	if !dara.IsNil(request.TrafficScenario) {
		query["TrafficScenario"] = request.TrafficScenario
	}

	if !dara.IsNil(request.TupleDimension) {
		query["TupleDimension"] = request.TupleDimension
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartNisTrafficRanking"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartNisTrafficRankingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取云网络流量指标排名
//
// @param request - StartNisTrafficRankingRequest
//
// @return StartNisTrafficRankingResponse
func (client *Client) StartNisTrafficRanking(request *StartNisTrafficRankingRequest) (_result *StartNisTrafficRankingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartNisTrafficRankingResponse{}
	_body, _err := client.StartNisTrafficRankingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改巡检项
//
// @param request - UpdateNisInspectionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNisInspectionTaskResponse
func (client *Client) UpdateNisInspectionTaskWithOptions(request *UpdateNisInspectionTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateNisInspectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InspectionTaskId) {
		query["InspectionTaskId"] = request.InspectionTaskId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNisInspectionTask"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNisInspectionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改巡检项
//
// @param request - UpdateNisInspectionTaskRequest
//
// @return UpdateNisInspectionTaskResponse
func (client *Client) UpdateNisInspectionTask(request *UpdateNisInspectionTaskRequest) (_result *UpdateNisInspectionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNisInspectionTaskResponse{}
	_body, _err := client.UpdateNisInspectionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
