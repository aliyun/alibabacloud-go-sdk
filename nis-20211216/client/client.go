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
		"cn-shanghai":    dara.String("nis.aliyuncs.com"),
		"ap-southeast-1": dara.String("nis-intl.aliyuncs.com"),
	}
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
// Creates a network path for reachability analysis.
//
// Description:
//
// - You can call the **CreateNetworkPath*	- operation to create network paths in multiple networking scenarios and between multiple resources. After a path is created, the path parameters are saved for repeated analysis.
//
// - You can create up to 100 network paths within one Alibaba Cloud account.
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
// Creates a network path for reachability analysis.
//
// Description:
//
// - You can call the **CreateNetworkPath*	- operation to create network paths in multiple networking scenarios and between multiple resources. After a path is created, the path parameters are saved for repeated analysis.
//
// - You can create up to 100 network paths within one Alibaba Cloud account.
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
// - The **CreateNetworkReachableAnalysis*	- operation is used to create a task for analyzing the reachability of the network path that is created by calling the **CreateNetworkPath*	- operation and record the analysis results.
//
// - The **CreateNetworkReachableAnalysis*	- operation can be called to repeatedly analyze the reachability of a network path.
//
// - You can create up to 1,000 reachability analysis records within one Alibaba Cloud account.
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
// - The **CreateNetworkReachableAnalysis*	- operation is used to create a task for analyzing the reachability of the network path that is created by calling the **CreateNetworkPath*	- operation and record the analysis results.
//
// - The **CreateNetworkReachableAnalysis*	- operation can be called to repeatedly analyze the reachability of a network path.
//
// - You can create up to 1,000 reachability analysis records within one Alibaba Cloud account.
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
// Deletes an inspection report.
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
// Deletes an inspection report.
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
// Deletes an inspection task.
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
// Deletes an inspection task.
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
// Queries the instances of abnormal items identified in an inspection report.
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
// Queries the instances of abnormal items identified in an inspection report.
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
// Describes the details of check items in an inspection report.
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
// Describes the details of check items in an inspection report.
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
// Describes the status of an inspection report.
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
// Describes the status of an inspection report.
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
// Queries an inspection report summary.
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
// Queries an inspection report summary.
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
// Queries the details of an inspection task.
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
// Queries the details of an inspection task.
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
// Queries the results of NIS traffic ranking analysis.
//
// Description:
//
// [Before using this operation, make sure that you fully understand the billing method and pricing of the NIS Traffic Analyzer.](https://www.alibabacloud.com/help/en/nis/product-overview/billing-method-new-version)
//
// Supported analysis scenarios:
//
// - All VPC network traffic analysis
//
// - Public VPC network traffic analysis
//
// - All TR network traffic analysis
//
// - Internet Shared Bandwidth metric analysis
//
// ## VPC flow log - All VPC flow log query and analysis results
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name               | Type   | Description                                                                                       | Example            | Valid values   |
//
// |--------------------|--------|---------------------------------------------------------------------------------------------------|--------------------|----------|
//
// | RegionId           | string | The region where the flow log resides.                                                            | ap-southeast-1     |          |
//
// | VpcId              | string | The VPC ID.                                                                                       | vpc-m5ec6i0h5xss**	- |         |
//
// | VSwitchId          | string | The vSwitch ID.                                                                                   | vsw-2zeekevlh***	- |          |
//
// | NetworkInterfaceId | string | The elastic network interface (ENI) ID.                                                           | eni-8vbf2jxul**	- |          |
//
// | EcsId              | string | The ECS instance ID of the management node.                                                       | i-uf6i1zi6yhq7h**	- |          |
//
// | TrafficPath        | string | The traffic path.                                                                                 | all                |          |
//
// | Direction          | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                 | -in / -out |
//
// | SourceIp           | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort         | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432               |          |
//
// | DestinationIp      | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0      |          |
//
// | DestinationPort    | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                 |          |
//
// | Protocol           | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                |          |
//
// | Bytes              | number | The bandwidth.                                                                                    | 100                |          |
//
// | Packets            | number | The number of packets.                                                                            | 100                |          |
//
// | RoundTripTime      | number | The TCP RTT.                                                                                      | 2                  |          |
//
// | BytesRate          | number | The traffic ratio.                                                                                | 0.2                |          |
//
// ---
//
// ## VPC flow log - Public VPC flow log query and analysis results
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name               | Type   | Description                                                                                       | Example                | Valid values   |
//
// |--------------------|--------|---------------------------------------------------------------------------------------------------|------------------------|----------|
//
// | RegionId           | string | The region where the flow log resides.                                                            | ap-southeast-1         |          |
//
// | VpcId              | string | The VPC ID.                                                                                       | vpc-m5ec6i0h5xss**	- |          |
//
// | VSwitchId          | string | The vSwitch ID.                                                                                   | vsw-2zeekevlh***	- |          |
//
// | NetworkInterfaceId | string | The elastic network interface (ENI) ID.                                                           | eni-8vbf2jxul**	- |          |
//
// | EcsId              | string | The ECS instance ID of the management node.                                                       | i-uf6i1zi6yhq7h**	- |          |
//
// | TrafficPath        | string | The traffic path.                                                                                 | all                    |          |
//
// | Direction          | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                     | -in / -out |
//
// | SourceIp           | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort         | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                   |          |
//
// | DestinationIp      | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0          |          |
//
// | DestinationPort    | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                     |          |
//
// | Protocol           | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                    |          |
//
// | ClientCountry      | string | The country of the client.                                                                        | China                  |          |
//
// | ClientProvince     | string | The province of the client.                                                                       | Hong Kong SAR          |          |
//
// | ClientCity         | string | The city of the client.                                                                           | Kowloon                |          |
//
// | ClientIsp          | string | The network service provider.                                                                     | Alibaba Cloud          |          |
//
// | ClientAsn          | string | The autonomous system number.                                                                     | 45102                  |          |
//
// | Bytes              | number | The bandwidth.                                                                                    | 100                    |          |
//
// | Packets            | number | The number of packets.                                                                            | 100                    |          |
//
// | RoundTripTime      | number | The TCP RTT.                                                                                      | 2                      |          |
//
// | BytesRate          | number | The traffic ratio.                                                                                | 0.2                    |          |
//
// ---
//
// ## TR flow log - TR cross-region scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                      | Type   | Description                                                                                       | Example                 | Valid values   |
//
// |---------------------------|--------|---------------------------------------------------------------------------------------------------|-------------------------|----------|
//
// | Direction                 | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                      | -in / -out |
//
// | SourceIp                  | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort                | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                    |          |
//
// | DestinationIp             | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0           |          |
//
// | DestinationPort           | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                      |          |
//
// | Protocol                  | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                     |          |
//
// | Dscp                      | string | The Differentiated Services Code Point (DSCP) value.                                              | 0                       |          |
//
// | DestinationRegionNo       | string | The destination region ID.                                                                        | ap-southeast-1          |          |
//
// | SourceRegionNo            | string | The source region ID.                                                                             | ap-southeast-1          |          |
//
// | TransitRouterAttachmentId | string | The network instance connection ID.                                                               | tr-attach-bfde1cd4cj**	- |          |
//
// | TransitRouterId           | string | The transit router instance ID.                                                                   | tr-2zefvwy2fz3444**	- |          |
//
// | TransitRouterPairAttachmentId | string | The transit router peering connection instance ID.                                            | tr-attach-okvj1cd4cjp**	- |         |
//
// | Bytes                     | number | The bandwidth.                                                                                    | 100                     |          |
//
// | Packets                   | number | The number of packets.                                                                            | 100                     |          |
//
// | BytesRate                 | number | The traffic ratio.                                                                                | 0.2                     |          |
//
// | PacketsLostNoRoute        | number | The number of packets dropped due to no route.                                                    | 2                       |          |
//
// | PacketsLostBlackhole      | number | The number of packets dropped due to blackhole routes.                                            | 4                       |          |
//
// | PacketsLostTTLExpired     | number | The number of packets dropped due to TTL expiration.                                              | 7                       |          |
//
// ---
//
// ## TR flow log - VPC connection traffic scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                                   | Type   | Description                                                                                       | Example                         | Valid values   |
//
// |----------------------------------------|--------|---------------------------------------------------------------------------------------------------|---------------------------------|----------|
//
// | Direction                              | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                              | -in / -out |
//
// | SourceIp                               | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort                             | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                            |          |
//
// | DestinationIp                          | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0                   |          |
//
// | DestinationPort                        | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                              |          |
//
// | Protocol                               | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                             |          |
//
// | Dscp                                   | string | The Differentiated Services Code Point (DSCP) value.                                              | 0                               |          |
//
// | DestinationRegionNo                    | string | The destination region ID.                                                                        | ap-southeast-1                  |          |
//
// | SourceRegionNo                         | string | The source region ID.                                                                             | ap-southeast-1                  |          |
//
// | TransitRouterAttachmentId              | string | The network instance connection ID.                                                               | tr-attach-bfde1cd4cj**	- |          |
//
// | TransitRouterId                        | string | The transit router instance ID.                                                                   | tr-2zefvwy2fz3444**	- |          |
//
// | TransitRouterPairAttachmentId          | string | The transit router peering connection instance ID.                                                 | tr-attach-okvj1cd4cjp**	- |          |
//
// | TransitRouterSourceResourceId          | string | The source network instance ID.                                                                   | tr-attach-hvve1cd4cjpj**	- |          |
//
// | TransitRouterSourceAccountId           | string | The account ID of the source network instance.                                                    | 1906814138**	- |          |
//
// | TransitRouterSourceVSwitchId           | string | The vSwitch ID of the source TR ENI.                                                              | vsw-ikfdkevlhxpqxuz***	- |          |
//
// | TransitRouterSourceNetworkInterface    | string | The source TR ENI.                                                                                | eni-8vbf2jxulma**	- |          |
//
// | TransitRouterSourceAvailableZone       | string | The source zone ID.                                                                               | ap-southeast-1-j                |          |
//
// | TransitRouterDestinationResourceId     | string | The destination network instance ID.                                                              | tr-attach-bfve1cd4cjp***	- |          |
//
// | TransitRouterDestinationAccountId      | string | The account ID of the destination network instance.                                               | 1906814138**	- |          |
//
// | TransitRouterDestinationVSwitchId      | string | The vSwitch ID of the destination TR ENI.                                                         | vsw-ikfdkevlhxpqxuz***	- |          |
//
// | TransitRouterDestinationNetworkInterface | string | The destination TR ENI.                                                                         | eni-7kcf2jxulma**	- |          |
//
// | TransitRouterDestinationAvailableZone  | string | The destination zone ID.                                                                          | ap-southeast-1-j                |          |
//
// | Bytes                                  | number | The bandwidth.                                                                                    | 100                             |          |
//
// | Packets                                | number | The number of packets.                                                                            | 100                             |          |
//
// | BytesRate                              | number | The traffic ratio.                                                                                | 0.2                             |          |
//
// | PacketsLostNoRoute                     | number | The number of packets dropped due to no route.                                                    | 2                               |          |
//
// | PacketsLostBlackhole                   | number | The number of packets dropped due to blackhole routes.                                            | 4                               |          |
//
// | PacketsLostTTLExpired                  | number | The number of packets dropped due to TTL expiration.                                              | 7                               |          |
//
// ---
//
// ## TR flow log - VBR traffic scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                              | Type   | Description                                                                                       | Example                         | Valid values   |
//
// |-----------------------------------|--------|---------------------------------------------------------------------------------------------------|---------------------------------|----------|
//
// | Direction                         | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                              | -in / -out |
//
// | SourceIp                          | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort                        | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                            |          |
//
// | DestinationIp                     | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0                   |          |
//
// | DestinationPort                   | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                              |          |
//
// | Protocol                          | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                             |          |
//
// | Dscp                              | string | The Differentiated Services Code Point (DSCP) value.                                              | 0                               |          |
//
// | DestinationRegionNo               | string | The destination region ID.                                                                        | ap-southeast-1                  |          |
//
// | SourceRegionNo                    | string | The source region ID.                                                                             | ap-southeast-1                  |          |
//
// | TransitRouterAttachmentId         | string | The network instance connection ID.                                                               | tr-attach-bfde1cd4cj**	- |          |
//
// | TransitRouterId                   | string | The transit router instance ID.                                                                   | tr-2zefvwy2fz3444**	- |          |
//
// | TransitRouterPairAttachmentId     | string | The transit router peering connection instance ID.                                                 | tr-attach-okvj1cd4cjp**	- |          |
//
// | TransitRouterSourceResourceId     | string | The source network instance ID.                                                                   | tr-attach-hvve1cd4cjpj**	- |          |
//
// | TransitRouterSourceAccountId      | string | The account ID of the source network instance.                                                    | 1906814138**	- |          |
//
// | TransitRouterDestinationResourceId| string | The destination network instance ID.                                                              | tr-attach-bfve1cd4cjp***	- |          |
//
// | TransitRouterDestinationAccountId | string | The account ID of the destination network instance.                                               | 1906814138**	- |          |
//
// | Bytes                             | number | The bandwidth.                                                                                    | 100                             |          |
//
// | Packets                           | number | The number of packets.                                                                            | 100                             |          |
//
// | BytesRate                         | number | The traffic ratio.                                                                                | 0.2                             |          |
//
// | PacketsLostNoRoute                | number | The number of packets dropped due to no route.                                                    | 2                               |          |
//
// | PacketsLostBlackhole              | number | The number of packets dropped due to blackhole routes.                                            | 4                               |          |
//
// | PacketsLostTTLExpired             | number | The number of packets dropped due to TTL expiration.                                              | 7                               |          |
//
// ---
//
// ## TR flow log - ECR traffic scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                              | Type   | Description                                                                                       | Example                         | Valid values   |
//
// |-----------------------------------|--------|---------------------------------------------------------------------------------------------------|---------------------------------|----------|
//
// | Direction                         | string | The traffic direction based on the
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
// Queries the results of NIS traffic ranking analysis.
//
// Description:
//
// [Before using this operation, make sure that you fully understand the billing method and pricing of the NIS Traffic Analyzer.](https://www.alibabacloud.com/help/en/nis/product-overview/billing-method-new-version)
//
// Supported analysis scenarios:
//
// - All VPC network traffic analysis
//
// - Public VPC network traffic analysis
//
// - All TR network traffic analysis
//
// - Internet Shared Bandwidth metric analysis
//
// ## VPC flow log - All VPC flow log query and analysis results
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name               | Type   | Description                                                                                       | Example            | Valid values   |
//
// |--------------------|--------|---------------------------------------------------------------------------------------------------|--------------------|----------|
//
// | RegionId           | string | The region where the flow log resides.                                                            | ap-southeast-1     |          |
//
// | VpcId              | string | The VPC ID.                                                                                       | vpc-m5ec6i0h5xss**	- |         |
//
// | VSwitchId          | string | The vSwitch ID.                                                                                   | vsw-2zeekevlh***	- |          |
//
// | NetworkInterfaceId | string | The elastic network interface (ENI) ID.                                                           | eni-8vbf2jxul**	- |          |
//
// | EcsId              | string | The ECS instance ID of the management node.                                                       | i-uf6i1zi6yhq7h**	- |          |
//
// | TrafficPath        | string | The traffic path.                                                                                 | all                |          |
//
// | Direction          | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                 | -in / -out |
//
// | SourceIp           | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort         | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432               |          |
//
// | DestinationIp      | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0      |          |
//
// | DestinationPort    | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                 |          |
//
// | Protocol           | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                |          |
//
// | Bytes              | number | The bandwidth.                                                                                    | 100                |          |
//
// | Packets            | number | The number of packets.                                                                            | 100                |          |
//
// | RoundTripTime      | number | The TCP RTT.                                                                                      | 2                  |          |
//
// | BytesRate          | number | The traffic ratio.                                                                                | 0.2                |          |
//
// ---
//
// ## VPC flow log - Public VPC flow log query and analysis results
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name               | Type   | Description                                                                                       | Example                | Valid values   |
//
// |--------------------|--------|---------------------------------------------------------------------------------------------------|------------------------|----------|
//
// | RegionId           | string | The region where the flow log resides.                                                            | ap-southeast-1         |          |
//
// | VpcId              | string | The VPC ID.                                                                                       | vpc-m5ec6i0h5xss**	- |          |
//
// | VSwitchId          | string | The vSwitch ID.                                                                                   | vsw-2zeekevlh***	- |          |
//
// | NetworkInterfaceId | string | The elastic network interface (ENI) ID.                                                           | eni-8vbf2jxul**	- |          |
//
// | EcsId              | string | The ECS instance ID of the management node.                                                       | i-uf6i1zi6yhq7h**	- |          |
//
// | TrafficPath        | string | The traffic path.                                                                                 | all                    |          |
//
// | Direction          | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                     | -in / -out |
//
// | SourceIp           | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort         | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                   |          |
//
// | DestinationIp      | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0          |          |
//
// | DestinationPort    | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                     |          |
//
// | Protocol           | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                    |          |
//
// | ClientCountry      | string | The country of the client.                                                                        | China                  |          |
//
// | ClientProvince     | string | The province of the client.                                                                       | Hong Kong SAR          |          |
//
// | ClientCity         | string | The city of the client.                                                                           | Kowloon                |          |
//
// | ClientIsp          | string | The network service provider.                                                                     | Alibaba Cloud          |          |
//
// | ClientAsn          | string | The autonomous system number.                                                                     | 45102                  |          |
//
// | Bytes              | number | The bandwidth.                                                                                    | 100                    |          |
//
// | Packets            | number | The number of packets.                                                                            | 100                    |          |
//
// | RoundTripTime      | number | The TCP RTT.                                                                                      | 2                      |          |
//
// | BytesRate          | number | The traffic ratio.                                                                                | 0.2                    |          |
//
// ---
//
// ## TR flow log - TR cross-region scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                      | Type   | Description                                                                                       | Example                 | Valid values   |
//
// |---------------------------|--------|---------------------------------------------------------------------------------------------------|-------------------------|----------|
//
// | Direction                 | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                      | -in / -out |
//
// | SourceIp                  | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort                | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                    |          |
//
// | DestinationIp             | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0           |          |
//
// | DestinationPort           | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                      |          |
//
// | Protocol                  | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                     |          |
//
// | Dscp                      | string | The Differentiated Services Code Point (DSCP) value.                                              | 0                       |          |
//
// | DestinationRegionNo       | string | The destination region ID.                                                                        | ap-southeast-1          |          |
//
// | SourceRegionNo            | string | The source region ID.                                                                             | ap-southeast-1          |          |
//
// | TransitRouterAttachmentId | string | The network instance connection ID.                                                               | tr-attach-bfde1cd4cj**	- |          |
//
// | TransitRouterId           | string | The transit router instance ID.                                                                   | tr-2zefvwy2fz3444**	- |          |
//
// | TransitRouterPairAttachmentId | string | The transit router peering connection instance ID.                                            | tr-attach-okvj1cd4cjp**	- |         |
//
// | Bytes                     | number | The bandwidth.                                                                                    | 100                     |          |
//
// | Packets                   | number | The number of packets.                                                                            | 100                     |          |
//
// | BytesRate                 | number | The traffic ratio.                                                                                | 0.2                     |          |
//
// | PacketsLostNoRoute        | number | The number of packets dropped due to no route.                                                    | 2                       |          |
//
// | PacketsLostBlackhole      | number | The number of packets dropped due to blackhole routes.                                            | 4                       |          |
//
// | PacketsLostTTLExpired     | number | The number of packets dropped due to TTL expiration.                                              | 7                       |          |
//
// ---
//
// ## TR flow log - VPC connection traffic scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                                   | Type   | Description                                                                                       | Example                         | Valid values   |
//
// |----------------------------------------|--------|---------------------------------------------------------------------------------------------------|---------------------------------|----------|
//
// | Direction                              | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                              | -in / -out |
//
// | SourceIp                               | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort                             | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                            |          |
//
// | DestinationIp                          | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0                   |          |
//
// | DestinationPort                        | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                              |          |
//
// | Protocol                               | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                             |          |
//
// | Dscp                                   | string | The Differentiated Services Code Point (DSCP) value.                                              | 0                               |          |
//
// | DestinationRegionNo                    | string | The destination region ID.                                                                        | ap-southeast-1                  |          |
//
// | SourceRegionNo                         | string | The source region ID.                                                                             | ap-southeast-1                  |          |
//
// | TransitRouterAttachmentId              | string | The network instance connection ID.                                                               | tr-attach-bfde1cd4cj**	- |          |
//
// | TransitRouterId                        | string | The transit router instance ID.                                                                   | tr-2zefvwy2fz3444**	- |          |
//
// | TransitRouterPairAttachmentId          | string | The transit router peering connection instance ID.                                                 | tr-attach-okvj1cd4cjp**	- |          |
//
// | TransitRouterSourceResourceId          | string | The source network instance ID.                                                                   | tr-attach-hvve1cd4cjpj**	- |          |
//
// | TransitRouterSourceAccountId           | string | The account ID of the source network instance.                                                    | 1906814138**	- |          |
//
// | TransitRouterSourceVSwitchId           | string | The vSwitch ID of the source TR ENI.                                                              | vsw-ikfdkevlhxpqxuz***	- |          |
//
// | TransitRouterSourceNetworkInterface    | string | The source TR ENI.                                                                                | eni-8vbf2jxulma**	- |          |
//
// | TransitRouterSourceAvailableZone       | string | The source zone ID.                                                                               | ap-southeast-1-j                |          |
//
// | TransitRouterDestinationResourceId     | string | The destination network instance ID.                                                              | tr-attach-bfve1cd4cjp***	- |          |
//
// | TransitRouterDestinationAccountId      | string | The account ID of the destination network instance.                                               | 1906814138**	- |          |
//
// | TransitRouterDestinationVSwitchId      | string | The vSwitch ID of the destination TR ENI.                                                         | vsw-ikfdkevlhxpqxuz***	- |          |
//
// | TransitRouterDestinationNetworkInterface | string | The destination TR ENI.                                                                         | eni-7kcf2jxulma**	- |          |
//
// | TransitRouterDestinationAvailableZone  | string | The destination zone ID.                                                                          | ap-southeast-1-j                |          |
//
// | Bytes                                  | number | The bandwidth.                                                                                    | 100                             |          |
//
// | Packets                                | number | The number of packets.                                                                            | 100                             |          |
//
// | BytesRate                              | number | The traffic ratio.                                                                                | 0.2                             |          |
//
// | PacketsLostNoRoute                     | number | The number of packets dropped due to no route.                                                    | 2                               |          |
//
// | PacketsLostBlackhole                   | number | The number of packets dropped due to blackhole routes.                                            | 4                               |          |
//
// | PacketsLostTTLExpired                  | number | The number of packets dropped due to TTL expiration.                                              | 7                               |          |
//
// ---
//
// ## TR flow log - VBR traffic scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                              | Type   | Description                                                                                       | Example                         | Valid values   |
//
// |-----------------------------------|--------|---------------------------------------------------------------------------------------------------|---------------------------------|----------|
//
// | Direction                         | string | The traffic direction based on the Alibaba Cloud network resource instance. Valid values:
//
// ● in: inbound traffic.
//
// ● out: outbound traffic. | in                              | -in / -out |
//
// | SourceIp                          | string | The source IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.    | 47.92.245.**	- |          |
//
// | SourcePort                        | string | The source port.
//
// ● This field is returned only when 5-tuple statistics are collected.          | 5432                            |          |
//
// | DestinationIp                     | string | The destination IP address.
//
// ● This field is returned only when 2-tuple statistics are collected.| 192.168.***.0                   |          |
//
// | DestinationPort                   | string | The destination port.
//
// ● This field is returned only when 5-tuple statistics are collected.     | 23                              |          |
//
// | Protocol                          | string | The network protocol.
//
// ● This field is returned only when 5-tuple statistics are collected.     | TCP                             |          |
//
// | Dscp                              | string | The Differentiated Services Code Point (DSCP) value.                                              | 0                               |          |
//
// | DestinationRegionNo               | string | The destination region ID.                                                                        | ap-southeast-1                  |          |
//
// | SourceRegionNo                    | string | The source region ID.                                                                             | ap-southeast-1                  |          |
//
// | TransitRouterAttachmentId         | string | The network instance connection ID.                                                               | tr-attach-bfde1cd4cj**	- |          |
//
// | TransitRouterId                   | string | The transit router instance ID.                                                                   | tr-2zefvwy2fz3444**	- |          |
//
// | TransitRouterPairAttachmentId     | string | The transit router peering connection instance ID.                                                 | tr-attach-okvj1cd4cjp**	- |          |
//
// | TransitRouterSourceResourceId     | string | The source network instance ID.                                                                   | tr-attach-hvve1cd4cjpj**	- |          |
//
// | TransitRouterSourceAccountId      | string | The account ID of the source network instance.                                                    | 1906814138**	- |          |
//
// | TransitRouterDestinationResourceId| string | The destination network instance ID.                                                              | tr-attach-bfve1cd4cjp***	- |          |
//
// | TransitRouterDestinationAccountId | string | The account ID of the destination network instance.                                               | 1906814138**	- |          |
//
// | Bytes                             | number | The bandwidth.                                                                                    | 100                             |          |
//
// | Packets                           | number | The number of packets.                                                                            | 100                             |          |
//
// | BytesRate                         | number | The traffic ratio.                                                                                | 0.2                             |          |
//
// | PacketsLostNoRoute                | number | The number of packets dropped due to no route.                                                    | 2                               |          |
//
// | PacketsLostBlackhole              | number | The number of packets dropped due to blackhole routes.                                            | 4                               |          |
//
// | PacketsLostTTLExpired             | number | The number of packets dropped due to TTL expiration.                                              | 7                               |          |
//
// ---
//
// ## TR flow log - ECR traffic scenario analysis results.
//
// ### Request parameters
//
// | Name                | Type    | Required | Description                                                                 | Example                                      | Valid values |
//
// |---------------------|---------|----------|-----------------------------------------------------------------------------|---------------------------------------------|--------|
//
// | NisTrafficRankingId | string  | Yes      | The ID of the network traffic analysis result.                              | task-6462a7b4c4a54b***	- |        |
//
// | NextToken           | string  | No       | The paging token. Set this parameter to the NextToken value returned in the previous API call. | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | No       | The paging size. Valid values: 1 to 100. Default value: 20.                 | 20                                          |        |
//
// ### Response parameters
//
// | Name                | Type    | Description                                    | Example                                      | Valid values |
//
// |---------------------|---------|------------------------------------------------|---------------------------------------------|--------|
//
// | -                   | object  | RpcResponse                                    |                                             |        |
//
// | RequestId           | string  | The request ID.                                | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941        |        |
//
// | NisTrafficRankingId | string  | The ID of the network traffic analysis result. | task-7619ecb1db9148bab9f4                   |        |
//
// | Status              | string  | The task running status.                       | Complete                                    |        |
//
// | NextToken           | string  | The token for the next query.                  | LoeJLhK0fsDqYoXkXieZUqB2vWnccJtVnsyKu9KxFFOMQxtV8XckOg5lk7F2bhC+ |        |
//
// | MaxResults          | integer | The paging size. Valid values: 1 to 100. Default value: 20. | 20                                          |        |
//
// | TotalCount          | integer | The total number of entries returned.          | 72                                          |        |
//
// | FlowRankingList     | array   | The list of network traffic analysis results.  |                                             |        |
//
// #### FlowRankingList element structure
//
// | Name                              | Type   | Description                                                                                       | Example                         | Valid values   |
//
// |-----------------------------------|--------|---------------------------------------------------------------------------------------------------|---------------------------------|----------|
//
// | Direction                         | string | The traffic direction based on the
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
// - The **init*	- state indicates that the task is in progress.
//
// - The **finish*	- state indicates that the task is complete. In this state, you can obtain the analysis result.
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
// - The **init*	- state indicates that the task is in progress.
//
// - The **finish*	- state indicates that the task is complete. In this state, you can obtain the analysis result.
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
// Retrieves detailed monitoring metric trends data for a specified NIS network analysis scenario, including traffic metric trends for network traffic analysis scenarios and performance metric trends for performance observation scenarios.
//
// Description:
//
// Data query result synchronization:
//
// - API data is synchronized every 6 hours. You can query current network analysis data only after 6 hours.
//
// Supported analysis scenarios:
//
// - Internet performance observation dashboard scenario
//
// - Cross-zone internal network performance observation dashboard scenario
//
// - Cross-region internal network performance observation dashboard scenario
//
// ## Internet performance observation dashboard scenario
//
// **Note**
//
// The maximum query time span is 24 hours. If not specified, the most recent 1 hour is queried by default. The query result contains one data point per minute.
//
// ### **Request parameters**
//
// | **Name*	- | **Type*	- | **Required*	- | **Description*	- |
//
// | --- | --- | --- | --- |
//
// | RegionNo | string | Yes | The Alibaba Cloud region where the probing source is located. |
//
// | ResourceType | string | Yes | Set to **InternetProbing*	- to query Internet performance monitoring trends. |
//
// | MetricName | string | Yes | The metric for which to query trends. Valid value: **rtt**, the round-trip time of probing. |
//
// | Dimensions | object[\\] | Yes | The filter information. |
//
// | \\-Name | string | Yes | The filter condition. Valid values:
//
//   - **Country**: the probing country.
//
//   - **Province**: the probing province.
//
//   - **Isp**: the probing ISP.
//
//     **Note:*	- Specify at least one of the preceding parameters. |
//
// | \\-Value | string | Yes | The filter value corresponding to the filter condition. Examples:
//
//   - Country: China
//
//   - Province: Zhejiang
//
// - Isp: Alibaba
//
// **Note:*	- Country and province values are capitalized. The ISP parameter value must match the name displayed in the console. |
//
// ### **Response parameters**
//
// | **Name*	- | **Type*	- | **Description*	- |
//
// | --- | --- | --- |
//
// | RequestId | string | The request ID. |
//
// | Data | object | The cloud network metric trends data object. |
//
// | Metrics | array | The collection of metric trends data. |
//
// | \\-TimeStamp | long | The UNIX timestamp in milliseconds. |
//
// | \\-Value | double | The metric value corresponding to the **MetricName*	- input parameter. |
//
// | Unit | String | The unit of **Value**. |
//
// ## Cross-zone internal network performance observation dashboard scenario
//
// **Note**
//
// -   If the query time span exceeds 5 days, the query result contains one data point per day.
//
// -   If the query time span exceeds 1 day, the query result contains one data point per hour.
//
// -   If the query time span is less than 1 day, the query result contains one data point per 5 minutes.
//
// ### **Request parameters**
//
// | **Name*	- | **Type*	- | **Required*	- | **Description*	- |
//
// | --- | --- | --- | --- |
//
// | RegionNo | string | Yes | The Alibaba Cloud region. |
//
// | ResourceType | string | Yes | Set to **IntranetProbing*	- to query cross-zone performance monitoring trends. |
//
// | MetricName | string | Yes | The metric for which to query trends. Valid value: **rtt**, the round-trip time of probing. |
//
// | Dimensions | object[\\] | Yes | The filter information. |
//
// | \\-Name | string | Yes | The filter condition. Valid values:
//
//   - **SourceZone**: the source zone for probing. This parameter is required.
//
//   - **DestinationZone**: the destination zone for probing. This parameter is required. |
//
// | \\-Value | string | Yes | The filter value corresponding to the filter condition. Examples:
//
//   - SourceZone: cn-hangzhou-j
//
//   - DestinationZone: cn-hangzhou-k |
//
// ### **Response parameters**
//
// | **Name*	- | **Type*	- | **Description*	- |
//
// | --- | --- | --- |
//
// | RequestId | string | The request ID. |
//
// | Data | object | The cloud network metric trends data object. |
//
// | Metrics | array | The collection of metric trends data. |
//
// | \\-TimeStamp | long | The UNIX timestamp in milliseconds. |
//
// | \\-Value | double | The metric value corresponding to the **MetricName*	- input parameter. |
//
// | Unit | String | The unit of **Value**. |
//
// ## **Cross-region internal network performance observation dashboard scenario**
//
// **Note**
//
// -   If the query time span exceeds 5 days, the query result contains one data point per day.
//
// -   If the query time span exceeds 1 day, the query result contains one data point per hour.
//
// -   If the query time span is less than 1 day, the query result contains one data point per 5 minutes.
//
// ### **Request parameters**
//
// | **Name*	- | **Type*	- | **Required*	- | **Description*	- |
//
// | --- | --- | --- | --- |
//
// | RegionNo | string | Yes | The Alibaba Cloud region. |
//
// | ResourceType | string | Yes | Set to **IntranetProbing*	- to query cross-region performance observation rankings. |
//
// | Direction | string | Yes | The probing direction. Valid values:
//
//   - **in**: probing with RegionNo as the destination.
//
//   - **out**: probing with RegionNo as the source. |
//
// | MetricName | string | Yes | The metric for which to query trends. Valid value: **rtt**, the round-trip time of probing. |
//
// | Dimensions | object[\\] | No | The filter information. |
//
// | \\-Name | string | No | The filter condition. Valid value: **DestinationRegionNo**, the destination region for probing. This parameter is required. |
//
// | \\-Value | string | No | The destination region ID. Example: DestinationRegionNo: cn-shenzhen |
//
// ### **Response parameters**
//
// | **Name*	- | **Type*	- | **Description*	- |
//
// | --- | --- | --- |
//
// | RequestId | string | The request ID. |
//
// | Data | object | The cloud network metric trends data object. |
//
// | Metrics | array | The collection of metric trends data. |
//
// | \\-TimeStamp | long | The UNIX timestamp in milliseconds. |
//
// | \\-Value | double | The metric value corresponding to the **MetricName*	- input parameter. |
//
// | Unit | String | The unit of **Value**. |
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
// Retrieves detailed monitoring metric trends data for a specified NIS network analysis scenario, including traffic metric trends for network traffic analysis scenarios and performance metric trends for performance observation scenarios.
//
// Description:
//
// Data query result synchronization:
//
// - API data is synchronized every 6 hours. You can query current network analysis data only after 6 hours.
//
// Supported analysis scenarios:
//
// - Internet performance observation dashboard scenario
//
// - Cross-zone internal network performance observation dashboard scenario
//
// - Cross-region internal network performance observation dashboard scenario
//
// ## Internet performance observation dashboard scenario
//
// **Note**
//
// The maximum query time span is 24 hours. If not specified, the most recent 1 hour is queried by default. The query result contains one data point per minute.
//
// ### **Request parameters**
//
// | **Name*	- | **Type*	- | **Required*	- | **Description*	- |
//
// | --- | --- | --- | --- |
//
// | RegionNo | string | Yes | The Alibaba Cloud region where the probing source is located. |
//
// | ResourceType | string | Yes | Set to **InternetProbing*	- to query Internet performance monitoring trends. |
//
// | MetricName | string | Yes | The metric for which to query trends. Valid value: **rtt**, the round-trip time of probing. |
//
// | Dimensions | object[\\] | Yes | The filter information. |
//
// | \\-Name | string | Yes | The filter condition. Valid values:
//
//   - **Country**: the probing country.
//
//   - **Province**: the probing province.
//
//   - **Isp**: the probing ISP.
//
//     **Note:*	- Specify at least one of the preceding parameters. |
//
// | \\-Value | string | Yes | The filter value corresponding to the filter condition. Examples:
//
//   - Country: China
//
//   - Province: Zhejiang
//
// - Isp: Alibaba
//
// **Note:*	- Country and province values are capitalized. The ISP parameter value must match the name displayed in the console. |
//
// ### **Response parameters**
//
// | **Name*	- | **Type*	- | **Description*	- |
//
// | --- | --- | --- |
//
// | RequestId | string | The request ID. |
//
// | Data | object | The cloud network metric trends data object. |
//
// | Metrics | array | The collection of metric trends data. |
//
// | \\-TimeStamp | long | The UNIX timestamp in milliseconds. |
//
// | \\-Value | double | The metric value corresponding to the **MetricName*	- input parameter. |
//
// | Unit | String | The unit of **Value**. |
//
// ## Cross-zone internal network performance observation dashboard scenario
//
// **Note**
//
// -   If the query time span exceeds 5 days, the query result contains one data point per day.
//
// -   If the query time span exceeds 1 day, the query result contains one data point per hour.
//
// -   If the query time span is less than 1 day, the query result contains one data point per 5 minutes.
//
// ### **Request parameters**
//
// | **Name*	- | **Type*	- | **Required*	- | **Description*	- |
//
// | --- | --- | --- | --- |
//
// | RegionNo | string | Yes | The Alibaba Cloud region. |
//
// | ResourceType | string | Yes | Set to **IntranetProbing*	- to query cross-zone performance monitoring trends. |
//
// | MetricName | string | Yes | The metric for which to query trends. Valid value: **rtt**, the round-trip time of probing. |
//
// | Dimensions | object[\\] | Yes | The filter information. |
//
// | \\-Name | string | Yes | The filter condition. Valid values:
//
//   - **SourceZone**: the source zone for probing. This parameter is required.
//
//   - **DestinationZone**: the destination zone for probing. This parameter is required. |
//
// | \\-Value | string | Yes | The filter value corresponding to the filter condition. Examples:
//
//   - SourceZone: cn-hangzhou-j
//
//   - DestinationZone: cn-hangzhou-k |
//
// ### **Response parameters**
//
// | **Name*	- | **Type*	- | **Description*	- |
//
// | --- | --- | --- |
//
// | RequestId | string | The request ID. |
//
// | Data | object | The cloud network metric trends data object. |
//
// | Metrics | array | The collection of metric trends data. |
//
// | \\-TimeStamp | long | The UNIX timestamp in milliseconds. |
//
// | \\-Value | double | The metric value corresponding to the **MetricName*	- input parameter. |
//
// | Unit | String | The unit of **Value**. |
//
// ## **Cross-region internal network performance observation dashboard scenario**
//
// **Note**
//
// -   If the query time span exceeds 5 days, the query result contains one data point per day.
//
// -   If the query time span exceeds 1 day, the query result contains one data point per hour.
//
// -   If the query time span is less than 1 day, the query result contains one data point per 5 minutes.
//
// ### **Request parameters**
//
// | **Name*	- | **Type*	- | **Required*	- | **Description*	- |
//
// | --- | --- | --- | --- |
//
// | RegionNo | string | Yes | The Alibaba Cloud region. |
//
// | ResourceType | string | Yes | Set to **IntranetProbing*	- to query cross-region performance observation rankings. |
//
// | Direction | string | Yes | The probing direction. Valid values:
//
//   - **in**: probing with RegionNo as the destination.
//
//   - **out**: probing with RegionNo as the source. |
//
// | MetricName | string | Yes | The metric for which to query trends. Valid value: **rtt**, the round-trip time of probing. |
//
// | Dimensions | object[\\] | No | The filter information. |
//
// | \\-Name | string | No | The filter condition. Valid value: **DestinationRegionNo**, the destination region for probing. This parameter is required. |
//
// | \\-Value | string | No | The destination region ID. Example: DestinationRegionNo: cn-shenzhen |
//
// ### **Response parameters**
//
// | **Name*	- | **Type*	- | **Description*	- |
//
// | --- | --- | --- |
//
// | RequestId | string | The request ID. |
//
// | Data | object | The cloud network metric trends data object. |
//
// | Metrics | array | The collection of metric trends data. |
//
// | \\-TimeStamp | long | The UNIX timestamp in milliseconds. |
//
// | \\-Value | double | The metric value corresponding to the **MetricName*	- input parameter. |
//
// | Unit | String | The unit of **Value**. |
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
// Supports ranking analysis of traffic flowing through different cloud network resources by using sorting algorithms and data processing capabilities, and sorts detection metrics of Alibaba Cloud performance observation to help you quickly identify network bottlenecks, optimize resource configurations, and improve overall network performance.
//
// Description:
//
// Supported analysis scenarios:
//
// ## Internet performance observation dashboard scenario.
//
// ### **Request parameters**
//
// ### **Response parameters**
//
// ## Cross-zone internal network performance observation dashboard scenario.
//
// ### **Request parameters**
//
// ### **Response parameters**
//
// ## **Cross-region internal network performance observation dashboard scenario**
//
// ### **Request parameters**
//
// ### **Response parameters**
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
// Supports ranking analysis of traffic flowing through different cloud network resources by using sorting algorithms and data processing capabilities, and sorts detection metrics of Alibaba Cloud performance observation to help you quickly identify network bottlenecks, optimize resource configurations, and improve overall network performance.
//
// Description:
//
// Supported analysis scenarios:
//
// ## Internet performance observation dashboard scenario.
//
// ### **Request parameters**
//
// ### **Response parameters**
//
// ## Cross-zone internal network performance observation dashboard scenario.
//
// ### **Request parameters**
//
// ### **Response parameters**
//
// ## **Cross-region internal network performance observation dashboard scenario**
//
// ### **Request parameters**
//
// ### **Response parameters**
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

// Summary:
//
// Retrieves the metric trends for network traffic analysis.
//
// @param tmpReq - GetNisTrafficMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNisTrafficMetricsResponse
func (client *Client) GetNisTrafficMetricsWithOptions(tmpReq *GetNisTrafficMetricsRequest, runtime *dara.RuntimeOptions) (_result *GetNisTrafficMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetNisTrafficMetricsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Filter) {
		request.FilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filter, dara.String("Filter"), dara.String("json"))
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

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MetricName) {
		query["MetricName"] = request.MetricName
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionNo) {
		query["RegionNo"] = request.RegionNo
	}

	if !dara.IsNil(request.ScanBy) {
		query["ScanBy"] = request.ScanBy
	}

	if !dara.IsNil(request.StepMinutes) {
		query["StepMinutes"] = request.StepMinutes
	}

	if !dara.IsNil(request.StorageInterval) {
		query["StorageInterval"] = request.StorageInterval
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
		Action:      dara.String("GetNisTrafficMetrics"),
		Version:     dara.String("2021-12-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNisTrafficMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the metric trends for network traffic analysis.
//
// @param request - GetNisTrafficMetricsRequest
//
// @return GetNisTrafficMetricsResponse
func (client *Client) GetNisTrafficMetrics(request *GetNisTrafficMetricsRequest) (_result *GetNisTrafficMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNisTrafficMetricsResponse{}
	_body, _err := client.GetNisTrafficMetricsWithOptions(request, runtime)
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
// Lists the resource types available for inspection.
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
// Lists the resource types available for inspection.
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
// Queries a list of inspection task reports.
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
// Queries a list of inspection task reports.
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
// Lists the network inspection tasks.
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
// Lists the network inspection tasks.
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
// Starts an inspection task to generate an inspection report.
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
// Starts an inspection task to generate an inspection report.
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
// Triggers the NIS traffic analyzer to execute a traffic sorting task within a specified scope.
//
// Description:
//
// [Before using this operation, make sure that you fully understand the billing method and pricing of the NIS traffic analyzer.](https://www.alibabacloud.com/help/en/nis/product-overview/billing-method-new-version)
//
// Before using this operation, create a traffic analyzer and add a data source.
//
// [Create a traffic analyzer](https://www.alibabacloud.com/help/en/nis/user-guide/traffic-analyzer-management#39d1693bce6yp)
//
// [Add a data source](https://www.alibabacloud.com/help/en/nis/user-guide/data-source-management#73845748bfstv)
//
// Supported analysis scenarios:
//
// - All VPC network traffic analysis
//
// - Internet VPC network traffic analysis
//
// - All TR network traffic analysis
//
// - Internet Shared Bandwidth metric analysis
//
// ## All VPC flow log analysis.
//
// ### Request parameters
//
// | Name              | Type     | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Example                                     | Valid values                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
//
// |-------------------|----------|------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
//
// | RegionNo          | string   | Yes   | The region where the resource resides.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | cn-shanghai                              | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | TrafficAnalyzerId | string   | Yes   | The traffic analyzer ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | nta-262****ca07f                         | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | TrafficScenario   | string   | Yes   | The supported analysis scenario:
//
// ● All VPC flow log analysis                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Fixed value: VpcFlowLogAll                    | VpcFlowLogAll                                                                                                                                                                                                                                                                                                                                                                                                                                         |
//
// | Direction         | string   | Yes   | The network traffic direction based on Alibaba Cloud resources. ● In: Traffic flowing into the elastic network interface (ENI).
//
// ● Out: Traffic flowing out of the elastic network interface (ENI).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | In                         |  - In
//
// - Out                                                                                                                                                                                                                                                                                                                                                                                                                                            |
//
// | TupleDimension    | string   | Yes   | The traffic storage aggregation dimension.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Tuple1                 | - Tuple1
//
// - Tuple2
//
// - Tuple5                                                                                                                                                                                                                                                                                                                                                                                                                              |
//
// | GroupBy           | array    | No   | Specifies multiple traffic dimensions for aggregation and sorting.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |  | -                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
//
// |                   | string   | No   | Based on the TupleDimension field, you can specify the following aggregation dimensions:
//
// ● When TupleDimension = Tuple1:
//
// 　■  VpcId: VPC ID
//
// 　■  VSwitchId: vSwitch ID
//
// 　■  NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■  EcsId: ECS instance ID
//
// 　■  CloudIp: Cloud IP
//
// ● When TupleDimension = Tuple2:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// ● When TupleDimension = Tuple5:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// 　■ SourcePort: Source port
//
// 　■ DestinationPort: Destination port
//
// 　■ Protocol: Network protocol | ["VpcId"]                                | - VpcId
//
// - VSwitchId
//
// - CloudIp
//
// - SourceIp
//
// - DestinationIp
//
// - Protocol
//
// - SourcePort
//
// - DestinationPort
//
// - TrafficPath
//
// - Country
//
// - Province
//
// - City
//
// - Isp
//
// - Asn                                                                                                                                                                                                                                          |
//
// | OrderBy           | string   | Yes   | Based on the TrafficScenario field, the following traffic metrics are supported for sorting:
//
// TrafficScenario = VpcFlowLogAll / VpcFlowLogInternet (VPC flow log scenario):
//
// ● Bytes: Bandwidth
//
// ● Packets: Packet count
//
// ● RoundTripTime: TCP RTT                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Bytes                                    | - Bytes
//
// - Packets
//
// - RoundTripTime                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | Filter            | array    | No   | Specifies additional filter conditions for focused traffic analysis.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |                                          | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// |                   | object   | No   | The filter condition for traffic, a Key-Value-Operator object.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                          | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | key               | string   | No   | The supported filter condition label keys are as follows:
//
// - When `TupleDimension` is 1-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (**required**, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `CloudIp`: Cloud IP (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 2-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (**required**, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP (the corresponding value supports multiple selections)
//
// 　- `DestinationIp`: Destination IP (the corresponding value supports multiple selections)
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 5-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (**required**, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP
//
// 　- `DestinationIp`: Destination IP
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// 　- `SourcePort`: Source port (the corresponding value supports multiple selections)
//
// 　- `DestinationPort`: Destination port (the corresponding value supports multiple selections)
//
// 　- `Protocol`: Network protocol (the corresponding value supports multiple selections)
//
// - In VPC scenarios, you can also filter by traffic metrics:
//
// 　- `MinBytes`: The minimum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
//
// 　- `MaxBytes`: The maximum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
//
// 　- `MinRoundTripTime`: The minimum RTT for sorting, in ms (the corresponding value does not support multiple selections)
//
// 　- `MaxRoundTripTime`: The maximum RTT for sorting, in ms (the corresponding value does not support multiple selections)
//
// 　- `MinPackages`: The minimum packet count for sorting (the corresponding value does not support multiple selections)
//
// 　- `MaxPackages`: The maximum packet count for sorting (the corresponding value does not support multiple selections) | FlowAction                               | - FlowAction
//
// - VpcId
//
// - VSwitchId
//
// - NetworkInterfaceId
//
// - CloudIp
//
// - DestinationIp
//
// - SourceIp
//
// - EcsId
//
// - TrafficPath
//
// - SourcePort
//
// - DestinationPort
//
// - Protocol
//
// - MinBytes
//
// - MaxBytes
//
// - MinRoundTripTime
//
// - MaxRoundTripTime
//
// - MinPackages
//
// - MaxPackages                                                                                                                |
//
// | value             | array    | No   | The filter condition values.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |                                          | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// |                   | string   | No   | The filter value corresponding to the specified key type.
//
// When the key is `FlowAction`, the valid values are:
//
// 　- `ACCEPT` (default: `Accept`): Traffic allowed by security groups and network ACLs
//
// 　- `REJECT`: Traffic denied by security groups and network ACLs                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Accept                                   | When the key is FlowAction: ACCEPT / REJECT                                                                                                                                                                                                                                                                                                                                                              |
//
// | Operator          | string   | No   | For specified key types, some support operators for character string matching on the passed value. The supported values are as follows (default: in):
//
// ● in: Equals
//
// ● not in: Not equals
//
// ● like: Contains
//
// Based on the TupleDimension and TrafficScenario fields, the support for `like` is as follows:
//
// TrafficScenario = VpcFlowLogAll / VpcFlowLogInternet (VPC flow log scenario):
//
// The following keys support `like`:
//
// 　○ CloudIp
//
// 　○ SourceIp
//
// 　○ DestinationIp
//
// All other keys support only `in` and `not in` operators.                                                                                                                                                                                                                                                                                                                                                                                                                           | in                                       | - not in
//
// - in
//
// - like                                                                                                                                                                                                                                                                                                                                                                                                                                   |
//
// | BeginTime         | long     | Yes   | The start time of the query as a millisecond UNIX timestamp.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | 1638239092000                            | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | EndTime           | long     | Yes   | The end time of the query as a millisecond UNIX timestamp.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | 1684373700099                            | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | Sort              | string   | No   | The sorting order for traffic analysis:
//
// ● ASC: Ascending order.
//
// ● DESC: Descending order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Desc                                     | - Desc
//
// - Asc                                                                                                                                                                                                                                                                                                                                                                                                                                           |
//
// | TopN              | integer  | No   | The number of entries to return for the traffic sorting query.
//
// You can specify a custom number. If this field is not specified, all traffic data that meets the specified conditions is sorted and analyzed within the performance limits of traffic analysis data.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | 10                                       | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | Language          | string   | No   | The language. Valid values: zh-CN, en-US.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | zh-CN                                    | - zh-CN
//
// - en-US                                                                                                                                                                                                                                                                                                                                                                                                                                        |
//
// | NextToken         | string   | No   | The pagination token. Leave this parameter empty for the first query or when no more results are available. If more results exist, set this to the NextToken value returned by the previous API call.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | MaxResults        | integer  | No   | The page size. Valid values: 1 to 100. Default value: 20.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | 20                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                         |
//
// | StorageInterval   | integer  | No   | The storage bucket granularity property.
//
// The storage bucket granularity specifies the storage aggregation period for the query. Two granularity levels are supported: high-precision (such as 1 minute) and long-period (such as 1 day). The specific granularity is determined by the traffic analysis sampling interval configured for high-precision or long-period traffic statistics when creating or editing the traffic analyzer.
//
// - The supported storage granularity values for the traffic analyzer tuple are:
//
// 　- `1`: In minutes (1 minute)
//
// 　- `10`: In minutes (10 minutes)
//
// 　- `60`: In minutes (60 minutes, i.e., 1 hour)
//
// 　- `1440`: In minutes (1440 minutes, i.e., 1 day)
//
// - The storage bucket granularity serves two typical purposes:
//
// 　- High-precision traffic statistics: such as 1-minute, 10-minute, or 60-minute aggregation
//
// 　- Long-period traffic statistics: such as 1440-minute (1-day) aggregation
//
// - Specify this field during the query to select the storage aggregation period. For example:
//
// 　- Pass `10`: Query short-period data aggregated at 10-minute granularity
//
// 　- Pass `1440`: Query long-period data aggregated at 1-day granularity | 10                                       | - 1
//
// - 10
//
// - 60
//
// - 1440                                                                                                                                                                                                                                                                                                                                                                                                                                   |
//
// ### Response elements
//
// | Name                | Type   | Description                                                                                       | Example value                    | Valid values |
//
// |---------------------|--------|--------------------------------------------------------------------------------------------|---------------------------|--------|
//
// |       | object | RpcResponse <ArrayList>                                                                    |                           | -      |
//
// | RequestId           | string | The request ID.                                                                                   | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941 | -      |
//
// | NisTrafficRankingId | string | The result ID of this traffic ranking analysis. Call the DescribeNisTrafficRanking operation to obtain the final analysis results.    | task-6462a7b4c4a54b***	- | -      |
//
// ## Internet VPC flow log analysis.
//
// ### Request parameters
//
// | Name              | Type    | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Example                                     | Valid values                                                                                                                                                                                                                                                                                                           |
//
// |-------------------|---------|------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
//
// | RegionNo          | string  | Yes   | The region where the resource resides.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | cn-shanghai                              | -                                                                                                                                                                                                                                                                                                                |
//
// | TrafficAnalyzerId | string  | Yes   | The traffic analyzer ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | nta-262****ca07f                         | -                                                                                                                                                                                                                                                                                                                |
//
// | TrafficScenario   | string  | Yes   | The supported analysis scenario:
//
// ● Internet VPC flow log analysis                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Fixed value: VpcFlowLogInternet              | VpcFlowLogInternet                                                                                                                                                                                                                                                                                               |
//
// | Direction         | string  | Yes   | The network traffic direction based on Alibaba Cloud resources. ● In: Traffic flowing into the elastic network interfaces (ENIs).
//
// ● Out: Traffic flowing out of the elastic network interfaces (ENIs).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | In                                       | - In
//
// - Out                                                                                                                                                                                                                                                                                                      |
//
// | TupleDimension    | string  | Yes   | The traffic storage aggregation dimension.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Tuple1                                   | - Tuple1
//
// - Tuple2
//
// - Tuple5                                                                                                                                                                                                                                                                                    |
//
// | GroupBy           | array   | No   | Specifies multiple traffic dimensions for aggregation and sorting.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |                                          | -                                                                                                                                                                                                                                                                                                                |
//
// |                   | string  | No   | Based on the TupleDimension field, you can specify the following aggregation dimensions:
//
// ● When TupleDimension = Tuple1:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ CloudIp: Cloud IP
//
// ● When TupleDimension = Tuple2:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// 　■ `Country`: Client country
//
// 　■ `Province`: Client province
//
// 　■ `City`: Client city
//
// 　■ `Isp`: Internet service provider
//
// 　■ `Asn`: Autonomous system number
//
// ● When TupleDimension = Tuple5:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// 　■ SourcePort: Source port
//
// 　■ DestinationPort: Destination port
//
// 　■ Protocol: Network protocol
//
// 　■ Country: Client country
//
// 　■ Province: Client province
//
// 　■ City: Client city
//
// 　■ Isp: Internet service provider
//
// 　■ Asn: Autonomous system number | ["VpcId"]                                | - VpcId
//
// - VSwitchId
//
// - CloudIp
//
// - SourceIp
//
// - DestinationIp
//
// - Protocol
//
// - SourcePort
//
// - DestinationPort
//
// - TrafficPath
//
// - Country
//
// - Province
//
// - City
//
// - Isp
//
// - Asn                                                                                         |
//
// | OrderBy           | string  | Yes   | The following traffic metrics are supported for sorting: Bytes, Packets, RoundTripTime.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Bytes                                    | - Bytes
//
// - Packets
//
// - RoundTripTime                                                                                                                                                                                                                                                                         |
//
// | Filter            | array   | No   | Specifies additional filter conditions for focused traffic analysis.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |                                          | -                                                                                                                                                                                                                                                                                                                |
//
// |                   | object  | No   | The filter condition for traffic, a Key-Value-Operator object.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |                                          | -                                                                                                                                                                                                                                                                                                                |
//
// | key               | string  | No   | The supported filter condition label keys are as follows:
//
// - When `TupleDimension` is 1-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (required, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `CloudIp`: Cloud IP (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 2-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (required, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP (the corresponding value supports multiple selections)
//
// 　- `DestinationIp`: Destination IP (the corresponding value supports multiple selections)
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// 　- `ClientCountry`: Filter traffic analysis scope by country (the corresponding value supports multiple selections)
//
// 　- `ClientCity`: Filter traffic analysis scope by city (the corresponding value supports multiple selections)
//
// 　- `ClientAsn`: Filter traffic analysis scope by ASN (the corresponding value supports multiple selections)
//
// 　- `ClientIsp`: Filter traffic analysis scope by client ISP (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 5-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (required, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP
//
// 　- `DestinationIp`: Destination IP
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// 　- `SourcePort`: Source port (the corresponding value supports multiple selections)
//
// 　- `DestinationPort`: Destination port (the corresponding value supports multiple selections)
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
// Triggers the NIS traffic analyzer to execute a traffic sorting task within a specified scope.
//
// Description:
//
// [Before using this operation, make sure that you fully understand the billing method and pricing of the NIS traffic analyzer.](https://www.alibabacloud.com/help/en/nis/product-overview/billing-method-new-version)
//
// Before using this operation, create a traffic analyzer and add a data source.
//
// [Create a traffic analyzer](https://www.alibabacloud.com/help/en/nis/user-guide/traffic-analyzer-management#39d1693bce6yp)
//
// [Add a data source](https://www.alibabacloud.com/help/en/nis/user-guide/data-source-management#73845748bfstv)
//
// Supported analysis scenarios:
//
// - All VPC network traffic analysis
//
// - Internet VPC network traffic analysis
//
// - All TR network traffic analysis
//
// - Internet Shared Bandwidth metric analysis
//
// ## All VPC flow log analysis.
//
// ### Request parameters
//
// | Name              | Type     | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Example                                     | Valid values                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
//
// |-------------------|----------|------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
//
// | RegionNo          | string   | Yes   | The region where the resource resides.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | cn-shanghai                              | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | TrafficAnalyzerId | string   | Yes   | The traffic analyzer ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | nta-262****ca07f                         | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | TrafficScenario   | string   | Yes   | The supported analysis scenario:
//
// ● All VPC flow log analysis                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Fixed value: VpcFlowLogAll                    | VpcFlowLogAll                                                                                                                                                                                                                                                                                                                                                                                                                                         |
//
// | Direction         | string   | Yes   | The network traffic direction based on Alibaba Cloud resources. ● In: Traffic flowing into the elastic network interface (ENI).
//
// ● Out: Traffic flowing out of the elastic network interface (ENI).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | In                         |  - In
//
// - Out                                                                                                                                                                                                                                                                                                                                                                                                                                            |
//
// | TupleDimension    | string   | Yes   | The traffic storage aggregation dimension.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Tuple1                 | - Tuple1
//
// - Tuple2
//
// - Tuple5                                                                                                                                                                                                                                                                                                                                                                                                                              |
//
// | GroupBy           | array    | No   | Specifies multiple traffic dimensions for aggregation and sorting.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |  | -                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
//
// |                   | string   | No   | Based on the TupleDimension field, you can specify the following aggregation dimensions:
//
// ● When TupleDimension = Tuple1:
//
// 　■  VpcId: VPC ID
//
// 　■  VSwitchId: vSwitch ID
//
// 　■  NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■  EcsId: ECS instance ID
//
// 　■  CloudIp: Cloud IP
//
// ● When TupleDimension = Tuple2:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// ● When TupleDimension = Tuple5:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// 　■ SourcePort: Source port
//
// 　■ DestinationPort: Destination port
//
// 　■ Protocol: Network protocol | ["VpcId"]                                | - VpcId
//
// - VSwitchId
//
// - CloudIp
//
// - SourceIp
//
// - DestinationIp
//
// - Protocol
//
// - SourcePort
//
// - DestinationPort
//
// - TrafficPath
//
// - Country
//
// - Province
//
// - City
//
// - Isp
//
// - Asn                                                                                                                                                                                                                                          |
//
// | OrderBy           | string   | Yes   | Based on the TrafficScenario field, the following traffic metrics are supported for sorting:
//
// TrafficScenario = VpcFlowLogAll / VpcFlowLogInternet (VPC flow log scenario):
//
// ● Bytes: Bandwidth
//
// ● Packets: Packet count
//
// ● RoundTripTime: TCP RTT                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Bytes                                    | - Bytes
//
// - Packets
//
// - RoundTripTime                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | Filter            | array    | No   | Specifies additional filter conditions for focused traffic analysis.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |                                          | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// |                   | object   | No   | The filter condition for traffic, a Key-Value-Operator object.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                          | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | key               | string   | No   | The supported filter condition label keys are as follows:
//
// - When `TupleDimension` is 1-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (**required**, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `CloudIp`: Cloud IP (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 2-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (**required**, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP (the corresponding value supports multiple selections)
//
// 　- `DestinationIp`: Destination IP (the corresponding value supports multiple selections)
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 5-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (**required**, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP
//
// 　- `DestinationIp`: Destination IP
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// 　- `SourcePort`: Source port (the corresponding value supports multiple selections)
//
// 　- `DestinationPort`: Destination port (the corresponding value supports multiple selections)
//
// 　- `Protocol`: Network protocol (the corresponding value supports multiple selections)
//
// - In VPC scenarios, you can also filter by traffic metrics:
//
// 　- `MinBytes`: The minimum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
//
// 　- `MaxBytes`: The maximum traffic volume for sorting, in bytes (the corresponding value does not support multiple selections)
//
// 　- `MinRoundTripTime`: The minimum RTT for sorting, in ms (the corresponding value does not support multiple selections)
//
// 　- `MaxRoundTripTime`: The maximum RTT for sorting, in ms (the corresponding value does not support multiple selections)
//
// 　- `MinPackages`: The minimum packet count for sorting (the corresponding value does not support multiple selections)
//
// 　- `MaxPackages`: The maximum packet count for sorting (the corresponding value does not support multiple selections) | FlowAction                               | - FlowAction
//
// - VpcId
//
// - VSwitchId
//
// - NetworkInterfaceId
//
// - CloudIp
//
// - DestinationIp
//
// - SourceIp
//
// - EcsId
//
// - TrafficPath
//
// - SourcePort
//
// - DestinationPort
//
// - Protocol
//
// - MinBytes
//
// - MaxBytes
//
// - MinRoundTripTime
//
// - MaxRoundTripTime
//
// - MinPackages
//
// - MaxPackages                                                                                                                |
//
// | value             | array    | No   | The filter condition values.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |                                          | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// |                   | string   | No   | The filter value corresponding to the specified key type.
//
// When the key is `FlowAction`, the valid values are:
//
// 　- `ACCEPT` (default: `Accept`): Traffic allowed by security groups and network ACLs
//
// 　- `REJECT`: Traffic denied by security groups and network ACLs                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Accept                                   | When the key is FlowAction: ACCEPT / REJECT                                                                                                                                                                                                                                                                                                                                                              |
//
// | Operator          | string   | No   | For specified key types, some support operators for character string matching on the passed value. The supported values are as follows (default: in):
//
// ● in: Equals
//
// ● not in: Not equals
//
// ● like: Contains
//
// Based on the TupleDimension and TrafficScenario fields, the support for `like` is as follows:
//
// TrafficScenario = VpcFlowLogAll / VpcFlowLogInternet (VPC flow log scenario):
//
// The following keys support `like`:
//
// 　○ CloudIp
//
// 　○ SourceIp
//
// 　○ DestinationIp
//
// All other keys support only `in` and `not in` operators.                                                                                                                                                                                                                                                                                                                                                                                                                           | in                                       | - not in
//
// - in
//
// - like                                                                                                                                                                                                                                                                                                                                                                                                                                   |
//
// | BeginTime         | long     | Yes   | The start time of the query as a millisecond UNIX timestamp.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | 1638239092000                            | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | EndTime           | long     | Yes   | The end time of the query as a millisecond UNIX timestamp.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | 1684373700099                            | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | Sort              | string   | No   | The sorting order for traffic analysis:
//
// ● ASC: Ascending order.
//
// ● DESC: Descending order.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Desc                                     | - Desc
//
// - Asc                                                                                                                                                                                                                                                                                                                                                                                                                                           |
//
// | TopN              | integer  | No   | The number of entries to return for the traffic sorting query.
//
// You can specify a custom number. If this field is not specified, all traffic data that meets the specified conditions is sorted and analyzed within the performance limits of traffic analysis data.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | 10                                       | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | Language          | string   | No   | The language. Valid values: zh-CN, en-US.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | zh-CN                                    | - zh-CN
//
// - en-US                                                                                                                                                                                                                                                                                                                                                                                                                                        |
//
// | NextToken         | string   | No   | The pagination token. Leave this parameter empty for the first query or when no more results are available. If more results exist, set this to the NextToken value returned by the previous API call.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+ | -                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
//
// | MaxResults        | integer  | No   | The page size. Valid values: 1 to 100. Default value: 20.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | 20                                       |                                                                                                                                                                                                                                                                                                                                                                                                                                         |
//
// | StorageInterval   | integer  | No   | The storage bucket granularity property.
//
// The storage bucket granularity specifies the storage aggregation period for the query. Two granularity levels are supported: high-precision (such as 1 minute) and long-period (such as 1 day). The specific granularity is determined by the traffic analysis sampling interval configured for high-precision or long-period traffic statistics when creating or editing the traffic analyzer.
//
// - The supported storage granularity values for the traffic analyzer tuple are:
//
// 　- `1`: In minutes (1 minute)
//
// 　- `10`: In minutes (10 minutes)
//
// 　- `60`: In minutes (60 minutes, i.e., 1 hour)
//
// 　- `1440`: In minutes (1440 minutes, i.e., 1 day)
//
// - The storage bucket granularity serves two typical purposes:
//
// 　- High-precision traffic statistics: such as 1-minute, 10-minute, or 60-minute aggregation
//
// 　- Long-period traffic statistics: such as 1440-minute (1-day) aggregation
//
// - Specify this field during the query to select the storage aggregation period. For example:
//
// 　- Pass `10`: Query short-period data aggregated at 10-minute granularity
//
// 　- Pass `1440`: Query long-period data aggregated at 1-day granularity | 10                                       | - 1
//
// - 10
//
// - 60
//
// - 1440                                                                                                                                                                                                                                                                                                                                                                                                                                   |
//
// ### Response elements
//
// | Name                | Type   | Description                                                                                       | Example value                    | Valid values |
//
// |---------------------|--------|--------------------------------------------------------------------------------------------|---------------------------|--------|
//
// |       | object | RpcResponse <ArrayList>                                                                    |                           | -      |
//
// | RequestId           | string | The request ID.                                                                                   | 4DAC4BE1-BEEA-5D84-BE06-E1B796F3B941 | -      |
//
// | NisTrafficRankingId | string | The result ID of this traffic ranking analysis. Call the DescribeNisTrafficRanking operation to obtain the final analysis results.    | task-6462a7b4c4a54b***	- | -      |
//
// ## Internet VPC flow log analysis.
//
// ### Request parameters
//
// | Name              | Type    | Required | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Example                                     | Valid values                                                                                                                                                                                                                                                                                                           |
//
// |-------------------|---------|------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
//
// | RegionNo          | string  | Yes   | The region where the resource resides.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | cn-shanghai                              | -                                                                                                                                                                                                                                                                                                                |
//
// | TrafficAnalyzerId | string  | Yes   | The traffic analyzer ID.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | nta-262****ca07f                         | -                                                                                                                                                                                                                                                                                                                |
//
// | TrafficScenario   | string  | Yes   | The supported analysis scenario:
//
// ● Internet VPC flow log analysis                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Fixed value: VpcFlowLogInternet              | VpcFlowLogInternet                                                                                                                                                                                                                                                                                               |
//
// | Direction         | string  | Yes   | The network traffic direction based on Alibaba Cloud resources. ● In: Traffic flowing into the elastic network interfaces (ENIs).
//
// ● Out: Traffic flowing out of the elastic network interfaces (ENIs).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | In                                       | - In
//
// - Out                                                                                                                                                                                                                                                                                                      |
//
// | TupleDimension    | string  | Yes   | The traffic storage aggregation dimension.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Tuple1                                   | - Tuple1
//
// - Tuple2
//
// - Tuple5                                                                                                                                                                                                                                                                                    |
//
// | GroupBy           | array   | No   | Specifies multiple traffic dimensions for aggregation and sorting.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |                                          | -                                                                                                                                                                                                                                                                                                                |
//
// |                   | string  | No   | Based on the TupleDimension field, you can specify the following aggregation dimensions:
//
// ● When TupleDimension = Tuple1:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ CloudIp: Cloud IP
//
// ● When TupleDimension = Tuple2:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// 　■ `Country`: Client country
//
// 　■ `Province`: Client province
//
// 　■ `City`: Client city
//
// 　■ `Isp`: Internet service provider
//
// 　■ `Asn`: Autonomous system number
//
// ● When TupleDimension = Tuple5:
//
// 　■ VpcId: VPC ID
//
// 　■ VSwitchId: vSwitch ID
//
// 　■ NetworkInterfaceId: elastic network interfaces (ENIs) ID
//
// 　■ EcsId: ECS instance ID
//
// 　■ SourceIp: Source IP
//
// 　■ DestinationIp: Destination IP
//
// 　■ TrafficPath: Traffic path
//
// 　■ SourcePort: Source port
//
// 　■ DestinationPort: Destination port
//
// 　■ Protocol: Network protocol
//
// 　■ Country: Client country
//
// 　■ Province: Client province
//
// 　■ City: Client city
//
// 　■ Isp: Internet service provider
//
// 　■ Asn: Autonomous system number | ["VpcId"]                                | - VpcId
//
// - VSwitchId
//
// - CloudIp
//
// - SourceIp
//
// - DestinationIp
//
// - Protocol
//
// - SourcePort
//
// - DestinationPort
//
// - TrafficPath
//
// - Country
//
// - Province
//
// - City
//
// - Isp
//
// - Asn                                                                                         |
//
// | OrderBy           | string  | Yes   | The following traffic metrics are supported for sorting: Bytes, Packets, RoundTripTime.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Bytes                                    | - Bytes
//
// - Packets
//
// - RoundTripTime                                                                                                                                                                                                                                                                         |
//
// | Filter            | array   | No   | Specifies additional filter conditions for focused traffic analysis.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |                                          | -                                                                                                                                                                                                                                                                                                                |
//
// |                   | object  | No   | The filter condition for traffic, a Key-Value-Operator object.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |                                          | -                                                                                                                                                                                                                                                                                                                |
//
// | key               | string  | No   | The supported filter condition label keys are as follows:
//
// - When `TupleDimension` is 1-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (required, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `CloudIp`: Cloud IP (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 2-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (required, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP (the corresponding value supports multiple selections)
//
// 　- `DestinationIp`: Destination IP (the corresponding value supports multiple selections)
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// 　- `ClientCountry`: Filter traffic analysis scope by country (the corresponding value supports multiple selections)
//
// 　- `ClientCity`: Filter traffic analysis scope by city (the corresponding value supports multiple selections)
//
// 　- `ClientAsn`: Filter traffic analysis scope by ASN (the corresponding value supports multiple selections)
//
// 　- `ClientIsp`: Filter traffic analysis scope by client ISP (the corresponding value supports multiple selections)
//
// - When `TupleDimension` is 5-tuple:
//
// 　- `FlowAction`: The action type executed on traffic after matching the corresponding rule or policy (required, the corresponding value does not support multiple selections)
//
// 　- `VpcId`: VPC ID (the corresponding value supports multiple selections)
//
// 　- `VSwitchId`: vSwitch ID (the corresponding value supports multiple selections)
//
// 　- `NetworkInterfaceId`: elastic network interfaces (ENIs) ID (the corresponding value supports multiple selections)
//
// 　- `EcsId`: ECS instance ID (the corresponding value supports multiple selections)
//
// 　- `SourceIp`: Source IP
//
// 　- `DestinationIp`: Destination IP
//
// 　- `TrafficPath`: Traffic path (the corresponding value supports multiple selections)
//
// 　- `SourcePort`: Source port (the corresponding value supports multiple selections)
//
// 　- `DestinationPort`: Destination port (the corresponding value supports multiple selections)
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
// Updates an inspection task.
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
// Updates an inspection task.
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
