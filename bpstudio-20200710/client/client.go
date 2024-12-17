// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BillingApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// 30GRJUY95TMYWBYJ
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 40
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	NextToken *int64 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1713597738
	RefId *string `json:"RefId,omitempty" xml:"RefId,omitempty"`
	// example:
	//
	// rg-aekz44tg3bnpyba
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// vsw-xxxxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// rgm-ecs
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// ecs
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetDeployDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0
	NextToken *int64 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E15B71B4-5D8F-5484-BC07-989E2987EE7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// example:
	//
	// EQ4W772D0VJ33IV1
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1645516991000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// DEPLOY
	CurrentProcess   *string                  `json:"CurrentProcess,omitempty" xml:"CurrentProcess,omitempty"`
	DeletingNodeList []map[string]interface{} `json:"DeletingNodeList,omitempty" xml:"DeletingNodeList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	DeployPercent     *float64                 `json:"DeployPercent,omitempty" xml:"DeployPercent,omitempty"`
	DeployedNodeList  []map[string]interface{} `json:"DeployedNodeList,omitempty" xml:"DeployedNodeList,omitempty" type:"Repeated"`
	DeployingNodeList []map[string]interface{} `json:"DeployingNodeList,omitempty" xml:"DeployingNodeList,omitempty" type:"Repeated"`
	// example:
	//
	// OperationNotSupport.SubscriptionCommodityCanNotDelete
	Error         *string   `json:"Error,omitempty" xml:"Error,omitempty"`
	ExecutionTime *int32    `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	FailStatus    *int32    `json:"FailStatus,omitempty" xml:"FailStatus,omitempty"`
	OrderIdList   []*string `json:"OrderIdList,omitempty" xml:"OrderIdList,omitempty" type:"Repeated"`
	// example:
	//
	// https://cadt-studio-publish.oss-cn-hangzhou.aliyuncs.com/1986207497633020/deployReport-EXN4FNUR12M35KJM.pdf?Expires=1716967763&OSSAccessKeyId=********nw4rvYAweFuQc3&Signature=*******fKo6164wykT9jBOsSGeQ%3D
	PdfUrl *string `json:"PdfUrl,omitempty" xml:"PdfUrl,omitempty"`
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceList    []*GetDeployDetailResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// example:
	//
	// Revised
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// https://cadt-studio-publish.oss-cn-hangzhou.aliyuncs.com/1986207497633020/app_EXN4FNUR12M35KJM.tf?Expires=1716967763&OSSAccessKeyId=*******nw4rvYAweFuQc3&Signature=%2********lYROqJLNvyA8g6qD9U%3D
	TerraformScriptUrl *string `json:"TerraformScriptUrl,omitempty" xml:"TerraformScriptUrl,omitempty"`
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
	BuyDuration *string `json:"BuyDuration,omitempty" xml:"BuyDuration,omitempty"`
	// example:
	//
	// Free
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// OpenApi
	ExecutionStrategy *string `json:"ExecutionStrategy,omitempty" xml:"ExecutionStrategy,omitempty"`
	ExpiredTime       *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// 1714031840000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// https://cloudmonitor.console.aliyun.com/?#/generalcharts/product=vpc&group=&type=&return&region=cn-hangzhou&dimension=instanceId:vpc-bp1n6uc5jqxtff2euhnx5
	MonitorURL *string `json:"MonitorURL,omitempty" xml:"MonitorURL,omitempty"`
	// example:
	//
	// vpc
	NodeName  *string                                               `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	Operation *GetDeployDetailResponseBodyDataResourceListOperation `json:"Operation,omitempty" xml:"Operation,omitempty" type:"Struct"`
	// example:
	//
	// 1716049421
	RefId *int64 `json:"RefId,omitempty" xml:"RefId,omitempty"`
	// example:
	//
	// {\\"Description\\":\\"\\",\\"ClassicLinkEnabled\\":false,\\"ResourceGroupId\\":\\"rg-acfm4mlwqjalz7a\\",\\"SecondaryCidrBlocks\\":[],\\"CidrBlock\\":\\"192.168.0.0/16\\",\\"UserCidrs\\":[],\\"NetworkAclNum\\":\\"0\\",\\"VRouterId\\":\\"vrt-m5ek7vcyld0x5ym8m9hix\\",\\"OwnerId\\":1986207497633020,\\"AssociatedCens\\":[],\\"id\\":\\"vpc-m5e3c9nz4lj19byan9m2g\\",\\"CloudResources\\":[{\\"ResourceCount\\":1,\\"ResourceType\\":\\"VSwitch\\"},{\\"ResourceCount\\":1,\\"ResourceType\\":\\"VRouter\\"},{\\"ResourceCount\\":1,\\"ResourceType\\":\\"RouteTable\\"}],\\"Tags\\":[],\\"Status\\":\\"Available\\",\\"IsDefault\\":false,\\"RequestId\\":\\"494646FB-14C0-5B4C-9684-B6EFBF4DF01C\\",\\"SupportIpv4Gateway\\":false,\\"Ipv4GatewayId\\":\\"\\",\\"VSwitchIds\\":[\\"vsw-m5egl9wtppiadysjwlgb1\\"],\\"VpcId\\":\\"vpc-m5e3c9nz4lj19byan9m2g\\",\\"CreationTime\\":\\"2024-04-25T07:56:59Z\\",\\"VpcName\\":\\"vpc\\",\\"refId\\":\\"1714031764_0\\",\\"RegionId\\":\\"cn-qingdao\\",\\"Ipv6CidrBlock\\":\\"\\",\\"Ipv6CidrBlocks\\":[]}
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// ecs
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// example:
	//
	// i-2zehnzxqixu1pywsfbx1
	ResourceId       *string                                                        `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceTimeList []*GetDeployDetailResponseBodyDataResourceListResourceTimeList `json:"ResourceTimeList,omitempty" xml:"ResourceTimeList,omitempty" type:"Repeated"`
	// example:
	//
	// vpc
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// Finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *GetDeployDetailResponseBodyDataResourceList) SetExpiredTime(v int64) *GetDeployDetailResponseBodyDataResourceList {
	s.ExpiredTime = &v
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

func (s *GetDeployDetailResponseBodyDataResourceList) SetResourceTimeList(v []*GetDeployDetailResponseBodyDataResourceListResourceTimeList) *GetDeployDetailResponseBodyDataResourceList {
	s.ResourceTimeList = v
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
	// example:
	//
	// ecsDelete
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

type GetDeployDetailResponseBodyDataResourceListResourceTimeList struct {
	BizId             *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CreationEndTime   *int64  `json:"CreationEndTime,omitempty" xml:"CreationEndTime,omitempty"`
	CreationStartTime *int64  `json:"CreationStartTime,omitempty" xml:"CreationStartTime,omitempty"`
	Id                *int64  `json:"id,omitempty" xml:"id,omitempty"`
}

func (s GetDeployDetailResponseBodyDataResourceListResourceTimeList) String() string {
	return tea.Prettify(s)
}

func (s GetDeployDetailResponseBodyDataResourceListResourceTimeList) GoString() string {
	return s.String()
}

func (s *GetDeployDetailResponseBodyDataResourceListResourceTimeList) SetBizId(v string) *GetDeployDetailResponseBodyDataResourceListResourceTimeList {
	s.BizId = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceListResourceTimeList) SetCreationEndTime(v int64) *GetDeployDetailResponseBodyDataResourceListResourceTimeList {
	s.CreationEndTime = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceListResourceTimeList) SetCreationStartTime(v int64) *GetDeployDetailResponseBodyDataResourceListResourceTimeList {
	s.CreationStartTime = &v
	return s
}

func (s *GetDeployDetailResponseBodyDataResourceListResourceTimeList) SetId(v int64) *GetDeployDetailResponseBodyDataResourceListResourceTimeList {
	s.Id = &v
	return s
}

type GetDeployDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeployDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	client.ProductId = tea.String("BPStudio")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
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

// Summary:
//
// BillingApplication
//
// @param request - BillingApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BillingApplicationResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &BillingApplicationResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &BillingApplicationResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// BillingApplication
//
// @param request - BillingApplicationRequest
//
// @return BillingApplicationResponse
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

// Summary:
//
// 分页查询部署清单
//
// @param request - GetDeployDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeployDetailResponse
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetDeployDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetDeployDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 分页查询部署清单
//
// @param request - GetDeployDetailRequest
//
// @return GetDeployDetailResponse
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
