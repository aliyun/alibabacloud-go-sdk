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
	client.Endpoint, _err = client.GetEndpoint(dara.String("aisearchengine"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Provides a search API operation for the CleverSee AI Search platform. After you create a search application on the platform, you can invoke this search API operation to retrieve images, documents, and audio/video content in datasets. The platform supports text (natural language), image, or combined text-and-image input, enabling fast adaptation to multi-modal large-scale data search scenarios and helping you efficiently locate target content.
//
// Description:
//
// ## Operation description
//
// This operation supports calling two types of search applications on the [CleverSee AI Search platform](https://aisearch.aliyun.com/web-search): **image search applications*	- (text-to-image, image-to-image, and combined text-and-image search) and **audio/video search applications*	- (text-to-audio/video, image-to-video, and combined text-and-image-to-video search).
//
// ### Data sources
//
// Audio and video data is supported. You can upload and update data through the [CleverSee AI Search platform](https://aisearch.aliyun.com/web-search) UI or through the [Dataset Data Add/Update API](https://www.alibabacloud.com/help/en/document_detail/3038471.html).
//
// # Authentication
//
// Call the CleverSee - Intelligent Search service by using the Alibaba Cloud SDK. For more information, see [AI search engine operation](https://api.aliyun.com/document/AiSearchEngine/2026-04-17/EngineSearch).
//
// @param request - EngineSearchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EngineSearchResponse
func (client *Client) EngineSearchWithOptions(request *EngineSearchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EngineSearchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Grey) {
		body["grey"] = request.Grey
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.User) {
		body["user"] = request.User
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EngineSearch"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/platform/app/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EngineSearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Provides a search API operation for the CleverSee AI Search platform. After you create a search application on the platform, you can invoke this search API operation to retrieve images, documents, and audio/video content in datasets. The platform supports text (natural language), image, or combined text-and-image input, enabling fast adaptation to multi-modal large-scale data search scenarios and helping you efficiently locate target content.
//
// Description:
//
// ## Operation description
//
// This operation supports calling two types of search applications on the [CleverSee AI Search platform](https://aisearch.aliyun.com/web-search): **image search applications*	- (text-to-image, image-to-image, and combined text-and-image search) and **audio/video search applications*	- (text-to-audio/video, image-to-video, and combined text-and-image-to-video search).
//
// ### Data sources
//
// Audio and video data is supported. You can upload and update data through the [CleverSee AI Search platform](https://aisearch.aliyun.com/web-search) UI or through the [Dataset Data Add/Update API](https://www.alibabacloud.com/help/en/document_detail/3038471.html).
//
// # Authentication
//
// Call the CleverSee - Intelligent Search service by using the Alibaba Cloud SDK. For more information, see [AI search engine operation](https://api.aliyun.com/document/AiSearchEngine/2026-04-17/EngineSearch).
//
// @param request - EngineSearchRequest
//
// @return EngineSearchResponse
func (client *Client) EngineSearch(request *EngineSearchRequest) (_result *EngineSearchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EngineSearchResponse{}
	_body, _err := client.EngineSearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a temporary public network access URL for an image, audio, or video resource in your dataset.
//
// Description:
//
// Retrieves a temporary secure access URL for an underlying media resource (such as audio, video, or image) stored in AI Search Platform, based on the dataset ID and data primary key. The URL can be used directly for frontend display or download.
//
// Key use case: When you make a Search API call for a search application created in AI Search Platform, the returned image, audio, and video result URLs are pre-signed links with a validity period of 24 hours. If your application persists these URLs in local storage, subsequent access may fail because the URLs have expired. In this case, invoke this operation with the corresponding dataset ID and data record primary key to retrieve the latest valid access URL for the resource.
//
// @param request - GetDatasetResourceUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResourceUrlResponse
func (client *Client) GetDatasetResourceUrlWithOptions(request *GetDatasetResourceUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResourceUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["datasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.PrimaryKey) {
		body["primaryKey"] = request.PrimaryKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatasetResourceUrl"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/dataset/open/resources"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResourceUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a temporary public network access URL for an image, audio, or video resource in your dataset.
//
// Description:
//
// Retrieves a temporary secure access URL for an underlying media resource (such as audio, video, or image) stored in AI Search Platform, based on the dataset ID and data primary key. The URL can be used directly for frontend display or download.
//
// Key use case: When you make a Search API call for a search application created in AI Search Platform, the returned image, audio, and video result URLs are pre-signed links with a validity period of 24 hours. If your application persists these URLs in local storage, subsequent access may fail because the URLs have expired. In this case, invoke this operation with the corresponding dataset ID and data record primary key to retrieve the latest valid access URL for the resource.
//
// @param request - GetDatasetResourceUrlRequest
//
// @return GetDatasetResourceUrlResponse
func (client *Client) GetDatasetResourceUrl(request *GetDatasetResourceUrlRequest) (_result *GetDatasetResourceUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetResourceUrlResponse{}
	_body, _err := client.GetDatasetResourceUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or updates data in a specific dataset in real time. The system matches records based on the primary key value of the target dataset and the primary key value of the new data record (such as "pk_id" = "2026aa01"). If a matching primary key value is found, the corresponding data record is updated. If no match is found, a new data record is added.
//
// Description:
//
// *Common scenarios**
//
// | Scenario | Description |
//
// | --- | --- |
//
// | Real-time data addition | Pushes new data to the AI search platform in real time when the business system generates new data. |
//
// | Status update | Promptly updates data when changes occur on the business side, such as title modifications or delisting. |
//
// **Before you begin**
//
// - **Primary key handling**: The system determines whether to add or update a record based on the primary key type (user-defined or system-generated) of the target dataset.
//
// - **Batch limit**: A maximum of 100 records can be processed per request.
//
// - **Schema matching**: The `records` field must strictly follow the schema configured for the target dataset in the console.
//
// - **Permission requirements**: Make sure you have sufficient permissions to write to or update the target dataset.
//
// - **Status check**: Before sending a request, confirm that the target dataset is in a writable state and not in read-only mode.
//
// @param request - ImportDatasetDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDatasetDataResponse
func (client *Client) ImportDatasetDataWithOptions(request *ImportDatasetDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ImportDatasetDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetId) {
		body["datasetId"] = request.DatasetId
	}

	if !dara.IsNil(request.Records) {
		body["records"] = request.Records
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportDatasetData"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/dataset/open/upsert"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportDatasetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or updates data in a specific dataset in real time. The system matches records based on the primary key value of the target dataset and the primary key value of the new data record (such as "pk_id" = "2026aa01"). If a matching primary key value is found, the corresponding data record is updated. If no match is found, a new data record is added.
//
// Description:
//
// *Common scenarios**
//
// | Scenario | Description |
//
// | --- | --- |
//
// | Real-time data addition | Pushes new data to the AI search platform in real time when the business system generates new data. |
//
// | Status update | Promptly updates data when changes occur on the business side, such as title modifications or delisting. |
//
// **Before you begin**
//
// - **Primary key handling**: The system determines whether to add or update a record based on the primary key type (user-defined or system-generated) of the target dataset.
//
// - **Batch limit**: A maximum of 100 records can be processed per request.
//
// - **Schema matching**: The `records` field must strictly follow the schema configured for the target dataset in the console.
//
// - **Permission requirements**: Make sure you have sufficient permissions to write to or update the target dataset.
//
// - **Status check**: Before sending a request, confirm that the target dataset is in a writable state and not in read-only mode.
//
// @param request - ImportDatasetDataRequest
//
// @return ImportDatasetDataResponse
func (client *Client) ImportDatasetData(request *ImportDatasetDataRequest) (_result *ImportDatasetDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportDatasetDataResponse{}
	_body, _err := client.ImportDatasetDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Build Q&A applications leveraging powerful omni-modal search and comprehension capabilities on the CleverSee AI Search Platform, enabling deep understanding and precise Q&A over images, documents, and video content in datasets. The platform supports flexible customization of the Q&A assistant\\"s response style and interaction mode, allowing rapid adaptation to diverse large-scale data Q&A scenarios. Users can ask questions via text, images, or a combination of both, and the platform performs deep semantic understanding across the complete dataset, producing answers in multiple formats including text, images, and video. For video content, the platform also provides template-based output capabilities for generating customized content summaries, information extraction, and video scripts.
//
// Description:
//
// Streaming API for [CleverSee AI Search Platform](https://aisearch.aliyun.com/web-search) intelligent Q&A applications, supporting multimodal input (text, images, structured data) and streaming output (text, images, video, sources, etc.). The API uses the SSE (Server-Sent Events) protocol to push response data, where each data stream is a JSON object with different data types identified by the `type` field.
//
// ### Integration Guide:
//
// Streaming API for intelligent Q&A, supporting multimodal input (text, images, structured data) and streaming output (text, images, video, sources, etc.). The API uses the SSE (Server-Sent Events) protocol to push response data, where each data stream is a JSON object with different data types identified by the `type` field.
//
// ### Data Sources:
//
// Supports video Q&A. Data can be uploaded and updated through the [CleverSee AI Search Platform](https://aisearch.aliyun.com/web-search) product interface or via the [Dataset Data Add/Update API](https://help.aliyun.com/zh/document_detail/3038471.html?spm=a2c4g.11186623.help-menu-3037946.d_0_2_1_0.54ed1e97NGXVV1&scm=20140722.H_3038471._.OR_help-T_cn~zh-V_1).
//
// # Authentication
//
// Call the CleverSee - Intelligent Q&A service through the Alibaba Cloud SDK. For the detailed calling guide, please refer to: [AI Q&A Engine API](https://api.aliyun.com/document/AiSearchEngine/2026-04-17/QaChat)
//
// @param request - QaChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QaChatResponse
func (client *Client) QaChatWithSSE(request *QaChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *QaChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.qaChatWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// Build Q&A applications leveraging powerful omni-modal search and comprehension capabilities on the CleverSee AI Search Platform, enabling deep understanding and precise Q&A over images, documents, and video content in datasets. The platform supports flexible customization of the Q&A assistant\\"s response style and interaction mode, allowing rapid adaptation to diverse large-scale data Q&A scenarios. Users can ask questions via text, images, or a combination of both, and the platform performs deep semantic understanding across the complete dataset, producing answers in multiple formats including text, images, and video. For video content, the platform also provides template-based output capabilities for generating customized content summaries, information extraction, and video scripts.
//
// Description:
//
// Streaming API for [CleverSee AI Search Platform](https://aisearch.aliyun.com/web-search) intelligent Q&A applications, supporting multimodal input (text, images, structured data) and streaming output (text, images, video, sources, etc.). The API uses the SSE (Server-Sent Events) protocol to push response data, where each data stream is a JSON object with different data types identified by the `type` field.
//
// ### Integration Guide:
//
// Streaming API for intelligent Q&A, supporting multimodal input (text, images, structured data) and streaming output (text, images, video, sources, etc.). The API uses the SSE (Server-Sent Events) protocol to push response data, where each data stream is a JSON object with different data types identified by the `type` field.
//
// ### Data Sources:
//
// Supports video Q&A. Data can be uploaded and updated through the [CleverSee AI Search Platform](https://aisearch.aliyun.com/web-search) product interface or via the [Dataset Data Add/Update API](https://help.aliyun.com/zh/document_detail/3038471.html?spm=a2c4g.11186623.help-menu-3037946.d_0_2_1_0.54ed1e97NGXVV1&scm=20140722.H_3038471._.OR_help-T_cn~zh-V_1).
//
// # Authentication
//
// Call the CleverSee - Intelligent Q&A service through the Alibaba Cloud SDK. For the detailed calling guide, please refer to: [AI Q&A Engine API](https://api.aliyun.com/document/AiSearchEngine/2026-04-17/QaChat)
//
// @param request - QaChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QaChatResponse
func (client *Client) QaChatWithOptions(request *QaChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QaChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Message) {
		body["message"] = request.Message
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QaChat"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/platform/app/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QaChatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Build Q&A applications leveraging powerful omni-modal search and comprehension capabilities on the CleverSee AI Search Platform, enabling deep understanding and precise Q&A over images, documents, and video content in datasets. The platform supports flexible customization of the Q&A assistant\\"s response style and interaction mode, allowing rapid adaptation to diverse large-scale data Q&A scenarios. Users can ask questions via text, images, or a combination of both, and the platform performs deep semantic understanding across the complete dataset, producing answers in multiple formats including text, images, and video. For video content, the platform also provides template-based output capabilities for generating customized content summaries, information extraction, and video scripts.
//
// Description:
//
// Streaming API for [CleverSee AI Search Platform](https://aisearch.aliyun.com/web-search) intelligent Q&A applications, supporting multimodal input (text, images, structured data) and streaming output (text, images, video, sources, etc.). The API uses the SSE (Server-Sent Events) protocol to push response data, where each data stream is a JSON object with different data types identified by the `type` field.
//
// ### Integration Guide:
//
// Streaming API for intelligent Q&A, supporting multimodal input (text, images, structured data) and streaming output (text, images, video, sources, etc.). The API uses the SSE (Server-Sent Events) protocol to push response data, where each data stream is a JSON object with different data types identified by the `type` field.
//
// ### Data Sources:
//
// Supports video Q&A. Data can be uploaded and updated through the [CleverSee AI Search Platform](https://aisearch.aliyun.com/web-search) product interface or via the [Dataset Data Add/Update API](https://help.aliyun.com/zh/document_detail/3038471.html?spm=a2c4g.11186623.help-menu-3037946.d_0_2_1_0.54ed1e97NGXVV1&scm=20140722.H_3038471._.OR_help-T_cn~zh-V_1).
//
// # Authentication
//
// Call the CleverSee - Intelligent Q&A service through the Alibaba Cloud SDK. For the detailed calling guide, please refer to: [AI Q&A Engine API](https://api.aliyun.com/document/AiSearchEngine/2026-04-17/QaChat)
//
// @param request - QaChatRequest
//
// @return QaChatResponse
func (client *Client) QaChat(request *QaChatRequest) (_result *QaChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QaChatResponse{}
	_body, _err := client.QaChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) qaChatWithSSE_opYieldFunc(_yield chan *QaChatResponse, _yieldErr chan error, request *QaChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Message) {
		body["message"] = request.Message
	}

	if !dara.IsNil(request.Options) {
		body["options"] = request.Options
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QaChat"),
		Version:     dara.String("2026-04-17"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/platform/app/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
