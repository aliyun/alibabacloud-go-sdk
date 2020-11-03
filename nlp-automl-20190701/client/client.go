// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type PredictMTModelByDocRequest struct {
	FileContent  *string `json:"FileContent,omitempty" xml:"FileContent,omitempty" require:"true"`
	FileType     *string `json:"FileType,omitempty" xml:"FileType,omitempty" require:"true"`
	ModelId      *int    `json:"ModelId,omitempty" xml:"ModelId,omitempty" require:"true"`
	NeedXLIFF    *bool   `json:"NeedXLIFF,omitempty" xml:"NeedXLIFF,omitempty"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty" require:"true"`
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

func (s *PredictMTModelByDocRequest) SetModelId(v int) *PredictMTModelByDocRequest {
	s.ModelId = &v
	return s
}

func (s *PredictMTModelByDocRequest) SetNeedXLIFF(v bool) *PredictMTModelByDocRequest {
	s.NeedXLIFF = &v
	return s
}

func (s *PredictMTModelByDocRequest) SetModelVersion(v string) *PredictMTModelByDocRequest {
	s.ModelVersion = &v
	return s
}

type PredictMTModelByDocResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DocId     *string `json:"DocId,omitempty" xml:"DocId,omitempty" require:"true"`
}

func (s PredictMTModelByDocResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictMTModelByDocResponse) GoString() string {
	return s.String()
}

func (s *PredictMTModelByDocResponse) SetRequestId(v string) *PredictMTModelByDocResponse {
	s.RequestId = &v
	return s
}

func (s *PredictMTModelByDocResponse) SetDocId(v string) *PredictMTModelByDocResponse {
	s.DocId = &v
	return s
}

type BindIntervenePackageAndModelRequest struct {
	PackageId    *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty" require:"true"`
	ModelId      *int64  `json:"ModelId,omitempty" xml:"ModelId,omitempty" require:"true"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	ProjectId    *int64  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" require:"true"`
}

func (s BindIntervenePackageAndModelRequest) String() string {
	return tea.Prettify(s)
}

func (s BindIntervenePackageAndModelRequest) GoString() string {
	return s.String()
}

func (s *BindIntervenePackageAndModelRequest) SetPackageId(v int64) *BindIntervenePackageAndModelRequest {
	s.PackageId = &v
	return s
}

func (s *BindIntervenePackageAndModelRequest) SetModelId(v int64) *BindIntervenePackageAndModelRequest {
	s.ModelId = &v
	return s
}

func (s *BindIntervenePackageAndModelRequest) SetModelVersion(v string) *BindIntervenePackageAndModelRequest {
	s.ModelVersion = &v
	return s
}

func (s *BindIntervenePackageAndModelRequest) SetProjectId(v int64) *BindIntervenePackageAndModelRequest {
	s.ProjectId = &v
	return s
}

type BindIntervenePackageAndModelResponse struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *int    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BindIntervenePackageAndModelResponse) String() string {
	return tea.Prettify(s)
}

func (s BindIntervenePackageAndModelResponse) GoString() string {
	return s.String()
}

func (s *BindIntervenePackageAndModelResponse) SetCode(v int) *BindIntervenePackageAndModelResponse {
	s.Code = &v
	return s
}

func (s *BindIntervenePackageAndModelResponse) SetMessage(v int) *BindIntervenePackageAndModelResponse {
	s.Message = &v
	return s
}

func (s *BindIntervenePackageAndModelResponse) SetSuccess(v string) *BindIntervenePackageAndModelResponse {
	s.Success = &v
	return s
}

func (s *BindIntervenePackageAndModelResponse) SetRequestId(v string) *BindIntervenePackageAndModelResponse {
	s.RequestId = &v
	return s
}

type AddMtIntervenePackageRequest struct {
	PackageName    *string `json:"PackageName,omitempty" xml:"PackageName,omitempty" require:"true"`
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty" require:"true"`
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty" require:"true"`
	ProjectId      *int    `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" require:"true"`
	SourceType     *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s AddMtIntervenePackageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMtIntervenePackageRequest) GoString() string {
	return s.String()
}

func (s *AddMtIntervenePackageRequest) SetPackageName(v string) *AddMtIntervenePackageRequest {
	s.PackageName = &v
	return s
}

func (s *AddMtIntervenePackageRequest) SetSourceLanguage(v string) *AddMtIntervenePackageRequest {
	s.SourceLanguage = &v
	return s
}

func (s *AddMtIntervenePackageRequest) SetTargetLanguage(v string) *AddMtIntervenePackageRequest {
	s.TargetLanguage = &v
	return s
}

func (s *AddMtIntervenePackageRequest) SetProjectId(v int) *AddMtIntervenePackageRequest {
	s.ProjectId = &v
	return s
}

func (s *AddMtIntervenePackageRequest) SetSourceType(v string) *AddMtIntervenePackageRequest {
	s.SourceType = &v
	return s
}

type AddMtIntervenePackageResponse struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *int    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddMtIntervenePackageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMtIntervenePackageResponse) GoString() string {
	return s.String()
}

func (s *AddMtIntervenePackageResponse) SetCode(v int) *AddMtIntervenePackageResponse {
	s.Code = &v
	return s
}

func (s *AddMtIntervenePackageResponse) SetMessage(v int) *AddMtIntervenePackageResponse {
	s.Message = &v
	return s
}

func (s *AddMtIntervenePackageResponse) SetPackageId(v string) *AddMtIntervenePackageResponse {
	s.PackageId = &v
	return s
}

func (s *AddMtIntervenePackageResponse) SetRequestId(v string) *AddMtIntervenePackageResponse {
	s.RequestId = &v
	return s
}

type GetPredictDocRequest struct {
	DocId   *int64  `json:"DocId,omitempty" xml:"DocId,omitempty" require:"true"`
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

type GetPredictDocResponse struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ResultContent *string `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" require:"true"`
	Status        *int    `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	XLIFFInfo     *string `json:"XLIFFInfo,omitempty" xml:"XLIFFInfo,omitempty" require:"true"`
}

func (s GetPredictDocResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPredictDocResponse) GoString() string {
	return s.String()
}

func (s *GetPredictDocResponse) SetRequestId(v string) *GetPredictDocResponse {
	s.RequestId = &v
	return s
}

func (s *GetPredictDocResponse) SetResultContent(v string) *GetPredictDocResponse {
	s.ResultContent = &v
	return s
}

func (s *GetPredictDocResponse) SetStatus(v int) *GetPredictDocResponse {
	s.Status = &v
	return s
}

func (s *GetPredictDocResponse) SetXLIFFInfo(v string) *GetPredictDocResponse {
	s.XLIFFInfo = &v
	return s
}

type AddMTInterveneWordRequest struct {
	SourceText *string `json:"SourceText,omitempty" xml:"SourceText,omitempty" require:"true"`
	TargetText *string `json:"TargetText,omitempty" xml:"TargetText,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty" require:"true"`
	PackageId  *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s AddMTInterveneWordRequest) String() string {
	return tea.Prettify(s)
}

func (s AddMTInterveneWordRequest) GoString() string {
	return s.String()
}

func (s *AddMTInterveneWordRequest) SetSourceText(v string) *AddMTInterveneWordRequest {
	s.SourceText = &v
	return s
}

func (s *AddMTInterveneWordRequest) SetTargetText(v string) *AddMTInterveneWordRequest {
	s.TargetText = &v
	return s
}

func (s *AddMTInterveneWordRequest) SetProjectId(v string) *AddMTInterveneWordRequest {
	s.ProjectId = &v
	return s
}

func (s *AddMTInterveneWordRequest) SetPackageId(v string) *AddMTInterveneWordRequest {
	s.PackageId = &v
	return s
}

type AddMTInterveneWordResponse struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *int    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	WordId    *string `json:"WordId,omitempty" xml:"WordId,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddMTInterveneWordResponse) String() string {
	return tea.Prettify(s)
}

func (s AddMTInterveneWordResponse) GoString() string {
	return s.String()
}

func (s *AddMTInterveneWordResponse) SetCode(v int) *AddMTInterveneWordResponse {
	s.Code = &v
	return s
}

func (s *AddMTInterveneWordResponse) SetMessage(v int) *AddMTInterveneWordResponse {
	s.Message = &v
	return s
}

func (s *AddMTInterveneWordResponse) SetWordId(v string) *AddMTInterveneWordResponse {
	s.WordId = &v
	return s
}

func (s *AddMTInterveneWordResponse) SetRequestId(v string) *AddMTInterveneWordResponse {
	s.RequestId = &v
	return s
}

type PredictMTModelRequest struct {
	ModelId      *string `json:"ModelId,omitempty" xml:"ModelId,omitempty" require:"true"`
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s PredictMTModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PredictMTModelRequest) GoString() string {
	return s.String()
}

func (s *PredictMTModelRequest) SetModelId(v string) *PredictMTModelRequest {
	s.ModelId = &v
	return s
}

func (s *PredictMTModelRequest) SetModelVersion(v string) *PredictMTModelRequest {
	s.ModelVersion = &v
	return s
}

func (s *PredictMTModelRequest) SetContent(v string) *PredictMTModelRequest {
	s.Content = &v
	return s
}

func (s *PredictMTModelRequest) SetProduct(v string) *PredictMTModelRequest {
	s.Product = &v
	return s
}

type PredictMTModelResponse struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *int    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s PredictMTModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PredictMTModelResponse) GoString() string {
	return s.String()
}

func (s *PredictMTModelResponse) SetCode(v int) *PredictMTModelResponse {
	s.Code = &v
	return s
}

func (s *PredictMTModelResponse) SetMessage(v int) *PredictMTModelResponse {
	s.Message = &v
	return s
}

func (s *PredictMTModelResponse) SetData(v string) *PredictMTModelResponse {
	s.Data = &v
	return s
}

func (s *PredictMTModelResponse) SetRequestId(v string) *PredictMTModelResponse {
	s.RequestId = &v
	return s
}

type InvokeActionRequest struct {
	InvokeProduct *string `json:"InvokeProduct,omitempty" xml:"InvokeProduct,omitempty"`
	InvokeRegion  *string `json:"InvokeRegion,omitempty" xml:"InvokeRegion,omitempty"`
	InvokeAction  *string `json:"InvokeAction,omitempty" xml:"InvokeAction,omitempty" require:"true"`
	InvokeParams  *string `json:"InvokeParams,omitempty" xml:"InvokeParams,omitempty" require:"true"`
}

func (s InvokeActionRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeActionRequest) GoString() string {
	return s.String()
}

func (s *InvokeActionRequest) SetInvokeProduct(v string) *InvokeActionRequest {
	s.InvokeProduct = &v
	return s
}

func (s *InvokeActionRequest) SetInvokeRegion(v string) *InvokeActionRequest {
	s.InvokeRegion = &v
	return s
}

func (s *InvokeActionRequest) SetInvokeAction(v string) *InvokeActionRequest {
	s.InvokeAction = &v
	return s
}

func (s *InvokeActionRequest) SetInvokeParams(v string) *InvokeActionRequest {
	s.InvokeParams = &v
	return s
}

type InvokeActionResponse struct {
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Message   *int    `json:"Message,omitempty" xml:"Message,omitempty" require:"true"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s InvokeActionResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeActionResponse) GoString() string {
	return s.String()
}

func (s *InvokeActionResponse) SetCode(v int) *InvokeActionResponse {
	s.Code = &v
	return s
}

func (s *InvokeActionResponse) SetMessage(v int) *InvokeActionResponse {
	s.Message = &v
	return s
}

func (s *InvokeActionResponse) SetData(v string) *InvokeActionResponse {
	s.Data = &v
	return s
}

func (s *InvokeActionResponse) SetRequestId(v string) *InvokeActionResponse {
	s.RequestId = &v
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

func (client *Client) PredictMTModelByDocWithOptions(request *PredictMTModelByDocRequest, runtime *util.RuntimeOptions) (_result *PredictMTModelByDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PredictMTModelByDocResponse{}
	_body, _err := client.DoRequest(tea.String("PredictMTModelByDoc"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) BindIntervenePackageAndModelWithOptions(request *BindIntervenePackageAndModelRequest, runtime *util.RuntimeOptions) (_result *BindIntervenePackageAndModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindIntervenePackageAndModelResponse{}
	_body, _err := client.DoRequest(tea.String("BindIntervenePackageAndModel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK,APP,PrivateKey"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindIntervenePackageAndModel(request *BindIntervenePackageAndModelRequest) (_result *BindIntervenePackageAndModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindIntervenePackageAndModelResponse{}
	_body, _err := client.BindIntervenePackageAndModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddMtIntervenePackageWithOptions(request *AddMtIntervenePackageRequest, runtime *util.RuntimeOptions) (_result *AddMtIntervenePackageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddMtIntervenePackageResponse{}
	_body, _err := client.DoRequest(tea.String("AddMtIntervenePackage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK,APP,PrivateKey"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddMtIntervenePackage(request *AddMtIntervenePackageRequest) (_result *AddMtIntervenePackageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddMtIntervenePackageResponse{}
	_body, _err := client.AddMtIntervenePackageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPredictDocWithOptions(request *GetPredictDocRequest, runtime *util.RuntimeOptions) (_result *GetPredictDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPredictDocResponse{}
	_body, _err := client.DoRequest(tea.String("GetPredictDoc"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK,APP,PrivateKey,BearerToken"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) AddMTInterveneWordWithOptions(request *AddMTInterveneWordRequest, runtime *util.RuntimeOptions) (_result *AddMTInterveneWordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddMTInterveneWordResponse{}
	_body, _err := client.DoRequest(tea.String("AddMTInterveneWord"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK,APP,PrivateKey"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) PredictMTModelWithOptions(request *PredictMTModelRequest, runtime *util.RuntimeOptions) (_result *PredictMTModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PredictMTModelResponse{}
	_body, _err := client.DoRequest(tea.String("PredictMTModel"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PredictMTModel(request *PredictMTModelRequest) (_result *PredictMTModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PredictMTModelResponse{}
	_body, _err := client.PredictMTModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeActionWithOptions(request *InvokeActionRequest, runtime *util.RuntimeOptions) (_result *InvokeActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &InvokeActionResponse{}
	_body, _err := client.DoRequest(tea.String("InvokeAction"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-07-01"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeAction(request *InvokeActionRequest) (_result *InvokeActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeActionResponse{}
	_body, _err := client.InvokeActionWithOptions(request, runtime)
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
