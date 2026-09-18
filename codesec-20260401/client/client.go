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
	client.Endpoint, _err = client.GetEndpoint(dara.String("codesec"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Finalizes a code bundle after the client completes an OSS PUT operation. This operation validates the uploaded object and sets the code bundle status to ready. If CI metadata that triggers an automatic scan was provided during creation, a scanId is returned.
//
// @param request - CompleteCodeBundleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteCodeBundleResponse
func (client *Client) CompleteCodeBundleWithOptions(projectId *string, codeBundleId *string, request *CompleteCodeBundleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CompleteCodeBundleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ByteSize) {
		query["byteSize"] = request.ByteSize
	}

	if !dara.IsNil(request.ContentType) {
		query["contentType"] = request.ContentType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteCodeBundle"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/codeBundles/" + dara.PercentEncode(dara.StringValue(codeBundleId)) + "/complete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteCodeBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Finalizes a code bundle after the client completes an OSS PUT operation. This operation validates the uploaded object and sets the code bundle status to ready. If CI metadata that triggers an automatic scan was provided during creation, a scanId is returned.
//
// @param request - CompleteCodeBundleRequest
//
// @return CompleteCodeBundleResponse
func (client *Client) CompleteCodeBundle(projectId *string, codeBundleId *string, request *CompleteCodeBundleRequest) (_result *CompleteCodeBundleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CompleteCodeBundleResponse{}
	_body, _err := client.CompleteCodeBundleWithOptions(projectId, codeBundleId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a function code package in pending status and returns a pre-signed OSS PUT upload credential.
//
// @param request - CreateCodeBundleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCodeBundleResponse
func (client *Client) CreateCodeBundleWithOptions(projectId *string, request *CreateCodeBundleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCodeBundleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CiMetadata) {
		query["ciMetadata"] = request.CiMetadata
	}

	if !dara.IsNil(request.Filename) {
		query["filename"] = request.Filename
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCodeBundle"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/codeBundles"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCodeBundleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a function code package in pending status and returns a pre-signed OSS PUT upload credential.
//
// @param request - CreateCodeBundleRequest
//
// @return CreateCodeBundleResponse
func (client *Client) CreateCodeBundle(projectId *string, request *CreateCodeBundleRequest) (_result *CreateCodeBundleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCodeBundleResponse{}
	_body, _err := client.CreateCodeBundleWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// @param tmpReq - CreateProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Engines) {
		request.EnginesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Engines, dara.String("engines"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Source) {
		request.SourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Source, dara.String("source"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.EnginesShrink) {
		query["engines"] = request.EnginesShrink
	}

	if !dara.IsNil(request.InstructionPrompt) {
		query["instructionPrompt"] = request.InstructionPrompt
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.SourceShrink) {
		query["source"] = request.SourceShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a project.
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a scan task based on a code package that is ready.
//
// @param request - CreateScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScanResponse
func (client *Client) CreateScanWithOptions(projectId *string, request *CreateScanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CodeBundleId) {
		query["codeBundleId"] = request.CodeBundleId
	}

	if !dara.IsNil(request.Kind) {
		query["kind"] = request.Kind
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScan"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a scan task based on a code package that is ready.
//
// @param request - CreateScanRequest
//
// @return CreateScanResponse
func (client *Client) CreateScan(projectId *string, request *CreateScanRequest) (_result *CreateScanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScanResponse{}
	_body, _err := client.CreateScanWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成 SBOM / 许可证清单的短时下载链接
//
// @param request - CreateScanSbomExportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScanSbomExportResponse
func (client *Client) CreateScanSbomExportWithOptions(projectId *string, scanId *string, request *CreateScanSbomExportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateScanSbomExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.Format) {
		query["format"] = request.Format
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScanSbomExport"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans/" + dara.PercentEncode(dara.StringValue(scanId)) + "/reports/sbomExports"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScanSbomExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成 SBOM / 许可证清单的短时下载链接
//
// @param request - CreateScanSbomExportRequest
//
// @return CreateScanSbomExportResponse
func (client *Client) CreateScanSbomExport(projectId *string, scanId *string, request *CreateScanSbomExportRequest) (_result *CreateScanSbomExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScanSbomExportResponse{}
	_body, _err := client.CreateScanSbomExportWithOptions(projectId, scanId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists projects under a tenant by page, with support for fuzzy search by name or prompt.
//
// @param request - DescribeProjectsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProjectsResponse
func (client *Client) DescribeProjectsWithOptions(request *DescribeProjectsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeProjectsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	if !dara.IsNil(request.SortBy) {
		query["sortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortOrder) {
		query["sortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeProjects"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists projects under a tenant by page, with support for fuzzy search by name or prompt.
//
// @param request - DescribeProjectsRequest
//
// @return DescribeProjectsResponse
func (client *Client) DescribeProjects(request *DescribeProjectsRequest) (_result *DescribeProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeProjectsResponse{}
	_body, _err := client.DescribeProjectsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a scan task.
//
// @param request - DescribeScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScanResponse
func (client *Client) DescribeScanWithOptions(projectId *string, scanId *string, request *DescribeScanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeScanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScan"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans/" + dara.PercentEncode(dara.StringValue(scanId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a scan task.
//
// @param request - DescribeScanRequest
//
// @return DescribeScanResponse
func (client *Client) DescribeScan(projectId *string, scanId *string, request *DescribeScanRequest) (_result *DescribeScanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeScanResponse{}
	_body, _err := client.DescribeScanWithOptions(projectId, scanId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the task result list to retrieve detailed SAST or SCA results of a specific scan.
//
// @param request - DescribeScanResultsByEngineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScanResultsByEngineResponse
func (client *Client) DescribeScanResultsByEngineWithOptions(projectId *string, scanId *string, engine *string, request *DescribeScanResultsByEngineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeScanResultsByEngineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaselineState) {
		query["baselineState"] = request.BaselineState
	}

	if !dara.IsNil(request.Lang) {
		query["lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PackageName) {
		query["packageName"] = request.PackageName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScanResultsByEngine"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans/" + dara.PercentEncode(dara.StringValue(scanId)) + "/results/" + dara.PercentEncode(dara.StringValue(engine))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScanResultsByEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task result list to retrieve detailed SAST or SCA results of a specific scan.
//
// @param request - DescribeScanResultsByEngineRequest
//
// @return DescribeScanResultsByEngineResponse
func (client *Client) DescribeScanResultsByEngine(projectId *string, scanId *string, engine *string, request *DescribeScanResultsByEngineRequest) (_result *DescribeScanResultsByEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeScanResultsByEngineResponse{}
	_body, _err := client.DescribeScanResultsByEngineWithOptions(projectId, scanId, engine, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists scan tasks under a specified project with pagination.
//
// @param request - DescribeScansRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScansResponse
func (client *Client) DescribeScansWithOptions(projectId *string, request *DescribeScansRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeScansResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeScans"),
		Version:     dara.String("2026-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/projects/" + dara.PercentEncode(dara.StringValue(projectId)) + "/scans"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeScansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists scan tasks under a specified project with pagination.
//
// @param request - DescribeScansRequest
//
// @return DescribeScansResponse
func (client *Client) DescribeScans(projectId *string, request *DescribeScansRequest) (_result *DescribeScansResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeScansResponse{}
	_body, _err := client.DescribeScansWithOptions(projectId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
