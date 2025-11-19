// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("wuying-personal-pc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建桌面镜像
//
// @param tmpReq - CreateDesktopImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDesktopImageResponse
func (client *Client) CreateDesktopImageWithOptions(tmpReq *CreateDesktopImageRequest, runtime *dara.RuntimeOptions) (_result *CreateDesktopImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDesktopImageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StartUpFilePathList) {
		request.StartUpFilePathListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartUpFilePathList, dara.String("StartUpFilePathList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.AutoCleanUserData) {
		query["AutoCleanUserData"] = request.AutoCleanUserData
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DiskType) {
		query["DiskType"] = request.DiskType
	}

	if !dara.IsNil(request.EnableStartUpConfig) {
		query["EnableStartUpConfig"] = request.EnableStartUpConfig
	}

	if !dara.IsNil(request.ImageName) {
		query["ImageName"] = request.ImageName
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.StartUpFilePathListShrink) {
		query["StartUpFilePathList"] = request.StartUpFilePathListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDesktopImage"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDesktopImageResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建桌面镜像
//
// @param request - CreateDesktopImageRequest
//
// @return CreateDesktopImageResponse
func (client *Client) CreateDesktopImage(request *CreateDesktopImageRequest) (_result *CreateDesktopImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDesktopImageResponse{}
	_body, _err := client.CreateDesktopImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 下单套餐包和核时包
//
// @param tmpReq - CreateOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithOptions(tmpReq *CreateOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DynamicProductParams) {
		request.DynamicProductParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DynamicProductParams, dara.String("DynamicProductParams"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.AutoPay) {
		query["AutoPay"] = request.AutoPay
	}

	if !dara.IsNil(request.DynamicProductParamsShrink) {
		query["DynamicProductParams"] = request.DynamicProductParamsShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.OrderFrom) {
		query["OrderFrom"] = request.OrderFrom
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrder"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下单套餐包和核时包
//
// @param request - CreateOrderRequest
//
// @return CreateOrderResponse
func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除云桌面
//
// @param request - DeleteDesktopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDesktopResponse
func (client *Client) DeleteDesktopWithOptions(request *DeleteDesktopRequest, runtime *dara.RuntimeOptions) (_result *DeleteDesktopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDesktop"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDesktopResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除云桌面
//
// @param request - DeleteDesktopRequest
//
// @return DeleteDesktopResponse
func (client *Client) DeleteDesktop(request *DeleteDesktopRequest) (_result *DeleteDesktopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDesktopResponse{}
	_body, _err := client.DeleteDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除桌面镜像
//
// @param request - DeleteDesktopImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDesktopImageResponse
func (client *Client) DeleteDesktopImageWithOptions(request *DeleteDesktopImageRequest, runtime *dara.RuntimeOptions) (_result *DeleteDesktopImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.IsCleanImageCode) {
		query["IsCleanImageCode"] = request.IsCleanImageCode
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDesktopImage"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDesktopImageResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除桌面镜像
//
// @param request - DeleteDesktopImageRequest
//
// @return DeleteDesktopImageResponse
func (client *Client) DeleteDesktopImage(request *DeleteDesktopImageRequest) (_result *DeleteDesktopImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDesktopImageResponse{}
	_body, _err := client.DeleteDesktopImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户可变更的镜像
//
// @param request - DescribeAccessibleImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAccessibleImagesResponse
func (client *Client) DescribeAccessibleImagesWithOptions(request *DescribeAccessibleImagesRequest, runtime *dara.RuntimeOptions) (_result *DescribeAccessibleImagesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.BizSource) {
		query["BizSource"] = request.BizSource
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.DesktopType) {
		query["DesktopType"] = request.DesktopType
	}

	if !dara.IsNil(request.ImageType) {
		query["ImageType"] = request.ImageType
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Scene) {
		query["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeAccessibleImages"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeAccessibleImagesResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户可变更的镜像
//
// @param request - DescribeAccessibleImagesRequest
//
// @return DescribeAccessibleImagesResponse
func (client *Client) DescribeAccessibleImages(request *DescribeAccessibleImagesRequest) (_result *DescribeAccessibleImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeAccessibleImagesResponse{}
	_body, _err := client.DescribeAccessibleImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询核时包包列表
//
// @param request - DescribeCorePackageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCorePackageListResponse
func (client *Client) DescribeCorePackageListWithOptions(request *DescribeCorePackageListRequest, runtime *dara.RuntimeOptions) (_result *DescribeCorePackageListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.QueryDeductionInstances) {
		query["QueryDeductionInstances"] = request.QueryDeductionInstances
	}

	if !dara.IsNil(request.QueryScenario) {
		query["QueryScenario"] = request.QueryScenario
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCorePackageList"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCorePackageListResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询核时包包列表
//
// @param request - DescribeCorePackageListRequest
//
// @return DescribeCorePackageListResponse
func (client *Client) DescribeCorePackageList(request *DescribeCorePackageListRequest) (_result *DescribeCorePackageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeCorePackageListResponse{}
	_body, _err := client.DescribeCorePackageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 软终端分tab查询云桌面列表 & 组织信息
//
// @param request - DescribeDesktopsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDesktopsResponse
func (client *Client) DescribeDesktopsWithOptions(request *DescribeDesktopsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDesktopsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.DisplayType) {
		query["DisplayType"] = request.DisplayType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDesktops"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 软终端分tab查询云桌面列表 & 组织信息
//
// @param request - DescribeDesktopsRequest
//
// @return DescribeDesktopsResponse
func (client *Client) DescribeDesktops(request *DescribeDesktopsRequest) (_result *DescribeDesktopsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDesktopsResponse{}
	_body, _err := client.DescribeDesktopsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据分享码查询镜像
//
// @param request - DescribeImageDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeImageDetailResponse
func (client *Client) DescribeImageDetailWithOptions(request *DescribeImageDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeImageDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ImageId) {
		query["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.ShareCode) {
		query["ShareCode"] = request.ShareCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeImageDetail"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeImageDetailResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据分享码查询镜像
//
// @param request - DescribeImageDetailRequest
//
// @return DescribeImageDetailResponse
func (client *Client) DescribeImageDetail(request *DescribeImageDetailRequest) (_result *DescribeImageDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeImageDetailResponse{}
	_body, _err := client.DescribeImageDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询套餐包订单列表
//
// @param tmpReq - DescribePackageOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePackageOrdersResponse
func (client *Client) DescribePackageOrdersWithOptions(tmpReq *DescribePackageOrdersRequest, runtime *dara.RuntimeOptions) (_result *DescribePackageOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePackageOrdersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DesktopIdList) {
		request.DesktopIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DesktopIdList, dara.String("DesktopIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OrderIdList) {
		request.OrderIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderIdList, dara.String("OrderIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OrderStatusList) {
		request.OrderStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderStatusList, dara.String("OrderStatusList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ProductTypeList) {
		request.ProductTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductTypeList, dara.String("ProductTypeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceIdList) {
		request.ResourceIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIdList, dara.String("ResourceIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DesktopIdListShrink) {
		query["DesktopIdList"] = request.DesktopIdListShrink
	}

	if !dara.IsNil(request.OrderIdListShrink) {
		query["OrderIdList"] = request.OrderIdListShrink
	}

	if !dara.IsNil(request.OrderStatusListShrink) {
		query["OrderStatusList"] = request.OrderStatusListShrink
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductTypeListShrink) {
		query["ProductTypeList"] = request.ProductTypeListShrink
	}

	if !dara.IsNil(request.ResourceIdListShrink) {
		query["ResourceIdList"] = request.ResourceIdListShrink
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePackageOrders"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePackageOrdersResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询套餐包订单列表
//
// @param request - DescribePackageOrdersRequest
//
// @return DescribePackageOrdersResponse
func (client *Client) DescribePackageOrders(request *DescribePackageOrdersRequest) (_result *DescribePackageOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePackageOrdersResponse{}
	_body, _err := client.DescribePackageOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 生成无影工作站的场景url
//
// @param request - GenerateWuyingServerSceneUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateWuyingServerSceneUrlResponse
func (client *Client) GenerateWuyingServerSceneUrlWithOptions(request *GenerateWuyingServerSceneUrlRequest, runtime *dara.RuntimeOptions) (_result *GenerateWuyingServerSceneUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientId) {
		body["ClientId"] = request.ClientId
	}

	if !dara.IsNil(request.ClientIp) {
		body["ClientIp"] = request.ClientIp
	}

	if !dara.IsNil(request.ClientOS) {
		body["ClientOS"] = request.ClientOS
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.ClientVersion) {
		body["ClientVersion"] = request.ClientVersion
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.LoginRegionId) {
		body["LoginRegionId"] = request.LoginRegionId
	}

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.ProductType) {
		body["ProductType"] = request.ProductType
	}

	if !dara.IsNil(request.Scene) {
		body["Scene"] = request.Scene
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	if !dara.IsNil(request.WuyingServerId) {
		body["WuyingServerId"] = request.WuyingServerId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateWuyingServerSceneUrl"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateWuyingServerSceneUrlResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成无影工作站的场景url
//
// @param request - GenerateWuyingServerSceneUrlRequest
//
// @return GenerateWuyingServerSceneUrlResponse
func (client *Client) GenerateWuyingServerSceneUrl(request *GenerateWuyingServerSceneUrlRequest) (_result *GenerateWuyingServerSceneUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateWuyingServerSceneUrlResponse{}
	_body, _err := client.GenerateWuyingServerSceneUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询展示商品列表
//
// @param request - GetProductListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductListResponse
func (client *Client) GetProductListWithOptions(request *GetProductListRequest, runtime *dara.RuntimeOptions) (_result *GetProductListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.ConfigVersion) {
		query["ConfigVersion"] = request.ConfigVersion
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProductList"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProductListResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询展示商品列表
//
// @param request - GetProductListRequest
//
// @return GetProductListResponse
func (client *Client) GetProductList(request *GetProductListRequest) (_result *GetProductListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProductListResponse{}
	_body, _err := client.GetProductListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户信息
//
// @param request - GetUserInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfoWithOptions(request *GetUserInfoRequest, runtime *dara.RuntimeOptions) (_result *GetUserInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserInfo"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserInfoResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户信息
//
// @param request - GetUserInfoRequest
//
// @return GetUserInfoResponse
func (client *Client) GetUserInfo(request *GetUserInfoRequest) (_result *GetUserInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserInfoResponse{}
	_body, _err := client.GetUserInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开机
//
// @param request - StartDesktopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDesktopResponse
func (client *Client) StartDesktopWithOptions(request *StartDesktopRequest, runtime *dara.RuntimeOptions) (_result *StartDesktopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDesktop"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDesktopResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开机
//
// @param request - StartDesktopRequest
//
// @return StartDesktopResponse
func (client *Client) StartDesktop(request *StartDesktopRequest) (_result *StartDesktopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDesktopResponse{}
	_body, _err := client.StartDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 关机
//
// @param request - StopDesktopRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDesktopResponse
func (client *Client) StopDesktopWithOptions(request *StopDesktopRequest, runtime *dara.RuntimeOptions) (_result *StopDesktopResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKey) {
		query["ApiKey"] = request.ApiKey
	}

	if !dara.IsNil(request.DesktopId) {
		query["DesktopId"] = request.DesktopId
	}

	if !dara.IsNil(request.SessionId) {
		query["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDesktop"),
		Version:     dara.String("2025-11-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDesktopResponse{}
	_body, _err := client.DoRPCRequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 关机
//
// @param request - StopDesktopRequest
//
// @return StopDesktopResponse
func (client *Client) StopDesktop(request *StopDesktopRequest) (_result *StopDesktopResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDesktopResponse{}
	_body, _err := client.StopDesktopWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
