// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AuthorizeFileUploadRequest struct {
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AuthorizeFileUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeFileUploadRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadRequest) SetProduct(v string) *AuthorizeFileUploadRequest {
	s.Product = &v
	return s
}

func (s *AuthorizeFileUploadRequest) SetRegionId(v string) *AuthorizeFileUploadRequest {
	s.RegionId = &v
	return s
}

type AuthorizeFileUploadResponseBody struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	Bucket        *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty"`
	Endpoint      *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	ObjectKey     *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	UseAccelerate *bool   `json:"UseAccelerate,omitempty" xml:"UseAccelerate,omitempty"`
}

func (s AuthorizeFileUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadResponseBody) SetAccessKeyId(v string) *AuthorizeFileUploadResponseBody {
	s.AccessKeyId = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetBucket(v string) *AuthorizeFileUploadResponseBody {
	s.Bucket = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetEncodedPolicy(v string) *AuthorizeFileUploadResponseBody {
	s.EncodedPolicy = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetEndpoint(v string) *AuthorizeFileUploadResponseBody {
	s.Endpoint = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetObjectKey(v string) *AuthorizeFileUploadResponseBody {
	s.ObjectKey = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetRequestId(v string) *AuthorizeFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetSignature(v string) *AuthorizeFileUploadResponseBody {
	s.Signature = &v
	return s
}

func (s *AuthorizeFileUploadResponseBody) SetUseAccelerate(v bool) *AuthorizeFileUploadResponseBody {
	s.UseAccelerate = &v
	return s
}

type AuthorizeFileUploadResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AuthorizeFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AuthorizeFileUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeFileUploadResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadResponse) SetHeaders(v map[string]*string) *AuthorizeFileUploadResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeFileUploadResponse) SetStatusCode(v int32) *AuthorizeFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetBody(v *AuthorizeFileUploadResponseBody) *AuthorizeFileUploadResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("openplatform"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthorizeFileUploadWithOptions(request *AuthorizeFileUploadRequest, runtime *util.RuntimeOptions) (_result *AuthorizeFileUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AuthorizeFileUpload"),
		Version:     tea.String("2019-12-19"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeFileUploadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthorizeFileUpload(request *AuthorizeFileUploadRequest) (_result *AuthorizeFileUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthorizeFileUploadResponse{}
	_body, _err := client.AuthorizeFileUploadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
