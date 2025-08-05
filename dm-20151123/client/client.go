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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dm"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Add IP Protection Information
//
// @param request - AddIpfilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddIpfilterResponse
func (client *Client) AddIpfilterWithOptions(request *AddIpfilterRequest, runtime *dara.RuntimeOptions) (_result *AddIpfilterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpAddress) {
		query["IpAddress"] = request.IpAddress
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
		Action:      dara.String("AddIpfilter"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddIpfilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Add IP Protection Information
//
// @param request - AddIpfilterRequest
//
// @return AddIpfilterResponse
func (client *Client) AddIpfilter(request *AddIpfilterRequest) (_result *AddIpfilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddIpfilterResponse{}
	_body, _err := client.AddIpfilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Verify Reply Address
//
// @param request - ApproveReplyMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveReplyMailAddressResponse
func (client *Client) ApproveReplyMailAddressWithOptions(request *ApproveReplyMailAddressRequest, runtime *dara.RuntimeOptions) (_result *ApproveReplyMailAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Ticket) {
		query["Ticket"] = request.Ticket
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveReplyMailAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveReplyMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Verify Reply Address
//
// @param request - ApproveReplyMailAddressRequest
//
// @return ApproveReplyMailAddressResponse
func (client *Client) ApproveReplyMailAddress(request *ApproveReplyMailAddressRequest) (_result *ApproveReplyMailAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApproveReplyMailAddressResponse{}
	_body, _err := client.ApproveReplyMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Batch Send Emails
//
// @param request - BatchSendMailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSendMailResponse
func (client *Client) BatchSendMailWithOptions(request *BatchSendMailRequest, runtime *dara.RuntimeOptions) (_result *BatchSendMailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AddressType) {
		query["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.ClickTrace) {
		query["ClickTrace"] = request.ClickTrace
	}

	if !dara.IsNil(request.Headers) {
		query["Headers"] = request.Headers
	}

	if !dara.IsNil(request.IpPoolId) {
		query["IpPoolId"] = request.IpPoolId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReceiversName) {
		query["ReceiversName"] = request.ReceiversName
	}

	if !dara.IsNil(request.ReplyAddress) {
		query["ReplyAddress"] = request.ReplyAddress
	}

	if !dara.IsNil(request.ReplyAddressAlias) {
		query["ReplyAddressAlias"] = request.ReplyAddressAlias
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.UnSubscribeFilterLevel) {
		query["UnSubscribeFilterLevel"] = request.UnSubscribeFilterLevel
	}

	if !dara.IsNil(request.UnSubscribeLinkType) {
		query["UnSubscribeLinkType"] = request.UnSubscribeLinkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSendMail"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSendMailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Batch Send Emails
//
// @param request - BatchSendMailRequest
//
// @return BatchSendMailResponse
func (client *Client) BatchSendMail(request *BatchSendMailRequest) (_result *BatchSendMailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchSendMailResponse{}
	_body, _err := client.BatchSendMailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改域名DKIM记录
//
// @param request - ChangeDomainDkimRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeDomainDkimRecordResponse
func (client *Client) ChangeDomainDkimRecordWithOptions(request *ChangeDomainDkimRecordRequest, runtime *dara.RuntimeOptions) (_result *ChangeDomainDkimRecordResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DkimRsaLength) {
		query["DkimRsaLength"] = request.DkimRsaLength
	}

	if !dara.IsNil(request.Domain) {
		query["Domain"] = request.Domain
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
		Action:      dara.String("ChangeDomainDkimRecord"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeDomainDkimRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改域名DKIM记录
//
// @param request - ChangeDomainDkimRecordRequest
//
// @return ChangeDomainDkimRecordResponse
func (client *Client) ChangeDomainDkimRecord(request *ChangeDomainDkimRecordRequest) (_result *ChangeDomainDkimRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeDomainDkimRecordResponse{}
	_body, _err := client.ChangeDomainDkimRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Check Domain Status
//
// @param request - CheckDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDomainResponse
func (client *Client) CheckDomainWithOptions(request *CheckDomainRequest, runtime *dara.RuntimeOptions) (_result *CheckDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
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
		Action:      dara.String("CheckDomain"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Check Domain Status
//
// @param request - CheckDomainRequest
//
// @return CheckDomainResponse
func (client *Client) CheckDomain(request *CheckDomainRequest) (_result *CheckDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDomainResponse{}
	_body, _err := client.CheckDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Validate Reply-To Address
//
// @param request - CheckReplyToMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckReplyToMailAddressResponse
func (client *Client) CheckReplyToMailAddressWithOptions(request *CheckReplyToMailAddressRequest, runtime *dara.RuntimeOptions) (_result *CheckReplyToMailAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MailAddressId) {
		query["MailAddressId"] = request.MailAddressId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
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
		Action:      dara.String("CheckReplyToMailAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckReplyToMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Validate Reply-To Address
//
// @param request - CheckReplyToMailAddressRequest
//
// @return CheckReplyToMailAddressResponse
func (client *Client) CheckReplyToMailAddress(request *CheckReplyToMailAddressRequest) (_result *CheckReplyToMailAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckReplyToMailAddressResponse{}
	_body, _err := client.CheckReplyToMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Domain
//
// @param request - CreateDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, runtime *dara.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
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

	if !dara.IsNil(request.DkimSelector) {
		query["dkimSelector"] = request.DkimSelector
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDomain"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Domain
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create a mail address.
//
// @param request - CreateMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMailAddressResponse
func (client *Client) CreateMailAddressWithOptions(request *CreateMailAddressRequest, runtime *dara.RuntimeOptions) (_result *CreateMailAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReplyAddress) {
		query["ReplyAddress"] = request.ReplyAddress
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Sendtype) {
		query["Sendtype"] = request.Sendtype
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMailAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a mail address.
//
// @param request - CreateMailAddressRequest
//
// @return CreateMailAddressResponse
func (client *Client) CreateMailAddress(request *CreateMailAddressRequest) (_result *CreateMailAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMailAddressResponse{}
	_body, _err := client.CreateMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Receiver List
//
// @param request - CreateReceiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReceiverResponse
func (client *Client) CreateReceiverWithOptions(request *CreateReceiverRequest, runtime *dara.RuntimeOptions) (_result *CreateReceiverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Desc) {
		query["Desc"] = request.Desc
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReceiversAlias) {
		query["ReceiversAlias"] = request.ReceiversAlias
	}

	if !dara.IsNil(request.ReceiversName) {
		query["ReceiversName"] = request.ReceiversName
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
		Action:      dara.String("CreateReceiver"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateReceiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Receiver List
//
// @param request - CreateReceiverRequest
//
// @return CreateReceiverResponse
func (client *Client) CreateReceiver(request *CreateReceiverRequest) (_result *CreateReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateReceiverResponse{}
	_body, _err := client.CreateReceiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Tag
//
// @param request - CreateTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTagResponse
func (client *Client) CreateTagWithOptions(request *CreateTagRequest, runtime *dara.RuntimeOptions) (_result *CreateTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagDescription) {
		query["TagDescription"] = request.TagDescription
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTag"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Tag
//
// @param request - CreateTagRequest
//
// @return CreateTagResponse
func (client *Client) CreateTag(request *CreateTagRequest) (_result *CreateTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTagResponse{}
	_body, _err := client.CreateTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create User\\"s Invalid Address
//
// @param request - CreateUserSuppressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserSuppressionResponse
func (client *Client) CreateUserSuppressionWithOptions(request *CreateUserSuppressionRequest, runtime *dara.RuntimeOptions) (_result *CreateUserSuppressionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
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
		Action:      dara.String("CreateUserSuppression"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUserSuppressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create User\\"s Invalid Address
//
// @param request - CreateUserSuppressionRequest
//
// @return CreateUserSuppressionResponse
func (client *Client) CreateUserSuppression(request *CreateUserSuppressionRequest) (_result *CreateUserSuppressionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUserSuppressionResponse{}
	_body, _err := client.CreateUserSuppressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Set Dedicated IP Auto Renewal
//
// @param request - DedicatedIpAutoRenewalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpAutoRenewalResponse
func (client *Client) DedicatedIpAutoRenewalWithOptions(request *DedicatedIpAutoRenewalRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpAutoRenewalResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoRenewal) {
		query["AutoRenewal"] = request.AutoRenewal
	}

	if !dara.IsNil(request.BuyResourceIds) {
		query["BuyResourceIds"] = request.BuyResourceIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpAutoRenewal"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpAutoRenewalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Set Dedicated IP Auto Renewal
//
// @param request - DedicatedIpAutoRenewalRequest
//
// @return DedicatedIpAutoRenewalResponse
func (client *Client) DedicatedIpAutoRenewal(request *DedicatedIpAutoRenewalRequest) (_result *DedicatedIpAutoRenewalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpAutoRenewalResponse{}
	_body, _err := client.DedicatedIpAutoRenewalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Change the warmup method for a dedicated IP
//
// @param request - DedicatedIpChangeWarmupTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpChangeWarmupTypeResponse
func (client *Client) DedicatedIpChangeWarmupTypeWithOptions(request *DedicatedIpChangeWarmupTypeRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpChangeWarmupTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.WarmupType) {
		query["WarmupType"] = request.WarmupType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpChangeWarmupType"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpChangeWarmupTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Change the warmup method for a dedicated IP
//
// @param request - DedicatedIpChangeWarmupTypeRequest
//
// @return DedicatedIpChangeWarmupTypeResponse
func (client *Client) DedicatedIpChangeWarmupType(request *DedicatedIpChangeWarmupTypeRequest) (_result *DedicatedIpChangeWarmupTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpChangeWarmupTypeResponse{}
	_body, _err := client.DedicatedIpChangeWarmupTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Dedicated IP User IP List
//
// @param request - DedicatedIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpListResponse
func (client *Client) DedicatedIpListWithOptions(request *DedicatedIpListRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpList"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Dedicated IP User IP List
//
// @param request - DedicatedIpListRequest
//
// @return DedicatedIpListResponse
func (client *Client) DedicatedIpList(request *DedicatedIpListRequest) (_result *DedicatedIpListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpListResponse{}
	_body, _err := client.DedicatedIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Purchased Independent IPs Not Added to Pool
//
// @param request - DedicatedIpNonePoolListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpNonePoolListResponse
func (client *Client) DedicatedIpNonePoolListWithOptions(runtime *dara.RuntimeOptions) (_result *DedicatedIpNonePoolListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpNonePoolList"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpNonePoolListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Purchased Independent IPs Not Added to Pool
//
// @return DedicatedIpNonePoolListResponse
func (client *Client) DedicatedIpNonePoolList() (_result *DedicatedIpNonePoolListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpNonePoolListResponse{}
	_body, _err := client.DedicatedIpNonePoolListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Creation of Independent IP Pool
//
// @param request - DedicatedIpPoolCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpPoolCreateResponse
func (client *Client) DedicatedIpPoolCreateWithOptions(request *DedicatedIpPoolCreateRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpPoolCreateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuyResourceIds) {
		query["BuyResourceIds"] = request.BuyResourceIds
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpPoolCreate"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpPoolCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Creation of Independent IP Pool
//
// @param request - DedicatedIpPoolCreateRequest
//
// @return DedicatedIpPoolCreateResponse
func (client *Client) DedicatedIpPoolCreate(request *DedicatedIpPoolCreateRequest) (_result *DedicatedIpPoolCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpPoolCreateResponse{}
	_body, _err := client.DedicatedIpPoolCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 独立IP池删除
//
// @param request - DedicatedIpPoolDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpPoolDeleteResponse
func (client *Client) DedicatedIpPoolDeleteWithOptions(request *DedicatedIpPoolDeleteRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpPoolDeleteResponse, _err error) {
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
		Action:      dara.String("DedicatedIpPoolDelete"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpPoolDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 独立IP池删除
//
// @param request - DedicatedIpPoolDeleteRequest
//
// @return DedicatedIpPoolDeleteResponse
func (client *Client) DedicatedIpPoolDelete(request *DedicatedIpPoolDeleteRequest) (_result *DedicatedIpPoolDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpPoolDeleteResponse{}
	_body, _err := client.DedicatedIpPoolDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Dedicated IP Pool List
//
// @param request - DedicatedIpPoolListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpPoolListResponse
func (client *Client) DedicatedIpPoolListWithOptions(request *DedicatedIpPoolListRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpPoolListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpPoolList"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpPoolListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Dedicated IP Pool List
//
// @param request - DedicatedIpPoolListRequest
//
// @return DedicatedIpPoolListResponse
func (client *Client) DedicatedIpPoolList(request *DedicatedIpPoolListRequest) (_result *DedicatedIpPoolListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpPoolListResponse{}
	_body, _err := client.DedicatedIpPoolListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update of dedicated IP Pool
//
// @param request - DedicatedIpPoolUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DedicatedIpPoolUpdateResponse
func (client *Client) DedicatedIpPoolUpdateWithOptions(request *DedicatedIpPoolUpdateRequest, runtime *dara.RuntimeOptions) (_result *DedicatedIpPoolUpdateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BuyResourceIds) {
		query["BuyResourceIds"] = request.BuyResourceIds
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.UpdateResource) {
		query["UpdateResource"] = request.UpdateResource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DedicatedIpPoolUpdate"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DedicatedIpPoolUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update of dedicated IP Pool
//
// @param request - DedicatedIpPoolUpdateRequest
//
// @return DedicatedIpPoolUpdateResponse
func (client *Client) DedicatedIpPoolUpdate(request *DedicatedIpPoolUpdateRequest) (_result *DedicatedIpPoolUpdateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DedicatedIpPoolUpdateResponse{}
	_body, _err := client.DedicatedIpPoolUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Domain
//
// @param request - DeleteDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *dara.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
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
		Action:      dara.String("DeleteDomain"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Domain
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Remove invalid addresses from the invalid address database
//
// @param request - DeleteInvalidAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInvalidAddressResponse
func (client *Client) DeleteInvalidAddressWithOptions(request *DeleteInvalidAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteInvalidAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ToAddress) {
		query["ToAddress"] = request.ToAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInvalidAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInvalidAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Remove invalid addresses from the invalid address database
//
// @param request - DeleteInvalidAddressRequest
//
// @return DeleteInvalidAddressResponse
func (client *Client) DeleteInvalidAddress(request *DeleteInvalidAddressRequest) (_result *DeleteInvalidAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInvalidAddressResponse{}
	_body, _err := client.DeleteInvalidAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete IP Protection Information
//
// @param request - DeleteIpfilterByEdmIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIpfilterByEdmIdResponse
func (client *Client) DeleteIpfilterByEdmIdWithOptions(request *DeleteIpfilterByEdmIdRequest, runtime *dara.RuntimeOptions) (_result *DeleteIpfilterByEdmIdResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FromType) {
		query["FromType"] = request.FromType
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
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
		Action:      dara.String("DeleteIpfilterByEdmId"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIpfilterByEdmIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete IP Protection Information
//
// @param request - DeleteIpfilterByEdmIdRequest
//
// @return DeleteIpfilterByEdmIdResponse
func (client *Client) DeleteIpfilterByEdmId(request *DeleteIpfilterByEdmIdRequest) (_result *DeleteIpfilterByEdmIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteIpfilterByEdmIdResponse{}
	_body, _err := client.DeleteIpfilterByEdmIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Mail Address
//
// @param request - DeleteMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMailAddressResponse
func (client *Client) DeleteMailAddressWithOptions(request *DeleteMailAddressRequest, runtime *dara.RuntimeOptions) (_result *DeleteMailAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MailAddressId) {
		query["MailAddressId"] = request.MailAddressId
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
		Action:      dara.String("DeleteMailAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Mail Address
//
// @param request - DeleteMailAddressRequest
//
// @return DeleteMailAddressResponse
func (client *Client) DeleteMailAddress(request *DeleteMailAddressRequest) (_result *DeleteMailAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMailAddressResponse{}
	_body, _err := client.DeleteMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Receiver List
//
// @param request - DeleteReceiverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReceiverResponse
func (client *Client) DeleteReceiverWithOptions(request *DeleteReceiverRequest, runtime *dara.RuntimeOptions) (_result *DeleteReceiverResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
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
		Action:      dara.String("DeleteReceiver"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReceiverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Receiver List
//
// @param request - DeleteReceiverRequest
//
// @return DeleteReceiverResponse
func (client *Client) DeleteReceiver(request *DeleteReceiverRequest) (_result *DeleteReceiverResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteReceiverResponse{}
	_body, _err := client.DeleteReceiverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete a Single Recipient
//
// @param request - DeleteReceiverDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteReceiverDetailResponse
func (client *Client) DeleteReceiverDetailWithOptions(request *DeleteReceiverDetailRequest, runtime *dara.RuntimeOptions) (_result *DeleteReceiverDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
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
		Action:      dara.String("DeleteReceiverDetail"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteReceiverDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete a Single Recipient
//
// @param request - DeleteReceiverDetailRequest
//
// @return DeleteReceiverDetailResponse
func (client *Client) DeleteReceiverDetail(request *DeleteReceiverDetailRequest) (_result *DeleteReceiverDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteReceiverDetailResponse{}
	_body, _err := client.DeleteReceiverDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Tag
//
// @param request - DeleteTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTagResponse
func (client *Client) DeleteTagWithOptions(request *DeleteTagRequest, runtime *dara.RuntimeOptions) (_result *DeleteTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTag"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Tag
//
// @param request - DeleteTagRequest
//
// @return DeleteTagResponse
func (client *Client) DeleteTag(request *DeleteTagRequest) (_result *DeleteTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTagResponse{}
	_body, _err := client.DeleteTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieve account information.
//
// @param request - DescAccountSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescAccountSummaryResponse
func (client *Client) DescAccountSummaryWithOptions(request *DescAccountSummaryRequest, runtime *dara.RuntimeOptions) (_result *DescAccountSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("DescAccountSummary"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescAccountSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieve account information.
//
// @param request - DescAccountSummaryRequest
//
// @return DescAccountSummaryResponse
func (client *Client) DescAccountSummary(request *DescAccountSummaryRequest) (_result *DescAccountSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescAccountSummaryResponse{}
	_body, _err := client.DescAccountSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Domain Details
//
// @param request - DescDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescDomainResponse
func (client *Client) DescDomainWithOptions(request *DescDomainRequest, runtime *dara.RuntimeOptions) (_result *DescDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainId) {
		query["DomainId"] = request.DomainId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RequireRealTimeDnsRecords) {
		query["RequireRealTimeDnsRecords"] = request.RequireRealTimeDnsRecords
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
		Action:      dara.String("DescDomain"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Domain Details
//
// @param request - DescDomainRequest
//
// @return DescDomainResponse
func (client *Client) DescDomain(request *DescDomainRequest) (_result *DescDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescDomainResponse{}
	_body, _err := client.DescDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取专属ip的预热详情信息
//
// @param request - GetDedicatedIpWarmUpDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDedicatedIpWarmUpDetailResponse
func (client *Client) GetDedicatedIpWarmUpDetailWithOptions(request *GetDedicatedIpWarmUpDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDedicatedIpWarmUpDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedIp) {
		query["DedicatedIp"] = request.DedicatedIp
	}

	if !dara.IsNil(request.EndDayMark) {
		query["EndDayMark"] = request.EndDayMark
	}

	if !dara.IsNil(request.Esp) {
		query["Esp"] = request.Esp
	}

	if !dara.IsNil(request.StartDayMark) {
		query["StartDayMark"] = request.StartDayMark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDedicatedIpWarmUpDetail"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDedicatedIpWarmUpDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取专属ip的预热详情信息
//
// @param request - GetDedicatedIpWarmUpDetailRequest
//
// @return GetDedicatedIpWarmUpDetailResponse
func (client *Client) GetDedicatedIpWarmUpDetail(request *GetDedicatedIpWarmUpDetailRequest) (_result *GetDedicatedIpWarmUpDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDedicatedIpWarmUpDetailResponse{}
	_body, _err := client.GetDedicatedIpWarmUpDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取专属ip的预热信息
//
// @param request - GetDedicatedIpWarmUpInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDedicatedIpWarmUpInfoResponse
func (client *Client) GetDedicatedIpWarmUpInfoWithOptions(request *GetDedicatedIpWarmUpInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDedicatedIpWarmUpInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DedicatedIp) {
		query["DedicatedIp"] = request.DedicatedIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDedicatedIpWarmUpInfo"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDedicatedIpWarmUpInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取专属ip的预热信息
//
// @param request - GetDedicatedIpWarmUpInfoRequest
//
// @return GetDedicatedIpWarmUpInfoResponse
func (client *Client) GetDedicatedIpWarmUpInfo(request *GetDedicatedIpWarmUpInfoRequest) (_result *GetDedicatedIpWarmUpInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDedicatedIpWarmUpInfoResponse{}
	_body, _err := client.GetDedicatedIpWarmUpInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get IP Protection Information
//
// @param request - GetIpProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpProtectionResponse
func (client *Client) GetIpProtectionWithOptions(request *GetIpProtectionRequest, runtime *dara.RuntimeOptions) (_result *GetIpProtectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetIpProtection"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIpProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get IP Protection Information
//
// @param request - GetIpProtectionRequest
//
// @return GetIpProtectionResponse
func (client *Client) GetIpProtection(request *GetIpProtectionRequest) (_result *GetIpProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIpProtectionResponse{}
	_body, _err := client.GetIpProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve IP Protection Information
//
// @param request - GetIpfilterListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIpfilterListResponse
func (client *Client) GetIpfilterListWithOptions(request *GetIpfilterListRequest, runtime *dara.RuntimeOptions) (_result *GetIpfilterListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetIpfilterList"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIpfilterListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve IP Protection Information
//
// @param request - GetIpfilterListRequest
//
// @return GetIpfilterListResponse
func (client *Client) GetIpfilterList(request *GetIpfilterListRequest) (_result *GetIpfilterListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIpfilterListResponse{}
	_body, _err := client.GetIpfilterListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户无效地址级别配置
//
// @param request - GetSuppressionListLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuppressionListLevelResponse
func (client *Client) GetSuppressionListLevelWithOptions(request *GetSuppressionListLevelRequest, runtime *dara.RuntimeOptions) (_result *GetSuppressionListLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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
		Action:      dara.String("GetSuppressionListLevel"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSuppressionListLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户无效地址级别配置
//
// @param request - GetSuppressionListLevelRequest
//
// @return GetSuppressionListLevelResponse
func (client *Client) GetSuppressionListLevel(request *GetSuppressionListLevelRequest) (_result *GetSuppressionListLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSuppressionListLevelResponse{}
	_body, _err := client.GetSuppressionListLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get tracking information
//
// @param request - GetTrackListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrackListResponse
func (client *Client) GetTrackListWithOptions(request *GetTrackListRequest, runtime *dara.RuntimeOptions) (_result *GetTrackListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DedicatedIp) {
		query["DedicatedIp"] = request.DedicatedIp
	}

	if !dara.IsNil(request.DedicatedIpPoolId) {
		query["DedicatedIpPoolId"] = request.DedicatedIpPoolId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Esp) {
		query["Esp"] = request.Esp
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OffsetCreateTime) {
		query["OffsetCreateTime"] = request.OffsetCreateTime
	}

	if !dara.IsNil(request.OffsetCreateTimeDesc) {
		query["OffsetCreateTimeDesc"] = request.OffsetCreateTimeDesc
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	if !dara.IsNil(request.Total) {
		query["Total"] = request.Total
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrackList"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get tracking information
//
// @param request - GetTrackListRequest
//
// @return GetTrackListResponse
func (client *Client) GetTrackList(request *GetTrackListRequest) (_result *GetTrackListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTrackListResponse{}
	_body, _err := client.GetTrackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get tracking information based on the sender address and tag name
//
// @param request - GetTrackListByMailFromAndTagNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTrackListByMailFromAndTagNameResponse
func (client *Client) GetTrackListByMailFromAndTagNameWithOptions(request *GetTrackListByMailFromAndTagNameRequest, runtime *dara.RuntimeOptions) (_result *GetTrackListByMailFromAndTagNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DedicatedIp) {
		query["DedicatedIp"] = request.DedicatedIp
	}

	if !dara.IsNil(request.DedicatedIpPoolId) {
		query["DedicatedIpPoolId"] = request.DedicatedIpPoolId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Esp) {
		query["Esp"] = request.Esp
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OffsetCreateTime) {
		query["OffsetCreateTime"] = request.OffsetCreateTime
	}

	if !dara.IsNil(request.OffsetCreateTimeDesc) {
		query["OffsetCreateTimeDesc"] = request.OffsetCreateTimeDesc
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	if !dara.IsNil(request.Total) {
		query["Total"] = request.Total
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTrackListByMailFromAndTagName"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTrackListByMailFromAndTagNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get tracking information based on the sender address and tag name
//
// @param request - GetTrackListByMailFromAndTagNameRequest
//
// @return GetTrackListByMailFromAndTagNameResponse
func (client *Client) GetTrackListByMailFromAndTagName(request *GetTrackListByMailFromAndTagNameRequest) (_result *GetTrackListByMailFromAndTagNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTrackListByMailFromAndTagNameResponse{}
	_body, _err := client.GetTrackListByMailFromAndTagNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Account Details
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Account Details
//
// @return GetUserResponse
func (client *Client) GetUser() (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取发信的黑名单列表
//
// @param request - ListBlockSendingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBlockSendingResponse
func (client *Client) ListBlockSendingWithOptions(request *ListBlockSendingRequest, runtime *dara.RuntimeOptions) (_result *ListBlockSendingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		query["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.BlockEmail) {
		query["BlockEmail"] = request.BlockEmail
	}

	if !dara.IsNil(request.BlockType) {
		query["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SenderEmail) {
		query["SenderEmail"] = request.SenderEmail
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBlockSending"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBlockSendingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取发信的黑名单列表
//
// @param request - ListBlockSendingRequest
//
// @return ListBlockSendingResponse
func (client *Client) ListBlockSending(request *ListBlockSendingRequest) (_result *ListBlockSendingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBlockSendingResponse{}
	_body, _err := client.ListBlockSendingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List User Invalid Addresses.
//
// @param request - ListUserSuppressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserSuppressionResponse
func (client *Client) ListUserSuppressionWithOptions(request *ListUserSuppressionRequest, runtime *dara.RuntimeOptions) (_result *ListUserSuppressionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Address) {
		query["Address"] = request.Address
	}

	if !dara.IsNil(request.EndBounceTime) {
		query["EndBounceTime"] = request.EndBounceTime
	}

	if !dara.IsNil(request.EndCreateTime) {
		query["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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

	if !dara.IsNil(request.StartBounceTime) {
		query["StartBounceTime"] = request.StartBounceTime
	}

	if !dara.IsNil(request.StartCreateTime) {
		query["StartCreateTime"] = request.StartCreateTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserSuppression"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserSuppressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List User Invalid Addresses.
//
// @param request - ListUserSuppressionRequest
//
// @return ListUserSuppressionResponse
func (client *Client) ListUserSuppression(request *ListUserSuppressionRequest) (_result *ListUserSuppressionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserSuppressionResponse{}
	_body, _err := client.ListUserSuppressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify the sending address
//
// @param request - ModifyMailAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMailAddressResponse
func (client *Client) ModifyMailAddressWithOptions(request *ModifyMailAddressRequest, runtime *dara.RuntimeOptions) (_result *ModifyMailAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MailAddressId) {
		query["MailAddressId"] = request.MailAddressId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ReplyAddress) {
		query["ReplyAddress"] = request.ReplyAddress
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
		Action:      dara.String("ModifyMailAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyMailAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify the sending address
//
// @param request - ModifyMailAddressRequest
//
// @return ModifyMailAddressResponse
func (client *Client) ModifyMailAddress(request *ModifyMailAddressRequest) (_result *ModifyMailAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyMailAddressResponse{}
	_body, _err := client.ModifyMailAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify the domain-level password
//
// @param request - ModifyPWByDomainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyPWByDomainResponse
func (client *Client) ModifyPWByDomainWithOptions(request *ModifyPWByDomainRequest, runtime *dara.RuntimeOptions) (_result *ModifyPWByDomainResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["DomainName"] = request.DomainName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
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
		Action:      dara.String("ModifyPWByDomain"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyPWByDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify the domain-level password
//
// @param request - ModifyPWByDomainRequest
//
// @return ModifyPWByDomainResponse
func (client *Client) ModifyPWByDomain(request *ModifyPWByDomainRequest) (_result *ModifyPWByDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyPWByDomainResponse{}
	_body, _err := client.ModifyPWByDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Modify Tag
//
// @param request - ModifyTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyTagResponse
func (client *Client) ModifyTagWithOptions(request *ModifyTagRequest, runtime *dara.RuntimeOptions) (_result *ModifyTagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TagDescription) {
		query["TagDescription"] = request.TagDescription
	}

	if !dara.IsNil(request.TagId) {
		query["TagId"] = request.TagId
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyTag"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Modify Tag
//
// @param request - ModifyTagRequest
//
// @return ModifyTagResponse
func (client *Client) ModifyTag(request *ModifyTagRequest) (_result *ModifyTagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyTagResponse{}
	_body, _err := client.ModifyTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query domain information
//
// @param request - QueryDomainByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDomainByParamResponse
func (client *Client) QueryDomainByParamWithOptions(request *QueryDomainByParamRequest, runtime *dara.RuntimeOptions) (_result *QueryDomainByParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDomainByParam"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDomainByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query domain information
//
// @param request - QueryDomainByParamRequest
//
// @return QueryDomainByParamResponse
func (client *Client) QueryDomainByParam(request *QueryDomainByParamRequest) (_result *QueryDomainByParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDomainByParamResponse{}
	_body, _err := client.QueryDomainByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # NextStart changed to string
//
// Description:
//
// Retrieve deduplicated invalid address information. If an email is sent to the same invalid address multiple times, only the first occurrence will be recorded. The query should be based on the time when the address was first classified as invalid.
//
// @param request - QueryInvalidAddressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInvalidAddressResponse
func (client *Client) QueryInvalidAddressWithOptions(request *QueryInvalidAddressRequest, runtime *dara.RuntimeOptions) (_result *QueryInvalidAddressResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.NextStart) {
		query["NextStart"] = request.NextStart
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInvalidAddress"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInvalidAddressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # NextStart changed to string
//
// Description:
//
// Retrieve deduplicated invalid address information. If an email is sent to the same invalid address multiple times, only the first occurrence will be recorded. The query should be based on the time when the address was first classified as invalid.
//
// @param request - QueryInvalidAddressRequest
//
// @return QueryInvalidAddressResponse
func (client *Client) QueryInvalidAddress(request *QueryInvalidAddressRequest) (_result *QueryInvalidAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInvalidAddressResponse{}
	_body, _err := client.QueryInvalidAddressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of sending addresses.
//
// @param request - QueryMailAddressByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMailAddressByParamResponse
func (client *Client) QueryMailAddressByParamWithOptions(request *QueryMailAddressByParamRequest, runtime *dara.RuntimeOptions) (_result *QueryMailAddressByParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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

	if !dara.IsNil(request.Sendtype) {
		query["Sendtype"] = request.Sendtype
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMailAddressByParam"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMailAddressByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of sending addresses.
//
// @param request - QueryMailAddressByParamRequest
//
// @return QueryMailAddressByParamResponse
func (client *Client) QueryMailAddressByParam(request *QueryMailAddressByParamRequest) (_result *QueryMailAddressByParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryMailAddressByParamResponse{}
	_body, _err := client.QueryMailAddressByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the details of the recipient list
//
// @param request - QueryReceiverByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReceiverByParamResponse
func (client *Client) QueryReceiverByParamWithOptions(request *QueryReceiverByParamRequest, runtime *dara.RuntimeOptions) (_result *QueryReceiverByParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryReceiverByParam"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReceiverByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the details of the recipient list
//
// @param request - QueryReceiverByParamRequest
//
// @return QueryReceiverByParamResponse
func (client *Client) QueryReceiverByParam(request *QueryReceiverByParamRequest) (_result *QueryReceiverByParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryReceiverByParamResponse{}
	_body, _err := client.QueryReceiverByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve detailed information about a recipient list
//
// @param request - QueryReceiverDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryReceiverDetailResponse
func (client *Client) QueryReceiverDetailWithOptions(request *QueryReceiverDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryReceiverDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.NextStart) {
		query["NextStart"] = request.NextStart
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
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
		Action:      dara.String("QueryReceiverDetail"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryReceiverDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve detailed information about a recipient list
//
// @param request - QueryReceiverDetailRequest
//
// @return QueryReceiverDetailResponse
func (client *Client) QueryReceiverDetail(request *QueryReceiverDetailRequest) (_result *QueryReceiverDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryReceiverDetailResponse{}
	_body, _err := client.QueryReceiverDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call QueryTagByParam to retrieve tags.
//
// @param request - QueryTagByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTagByParamResponse
func (client *Client) QueryTagByParamWithOptions(request *QueryTagByParamRequest, runtime *dara.RuntimeOptions) (_result *QueryTagByParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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
		Action:      dara.String("QueryTagByParam"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTagByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call QueryTagByParam to retrieve tags.
//
// @param request - QueryTagByParamRequest
//
// @return QueryTagByParamResponse
func (client *Client) QueryTagByParam(request *QueryTagByParamRequest) (_result *QueryTagByParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTagByParamResponse{}
	_body, _err := client.QueryTagByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query task list
//
// @param request - QueryTaskByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTaskByParamResponse
func (client *Client) QueryTaskByParamWithOptions(request *QueryTaskByParamRequest, runtime *dara.RuntimeOptions) (_result *QueryTaskByParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KeyWord) {
		query["KeyWord"] = request.KeyWord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
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

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTaskByParam"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTaskByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query task list
//
// @param request - QueryTaskByParamRequest
//
// @return QueryTaskByParamResponse
func (client *Client) QueryTaskByParam(request *QueryTaskByParamRequest) (_result *QueryTaskByParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryTaskByParamResponse{}
	_body, _err := client.QueryTaskByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户无效地址
//
// @param request - RemoveUserSuppressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserSuppressionResponse
func (client *Client) RemoveUserSuppressionWithOptions(request *RemoveUserSuppressionRequest, runtime *dara.RuntimeOptions) (_result *RemoveUserSuppressionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SuppressionIds) {
		query["SuppressionIds"] = request.SuppressionIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveUserSuppression"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveUserSuppressionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户无效地址
//
// @param request - RemoveUserSuppressionRequest
//
// @return RemoveUserSuppressionResponse
func (client *Client) RemoveUserSuppression(request *RemoveUserSuppressionRequest) (_result *RemoveUserSuppressionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveUserSuppressionResponse{}
	_body, _err := client.RemoveUserSuppressionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a Single Recipient
//
// @param request - SaveReceiverDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveReceiverDetailResponse
func (client *Client) SaveReceiverDetailWithOptions(request *SaveReceiverDetailRequest, runtime *dara.RuntimeOptions) (_result *SaveReceiverDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Detail) {
		query["Detail"] = request.Detail
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReceiverId) {
		query["ReceiverId"] = request.ReceiverId
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
		Action:      dara.String("SaveReceiverDetail"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveReceiverDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a Single Recipient
//
// @param request - SaveReceiverDetailRequest
//
// @return SaveReceiverDetailResponse
func (client *Client) SaveReceiverDetail(request *SaveReceiverDetailRequest) (_result *SaveReceiverDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveReceiverDetailResponse{}
	_body, _err := client.SaveReceiverDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Send Template Test Email
//
// @param request - SendTestByTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTestByTemplateResponse
func (client *Client) SendTestByTemplateWithOptions(request *SendTestByTemplateRequest, runtime *dara.RuntimeOptions) (_result *SendTestByTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.Birthday) {
		query["Birthday"] = request.Birthday
	}

	if !dara.IsNil(request.Email) {
		query["Email"] = request.Email
	}

	if !dara.IsNil(request.Gender) {
		query["Gender"] = request.Gender
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.NickName) {
		query["NickName"] = request.NickName
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendTestByTemplate"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendTestByTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Send Template Test Email
//
// @param request - SendTestByTemplateRequest
//
// @return SendTestByTemplateResponse
func (client *Client) SendTestByTemplate(request *SendTestByTemplateRequest) (_result *SendTestByTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendTestByTemplateResponse{}
	_body, _err := client.SendTestByTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve Sending Data under Specified Conditions
//
// @param request - SenderStatisticsByTagNameAndBatchIDRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SenderStatisticsByTagNameAndBatchIDResponse
func (client *Client) SenderStatisticsByTagNameAndBatchIDWithOptions(request *SenderStatisticsByTagNameAndBatchIDRequest, runtime *dara.RuntimeOptions) (_result *SenderStatisticsByTagNameAndBatchIDResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.DedicatedIp) {
		query["DedicatedIp"] = request.DedicatedIp
	}

	if !dara.IsNil(request.DedicatedIpPoolId) {
		query["DedicatedIpPoolId"] = request.DedicatedIpPoolId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Esp) {
		query["Esp"] = request.Esp
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SenderStatisticsByTagNameAndBatchID"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SenderStatisticsByTagNameAndBatchIDResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve Sending Data under Specified Conditions
//
// @param request - SenderStatisticsByTagNameAndBatchIDRequest
//
// @return SenderStatisticsByTagNameAndBatchIDResponse
func (client *Client) SenderStatisticsByTagNameAndBatchID(request *SenderStatisticsByTagNameAndBatchIDRequest) (_result *SenderStatisticsByTagNameAndBatchIDResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SenderStatisticsByTagNameAndBatchIDResponse{}
	_body, _err := client.SenderStatisticsByTagNameAndBatchIDWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Delivery Result Details
//
// @param request - SenderStatisticsDetailByParamRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SenderStatisticsDetailByParamResponse
func (client *Client) SenderStatisticsDetailByParamWithOptions(request *SenderStatisticsDetailByParamRequest, runtime *dara.RuntimeOptions) (_result *SenderStatisticsDetailByParamResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Length) {
		query["Length"] = request.Length
	}

	if !dara.IsNil(request.NextStart) {
		query["NextStart"] = request.NextStart
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TagName) {
		query["TagName"] = request.TagName
	}

	if !dara.IsNil(request.ToAddress) {
		query["ToAddress"] = request.ToAddress
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SenderStatisticsDetailByParam"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SenderStatisticsDetailByParamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Delivery Result Details
//
// @param request - SenderStatisticsDetailByParamRequest
//
// @return SenderStatisticsDetailByParamResponse
func (client *Client) SenderStatisticsDetailByParam(request *SenderStatisticsDetailByParamRequest) (_result *SenderStatisticsDetailByParamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SenderStatisticsDetailByParamResponse{}
	_body, _err := client.SenderStatisticsDetailByParamWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置用户无效地址级别配置
//
// @param request - SetSuppressionListLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSuppressionListLevelResponse
func (client *Client) SetSuppressionListLevelWithOptions(request *SetSuppressionListLevelRequest, runtime *dara.RuntimeOptions) (_result *SetSuppressionListLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SuppressionListLevel) {
		query["SuppressionListLevel"] = request.SuppressionListLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSuppressionListLevel"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSuppressionListLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置用户无效地址级别配置
//
// @param request - SetSuppressionListLevelRequest
//
// @return SetSuppressionListLevelResponse
func (client *Client) SetSuppressionListLevel(request *SetSuppressionListLevelRequest) (_result *SetSuppressionListLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSuppressionListLevelResponse{}
	_body, _err := client.SetSuppressionListLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # API for Sending Emails
//
// @param request - SingleSendMailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SingleSendMailResponse
func (client *Client) SingleSendMailWithOptions(request *SingleSendMailRequest, runtime *dara.RuntimeOptions) (_result *SingleSendMailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AddressType) {
		body["AddressType"] = request.AddressType
	}

	if !dara.IsNil(request.Attachments) {
		body["Attachments"] = request.Attachments
	}

	if !dara.IsNil(request.ClickTrace) {
		body["ClickTrace"] = request.ClickTrace
	}

	if !dara.IsNil(request.FromAlias) {
		body["FromAlias"] = request.FromAlias
	}

	if !dara.IsNil(request.Headers) {
		body["Headers"] = request.Headers
	}

	if !dara.IsNil(request.HtmlBody) {
		body["HtmlBody"] = request.HtmlBody
	}

	if !dara.IsNil(request.IpPoolId) {
		body["IpPoolId"] = request.IpPoolId
	}

	if !dara.IsNil(request.ReplyAddress) {
		body["ReplyAddress"] = request.ReplyAddress
	}

	if !dara.IsNil(request.ReplyAddressAlias) {
		body["ReplyAddressAlias"] = request.ReplyAddressAlias
	}

	if !dara.IsNil(request.ReplyToAddress) {
		body["ReplyToAddress"] = request.ReplyToAddress
	}

	if !dara.IsNil(request.Subject) {
		body["Subject"] = request.Subject
	}

	if !dara.IsNil(request.TagName) {
		body["TagName"] = request.TagName
	}

	if !dara.IsNil(request.TextBody) {
		body["TextBody"] = request.TextBody
	}

	if !dara.IsNil(request.ToAddress) {
		body["ToAddress"] = request.ToAddress
	}

	if !dara.IsNil(request.UnSubscribeFilterLevel) {
		body["UnSubscribeFilterLevel"] = request.UnSubscribeFilterLevel
	}

	if !dara.IsNil(request.UnSubscribeLinkType) {
		body["UnSubscribeLinkType"] = request.UnSubscribeLinkType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SingleSendMail"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SingleSendMailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # API for Sending Emails
//
// @param request - SingleSendMailRequest
//
// @return SingleSendMailResponse
func (client *Client) SingleSendMail(request *SingleSendMailRequest) (_result *SingleSendMailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SingleSendMailResponse{}
	_body, _err := client.SingleSendMailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SingleSendMailAdvance(request *SingleSendMailAdvanceRequest, runtime *dara.RuntimeOptions) (_result *SingleSendMailResponse, _err error) {
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
		"Product":  dara.String("Dm"),
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
	singleSendMailReq := &SingleSendMailRequest{}
	openapiutil.Convert(request, singleSendMailReq)
	if !dara.IsNil(request.Attachments) {
		i0 := 0
		for _, item0 := range request.Attachments {
			if !dara.IsNil(item0.AttachmentUrlObject) {
				authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
				if _err != nil {
					return _result, _err
				}

				tmpBody = dara.ToMap(authResponse["body"])
				useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
				authResponseBody = openapiutil.StringifyMapValue(tmpBody)
				fileObj = &dara.FileField{
					Filename:    authResponseBody["ObjectKey"],
					Content:     item0.AttachmentUrlObject,
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
				tmpObj := singleSendMailReq.Attachments[i0]
				tmpObj.AttachmentUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
				i0++
			}

		}
	}

	singleSendMailResp, _err := client.SingleSendMailWithOptions(singleSendMailReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = singleSendMailResp
	return _result, _err
}

// Summary:
//
// Lift sending restrictions due to unsubscription, reporting, etc.
//
// @param request - UnblockSendingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnblockSendingResponse
func (client *Client) UnblockSendingWithOptions(request *UnblockSendingRequest, runtime *dara.RuntimeOptions) (_result *UnblockSendingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BlockEmail) {
		query["BlockEmail"] = request.BlockEmail
	}

	if !dara.IsNil(request.BlockType) {
		query["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.SenderEmail) {
		query["SenderEmail"] = request.SenderEmail
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnblockSending"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnblockSendingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lift sending restrictions due to unsubscription, reporting, etc.
//
// @param request - UnblockSendingRequest
//
// @return UnblockSendingResponse
func (client *Client) UnblockSending(request *UnblockSendingRequest) (_result *UnblockSendingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnblockSendingResponse{}
	_body, _err := client.UnblockSendingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update IP Protection API
//
// @param request - UpdateIpProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIpProtectionResponse
func (client *Client) UpdateIpProtectionWithOptions(request *UpdateIpProtectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateIpProtectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IpProtection) {
		query["IpProtection"] = request.IpProtection
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
		Action:      dara.String("UpdateIpProtection"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIpProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update IP Protection API
//
// @param request - UpdateIpProtectionRequest
//
// @return UpdateIpProtectionResponse
func (client *Client) UpdateIpProtection(request *UpdateIpProtectionRequest) (_result *UpdateIpProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIpProtectionResponse{}
	_body, _err := client.UpdateIpProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update account information
//
// @param tmpReq - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(tmpReq *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateUserShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.User) {
		request.UserShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.User, dara.String("User"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.UserShrink) {
		body["User"] = request.UserShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2015-11-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update account information
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
