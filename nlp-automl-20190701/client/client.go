// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddMTInterveneWordRequest struct {
	// example:
	//
	// 1
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// hello
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty"`
	TargetText *string `json:"TargetText,omitempty" xml:"TargetText,omitempty"`
}

func (s AddMTInterveneWordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMTInterveneWordRequest) GoString() string {
	return s.String()
}

func (s *AddMTInterveneWordRequest) SetPackageId(v string) *AddMTInterveneWordRequest {
	s.PackageId = &v
	return s
}

func (s *AddMTInterveneWordRequest) SetProjectId(v string) *AddMTInterveneWordRequest {
	s.ProjectId = &v
	return s
}

func (s *AddMTInterveneWordRequest) SetSourceText(v string) *AddMTInterveneWordRequest {
	s.SourceText = &v
	return s
}

func (s *AddMTInterveneWordRequest) SetTargetText(v string) *AddMTInterveneWordRequest {
	s.TargetText = &v
	return s
}

type AddMTInterveneWordResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// parameterError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 46E6B40D-2B6C-571B-AC41-86207DE288A5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	WordId *int64 `json:"WordId,omitempty" xml:"WordId,omitempty"`
}

func (s AddMTInterveneWordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddMTInterveneWordResponseBody) GoString() string {
	return s.String()
}

func (s *AddMTInterveneWordResponseBody) SetCode(v int32) *AddMTInterveneWordResponseBody {
	s.Code = &v
	return s
}

func (s *AddMTInterveneWordResponseBody) SetMessage(v string) *AddMTInterveneWordResponseBody {
	s.Message = &v
	return s
}

func (s *AddMTInterveneWordResponseBody) SetRequestId(v string) *AddMTInterveneWordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMTInterveneWordResponseBody) SetWordId(v int64) *AddMTInterveneWordResponseBody {
	s.WordId = &v
	return s
}

type AddMTInterveneWordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddMTInterveneWordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddMTInterveneWordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMTInterveneWordResponse) GoString() string {
	return s.String()
}

func (s *AddMTInterveneWordResponse) SetHeaders(v map[string]*string) *AddMTInterveneWordResponse {
	s.Headers = v
	return s
}

func (s *AddMTInterveneWordResponse) SetStatusCode(v int32) *AddMTInterveneWordResponse {
	s.StatusCode = &v
	return s
}

func (s *AddMTInterveneWordResponse) SetBody(v *AddMTInterveneWordResponseBody) *AddMTInterveneWordResponse {
	s.Body = v
	return s
}

type GetPredictDocRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1000
	DocId   *int64  `json:"DocId,omitempty" xml:"DocId,omitempty"`
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s GetPredictDocRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPredictDocRequest) GoString() string {
	return s.String()
}

func (s *GetPredictDocRequest) SetDocId(v int64) *GetPredictDocRequest {
	s.DocId = &v
	return s
}

func (s *GetPredictDocRequest) SetProduct(v string) *GetPredictDocRequest {
	s.Product = &v
	return s
}

type GetPredictDocResponseBody struct {
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// xxx
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xxx
	XLIFFInfo *string `json:"XLIFFInfo,omitempty" xml:"XLIFFInfo,omitempty"`
}

func (s GetPredictDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPredictDocResponseBody) GoString() string {
	return s.String()
}

func (s *GetPredictDocResponseBody) SetRequestId(v string) *GetPredictDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPredictDocResponseBody) SetResultContent(v string) *GetPredictDocResponseBody {
	s.ResultContent = &v
	return s
}

func (s *GetPredictDocResponseBody) SetStatus(v int32) *GetPredictDocResponseBody {
	s.Status = &v
	return s
}

func (s *GetPredictDocResponseBody) SetXLIFFInfo(v string) *GetPredictDocResponseBody {
	s.XLIFFInfo = &v
	return s
}

type GetPredictDocResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPredictDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPredictDocResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPredictDocResponse) GoString() string {
	return s.String()
}

func (s *GetPredictDocResponse) SetHeaders(v map[string]*string) *GetPredictDocResponse {
	s.Headers = v
	return s
}

func (s *GetPredictDocResponse) SetStatusCode(v int32) *GetPredictDocResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPredictDocResponse) SetBody(v *GetPredictDocResponseBody) *GetPredictDocResponse {
	s.Body = v
	return s
}

type PredictMTModelByDocRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [base64 encode content]
	FileContent *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// docx
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ModelId *int32 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// v1
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// true
	NeedXLIFF *bool `json:"NeedXLIFF,omitempty" xml:"NeedXLIFF,omitempty"`
}

func (s PredictMTModelByDocRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictMTModelByDocRequest) GoString() string {
	return s.String()
}

func (s *PredictMTModelByDocRequest) SetFileContent(v string) *PredictMTModelByDocRequest {
	s.FileContent = &v
	return s
}

func (s *PredictMTModelByDocRequest) SetFileType(v string) *PredictMTModelByDocRequest {
	s.FileType = &v
	return s
}

func (s *PredictMTModelByDocRequest) SetModelId(v int32) *PredictMTModelByDocRequest {
	s.ModelId = &v
	return s
}

func (s *PredictMTModelByDocRequest) SetModelVersion(v string) *PredictMTModelByDocRequest {
	s.ModelVersion = &v
	return s
}

func (s *PredictMTModelByDocRequest) SetNeedXLIFF(v bool) *PredictMTModelByDocRequest {
	s.NeedXLIFF = &v
	return s
}

type PredictMTModelByDocResponseBody struct {
	// example:
	//
	// 1
	DocId *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	// example:
	//
	// 86D18195-D89C-4C8C-9DC4-5FCE789CE6D5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PredictMTModelByDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PredictMTModelByDocResponseBody) GoString() string {
	return s.String()
}

func (s *PredictMTModelByDocResponseBody) SetDocId(v string) *PredictMTModelByDocResponseBody {
	s.DocId = &v
	return s
}

func (s *PredictMTModelByDocResponseBody) SetRequestId(v string) *PredictMTModelByDocResponseBody {
	s.RequestId = &v
	return s
}

type PredictMTModelByDocResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PredictMTModelByDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PredictMTModelByDocResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictMTModelByDocResponse) GoString() string {
	return s.String()
}

func (s *PredictMTModelByDocResponse) SetHeaders(v map[string]*string) *PredictMTModelByDocResponse {
	s.Headers = v
	return s
}

func (s *PredictMTModelByDocResponse) SetStatusCode(v int32) *PredictMTModelByDocResponse {
	s.StatusCode = &v
	return s
}

func (s *PredictMTModelByDocResponse) SetBody(v *PredictMTModelByDocResponseBody) *PredictMTModelByDocResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nlp-automl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AddMTInterveneWordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMTInterveneWordResponse
func (client *Client) AddMTInterveneWordWithOptions(request *AddMTInterveneWordRequest, runtime *util.RuntimeOptions) (_result *AddMTInterveneWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PackageId)) {
		query["PackageId"] = request.PackageId
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceText)) {
		query["SourceText"] = request.SourceText
	}

	if !tea.BoolValue(util.IsUnset(request.TargetText)) {
		query["TargetText"] = request.TargetText
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddMTInterveneWord"),
		Version:     tea.String("2019-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddMTInterveneWordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddMTInterveneWordRequest
//
// @return AddMTInterveneWordResponse
func (client *Client) AddMTInterveneWord(request *AddMTInterveneWordRequest) (_result *AddMTInterveneWordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMTInterveneWordResponse{}
	_body, _err := client.AddMTInterveneWordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetPredictDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPredictDocResponse
func (client *Client) GetPredictDocWithOptions(request *GetPredictDocRequest, runtime *util.RuntimeOptions) (_result *GetPredictDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocId)) {
		query["DocId"] = request.DocId
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPredictDoc"),
		Version:     tea.String("2019-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPredictDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetPredictDocRequest
//
// @return GetPredictDocResponse
func (client *Client) GetPredictDoc(request *GetPredictDocRequest) (_result *GetPredictDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPredictDocResponse{}
	_body, _err := client.GetPredictDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - PredictMTModelByDocRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PredictMTModelByDocResponse
func (client *Client) PredictMTModelByDocWithOptions(request *PredictMTModelByDocRequest, runtime *util.RuntimeOptions) (_result *PredictMTModelByDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileType)) {
		query["FileType"] = request.FileType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		query["ModelVersion"] = request.ModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.NeedXLIFF)) {
		query["NeedXLIFF"] = request.NeedXLIFF
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileContent)) {
		body["FileContent"] = request.FileContent
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PredictMTModelByDoc"),
		Version:     tea.String("2019-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PredictMTModelByDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - PredictMTModelByDocRequest
//
// @return PredictMTModelByDocResponse
func (client *Client) PredictMTModelByDoc(request *PredictMTModelByDocRequest) (_result *PredictMTModelByDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictMTModelByDocResponse{}
	_body, _err := client.PredictMTModelByDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
