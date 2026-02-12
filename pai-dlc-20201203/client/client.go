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
		"ap-northeast-2-pop":          dara.String("pai-dlc.aliyuncs.com"),
		"ap-south-1":                  dara.String("pai-dlc.aliyuncs.com"),
		"ap-southeast-2":              dara.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("pai-dlc.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("pai-dlc.aliyuncs.com"),
		"cn-chengdu":                  dara.String("pai-dlc.aliyuncs.com"),
		"cn-edge-1":                   dara.String("pai-dlc.aliyuncs.com"),
		"cn-fujian":                   dara.String("pai-dlc.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("pai-dlc.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("pai-dlc.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("pai-dlc.aliyuncs.com"),
		"cn-huhehaote":                dara.String("pai-dlc.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("pai-dlc.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("pai-dlc.aliyuncs.com"),
		"cn-qingdao":                  dara.String("pai-dlc.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("pai-dlc.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("pai-dlc.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("pai-dlc.aliyuncs.com"),
		"cn-wuhan":                    dara.String("pai-dlc.aliyuncs.com"),
		"cn-yushanfang":               dara.String("pai-dlc.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("pai-dlc.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("pai-dlc.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("pai-dlc.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("pai-dlc.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("pai-dlc.aliyuncs.com"),
		"eu-west-1":                   dara.String("pai-dlc.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("pai-dlc.aliyuncs.com"),
		"me-east-1":                   dara.String("pai-dlc.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("pai-dlc.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("pai-dlc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Creates a job that runs in a cluster. You can configure the data source, code source, startup command, and computing resources of each node on which a job runs.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithOptions(request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.CodeSource) {
		body["CodeSource"] = request.CodeSource
	}

	if !dara.IsNil(request.CredentialConfig) {
		body["CredentialConfig"] = request.CredentialConfig
	}

	if !dara.IsNil(request.CustomEnvs) {
		body["CustomEnvs"] = request.CustomEnvs
	}

	if !dara.IsNil(request.DataSources) {
		body["DataSources"] = request.DataSources
	}

	if !dara.IsNil(request.DebuggerConfigContent) {
		body["DebuggerConfigContent"] = request.DebuggerConfigContent
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.ElasticSpec) {
		body["ElasticSpec"] = request.ElasticSpec
	}

	if !dara.IsNil(request.Envs) {
		body["Envs"] = request.Envs
	}

	if !dara.IsNil(request.JobMaxRunningTimeMinutes) {
		body["JobMaxRunningTimeMinutes"] = request.JobMaxRunningTimeMinutes
	}

	if !dara.IsNil(request.JobSpecs) {
		body["JobSpecs"] = request.JobSpecs
	}

	if !dara.IsNil(request.JobType) {
		body["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Settings) {
		body["Settings"] = request.Settings
	}

	if !dara.IsNil(request.SuccessPolicy) {
		body["SuccessPolicy"] = request.SuccessPolicy
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateVersion) {
		body["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.ThirdpartyLibDir) {
		body["ThirdpartyLibDir"] = request.ThirdpartyLibDir
	}

	if !dara.IsNil(request.ThirdpartyLibs) {
		body["ThirdpartyLibs"] = request.ThirdpartyLibs
	}

	if !dara.IsNil(request.UserCommand) {
		body["UserCommand"] = request.UserCommand
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a job that runs in a cluster. You can configure the data source, code source, startup command, and computing resources of each node on which a job runs.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - CreateJobRequest
//
// @return CreateJobResponse
func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a TensorBoard by using a job or specifying a data source configuration.
//
// @param request - CreateTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTensorboardResponse
func (client *Client) CreateTensorboardWithOptions(request *CreateTensorboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTensorboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Cpu) {
		body["Cpu"] = request.Cpu
	}

	if !dara.IsNil(request.DataSourceId) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DataSourceType) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DataSources) {
		body["DataSources"] = request.DataSources
	}

	if !dara.IsNil(request.DisplayName) {
		body["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MaxRunningTimeMinutes) {
		body["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !dara.IsNil(request.Memory) {
		body["Memory"] = request.Memory
	}

	if !dara.IsNil(request.Options) {
		body["Options"] = request.Options
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.QuotaId) {
		body["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.SourceId) {
		body["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		body["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.SummaryPath) {
		body["SummaryPath"] = request.SummaryPath
	}

	if !dara.IsNil(request.SummaryRelativePath) {
		body["SummaryRelativePath"] = request.SummaryRelativePath
	}

	if !dara.IsNil(request.TensorboardDataSources) {
		body["TensorboardDataSources"] = request.TensorboardDataSources
	}

	if !dara.IsNil(request.TensorboardSpec) {
		body["TensorboardSpec"] = request.TensorboardSpec
	}

	if !dara.IsNil(request.Uri) {
		body["Uri"] = request.Uri
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTensorboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a TensorBoard by using a job or specifying a data source configuration.
//
// @param request - CreateTensorboardRequest
//
// @return CreateTensorboardResponse
func (client *Client) CreateTensorboard(request *CreateTensorboardRequest) (_result *CreateTensorboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTensorboardResponse{}
	_body, _err := client.CreateTensorboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a completed or stopped job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithOptions(JobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJob"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a completed or stopped job.
//
// @return DeleteJobResponse
func (client *Client) DeleteJob(JobId *string) (_result *DeleteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a stopped TensorBoard.
//
// @param request - DeleteTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTensorboardResponse
func (client *Client) DeleteTensorboardWithOptions(TensorboardId *string, request *DeleteTensorboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTensorboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTensorboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards/" + dara.PercentEncode(dara.StringValue(TensorboardId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stopped TensorBoard.
//
// @param request - DeleteTensorboardRequest
//
// @return DeleteTensorboardResponse
func (client *Client) DeleteTensorboard(TensorboardId *string, request *DeleteTensorboardRequest) (_result *DeleteTensorboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.DeleteTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the DLC task\\"s Dashboard URL, if one exists.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - GetDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDashboardResponse
func (client *Client) GetDashboardWithOptions(jobId *string, request *GetDashboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDashboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsShared) {
		query["isShared"] = request.IsShared
	}

	if !dara.IsNil(request.Token) {
		query["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDashboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/dashboard"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the DLC task\\"s Dashboard URL, if one exists.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - GetDashboardRequest
//
// @return GetDashboardResponse
func (client *Client) GetDashboard(jobId *string, request *GetDashboardRequest) (_result *GetDashboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDashboardResponse{}
	_body, _err := client.GetDashboardWithOptions(jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the configuration and runtime information of a job.
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithOptions(JobId *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NeedDetail) {
		query["NeedDetail"] = request.NeedDetail
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the configuration and runtime information of a job.
//
// @param request - GetJobRequest
//
// @return GetJobResponse
func (client *Client) GetJob(JobId *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the system events of a job.
//
// @param request - GetJobEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobEventsResponse
func (client *Client) GetJobEventsWithOptions(JobId *string, request *GetJobEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxEventsNum) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobEvents"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the system events of a job.
//
// @param request - GetJobEventsRequest
//
// @return GetJobEventsResponse
func (client *Client) GetJobEvents(JobId *string, request *GetJobEventsRequest) (_result *GetJobEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobEventsResponse{}
	_body, _err := client.GetJobEventsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the monitoring data of a job, including the CPU, GPU, and memory utilization, network, and disk read/write rate. ⚠️ Note: Except for pay-as-you-go tasks based on general-purpose computing resources, all task types are connected to CloudMonitor. Use the CloudMonitor API to call related monitoring. The overwritten features in the original API are no longer maintained. For more information, see \\[Training monitoring and alerting]\\\\(https://www.alibabacloud.com/help/zh/pai/user-guide/training-monitoring-and-alerting).
//
// @param request - GetJobMetricsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobMetricsResponse
func (client *Client) GetJobMetricsWithOptions(JobId *string, request *GetJobMetricsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobMetricsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MetricType) {
		query["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeStep) {
		query["TimeStep"] = request.TimeStep
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobMetrics"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/metrics"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the monitoring data of a job, including the CPU, GPU, and memory utilization, network, and disk read/write rate. ⚠️ Note: Except for pay-as-you-go tasks based on general-purpose computing resources, all task types are connected to CloudMonitor. Use the CloudMonitor API to call related monitoring. The overwritten features in the original API are no longer maintained. For more information, see \\[Training monitoring and alerting]\\\\(https://www.alibabacloud.com/help/zh/pai/user-guide/training-monitoring-and-alerting).
//
// @param request - GetJobMetricsRequest
//
// @return GetJobMetricsResponse
func (client *Client) GetJobMetrics(JobId *string, request *GetJobMetricsRequest) (_result *GetJobMetricsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobMetricsResponse{}
	_body, _err := client.GetJobMetricsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains specified job sanity check result in a Deep Learning Containers (DLC) job.
//
// @param request - GetJobSanityCheckResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobSanityCheckResultResponse
func (client *Client) GetJobSanityCheckResultWithOptions(JobId *string, request *GetJobSanityCheckResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobSanityCheckResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SanityCheckNumber) {
		query["SanityCheckNumber"] = request.SanityCheckNumber
	}

	if !dara.IsNil(request.SanityCheckPhase) {
		query["SanityCheckPhase"] = request.SanityCheckPhase
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobSanityCheckResult"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/sanitycheckresult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetJobSanityCheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains specified job sanity check result in a Deep Learning Containers (DLC) job.
//
// @param request - GetJobSanityCheckResultRequest
//
// @return GetJobSanityCheckResultResponse
func (client *Client) GetJobSanityCheckResult(JobId *string, request *GetJobSanityCheckResultRequest) (_result *GetJobSanityCheckResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobSanityCheckResultResponse{}
	_body, _err := client.GetJobSanityCheckResultWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the system events of a specific node in a job to locate and troubleshoot issues.
//
// @param request - GetPodEventsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPodEventsResponse
func (client *Client) GetPodEventsWithOptions(JobId *string, PodId *string, request *GetPodEventsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPodEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxEventsNum) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !dara.IsNil(request.PodUid) {
		query["PodUid"] = request.PodUid
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPodEvents"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/pods/" + dara.PercentEncode(dara.StringValue(PodId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPodEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the system events of a specific node in a job to locate and troubleshoot issues.
//
// @param request - GetPodEventsRequest
//
// @return GetPodEventsResponse
func (client *Client) GetPodEvents(JobId *string, PodId *string, request *GetPodEventsRequest) (_result *GetPodEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPodEventsResponse{}
	_body, _err := client.GetPodEventsWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains or downloads the logs of a node for a task. The logs are from the stdout and stderr of the system and user scripts.
//
// @param request - GetPodLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPodLogsResponse
func (client *Client) GetPodLogsWithOptions(JobId *string, PodId *string, request *GetPodLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPodLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DownloadToFile) {
		query["DownloadToFile"] = request.DownloadToFile
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxLines) {
		query["MaxLines"] = request.MaxLines
	}

	if !dara.IsNil(request.PodUid) {
		query["PodUid"] = request.PodUid
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPodLogs"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/pods/" + dara.PercentEncode(dara.StringValue(PodId)) + "/logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPodLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains or downloads the logs of a node for a task. The logs are from the stdout and stderr of the system and user scripts.
//
// @param request - GetPodLogsRequest
//
// @return GetPodLogsResponse
func (client *Client) GetPodLogs(JobId *string, PodId *string, request *GetPodLogsRequest) (_result *GetPodLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPodLogsResponse{}
	_body, _err := client.GetPodLogsWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a Ray Dashboard URL.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - GetRayDashboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRayDashboardResponse
func (client *Client) GetRayDashboardWithOptions(jobId *string, request *GetRayDashboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRayDashboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsShared) {
		query["isShared"] = request.IsShared
	}

	if !dara.IsNil(request.Token) {
		query["token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRayDashboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(jobId)) + "/rayDashboard"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRayDashboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a Ray Dashboard URL.
//
// Description:
//
// Before you call this operation, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/document_detail/171758.html) of Deep Learning Containers (DLC) of Platform for AI (PAI).
//
// @param request - GetRayDashboardRequest
//
// @return GetRayDashboardResponse
func (client *Client) GetRayDashboard(jobId *string, request *GetRayDashboardRequest) (_result *GetRayDashboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRayDashboardResponse{}
	_body, _err := client.GetRayDashboardWithOptions(jobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a TensorBoard instance.
//
// @param request - GetTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTensorboardResponse
func (client *Client) GetTensorboardWithOptions(TensorboardId *string, request *GetTensorboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTensorboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JodId) {
		query["JodId"] = request.JodId
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTensorboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards/" + dara.PercentEncode(dara.StringValue(TensorboardId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a TensorBoard instance.
//
// @param request - GetTensorboardRequest
//
// @return GetTensorboardResponse
func (client *Client) GetTensorboard(TensorboardId *string, request *GetTensorboardRequest) (_result *GetTensorboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTensorboardResponse{}
	_body, _err := client.GetTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the shareable link of a TensorBoard task. The link contains digital tokens. You can use a shareable link to access a TensorBoard task.
//
// @param request - GetTensorboardSharedUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTensorboardSharedUrlResponse
func (client *Client) GetTensorboardSharedUrlWithOptions(TensorboardId *string, request *GetTensorboardSharedUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTensorboardSharedUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTimeSeconds) {
		query["ExpireTimeSeconds"] = request.ExpireTimeSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTensorboardSharedUrl"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards/" + dara.PercentEncode(dara.StringValue(TensorboardId)) + "/sharedurl"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTensorboardSharedUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the shareable link of a TensorBoard task. The link contains digital tokens. You can use a shareable link to access a TensorBoard task.
//
// @param request - GetTensorboardSharedUrlRequest
//
// @return GetTensorboardSharedUrlResponse
func (client *Client) GetTensorboardSharedUrl(TensorboardId *string, request *GetTensorboardSharedUrlRequest) (_result *GetTensorboardSharedUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTensorboardSharedUrlResponse{}
	_body, _err := client.GetTensorboardSharedUrlWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the sharing token of a DLC job. This token is used to view the information about the shared job.
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.TargetId) {
		query["TargetId"] = request.TargetId
	}

	if !dara.IsNil(request.TargetType) {
		query["TargetType"] = request.TargetType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tokens"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the sharing token of a DLC job. This token is used to view the information about the shared job.
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Provides methods and steps to obtain a HTTP link for accessing a container.
//
// @param request - GetWebTerminalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWebTerminalResponse
func (client *Client) GetWebTerminalWithOptions(JobId *string, PodId *string, request *GetWebTerminalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWebTerminalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IsShared) {
		query["IsShared"] = request.IsShared
	}

	if !dara.IsNil(request.PodUid) {
		query["PodUid"] = request.PodUid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWebTerminal"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/pods/" + dara.PercentEncode(dara.StringValue(PodId)) + "/webterminal"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWebTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides methods and steps to obtain a HTTP link for accessing a container.
//
// @param request - GetWebTerminalRequest
//
// @return GetWebTerminalResponse
func (client *Client) GetWebTerminal(JobId *string, PodId *string, request *GetWebTerminalRequest) (_result *GetWebTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWebTerminalResponse{}
	_body, _err := client.GetWebTerminalWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of supported instance types.
//
// @param request - ListEcsSpecsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEcsSpecsResponse
func (client *Client) ListEcsSpecsWithOptions(request *ListEcsSpecsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEcsSpecsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceleratorType) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !dara.IsNil(request.InstanceTypes) {
		query["InstanceTypes"] = request.InstanceTypes
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEcsSpecs"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/ecsspecs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of supported instance types.
//
// @param request - ListEcsSpecsRequest
//
// @return ListEcsSpecsResponse
func (client *Client) ListEcsSpecs(request *ListEcsSpecsRequest) (_result *ListEcsSpecsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.ListEcsSpecsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the results of all sanity checks for a DLC job.
//
// @param request - ListJobSanityCheckResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobSanityCheckResultsResponse
func (client *Client) ListJobSanityCheckResultsWithOptions(JobId *string, request *ListJobSanityCheckResultsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobSanityCheckResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobSanityCheckResults"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/sanitycheckresults"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobSanityCheckResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the results of all sanity checks for a DLC job.
//
// @param request - ListJobSanityCheckResultsRequest
//
// @return ListJobSanityCheckResultsResponse
func (client *Client) ListJobSanityCheckResults(JobId *string, request *ListJobSanityCheckResultsRequest) (_result *ListJobSanityCheckResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobSanityCheckResultsResponse{}
	_body, _err := client.ListJobSanityCheckResultsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of jobs and supports pagination, sorting, and filtering by conditions.
//
// @param tmpReq - ListJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobsResponse
func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.BusinessUserId) {
		query["BusinessUserId"] = request.BusinessUserId
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.DisplayNameSearchMode) {
		query["DisplayNameSearchMode"] = request.DisplayNameSearchMode
	}

	if !dara.IsNil(request.EnableAssignNode) {
		query["EnableAssignNode"] = request.EnableAssignNode
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.FromAllWorkspaces) {
		query["FromAllWorkspaces"] = request.FromAllWorkspaces
	}

	if !dara.IsNil(request.ImageSearch) {
		query["ImageSearch"] = request.ImageSearch
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.NumericRangeField) {
		query["NumericRangeField"] = request.NumericRangeField
	}

	if !dara.IsNil(request.NumericRangeMax) {
		query["NumericRangeMax"] = request.NumericRangeMax
	}

	if !dara.IsNil(request.NumericRangeMin) {
		query["NumericRangeMin"] = request.NumericRangeMin
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OversoldInfo) {
		query["OversoldInfo"] = request.OversoldInfo
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ReasonSearch) {
		query["ReasonSearch"] = request.ReasonSearch
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceQuotaName) {
		query["ResourceQuotaName"] = request.ResourceQuotaName
	}

	if !dara.IsNil(request.ShowOwn) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TimeRangeField) {
		query["TimeRangeField"] = request.TimeRangeField
	}

	if !dara.IsNil(request.UserCommandSearch) {
		query["UserCommandSearch"] = request.UserCommandSearch
	}

	if !dara.IsNil(request.UserIdForFilter) {
		query["UserIdForFilter"] = request.UserIdForFilter
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobs"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of jobs and supports pagination, sorting, and filtering by conditions.
//
// @param request - ListJobsRequest
//
// @return ListJobsResponse
func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of TensorBoard instances.
//
// @param request - ListTensorboardsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTensorboardsResponse
func (client *Client) ListTensorboardsWithOptions(request *ListTensorboardsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTensorboardsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PaymentType) {
		query["PaymentType"] = request.PaymentType
	}

	if !dara.IsNil(request.QuotaId) {
		query["QuotaId"] = request.QuotaId
	}

	if !dara.IsNil(request.ShowOwn) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SourceId) {
		query["SourceId"] = request.SourceId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TensorboardId) {
		query["TensorboardId"] = request.TensorboardId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	if !dara.IsNil(request.Verbose) {
		query["Verbose"] = request.Verbose
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTensorboards"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTensorboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of TensorBoard instances.
//
// @param request - ListTensorboardsRequest
//
// @return ListTensorboardsResponse
func (client *Client) ListTensorboards(request *ListTensorboardsRequest) (_result *ListTensorboardsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTensorboardsResponse{}
	_body, _err := client.ListTensorboardsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a TensorBoard instance.
//
// @param request - StartTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTensorboardResponse
func (client *Client) StartTensorboardWithOptions(TensorboardId *string, request *StartTensorboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartTensorboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartTensorboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards/" + dara.PercentEncode(dara.StringValue(TensorboardId)) + "/start"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a TensorBoard instance.
//
// @param request - StartTensorboardRequest
//
// @return StartTensorboardResponse
func (client *Client) StartTensorboard(TensorboardId *string, request *StartTensorboardRequest) (_result *StartTensorboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTensorboardResponse{}
	_body, _err := client.StartTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a running job.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopJobResponse
func (client *Client) StopJobWithOptions(JobId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopJobResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopJob"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a running job.
//
// @return StopJobResponse
func (client *Client) StopJob(JobId *string) (_result *StopJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a TensorBoard instance.
//
// @param request - StopTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTensorboardResponse
func (client *Client) StopTensorboardWithOptions(TensorboardId *string, request *StopTensorboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTensorboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTensorboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards/" + dara.PercentEncode(dara.StringValue(TensorboardId)) + "/stop"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a TensorBoard instance.
//
// @param request - StopTensorboardRequest
//
// @return StopTensorboardResponse
func (client *Client) StopTensorboard(TensorboardId *string, request *StopTensorboardRequest) (_result *StopTensorboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTensorboardResponse{}
	_body, _err := client.StopTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration information of a job. For example, you can modify the priority of a job in a queue.
//
// @param request - UpdateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithOptions(JobId *string, request *UpdateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.JobSpecs) {
		body["JobSpecs"] = request.JobSpecs
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJob"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/jobs/" + dara.PercentEncode(dara.StringValue(JobId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration information of a job. For example, you can modify the priority of a job in a queue.
//
// @param request - UpdateJobRequest
//
// @return UpdateJobResponse
func (client *Client) UpdateJob(JobId *string, request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a TensorBoard instance.
//
// @param request - UpdateTensorboardRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTensorboardResponse
func (client *Client) UpdateTensorboardWithOptions(TensorboardId *string, request *UpdateTensorboardRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTensorboardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		query["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.MaxRunningTimeMinutes) {
		query["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTensorboard"),
		Version:     dara.String("2020-12-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/tensorboards/" + dara.PercentEncode(dara.StringValue(TensorboardId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a TensorBoard instance.
//
// @param request - UpdateTensorboardRequest
//
// @return UpdateTensorboardResponse
func (client *Client) UpdateTensorboard(TensorboardId *string, request *UpdateTensorboardRequest) (_result *UpdateTensorboardResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.UpdateTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
