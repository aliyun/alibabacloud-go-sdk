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
	client.EndpointRule = dara.String("central")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1": dara.String("workorder.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1": dara.String("workorder.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("workorder"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Closes a ticket.
//
// @param request - CloseTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseTicketResponse
func (client *Client) CloseTicketWithOptions(request *CloseTicketRequest, runtime *dara.RuntimeOptions) (_result *CloseTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes a ticket.
//
// @param request - CloseTicketRequest
//
// @return CloseTicketResponse
func (client *Client) CloseTicket(request *CloseTicketRequest) (_result *CloseTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseTicketResponse{}
	_body, _err := client.CloseTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ticket.
//
// @param tmpReq - CreateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTicketResponse
func (client *Client) CreateTicketWithOptions(tmpReq *CreateTicketRequest, runtime *dara.RuntimeOptions) (_result *CreateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileNameList) {
		request.FileNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileNameList, dara.String("FileNameList"), dara.String("simple"))
	}

	if !dara.IsNil(tmpReq.SecretInfo) {
		request.SecretInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SecretInfo, dara.String("SecretInfo"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SecretInfoShrink) {
		query["SecretInfo"] = request.SecretInfoShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Email) {
		body["Email"] = request.Email
	}

	if !dara.IsNil(request.FileNameListShrink) {
		body["FileNameList"] = request.FileNameListShrink
	}

	if !dara.IsNil(request.Severity) {
		body["Severity"] = request.Severity
	}

	if !dara.IsNil(request.Title) {
		body["Title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket.
//
// @param request - CreateTicketRequest
//
// @return CreateTicketResponse
func (client *Client) CreateTicket(request *CreateTicketRequest) (_result *CreateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTicketResponse{}
	_body, _err := client.CreateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Evaluates a ticket.
//
// @param request - EvaluateTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvaluateTicketResponse
func (client *Client) EvaluateTicketWithOptions(request *EvaluateTicketRequest, runtime *dara.RuntimeOptions) (_result *EvaluateTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Score) {
		body["Score"] = request.Score
	}

	if !dara.IsNil(request.Solved) {
		body["Solved"] = request.Solved
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EvaluateTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EvaluateTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Evaluates a ticket.
//
// @param request - EvaluateTicketRequest
//
// @return EvaluateTicketResponse
func (client *Client) EvaluateTicket(request *EvaluateTicketRequest) (_result *EvaluateTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EvaluateTicketResponse{}
	_body, _err := client.EvaluateTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Object Storage Service (OSS) URL that is used to upload attachments.
//
// Description:
//
// > Attachment upload
//
// >
//
// > - Use a presigned URL to upload the file. For more information, see the OSS documentation.
//
// After you obtain the temporary upload URL, you can upload the file as follows:
//
// <details>
//
// <summary>
//
// # Sample
//
// </summary>
//
// curl -X PUT -H "Content-Type:application/octet-stream" -H "x-oss-meta-author:aliy" -T /home/admin/file.png "https\\://xx.oss-cn-zhangjiakou.aliyuncs.com/xx/xx.png"
//
// </details>
//
// @param request - GetAttachmentUploadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAttachmentUploadUrlResponse
func (client *Client) GetAttachmentUploadUrlWithOptions(request *GetAttachmentUploadUrlRequest, runtime *dara.RuntimeOptions) (_result *GetAttachmentUploadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAttachmentUploadUrl"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAttachmentUploadUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Object Storage Service (OSS) URL that is used to upload attachments.
//
// Description:
//
// > Attachment upload
//
// >
//
// > - Use a presigned URL to upload the file. For more information, see the OSS documentation.
//
// After you obtain the temporary upload URL, you can upload the file as follows:
//
// <details>
//
// <summary>
//
// # Sample
//
// </summary>
//
// curl -X PUT -H "Content-Type:application/octet-stream" -H "x-oss-meta-author:aliy" -T /home/admin/file.png "https\\://xx.oss-cn-zhangjiakou.aliyuncs.com/xx/xx.png"
//
// </details>
//
// @param request - GetAttachmentUploadUrlRequest
//
// @return GetAttachmentUploadUrlResponse
func (client *Client) GetAttachmentUploadUrl(request *GetAttachmentUploadUrlRequest) (_result *GetAttachmentUploadUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAttachmentUploadUrlResponse{}
	_body, _err := client.GetAttachmentUploadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the tag for a ticket\\"s RocketMQ message listener.
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMqConsumerTagResponse
func (client *Client) GetMqConsumerTagWithOptions(runtime *dara.RuntimeOptions) (_result *GetMqConsumerTagResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetMqConsumerTag"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMqConsumerTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the tag for a ticket\\"s RocketMQ message listener.
//
// @return GetMqConsumerTagResponse
func (client *Client) GetMqConsumerTag() (_result *GetMqConsumerTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMqConsumerTagResponse{}
	_body, _err := client.GetMqConsumerTagWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a ticket.
//
// @param request - GetTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTicketResponse
func (client *Client) GetTicketWithOptions(request *GetTicketRequest, runtime *dara.RuntimeOptions) (_result *GetTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a ticket.
//
// @param request - GetTicketRequest
//
// @return GetTicketResponse
func (client *Client) GetTicket(request *GetTicketRequest) (_result *GetTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTicketResponse{}
	_body, _err := client.GetTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the list data of ticket problem categories.
//
// @param request - ListCategoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCategoriesResponse
func (client *Client) ListCategoriesWithOptions(request *ListCategoriesRequest, runtime *dara.RuntimeOptions) (_result *ListCategoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProductId) {
		body["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCategories"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list data of ticket problem categories.
//
// @param request - ListCategoriesRequest
//
// @return ListCategoriesResponse
func (client *Client) ListCategories(request *ListCategoriesRequest) (_result *ListCategoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCategoriesResponse{}
	_body, _err := client.ListCategoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the data of the Alibaba Cloud product list.
//
// @param request - ListProductsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProductsResponse
func (client *Client) ListProductsWithOptions(request *ListProductsRequest, runtime *dara.RuntimeOptions) (_result *ListProductsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProducts"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the data of the Alibaba Cloud product list.
//
// @param request - ListProductsRequest
//
// @return ListProductsResponse
func (client *Client) ListProducts(request *ListProductsRequest) (_result *ListProductsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProductsResponse{}
	_body, _err := client.ListProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the communication records for a ticket.
//
// @param request - ListTicketNotesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketNotesResponse
func (client *Client) ListTicketNotesWithOptions(request *ListTicketNotesRequest, runtime *dara.RuntimeOptions) (_result *ListTicketNotesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TicketId) {
		query["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTicketNotes"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the communication records for a ticket.
//
// @param request - ListTicketNotesRequest
//
// @return ListTicketNotesResponse
func (client *Client) ListTicketNotes(request *ListTicketNotesRequest) (_result *ListTicketNotesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTicketNotesResponse{}
	_body, _err := client.ListTicketNotesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of your tickets.
//
// @param tmpReq - ListTicketsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTicketsResponse
func (client *Client) ListTicketsWithOptions(tmpReq *ListTicketsRequest, runtime *dara.RuntimeOptions) (_result *ListTicketsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTicketsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TicketIdList) {
		request.TicketIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TicketIdList, dara.String("TicketIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.StatusList) {
		body["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.TicketIdListShrink) {
		body["TicketIdList"] = request.TicketIdListShrink
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTickets"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTicketsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of your tickets.
//
// @param request - ListTicketsRequest
//
// @return ListTicketsResponse
func (client *Client) ListTickets(request *ListTicketsRequest) (_result *ListTicketsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTicketsResponse{}
	_body, _err := client.ListTicketsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reopens a ticket.
//
// @param request - ReopenTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReopenTicketResponse
func (client *Client) ReopenTicketWithOptions(request *ReopenTicketRequest, runtime *dara.RuntimeOptions) (_result *ReopenTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReopenTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReopenTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reopens a ticket.
//
// @param request - ReopenTicketRequest
//
// @return ReopenTicketResponse
func (client *Client) ReopenTicket(request *ReopenTicketRequest) (_result *ReopenTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReopenTicketResponse{}
	_body, _err := client.ReopenTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Replies to a ticket. You can call the ListTicketNotes operation to retrieve the reply.
//
// @param tmpReq - ReplyTicketRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplyTicketResponse
func (client *Client) ReplyTicketWithOptions(tmpReq *ReplyTicketRequest, runtime *dara.RuntimeOptions) (_result *ReplyTicketResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ReplyTicketShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FileNameList) {
		request.FileNameListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FileNameList, dara.String("FileNameList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FileNameListShrink) {
		query["FileNameList"] = request.FileNameListShrink
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.Encrypt) {
		body["Encrypt"] = request.Encrypt
	}

	if !dara.IsNil(request.TicketId) {
		body["TicketId"] = request.TicketId
	}

	if !dara.IsNil(request.Uid) {
		body["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReplyTicket"),
		Version:     dara.String("2021-06-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReplyTicketResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Replies to a ticket. You can call the ListTicketNotes operation to retrieve the reply.
//
// @param request - ReplyTicketRequest
//
// @return ReplyTicketResponse
func (client *Client) ReplyTicket(request *ReplyTicketRequest) (_result *ReplyTicketResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReplyTicketResponse{}
	_body, _err := client.ReplyTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
