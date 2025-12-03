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
	client.Endpoint, _err = client.GetEndpoint(dara.String("governance"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Applies an account baseline to multiple existing resource accounts at a time.
//
// Description:
//
// You can call this operation to apply an account baseline to existing resource accounts.
//
// Accounts are enrolled in the account factory in asynchronous mode. After a resource account is created, an account baseline is applied to the account. You can call the [GetEnrolledAccount](https://help.aliyun.com/document_detail/609062.html) operation to query the details of the account enrolled in the account factory and check whether the account baseline is applied to the account.
//
// @param request - BatchEnrollAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchEnrollAccountsResponse
func (client *Client) BatchEnrollAccountsWithOptions(request *BatchEnrollAccountsRequest, runtime *dara.RuntimeOptions) (_result *BatchEnrollAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Accounts) {
		query["Accounts"] = request.Accounts
	}

	if !dara.IsNil(request.BaselineId) {
		query["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.BaselineItems) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchEnrollAccounts"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchEnrollAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies an account baseline to multiple existing resource accounts at a time.
//
// Description:
//
// You can call this operation to apply an account baseline to existing resource accounts.
//
// Accounts are enrolled in the account factory in asynchronous mode. After a resource account is created, an account baseline is applied to the account. You can call the [GetEnrolledAccount](https://help.aliyun.com/document_detail/609062.html) operation to query the details of the account enrolled in the account factory and check whether the account baseline is applied to the account.
//
// @param request - BatchEnrollAccountsRequest
//
// @return BatchEnrollAccountsResponse
func (client *Client) BatchEnrollAccounts(request *BatchEnrollAccountsRequest) (_result *BatchEnrollAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchEnrollAccountsResponse{}
	_body, _err := client.BatchEnrollAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a baseline of the account factory.
//
// @param request - CreateAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAccountFactoryBaselineResponse
func (client *Client) CreateAccountFactoryBaselineWithOptions(request *CreateAccountFactoryBaselineRequest, runtime *dara.RuntimeOptions) (_result *CreateAccountFactoryBaselineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaselineItems) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !dara.IsNil(request.BaselineName) {
		query["BaselineName"] = request.BaselineName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAccountFactoryBaseline"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a baseline of the account factory.
//
// @param request - CreateAccountFactoryBaselineRequest
//
// @return CreateAccountFactoryBaselineResponse
func (client *Client) CreateAccountFactoryBaseline(request *CreateAccountFactoryBaselineRequest) (_result *CreateAccountFactoryBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAccountFactoryBaselineResponse{}
	_body, _err := client.CreateAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an account factory baseline.
//
// @param request - DeleteAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAccountFactoryBaselineResponse
func (client *Client) DeleteAccountFactoryBaselineWithOptions(request *DeleteAccountFactoryBaselineRequest, runtime *dara.RuntimeOptions) (_result *DeleteAccountFactoryBaselineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		query["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAccountFactoryBaseline"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an account factory baseline.
//
// @param request - DeleteAccountFactoryBaselineRequest
//
// @return DeleteAccountFactoryBaselineResponse
func (client *Client) DeleteAccountFactoryBaseline(request *DeleteAccountFactoryBaselineRequest) (_result *DeleteAccountFactoryBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAccountFactoryBaselineResponse{}
	_body, _err := client.DeleteAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enrolls an account. You can create a new account or manage an existing account in the account factory.
//
// Description:
//
// You can call this API operation to create a new account or manage an existing account and apply the account baseline to the account.
//
// Accounts are created in asynchronous mode. After you create an account, you can apply the account baseline to the account. You can call the [GetEnrolledAccount API](~~GetEnrolledAccount~~) operation to view the details about the account to obtain the result of applying the account baseline to the account.
//
// @param tmpReq - EnrollAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnrollAccountResponse
func (client *Client) EnrollAccountWithOptions(tmpReq *EnrollAccountRequest, runtime *dara.RuntimeOptions) (_result *EnrollAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnrollAccountShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tag) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, dara.String("Tag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountNamePrefix) {
		query["AccountNamePrefix"] = request.AccountNamePrefix
	}

	if !dara.IsNil(request.AccountUid) {
		query["AccountUid"] = request.AccountUid
	}

	if !dara.IsNil(request.BaselineId) {
		query["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.BaselineItems) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !dara.IsNil(request.DisplayName) {
		query["DisplayName"] = request.DisplayName
	}

	if !dara.IsNil(request.FolderId) {
		query["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.PayerAccountUid) {
		query["PayerAccountUid"] = request.PayerAccountUid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResellAccountType) {
		query["ResellAccountType"] = request.ResellAccountType
	}

	if !dara.IsNil(request.TagShrink) {
		query["Tag"] = request.TagShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnrollAccount"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnrollAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enrolls an account. You can create a new account or manage an existing account in the account factory.
//
// Description:
//
// You can call this API operation to create a new account or manage an existing account and apply the account baseline to the account.
//
// Accounts are created in asynchronous mode. After you create an account, you can apply the account baseline to the account. You can call the [GetEnrolledAccount API](~~GetEnrolledAccount~~) operation to view the details about the account to obtain the result of applying the account baseline to the account.
//
// @param request - EnrollAccountRequest
//
// @return EnrollAccountResponse
func (client *Client) EnrollAccount(request *EnrollAccountRequest) (_result *EnrollAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnrollAccountResponse{}
	_body, _err := client.EnrollAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of an account factory baseline.
//
// @param request - GetAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccountFactoryBaselineResponse
func (client *Client) GetAccountFactoryBaselineWithOptions(request *GetAccountFactoryBaselineRequest, runtime *dara.RuntimeOptions) (_result *GetAccountFactoryBaselineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		query["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccountFactoryBaseline"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of an account factory baseline.
//
// @param request - GetAccountFactoryBaselineRequest
//
// @return GetAccountFactoryBaselineResponse
func (client *Client) GetAccountFactoryBaseline(request *GetAccountFactoryBaselineRequest) (_result *GetAccountFactoryBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAccountFactoryBaselineResponse{}
	_body, _err := client.GetAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about an account that is enrolled in the account factory.
//
// @param request - GetEnrolledAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEnrolledAccountResponse
func (client *Client) GetEnrolledAccountWithOptions(request *GetEnrolledAccountRequest, runtime *dara.RuntimeOptions) (_result *GetEnrolledAccountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountUid) {
		query["AccountUid"] = request.AccountUid
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEnrolledAccount"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEnrolledAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about an account that is enrolled in the account factory.
//
// @param request - GetEnrolledAccountRequest
//
// @return GetEnrolledAccountResponse
func (client *Client) GetEnrolledAccount(request *GetEnrolledAccountRequest) (_result *GetEnrolledAccountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetEnrolledAccountResponse{}
	_body, _err := client.GetEnrolledAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of baseline items that are supported by the account factory of Cloud Governance Center (CGC).
//
// @param request - ListAccountFactoryBaselineItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountFactoryBaselineItemsResponse
func (client *Client) ListAccountFactoryBaselineItemsWithOptions(request *ListAccountFactoryBaselineItemsRequest, runtime *dara.RuntimeOptions) (_result *ListAccountFactoryBaselineItemsResponse, _err error) {
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

	if !dara.IsNil(request.Names) {
		query["Names"] = request.Names
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Versions) {
		query["Versions"] = request.Versions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccountFactoryBaselineItems"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountFactoryBaselineItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of baseline items that are supported by the account factory of Cloud Governance Center (CGC).
//
// @param request - ListAccountFactoryBaselineItemsRequest
//
// @return ListAccountFactoryBaselineItemsResponse
func (client *Client) ListAccountFactoryBaselineItems(request *ListAccountFactoryBaselineItemsRequest) (_result *ListAccountFactoryBaselineItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountFactoryBaselineItemsResponse{}
	_body, _err := client.ListAccountFactoryBaselineItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of baselines in the account factory.
//
// @param request - ListAccountFactoryBaselinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAccountFactoryBaselinesResponse
func (client *Client) ListAccountFactoryBaselinesWithOptions(request *ListAccountFactoryBaselinesRequest, runtime *dara.RuntimeOptions) (_result *ListAccountFactoryBaselinesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAccountFactoryBaselines"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAccountFactoryBaselinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of baselines in the account factory.
//
// @param request - ListAccountFactoryBaselinesRequest
//
// @return ListAccountFactoryBaselinesResponse
func (client *Client) ListAccountFactoryBaselines(request *ListAccountFactoryBaselinesRequest) (_result *ListAccountFactoryBaselinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAccountFactoryBaselinesResponse{}
	_body, _err := client.ListAccountFactoryBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of accounts that are enrolled in the account factory.
//
// @param request - ListEnrolledAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnrolledAccountsResponse
func (client *Client) ListEnrolledAccountsWithOptions(request *ListEnrolledAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListEnrolledAccountsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnrolledAccounts"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnrolledAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of accounts that are enrolled in the account factory.
//
// @param request - ListEnrolledAccountsRequest
//
// @return ListEnrolledAccountsResponse
func (client *Client) ListEnrolledAccounts(request *ListEnrolledAccountsRequest) (_result *ListEnrolledAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnrolledAccountsResponse{}
	_body, _err := client.ListEnrolledAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all available information about check items in a governance maturity check, including the name, ID, description, stage, resource metadata, and fixing guide.
//
// @param request - ListEvaluationMetadataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationMetadataResponse
func (client *Client) ListEvaluationMetadataWithOptions(request *ListEvaluationMetadataRequest, runtime *dara.RuntimeOptions) (_result *ListEvaluationMetadataResponse, _err error) {
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

	if !dara.IsNil(request.LensCode) {
		query["LensCode"] = request.LensCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TopicCode) {
		query["TopicCode"] = request.TopicCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluationMetadata"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluationMetadataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all available information about check items in a governance maturity check, including the name, ID, description, stage, resource metadata, and fixing guide.
//
// @param request - ListEvaluationMetadataRequest
//
// @return ListEvaluationMetadataResponse
func (client *Client) ListEvaluationMetadata(request *ListEvaluationMetadataRequest) (_result *ListEvaluationMetadataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEvaluationMetadataResponse{}
	_body, _err := client.ListEvaluationMetadataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the non-compliant resource information of a check item, including the name, ID, category, type, region, and related metadata of non-compliant resources.
//
// @param request - ListEvaluationMetricDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationMetricDetailsResponse
func (client *Client) ListEvaluationMetricDetailsWithOptions(request *ListEvaluationMetricDetailsRequest, runtime *dara.RuntimeOptions) (_result *ListEvaluationMetricDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluationMetricDetails"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluationMetricDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the non-compliant resource information of a check item, including the name, ID, category, type, region, and related metadata of non-compliant resources.
//
// @param request - ListEvaluationMetricDetailsRequest
//
// @return ListEvaluationMetricDetailsResponse
func (client *Client) ListEvaluationMetricDetails(request *ListEvaluationMetricDetailsRequest) (_result *ListEvaluationMetricDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEvaluationMetricDetailsResponse{}
	_body, _err := client.ListEvaluationMetricDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result and status of a governance check.
//
// @param request - ListEvaluationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationResultsResponse
func (client *Client) ListEvaluationResultsWithOptions(request *ListEvaluationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListEvaluationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.LensCode) {
		query["LensCode"] = request.LensCode
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	if !dara.IsNil(request.SnapshotId) {
		query["SnapshotId"] = request.SnapshotId
	}

	if !dara.IsNil(request.TopicCode) {
		query["TopicCode"] = request.TopicCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluationResults"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result and status of a governance check.
//
// @param request - ListEvaluationResultsRequest
//
// @return ListEvaluationResultsResponse
func (client *Client) ListEvaluationResults(request *ListEvaluationResultsRequest) (_result *ListEvaluationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEvaluationResultsResponse{}
	_body, _err := client.ListEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical scores of a governance maturity check.
//
// @param request - ListEvaluationScoreHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationScoreHistoryResponse
func (client *Client) ListEvaluationScoreHistoryWithOptions(request *ListEvaluationScoreHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListEvaluationScoreHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluationScoreHistory"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluationScoreHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical scores of a governance maturity check.
//
// @param request - ListEvaluationScoreHistoryRequest
//
// @return ListEvaluationScoreHistoryResponse
func (client *Client) ListEvaluationScoreHistory(request *ListEvaluationScoreHistoryRequest) (_result *ListEvaluationScoreHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEvaluationScoreHistoryResponse{}
	_body, _err := client.ListEvaluationScoreHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs a governance maturity check.
//
// @param tmpReq - RunEvaluationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunEvaluationResponse
func (client *Client) RunEvaluationWithOptions(tmpReq *RunEvaluationRequest, runtime *dara.RuntimeOptions) (_result *RunEvaluationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &RunEvaluationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MetricIds) {
		request.MetricIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MetricIds, dara.String("MetricIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountId) {
		query["AccountId"] = request.AccountId
	}

	if !dara.IsNil(request.MetricIdsShrink) {
		query["MetricIds"] = request.MetricIdsShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Scope) {
		query["Scope"] = request.Scope
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunEvaluation"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunEvaluationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs a governance maturity check.
//
// @param request - RunEvaluationRequest
//
// @return RunEvaluationResponse
func (client *Client) RunEvaluation(request *RunEvaluationRequest) (_result *RunEvaluationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunEvaluationResponse{}
	_body, _err := client.RunEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a baseline of the account factory.
//
// @param request - UpdateAccountFactoryBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAccountFactoryBaselineResponse
func (client *Client) UpdateAccountFactoryBaselineWithOptions(request *UpdateAccountFactoryBaselineRequest, runtime *dara.RuntimeOptions) (_result *UpdateAccountFactoryBaselineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		query["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.BaselineItems) {
		query["BaselineItems"] = request.BaselineItems
	}

	if !dara.IsNil(request.BaselineName) {
		query["BaselineName"] = request.BaselineName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAccountFactoryBaseline"),
		Version:     dara.String("2021-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAccountFactoryBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a baseline of the account factory.
//
// @param request - UpdateAccountFactoryBaselineRequest
//
// @return UpdateAccountFactoryBaselineResponse
func (client *Client) UpdateAccountFactoryBaseline(request *UpdateAccountFactoryBaselineRequest) (_result *UpdateAccountFactoryBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAccountFactoryBaselineResponse{}
	_body, _err := client.UpdateAccountFactoryBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
