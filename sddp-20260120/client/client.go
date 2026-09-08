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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hongkong": dara.String("sddp-api.cn-hongkong.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("sddp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 校验列加密实例权限
//
// Description:
//
// 创建列加密前检查实例的权限、引擎及运行条件。应同时检查响应中的 ErrorCode 和 ErrorMessage；仅 ErrorCode=Success 表示检查通过，请求成功本身不表示实例满足全部加密条件。此检查不会创建列加密规则。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - CheckDataMaskingInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckDataMaskingInstanceResponse
func (client *Client) CheckDataMaskingInstanceWithOptions(request *CheckDataMaskingInstanceRequest, runtime *dara.RuntimeOptions) (_result *CheckDataMaskingInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckDataMaskingInstance"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckDataMaskingInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验列加密实例权限
//
// Description:
//
// 创建列加密前检查实例的权限、引擎及运行条件。应同时检查响应中的 ErrorCode 和 ErrorMessage；仅 ErrorCode=Success 表示检查通过，请求成功本身不表示实例满足全部加密条件。此检查不会创建列加密规则。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - CheckDataMaskingInstanceRequest
//
// @return CheckDataMaskingInstanceResponse
func (client *Client) CheckDataMaskingInstance(request *CheckDataMaskingInstanceRequest) (_result *CheckDataMaskingInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckDataMaskingInstanceResponse{}
	_body, _err := client.CheckDataMaskingInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建列加密策略
//
// Description:
//
// 为同一实例下指定数据库、表和列创建列加密规则。调用前检查实例状态、支持算法、密钥、目标列和账号。SubRuleList 为按表分组的目标列表，Columns 是以英文逗号分隔的列名字符串；UserList 中的账号被授予 fullAccess 明文权限。请求成功仅表示已受理，必须回读 ListDataMaskingColumns 和 ListDataAssetAccounts 确认列状态、账号权限及期限。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param tmpReq - CreateDataMaskingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataMaskingRuleResponse
func (client *Client) CreateDataMaskingRuleWithOptions(tmpReq *CreateDataMaskingRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDataMaskingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataMaskingRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubRuleList) {
		request.SubRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubRuleList, dara.String("SubRuleList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.UserList) {
		request.UserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserList, dara.String("UserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EncAlgorithm) {
		query["EncAlgorithm"] = request.EncAlgorithm
	}

	if !dara.IsNil(request.EncryptionKeyId) {
		query["EncryptionKeyId"] = request.EncryptionKeyId
	}

	if !dara.IsNil(request.EncryptionKeyMode) {
		query["EncryptionKeyMode"] = request.EncryptionKeyMode
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.ExpireTimeOperation) {
		query["ExpireTimeOperation"] = request.ExpireTimeOperation
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskHandleId) {
		query["RiskHandleId"] = request.RiskHandleId
	}

	if !dara.IsNil(request.SubRuleListShrink) {
		query["SubRuleList"] = request.SubRuleListShrink
	}

	if !dara.IsNil(request.UserListShrink) {
		query["UserList"] = request.UserListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataMaskingRule"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataMaskingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建列加密策略
//
// Description:
//
// 为同一实例下指定数据库、表和列创建列加密规则。调用前检查实例状态、支持算法、密钥、目标列和账号。SubRuleList 为按表分组的目标列表，Columns 是以英文逗号分隔的列名字符串；UserList 中的账号被授予 fullAccess 明文权限。请求成功仅表示已受理，必须回读 ListDataMaskingColumns 和 ListDataAssetAccounts 确认列状态、账号权限及期限。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - CreateDataMaskingRuleRequest
//
// @return CreateDataMaskingRuleResponse
func (client *Client) CreateDataMaskingRule(request *CreateDataMaskingRuleRequest) (_result *CreateDataMaskingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataMaskingRuleResponse{}
	_body, _err := client.CreateDataMaskingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除列加密策略
//
// @param tmpReq - DeleteDataMaskingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataMaskingRuleResponse
func (client *Client) DeleteDataMaskingRuleWithOptions(tmpReq *DeleteDataMaskingRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataMaskingRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataMaskingRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SubRuleList) {
		request.SubRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubRuleList, dara.String("SubRuleList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.SubRuleListShrink) {
		query["SubRuleList"] = request.SubRuleListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataMaskingRule"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataMaskingRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除列加密策略
//
// @param request - DeleteDataMaskingRuleRequest
//
// @return DeleteDataMaskingRuleResponse
func (client *Client) DeleteDataMaskingRule(request *DeleteDataMaskingRuleRequest) (_result *DeleteDataMaskingRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataMaskingRuleResponse{}
	_body, _err := client.DeleteDataMaskingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 检查实例列加密状态
//
// Description:
//
// 按产品范围查询当前地域数据库账号的权限统计。明文账号对应 FullAccessCount；限制访问和禁止解密账号分别计入 RestrictedAccessCount、NoneAccessCount。未配置权限账号数可由 TotalCount 减去上述三类账号数得到。本接口不接受实例、库、表、列等筛选条件。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - GetDataMaskingAccountCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataMaskingAccountCountResponse
func (client *Client) GetDataMaskingAccountCountWithOptions(request *GetDataMaskingAccountCountRequest, runtime *dara.RuntimeOptions) (_result *GetDataMaskingAccountCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductIds) {
		query["ProductIds"] = request.ProductIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataMaskingAccountCount"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataMaskingAccountCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查实例列加密状态
//
// Description:
//
// 按产品范围查询当前地域数据库账号的权限统计。明文账号对应 FullAccessCount；限制访问和禁止解密账号分别计入 RestrictedAccessCount、NoneAccessCount。未配置权限账号数可由 TotalCount 减去上述三类账号数得到。本接口不接受实例、库、表、列等筛选条件。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - GetDataMaskingAccountCountRequest
//
// @return GetDataMaskingAccountCountResponse
func (client *Client) GetDataMaskingAccountCount(request *GetDataMaskingAccountCountRequest) (_result *GetDataMaskingAccountCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataMaskingAccountCountResponse{}
	_body, _err := client.GetDataMaskingAccountCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取列加密统计信息
//
// Description:
//
// 按地域、产品和分类分级模板查询总列数、敏感列数、已加密列数和加密失败列数。本接口返回统计卡片数据，不跟随列列表中的实例、库名、表名、列名或模型筛选。
//
// 为兼容历史识别结果，指定模板无首屏结果且未按敏感等级或识别模型过滤时，结果可能回退到内置通用识别结果。要求严格模板归属时，请同时使用该模板内的等级或模型条件核验。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - GetDataMaskingColumnCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataMaskingColumnCountResponse
func (client *Client) GetDataMaskingColumnCountWithOptions(request *GetDataMaskingColumnCountRequest, runtime *dara.RuntimeOptions) (_result *GetDataMaskingColumnCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductIds) {
		query["ProductIds"] = request.ProductIds
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataMaskingColumnCount"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataMaskingColumnCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取列加密统计信息
//
// Description:
//
// 按地域、产品和分类分级模板查询总列数、敏感列数、已加密列数和加密失败列数。本接口返回统计卡片数据，不跟随列列表中的实例、库名、表名、列名或模型筛选。
//
// 为兼容历史识别结果，指定模板无首屏结果且未按敏感等级或识别模型过滤时，结果可能回退到内置通用识别结果。要求严格模板归属时，请同时使用该模板内的等级或模型条件核验。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - GetDataMaskingColumnCountRequest
//
// @return GetDataMaskingColumnCountResponse
func (client *Client) GetDataMaskingColumnCount(request *GetDataMaskingColumnCountRequest) (_result *GetDataMaskingColumnCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataMaskingColumnCountResponse{}
	_body, _err := client.GetDataMaskingColumnCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例属性
//
// @param request - GetInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceAttributeResponse
func (client *Client) GetInstanceAttributeWithOptions(request *GetInstanceAttributeRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceAttribute"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例属性
//
// @param request - GetInstanceAttributeRequest
//
// @return GetInstanceAttributeResponse
func (client *Client) GetInstanceAttribute(request *GetInstanceAttributeRequest) (_result *GetInstanceAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceAttributeResponse{}
	_body, _err := client.GetInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据安全中心连接授权的MaxCompute、RDS等数据资产表中列的数据
//
// @param request - ListColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListColumnsResponse
func (client *Client) ListColumnsWithOptions(request *ListColumnsRequest, runtime *dara.RuntimeOptions) (_result *ListColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataAssetSourceId) {
		query["DataAssetSourceId"] = request.DataAssetSourceId
	}

	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceName) {
		query["InstanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListColumns"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据安全中心连接授权的MaxCompute、RDS等数据资产表中列的数据
//
// @param request - ListColumnsRequest
//
// @return ListColumnsResponse
func (client *Client) ListColumns(request *ListColumnsRequest) (_result *ListColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListColumnsResponse{}
	_body, _err := client.ListColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询账号列表
//
// @param request - ListDataAssetAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataAssetAccountsResponse
func (client *Client) ListDataAssetAccountsWithOptions(request *ListDataAssetAccountsRequest, runtime *dara.RuntimeOptions) (_result *ListDataAssetAccountsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		query["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.AuthRole) {
		query["AuthRole"] = request.AuthRole
	}

	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductIds) {
		query["ProductIds"] = request.ProductIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataAssetAccounts"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataAssetAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询账号列表
//
// @param request - ListDataAssetAccountsRequest
//
// @return ListDataAssetAccountsResponse
func (client *Client) ListDataAssetAccounts(request *ListDataAssetAccountsRequest) (_result *ListDataAssetAccountsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataAssetAccountsResponse{}
	_body, _err := client.ListDataAssetAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取列加密列列表
//
// Description:
//
// 分页查询列及其加密状态、敏感等级和识别模型。支持模板、产品、实例、数据库、表、列名、模型和敏感等级组合筛选。创建或关闭规则后使用本接口回读；Processing、Deleting 为中间状态，Failed、DeleteFailed 表示操作失败。RiskLeveLId 的参数名大小写应原样保留。
//
// 按产品查询时使用单个 ProductId 或对应 ProductCode。当前列查询不能依赖 ProductIds 实现多产品筛选；多产品应分别查询。
//
// 为兼容历史识别结果，指定模板无首屏结果且未按敏感等级或识别模型过滤时，结果可能回退到内置通用识别结果。要求严格模板归属时，请同时使用该模板内的等级或模型条件核验。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataMaskingColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataMaskingColumnsResponse
func (client *Client) ListDataMaskingColumnsWithOptions(request *ListDataMaskingColumnsRequest, runtime *dara.RuntimeOptions) (_result *ListDataMaskingColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnName) {
		query["ColumnName"] = request.ColumnName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaskingStatus) {
		query["MaskingStatus"] = request.MaskingStatus
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductIds) {
		query["ProductIds"] = request.ProductIds
	}

	if !dara.IsNil(request.RiskLeveLId) {
		query["RiskLeveLId"] = request.RiskLeveLId
	}

	if !dara.IsNil(request.RiskLevelIds) {
		query["RiskLevelIds"] = request.RiskLevelIds
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateRuleIds) {
		query["TemplateRuleIds"] = request.TemplateRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataMaskingColumns"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataMaskingColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取列加密列列表
//
// Description:
//
// 分页查询列及其加密状态、敏感等级和识别模型。支持模板、产品、实例、数据库、表、列名、模型和敏感等级组合筛选。创建或关闭规则后使用本接口回读；Processing、Deleting 为中间状态，Failed、DeleteFailed 表示操作失败。RiskLeveLId 的参数名大小写应原样保留。
//
// 按产品查询时使用单个 ProductId 或对应 ProductCode。当前列查询不能依赖 ProductIds 实现多产品筛选；多产品应分别查询。
//
// 为兼容历史识别结果，指定模板无首屏结果且未按敏感等级或识别模型过滤时，结果可能回退到内置通用识别结果。要求严格模板归属时，请同时使用该模板内的等级或模型条件核验。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataMaskingColumnsRequest
//
// @return ListDataMaskingColumnsResponse
func (client *Client) ListDataMaskingColumns(request *ListDataMaskingColumnsRequest) (_result *ListDataMaskingColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataMaskingColumnsResponse{}
	_body, _err := client.ListDataMaskingColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询实例支持的列加密算法列表
//
// Description:
//
// 查询指定实例可选的列加密算法及各算法的限制原因。选择算法前检查对应项的 ErrorCode 和 ErrorMessage；不要把所有实例都视为支持同一组算法。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataMaskingEncryptionAlgorithmsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataMaskingEncryptionAlgorithmsResponse
func (client *Client) ListDataMaskingEncryptionAlgorithmsWithOptions(request *ListDataMaskingEncryptionAlgorithmsRequest, runtime *dara.RuntimeOptions) (_result *ListDataMaskingEncryptionAlgorithmsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataMaskingEncryptionAlgorithms"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataMaskingEncryptionAlgorithmsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询实例支持的列加密算法列表
//
// Description:
//
// 查询指定实例可选的列加密算法及各算法的限制原因。选择算法前检查对应项的 ErrorCode 和 ErrorMessage；不要把所有实例都视为支持同一组算法。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataMaskingEncryptionAlgorithmsRequest
//
// @return ListDataMaskingEncryptionAlgorithmsResponse
func (client *Client) ListDataMaskingEncryptionAlgorithms(request *ListDataMaskingEncryptionAlgorithmsRequest) (_result *ListDataMaskingEncryptionAlgorithmsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataMaskingEncryptionAlgorithmsResponse{}
	_body, _err := client.ListDataMaskingEncryptionAlgorithmsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取列加密实例列表
//
// Description:
//
// 分页查询列加密实例，并支持产品、识别模板及库表列等条件。实例级状态筛选与列级状态筛选范围不同：查找全部未加密敏感列时，先枚举目标产品的实例，再使用 ListDataMaskingColumns 按 NotEncrypted 筛选，避免遗漏已经部分加密的实例。
//
// 需要按 EngineType 精确筛选时，应完整分页读取候选后按返回值过滤；部分查询路径不应用此参数。InstanceId 的匹配语义随查询组合变化，精确定位时应核对返回的完整实例标识。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataMaskingInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataMaskingInstancesResponse
func (client *Client) ListDataMaskingInstancesWithOptions(request *ListDataMaskingInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListDataMaskingInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnName) {
		query["ColumnName"] = request.ColumnName
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaskingStatus) {
		query["MaskingStatus"] = request.MaskingStatus
	}

	if !dara.IsNil(request.ModelTagId) {
		query["ModelTagId"] = request.ModelTagId
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductIds) {
		query["ProductIds"] = request.ProductIds
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RiskLevelIds) {
		query["RiskLevelIds"] = request.RiskLevelIds
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateRuleIds) {
		query["TemplateRuleIds"] = request.TemplateRuleIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataMaskingInstances"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataMaskingInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取列加密实例列表
//
// Description:
//
// 分页查询列加密实例，并支持产品、识别模板及库表列等条件。实例级状态筛选与列级状态筛选范围不同：查找全部未加密敏感列时，先枚举目标产品的实例，再使用 ListDataMaskingColumns 按 NotEncrypted 筛选，避免遗漏已经部分加密的实例。
//
// 需要按 EngineType 精确筛选时，应完整分页读取候选后按返回值过滤；部分查询路径不应用此参数。InstanceId 的匹配语义随查询组合变化，精确定位时应核对返回的完整实例标识。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataMaskingInstancesRequest
//
// @return ListDataMaskingInstancesResponse
func (client *Client) ListDataMaskingInstances(request *ListDataMaskingInstancesRequest) (_result *ListDataMaskingInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataMaskingInstancesResponse{}
	_body, _err := client.ListDataMaskingInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询数据源列表
//
// Description:
//
// 分页查询已纳管的数据源。关系型数据库场景通过 ProductCode、InstanceId 定位实例，从 Items[].DbName 选择数据库。本接口查询数据库库存，不使用分类分级模板筛选。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(request *ListDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectStatus) {
		query["ConnectStatus"] = request.ConnectStatus
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataAssetId) {
		query["DataAssetId"] = request.DataAssetId
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IdentifyStatus) {
		query["IdentifyStatus"] = request.IdentifyStatus
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.SourceIp) {
		query["SourceIp"] = request.SourceIp
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据源列表
//
// Description:
//
// 分页查询已纳管的数据源。关系型数据库场景通过 ProductCode、InstanceId 定位实例，从 Items[].DbName 选择数据库。本接口查询数据库库存，不使用分类分级模板筛选。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListDataSourcesRequest
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSources(request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询模板的所有规则
//
// Description:
//
// 查询指定模板内有效的识别模型，返回 RuleList。未指定 TemplateId 时使用当前启用模板。模型标识取 RuleList[].Id，可用于 TemplateRuleIds 筛选。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// @param request - ListIdentifyModelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentifyModelsResponse
func (client *Client) ListIdentifyModelsWithOptions(request *ListIdentifyModelsRequest, runtime *dara.RuntimeOptions) (_result *ListIdentifyModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.FilterAuditModel) {
		query["FilterAuditModel"] = request.FilterAuditModel
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIdentifyModels"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIdentifyModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询模板的所有规则
//
// Description:
//
// 查询指定模板内有效的识别模型，返回 RuleList。未指定 TemplateId 时使用当前启用模板。模型标识取 RuleList[].Id，可用于 TemplateRuleIds 筛选。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// @param request - ListIdentifyModelsRequest
//
// @return ListIdentifyModelsResponse
func (client *Client) ListIdentifyModels(request *ListIdentifyModelsRequest) (_result *ListIdentifyModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIdentifyModelsResponse{}
	_body, _err := client.ListIdentifyModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询可用的KMS主密钥
//
// @param request - ListKmsKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKmsKeysResponse
func (client *Client) ListKmsKeysWithOptions(request *ListKmsKeysRequest, runtime *dara.RuntimeOptions) (_result *ListKmsKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKmsKeys"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKmsKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可用的KMS主密钥
//
// @param request - ListKmsKeysRequest
//
// @return ListKmsKeysResponse
func (client *Client) ListKmsKeys(request *ListKmsKeysRequest) (_result *ListKmsKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListKmsKeysResponse{}
	_body, _err := client.ListKmsKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例内核版本列表
//
// Description:
//
// 查询实例可选择的内核版本。升级时从 KernelVersions[].KernelVersion 选择目标，不应手工构造版本号。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListMiniEngineVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMiniEngineVersionsResponse
func (client *Client) ListMiniEngineVersionsWithOptions(request *ListMiniEngineVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListMiniEngineVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMiniEngineVersions"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMiniEngineVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例内核版本列表
//
// Description:
//
// 查询实例可选择的内核版本。升级时从 KernelVersions[].KernelVersion 选择目标，不应手工构造版本号。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListMiniEngineVersionsRequest
//
// @return ListMiniEngineVersionsResponse
func (client *Client) ListMiniEngineVersions(request *ListMiniEngineVersionsRequest) (_result *ListMiniEngineVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMiniEngineVersionsResponse{}
	_body, _err := client.ListMiniEngineVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询区域列表
//
// Description:
//
// 查询当前账号拥有资产的地域列表，用于选择后续地域化接口的 RegionId。可按是否开启审计、是否开启识别筛选。返回的地域列表不等同于所有云产品支持地域清单。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *dara.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Audited) {
		query["Audited"] = request.Audited
	}

	if !dara.IsNil(request.Identified) {
		query["Identified"] = request.Identified
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRegions"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询区域列表
//
// Description:
//
// 查询当前账号拥有资产的地域列表，用于选择后续地域化接口的 RegionId。可按是否开启审计、是否开启识别筛选。返回的地域列表不等同于所有云产品支持地域清单。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询风险级别列表
//
// @param request - ListRiskLevelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRiskLevelsResponse
func (client *Client) ListRiskLevelsWithOptions(request *ListRiskLevelsRequest, runtime *dara.RuntimeOptions) (_result *ListRiskLevelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRiskLevels"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRiskLevelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询风险级别列表
//
// @param request - ListRiskLevelsRequest
//
// @return ListRiskLevelsResponse
func (client *Client) ListRiskLevels(request *ListRiskLevelsRequest) (_result *ListRiskLevelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRiskLevelsResponse{}
	_body, _err := client.ListRiskLevelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询表识别结果列表
//
// Description:
//
// 分页查询数据库中的数据表及识别结果。关系型数据库场景传入 DataAssetSourceId=实例标识、DataSourceName=数据库名称。PostgreSQL 和 Oracle 的表名可能包含 schema 前缀，后续列查询应原样传递表名。
//
// 为兼容历史识别结果，指定模板无首屏结果且未按敏感等级或识别模型过滤时，结果可能回退到内置通用识别结果。要求严格模板归属时，请同时使用该模板内的等级或模型条件核验。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(request *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.DataAssetSourceId) {
		query["DataAssetSourceId"] = request.DataAssetSourceId
	}

	if !dara.IsNil(request.DataSourceName) {
		query["DataSourceName"] = request.DataSourceName
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Marker) {
		query["Marker"] = request.Marker
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.RiskLevelId) {
		query["RiskLevelId"] = request.RiskLevelId
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表识别结果列表
//
// Description:
//
// 分页查询数据库中的数据表及识别结果。关系型数据库场景传入 DataAssetSourceId=实例标识、DataSourceName=数据库名称。PostgreSQL 和 Oracle 的表名可能包含 schema 前缀，后续列查询应原样传递表名。
//
// 为兼容历史识别结果，指定模板无首屏结果且未按敏感等级或识别模型过滤时，结果可能回退到内置通用识别结果。要求严格模板归属时，请同时使用该模板内的等级或模型条件核验。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// @param request - ListTablesRequest
//
// @return ListTablesResponse
func (client *Client) ListTables(request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询识别模版列表
//
// @param request - ListTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.FeatureType) {
		query["FeatureType"] = request.FeatureType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UsageScenario) {
		query["UsageScenario"] = request.UsageScenario
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询识别模版列表
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提供数据目录、总览页面的敏感数据汇总
//
// Description:
//
// 查询总览、数据目录或按地域汇总的敏感数据统计。CountType=41 返回总览统计，42 返回指定产品的数据目录统计，43 返回按地域和模板聚合的数据。列加密控制台使用 CountType=43 与 ProductCodeList 获取地域和模板候选。统计数据可能来自已生成的汇总结果，不代表刚发起的同步或加密操作已经完成。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// @param request - ListTotalSensitiveInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTotalSensitiveInfoResponse
func (client *Client) ListTotalSensitiveInfoWithOptions(request *ListTotalSensitiveInfoRequest, runtime *dara.RuntimeOptions) (_result *ListTotalSensitiveInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CountType) {
		query["CountType"] = request.CountType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductCodeList) {
		query["ProductCodeList"] = request.ProductCodeList
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTotalSensitiveInfo"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTotalSensitiveInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提供数据目录、总览页面的敏感数据汇总
//
// Description:
//
// 查询总览、数据目录或按地域汇总的敏感数据统计。CountType=41 返回总览统计，42 返回指定产品的数据目录统计，43 返回按地域和模板聚合的数据。列加密控制台使用 CountType=43 与 ProductCodeList 获取地域和模板候选。统计数据可能来自已生成的汇总结果，不代表刚发起的同步或加密操作已经完成。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// @param request - ListTotalSensitiveInfoRequest
//
// @return ListTotalSensitiveInfoResponse
func (client *Client) ListTotalSensitiveInfo(request *ListTotalSensitiveInfoRequest) (_result *ListTotalSensitiveInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTotalSensitiveInfoResponse{}
	_body, _err := client.ListTotalSensitiveInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资产同步
//
// Description:
//
// 触发当前地域内指定产品的资产同步，作用范围是产品和地域，不限于某个实例。请求成功仅表示已受理；随后通过实例、数据库、表和列列表核对资产变化。该接口不返回可供轮询的公开任务标识。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - SyncDataAssetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDataAssetsResponse
func (client *Client) SyncDataAssetsWithOptions(request *SyncDataAssetsRequest, runtime *dara.RuntimeOptions) (_result *SyncDataAssetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDataAssets"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDataAssetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资产同步
//
// Description:
//
// 触发当前地域内指定产品的资产同步，作用范围是产品和地域，不限于某个实例。请求成功仅表示已受理；随后通过实例、数据库、表和列列表核对资产变化。该接口不返回可供轮询的公开任务标识。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - SyncDataAssetsRequest
//
// @return SyncDataAssetsResponse
func (client *Client) SyncDataAssets(request *SyncDataAssetsRequest) (_result *SyncDataAssetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncDataAssetsResponse{}
	_body, _err := client.SyncDataAssetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新列加密算法
//
// Description:
//
// 修改实例级列加密算法及密钥配置，影响该实例的列加密配置。当前实现支持 RDS MySQL 和 PolarDB-X 2.0；先查询实例支持算法和密钥。该接口没有 EncryptionKeyMode 参数，应按 EncryptionKeyId 的使用条件配置。请求成功后回读实例加密配置和列状态，不能仅凭 RequestId 判断完成。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - UpdateDataMaskingEncryptionAlgorithmRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataMaskingEncryptionAlgorithmResponse
func (client *Client) UpdateDataMaskingEncryptionAlgorithmWithOptions(request *UpdateDataMaskingEncryptionAlgorithmRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataMaskingEncryptionAlgorithmResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EncryptionAlgorithm) {
		query["EncryptionAlgorithm"] = request.EncryptionAlgorithm
	}

	if !dara.IsNil(request.EncryptionKeyId) {
		query["EncryptionKeyId"] = request.EncryptionKeyId
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataMaskingEncryptionAlgorithm"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataMaskingEncryptionAlgorithmResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新列加密算法
//
// Description:
//
// 修改实例级列加密算法及密钥配置，影响该实例的列加密配置。当前实现支持 RDS MySQL 和 PolarDB-X 2.0；先查询实例支持算法和密钥。该接口没有 EncryptionKeyMode 参数，应按 EncryptionKeyId 的使用条件配置。请求成功后回读实例加密配置和列状态，不能仅凭 RequestId 判断完成。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - UpdateDataMaskingEncryptionAlgorithmRequest
//
// @return UpdateDataMaskingEncryptionAlgorithmResponse
func (client *Client) UpdateDataMaskingEncryptionAlgorithm(request *UpdateDataMaskingEncryptionAlgorithmRequest) (_result *UpdateDataMaskingEncryptionAlgorithmResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataMaskingEncryptionAlgorithmResponse{}
	_body, _err := client.UpdateDataMaskingEncryptionAlgorithmWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量修改账号权限
//
// Description:
//
// 为一批数据库账号设置同一种列加密访问权限。UserList 可包含同一产品、同一地域下不同实例的账号，每项必须携带 InstanceId 和 AccountId。fullAccess 允许明文访问，restrictedAccess 允许受限访问，noneAccess 禁止解密且仅支持特定引擎和密钥模式。仅 fullAccess 可配置有效期。请求成功后使用 ListDataAssetAccounts 回读。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param tmpReq - UpdateDataMaskingUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataMaskingUsersResponse
func (client *Client) UpdateDataMaskingUsersWithOptions(tmpReq *UpdateDataMaskingUsersRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataMaskingUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataMaskingUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserList) {
		request.UserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserList, dara.String("UserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthRole) {
		query["AuthRole"] = request.AuthRole
	}

	if !dara.IsNil(request.ExpireTime) {
		query["ExpireTime"] = request.ExpireTime
	}

	if !dara.IsNil(request.ExpireTimeOperation) {
		query["ExpireTimeOperation"] = request.ExpireTimeOperation
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.UserListShrink) {
		query["UserList"] = request.UserListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataMaskingUsers"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataMaskingUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改账号权限
//
// Description:
//
// 为一批数据库账号设置同一种列加密访问权限。UserList 可包含同一产品、同一地域下不同实例的账号，每项必须携带 InstanceId 和 AccountId。fullAccess 允许明文访问，restrictedAccess 允许受限访问，noneAccess 禁止解密且仅支持特定引擎和密钥模式。仅 fullAccess 可配置有效期。请求成功后使用 ListDataAssetAccounts 回读。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - UpdateDataMaskingUsersRequest
//
// @return UpdateDataMaskingUsersResponse
func (client *Client) UpdateDataMaskingUsers(request *UpdateDataMaskingUsersRequest) (_result *UpdateDataMaskingUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataMaskingUsersResponse{}
	_body, _err := client.UpdateDataMaskingUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 升级内核版本
//
// Description:
//
// 按所选时间升级指定实例的数据库内核。先调用 ListMiniEngineVersions 选择版本，并检查实例当前状态和维护窗口。SpecifyTime 需要提供未来的 SwitchTime，PolarDB-X 2.0 不支持该时间选项。升级为异步操作，回读 GetInstanceAttribute 的 CurrentKernelVersion 核对结果。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - UpgradeKernelVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeKernelVersionResponse
func (client *Client) UpgradeKernelVersionWithOptions(request *UpgradeKernelVersionRequest, runtime *dara.RuntimeOptions) (_result *UpgradeKernelVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.KernelVersion) {
		query["KernelVersion"] = request.KernelVersion
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.ProductCode) {
		query["ProductCode"] = request.ProductCode
	}

	if !dara.IsNil(request.ProductId) {
		query["ProductId"] = request.ProductId
	}

	if !dara.IsNil(request.SwitchTime) {
		query["SwitchTime"] = request.SwitchTime
	}

	if !dara.IsNil(request.UpgradeTime) {
		query["UpgradeTime"] = request.UpgradeTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeKernelVersion"),
		Version:     dara.String("2026-01-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeKernelVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 升级内核版本
//
// Description:
//
// 按所选时间升级指定实例的数据库内核。先调用 ListMiniEngineVersions 选择版本，并检查实例当前状态和维护窗口。SpecifyTime 需要提供未来的 SwitchTime，PolarDB-X 2.0 不支持该时间选项。升级为异步操作，回读 GetInstanceAttribute 的 CurrentKernelVersion 核对结果。
//
// 参数示例仅用于说明格式。调用时请替换为当前账号查询得到的地域、资源标识和配置值。
//
// 使用目标资产所在地域的服务接入点，并设置公共参数 RegionId，例如 cn-zhangjiakou。产品编码和实例标识必须与目标资产一致。
//
// 本接口仅返回 RequestId。请求受理不等于业务操作完成，应按接口说明回读状态。
//
// @param request - UpgradeKernelVersionRequest
//
// @return UpgradeKernelVersionResponse
func (client *Client) UpgradeKernelVersion(request *UpgradeKernelVersionRequest) (_result *UpgradeKernelVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeKernelVersionResponse{}
	_body, _err := client.UpgradeKernelVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
