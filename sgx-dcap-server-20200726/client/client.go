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

type GetQeIdentityRequest struct {
	AcsHost     *string `json:"AcsHost,omitempty" xml:"AcsHost,omitempty"`
	ClientVpcId *string `json:"ClientVpcId,omitempty" xml:"ClientVpcId,omitempty"`
}

func (s GetQeIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQeIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetQeIdentityRequest) SetAcsHost(v string) *GetQeIdentityRequest {
	s.AcsHost = &v
	return s
}

func (s *GetQeIdentityRequest) SetClientVpcId(v string) *GetQeIdentityRequest {
	s.ClientVpcId = &v
	return s
}

type GetQeIdentityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQeIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQeIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetQeIdentityResponse) SetHeaders(v map[string]*string) *GetQeIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetQeIdentityResponse) SetStatusCode(v int32) *GetQeIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQeIdentityResponse) SetBody(v string) *GetQeIdentityResponse {
	s.Body = &v
	return s
}

type GetQveIdentityRequest struct {
	AcsHost     *string `json:"AcsHost,omitempty" xml:"AcsHost,omitempty"`
	ClientVpcId *string `json:"ClientVpcId,omitempty" xml:"ClientVpcId,omitempty"`
}

func (s GetQveIdentityRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQveIdentityRequest) GoString() string {
	return s.String()
}

func (s *GetQveIdentityRequest) SetAcsHost(v string) *GetQveIdentityRequest {
	s.AcsHost = &v
	return s
}

func (s *GetQveIdentityRequest) SetClientVpcId(v string) *GetQveIdentityRequest {
	s.ClientVpcId = &v
	return s
}

type GetQveIdentityResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQveIdentityResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQveIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetQveIdentityResponse) SetHeaders(v map[string]*string) *GetQveIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetQveIdentityResponse) SetStatusCode(v int32) *GetQveIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQveIdentityResponse) SetBody(v string) *GetQveIdentityResponse {
	s.Body = &v
	return s
}

type GetTcbInfoRequest struct {
	AcsHost     *string `json:"AcsHost,omitempty" xml:"AcsHost,omitempty"`
	ClientVpcId *string `json:"ClientVpcId,omitempty" xml:"ClientVpcId,omitempty"`
	Fmspc       *string `json:"fmspc,omitempty" xml:"fmspc,omitempty"`
}

func (s GetTcbInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTcbInfoRequest) GoString() string {
	return s.String()
}

func (s *GetTcbInfoRequest) SetAcsHost(v string) *GetTcbInfoRequest {
	s.AcsHost = &v
	return s
}

func (s *GetTcbInfoRequest) SetClientVpcId(v string) *GetTcbInfoRequest {
	s.ClientVpcId = &v
	return s
}

func (s *GetTcbInfoRequest) SetFmspc(v string) *GetTcbInfoRequest {
	s.Fmspc = &v
	return s
}

type GetTcbInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTcbInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTcbInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTcbInfoResponse) SetHeaders(v map[string]*string) *GetTcbInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTcbInfoResponse) SetStatusCode(v int32) *GetTcbInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTcbInfoResponse) SetBody(v string) *GetTcbInfoResponse {
	s.Body = &v
	return s
}

type PckCrlRequest struct {
	AcsHost     *string `json:"AcsHost,omitempty" xml:"AcsHost,omitempty"`
	ClientVpcId *string `json:"ClientVpcId,omitempty" xml:"ClientVpcId,omitempty"`
	Ca          *string `json:"ca,omitempty" xml:"ca,omitempty"`
}

func (s PckCrlRequest) String() string {
	return tea.Prettify(s)
}

func (s PckCrlRequest) GoString() string {
	return s.String()
}

func (s *PckCrlRequest) SetAcsHost(v string) *PckCrlRequest {
	s.AcsHost = &v
	return s
}

func (s *PckCrlRequest) SetClientVpcId(v string) *PckCrlRequest {
	s.ClientVpcId = &v
	return s
}

func (s *PckCrlRequest) SetCa(v string) *PckCrlRequest {
	s.Ca = &v
	return s
}

type PckCrlResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PckCrlResponse) String() string {
	return tea.Prettify(s)
}

func (s PckCrlResponse) GoString() string {
	return s.String()
}

func (s *PckCrlResponse) SetHeaders(v map[string]*string) *PckCrlResponse {
	s.Headers = v
	return s
}

func (s *PckCrlResponse) SetStatusCode(v int32) *PckCrlResponse {
	s.StatusCode = &v
	return s
}

func (s *PckCrlResponse) SetBody(v string) *PckCrlResponse {
	s.Body = &v
	return s
}

type RootCaCrlRequest struct {
	AcsHost     *string `json:"AcsHost,omitempty" xml:"AcsHost,omitempty"`
	ClientVpcId *string `json:"ClientVpcId,omitempty" xml:"ClientVpcId,omitempty"`
}

func (s RootCaCrlRequest) String() string {
	return tea.Prettify(s)
}

func (s RootCaCrlRequest) GoString() string {
	return s.String()
}

func (s *RootCaCrlRequest) SetAcsHost(v string) *RootCaCrlRequest {
	s.AcsHost = &v
	return s
}

func (s *RootCaCrlRequest) SetClientVpcId(v string) *RootCaCrlRequest {
	s.ClientVpcId = &v
	return s
}

type RootCaCrlResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RootCaCrlResponse) String() string {
	return tea.Prettify(s)
}

func (s RootCaCrlResponse) GoString() string {
	return s.String()
}

func (s *RootCaCrlResponse) SetHeaders(v map[string]*string) *RootCaCrlResponse {
	s.Headers = v
	return s
}

func (s *RootCaCrlResponse) SetStatusCode(v int32) *RootCaCrlResponse {
	s.StatusCode = &v
	return s
}

func (s *RootCaCrlResponse) SetBody(v string) *RootCaCrlResponse {
	s.Body = &v
	return s
}

type SimplePackagePckCertRequest struct {
	AcsHost       *string `json:"AcsHost,omitempty" xml:"AcsHost,omitempty"`
	ClientVpcId   *string `json:"ClientVpcId,omitempty" xml:"ClientVpcId,omitempty"`
	Cpusvn        *string `json:"cpusvn,omitempty" xml:"cpusvn,omitempty"`
	EncryptedPpid *string `json:"encrypted_ppid,omitempty" xml:"encrypted_ppid,omitempty"`
	Pceid         *string `json:"pceid,omitempty" xml:"pceid,omitempty"`
	Pcesvn        *string `json:"pcesvn,omitempty" xml:"pcesvn,omitempty"`
	Qeid          *string `json:"qeid,omitempty" xml:"qeid,omitempty"`
}

func (s SimplePackagePckCertRequest) String() string {
	return tea.Prettify(s)
}

func (s SimplePackagePckCertRequest) GoString() string {
	return s.String()
}

func (s *SimplePackagePckCertRequest) SetAcsHost(v string) *SimplePackagePckCertRequest {
	s.AcsHost = &v
	return s
}

func (s *SimplePackagePckCertRequest) SetClientVpcId(v string) *SimplePackagePckCertRequest {
	s.ClientVpcId = &v
	return s
}

func (s *SimplePackagePckCertRequest) SetCpusvn(v string) *SimplePackagePckCertRequest {
	s.Cpusvn = &v
	return s
}

func (s *SimplePackagePckCertRequest) SetEncryptedPpid(v string) *SimplePackagePckCertRequest {
	s.EncryptedPpid = &v
	return s
}

func (s *SimplePackagePckCertRequest) SetPceid(v string) *SimplePackagePckCertRequest {
	s.Pceid = &v
	return s
}

func (s *SimplePackagePckCertRequest) SetPcesvn(v string) *SimplePackagePckCertRequest {
	s.Pcesvn = &v
	return s
}

func (s *SimplePackagePckCertRequest) SetQeid(v string) *SimplePackagePckCertRequest {
	s.Qeid = &v
	return s
}

type SimplePackagePckCertResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SimplePackagePckCertResponse) String() string {
	return tea.Prettify(s)
}

func (s SimplePackagePckCertResponse) GoString() string {
	return s.String()
}

func (s *SimplePackagePckCertResponse) SetHeaders(v map[string]*string) *SimplePackagePckCertResponse {
	s.Headers = v
	return s
}

func (s *SimplePackagePckCertResponse) SetStatusCode(v int32) *SimplePackagePckCertResponse {
	s.StatusCode = &v
	return s
}

func (s *SimplePackagePckCertResponse) SetBody(v string) *SimplePackagePckCertResponse {
	s.Body = &v
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("sgx-dcap-server"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) GetQeIdentityWithOptions(request *GetQeIdentityRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQeIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsHost)) {
		query["AcsHost"] = request.AcsHost
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVpcId)) {
		query["ClientVpcId"] = request.ClientVpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQeIdentity"),
		Version:     tea.String("2020-07-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sgx/certification/v3/qe/identity"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &GetQeIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQeIdentity(request *GetQeIdentityRequest) (_result *GetQeIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQeIdentityResponse{}
	_body, _err := client.GetQeIdentityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQveIdentityWithOptions(request *GetQveIdentityRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQveIdentityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsHost)) {
		query["AcsHost"] = request.AcsHost
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVpcId)) {
		query["ClientVpcId"] = request.ClientVpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQveIdentity"),
		Version:     tea.String("2020-07-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sgx/certification/v3/qve/identity"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &GetQveIdentityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQveIdentity(request *GetQveIdentityRequest) (_result *GetQveIdentityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQveIdentityResponse{}
	_body, _err := client.GetQveIdentityWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTcbInfoWithOptions(request *GetTcbInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTcbInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsHost)) {
		query["AcsHost"] = request.AcsHost
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVpcId)) {
		query["ClientVpcId"] = request.ClientVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Fmspc)) {
		query["fmspc"] = request.Fmspc
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTcbInfo"),
		Version:     tea.String("2020-07-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sgx/certification/v3/tcb"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &GetTcbInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTcbInfo(request *GetTcbInfoRequest) (_result *GetTcbInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTcbInfoResponse{}
	_body, _err := client.GetTcbInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PckCrlWithOptions(request *PckCrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PckCrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsHost)) {
		query["AcsHost"] = request.AcsHost
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVpcId)) {
		query["ClientVpcId"] = request.ClientVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Ca)) {
		query["ca"] = request.Ca
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PckCrl"),
		Version:     tea.String("2020-07-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sgx/certification/v3/pckcrl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &PckCrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PckCrl(request *PckCrlRequest) (_result *PckCrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PckCrlResponse{}
	_body, _err := client.PckCrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RootCaCrlWithOptions(request *RootCaCrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RootCaCrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsHost)) {
		query["AcsHost"] = request.AcsHost
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVpcId)) {
		query["ClientVpcId"] = request.ClientVpcId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RootCaCrl"),
		Version:     tea.String("2020-07-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sgx/certification/v3/rootcacrl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &RootCaCrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RootCaCrl(request *RootCaCrlRequest) (_result *RootCaCrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RootCaCrlResponse{}
	_body, _err := client.RootCaCrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SimplePackagePckCertWithOptions(request *SimplePackagePckCertRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SimplePackagePckCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsHost)) {
		query["AcsHost"] = request.AcsHost
	}

	if !tea.BoolValue(util.IsUnset(request.ClientVpcId)) {
		query["ClientVpcId"] = request.ClientVpcId
	}

	if !tea.BoolValue(util.IsUnset(request.Cpusvn)) {
		query["cpusvn"] = request.Cpusvn
	}

	if !tea.BoolValue(util.IsUnset(request.EncryptedPpid)) {
		query["encrypted_ppid"] = request.EncryptedPpid
	}

	if !tea.BoolValue(util.IsUnset(request.Pceid)) {
		query["pceid"] = request.Pceid
	}

	if !tea.BoolValue(util.IsUnset(request.Pcesvn)) {
		query["pcesvn"] = request.Pcesvn
	}

	if !tea.BoolValue(util.IsUnset(request.Qeid)) {
		query["qeid"] = request.Qeid
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SimplePackagePckCert"),
		Version:     tea.String("2020-07-26"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/sgx/certification/v3/pckcert"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &SimplePackagePckCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SimplePackagePckCert(request *SimplePackagePckCertRequest) (_result *SimplePackagePckCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SimplePackagePckCertResponse{}
	_body, _err := client.SimplePackagePckCertWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
