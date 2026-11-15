// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("MilvusKnowledgeBase")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("milvusknowledgebase"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Registers files that have been uploaded to the knowledge base storage as knowledge base documents and **automatically triggers parsing*	- (chunking and embedding). Supports two types of import:
//
// - `LOCAL_UPLOAD`: Used with the `GetKnowledgeBasePreSignedUrl` direct upload process. This operation only registers the files and does not verify whether the files have been actually uploaded. Therefore, you must complete the PUT upload first.
//
// - `OSS_IMPORT`: Imports files from an external OSS bucket. This operation creates an asynchronous import task and returns the `knowledge_import_task_id`. The system downloads and registers the files in the background.
//
// A maximum of 100 files can be processed at a time.
//
// @param request - AddDocumentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDocumentsResponse
func (client *Client) AddDocumentsWithOptions(datasetId *string, request *AddDocumentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddDocumentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dedup) {
		body["Dedup"] = request.Dedup
	}

	if !dara.IsNil(request.Documents) {
		body["Documents"] = request.Documents
	}

	if !dara.IsNil(request.ImportType) {
		body["ImportType"] = request.ImportType
	}

	if !dara.IsNil(request.KnowledgeBaseId) {
		body["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !dara.IsNil(request.MetaFields) {
		body["MetaFields"] = request.MetaFields
	}

	if !dara.IsNil(request.StrategyId) {
		body["StrategyId"] = request.StrategyId
	}

	if !dara.IsNil(request.DingTalkConfiguration) {
		body["dingTalkConfiguration"] = request.DingTalkConfiguration
	}

	if !dara.IsNil(request.ParentId) {
		body["parentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDocuments"),
		Version:     dara.String("2026-06-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(datasetId)) + "/documents/addDocuments"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers files that have been uploaded to the knowledge base storage as knowledge base documents and **automatically triggers parsing*	- (chunking and embedding). Supports two types of import:
//
// - `LOCAL_UPLOAD`: Used with the `GetKnowledgeBasePreSignedUrl` direct upload process. This operation only registers the files and does not verify whether the files have been actually uploaded. Therefore, you must complete the PUT upload first.
//
// - `OSS_IMPORT`: Imports files from an external OSS bucket. This operation creates an asynchronous import task and returns the `knowledge_import_task_id`. The system downloads and registers the files in the background.
//
// A maximum of 100 files can be processed at a time.
//
// @param request - AddDocumentsRequest
//
// @return AddDocumentsResponse
func (client *Client) AddDocuments(datasetId *string, request *AddDocumentsRequest) (_result *AddDocumentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddDocumentsResponse{}
	_body, _err := client.AddDocumentsWithOptions(datasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an **Object Storage Service (OSS) pre-signed PUT URL*	- pointing to the dedicated storage of the knowledge base for each file in `Documents`. Callers use the URL to upload file content directly to OSS and then call `AddDocuments` to complete the registration. A maximum of 100 files can be processed at a time.
//
// @param request - GetKnowledgeBasePreSignedUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBasePreSignedUrlResponse
func (client *Client) GetKnowledgeBasePreSignedUrlWithOptions(datasetId *string, request *GetKnowledgeBasePreSignedUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBasePreSignedUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Documents) {
		body["Documents"] = request.Documents
	}

	if !dara.IsNil(request.ExpiresIn) {
		body["ExpiresIn"] = request.ExpiresIn
	}

	if !dara.IsNil(request.KnowledgeBaseId) {
		body["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeBasePreSignedUrl"),
		Version:     dara.String("2026-06-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/datasets/" + dara.PercentEncode(dara.StringValue(datasetId)) + "/getKnowledgeBasePreSignedUrl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKnowledgeBasePreSignedUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an **Object Storage Service (OSS) pre-signed PUT URL*	- pointing to the dedicated storage of the knowledge base for each file in `Documents`. Callers use the URL to upload file content directly to OSS and then call `AddDocuments` to complete the registration. A maximum of 100 files can be processed at a time.
//
// @param request - GetKnowledgeBasePreSignedUrlRequest
//
// @return GetKnowledgeBasePreSignedUrlResponse
func (client *Client) GetKnowledgeBasePreSignedUrl(datasetId *string, request *GetKnowledgeBasePreSignedUrlRequest) (_result *GetKnowledgeBasePreSignedUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeBasePreSignedUrlResponse{}
	_body, _err := client.GetKnowledgeBasePreSignedUrlWithOptions(datasetId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches a knowledge base.
//
// Description:
//
// Searches a specified knowledge base by using a text query or an image. For the version parameter, use the display name DRAFT, LATEST_PUBLISHED, or vN. Pass tag filter conditions based on the actual operators used by the backend.
//
// @param request - SearchKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchKnowledgeBaseResponse
func (client *Client) SearchKnowledgeBaseWithOptions(knowledgeBaseId *string, request *SearchKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIds) {
		body["documentIds"] = request.DocumentIds
	}

	if !dara.IsNil(request.EnableKnowledgeGraph) {
		body["enableKnowledgeGraph"] = request.EnableKnowledgeGraph
	}

	if !dara.IsNil(request.EvidencePageToken) {
		body["evidencePageToken"] = request.EvidencePageToken
	}

	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	if !dara.IsNil(request.IncludeWiki) {
		body["includeWiki"] = request.IncludeWiki
	}

	if !dara.IsNil(request.PageNumber) {
		body["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.RerankModelId) {
		body["rerankModelId"] = request.RerankModelId
	}

	if !dara.IsNil(request.RerankModelName) {
		body["rerankModelName"] = request.RerankModelName
	}

	if !dara.IsNil(request.RetrievalConfig) {
		body["retrievalConfig"] = request.RetrievalConfig
	}

	if !dara.IsNil(request.TagFilter) {
		body["tagFilter"] = request.TagFilter
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchKnowledgeBase"),
		Version:     dara.String("2026-06-04"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/knowledge-bases/" + dara.PercentEncode(dara.StringValue(knowledgeBaseId)) + "/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchKnowledgeBaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches a knowledge base.
//
// Description:
//
// Searches a specified knowledge base by using a text query or an image. For the version parameter, use the display name DRAFT, LATEST_PUBLISHED, or vN. Pass tag filter conditions based on the actual operators used by the backend.
//
// @param request - SearchKnowledgeBaseRequest
//
// @return SearchKnowledgeBaseResponse
func (client *Client) SearchKnowledgeBase(knowledgeBaseId *string, request *SearchKnowledgeBaseRequest) (_result *SearchKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchKnowledgeBaseResponse{}
	_body, _err := client.SearchKnowledgeBaseWithOptions(knowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
