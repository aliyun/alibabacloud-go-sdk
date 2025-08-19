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
		"ap-northeast-1":              dara.String("ice.aliyuncs.com"),
		"ap-northeast-2-pop":          dara.String("ice.aliyuncs.com"),
		"ap-south-1":                  dara.String("ice.aliyuncs.com"),
		"ap-southeast-1":              dara.String("ice.aliyuncs.com"),
		"ap-southeast-2":              dara.String("ice.aliyuncs.com"),
		"ap-southeast-3":              dara.String("ice.aliyuncs.com"),
		"ap-southeast-5":              dara.String("ice.aliyuncs.com"),
		"cn-beijing":                  dara.String("ice.aliyuncs.com"),
		"cn-beijing-finance-1":        dara.String("ice.aliyuncs.com"),
		"cn-beijing-finance-pop":      dara.String("ice.aliyuncs.com"),
		"cn-beijing-gov-1":            dara.String("ice.aliyuncs.com"),
		"cn-beijing-nu16-b01":         dara.String("ice.aliyuncs.com"),
		"cn-chengdu":                  dara.String("ice.aliyuncs.com"),
		"cn-edge-1":                   dara.String("ice.aliyuncs.com"),
		"cn-fujian":                   dara.String("ice.aliyuncs.com"),
		"cn-haidian-cm12-c01":         dara.String("ice.aliyuncs.com"),
		"cn-hangzhou":                 dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-finance":         dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": dara.String("ice.aliyuncs.com"),
		"cn-hangzhou-test-306":        dara.String("ice.aliyuncs.com"),
		"cn-hongkong":                 dara.String("ice.aliyuncs.com"),
		"cn-hongkong-finance-pop":     dara.String("ice.aliyuncs.com"),
		"cn-huhehaote":                dara.String("ice.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       dara.String("ice.aliyuncs.com"),
		"cn-north-2-gov-1":            dara.String("ice.aliyuncs.com"),
		"cn-qingdao":                  dara.String("ice.aliyuncs.com"),
		"cn-qingdao-nebula":           dara.String("ice.aliyuncs.com"),
		"cn-shanghai-et15-b01":        dara.String("ice.aliyuncs.com"),
		"cn-shanghai-et2-b01":         dara.String("ice.aliyuncs.com"),
		"cn-shanghai-finance-1":       dara.String("ice.aliyuncs.com"),
		"cn-shanghai-inner":           dara.String("ice.aliyuncs.com"),
		"cn-shanghai-internal-test-1": dara.String("ice.aliyuncs.com"),
		"cn-shenzhen":                 dara.String("ice.aliyuncs.com"),
		"cn-shenzhen-finance-1":       dara.String("ice.aliyuncs.com"),
		"cn-shenzhen-inner":           dara.String("ice.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         dara.String("ice.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        dara.String("ice.aliyuncs.com"),
		"cn-wuhan":                    dara.String("ice.aliyuncs.com"),
		"cn-wulanchabu":               dara.String("ice.aliyuncs.com"),
		"cn-yushanfang":               dara.String("ice.aliyuncs.com"),
		"cn-zhangbei":                 dara.String("ice.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        dara.String("ice.aliyuncs.com"),
		"cn-zhangjiakou":              dara.String("ice.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     dara.String("ice.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       dara.String("ice.aliyuncs.com"),
		"eu-central-1":                dara.String("ice.aliyuncs.com"),
		"eu-west-1":                   dara.String("ice.aliyuncs.com"),
		"eu-west-1-oxs":               dara.String("ice.aliyuncs.com"),
		"me-east-1":                   dara.String("ice.aliyuncs.com"),
		"rus-west-1-pop":              dara.String("ice.aliyuncs.com"),
		"us-east-1":                   dara.String("ice.aliyuncs.com"),
		"us-west-1":                   dara.String("ice.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds an ad insertion configuration.
//
// @param request - AddAdInsertionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAdInsertionResponse
func (client *Client) AddAdInsertionWithOptions(request *AddAdInsertionRequest, runtime *dara.RuntimeOptions) (_result *AddAdInsertionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdMarkerPassthrough) {
		body["AdMarkerPassthrough"] = request.AdMarkerPassthrough
	}

	if !dara.IsNil(request.AdsUrl) {
		body["AdsUrl"] = request.AdsUrl
	}

	if !dara.IsNil(request.CdnAdSegmentUrlPrefix) {
		body["CdnAdSegmentUrlPrefix"] = request.CdnAdSegmentUrlPrefix
	}

	if !dara.IsNil(request.CdnContentSegmentUrlPrefix) {
		body["CdnContentSegmentUrlPrefix"] = request.CdnContentSegmentUrlPrefix
	}

	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ConfigAliases) {
		body["ConfigAliases"] = request.ConfigAliases
	}

	if !dara.IsNil(request.ContentUrlPrefix) {
		body["ContentUrlPrefix"] = request.ContentUrlPrefix
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PersonalizationThreshold) {
		body["PersonalizationThreshold"] = request.PersonalizationThreshold
	}

	if !dara.IsNil(request.SlateAdUrl) {
		body["SlateAdUrl"] = request.SlateAdUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAdInsertion"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAdInsertionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an ad insertion configuration.
//
// @param request - AddAdInsertionRequest
//
// @return AddAdInsertionResponse
func (client *Client) AddAdInsertion(request *AddAdInsertionRequest) (_result *AddAdInsertionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAdInsertionResponse{}
	_body, _err := client.AddAdInsertionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a category.
//
// Description:
//
// You can create at most three levels of categories. Each category level can contain a maximum of 100 subcategories.
//
// @param request - AddCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCategoryResponse
func (client *Client) AddCategoryWithOptions(request *AddCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateName) {
		query["CateName"] = request.CateName
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddCategory"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a category.
//
// Description:
//
// You can create at most three levels of categories. Each category level can contain a maximum of 100 subcategories.
//
// @param request - AddCategoryRequest
//
// @return AddCategoryResponse
func (client *Client) AddCategory(request *AddCategoryRequest) (_result *AddCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddCategoryResponse{}
	_body, _err := client.AddCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds one or more materials to an online editing project.
//
// @param request - AddEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddEditingProjectMaterialsResponse
func (client *Client) AddEditingProjectMaterialsWithOptions(request *AddEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *AddEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaterialMaps) {
		query["MaterialMaps"] = request.MaterialMaps
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddEditingProjectMaterials"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddEditingProjectMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds one or more materials to an online editing project.
//
// @param request - AddEditingProjectMaterialsRequest
//
// @return AddEditingProjectMaterialsResponse
func (client *Client) AddEditingProjectMaterials(request *AddEditingProjectMaterialsRequest) (_result *AddEditingProjectMaterialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddEditingProjectMaterialsResponse{}
	_body, _err := client.AddEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 收藏公共媒资
//
// @param request - AddFavoritePublicMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFavoritePublicMediaResponse
func (client *Client) AddFavoritePublicMediaWithOptions(request *AddFavoritePublicMediaRequest, runtime *dara.RuntimeOptions) (_result *AddFavoritePublicMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFavoritePublicMedia"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFavoritePublicMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 收藏公共媒资
//
// @param request - AddFavoritePublicMediaRequest
//
// @return AddFavoritePublicMediaResponse
func (client *Client) AddFavoritePublicMedia(request *AddFavoritePublicMediaRequest) (_result *AddFavoritePublicMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddFavoritePublicMediaResponse{}
	_body, _err := client.AddFavoritePublicMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a source for a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- A flow can have only one source.
//
// ### [](#)Source type
//
//   - RTMP-PUSH: An input that you can push to the returned URL over the RTMP protocol.
//
//   - RTMP-PULL: An input that the MediaConnect flow pulls from the specified server over the RTMP protocol.
//
//   - SRT-Listener: An input that you can push to the returned URL over the SRT protocol.
//
//   - SRT-Caller: An input that the MediaConnect flow pulls from the specified server over the SRT protocol.
//
//   - Flow: An input that uses the output of another upstream flow. You must specify an upstream flow and its output. The output type of the upstream flow must be SRT-Listener or RTMP-PULL. By default, a dedicated line is used when flows are cascaded. This allows for cross-region distribution among multiple flows.
//
// @param request - AddMediaConnectFlowInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMediaConnectFlowInputResponse
func (client *Client) AddMediaConnectFlowInputWithOptions(request *AddMediaConnectFlowInputRequest, runtime *dara.RuntimeOptions) (_result *AddMediaConnectFlowInputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidrs) {
		query["Cidrs"] = request.Cidrs
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.InputFromUrl) {
		query["InputFromUrl"] = request.InputFromUrl
	}

	if !dara.IsNil(request.InputName) {
		query["InputName"] = request.InputName
	}

	if !dara.IsNil(request.InputProtocol) {
		query["InputProtocol"] = request.InputProtocol
	}

	if !dara.IsNil(request.MaxBitrate) {
		query["MaxBitrate"] = request.MaxBitrate
	}

	if !dara.IsNil(request.PairFlowId) {
		query["PairFlowId"] = request.PairFlowId
	}

	if !dara.IsNil(request.PairOutputName) {
		query["PairOutputName"] = request.PairOutputName
	}

	if !dara.IsNil(request.SrtLatency) {
		query["SrtLatency"] = request.SrtLatency
	}

	if !dara.IsNil(request.SrtPassphrase) {
		query["SrtPassphrase"] = request.SrtPassphrase
	}

	if !dara.IsNil(request.SrtPbkeyLen) {
		query["SrtPbkeyLen"] = request.SrtPbkeyLen
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMediaConnectFlowInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMediaConnectFlowInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a source for a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- A flow can have only one source.
//
// ### [](#)Source type
//
//   - RTMP-PUSH: An input that you can push to the returned URL over the RTMP protocol.
//
//   - RTMP-PULL: An input that the MediaConnect flow pulls from the specified server over the RTMP protocol.
//
//   - SRT-Listener: An input that you can push to the returned URL over the SRT protocol.
//
//   - SRT-Caller: An input that the MediaConnect flow pulls from the specified server over the SRT protocol.
//
//   - Flow: An input that uses the output of another upstream flow. You must specify an upstream flow and its output. The output type of the upstream flow must be SRT-Listener or RTMP-PULL. By default, a dedicated line is used when flows are cascaded. This allows for cross-region distribution among multiple flows.
//
// @param request - AddMediaConnectFlowInputRequest
//
// @return AddMediaConnectFlowInputResponse
func (client *Client) AddMediaConnectFlowInput(request *AddMediaConnectFlowInputRequest) (_result *AddMediaConnectFlowInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMediaConnectFlowInputResponse{}
	_body, _err := client.AddMediaConnectFlowInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an output for a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- A flow can have a maximum of four outputs.
//
//		- The output names in the same flow cannot be duplicated.
//
//		- You can set an upper limit on the number of concurrent viewers for each output. If this limit is exceeded, any new playback requests will fail. Each output supports up to five streams.
//
// ### [](#)Output types
//
//   - RTMP-PUSH: An output that the MediaConnect flow pushes to the server you specified over the RTMP protocol.
//
//   - RTMP-PULL: An output that you can pull using the returned streaming URL over the RTMP protocol.
//
//   - SRT-Caller: An output that the MediaConnect flow pushes to the server you specified over the SRT protocol.
//
//   - SRT-Listener: An output that you can pull using the returned streaming URL over the SRT protocol.
//
//   - Flow: An output that is pushed to the source URL of another MediaConnect flow. The source type of the destination flow must be SRT-Listener or RTMP-PUSH. By default, a dedicated line is used when flows are cascaded. This allows for cross-region distribution among multiple flows.
//
// @param request - AddMediaConnectFlowOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMediaConnectFlowOutputResponse
func (client *Client) AddMediaConnectFlowOutputWithOptions(request *AddMediaConnectFlowOutputRequest, runtime *dara.RuntimeOptions) (_result *AddMediaConnectFlowOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidrs) {
		query["Cidrs"] = request.Cidrs
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	if !dara.IsNil(request.OutputProtocol) {
		query["OutputProtocol"] = request.OutputProtocol
	}

	if !dara.IsNil(request.OutputToUrl) {
		query["OutputToUrl"] = request.OutputToUrl
	}

	if !dara.IsNil(request.PairFlowId) {
		query["PairFlowId"] = request.PairFlowId
	}

	if !dara.IsNil(request.PairInputName) {
		query["PairInputName"] = request.PairInputName
	}

	if !dara.IsNil(request.PlayerLimit) {
		query["PlayerLimit"] = request.PlayerLimit
	}

	if !dara.IsNil(request.SrtLatency) {
		query["SrtLatency"] = request.SrtLatency
	}

	if !dara.IsNil(request.SrtPassphrase) {
		query["SrtPassphrase"] = request.SrtPassphrase
	}

	if !dara.IsNil(request.SrtPbkeyLen) {
		query["SrtPbkeyLen"] = request.SrtPbkeyLen
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMediaConnectFlowOutput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMediaConnectFlowOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an output for a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- A flow can have a maximum of four outputs.
//
//		- The output names in the same flow cannot be duplicated.
//
//		- You can set an upper limit on the number of concurrent viewers for each output. If this limit is exceeded, any new playback requests will fail. Each output supports up to five streams.
//
// ### [](#)Output types
//
//   - RTMP-PUSH: An output that the MediaConnect flow pushes to the server you specified over the RTMP protocol.
//
//   - RTMP-PULL: An output that you can pull using the returned streaming URL over the RTMP protocol.
//
//   - SRT-Caller: An output that the MediaConnect flow pushes to the server you specified over the SRT protocol.
//
//   - SRT-Listener: An output that you can pull using the returned streaming URL over the SRT protocol.
//
//   - Flow: An output that is pushed to the source URL of another MediaConnect flow. The source type of the destination flow must be SRT-Listener or RTMP-PUSH. By default, a dedicated line is used when flows are cascaded. This allows for cross-region distribution among multiple flows.
//
// @param request - AddMediaConnectFlowOutputRequest
//
// @return AddMediaConnectFlowOutputResponse
func (client *Client) AddMediaConnectFlowOutput(request *AddMediaConnectFlowOutputRequest) (_result *AddMediaConnectFlowOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMediaConnectFlowOutputResponse{}
	_body, _err := client.AddMediaConnectFlowOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds marks for a media asset.
//
// @param request - AddMediaMarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMediaMarksResponse
func (client *Client) AddMediaMarksWithOptions(request *AddMediaMarksRequest, runtime *dara.RuntimeOptions) (_result *AddMediaMarksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaMarks) {
		query["MediaMarks"] = request.MediaMarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMediaMarks"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMediaMarksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds marks for a media asset.
//
// @param request - AddMediaMarksRequest
//
// @return AddMediaMarksResponse
func (client *Client) AddMediaMarks(request *AddMediaMarksRequest) (_result *AddMediaMarksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMediaMarksResponse{}
	_body, _err := client.AddMediaMarksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// Description:
//
//	  For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//		- For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
//		- After an advanced template is created, it enters the Processing state. In this case, the template is unavailable. The template can be used only when it is in the Available state. The time required for template processing varies based on the size of the template file. Generally, it ranges from 10 seconds to 5 minutes.
//
// @param request - AddTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTemplateResponse
func (client *Client) AddTemplateWithOptions(request *AddTemplateRequest, runtime *dara.RuntimeOptions) (_result *AddTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoverUrl) {
		query["CoverUrl"] = request.CoverUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PreviewMedia) {
		query["PreviewMedia"] = request.PreviewMedia
	}

	if !dara.IsNil(request.RelatedMediaids) {
		query["RelatedMediaids"] = request.RelatedMediaids
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// Description:
//
//	  For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//		- For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
//		- After an advanced template is created, it enters the Processing state. In this case, the template is unavailable. The template can be used only when it is in the Available state. The time required for template processing varies based on the size of the template file. Generally, it ranges from 10 seconds to 5 minutes.
//
// @param request - AddTemplateRequest
//
// @return AddTemplateResponse
func (client *Client) AddTemplate(request *AddTemplateRequest) (_result *AddTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTemplateResponse{}
	_body, _err := client.AddTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies search index information including index status and configurations.
//
// @param request - AlterSearchIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AlterSearchIndexResponse
func (client *Client) AlterSearchIndexWithOptions(request *AlterSearchIndexRequest, runtime *dara.RuntimeOptions) (_result *AlterSearchIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexConfig) {
		query["IndexConfig"] = request.IndexConfig
	}

	if !dara.IsNil(request.IndexStatus) {
		query["IndexStatus"] = request.IndexStatus
	}

	if !dara.IsNil(request.IndexType) {
		query["IndexType"] = request.IndexType
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AlterSearchIndex"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AlterSearchIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies search index information including index status and configurations.
//
// @param request - AlterSearchIndexRequest
//
// @return AlterSearchIndexResponse
func (client *Client) AlterSearchIndex(request *AlterSearchIndexRequest) (_result *AlterSearchIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AlterSearchIndexResponse{}
	_body, _err := client.AlterSearchIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ingests multiple assets for VOD packaging.
//
// @param tmpReq - BatchCreateVodPackagingAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateVodPackagingAssetResponse
func (client *Client) BatchCreateVodPackagingAssetWithOptions(tmpReq *BatchCreateVodPackagingAssetRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateVodPackagingAssetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &BatchCreateVodPackagingAssetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Assets) {
		request.AssetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assets, dara.String("Assets"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetsShrink) {
		query["Assets"] = request.AssetsShrink
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateVodPackagingAsset"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateVodPackagingAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ingests multiple assets for VOD packaging.
//
// @param request - BatchCreateVodPackagingAssetRequest
//
// @return BatchCreateVodPackagingAssetResponse
func (client *Client) BatchCreateVodPackagingAsset(request *BatchCreateVodPackagingAssetRequest) (_result *BatchCreateVodPackagingAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateVodPackagingAssetResponse{}
	_body, _err := client.BatchCreateVodPackagingAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about multiple media assets at a time based on media asset IDs.
//
// @param request - BatchGetMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetMediaInfosResponse
func (client *Client) BatchGetMediaInfosWithOptions(request *BatchGetMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *BatchGetMediaInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdditionType) {
		query["AdditionType"] = request.AdditionType
	}

	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetMediaInfos"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetMediaInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about multiple media assets at a time based on media asset IDs.
//
// @param request - BatchGetMediaInfosRequest
//
// @return BatchGetMediaInfosResponse
func (client *Client) BatchGetMediaInfos(request *BatchGetMediaInfosRequest) (_result *BatchGetMediaInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchGetMediaInfosResponse{}
	_body, _err := client.BatchGetMediaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a media fingerprint analysis job.
//
// Description:
//
//	  You can cancel a media fingerprint analysis job only if the job is in the Queuing state.
//
//		- We recommend that you call the **UpdatePipeline*	- operation to set the status of the ApsaraVideo Media Processing (MPS) queue to Paused before you cancel a job. This suspends job scheduling in the MPS queue. After the job is canceled, you must set the status of the MPS queue back to Active so that the other jobs in the MPS queue can be scheduled.
//
// @param request - CancelDNAJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelDNAJobResponse
func (client *Client) CancelDNAJobWithOptions(request *CancelDNAJobRequest, runtime *dara.RuntimeOptions) (_result *CancelDNAJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelDNAJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelDNAJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a media fingerprint analysis job.
//
// Description:
//
//	  You can cancel a media fingerprint analysis job only if the job is in the Queuing state.
//
//		- We recommend that you call the **UpdatePipeline*	- operation to set the status of the ApsaraVideo Media Processing (MPS) queue to Paused before you cancel a job. This suspends job scheduling in the MPS queue. After the job is canceled, you must set the status of the MPS queue back to Active so that the other jobs in the MPS queue can be scheduled.
//
// @param request - CancelDNAJobRequest
//
// @return CancelDNAJobResponse
func (client *Client) CancelDNAJob(request *CancelDNAJobRequest) (_result *CancelDNAJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelDNAJobResponse{}
	_body, _err := client.CancelDNAJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消收藏公共媒资
//
// @param request - CancelFavoritePublicMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelFavoritePublicMediaResponse
func (client *Client) CancelFavoritePublicMediaWithOptions(request *CancelFavoritePublicMediaRequest, runtime *dara.RuntimeOptions) (_result *CancelFavoritePublicMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelFavoritePublicMedia"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelFavoritePublicMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消收藏公共媒资
//
// @param request - CancelFavoritePublicMediaRequest
//
// @return CancelFavoritePublicMediaResponse
func (client *Client) CancelFavoritePublicMedia(request *CancelFavoritePublicMediaRequest) (_result *CancelFavoritePublicMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelFavoritePublicMediaResponse{}
	_body, _err := client.CancelFavoritePublicMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消智能生产任务
//
// @param request - CancelIProductionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelIProductionJobResponse
func (client *Client) CancelIProductionJobWithOptions(request *CancelIProductionJobRequest, runtime *dara.RuntimeOptions) (_result *CancelIProductionJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelIProductionJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelIProductionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消智能生产任务
//
// @param request - CancelIProductionJobRequest
//
// @return CancelIProductionJobResponse
func (client *Client) CancelIProductionJob(request *CancelIProductionJobRequest) (_result *CancelIProductionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelIProductionJobResponse{}
	_body, _err := client.CancelIProductionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 清除声纹ID对应的声纹
//
// @param request - ClearAIAgentVoiceprintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearAIAgentVoiceprintResponse
func (client *Client) ClearAIAgentVoiceprintWithOptions(request *ClearAIAgentVoiceprintRequest, runtime *dara.RuntimeOptions) (_result *ClearAIAgentVoiceprintResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VoiceprintId) {
		query["VoiceprintId"] = request.VoiceprintId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClearAIAgentVoiceprint"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClearAIAgentVoiceprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 清除声纹ID对应的声纹
//
// @param request - ClearAIAgentVoiceprintRequest
//
// @return ClearAIAgentVoiceprintResponse
func (client *Client) ClearAIAgentVoiceprint(request *ClearAIAgentVoiceprintRequest) (_result *ClearAIAgentVoiceprintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClearAIAgentVoiceprintResponse{}
	_body, _err := client.ClearAIAgentVoiceprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关闭MediaConnect实例的双流灾备功能
//
// @param request - CloseMediaConnectFlowFailoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseMediaConnectFlowFailoverResponse
func (client *Client) CloseMediaConnectFlowFailoverWithOptions(request *CloseMediaConnectFlowFailoverRequest, runtime *dara.RuntimeOptions) (_result *CloseMediaConnectFlowFailoverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseMediaConnectFlowFailover"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseMediaConnectFlowFailoverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关闭MediaConnect实例的双流灾备功能
//
// @param request - CloseMediaConnectFlowFailoverRequest
//
// @return CloseMediaConnectFlowFailoverResponse
func (client *Client) CloseMediaConnectFlowFailover(request *CloseMediaConnectFlowFailoverRequest) (_result *CloseMediaConnectFlowFailoverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseMediaConnectFlowFailoverResponse{}
	_body, _err := client.CloseMediaConnectFlowFailoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits manual review results for media assets.
//
// @param request - CreateAuditRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuditResponse
func (client *Client) CreateAuditWithOptions(request *CreateAuditRequest, runtime *dara.RuntimeOptions) (_result *CreateAuditResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditContent) {
		query["AuditContent"] = request.AuditContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAudit"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits manual review results for media assets.
//
// @param request - CreateAuditRequest
//
// @return CreateAuditResponse
func (client *Client) CreateAudit(request *CreateAuditRequest) (_result *CreateAuditResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAuditResponse{}
	_body, _err := client.CreateAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a digital human training job. You can configure the basic information of the digital human and the materials required for the training. Note: This operation is used to initialize the training job. It does not submit the training job. To submit the training job, call the SubmitAvatarTrainingJob operation.
//
// @param request - CreateAvatarTrainingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAvatarTrainingJobResponse
func (client *Client) CreateAvatarTrainingJobWithOptions(request *CreateAvatarTrainingJobRequest, runtime *dara.RuntimeOptions) (_result *CreateAvatarTrainingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarDescription) {
		query["AvatarDescription"] = request.AvatarDescription
	}

	if !dara.IsNil(request.AvatarName) {
		query["AvatarName"] = request.AvatarName
	}

	if !dara.IsNil(request.AvatarType) {
		query["AvatarType"] = request.AvatarType
	}

	if !dara.IsNil(request.Portrait) {
		query["Portrait"] = request.Portrait
	}

	if !dara.IsNil(request.Thumbnail) {
		query["Thumbnail"] = request.Thumbnail
	}

	if !dara.IsNil(request.Transparent) {
		query["Transparent"] = request.Transparent
	}

	if !dara.IsNil(request.Video) {
		query["Video"] = request.Video
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAvatarTrainingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAvatarTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a digital human training job. You can configure the basic information of the digital human and the materials required for the training. Note: This operation is used to initialize the training job. It does not submit the training job. To submit the training job, call the SubmitAvatarTrainingJob operation.
//
// @param request - CreateAvatarTrainingJobRequest
//
// @return CreateAvatarTrainingJobResponse
func (client *Client) CreateAvatarTrainingJob(request *CreateAvatarTrainingJobRequest) (_result *CreateAvatarTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAvatarTrainingJobResponse{}
	_body, _err := client.CreateAvatarTrainingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a channel in MediaWeaver.
//
// @param request - CreateChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChannelResponse
func (client *Client) CreateChannelWithOptions(request *CreateChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPolicy) {
		query["AccessPolicy"] = request.AccessPolicy
	}

	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ChannelTier) {
		query["ChannelTier"] = request.ChannelTier
	}

	if !dara.IsNil(request.FillerSourceLocationName) {
		query["FillerSourceLocationName"] = request.FillerSourceLocationName
	}

	if !dara.IsNil(request.FillerSourceName) {
		query["FillerSourceName"] = request.FillerSourceName
	}

	if !dara.IsNil(request.OutPutConfigList) {
		query["OutPutConfigList"] = request.OutPutConfigList
	}

	if !dara.IsNil(request.PlaybackMode) {
		query["PlaybackMode"] = request.PlaybackMode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a channel in MediaWeaver.
//
// @param request - CreateChannelRequest
//
// @return CreateChannelResponse
func (client *Client) CreateChannel(request *CreateChannelRequest) (_result *CreateChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChannelResponse{}
	_body, _err := client.CreateChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom template.
//
// @param request - CreateCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomTemplateResponse
func (client *Client) CreateCustomTemplateWithOptions(request *CreateCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Subtype) {
		query["Subtype"] = request.Subtype
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom template.
//
// @param request - CreateCustomTemplateRequest
//
// @return CreateCustomTemplateResponse
func (client *Client) CreateCustomTemplate(request *CreateCustomTemplateRequest) (_result *CreateCustomTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomTemplateResponse{}
	_body, _err := client.CreateCustomTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a human voice cloning job. You can configure the basic information of the human voice cloning job.
//
// @param request - CreateCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomizedVoiceJobResponse
func (client *Client) CreateCustomizedVoiceJobWithOptions(request *CreateCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Gender) {
		query["Gender"] = request.Gender
	}

	if !dara.IsNil(request.Scenario) {
		query["Scenario"] = request.Scenario
	}

	if !dara.IsNil(request.VoiceDesc) {
		query["VoiceDesc"] = request.VoiceDesc
	}

	if !dara.IsNil(request.VoiceId) {
		query["VoiceId"] = request.VoiceId
	}

	if !dara.IsNil(request.VoiceName) {
		query["VoiceName"] = request.VoiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a human voice cloning job. You can configure the basic information of the human voice cloning job.
//
// @param request - CreateCustomizedVoiceJobRequest
//
// @return CreateCustomizedVoiceJobResponse
func (client *Client) CreateCustomizedVoiceJob(request *CreateCustomizedVoiceJobRequest) (_result *CreateCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCustomizedVoiceJobResponse{}
	_body, _err := client.CreateCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates media fingerprint libraries.
//
// Description:
//
//	You can create up to five media fingerprint libraries within an account. To increase the quota, submit a ticket. You can call the DeleteDNADB operation to delete the fingerprint libraries that you no longer need.
//
// @param request - CreateDNADBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDNADBResponse
func (client *Client) CreateDNADBWithOptions(request *CreateDNADBRequest, runtime *dara.RuntimeOptions) (_result *CreateDNADBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDNADB"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDNADBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates media fingerprint libraries.
//
// Description:
//
//	You can create up to five media fingerprint libraries within an account. To increase the quota, submit a ticket. You can call the DeleteDNADB operation to delete the fingerprint libraries that you no longer need.
//
// @param request - CreateDNADBRequest
//
// @return CreateDNADBResponse
func (client *Client) CreateDNADB(request *CreateDNADBRequest) (_result *CreateDNADBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDNADBResponse{}
	_body, _err := client.CreateDNADBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an online editing project. You can specify configurations such as the title, description, timeline, and thumbnail for the project.
//
// @param request - CreateEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEditingProjectResponse
func (client *Client) CreateEditingProjectWithOptions(request *CreateEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessConfig) {
		query["BusinessConfig"] = request.BusinessConfig
	}

	if !dara.IsNil(request.ClipsParam) {
		query["ClipsParam"] = request.ClipsParam
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.MaterialMaps) {
		query["MaterialMaps"] = request.MaterialMaps
	}

	if !dara.IsNil(request.ProjectType) {
		query["ProjectType"] = request.ProjectType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Timeline) {
		body["Timeline"] = request.Timeline
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEditingProject"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEditingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an online editing project. You can specify configurations such as the title, description, timeline, and thumbnail for the project.
//
// @param request - CreateEditingProjectRequest
//
// @return CreateEditingProjectResponse
func (client *Client) CreateEditingProject(request *CreateEditingProjectRequest) (_result *CreateEditingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateEditingProjectResponse{}
	_body, _err := client.CreateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建热词库
//
// @param tmpReq - CreateHotwordLibraryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateHotwordLibraryResponse
func (client *Client) CreateHotwordLibraryWithOptions(tmpReq *CreateHotwordLibraryRequest, runtime *dara.RuntimeOptions) (_result *CreateHotwordLibraryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateHotwordLibraryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Hotwords) {
		request.HotwordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hotwords, dara.String("Hotwords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HotwordsShrink) {
		query["Hotwords"] = request.HotwordsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.UsageScenario) {
		query["UsageScenario"] = request.UsageScenario
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateHotwordLibrary"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateHotwordLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热词库
//
// @param request - CreateHotwordLibraryRequest
//
// @return CreateHotwordLibraryResponse
func (client *Client) CreateHotwordLibrary(request *CreateHotwordLibraryRequest) (_result *CreateHotwordLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateHotwordLibraryResponse{}
	_body, _err := client.CreateHotwordLibraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// After you call this operation to create a live package channel, the system will automatically generate the ingest endpoint URL, and username and password required for authentication.
//
// ### [](#)Precautions
//
//   - Channel group names and channel names can contain only letters, digits, underscores (_), and hyphens (-).
//
//   - Only `HLS` is supported.
//
//   - The segment duration must be from 1 to 30 seconds.
//
//   - The number of M3U8 segments must be from 2 to 100.
//
// If the request succeeds, the system will return the details of the newly created channel, including the channel name, creation time, modification time, and ingest endpoint details.
//
// @param request - CreateLivePackageChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivePackageChannelResponse
func (client *Client) CreateLivePackageChannelWithOptions(request *CreateLivePackageChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateLivePackageChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		body["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.SegmentCount) {
		body["SegmentCount"] = request.SegmentCount
	}

	if !dara.IsNil(request.SegmentDuration) {
		body["SegmentDuration"] = request.SegmentDuration
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivePackageChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivePackageChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// After you call this operation to create a live package channel, the system will automatically generate the ingest endpoint URL, and username and password required for authentication.
//
// ### [](#)Precautions
//
//   - Channel group names and channel names can contain only letters, digits, underscores (_), and hyphens (-).
//
//   - Only `HLS` is supported.
//
//   - The segment duration must be from 1 to 30 seconds.
//
//   - The number of M3U8 segments must be from 2 to 100.
//
// If the request succeeds, the system will return the details of the newly created channel, including the channel name, creation time, modification time, and ingest endpoint details.
//
// @param request - CreateLivePackageChannelRequest
//
// @return CreateLivePackageChannelResponse
func (client *Client) CreateLivePackageChannel(request *CreateLivePackageChannelRequest) (_result *CreateLivePackageChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLivePackageChannelResponse{}
	_body, _err := client.CreateLivePackageChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a live package channel group with a custom name and description.
//
// Description:
//
// After you create a channel group, the assigned origin domain is returned.
//
// @param request - CreateLivePackageChannelGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivePackageChannelGroupResponse
func (client *Client) CreateLivePackageChannelGroupWithOptions(request *CreateLivePackageChannelGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateLivePackageChannelGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivePackageChannelGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivePackageChannelGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a live package channel group with a custom name and description.
//
// Description:
//
// After you create a channel group, the assigned origin domain is returned.
//
// @param request - CreateLivePackageChannelGroupRequest
//
// @return CreateLivePackageChannelGroupResponse
func (client *Client) CreateLivePackageChannelGroup(request *CreateLivePackageChannelGroupRequest) (_result *CreateLivePackageChannelGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLivePackageChannelGroupResponse{}
	_body, _err := client.CreateLivePackageChannelGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an origin endpoint for a live package channel to deliver live streams in HLS format.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation is mainly used to configure origin settings, security policies including the IP address blacklist and whitelist and authorization code, and time shifting settings for channels. Before you create an origin endpoint, you must create a live package channel group and channel. After you create the endpoint, the endpoint URL and other configuration details are returned.
//
// @param tmpReq - CreateLivePackageOriginEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLivePackageOriginEndpointResponse
func (client *Client) CreateLivePackageOriginEndpointWithOptions(tmpReq *CreateLivePackageOriginEndpointRequest, runtime *dara.RuntimeOptions) (_result *CreateLivePackageOriginEndpointResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateLivePackageOriginEndpointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LivePackagingConfig) {
		request.LivePackagingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LivePackagingConfig, dara.String("LivePackagingConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationCode) {
		body["AuthorizationCode"] = request.AuthorizationCode
	}

	if !dara.IsNil(request.ChannelName) {
		body["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndpointName) {
		body["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IpBlacklist) {
		body["IpBlacklist"] = request.IpBlacklist
	}

	if !dara.IsNil(request.IpWhitelist) {
		body["IpWhitelist"] = request.IpWhitelist
	}

	if !dara.IsNil(request.LivePackagingConfigShrink) {
		body["LivePackagingConfig"] = request.LivePackagingConfigShrink
	}

	if !dara.IsNil(request.ManifestName) {
		body["ManifestName"] = request.ManifestName
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TimeshiftVision) {
		body["TimeshiftVision"] = request.TimeshiftVision
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLivePackageOriginEndpoint"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLivePackageOriginEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an origin endpoint for a live package channel to deliver live streams in HLS format.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation is mainly used to configure origin settings, security policies including the IP address blacklist and whitelist and authorization code, and time shifting settings for channels. Before you create an origin endpoint, you must create a live package channel group and channel. After you create the endpoint, the endpoint URL and other configuration details are returned.
//
// @param request - CreateLivePackageOriginEndpointRequest
//
// @return CreateLivePackageOriginEndpointResponse
func (client *Client) CreateLivePackageOriginEndpoint(request *CreateLivePackageOriginEndpointRequest) (_result *CreateLivePackageOriginEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLivePackageOriginEndpointResponse{}
	_body, _err := client.CreateLivePackageOriginEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a live stream recording template to submit live stream recording jobs.
//
// Description:
//
// You must specify a recording template for live stream recording. You can configure information such as the format and duration of a recording in a recording template. The recording format can be M3U8, MP4, or FLV.
//
// @param tmpReq - CreateLiveRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveRecordTemplateResponse
func (client *Client) CreateLiveRecordTemplateWithOptions(tmpReq *CreateLiveRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveRecordTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateLiveRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordFormat) {
		request.RecordFormatShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordFormat, dara.String("RecordFormat"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecordFormatShrink) {
		body["RecordFormat"] = request.RecordFormatShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveRecordTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a live stream recording template to submit live stream recording jobs.
//
// Description:
//
// You must specify a recording template for live stream recording. You can configure information such as the format and duration of a recording in a recording template. The recording format can be M3U8, MP4, or FLV.
//
// @param request - CreateLiveRecordTemplateRequest
//
// @return CreateLiveRecordTemplateResponse
func (client *Client) CreateLiveRecordTemplate(request *CreateLiveRecordTemplateRequest) (_result *CreateLiveRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLiveRecordTemplateResponse{}
	_body, _err := client.CreateLiveRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a live stream snapshot template to facilitate the creation of snapshot jobs.
//
// @param request - CreateLiveSnapshotTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveSnapshotTemplateResponse
func (client *Client) CreateLiveSnapshotTemplateWithOptions(request *CreateLiveSnapshotTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveSnapshotTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OverwriteFormat) {
		body["OverwriteFormat"] = request.OverwriteFormat
	}

	if !dara.IsNil(request.SequenceFormat) {
		body["SequenceFormat"] = request.SequenceFormat
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TimeInterval) {
		body["TimeInterval"] = request.TimeInterval
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveSnapshotTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveSnapshotTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a live stream snapshot template to facilitate the creation of snapshot jobs.
//
// @param request - CreateLiveSnapshotTemplateRequest
//
// @return CreateLiveSnapshotTemplateResponse
func (client *Client) CreateLiveSnapshotTemplate(request *CreateLiveSnapshotTemplateRequest) (_result *CreateLiveSnapshotTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLiveSnapshotTemplateResponse{}
	_body, _err := client.CreateLiveSnapshotTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a live stream transcoding template to submit live stream transcoding jobs.
//
// @param tmpReq - CreateLiveTranscodeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLiveTranscodeTemplateResponse
func (client *Client) CreateLiveTranscodeTemplateWithOptions(tmpReq *CreateLiveTranscodeTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateLiveTranscodeTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateLiveTranscodeTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLiveTranscodeTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLiveTranscodeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a live stream transcoding template to submit live stream transcoding jobs.
//
// @param request - CreateLiveTranscodeTemplateRequest
//
// @return CreateLiveTranscodeTemplateResponse
func (client *Client) CreateLiveTranscodeTemplate(request *CreateLiveTranscodeTemplateRequest) (_result *CreateLiveTranscodeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLiveTranscodeTemplateResponse{}
	_body, _err := client.CreateLiveTranscodeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a MediaConnect flow.
//
// Description:
//
//	  The flow names cannot be duplicated in the same region.
//
//		- Take note of the returned flow ID. You may reference it in other API operations.
//
// @param request - CreateMediaConnectFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMediaConnectFlowResponse
func (client *Client) CreateMediaConnectFlowWithOptions(request *CreateMediaConnectFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateMediaConnectFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowName) {
		query["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.FlowRegion) {
		query["FlowRegion"] = request.FlowRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMediaConnectFlow"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMediaConnectFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a MediaConnect flow.
//
// Description:
//
//	  The flow names cannot be duplicated in the same region.
//
//		- Take note of the returned flow ID. You may reference it in other API operations.
//
// @param request - CreateMediaConnectFlowRequest
//
// @return CreateMediaConnectFlowResponse
func (client *Client) CreateMediaConnectFlow(request *CreateMediaConnectFlowRequest) (_result *CreateMediaConnectFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMediaConnectFlowResponse{}
	_body, _err := client.CreateMediaConnectFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a MediaLive channel.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - CreateMediaLiveChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMediaLiveChannelResponse
func (client *Client) CreateMediaLiveChannelWithOptions(tmpReq *CreateMediaLiveChannelRequest, runtime *dara.RuntimeOptions) (_result *CreateMediaLiveChannelResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateMediaLiveChannelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AudioSettings) {
		request.AudioSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AudioSettings, dara.String("AudioSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InputAttachments) {
		request.InputAttachmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputAttachments, dara.String("InputAttachments"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutputGroups) {
		request.OutputGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputGroups, dara.String("OutputGroups"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoSettings) {
		request.VideoSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoSettings, dara.String("VideoSettings"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AudioSettingsShrink) {
		body["AudioSettings"] = request.AudioSettingsShrink
	}

	if !dara.IsNil(request.InputAttachmentsShrink) {
		body["InputAttachments"] = request.InputAttachmentsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputGroupsShrink) {
		body["OutputGroups"] = request.OutputGroupsShrink
	}

	if !dara.IsNil(request.VideoSettingsShrink) {
		body["VideoSettings"] = request.VideoSettingsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMediaLiveChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMediaLiveChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a MediaLive channel.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - CreateMediaLiveChannelRequest
//
// @return CreateMediaLiveChannelResponse
func (client *Client) CreateMediaLiveChannel(request *CreateMediaLiveChannelRequest) (_result *CreateMediaLiveChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMediaLiveChannelResponse{}
	_body, _err := client.CreateMediaLiveChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a MediaLive input.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - CreateMediaLiveInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMediaLiveInputResponse
func (client *Client) CreateMediaLiveInputWithOptions(tmpReq *CreateMediaLiveInputRequest, runtime *dara.RuntimeOptions) (_result *CreateMediaLiveInputResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateMediaLiveInputShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputSettings) {
		request.InputSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputSettings, dara.String("InputSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityGroupIds) {
		request.SecurityGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityGroupIds, dara.String("SecurityGroupIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InputSettingsShrink) {
		body["InputSettings"] = request.InputSettingsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SecurityGroupIdsShrink) {
		body["SecurityGroupIds"] = request.SecurityGroupIdsShrink
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMediaLiveInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMediaLiveInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a MediaLive input.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - CreateMediaLiveInputRequest
//
// @return CreateMediaLiveInputResponse
func (client *Client) CreateMediaLiveInput(request *CreateMediaLiveInputRequest) (_result *CreateMediaLiveInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMediaLiveInputResponse{}
	_body, _err := client.CreateMediaLiveInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a security group in MediaLive.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - CreateMediaLiveInputSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMediaLiveInputSecurityGroupResponse
func (client *Client) CreateMediaLiveInputSecurityGroupWithOptions(tmpReq *CreateMediaLiveInputSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateMediaLiveInputSecurityGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateMediaLiveInputSecurityGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WhitelistRules) {
		request.WhitelistRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WhitelistRules, dara.String("WhitelistRules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.WhitelistRulesShrink) {
		body["WhitelistRules"] = request.WhitelistRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMediaLiveInputSecurityGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a security group in MediaLive.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - CreateMediaLiveInputSecurityGroupRequest
//
// @return CreateMediaLiveInputSecurityGroupResponse
func (client *Client) CreateMediaLiveInputSecurityGroup(request *CreateMediaLiveInputSecurityGroupRequest) (_result *CreateMediaLiveInputSecurityGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.CreateMediaLiveInputSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - CreatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePipelineResponse
func (client *Client) CreatePipelineWithOptions(request *CreatePipelineRequest, runtime *dara.RuntimeOptions) (_result *CreatePipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePipeline"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - CreatePipelineRequest
//
// @return CreatePipelineResponse
func (client *Client) CreatePipeline(request *CreatePipelineRequest) (_result *CreatePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePipelineResponse{}
	_body, _err := client.CreatePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a program in a MediaWeaver channel.
//
// @param request - CreateProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProgramResponse
func (client *Client) CreateProgramWithOptions(request *CreateProgramRequest, runtime *dara.RuntimeOptions) (_result *CreateProgramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdBreaks) {
		query["AdBreaks"] = request.AdBreaks
	}

	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ClipRange) {
		query["ClipRange"] = request.ClipRange
	}

	if !dara.IsNil(request.ProgramName) {
		query["ProgramName"] = request.ProgramName
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Transition) {
		query["Transition"] = request.Transition
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProgram"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProgramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a program in a MediaWeaver channel.
//
// @param request - CreateProgramRequest
//
// @return CreateProgramResponse
func (client *Client) CreateProgram(request *CreateProgramRequest) (_result *CreateProgramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProgramResponse{}
	_body, _err := client.CreateProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义实体
//
// @param request - CreateRecognitionEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecognitionEntityResponse
func (client *Client) CreateRecognitionEntityWithOptions(request *CreateRecognitionEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateRecognitionEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.EntityInfo) {
		query["EntityInfo"] = request.EntityInfo
	}

	if !dara.IsNil(request.EntityName) {
		query["EntityName"] = request.EntityName
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecognitionEntity"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecognitionEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义实体
//
// @param request - CreateRecognitionEntityRequest
//
// @return CreateRecognitionEntityResponse
func (client *Client) CreateRecognitionEntity(request *CreateRecognitionEntityRequest) (_result *CreateRecognitionEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecognitionEntityResponse{}
	_body, _err := client.CreateRecognitionEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义库。
//
// @param request - CreateRecognitionLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecognitionLibResponse
func (client *Client) CreateRecognitionLibWithOptions(request *CreateRecognitionLibRequest, runtime *dara.RuntimeOptions) (_result *CreateRecognitionLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.LibDescription) {
		query["LibDescription"] = request.LibDescription
	}

	if !dara.IsNil(request.LibName) {
		query["LibName"] = request.LibName
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecognitionLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecognitionLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义库。
//
// @param request - CreateRecognitionLibRequest
//
// @return CreateRecognitionLibResponse
func (client *Client) CreateRecognitionLib(request *CreateRecognitionLibRequest) (_result *CreateRecognitionLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecognitionLibResponse{}
	_body, _err := client.CreateRecognitionLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 注册自定义实体照片
//
// @param request - CreateRecognitionSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRecognitionSampleResponse
func (client *Client) CreateRecognitionSampleWithOptions(request *CreateRecognitionSampleRequest, runtime *dara.RuntimeOptions) (_result *CreateRecognitionSampleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ImageUrl) {
		query["ImageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.LabelPrompt) {
		query["LabelPrompt"] = request.LabelPrompt
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRecognitionSample"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRecognitionSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 注册自定义实体照片
//
// @param request - CreateRecognitionSampleRequest
//
// @return CreateRecognitionSampleResponse
func (client *Client) CreateRecognitionSample(request *CreateRecognitionSampleRequest) (_result *CreateRecognitionSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRecognitionSampleResponse{}
	_body, _err := client.CreateRecognitionSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建搜索索引
//
// Description:
//
// The large visual model feature is still in the public preview phase. You can use this feature for free for 1,000 hours of videos.
//
// @param request - CreateSearchIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSearchIndexResponse
func (client *Client) CreateSearchIndexWithOptions(request *CreateSearchIndexRequest, runtime *dara.RuntimeOptions) (_result *CreateSearchIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexConfig) {
		query["IndexConfig"] = request.IndexConfig
	}

	if !dara.IsNil(request.IndexStatus) {
		query["IndexStatus"] = request.IndexStatus
	}

	if !dara.IsNil(request.IndexType) {
		query["IndexType"] = request.IndexType
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSearchIndex"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSearchIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建搜索索引
//
// Description:
//
// The large visual model feature is still in the public preview phase. You can use this feature for free for 1,000 hours of videos.
//
// @param request - CreateSearchIndexRequest
//
// @return CreateSearchIndexResponse
func (client *Client) CreateSearchIndex(request *CreateSearchIndexRequest) (_result *CreateSearchIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSearchIndexResponse{}
	_body, _err := client.CreateSearchIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a search library to store media assets.
//
// @param request - CreateSearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSearchLibResponse
func (client *Client) CreateSearchLibWithOptions(request *CreateSearchLibRequest, runtime *dara.RuntimeOptions) (_result *CreateSearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a search library to store media assets.
//
// @param request - CreateSearchLibRequest
//
// @return CreateSearchLibResponse
func (client *Client) CreateSearchLib(request *CreateSearchLibRequest) (_result *CreateSearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSearchLibResponse{}
	_body, _err := client.CreateSearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a source in MediaWeaver.
//
// @param request - CreateSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSourceResponse
func (client *Client) CreateSourceWithOptions(request *CreateSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HttpPackageConfigurations) {
		query["HttpPackageConfigurations"] = request.HttpPackageConfigurations
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSource"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a source in MediaWeaver.
//
// @param request - CreateSourceRequest
//
// @return CreateSourceResponse
func (client *Client) CreateSource(request *CreateSourceRequest) (_result *CreateSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSourceResponse{}
	_body, _err := client.CreateSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a source location.
//
// @param request - CreateSourceLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSourceLocationResponse
func (client *Client) CreateSourceLocationWithOptions(request *CreateSourceLocationRequest, runtime *dara.RuntimeOptions) (_result *CreateSourceLocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseUrl) {
		query["BaseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.EnableSegmentDelivery) {
		query["EnableSegmentDelivery"] = request.EnableSegmentDelivery
	}

	if !dara.IsNil(request.SegmentDeliveryUrl) {
		query["SegmentDeliveryUrl"] = request.SegmentDeliveryUrl
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSourceLocation"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSourceLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a source location.
//
// @param request - CreateSourceLocationRequest
//
// @return CreateSourceLocationResponse
func (client *Client) CreateSourceLocation(request *CreateSourceLocationRequest) (_result *CreateSourceLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSourceLocationResponse{}
	_body, _err := client.CreateSourceLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the upload URL and credential of a media asset and creates information about the media asset.
//
// Description:
//
//	  You can call this operation to obtain the upload URLs and credentials of audio and video files. You can also call this operation to obtain the upload URLs and credentials of images and auxiliary media assets.
//
//		- Obtaining an upload URL and credential is essential for Intelligent Media Services (IMS) and is required in each upload operation.
//
//		- If the video upload credential expires, you can call the RefreshUploadMedia operation to obtain a new upload credential. The default validity period of a video upload credential is 3,000 seconds.
//
//		- After you upload a media asset, you can configure a callback to receive upload event notifications or call the GetMediaInfo operation to determine whether the media asset is uploaded based on the returned status.
//
//		- The MediaId parameter returned by this operation can be used for media asset lifecycle management or media processing.
//
//		- You can call this operation to upload media assets only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media asset to your own OSS bucket, you can upload the file to your OSS bucket by using [OSS SDK](https://help.aliyun.com/document_detail/32006.html), and then call the [RegisterMediaInfo](https://help.aliyun.com/document_detail/441152.html) operation to register the file in the OSS bucket with the media asset library.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
// @param request - CreateUploadMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadMediaResponse
func (client *Client) CreateUploadMediaWithOptions(request *CreateUploadMediaRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.FileInfo) {
		query["FileInfo"] = request.FileInfo
	}

	if !dara.IsNil(request.MediaMetaData) {
		query["MediaMetaData"] = request.MediaMetaData
	}

	if !dara.IsNil(request.PostProcessConfig) {
		query["PostProcessConfig"] = request.PostProcessConfig
	}

	if !dara.IsNil(request.UploadTargetConfig) {
		query["UploadTargetConfig"] = request.UploadTargetConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUploadMedia"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the upload URL and credential of a media asset and creates information about the media asset.
//
// Description:
//
//	  You can call this operation to obtain the upload URLs and credentials of audio and video files. You can also call this operation to obtain the upload URLs and credentials of images and auxiliary media assets.
//
//		- Obtaining an upload URL and credential is essential for Intelligent Media Services (IMS) and is required in each upload operation.
//
//		- If the video upload credential expires, you can call the RefreshUploadMedia operation to obtain a new upload credential. The default validity period of a video upload credential is 3,000 seconds.
//
//		- After you upload a media asset, you can configure a callback to receive upload event notifications or call the GetMediaInfo operation to determine whether the media asset is uploaded based on the returned status.
//
//		- The MediaId parameter returned by this operation can be used for media asset lifecycle management or media processing.
//
//		- You can call this operation to upload media assets only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media asset to your own OSS bucket, you can upload the file to your OSS bucket by using [OSS SDK](https://help.aliyun.com/document_detail/32006.html), and then call the [RegisterMediaInfo](https://help.aliyun.com/document_detail/441152.html) operation to register the file in the OSS bucket with the media asset library.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
// @param request - CreateUploadMediaRequest
//
// @return CreateUploadMediaResponse
func (client *Client) CreateUploadMedia(request *CreateUploadMediaRequest) (_result *CreateUploadMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUploadMediaResponse{}
	_body, _err := client.CreateUploadMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the upload URL and credential of a media stream.
//
// Description:
//
//	  You can call this operation to upload only a local media stream. After the media stream is uploaded, it is associated with the specified media asset ID.
//
//		- You can call this operation to upload media streams only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media stream to your own OSS bucket, you can upload the file to your OSS bucket by using [OSS SDK](https://help.aliyun.com/document_detail/32006.html), and then call the [RegisterMediaStream](https://help.aliyun.com/document_detail/440765.html) operation to associate the media stream with the specified media asset ID.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
// @param request - CreateUploadStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadStreamResponse
func (client *Client) CreateUploadStreamWithOptions(request *CreateUploadStreamRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		query["Definition"] = request.Definition
	}

	if !dara.IsNil(request.FileExtension) {
		query["FileExtension"] = request.FileExtension
	}

	if !dara.IsNil(request.HDRType) {
		query["HDRType"] = request.HDRType
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUploadStream"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the upload URL and credential of a media stream.
//
// Description:
//
//	  You can call this operation to upload only a local media stream. After the media stream is uploaded, it is associated with the specified media asset ID.
//
//		- You can call this operation to upload media streams only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media stream to your own OSS bucket, you can upload the file to your OSS bucket by using [OSS SDK](https://help.aliyun.com/document_detail/32006.html), and then call the [RegisterMediaStream](https://help.aliyun.com/document_detail/440765.html) operation to associate the media stream with the specified media asset ID.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
// @param request - CreateUploadStreamRequest
//
// @return CreateUploadStreamResponse
func (client *Client) CreateUploadStream(request *CreateUploadStreamRequest) (_result *CreateUploadStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUploadStreamResponse{}
	_body, _err := client.CreateUploadStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Ingests an asset for VOD packaging.
//
// @param tmpReq - CreateVodPackagingAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVodPackagingAssetResponse
func (client *Client) CreateVodPackagingAssetWithOptions(tmpReq *CreateVodPackagingAssetRequest, runtime *dara.RuntimeOptions) (_result *CreateVodPackagingAssetResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateVodPackagingAssetShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetName) {
		query["AssetName"] = request.AssetName
	}

	if !dara.IsNil(request.ContentId) {
		query["ContentId"] = request.ContentId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVodPackagingAsset"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVodPackagingAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Ingests an asset for VOD packaging.
//
// @param request - CreateVodPackagingAssetRequest
//
// @return CreateVodPackagingAssetResponse
func (client *Client) CreateVodPackagingAsset(request *CreateVodPackagingAssetRequest) (_result *CreateVodPackagingAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVodPackagingAssetResponse{}
	_body, _err := client.CreateVodPackagingAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a packaging configuration.
//
// @param tmpReq - CreateVodPackagingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVodPackagingConfigurationResponse
func (client *Client) CreateVodPackagingConfigurationWithOptions(tmpReq *CreateVodPackagingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *CreateVodPackagingConfigurationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateVodPackagingConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PackageConfig) {
		request.PackageConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PackageConfig, dara.String("PackageConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationName) {
		query["ConfigurationName"] = request.ConfigurationName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.PackageConfigShrink) {
		query["PackageConfig"] = request.PackageConfigShrink
	}

	if !dara.IsNil(request.Protocol) {
		query["Protocol"] = request.Protocol
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVodPackagingConfiguration"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVodPackagingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a packaging configuration.
//
// @param request - CreateVodPackagingConfigurationRequest
//
// @return CreateVodPackagingConfigurationResponse
func (client *Client) CreateVodPackagingConfiguration(request *CreateVodPackagingConfigurationRequest) (_result *CreateVodPackagingConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVodPackagingConfigurationResponse{}
	_body, _err := client.CreateVodPackagingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a packaging group.
//
// @param request - CreateVodPackagingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVodPackagingGroupResponse
func (client *Client) CreateVodPackagingGroupWithOptions(request *CreateVodPackagingGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateVodPackagingGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVodPackagingGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVodPackagingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a packaging group.
//
// @param request - CreateVodPackagingGroupRequest
//
// @return CreateVodPackagingGroupResponse
func (client *Client) CreateVodPackagingGroup(request *CreateVodPackagingGroupRequest) (_result *CreateVodPackagingGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateVodPackagingGroupResponse{}
	_body, _err := client.CreateVodPackagingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Decrypts the ciphertext specified by CiphertextBlob in the Key Management Service (KMS) data key.
//
// @param request - DecryptKMSDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DecryptKMSDataKeyResponse
func (client *Client) DecryptKMSDataKeyWithOptions(request *DecryptKMSDataKeyRequest, runtime *dara.RuntimeOptions) (_result *DecryptKMSDataKeyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CiphertextBlob) {
		query["CiphertextBlob"] = request.CiphertextBlob
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DecryptKMSDataKey"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DecryptKMSDataKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Decrypts the ciphertext specified by CiphertextBlob in the Key Management Service (KMS) data key.
//
// @param request - DecryptKMSDataKeyRequest
//
// @return DecryptKMSDataKeyResponse
func (client *Client) DecryptKMSDataKey(request *DecryptKMSDataKeyRequest) (_result *DecryptKMSDataKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DecryptKMSDataKeyResponse{}
	_body, _err := client.DecryptKMSDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除智能体的对话历史记录。
//
// @param request - DeleteAIAgentDialogueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAIAgentDialogueResponse
func (client *Client) DeleteAIAgentDialogueWithOptions(request *DeleteAIAgentDialogueRequest, runtime *dara.RuntimeOptions) (_result *DeleteAIAgentDialogueResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DialogueId) {
		query["DialogueId"] = request.DialogueId
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAIAgentDialogue"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAIAgentDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除智能体的对话历史记录。
//
// @param request - DeleteAIAgentDialogueRequest
//
// @return DeleteAIAgentDialogueResponse
func (client *Client) DeleteAIAgentDialogue(request *DeleteAIAgentDialogueRequest) (_result *DeleteAIAgentDialogueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAIAgentDialogueResponse{}
	_body, _err := client.DeleteAIAgentDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ad insertion configuration.
//
// @param request - DeleteAdInsertionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAdInsertionResponse
func (client *Client) DeleteAdInsertionWithOptions(request *DeleteAdInsertionRequest, runtime *dara.RuntimeOptions) (_result *DeleteAdInsertionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAdInsertion"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAdInsertionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ad insertion configuration.
//
// @param request - DeleteAdInsertionRequest
//
// @return DeleteAdInsertionResponse
func (client *Client) DeleteAdInsertion(request *DeleteAdInsertionRequest) (_result *DeleteAdInsertionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAdInsertionResponse{}
	_body, _err := client.DeleteAdInsertionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a digital human training job that is in the Init or Fail state.
//
// @param request - DeleteAvatarTrainingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAvatarTrainingJobResponse
func (client *Client) DeleteAvatarTrainingJobWithOptions(request *DeleteAvatarTrainingJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteAvatarTrainingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAvatarTrainingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAvatarTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a digital human training job that is in the Init or Fail state.
//
// @param request - DeleteAvatarTrainingJobRequest
//
// @return DeleteAvatarTrainingJobResponse
func (client *Client) DeleteAvatarTrainingJob(request *DeleteAvatarTrainingJobRequest) (_result *DeleteAvatarTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAvatarTrainingJobResponse{}
	_body, _err := client.DeleteAvatarTrainingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a media asset category.
//
// Description:
//
// This operation also deletes the subcategories, including the level-2 and level-3 categories, of the category.
//
// @param request - DeleteCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategoryWithOptions(request *DeleteCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCategory"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a media asset category.
//
// Description:
//
// This operation also deletes the subcategories, including the level-2 and level-3 categories, of the category.
//
// @param request - DeleteCategoryRequest
//
// @return DeleteCategoryResponse
func (client *Client) DeleteCategory(request *DeleteCategoryRequest) (_result *DeleteCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCategoryResponse{}
	_body, _err := client.DeleteCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a channel in MediaWeaver.
//
// @param request - DeleteChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannelWithOptions(request *DeleteChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a channel in MediaWeaver.
//
// @param request - DeleteChannelRequest
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannel(request *DeleteChannelRequest) (_result *DeleteChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChannelResponse{}
	_body, _err := client.DeleteChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom template.
//
// @param request - DeleteCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomTemplateResponse
func (client *Client) DeleteCustomTemplateWithOptions(request *DeleteCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom template.
//
// @param request - DeleteCustomTemplateRequest
//
// @return DeleteCustomTemplateResponse
func (client *Client) DeleteCustomTemplate(request *DeleteCustomTemplateRequest) (_result *DeleteCustomTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomTemplateResponse{}
	_body, _err := client.DeleteCustomTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a human voice cloning job that is not in the Training or Success state.
//
// @param request - DeleteCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomizedVoiceJobResponse
func (client *Client) DeleteCustomizedVoiceJobWithOptions(request *DeleteCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a human voice cloning job that is not in the Training or Success state.
//
// @param request - DeleteCustomizedVoiceJobRequest
//
// @return DeleteCustomizedVoiceJobResponse
func (client *Client) DeleteCustomizedVoiceJob(request *DeleteCustomizedVoiceJobRequest) (_result *DeleteCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteCustomizedVoiceJobResponse{}
	_body, _err := client.DeleteCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a media fingerprint library.
//
// @param request - DeleteDNADBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDNADBResponse
func (client *Client) DeleteDNADBWithOptions(request *DeleteDNADBRequest, runtime *dara.RuntimeOptions) (_result *DeleteDNADBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBId) {
		query["DBId"] = request.DBId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDNADB"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDNADBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a media fingerprint library.
//
// @param request - DeleteDNADBRequest
//
// @return DeleteDNADBResponse
func (client *Client) DeleteDNADB(request *DeleteDNADBRequest) (_result *DeleteDNADBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDNADBResponse{}
	_body, _err := client.DeleteDNADBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes files from a media fingerprint library.
//
// @param request - DeleteDNAFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDNAFilesResponse
func (client *Client) DeleteDNAFilesWithOptions(request *DeleteDNAFilesRequest, runtime *dara.RuntimeOptions) (_result *DeleteDNAFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBId) {
		query["DBId"] = request.DBId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PrimaryKeys) {
		query["PrimaryKeys"] = request.PrimaryKeys
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDNAFiles"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDNAFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes files from a media fingerprint library.
//
// @param request - DeleteDNAFilesRequest
//
// @return DeleteDNAFilesResponse
func (client *Client) DeleteDNAFiles(request *DeleteDNAFilesRequest) (_result *DeleteDNAFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDNAFilesResponse{}
	_body, _err := client.DeleteDNAFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes one or more materials from an online editing project.
//
// @param request - DeleteEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEditingProjectMaterialsResponse
func (client *Client) DeleteEditingProjectMaterialsWithOptions(request *DeleteEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaterialIds) {
		query["MaterialIds"] = request.MaterialIds
	}

	if !dara.IsNil(request.MaterialType) {
		query["MaterialType"] = request.MaterialType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEditingProjectMaterials"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEditingProjectMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more materials from an online editing project.
//
// @param request - DeleteEditingProjectMaterialsRequest
//
// @return DeleteEditingProjectMaterialsResponse
func (client *Client) DeleteEditingProjectMaterials(request *DeleteEditingProjectMaterialsRequest) (_result *DeleteEditingProjectMaterialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEditingProjectMaterialsResponse{}
	_body, _err := client.DeleteEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes one or more online editing project.
//
// @param request - DeleteEditingProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEditingProjectsResponse
func (client *Client) DeleteEditingProjectsWithOptions(request *DeleteEditingProjectsRequest, runtime *dara.RuntimeOptions) (_result *DeleteEditingProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectIds) {
		query["ProjectIds"] = request.ProjectIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEditingProjects"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEditingProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes one or more online editing project.
//
// @param request - DeleteEditingProjectsRequest
//
// @return DeleteEditingProjectsResponse
func (client *Client) DeleteEditingProjects(request *DeleteEditingProjectsRequest) (_result *DeleteEditingProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteEditingProjectsResponse{}
	_body, _err := client.DeleteEditingProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除热词词库
//
// @param request - DeleteHotwordLibraryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteHotwordLibraryResponse
func (client *Client) DeleteHotwordLibraryWithOptions(request *DeleteHotwordLibraryRequest, runtime *dara.RuntimeOptions) (_result *DeleteHotwordLibraryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HotwordLibraryId) {
		query["HotwordLibraryId"] = request.HotwordLibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteHotwordLibrary"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteHotwordLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除热词词库
//
// @param request - DeleteHotwordLibraryRequest
//
// @return DeleteHotwordLibraryResponse
func (client *Client) DeleteHotwordLibrary(request *DeleteHotwordLibraryRequest) (_result *DeleteHotwordLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteHotwordLibraryResponse{}
	_body, _err := client.DeleteHotwordLibraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a live package channel by GroupName and ChannelName.
//
// Description:
//
// ## [](#)Usage notes
//
// You need to provide GroupName and ChannelName as parameters to specify exactly which channel to delete. Before you delete a channel, you must delete the origin endpoints associated with the channel.
//
// @param request - DeleteLivePackageChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePackageChannelResponse
func (client *Client) DeleteLivePackageChannelWithOptions(request *DeleteLivePackageChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePackageChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivePackageChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePackageChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live package channel by GroupName and ChannelName.
//
// Description:
//
// ## [](#)Usage notes
//
// You need to provide GroupName and ChannelName as parameters to specify exactly which channel to delete. Before you delete a channel, you must delete the origin endpoints associated with the channel.
//
// @param request - DeleteLivePackageChannelRequest
//
// @return DeleteLivePackageChannelResponse
func (client *Client) DeleteLivePackageChannel(request *DeleteLivePackageChannelRequest) (_result *DeleteLivePackageChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLivePackageChannelResponse{}
	_body, _err := client.DeleteLivePackageChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a live package channel group by name.
//
// Description:
//
// ## [](#)Usage notes
//
// Make sure that no channels are included in the channel group before you delete it.
//
// @param request - DeleteLivePackageChannelGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePackageChannelGroupResponse
func (client *Client) DeleteLivePackageChannelGroupWithOptions(request *DeleteLivePackageChannelGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePackageChannelGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivePackageChannelGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePackageChannelGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live package channel group by name.
//
// Description:
//
// ## [](#)Usage notes
//
// Make sure that no channels are included in the channel group before you delete it.
//
// @param request - DeleteLivePackageChannelGroupRequest
//
// @return DeleteLivePackageChannelGroupResponse
func (client *Client) DeleteLivePackageChannelGroup(request *DeleteLivePackageChannelGroupRequest) (_result *DeleteLivePackageChannelGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLivePackageChannelGroupResponse{}
	_body, _err := client.DeleteLivePackageChannelGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an origin endpoint associated with a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation is used to delete an origin endpoint associated with a live package channel by specifying `GroupName`, `ChannelName`, and `EndpointName`. This operation will permanently delete the relevant configurations. Exercise caution when you perform this operation.
//
// @param request - DeleteLivePackageOriginEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLivePackageOriginEndpointResponse
func (client *Client) DeleteLivePackageOriginEndpointWithOptions(request *DeleteLivePackageOriginEndpointRequest, runtime *dara.RuntimeOptions) (_result *DeleteLivePackageOriginEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.EndpointName) {
		query["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLivePackageOriginEndpoint"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLivePackageOriginEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an origin endpoint associated with a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation is used to delete an origin endpoint associated with a live package channel by specifying `GroupName`, `ChannelName`, and `EndpointName`. This operation will permanently delete the relevant configurations. Exercise caution when you perform this operation.
//
// @param request - DeleteLivePackageOriginEndpointRequest
//
// @return DeleteLivePackageOriginEndpointResponse
func (client *Client) DeleteLivePackageOriginEndpoint(request *DeleteLivePackageOriginEndpointRequest) (_result *DeleteLivePackageOriginEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLivePackageOriginEndpointResponse{}
	_body, _err := client.DeleteLivePackageOriginEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes live stream recording files. You can choose to delete only the recording files or delete both the recording files and the original Object Storage Service (OSS) files.
//
// @param request - DeleteLiveRecordFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveRecordFilesResponse
func (client *Client) DeleteLiveRecordFilesWithOptions(request *DeleteLiveRecordFilesRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveRecordFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RecordIds) {
		query["RecordIds"] = request.RecordIds
	}

	if !dara.IsNil(request.RemoveFile) {
		query["RemoveFile"] = request.RemoveFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveRecordFiles"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveRecordFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes live stream recording files. You can choose to delete only the recording files or delete both the recording files and the original Object Storage Service (OSS) files.
//
// @param request - DeleteLiveRecordFilesRequest
//
// @return DeleteLiveRecordFilesResponse
func (client *Client) DeleteLiveRecordFiles(request *DeleteLiveRecordFilesRequest) (_result *DeleteLiveRecordFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLiveRecordFilesResponse{}
	_body, _err := client.DeleteLiveRecordFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a live stream recording template without affecting existing jobs.
//
// @param request - DeleteLiveRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveRecordTemplateResponse
func (client *Client) DeleteLiveRecordTemplateWithOptions(request *DeleteLiveRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveRecordTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveRecordTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live stream recording template without affecting existing jobs.
//
// @param request - DeleteLiveRecordTemplateRequest
//
// @return DeleteLiveRecordTemplateResponse
func (client *Client) DeleteLiveRecordTemplate(request *DeleteLiveRecordTemplateRequest) (_result *DeleteLiveRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLiveRecordTemplateResponse{}
	_body, _err := client.DeleteLiveRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes live stream snapshot files. You can choose to delete only the snapshot files or delete both the snapshot files and the original Object Storage Service (OSS) files.
//
// @param tmpReq - DeleteLiveSnapshotFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveSnapshotFilesResponse
func (client *Client) DeleteLiveSnapshotFilesWithOptions(tmpReq *DeleteLiveSnapshotFilesRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveSnapshotFilesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DeleteLiveSnapshotFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CreateTimestampList) {
		request.CreateTimestampListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CreateTimestampList, dara.String("CreateTimestampList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimestampListShrink) {
		query["CreateTimestampList"] = request.CreateTimestampListShrink
	}

	if !dara.IsNil(request.DeleteOriginalFile) {
		query["DeleteOriginalFile"] = request.DeleteOriginalFile
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveSnapshotFiles"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveSnapshotFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes live stream snapshot files. You can choose to delete only the snapshot files or delete both the snapshot files and the original Object Storage Service (OSS) files.
//
// @param request - DeleteLiveSnapshotFilesRequest
//
// @return DeleteLiveSnapshotFilesResponse
func (client *Client) DeleteLiveSnapshotFiles(request *DeleteLiveSnapshotFilesRequest) (_result *DeleteLiveSnapshotFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLiveSnapshotFilesResponse{}
	_body, _err := client.DeleteLiveSnapshotFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a live stream snapshot template.
//
// @param request - DeleteLiveSnapshotTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveSnapshotTemplateResponse
func (client *Client) DeleteLiveSnapshotTemplateWithOptions(request *DeleteLiveSnapshotTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveSnapshotTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveSnapshotTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveSnapshotTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live stream snapshot template.
//
// @param request - DeleteLiveSnapshotTemplateRequest
//
// @return DeleteLiveSnapshotTemplateResponse
func (client *Client) DeleteLiveSnapshotTemplate(request *DeleteLiveSnapshotTemplateRequest) (_result *DeleteLiveSnapshotTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLiveSnapshotTemplateResponse{}
	_body, _err := client.DeleteLiveSnapshotTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定转码任务
//
// @param request - DeleteLiveTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveTranscodeJobResponse
func (client *Client) DeleteLiveTranscodeJobWithOptions(request *DeleteLiveTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveTranscodeJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveTranscodeJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveTranscodeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定转码任务
//
// @param request - DeleteLiveTranscodeJobRequest
//
// @return DeleteLiveTranscodeJobResponse
func (client *Client) DeleteLiveTranscodeJob(request *DeleteLiveTranscodeJobRequest) (_result *DeleteLiveTranscodeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLiveTranscodeJobResponse{}
	_body, _err := client.DeleteLiveTranscodeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a live stream transcoding template.
//
// @param request - DeleteLiveTranscodeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLiveTranscodeTemplateResponse
func (client *Client) DeleteLiveTranscodeTemplateWithOptions(request *DeleteLiveTranscodeTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteLiveTranscodeTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLiveTranscodeTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLiveTranscodeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a live stream transcoding template.
//
// @param request - DeleteLiveTranscodeTemplateRequest
//
// @return DeleteLiveTranscodeTemplateResponse
func (client *Client) DeleteLiveTranscodeTemplate(request *DeleteLiveTranscodeTemplateRequest) (_result *DeleteLiveTranscodeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLiveTranscodeTemplateResponse{}
	_body, _err := client.DeleteLiveTranscodeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- When a flow is deleted, its source and outputs are also deleted.
//
//		- When a flow is in the online state, it cannot be deleted.
//
// @param request - DeleteMediaConnectFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaConnectFlowResponse
func (client *Client) DeleteMediaConnectFlowWithOptions(request *DeleteMediaConnectFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaConnectFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaConnectFlow"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaConnectFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- When a flow is deleted, its source and outputs are also deleted.
//
//		- When a flow is in the online state, it cannot be deleted.
//
// @param request - DeleteMediaConnectFlowRequest
//
// @return DeleteMediaConnectFlowResponse
func (client *Client) DeleteMediaConnectFlow(request *DeleteMediaConnectFlowRequest) (_result *DeleteMediaConnectFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaConnectFlowResponse{}
	_body, _err := client.DeleteMediaConnectFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the source of a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- When a flow is in the online state, its source cannot be deleted.
//
//		- You can delete the source only after all outputs of the flow have been deleted.
//
// @param request - DeleteMediaConnectFlowInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaConnectFlowInputResponse
func (client *Client) DeleteMediaConnectFlowInputWithOptions(request *DeleteMediaConnectFlowInputRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaConnectFlowInputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.InputName) {
		query["InputName"] = request.InputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaConnectFlowInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaConnectFlowInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the source of a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- When a flow is in the online state, its source cannot be deleted.
//
//		- You can delete the source only after all outputs of the flow have been deleted.
//
// @param request - DeleteMediaConnectFlowInputRequest
//
// @return DeleteMediaConnectFlowInputResponse
func (client *Client) DeleteMediaConnectFlowInput(request *DeleteMediaConnectFlowInputRequest) (_result *DeleteMediaConnectFlowInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaConnectFlowInputResponse{}
	_body, _err := client.DeleteMediaConnectFlowInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an output of a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- When a flow is in the online state, its outputs cannot be deleted.
//
// @param request - DeleteMediaConnectFlowOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaConnectFlowOutputResponse
func (client *Client) DeleteMediaConnectFlowOutputWithOptions(request *DeleteMediaConnectFlowOutputRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaConnectFlowOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaConnectFlowOutput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaConnectFlowOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an output of a MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- When a flow is in the online state, its outputs cannot be deleted.
//
// @param request - DeleteMediaConnectFlowOutputRequest
//
// @return DeleteMediaConnectFlowOutputResponse
func (client *Client) DeleteMediaConnectFlowOutput(request *DeleteMediaConnectFlowOutputRequest) (_result *DeleteMediaConnectFlowOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaConnectFlowOutputResponse{}
	_body, _err := client.DeleteMediaConnectFlowOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific media asset from a search library.
//
// @param request - DeleteMediaFromSearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaFromSearchLibResponse
func (client *Client) DeleteMediaFromSearchLibWithOptions(request *DeleteMediaFromSearchLibRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaFromSearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MsgBody) {
		query["MsgBody"] = request.MsgBody
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaFromSearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaFromSearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific media asset from a search library.
//
// @param request - DeleteMediaFromSearchLibRequest
//
// @return DeleteMediaFromSearchLibResponse
func (client *Client) DeleteMediaFromSearchLib(request *DeleteMediaFromSearchLibRequest) (_result *DeleteMediaFromSearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaFromSearchLibResponse{}
	_body, _err := client.DeleteMediaFromSearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple media assets at a time. You can delete at most 20 media assets at a time. If MediaIds is specified, it is preferentially used. If MediaIds is empty, InputURLs must be specified.
//
// @param request - DeleteMediaInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaInfosResponse
func (client *Client) DeleteMediaInfosWithOptions(request *DeleteMediaInfosRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletePhysicalFiles) {
		query["DeletePhysicalFiles"] = request.DeletePhysicalFiles
	}

	if !dara.IsNil(request.InputURLs) {
		query["InputURLs"] = request.InputURLs
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaInfos"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple media assets at a time. You can delete at most 20 media assets at a time. If MediaIds is specified, it is preferentially used. If MediaIds is empty, InputURLs must be specified.
//
// @param request - DeleteMediaInfosRequest
//
// @return DeleteMediaInfosResponse
func (client *Client) DeleteMediaInfos(request *DeleteMediaInfosRequest) (_result *DeleteMediaInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaInfosResponse{}
	_body, _err := client.DeleteMediaInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a MediaLive channel.
//
// Description:
//
//	You can only delete a channel that is not running.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteMediaLiveChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaLiveChannelResponse
func (client *Client) DeleteMediaLiveChannelWithOptions(request *DeleteMediaLiveChannelRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaLiveChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaLiveChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaLiveChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a MediaLive channel.
//
// Description:
//
//	You can only delete a channel that is not running.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteMediaLiveChannelRequest
//
// @return DeleteMediaLiveChannelResponse
func (client *Client) DeleteMediaLiveChannel(request *DeleteMediaLiveChannelRequest) (_result *DeleteMediaLiveChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaLiveChannelResponse{}
	_body, _err := client.DeleteMediaLiveChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a MediaLive input.
//
// Description:
//
//	You can delete an input only when it is not associated with a MediaLive channel.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteMediaLiveInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaLiveInputResponse
func (client *Client) DeleteMediaLiveInputWithOptions(request *DeleteMediaLiveInputRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaLiveInputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InputId) {
		body["InputId"] = request.InputId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaLiveInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaLiveInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a MediaLive input.
//
// Description:
//
//	You can delete an input only when it is not associated with a MediaLive channel.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteMediaLiveInputRequest
//
// @return DeleteMediaLiveInputResponse
func (client *Client) DeleteMediaLiveInput(request *DeleteMediaLiveInputRequest) (_result *DeleteMediaLiveInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaLiveInputResponse{}
	_body, _err := client.DeleteMediaLiveInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a security group in MediaLive.
//
// Description:
//
//	You can only delete a security group not associated with an input.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteMediaLiveInputSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaLiveInputSecurityGroupResponse
func (client *Client) DeleteMediaLiveInputSecurityGroupWithOptions(request *DeleteMediaLiveInputSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaLiveInputSecurityGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaLiveInputSecurityGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security group in MediaLive.
//
// Description:
//
//	You can only delete a security group not associated with an input.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - DeleteMediaLiveInputSecurityGroupRequest
//
// @return DeleteMediaLiveInputSecurityGroupResponse
func (client *Client) DeleteMediaLiveInputSecurityGroup(request *DeleteMediaLiveInputSecurityGroupRequest) (_result *DeleteMediaLiveInputSecurityGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.DeleteMediaLiveInputSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the marks of a media asset.
//
// @param request - DeleteMediaMarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMediaMarksResponse
func (client *Client) DeleteMediaMarksWithOptions(request *DeleteMediaMarksRequest, runtime *dara.RuntimeOptions) (_result *DeleteMediaMarksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaMarkIds) {
		query["MediaMarkIds"] = request.MediaMarkIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMediaMarks"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMediaMarksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the marks of a media asset.
//
// @param request - DeleteMediaMarksRequest
//
// @return DeleteMediaMarksResponse
func (client *Client) DeleteMediaMarks(request *DeleteMediaMarksRequest) (_result *DeleteMediaMarksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMediaMarksResponse{}
	_body, _err := client.DeleteMediaMarksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - DeletePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineResponse
func (client *Client) DeletePipelineWithOptions(request *DeletePipelineRequest, runtime *dara.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePipeline"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - DeletePipelineRequest
//
// @return DeletePipelineResponse
func (client *Client) DeletePipeline(request *DeletePipelineRequest) (_result *DeletePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePipelineResponse{}
	_body, _err := client.DeletePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes media streams such as video streams and audio streams.
//
// Description:
//
// You can call this operation to delete multiple media streams at a time.
//
// @param request - DeletePlayInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePlayInfoResponse
func (client *Client) DeletePlayInfoWithOptions(request *DeletePlayInfoRequest, runtime *dara.RuntimeOptions) (_result *DeletePlayInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletePhysicalFiles) {
		query["DeletePhysicalFiles"] = request.DeletePhysicalFiles
	}

	if !dara.IsNil(request.FileURLs) {
		query["FileURLs"] = request.FileURLs
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeletePlayInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePlayInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes media streams such as video streams and audio streams.
//
// Description:
//
// You can call this operation to delete multiple media streams at a time.
//
// @param request - DeletePlayInfoRequest
//
// @return DeletePlayInfoResponse
func (client *Client) DeletePlayInfo(request *DeletePlayInfoRequest) (_result *DeletePlayInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeletePlayInfoResponse{}
	_body, _err := client.DeletePlayInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a program from a channel.
//
// @param request - DeleteProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProgramResponse
func (client *Client) DeleteProgramWithOptions(request *DeleteProgramRequest, runtime *dara.RuntimeOptions) (_result *DeleteProgramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ProgramName) {
		query["ProgramName"] = request.ProgramName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProgram"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProgramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a program from a channel.
//
// @param request - DeleteProgramRequest
//
// @return DeleteProgramResponse
func (client *Client) DeleteProgram(request *DeleteProgramRequest) (_result *DeleteProgramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProgramResponse{}
	_body, _err := client.DeleteProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义库实体
//
// @param request - DeleteRecognitionEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecognitionEntityResponse
func (client *Client) DeleteRecognitionEntityWithOptions(request *DeleteRecognitionEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecognitionEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecognitionEntity"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecognitionEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义库实体
//
// @param request - DeleteRecognitionEntityRequest
//
// @return DeleteRecognitionEntityResponse
func (client *Client) DeleteRecognitionEntity(request *DeleteRecognitionEntityRequest) (_result *DeleteRecognitionEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecognitionEntityResponse{}
	_body, _err := client.DeleteRecognitionEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义库
//
// @param request - DeleteRecognitionLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecognitionLibResponse
func (client *Client) DeleteRecognitionLibWithOptions(request *DeleteRecognitionLibRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecognitionLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecognitionLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecognitionLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义库
//
// @param request - DeleteRecognitionLibRequest
//
// @return DeleteRecognitionLibResponse
func (client *Client) DeleteRecognitionLib(request *DeleteRecognitionLibRequest) (_result *DeleteRecognitionLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecognitionLibResponse{}
	_body, _err := client.DeleteRecognitionLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义实体照片
//
// @param request - DeleteRecognitionSampleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecognitionSampleResponse
func (client *Client) DeleteRecognitionSampleWithOptions(request *DeleteRecognitionSampleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecognitionSampleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SampleId) {
		query["SampleId"] = request.SampleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecognitionSample"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecognitionSampleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义实体照片
//
// @param request - DeleteRecognitionSampleRequest
//
// @return DeleteRecognitionSampleResponse
func (client *Client) DeleteRecognitionSample(request *DeleteRecognitionSampleRequest) (_result *DeleteRecognitionSampleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecognitionSampleResponse{}
	_body, _err := client.DeleteRecognitionSampleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes intelligent jobs based on job IDs.
//
// @param request - DeleteSmartJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSmartJobResponse
func (client *Client) DeleteSmartJobWithOptions(request *DeleteSmartJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteSmartJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSmartJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSmartJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes intelligent jobs based on job IDs.
//
// @param request - DeleteSmartJobRequest
//
// @return DeleteSmartJobResponse
func (client *Client) DeleteSmartJob(request *DeleteSmartJobRequest) (_result *DeleteSmartJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSmartJobResponse{}
	_body, _err := client.DeleteSmartJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a source from MediaWeaver.
//
// @param request - DeleteSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceResponse
func (client *Client) DeleteSourceWithOptions(request *DeleteSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SoftDelete) {
		query["SoftDelete"] = request.SoftDelete
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSource"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a source from MediaWeaver.
//
// @param request - DeleteSourceRequest
//
// @return DeleteSourceResponse
func (client *Client) DeleteSource(request *DeleteSourceRequest) (_result *DeleteSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSourceResponse{}
	_body, _err := client.DeleteSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a source location.
//
// @param request - DeleteSourceLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSourceLocationResponse
func (client *Client) DeleteSourceLocationWithOptions(request *DeleteSourceLocationRequest, runtime *dara.RuntimeOptions) (_result *DeleteSourceLocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SoftDelete) {
		query["SoftDelete"] = request.SoftDelete
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSourceLocation"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSourceLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a source location.
//
// @param request - DeleteSourceLocationRequest
//
// @return DeleteSourceLocationResponse
func (client *Client) DeleteSourceLocation(request *DeleteSourceLocationRequest) (_result *DeleteSourceLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteSourceLocationResponse{}
	_body, _err := client.DeleteSourceLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes templates.
//
// Description:
//
// A template is an encapsulation of the timeline of a media editing and production job. You can define a common timeline as a template. When you have the same requirements, you need to only specify key parameters and materials to produce videos.
//
//   - For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//   - For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes templates.
//
// Description:
//
// A template is an encapsulation of the timeline of a media editing and production job. You can define a common timeline as a template. When you have the same requirements, you need to only specify key parameters and materials to produce videos.
//
//   - For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//   - For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a VOD packaging asset.
//
// @param request - DeleteVodPackagingAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodPackagingAssetResponse
func (client *Client) DeleteVodPackagingAssetWithOptions(request *DeleteVodPackagingAssetRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodPackagingAssetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetName) {
		query["AssetName"] = request.AssetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVodPackagingAsset"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVodPackagingAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a VOD packaging asset.
//
// @param request - DeleteVodPackagingAssetRequest
//
// @return DeleteVodPackagingAssetResponse
func (client *Client) DeleteVodPackagingAsset(request *DeleteVodPackagingAssetRequest) (_result *DeleteVodPackagingAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVodPackagingAssetResponse{}
	_body, _err := client.DeleteVodPackagingAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a packaging configuration.
//
// @param request - DeleteVodPackagingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodPackagingConfigurationResponse
func (client *Client) DeleteVodPackagingConfigurationWithOptions(request *DeleteVodPackagingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodPackagingConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationName) {
		query["ConfigurationName"] = request.ConfigurationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVodPackagingConfiguration"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVodPackagingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a packaging configuration.
//
// @param request - DeleteVodPackagingConfigurationRequest
//
// @return DeleteVodPackagingConfigurationResponse
func (client *Client) DeleteVodPackagingConfiguration(request *DeleteVodPackagingConfigurationRequest) (_result *DeleteVodPackagingConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVodPackagingConfigurationResponse{}
	_body, _err := client.DeleteVodPackagingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a packaging group.
//
// @param request - DeleteVodPackagingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVodPackagingGroupResponse
func (client *Client) DeleteVodPackagingGroupWithOptions(request *DeleteVodPackagingGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteVodPackagingGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVodPackagingGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVodPackagingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a packaging group.
//
// @param request - DeleteVodPackagingGroupRequest
//
// @return DeleteVodPackagingGroupResponse
func (client *Client) DeleteVodPackagingGroup(request *DeleteVodPackagingGroupRequest) (_result *DeleteVodPackagingGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteVodPackagingGroupResponse{}
	_body, _err := client.DeleteVodPackagingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an AI agent.
//
// Description:
//
// ## [](#)Request description
//
//   - **Feature**: You can call this operation to query the information about an AI agent.
//
//   - **Scenario**: If you need to monitor or analyze the performance of an AI agent in a call or debug the agent configurations, you can call this operation to obtain required data.
//
// @param request - DescribeAIAgentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAIAgentInstanceResponse
func (client *Client) DescribeAIAgentInstanceWithOptions(request *DescribeAIAgentInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeAIAgentInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAIAgentInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAIAgentInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an AI agent.
//
// Description:
//
// ## [](#)Request description
//
//   - **Feature**: You can call this operation to query the information about an AI agent.
//
//   - **Scenario**: If you need to monitor or analyze the performance of an AI agent in a call or debug the agent configurations, you can call this operation to obtain required data.
//
// @param request - DescribeAIAgentInstanceRequest
//
// @return DescribeAIAgentInstanceResponse
func (client *Client) DescribeAIAgentInstance(request *DescribeAIAgentInstanceRequest) (_result *DescribeAIAgentInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAIAgentInstanceResponse{}
	_body, _err := client.DescribeAIAgentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on video-on-demand (VOD) editing. The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsEditUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeterImsEditUsageResponse
func (client *Client) DescribeMeterImsEditUsageWithOptions(request *DescribeMeterImsEditUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeterImsEditUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeterImsEditUsage"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeterImsEditUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on video-on-demand (VOD) editing. The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsEditUsageRequest
//
// @return DescribeMeterImsEditUsageResponse
func (client *Client) DescribeMeterImsEditUsage(request *DescribeMeterImsEditUsageRequest) (_result *DescribeMeterImsEditUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMeterImsEditUsageResponse{}
	_body, _err := client.DescribeMeterImsEditUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on ultra high definition (UHD) transcoding of ApsaraVideo Media Processing (MPS). The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsMediaConvertUHDUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeterImsMediaConvertUHDUsageResponse
func (client *Client) DescribeMeterImsMediaConvertUHDUsageWithOptions(request *DescribeMeterImsMediaConvertUHDUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeterImsMediaConvertUHDUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeterImsMediaConvertUHDUsage"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeterImsMediaConvertUHDUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on ultra high definition (UHD) transcoding of ApsaraVideo Media Processing (MPS). The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsMediaConvertUHDUsageRequest
//
// @return DescribeMeterImsMediaConvertUHDUsageResponse
func (client *Client) DescribeMeterImsMediaConvertUHDUsage(request *DescribeMeterImsMediaConvertUHDUsageRequest) (_result *DescribeMeterImsMediaConvertUHDUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMeterImsMediaConvertUHDUsageResponse{}
	_body, _err := client.DescribeMeterImsMediaConvertUHDUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on video-on-demand (VOD) transcoding. The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsMediaConvertUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeterImsMediaConvertUsageResponse
func (client *Client) DescribeMeterImsMediaConvertUsageWithOptions(request *DescribeMeterImsMediaConvertUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeterImsMediaConvertUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeterImsMediaConvertUsage"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeterImsMediaConvertUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on video-on-demand (VOD) transcoding. The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsMediaConvertUsageRequest
//
// @return DescribeMeterImsMediaConvertUsageResponse
func (client *Client) DescribeMeterImsMediaConvertUsage(request *DescribeMeterImsMediaConvertUsageRequest) (_result *DescribeMeterImsMediaConvertUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMeterImsMediaConvertUsageResponse{}
	_body, _err := client.DescribeMeterImsMediaConvertUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on AI processing of ApsaraVideo Media Processing (MPS). The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsMpsAiUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeterImsMpsAiUsageResponse
func (client *Client) DescribeMeterImsMpsAiUsageWithOptions(request *DescribeMeterImsMpsAiUsageRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeterImsMpsAiUsageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeterImsMpsAiUsage"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeterImsMpsAiUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS) on AI processing of ApsaraVideo Media Processing (MPS). The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsMpsAiUsageRequest
//
// @return DescribeMeterImsMpsAiUsageResponse
func (client *Client) DescribeMeterImsMpsAiUsage(request *DescribeMeterImsMpsAiUsageRequest) (_result *DescribeMeterImsMpsAiUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMeterImsMpsAiUsageResponse{}
	_body, _err := client.DescribeMeterImsMpsAiUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS). The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMeterImsSummaryResponse
func (client *Client) DescribeMeterImsSummaryWithOptions(request *DescribeMeterImsSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescribeMeterImsSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.StartTs) {
		query["StartTs"] = request.StartTs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeMeterImsSummary"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeMeterImsSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the usage statistics of Intelligent Media Services (IMS). The maximum query range is 31 days. You can query data within the last 90 days.
//
// @param request - DescribeMeterImsSummaryRequest
//
// @return DescribeMeterImsSummaryResponse
func (client *Client) DescribeMeterImsSummary(request *DescribeMeterImsSummaryRequest) (_result *DescribeMeterImsSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeMeterImsSummaryResponse{}
	_body, _err := client.DescribeMeterImsSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the event callback configurations of an AI agent.
//
// Description:
//
// You can call this operation to query the detailed callback configurations of an AI agent.
//
// @param request - DescribeNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNotifyConfigResponse
func (client *Client) DescribeNotifyConfigWithOptions(request *DescribeNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *DescribeNotifyConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNotifyConfig"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeNotifyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the event callback configurations of an AI agent.
//
// Description:
//
// You can call this operation to query the detailed callback configurations of an AI agent.
//
// @param request - DescribeNotifyConfigRequest
//
// @return DescribeNotifyConfigResponse
func (client *Client) DescribeNotifyConfig(request *DescribeNotifyConfigRequest) (_result *DescribeNotifyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeNotifyConfigResponse{}
	_body, _err := client.DescribeNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribePlayListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePlayListResponse
func (client *Client) DescribePlayListWithOptions(request *DescribePlayListRequest, runtime *dara.RuntimeOptions) (_result *DescribePlayListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTs) {
		query["BeginTs"] = request.BeginTs
	}

	if !dara.IsNil(request.EndTs) {
		query["EndTs"] = request.EndTs
	}

	if !dara.IsNil(request.OrderName) {
		query["OrderName"] = request.OrderName
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlayType) {
		query["PlayType"] = request.PlayType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TraceId) {
		query["TraceId"] = request.TraceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePlayList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePlayListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribePlayListRequest
//
// @return DescribePlayListResponse
func (client *Client) DescribePlayList(request *DescribePlayListRequest) (_result *DescribePlayListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePlayListResponse{}
	_body, _err := client.DescribePlayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例
//
// @param request - DescribeRtcRobotInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcRobotInstanceResponse
func (client *Client) DescribeRtcRobotInstanceWithOptions(request *DescribeRtcRobotInstanceRequest, runtime *dara.RuntimeOptions) (_result *DescribeRtcRobotInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRtcRobotInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRtcRobotInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例
//
// @param request - DescribeRtcRobotInstanceRequest
//
// @return DescribeRtcRobotInstanceResponse
func (client *Client) DescribeRtcRobotInstance(request *DescribeRtcRobotInstanceRequest) (_result *DescribeRtcRobotInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRtcRobotInstanceResponse{}
	_body, _err := client.DescribeRtcRobotInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether the reading of users has issues, such as noticeable pronunciation errors or background noise. After the audio is checked on the cloud, the qualified audio is temporarily stored on the cloud for subsequent training. Do not skip this step.
//
// @param request - DetectAudioForCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectAudioForCustomizedVoiceJobResponse
func (client *Client) DetectAudioForCustomizedVoiceJobWithOptions(request *DetectAudioForCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *DetectAudioForCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AudioRecordId) {
		query["AudioRecordId"] = request.AudioRecordId
	}

	if !dara.IsNil(request.RecordUrl) {
		query["RecordUrl"] = request.RecordUrl
	}

	if !dara.IsNil(request.VoiceId) {
		query["VoiceId"] = request.VoiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectAudioForCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectAudioForCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether the reading of users has issues, such as noticeable pronunciation errors or background noise. After the audio is checked on the cloud, the qualified audio is temporarily stored on the cloud for subsequent training. Do not skip this step.
//
// @param request - DetectAudioForCustomizedVoiceJobRequest
//
// @return DetectAudioForCustomizedVoiceJobResponse
func (client *Client) DetectAudioForCustomizedVoiceJob(request *DetectAudioForCustomizedVoiceJobRequest) (_result *DetectAudioForCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectAudioForCustomizedVoiceJobResponse{}
	_body, _err := client.DetectAudioForCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a search index. After you delete a search index, the existing index data is cleared and index-based analysis, storage, and query are not supported for subsequent media assets.
//
// @param request - DropSearchIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropSearchIndexResponse
func (client *Client) DropSearchIndexWithOptions(request *DropSearchIndexRequest, runtime *dara.RuntimeOptions) (_result *DropSearchIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexType) {
		query["IndexType"] = request.IndexType
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropSearchIndex"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DropSearchIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a search index. After you delete a search index, the existing index data is cleared and index-based analysis, storage, and query are not supported for subsequent media assets.
//
// @param request - DropSearchIndexRequest
//
// @return DropSearchIndexResponse
func (client *Client) DropSearchIndex(request *DropSearchIndexRequest) (_result *DropSearchIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DropSearchIndexResponse{}
	_body, _err := client.DropSearchIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a search library and all media assets in the library.
//
// @param request - DropSearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DropSearchLibResponse
func (client *Client) DropSearchLibWithOptions(request *DropSearchLibRequest, runtime *dara.RuntimeOptions) (_result *DropSearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DropSearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DropSearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a search library and all media assets in the library.
//
// @param request - DropSearchLibRequest
//
// @return DropSearchLibResponse
func (client *Client) DropSearchLib(request *DropSearchLibRequest) (_result *DropSearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DropSearchLibResponse{}
	_body, _err := client.DropSearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 禁用某个MediaConnect实例的某个输出
//
// @param request - ForbidMediaConnectFlowOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForbidMediaConnectFlowOutputResponse
func (client *Client) ForbidMediaConnectFlowOutputWithOptions(request *ForbidMediaConnectFlowOutputRequest, runtime *dara.RuntimeOptions) (_result *ForbidMediaConnectFlowOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForbidMediaConnectFlowOutput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForbidMediaConnectFlowOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 禁用某个MediaConnect实例的某个输出
//
// @param request - ForbidMediaConnectFlowOutputRequest
//
// @return ForbidMediaConnectFlowOutputResponse
func (client *Client) ForbidMediaConnectFlowOutput(request *ForbidMediaConnectFlowOutputRequest) (_result *ForbidMediaConnectFlowOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForbidMediaConnectFlowOutputResponse{}
	_body, _err := client.ForbidMediaConnectFlowOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 转呼通话到目标电话
//
// @param request - ForwardAIAgentCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ForwardAIAgentCallResponse
func (client *Client) ForwardAIAgentCallWithOptions(request *ForwardAIAgentCallRequest, runtime *dara.RuntimeOptions) (_result *ForwardAIAgentCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ForwardAIAgentCall"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ForwardAIAgentCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 转呼通话到目标电话
//
// @param request - ForwardAIAgentCallRequest
//
// @return ForwardAIAgentCallResponse
func (client *Client) ForwardAIAgentCall(request *ForwardAIAgentCallRequest) (_result *ForwardAIAgentCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ForwardAIAgentCallResponse{}
	_body, _err := client.ForwardAIAgentCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AI agent. This operation returns the channel in which the AI agent resides, the username of the AI agent in the channel, and the token that you can use to join the channel.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to create an AI agent based on the provided ID. You can join the channel based on the returned information and talk to the agent.
//
// **Note:*	- Make sure that the provided AI agent ID is valid and configure optional parameters based on your business requirements.
//
// @param tmpReq - GenerateAIAgentCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAIAgentCallResponse
func (client *Client) GenerateAIAgentCallWithOptions(tmpReq *GenerateAIAgentCallRequest, runtime *dara.RuntimeOptions) (_result *GenerateAIAgentCallResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GenerateAIAgentCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentConfig) {
		request.AgentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentConfig, dara.String("AgentConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChatSyncConfig) {
		request.ChatSyncConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChatSyncConfig, dara.String("ChatSyncConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.AgentConfigShrink) {
		query["AgentConfig"] = request.AgentConfigShrink
	}

	if !dara.IsNil(request.ChatSyncConfigShrink) {
		query["ChatSyncConfig"] = request.ChatSyncConfigShrink
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAIAgentCall"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateAIAgentCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an AI agent. This operation returns the channel in which the AI agent resides, the username of the AI agent in the channel, and the token that you can use to join the channel.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to create an AI agent based on the provided ID. You can join the channel based on the returned information and talk to the agent.
//
// **Note:*	- Make sure that the provided AI agent ID is valid and configure optional parameters based on your business requirements.
//
// @param request - GenerateAIAgentCallRequest
//
// @return GenerateAIAgentCallResponse
func (client *Client) GenerateAIAgentCall(request *GenerateAIAgentCallRequest) (_result *GenerateAIAgentCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateAIAgentCallResponse{}
	_body, _err := client.GenerateAIAgentCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a random Key Management Service (KMS) data key used for HTTP Live Streaming (HLS) encryption and transcoding of videos.
//
// @param request - GenerateKMSDataKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateKMSDataKeyResponse
func (client *Client) GenerateKMSDataKeyWithOptions(runtime *dara.RuntimeOptions) (_result *GenerateKMSDataKeyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateKMSDataKey"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateKMSDataKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a random Key Management Service (KMS) data key used for HTTP Live Streaming (HLS) encryption and transcoding of videos.
//
// @return GenerateKMSDataKeyResponse
func (client *Client) GenerateKMSDataKey() (_result *GenerateKMSDataKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateKMSDataKeyResponse{}
	_body, _err := client.GenerateKMSDataKeyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成直播互动消息所需的token
//
// @param request - GenerateMessageChatTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateMessageChatTokenResponse
func (client *Client) GenerateMessageChatTokenWithOptions(request *GenerateMessageChatTokenRequest, runtime *dara.RuntimeOptions) (_result *GenerateMessageChatTokenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.Expire) {
		query["Expire"] = request.Expire
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateMessageChatToken"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateMessageChatTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成直播互动消息所需的token
//
// @param request - GenerateMessageChatTokenRequest
//
// @return GenerateMessageChatTokenResponse
func (client *Client) GenerateMessageChatToken(request *GenerateMessageChatTokenRequest) (_result *GenerateMessageChatTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateMessageChatTokenResponse{}
	_body, _err := client.GenerateMessageChatTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains details of an ad insertion configuration.
//
// @param request - GetAdInsertionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAdInsertionResponse
func (client *Client) GetAdInsertionWithOptions(request *GetAdInsertionRequest, runtime *dara.RuntimeOptions) (_result *GetAdInsertionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAdInsertion"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAdInsertionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains details of an ad insertion configuration.
//
// @param request - GetAdInsertionRequest
//
// @return GetAdInsertionResponse
func (client *Client) GetAdInsertion(request *GetAdInsertionRequest) (_result *GetAdInsertionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAdInsertionResponse{}
	_body, _err := client.GetAdInsertionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a trained digital human.
//
// @param request - GetAvatarRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAvatarResponse
func (client *Client) GetAvatarWithOptions(request *GetAvatarRequest, runtime *dara.RuntimeOptions) (_result *GetAvatarResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarId) {
		query["AvatarId"] = request.AvatarId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAvatar"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAvatarResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a trained digital human.
//
// @param request - GetAvatarRequest
//
// @return GetAvatarResponse
func (client *Client) GetAvatar(request *GetAvatarRequest) (_result *GetAvatarResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAvatarResponse{}
	_body, _err := client.GetAvatarWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a digital human training job.
//
// @param request - GetAvatarTrainingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAvatarTrainingJobResponse
func (client *Client) GetAvatarTrainingJobWithOptions(request *GetAvatarTrainingJobRequest, runtime *dara.RuntimeOptions) (_result *GetAvatarTrainingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAvatarTrainingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAvatarTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a digital human training job.
//
// @param request - GetAvatarTrainingJobRequest
//
// @return GetAvatarTrainingJobResponse
func (client *Client) GetAvatarTrainingJob(request *GetAvatarTrainingJobRequest) (_result *GetAvatarTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAvatarTrainingJobResponse{}
	_body, _err := client.GetAvatarTrainingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a quick video production job, including the input parameters, job state, and the IDs and URLs of the output media assets. You can call this operation to query only quick video production jobs created within the past year.
//
// @param request - GetBatchMediaProducingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBatchMediaProducingJobResponse
func (client *Client) GetBatchMediaProducingJobWithOptions(request *GetBatchMediaProducingJobRequest, runtime *dara.RuntimeOptions) (_result *GetBatchMediaProducingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBatchMediaProducingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBatchMediaProducingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a quick video production job, including the input parameters, job state, and the IDs and URLs of the output media assets. You can call this operation to query only quick video production jobs created within the past year.
//
// @param request - GetBatchMediaProducingJobRequest
//
// @return GetBatchMediaProducingJobResponse
func (client *Client) GetBatchMediaProducingJob(request *GetBatchMediaProducingJobRequest) (_result *GetBatchMediaProducingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBatchMediaProducingJobResponse{}
	_body, _err := client.GetBatchMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a category and its subcategories.
//
// Description:
//
// You can call this operation to query the information about a category and its subcategories based on the category ID and category type.
//
// @param request - GetCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCategoriesResponse
func (client *Client) GetCategoriesWithOptions(request *GetCategoriesRequest, runtime *dara.RuntimeOptions) (_result *GetCategoriesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCategories"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a category and its subcategories.
//
// Description:
//
// You can call this operation to query the information about a category and its subcategories based on the category ID and category type.
//
// @param request - GetCategoriesRequest
//
// @return GetCategoriesResponse
func (client *Client) GetCategories(request *GetCategoriesRequest) (_result *GetCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCategoriesResponse{}
	_body, _err := client.GetCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a channel in MediaWeaver.
//
// @param request - GetChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChannelResponse
func (client *Client) GetChannelWithOptions(request *GetChannelRequest, runtime *dara.RuntimeOptions) (_result *GetChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a channel in MediaWeaver.
//
// @param request - GetChannelRequest
//
// @return GetChannelResponse
func (client *Client) GetChannel(request *GetChannelRequest) (_result *GetChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChannelResponse{}
	_body, _err := client.GetChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取内容分析搜索配置
//
// @param request - GetContentAnalyzeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContentAnalyzeConfigResponse
func (client *Client) GetContentAnalyzeConfigWithOptions(runtime *dara.RuntimeOptions) (_result *GetContentAnalyzeConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetContentAnalyzeConfig"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContentAnalyzeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取内容分析搜索配置
//
// @return GetContentAnalyzeConfigResponse
func (client *Client) GetContentAnalyzeConfig() (_result *GetContentAnalyzeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetContentAnalyzeConfigResponse{}
	_body, _err := client.GetContentAnalyzeConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a custom template.
//
// Description:
//
// You can call this operation to query the information about a template with the ID specified by the TemplateId parameter. You can also query the information about the default template. If TemplateId is specified, other parameters are ignored and the template whose ID is specified is queried. If TemplateId is not specified, the default template is queried based on other parameters. In this case, Type is required.
//
// Template types:
//
// 1.  1: transcoding template.
//
// 2.  2: snapshot template.
//
// 3.  3: animated image template.
//
// 4.  4\\. image watermark template.
//
// 5.  5: text watermark template.
//
// 6.  6: subtitle template.
//
// 7.  7: AI-assisted content moderation template.
//
// 8.  8: AI-assisted intelligent thumbnail template.
//
// 9.  9: AI-assisted intelligent erasure template.
//
// Subtypes of transcoding templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (AudioTranscode): audio transcoding template.
//
// 3.  3 (Remux): container format conversion template.
//
// 4.  4 (NarrowBandV1): Narrowband HD 1.0 template.
//
// 5.  5 (NarrowBandV2): Narrowband HD 2.0 template.
//
// Subtypes of snapshot templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (Sprite): sprite template.
//
// 3.  3 (WebVtt): WebVTT template.
//
// Subtypes of AI-assisted content moderation templates:
//
// 1.  1 (Video): video moderation template.
//
// 2.  2 (Audio): audio moderation template.
//
// 3.  3 (Image): image moderation template.
//
// Subtypes of AI-assisted intelligent erasure templates:
//
// 1.  1 (VideoDelogo): logo erasure template.
//
// 2.  2 (VideoDetext): subtitle erasure template.
//
// @param request - GetCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomTemplateResponse
func (client *Client) GetCustomTemplateWithOptions(request *GetCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetCustomTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Subtype) {
		query["Subtype"] = request.Subtype
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a custom template.
//
// Description:
//
// You can call this operation to query the information about a template with the ID specified by the TemplateId parameter. You can also query the information about the default template. If TemplateId is specified, other parameters are ignored and the template whose ID is specified is queried. If TemplateId is not specified, the default template is queried based on other parameters. In this case, Type is required.
//
// Template types:
//
// 1.  1: transcoding template.
//
// 2.  2: snapshot template.
//
// 3.  3: animated image template.
//
// 4.  4\\. image watermark template.
//
// 5.  5: text watermark template.
//
// 6.  6: subtitle template.
//
// 7.  7: AI-assisted content moderation template.
//
// 8.  8: AI-assisted intelligent thumbnail template.
//
// 9.  9: AI-assisted intelligent erasure template.
//
// Subtypes of transcoding templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (AudioTranscode): audio transcoding template.
//
// 3.  3 (Remux): container format conversion template.
//
// 4.  4 (NarrowBandV1): Narrowband HD 1.0 template.
//
// 5.  5 (NarrowBandV2): Narrowband HD 2.0 template.
//
// Subtypes of snapshot templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (Sprite): sprite template.
//
// 3.  3 (WebVtt): WebVTT template.
//
// Subtypes of AI-assisted content moderation templates:
//
// 1.  1 (Video): video moderation template.
//
// 2.  2 (Audio): audio moderation template.
//
// 3.  3 (Image): image moderation template.
//
// Subtypes of AI-assisted intelligent erasure templates:
//
// 1.  1 (VideoDelogo): logo erasure template.
//
// 2.  2 (VideoDetext): subtitle erasure template.
//
// @param request - GetCustomTemplateRequest
//
// @return GetCustomTemplateResponse
func (client *Client) GetCustomTemplate(request *GetCustomTemplateRequest) (_result *GetCustomTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomTemplateResponse{}
	_body, _err := client.GetCustomTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a personalized human voice.
//
// @param request - GetCustomizedVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomizedVoiceResponse
func (client *Client) GetCustomizedVoiceWithOptions(request *GetCustomizedVoiceRequest, runtime *dara.RuntimeOptions) (_result *GetCustomizedVoiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VoiceId) {
		query["VoiceId"] = request.VoiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomizedVoice"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomizedVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a personalized human voice.
//
// @param request - GetCustomizedVoiceRequest
//
// @return GetCustomizedVoiceResponse
func (client *Client) GetCustomizedVoice(request *GetCustomizedVoiceRequest) (_result *GetCustomizedVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomizedVoiceResponse{}
	_body, _err := client.GetCustomizedVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a human voice cloning job.
//
// @param request - GetCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomizedVoiceJobResponse
func (client *Client) GetCustomizedVoiceJobWithOptions(request *GetCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *GetCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a human voice cloning job.
//
// @param request - GetCustomizedVoiceJobRequest
//
// @return GetCustomizedVoiceJobResponse
func (client *Client) GetCustomizedVoiceJob(request *GetCustomizedVoiceJobRequest) (_result *GetCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCustomizedVoiceJobResponse{}
	_body, _err := client.GetCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户默认存储地址
//
// @param request - GetDefaultStorageLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultStorageLocationResponse
func (client *Client) GetDefaultStorageLocationWithOptions(runtime *dara.RuntimeOptions) (_result *GetDefaultStorageLocationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetDefaultStorageLocation"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDefaultStorageLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户默认存储地址
//
// @return GetDefaultStorageLocationResponse
func (client *Client) GetDefaultStorageLocation() (_result *GetDefaultStorageLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDefaultStorageLocationResponse{}
	_body, _err := client.GetDefaultStorageLocationWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the text to be read and sample audio for training a personalized human voice.
//
// @param request - GetDemonstrationForCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDemonstrationForCustomizedVoiceJobResponse
func (client *Client) GetDemonstrationForCustomizedVoiceJobWithOptions(request *GetDemonstrationForCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *GetDemonstrationForCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Scenario) {
		query["Scenario"] = request.Scenario
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDemonstrationForCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDemonstrationForCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the text to be read and sample audio for training a personalized human voice.
//
// @param request - GetDemonstrationForCustomizedVoiceJobRequest
//
// @return GetDemonstrationForCustomizedVoiceJobResponse
func (client *Client) GetDemonstrationForCustomizedVoiceJob(request *GetDemonstrationForCustomizedVoiceJobRequest) (_result *GetDemonstrationForCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDemonstrationForCustomizedVoiceJobResponse{}
	_body, _err := client.GetDemonstrationForCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an image animation job.
//
// @param request - GetDynamicImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDynamicImageJobResponse
func (client *Client) GetDynamicImageJobWithOptions(request *GetDynamicImageJobRequest, runtime *dara.RuntimeOptions) (_result *GetDynamicImageJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDynamicImageJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDynamicImageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an image animation job.
//
// @param request - GetDynamicImageJobRequest
//
// @return GetDynamicImageJobResponse
func (client *Client) GetDynamicImageJob(request *GetDynamicImageJobRequest) (_result *GetDynamicImageJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDynamicImageJobResponse{}
	_body, _err := client.GetDynamicImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an online editing project.
//
// @param request - GetEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingProjectResponse
func (client *Client) GetEditingProjectWithOptions(request *GetEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *GetEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RequestSource) {
		query["RequestSource"] = request.RequestSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEditingProject"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an online editing project.
//
// @param request - GetEditingProjectRequest
//
// @return GetEditingProjectResponse
func (client *Client) GetEditingProject(request *GetEditingProjectRequest) (_result *GetEditingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEditingProjectResponse{}
	_body, _err := client.GetEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all materials associated with an online editing project.
//
// @param request - GetEditingProjectMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEditingProjectMaterialsResponse
func (client *Client) GetEditingProjectMaterialsWithOptions(request *GetEditingProjectMaterialsRequest, runtime *dara.RuntimeOptions) (_result *GetEditingProjectMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEditingProjectMaterials"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all materials associated with an online editing project.
//
// @param request - GetEditingProjectMaterialsRequest
//
// @return GetEditingProjectMaterialsResponse
func (client *Client) GetEditingProjectMaterials(request *GetEditingProjectMaterialsRequest) (_result *GetEditingProjectMaterialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEditingProjectMaterialsResponse{}
	_body, _err := client.GetEditingProjectMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries event callback configurations.
//
// @param request - GetEventCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEventCallbackResponse
func (client *Client) GetEventCallbackWithOptions(runtime *dara.RuntimeOptions) (_result *GetEventCallbackResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetEventCallback"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEventCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries event callback configurations.
//
// @return GetEventCallbackResponse
func (client *Client) GetEventCallback() (_result *GetEventCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEventCallbackResponse{}
	_body, _err := client.GetEventCallbackWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询热词词库
//
// @param request - GetHotwordLibraryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotwordLibraryResponse
func (client *Client) GetHotwordLibraryWithOptions(request *GetHotwordLibraryRequest, runtime *dara.RuntimeOptions) (_result *GetHotwordLibraryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HotwordLibraryId) {
		query["HotwordLibraryId"] = request.HotwordLibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotwordLibrary"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotwordLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询热词词库
//
// @param request - GetHotwordLibraryRequest
//
// @return GetHotwordLibraryResponse
func (client *Client) GetHotwordLibrary(request *GetHotwordLibraryRequest) (_result *GetHotwordLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHotwordLibraryResponse{}
	_body, _err := client.GetHotwordLibraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the index file of a live stream. The index file is used to preview an editing project in the console.
//
// @param request - GetLiveEditingIndexFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveEditingIndexFileResponse
func (client *Client) GetLiveEditingIndexFileWithOptions(request *GetLiveEditingIndexFileRequest, runtime *dara.RuntimeOptions) (_result *GetLiveEditingIndexFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StreamName) {
		query["StreamName"] = request.StreamName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveEditingIndexFile"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveEditingIndexFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the index file of a live stream. The index file is used to preview an editing project in the console.
//
// @param request - GetLiveEditingIndexFileRequest
//
// @return GetLiveEditingIndexFileResponse
func (client *Client) GetLiveEditingIndexFile(request *GetLiveEditingIndexFileRequest) (_result *GetLiveEditingIndexFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveEditingIndexFileResponse{}
	_body, _err := client.GetLiveEditingIndexFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a live editing job. The requested information includes the state, timeline, and template of the job, the ID and URL of the output file, and the configurations of the job. You can call this operation to query only live editing jobs created within the past year.
//
// @param request - GetLiveEditingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveEditingJobResponse
func (client *Client) GetLiveEditingJobWithOptions(request *GetLiveEditingJobRequest, runtime *dara.RuntimeOptions) (_result *GetLiveEditingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveEditingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveEditingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a live editing job. The requested information includes the state, timeline, and template of the job, the ID and URL of the output file, and the configurations of the job. You can call this operation to query only live editing jobs created within the past year.
//
// @param request - GetLiveEditingJobRequest
//
// @return GetLiveEditingJobResponse
func (client *Client) GetLiveEditingJob(request *GetLiveEditingJobRequest) (_result *GetLiveEditingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveEditingJobResponse{}
	_body, _err := client.GetLiveEditingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to query the details of a live package channel, including the creation time, description, ingest endpoint, protocol, number of segments, and segment duration.
//
// @param request - GetLivePackageChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivePackageChannelResponse
func (client *Client) GetLivePackageChannelWithOptions(request *GetLivePackageChannelRequest, runtime *dara.RuntimeOptions) (_result *GetLivePackageChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLivePackageChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLivePackageChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to query the details of a live package channel, including the creation time, description, ingest endpoint, protocol, number of segments, and segment duration.
//
// @param request - GetLivePackageChannelRequest
//
// @return GetLivePackageChannelResponse
func (client *Client) GetLivePackageChannel(request *GetLivePackageChannelRequest) (_result *GetLivePackageChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLivePackageChannelResponse{}
	_body, _err := client.GetLivePackageChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a live package channel group by name.
//
// Description:
//
// ## [](#)Usage notes
//
// You can call this API operation to query the details of a specific channel group, including its name, description, origin domain, and creation and last modification timestamps.
//
// @param request - GetLivePackageChannelGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivePackageChannelGroupResponse
func (client *Client) GetLivePackageChannelGroupWithOptions(request *GetLivePackageChannelGroupRequest, runtime *dara.RuntimeOptions) (_result *GetLivePackageChannelGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLivePackageChannelGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLivePackageChannelGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a live package channel group by name.
//
// Description:
//
// ## [](#)Usage notes
//
// You can call this API operation to query the details of a specific channel group, including its name, description, origin domain, and creation and last modification timestamps.
//
// @param request - GetLivePackageChannelGroupRequest
//
// @return GetLivePackageChannelGroupResponse
func (client *Client) GetLivePackageChannelGroup(request *GetLivePackageChannelGroupRequest) (_result *GetLivePackageChannelGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLivePackageChannelGroupResponse{}
	_body, _err := client.GetLivePackageChannelGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries origin endpoints associated with a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// @param request - GetLivePackageOriginEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLivePackageOriginEndpointResponse
func (client *Client) GetLivePackageOriginEndpointWithOptions(request *GetLivePackageOriginEndpointRequest, runtime *dara.RuntimeOptions) (_result *GetLivePackageOriginEndpointResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.EndpointName) {
		query["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLivePackageOriginEndpoint"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLivePackageOriginEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries origin endpoints associated with a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// @param request - GetLivePackageOriginEndpointRequest
//
// @return GetLivePackageOriginEndpointResponse
func (client *Client) GetLivePackageOriginEndpoint(request *GetLivePackageOriginEndpointRequest) (_result *GetLivePackageOriginEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLivePackageOriginEndpointResponse{}
	_body, _err := client.GetLivePackageOriginEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream recording job.
//
// @param request - GetLiveRecordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveRecordJobResponse
func (client *Client) GetLiveRecordJobWithOptions(request *GetLiveRecordJobRequest, runtime *dara.RuntimeOptions) (_result *GetLiveRecordJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveRecordJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveRecordJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream recording job.
//
// @param request - GetLiveRecordJobRequest
//
// @return GetLiveRecordJobResponse
func (client *Client) GetLiveRecordJob(request *GetLiveRecordJobRequest) (_result *GetLiveRecordJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveRecordJobResponse{}
	_body, _err := client.GetLiveRecordJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream recording template or a snapshot of the template.
//
// @param request - GetLiveRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveRecordTemplateResponse
func (client *Client) GetLiveRecordTemplateWithOptions(request *GetLiveRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetLiveRecordTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveRecordTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream recording template or a snapshot of the template.
//
// @param request - GetLiveRecordTemplateRequest
//
// @return GetLiveRecordTemplateResponse
func (client *Client) GetLiveRecordTemplate(request *GetLiveRecordTemplateRequest) (_result *GetLiveRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveRecordTemplateResponse{}
	_body, _err := client.GetLiveRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information a live stream snapshot job.
//
// @param request - GetLiveSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveSnapshotJobResponse
func (client *Client) GetLiveSnapshotJobWithOptions(request *GetLiveSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *GetLiveSnapshotJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveSnapshotJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information a live stream snapshot job.
//
// @param request - GetLiveSnapshotJobRequest
//
// @return GetLiveSnapshotJobResponse
func (client *Client) GetLiveSnapshotJob(request *GetLiveSnapshotJobRequest) (_result *GetLiveSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveSnapshotJobResponse{}
	_body, _err := client.GetLiveSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream snapshot template.
//
// @param request - GetLiveSnapshotTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveSnapshotTemplateResponse
func (client *Client) GetLiveSnapshotTemplateWithOptions(request *GetLiveSnapshotTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetLiveSnapshotTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveSnapshotTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveSnapshotTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream snapshot template.
//
// @param request - GetLiveSnapshotTemplateRequest
//
// @return GetLiveSnapshotTemplateResponse
func (client *Client) GetLiveSnapshotTemplate(request *GetLiveSnapshotTemplateRequest) (_result *GetLiveSnapshotTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveSnapshotTemplateResponse{}
	_body, _err := client.GetLiveSnapshotTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream transcoding job.
//
// @param request - GetLiveTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveTranscodeJobResponse
func (client *Client) GetLiveTranscodeJobWithOptions(request *GetLiveTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *GetLiveTranscodeJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveTranscodeJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveTranscodeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a live stream transcoding job.
//
// @param request - GetLiveTranscodeJobRequest
//
// @return GetLiveTranscodeJobResponse
func (client *Client) GetLiveTranscodeJob(request *GetLiveTranscodeJobRequest) (_result *GetLiveTranscodeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveTranscodeJobResponse{}
	_body, _err := client.GetLiveTranscodeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information a live stream transcoding template.
//
// @param request - GetLiveTranscodeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLiveTranscodeTemplateResponse
func (client *Client) GetLiveTranscodeTemplateWithOptions(request *GetLiveTranscodeTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetLiveTranscodeTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLiveTranscodeTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLiveTranscodeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information a live stream transcoding template.
//
// @param request - GetLiveTranscodeTemplateRequest
//
// @return GetLiveTranscodeTemplateResponse
func (client *Client) GetLiveTranscodeTemplate(request *GetLiveTranscodeTemplateRequest) (_result *GetLiveTranscodeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLiveTranscodeTemplateResponse{}
	_body, _err := client.GetLiveTranscodeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取MediaConnect可用的区域
//
// @param request - GetMediaConnectAvailableRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaConnectAvailableRegionResponse
func (client *Client) GetMediaConnectAvailableRegionWithOptions(runtime *dara.RuntimeOptions) (_result *GetMediaConnectAvailableRegionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaConnectAvailableRegion"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaConnectAvailableRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取MediaConnect可用的区域
//
// @return GetMediaConnectAvailableRegionResponse
func (client *Client) GetMediaConnectAvailableRegion() (_result *GetMediaConnectAvailableRegionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaConnectAvailableRegionResponse{}
	_body, _err := client.GetMediaConnectAvailableRegionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains information about a specific MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- The returned StartTime is valid only when the flow is in the online state.
//
// @param request - GetMediaConnectFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaConnectFlowResponse
func (client *Client) GetMediaConnectFlowWithOptions(request *GetMediaConnectFlowRequest, runtime *dara.RuntimeOptions) (_result *GetMediaConnectFlowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaConnectFlow"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaConnectFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about a specific MediaConnect flow.
//
// Description:
//
//	  When the specified flow ID is not available, an error code is returned.
//
//		- The returned StartTime is valid only when the flow is in the online state.
//
// @param request - GetMediaConnectFlowRequest
//
// @return GetMediaConnectFlowResponse
func (client *Client) GetMediaConnectFlow(request *GetMediaConnectFlowRequest) (_result *GetMediaConnectFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaConnectFlowResponse{}
	_body, _err := client.GetMediaConnectFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取某个MediaConnect实例下的所有output的名字
//
// @param request - GetMediaConnectFlowAllOutputNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaConnectFlowAllOutputNameResponse
func (client *Client) GetMediaConnectFlowAllOutputNameWithOptions(request *GetMediaConnectFlowAllOutputNameRequest, runtime *dara.RuntimeOptions) (_result *GetMediaConnectFlowAllOutputNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaConnectFlowAllOutputName"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaConnectFlowAllOutputNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取某个MediaConnect实例下的所有output的名字
//
// @param request - GetMediaConnectFlowAllOutputNameRequest
//
// @return GetMediaConnectFlowAllOutputNameResponse
func (client *Client) GetMediaConnectFlowAllOutputName(request *GetMediaConnectFlowAllOutputNameRequest) (_result *GetMediaConnectFlowAllOutputNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaConnectFlowAllOutputNameResponse{}
	_body, _err := client.GetMediaConnectFlowAllOutputNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains information about the source of a MediaConnect flow.
//
// Description:
//
//	When the specified flow ID is not available, an error code is returned.
//
// @param request - GetMediaConnectFlowInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaConnectFlowInputResponse
func (client *Client) GetMediaConnectFlowInputWithOptions(request *GetMediaConnectFlowInputRequest, runtime *dara.RuntimeOptions) (_result *GetMediaConnectFlowInputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaConnectFlowInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaConnectFlowInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about the source of a MediaConnect flow.
//
// Description:
//
//	When the specified flow ID is not available, an error code is returned.
//
// @param request - GetMediaConnectFlowInputRequest
//
// @return GetMediaConnectFlowInputResponse
func (client *Client) GetMediaConnectFlowInput(request *GetMediaConnectFlowInputRequest) (_result *GetMediaConnectFlowInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaConnectFlowInputResponse{}
	_body, _err := client.GetMediaConnectFlowInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains information about an output of a MediaConnect flow.
//
// Description:
//
//	When the specified flow ID is not available, an error code is returned.
//
// @param request - GetMediaConnectFlowOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaConnectFlowOutputResponse
func (client *Client) GetMediaConnectFlowOutputWithOptions(request *GetMediaConnectFlowOutputRequest, runtime *dara.RuntimeOptions) (_result *GetMediaConnectFlowOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaConnectFlowOutput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaConnectFlowOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains information about an output of a MediaConnect flow.
//
// Description:
//
//	When the specified flow ID is not available, an error code is returned.
//
// @param request - GetMediaConnectFlowOutputRequest
//
// @return GetMediaConnectFlowOutputResponse
func (client *Client) GetMediaConnectFlowOutput(request *GetMediaConnectFlowOutputRequest) (_result *GetMediaConnectFlowOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaConnectFlowOutputResponse{}
	_body, _err := client.GetMediaConnectFlowOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a transcoding task.
//
// @param request - GetMediaConvertJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaConvertJobResponse
func (client *Client) GetMediaConvertJobWithOptions(request *GetMediaConvertJobRequest, runtime *dara.RuntimeOptions) (_result *GetMediaConvertJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaConvertJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaConvertJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a transcoding task.
//
// @param request - GetMediaConvertJobRequest
//
// @return GetMediaConvertJobResponse
func (client *Client) GetMediaConvertJob(request *GetMediaConvertJobRequest) (_result *GetMediaConvertJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaConvertJobResponse{}
	_body, _err := client.GetMediaConvertJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about a media asset based on the ID of the media asset in Intelligent Media Services (IMS) or the input URL of the media asset.
//
// Description:
//
// If the MediaId parameter is specified, the MediaId parameter is preferentially used for the query. If the MediaId parameter is left empty, the InputURL parameter must be specified.
//
// @param request - GetMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaInfoResponse
func (client *Client) GetMediaInfoWithOptions(request *GetMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMediaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.OutputType) {
		query["OutputType"] = request.OutputType
	}

	if !dara.IsNil(request.ReturnDetailedInfo) {
		query["ReturnDetailedInfo"] = request.ReturnDetailedInfo
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about a media asset based on the ID of the media asset in Intelligent Media Services (IMS) or the input URL of the media asset.
//
// Description:
//
// If the MediaId parameter is specified, the MediaId parameter is preferentially used for the query. If the MediaId parameter is left empty, the InputURL parameter must be specified.
//
// @param request - GetMediaInfoRequest
//
// @return GetMediaInfoResponse
func (client *Client) GetMediaInfo(request *GetMediaInfoRequest) (_result *GetMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaInfoResponse{}
	_body, _err := client.GetMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a media information analysis job.
//
// @param request - GetMediaInfoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaInfoJobResponse
func (client *Client) GetMediaInfoJobWithOptions(request *GetMediaInfoJobRequest, runtime *dara.RuntimeOptions) (_result *GetMediaInfoJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaInfoJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaInfoJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a media information analysis job.
//
// @param request - GetMediaInfoJobRequest
//
// @return GetMediaInfoJobResponse
func (client *Client) GetMediaInfoJob(request *GetMediaInfoJobRequest) (_result *GetMediaInfoJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaInfoJobResponse{}
	_body, _err := client.GetMediaInfoJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a MediaLive channel.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - GetMediaLiveChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaLiveChannelResponse
func (client *Client) GetMediaLiveChannelWithOptions(request *GetMediaLiveChannelRequest, runtime *dara.RuntimeOptions) (_result *GetMediaLiveChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaLiveChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaLiveChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a MediaLive channel.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - GetMediaLiveChannelRequest
//
// @return GetMediaLiveChannelResponse
func (client *Client) GetMediaLiveChannel(request *GetMediaLiveChannelRequest) (_result *GetMediaLiveChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaLiveChannelResponse{}
	_body, _err := client.GetMediaLiveChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a MediaLive input.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - GetMediaLiveInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaLiveInputResponse
func (client *Client) GetMediaLiveInputWithOptions(request *GetMediaLiveInputRequest, runtime *dara.RuntimeOptions) (_result *GetMediaLiveInputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InputId) {
		body["InputId"] = request.InputId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaLiveInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaLiveInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a MediaLive input.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - GetMediaLiveInputRequest
//
// @return GetMediaLiveInputResponse
func (client *Client) GetMediaLiveInput(request *GetMediaLiveInputRequest) (_result *GetMediaLiveInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaLiveInputResponse{}
	_body, _err := client.GetMediaLiveInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a security group in MediaLive.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - GetMediaLiveInputSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaLiveInputSecurityGroupResponse
func (client *Client) GetMediaLiveInputSecurityGroupWithOptions(request *GetMediaLiveInputSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *GetMediaLiveInputSecurityGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaLiveInputSecurityGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a security group in MediaLive.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - GetMediaLiveInputSecurityGroupRequest
//
// @return GetMediaLiveInputSecurityGroupResponse
func (client *Client) GetMediaLiveInputSecurityGroup(request *GetMediaLiveInputSecurityGroupRequest) (_result *GetMediaLiveInputSecurityGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.GetMediaLiveInputSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about marks based on mark IDs.
//
// @param request - GetMediaMarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaMarksResponse
func (client *Client) GetMediaMarksWithOptions(request *GetMediaMarksRequest, runtime *dara.RuntimeOptions) (_result *GetMediaMarksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaMarkIds) {
		query["MediaMarkIds"] = request.MediaMarkIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaMarks"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaMarksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about marks based on mark IDs.
//
// @param request - GetMediaMarksRequest
//
// @return GetMediaMarksResponse
func (client *Client) GetMediaMarks(request *GetMediaMarksRequest) (_result *GetMediaMarksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaMarksResponse{}
	_body, _err := client.GetMediaMarksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a media editing and production job. The requested information includes the state, timeline, template, and data of the job. You can call this operation to query only media editing and production jobs created within the past year.
//
// @param request - GetMediaProducingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMediaProducingJobResponse
func (client *Client) GetMediaProducingJobWithOptions(request *GetMediaProducingJobRequest, runtime *dara.RuntimeOptions) (_result *GetMediaProducingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMediaProducingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMediaProducingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a media editing and production job. The requested information includes the state, timeline, template, and data of the job. You can call this operation to query only media editing and production jobs created within the past year.
//
// @param request - GetMediaProducingJobRequest
//
// @return GetMediaProducingJobResponse
func (client *Client) GetMediaProducingJob(request *GetMediaProducingJobRequest) (_result *GetMediaProducingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMediaProducingJobResponse{}
	_body, _err := client.GetMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a packaging job.
//
// @param request - GetPackageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPackageJobResponse
func (client *Client) GetPackageJobWithOptions(request *GetPackageJobRequest, runtime *dara.RuntimeOptions) (_result *GetPackageJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPackageJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPackageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a packaging job.
//
// @param request - GetPackageJobRequest
//
// @return GetPackageJobResponse
func (client *Client) GetPackageJob(request *GetPackageJobRequest) (_result *GetPackageJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPackageJobResponse{}
	_body, _err := client.GetPackageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - GetPipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineResponse
func (client *Client) GetPipelineWithOptions(request *GetPipelineRequest, runtime *dara.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipeline"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - GetPipelineRequest
//
// @return GetPipelineResponse
func (client *Client) GetPipeline(request *GetPipelineRequest) (_result *GetPipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPipelineResponse{}
	_body, _err := client.GetPipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the playback URL of a video or audio file based on its ID.
//
// Description:
//
// You use the ID of a video or audio file to query the playback URL of the file. Then, you can use the playback URL to play the audio or video in ApsaraVideo Player SDK (for URL-based playback) or a third-party player.
//
// @param request - GetPlayInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPlayInfoResponse
func (client *Client) GetPlayInfoWithOptions(request *GetPlayInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPlayInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPlayInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPlayInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the playback URL of a video or audio file based on its ID.
//
// Description:
//
// You use the ID of a video or audio file to query the playback URL of the file. Then, you can use the playback URL to play the audio or video in ApsaraVideo Player SDK (for URL-based playback) or a third-party player.
//
// @param request - GetPlayInfoRequest
//
// @return GetPlayInfoResponse
func (client *Client) GetPlayInfo(request *GetPlayInfoRequest) (_result *GetPlayInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPlayInfoResponse{}
	_body, _err := client.GetPlayInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a program.
//
// @param request - GetProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProgramResponse
func (client *Client) GetProgramWithOptions(request *GetProgramRequest, runtime *dara.RuntimeOptions) (_result *GetProgramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ProgramName) {
		query["ProgramName"] = request.ProgramName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProgram"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProgramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a program.
//
// @param request - GetProgramRequest
//
// @return GetProgramResponse
func (client *Client) GetProgram(request *GetProgramRequest) (_result *GetProgramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProgramResponse{}
	_body, _err := client.GetProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询工程导出任务
//
// @param request - GetProjectExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectExportJobResponse
func (client *Client) GetProjectExportJobWithOptions(request *GetProjectExportJobRequest, runtime *dara.RuntimeOptions) (_result *GetProjectExportJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectExportJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询工程导出任务
//
// @param request - GetProjectExportJobRequest
//
// @return GetProjectExportJobResponse
func (client *Client) GetProjectExportJob(request *GetProjectExportJobRequest) (_result *GetProjectExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectExportJobResponse{}
	_body, _err := client.GetProjectExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取公共媒资内容信息
//
// @param request - GetPublicMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPublicMediaInfoResponse
func (client *Client) GetPublicMediaInfoWithOptions(request *GetPublicMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *GetPublicMediaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPublicMediaInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPublicMediaInfoResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取公共媒资内容信息
//
// @param request - GetPublicMediaInfoRequest
//
// @return GetPublicMediaInfoResponse
func (client *Client) GetPublicMediaInfo(request *GetPublicMediaInfoRequest) (_result *GetPublicMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPublicMediaInfoResponse{}
	_body, _err := client.GetPublicMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an intelligent job and the execution results of the job based the job ID. You can call this operation to query only intelligent jobs created within the past year.
//
// @param request - GetSmartHandleJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSmartHandleJobResponse
func (client *Client) GetSmartHandleJobWithOptions(request *GetSmartHandleJobRequest, runtime *dara.RuntimeOptions) (_result *GetSmartHandleJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSmartHandleJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSmartHandleJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an intelligent job and the execution results of the job based the job ID. You can call this operation to query only intelligent jobs created within the past year.
//
// @param request - GetSmartHandleJobRequest
//
// @return GetSmartHandleJobResponse
func (client *Client) GetSmartHandleJob(request *GetSmartHandleJobRequest) (_result *GetSmartHandleJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSmartHandleJobResponse{}
	_body, _err := client.GetSmartHandleJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a snapshot job.
//
// @param request - GetSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSnapshotJobResponse
func (client *Client) GetSnapshotJobWithOptions(request *GetSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *GetSnapshotJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSnapshotJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a snapshot job.
//
// @param request - GetSnapshotJobRequest
//
// @return GetSnapshotJobResponse
func (client *Client) GetSnapshotJob(request *GetSnapshotJobRequest) (_result *GetSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSnapshotJobResponse{}
	_body, _err := client.GetSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the accessible URLs of the output images of a snapshot job.
//
// @param request - GetSnapshotUrlsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSnapshotUrlsResponse
func (client *Client) GetSnapshotUrlsWithOptions(request *GetSnapshotUrlsRequest, runtime *dara.RuntimeOptions) (_result *GetSnapshotUrlsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSnapshotUrls"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSnapshotUrlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the accessible URLs of the output images of a snapshot job.
//
// @param request - GetSnapshotUrlsRequest
//
// @return GetSnapshotUrlsResponse
func (client *Client) GetSnapshotUrls(request *GetSnapshotUrlsRequest) (_result *GetSnapshotUrlsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSnapshotUrlsResponse{}
	_body, _err := client.GetSnapshotUrlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a source in MediaWeaver.
//
// @param request - GetSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSourceResponse
func (client *Client) GetSourceWithOptions(request *GetSourceRequest, runtime *dara.RuntimeOptions) (_result *GetSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSource"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a source in MediaWeaver.
//
// @param request - GetSourceRequest
//
// @return GetSourceResponse
func (client *Client) GetSource(request *GetSourceRequest) (_result *GetSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSourceResponse{}
	_body, _err := client.GetSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a source location.
//
// @param request - GetSourceLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSourceLocationResponse
func (client *Client) GetSourceLocationWithOptions(request *GetSourceLocationRequest, runtime *dara.RuntimeOptions) (_result *GetSourceLocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSourceLocation"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSourceLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a source location.
//
// @param request - GetSourceLocationRequest
//
// @return GetSourceLocationResponse
func (client *Client) GetSourceLocation(request *GetSourceLocationRequest) (_result *GetSourceLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSourceLocationResponse{}
	_body, _err := client.GetSourceLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取存储地址列表
//
// @param request - GetStorageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStorageListResponse
func (client *Client) GetStorageListWithOptions(request *GetStorageListRequest, runtime *dara.RuntimeOptions) (_result *GetStorageListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStorageList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStorageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取存储地址列表
//
// @param request - GetStorageListRequest
//
// @return GetStorageListResponse
func (client *Client) GetStorageList(request *GetStorageListRequest) (_result *GetStorageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStorageListResponse{}
	_body, _err := client.GetStorageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a system template.
//
// @param request - GetSystemTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSystemTemplateResponse
func (client *Client) GetSystemTemplateWithOptions(request *GetSystemTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetSystemTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSystemTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSystemTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a system template.
//
// @param request - GetSystemTemplateRequest
//
// @return GetSystemTemplateResponse
func (client *Client) GetSystemTemplate(request *GetSystemTemplateRequest) (_result *GetSystemTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSystemTemplateResponse{}
	_body, _err := client.GetSystemTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a template based on the template ID. You can call this operation to query the information about an advanced template if the template is in the Available state.
//
// Description:
//
// A template is an encapsulation of the timeline of a media editing and production job. You can define a common timeline as a template. When you have the same requirements, you need to only specify key parameters and materials to produce videos.
//
//   - For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//   - For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RelatedMediaidFlag) {
		query["RelatedMediaidFlag"] = request.RelatedMediaidFlag
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a template based on the template ID. You can call this operation to query the information about an advanced template if the template is in the Available state.
//
// Description:
//
// A template is an encapsulation of the timeline of a media editing and production job. You can define a common timeline as a template. When you have the same requirements, you need to only specify key parameters and materials to produce videos.
//
//   - For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//   - For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the URLs of materials associated with an advanced template for use by the advanced template editor. The URLs expire in 30 minutes. FileList is an array of materials that you want to query. If you do not specify this parameter, the URLs of all materials are returned. A maximum of 400 URLs can be returned.
//
// @param request - GetTemplateMaterialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateMaterialsResponse
func (client *Client) GetTemplateMaterialsWithOptions(request *GetTemplateMaterialsRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateMaterialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileList) {
		query["FileList"] = request.FileList
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateMaterials"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateMaterialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the URLs of materials associated with an advanced template for use by the advanced template editor. The URLs expire in 30 minutes. FileList is an array of materials that you want to query. If you do not specify this parameter, the URLs of all materials are returned. A maximum of 400 URLs can be returned.
//
// @param request - GetTemplateMaterialsRequest
//
// @return GetTemplateMaterialsResponse
func (client *Client) GetTemplateMaterials(request *GetTemplateMaterialsRequest) (_result *GetTemplateMaterialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateMaterialsResponse{}
	_body, _err := client.GetTemplateMaterialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters for replaceable materials in a template, including the parameter names, default values, and material thumbnails. Only advanced templates are supported.
//
// @param request - GetTemplateParamsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateParamsResponse
func (client *Client) GetTemplateParamsWithOptions(request *GetTemplateParamsRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateParamsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateParams"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters for replaceable materials in a template, including the parameter names, default values, and material thumbnails. Only advanced templates are supported.
//
// @param request - GetTemplateParamsRequest
//
// @return GetTemplateParamsResponse
func (client *Client) GetTemplateParams(request *GetTemplateParamsRequest) (_result *GetTemplateParamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateParamsResponse{}
	_body, _err := client.GetTemplateParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a transcoding job.
//
// @param request - GetTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscodeJobResponse
func (client *Client) GetTranscodeJobWithOptions(request *GetTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *GetTranscodeJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ParentJobId) {
		query["ParentJobId"] = request.ParentJobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranscodeJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranscodeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a transcoding job.
//
// @param request - GetTranscodeJobRequest
//
// @return GetTranscodeJobResponse
func (client *Client) GetTranscodeJob(request *GetTranscodeJobRequest) (_result *GetTranscodeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTranscodeJobResponse{}
	_body, _err := client.GetTranscodeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about URL-based upload jobs.
//
// Description:
//
// You can call this operation to query the information, including the upload status, user data, creation time, and completion time, about URL-based upload jobs based on the returned job IDs or the URLs used during the upload.
//
// If an upload job fails, you can view the error code and error message. If an upload job is successful, you can obtain the video ID.
//
// @param request - GetUrlUploadInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUrlUploadInfosResponse
func (client *Client) GetUrlUploadInfosWithOptions(request *GetUrlUploadInfosRequest, runtime *dara.RuntimeOptions) (_result *GetUrlUploadInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.UploadURLs) {
		query["UploadURLs"] = request.UploadURLs
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUrlUploadInfos"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUrlUploadInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about URL-based upload jobs.
//
// Description:
//
// You can call this operation to query the information, including the upload status, user data, creation time, and completion time, about URL-based upload jobs based on the returned job IDs or the URLs used during the upload.
//
// If an upload job fails, you can view the error code and error message. If an upload job is successful, you can obtain the video ID.
//
// @param request - GetUrlUploadInfosRequest
//
// @return GetUrlUploadInfosResponse
func (client *Client) GetUrlUploadInfos(request *GetUrlUploadInfosRequest) (_result *GetUrlUploadInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUrlUploadInfosResponse{}
	_body, _err := client.GetUrlUploadInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about video and audio files.
//
// Description:
//
// You can call this operation to query information about up to the first 5,000 audio and video files based on the filter condition, such as the status or category ID of the file. We recommend that you set the StartTime and EndTime parameters to narrow down the time range and perform multiple queries to obtain data.
//
// @param request - GetVideoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoListResponse
func (client *Client) GetVideoListWithOptions(request *GetVideoListRequest, runtime *dara.RuntimeOptions) (_result *GetVideoListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about video and audio files.
//
// Description:
//
// You can call this operation to query information about up to the first 5,000 audio and video files based on the filter condition, such as the status or category ID of the file. We recommend that you set the StartTime and EndTime parameters to narrow down the time range and perform multiple queries to obtain data.
//
// @param request - GetVideoListRequest
//
// @return GetVideoListResponse
func (client *Client) GetVideoList(request *GetVideoListRequest) (_result *GetVideoListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoListResponse{}
	_body, _err := client.GetVideoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a VOD packaging asset.
//
// @param request - GetVodPackagingAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVodPackagingAssetResponse
func (client *Client) GetVodPackagingAssetWithOptions(request *GetVodPackagingAssetRequest, runtime *dara.RuntimeOptions) (_result *GetVodPackagingAssetResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetName) {
		query["AssetName"] = request.AssetName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVodPackagingAsset"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVodPackagingAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a VOD packaging asset.
//
// @param request - GetVodPackagingAssetRequest
//
// @return GetVodPackagingAssetResponse
func (client *Client) GetVodPackagingAsset(request *GetVodPackagingAssetRequest) (_result *GetVodPackagingAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVodPackagingAssetResponse{}
	_body, _err := client.GetVodPackagingAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a packaging configuration.
//
// @param request - GetVodPackagingConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVodPackagingConfigurationResponse
func (client *Client) GetVodPackagingConfigurationWithOptions(request *GetVodPackagingConfigurationRequest, runtime *dara.RuntimeOptions) (_result *GetVodPackagingConfigurationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigurationName) {
		query["ConfigurationName"] = request.ConfigurationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVodPackagingConfiguration"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVodPackagingConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a packaging configuration.
//
// @param request - GetVodPackagingConfigurationRequest
//
// @return GetVodPackagingConfigurationResponse
func (client *Client) GetVodPackagingConfiguration(request *GetVodPackagingConfigurationRequest) (_result *GetVodPackagingConfigurationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVodPackagingConfigurationResponse{}
	_body, _err := client.GetVodPackagingConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a packaging group.
//
// @param request - GetVodPackagingGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVodPackagingGroupResponse
func (client *Client) GetVodPackagingGroupWithOptions(request *GetVodPackagingGroupRequest, runtime *dara.RuntimeOptions) (_result *GetVodPackagingGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVodPackagingGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVodPackagingGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a packaging group.
//
// @param request - GetVodPackagingGroupRequest
//
// @return GetVodPackagingGroupResponse
func (client *Client) GetVodPackagingGroup(request *GetVodPackagingGroupRequest) (_result *GetVodPackagingGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVodPackagingGroupResponse{}
	_body, _err := client.GetVodPackagingGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow task by task ID, including the workflow ID and the status and result of the task. You can query only the workflow task data of the last year.
//
// @param request - GetWorkflowTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkflowTaskResponse
func (client *Client) GetWorkflowTaskWithOptions(request *GetWorkflowTaskRequest, runtime *dara.RuntimeOptions) (_result *GetWorkflowTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkflowTask"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkflowTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow task by task ID, including the workflow ID and the status and result of the task. You can query only the workflow task data of the last year.
//
// @param request - GetWorkflowTaskRequest
//
// @return GetWorkflowTaskResponse
func (client *Client) GetWorkflowTask(request *GetWorkflowTaskRequest) (_result *GetWorkflowTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkflowTaskResponse{}
	_body, _err := client.GetWorkflowTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a media asset in a search library. Before you call this operation, you must create a search library.
//
// @param request - InsertMediaToSearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InsertMediaToSearchLibResponse
func (client *Client) InsertMediaToSearchLibWithOptions(request *InsertMediaToSearchLibRequest, runtime *dara.RuntimeOptions) (_result *InsertMediaToSearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ImagesInput) {
		query["ImagesInput"] = request.ImagesInput
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.MsgBody) {
		query["MsgBody"] = request.MsgBody
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InsertMediaToSearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InsertMediaToSearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a media asset in a search library. Before you call this operation, you must create a search library.
//
// @param request - InsertMediaToSearchLibRequest
//
// @return InsertMediaToSearchLibResponse
func (client *Client) InsertMediaToSearchLib(request *InsertMediaToSearchLibRequest) (_result *InsertMediaToSearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InsertMediaToSearchLibResponse{}
	_body, _err := client.InsertMediaToSearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出智能体的对话历史记录。
//
// @param request - ListAIAgentDialoguesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIAgentDialoguesResponse
func (client *Client) ListAIAgentDialoguesWithOptions(request *ListAIAgentDialoguesRequest, runtime *dara.RuntimeOptions) (_result *ListAIAgentDialoguesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
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

	if !dara.IsNil(request.RoundLimit) {
		query["RoundLimit"] = request.RoundLimit
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIAgentDialogues"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIAgentDialoguesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出智能体的对话历史记录。
//
// @param request - ListAIAgentDialoguesRequest
//
// @return ListAIAgentDialoguesResponse
func (client *Client) ListAIAgentDialogues(request *ListAIAgentDialoguesRequest) (_result *ListAIAgentDialoguesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIAgentDialoguesResponse{}
	_body, _err := client.ListAIAgentDialoguesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AI agents.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to query a list of AI agents based on the `AIAgentId`. The optional parameters include `StartTime`, `EndTime`, `PageSize`, and `PageNumber`. The returned result includes the status, runtime configurations, template configurations, custom information, and the URL of call log file for each AI agent.
//
// **Note**:
//
//   - The default value of `PageSize` is 10, and the default value of `PageNumber` is 1.
//
// @param request - ListAIAgentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIAgentInstanceResponse
func (client *Client) ListAIAgentInstanceWithOptions(request *ListAIAgentInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListAIAgentInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
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
		Action:      dara.String("ListAIAgentInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIAgentInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AI agents.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to query a list of AI agents based on the `AIAgentId`. The optional parameters include `StartTime`, `EndTime`, `PageSize`, and `PageNumber`. The returned result includes the status, runtime configurations, template configurations, custom information, and the URL of call log file for each AI agent.
//
// **Note**:
//
//   - The default value of `PageSize` is 10, and the default value of `PageNumber` is 1.
//
// @param request - ListAIAgentInstanceRequest
//
// @return ListAIAgentInstanceResponse
func (client *Client) ListAIAgentInstance(request *ListAIAgentInstanceRequest) (_result *ListAIAgentInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIAgentInstanceResponse{}
	_body, _err := client.ListAIAgentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 罗列用户电话资源接口
//
// @param request - ListAIAgentPhoneNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIAgentPhoneNumberResponse
func (client *Client) ListAIAgentPhoneNumberWithOptions(request *ListAIAgentPhoneNumberRequest, runtime *dara.RuntimeOptions) (_result *ListAIAgentPhoneNumberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIAgentPhoneNumber"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIAgentPhoneNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 罗列用户电话资源接口
//
// @param request - ListAIAgentPhoneNumberRequest
//
// @return ListAIAgentPhoneNumberResponse
func (client *Client) ListAIAgentPhoneNumber(request *ListAIAgentPhoneNumberRequest) (_result *ListAIAgentPhoneNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIAgentPhoneNumberResponse{}
	_body, _err := client.ListAIAgentPhoneNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新回调配置
//
// @param request - ListAIAgentVoiceprintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIAgentVoiceprintsResponse
func (client *Client) ListAIAgentVoiceprintsWithOptions(request *ListAIAgentVoiceprintsRequest, runtime *dara.RuntimeOptions) (_result *ListAIAgentVoiceprintsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VoiceprintId) {
		query["VoiceprintId"] = request.VoiceprintId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIAgentVoiceprints"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIAgentVoiceprintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新回调配置
//
// @param request - ListAIAgentVoiceprintsRequest
//
// @return ListAIAgentVoiceprintsResponse
func (client *Client) ListAIAgentVoiceprints(request *ListAIAgentVoiceprintsRequest) (_result *ListAIAgentVoiceprintsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIAgentVoiceprintsResponse{}
	_body, _err := client.ListAIAgentVoiceprintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains ad insertion configurations.
//
// @param request - ListAdInsertionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAdInsertionsResponse
func (client *Client) ListAdInsertionsWithOptions(request *ListAdInsertionsRequest, runtime *dara.RuntimeOptions) (_result *ListAdInsertionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAdInsertions"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAdInsertionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains ad insertion configurations.
//
// @param request - ListAdInsertionsRequest
//
// @return ListAdInsertionsResponse
func (client *Client) ListAdInsertions(request *ListAdInsertionsRequest) (_result *ListAdInsertionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAdInsertionsResponse{}
	_body, _err := client.ListAdInsertionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists alerts received in MediaWeaver.
//
// @param request - ListAlertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertsResponse
func (client *Client) ListAlertsWithOptions(request *ListAlertsRequest, runtime *dara.RuntimeOptions) (_result *ListAlertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.GmtEnd) {
		query["GmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		query["GmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceArn) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortByModifiedTime) {
		query["SortByModifiedTime"] = request.SortByModifiedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlerts"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists alerts received in MediaWeaver.
//
// @param request - ListAlertsRequest
//
// @return ListAlertsResponse
func (client *Client) ListAlerts(request *ListAlertsRequest) (_result *ListAlertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlertsResponse{}
	_body, _err := client.ListAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags of media assets in the public media library.
//
// @param request - ListAllPublicMediaTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAllPublicMediaTagsResponse
func (client *Client) ListAllPublicMediaTagsWithOptions(request *ListAllPublicMediaTagsRequest, runtime *dara.RuntimeOptions) (_result *ListAllPublicMediaTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAllPublicMediaTags"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAllPublicMediaTagsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tags of media assets in the public media library.
//
// @param request - ListAllPublicMediaTagsRequest
//
// @return ListAllPublicMediaTagsResponse
func (client *Client) ListAllPublicMediaTags(request *ListAllPublicMediaTagsRequest) (_result *ListAllPublicMediaTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAllPublicMediaTagsResponse{}
	_body, _err := client.ListAllPublicMediaTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of digital human training jobs.
//
// @param request - ListAvatarTrainingJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvatarTrainingJobsResponse
func (client *Client) ListAvatarTrainingJobsWithOptions(request *ListAvatarTrainingJobsRequest, runtime *dara.RuntimeOptions) (_result *ListAvatarTrainingJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvatarTrainingJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvatarTrainingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of digital human training jobs.
//
// @param request - ListAvatarTrainingJobsRequest
//
// @return ListAvatarTrainingJobsResponse
func (client *Client) ListAvatarTrainingJobs(request *ListAvatarTrainingJobsRequest) (_result *ListAvatarTrainingJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAvatarTrainingJobsResponse{}
	_body, _err := client.ListAvatarTrainingJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of trained digital humans.
//
// @param request - ListAvatarsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvatarsResponse
func (client *Client) ListAvatarsWithOptions(request *ListAvatarsRequest, runtime *dara.RuntimeOptions) (_result *ListAvatarsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarType) {
		query["AvatarType"] = request.AvatarType
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvatars"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvatarsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of trained digital humans.
//
// @param request - ListAvatarsRequest
//
// @return ListAvatarsResponse
func (client *Client) ListAvatars(request *ListAvatarsRequest) (_result *ListAvatarsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAvatarsResponse{}
	_body, _err := client.ListAvatarsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of quick video production jobs based on conditions such as the job type and state.
//
// @param request - ListBatchMediaProducingJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBatchMediaProducingJobsResponse
func (client *Client) ListBatchMediaProducingJobsWithOptions(request *ListBatchMediaProducingJobsRequest, runtime *dara.RuntimeOptions) (_result *ListBatchMediaProducingJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBatchMediaProducingJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBatchMediaProducingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of quick video production jobs based on conditions such as the job type and state.
//
// @param request - ListBatchMediaProducingJobsRequest
//
// @return ListBatchMediaProducingJobsResponse
func (client *Client) ListBatchMediaProducingJobs(request *ListBatchMediaProducingJobsRequest) (_result *ListBatchMediaProducingJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBatchMediaProducingJobsResponse{}
	_body, _err := client.ListBatchMediaProducingJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists alerts for resources in a MediaWeaver channel.
//
// @param request - ListChannelAlertsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChannelAlertsResponse
func (client *Client) ListChannelAlertsWithOptions(request *ListChannelAlertsRequest, runtime *dara.RuntimeOptions) (_result *ListChannelAlertsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.GmtEnd) {
		query["GmtEnd"] = request.GmtEnd
	}

	if !dara.IsNil(request.GmtStart) {
		query["GmtStart"] = request.GmtStart
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceArn) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !dara.IsNil(request.SortByModifiedTime) {
		query["SortByModifiedTime"] = request.SortByModifiedTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChannelAlerts"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChannelAlertsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists alerts for resources in a MediaWeaver channel.
//
// @param request - ListChannelAlertsRequest
//
// @return ListChannelAlertsResponse
func (client *Client) ListChannelAlerts(request *ListChannelAlertsRequest) (_result *ListChannelAlertsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChannelAlertsResponse{}
	_body, _err := client.ListChannelAlertsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists MediaWeaver channels.
//
// @param request - ListChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChannelsResponse
func (client *Client) ListChannelsWithOptions(request *ListChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListChannelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ChannelTier) {
		query["ChannelTier"] = request.ChannelTier
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PlaybackMode) {
		query["PlaybackMode"] = request.PlaybackMode
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortByModifiedTime) {
		query["SortByModifiedTime"] = request.SortByModifiedTime
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChannels"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists MediaWeaver channels.
//
// @param request - ListChannelsRequest
//
// @return ListChannelsResponse
func (client *Client) ListChannels(request *ListChannelsRequest) (_result *ListChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChannelsResponse{}
	_body, _err := client.ListChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of custom templates.
//
// @param request - ListCustomTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomTemplatesResponse
func (client *Client) ListCustomTemplatesWithOptions(request *ListCustomTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Subtype) {
		query["Subtype"] = request.Subtype
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomTemplates"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom templates.
//
// @param request - ListCustomTemplatesRequest
//
// @return ListCustomTemplatesResponse
func (client *Client) ListCustomTemplates(request *ListCustomTemplatesRequest) (_result *ListCustomTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomTemplatesResponse{}
	_body, _err := client.ListCustomTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of human voice cloning jobs.
//
// @param request - ListCustomizedVoiceJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomizedVoiceJobsResponse
func (client *Client) ListCustomizedVoiceJobsWithOptions(request *ListCustomizedVoiceJobsRequest, runtime *dara.RuntimeOptions) (_result *ListCustomizedVoiceJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomizedVoiceJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomizedVoiceJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of human voice cloning jobs.
//
// @param request - ListCustomizedVoiceJobsRequest
//
// @return ListCustomizedVoiceJobsResponse
func (client *Client) ListCustomizedVoiceJobs(request *ListCustomizedVoiceJobsRequest) (_result *ListCustomizedVoiceJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomizedVoiceJobsResponse{}
	_body, _err := client.ListCustomizedVoiceJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of personalized human voices.
//
// @param request - ListCustomizedVoicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomizedVoicesResponse
func (client *Client) ListCustomizedVoicesWithOptions(request *ListCustomizedVoicesRequest, runtime *dara.RuntimeOptions) (_result *ListCustomizedVoicesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomizedVoices"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomizedVoicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of personalized human voices.
//
// @param request - ListCustomizedVoicesRequest
//
// @return ListCustomizedVoicesResponse
func (client *Client) ListCustomizedVoices(request *ListCustomizedVoicesRequest) (_result *ListCustomizedVoicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCustomizedVoicesResponse{}
	_body, _err := client.ListCustomizedVoicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of media fingerprint libraries.
//
// @param request - ListDNADBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDNADBResponse
func (client *Client) ListDNADBWithOptions(request *ListDNADBRequest, runtime *dara.RuntimeOptions) (_result *ListDNADBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBIds) {
		query["DBIds"] = request.DBIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDNADB"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDNADBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of media fingerprint libraries.
//
// @param request - ListDNADBRequest
//
// @return ListDNADBResponse
func (client *Client) ListDNADB(request *ListDNADBRequest) (_result *ListDNADBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDNADBResponse{}
	_body, _err := client.ListDNADBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of files in a media fingerprint library.
//
// Description:
//
// You can call this operation to query files in a media fingerprint library based on the library ID. The queried results can be paginated.
//
// @param request - ListDNAFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDNAFilesResponse
func (client *Client) ListDNAFilesWithOptions(request *ListDNAFilesRequest, runtime *dara.RuntimeOptions) (_result *ListDNAFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBId) {
		query["DBId"] = request.DBId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDNAFiles"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDNAFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files in a media fingerprint library.
//
// Description:
//
// You can call this operation to query files in a media fingerprint library based on the library ID. The queried results can be paginated.
//
// @param request - ListDNAFilesRequest
//
// @return ListDNAFilesResponse
func (client *Client) ListDNAFiles(request *ListDNAFilesRequest) (_result *ListDNAFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDNAFilesResponse{}
	_body, _err := client.ListDNAFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of image animation jobs.
//
// @param request - ListDynamicImageJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicImageJobsResponse
func (client *Client) ListDynamicImageJobsWithOptions(request *ListDynamicImageJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDynamicImageJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDynamicImageJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDynamicImageJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of image animation jobs.
//
// @param request - ListDynamicImageJobsRequest
//
// @return ListDynamicImageJobsResponse
func (client *Client) ListDynamicImageJobs(request *ListDynamicImageJobsRequest) (_result *ListDynamicImageJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDynamicImageJobsResponse{}
	_body, _err := client.ListDynamicImageJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of projects that meet the specified conditions. You can filter projects by project creation time.
//
// @param request - ListEditingProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEditingProjectsResponse
func (client *Client) ListEditingProjectsWithOptions(request *ListEditingProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListEditingProjectsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSource) {
		query["CreateSource"] = request.CreateSource
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectType) {
		query["ProjectType"] = request.ProjectType
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

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEditingProjects"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEditingProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of projects that meet the specified conditions. You can filter projects by project creation time.
//
// @param request - ListEditingProjectsRequest
//
// @return ListEditingProjectsResponse
func (client *Client) ListEditingProjects(request *ListEditingProjectsRequest) (_result *ListEditingProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEditingProjectsResponse{}
	_body, _err := client.ListEditingProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询所有热词库的列表
//
// @param request - ListHotwordLibrariesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotwordLibrariesResponse
func (client *Client) ListHotwordLibrariesWithOptions(request *ListHotwordLibrariesRequest, runtime *dara.RuntimeOptions) (_result *ListHotwordLibrariesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UsageScenario) {
		query["UsageScenario"] = request.UsageScenario
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotwordLibraries"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotwordLibrariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询所有热词库的列表
//
// @param request - ListHotwordLibrariesRequest
//
// @return ListHotwordLibrariesResponse
func (client *Client) ListHotwordLibraries(request *ListHotwordLibrariesRequest) (_result *ListHotwordLibrariesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHotwordLibrariesResponse{}
	_body, _err := client.ListHotwordLibrariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries live package channel groups by page. Fuzzy search by name or description and sorting are supported.
//
// Description:
//
// ## [](#)Usage notes
//
// @param request - ListLivePackageChannelGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivePackageChannelGroupsResponse
func (client *Client) ListLivePackageChannelGroupsWithOptions(request *ListLivePackageChannelGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListLivePackageChannelGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLivePackageChannelGroups"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLivePackageChannelGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries live package channel groups by page. Fuzzy search by name or description and sorting are supported.
//
// Description:
//
// ## [](#)Usage notes
//
// @param request - ListLivePackageChannelGroupsRequest
//
// @return ListLivePackageChannelGroupsResponse
func (client *Client) ListLivePackageChannelGroups(request *ListLivePackageChannelGroupsRequest) (_result *ListLivePackageChannelGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLivePackageChannelGroupsResponse{}
	_body, _err := client.ListLivePackageChannelGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries live package channels by channel group and keyword. Paging and sorting are supported.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to query live package channels by **GroupName*	- and **Keyword**. Keyword is optional. You can sort the channels by creation time in ascending or descending order and paginate the results. This facilitates the management of channels and retrieval of channel information.
//
//   - **GroupName*	- is required to specify the channel group to which the channel belongs.
//
//   - **Keyword*	- supports fuzzy match of channel names or descriptions, which helps quickly filter desired channels.
//
//   - **PageNo*	- and **PageSize*	- can help control the paging of returned results to facilitate batch processing of data.
//
//   - **SortBy*	- allows you to customize how the results are sorted. By default, the results are sorted in descending order.
//
// **RequestId*	- in the response is used for subsequent troubleshooting. **TotalCount*	- indicates the total number of channels that meet the conditions.
//
// @param request - ListLivePackageChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivePackageChannelsResponse
func (client *Client) ListLivePackageChannelsWithOptions(request *ListLivePackageChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListLivePackageChannelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLivePackageChannels"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLivePackageChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries live package channels by channel group and keyword. Paging and sorting are supported.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to query live package channels by **GroupName*	- and **Keyword**. Keyword is optional. You can sort the channels by creation time in ascending or descending order and paginate the results. This facilitates the management of channels and retrieval of channel information.
//
//   - **GroupName*	- is required to specify the channel group to which the channel belongs.
//
//   - **Keyword*	- supports fuzzy match of channel names or descriptions, which helps quickly filter desired channels.
//
//   - **PageNo*	- and **PageSize*	- can help control the paging of returned results to facilitate batch processing of data.
//
//   - **SortBy*	- allows you to customize how the results are sorted. By default, the results are sorted in descending order.
//
// **RequestId*	- in the response is used for subsequent troubleshooting. **TotalCount*	- indicates the total number of channels that meet the conditions.
//
// @param request - ListLivePackageChannelsRequest
//
// @return ListLivePackageChannelsResponse
func (client *Client) ListLivePackageChannels(request *ListLivePackageChannelsRequest) (_result *ListLivePackageChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLivePackageChannelsResponse{}
	_body, _err := client.ListLivePackageChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries origin endpoints by channel group and channel name. Paging and sorting are supported.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to query origin endpoints associated with a live package channel. The results include detailed configurations about the origin endpoints, such as access URL, protocol, and security policies. Paging and sorting by creation time are supported.
//
// @param request - ListLivePackageOriginEndpointsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLivePackageOriginEndpointsResponse
func (client *Client) ListLivePackageOriginEndpointsWithOptions(request *ListLivePackageOriginEndpointsRequest, runtime *dara.RuntimeOptions) (_result *ListLivePackageOriginEndpointsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLivePackageOriginEndpoints"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLivePackageOriginEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries origin endpoints by channel group and channel name. Paging and sorting are supported.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to query origin endpoints associated with a live package channel. The results include detailed configurations about the origin endpoints, such as access URL, protocol, and security policies. Paging and sorting by creation time are supported.
//
// @param request - ListLivePackageOriginEndpointsRequest
//
// @return ListLivePackageOriginEndpointsResponse
func (client *Client) ListLivePackageOriginEndpoints(request *ListLivePackageOriginEndpointsRequest) (_result *ListLivePackageOriginEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLivePackageOriginEndpointsResponse{}
	_body, _err := client.ListLivePackageOriginEndpointsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all recording index files in the specified period of time.
//
// @param request - ListLiveRecordFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRecordFilesResponse
func (client *Client) ListLiveRecordFilesWithOptions(request *ListLiveRecordFilesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRecordFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveRecordFiles"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRecordFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all recording index files in the specified period of time.
//
// @param request - ListLiveRecordFilesRequest
//
// @return ListLiveRecordFilesResponse
func (client *Client) ListLiveRecordFiles(request *ListLiveRecordFilesRequest) (_result *ListLiveRecordFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveRecordFilesResponse{}
	_body, _err := client.ListLiveRecordFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream recording jobs by page.
//
// @param request - ListLiveRecordJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRecordJobsResponse
func (client *Client) ListLiveRecordJobsWithOptions(request *ListLiveRecordJobsRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRecordJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveRecordJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRecordJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream recording jobs by page.
//
// @param request - ListLiveRecordJobsRequest
//
// @return ListLiveRecordJobsResponse
func (client *Client) ListLiveRecordJobs(request *ListLiveRecordJobsRequest) (_result *ListLiveRecordJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveRecordJobsResponse{}
	_body, _err := client.ListLiveRecordJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream recording templates.
//
// @param request - ListLiveRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveRecordTemplatesResponse
func (client *Client) ListLiveRecordTemplatesWithOptions(request *ListLiveRecordTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveRecordTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveRecordTemplates"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveRecordTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream recording templates.
//
// @param request - ListLiveRecordTemplatesRequest
//
// @return ListLiveRecordTemplatesResponse
func (client *Client) ListLiveRecordTemplates(request *ListLiveRecordTemplatesRequest) (_result *ListLiveRecordTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveRecordTemplatesResponse{}
	_body, _err := client.ListLiveRecordTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream snapshot files by page.
//
// @param request - ListLiveSnapshotFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveSnapshotFilesResponse
func (client *Client) ListLiveSnapshotFilesWithOptions(request *ListLiveSnapshotFilesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveSnapshotFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveSnapshotFiles"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveSnapshotFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream snapshot files by page.
//
// @param request - ListLiveSnapshotFilesRequest
//
// @return ListLiveSnapshotFilesResponse
func (client *Client) ListLiveSnapshotFiles(request *ListLiveSnapshotFilesRequest) (_result *ListLiveSnapshotFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveSnapshotFilesResponse{}
	_body, _err := client.ListLiveSnapshotFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream snapshot jobs by page.
//
// @param request - ListLiveSnapshotJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveSnapshotJobsResponse
func (client *Client) ListLiveSnapshotJobsWithOptions(request *ListLiveSnapshotJobsRequest, runtime *dara.RuntimeOptions) (_result *ListLiveSnapshotJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveSnapshotJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveSnapshotJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream snapshot jobs by page.
//
// @param request - ListLiveSnapshotJobsRequest
//
// @return ListLiveSnapshotJobsResponse
func (client *Client) ListLiveSnapshotJobs(request *ListLiveSnapshotJobsRequest) (_result *ListLiveSnapshotJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveSnapshotJobsResponse{}
	_body, _err := client.ListLiveSnapshotJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream snapshot templates by page.
//
// @param request - ListLiveSnapshotTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveSnapshotTemplatesResponse
func (client *Client) ListLiveSnapshotTemplatesWithOptions(request *ListLiveSnapshotTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveSnapshotTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveSnapshotTemplates"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveSnapshotTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream snapshot templates by page.
//
// @param request - ListLiveSnapshotTemplatesRequest
//
// @return ListLiveSnapshotTemplatesResponse
func (client *Client) ListLiveSnapshotTemplates(request *ListLiveSnapshotTemplatesRequest) (_result *ListLiveSnapshotTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveSnapshotTemplatesResponse{}
	_body, _err := client.ListLiveSnapshotTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream transcoding jobs.
//
// @param request - ListLiveTranscodeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveTranscodeJobsResponse
func (client *Client) ListLiveTranscodeJobsWithOptions(request *ListLiveTranscodeJobsRequest, runtime *dara.RuntimeOptions) (_result *ListLiveTranscodeJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.StartMode) {
		query["StartMode"] = request.StartMode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveTranscodeJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveTranscodeJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream transcoding jobs.
//
// @param request - ListLiveTranscodeJobsRequest
//
// @return ListLiveTranscodeJobsResponse
func (client *Client) ListLiveTranscodeJobs(request *ListLiveTranscodeJobsRequest) (_result *ListLiveTranscodeJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveTranscodeJobsResponse{}
	_body, _err := client.ListLiveTranscodeJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of live stream transcoding templates.
//
// @param request - ListLiveTranscodeTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLiveTranscodeTemplatesResponse
func (client *Client) ListLiveTranscodeTemplatesWithOptions(request *ListLiveTranscodeTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListLiveTranscodeTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.VideoCodec) {
		query["VideoCodec"] = request.VideoCodec
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLiveTranscodeTemplates"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLiveTranscodeTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of live stream transcoding templates.
//
// @param request - ListLiveTranscodeTemplatesRequest
//
// @return ListLiveTranscodeTemplatesResponse
func (client *Client) ListLiveTranscodeTemplates(request *ListLiveTranscodeTemplatesRequest) (_result *ListLiveTranscodeTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLiveTranscodeTemplatesResponse{}
	_body, _err := client.ListLiveTranscodeTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information of all media assets that meet the specified conditions.
//
// Description:
//
// If includeFileBasicInfo is set to true, the basic information, such as the duration and file size, of the source file is also returned. At most the first 100 entries that meet the specified conditions are returned. All media assets must exactly match all non-empty fields. The fields that support exact match include MediaType, Source, BusinessType, Category, and Status. If all information cannot be returned at a time, you can use NextToken to initiate a request to retrieve a new page of results.
//
// @param request - ListMediaBasicInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaBasicInfosResponse
func (client *Client) ListMediaBasicInfosWithOptions(request *ListMediaBasicInfosRequest, runtime *dara.RuntimeOptions) (_result *ListMediaBasicInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthTimeout) {
		query["AuthTimeout"] = request.AuthTimeout
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IncludeFileBasicInfo) {
		query["IncludeFileBasicInfo"] = request.IncludeFileBasicInfo
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaBasicInfos"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaBasicInfosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information of all media assets that meet the specified conditions.
//
// Description:
//
// If includeFileBasicInfo is set to true, the basic information, such as the duration and file size, of the source file is also returned. At most the first 100 entries that meet the specified conditions are returned. All media assets must exactly match all non-empty fields. The fields that support exact match include MediaType, Source, BusinessType, Category, and Status. If all information cannot be returned at a time, you can use NextToken to initiate a request to retrieve a new page of results.
//
// @param request - ListMediaBasicInfosRequest
//
// @return ListMediaBasicInfosResponse
func (client *Client) ListMediaBasicInfos(request *ListMediaBasicInfosRequest) (_result *ListMediaBasicInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaBasicInfosResponse{}
	_body, _err := client.ListMediaBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 MediaConvertJob 列表
//
// @param request - ListMediaConvertJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaConvertJobsResponse
func (client *Client) ListMediaConvertJobsWithOptions(request *ListMediaConvertJobsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaConvertJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaConvertJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaConvertJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 MediaConvertJob 列表
//
// @param request - ListMediaConvertJobsRequest
//
// @return ListMediaConvertJobsResponse
func (client *Client) ListMediaConvertJobs(request *ListMediaConvertJobsRequest) (_result *ListMediaConvertJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaConvertJobsResponse{}
	_body, _err := client.ListMediaConvertJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of media information analysis jobs.
//
// @param request - ListMediaInfoJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaInfoJobsResponse
func (client *Client) ListMediaInfoJobsWithOptions(request *ListMediaInfoJobsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaInfoJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaInfoJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaInfoJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of media information analysis jobs.
//
// @param request - ListMediaInfoJobsRequest
//
// @return ListMediaInfoJobsResponse
func (client *Client) ListMediaInfoJobs(request *ListMediaInfoJobsRequest) (_result *ListMediaInfoJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaInfoJobsResponse{}
	_body, _err := client.ListMediaInfoJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries MediaLive channels.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListMediaLiveChannelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaLiveChannelsResponse
func (client *Client) ListMediaLiveChannelsWithOptions(request *ListMediaLiveChannelsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaLiveChannelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.States) {
		body["States"] = request.States
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaLiveChannels"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaLiveChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries MediaLive channels.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListMediaLiveChannelsRequest
//
// @return ListMediaLiveChannelsResponse
func (client *Client) ListMediaLiveChannels(request *ListMediaLiveChannelsRequest) (_result *ListMediaLiveChannelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaLiveChannelsResponse{}
	_body, _err := client.ListMediaLiveChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the security groups in MediaLive.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListMediaLiveInputSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaLiveInputSecurityGroupsResponse
func (client *Client) ListMediaLiveInputSecurityGroupsWithOptions(request *ListMediaLiveInputSecurityGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaLiveInputSecurityGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaLiveInputSecurityGroups"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaLiveInputSecurityGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the security groups in MediaLive.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListMediaLiveInputSecurityGroupsRequest
//
// @return ListMediaLiveInputSecurityGroupsResponse
func (client *Client) ListMediaLiveInputSecurityGroups(request *ListMediaLiveInputSecurityGroupsRequest) (_result *ListMediaLiveInputSecurityGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaLiveInputSecurityGroupsResponse{}
	_body, _err := client.ListMediaLiveInputSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries MediaLive inputs.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListMediaLiveInputsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaLiveInputsResponse
func (client *Client) ListMediaLiveInputsWithOptions(request *ListMediaLiveInputsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaLiveInputsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		body["Skip"] = request.Skip
	}

	if !dara.IsNil(request.SortOrder) {
		body["SortOrder"] = request.SortOrder
	}

	if !dara.IsNil(request.Types) {
		body["Types"] = request.Types
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaLiveInputs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaLiveInputsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries MediaLive inputs.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - ListMediaLiveInputsRequest
//
// @return ListMediaLiveInputsResponse
func (client *Client) ListMediaLiveInputs(request *ListMediaLiveInputsRequest) (_result *ListMediaLiveInputsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaLiveInputsResponse{}
	_body, _err := client.ListMediaLiveInputsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of marks of a media asset.
//
// @param request - ListMediaMarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaMarksResponse
func (client *Client) ListMediaMarksWithOptions(request *ListMediaMarksRequest, runtime *dara.RuntimeOptions) (_result *ListMediaMarksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaMarkIds) {
		query["MediaMarkIds"] = request.MediaMarkIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaMarks"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaMarksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of marks of a media asset.
//
// @param request - ListMediaMarksRequest
//
// @return ListMediaMarksResponse
func (client *Client) ListMediaMarks(request *ListMediaMarksRequest) (_result *ListMediaMarksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaMarksResponse{}
	_body, _err := client.ListMediaMarksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of media editing and production jobs that meet the specified conditions. You can query the jobs based on the job state and type.
//
// @param request - ListMediaProducingJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMediaProducingJobsResponse
func (client *Client) ListMediaProducingJobsWithOptions(request *ListMediaProducingJobsRequest, runtime *dara.RuntimeOptions) (_result *ListMediaProducingJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.JobType) {
		query["JobType"] = request.JobType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.MasterJobId) {
		query["MasterJobId"] = request.MasterJobId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMediaProducingJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMediaProducingJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of media editing and production jobs that meet the specified conditions. You can query the jobs based on the job state and type.
//
// @param request - ListMediaProducingJobsRequest
//
// @return ListMediaProducingJobsResponse
func (client *Client) ListMediaProducingJobs(request *ListMediaProducingJobsRequest) (_result *ListMediaProducingJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMediaProducingJobsResponse{}
	_body, _err := client.ListMediaProducingJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of packaging jobs.
//
// @param request - ListPackageJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPackageJobsResponse
func (client *Client) ListPackageJobsWithOptions(request *ListPackageJobsRequest, runtime *dara.RuntimeOptions) (_result *ListPackageJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPackageJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPackageJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of packaging jobs.
//
// @param request - ListPackageJobsRequest
//
// @return ListPackageJobsResponse
func (client *Client) ListPackageJobs(request *ListPackageJobsRequest) (_result *ListPackageJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPackageJobsResponse{}
	_body, _err := client.ListPackageJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ApsaraVideo Media Processing (MPS) queues.
//
// @param request - ListPipelinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelinesResponse
func (client *Client) ListPipelinesWithOptions(request *ListPipelinesRequest, runtime *dara.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelines"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ApsaraVideo Media Processing (MPS) queues.
//
// @param request - ListPipelinesRequest
//
// @return ListPipelinesResponse
func (client *Client) ListPipelines(request *ListPipelinesRequest) (_result *ListPipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPipelinesResponse{}
	_body, _err := client.ListPipelinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists programs.
//
// @param request - ListProgramsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProgramsResponse
func (client *Client) ListProgramsWithOptions(request *ListProgramsRequest, runtime *dara.RuntimeOptions) (_result *ListProgramsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProgramName) {
		query["ProgramName"] = request.ProgramName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPrograms"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProgramsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists programs.
//
// @param request - ListProgramsRequest
//
// @return ListProgramsResponse
func (client *Client) ListPrograms(request *ListProgramsRequest) (_result *ListProgramsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProgramsResponse{}
	_body, _err := client.ListProgramsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of media assets in the public media library that meet the specified conditions. A maximum of 100 media assets can be returned.
//
// @param request - ListPublicMediaBasicInfosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPublicMediaBasicInfosResponse
func (client *Client) ListPublicMediaBasicInfosWithOptions(request *ListPublicMediaBasicInfosRequest, runtime *dara.RuntimeOptions) (_result *ListPublicMediaBasicInfosResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.IncludeFileBasicInfo) {
		query["IncludeFileBasicInfo"] = request.IncludeFileBasicInfo
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.MediaTagId) {
		query["MediaTagId"] = request.MediaTagId
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPublicMediaBasicInfos"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPublicMediaBasicInfosResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of media assets in the public media library that meet the specified conditions. A maximum of 100 media assets can be returned.
//
// @param request - ListPublicMediaBasicInfosRequest
//
// @return ListPublicMediaBasicInfosResponse
func (client *Client) ListPublicMediaBasicInfos(request *ListPublicMediaBasicInfosRequest) (_result *ListPublicMediaBasicInfosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPublicMediaBasicInfosResponse{}
	_body, _err := client.ListPublicMediaBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义库下实体列表
//
// @param request - ListRecognitionEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecognitionEntitiesResponse
func (client *Client) ListRecognitionEntitiesWithOptions(request *ListRecognitionEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListRecognitionEntitiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecognitionEntities"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecognitionEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义库下实体列表
//
// @param request - ListRecognitionEntitiesRequest
//
// @return ListRecognitionEntitiesResponse
func (client *Client) ListRecognitionEntities(request *ListRecognitionEntitiesRequest) (_result *ListRecognitionEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecognitionEntitiesResponse{}
	_body, _err := client.ListRecognitionEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义库列表
//
// @param request - ListRecognitionLibsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecognitionLibsResponse
func (client *Client) ListRecognitionLibsWithOptions(request *ListRecognitionLibsRequest, runtime *dara.RuntimeOptions) (_result *ListRecognitionLibsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecognitionLibs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecognitionLibsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义库列表
//
// @param request - ListRecognitionLibsRequest
//
// @return ListRecognitionLibsResponse
func (client *Client) ListRecognitionLibs(request *ListRecognitionLibsRequest) (_result *ListRecognitionLibsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecognitionLibsResponse{}
	_body, _err := client.ListRecognitionLibsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取自定义库实体的图片列表
//
// @param request - ListRecognitionSamplesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecognitionSamplesResponse
func (client *Client) ListRecognitionSamplesWithOptions(request *ListRecognitionSamplesRequest, runtime *dara.RuntimeOptions) (_result *ListRecognitionSamplesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Algorithm) {
		query["Algorithm"] = request.Algorithm
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.LibId) {
		query["LibId"] = request.LibId
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRecognitionSamples"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRecognitionSamplesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取自定义库实体的图片列表
//
// @param request - ListRecognitionSamplesRequest
//
// @return ListRecognitionSamplesResponse
func (client *Client) ListRecognitionSamples(request *ListRecognitionSamplesRequest) (_result *ListRecognitionSamplesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRecognitionSamplesResponse{}
	_body, _err := client.ListRecognitionSamplesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the program schedule of a MediaWeaver channel.
//
// @param request - ListSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSchedulesResponse
func (client *Client) ListSchedulesWithOptions(request *ListSchedulesRequest, runtime *dara.RuntimeOptions) (_result *ListSchedulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WindowDurationSeconds) {
		query["WindowDurationSeconds"] = request.WindowDurationSeconds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSchedules"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the program schedule of a MediaWeaver channel.
//
// @param request - ListSchedulesRequest
//
// @return ListSchedulesResponse
func (client *Client) ListSchedules(request *ListSchedulesRequest) (_result *ListSchedulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSchedulesResponse{}
	_body, _err := client.ListSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about search libraries.
//
// @param request - ListSearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchLibResponse
func (client *Client) ListSearchLibWithOptions(request *ListSearchLibRequest, runtime *dara.RuntimeOptions) (_result *ListSearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about search libraries.
//
// @param request - ListSearchLibRequest
//
// @return ListSearchLibResponse
func (client *Client) ListSearchLib(request *ListSearchLibRequest) (_result *ListSearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSearchLibResponse{}
	_body, _err := client.ListSearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of intelligent jobs based on specified parameters.
//
// @param request - ListSmartJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmartJobsResponse
func (client *Client) ListSmartJobsWithOptions(request *ListSmartJobsRequest, runtime *dara.RuntimeOptions) (_result *ListSmartJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSmartJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmartJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of intelligent jobs based on specified parameters.
//
// @param request - ListSmartJobsRequest
//
// @return ListSmartJobsResponse
func (client *Client) ListSmartJobs(request *ListSmartJobsRequest) (_result *ListSmartJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSmartJobsResponse{}
	_body, _err := client.ListSmartJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of system digital humans. This operation supports paged queries.
//
// @param request - ListSmartSysAvatarModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmartSysAvatarModelsResponse
func (client *Client) ListSmartSysAvatarModelsWithOptions(request *ListSmartSysAvatarModelsRequest, runtime *dara.RuntimeOptions) (_result *ListSmartSysAvatarModelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SdkVersion) {
		query["SdkVersion"] = request.SdkVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSmartSysAvatarModels"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmartSysAvatarModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of system digital humans. This operation supports paged queries.
//
// @param request - ListSmartSysAvatarModelsRequest
//
// @return ListSmartSysAvatarModelsResponse
func (client *Client) ListSmartSysAvatarModels(request *ListSmartSysAvatarModelsRequest) (_result *ListSmartSysAvatarModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSmartSysAvatarModelsResponse{}
	_body, _err := client.ListSmartSysAvatarModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of speaker groups, including the name, gender, and sample audio of each speaker. The list is grouped by scenario.
//
// @param request - ListSmartVoiceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSmartVoiceGroupsResponse
func (client *Client) ListSmartVoiceGroupsWithOptions(request *ListSmartVoiceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListSmartVoiceGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSmartVoiceGroups"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSmartVoiceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of speaker groups, including the name, gender, and sample audio of each speaker. The list is grouped by scenario.
//
// @param request - ListSmartVoiceGroupsRequest
//
// @return ListSmartVoiceGroupsResponse
func (client *Client) ListSmartVoiceGroups(request *ListSmartVoiceGroupsRequest) (_result *ListSmartVoiceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSmartVoiceGroupsResponse{}
	_body, _err := client.ListSmartVoiceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of snapshot jobs.
//
// @param request - ListSnapshotJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotJobsResponse
func (client *Client) ListSnapshotJobsWithOptions(request *ListSnapshotJobsRequest, runtime *dara.RuntimeOptions) (_result *ListSnapshotJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSnapshotJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSnapshotJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of snapshot jobs.
//
// @param request - ListSnapshotJobsRequest
//
// @return ListSnapshotJobsResponse
func (client *Client) ListSnapshotJobs(request *ListSnapshotJobsRequest) (_result *ListSnapshotJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSnapshotJobsResponse{}
	_body, _err := client.ListSnapshotJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists source locations.
//
// @param request - ListSourceLocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSourceLocationsResponse
func (client *Client) ListSourceLocationsWithOptions(request *ListSourceLocationsRequest, runtime *dara.RuntimeOptions) (_result *ListSourceLocationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterState) {
		query["FilterState"] = request.FilterState
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortByModifiedTime) {
		query["SortByModifiedTime"] = request.SortByModifiedTime
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSourceLocations"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSourceLocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists source locations.
//
// @param request - ListSourceLocationsRequest
//
// @return ListSourceLocationsResponse
func (client *Client) ListSourceLocations(request *ListSourceLocationsRequest) (_result *ListSourceLocationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSourceLocationsResponse{}
	_body, _err := client.ListSourceLocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists sources in MediaWeaver.
//
// @param request - ListSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSourcesResponse
func (client *Client) ListSourcesWithOptions(request *ListSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilterState) {
		query["FilterState"] = request.FilterState
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SortByModifiedTime) {
		query["SortByModifiedTime"] = request.SortByModifiedTime
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSources"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists sources in MediaWeaver.
//
// @param request - ListSourcesRequest
//
// @return ListSourcesResponse
func (client *Client) ListSources(request *ListSourcesRequest) (_result *ListSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSourcesResponse{}
	_body, _err := client.ListSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of system templates.
//
// Description:
//
// Template types:
//
// 1.  1: transcoding template.
//
// 2.  2: snapshot template.
//
// 3.  3: animated image template.
//
// 4.  4\\. image watermark template.
//
// 5.  5: text watermark template.
//
// 6.  6: subtitle template.
//
// 7.  7: AI-assisted content moderation template.
//
// 8.  8: AI-assisted intelligent thumbnail template.
//
// 9.  9: AI-assisted intelligent erasure template.
//
// Subtypes of transcoding templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (AudioTranscode): audio transcoding template.
//
// 3.  3 (Remux): container format conversion template.
//
// 4.  4 (NarrowBandV1): Narrowband HD 1.0 template.
//
// 5.  5 (NarrowBandV2): Narrowband HD 2.0 template.
//
// Subtypes of snapshot templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (Sprite): sprite template.
//
// 3.  3 (WebVtt): WebVTT template.
//
// Subtypes of AI-assisted content moderation templates:
//
// 1.  1 (Video): video moderation template.
//
// 2.  2 (Audio): audio moderation template.
//
// 3.  3 (Image): image moderation template.
//
// Subtypes of AI-assisted intelligent erasure templates:
//
// 1.  1 (VideoDelogo): logo erasure template.
//
// 2.  2 (VideoDetext): subtitle erasure template.
//
// @param request - ListSystemTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSystemTemplatesResponse
func (client *Client) ListSystemTemplatesWithOptions(request *ListSystemTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListSystemTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Subtype) {
		query["Subtype"] = request.Subtype
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSystemTemplates"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSystemTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of system templates.
//
// Description:
//
// Template types:
//
// 1.  1: transcoding template.
//
// 2.  2: snapshot template.
//
// 3.  3: animated image template.
//
// 4.  4\\. image watermark template.
//
// 5.  5: text watermark template.
//
// 6.  6: subtitle template.
//
// 7.  7: AI-assisted content moderation template.
//
// 8.  8: AI-assisted intelligent thumbnail template.
//
// 9.  9: AI-assisted intelligent erasure template.
//
// Subtypes of transcoding templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (AudioTranscode): audio transcoding template.
//
// 3.  3 (Remux): container format conversion template.
//
// 4.  4 (NarrowBandV1): Narrowband HD 1.0 template.
//
// 5.  5 (NarrowBandV2): Narrowband HD 2.0 template.
//
// Subtypes of snapshot templates:
//
// 1.  1 (Normal): regular template.
//
// 2.  2 (Sprite): sprite template.
//
// 3.  3 (WebVtt): WebVTT template.
//
// Subtypes of AI-assisted content moderation templates:
//
// 1.  1 (Video): video moderation template.
//
// 2.  2 (Audio): audio moderation template.
//
// 3.  3 (Image): image moderation template.
//
// Subtypes of AI-assisted intelligent erasure templates:
//
// 1.  1 (VideoDelogo): logo erasure template.
//
// 2.  2 (VideoDetext): subtitle erasure template.
//
// @param request - ListSystemTemplatesRequest
//
// @return ListSystemTemplatesResponse
func (client *Client) ListSystemTemplates(request *ListSystemTemplatesRequest) (_result *ListSystemTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSystemTemplatesResponse{}
	_body, _err := client.ListSystemTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of templates that meet the specified conditions. You can query templates based on information such as the template status and creation source.
//
// Description:
//
// A template is an encapsulation of the timeline of a media editing and production job. You can define a common timeline as a template. When you have the same requirements, you need to only specify key parameters and materials to produce videos.
//
//   - For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//   - For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - ListTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSource) {
		query["CreateSource"] = request.CreateSource
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortType) {
		query["SortType"] = request.SortType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of templates that meet the specified conditions. You can query templates based on information such as the template status and creation source.
//
// Description:
//
// A template is an encapsulation of the timeline of a media editing and production job. You can define a common timeline as a template. When you have the same requirements, you need to only specify key parameters and materials to produce videos.
//
//   - For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//   - For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of transcoding jobs.
//
// @param request - ListTranscodeJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscodeJobsResponse
func (client *Client) ListTranscodeJobsWithOptions(request *ListTranscodeJobsRequest, runtime *dara.RuntimeOptions) (_result *ListTranscodeJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentJobId) {
		query["ParentJobId"] = request.ParentJobId
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTranscodeJobs"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTranscodeJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of transcoding jobs.
//
// @param request - ListTranscodeJobsRequest
//
// @return ListTranscodeJobsResponse
func (client *Client) ListTranscodeJobs(request *ListTranscodeJobsRequest) (_result *ListTranscodeJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTranscodeJobsResponse{}
	_body, _err := client.ListTranscodeJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists VOD packaging assets.
//
// @param request - ListVodPackagingAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVodPackagingAssetsResponse
func (client *Client) ListVodPackagingAssetsWithOptions(request *ListVodPackagingAssetsRequest, runtime *dara.RuntimeOptions) (_result *ListVodPackagingAssetsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVodPackagingAssets"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVodPackagingAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists VOD packaging assets.
//
// @param request - ListVodPackagingAssetsRequest
//
// @return ListVodPackagingAssetsResponse
func (client *Client) ListVodPackagingAssets(request *ListVodPackagingAssetsRequest) (_result *ListVodPackagingAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVodPackagingAssetsResponse{}
	_body, _err := client.ListVodPackagingAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists packaging configurations.
//
// @param request - ListVodPackagingConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVodPackagingConfigurationsResponse
func (client *Client) ListVodPackagingConfigurationsWithOptions(request *ListVodPackagingConfigurationsRequest, runtime *dara.RuntimeOptions) (_result *ListVodPackagingConfigurationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVodPackagingConfigurations"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVodPackagingConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists packaging configurations.
//
// @param request - ListVodPackagingConfigurationsRequest
//
// @return ListVodPackagingConfigurationsResponse
func (client *Client) ListVodPackagingConfigurations(request *ListVodPackagingConfigurationsRequest) (_result *ListVodPackagingConfigurationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVodPackagingConfigurationsResponse{}
	_body, _err := client.ListVodPackagingConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists packaging groups.
//
// @param request - ListVodPackagingGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVodPackagingGroupsResponse
func (client *Client) ListVodPackagingGroupsWithOptions(request *ListVodPackagingGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListVodPackagingGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVodPackagingGroups"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVodPackagingGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists packaging groups.
//
// @param request - ListVodPackagingGroupsRequest
//
// @return ListVodPackagingGroupsResponse
func (client *Client) ListVodPackagingGroups(request *ListVodPackagingGroupsRequest) (_result *ListVodPackagingGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListVodPackagingGroupsResponse{}
	_body, _err := client.ListVodPackagingGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 工作流任务实例列表
//
// @param request - ListWorkflowTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkflowTasksResponse
func (client *Client) ListWorkflowTasksWithOptions(request *ListWorkflowTasksRequest, runtime *dara.RuntimeOptions) (_result *ListWorkflowTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfCreateTime) {
		query["EndOfCreateTime"] = request.EndOfCreateTime
	}

	if !dara.IsNil(request.KeyText) {
		query["KeyText"] = request.KeyText
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartOfCreateTime) {
		query["StartOfCreateTime"] = request.StartOfCreateTime
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	if !dara.IsNil(request.WorkflowName) {
		query["WorkflowName"] = request.WorkflowName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkflowTasks"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkflowTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 工作流任务实例列表
//
// @param request - ListWorkflowTasksRequest
//
// @return ListWorkflowTasksResponse
func (client *Client) ListWorkflowTasks(request *ListWorkflowTasksRequest) (_result *ListWorkflowTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkflowTasksResponse{}
	_body, _err := client.ListWorkflowTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启MediaConnect双流灾备
//
// @param request - OpenMediaConnectFlowFailoverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenMediaConnectFlowFailoverResponse
func (client *Client) OpenMediaConnectFlowFailoverWithOptions(request *OpenMediaConnectFlowFailoverRequest, runtime *dara.RuntimeOptions) (_result *OpenMediaConnectFlowFailoverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenMediaConnectFlowFailover"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenMediaConnectFlowFailoverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启MediaConnect双流灾备
//
// @param request - OpenMediaConnectFlowFailoverRequest
//
// @return OpenMediaConnectFlowFailoverResponse
func (client *Client) OpenMediaConnectFlowFailover(request *OpenMediaConnectFlowFailoverRequest) (_result *OpenMediaConnectFlowFailoverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OpenMediaConnectFlowFailoverResponse{}
	_body, _err := client.OpenMediaConnectFlowFailoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询版权水印提取任务
//
// @param request - QueryCopyrightExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopyrightExtractJobResponse
func (client *Client) QueryCopyrightExtractJobWithOptions(request *QueryCopyrightExtractJobRequest, runtime *dara.RuntimeOptions) (_result *QueryCopyrightExtractJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCopyrightExtractJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCopyrightExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询版权水印提取任务
//
// @param request - QueryCopyrightExtractJobRequest
//
// @return QueryCopyrightExtractJobResponse
func (client *Client) QueryCopyrightExtractJob(request *QueryCopyrightExtractJobRequest) (_result *QueryCopyrightExtractJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCopyrightExtractJobResponse{}
	_body, _err := client.QueryCopyrightExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询视频版权水印任务列表
//
// @param request - QueryCopyrightJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCopyrightJobListResponse
func (client *Client) QueryCopyrightJobListWithOptions(request *QueryCopyrightJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryCopyrightJobListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
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
		Action:      dara.String("QueryCopyrightJobList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCopyrightJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频版权水印任务列表
//
// @param request - QueryCopyrightJobListRequest
//
// @return QueryCopyrightJobListResponse
func (client *Client) QueryCopyrightJobList(request *QueryCopyrightJobListRequest) (_result *QueryCopyrightJobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCopyrightJobListResponse{}
	_body, _err := client.QueryCopyrightJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of media fingerprint analysis jobs.
//
// @param request - QueryDNAJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDNAJobListResponse
func (client *Client) QueryDNAJobListWithOptions(request *QueryDNAJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryDNAJobListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDNAJobList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDNAJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of media fingerprint analysis jobs.
//
// @param request - QueryDNAJobListRequest
//
// @return QueryDNAJobListResponse
func (client *Client) QueryDNAJobList(request *QueryDNAJobListRequest) (_result *QueryDNAJobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDNAJobListResponse{}
	_body, _err := client.QueryDNAJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status and result of an intelligent production job.
//
// @param request - QueryIProductionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIProductionJobResponse
func (client *Client) QueryIProductionJobWithOptions(request *QueryIProductionJobRequest, runtime *dara.RuntimeOptions) (_result *QueryIProductionJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryIProductionJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryIProductionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status and result of an intelligent production job.
//
// @param request - QueryIProductionJobRequest
//
// @return QueryIProductionJobResponse
func (client *Client) QueryIProductionJob(request *QueryIProductionJobRequest) (_result *QueryIProductionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryIProductionJobResponse{}
	_body, _err := client.QueryIProductionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a content moderation job.
//
// Description:
//
// In the content moderation results, the moderation results of the video are sorted in ascending order by time into a timeline. If the video is long, the content moderation results are paginated, and the first page is returned. You can call this operation again to query the remaining moderation results of the video.
//
// @param request - QueryMediaCensorJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaCensorJobDetailResponse
func (client *Client) QueryMediaCensorJobDetailWithOptions(request *QueryMediaCensorJobDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaCensorJobDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaCensorJobDetail"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaCensorJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a content moderation job.
//
// Description:
//
// In the content moderation results, the moderation results of the video are sorted in ascending order by time into a timeline. If the video is long, the content moderation results are paginated, and the first page is returned. You can call this operation again to query the remaining moderation results of the video.
//
// @param request - QueryMediaCensorJobDetailRequest
//
// @return QueryMediaCensorJobDetailResponse
func (client *Client) QueryMediaCensorJobDetail(request *QueryMediaCensorJobDetailRequest) (_result *QueryMediaCensorJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMediaCensorJobDetailResponse{}
	_body, _err := client.QueryMediaCensorJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of content moderation jobs.
//
// Description:
//
// You can call this operation to query only the content moderation jobs within the most recent three months.
//
// @param request - QueryMediaCensorJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaCensorJobListResponse
func (client *Client) QueryMediaCensorJobListWithOptions(request *QueryMediaCensorJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaCensorJobListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndOfJobCreatedTimeRange) {
		query["EndOfJobCreatedTimeRange"] = request.EndOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.JobIds) {
		query["JobIds"] = request.JobIds
	}

	if !dara.IsNil(request.MaximumPageSize) {
		query["MaximumPageSize"] = request.MaximumPageSize
	}

	if !dara.IsNil(request.NextPageToken) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartOfJobCreatedTimeRange) {
		query["StartOfJobCreatedTimeRange"] = request.StartOfJobCreatedTimeRange
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaCensorJobList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaCensorJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of content moderation jobs.
//
// Description:
//
// You can call this operation to query only the content moderation jobs within the most recent three months.
//
// @param request - QueryMediaCensorJobListRequest
//
// @return QueryMediaCensorJobListResponse
func (client *Client) QueryMediaCensorJobList(request *QueryMediaCensorJobListRequest) (_result *QueryMediaCensorJobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMediaCensorJobListResponse{}
	_body, _err := client.QueryMediaCensorJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the indexing jobs enabled for a media asset.
//
// @param request - QueryMediaIndexJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMediaIndexJobResponse
func (client *Client) QueryMediaIndexJobWithOptions(request *QueryMediaIndexJobRequest, runtime *dara.RuntimeOptions) (_result *QueryMediaIndexJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMediaIndexJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMediaIndexJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the indexing jobs enabled for a media asset.
//
// @param request - QueryMediaIndexJobRequest
//
// @return QueryMediaIndexJobResponse
func (client *Client) QueryMediaIndexJob(request *QueryMediaIndexJobRequest) (_result *QueryMediaIndexJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMediaIndexJobResponse{}
	_body, _err := client.QueryMediaIndexJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a search index.
//
// @param request - QuerySearchIndexRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySearchIndexResponse
func (client *Client) QuerySearchIndexWithOptions(request *QuerySearchIndexRequest, runtime *dara.RuntimeOptions) (_result *QuerySearchIndexResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndexType) {
		query["IndexType"] = request.IndexType
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySearchIndex"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySearchIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a search index.
//
// @param request - QuerySearchIndexRequest
//
// @return QuerySearchIndexResponse
func (client *Client) QuerySearchIndex(request *QuerySearchIndexRequest) (_result *QuerySearchIndexResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySearchIndexResponse{}
	_body, _err := client.QuerySearchIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a search library.
//
// @param request - QuerySearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySearchLibResponse
func (client *Client) QuerySearchLibWithOptions(request *QuerySearchLibRequest, runtime *dara.RuntimeOptions) (_result *QuerySearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a search library.
//
// @param request - QuerySearchLibRequest
//
// @return QuerySearchLibResponse
func (client *Client) QuerySearchLib(request *QuerySearchLibRequest) (_result *QuerySearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySearchLibResponse{}
	_body, _err := client.QuerySearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a smart tagging job.
//
// @param request - QuerySmarttagJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySmarttagJobResponse
func (client *Client) QuerySmarttagJobWithOptions(request *QuerySmarttagJobRequest, runtime *dara.RuntimeOptions) (_result *QuerySmarttagJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySmarttagJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySmarttagJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a smart tagging job.
//
// @param request - QuerySmarttagJobRequest
//
// @return QuerySmarttagJobResponse
func (client *Client) QuerySmarttagJob(request *QuerySmarttagJobRequest) (_result *QuerySmarttagJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySmarttagJobResponse{}
	_body, _err := client.QuerySmarttagJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询视频溯源水印ab流任务
//
// @param request - QueryTraceAbJobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTraceAbJobListResponse
func (client *Client) QueryTraceAbJobListWithOptions(request *QueryTraceAbJobListRequest, runtime *dara.RuntimeOptions) (_result *QueryTraceAbJobListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TraceMediaId) {
		query["TraceMediaId"] = request.TraceMediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTraceAbJobList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTraceAbJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频溯源水印ab流任务
//
// @param request - QueryTraceAbJobListRequest
//
// @return QueryTraceAbJobListResponse
func (client *Client) QueryTraceAbJobList(request *QueryTraceAbJobListRequest) (_result *QueryTraceAbJobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTraceAbJobListResponse{}
	_body, _err := client.QueryTraceAbJobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询溯源水印提取任务
//
// @param request - QueryTraceExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTraceExtractJobResponse
func (client *Client) QueryTraceExtractJobWithOptions(request *QueryTraceExtractJobRequest, runtime *dara.RuntimeOptions) (_result *QueryTraceExtractJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTraceExtractJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTraceExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询溯源水印提取任务
//
// @param request - QueryTraceExtractJobRequest
//
// @return QueryTraceExtractJobResponse
func (client *Client) QueryTraceExtractJob(request *QueryTraceExtractJobRequest) (_result *QueryTraceExtractJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTraceExtractJobResponse{}
	_body, _err := client.QueryTraceExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询视频溯源水印m3u8任务
//
// @param request - QueryTraceM3u8JobListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTraceM3u8JobListResponse
func (client *Client) QueryTraceM3u8JobListWithOptions(request *QueryTraceM3u8JobListRequest, runtime *dara.RuntimeOptions) (_result *QueryTraceM3u8JobListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTimeEnd) {
		query["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !dara.IsNil(request.CreateTimeStart) {
		query["CreateTimeStart"] = request.CreateTimeStart
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
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
		Action:      dara.String("QueryTraceM3u8JobList"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTraceM3u8JobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频溯源水印m3u8任务
//
// @param request - QueryTraceM3u8JobListRequest
//
// @return QueryTraceM3u8JobListResponse
func (client *Client) QueryTraceM3u8JobList(request *QueryTraceM3u8JobListRequest) (_result *QueryTraceM3u8JobListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTraceM3u8JobListResponse{}
	_body, _err := client.QueryTraceM3u8JobListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询视频理解任务结果
//
// @param tmpReq - QueryVideoCognitionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVideoCognitionJobResponse
func (client *Client) QueryVideoCognitionJobWithOptions(tmpReq *QueryVideoCognitionJobRequest, runtime *dara.RuntimeOptions) (_result *QueryVideoCognitionJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryVideoCognitionJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IncludeResults) {
		request.IncludeResultsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IncludeResults, dara.String("IncludeResults"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IncludeResultsShrink) {
		query["IncludeResults"] = request.IncludeResultsShrink
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVideoCognitionJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVideoCognitionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询视频理解任务结果
//
// @param request - QueryVideoCognitionJobRequest
//
// @return QueryVideoCognitionJobResponse
func (client *Client) QueryVideoCognitionJob(request *QueryVideoCognitionJobRequest) (_result *QueryVideoCognitionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVideoCognitionJobResponse{}
	_body, _err := client.QueryVideoCognitionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain a new upload credential for a media asset after its upload 
//
// Description:
//
// You can also call this operation to overwrite media files. After you obtain the upload URL of a media file, you can upload the media file again without changing the audio or video ID.
//
// @param request - RefreshUploadMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefreshUploadMediaResponse
func (client *Client) RefreshUploadMediaWithOptions(request *RefreshUploadMediaRequest, runtime *dara.RuntimeOptions) (_result *RefreshUploadMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefreshUploadMedia"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefreshUploadMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain a new upload credential for a media asset after its upload credential expires.
//
// Description:
//
// You can also call this operation to overwrite media files. After you obtain the upload URL of a media file, you can upload the media file again without changing the audio or video ID.
//
// @param request - RefreshUploadMediaRequest
//
// @return RefreshUploadMediaResponse
func (client *Client) RefreshUploadMedia(request *RefreshUploadMediaRequest) (_result *RefreshUploadMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefreshUploadMediaResponse{}
	_body, _err := client.RefreshUploadMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a media asset with Intelligent Media Services (IMS). IMS assigns an ID to the media asset. This operation asynchronously accesses the media asset service in which the media asset is stored to obtain the file information of the media asset based on the input URL. You can also specify basic information, such as the title, tags, and description, for the media asset. This operation returns the ID of the media asset. You can call the GetMediaInfo operation based on the ID to query the details of the media asset. You can set InputURL only to the URL of an Object Storage Service (OSS) file or an ApsaraVideo VOD media asset.
//
// Description:
//
// Registering a media asset is an asynchronous job that takes 2 to 3 seconds. When the operation returns the ID of the media asset, the registration may have not be completed. If you call the GetMediaInfo operation at this time, you may fail to obtain the information about the media asset.
//
// @param request - RegisterMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterMediaInfoResponse
func (client *Client) RegisterMediaInfoWithOptions(request *RegisterMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *RegisterMediaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaTags) {
		query["MediaTags"] = request.MediaTags
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
	}

	if !dara.IsNil(request.RegisterConfig) {
		query["RegisterConfig"] = request.RegisterConfig
	}

	if !dara.IsNil(request.SmartTagTemplateId) {
		query["SmartTagTemplateId"] = request.SmartTagTemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterMediaInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterMediaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a media asset with Intelligent Media Services (IMS). IMS assigns an ID to the media asset. This operation asynchronously accesses the media asset service in which the media asset is stored to obtain the file information of the media asset based on the input URL. You can also specify basic information, such as the title, tags, and description, for the media asset. This operation returns the ID of the media asset. You can call the GetMediaInfo operation based on the ID to query the details of the media asset. You can set InputURL only to the URL of an Object Storage Service (OSS) file or an ApsaraVideo VOD media asset.
//
// Description:
//
// Registering a media asset is an asynchronous job that takes 2 to 3 seconds. When the operation returns the ID of the media asset, the registration may have not be completed. If you call the GetMediaInfo operation at this time, you may fail to obtain the information about the media asset.
//
// @param request - RegisterMediaInfoRequest
//
// @return RegisterMediaInfoResponse
func (client *Client) RegisterMediaInfo(request *RegisterMediaInfoRequest) (_result *RegisterMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterMediaInfoResponse{}
	_body, _err := client.RegisterMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a media stream.
//
// Description:
//
// You can call this operation to register a media stream file in an Object Storage Service (OSS) bucket with Intelligent Media Services (IMS) and associate the media stream with the specified media asset ID.
//
// @param request - RegisterMediaStreamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterMediaStreamResponse
func (client *Client) RegisterMediaStreamWithOptions(request *RegisterMediaStreamRequest, runtime *dara.RuntimeOptions) (_result *RegisterMediaStreamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.StreamTags) {
		query["StreamTags"] = request.StreamTags
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterMediaStream"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterMediaStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a media stream.
//
// Description:
//
// You can call this operation to register a media stream file in an Object Storage Service (OSS) bucket with Intelligent Media Services (IMS) and associate the media stream with the specified media asset ID.
//
// @param request - RegisterMediaStreamRequest
//
// @return RegisterMediaStreamResponse
func (client *Client) RegisterMediaStream(request *RegisterMediaStreamRequest) (_result *RegisterMediaStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterMediaStreamResponse{}
	_body, _err := client.RegisterMediaStreamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 恢复某个MediaConnect实例的某个输出
//
// @param request - ResumeMediaConnectFlowOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeMediaConnectFlowOutputResponse
func (client *Client) ResumeMediaConnectFlowOutputWithOptions(request *ResumeMediaConnectFlowOutputRequest, runtime *dara.RuntimeOptions) (_result *ResumeMediaConnectFlowOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeMediaConnectFlowOutput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeMediaConnectFlowOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 恢复某个MediaConnect实例的某个输出
//
// @param request - ResumeMediaConnectFlowOutputRequest
//
// @return ResumeMediaConnectFlowOutputResponse
func (client *Client) ResumeMediaConnectFlowOutput(request *ResumeMediaConnectFlowOutputRequest) (_result *ResumeMediaConnectFlowOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeMediaConnectFlowOutputResponse{}
	_body, _err := client.ResumeMediaConnectFlowOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries online editing projects by creation time and status.
//
// @param request - SearchEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchEditingProjectResponse
func (client *Client) SearchEditingProjectWithOptions(request *SearchEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *SearchEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateSource) {
		query["CreateSource"] = request.CreateSource
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectType) {
		query["ProjectType"] = request.ProjectType
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

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchEditingProject"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries online editing projects by creation time and status.
//
// @param request - SearchEditingProjectRequest
//
// @return SearchEditingProjectResponse
func (client *Client) SearchEditingProject(request *SearchEditingProjectRequest) (_result *SearchEditingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchEditingProjectResponse{}
	_body, _err := client.SearchEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Re-analyzes the search index jobs of media assets. You can re-run the search index jobs of up to 20 media assets in each request.
//
// @param request - SearchIndexJobRerunRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchIndexJobRerunResponse
func (client *Client) SearchIndexJobRerunWithOptions(request *SearchIndexJobRerunRequest, runtime *dara.RuntimeOptions) (_result *SearchIndexJobRerunResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	if !dara.IsNil(request.Task) {
		query["Task"] = request.Task
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchIndexJobRerun"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchIndexJobRerunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Re-analyzes the search index jobs of media assets. You can re-run the search index jobs of up to 20 media assets in each request.
//
// @param request - SearchIndexJobRerunRequest
//
// @return SearchIndexJobRerunResponse
func (client *Client) SearchIndexJobRerun(request *SearchIndexJobRerunRequest) (_result *SearchIndexJobRerunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchIndexJobRerunResponse{}
	_body, _err := client.SearchIndexJobRerunWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about media assets based on the request parameters.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaResponse
func (client *Client) SearchMediaWithOptions(request *SearchMediaRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.Match) {
		query["Match"] = request.Match
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScrollToken) {
		query["ScrollToken"] = request.ScrollToken
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMedia"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about media assets based on the request parameters.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaRequest
//
// @return SearchMediaResponse
func (client *Client) SearchMedia(request *SearchMediaRequest) (_result *SearchMediaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaResponse{}
	_body, _err := client.SearchMediaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries media assets based on character names, subtitles, or AI categories.
//
// Description:
//
// You can call this operation to query media assets or media asset clips based on character names, subtitles, or AI categories.
//
// @param request - SearchMediaByAILabelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaByAILabelResponse
func (client *Client) SearchMediaByAILabelWithOptions(request *SearchMediaByAILabelRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaByAILabelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MatchingMode) {
		query["MatchingMode"] = request.MatchingMode
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.MultimodalSearchType) {
		query["MultimodalSearchType"] = request.MultimodalSearchType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.SpecificSearch) {
		query["SpecificSearch"] = request.SpecificSearch
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMediaByAILabel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaByAILabelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media assets based on character names, subtitles, or AI categories.
//
// Description:
//
// You can call this operation to query media assets or media asset clips based on character names, subtitles, or AI categories.
//
// @param request - SearchMediaByAILabelRequest
//
// @return SearchMediaByAILabelResponse
func (client *Client) SearchMediaByAILabel(request *SearchMediaByAILabelRequest) (_result *SearchMediaByAILabelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaByAILabelResponse{}
	_body, _err := client.SearchMediaByAILabelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about media assets that are related to a specific face.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaByFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaByFaceResponse
func (client *Client) SearchMediaByFaceWithOptions(request *SearchMediaByFaceRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaByFaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.FaceSearchToken) {
		query["FaceSearchToken"] = request.FaceSearchToken
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PersonImageUrl) {
		query["PersonImageUrl"] = request.PersonImageUrl
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMediaByFace"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaByFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about media assets that are related to a specific face.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaByFaceRequest
//
// @return SearchMediaByFaceResponse
func (client *Client) SearchMediaByFace(request *SearchMediaByFaceRequest) (_result *SearchMediaByFaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaByFaceResponse{}
	_body, _err := client.SearchMediaByFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries media assets by using the hybrid search feature. This operation allows you to search for media assets by using natural language based on intelligent tag text search and the search capabilities of large language models (LLMs). This implements multimodal retrieval.
//
// @param request - SearchMediaByHybridRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaByHybridResponse
func (client *Client) SearchMediaByHybridWithOptions(request *SearchMediaByHybridRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaByHybridResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMediaByHybrid"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaByHybridResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media assets by using the hybrid search feature. This operation allows you to search for media assets by using natural language based on intelligent tag text search and the search capabilities of large language models (LLMs). This implements multimodal retrieval.
//
// @param request - SearchMediaByHybridRequest
//
// @return SearchMediaByHybridResponse
func (client *Client) SearchMediaByHybrid(request *SearchMediaByHybridRequest) (_result *SearchMediaByHybridResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaByHybridResponse{}
	_body, _err := client.SearchMediaByHybridWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries media assets by using the large visual model. You can use natural language for the query.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaByMultimodalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaByMultimodalResponse
func (client *Client) SearchMediaByMultimodalWithOptions(request *SearchMediaByMultimodalRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaByMultimodalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMediaByMultimodal"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaByMultimodalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries media assets by using the large visual model. You can use natural language for the query.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaByMultimodalRequest
//
// @return SearchMediaByMultimodalResponse
func (client *Client) SearchMediaByMultimodal(request *SearchMediaByMultimodalRequest) (_result *SearchMediaByMultimodalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaByMultimodalResponse{}
	_body, _err := client.SearchMediaByMultimodalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about media asset clips that are related to a specific face based on the response to the SearchMediaByFace operation.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaClipByFaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMediaClipByFaceResponse
func (client *Client) SearchMediaClipByFaceWithOptions(request *SearchMediaClipByFaceRequest, runtime *dara.RuntimeOptions) (_result *SearchMediaClipByFaceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.FaceSearchToken) {
		query["FaceSearchToken"] = request.FaceSearchToken
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMediaClipByFace"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMediaClipByFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about media asset clips that are related to a specific face based on the response to the SearchMediaByFace operation.
//
// Description:
//
// If you have questions about how to use the media asset search feature in Intelligent Media Services (IMS), contact technical support in the DingTalk group (ID 30415005038).
//
// @param request - SearchMediaClipByFaceRequest
//
// @return SearchMediaClipByFaceResponse
func (client *Client) SearchMediaClipByFace(request *SearchMediaClipByFaceRequest) (_result *SearchMediaClipByFaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMediaClipByFaceResponse{}
	_body, _err := client.SearchMediaClipByFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 搜索公共媒资信息
//
// @param request - SearchPublicMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchPublicMediaInfoResponse
func (client *Client) SearchPublicMediaInfoWithOptions(request *SearchPublicMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *SearchPublicMediaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Authorized) {
		query["Authorized"] = request.Authorized
	}

	if !dara.IsNil(request.DynamicMetaDataMatchFields) {
		query["DynamicMetaDataMatchFields"] = request.DynamicMetaDataMatchFields
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.Favorite) {
		query["Favorite"] = request.Favorite
	}

	if !dara.IsNil(request.MediaIds) {
		query["MediaIds"] = request.MediaIds
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchPublicMediaInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchPublicMediaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索公共媒资信息
//
// @param request - SearchPublicMediaInfoRequest
//
// @return SearchPublicMediaInfoResponse
func (client *Client) SearchPublicMediaInfo(request *SearchPublicMediaInfoRequest) (_result *SearchPublicMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchPublicMediaInfoResponse{}
	_body, _err := client.SearchPublicMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a DataChannel message to an AI agent.
//
// @param request - SendAIAgentDataChannelMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAIAgentDataChannelMessageResponse
func (client *Client) SendAIAgentDataChannelMessageWithOptions(request *SendAIAgentDataChannelMessageRequest, runtime *dara.RuntimeOptions) (_result *SendAIAgentDataChannelMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendAIAgentDataChannelMessage"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendAIAgentDataChannelMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a DataChannel message to an AI agent.
//
// @param request - SendAIAgentDataChannelMessageRequest
//
// @return SendAIAgentDataChannelMessageResponse
func (client *Client) SendAIAgentDataChannelMessage(request *SendAIAgentDataChannelMessageRequest) (_result *SendAIAgentDataChannelMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendAIAgentDataChannelMessageResponse{}
	_body, _err := client.SendAIAgentDataChannelMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Instructs an AI agent to immediately broadcast a text message and supports interruption settings.
//
// Description:
//
// You can call this operation to instruct an AI agent to broadcast the content that you specify. You can determine whether this broadcast can immediately interrupt the ongoing speech. The interruption is allowed by default.
//
// **Note**
//
//   - Make sure that the `InstanceId` is valid and corresponds to an existing AI agent.
//
//   - The content of `Text` must comply with the specifications and does not contain sensitive or inappropriate information.
//
//   - If you do not want the new broadcast to interrupt the ongoing speech, you must set `EnableInterrupt` to `false`.
//
// @param request - SendAIAgentSpeechRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAIAgentSpeechResponse
func (client *Client) SendAIAgentSpeechWithOptions(request *SendAIAgentSpeechRequest, runtime *dara.RuntimeOptions) (_result *SendAIAgentSpeechResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnableInterrupt) {
		query["EnableInterrupt"] = request.EnableInterrupt
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendAIAgentSpeech"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendAIAgentSpeechResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Instructs an AI agent to immediately broadcast a text message and supports interruption settings.
//
// Description:
//
// You can call this operation to instruct an AI agent to broadcast the content that you specify. You can determine whether this broadcast can immediately interrupt the ongoing speech. The interruption is allowed by default.
//
// **Note**
//
//   - Make sure that the `InstanceId` is valid and corresponds to an existing AI agent.
//
//   - The content of `Text` must comply with the specifications and does not contain sensitive or inappropriate information.
//
//   - If you do not want the new broadcast to interrupt the ongoing speech, you must set `EnableInterrupt` to `false`.
//
// @param request - SendAIAgentSpeechRequest
//
// @return SendAIAgentSpeechResponse
func (client *Client) SendAIAgentSpeech(request *SendAIAgentSpeechRequest) (_result *SendAIAgentSpeechResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendAIAgentSpeechResponse{}
	_body, _err := client.SendAIAgentSpeechWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 传入消息作为LLM输入。
//
// @param request - SendAIAgentTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendAIAgentTextResponse
func (client *Client) SendAIAgentTextWithOptions(request *SendAIAgentTextRequest, runtime *dara.RuntimeOptions) (_result *SendAIAgentTextResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendAIAgentText"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendAIAgentTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 传入消息作为LLM输入。
//
// @param request - SendAIAgentTextRequest
//
// @return SendAIAgentTextResponse
func (client *Client) SendAIAgentText(request *SendAIAgentTextRequest) (_result *SendAIAgentTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendAIAgentTextResponse{}
	_body, _err := client.SendAIAgentTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a command to process a live stream snapshot job.
//
// @param request - SendLiveSnapshotJobCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLiveSnapshotJobCommandResponse
func (client *Client) SendLiveSnapshotJobCommandWithOptions(request *SendLiveSnapshotJobCommandRequest, runtime *dara.RuntimeOptions) (_result *SendLiveSnapshotJobCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Command) {
		body["Command"] = request.Command
	}

	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendLiveSnapshotJobCommand"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendLiveSnapshotJobCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a command to process a live stream snapshot job.
//
// @param request - SendLiveSnapshotJobCommandRequest
//
// @return SendLiveSnapshotJobCommandResponse
func (client *Client) SendLiveSnapshotJobCommand(request *SendLiveSnapshotJobCommandRequest) (_result *SendLiveSnapshotJobCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendLiveSnapshotJobCommandResponse{}
	_body, _err := client.SendLiveSnapshotJobCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a command to process a live stream transcoding job.
//
// @param request - SendLiveTranscodeJobCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendLiveTranscodeJobCommandResponse
func (client *Client) SendLiveTranscodeJobCommandWithOptions(request *SendLiveTranscodeJobCommandRequest, runtime *dara.RuntimeOptions) (_result *SendLiveTranscodeJobCommandResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendLiveTranscodeJobCommand"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendLiveTranscodeJobCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a command to process a live stream transcoding job.
//
// @param request - SendLiveTranscodeJobCommandRequest
//
// @return SendLiveTranscodeJobCommandResponse
func (client *Client) SendLiveTranscodeJobCommand(request *SendLiveTranscodeJobCommandRequest) (_result *SendLiveTranscodeJobCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendLiveTranscodeJobCommandResponse{}
	_body, _err := client.SendLiveTranscodeJobCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向IM客户端发送消息。
//
// @param request - SendMessageChatTextRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendMessageChatTextResponse
func (client *Client) SendMessageChatTextWithOptions(request *SendMessageChatTextRequest, runtime *dara.RuntimeOptions) (_result *SendMessageChatTextResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.NeedArchiving) {
		query["NeedArchiving"] = request.NeedArchiving
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Text) {
		query["Text"] = request.Text
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendMessageChatText"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendMessageChatTextResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向IM客户端发送消息。
//
// @param request - SendMessageChatTextRequest
//
// @return SendMessageChatTextResponse
func (client *Client) SendMessageChatText(request *SendMessageChatTextRequest) (_result *SendMessageChatTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendMessageChatTextResponse{}
	_body, _err := client.SendMessageChatTextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置声纹
//
// @param tmpReq - SetAIAgentVoiceprintRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetAIAgentVoiceprintResponse
func (client *Client) SetAIAgentVoiceprintWithOptions(tmpReq *SetAIAgentVoiceprintRequest, runtime *dara.RuntimeOptions) (_result *SetAIAgentVoiceprintResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetAIAgentVoiceprintShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.VoiceprintId) {
		query["VoiceprintId"] = request.VoiceprintId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetAIAgentVoiceprint"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetAIAgentVoiceprintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置声纹
//
// @param request - SetAIAgentVoiceprintRequest
//
// @return SetAIAgentVoiceprintResponse
func (client *Client) SetAIAgentVoiceprint(request *SetAIAgentVoiceprintRequest) (_result *SetAIAgentVoiceprintResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetAIAgentVoiceprintResponse{}
	_body, _err := client.SetAIAgentVoiceprintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置内容分析搜索配置
//
// @param request - SetContentAnalyzeConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetContentAnalyzeConfigResponse
func (client *Client) SetContentAnalyzeConfigWithOptions(request *SetContentAnalyzeConfigRequest, runtime *dara.RuntimeOptions) (_result *SetContentAnalyzeConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Auto) {
		query["Auto"] = request.Auto
	}

	if !dara.IsNil(request.SaveType) {
		query["SaveType"] = request.SaveType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetContentAnalyzeConfig"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetContentAnalyzeConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置内容分析搜索配置
//
// @param request - SetContentAnalyzeConfigRequest
//
// @return SetContentAnalyzeConfigResponse
func (client *Client) SetContentAnalyzeConfig(request *SetContentAnalyzeConfigRequest) (_result *SetContentAnalyzeConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetContentAnalyzeConfigResponse{}
	_body, _err := client.SetContentAnalyzeConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets a custom template as the default template.
//
// @param request - SetDefaultCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultCustomTemplateResponse
func (client *Client) SetDefaultCustomTemplateWithOptions(request *SetDefaultCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultCustomTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultCustomTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultCustomTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets a custom template as the default template.
//
// @param request - SetDefaultCustomTemplateRequest
//
// @return SetDefaultCustomTemplateResponse
func (client *Client) SetDefaultCustomTemplate(request *SetDefaultCustomTemplateRequest) (_result *SetDefaultCustomTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDefaultCustomTemplateResponse{}
	_body, _err := client.SetDefaultCustomTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置默认存储路径
//
// @param request - SetDefaultStorageLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultStorageLocationResponse
func (client *Client) SetDefaultStorageLocationWithOptions(request *SetDefaultStorageLocationRequest, runtime *dara.RuntimeOptions) (_result *SetDefaultStorageLocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Bucket) {
		query["Bucket"] = request.Bucket
	}

	if !dara.IsNil(request.Path) {
		query["Path"] = request.Path
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDefaultStorageLocation"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDefaultStorageLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置默认存储路径
//
// @param request - SetDefaultStorageLocationRequest
//
// @return SetDefaultStorageLocationResponse
func (client *Client) SetDefaultStorageLocation(request *SetDefaultStorageLocationRequest) (_result *SetDefaultStorageLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDefaultStorageLocationResponse{}
	_body, _err := client.SetDefaultStorageLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures a callback method for one or more events.
//
// @param request - SetEventCallbackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetEventCallbackResponse
func (client *Client) SetEventCallbackWithOptions(request *SetEventCallbackRequest, runtime *dara.RuntimeOptions) (_result *SetEventCallbackResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthKey) {
		query["AuthKey"] = request.AuthKey
	}

	if !dara.IsNil(request.AuthSwitch) {
		query["AuthSwitch"] = request.AuthSwitch
	}

	if !dara.IsNil(request.CallbackQueueName) {
		query["CallbackQueueName"] = request.CallbackQueueName
	}

	if !dara.IsNil(request.CallbackType) {
		query["CallbackType"] = request.CallbackType
	}

	if !dara.IsNil(request.CallbackURL) {
		query["CallbackURL"] = request.CallbackURL
	}

	if !dara.IsNil(request.EventTypeList) {
		query["EventTypeList"] = request.EventTypeList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetEventCallback"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetEventCallbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures a callback method for one or more events.
//
// @param request - SetEventCallbackRequest
//
// @return SetEventCallbackResponse
func (client *Client) SetEventCallback(request *SetEventCallbackRequest) (_result *SetEventCallbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetEventCallbackResponse{}
	_body, _err := client.SetEventCallbackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables event notifications for an AI agent and configures the callback URL and event types.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to configure event notifications for an AI agent. You can configure `EnableNotify` to enable or disable event notifications, configure `CallbackUrl` to specify a callback URL, and configure `EventTypes` to specify event types. You can also configure `Token` to specify an authentication token for enhanced security. The system returns a unique `RequestId` for subsequent tracing after a successful request.
//
// @param request - SetNotifyConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetNotifyConfigResponse
func (client *Client) SetNotifyConfigWithOptions(request *SetNotifyConfigRequest, runtime *dara.RuntimeOptions) (_result *SetNotifyConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.AudioOssPath) {
		query["AudioOssPath"] = request.AudioOssPath
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.EnableAudioRecording) {
		query["EnableAudioRecording"] = request.EnableAudioRecording
	}

	if !dara.IsNil(request.EnableNotify) {
		query["EnableNotify"] = request.EnableNotify
	}

	if !dara.IsNil(request.EventTypes) {
		query["EventTypes"] = request.EventTypes
	}

	if !dara.IsNil(request.Token) {
		query["Token"] = request.Token
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetNotifyConfig"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetNotifyConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables event notifications for an AI agent and configures the callback URL and event types.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to configure event notifications for an AI agent. You can configure `EnableNotify` to enable or disable event notifications, configure `CallbackUrl` to specify a callback URL, and configure `EventTypes` to specify event types. You can also configure `Token` to specify an authentication token for enhanced security. The system returns a unique `RequestId` for subsequent tracing after a successful request.
//
// @param request - SetNotifyConfigRequest
//
// @return SetNotifyConfigResponse
func (client *Client) SetNotifyConfig(request *SetNotifyConfigRequest) (_result *SetNotifyConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetNotifyConfigResponse{}
	_body, _err := client.SetNotifyConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an AI agent that is configured in the Intelligent Media Services (IMS) console.
//
// Description:
//
// You can call this operation to start an AI agent instance for a conversation. ``````“When the AI agent is started, the system returns a unique `InstanceId` for subsequent tracking and operations.
//
// @param tmpReq - StartAIAgentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIAgentInstanceResponse
func (client *Client) StartAIAgentInstanceWithOptions(tmpReq *StartAIAgentInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartAIAgentInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartAIAgentInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentConfig) {
		request.AgentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentConfig, dara.String("AgentConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChatSyncConfig) {
		request.ChatSyncConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChatSyncConfig, dara.String("ChatSyncConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RuntimeConfig) {
		request.RuntimeConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuntimeConfig, dara.String("RuntimeConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.AgentConfigShrink) {
		query["AgentConfig"] = request.AgentConfigShrink
	}

	if !dara.IsNil(request.ChatSyncConfigShrink) {
		query["ChatSyncConfig"] = request.ChatSyncConfigShrink
	}

	if !dara.IsNil(request.RuntimeConfigShrink) {
		query["RuntimeConfig"] = request.RuntimeConfigShrink
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAIAgentInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAIAgentInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an AI agent that is configured in the Intelligent Media Services (IMS) console.
//
// Description:
//
// You can call this operation to start an AI agent instance for a conversation. ``````“When the AI agent is started, the system returns a unique `InstanceId` for subsequent tracking and operations.
//
// @param request - StartAIAgentInstanceRequest
//
// @return StartAIAgentInstanceResponse
func (client *Client) StartAIAgentInstance(request *StartAIAgentInstanceRequest) (_result *StartAIAgentInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartAIAgentInstanceResponse{}
	_body, _err := client.StartAIAgentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一个智能体实例，返回智能体所在的频道、频道内名称以及进入频道所需的token。
//
// @param tmpReq - StartAIAgentOutboundCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAIAgentOutboundCallResponse
func (client *Client) StartAIAgentOutboundCallWithOptions(tmpReq *StartAIAgentOutboundCallRequest, runtime *dara.RuntimeOptions) (_result *StartAIAgentOutboundCallResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartAIAgentOutboundCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CallerNumber) {
		query["CallerNumber"] = request.CallerNumber
	}

	if !dara.IsNil(request.ConfigShrink) {
		query["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.ImsAIAgentFreeObCall) {
		query["ImsAIAgentFreeObCall"] = request.ImsAIAgentFreeObCall
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAIAgentOutboundCall"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAIAgentOutboundCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一个智能体实例，返回智能体所在的频道、频道内名称以及进入频道所需的token。
//
// @param request - StartAIAgentOutboundCallRequest
//
// @return StartAIAgentOutboundCallResponse
func (client *Client) StartAIAgentOutboundCall(request *StartAIAgentOutboundCallRequest) (_result *StartAIAgentOutboundCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartAIAgentOutboundCallResponse{}
	_body, _err := client.StartAIAgentOutboundCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a channel.
//
// @param request - StartChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartChannelResponse
func (client *Client) StartChannelWithOptions(request *StartChannelRequest, runtime *dara.RuntimeOptions) (_result *StartChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a channel.
//
// @param request - StartChannelRequest
//
// @return StartChannelResponse
func (client *Client) StartChannel(request *StartChannelRequest) (_result *StartChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartChannelResponse{}
	_body, _err := client.StartChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a MediaLive channel.
//
// Description:
//
//	You can call this operation only when the channel is idle. You cannot start a channel repeatedly.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - StartMediaLiveChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMediaLiveChannelResponse
func (client *Client) StartMediaLiveChannelWithOptions(request *StartMediaLiveChannelRequest, runtime *dara.RuntimeOptions) (_result *StartMediaLiveChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMediaLiveChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMediaLiveChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a MediaLive channel.
//
// Description:
//
//	You can call this operation only when the channel is idle. You cannot start a channel repeatedly.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - StartMediaLiveChannelRequest
//
// @return StartMediaLiveChannelResponse
func (client *Client) StartMediaLiveChannel(request *StartMediaLiveChannelRequest) (_result *StartMediaLiveChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartMediaLiveChannelResponse{}
	_body, _err := client.StartMediaLiveChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启一个机器人实例
//
// @param tmpReq - StartRtcRobotInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRtcRobotInstanceResponse
func (client *Client) StartRtcRobotInstanceWithOptions(tmpReq *StartRtcRobotInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartRtcRobotInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartRtcRobotInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthToken) {
		query["AuthToken"] = request.AuthToken
	}

	if !dara.IsNil(request.ChannelId) {
		query["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.ConfigShrink) {
		query["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.RobotId) {
		query["RobotId"] = request.RobotId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRtcRobotInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRtcRobotInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启一个机器人实例
//
// @param request - StartRtcRobotInstanceRequest
//
// @return StartRtcRobotInstanceResponse
func (client *Client) StartRtcRobotInstance(request *StartRtcRobotInstanceRequest) (_result *StartRtcRobotInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRtcRobotInstanceResponse{}
	_body, _err := client.StartRtcRobotInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a workflow task. You can submit a workflow task to implement automated media processing based on a workflow template.
//
// Description:
//
//	  Only media assets from Intelligent Media Services (IMS) or ApsaraVideo VOD can be used as the input of a workflow.
//
//		- When you submit a workflow task, you must specify a workflow template. You can create a workflow template in the [IMS console](https://ims.console.aliyun.com/settings/workflow/list) or use a preset workflow template.
//
// @param request - StartWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartWorkflowResponse
func (client *Client) StartWorkflowWithOptions(request *StartWorkflowRequest, runtime *dara.RuntimeOptions) (_result *StartWorkflowResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SkipInputVerification) {
		query["SkipInputVerification"] = request.SkipInputVerification
	}

	if !dara.IsNil(request.TaskInput) {
		query["TaskInput"] = request.TaskInput
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	if !dara.IsNil(request.WorkflowId) {
		query["WorkflowId"] = request.WorkflowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartWorkflow"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a workflow task. You can submit a workflow task to implement automated media processing based on a workflow template.
//
// Description:
//
//	  Only media assets from Intelligent Media Services (IMS) or ApsaraVideo VOD can be used as the input of a workflow.
//
//		- When you submit a workflow task, you must specify a workflow template. You can create a workflow template in the [IMS console](https://ims.console.aliyun.com/settings/workflow/list) or use a preset workflow template.
//
// @param request - StartWorkflowRequest
//
// @return StartWorkflowResponse
func (client *Client) StartWorkflow(request *StartWorkflowRequest) (_result *StartWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartWorkflowResponse{}
	_body, _err := client.StartWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops an AI agent instance.
//
// Description:
//
//	  When you no longer need an AI agent to participate in a conversation or task, you can call this operation to stop the running agent and release relevant resources.****
//
//		- You must specify the unique ID of the AI agent that you want to stop by using InstanceId.****
//
//		- ****
//
// @param request - StopAIAgentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAIAgentInstanceResponse
func (client *Client) StopAIAgentInstanceWithOptions(request *StopAIAgentInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopAIAgentInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAIAgentInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAIAgentInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops an AI agent instance.
//
// Description:
//
//	  When you no longer need an AI agent to participate in a conversation or task, you can call this operation to stop the running agent and release relevant resources.****
//
//		- You must specify the unique ID of the AI agent that you want to stop by using InstanceId.****
//
//		- ****
//
// @param request - StopAIAgentInstanceRequest
//
// @return StopAIAgentInstanceResponse
func (client *Client) StopAIAgentInstance(request *StopAIAgentInstanceRequest) (_result *StopAIAgentInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopAIAgentInstanceResponse{}
	_body, _err := client.StopAIAgentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a MediaWeaver channel.
//
// @param request - StopChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopChannelResponse
func (client *Client) StopChannelWithOptions(request *StopChannelRequest, runtime *dara.RuntimeOptions) (_result *StopChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a MediaWeaver channel.
//
// @param request - StopChannelRequest
//
// @return StopChannelResponse
func (client *Client) StopChannel(request *StopChannelRequest) (_result *StopChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopChannelResponse{}
	_body, _err := client.StopChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a MediaLive channel.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - StopMediaLiveChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMediaLiveChannelResponse
func (client *Client) StopMediaLiveChannelWithOptions(request *StopMediaLiveChannelRequest, runtime *dara.RuntimeOptions) (_result *StopMediaLiveChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopMediaLiveChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopMediaLiveChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a MediaLive channel.
//
// Description:
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - StopMediaLiveChannelRequest
//
// @return StopMediaLiveChannelResponse
func (client *Client) StopMediaLiveChannel(request *StopMediaLiveChannelRequest) (_result *StopMediaLiveChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopMediaLiveChannelResponse{}
	_body, _err := client.StopMediaLiveChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止一个机器人实例
//
// @param request - StopRtcRobotInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRtcRobotInstanceResponse
func (client *Client) StopRtcRobotInstanceWithOptions(request *StopRtcRobotInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopRtcRobotInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRtcRobotInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRtcRobotInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止一个机器人实例
//
// @param request - StopRtcRobotInstanceRequest
//
// @return StopRtcRobotInstanceResponse
func (client *Client) StopRtcRobotInstance(request *StopRtcRobotInstanceRequest) (_result *StopRtcRobotInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopRtcRobotInstanceResponse{}
	_body, _err := client.StopRtcRobotInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交视频送审任务
//
// @param tmpReq - SubmitAIAgentVideoAuditTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAIAgentVideoAuditTaskResponse
func (client *Client) SubmitAIAgentVideoAuditTaskWithOptions(tmpReq *SubmitAIAgentVideoAuditTaskRequest, runtime *dara.RuntimeOptions) (_result *SubmitAIAgentVideoAuditTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitAIAgentVideoAuditTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallbackConfig) {
		request.CallbackConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallbackConfig, dara.String("CallbackConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CapturePolicies) {
		request.CapturePoliciesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CapturePolicies, dara.String("CapturePolicies"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AIAgentId) {
		query["AIAgentId"] = request.AIAgentId
	}

	if !dara.IsNil(request.AuditInterval) {
		query["AuditInterval"] = request.AuditInterval
	}

	if !dara.IsNil(request.CallbackConfigShrink) {
		query["CallbackConfig"] = request.CallbackConfigShrink
	}

	if !dara.IsNil(request.CapturePoliciesShrink) {
		query["CapturePolicies"] = request.CapturePoliciesShrink
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAIAgentVideoAuditTask"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAIAgentVideoAuditTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频送审任务
//
// @param request - SubmitAIAgentVideoAuditTaskRequest
//
// @return SubmitAIAgentVideoAuditTaskResponse
func (client *Client) SubmitAIAgentVideoAuditTask(request *SubmitAIAgentVideoAuditTaskRequest) (_result *SubmitAIAgentVideoAuditTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAIAgentVideoAuditTaskResponse{}
	_body, _err := client.SubmitAIAgentVideoAuditTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an automatic speech recognition (ASR) job to extract the start and end time and the corresponding text information of a speech in a video.
//
// @param request - SubmitASRJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitASRJobResponse
func (client *Client) SubmitASRJobWithOptions(request *SubmitASRJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitASRJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Duration) {
		query["Duration"] = request.Duration
	}

	if !dara.IsNil(request.EditingConfig) {
		query["EditingConfig"] = request.EditingConfig
	}

	if !dara.IsNil(request.InputFile) {
		query["InputFile"] = request.InputFile
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitASRJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitASRJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an automatic speech recognition (ASR) job to extract the start and end time and the corresponding text information of a speech in a video.
//
// @param request - SubmitASRJobRequest
//
// @return SubmitASRJobResponse
func (client *Client) SubmitASRJob(request *SubmitASRJobRequest) (_result *SubmitASRJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitASRJobResponse{}
	_body, _err := client.SubmitASRJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an audio production job that converts text into an audio file.
//
// @param request - SubmitAudioProduceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAudioProduceJobResponse
func (client *Client) SubmitAudioProduceJobWithOptions(request *SubmitAudioProduceJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAudioProduceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EditingConfig) {
		query["EditingConfig"] = request.EditingConfig
	}

	if !dara.IsNil(request.InputConfig) {
		query["InputConfig"] = request.InputConfig
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.Overwrite) {
		query["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAudioProduceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAudioProduceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an audio production job that converts text into an audio file.
//
// @param request - SubmitAudioProduceJobRequest
//
// @return SubmitAudioProduceJobResponse
func (client *Client) SubmitAudioProduceJob(request *SubmitAudioProduceJobRequest) (_result *SubmitAudioProduceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAudioProduceJobResponse{}
	_body, _err := client.SubmitAudioProduceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a digital human training job. You can call this operation to submit a job the first time or submit a job again with updated parameters if the training failed.
//
// @param request - SubmitAvatarTrainingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAvatarTrainingJobResponse
func (client *Client) SubmitAvatarTrainingJobWithOptions(request *SubmitAvatarTrainingJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAvatarTrainingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAvatarTrainingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAvatarTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a digital human training job. You can call this operation to submit a job the first time or submit a job again with updated parameters if the training failed.
//
// @param request - SubmitAvatarTrainingJobRequest
//
// @return SubmitAvatarTrainingJobResponse
func (client *Client) SubmitAvatarTrainingJob(request *SubmitAvatarTrainingJobRequest) (_result *SubmitAvatarTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAvatarTrainingJobResponse{}
	_body, _err := client.SubmitAvatarTrainingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a video rendering job for a digitized virtual human based on text or an audio file of a human voice.
//
// @param request - SubmitAvatarVideoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAvatarVideoJobResponse
func (client *Client) SubmitAvatarVideoJobWithOptions(request *SubmitAvatarVideoJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitAvatarVideoJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EditingConfig) {
		query["EditingConfig"] = request.EditingConfig
	}

	if !dara.IsNil(request.InputConfig) {
		query["InputConfig"] = request.InputConfig
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAvatarVideoJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAvatarVideoJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video rendering job for a digitized virtual human based on text or an audio file of a human voice.
//
// @param request - SubmitAvatarVideoJobRequest
//
// @return SubmitAvatarVideoJobResponse
func (client *Client) SubmitAvatarVideoJob(request *SubmitAvatarVideoJobRequest) (_result *SubmitAvatarVideoJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitAvatarVideoJobResponse{}
	_body, _err := client.SubmitAvatarVideoJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a quick video production job that intelligently edits multiple video, audio, and image assets to generate multiple videos at a time.
//
// @param request - SubmitBatchMediaProducingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitBatchMediaProducingJobResponse
func (client *Client) SubmitBatchMediaProducingJobWithOptions(request *SubmitBatchMediaProducingJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitBatchMediaProducingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EditingConfig) {
		body["EditingConfig"] = request.EditingConfig
	}

	if !dara.IsNil(request.InputConfig) {
		body["InputConfig"] = request.InputConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitBatchMediaProducingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitBatchMediaProducingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a quick video production job that intelligently edits multiple video, audio, and image assets to generate multiple videos at a time.
//
// @param request - SubmitBatchMediaProducingJobRequest
//
// @return SubmitBatchMediaProducingJobResponse
func (client *Client) SubmitBatchMediaProducingJob(request *SubmitBatchMediaProducingJobRequest) (_result *SubmitBatchMediaProducingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitBatchMediaProducingJobResponse{}
	_body, _err := client.SubmitBatchMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交版权水印提取作业
//
// @param tmpReq - SubmitCopyrightExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCopyrightExtractJobResponse
func (client *Client) SubmitCopyrightExtractJobWithOptions(tmpReq *SubmitCopyrightExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCopyrightExtractJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitCopyrightExtractJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCopyrightExtractJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCopyrightExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交版权水印提取作业
//
// @param request - SubmitCopyrightExtractJobRequest
//
// @return SubmitCopyrightExtractJobResponse
func (client *Client) SubmitCopyrightExtractJob(request *SubmitCopyrightExtractJobRequest) (_result *SubmitCopyrightExtractJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCopyrightExtractJobResponse{}
	_body, _err := client.SubmitCopyrightExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交版权水印任务
//
// @param tmpReq - SubmitCopyrightJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCopyrightJobResponse
func (client *Client) SubmitCopyrightJobWithOptions(tmpReq *SubmitCopyrightJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCopyrightJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitCopyrightJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Message) {
		query["Message"] = request.Message
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TotalTime) {
		query["TotalTime"] = request.TotalTime
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCopyrightJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCopyrightJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交版权水印任务
//
// @param request - SubmitCopyrightJobRequest
//
// @return SubmitCopyrightJobResponse
func (client *Client) SubmitCopyrightJob(request *SubmitCopyrightJobRequest) (_result *SubmitCopyrightJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCopyrightJobResponse{}
	_body, _err := client.SubmitCopyrightJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a human voice cloning job. The value of VoiceId must be the one used during audio check. The system uses this ID to find the cached audio file for training. After you call this operation, the JobId is returned. The training process is asynchronous. During training, you can call the GetCustomizedVoiceJob operation to query information such as the job state.
//
// @param request - SubmitCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitCustomizedVoiceJobResponse
func (client *Client) SubmitCustomizedVoiceJobWithOptions(request *SubmitCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DemoAudioMediaURL) {
		query["DemoAudioMediaURL"] = request.DemoAudioMediaURL
	}

	if !dara.IsNil(request.VoiceId) {
		query["VoiceId"] = request.VoiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a human voice cloning job. The value of VoiceId must be the one used during audio check. The system uses this ID to find the cached audio file for training. After you call this operation, the JobId is returned. The training process is asynchronous. During training, you can call the GetCustomizedVoiceJob operation to query information such as the job state.
//
// @param request - SubmitCustomizedVoiceJobRequest
//
// @return SubmitCustomizedVoiceJobResponse
func (client *Client) SubmitCustomizedVoiceJob(request *SubmitCustomizedVoiceJobRequest) (_result *SubmitCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitCustomizedVoiceJobResponse{}
	_body, _err := client.SubmitCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a media fingerprint analysis job.
//
// Description:
//
//	  SubmitDNAJob is an asynchronous operation. After a request is sent, the system returns a request ID and a job ID and runs the task in the background.
//
//		- You can call this operation only in the China (Beijing), China (Hangzhou), and China (Shanghai) regions.
//
//		- You can submit a text fingerprint analysis job only in the China (Shanghai) region.
//
// @param tmpReq - SubmitDNAJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDNAJobResponse
func (client *Client) SubmitDNAJobWithOptions(tmpReq *SubmitDNAJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDNAJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitDNAJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.DBId) {
		query["DBId"] = request.DBId
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.OwnerAccount) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.PrimaryKey) {
		query["PrimaryKey"] = request.PrimaryKey
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDNAJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDNAJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a media fingerprint analysis job.
//
// Description:
//
//	  SubmitDNAJob is an asynchronous operation. After a request is sent, the system returns a request ID and a job ID and runs the task in the background.
//
//		- You can call this operation only in the China (Beijing), China (Hangzhou), and China (Shanghai) regions.
//
//		- You can submit a text fingerprint analysis job only in the China (Shanghai) region.
//
// @param request - SubmitDNAJobRequest
//
// @return SubmitDNAJobResponse
func (client *Client) SubmitDNAJob(request *SubmitDNAJobRequest) (_result *SubmitDNAJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDNAJobResponse{}
	_body, _err := client.SubmitDNAJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates animated charts based on Excel datasheets, such as line, pie, and bar charts. You can modify the line color and font.
//
// Description:
//
// This feature is available only in the China (Shanghai) region.
//
//   - You can add a title, subtitle, data source, and unit to a chart and specify the font and font size. For supported fonts, see [Fonts](https://help.aliyun.com/document_detail/449567.html).
//
//   - This feature provides five styles of animated charts: normal, mystery, lively, business, and green.
//
//   - You can set the background color or image.
//
//   - You can set the animation duration, size, and bitrate.
//
// Examples
//
//   - Line chart: [Sample datasheet](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/line.xlsx), [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/line.mp4)
//
//   - Bar chart: [Sample datasheet](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/histgram.xlsx), [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/histgram.mp4)
//
//   - Pie chart: [Sample datasheet](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/pie.xlsx), [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/pie.mp4)
//
//   - Normal: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Normal.mp4)
//
//   - Mystery: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Mystery.mp4)
//
//   - Lively: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Lively.mp4)
//
//   - Business: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Business.mp4)
//
//   - Green: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Green.mp4)
//
// @param request - SubmitDynamicChartJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDynamicChartJobResponse
func (client *Client) SubmitDynamicChartJobWithOptions(request *SubmitDynamicChartJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDynamicChartJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AxisParams) {
		query["AxisParams"] = request.AxisParams
	}

	if !dara.IsNil(request.Background) {
		query["Background"] = request.Background
	}

	if !dara.IsNil(request.ChartConfig) {
		query["ChartConfig"] = request.ChartConfig
	}

	if !dara.IsNil(request.ChartTitle) {
		query["ChartTitle"] = request.ChartTitle
	}

	if !dara.IsNil(request.ChartType) {
		query["ChartType"] = request.ChartType
	}

	if !dara.IsNil(request.DataSource) {
		query["DataSource"] = request.DataSource
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.Subtitle) {
		query["Subtitle"] = request.Subtitle
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.Unit) {
		query["Unit"] = request.Unit
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDynamicChartJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDynamicChartJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates animated charts based on Excel datasheets, such as line, pie, and bar charts. You can modify the line color and font.
//
// Description:
//
// This feature is available only in the China (Shanghai) region.
//
//   - You can add a title, subtitle, data source, and unit to a chart and specify the font and font size. For supported fonts, see [Fonts](https://help.aliyun.com/document_detail/449567.html).
//
//   - This feature provides five styles of animated charts: normal, mystery, lively, business, and green.
//
//   - You can set the background color or image.
//
//   - You can set the animation duration, size, and bitrate.
//
// Examples
//
//   - Line chart: [Sample datasheet](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/line.xlsx), [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/line.mp4)
//
//   - Bar chart: [Sample datasheet](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/histgram.xlsx), [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/histgram.mp4)
//
//   - Pie chart: [Sample datasheet](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/pie.xlsx), [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/pie.mp4)
//
//   - Normal: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Normal.mp4)
//
//   - Mystery: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Mystery.mp4)
//
//   - Lively: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Lively.mp4)
//
//   - Business: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Business.mp4)
//
//   - Green: [Effect](https://ice-public-media.oss-cn-shanghai.aliyuncs.com/smart/dynamicChart/Green.mp4)
//
// @param request - SubmitDynamicChartJobRequest
//
// @return SubmitDynamicChartJobResponse
func (client *Client) SubmitDynamicChartJob(request *SubmitDynamicChartJobRequest) (_result *SubmitDynamicChartJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDynamicChartJobResponse{}
	_body, _err := client.SubmitDynamicChartJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an image animation job.
//
// @param tmpReq - SubmitDynamicImageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDynamicImageJobResponse
func (client *Client) SubmitDynamicImageJobWithOptions(tmpReq *SubmitDynamicImageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitDynamicImageJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitDynamicImageJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDynamicImageJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDynamicImageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an image animation job.
//
// @param request - SubmitDynamicImageJobRequest
//
// @return SubmitDynamicImageJobResponse
func (client *Client) SubmitDynamicImageJob(request *SubmitDynamicImageJobRequest) (_result *SubmitDynamicImageJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDynamicImageJobResponse{}
	_body, _err := client.SubmitDynamicImageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a highlight extraction task.
//
// @param request - SubmitHighlightExtractionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHighlightExtractionJobResponse
func (client *Client) SubmitHighlightExtractionJobWithOptions(request *SubmitHighlightExtractionJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitHighlightExtractionJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InputConfig) {
		body["InputConfig"] = request.InputConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitHighlightExtractionJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitHighlightExtractionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a highlight extraction task.
//
// @param request - SubmitHighlightExtractionJobRequest
//
// @return SubmitHighlightExtractionJobResponse
func (client *Client) SubmitHighlightExtractionJob(request *SubmitHighlightExtractionJobRequest) (_result *SubmitHighlightExtractionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitHighlightExtractionJobResponse{}
	_body, _err := client.SubmitHighlightExtractionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an intelligent production job.
//
// @param tmpReq - SubmitIProductionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitIProductionJobResponse
func (client *Client) SubmitIProductionJobWithOptions(tmpReq *SubmitIProductionJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitIProductionJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitIProductionJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.JobParams) {
		query["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitIProductionJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitIProductionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an intelligent production job.
//
// @param request - SubmitIProductionJobRequest
//
// @return SubmitIProductionJobResponse
func (client *Client) SubmitIProductionJob(request *SubmitIProductionJobRequest) (_result *SubmitIProductionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitIProductionJobResponse{}
	_body, _err := client.SubmitIProductionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a live editing job to merge one or more live stream clips into one video. After a live editing job is submitted, the job is queued in the background for asynchronous processing. You can call the GeLiveEditingJob operation to query the state of the job based on the job ID. You can also call the GetMediaInfo operation to query the information about the generated media asset based on the media asset ID.
//
// Description:
//
// Live editing is supported for live streams that are recorded and stored in Object Storage Service (OSS) and ApsaraVideo VOD. If multiple live streams are involved in a single job, only those recorded within the same application are supported for mixed editing. The streams must all be recorded either in OSS or ApsaraVideo VOD.
//
// @param request - SubmitLiveEditingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitLiveEditingJobResponse
func (client *Client) SubmitLiveEditingJobWithOptions(request *SubmitLiveEditingJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitLiveEditingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Clips) {
		query["Clips"] = request.Clips
	}

	if !dara.IsNil(request.LiveStreamConfig) {
		query["LiveStreamConfig"] = request.LiveStreamConfig
	}

	if !dara.IsNil(request.MediaProduceConfig) {
		query["MediaProduceConfig"] = request.MediaProduceConfig
	}

	if !dara.IsNil(request.OutputMediaConfig) {
		query["OutputMediaConfig"] = request.OutputMediaConfig
	}

	if !dara.IsNil(request.OutputMediaTarget) {
		query["OutputMediaTarget"] = request.OutputMediaTarget
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitLiveEditingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitLiveEditingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a live editing job to merge one or more live stream clips into one video. After a live editing job is submitted, the job is queued in the background for asynchronous processing. You can call the GeLiveEditingJob operation to query the state of the job based on the job ID. You can also call the GetMediaInfo operation to query the information about the generated media asset based on the media asset ID.
//
// Description:
//
// Live editing is supported for live streams that are recorded and stored in Object Storage Service (OSS) and ApsaraVideo VOD. If multiple live streams are involved in a single job, only those recorded within the same application are supported for mixed editing. The streams must all be recorded either in OSS or ApsaraVideo VOD.
//
// @param request - SubmitLiveEditingJobRequest
//
// @return SubmitLiveEditingJobResponse
func (client *Client) SubmitLiveEditingJob(request *SubmitLiveEditingJobRequest) (_result *SubmitLiveEditingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitLiveEditingJobResponse{}
	_body, _err := client.SubmitLiveEditingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a live stream recording job.
//
// Description:
//
// You can call this operation to record live streams of ApsaraVideo Live or third-party Real-Time Messaging Protocol (RTMP) live streams. We recommend that you ingest a stream before you call this operation to submit a recording job. If no stream is pulled from the streaming URL, the job attempts to pull a stream for 3 minutes. If the attempt times out, the recording service stops.
//
// Before you submit a recording job, you must prepare an Object Storage Service (OSS) or ApsaraVideo VOD bucket. We recommend that you use a storage address configured in Intelligent Media Services (IMS) to facilitate the management and processing of generated recording files.
//
// If the preset recording template does not meet your requirements, you can create a custom recording template.
//
// @param tmpReq - SubmitLiveRecordJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitLiveRecordJobResponse
func (client *Client) SubmitLiveRecordJobWithOptions(tmpReq *SubmitLiveRecordJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitLiveRecordJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitLiveRecordJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordOutput) {
		request.RecordOutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordOutput, dara.String("RecordOutput"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StreamInput) {
		request.StreamInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamInput, dara.String("StreamInput"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NotifyUrl) {
		body["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.RecordOutputShrink) {
		body["RecordOutput"] = request.RecordOutputShrink
	}

	if !dara.IsNil(request.StreamInputShrink) {
		body["StreamInput"] = request.StreamInputShrink
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitLiveRecordJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitLiveRecordJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a live stream recording job.
//
// Description:
//
// You can call this operation to record live streams of ApsaraVideo Live or third-party Real-Time Messaging Protocol (RTMP) live streams. We recommend that you ingest a stream before you call this operation to submit a recording job. If no stream is pulled from the streaming URL, the job attempts to pull a stream for 3 minutes. If the attempt times out, the recording service stops.
//
// Before you submit a recording job, you must prepare an Object Storage Service (OSS) or ApsaraVideo VOD bucket. We recommend that you use a storage address configured in Intelligent Media Services (IMS) to facilitate the management and processing of generated recording files.
//
// If the preset recording template does not meet your requirements, you can create a custom recording template.
//
// @param request - SubmitLiveRecordJobRequest
//
// @return SubmitLiveRecordJobResponse
func (client *Client) SubmitLiveRecordJob(request *SubmitLiveRecordJobRequest) (_result *SubmitLiveRecordJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitLiveRecordJobResponse{}
	_body, _err := client.SubmitLiveRecordJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a live stream snapshot job. If the job is submitted during stream ingest, it automatically starts in asynchronous mode. Otherwise, it does not start.
//
// @param tmpReq - SubmitLiveSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitLiveSnapshotJobResponse
func (client *Client) SubmitLiveSnapshotJobWithOptions(tmpReq *SubmitLiveSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitLiveSnapshotJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitLiveSnapshotJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SnapshotOutput) {
		request.SnapshotOutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SnapshotOutput, dara.String("SnapshotOutput"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StreamInput) {
		request.StreamInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamInput, dara.String("StreamInput"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallbackUrl) {
		body["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.JobName) {
		body["JobName"] = request.JobName
	}

	if !dara.IsNil(request.SnapshotOutputShrink) {
		body["SnapshotOutput"] = request.SnapshotOutputShrink
	}

	if !dara.IsNil(request.StreamInputShrink) {
		body["StreamInput"] = request.StreamInputShrink
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitLiveSnapshotJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitLiveSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a live stream snapshot job. If the job is submitted during stream ingest, it automatically starts in asynchronous mode. Otherwise, it does not start.
//
// @param request - SubmitLiveSnapshotJobRequest
//
// @return SubmitLiveSnapshotJobResponse
func (client *Client) SubmitLiveSnapshotJob(request *SubmitLiveSnapshotJobRequest) (_result *SubmitLiveSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitLiveSnapshotJobResponse{}
	_body, _err := client.SubmitLiveSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a live stream transcoding job.
//
// Description:
//
//	  When you submit a transcoding job that immediately takes effect, make sure that the input stream can be streamed.
//
//		- When you submit a timed transcoding job, make sure that the input stream can be streamed before the specified time.
//
// @param tmpReq - SubmitLiveTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitLiveTranscodeJobResponse
func (client *Client) SubmitLiveTranscodeJobWithOptions(tmpReq *SubmitLiveTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitLiveTranscodeJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitLiveTranscodeJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamInput) {
		request.StreamInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamInput, dara.String("StreamInput"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TimedConfig) {
		request.TimedConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TimedConfig, dara.String("TimedConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TranscodeOutput) {
		request.TranscodeOutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TranscodeOutput, dara.String("TranscodeOutput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.StartMode) {
		query["StartMode"] = request.StartMode
	}

	if !dara.IsNil(request.StreamInputShrink) {
		query["StreamInput"] = request.StreamInputShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TimedConfigShrink) {
		query["TimedConfig"] = request.TimedConfigShrink
	}

	if !dara.IsNil(request.TranscodeOutputShrink) {
		query["TranscodeOutput"] = request.TranscodeOutputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitLiveTranscodeJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitLiveTranscodeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a live stream transcoding job.
//
// Description:
//
//	  When you submit a transcoding job that immediately takes effect, make sure that the input stream can be streamed.
//
//		- When you submit a timed transcoding job, make sure that the input stream can be streamed before the specified time.
//
// @param request - SubmitLiveTranscodeJobRequest
//
// @return SubmitLiveTranscodeJobResponse
func (client *Client) SubmitLiveTranscodeJob(request *SubmitLiveTranscodeJobRequest) (_result *SubmitLiveTranscodeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitLiveTranscodeJobResponse{}
	_body, _err := client.SubmitLiveTranscodeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a structural analysis job for a media asset. For example, you can submit a job to analyze the speaker, translate the video, and obtain the paragraph summary.
//
// @param request - SubmitMediaAiAnalysisJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaAiAnalysisJobResponse
func (client *Client) SubmitMediaAiAnalysisJobWithOptions(request *SubmitMediaAiAnalysisJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaAiAnalysisJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisParams) {
		query["AnalysisParams"] = request.AnalysisParams
	}

	if !dara.IsNil(request.Input) {
		query["Input"] = request.Input
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaAiAnalysisJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaAiAnalysisJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a structural analysis job for a media asset. For example, you can submit a job to analyze the speaker, translate the video, and obtain the paragraph summary.
//
// @param request - SubmitMediaAiAnalysisJobRequest
//
// @return SubmitMediaAiAnalysisJobResponse
func (client *Client) SubmitMediaAiAnalysisJob(request *SubmitMediaAiAnalysisJobRequest) (_result *SubmitMediaAiAnalysisJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMediaAiAnalysisJobResponse{}
	_body, _err := client.SubmitMediaAiAnalysisJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a content moderation job.
//
// Description:
//
// The job that you submit by calling this operation is run in asynchronous mode. The job is added to an ApsaraVideo Media Processing (MPS) queue to be scheduled and run. You can call the [QueryMediaCensorJobDetail](https://help.aliyun.com/document_detail/444847.html) operation or configure an asynchronous notification to obtain the job results.
//
// @param tmpReq - SubmitMediaCensorJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaCensorJobResponse
func (client *Client) SubmitMediaCensorJobWithOptions(tmpReq *SubmitMediaCensorJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaCensorJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitMediaCensorJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Barrages) {
		query["Barrages"] = request.Barrages
	}

	if !dara.IsNil(request.CoverImages) {
		query["CoverImages"] = request.CoverImages
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.Output) {
		query["Output"] = request.Output
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaCensorJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaCensorJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a content moderation job.
//
// Description:
//
// The job that you submit by calling this operation is run in asynchronous mode. The job is added to an ApsaraVideo Media Processing (MPS) queue to be scheduled and run. You can call the [QueryMediaCensorJobDetail](https://help.aliyun.com/document_detail/444847.html) operation or configure an asynchronous notification to obtain the job results.
//
// @param request - SubmitMediaCensorJobRequest
//
// @return SubmitMediaCensorJobResponse
func (client *Client) SubmitMediaCensorJob(request *SubmitMediaCensorJobRequest) (_result *SubmitMediaCensorJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMediaCensorJobResponse{}
	_body, _err := client.SubmitMediaCensorJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a transcoding task.
//
// @param request - SubmitMediaConvertJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaConvertJobResponse
func (client *Client) SubmitMediaConvertJobWithOptions(request *SubmitMediaConvertJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaConvertJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Config) {
		query["Config"] = request.Config
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaConvertJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaConvertJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a transcoding task.
//
// @param request - SubmitMediaConvertJobRequest
//
// @return SubmitMediaConvertJobResponse
func (client *Client) SubmitMediaConvertJob(request *SubmitMediaConvertJobRequest) (_result *SubmitMediaConvertJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMediaConvertJobResponse{}
	_body, _err := client.SubmitMediaConvertJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a media information analysis job in asynchronous mode.
//
// Description:
//
// You can call this operation to analyze an input media file by using a callback mechanism or initiating subsequent queries. This operation is suitable for scenarios in which real-time performance is less critical and high concurrency is expected.
//
// @param tmpReq - SubmitMediaInfoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaInfoJobResponse
func (client *Client) SubmitMediaInfoJobWithOptions(tmpReq *SubmitMediaInfoJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaInfoJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitMediaInfoJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaInfoJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaInfoJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a media information analysis job in asynchronous mode.
//
// Description:
//
// You can call this operation to analyze an input media file by using a callback mechanism or initiating subsequent queries. This operation is suitable for scenarios in which real-time performance is less critical and high concurrency is expected.
//
// @param request - SubmitMediaInfoJobRequest
//
// @return SubmitMediaInfoJobResponse
func (client *Client) SubmitMediaInfoJob(request *SubmitMediaInfoJobRequest) (_result *SubmitMediaInfoJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMediaInfoJobResponse{}
	_body, _err := client.SubmitMediaInfoJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a media editing and production job. If you need to perform any form of post-production such as editing and production on video or audio materials, you can call this operation to automate the process.
//
// Description:
//
//	  This operation returns only the submission result of a media editing and production job. When the submission result is returned, the job may still be in progress. After a media editing and production job is submitted, the job is queued in the background for asynchronous processing.
//
//		- The materials referenced in the timeline of an online editing project can be media assets in the media asset library or Object Storage Service (OSS) objects. External URLs or Alibaba Cloud Content Delivery Network (CDN) URLs are not supported. To use an OSS object as a material, you must set MediaUrl to an OSS URL, such as https://your-bucket.oss-region-name.aliyuncs.com/your-object.ext.
//
//		- After the production is complete, the output file is automatically registered as a media asset. The media asset first needs to be analyzed. After the media asset is analyzed, you can query the duration and resolution information based on the media asset ID.
//
// ## [](#)Limits
//
//   - The throttling threshold of this operation is 30 queries per second (QPS).
//
//     **
//
//     **Note*	- If the threshold is exceeded, a "Throttling.User" error is returned when you submit an editing job. For more information about how to resolve this issue, see the [FAQ](https://help.aliyun.com/document_detail/453484.html).
//
//   - You can create up to 100 video tracks, 100 image tracks, and 100 subtitle tracks in a project.
//
//   - The total size of material files cannot exceed 1 TB.
//
//   - The OSS buckets in which the materials reside and where the output media assets are stored must be in the same region as the region in which Intelligent Media Services (IMS) is activated.
//
//   - An output video must meet the following requirements:
//
//   - Both the width and height must be at least 128 pixels.
//
//   - Both the width and height cannot exceed 4,096 pixels.
//
//   - The shorter side of the video cannot exceed 2,160 pixels.
//
// @param request - SubmitMediaProducingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitMediaProducingJobResponse
func (client *Client) SubmitMediaProducingJobWithOptions(request *SubmitMediaProducingJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitMediaProducingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ClipsParam) {
		query["ClipsParam"] = request.ClipsParam
	}

	if !dara.IsNil(request.EditingProduceConfig) {
		query["EditingProduceConfig"] = request.EditingProduceConfig
	}

	if !dara.IsNil(request.MediaMetadata) {
		query["MediaMetadata"] = request.MediaMetadata
	}

	if !dara.IsNil(request.OutputMediaConfig) {
		query["OutputMediaConfig"] = request.OutputMediaConfig
	}

	if !dara.IsNil(request.OutputMediaTarget) {
		query["OutputMediaTarget"] = request.OutputMediaTarget
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectMetadata) {
		query["ProjectMetadata"] = request.ProjectMetadata
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Timeline) {
		body["Timeline"] = request.Timeline
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitMediaProducingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitMediaProducingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a media editing and production job. If you need to perform any form of post-production such as editing and production on video or audio materials, you can call this operation to automate the process.
//
// Description:
//
//	  This operation returns only the submission result of a media editing and production job. When the submission result is returned, the job may still be in progress. After a media editing and production job is submitted, the job is queued in the background for asynchronous processing.
//
//		- The materials referenced in the timeline of an online editing project can be media assets in the media asset library or Object Storage Service (OSS) objects. External URLs or Alibaba Cloud Content Delivery Network (CDN) URLs are not supported. To use an OSS object as a material, you must set MediaUrl to an OSS URL, such as https://your-bucket.oss-region-name.aliyuncs.com/your-object.ext.
//
//		- After the production is complete, the output file is automatically registered as a media asset. The media asset first needs to be analyzed. After the media asset is analyzed, you can query the duration and resolution information based on the media asset ID.
//
// ## [](#)Limits
//
//   - The throttling threshold of this operation is 30 queries per second (QPS).
//
//     **
//
//     **Note*	- If the threshold is exceeded, a "Throttling.User" error is returned when you submit an editing job. For more information about how to resolve this issue, see the [FAQ](https://help.aliyun.com/document_detail/453484.html).
//
//   - You can create up to 100 video tracks, 100 image tracks, and 100 subtitle tracks in a project.
//
//   - The total size of material files cannot exceed 1 TB.
//
//   - The OSS buckets in which the materials reside and where the output media assets are stored must be in the same region as the region in which Intelligent Media Services (IMS) is activated.
//
//   - An output video must meet the following requirements:
//
//   - Both the width and height must be at least 128 pixels.
//
//   - Both the width and height cannot exceed 4,096 pixels.
//
//   - The shorter side of the video cannot exceed 2,160 pixels.
//
// @param request - SubmitMediaProducingJobRequest
//
// @return SubmitMediaProducingJobResponse
func (client *Client) SubmitMediaProducingJob(request *SubmitMediaProducingJobRequest) (_result *SubmitMediaProducingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitMediaProducingJobResponse{}
	_body, _err := client.SubmitMediaProducingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a packaging job.
//
// @param tmpReq - SubmitPackageJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitPackageJobResponse
func (client *Client) SubmitPackageJobWithOptions(tmpReq *SubmitPackageJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitPackageJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitPackageJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Inputs) {
		request.InputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Inputs, dara.String("Inputs"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputsShrink) {
		query["Inputs"] = request.InputsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitPackageJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitPackageJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a packaging job.
//
// @param request - SubmitPackageJobRequest
//
// @return SubmitPackageJobResponse
func (client *Client) SubmitPackageJob(request *SubmitPackageJobRequest) (_result *SubmitPackageJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitPackageJobResponse{}
	_body, _err := client.SubmitPackageJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交工程导出任务
//
// @param request - SubmitProjectExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitProjectExportJobResponse
func (client *Client) SubmitProjectExportJobWithOptions(request *SubmitProjectExportJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitProjectExportJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExportType) {
		query["ExportType"] = request.ExportType
	}

	if !dara.IsNil(request.OutputMediaConfig) {
		query["OutputMediaConfig"] = request.OutputMediaConfig
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Timeline) {
		body["Timeline"] = request.Timeline
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitProjectExportJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitProjectExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交工程导出任务
//
// @param request - SubmitProjectExportJobRequest
//
// @return SubmitProjectExportJobResponse
func (client *Client) SubmitProjectExportJob(request *SubmitProjectExportJobRequest) (_result *SubmitProjectExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitProjectExportJobResponse{}
	_body, _err := client.SubmitProjectExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a task to automatically recognize the highlight segments in the video input and compile them into a dramatic and engaging clip.
//
// @param request - SubmitScreenMediaHighlightsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitScreenMediaHighlightsJobResponse
func (client *Client) SubmitScreenMediaHighlightsJobWithOptions(request *SubmitScreenMediaHighlightsJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitScreenMediaHighlightsJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EditingConfig) {
		body["EditingConfig"] = request.EditingConfig
	}

	if !dara.IsNil(request.InputConfig) {
		body["InputConfig"] = request.InputConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitScreenMediaHighlightsJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitScreenMediaHighlightsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a task to automatically recognize the highlight segments in the video input and compile them into a dramatic and engaging clip.
//
// @param request - SubmitScreenMediaHighlightsJobRequest
//
// @return SubmitScreenMediaHighlightsJobResponse
func (client *Client) SubmitScreenMediaHighlightsJob(request *SubmitScreenMediaHighlightsJobRequest) (_result *SubmitScreenMediaHighlightsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitScreenMediaHighlightsJobResponse{}
	_body, _err := client.SubmitScreenMediaHighlightsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交拆条任务
//
// @param request - SubmitSegmentationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSegmentationJobResponse
func (client *Client) SubmitSegmentationJobWithOptions(request *SubmitSegmentationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSegmentationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.JobParams) {
		query["JobParams"] = request.JobParams
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InputConfig) {
		body["InputConfig"] = request.InputConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSegmentationJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSegmentationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交拆条任务
//
// @param request - SubmitSegmentationJobRequest
//
// @return SubmitSegmentationJobResponse
func (client *Client) SubmitSegmentationJob(request *SubmitSegmentationJobRequest) (_result *SubmitSegmentationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSegmentationJobResponse{}
	_body, _err := client.SubmitSegmentationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a smart tagging job.
//
// Description:
//
// Before you call this operation to submit a smart tagging job, you must add a smart tagging template and specify the analysis types that you want to use in the template. For more information, see CreateCustomTemplate. You can use the smart tagging feature only in the China (Beijing), China (Shanghai), and China (Hangzhou) regions. By default, an ApsaraVideo Media Processing (MPS) queue can process a maximum of two concurrent smart tagging jobs. If you need to process more concurrent smart tagging jobs, submit a ticket to contact Alibaba Cloud Technical Support for evaluation and configuration.
//
// @param tmpReq - SubmitSmarttagJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSmarttagJobResponse
func (client *Client) SubmitSmarttagJobWithOptions(tmpReq *SubmitSmarttagJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSmarttagJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitSmarttagJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentAddr) {
		query["ContentAddr"] = request.ContentAddr
	}

	if !dara.IsNil(request.ContentType) {
		query["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.NotifyUrl) {
		query["NotifyUrl"] = request.NotifyUrl
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSmarttagJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSmarttagJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a smart tagging job.
//
// Description:
//
// Before you call this operation to submit a smart tagging job, you must add a smart tagging template and specify the analysis types that you want to use in the template. For more information, see CreateCustomTemplate. You can use the smart tagging feature only in the China (Beijing), China (Shanghai), and China (Hangzhou) regions. By default, an ApsaraVideo Media Processing (MPS) queue can process a maximum of two concurrent smart tagging jobs. If you need to process more concurrent smart tagging jobs, submit a ticket to contact Alibaba Cloud Technical Support for evaluation and configuration.
//
// @param request - SubmitSmarttagJobRequest
//
// @return SubmitSmarttagJobResponse
func (client *Client) SubmitSmarttagJob(request *SubmitSmarttagJobRequest) (_result *SubmitSmarttagJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSmarttagJobResponse{}
	_body, _err := client.SubmitSmarttagJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a snapshot job.
//
// @param tmpReq - SubmitSnapshotJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSnapshotJobResponse
func (client *Client) SubmitSnapshotJobWithOptions(tmpReq *SubmitSnapshotJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSnapshotJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitSnapshotJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSnapshotJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSnapshotJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a snapshot job.
//
// @param request - SubmitSnapshotJobRequest
//
// @return SubmitSnapshotJobResponse
func (client *Client) SubmitSnapshotJob(request *SubmitSnapshotJobRequest) (_result *SubmitSnapshotJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSnapshotJobResponse{}
	_body, _err := client.SubmitSnapshotJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a sports highlights job to generate a highlights video of an event based on event materials that contain commentary.
//
// @param request - SubmitSportsHighlightsJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSportsHighlightsJobResponse
func (client *Client) SubmitSportsHighlightsJobWithOptions(request *SubmitSportsHighlightsJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSportsHighlightsJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InputConfig) {
		body["InputConfig"] = request.InputConfig
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSportsHighlightsJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSportsHighlightsJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a sports highlights job to generate a highlights video of an event based on event materials that contain commentary.
//
// @param request - SubmitSportsHighlightsJobRequest
//
// @return SubmitSportsHighlightsJobResponse
func (client *Client) SubmitSportsHighlightsJob(request *SubmitSportsHighlightsJobRequest) (_result *SubmitSportsHighlightsJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSportsHighlightsJobResponse{}
	_body, _err := client.SubmitSportsHighlightsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a standard human voice cloning job. After you call this operation, the JobId is returned. The training process is asynchronous. During training, you can call the GetCustomizedVoiceJob operation to query information such as the job state.
//
// @param request - SubmitStandardCustomizedVoiceJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitStandardCustomizedVoiceJobResponse
func (client *Client) SubmitStandardCustomizedVoiceJobWithOptions(request *SubmitStandardCustomizedVoiceJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitStandardCustomizedVoiceJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Audios) {
		query["Audios"] = request.Audios
	}

	if !dara.IsNil(request.Authentication) {
		query["Authentication"] = request.Authentication
	}

	if !dara.IsNil(request.DemoAudioMediaURL) {
		query["DemoAudioMediaURL"] = request.DemoAudioMediaURL
	}

	if !dara.IsNil(request.Gender) {
		query["Gender"] = request.Gender
	}

	if !dara.IsNil(request.VoiceName) {
		query["VoiceName"] = request.VoiceName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitStandardCustomizedVoiceJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitStandardCustomizedVoiceJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a standard human voice cloning job. After you call this operation, the JobId is returned. The training process is asynchronous. During training, you can call the GetCustomizedVoiceJob operation to query information such as the job state.
//
// @param request - SubmitStandardCustomizedVoiceJobRequest
//
// @return SubmitStandardCustomizedVoiceJobResponse
func (client *Client) SubmitStandardCustomizedVoiceJob(request *SubmitStandardCustomizedVoiceJobRequest) (_result *SubmitStandardCustomizedVoiceJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitStandardCustomizedVoiceJobResponse{}
	_body, _err := client.SubmitStandardCustomizedVoiceJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a media file in synchronous mode for media information analysis.
//
// Description:
//
// You can call this operation to analyze an input media file in synchronous mode. This operation is suitable for scenarios that require high real-time performance and low concurrency. If it takes an extended period of time to obtain the media information about the input media file, the request may time out or the obtained information may be inaccurate. We recommend that you call the [SubmitMediaInfoJob](https://help.aliyun.com/document_detail/441222.html) operation to obtain media information.
//
// @param tmpReq - SubmitSyncMediaInfoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitSyncMediaInfoJobResponse
func (client *Client) SubmitSyncMediaInfoJobWithOptions(tmpReq *SubmitSyncMediaInfoJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitSyncMediaInfoJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitSyncMediaInfoJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitSyncMediaInfoJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitSyncMediaInfoJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a media file in synchronous mode for media information analysis.
//
// Description:
//
// You can call this operation to analyze an input media file in synchronous mode. This operation is suitable for scenarios that require high real-time performance and low concurrency. If it takes an extended period of time to obtain the media information about the input media file, the request may time out or the obtained information may be inaccurate. We recommend that you call the [SubmitMediaInfoJob](https://help.aliyun.com/document_detail/441222.html) operation to obtain media information.
//
// @param request - SubmitSyncMediaInfoJobRequest
//
// @return SubmitSyncMediaInfoJobResponse
func (client *Client) SubmitSyncMediaInfoJob(request *SubmitSyncMediaInfoJobRequest) (_result *SubmitSyncMediaInfoJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitSyncMediaInfoJobResponse{}
	_body, _err := client.SubmitSyncMediaInfoJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a text generation job to generate marketing copies based on keywords and the requirements for the word count and number of output copies. The word count of the output copies may differ from the specified word count. After the job is submitted, you can call the GetSmartHandleJob operation to obtain the job state and result based on the job ID.
//
// @param request - SubmitTextGenerateJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTextGenerateJobResponse
func (client *Client) SubmitTextGenerateJobWithOptions(request *SubmitTextGenerateJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTextGenerateJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GenerateConfig) {
		query["GenerateConfig"] = request.GenerateConfig
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTextGenerateJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTextGenerateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a text generation job to generate marketing copies based on keywords and the requirements for the word count and number of output copies. The word count of the output copies may differ from the specified word count. After the job is submitted, you can call the GetSmartHandleJob operation to obtain the job state and result based on the job ID.
//
// @param request - SubmitTextGenerateJobRequest
//
// @return SubmitTextGenerateJobResponse
func (client *Client) SubmitTextGenerateJob(request *SubmitTextGenerateJobRequest) (_result *SubmitTextGenerateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTextGenerateJobResponse{}
	_body, _err := client.SubmitTextGenerateJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交视频溯源水印ab流任务
//
// @param tmpReq - SubmitTraceAbJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTraceAbJobResponse
func (client *Client) SubmitTraceAbJobWithOptions(tmpReq *SubmitTraceAbJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTraceAbJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTraceAbJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CipherBase64ed) {
		query["CipherBase64ed"] = request.CipherBase64ed
	}

	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TotalTime) {
		query["TotalTime"] = request.TotalTime
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTraceAbJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTraceAbJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频溯源水印ab流任务
//
// @param request - SubmitTraceAbJobRequest
//
// @return SubmitTraceAbJobResponse
func (client *Client) SubmitTraceAbJob(request *SubmitTraceAbJobRequest) (_result *SubmitTraceAbJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTraceAbJobResponse{}
	_body, _err := client.SubmitTraceAbJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交溯源水印提取任务
//
// @param tmpReq - SubmitTraceExtractJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTraceExtractJobResponse
func (client *Client) SubmitTraceExtractJobWithOptions(tmpReq *SubmitTraceExtractJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTraceExtractJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTraceExtractJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTraceExtractJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTraceExtractJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交溯源水印提取任务
//
// @param request - SubmitTraceExtractJobRequest
//
// @return SubmitTraceExtractJobResponse
func (client *Client) SubmitTraceExtractJob(request *SubmitTraceExtractJobRequest) (_result *SubmitTraceExtractJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTraceExtractJobResponse{}
	_body, _err := client.SubmitTraceExtractJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交视频溯源水印m3u8文件任务
//
// @param tmpReq - SubmitTraceM3u8JobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTraceM3u8JobResponse
func (client *Client) SubmitTraceM3u8JobWithOptions(tmpReq *SubmitTraceM3u8JobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTraceM3u8JobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTraceM3u8JobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Output) {
		request.OutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Output, dara.String("Output"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyUri) {
		query["KeyUri"] = request.KeyUri
	}

	if !dara.IsNil(request.OutputShrink) {
		query["Output"] = request.OutputShrink
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.Trace) {
		query["Trace"] = request.Trace
	}

	if !dara.IsNil(request.TraceMediaId) {
		query["TraceMediaId"] = request.TraceMediaId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTraceM3u8Job"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTraceM3u8JobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频溯源水印m3u8文件任务
//
// @param request - SubmitTraceM3u8JobRequest
//
// @return SubmitTraceM3u8JobResponse
func (client *Client) SubmitTraceM3u8Job(request *SubmitTraceM3u8JobRequest) (_result *SubmitTraceM3u8JobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTraceM3u8JobResponse{}
	_body, _err := client.SubmitTraceM3u8JobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a transcoding job.
//
// @param tmpReq - SubmitTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTranscodeJobResponse
func (client *Client) SubmitTranscodeJobWithOptions(tmpReq *SubmitTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitTranscodeJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitTranscodeJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputGroup) {
		request.InputGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputGroup, dara.String("InputGroup"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutputGroup) {
		request.OutputGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputGroup, dara.String("OutputGroup"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ScheduleConfig) {
		request.ScheduleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScheduleConfig, dara.String("ScheduleConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.InputGroupShrink) {
		query["InputGroup"] = request.InputGroupShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputGroupShrink) {
		query["OutputGroup"] = request.OutputGroupShrink
	}

	if !dara.IsNil(request.ScheduleConfigShrink) {
		query["ScheduleConfig"] = request.ScheduleConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTranscodeJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTranscodeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a transcoding job.
//
// @param request - SubmitTranscodeJobRequest
//
// @return SubmitTranscodeJobResponse
func (client *Client) SubmitTranscodeJob(request *SubmitTranscodeJobRequest) (_result *SubmitTranscodeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitTranscodeJobResponse{}
	_body, _err := client.SubmitTranscodeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交视频理解任务
//
// @param tmpReq - SubmitVideoCognitionJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoCognitionJobResponse
func (client *Client) SubmitVideoCognitionJobWithOptions(tmpReq *SubmitVideoCognitionJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoCognitionJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SubmitVideoCognitionJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputShrink) {
		query["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoCognitionJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoCognitionJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频理解任务
//
// @param request - SubmitVideoCognitionJobRequest
//
// @return SubmitVideoCognitionJobResponse
func (client *Client) SubmitVideoCognitionJob(request *SubmitVideoCognitionJobRequest) (_result *SubmitVideoCognitionJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitVideoCognitionJobResponse{}
	_body, _err := client.SubmitVideoCognitionJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a video translation job. You can call this operation to translate subtitles in a video and audio to a specific language. Lip-sync adaptation will be supported in the future.
//
// Description:
//
// After you call this operation to submit a video translation job, the system returns a job ID. You can call the GetSmartHandleJob operation based on the job ID to obtain the status and result information of the job.
//
// @param request - SubmitVideoTranslationJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVideoTranslationJobResponse
func (client *Client) SubmitVideoTranslationJobWithOptions(request *SubmitVideoTranslationJobRequest, runtime *dara.RuntimeOptions) (_result *SubmitVideoTranslationJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EditingConfig) {
		query["EditingConfig"] = request.EditingConfig
	}

	if !dara.IsNil(request.InputConfig) {
		query["InputConfig"] = request.InputConfig
	}

	if !dara.IsNil(request.OutputConfig) {
		query["OutputConfig"] = request.OutputConfig
	}

	if !dara.IsNil(request.Signature) {
		query["Signature"] = request.Signature
	}

	if !dara.IsNil(request.SignatureMehtod) {
		query["SignatureMehtod"] = request.SignatureMehtod
	}

	if !dara.IsNil(request.SignatureNonce) {
		query["SignatureNonce"] = request.SignatureNonce
	}

	if !dara.IsNil(request.SignatureType) {
		query["SignatureType"] = request.SignatureType
	}

	if !dara.IsNil(request.SignatureVersion) {
		query["SignatureVersion"] = request.SignatureVersion
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVideoTranslationJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVideoTranslationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a video translation job. You can call this operation to translate subtitles in a video and audio to a specific language. Lip-sync adaptation will be supported in the future.
//
// Description:
//
// After you call this operation to submit a video translation job, the system returns a job ID. You can call the GetSmartHandleJob operation based on the job ID to obtain the status and result information of the job.
//
// @param request - SubmitVideoTranslationJobRequest
//
// @return SubmitVideoTranslationJobResponse
func (client *Client) SubmitVideoTranslationJob(request *SubmitVideoTranslationJobRequest) (_result *SubmitVideoTranslationJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitVideoTranslationJobResponse{}
	_body, _err := client.SubmitVideoTranslationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Hands off a conversation to a human agent.
//
// @param request - TakeoverAIAgentCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TakeoverAIAgentCallResponse
func (client *Client) TakeoverAIAgentCallWithOptions(request *TakeoverAIAgentCallRequest, runtime *dara.RuntimeOptions) (_result *TakeoverAIAgentCallResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HumanAgentUserId) {
		query["HumanAgentUserId"] = request.HumanAgentUserId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RequireToken) {
		query["RequireToken"] = request.RequireToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TakeoverAIAgentCall"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TakeoverAIAgentCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Hands off a conversation to a human agent.
//
// @param request - TakeoverAIAgentCallRequest
//
// @return TakeoverAIAgentCallResponse
func (client *Client) TakeoverAIAgentCall(request *TakeoverAIAgentCallRequest) (_result *TakeoverAIAgentCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TakeoverAIAgentCallResponse{}
	_body, _err := client.TakeoverAIAgentCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of an AI agent.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to update the configurations of an AI agent, such as the tone, by specifying the agent ID and configurations.
//
// @param tmpReq - UpdateAIAgentInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAIAgentInstanceResponse
func (client *Client) UpdateAIAgentInstanceWithOptions(tmpReq *UpdateAIAgentInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateAIAgentInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAIAgentInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentConfig) {
		request.AgentConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentConfig, dara.String("AgentConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentConfigShrink) {
		query["AgentConfig"] = request.AgentConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAIAgentInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAIAgentInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of an AI agent.
//
// Description:
//
// ## [](#)Request description
//
// You can call this operation to update the configurations of an AI agent, such as the tone, by specifying the agent ID and configurations.
//
// @param request - UpdateAIAgentInstanceRequest
//
// @return UpdateAIAgentInstanceResponse
func (client *Client) UpdateAIAgentInstance(request *UpdateAIAgentInstanceRequest) (_result *UpdateAIAgentInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAIAgentInstanceResponse{}
	_body, _err := client.UpdateAIAgentInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an ad insertion configuration.
//
// @param request - UpdateAdInsertionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAdInsertionResponse
func (client *Client) UpdateAdInsertionWithOptions(request *UpdateAdInsertionRequest, runtime *dara.RuntimeOptions) (_result *UpdateAdInsertionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdMarkerPassthrough) {
		body["AdMarkerPassthrough"] = request.AdMarkerPassthrough
	}

	if !dara.IsNil(request.AdsUrl) {
		body["AdsUrl"] = request.AdsUrl
	}

	if !dara.IsNil(request.CdnAdSegmentUrlPrefix) {
		body["CdnAdSegmentUrlPrefix"] = request.CdnAdSegmentUrlPrefix
	}

	if !dara.IsNil(request.CdnContentSegmentUrlPrefix) {
		body["CdnContentSegmentUrlPrefix"] = request.CdnContentSegmentUrlPrefix
	}

	if !dara.IsNil(request.ConfigAliases) {
		body["ConfigAliases"] = request.ConfigAliases
	}

	if !dara.IsNil(request.ContentUrlPrefix) {
		body["ContentUrlPrefix"] = request.ContentUrlPrefix
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PersonalizationThreshold) {
		body["PersonalizationThreshold"] = request.PersonalizationThreshold
	}

	if !dara.IsNil(request.SlateAdUrl) {
		body["SlateAdUrl"] = request.SlateAdUrl
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAdInsertion"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAdInsertionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an ad insertion configuration.
//
// @param request - UpdateAdInsertionRequest
//
// @return UpdateAdInsertionResponse
func (client *Client) UpdateAdInsertion(request *UpdateAdInsertionRequest) (_result *UpdateAdInsertionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAdInsertionResponse{}
	_body, _err := client.UpdateAdInsertionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a digital human training job. You can modify the basic information or update parameters such as Video and Transparent for retraining if the training failed.
//
// @param request - UpdateAvatarTrainingJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAvatarTrainingJobResponse
func (client *Client) UpdateAvatarTrainingJobWithOptions(request *UpdateAvatarTrainingJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateAvatarTrainingJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AvatarDescription) {
		query["AvatarDescription"] = request.AvatarDescription
	}

	if !dara.IsNil(request.AvatarName) {
		query["AvatarName"] = request.AvatarName
	}

	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Portrait) {
		query["Portrait"] = request.Portrait
	}

	if !dara.IsNil(request.Thumbnail) {
		query["Thumbnail"] = request.Thumbnail
	}

	if !dara.IsNil(request.Transparent) {
		query["Transparent"] = request.Transparent
	}

	if !dara.IsNil(request.Video) {
		query["Video"] = request.Video
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAvatarTrainingJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAvatarTrainingJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a digital human training job. You can modify the basic information or update parameters such as Video and Transparent for retraining if the training failed.
//
// @param request - UpdateAvatarTrainingJobRequest
//
// @return UpdateAvatarTrainingJobResponse
func (client *Client) UpdateAvatarTrainingJob(request *UpdateAvatarTrainingJobRequest) (_result *UpdateAvatarTrainingJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAvatarTrainingJobResponse{}
	_body, _err := client.UpdateAvatarTrainingJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a category.
//
// Description:
//
// After you create a media asset category, you can call this operation to find the category based on the category ID and change the name of the category.
//
// @param request - UpdateCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategoryWithOptions(request *UpdateCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.CateName) {
		query["CateName"] = request.CateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCategory"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a category.
//
// Description:
//
// After you create a media asset category, you can call this operation to find the category based on the category ID and change the name of the category.
//
// @param request - UpdateCategoryRequest
//
// @return UpdateCategoryResponse
func (client *Client) UpdateCategory(request *UpdateCategoryRequest) (_result *UpdateCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCategoryResponse{}
	_body, _err := client.UpdateCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a MediaWeaver channel.
//
// @param request - UpdateChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChannelResponse
func (client *Client) UpdateChannelWithOptions(request *UpdateChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccessPolicy) {
		query["AccessPolicy"] = request.AccessPolicy
	}

	if !dara.IsNil(request.AccessToken) {
		query["AccessToken"] = request.AccessToken
	}

	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.FillerSourceLocationName) {
		query["FillerSourceLocationName"] = request.FillerSourceLocationName
	}

	if !dara.IsNil(request.FillerSourceName) {
		query["FillerSourceName"] = request.FillerSourceName
	}

	if !dara.IsNil(request.OutPutConfigList) {
		query["OutPutConfigList"] = request.OutPutConfigList
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a MediaWeaver channel.
//
// @param request - UpdateChannelRequest
//
// @return UpdateChannelResponse
func (client *Client) UpdateChannel(request *UpdateChannelRequest) (_result *UpdateChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateChannelResponse{}
	_body, _err := client.UpdateChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom template.
//
// @param request - UpdateCustomTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomTemplateResponse
func (client *Client) UpdateCustomTemplateWithOptions(request *UpdateCustomTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateConfig) {
		query["TemplateConfig"] = request.TemplateConfig
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom template.
//
// @param request - UpdateCustomTemplateRequest
//
// @return UpdateCustomTemplateResponse
func (client *Client) UpdateCustomTemplate(request *UpdateCustomTemplateRequest) (_result *UpdateCustomTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomTemplateResponse{}
	_body, _err := client.UpdateCustomTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a personalized human voice. Only the media asset ID of the sample audio file can be modified.
//
// @param request - UpdateCustomizedVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomizedVoiceResponse
func (client *Client) UpdateCustomizedVoiceWithOptions(request *UpdateCustomizedVoiceRequest, runtime *dara.RuntimeOptions) (_result *UpdateCustomizedVoiceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DemoAudioMediaId) {
		query["DemoAudioMediaId"] = request.DemoAudioMediaId
	}

	if !dara.IsNil(request.VoiceId) {
		query["VoiceId"] = request.VoiceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCustomizedVoice"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomizedVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a personalized human voice. Only the media asset ID of the sample audio file can be modified.
//
// @param request - UpdateCustomizedVoiceRequest
//
// @return UpdateCustomizedVoiceResponse
func (client *Client) UpdateCustomizedVoice(request *UpdateCustomizedVoiceRequest) (_result *UpdateCustomizedVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateCustomizedVoiceResponse{}
	_body, _err := client.UpdateCustomizedVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an online editing project. You can call this operation to modify the configurations such as the title, timeline, and thumbnail of an online editing project.
//
// @param request - UpdateEditingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEditingProjectResponse
func (client *Client) UpdateEditingProjectWithOptions(request *UpdateEditingProjectRequest, runtime *dara.RuntimeOptions) (_result *UpdateEditingProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessStatus) {
		query["BusinessStatus"] = request.BusinessStatus
	}

	if !dara.IsNil(request.ClipsParam) {
		query["ClipsParam"] = request.ClipsParam
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Timeline) {
		body["Timeline"] = request.Timeline
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEditingProject"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an online editing project. You can call this operation to modify the configurations such as the title, timeline, and thumbnail of an online editing project.
//
// @param request - UpdateEditingProjectRequest
//
// @return UpdateEditingProjectResponse
func (client *Client) UpdateEditingProject(request *UpdateEditingProjectRequest) (_result *UpdateEditingProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateEditingProjectResponse{}
	_body, _err := client.UpdateEditingProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新热词库
//
// @param tmpReq - UpdateHotwordLibraryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateHotwordLibraryResponse
func (client *Client) UpdateHotwordLibraryWithOptions(tmpReq *UpdateHotwordLibraryRequest, runtime *dara.RuntimeOptions) (_result *UpdateHotwordLibraryResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateHotwordLibraryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Hotwords) {
		request.HotwordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Hotwords, dara.String("Hotwords"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.HotwordLibraryId) {
		query["HotwordLibraryId"] = request.HotwordLibraryId
	}

	if !dara.IsNil(request.HotwordsShrink) {
		query["Hotwords"] = request.HotwordsShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateHotwordLibrary"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateHotwordLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新热词库
//
// @param request - UpdateHotwordLibraryRequest
//
// @return UpdateHotwordLibraryResponse
func (client *Client) UpdateHotwordLibrary(request *UpdateHotwordLibraryRequest) (_result *UpdateHotwordLibraryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateHotwordLibraryResponse{}
	_body, _err := client.UpdateHotwordLibraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a live package channel, including the protocol, segment duration, and number of segments.
//
// Description:
//
// ## [](#)Usage notes
//
// You need to provide the name of the channel group to which the channel belongs, channel name, protocol, segment duration, and number of segments to update. In addition, you can choose to add or modify the description of the channel. Make sure that the provided channel group name and channel name conform to the naming conventions.
//
// @param request - UpdateLivePackageChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePackageChannelResponse
func (client *Client) UpdateLivePackageChannelWithOptions(request *UpdateLivePackageChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePackageChannelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		body["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.SegmentCount) {
		body["SegmentCount"] = request.SegmentCount
	}

	if !dara.IsNil(request.SegmentDuration) {
		body["SegmentDuration"] = request.SegmentDuration
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivePackageChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePackageChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a live package channel, including the protocol, segment duration, and number of segments.
//
// Description:
//
// ## [](#)Usage notes
//
// You need to provide the name of the channel group to which the channel belongs, channel name, protocol, segment duration, and number of segments to update. In addition, you can choose to add or modify the description of the channel. Make sure that the provided channel group name and channel name conform to the naming conventions.
//
// @param request - UpdateLivePackageChannelRequest
//
// @return UpdateLivePackageChannelResponse
func (client *Client) UpdateLivePackageChannel(request *UpdateLivePackageChannelRequest) (_result *UpdateLivePackageChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLivePackageChannelResponse{}
	_body, _err := client.UpdateLivePackageChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the credentials of ingest endpoints associated with a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// You can choose to update the primary endpoint, secondary endpoint, or both. The response includes the updated ingest endpoint URL, username, and password for the ingest device to reconfigure.
//
// @param request - UpdateLivePackageChannelCredentialsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePackageChannelCredentialsResponse
func (client *Client) UpdateLivePackageChannelCredentialsWithOptions(request *UpdateLivePackageChannelCredentialsRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePackageChannelCredentialsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelName) {
		body["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.RotateCredentials) {
		body["RotateCredentials"] = request.RotateCredentials
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivePackageChannelCredentials"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePackageChannelCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the credentials of ingest endpoints associated with a live package channel.
//
// Description:
//
// ## [](#)Usage notes
//
// You can choose to update the primary endpoint, secondary endpoint, or both. The response includes the updated ingest endpoint URL, username, and password for the ingest device to reconfigure.
//
// @param request - UpdateLivePackageChannelCredentialsRequest
//
// @return UpdateLivePackageChannelCredentialsResponse
func (client *Client) UpdateLivePackageChannelCredentials(request *UpdateLivePackageChannelCredentialsRequest) (_result *UpdateLivePackageChannelCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLivePackageChannelCredentialsResponse{}
	_body, _err := client.UpdateLivePackageChannelCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a live package channel group including its description.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to modify the name and description of a live package channel group. The channel group name must conform to the naming conventions and can be up to 1,000 characters. The API response includes the updated channel group details and unique identifier of the request.
//
// @param request - UpdateLivePackageChannelGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePackageChannelGroupResponse
func (client *Client) UpdateLivePackageChannelGroupWithOptions(request *UpdateLivePackageChannelGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePackageChannelGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivePackageChannelGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePackageChannelGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a live package channel group including its description.
//
// Description:
//
// ## [](#)Usage notes
//
// This API operation allows you to modify the name and description of a live package channel group. The channel group name must conform to the naming conventions and can be up to 1,000 characters. The API response includes the updated channel group details and unique identifier of the request.
//
// @param request - UpdateLivePackageChannelGroupRequest
//
// @return UpdateLivePackageChannelGroupResponse
func (client *Client) UpdateLivePackageChannelGroup(request *UpdateLivePackageChannelGroupRequest) (_result *UpdateLivePackageChannelGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLivePackageChannelGroupResponse{}
	_body, _err := client.UpdateLivePackageChannelGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the origin endpoint settings including the protocol, time shifting, and access control settings.
//
// Description:
//
// ## [](#)Usage notes
//
// You can call this operation to modify the origin protocol, set the number of days that time-shifted content is available, define playlist names, and configure the IP address blacklist and whitelist, allowing for fine-grained control over streaming media distribution. Some parameters are required. You must configure IpWhitelist, AuthorizationCode, or both.
//
// @param tmpReq - UpdateLivePackageOriginEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLivePackageOriginEndpointResponse
func (client *Client) UpdateLivePackageOriginEndpointWithOptions(tmpReq *UpdateLivePackageOriginEndpointRequest, runtime *dara.RuntimeOptions) (_result *UpdateLivePackageOriginEndpointResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLivePackageOriginEndpointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LivePackagingConfig) {
		request.LivePackagingConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LivePackagingConfig, dara.String("LivePackagingConfig"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationCode) {
		body["AuthorizationCode"] = request.AuthorizationCode
	}

	if !dara.IsNil(request.ChannelName) {
		body["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndpointName) {
		body["EndpointName"] = request.EndpointName
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.IpBlacklist) {
		body["IpBlacklist"] = request.IpBlacklist
	}

	if !dara.IsNil(request.IpWhitelist) {
		body["IpWhitelist"] = request.IpWhitelist
	}

	if !dara.IsNil(request.LivePackagingConfigShrink) {
		body["LivePackagingConfig"] = request.LivePackagingConfigShrink
	}

	if !dara.IsNil(request.ManifestName) {
		body["ManifestName"] = request.ManifestName
	}

	if !dara.IsNil(request.Protocol) {
		body["Protocol"] = request.Protocol
	}

	if !dara.IsNil(request.TimeshiftVision) {
		body["TimeshiftVision"] = request.TimeshiftVision
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLivePackageOriginEndpoint"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLivePackageOriginEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the origin endpoint settings including the protocol, time shifting, and access control settings.
//
// Description:
//
// ## [](#)Usage notes
//
// You can call this operation to modify the origin protocol, set the number of days that time-shifted content is available, define playlist names, and configure the IP address blacklist and whitelist, allowing for fine-grained control over streaming media distribution. Some parameters are required. You must configure IpWhitelist, AuthorizationCode, or both.
//
// @param request - UpdateLivePackageOriginEndpointRequest
//
// @return UpdateLivePackageOriginEndpointResponse
func (client *Client) UpdateLivePackageOriginEndpoint(request *UpdateLivePackageOriginEndpointRequest) (_result *UpdateLivePackageOriginEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLivePackageOriginEndpointResponse{}
	_body, _err := client.UpdateLivePackageOriginEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream recording template.
//
// Description:
//
// Only user-created templates can be updated. The preset template cannot be updated.
//
// @param tmpReq - UpdateLiveRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveRecordTemplateResponse
func (client *Client) UpdateLiveRecordTemplateWithOptions(tmpReq *UpdateLiveRecordTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveRecordTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLiveRecordTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RecordFormat) {
		request.RecordFormatShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RecordFormat, dara.String("RecordFormat"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.RecordFormatShrink) {
		body["RecordFormat"] = request.RecordFormatShrink
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveRecordTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream recording template.
//
// Description:
//
// Only user-created templates can be updated. The preset template cannot be updated.
//
// @param request - UpdateLiveRecordTemplateRequest
//
// @return UpdateLiveRecordTemplateResponse
func (client *Client) UpdateLiveRecordTemplate(request *UpdateLiveRecordTemplateRequest) (_result *UpdateLiveRecordTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLiveRecordTemplateResponse{}
	_body, _err := client.UpdateLiveRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream snapshot template.
//
// @param request - UpdateLiveSnapshotTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveSnapshotTemplateResponse
func (client *Client) UpdateLiveSnapshotTemplateWithOptions(request *UpdateLiveSnapshotTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveSnapshotTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OverwriteFormat) {
		body["OverwriteFormat"] = request.OverwriteFormat
	}

	if !dara.IsNil(request.SequenceFormat) {
		body["SequenceFormat"] = request.SequenceFormat
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		body["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TimeInterval) {
		body["TimeInterval"] = request.TimeInterval
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveSnapshotTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveSnapshotTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream snapshot template.
//
// @param request - UpdateLiveSnapshotTemplateRequest
//
// @return UpdateLiveSnapshotTemplateResponse
func (client *Client) UpdateLiveSnapshotTemplate(request *UpdateLiveSnapshotTemplateRequest) (_result *UpdateLiveSnapshotTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLiveSnapshotTemplateResponse{}
	_body, _err := client.UpdateLiveSnapshotTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream transcoding job.
//
// Description:
//
//	  For a non-timed transcoding job, you can modify the Name parameter of the job, regardless of the job state.
//
//		- For a timed job, you can modify the Name, StreamInput, TranscodeOutput, and TimedConfig parameters. However, the StreamInput, TranscodeOutput, and TimedConfig parameters can be modified only when the job is not started.
//
// @param tmpReq - UpdateLiveTranscodeJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveTranscodeJobResponse
func (client *Client) UpdateLiveTranscodeJobWithOptions(tmpReq *UpdateLiveTranscodeJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveTranscodeJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLiveTranscodeJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StreamInput) {
		request.StreamInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamInput, dara.String("StreamInput"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TimedConfig) {
		request.TimedConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TimedConfig, dara.String("TimedConfig"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TranscodeOutput) {
		request.TranscodeOutputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TranscodeOutput, dara.String("TranscodeOutput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.StreamInputShrink) {
		query["StreamInput"] = request.StreamInputShrink
	}

	if !dara.IsNil(request.TimedConfigShrink) {
		query["TimedConfig"] = request.TimedConfigShrink
	}

	if !dara.IsNil(request.TranscodeOutputShrink) {
		query["TranscodeOutput"] = request.TranscodeOutputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveTranscodeJob"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveTranscodeJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream transcoding job.
//
// Description:
//
//	  For a non-timed transcoding job, you can modify the Name parameter of the job, regardless of the job state.
//
//		- For a timed job, you can modify the Name, StreamInput, TranscodeOutput, and TimedConfig parameters. However, the StreamInput, TranscodeOutput, and TimedConfig parameters can be modified only when the job is not started.
//
// @param request - UpdateLiveTranscodeJobRequest
//
// @return UpdateLiveTranscodeJobResponse
func (client *Client) UpdateLiveTranscodeJob(request *UpdateLiveTranscodeJobRequest) (_result *UpdateLiveTranscodeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLiveTranscodeJobResponse{}
	_body, _err := client.UpdateLiveTranscodeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream transcoding template.
//
// @param tmpReq - UpdateLiveTranscodeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLiveTranscodeTemplateResponse
func (client *Client) UpdateLiveTranscodeTemplateWithOptions(tmpReq *UpdateLiveTranscodeTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateLiveTranscodeTemplateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateLiveTranscodeTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TemplateConfig) {
		request.TemplateConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplateConfig, dara.String("TemplateConfig"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateConfigShrink) {
		query["TemplateConfig"] = request.TemplateConfigShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLiveTranscodeTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLiveTranscodeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a live stream transcoding template.
//
// @param request - UpdateLiveTranscodeTemplateRequest
//
// @return UpdateLiveTranscodeTemplateResponse
func (client *Client) UpdateLiveTranscodeTemplate(request *UpdateLiveTranscodeTemplateRequest) (_result *UpdateLiveTranscodeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateLiveTranscodeTemplateResponse{}
	_body, _err := client.UpdateLiveTranscodeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the source of a MediaConnect flow.
//
// Description:
//
//	  You can modify the source only when the flow is in the offline state.
//
//		- The source type cannot be modified.
//
// @param request - UpdateMediaConnectFlowInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaConnectFlowInputResponse
func (client *Client) UpdateMediaConnectFlowInputWithOptions(request *UpdateMediaConnectFlowInputRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaConnectFlowInputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidrs) {
		query["Cidrs"] = request.Cidrs
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.InputFromUrl) {
		query["InputFromUrl"] = request.InputFromUrl
	}

	if !dara.IsNil(request.InputName) {
		query["InputName"] = request.InputName
	}

	if !dara.IsNil(request.MaxBitrate) {
		query["MaxBitrate"] = request.MaxBitrate
	}

	if !dara.IsNil(request.SrtLatency) {
		query["SrtLatency"] = request.SrtLatency
	}

	if !dara.IsNil(request.SrtPassphrase) {
		query["SrtPassphrase"] = request.SrtPassphrase
	}

	if !dara.IsNil(request.SrtPbkeyLen) {
		query["SrtPbkeyLen"] = request.SrtPbkeyLen
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaConnectFlowInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaConnectFlowInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the source of a MediaConnect flow.
//
// Description:
//
//	  You can modify the source only when the flow is in the offline state.
//
//		- The source type cannot be modified.
//
// @param request - UpdateMediaConnectFlowInputRequest
//
// @return UpdateMediaConnectFlowInputResponse
func (client *Client) UpdateMediaConnectFlowInput(request *UpdateMediaConnectFlowInputRequest) (_result *UpdateMediaConnectFlowInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaConnectFlowInputResponse{}
	_body, _err := client.UpdateMediaConnectFlowInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an output of a MediaConnect flow.
//
// Description:
//
//	  You can modify an output only when the flow is in the offline state.
//
//		- The output type cannot be modified.
//
// @param request - UpdateMediaConnectFlowOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaConnectFlowOutputResponse
func (client *Client) UpdateMediaConnectFlowOutputWithOptions(request *UpdateMediaConnectFlowOutputRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaConnectFlowOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cidrs) {
		query["Cidrs"] = request.Cidrs
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.OutputName) {
		query["OutputName"] = request.OutputName
	}

	if !dara.IsNil(request.OutputToUrl) {
		query["OutputToUrl"] = request.OutputToUrl
	}

	if !dara.IsNil(request.PlayerLimit) {
		query["PlayerLimit"] = request.PlayerLimit
	}

	if !dara.IsNil(request.SrtLatency) {
		query["SrtLatency"] = request.SrtLatency
	}

	if !dara.IsNil(request.SrtPassphrase) {
		query["SrtPassphrase"] = request.SrtPassphrase
	}

	if !dara.IsNil(request.SrtPbkeyLen) {
		query["SrtPbkeyLen"] = request.SrtPbkeyLen
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaConnectFlowOutput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaConnectFlowOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an output of a MediaConnect flow.
//
// Description:
//
//	  You can modify an output only when the flow is in the offline state.
//
//		- The output type cannot be modified.
//
// @param request - UpdateMediaConnectFlowOutputRequest
//
// @return UpdateMediaConnectFlowOutputResponse
func (client *Client) UpdateMediaConnectFlowOutput(request *UpdateMediaConnectFlowOutputRequest) (_result *UpdateMediaConnectFlowOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaConnectFlowOutputResponse{}
	_body, _err := client.UpdateMediaConnectFlowOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the state of a MediaConnect flow.
//
// @param request - UpdateMediaConnectFlowStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaConnectFlowStatusResponse
func (client *Client) UpdateMediaConnectFlowStatusWithOptions(request *UpdateMediaConnectFlowStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaConnectFlowStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaConnectFlowStatus"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaConnectFlowStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the state of a MediaConnect flow.
//
// @param request - UpdateMediaConnectFlowStatusRequest
//
// @return UpdateMediaConnectFlowStatusResponse
func (client *Client) UpdateMediaConnectFlowStatus(request *UpdateMediaConnectFlowStatusRequest) (_result *UpdateMediaConnectFlowStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaConnectFlowStatusResponse{}
	_body, _err := client.UpdateMediaConnectFlowStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates information about a media asset based on the ID of the media asset in Intelligent Media Services (IMS) or the input URL of the media asset.
//
// Description:
//
// If the MediaId parameter is specified, the MediaId parameter is preferentially used for the query. If the MediaId parameter is left empty, the InputURL parameter must be specified. The request ID and media asset ID are returned. You cannot modify the input URL of a media asset by specifying the ID of the media asset.
//
// @param request - UpdateMediaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaInfoResponse
func (client *Client) UpdateMediaInfoWithOptions(request *UpdateMediaInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppendTags) {
		query["AppendTags"] = request.AppendTags
	}

	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.CateId) {
		query["CateId"] = request.CateId
	}

	if !dara.IsNil(request.Category) {
		query["Category"] = request.Category
	}

	if !dara.IsNil(request.CoverURL) {
		query["CoverURL"] = request.CoverURL
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.InputURL) {
		query["InputURL"] = request.InputURL
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaTags) {
		query["MediaTags"] = request.MediaTags
	}

	if !dara.IsNil(request.ReferenceId) {
		query["ReferenceId"] = request.ReferenceId
	}

	if !dara.IsNil(request.Title) {
		query["Title"] = request.Title
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaInfo"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates information about a media asset based on the ID of the media asset in Intelligent Media Services (IMS) or the input URL of the media asset.
//
// Description:
//
// If the MediaId parameter is specified, the MediaId parameter is preferentially used for the query. If the MediaId parameter is left empty, the InputURL parameter must be specified. The request ID and media asset ID are returned. You cannot modify the input URL of a media asset by specifying the ID of the media asset.
//
// @param request - UpdateMediaInfoRequest
//
// @return UpdateMediaInfoResponse
func (client *Client) UpdateMediaInfo(request *UpdateMediaInfoRequest) (_result *UpdateMediaInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaInfoResponse{}
	_body, _err := client.UpdateMediaInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a MediaLive channel.
//
// Description:
//
//	You can modify a MediaLive channel only when it is not running.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - UpdateMediaLiveChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaLiveChannelResponse
func (client *Client) UpdateMediaLiveChannelWithOptions(tmpReq *UpdateMediaLiveChannelRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaLiveChannelResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMediaLiveChannelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AudioSettings) {
		request.AudioSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AudioSettings, dara.String("AudioSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.InputAttachments) {
		request.InputAttachmentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputAttachments, dara.String("InputAttachments"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OutputGroups) {
		request.OutputGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OutputGroups, dara.String("OutputGroups"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VideoSettings) {
		request.VideoSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VideoSettings, dara.String("VideoSettings"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AudioSettingsShrink) {
		body["AudioSettings"] = request.AudioSettingsShrink
	}

	if !dara.IsNil(request.ChannelId) {
		body["ChannelId"] = request.ChannelId
	}

	if !dara.IsNil(request.InputAttachmentsShrink) {
		body["InputAttachments"] = request.InputAttachmentsShrink
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputGroupsShrink) {
		body["OutputGroups"] = request.OutputGroupsShrink
	}

	if !dara.IsNil(request.VideoSettingsShrink) {
		body["VideoSettings"] = request.VideoSettingsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaLiveChannel"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaLiveChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a MediaLive channel.
//
// Description:
//
//	You can modify a MediaLive channel only when it is not running.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateMediaLiveChannelRequest
//
// @return UpdateMediaLiveChannelResponse
func (client *Client) UpdateMediaLiveChannel(request *UpdateMediaLiveChannelRequest) (_result *UpdateMediaLiveChannelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaLiveChannelResponse{}
	_body, _err := client.UpdateMediaLiveChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an input of MediaLive.
//
// Description:
//
//	You can modify an input only when it is not associated with a MediaLive channel.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - UpdateMediaLiveInputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaLiveInputResponse
func (client *Client) UpdateMediaLiveInputWithOptions(tmpReq *UpdateMediaLiveInputRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaLiveInputResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMediaLiveInputShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.InputSettings) {
		request.InputSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputSettings, dara.String("InputSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SecurityGroupIds) {
		request.SecurityGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecurityGroupIds, dara.String("SecurityGroupIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.InputSettingsShrink) {
		query["InputSettings"] = request.InputSettingsShrink
	}

	if !dara.IsNil(request.SecurityGroupIdsShrink) {
		query["SecurityGroupIds"] = request.SecurityGroupIdsShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InputId) {
		body["InputId"] = request.InputId
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaLiveInput"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaLiveInputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an input of MediaLive.
//
// Description:
//
//	You can modify an input only when it is not associated with a MediaLive channel.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateMediaLiveInputRequest
//
// @return UpdateMediaLiveInputResponse
func (client *Client) UpdateMediaLiveInput(request *UpdateMediaLiveInputRequest) (_result *UpdateMediaLiveInputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaLiveInputResponse{}
	_body, _err := client.UpdateMediaLiveInputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a security group created in MediaLive.
//
// Description:
//
//	You can modify a security group only when it is not associated with a MediaLive input.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param tmpReq - UpdateMediaLiveInputSecurityGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaLiveInputSecurityGroupResponse
func (client *Client) UpdateMediaLiveInputSecurityGroupWithOptions(tmpReq *UpdateMediaLiveInputSecurityGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaLiveInputSecurityGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMediaLiveInputSecurityGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WhitelistRules) {
		request.WhitelistRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WhitelistRules, dara.String("WhitelistRules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.SecurityGroupId) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.WhitelistRulesShrink) {
		body["WhitelistRules"] = request.WhitelistRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaLiveInputSecurityGroup"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a security group created in MediaLive.
//
// Description:
//
//	You can modify a security group only when it is not associated with a MediaLive input.
//
// ## QPS limit
//
// This operation can be called up to 50 times per second for each Alibaba Cloud account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation.
//
// @param request - UpdateMediaLiveInputSecurityGroupRequest
//
// @return UpdateMediaLiveInputSecurityGroupResponse
func (client *Client) UpdateMediaLiveInputSecurityGroup(request *UpdateMediaLiveInputSecurityGroupRequest) (_result *UpdateMediaLiveInputSecurityGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaLiveInputSecurityGroupResponse{}
	_body, _err := client.UpdateMediaLiveInputSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the marks of a media asset.
//
// @param request - UpdateMediaMarksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaMarksResponse
func (client *Client) UpdateMediaMarksWithOptions(request *UpdateMediaMarksRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaMarksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MediaMarks) {
		query["MediaMarks"] = request.MediaMarks
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaMarks"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaMarksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the marks of a media asset.
//
// @param request - UpdateMediaMarksRequest
//
// @return UpdateMediaMarksResponse
func (client *Client) UpdateMediaMarks(request *UpdateMediaMarksRequest) (_result *UpdateMediaMarksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaMarksResponse{}
	_body, _err := client.UpdateMediaMarksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the media asset information in a search library.
//
// @param request - UpdateMediaToSearchLibRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMediaToSearchLibResponse
func (client *Client) UpdateMediaToSearchLibWithOptions(request *UpdateMediaToSearchLibRequest, runtime *dara.RuntimeOptions) (_result *UpdateMediaToSearchLibResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.MsgBody) {
		query["MsgBody"] = request.MsgBody
	}

	if !dara.IsNil(request.Namespace) {
		query["Namespace"] = request.Namespace
	}

	if !dara.IsNil(request.SearchLibName) {
		query["SearchLibName"] = request.SearchLibName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMediaToSearchLib"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMediaToSearchLibResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the media asset information in a search library.
//
// @param request - UpdateMediaToSearchLibRequest
//
// @return UpdateMediaToSearchLibResponse
func (client *Client) UpdateMediaToSearchLib(request *UpdateMediaToSearchLibRequest) (_result *UpdateMediaToSearchLibResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMediaToSearchLibResponse{}
	_body, _err := client.UpdateMediaToSearchLibWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - UpdatePipelineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipelineWithOptions(request *UpdatePipelineRequest, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PipelineId) {
		query["PipelineId"] = request.PipelineId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipeline"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an ApsaraVideo Media Processing (MPS) queue.
//
// @param request - UpdatePipelineRequest
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipeline(request *UpdatePipelineRequest) (_result *UpdatePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.UpdatePipelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a program in a MediaWeaver channel.
//
// @param request - UpdateProgramRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateProgramResponse
func (client *Client) UpdateProgramWithOptions(request *UpdateProgramRequest, runtime *dara.RuntimeOptions) (_result *UpdateProgramResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdBreaks) {
		query["AdBreaks"] = request.AdBreaks
	}

	if !dara.IsNil(request.ChannelName) {
		query["ChannelName"] = request.ChannelName
	}

	if !dara.IsNil(request.ClipRange) {
		query["ClipRange"] = request.ClipRange
	}

	if !dara.IsNil(request.ProgramName) {
		query["ProgramName"] = request.ProgramName
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	if !dara.IsNil(request.Transition) {
		query["Transition"] = request.Transition
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateProgram"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateProgramResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a program in a MediaWeaver channel.
//
// @param request - UpdateProgramRequest
//
// @return UpdateProgramResponse
func (client *Client) UpdateProgram(request *UpdateProgramRequest) (_result *UpdateProgramResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateProgramResponse{}
	_body, _err := client.UpdateProgramWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改实例的配置
//
// @param tmpReq - UpdateRtcRobotInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRtcRobotInstanceResponse
func (client *Client) UpdateRtcRobotInstanceWithOptions(tmpReq *UpdateRtcRobotInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateRtcRobotInstanceResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRtcRobotInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Config) {
		request.ConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Config, dara.String("Config"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ConfigShrink) {
		query["Config"] = request.ConfigShrink
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRtcRobotInstance"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRtcRobotInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改实例的配置
//
// @param request - UpdateRtcRobotInstanceRequest
//
// @return UpdateRtcRobotInstanceResponse
func (client *Client) UpdateRtcRobotInstance(request *UpdateRtcRobotInstanceRequest) (_result *UpdateRtcRobotInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRtcRobotInstanceResponse{}
	_body, _err := client.UpdateRtcRobotInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a source in MediaWeaver.
//
// @param request - UpdateSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSourceResponse
func (client *Client) UpdateSourceWithOptions(request *UpdateSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HttpPackageConfigurations) {
		query["HttpPackageConfigurations"] = request.HttpPackageConfigurations
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	if !dara.IsNil(request.SourceName) {
		query["SourceName"] = request.SourceName
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSource"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a source in MediaWeaver.
//
// @param request - UpdateSourceRequest
//
// @return UpdateSourceResponse
func (client *Client) UpdateSource(request *UpdateSourceRequest) (_result *UpdateSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSourceResponse{}
	_body, _err := client.UpdateSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a source location.
//
// @param request - UpdateSourceLocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSourceLocationResponse
func (client *Client) UpdateSourceLocationWithOptions(request *UpdateSourceLocationRequest, runtime *dara.RuntimeOptions) (_result *UpdateSourceLocationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaseUrl) {
		query["BaseUrl"] = request.BaseUrl
	}

	if !dara.IsNil(request.EnableSegmentDelivery) {
		query["EnableSegmentDelivery"] = request.EnableSegmentDelivery
	}

	if !dara.IsNil(request.SegmentDeliveryUrl) {
		query["SegmentDeliveryUrl"] = request.SegmentDeliveryUrl
	}

	if !dara.IsNil(request.SourceLocationName) {
		query["SourceLocationName"] = request.SourceLocationName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSourceLocation"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSourceLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a source location.
//
// @param request - UpdateSourceLocationRequest
//
// @return UpdateSourceLocationResponse
func (client *Client) UpdateSourceLocation(request *UpdateSourceLocationRequest) (_result *UpdateSourceLocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSourceLocationResponse{}
	_body, _err := client.UpdateSourceLocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an online editing template. You can modify the template title and template configurations.
//
// Description:
//
//	  For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//		- For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - UpdateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(request *UpdateTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CoverUrl) {
		query["CoverUrl"] = request.CoverUrl
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PreviewMedia) {
		query["PreviewMedia"] = request.PreviewMedia
	}

	if !dara.IsNil(request.RelatedMediaids) {
		query["RelatedMediaids"] = request.RelatedMediaids
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["Config"] = request.Config
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an online editing template. You can modify the template title and template configurations.
//
// Description:
//
//	  For more information about how to use a regular template, see [Create and use a regular template](https://help.aliyun.com/document_detail/445399.html).
//
//		- For more information about how to use an advanced template, see [Create and use advanced templates](https://help.aliyun.com/document_detail/445389.html).
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads an audio or video file based on the URL of the source file. You can upload multiple media files at a time.
//
// Description:
//
//	  If a callback is configured, you will receive an UploadByURLComplete event notification after the file is uploaded. You can query the upload status by calling the GetURLUploadInfos operation.
//
//		- After a request is submitted, the upload job is queued as an asynchronous job in the cloud. You can query the status of the upload job based on information such as the URL and media asset ID that are returned in the event notification.
//
//		- You can call this operation to upload media files that are not stored on a local server or device and must be uploaded by using URLs that are accessible over the Internet.
//
//		- You can call this operation to upload media files only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media file to an OSS bucket, pull the file to a local directory, use [OSS SDK](https://help.aliyun.com/document_detail/32006.html) to upload the file to an OSS bucket, and then call the [RegisterMediaInfo](https://help.aliyun.com/document_detail/441152.html) operation to register the file in the OSS bucket with the media asset library.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
//		- You can call this operation to upload only audio and video files.
//
// @param request - UploadMediaByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadMediaByURLResponse
func (client *Client) UploadMediaByURLWithOptions(request *UploadMediaByURLRequest, runtime *dara.RuntimeOptions) (_result *UploadMediaByURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.EntityId) {
		query["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.MediaMetaData) {
		query["MediaMetaData"] = request.MediaMetaData
	}

	if !dara.IsNil(request.PostProcessConfig) {
		query["PostProcessConfig"] = request.PostProcessConfig
	}

	if !dara.IsNil(request.UploadTargetConfig) {
		query["UploadTargetConfig"] = request.UploadTargetConfig
	}

	if !dara.IsNil(request.UploadURLs) {
		query["UploadURLs"] = request.UploadURLs
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadMediaByURL"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads an audio or video file based on the URL of the source file. You can upload multiple media files at a time.
//
// Description:
//
//	  If a callback is configured, you will receive an UploadByURLComplete event notification after the file is uploaded. You can query the upload status by calling the GetURLUploadInfos operation.
//
//		- After a request is submitted, the upload job is queued as an asynchronous job in the cloud. You can query the status of the upload job based on information such as the URL and media asset ID that are returned in the event notification.
//
//		- You can call this operation to upload media files that are not stored on a local server or device and must be uploaded by using URLs that are accessible over the Internet.
//
//		- You can call this operation to upload media files only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media file to an OSS bucket, pull the file to a local directory, use [OSS SDK](https://help.aliyun.com/document_detail/32006.html) to upload the file to an OSS bucket, and then call the [RegisterMediaInfo](https://help.aliyun.com/document_detail/441152.html) operation to register the file in the OSS bucket with the media asset library.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
//		- You can call this operation to upload only audio and video files.
//
// @param request - UploadMediaByURLRequest
//
// @return UploadMediaByURLResponse
func (client *Client) UploadMediaByURL(request *UploadMediaByURLRequest) (_result *UploadMediaByURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadMediaByURLResponse{}
	_body, _err := client.UploadMediaByURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads a media stream file based on the URL of the source file.
//
// Description:
//
//	  You can call this operation to pull a media stream file based on a URL and upload the file. After the media stream file is uploaded, the media stream is associated with the specified media asset ID.
//
//		- You can call this operation to upload media stream files only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media stream file to an OSS bucket, pull the file to a local directory, use [OSS SDK](https://help.aliyun.com/document_detail/32006.html) to upload the file to an OSS bucket, and then call the [RegisterMediaStream](https://help.aliyun.com/document_detail/440765.html) operation to associate the media stream with the specified media asset ID.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
// @param request - UploadStreamByURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadStreamByURLResponse
func (client *Client) UploadStreamByURLWithOptions(request *UploadStreamByURLRequest, runtime *dara.RuntimeOptions) (_result *UploadStreamByURLResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Definition) {
		query["Definition"] = request.Definition
	}

	if !dara.IsNil(request.FileExtension) {
		query["FileExtension"] = request.FileExtension
	}

	if !dara.IsNil(request.HDRType) {
		query["HDRType"] = request.HDRType
	}

	if !dara.IsNil(request.MediaId) {
		query["MediaId"] = request.MediaId
	}

	if !dara.IsNil(request.StreamURL) {
		query["StreamURL"] = request.StreamURL
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadStreamByURL"),
		Version:     dara.String("2020-11-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadStreamByURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads a media stream file based on the URL of the source file.
//
// Description:
//
//	  You can call this operation to pull a media stream file based on a URL and upload the file. After the media stream file is uploaded, the media stream is associated with the specified media asset ID.
//
//		- You can call this operation to upload media stream files only to ApsaraVideo VOD, but not to your own Object Storage Service (OSS) buckets. To upload a media stream file to an OSS bucket, pull the file to a local directory, use [OSS SDK](https://help.aliyun.com/document_detail/32006.html) to upload the file to an OSS bucket, and then call the [RegisterMediaStream](https://help.aliyun.com/document_detail/440765.html) operation to associate the media stream with the specified media asset ID.
//
//		- This operation is available only in the China (Shanghai), China (Beijing), and China (Shenzhen) regions.
//
// @param request - UploadStreamByURLRequest
//
// @return UploadStreamByURLResponse
func (client *Client) UploadStreamByURL(request *UploadStreamByURLRequest) (_result *UploadStreamByURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadStreamByURLResponse{}
	_body, _err := client.UploadStreamByURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
