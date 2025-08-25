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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              dara.String("docmind-api.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("docmind-api.aliyuncs.com"),
		"ap-south-1":                  dara.String("docmind-api.aliyuncs.com"),
		"ap-southeast-1":              dara.String("docmind-api.aliyuncs.com"),
		"ap-southeast-2":              dara.String("docmind-api.aliyuncs.com"),
		"ap-southeast-3":              dara.String("docmind-api.aliyuncs.com"),
		"ap-southeast-5":              dara.String("docmind-api.aliyuncs.com"),
		"cn-beijing":                  dara.String("docmind-api.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("docmind-api.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("docmind-api.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("docmind-api.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("docmind-api.aliyuncs.com"),
		"cn-chengdu":                  dara.String("docmind-api.aliyuncs.com"),
		"cn-edge-1":                   dara.String("docmind-api.aliyuncs.com"),
		"cn-fujian":                   dara.String("docmind-api.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("docmind-api.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("docmind-api.aliyuncs.com"),
		"cn-hongkong":                 dara.String("docmind-api.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("docmind-api.aliyuncs.com"),
		"cn-huhehaote":                dara.String("docmind-api.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("docmind-api.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("docmind-api.aliyuncs.com"),
		"cn-qingdao":                  dara.String("docmind-api.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("docmind-api.aliyuncs.com"),
		"cn-shanghai":                 dara.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("docmind-api.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("docmind-api.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("docmind-api.aliyuncs.com"),
		"cn-wuhan":                    dara.String("docmind-api.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("docmind-api.aliyuncs.com"),
		"cn-yushanfang":               dara.String("docmind-api.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("docmind-api.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("docmind-api.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("docmind-api.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("docmind-api.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("docmind-api.aliyuncs.com"),
		"eu-central-1":                dara.String("docmind-api.aliyuncs.com"),
		"eu-west-1":                   dara.String("docmind-api.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("docmind-api.aliyuncs.com"),
		"me-east-1":                   dara.String("docmind-api.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("docmind-api.aliyuncs.com"),
		"us-east-1":                   dara.String("docmind-api.aliyuncs.com"),
		"us-west-1":                   dara.String("docmind-api.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("docmind-api"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// 整票识别
//
// @param tmpReq - AyncTradeDocumentPackageExtractSmartAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AyncTradeDocumentPackageExtractSmartAppResponse
func (client *Client) AyncTradeDocumentPackageExtractSmartAppWithOptions(tmpReq *AyncTradeDocumentPackageExtractSmartAppRequest, runtime *dara.RuntimeOptions) (_result *AyncTradeDocumentPackageExtractSmartAppResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &AyncTradeDocumentPackageExtractSmartAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CustomExtractionRange) {
		request.CustomExtractionRangeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomExtractionRange, dara.String("CustomExtractionRange"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomExtractionRangeShrink) {
		query["CustomExtractionRange"] = request.CustomExtractionRangeShrink
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AyncTradeDocumentPackageExtractSmartApp"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AyncTradeDocumentPackageExtractSmartAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 整票识别
//
// @param request - AyncTradeDocumentPackageExtractSmartAppRequest
//
// @return AyncTradeDocumentPackageExtractSmartAppResponse
func (client *Client) AyncTradeDocumentPackageExtractSmartApp(request *AyncTradeDocumentPackageExtractSmartAppRequest) (_result *AyncTradeDocumentPackageExtractSmartAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AyncTradeDocumentPackageExtractSmartAppResponse{}
	_body, _err := client.AyncTradeDocumentPackageExtractSmartAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档结构化流式接口
//
// @param request - GetDocParserResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocParserResultResponse
func (client *Client) GetDocParserResultWithOptions(request *GetDocParserResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocParserResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.LayoutNum) {
		query["LayoutNum"] = request.LayoutNum
	}

	if !dara.IsNil(request.LayoutStepSize) {
		query["LayoutStepSize"] = request.LayoutStepSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocParserResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocParserResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档结构化流式接口
//
// @param request - GetDocParserResultRequest
//
// @return GetDocParserResultResponse
func (client *Client) GetDocParserResult(request *GetDocParserResultRequest) (_result *GetDocParserResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocParserResultResponse{}
	_body, _err := client.GetDocParserResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档智能解析结果查询
//
// @param request - GetDocStructureResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocStructureResultResponse
func (client *Client) GetDocStructureResultWithOptions(request *GetDocStructureResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocStructureResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.ImageStrategy) {
		query["ImageStrategy"] = request.ImageStrategy
	}

	if !dara.IsNil(request.RevealMarkdown) {
		query["RevealMarkdown"] = request.RevealMarkdown
	}

	if !dara.IsNil(request.UseUrlResponseBody) {
		query["UseUrlResponseBody"] = request.UseUrlResponseBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocStructureResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocStructureResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能解析结果查询
//
// @param request - GetDocStructureResultRequest
//
// @return GetDocStructureResultResponse
func (client *Client) GetDocStructureResult(request *GetDocStructureResultRequest) (_result *GetDocStructureResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocStructureResultResponse{}
	_body, _err := client.GetDocStructureResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档对比结果查询
//
// @param request - GetDocumentCompareResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentCompareResultResponse
func (client *Client) GetDocumentCompareResultWithOptions(request *GetDocumentCompareResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentCompareResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentCompareResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentCompareResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档对比结果查询
//
// @param request - GetDocumentCompareResultRequest
//
// @return GetDocumentCompareResultResponse
func (client *Client) GetDocumentCompareResult(request *GetDocumentCompareResultRequest) (_result *GetDocumentCompareResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocumentCompareResultResponse{}
	_body, _err := client.GetDocumentCompareResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档转换结果查询
//
// @param request - GetDocumentConvertResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentConvertResultResponse
func (client *Client) GetDocumentConvertResultWithOptions(request *GetDocumentConvertResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentConvertResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentConvertResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentConvertResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档转换结果查询
//
// @param request - GetDocumentConvertResultRequest
//
// @return GetDocumentConvertResultResponse
func (client *Client) GetDocumentConvertResult(request *GetDocumentConvertResultRequest) (_result *GetDocumentConvertResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocumentConvertResultResponse{}
	_body, _err := client.GetDocumentConvertResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 文档抽取结果查询
//
// @param request - GetDocumentExtractResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentExtractResultResponse
func (client *Client) GetDocumentExtractResultWithOptions(request *GetDocumentExtractResultRequest, runtime *dara.RuntimeOptions) (_result *GetDocumentExtractResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentExtractResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentExtractResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档抽取结果查询
//
// @param request - GetDocumentExtractResultRequest
//
// @return GetDocumentExtractResultResponse
func (client *Client) GetDocumentExtractResult(request *GetDocumentExtractResultRequest) (_result *GetDocumentExtractResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDocumentExtractResultResponse{}
	_body, _err := client.GetDocumentExtractResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// openmind
//
// @param request - GetPageNumRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPageNumResponse
func (client *Client) GetPageNumWithOptions(request *GetPageNumRequest, runtime *dara.RuntimeOptions) (_result *GetPageNumResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPageNum"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPageNumResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// openmind
//
// @param request - GetPageNumRequest
//
// @return GetPageNumResponse
func (client *Client) GetPageNum(request *GetPageNumRequest) (_result *GetPageNumResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPageNumResponse{}
	_body, _err := client.GetPageNumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 表格智能解析结果查询
//
// @param request - GetTableUnderstandingResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableUnderstandingResultResponse
func (client *Client) GetTableUnderstandingResultWithOptions(request *GetTableUnderstandingResultRequest, runtime *dara.RuntimeOptions) (_result *GetTableUnderstandingResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableUnderstandingResult"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableUnderstandingResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 表格智能解析结果查询
//
// @param request - GetTableUnderstandingResultRequest
//
// @return GetTableUnderstandingResultResponse
func (client *Client) GetTableUnderstandingResult(request *GetTableUnderstandingResultRequest) (_result *GetTableUnderstandingResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableUnderstandingResultResponse{}
	_body, _err := client.GetTableUnderstandingResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文档智能解析处理状态
//
// @param request - QueryDocParserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDocParserStatusResponse
func (client *Client) QueryDocParserStatusWithOptions(request *QueryDocParserStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDocParserStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDocParserStatus"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDocParserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档智能解析处理状态
//
// @param request - QueryDocParserStatusRequest
//
// @return QueryDocParserStatusResponse
func (client *Client) QueryDocParserStatus(request *QueryDocParserStatusRequest) (_result *QueryDocParserStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDocParserStatusResponse{}
	_body, _err := client.QueryDocParserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片转excel
//
// @param tmpReq - SubmitConvertImageToExcelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToExcelJobResponse
func (client *Client) SubmitConvertImageToExcelJobWithOptions(tmpReq *SubmitConvertImageToExcelJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToExcelJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToExcelJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceMergeExcel) {
		query["ForceMergeExcel"] = request.ForceMergeExcel
	}

	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToExcelJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToExcelJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转excel
//
// @param request - SubmitConvertImageToExcelJobRequest
//
// @return SubmitConvertImageToExcelJobResponse
func (client *Client) SubmitConvertImageToExcelJob(request *SubmitConvertImageToExcelJobRequest) (_result *SubmitConvertImageToExcelJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertImageToExcelJobResponse{}
	_body, _err := client.SubmitConvertImageToExcelJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片转markdown
//
// @param tmpReq - SubmitConvertImageToMarkdownJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToMarkdownJobResponse
func (client *Client) SubmitConvertImageToMarkdownJobWithOptions(tmpReq *SubmitConvertImageToMarkdownJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToMarkdownJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToMarkdownJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToMarkdownJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToMarkdownJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转markdown
//
// @param request - SubmitConvertImageToMarkdownJobRequest
//
// @return SubmitConvertImageToMarkdownJobResponse
func (client *Client) SubmitConvertImageToMarkdownJob(request *SubmitConvertImageToMarkdownJobRequest) (_result *SubmitConvertImageToMarkdownJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertImageToMarkdownJobResponse{}
	_body, _err := client.SubmitConvertImageToMarkdownJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片转pdf
//
// @param tmpReq - SubmitConvertImageToPdfJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToPdfJobResponse
func (client *Client) SubmitConvertImageToPdfJobWithOptions(tmpReq *SubmitConvertImageToPdfJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToPdfJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToPdfJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToPdfJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToPdfJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转pdf
//
// @param request - SubmitConvertImageToPdfJobRequest
//
// @return SubmitConvertImageToPdfJobResponse
func (client *Client) SubmitConvertImageToPdfJob(request *SubmitConvertImageToPdfJobRequest) (_result *SubmitConvertImageToPdfJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertImageToPdfJobResponse{}
	_body, _err := client.SubmitConvertImageToPdfJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片转word
//
// @param tmpReq - SubmitConvertImageToWordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertImageToWordJobResponse
func (client *Client) SubmitConvertImageToWordJobWithOptions(tmpReq *SubmitConvertImageToWordJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertImageToWordJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitConvertImageToWordJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ImageNames) {
		request.ImageNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageNames, dara.String("ImageNames"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.ImageUrls) {
		request.ImageUrlsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ImageUrls, dara.String("ImageUrls"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ImageNameExtension) {
		query["ImageNameExtension"] = request.ImageNameExtension
	}

	if !dara.IsNil(request.ImageNamesShrink) {
		query["ImageNames"] = request.ImageNamesShrink
	}

	if !dara.IsNil(request.ImageUrlsShrink) {
		query["ImageUrls"] = request.ImageUrlsShrink
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertImageToWordJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertImageToWordJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片转word
//
// @param request - SubmitConvertImageToWordJobRequest
//
// @return SubmitConvertImageToWordJobResponse
func (client *Client) SubmitConvertImageToWordJob(request *SubmitConvertImageToWordJobRequest) (_result *SubmitConvertImageToWordJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertImageToWordJobResponse{}
	_body, _err := client.SubmitConvertImageToWordJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// pdf转excel
//
// @param request - SubmitConvertPdfToExcelJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToExcelJobResponse
func (client *Client) SubmitConvertPdfToExcelJobWithOptions(request *SubmitConvertPdfToExcelJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ForceExportInnerImage) {
		query["ForceExportInnerImage"] = request.ForceExportInnerImage
	}

	if !dara.IsNil(request.ForceMergeExcel) {
		query["ForceMergeExcel"] = request.ForceMergeExcel
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToExcelJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToExcelJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转excel
//
// @param request - SubmitConvertPdfToExcelJobRequest
//
// @return SubmitConvertPdfToExcelJobResponse
func (client *Client) SubmitConvertPdfToExcelJob(request *SubmitConvertPdfToExcelJobRequest) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertPdfToExcelJobResponse{}
	_body, _err := client.SubmitConvertPdfToExcelJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToExcelJobAdvance(request *SubmitConvertPdfToExcelJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToExcelJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitConvertPdfToExcelJobReq := &SubmitConvertPdfToExcelJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToExcelJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToExcelJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitConvertPdfToExcelJobResp, _err := client.SubmitConvertPdfToExcelJobWithOptions(submitConvertPdfToExcelJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToExcelJobResp
	return _result, _err
}

// Summary:
//
// pdf转图片
//
// @param request - SubmitConvertPdfToImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToImageJobResponse
func (client *Client) SubmitConvertPdfToImageJobWithOptions(request *SubmitConvertPdfToImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToImageJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToImageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转图片
//
// @param request - SubmitConvertPdfToImageJobRequest
//
// @return SubmitConvertPdfToImageJobResponse
func (client *Client) SubmitConvertPdfToImageJob(request *SubmitConvertPdfToImageJobRequest) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertPdfToImageJobResponse{}
	_body, _err := client.SubmitConvertPdfToImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToImageJobAdvance(request *SubmitConvertPdfToImageJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToImageJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitConvertPdfToImageJobReq := &SubmitConvertPdfToImageJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToImageJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToImageJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitConvertPdfToImageJobResp, _err := client.SubmitConvertPdfToImageJobWithOptions(submitConvertPdfToImageJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToImageJobResp
	return _result, _err
}

// Summary:
//
// pdf转markdown
//
// @param request - SubmitConvertPdfToMarkdownJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToMarkdownJobResponse
func (client *Client) SubmitConvertPdfToMarkdownJobWithOptions(request *SubmitConvertPdfToMarkdownJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToMarkdownJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToMarkdownJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转markdown
//
// @param request - SubmitConvertPdfToMarkdownJobRequest
//
// @return SubmitConvertPdfToMarkdownJobResponse
func (client *Client) SubmitConvertPdfToMarkdownJob(request *SubmitConvertPdfToMarkdownJobRequest) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertPdfToMarkdownJobResponse{}
	_body, _err := client.SubmitConvertPdfToMarkdownJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToMarkdownJobAdvance(request *SubmitConvertPdfToMarkdownJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToMarkdownJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitConvertPdfToMarkdownJobReq := &SubmitConvertPdfToMarkdownJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToMarkdownJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToMarkdownJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitConvertPdfToMarkdownJobResp, _err := client.SubmitConvertPdfToMarkdownJobWithOptions(submitConvertPdfToMarkdownJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToMarkdownJobResp
	return _result, _err
}

// Summary:
//
// pdf转word
//
// @param request - SubmitConvertPdfToWordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitConvertPdfToWordJobResponse
func (client *Client) SubmitConvertPdfToWordJobWithOptions(request *SubmitConvertPdfToWordJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ForceExportInnerImage) {
		query["ForceExportInnerImage"] = request.ForceExportInnerImage
	}

	if !dara.IsNil(request.FormulaEnhancement) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitConvertPdfToWordJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitConvertPdfToWordJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pdf转word
//
// @param request - SubmitConvertPdfToWordJobRequest
//
// @return SubmitConvertPdfToWordJobResponse
func (client *Client) SubmitConvertPdfToWordJob(request *SubmitConvertPdfToWordJobRequest) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitConvertPdfToWordJobResponse{}
	_body, _err := client.SubmitConvertPdfToWordJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitConvertPdfToWordJobAdvance(request *SubmitConvertPdfToWordJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitConvertPdfToWordJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitConvertPdfToWordJobReq := &SubmitConvertPdfToWordJobRequest{}
	openapiutil.Convert(request, submitConvertPdfToWordJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitConvertPdfToWordJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitConvertPdfToWordJobResp, _err := client.SubmitConvertPdfToWordJobWithOptions(submitConvertPdfToWordJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitConvertPdfToWordJobResp
	return _result, _err
}

// Summary:
//
// 电子解析
//
// @param request - SubmitDigitalDocStructureJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDigitalDocStructureJobResponse
func (client *Client) SubmitDigitalDocStructureJobWithOptions(request *SubmitDigitalDocStructureJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.ImageStrategy) {
		query["ImageStrategy"] = request.ImageStrategy
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.RevealMarkdown) {
		query["RevealMarkdown"] = request.RevealMarkdown
	}

	if !dara.IsNil(request.UseUrlResponseBody) {
		query["UseUrlResponseBody"] = request.UseUrlResponseBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDigitalDocStructureJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDigitalDocStructureJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 电子解析
//
// @param request - SubmitDigitalDocStructureJobRequest
//
// @return SubmitDigitalDocStructureJobResponse
func (client *Client) SubmitDigitalDocStructureJob(request *SubmitDigitalDocStructureJobRequest) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDigitalDocStructureJobResponse{}
	_body, _err := client.SubmitDigitalDocStructureJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDigitalDocStructureJobAdvance(request *SubmitDigitalDocStructureJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitDigitalDocStructureJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDigitalDocStructureJobReq := &SubmitDigitalDocStructureJobRequest{}
	openapiutil.Convert(request, submitDigitalDocStructureJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDigitalDocStructureJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDigitalDocStructureJobResp, _err := client.SubmitDigitalDocStructureJobWithOptions(submitDigitalDocStructureJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDigitalDocStructureJobResp
	return _result, _err
}

// Summary:
//
// 文档智能解析流式输出
//
// @param tmpReq - SubmitDocParserJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocParserJobResponse
func (client *Client) SubmitDocParserJobWithOptions(tmpReq *SubmitDocParserJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocParserJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitDocParserJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MultimediaParameters) {
		request.MultimediaParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MultimediaParameters, dara.String("MultimediaParameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnhancementMode) {
		query["EnhancementMode"] = request.EnhancementMode
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FormulaEnhancement) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !dara.IsNil(request.LlmEnhancement) {
		query["LlmEnhancement"] = request.LlmEnhancement
	}

	if !dara.IsNil(request.MultimediaParametersShrink) {
		query["MultimediaParameters"] = request.MultimediaParametersShrink
	}

	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.OutputHtmlTable) {
		query["OutputHtmlTable"] = request.OutputHtmlTable
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocParserJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocParserJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能解析流式输出
//
// @param request - SubmitDocParserJobRequest
//
// @return SubmitDocParserJobResponse
func (client *Client) SubmitDocParserJob(request *SubmitDocParserJobRequest) (_result *SubmitDocParserJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDocParserJobResponse{}
	_body, _err := client.SubmitDocParserJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocParserJobAdvance(request *SubmitDocParserJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocParserJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDocParserJobReq := &SubmitDocParserJobRequest{}
	openapiutil.Convert(request, submitDocParserJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocParserJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocParserJobResp, _err := client.SubmitDocParserJobWithOptions(submitDocParserJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocParserJobResp
	return _result, _err
}

// Summary:
//
// 文档智能解析
//
// @param request - SubmitDocStructureJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocStructureJobResponse
func (client *Client) SubmitDocStructureJobWithOptions(request *SubmitDocStructureJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocStructureJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowPptFormat) {
		query["AllowPptFormat"] = request.AllowPptFormat
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FormulaEnhancement) {
		query["FormulaEnhancement"] = request.FormulaEnhancement
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.StructureType) {
		query["StructureType"] = request.StructureType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocStructureJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocStructureJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档智能解析
//
// @param request - SubmitDocStructureJobRequest
//
// @return SubmitDocStructureJobResponse
func (client *Client) SubmitDocStructureJob(request *SubmitDocStructureJobRequest) (_result *SubmitDocStructureJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDocStructureJobResponse{}
	_body, _err := client.SubmitDocStructureJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocStructureJobAdvance(request *SubmitDocStructureJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocStructureJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDocStructureJobReq := &SubmitDocStructureJobRequest{}
	openapiutil.Convert(request, submitDocStructureJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocStructureJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocStructureJobResp, _err := client.SubmitDocStructureJobWithOptions(submitDocStructureJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocStructureJobResp
	return _result, _err
}

// Summary:
//
// 文档抽取
//
// @param request - SubmitDocumentExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocumentExtractJobResponse
func (client *Client) SubmitDocumentExtractJobWithOptions(request *SubmitDocumentExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocumentExtractJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocumentExtractJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocumentExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档抽取
//
// @param request - SubmitDocumentExtractJobRequest
//
// @return SubmitDocumentExtractJobResponse
func (client *Client) SubmitDocumentExtractJob(request *SubmitDocumentExtractJobRequest) (_result *SubmitDocumentExtractJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDocumentExtractJobResponse{}
	_body, _err := client.SubmitDocumentExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocumentExtractJobAdvance(request *SubmitDocumentExtractJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitDocumentExtractJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDocumentExtractJobReq := &SubmitDocumentExtractJobRequest{}
	openapiutil.Convert(request, submitDocumentExtractJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocumentExtractJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocumentExtractJobResp, _err := client.SubmitDocumentExtractJobWithOptions(submitDocumentExtractJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocumentExtractJobResp
	return _result, _err
}

// Summary:
//
// 表格智能解析
//
// @param request - SubmitTableUnderstandingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTableUnderstandingJobResponse
func (client *Client) SubmitTableUnderstandingJobWithOptions(request *SubmitTableUnderstandingJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileNameExtension) {
		query["FileNameExtension"] = request.FileNameExtension
	}

	if !dara.IsNil(request.FileUrl) {
		query["FileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.OssBucket) {
		query["OssBucket"] = request.OssBucket
	}

	if !dara.IsNil(request.OssEndpoint) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTableUnderstandingJob"),
		Version:     dara.String("2022-07-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTableUnderstandingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 表格智能解析
//
// @param request - SubmitTableUnderstandingJobRequest
//
// @return SubmitTableUnderstandingJobResponse
func (client *Client) SubmitTableUnderstandingJob(request *SubmitTableUnderstandingJobRequest) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTableUnderstandingJobResponse{}
	_body, _err := client.SubmitTableUnderstandingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitTableUnderstandingJobAdvance(request *SubmitTableUnderstandingJobAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SubmitTableUnderstandingJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("docmind-api"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitTableUnderstandingJobReq := &SubmitTableUnderstandingJobRequest{}
	openapiutil.Convert(request, submitTableUnderstandingJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitTableUnderstandingJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitTableUnderstandingJobResp, _err := client.SubmitTableUnderstandingJobWithOptions(submitTableUnderstandingJobReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitTableUnderstandingJobResp
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
