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

type CheckResultRequest struct {
	ExtraImageControlList         *string `json:"ExtraImageControlList,omitempty" xml:"ExtraImageControlList,omitempty"`
	IsReturnImage                 *string `json:"IsReturnImage,omitempty" xml:"IsReturnImage,omitempty"`
	MerchantBizId                 *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	ReturnFiveCategorySpoofResult *string `json:"ReturnFiveCategorySpoofResult,omitempty" xml:"ReturnFiveCategorySpoofResult,omitempty"`
	TransactionId                 *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s CheckResultRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckResultRequest) GoString() string {
	return s.String()
}

func (s *CheckResultRequest) SetExtraImageControlList(v string) *CheckResultRequest {
	s.ExtraImageControlList = &v
	return s
}

func (s *CheckResultRequest) SetIsReturnImage(v string) *CheckResultRequest {
	s.IsReturnImage = &v
	return s
}

func (s *CheckResultRequest) SetMerchantBizId(v string) *CheckResultRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CheckResultRequest) SetReturnFiveCategorySpoofResult(v string) *CheckResultRequest {
	s.ReturnFiveCategorySpoofResult = &v
	return s
}

func (s *CheckResultRequest) SetTransactionId(v string) *CheckResultRequest {
	s.TransactionId = &v
	return s
}

type CheckResultResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *CheckResultResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CheckResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBody) SetCode(v string) *CheckResultResponseBody {
	s.Code = &v
	return s
}

func (s *CheckResultResponseBody) SetMessage(v string) *CheckResultResponseBody {
	s.Message = &v
	return s
}

func (s *CheckResultResponseBody) SetRequestId(v string) *CheckResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResultResponseBody) SetResult(v *CheckResultResponseBodyResult) *CheckResultResponseBody {
	s.Result = v
	return s
}

type CheckResultResponseBodyResult struct {
	EkycResult   *string `json:"EkycResult,omitempty" xml:"EkycResult,omitempty"`
	ExtBasicInfo *string `json:"ExtBasicInfo,omitempty" xml:"ExtBasicInfo,omitempty"`
	ExtFaceInfo  *string `json:"ExtFaceInfo,omitempty" xml:"ExtFaceInfo,omitempty"`
	ExtIdInfo    *string `json:"ExtIdInfo,omitempty" xml:"ExtIdInfo,omitempty"`
	ExtRiskInfo  *string `json:"ExtRiskInfo,omitempty" xml:"ExtRiskInfo,omitempty"`
	Passed       *string `json:"Passed,omitempty" xml:"Passed,omitempty"`
	SubCode      *string `json:"SubCode,omitempty" xml:"SubCode,omitempty"`
}

func (s CheckResultResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CheckResultResponseBodyResult) SetEkycResult(v string) *CheckResultResponseBodyResult {
	s.EkycResult = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtBasicInfo(v string) *CheckResultResponseBodyResult {
	s.ExtBasicInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtFaceInfo(v string) *CheckResultResponseBodyResult {
	s.ExtFaceInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtIdInfo(v string) *CheckResultResponseBodyResult {
	s.ExtIdInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetExtRiskInfo(v string) *CheckResultResponseBodyResult {
	s.ExtRiskInfo = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetPassed(v string) *CheckResultResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *CheckResultResponseBodyResult) SetSubCode(v string) *CheckResultResponseBodyResult {
	s.SubCode = &v
	return s
}

type CheckResultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckResultResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckResultResponse) GoString() string {
	return s.String()
}

func (s *CheckResultResponse) SetHeaders(v map[string]*string) *CheckResultResponse {
	s.Headers = v
	return s
}

func (s *CheckResultResponse) SetStatusCode(v int32) *CheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResultResponse) SetBody(v *CheckResultResponseBody) *CheckResultResponse {
	s.Body = v
	return s
}

type FaceCompareRequest struct {
	MerchantBizId        *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	SourceFacePicture    *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	TargetFacePicture    *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
}

func (s FaceCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareRequest) SetMerchantBizId(v string) *FaceCompareRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePicture(v string) *FaceCompareRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePictureUrl(v string) *FaceCompareRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePicture(v string) *FaceCompareRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePictureUrl(v string) *FaceCompareRequest {
	s.TargetFacePictureUrl = &v
	return s
}

type FaceCompareResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FaceCompareResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s FaceCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBody) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBody) SetCode(v string) *FaceCompareResponseBody {
	s.Code = &v
	return s
}

func (s *FaceCompareResponseBody) SetMessage(v string) *FaceCompareResponseBody {
	s.Message = &v
	return s
}

func (s *FaceCompareResponseBody) SetRequestId(v string) *FaceCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *FaceCompareResponseBody) SetResult(v *FaceCompareResponseBodyResult) *FaceCompareResponseBody {
	s.Result = v
	return s
}

type FaceCompareResponseBodyResult struct {
	FaceComparisonScore *float64 `json:"FaceComparisonScore,omitempty" xml:"FaceComparisonScore,omitempty"`
	Passed              *string  `json:"Passed,omitempty" xml:"Passed,omitempty"`
	TransactionId       *string  `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s FaceCompareResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FaceCompareResponseBodyResult) SetFaceComparisonScore(v float64) *FaceCompareResponseBodyResult {
	s.FaceComparisonScore = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetPassed(v string) *FaceCompareResponseBodyResult {
	s.Passed = &v
	return s
}

func (s *FaceCompareResponseBodyResult) SetTransactionId(v string) *FaceCompareResponseBodyResult {
	s.TransactionId = &v
	return s
}

type FaceCompareResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FaceCompareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FaceCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s FaceCompareResponse) GoString() string {
	return s.String()
}

func (s *FaceCompareResponse) SetHeaders(v map[string]*string) *FaceCompareResponse {
	s.Headers = v
	return s
}

func (s *FaceCompareResponse) SetStatusCode(v int32) *FaceCompareResponse {
	s.StatusCode = &v
	return s
}

func (s *FaceCompareResponse) SetBody(v *FaceCompareResponseBody) *FaceCompareResponse {
	s.Body = v
	return s
}

type InitializeRequest struct {
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	FacePictureUrl    *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	FlowType          *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	MerchantUserId    *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	MetaInfo          *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
	OperationMode     *string `json:"OperationMode,omitempty" xml:"OperationMode,omitempty"`
	Pages             *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	ProductCode       *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductConfig     *string `json:"ProductConfig,omitempty" xml:"ProductConfig,omitempty"`
	SceneCode         *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	ServiceLevel      *string `json:"ServiceLevel,omitempty" xml:"ServiceLevel,omitempty"`
}

func (s InitializeRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeRequest) GoString() string {
	return s.String()
}

func (s *InitializeRequest) SetDocType(v string) *InitializeRequest {
	s.DocType = &v
	return s
}

func (s *InitializeRequest) SetFacePictureBase64(v string) *InitializeRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *InitializeRequest) SetFacePictureUrl(v string) *InitializeRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *InitializeRequest) SetFlowType(v string) *InitializeRequest {
	s.FlowType = &v
	return s
}

func (s *InitializeRequest) SetMerchantBizId(v string) *InitializeRequest {
	s.MerchantBizId = &v
	return s
}

func (s *InitializeRequest) SetMerchantUserId(v string) *InitializeRequest {
	s.MerchantUserId = &v
	return s
}

func (s *InitializeRequest) SetMetaInfo(v string) *InitializeRequest {
	s.MetaInfo = &v
	return s
}

func (s *InitializeRequest) SetOperationMode(v string) *InitializeRequest {
	s.OperationMode = &v
	return s
}

func (s *InitializeRequest) SetPages(v string) *InitializeRequest {
	s.Pages = &v
	return s
}

func (s *InitializeRequest) SetProductCode(v string) *InitializeRequest {
	s.ProductCode = &v
	return s
}

func (s *InitializeRequest) SetProductConfig(v string) *InitializeRequest {
	s.ProductConfig = &v
	return s
}

func (s *InitializeRequest) SetSceneCode(v string) *InitializeRequest {
	s.SceneCode = &v
	return s
}

func (s *InitializeRequest) SetServiceLevel(v string) *InitializeRequest {
	s.ServiceLevel = &v
	return s
}

type InitializeResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *InitializeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s InitializeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeResponseBody) SetCode(v string) *InitializeResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeResponseBody) SetMessage(v string) *InitializeResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeResponseBody) SetRequestId(v string) *InitializeResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeResponseBody) SetResult(v *InitializeResponseBodyResult) *InitializeResponseBody {
	s.Result = v
	return s
}

type InitializeResponseBodyResult struct {
	ClientCfg     *string `json:"ClientCfg,omitempty" xml:"ClientCfg,omitempty"`
	TransactionId *string `json:"TransactionId,omitempty" xml:"TransactionId,omitempty"`
}

func (s InitializeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *InitializeResponseBodyResult) SetClientCfg(v string) *InitializeResponseBodyResult {
	s.ClientCfg = &v
	return s
}

func (s *InitializeResponseBodyResult) SetTransactionId(v string) *InitializeResponseBodyResult {
	s.TransactionId = &v
	return s
}

type InitializeResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitializeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeResponse) GoString() string {
	return s.String()
}

func (s *InitializeResponse) SetHeaders(v map[string]*string) *InitializeResponse {
	s.Headers = v
	return s
}

func (s *InitializeResponse) SetStatusCode(v int32) *InitializeResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeResponse) SetBody(v *InitializeResponseBody) *InitializeResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudauth-intl"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CheckResultWithOptions(request *CheckResultRequest, runtime *util.RuntimeOptions) (_result *CheckResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExtraImageControlList)) {
		query["ExtraImageControlList"] = request.ExtraImageControlList
	}

	if !tea.BoolValue(util.IsUnset(request.IsReturnImage)) {
		query["IsReturnImage"] = request.IsReturnImage
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnFiveCategorySpoofResult)) {
		query["ReturnFiveCategorySpoofResult"] = request.ReturnFiveCategorySpoofResult
	}

	if !tea.BoolValue(util.IsUnset(request.TransactionId)) {
		query["TransactionId"] = request.TransactionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckResult"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckResult(request *CheckResultRequest) (_result *CheckResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckResultResponse{}
	_body, _err := client.CheckResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FaceCompareWithOptions(request *FaceCompareRequest, runtime *util.RuntimeOptions) (_result *FaceCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFacePicture)) {
		query["SourceFacePicture"] = request.SourceFacePicture
	}

	if !tea.BoolValue(util.IsUnset(request.SourceFacePictureUrl)) {
		query["SourceFacePictureUrl"] = request.SourceFacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePicture)) {
		query["TargetFacePicture"] = request.TargetFacePicture
	}

	if !tea.BoolValue(util.IsUnset(request.TargetFacePictureUrl)) {
		query["TargetFacePictureUrl"] = request.TargetFacePictureUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FaceCompare"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FaceCompareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FaceCompare(request *FaceCompareRequest) (_result *FaceCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FaceCompareResponse{}
	_body, _err := client.FaceCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeWithOptions(request *InitializeRequest, runtime *util.RuntimeOptions) (_result *InitializeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		query["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureBase64)) {
		query["FacePictureBase64"] = request.FacePictureBase64
	}

	if !tea.BoolValue(util.IsUnset(request.FacePictureUrl)) {
		query["FacePictureUrl"] = request.FacePictureUrl
	}

	if !tea.BoolValue(util.IsUnset(request.FlowType)) {
		query["FlowType"] = request.FlowType
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantBizId)) {
		query["MerchantBizId"] = request.MerchantBizId
	}

	if !tea.BoolValue(util.IsUnset(request.MerchantUserId)) {
		query["MerchantUserId"] = request.MerchantUserId
	}

	if !tea.BoolValue(util.IsUnset(request.MetaInfo)) {
		query["MetaInfo"] = request.MetaInfo
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMode)) {
		query["OperationMode"] = request.OperationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Pages)) {
		query["Pages"] = request.Pages
	}

	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		query["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProductConfig)) {
		query["ProductConfig"] = request.ProductConfig
	}

	if !tea.BoolValue(util.IsUnset(request.SceneCode)) {
		query["SceneCode"] = request.SceneCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceLevel)) {
		query["ServiceLevel"] = request.ServiceLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Initialize"),
		Version:     tea.String("2022-08-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitializeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Initialize(request *InitializeRequest) (_result *InitializeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeResponse{}
	_body, _err := client.InitializeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
