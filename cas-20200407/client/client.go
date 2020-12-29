// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateCertificateRequestRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s CreateCertificateRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestRequest) SetSourceIp(v string) *CreateCertificateRequestRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetProductCode(v string) *CreateCertificateRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetUsername(v string) *CreateCertificateRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetPhone(v string) *CreateCertificateRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetEmail(v string) *CreateCertificateRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetDomain(v string) *CreateCertificateRequestRequest {
	s.Domain = &v
	return s
}

func (s *CreateCertificateRequestRequest) SetValidateType(v string) *CreateCertificateRequestRequest {
	s.ValidateType = &v
	return s
}

type CreateCertificateRequestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateCertificateRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestResponseBody) SetRequestId(v string) *CreateCertificateRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertificateRequestResponseBody) SetOrderId(v int64) *CreateCertificateRequestResponseBody {
	s.OrderId = &v
	return s
}

type CreateCertificateRequestResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertificateRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateRequestResponse) SetBody(v *CreateCertificateRequestResponseBody) *CreateCertificateRequestResponse {
	s.Body = v
	return s
}

type CreateCertificateWithCsrRequestRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Csr          *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	ProductCode  *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
	Phone        *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Email        *string `json:"Email,omitempty" xml:"Email,omitempty"`
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
}

func (s CreateCertificateWithCsrRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithCsrRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestRequest) SetSourceIp(v string) *CreateCertificateWithCsrRequestRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetCsr(v string) *CreateCertificateWithCsrRequestRequest {
	s.Csr = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetProductCode(v string) *CreateCertificateWithCsrRequestRequest {
	s.ProductCode = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetUsername(v string) *CreateCertificateWithCsrRequestRequest {
	s.Username = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetPhone(v string) *CreateCertificateWithCsrRequestRequest {
	s.Phone = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetEmail(v string) *CreateCertificateWithCsrRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCertificateWithCsrRequestRequest) SetValidateType(v string) *CreateCertificateWithCsrRequestRequest {
	s.ValidateType = &v
	return s
}

type CreateCertificateWithCsrRequestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OrderId   *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s CreateCertificateWithCsrRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithCsrRequestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestResponseBody) SetRequestId(v string) *CreateCertificateWithCsrRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCertificateWithCsrRequestResponseBody) SetOrderId(v int64) *CreateCertificateWithCsrRequestResponseBody {
	s.OrderId = &v
	return s
}

type CreateCertificateWithCsrRequestResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCertificateWithCsrRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCertificateWithCsrRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCertificateWithCsrRequestResponse) GoString() string {
	return s.String()
}

func (s *CreateCertificateWithCsrRequestResponse) SetHeaders(v map[string]*string) *CreateCertificateWithCsrRequestResponse {
	s.Headers = v
	return s
}

func (s *CreateCertificateWithCsrRequestResponse) SetBody(v *CreateCertificateWithCsrRequestResponseBody) *CreateCertificateWithCsrRequestResponse {
	s.Body = v
	return s
}

type DeleteCertificateRequestRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrderId  *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DeleteCertificateRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequestRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestRequest) SetSourceIp(v string) *DeleteCertificateRequestRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteCertificateRequestRequest) SetOrderId(v int64) *DeleteCertificateRequestRequest {
	s.OrderId = &v
	return s
}

type DeleteCertificateRequestResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCertificateRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequestResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestResponseBody) SetRequestId(v string) *DeleteCertificateRequestResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCertificateRequestResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCertificateRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCertificateRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCertificateRequestResponse) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequestResponse) SetHeaders(v map[string]*string) *DeleteCertificateRequestResponse {
	s.Headers = v
	return s
}

func (s *DeleteCertificateRequestResponse) SetBody(v *DeleteCertificateRequestResponseBody) *DeleteCertificateRequestResponse {
	s.Body = v
	return s
}

type DescribeCertificateStateRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	OrderId  *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s DescribeCertificateStateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateStateRequest) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateRequest) SetSourceIp(v string) *DescribeCertificateStateRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCertificateStateRequest) SetOrderId(v int64) *DescribeCertificateStateRequest {
	s.OrderId = &v
	return s
}

type DescribeCertificateStateResponseBody struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	PrivateKey   *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RecordType   *string `json:"RecordType,omitempty" xml:"RecordType,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	RecordDomain *string `json:"RecordDomain,omitempty" xml:"RecordDomain,omitempty"`
	RecordValue  *string `json:"RecordValue,omitempty" xml:"RecordValue,omitempty"`
	Domain       *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ValidateType *string `json:"ValidateType,omitempty" xml:"ValidateType,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	Certificate  *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
}

func (s DescribeCertificateStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateResponseBody) SetType(v string) *DescribeCertificateStateResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetPrivateKey(v string) *DescribeCertificateStateResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordType(v string) *DescribeCertificateStateResponseBody {
	s.RecordType = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRequestId(v string) *DescribeCertificateStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetContent(v string) *DescribeCertificateStateResponseBody {
	s.Content = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordDomain(v string) *DescribeCertificateStateResponseBody {
	s.RecordDomain = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetRecordValue(v string) *DescribeCertificateStateResponseBody {
	s.RecordValue = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetDomain(v string) *DescribeCertificateStateResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetValidateType(v string) *DescribeCertificateStateResponseBody {
	s.ValidateType = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetUri(v string) *DescribeCertificateStateResponseBody {
	s.Uri = &v
	return s
}

func (s *DescribeCertificateStateResponseBody) SetCertificate(v string) *DescribeCertificateStateResponseBody {
	s.Certificate = &v
	return s
}

type DescribeCertificateStateResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeCertificateStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCertificateStateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCertificateStateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCertificateStateResponse) SetHeaders(v map[string]*string) *DescribeCertificateStateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCertificateStateResponse) SetBody(v *DescribeCertificateStateResponseBody) *DescribeCertificateStateResponse {
	s.Body = v
	return s
}

type DescribePackageStateRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s DescribePackageStateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageStateRequest) GoString() string {
	return s.String()
}

func (s *DescribePackageStateRequest) SetSourceIp(v string) *DescribePackageStateRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribePackageStateRequest) SetProductCode(v string) *DescribePackageStateRequest {
	s.ProductCode = &v
	return s
}

type DescribePackageStateResponseBody struct {
	TotalCount  *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UsedCount   *int64  `json:"UsedCount,omitempty" xml:"UsedCount,omitempty"`
	IssuedCount *int64  `json:"IssuedCount,omitempty" xml:"IssuedCount,omitempty"`
}

func (s DescribePackageStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageStateResponseBody) SetTotalCount(v int64) *DescribePackageStateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetRequestId(v string) *DescribePackageStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetProductCode(v string) *DescribePackageStateResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetUsedCount(v int64) *DescribePackageStateResponseBody {
	s.UsedCount = &v
	return s
}

func (s *DescribePackageStateResponseBody) SetIssuedCount(v int64) *DescribePackageStateResponseBody {
	s.IssuedCount = &v
	return s
}

type DescribePackageStateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePackageStateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePackageStateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePackageStateResponse) GoString() string {
	return s.String()
}

func (s *DescribePackageStateResponse) SetHeaders(v map[string]*string) *DescribePackageStateResponse {
	s.Headers = v
	return s
}

func (s *DescribePackageStateResponse) SetBody(v *DescribePackageStateResponseBody) *DescribePackageStateResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou":                 tea.String("cas.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("cas.aliyuncs.com"),
		"ap-southeast-1":              tea.String("cas.aliyuncs.com"),
		"ap-southeast-3":              tea.String("cas.aliyuncs.com"),
		"ap-southeast-5":              tea.String("cas.aliyuncs.com"),
		"cn-beijing":                  tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cas.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cas.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cas.aliyuncs.com"),
		"cn-chengdu":                  tea.String("cas.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cas.aliyuncs.com"),
		"cn-fujian":                   tea.String("cas.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cas.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cas.aliyuncs.com"),
		"cn-hongkong":                 tea.String("cas.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cas.aliyuncs.com"),
		"cn-huhehaote":                tea.String("cas.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("cas.aliyuncs.com"),
		"cn-qingdao":                  tea.String("cas.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai":                 tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cas.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cas.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cas.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cas.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cas.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cas.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cas.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("cas.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cas.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cas.aliyuncs.com"),
		"eu-west-1":                   tea.String("cas.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cas.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cas.aliyuncs.com"),
		"us-east-1":                   tea.String("cas.aliyuncs.com"),
		"us-west-1":                   tea.String("cas.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateCertificateRequestWithOptions(request *CreateCertificateRequestRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCertificateRequestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCertificateRequest"), tea.String("2020-04-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertificateRequest(request *CreateCertificateRequestRequest) (_result *CreateCertificateRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateRequestResponse{}
	_body, _err := client.CreateCertificateRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCertificateWithCsrRequestWithOptions(request *CreateCertificateWithCsrRequestRequest, runtime *util.RuntimeOptions) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateCertificateWithCsrRequestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCertificateWithCsrRequest"), tea.String("2020-04-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCertificateWithCsrRequest(request *CreateCertificateWithCsrRequestRequest) (_result *CreateCertificateWithCsrRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCertificateWithCsrRequestResponse{}
	_body, _err := client.CreateCertificateWithCsrRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCertificateRequestWithOptions(request *DeleteCertificateRequestRequest, runtime *util.RuntimeOptions) (_result *DeleteCertificateRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteCertificateRequestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCertificateRequest"), tea.String("2020-04-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCertificateRequest(request *DeleteCertificateRequestRequest) (_result *DeleteCertificateRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCertificateRequestResponse{}
	_body, _err := client.DeleteCertificateRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCertificateStateWithOptions(request *DescribeCertificateStateRequest, runtime *util.RuntimeOptions) (_result *DescribeCertificateStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeCertificateStateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCertificateState"), tea.String("2020-04-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCertificateState(request *DescribeCertificateStateRequest) (_result *DescribeCertificateStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCertificateStateResponse{}
	_body, _err := client.DescribeCertificateStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePackageStateWithOptions(request *DescribePackageStateRequest, runtime *util.RuntimeOptions) (_result *DescribePackageStateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePackageStateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePackageState"), tea.String("2020-04-07"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePackageState(request *DescribePackageStateRequest) (_result *DescribePackageStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePackageStateResponse{}
	_body, _err := client.DescribePackageStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
