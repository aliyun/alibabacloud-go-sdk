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

type BillingApplicationRequest struct {
	Month           *int32  `json:"Month,omitempty" xml:"Month,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Year            *int32  `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s BillingApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s BillingApplicationRequest) GoString() string {
	return s.String()
}

func (s *BillingApplicationRequest) SetMonth(v int32) *BillingApplicationRequest {
	s.Month = &v
	return s
}

func (s *BillingApplicationRequest) SetResourceGroupId(v string) *BillingApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BillingApplicationRequest) SetYear(v int32) *BillingApplicationRequest {
	s.Year = &v
	return s
}

type BillingApplicationResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BillingApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BillingApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *BillingApplicationResponseBody) SetCode(v int32) *BillingApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *BillingApplicationResponseBody) SetData(v int32) *BillingApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *BillingApplicationResponseBody) SetMessage(v string) *BillingApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *BillingApplicationResponseBody) SetRequestId(v string) *BillingApplicationResponseBody {
	s.RequestId = &v
	return s
}

type BillingApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BillingApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BillingApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s BillingApplicationResponse) GoString() string {
	return s.String()
}

func (s *BillingApplicationResponse) SetHeaders(v map[string]*string) *BillingApplicationResponse {
	s.Headers = v
	return s
}

func (s *BillingApplicationResponse) SetStatusCode(v int32) *BillingApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *BillingApplicationResponse) SetBody(v *BillingApplicationResponseBody) *BillingApplicationResponse {
	s.Body = v
	return s
}

type GetDeployDetailRequest struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	MaxResults      *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *int64  `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RefId           *string `json:"RefId,omitempty" xml:"RefId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName    *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDeployDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDeployDetailRequest) SetAppId(v string) *GetDeployDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetDeployDetailRequest) SetMaxResults(v int64) *GetDeployDetailRequest {
	s.MaxResults = &v
	return s
}

func (s *GetDeployDetailRequest) SetNextToken(v int64) *GetDeployDetailRequest {
	s.NextToken = &v
	return s
}

func (s *GetDeployDetailRequest) SetRefId(v string) *GetDeployDetailRequest {
	s.RefId = &v
	return s
}

func (s *GetDeployDetailRequest) SetResourceGroupId(v string) *GetDeployDetailRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDeployDetailRequest) SetResourceId(v string) *GetDeployDetailRequest {
	s.ResourceId = &v
	return s
}

func (s *GetDeployDetailRequest) SetResourceName(v string) *GetDeployDetailRequest {
	s.ResourceName = &v
	return s
}

func (s *GetDeployDetailRequest) SetResourceType(v string) *GetDeployDetailRequest {
	s.ResourceType = &v
	return s
}

type GetDeployDetailResponseBody struct {
	Code       *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*GetDeployDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken  *int64                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetDeployDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeployDetailResponseBody) SetCode(v int32) *GetDeployDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetDeployDetailResponseBody) SetData(v []*GetDeployDetailResponseBodyData) *GetDeployDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetDeployDetailResponseBody) SetMessage(v string) *GetDeployDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetDeployDetailResponseBody) SetNextToken(v int64) *GetDeployDetailResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetDeployDetailResponseBody) SetRequestId(v string) *GetDeployDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeployDetailResponseBody) SetTotalCount(v int64) *GetDeployDetailResponseBody {
	s.TotalCount = &v
	return s
}

type GetDeployDetailResponseBodyData struct {
	AppId              *string                                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime         *int64                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CurrentProcess     *string                                        `json:"CurrentProcess,omitempty" xml:"CurrentProcess,omitempty"`
	DeletingNodeList   []map[string]interface{}                       `json:"DeletingNodeList,omitempty" xml:"DeletingNodeList,omitempty" type:"Repeated"`
	DeployPercent      *float64                                       `json:"DeployPercent,omitempty" xml:"DeployPercent,omitempty"`
	DeployedNodeList   []map[string]interface{}                       `json:"DeployedNodeList,omitempty" xml:"DeployedNodeList,omitempty" type:"Repeated"`
	DeployingNodeList  []map[string]interface{}                       `json:"DeployingNodeList,omitempty" xml:"DeployingNodeList,omitempty" type:"Repeated"`
	Error              *string                                        `json:"Error,omitempty" xml:"Error,omitempty"`
	ExecutionTime      *int32                                         `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	FailStatus         *int32                                         `json:"FailStatus,omitempty" xml:"FailStatus,omitempty"`
	OrderIdList        []*string                                      `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty" type:"Repeated"`
	PdfUrl             *string                                        `json:"PdfUrl,omitempty" xml:"PdfUrl,omitempty"`
	ResourceGroupId    *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceList       []*GetDeployDetailResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	Status             *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	TerraformScriptUrl *string                                        `json:"TerraformScriptUrl,omitempty" xml:"TerraformScriptUrl,omitempty"`
}

func (s GetDeployDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDeployDetailResponseBodyData) SetAppId(v string) *GetDeployDetailResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetCreateTime(v int64) *GetDeployDetailResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetCurrentProcess(v string) *GetDeployDetailResponseBodyData {
	s.CurrentProcess = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetDeletingNodeList(v []map[string]interface{}) *GetDeployDetailResponseBodyData {
	s.DeletingNodeList = v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetDeployPercent(v float64) *GetDeployDetailResponseBodyData {
	s.DeployPercent = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetDeployedNodeList(v []map[string]interface{}) *GetDeployDetailResponseBodyData {
	s.DeployedNodeList = v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetDeployingNodeList(v []map[string]interface{}) *GetDeployDetailResponseBodyData {
	s.DeployingNodeList = v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetError(v string) *GetDeployDetailResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetExecutionTime(v int32) *GetDeployDetailResponseBodyData {
	s.ExecutionTime = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetFailStatus(v int32) *GetDeployDetailResponseBodyData {
	s.FailStatus = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetOrderIdList(v []*string) *GetDeployDetailResponseBodyData {
	s.OrderIdList = v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetPdfUrl(v string) *GetDeployDetailResponseBodyData {
	s.PdfUrl = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetResourceGroupId(v string) *GetDeployDetailResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetResourceList(v []*GetDeployDetailResponseBodyDataResourceList) *GetDeployDetailResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetStatus(v string) *GetDeployDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDeployDetailResponseBodyData) SetTerraformScriptUrl(v string) *GetDeployDetailResponseBodyData {
	s.TerraformScriptUrl = &v
	return s
}

type GetDeployDetailResponseBodyDataResourceList struct {
	BuyDuration       *string                                               `json:"BuyDuration,omitempty" xml:"BuyDuration,omitempty"`
	ChargeType        *string                                               `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ExecutionStrategy *string                                               `json:"ExecutionStrategy,omitempty" xml:"ExecutionStrategy,omitempty"`
	ModifiedTime      *int64                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	MonitorURL        *string                                               `json:"MonitorURL,omitempty" xml:"MonitorURL,omitempty"`
	NodeName          *string                                               `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Operation         *GetDeployDetailResponseBodyDataResourceListOperation `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Struct"`
	RefId             *int64                                                `json:"RefId,omitempty" xml:"RefId,omitempty"`
	Remark            *string                                               `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceCode      *string                                               `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	ResourceId        *string                                               `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType      *string                                               `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status            *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetDeployDetailResponseBodyDataResourceList) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetBuyDuration(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.BuyDuration = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetChargeType(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.ChargeType = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetExecutionStrategy(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.ExecutionStrategy = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetModifiedTime(v int64) *GetDeployDetailResponseBodyDataResourceList {
	s.ModifiedTime = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetMonitorURL(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.MonitorURL = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetNodeName(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.NodeName = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetOperation(v *GetDeployDetailResponseBodyDataResourceListOperation) *GetDeployDetailResponseBodyDataResourceList {
	s.Operation = v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetRefId(v int64) *GetDeployDetailResponseBodyDataResourceList {
	s.RefId = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetRemark(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.Remark = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetResourceCode(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.ResourceCode = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetResourceId(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.ResourceId = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetResourceType(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.ResourceType = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceList) SetStatus(v string) *GetDeployDetailResponseBodyDataResourceList {
	s.Status = &v
	return s
}

type GetDeployDetailResponseBodyDataResourceListOperation struct {
	Name       *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Operations map[string]interface{} `json:"Operations,omitempty" xml:"Operations,omitempty"`
}

func (s GetDeployDetailResponseBodyDataResourceListOperation) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailResponseBodyDataResourceListOperation) GoString() string {
	return s.String()
}

func (s *GetDeployDetailResponseBodyDataResourceListOperation) SetName(v string) *GetDeployDetailResponseBodyDataResourceListOperation {
	s.Name = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceListOperation) SetOperations(v map[string]interface{}) *GetDeployDetailResponseBodyDataResourceListOperation {
	s.Operations = v
	return s
}

type GetDeployDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDeployDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDeployDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDeployDetailResponse) SetHeaders(v map[string]*string) *GetDeployDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDeployDetailResponse) SetStatusCode(v int32) *GetDeployDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeployDetailResponse) SetBody(v *GetDeployDetailResponseBody) *GetDeployDetailResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("bpstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BillingApplicationWithOptions(request *BillingApplicationRequest, runtime *util.RuntimeOptions) (_result *BillingApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Month)) {
		body["Month"] = request.Month
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Year)) {
		body["Year"] = request.Year
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BillingApplication"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BillingApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BillingApplication(request *BillingApplicationRequest) (_result *BillingApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BillingApplicationResponse{}
	_body, _err := client.BillingApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDeployDetailWithOptions(request *GetDeployDetailRequest, runtime *util.RuntimeOptions) (_result *GetDeployDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RefId)) {
		body["RefId"] = request.RefId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDeployDetail"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDeployDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDeployDetail(request *GetDeployDetailRequest) (_result *GetDeployDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDeployDetailResponse{}
	_body, _err := client.GetDeployDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
