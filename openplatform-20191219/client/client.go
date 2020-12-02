// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AuthorizeFileUploadRequest struct {
	Product  *string `json:"Product,omitempty" xml:"Product,omitempty" require:"true"`
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

type AuthorizeFileUploadResponse struct {
	AccessKeyId   *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty" require:"true"`
	Bucket        *string `json:"Bucket,omitempty" xml:"Bucket,omitempty" require:"true"`
	EncodedPolicy *string `json:"EncodedPolicy,omitempty" xml:"EncodedPolicy,omitempty" require:"true"`
	Endpoint      *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" require:"true"`
	ObjectKey     *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Signature     *string `json:"Signature,omitempty" xml:"Signature,omitempty" require:"true"`
	UseAccelerate *bool   `json:"UseAccelerate,omitempty" xml:"UseAccelerate,omitempty" require:"true"`
}

func (s AuthorizeFileUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeFileUploadResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeFileUploadResponse) SetAccessKeyId(v string) *AuthorizeFileUploadResponse {
	s.AccessKeyId = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetBucket(v string) *AuthorizeFileUploadResponse {
	s.Bucket = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetEncodedPolicy(v string) *AuthorizeFileUploadResponse {
	s.EncodedPolicy = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetEndpoint(v string) *AuthorizeFileUploadResponse {
	s.Endpoint = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetObjectKey(v string) *AuthorizeFileUploadResponse {
	s.ObjectKey = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetRequestId(v string) *AuthorizeFileUploadResponse {
	s.RequestId = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetSignature(v string) *AuthorizeFileUploadResponse {
	s.Signature = &v
	return s
}

func (s *AuthorizeFileUploadResponse) SetUseAccelerate(v bool) *AuthorizeFileUploadResponse {
	s.UseAccelerate = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
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

func (client *Client) AuthorizeFileUploadWithOptions(request *AuthorizeFileUploadRequest, runtime *util.RuntimeOptions) (_result *AuthorizeFileUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthorizeFileUploadResponse{}
	_body, _err := client.DoRequest(tea.String("AuthorizeFileUpload"), tea.String("HTTPS"), tea.String("GET"), tea.String("2019-12-19"), tea.String("AK"), tea.ToMap(request), nil, runtime)
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
