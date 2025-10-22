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
		"ap-northeast-1":              dara.String("companyreg.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("companyreg.aliyuncs.com"),
		"ap-south-1":                  dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-1":              dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-2":              dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-3":              dara.String("companyreg.aliyuncs.com"),
		"ap-southeast-5":              dara.String("companyreg.aliyuncs.com"),
		"cn-beijing":                  dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("companyreg.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("companyreg.aliyuncs.com"),
		"cn-chengdu":                  dara.String("companyreg.aliyuncs.com"),
		"cn-edge-1":                   dara.String("companyreg.aliyuncs.com"),
		"cn-fujian":                   dara.String("companyreg.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("companyreg.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("companyreg.aliyuncs.com"),
		"cn-hongkong":                 dara.String("companyreg.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("companyreg.aliyuncs.com"),
		"cn-huhehaote":                dara.String("companyreg.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("companyreg.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("companyreg.aliyuncs.com"),
		"cn-qingdao":                  dara.String("companyreg.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai":                 dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("companyreg.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("companyreg.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("companyreg.aliyuncs.com"),
		"cn-wuhan":                    dara.String("companyreg.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("companyreg.aliyuncs.com"),
		"cn-yushanfang":               dara.String("companyreg.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("companyreg.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("companyreg.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("companyreg.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("companyreg.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("companyreg.aliyuncs.com"),
		"eu-central-1":                dara.String("companyreg.aliyuncs.com"),
		"eu-west-1":                   dara.String("companyreg.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("companyreg.aliyuncs.com"),
		"me-east-1":                   dara.String("companyreg.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("companyreg.aliyuncs.com"),
		"us-east-1":                   dara.String("companyreg.aliyuncs.com"),
		"us-west-1":                   dara.String("companyreg.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("companyreg"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - BindProduceAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindProduceAuthorizationResponse
func (client *Client) BindProduceAuthorizationWithOptions(request *BindProduceAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *BindProduceAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedUserIds) {
		body["AuthorizedUserIds"] = request.AuthorizedUserIds
	}

	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindProduceAuthorization"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindProduceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - BindProduceAuthorizationRequest
//
// @return BindProduceAuthorizationResponse
func (client *Client) BindProduceAuthorization(request *BindProduceAuthorizationRequest) (_result *BindProduceAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindProduceAuthorizationResponse{}
	_body, _err := client.BindProduceAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CloseIntentionForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseIntentionForPartnerResponse
func (client *Client) CloseIntentionForPartnerWithOptions(request *CloseIntentionForPartnerRequest, runtime *dara.RuntimeOptions) (_result *CloseIntentionForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseIntentionForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseIntentionForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloseIntentionForPartnerRequest
//
// @return CloseIntentionForPartnerResponse
func (client *Client) CloseIntentionForPartner(request *CloseIntentionForPartnerRequest) (_result *CloseIntentionForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseIntentionForPartnerResponse{}
	_body, _err := client.CloseIntentionForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CloseUserIntentionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseUserIntentionResponse
func (client *Client) CloseUserIntentionWithOptions(request *CloseUserIntentionRequest, runtime *dara.RuntimeOptions) (_result *CloseUserIntentionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseUserIntention"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseUserIntentionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CloseUserIntentionRequest
//
// @return CloseUserIntentionResponse
func (client *Client) CloseUserIntention(request *CloseUserIntentionRequest) (_result *CloseUserIntentionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseUserIntentionResponse{}
	_body, _err := client.CloseUserIntentionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CreateBusinessOpportunity
//
// @param request - CreateBusinessOpportunityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBusinessOpportunityResponse
func (client *Client) CreateBusinessOpportunityWithOptions(request *CreateBusinessOpportunityRequest, runtime *dara.RuntimeOptions) (_result *CreateBusinessOpportunityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.VCode) {
		query["VCode"] = request.VCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBusinessOpportunity"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBusinessOpportunityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateBusinessOpportunity
//
// @param request - CreateBusinessOpportunityRequest
//
// @return CreateBusinessOpportunityResponse
func (client *Client) CreateBusinessOpportunity(request *CreateBusinessOpportunityRequest) (_result *CreateBusinessOpportunityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBusinessOpportunityResponse{}
	_body, _err := client.CreateBusinessOpportunityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateProduceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProduceForPartnerResponse
func (client *Client) CreateProduceForPartnerWithOptions(request *CreateProduceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *CreateProduceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProduceForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProduceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateProduceForPartnerRequest
//
// @return CreateProduceForPartnerResponse
func (client *Client) CreateProduceForPartner(request *CreateProduceForPartnerRequest) (_result *CreateProduceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProduceForPartnerResponse{}
	_body, _err := client.CreateProduceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribePartnerConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePartnerConfigResponse
func (client *Client) DescribePartnerConfigWithOptions(request *DescribePartnerConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribePartnerConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PartnerCode) {
		query["PartnerCode"] = request.PartnerCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePartnerConfig"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePartnerConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePartnerConfigRequest
//
// @return DescribePartnerConfigResponse
func (client *Client) DescribePartnerConfig(request *DescribePartnerConfigRequest) (_result *DescribePartnerConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePartnerConfigResponse{}
	_body, _err := client.DescribePartnerConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GenerateUploadFilePolicy
//
// @param request - GenerateUploadFilePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateUploadFilePolicyResponse
func (client *Client) GenerateUploadFilePolicyWithOptions(request *GenerateUploadFilePolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateUploadFilePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["FileType"] = request.FileType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateUploadFilePolicy"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GenerateUploadFilePolicy
//
// @param request - GenerateUploadFilePolicyRequest
//
// @return GenerateUploadFilePolicyResponse
func (client *Client) GenerateUploadFilePolicy(request *GenerateUploadFilePolicyRequest) (_result *GenerateUploadFilePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateUploadFilePolicyResponse{}
	_body, _err := client.GenerateUploadFilePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAlipayUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlipayUrlResponse
func (client *Client) GetAlipayUrlWithOptions(request *GetAlipayUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAlipayUrlResponse, _err error) {
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
		Action:      dara.String("GetAlipayUrl"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetAlipayUrlRequest
//
// @return GetAlipayUrlResponse
func (client *Client) GetAlipayUrl(request *GetAlipayUrlRequest) (_result *GetAlipayUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAlipayUrlResponse{}
	_body, _err := client.GetAlipayUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListIntentionNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIntentionNoteResponse
func (client *Client) ListIntentionNoteWithOptions(request *ListIntentionNoteRequest, runtime *dara.RuntimeOptions) (_result *ListIntentionNoteResponse, _err error) {
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

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIntentionNote"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIntentionNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListIntentionNoteRequest
//
// @return ListIntentionNoteResponse
func (client *Client) ListIntentionNote(request *ListIntentionNoteRequest) (_result *ListIntentionNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIntentionNoteResponse{}
	_body, _err := client.ListIntentionNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListProduceAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProduceAuthorizationResponse
func (client *Client) ListProduceAuthorizationWithOptions(request *ListProduceAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ListProduceAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProduceAuthorization"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProduceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListProduceAuthorizationRequest
//
// @return ListProduceAuthorizationResponse
func (client *Client) ListProduceAuthorization(request *ListProduceAuthorizationRequest) (_result *ListProduceAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProduceAuthorizationResponse{}
	_body, _err := client.ListProduceAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserDetailSolutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDetailSolutionsResponse
func (client *Client) ListUserDetailSolutionsWithOptions(request *ListUserDetailSolutionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserDetailSolutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserDetailSolutions"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserDetailSolutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserDetailSolutionsRequest
//
// @return ListUserDetailSolutionsResponse
func (client *Client) ListUserDetailSolutions(request *ListUserDetailSolutionsRequest) (_result *ListUserDetailSolutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserDetailSolutionsResponse{}
	_body, _err := client.ListUserDetailSolutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserIntentionNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserIntentionNotesResponse
func (client *Client) ListUserIntentionNotesWithOptions(request *ListUserIntentionNotesRequest, runtime *dara.RuntimeOptions) (_result *ListUserIntentionNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserIntentionNotes"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserIntentionNotesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserIntentionNotesRequest
//
// @return ListUserIntentionNotesResponse
func (client *Client) ListUserIntentionNotes(request *ListUserIntentionNotesRequest) (_result *ListUserIntentionNotesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserIntentionNotesResponse{}
	_body, _err := client.ListUserIntentionNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户控制天需求列表查询
//
// @param request - ListUserIntentionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserIntentionsResponse
func (client *Client) ListUserIntentionsWithOptions(request *ListUserIntentionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserIntentionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.BizTypes) {
		query["BizTypes"] = request.BizTypes
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortFiled) {
		query["SortFiled"] = request.SortFiled
	}

	if !dara.IsNil(request.SortOrder) {
		query["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WithExtInfo) {
		query["WithExtInfo"] = request.WithExtInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserIntentions"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserIntentionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户控制天需求列表查询
//
// @param request - ListUserIntentionsRequest
//
// @return ListUserIntentionsResponse
func (client *Client) ListUserIntentions(request *ListUserIntentionsRequest) (_result *ListUserIntentionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserIntentionsResponse{}
	_body, _err := client.ListUserIntentionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListUserProduceOperateLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserProduceOperateLogsResponse
func (client *Client) ListUserProduceOperateLogsWithOptions(request *ListUserProduceOperateLogsRequest, runtime *dara.RuntimeOptions) (_result *ListUserProduceOperateLogsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserProduceOperateLogs"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserProduceOperateLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserProduceOperateLogsRequest
//
// @return ListUserProduceOperateLogsResponse
func (client *Client) ListUserProduceOperateLogs(request *ListUserProduceOperateLogsRequest) (_result *ListUserProduceOperateLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserProduceOperateLogsResponse{}
	_body, _err := client.ListUserProduceOperateLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - ListUserSolutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserSolutionsResponse
func (client *Client) ListUserSolutionsWithOptions(tmpReq *ListUserSolutionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserSolutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListUserSolutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ExistStatus) {
		request.ExistStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExistStatus, dara.String("ExistStatus"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ExistStatusShrink) {
		query["ExistStatus"] = request.ExistStatusShrink
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserSolutions"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserSolutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListUserSolutionsRequest
//
// @return ListUserSolutionsResponse
func (client *Client) ListUserSolutions(request *ListUserSolutionsRequest) (_result *ListUserSolutionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserSolutionsResponse{}
	_body, _err := client.ListUserSolutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务商玄坛呼叫中心操作
//
// @param request - OperateCallCenterForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateCallCenterForPartnerResponse
func (client *Client) OperateCallCenterForPartnerWithOptions(request *OperateCallCenterForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateCallCenterForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CallAction) {
		query["CallAction"] = request.CallAction
	}

	if !dara.IsNil(request.EmployeeCode) {
		query["EmployeeCode"] = request.EmployeeCode
	}

	if !dara.IsNil(request.Request) {
		query["Request"] = request.Request
	}

	if !dara.IsNil(request.TenantId) {
		query["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateCallCenterForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateCallCenterForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商玄坛呼叫中心操作
//
// @param request - OperateCallCenterForPartnerRequest
//
// @return OperateCallCenterForPartnerResponse
func (client *Client) OperateCallCenterForPartner(request *OperateCallCenterForPartnerRequest) (_result *OperateCallCenterForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateCallCenterForPartnerResponse{}
	_body, _err := client.OperateCallCenterForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OperateProduceForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateProduceForPartnerResponse
func (client *Client) OperateProduceForPartnerWithOptions(request *OperateProduceForPartnerRequest, runtime *dara.RuntimeOptions) (_result *OperateProduceForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateProduceForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateProduceForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - OperateProduceForPartnerRequest
//
// @return OperateProduceForPartnerResponse
func (client *Client) OperateProduceForPartner(request *OperateProduceForPartnerRequest) (_result *OperateProduceForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OperateProduceForPartnerResponse{}
	_body, _err := client.OperateProduceForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PutMeasureDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMeasureDataResponse
func (client *Client) PutMeasureDataWithOptions(request *PutMeasureDataRequest, runtime *dara.RuntimeOptions) (_result *PutMeasureDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Data) {
		body["Data"] = request.Data
	}

	if !dara.IsNil(request.DataType) {
		body["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMeasureData"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutMeasureDataRequest
//
// @return PutMeasureDataResponse
func (client *Client) PutMeasureData(request *PutMeasureDataRequest) (_result *PutMeasureDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutMeasureDataResponse{}
	_body, _err := client.PutMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PutMeasureReadyFlagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PutMeasureReadyFlagResponse
func (client *Client) PutMeasureReadyFlagWithOptions(request *PutMeasureReadyFlagRequest, runtime *dara.RuntimeOptions) (_result *PutMeasureReadyFlagResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ReadyFlag) {
		query["ReadyFlag"] = request.ReadyFlag
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PutMeasureReadyFlag"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - PutMeasureReadyFlagRequest
//
// @return PutMeasureReadyFlagResponse
func (client *Client) PutMeasureReadyFlag(request *PutMeasureReadyFlagRequest) (_result *PutMeasureReadyFlagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PutMeasureReadyFlagResponse{}
	_body, _err := client.PutMeasureReadyFlagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取玄坛合作伙伴双呼时可用外呼号码
//
// @param request - QueryAvailableNumbersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvailableNumbersResponse
func (client *Client) QueryAvailableNumbersWithOptions(request *QueryAvailableNumbersRequest, runtime *dara.RuntimeOptions) (_result *QueryAvailableNumbersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvailableNumbers"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvailableNumbersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取玄坛合作伙伴双呼时可用外呼号码
//
// @param request - QueryAvailableNumbersRequest
//
// @return QueryAvailableNumbersResponse
func (client *Client) QueryAvailableNumbers(request *QueryAvailableNumbersRequest) (_result *QueryAvailableNumbersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAvailableNumbersResponse{}
	_body, _err := client.QueryAvailableNumbersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryBagRemainingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryBagRemainingResponse
func (client *Client) QueryBagRemainingWithOptions(request *QueryBagRemainingRequest, runtime *dara.RuntimeOptions) (_result *QueryBagRemainingResponse, _err error) {
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
		Action:      dara.String("QueryBagRemaining"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryBagRemainingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryBagRemainingRequest
//
// @return QueryBagRemainingResponse
func (client *Client) QueryBagRemaining(request *QueryBagRemainingRequest) (_result *QueryBagRemainingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryBagRemainingResponse{}
	_body, _err := client.QueryBagRemainingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询玄坛外呼语音文件
//
// @param request - QueryCallRecordListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallRecordListResponse
func (client *Client) QueryCallRecordListWithOptions(request *QueryCallRecordListRequest, runtime *dara.RuntimeOptions) (_result *QueryCallRecordListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.SkillType) {
		query["SkillType"] = request.SkillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallRecordList"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallRecordListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询玄坛外呼语音文件
//
// @param request - QueryCallRecordListRequest
//
// @return QueryCallRecordListResponse
func (client *Client) QueryCallRecordList(request *QueryCallRecordListRequest) (_result *QueryCallRecordListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallRecordListResponse{}
	_body, _err := client.QueryCallRecordListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceResponse
func (client *Client) QueryInstanceWithOptions(request *QueryInstanceRequest, runtime *dara.RuntimeOptions) (_result *QueryInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInstance"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryInstanceRequest
//
// @return QueryInstanceResponse
func (client *Client) QueryInstance(request *QueryInstanceRequest) (_result *QueryInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInstanceResponse{}
	_body, _err := client.QueryInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryPartnerIntentionList
//
// @param request - QueryPartnerIntentionListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPartnerIntentionListResponse
func (client *Client) QueryPartnerIntentionListWithOptions(request *QueryPartnerIntentionListRequest, runtime *dara.RuntimeOptions) (_result *QueryPartnerIntentionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPartnerIntentionList"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPartnerIntentionListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryPartnerIntentionList
//
// @param request - QueryPartnerIntentionListRequest
//
// @return QueryPartnerIntentionListResponse
func (client *Client) QueryPartnerIntentionList(request *QueryPartnerIntentionListRequest) (_result *QueryPartnerIntentionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPartnerIntentionListResponse{}
	_body, _err := client.QueryPartnerIntentionListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # QueryPartnerProduceList
//
// @param request - QueryPartnerProduceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPartnerProduceListResponse
func (client *Client) QueryPartnerProduceListWithOptions(request *QueryPartnerProduceListRequest, runtime *dara.RuntimeOptions) (_result *QueryPartnerProduceListResponse, _err error) {
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
		Action:      dara.String("QueryPartnerProduceList"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPartnerProduceListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # QueryPartnerProduceList
//
// @param request - QueryPartnerProduceListRequest
//
// @return QueryPartnerProduceListResponse
func (client *Client) QueryPartnerProduceList(request *QueryPartnerProduceListRequest) (_result *QueryPartnerProduceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPartnerProduceListResponse{}
	_body, _err := client.QueryPartnerProduceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryUserNeedAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryUserNeedAuthResponse
func (client *Client) QueryUserNeedAuthWithOptions(runtime *dara.RuntimeOptions) (_result *QueryUserNeedAuthResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryUserNeedAuth"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryUserNeedAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @return QueryUserNeedAuthResponse
func (client *Client) QueryUserNeedAuth() (_result *QueryUserNeedAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryUserNeedAuthResponse{}
	_body, _err := client.QueryUserNeedAuthWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 服务商玄坛外呼呼叫中心事件回传
//
// @param request - RecordCallCenterEventForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordCallCenterEventForPartnerResponse
func (client *Client) RecordCallCenterEventForPartnerWithOptions(request *RecordCallCenterEventForPartnerRequest, runtime *dara.RuntimeOptions) (_result *RecordCallCenterEventForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CallAction) {
		query["CallAction"] = request.CallAction
	}

	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.ConnId) {
		query["ConnId"] = request.ConnId
	}

	if !dara.IsNil(request.ContactId) {
		query["ContactId"] = request.ContactId
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.RelatedId) {
		query["RelatedId"] = request.RelatedId
	}

	if !dara.IsNil(request.SecretMobile) {
		query["SecretMobile"] = request.SecretMobile
	}

	if !dara.IsNil(request.SkillType) {
		query["SkillType"] = request.SkillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecordCallCenterEventForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecordCallCenterEventForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 服务商玄坛外呼呼叫中心事件回传
//
// @param request - RecordCallCenterEventForPartnerRequest
//
// @return RecordCallCenterEventForPartnerResponse
func (client *Client) RecordCallCenterEventForPartner(request *RecordCallCenterEventForPartnerRequest) (_result *RecordCallCenterEventForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecordCallCenterEventForPartnerResponse{}
	_body, _err := client.RecordCallCenterEventForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # RecordPostBack
//
// @param request - RecordPostBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecordPostBackResponse
func (client *Client) RecordPostBackWithOptions(request *RecordPostBackRequest, runtime *dara.RuntimeOptions) (_result *RecordPostBackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["bizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["bizType"] = request.BizType
	}

	if !dara.IsNil(request.Contactor) {
		query["contactor"] = request.Contactor
	}

	if !dara.IsNil(request.Content) {
		query["content"] = request.Content
	}

	if !dara.IsNil(request.EntityKey) {
		query["entityKey"] = request.EntityKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecordPostBack"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecordPostBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RecordPostBack
//
// @param request - RecordPostBackRequest
//
// @return RecordPostBackResponse
func (client *Client) RecordPostBack(request *RecordPostBackRequest) (_result *RecordPostBackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecordPostBackResponse{}
	_body, _err := client.RecordPostBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RejectSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectSolutionResponse
func (client *Client) RejectSolutionWithOptions(request *RejectSolutionRequest, runtime *dara.RuntimeOptions) (_result *RejectSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.SolutionBizId) {
		query["SolutionBizId"] = request.SolutionBizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectSolution"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RejectSolutionRequest
//
// @return RejectSolutionResponse
func (client *Client) RejectSolution(request *RejectSolutionRequest) (_result *RejectSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectSolutionResponse{}
	_body, _err := client.RejectSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RejectUserSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RejectUserSolutionResponse
func (client *Client) RejectUserSolutionWithOptions(request *RejectUserSolutionRequest, runtime *dara.RuntimeOptions) (_result *RejectUserSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	if !dara.IsNil(request.SolutionBizId) {
		query["SolutionBizId"] = request.SolutionBizId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RejectUserSolution"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RejectUserSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - RejectUserSolutionRequest
//
// @return RejectUserSolutionResponse
func (client *Client) RejectUserSolution(request *RejectUserSolutionRequest) (_result *RejectUserSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RejectUserSolutionResponse{}
	_body, _err := client.RejectUserSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseProduceAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseProduceAuthorizationResponse
func (client *Client) ReleaseProduceAuthorizationWithOptions(request *ReleaseProduceAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *ReleaseProduceAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizedUserId) {
		body["AuthorizedUserId"] = request.AuthorizedUserId
	}

	if !dara.IsNil(request.BizId) {
		body["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		body["BizType"] = request.BizType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseProduceAuthorization"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseProduceAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ReleaseProduceAuthorizationRequest
//
// @return ReleaseProduceAuthorizationResponse
func (client *Client) ReleaseProduceAuthorization(request *ReleaseProduceAuthorizationRequest) (_result *ReleaseProduceAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReleaseProduceAuthorizationResponse{}
	_body, _err := client.ReleaseProduceAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 玄坛双呼外呼发起
//
// @param request - StartBackToBackCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartBackToBackCallResponse
func (client *Client) StartBackToBackCallWithOptions(request *StartBackToBackCallRequest, runtime *dara.RuntimeOptions) (_result *StartBackToBackCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CallCenterNumber) {
		query["CallCenterNumber"] = request.CallCenterNumber
	}

	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.MobileKey) {
		query["MobileKey"] = request.MobileKey
	}

	if !dara.IsNil(request.SkillType) {
		query["SkillType"] = request.SkillType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartBackToBackCall"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartBackToBackCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 玄坛双呼外呼发起
//
// @param request - StartBackToBackCallRequest
//
// @return StartBackToBackCallResponse
func (client *Client) StartBackToBackCall(request *StartBackToBackCallRequest) (_result *StartBackToBackCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartBackToBackCallResponse{}
	_body, _err := client.StartBackToBackCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 合作伙伴提交需求单
//
// @param request - SubmitIntentionForPartnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIntentionForPartnerResponse
func (client *Client) SubmitIntentionForPartnerWithOptions(request *SubmitIntentionForPartnerRequest, runtime *dara.RuntimeOptions) (_result *SubmitIntentionForPartnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Area) {
		query["Area"] = request.Area
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.CommodityType) {
		query["CommodityType"] = request.CommodityType
	}

	if !dara.IsNil(request.ContactName) {
		query["ContactName"] = request.ContactName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExtInfo) {
		query["ExtInfo"] = request.ExtInfo
	}

	if !dara.IsNil(request.Grade) {
		query["Grade"] = request.Grade
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIntentionForPartner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIntentionForPartnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 合作伙伴提交需求单
//
// @param request - SubmitIntentionForPartnerRequest
//
// @return SubmitIntentionForPartnerResponse
func (client *Client) SubmitIntentionForPartner(request *SubmitIntentionForPartnerRequest) (_result *SubmitIntentionForPartnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitIntentionForPartnerResponse{}
	_body, _err := client.SubmitIntentionForPartnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitIntentionNoteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIntentionNoteResponse
func (client *Client) SubmitIntentionNoteWithOptions(request *SubmitIntentionNoteRequest, runtime *dara.RuntimeOptions) (_result *SubmitIntentionNoteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.Note) {
		query["Note"] = request.Note
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIntentionNote"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIntentionNoteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitIntentionNoteRequest
//
// @return SubmitIntentionNoteResponse
func (client *Client) SubmitIntentionNote(request *SubmitIntentionNoteRequest) (_result *SubmitIntentionNoteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitIntentionNoteResponse{}
	_body, _err := client.SubmitIntentionNoteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SubmitSolutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSolutionResponse
func (client *Client) SubmitSolutionWithOptions(request *SubmitSolutionRequest, runtime *dara.RuntimeOptions) (_result *SubmitSolutionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.IntentionBizId) {
		query["IntentionBizId"] = request.IntentionBizId
	}

	if !dara.IsNil(request.OperateType) {
		query["OperateType"] = request.OperateType
	}

	if !dara.IsNil(request.Solution) {
		query["Solution"] = request.Solution
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSolution"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - SubmitSolutionRequest
//
// @return SubmitSolutionResponse
func (client *Client) SubmitSolution(request *SubmitSolutionRequest) (_result *SubmitSolutionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSolutionResponse{}
	_body, _err := client.SubmitSolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 玄坛需求单转派小二
//
// @param request - TransferIntentionOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferIntentionOwnerResponse
func (client *Client) TransferIntentionOwnerWithOptions(request *TransferIntentionOwnerRequest, runtime *dara.RuntimeOptions) (_result *TransferIntentionOwnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferIntentionOwner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferIntentionOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 玄坛需求单转派小二
//
// @param request - TransferIntentionOwnerRequest
//
// @return TransferIntentionOwnerResponse
func (client *Client) TransferIntentionOwner(request *TransferIntentionOwnerRequest) (_result *TransferIntentionOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferIntentionOwnerResponse{}
	_body, _err := client.TransferIntentionOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 玄坛服务单转派小二
//
// @param request - TransferProduceOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferProduceOwnerResponse
func (client *Client) TransferProduceOwnerWithOptions(request *TransferProduceOwnerRequest, runtime *dara.RuntimeOptions) (_result *TransferProduceOwnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["BizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.PersonId) {
		query["PersonId"] = request.PersonId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferProduceOwner"),
		Version:     dara.String("2020-03-06"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferProduceOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 玄坛服务单转派小二
//
// @param request - TransferProduceOwnerRequest
//
// @return TransferProduceOwnerResponse
func (client *Client) TransferProduceOwner(request *TransferProduceOwnerRequest) (_result *TransferProduceOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TransferProduceOwnerResponse{}
	_body, _err := client.TransferProduceOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
