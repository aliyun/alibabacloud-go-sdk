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

type DescribeDomainReportRequest struct {
	Domain      *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Field       *string `json:"Field,omitempty" xml:"Field,omitempty"`
	ServiceLang *string `json:"ServiceLang,omitempty" xml:"ServiceLang,omitempty"`
}

func (s DescribeDomainReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainReportRequest) SetDomain(v string) *DescribeDomainReportRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainReportRequest) SetField(v string) *DescribeDomainReportRequest {
	s.Field = &v
	return s
}

func (s *DescribeDomainReportRequest) SetServiceLang(v string) *DescribeDomainReportRequest {
	s.ServiceLang = &v
	return s
}

type DescribeDomainReportResponseBody struct {
	AttackCntByThreatType *string `json:"AttackCntByThreatType,omitempty" xml:"AttackCntByThreatType,omitempty"`
	AttackPreferenceTop5  *string `json:"AttackPreferenceTop5,omitempty" xml:"AttackPreferenceTop5,omitempty"`
	Basic                 *string `json:"Basic,omitempty" xml:"Basic,omitempty"`
	Confidence            *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Context               *string `json:"Context,omitempty" xml:"Context,omitempty"`
	Domain                *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Group                 *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Intelligences         *string `json:"Intelligences,omitempty" xml:"Intelligences,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenario              *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	SslCert               *string `json:"SslCert,omitempty" xml:"SslCert,omitempty"`
	ThreatLevel           *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatTypes           *string `json:"ThreatTypes,omitempty" xml:"ThreatTypes,omitempty"`
	Whois                 *string `json:"Whois,omitempty" xml:"Whois,omitempty"`
}

func (s DescribeDomainReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainReportResponseBody) SetAttackCntByThreatType(v string) *DescribeDomainReportResponseBody {
	s.AttackCntByThreatType = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetAttackPreferenceTop5(v string) *DescribeDomainReportResponseBody {
	s.AttackPreferenceTop5 = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetBasic(v string) *DescribeDomainReportResponseBody {
	s.Basic = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetConfidence(v string) *DescribeDomainReportResponseBody {
	s.Confidence = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetContext(v string) *DescribeDomainReportResponseBody {
	s.Context = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetDomain(v string) *DescribeDomainReportResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetGroup(v string) *DescribeDomainReportResponseBody {
	s.Group = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetIntelligences(v string) *DescribeDomainReportResponseBody {
	s.Intelligences = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetRequestId(v string) *DescribeDomainReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetScenario(v string) *DescribeDomainReportResponseBody {
	s.Scenario = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetSslCert(v string) *DescribeDomainReportResponseBody {
	s.SslCert = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetThreatLevel(v string) *DescribeDomainReportResponseBody {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetThreatTypes(v string) *DescribeDomainReportResponseBody {
	s.ThreatTypes = &v
	return s
}

func (s *DescribeDomainReportResponseBody) SetWhois(v string) *DescribeDomainReportResponseBody {
	s.Whois = &v
	return s
}

type DescribeDomainReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainReportResponse) SetHeaders(v map[string]*string) *DescribeDomainReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainReportResponse) SetStatusCode(v int32) *DescribeDomainReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainReportResponse) SetBody(v *DescribeDomainReportResponseBody) *DescribeDomainReportResponse {
	s.Body = v
	return s
}

type DescribeFileReportRequest struct {
	Field       *string `json:"Field,omitempty" xml:"Field,omitempty"`
	FileHash    *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	ServiceLang *string `json:"ServiceLang,omitempty" xml:"ServiceLang,omitempty"`
}

func (s DescribeFileReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileReportRequest) SetField(v string) *DescribeFileReportRequest {
	s.Field = &v
	return s
}

func (s *DescribeFileReportRequest) SetFileHash(v string) *DescribeFileReportRequest {
	s.FileHash = &v
	return s
}

func (s *DescribeFileReportRequest) SetServiceLang(v string) *DescribeFileReportRequest {
	s.ServiceLang = &v
	return s
}

type DescribeFileReportResponseBody struct {
	Basic         *string `json:"Basic,omitempty" xml:"Basic,omitempty"`
	FileHash      *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	Intelligences *string `json:"Intelligences,omitempty" xml:"Intelligences,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sandbox       *string `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	ThreatLevel   *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatTypes   *string `json:"ThreatTypes,omitempty" xml:"ThreatTypes,omitempty"`
}

func (s DescribeFileReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFileReportResponseBody) SetBasic(v string) *DescribeFileReportResponseBody {
	s.Basic = &v
	return s
}

func (s *DescribeFileReportResponseBody) SetFileHash(v string) *DescribeFileReportResponseBody {
	s.FileHash = &v
	return s
}

func (s *DescribeFileReportResponseBody) SetIntelligences(v string) *DescribeFileReportResponseBody {
	s.Intelligences = &v
	return s
}

func (s *DescribeFileReportResponseBody) SetRequestId(v string) *DescribeFileReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFileReportResponseBody) SetSandbox(v string) *DescribeFileReportResponseBody {
	s.Sandbox = &v
	return s
}

func (s *DescribeFileReportResponseBody) SetThreatLevel(v string) *DescribeFileReportResponseBody {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeFileReportResponseBody) SetThreatTypes(v string) *DescribeFileReportResponseBody {
	s.ThreatTypes = &v
	return s
}

type DescribeFileReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFileReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFileReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFileReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeFileReportResponse) SetHeaders(v map[string]*string) *DescribeFileReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeFileReportResponse) SetStatusCode(v int32) *DescribeFileReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFileReportResponse) SetBody(v *DescribeFileReportResponseBody) *DescribeFileReportResponse {
	s.Body = v
	return s
}

type DescribeIpReportRequest struct {
	Field       *string `json:"Field,omitempty" xml:"Field,omitempty"`
	Ip          *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	ServiceLang *string `json:"ServiceLang,omitempty" xml:"ServiceLang,omitempty"`
}

func (s DescribeIpReportRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpReportRequest) SetField(v string) *DescribeIpReportRequest {
	s.Field = &v
	return s
}

func (s *DescribeIpReportRequest) SetIp(v string) *DescribeIpReportRequest {
	s.Ip = &v
	return s
}

func (s *DescribeIpReportRequest) SetServiceLang(v string) *DescribeIpReportRequest {
	s.ServiceLang = &v
	return s
}

type DescribeIpReportResponseBody struct {
	AttackCntByThreatType *string `json:"AttackCntByThreatType,omitempty" xml:"AttackCntByThreatType,omitempty"`
	AttackPreferenceTop5  *string `json:"AttackPreferenceTop5,omitempty" xml:"AttackPreferenceTop5,omitempty"`
	Confidence            *string `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	Context               *string `json:"Context,omitempty" xml:"Context,omitempty"`
	Group                 *string `json:"Group,omitempty" xml:"Group,omitempty"`
	Intelligences         *string `json:"Intelligences,omitempty" xml:"Intelligences,omitempty"`
	Ip                    *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scenario              *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	ThreatLevel           *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
	ThreatTypes           *string `json:"ThreatTypes,omitempty" xml:"ThreatTypes,omitempty"`
	Whois                 *string `json:"Whois,omitempty" xml:"Whois,omitempty"`
}

func (s DescribeIpReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpReportResponseBody) SetAttackCntByThreatType(v string) *DescribeIpReportResponseBody {
	s.AttackCntByThreatType = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetAttackPreferenceTop5(v string) *DescribeIpReportResponseBody {
	s.AttackPreferenceTop5 = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetConfidence(v string) *DescribeIpReportResponseBody {
	s.Confidence = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetContext(v string) *DescribeIpReportResponseBody {
	s.Context = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetGroup(v string) *DescribeIpReportResponseBody {
	s.Group = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetIntelligences(v string) *DescribeIpReportResponseBody {
	s.Intelligences = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetIp(v string) *DescribeIpReportResponseBody {
	s.Ip = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetRequestId(v string) *DescribeIpReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetScenario(v string) *DescribeIpReportResponseBody {
	s.Scenario = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetThreatLevel(v string) *DescribeIpReportResponseBody {
	s.ThreatLevel = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetThreatTypes(v string) *DescribeIpReportResponseBody {
	s.ThreatTypes = &v
	return s
}

func (s *DescribeIpReportResponseBody) SetWhois(v string) *DescribeIpReportResponseBody {
	s.Whois = &v
	return s
}

type DescribeIpReportResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIpReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIpReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpReportResponse) SetHeaders(v map[string]*string) *DescribeIpReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpReportResponse) SetStatusCode(v int32) *DescribeIpReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpReportResponse) SetBody(v *DescribeIpReportResponseBody) *DescribeIpReportResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("sasti"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) DescribeDomainReportWithOptions(request *DescribeDomainReportRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["Domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.Field)) {
		query["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLang)) {
		query["ServiceLang"] = request.ServiceLang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainReport"),
		Version:     tea.String("2020-05-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainReport(request *DescribeDomainReportRequest) (_result *DescribeDomainReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainReportResponse{}
	_body, _err := client.DescribeDomainReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFileReportWithOptions(request *DescribeFileReportRequest, runtime *util.RuntimeOptions) (_result *DescribeFileReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Field)) {
		query["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.FileHash)) {
		query["FileHash"] = request.FileHash
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLang)) {
		query["ServiceLang"] = request.ServiceLang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFileReport"),
		Version:     tea.String("2020-05-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFileReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFileReport(request *DescribeFileReportRequest) (_result *DescribeFileReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFileReportResponse{}
	_body, _err := client.DescribeFileReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpReportWithOptions(request *DescribeIpReportRequest, runtime *util.RuntimeOptions) (_result *DescribeIpReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Field)) {
		query["Field"] = request.Field
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLang)) {
		query["ServiceLang"] = request.ServiceLang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpReport"),
		Version:     tea.String("2020-05-12"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpReport(request *DescribeIpReportRequest) (_result *DescribeIpReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpReportResponse{}
	_body, _err := client.DescribeIpReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
