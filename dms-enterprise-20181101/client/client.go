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
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dms-enterprise"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Adds instances, databases, and tables to the specified permission template.
//
// @param tmpReq - AddAuthorityTemplateItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAuthorityTemplateItemsResponse
func (client *Client) AddAuthorityTemplateItemsWithOptions(tmpReq *AddAuthorityTemplateItemsRequest, runtime *dara.RuntimeOptions) (_result *AddAuthorityTemplateItemsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddAuthorityTemplateItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Items) {
		request.ItemsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Items, dara.String("Items"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ItemsShrink) {
		query["Items"] = request.ItemsShrink
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAuthorityTemplateItems"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAuthorityTemplateItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds instances, databases, and tables to the specified permission template.
//
// @param request - AddAuthorityTemplateItemsRequest
//
// @return AddAuthorityTemplateItemsResponse
func (client *Client) AddAuthorityTemplateItems(request *AddAuthorityTemplateItemsRequest) (_result *AddAuthorityTemplateItemsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAuthorityTemplateItemsResponse{}
	_body, _err := client.AddAuthorityTemplateItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a masking rule.
//
// @param request - AddDesensitizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDesensitizationRuleResponse
func (client *Client) AddDesensitizationRuleWithOptions(request *AddDesensitizationRuleRequest, runtime *dara.RuntimeOptions) (_result *AddDesensitizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionType) {
		query["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.RuleDescription) {
		query["RuleDescription"] = request.RuleDescription
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.FunctionParams) {
		bodyFlat["FunctionParams"] = request.FunctionParams
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDesensitizationRule"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDesensitizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a masking rule.
//
// @param request - AddDesensitizationRuleRequest
//
// @return AddDesensitizationRuleResponse
func (client *Client) AddDesensitizationRule(request *AddDesensitizationRuleRequest) (_result *AddDesensitizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddDesensitizationRuleResponse{}
	_body, _err := client.AddDesensitizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a database instance to Data Management (DMS) and enable security hosting for the instance.
//
// Description:
//
// You must be a database administrator (DBA) or a DMS administrator. For more information, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - AddInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddInstanceResponse
func (client *Client) AddInstanceWithOptions(request *AddInstanceRequest, runtime *dara.RuntimeOptions) (_result *AddInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataLinkName) {
		query["DataLinkName"] = request.DataLinkName
	}

	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.DatabaseUser) {
		query["DatabaseUser"] = request.DatabaseUser
	}

	if !dara.IsNil(request.DbaId) {
		query["DbaId"] = request.DbaId
	}

	if !dara.IsNil(request.DdlOnline) {
		query["DdlOnline"] = request.DdlOnline
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.EcsRegion) {
		query["EcsRegion"] = request.EcsRegion
	}

	if !dara.IsNil(request.EnableSellCommon) {
		query["EnableSellCommon"] = request.EnableSellCommon
	}

	if !dara.IsNil(request.EnableSellSitd) {
		query["EnableSellSitd"] = request.EnableSellSitd
	}

	if !dara.IsNil(request.EnableSellStable) {
		query["EnableSellStable"] = request.EnableSellStable
	}

	if !dara.IsNil(request.EnableSellTrust) {
		query["EnableSellTrust"] = request.EnableSellTrust
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ExportTimeout) {
		query["ExportTimeout"] = request.ExportTimeout
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceSource) {
		query["InstanceSource"] = request.InstanceSource
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.QueryTimeout) {
		query["QueryTimeout"] = request.QueryTimeout
	}

	if !dara.IsNil(request.SafeRule) {
		query["SafeRule"] = request.SafeRule
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.SkipTest) {
		query["SkipTest"] = request.SkipTest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UseDsql) {
		query["UseDsql"] = request.UseDsql
	}

	if !dara.IsNil(request.UseSsl) {
		query["UseSsl"] = request.UseSsl
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a database instance to Data Management (DMS) and enable security hosting for the instance.
//
// Description:
//
// You must be a database administrator (DBA) or a DMS administrator. For more information, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - AddInstanceRequest
//
// @return AddInstanceResponse
func (client *Client) AddInstance(request *AddInstanceRequest) (_result *AddInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddInstanceResponse{}
	_body, _err := client.AddInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a workspace administrator, a workspace member, or a task flow developer in Data Management (DMS).
//
// Description:
//
// You must call this operation as a DMS administrator, a database administrator (DBA), or a workspace administrator.
//
// Usage notes:
//
//   - Before you call this operation to add a user as a task flow developer, make sure that you have added the user as a workspace member.
//
//   - You cannot call this operation to transfer the ownership of a task flow. To transfer the ownership of a task flow, call the [ChangLhDagOwner](https://help.aliyun.com/document_detail/424761.html) operation.
//
//   - For more information about workspace roles and permissions, see [Manage permissions on a workspace](https://help.aliyun.com/document_detail/410893.html).
//
// @param tmpReq - AddLhMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLhMembersResponse
func (client *Client) AddLhMembersWithOptions(tmpReq *AddLhMembersRequest, runtime *dara.RuntimeOptions) (_result *AddLhMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddLhMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Members) {
		request.MembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Members, dara.String("Members"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MembersShrink) {
		query["Members"] = request.MembersShrink
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLhMembers"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLhMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a workspace administrator, a workspace member, or a task flow developer in Data Management (DMS).
//
// Description:
//
// You must call this operation as a DMS administrator, a database administrator (DBA), or a workspace administrator.
//
// Usage notes:
//
//   - Before you call this operation to add a user as a task flow developer, make sure that you have added the user as a workspace member.
//
//   - You cannot call this operation to transfer the ownership of a task flow. To transfer the ownership of a task flow, call the [ChangLhDagOwner](https://help.aliyun.com/document_detail/424761.html) operation.
//
//   - For more information about workspace roles and permissions, see [Manage permissions on a workspace](https://help.aliyun.com/document_detail/410893.html).
//
// @param request - AddLhMembersRequest
//
// @return AddLhMembersResponse
func (client *Client) AddLhMembers(request *AddLhMembersRequest) (_result *AddLhMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddLhMembersResponse{}
	_body, _err := client.AddLhMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a routing algorithm to a logical table.
//
// @param request - AddLogicTableRouteConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddLogicTableRouteConfigResponse
func (client *Client) AddLogicTableRouteConfigWithOptions(request *AddLogicTableRouteConfigRequest, runtime *dara.RuntimeOptions) (_result *AddLogicTableRouteConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RouteExpr) {
		query["RouteExpr"] = request.RouteExpr
	}

	if !dara.IsNil(request.RouteKey) {
		query["RouteKey"] = request.RouteKey
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddLogicTableRouteConfig"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddLogicTableRouteConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a routing algorithm to a logical table.
//
// @param request - AddLogicTableRouteConfigRequest
//
// @return AddLogicTableRouteConfigResponse
func (client *Client) AddLogicTableRouteConfig(request *AddLogicTableRouteConfigRequest) (_result *AddLogicTableRouteConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddLogicTableRouteConfigResponse{}
	_body, _err := client.AddLogicTableRouteConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a table to an asset category.
//
// @param request - AddTableToCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTableToCategoryResponse
func (client *Client) AddTableToCategoryWithOptions(request *AddTableToCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddTableToCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableSchemaName) {
		query["TableSchemaName"] = request.TableSchemaName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTableToCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTableToCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a table to an asset category.
//
// @param request - AddTableToCategoryRequest
//
// @return AddTableToCategoryResponse
func (client *Client) AddTableToCategory(request *AddTableToCategoryRequest) (_result *AddTableToCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTableToCategoryResponse{}
	_body, _err := client.AddTableToCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds directed edges for an existing task node.
//
// Description:
//
// When you add directed edges for a task node, take note of the following limits:
//
// 1. The endpoints of the specified edge exist in the Directed Acyclic Graph (DAG) of the task flow specified by DagId.
//
// 2. After a backward edge is added, the DAG does not contain loops.
//
// @param tmpReq - AddTaskFlowEdgesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTaskFlowEdgesResponse
func (client *Client) AddTaskFlowEdgesWithOptions(tmpReq *AddTaskFlowEdgesRequest, runtime *dara.RuntimeOptions) (_result *AddTaskFlowEdgesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddTaskFlowEdgesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Edges) {
		request.EdgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Edges, dara.String("Edges"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.EdgesShrink) {
		query["Edges"] = request.EdgesShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTaskFlowEdges"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTaskFlowEdgesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds directed edges for an existing task node.
//
// Description:
//
// When you add directed edges for a task node, take note of the following limits:
//
// 1. The endpoints of the specified edge exist in the Directed Acyclic Graph (DAG) of the task flow specified by DagId.
//
// 2. After a backward edge is added, the DAG does not contain loops.
//
// @param request - AddTaskFlowEdgesRequest
//
// @return AddTaskFlowEdgesResponse
func (client *Client) AddTaskFlowEdges(request *AddTaskFlowEdgesRequest) (_result *AddTaskFlowEdgesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTaskFlowEdgesResponse{}
	_body, _err := client.AddTaskFlowEdgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Analyzes the lineage (dependencies and influence) between tables and between fields in SQL statements.
//
// Description:
//
// The following conditions must be met before you call this API operation.
//
//   - The database instance is of one of the following types: ApsaraDB RDS for MySQL, PolarDB for MySQL, AnalyticDB for MySQL, ApsaraDB RDS for PostgreSQL, PolarDB for PostgreSQL, AnalyticDB for PostgreSQL, Oracle, and openGauss.
//
//   - A database instance is managed in Security Collaboration mode. For more information about control modes, see [Control modes](https://help.aliyun.com/document_detail/151629.html).
//
// @param request - AnalyzeSQLLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeSQLLineageResponse
func (client *Client) AnalyzeSQLLineageWithOptions(request *AnalyzeSQLLineageRequest, runtime *dara.RuntimeOptions) (_result *AnalyzeSQLLineageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.SqlContent) {
		query["SqlContent"] = request.SqlContent
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeSQLLineage"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeSQLLineageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Analyzes the lineage (dependencies and influence) between tables and between fields in SQL statements.
//
// Description:
//
// The following conditions must be met before you call this API operation.
//
//   - The database instance is of one of the following types: ApsaraDB RDS for MySQL, PolarDB for MySQL, AnalyticDB for MySQL, ApsaraDB RDS for PostgreSQL, PolarDB for PostgreSQL, AnalyticDB for PostgreSQL, Oracle, and openGauss.
//
//   - A database instance is managed in Security Collaboration mode. For more information about control modes, see [Control modes](https://help.aliyun.com/document_detail/151629.html).
//
// @param request - AnalyzeSQLLineageRequest
//
// @return AnalyzeSQLLineageResponse
func (client *Client) AnalyzeSQLLineage(request *AnalyzeSQLLineageRequest) (_result *AnalyzeSQLLineageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AnalyzeSQLLineageResponse{}
	_body, _err := client.AnalyzeSQLLineageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Answers the syntax questions of the corresponding engine according to the specified database ID.
//
// Description:
//
// You can call this operation to answer the syntax questions of the corresponding engine according to the specified database ID.
//
// @param request - AnswerSqlSyntaxByMetaAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnswerSqlSyntaxByMetaAgentResponse
func (client *Client) AnswerSqlSyntaxByMetaAgentWithOptions(request *AnswerSqlSyntaxByMetaAgentRequest, runtime *dara.RuntimeOptions) (_result *AnswerSqlSyntaxByMetaAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnswerSqlSyntaxByMetaAgent"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AnswerSqlSyntaxByMetaAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Answers the syntax questions of the corresponding engine according to the specified database ID.
//
// Description:
//
// You can call this operation to answer the syntax questions of the corresponding engine according to the specified database ID.
//
// @param request - AnswerSqlSyntaxByMetaAgentRequest
//
// @return AnswerSqlSyntaxByMetaAgentResponse
func (client *Client) AnswerSqlSyntaxByMetaAgent(request *AnswerSqlSyntaxByMetaAgentRequest) (_result *AnswerSqlSyntaxByMetaAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AnswerSqlSyntaxByMetaAgentResponse{}
	_body, _err := client.AnswerSqlSyntaxByMetaAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reviews a ticket.
//
// @param request - ApproveOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveOrderResponse
func (client *Client) ApproveOrderWithOptions(request *ApproveOrderRequest, runtime *dara.RuntimeOptions) (_result *ApproveOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApprovalNodeId) {
		query["ApprovalNodeId"] = request.ApprovalNodeId
	}

	if !dara.IsNil(request.ApprovalNodePos) {
		query["ApprovalNodePos"] = request.ApprovalNodePos
	}

	if !dara.IsNil(request.ApprovalType) {
		query["ApprovalType"] = request.ApprovalType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.NewApprover) {
		query["NewApprover"] = request.NewApprover
	}

	if !dara.IsNil(request.NewApproverList) {
		query["NewApproverList"] = request.NewApproverList
	}

	if !dara.IsNil(request.OldApprover) {
		query["OldApprover"] = request.OldApprover
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		query["WorkflowInstanceId"] = request.WorkflowInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApproveOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reviews a ticket.
//
// @param request - ApproveOrderRequest
//
// @return ApproveOrderResponse
func (client *Client) ApproveOrder(request *ApproveOrderRequest) (_result *ApproveOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApproveOrderResponse{}
	_body, _err := client.ApproveOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Backfills data for task orchestration.
//
// Description:
//
// During a data backfill, task flows are run in sequence based on their dates. You can specify whether task flows are run in chronological or reverse chronological order. After the data backfill is complete, you can specify a date or date range, and a node range to run task flows.
//
// @param tmpReq - BackFillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BackFillResponse
func (client *Client) BackFillWithOptions(tmpReq *BackFillRequest, runtime *dara.RuntimeOptions) (_result *BackFillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BackFillShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FilterNodeIds) {
		request.FilterNodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterNodeIds, dara.String("FilterNodeIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StartNodeIds) {
		request.StartNodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StartNodeIds, dara.String("StartNodeIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		query["Asc"] = request.Asc
	}

	if !dara.IsNil(request.BackFillDate) {
		query["BackFillDate"] = request.BackFillDate
	}

	if !dara.IsNil(request.BackFillDateBegin) {
		query["BackFillDateBegin"] = request.BackFillDateBegin
	}

	if !dara.IsNil(request.BackFillDateEnd) {
		query["BackFillDateEnd"] = request.BackFillDateEnd
	}

	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.FilterNodeIdsShrink) {
		query["FilterNodeIds"] = request.FilterNodeIdsShrink
	}

	if !dara.IsNil(request.HistoryDagId) {
		query["HistoryDagId"] = request.HistoryDagId
	}

	if !dara.IsNil(request.Interval) {
		query["Interval"] = request.Interval
	}

	if !dara.IsNil(request.IsTriggerSubTree) {
		query["IsTriggerSubTree"] = request.IsTriggerSubTree
	}

	if !dara.IsNil(request.StartNodeIdsShrink) {
		query["StartNodeIds"] = request.StartNodeIdsShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BackFill"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BackFillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Backfills data for task orchestration.
//
// Description:
//
// During a data backfill, task flows are run in sequence based on their dates. You can specify whether task flows are run in chronological or reverse chronological order. After the data backfill is complete, you can specify a date or date range, and a node range to run task flows.
//
// @param request - BackFillRequest
//
// @return BackFillResponse
func (client *Client) BackFill(request *BackFillRequest) (_result *BackFillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BackFillResponse{}
	_body, _err := client.BackFillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple partitions for a table in a data lakehouse at a time.
//
// @param tmpReq - BatchCreateDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateDataLakePartitionsResponse
func (client *Client) BatchCreateDataLakePartitionsWithOptions(tmpReq *BatchCreateDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchCreateDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchCreateDataLakePartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInputs) {
		request.PartitionInputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInputs, dara.String("PartitionInputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfNotExists) {
		query["IfNotExists"] = request.IfNotExists
	}

	if !dara.IsNil(request.NeedResult) {
		query["NeedResult"] = request.NeedResult
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputsShrink) {
		body["PartitionInputs"] = request.PartitionInputsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateDataLakePartitions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateDataLakePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple partitions for a table in a data lakehouse at a time.
//
// @param request - BatchCreateDataLakePartitionsRequest
//
// @return BatchCreateDataLakePartitionsResponse
func (client *Client) BatchCreateDataLakePartitions(request *BatchCreateDataLakePartitionsRequest) (_result *BatchCreateDataLakePartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchCreateDataLakePartitionsResponse{}
	_body, _err := client.BatchCreateDataLakePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple partitions of a table in a data lakehouse at a time.
//
// @param request - BatchDeleteDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeleteDataLakePartitionsResponse
func (client *Client) BatchDeleteDataLakePartitionsWithOptions(request *BatchDeleteDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchDeleteDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfExists) {
		query["IfExists"] = request.IfExists
	}

	if !dara.IsNil(request.PartitionValuesList) {
		query["PartitionValuesList"] = request.PartitionValuesList
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeleteDataLakePartitions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeleteDataLakePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple partitions of a table in a data lakehouse at a time.
//
// @param request - BatchDeleteDataLakePartitionsRequest
//
// @return BatchDeleteDataLakePartitionsResponse
func (client *Client) BatchDeleteDataLakePartitions(request *BatchDeleteDataLakePartitionsRequest) (_result *BatchDeleteDataLakePartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchDeleteDataLakePartitionsResponse{}
	_body, _err := client.BatchDeleteDataLakePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates multiple partitions of a table in a data lakehouse at a time.
//
// @param tmpReq - BatchUpdateDataLakePartitionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateDataLakePartitionsResponse
func (client *Client) BatchUpdateDataLakePartitionsWithOptions(tmpReq *BatchUpdateDataLakePartitionsRequest, runtime *dara.RuntimeOptions) (_result *BatchUpdateDataLakePartitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchUpdateDataLakePartitionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInputs) {
		request.PartitionInputsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInputs, dara.String("PartitionInputs"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputsShrink) {
		body["PartitionInputs"] = request.PartitionInputsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateDataLakePartitions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateDataLakePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates multiple partitions of a table in a data lakehouse at a time.
//
// @param request - BatchUpdateDataLakePartitionsRequest
//
// @return BatchUpdateDataLakePartitionsResponse
func (client *Client) BatchUpdateDataLakePartitions(request *BatchUpdateDataLakePartitionsRequest) (_result *BatchUpdateDataLakePartitionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchUpdateDataLakePartitionsResponse{}
	_body, _err := client.BatchUpdateDataLakePartitionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purchases a pay-as-you-go Data Management (DMS) resource.
//
// @param request - BuyPayAsYouGoOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuyPayAsYouGoOrderResponse
func (client *Client) BuyPayAsYouGoOrderWithOptions(request *BuyPayAsYouGoOrderRequest, runtime *dara.RuntimeOptions) (_result *BuyPayAsYouGoOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CommodityType) {
		query["CommodityType"] = request.CommodityType
	}

	if !dara.IsNil(request.InsNum) {
		query["InsNum"] = request.InsNum
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.VersionType) {
		query["VersionType"] = request.VersionType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuyPayAsYouGoOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BuyPayAsYouGoOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purchases a pay-as-you-go Data Management (DMS) resource.
//
// @param request - BuyPayAsYouGoOrderRequest
//
// @return BuyPayAsYouGoOrderResponse
func (client *Client) BuyPayAsYouGoOrder(request *BuyPayAsYouGoOrderRequest) (_result *BuyPayAsYouGoOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BuyPayAsYouGoOrderResponse{}
	_body, _err := client.BuyPayAsYouGoOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adjusts the sensitivity level of one or more fields.
//
// @param request - ChangeColumnSecLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeColumnSecLevelResponse
func (client *Client) ChangeColumnSecLevelWithOptions(request *ChangeColumnSecLevelRequest, runtime *dara.RuntimeOptions) (_result *ChangeColumnSecLevelResponse, _err error) {
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

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.IsLogic) {
		query["IsLogic"] = request.IsLogic
	}

	if !dara.IsNil(request.NewLevel) {
		query["NewLevel"] = request.NewLevel
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeColumnSecLevel"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeColumnSecLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adjusts the sensitivity level of one or more fields.
//
// @param request - ChangeColumnSecLevelRequest
//
// @return ChangeColumnSecLevelResponse
func (client *Client) ChangeColumnSecLevel(request *ChangeColumnSecLevelRequest) (_result *ChangeColumnSecLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeColumnSecLevelResponse{}
	_body, _err := client.ChangeColumnSecLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the security level of a column.
//
// @param request - ChangeColumnSecurityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeColumnSecurityLevelResponse
func (client *Client) ChangeColumnSecurityLevelWithOptions(request *ChangeColumnSecurityLevelRequest, runtime *dara.RuntimeOptions) (_result *ChangeColumnSecurityLevelResponse, _err error) {
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

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.IsLogic) {
		query["IsLogic"] = request.IsLogic
	}

	if !dara.IsNil(request.NewSensitivityLevel) {
		query["NewSensitivityLevel"] = request.NewSensitivityLevel
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeColumnSecurityLevel"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeColumnSecurityLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the security level of a column.
//
// @param request - ChangeColumnSecurityLevelRequest
//
// @return ChangeColumnSecurityLevelResponse
func (client *Client) ChangeColumnSecurityLevel(request *ChangeColumnSecurityLevelRequest) (_result *ChangeColumnSecurityLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeColumnSecurityLevelResponse{}
	_body, _err := client.ChangeColumnSecurityLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移交数仓开发任务流
//
// Description:
//
// Usage notes:
//
//   - If you call this operation to transfer the ownership of a published task flow, the ownership transfer does not take effect.
//
//   - You can call the [ReDeployLhDagVersion](https://help.aliyun.com/document_detail/424712.html) operation to redeploy a published version of a task flow.
//
// @param request - ChangeLhDagOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeLhDagOwnerResponse
func (client *Client) ChangeLhDagOwnerWithOptions(request *ChangeLhDagOwnerRequest, runtime *dara.RuntimeOptions) (_result *ChangeLhDagOwnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.OwnerUserId) {
		query["OwnerUserId"] = request.OwnerUserId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeLhDagOwner"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeLhDagOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移交数仓开发任务流
//
// Description:
//
// Usage notes:
//
//   - If you call this operation to transfer the ownership of a published task flow, the ownership transfer does not take effect.
//
//   - You can call the [ReDeployLhDagVersion](https://help.aliyun.com/document_detail/424712.html) operation to redeploy a published version of a task flow.
//
// @param request - ChangeLhDagOwnerRequest
//
// @return ChangeLhDagOwnerResponse
func (client *Client) ChangeLhDagOwner(request *ChangeLhDagOwnerRequest) (_result *ChangeLhDagOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeLhDagOwnerResponse{}
	_body, _err := client.ChangeLhDagOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 带有DMS脱敏能力的聊天API接口
//
// Description:
//
// 带有DMS脱敏能力的聊天API接口
//
// @param tmpReq - ChatWithDesensitizeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatWithDesensitizeResponse
func (client *Client) ChatWithDesensitizeWithOptions(tmpReq *ChatWithDesensitizeRequest, runtime *dara.RuntimeOptions) (_result *ChatWithDesensitizeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChatWithDesensitizeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Messages) {
		request.MessagesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Messages, dara.String("Messages"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Stop) {
		request.StopShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Stop, dara.String("Stop"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DesensitizationRule) {
		query["DesensitizationRule"] = request.DesensitizationRule
	}

	if !dara.IsNil(request.EnableThinking) {
		query["EnableThinking"] = request.EnableThinking
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.MaxTokens) {
		query["MaxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.NeedDesensitization) {
		query["NeedDesensitization"] = request.NeedDesensitization
	}

	if !dara.IsNil(request.PresencePenalty) {
		query["PresencePenalty"] = request.PresencePenalty
	}

	if !dara.IsNil(request.ResponseFormat) {
		query["ResponseFormat"] = request.ResponseFormat
	}

	if !dara.IsNil(request.Seed) {
		query["Seed"] = request.Seed
	}

	if !dara.IsNil(request.StopShrink) {
		query["Stop"] = request.StopShrink
	}

	if !dara.IsNil(request.Temperature) {
		query["Temperature"] = request.Temperature
	}

	if !dara.IsNil(request.ThinkingBudget) {
		query["ThinkingBudget"] = request.ThinkingBudget
	}

	if !dara.IsNil(request.TopK) {
		query["TopK"] = request.TopK
	}

	if !dara.IsNil(request.TopLogprobs) {
		query["TopLogprobs"] = request.TopLogprobs
	}

	if !dara.IsNil(request.TopP) {
		query["TopP"] = request.TopP
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MessagesShrink) {
		body["Messages"] = request.MessagesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatWithDesensitize"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatWithDesensitizeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 带有DMS脱敏能力的聊天API接口
//
// Description:
//
// 带有DMS脱敏能力的聊天API接口
//
// @param request - ChatWithDesensitizeRequest
//
// @return ChatWithDesensitizeResponse
func (client *Client) ChatWithDesensitize(request *ChatWithDesensitizeRequest) (_result *ChatWithDesensitizeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChatWithDesensitizeResponse{}
	_body, _err := client.ChatWithDesensitizeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Closes a ticket.
//
// @param request - CloseOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseOrderResponse
func (client *Client) CloseOrderWithOptions(request *CloseOrderRequest, runtime *dara.RuntimeOptions) (_result *CloseOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CloseReason) {
		query["CloseReason"] = request.CloseReason
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Closes a ticket.
//
// @param request - CloseOrderRequest
//
// @return CloseOrderResponse
func (client *Client) CloseOrder(request *CloseOrderRequest) (_result *CloseOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloseOrderResponse{}
	_body, _err := client.CloseOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attaches a system policy or custom policy to a DMS user or custom role.
//
// @param request - CreateAbacAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAbacAuthorizationResponse
func (client *Client) CreateAbacAuthorizationWithOptions(request *CreateAbacAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *CreateAbacAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.RoleId) {
		query["RoleId"] = request.RoleId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAbacAuthorization"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAbacAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attaches a system policy or custom policy to a DMS user or custom role.
//
// @param request - CreateAbacAuthorizationRequest
//
// @return CreateAbacAuthorizationResponse
func (client *Client) CreateAbacAuthorization(request *CreateAbacAuthorizationRequest) (_result *CreateAbacAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAbacAuthorizationResponse{}
	_body, _err := client.CreateAbacAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a policy to manage the permissions on DMS features and data resources managed in DMS in a fine-grained manner.
//
// Description:
//
// Currently, this feature is in the phased release phase. Only users who have phased out can use this feature and related APIs. For more information about the policy feature, see [Manage policies](https://help.aliyun.com/document_detail/2848565.html).
//
// @param request - CreateAbacPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAbacPolicyResponse
func (client *Client) CreateAbacPolicyWithOptions(request *CreateAbacPolicyRequest, runtime *dara.RuntimeOptions) (_result *CreateAbacPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbacPolicyContent) {
		query["AbacPolicyContent"] = request.AbacPolicyContent
	}

	if !dara.IsNil(request.AbacPolicyDesc) {
		query["AbacPolicyDesc"] = request.AbacPolicyDesc
	}

	if !dara.IsNil(request.AbacPolicyName) {
		query["AbacPolicyName"] = request.AbacPolicyName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAbacPolicy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAbacPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a policy to manage the permissions on DMS features and data resources managed in DMS in a fine-grained manner.
//
// Description:
//
// Currently, this feature is in the phased release phase. Only users who have phased out can use this feature and related APIs. For more information about the policy feature, see [Manage policies](https://help.aliyun.com/document_detail/2848565.html).
//
// @param request - CreateAbacPolicyRequest
//
// @return CreateAbacPolicyResponse
func (client *Client) CreateAbacPolicy(request *CreateAbacPolicyRequest) (_result *CreateAbacPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAbacPolicyResponse{}
	_body, _err := client.CreateAbacPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Creates a permission template
//
// Description:
//
// You are a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - CreateAuthorityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAuthorityTemplateResponse
func (client *Client) CreateAuthorityTemplateWithOptions(request *CreateAuthorityTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateAuthorityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAuthorityTemplate"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAuthorityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Creates a permission template
//
// Description:
//
// You are a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - CreateAuthorityTemplateRequest
//
// @return CreateAuthorityTemplateResponse
func (client *Client) CreateAuthorityTemplate(request *CreateAuthorityTemplateRequest) (_result *CreateAuthorityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAuthorityTemplateResponse{}
	_body, _err := client.CreateAuthorityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data archiving ticket to archive data to destinations such as dedicated storage space or ApsaraDB RDS for MySQL instances.
//
// Description:
//
// You can call this API operation only for database instances that are managed in Security Collaboration mode.
//
// @param tmpReq - CreateDataArchiveOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataArchiveOrderResponse
func (client *Client) CreateDataArchiveOrderWithOptions(tmpReq *CreateDataArchiveOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataArchiveOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataArchiveOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataArchiveOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataArchiveOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data archiving ticket to archive data to destinations such as dedicated storage space or ApsaraDB RDS for MySQL instances.
//
// Description:
//
// You can call this API operation only for database instances that are managed in Security Collaboration mode.
//
// @param request - CreateDataArchiveOrderRequest
//
// @return CreateDataArchiveOrderResponse
func (client *Client) CreateDataArchiveOrder(request *CreateDataArchiveOrderRequest) (_result *CreateDataArchiveOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataArchiveOrderResponse{}
	_body, _err := client.CreateDataArchiveOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ticket for changing data in Data Management (DMS).
//
// Description:
//
// For more information about the Normal Data Modify feature, see [Change regular data](https://help.aliyun.com/document_detail/58419.html).
//
// @param tmpReq - CreateDataCorrectOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataCorrectOrderResponse
func (client *Client) CreateDataCorrectOrderWithOptions(tmpReq *CreateDataCorrectOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataCorrectOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataCorrectOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataCorrectOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataCorrectOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket for changing data in Data Management (DMS).
//
// Description:
//
// For more information about the Normal Data Modify feature, see [Change regular data](https://help.aliyun.com/document_detail/58419.html).
//
// @param request - CreateDataCorrectOrderRequest
//
// @return CreateDataCorrectOrderResponse
func (client *Client) CreateDataCorrectOrder(request *CreateDataCorrectOrderRequest) (_result *CreateDataCorrectOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataCorrectOrderResponse{}
	_body, _err := client.CreateDataCorrectOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ticket for clearing historical data.
//
// Description:
//
// For more information about the historical data cleaning, see [Clear historical data](https://help.aliyun.com/document_detail/162507.html).
//
// This operation can be used only for MySQL databases.
//
// @param tmpReq - CreateDataCronClearOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataCronClearOrderResponse
func (client *Client) CreateDataCronClearOrderWithOptions(tmpReq *CreateDataCronClearOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataCronClearOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataCronClearOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataCronClearOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataCronClearOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket for clearing historical data.
//
// Description:
//
// For more information about the historical data cleaning, see [Clear historical data](https://help.aliyun.com/document_detail/162507.html).
//
// This operation can be used only for MySQL databases.
//
// @param request - CreateDataCronClearOrderRequest
//
// @return CreateDataCronClearOrderResponse
func (client *Client) CreateDataCronClearOrder(request *CreateDataCronClearOrderRequest) (_result *CreateDataCronClearOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataCronClearOrderResponse{}
	_body, _err := client.CreateDataCronClearOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ticket to export an SQL result set.
//
// @param tmpReq - CreateDataExportOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataExportOrderResponse
func (client *Client) CreateDataExportOrderWithOptions(tmpReq *CreateDataExportOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataExportOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataExportOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PluginParam) {
		request.PluginParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PluginParam, dara.String("PluginParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.PluginParamShrink) {
		query["PluginParam"] = request.PluginParamShrink
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataExportOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataExportOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket to export an SQL result set.
//
// @param request - CreateDataExportOrderRequest
//
// @return CreateDataExportOrderResponse
func (client *Client) CreateDataExportOrder(request *CreateDataExportOrderRequest) (_result *CreateDataExportOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataExportOrderResponse{}
	_body, _err := client.CreateDataExportOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ticket for importing data to Data Management (DMS).
//
// Description:
//
// For more information about the Large Data Import feature, see [Import data](https://help.aliyun.com/document_detail/161439.html).
//
// @param tmpReq - CreateDataImportOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataImportOrderResponse
func (client *Client) CreateDataImportOrderWithOptions(tmpReq *CreateDataImportOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataImportOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataImportOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataImportOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataImportOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket for importing data to Data Management (DMS).
//
// Description:
//
// For more information about the Large Data Import feature, see [Import data](https://help.aliyun.com/document_detail/161439.html).
//
// @param request - CreateDataImportOrderRequest
//
// @return CreateDataImportOrderResponse
func (client *Client) CreateDataImportOrder(request *CreateDataImportOrderRequest) (_result *CreateDataImportOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataImportOrderResponse{}
	_body, _err := client.CreateDataImportOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database in a data lakehouse.
//
// @param tmpReq - CreateDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeDatabaseResponse
func (client *Client) CreateDataLakeDatabaseWithOptions(tmpReq *CreateDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database in a data lakehouse.
//
// @param request - CreateDataLakeDatabaseRequest
//
// @return CreateDataLakeDatabaseResponse
func (client *Client) CreateDataLakeDatabase(request *CreateDataLakeDatabaseRequest) (_result *CreateDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakeDatabaseResponse{}
	_body, _err := client.CreateDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom function in a data lakehouse.
//
// @param tmpReq - CreateDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeFunctionResponse
func (client *Client) CreateDataLakeFunctionWithOptions(tmpReq *CreateDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FunctionInput) {
		request.FunctionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FunctionInput, dara.String("FunctionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionInputShrink) {
		body["FunctionInput"] = request.FunctionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeFunction"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom function in a data lakehouse.
//
// @param request - CreateDataLakeFunctionRequest
//
// @return CreateDataLakeFunctionResponse
func (client *Client) CreateDataLakeFunction(request *CreateDataLakeFunctionRequest) (_result *CreateDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakeFunctionResponse{}
	_body, _err := client.CreateDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a partition for a table in a data lakehouse.
//
// @param tmpReq - CreateDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakePartitionResponse
func (client *Client) CreateDataLakePartitionWithOptions(tmpReq *CreateDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInput) {
		request.PartitionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInput, dara.String("PartitionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfNotExists) {
		query["IfNotExists"] = request.IfNotExists
	}

	if !dara.IsNil(request.NeedResult) {
		query["NeedResult"] = request.NeedResult
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputShrink) {
		body["PartitionInput"] = request.PartitionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakePartition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a partition for a table in a data lakehouse.
//
// @param request - CreateDataLakePartitionRequest
//
// @return CreateDataLakePartitionResponse
func (client *Client) CreateDataLakePartition(request *CreateDataLakePartitionRequest) (_result *CreateDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakePartitionResponse{}
	_body, _err := client.CreateDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a table in a data lakehouse.
//
// @param tmpReq - CreateDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataLakeTableResponse
func (client *Client) CreateDataLakeTableWithOptions(tmpReq *CreateDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *CreateDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataLakeTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableInput) {
		request.TableInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableInput, dara.String("TableInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableInputShrink) {
		body["TableInput"] = request.TableInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataLakeTable"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a table in a data lakehouse.
//
// @param request - CreateDataLakeTableRequest
//
// @return CreateDataLakeTableResponse
func (client *Client) CreateDataLakeTable(request *CreateDataLakeTableRequest) (_result *CreateDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataLakeTableResponse{}
	_body, _err := client.CreateDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data tracking ticket.
//
// Description:
//
// This operation is available only for instances that are managed in Security Collaboration mode.
//
// @param tmpReq - CreateDataTrackOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataTrackOrderResponse
func (client *Client) CreateDataTrackOrderWithOptions(tmpReq *CreateDataTrackOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataTrackOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDataTrackOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataTrackOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataTrackOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data tracking ticket.
//
// Description:
//
// This operation is available only for instances that are managed in Security Collaboration mode.
//
// @param request - CreateDataTrackOrderRequest
//
// @return CreateDataTrackOrderResponse
func (client *Client) CreateDataTrackOrder(request *CreateDataTrackOrderRequest) (_result *CreateDataTrackOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataTrackOrderResponse{}
	_body, _err := client.CreateDataTrackOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a database export ticket.
//
// @param tmpReq - CreateDatabaseExportOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatabaseExportOrderResponse
func (client *Client) CreateDatabaseExportOrderWithOptions(tmpReq *CreateDatabaseExportOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateDatabaseExportOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateDatabaseExportOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PluginParam) {
		request.PluginParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PluginParam, dara.String("PluginParam"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.PluginParamShrink) {
		query["PluginParam"] = request.PluginParamShrink
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDatabaseExportOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatabaseExportOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a database export ticket.
//
// @param request - CreateDatabaseExportOrderRequest
//
// @return CreateDatabaseExportOrderResponse
func (client *Client) CreateDatabaseExportOrder(request *CreateDatabaseExportOrderRequest) (_result *CreateDatabaseExportOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDatabaseExportOrderResponse{}
	_body, _err := client.CreateDatabaseExportOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - CreateDifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDifyInstanceResponse
func (client *Client) CreateDifyInstanceWithOptions(request *CreateDifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *CreateDifyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AdbpgInstanceMode) {
		query["AdbpgInstanceMode"] = request.AdbpgInstanceMode
	}

	if !dara.IsNil(request.BackupVSwitchId) {
		query["BackupVSwitchId"] = request.BackupVSwitchId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DatabaseOption) {
		query["DatabaseOption"] = request.DatabaseOption
	}

	if !dara.IsNil(request.DbEngineType) {
		query["DbEngineType"] = request.DbEngineType
	}

	if !dara.IsNil(request.DbEngineVersion) {
		query["DbEngineVersion"] = request.DbEngineVersion
	}

	if !dara.IsNil(request.DbInstanceAccount) {
		query["DbInstanceAccount"] = request.DbInstanceAccount
	}

	if !dara.IsNil(request.DbInstanceCategory) {
		query["DbInstanceCategory"] = request.DbInstanceCategory
	}

	if !dara.IsNil(request.DbInstanceClass) {
		query["DbInstanceClass"] = request.DbInstanceClass
	}

	if !dara.IsNil(request.DbInstancePassword) {
		query["DbInstancePassword"] = request.DbInstancePassword
	}

	if !dara.IsNil(request.DbResourceId) {
		query["DbResourceId"] = request.DbResourceId
	}

	if !dara.IsNil(request.DbStorageSize) {
		query["DbStorageSize"] = request.DbStorageSize
	}

	if !dara.IsNil(request.DbStorageType) {
		query["DbStorageType"] = request.DbStorageType
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Edition) {
		query["Edition"] = request.Edition
	}

	if !dara.IsNil(request.EnableExtraEndpoint) {
		query["EnableExtraEndpoint"] = request.EnableExtraEndpoint
	}

	if !dara.IsNil(request.GpuNodeSpec) {
		query["GpuNodeSpec"] = request.GpuNodeSpec
	}

	if !dara.IsNil(request.KvStoreAccount) {
		query["KvStoreAccount"] = request.KvStoreAccount
	}

	if !dara.IsNil(request.KvStoreEngineVersion) {
		query["KvStoreEngineVersion"] = request.KvStoreEngineVersion
	}

	if !dara.IsNil(request.KvStoreInstanceClass) {
		query["KvStoreInstanceClass"] = request.KvStoreInstanceClass
	}

	if !dara.IsNil(request.KvStoreNodeType) {
		query["KvStoreNodeType"] = request.KvStoreNodeType
	}

	if !dara.IsNil(request.KvStoreOption) {
		query["KvStoreOption"] = request.KvStoreOption
	}

	if !dara.IsNil(request.KvStorePassword) {
		query["KvStorePassword"] = request.KvStorePassword
	}

	if !dara.IsNil(request.KvStoreResourceId) {
		query["KvStoreResourceId"] = request.KvStoreResourceId
	}

	if !dara.IsNil(request.KvStoreType) {
		query["KvStoreType"] = request.KvStoreType
	}

	if !dara.IsNil(request.MajorVersion) {
		query["MajorVersion"] = request.MajorVersion
	}

	if !dara.IsNil(request.ModelId) {
		query["ModelId"] = request.ModelId
	}

	if !dara.IsNil(request.ModelOption) {
		query["ModelOption"] = request.ModelOption
	}

	if !dara.IsNil(request.NatGatewayOption) {
		query["NatGatewayOption"] = request.NatGatewayOption
	}

	if !dara.IsNil(request.OssPath) {
		query["OssPath"] = request.OssPath
	}

	if !dara.IsNil(request.OssResourceId) {
		query["OssResourceId"] = request.OssResourceId
	}

	if !dara.IsNil(request.PayPeriod) {
		query["PayPeriod"] = request.PayPeriod
	}

	if !dara.IsNil(request.PayPeriodType) {
		query["PayPeriodType"] = request.PayPeriodType
	}

	if !dara.IsNil(request.PayType) {
		query["PayType"] = request.PayType
	}

	if !dara.IsNil(request.Replicas) {
		query["Replicas"] = request.Replicas
	}

	if !dara.IsNil(request.ResourceQuota) {
		query["ResourceQuota"] = request.ResourceQuota
	}

	if !dara.IsNil(request.SecurityGroupId) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !dara.IsNil(request.SegDiskPerformanceLevel) {
		query["SegDiskPerformanceLevel"] = request.SegDiskPerformanceLevel
	}

	if !dara.IsNil(request.SegNodeNum) {
		query["SegNodeNum"] = request.SegNodeNum
	}

	if !dara.IsNil(request.StorageType) {
		query["StorageType"] = request.StorageType
	}

	if !dara.IsNil(request.VSwitchId) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !dara.IsNil(request.VectordbAccount) {
		query["VectordbAccount"] = request.VectordbAccount
	}

	if !dara.IsNil(request.VectordbCategory) {
		query["VectordbCategory"] = request.VectordbCategory
	}

	if !dara.IsNil(request.VectordbEngineVersion) {
		query["VectordbEngineVersion"] = request.VectordbEngineVersion
	}

	if !dara.IsNil(request.VectordbInstanceSpec) {
		query["VectordbInstanceSpec"] = request.VectordbInstanceSpec
	}

	if !dara.IsNil(request.VectordbOption) {
		query["VectordbOption"] = request.VectordbOption
	}

	if !dara.IsNil(request.VectordbPassword) {
		query["VectordbPassword"] = request.VectordbPassword
	}

	if !dara.IsNil(request.VectordbResourceId) {
		query["VectordbResourceId"] = request.VectordbResourceId
	}

	if !dara.IsNil(request.VectordbStorageSize) {
		query["VectordbStorageSize"] = request.VectordbStorageSize
	}

	if !dara.IsNil(request.VectordbStorageType) {
		query["VectordbStorageType"] = request.VectordbStorageType
	}

	if !dara.IsNil(request.VectordbType) {
		query["VectordbType"] = request.VectordbType
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkspaceDescription) {
		query["WorkspaceDescription"] = request.WorkspaceDescription
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	if !dara.IsNil(request.WorkspaceOption) {
		query["WorkspaceOption"] = request.WorkspaceOption
	}

	if !dara.IsNil(request.ZoneId) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDifyInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - CreateDifyInstanceRequest
//
// @return CreateDifyInstanceResponse
func (client *Client) CreateDifyInstance(request *CreateDifyInstanceRequest) (_result *CreateDifyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDifyInstanceResponse{}
	_body, _err := client.CreateDifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建无锁变更工单
//
// Description:
//
// For more information about the lock-free change feature, see [Overview](https://help.aliyun.com/document_detail/207847.html).
//
// This operation can be used only for instances that are managed in Stable Change or Security Collaboration mode. For more information, see [Change data without the need to lock tables](https://help.aliyun.com/document_detail/96145.html) and [Change schemas without locking tables](https://help.aliyun.com/document_detail/98373.html).
//
// @param tmpReq - CreateFreeLockCorrectOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFreeLockCorrectOrderResponse
func (client *Client) CreateFreeLockCorrectOrderWithOptions(tmpReq *CreateFreeLockCorrectOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateFreeLockCorrectOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateFreeLockCorrectOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFreeLockCorrectOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFreeLockCorrectOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建无锁变更工单
//
// Description:
//
// For more information about the lock-free change feature, see [Overview](https://help.aliyun.com/document_detail/207847.html).
//
// This operation can be used only for instances that are managed in Stable Change or Security Collaboration mode. For more information, see [Change data without the need to lock tables](https://help.aliyun.com/document_detail/96145.html) and [Change schemas without locking tables](https://help.aliyun.com/document_detail/98373.html).
//
// @param request - CreateFreeLockCorrectOrderRequest
//
// @return CreateFreeLockCorrectOrderResponse
func (client *Client) CreateFreeLockCorrectOrder(request *CreateFreeLockCorrectOrderRequest) (_result *CreateFreeLockCorrectOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFreeLockCorrectOrderResponse{}
	_body, _err := client.CreateFreeLockCorrectOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace for data warehouse development in Data Management (DMS).
//
// Description:
//
//	  The workspace name must be unique within a tenant. If a workspace with the same name already exists within the tenant, the call may fail.
//
//		- You can call the [GetLhSpaceByName](https://help.aliyun.com/document_detail/424379.html) operation to query whether a workspace with a specific name already exists as a DMS administrator or database administrator (DBA).
//
// @param request - CreateLakeHouseSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLakeHouseSpaceResponse
func (client *Client) CreateLakeHouseSpaceWithOptions(request *CreateLakeHouseSpaceRequest, runtime *dara.RuntimeOptions) (_result *CreateLakeHouseSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DevDbId) {
		query["DevDbId"] = request.DevDbId
	}

	if !dara.IsNil(request.DwDbType) {
		query["DwDbType"] = request.DwDbType
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.ProdDbId) {
		query["ProdDbId"] = request.ProdDbId
	}

	if !dara.IsNil(request.SpaceConfig) {
		query["SpaceConfig"] = request.SpaceConfig
	}

	if !dara.IsNil(request.SpaceName) {
		query["SpaceName"] = request.SpaceName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLakeHouseSpace"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLakeHouseSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace for data warehouse development in Data Management (DMS).
//
// Description:
//
//	  The workspace name must be unique within a tenant. If a workspace with the same name already exists within the tenant, the call may fail.
//
//		- You can call the [GetLhSpaceByName](https://help.aliyun.com/document_detail/424379.html) operation to query whether a workspace with a specific name already exists as a DMS administrator or database administrator (DBA).
//
// @param request - CreateLakeHouseSpaceRequest
//
// @return CreateLakeHouseSpaceResponse
func (client *Client) CreateLakeHouseSpace(request *CreateLakeHouseSpaceRequest) (_result *CreateLakeHouseSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLakeHouseSpaceResponse{}
	_body, _err := client.CreateLakeHouseSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a logical database in Database Management (DMS).
//
// @param tmpReq - CreateLogicDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLogicDatabaseResponse
func (client *Client) CreateLogicDatabaseWithOptions(tmpReq *CreateLogicDatabaseRequest, runtime *dara.RuntimeOptions) (_result *CreateLogicDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateLogicDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatabaseIds) {
		request.DatabaseIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatabaseIds, dara.String("DatabaseIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.DatabaseIdsShrink) {
		query["DatabaseIds"] = request.DatabaseIdsShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLogicDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLogicDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a logical database in Database Management (DMS).
//
// @param request - CreateLogicDatabaseRequest
//
// @return CreateLogicDatabaseResponse
func (client *Client) CreateLogicDatabase(request *CreateLogicDatabaseRequest) (_result *CreateLogicDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateLogicDatabaseResponse{}
	_body, _err := client.CreateLogicDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Asset Category
//
// @param request - CreateMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetaCategoryResponse
func (client *Client) CreateMetaCategoryWithOptions(request *CreateMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *CreateMetaCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentCategoryId) {
		query["ParentCategoryId"] = request.ParentCategoryId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetaCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Asset Category
//
// @param request - CreateMetaCategoryRequest
//
// @return CreateMetaCategoryResponse
func (client *Client) CreateMetaCategory(request *CreateMetaCategoryRequest) (_result *CreateMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMetaCategoryResponse{}
	_body, _err := client.CreateMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a ticket in Data Management (DMS).
//
// Description:
//
// To facilitate ticket creation, you can call the following dedicated operations to create some types of tickets:
//
//   - [CreateDataCorrectOrder](https://help.aliyun.com/document_detail/208388.html): creates a regular data change ticket.
//
//   - [CreateDataCronClearOrder](https://help.aliyun.com/document_detail/208385.html): creates a ticket to clear historical data.
//
//   - [CreateDataImportOrder](https://help.aliyun.com/document_detail/208387.html): creates a data import ticket.
//
//   - [CreateFreeLockCorrectOrder](https://help.aliyun.com/document_detail/208386.html): creates a lock-free change ticket.
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
	if !dara.IsNil(tmpReq.PluginParam) {
		request.PluginParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PluginParam, dara.String("PluginParam"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.RelatedUserList) {
		query["RelatedUserList"] = request.RelatedUserList
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PluginParamShrink) {
		body["PluginParam"] = request.PluginParamShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a ticket in Data Management (DMS).
//
// Description:
//
// To facilitate ticket creation, you can call the following dedicated operations to create some types of tickets:
//
//   - [CreateDataCorrectOrder](https://help.aliyun.com/document_detail/208388.html): creates a regular data change ticket.
//
//   - [CreateDataCronClearOrder](https://help.aliyun.com/document_detail/208385.html): creates a ticket to clear historical data.
//
//   - [CreateDataImportOrder](https://help.aliyun.com/document_detail/208387.html): creates a data import ticket.
//
//   - [CreateFreeLockCorrectOrder](https://help.aliyun.com/document_detail/208386.html): creates a lock-free change ticket.
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
// Creates a programmable object ticket.
//
// Description:
//
// You can call this API operation only for database instances that are managed in Security Collaboration mode.
//
// @param tmpReq - CreateProcCorrectOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProcCorrectOrderResponse
func (client *Client) CreateProcCorrectOrderWithOptions(tmpReq *CreateProcCorrectOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateProcCorrectOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateProcCorrectOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProcCorrectOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProcCorrectOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a programmable object ticket.
//
// Description:
//
// You can call this API operation only for database instances that are managed in Security Collaboration mode.
//
// @param request - CreateProcCorrectOrderRequest
//
// @return CreateProcCorrectOrderResponse
func (client *Client) CreateProcCorrectOrder(request *CreateProcCorrectOrderRequest) (_result *CreateProcCorrectOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProcCorrectOrderResponse{}
	_body, _err := client.CreateProcCorrectOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables the secure access proxy feature for a database instance.
//
// Description:
//
// - The database instance runs the MySQL or MariaDB database engine. For example, the database instance can be an ApsaraDB RDS for MySQL instance, a PolarDB for MySQL cluster, a Distributed Relational Database Service (DRDS) cluster, or an AnalyticDB for MySQL cluster. The database instance can also be a self-managed MySQL or MariaDB database, or a MySQL or MariaDB database in a third-party cloud.
//
// - The database instance resides in the China (Hangzhou) or China (Beijing) region.
//
// - You are a Data Management (DMS) administrator, a database administrator (DBA), or the owner of the database instance.
//
// @param request - CreateProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProxyResponse
func (client *Client) CreateProxyWithOptions(request *CreateProxyRequest, runtime *dara.RuntimeOptions) (_result *CreateProxyResponse, _err error) {
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

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Username) {
		query["Username"] = request.Username
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProxy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables the secure access proxy feature for a database instance.
//
// Description:
//
// - The database instance runs the MySQL or MariaDB database engine. For example, the database instance can be an ApsaraDB RDS for MySQL instance, a PolarDB for MySQL cluster, a Distributed Relational Database Service (DRDS) cluster, or an AnalyticDB for MySQL cluster. The database instance can also be a self-managed MySQL or MariaDB database, or a MySQL or MariaDB database in a third-party cloud.
//
// - The database instance resides in the China (Hangzhou) or China (Beijing) region.
//
// - You are a Data Management (DMS) administrator, a database administrator (DBA), or the owner of the database instance.
//
// @param request - CreateProxyRequest
//
// @return CreateProxyResponse
func (client *Client) CreateProxy(request *CreateProxyRequest) (_result *CreateProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProxyResponse{}
	_body, _err := client.CreateProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call the CreateProxyAccess to authorize users to access the DB instance through the Data Security Protection agent.
//
// Description:
//
// - The data security protection feature is enabled for the instance.
//
// - Your user role is the administrator role, DBA role, or the owner of data security protection for the current instance.
//
// @param request - CreateProxyAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProxyAccessResponse
func (client *Client) CreateProxyAccessWithOptions(request *CreateProxyAccessRequest, runtime *dara.RuntimeOptions) (_result *CreateProxyAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IndepAccount) {
		query["IndepAccount"] = request.IndepAccount
	}

	if !dara.IsNil(request.IndepPassword) {
		query["IndepPassword"] = request.IndepPassword
	}

	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProxyAccess"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProxyAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call the CreateProxyAccess to authorize users to access the DB instance through the Data Security Protection agent.
//
// Description:
//
// - The data security protection feature is enabled for the instance.
//
// - Your user role is the administrator role, DBA role, or the owner of data security protection for the current instance.
//
// @param request - CreateProxyAccessRequest
//
// @return CreateProxyAccessResponse
func (client *Client) CreateProxyAccess(request *CreateProxyAccessRequest) (_result *CreateProxyAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProxyAccessResponse{}
	_body, _err := client.CreateProxyAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a schema design ticket.
//
// @param request - CreatePublishGroupTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePublishGroupTaskResponse
func (client *Client) CreatePublishGroupTaskWithOptions(request *CreatePublishGroupTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatePublishGroupTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PlanTime) {
		query["PlanTime"] = request.PlanTime
	}

	if !dara.IsNil(request.PublishStrategy) {
		query["PublishStrategy"] = request.PublishStrategy
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePublishGroupTask"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePublishGroupTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a schema design ticket.
//
// @param request - CreatePublishGroupTaskRequest
//
// @return CreatePublishGroupTaskResponse
func (client *Client) CreatePublishGroupTask(request *CreatePublishGroupTaskRequest) (_result *CreatePublishGroupTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePublishGroupTaskResponse{}
	_body, _err := client.CreatePublishGroupTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a SQL review ticket.
//
// Description:
//
// You can call this operation only for database instances that are managed in Security Collaboration mode.
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param tmpReq - CreateSQLReviewOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSQLReviewOrderResponse
func (client *Client) CreateSQLReviewOrderWithOptions(tmpReq *CreateSQLReviewOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateSQLReviewOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateSQLReviewOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSQLReviewOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSQLReviewOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a SQL review ticket.
//
// Description:
//
// You can call this operation only for database instances that are managed in Security Collaboration mode.
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param request - CreateSQLReviewOrderRequest
//
// @return CreateSQLReviewOrderResponse
func (client *Client) CreateSQLReviewOrder(request *CreateSQLReviewOrderRequest) (_result *CreateSQLReviewOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSQLReviewOrderResponse{}
	_body, _err := client.CreateSQLReviewOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a business scenario to group task flows by business scenario.
//
// @param request - CreateScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScenarioResponse
func (client *Client) CreateScenarioWithOptions(request *CreateScenarioRequest, runtime *dara.RuntimeOptions) (_result *CreateScenarioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ScenarioName) {
		query["ScenarioName"] = request.ScenarioName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateScenario"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a business scenario to group task flows by business scenario.
//
// @param request - CreateScenarioRequest
//
// @return CreateScenarioResponse
func (client *Client) CreateScenario(request *CreateScenarioRequest) (_result *CreateScenarioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateScenarioResponse{}
	_body, _err := client.CreateScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a security rule set.
//
// @param request - CreateStandardGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStandardGroupResponse
func (client *Client) CreateStandardGroupWithOptions(request *CreateStandardGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateStandardGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStandardGroup"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStandardGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a security rule set.
//
// @param request - CreateStandardGroupRequest
//
// @return CreateStandardGroupResponse
func (client *Client) CreateStandardGroup(request *CreateStandardGroupRequest) (_result *CreateStandardGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStandardGroupResponse{}
	_body, _err := client.CreateStandardGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a schema synchronization ticket.
//
// @param tmpReq - CreateStructSyncOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStructSyncOrderResponse
func (client *Client) CreateStructSyncOrderWithOptions(tmpReq *CreateStructSyncOrderRequest, runtime *dara.RuntimeOptions) (_result *CreateStructSyncOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStructSyncOrderShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Param) {
		request.ParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Param, dara.String("Param"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RelatedUserList) {
		request.RelatedUserListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RelatedUserList, dara.String("RelatedUserList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AttachmentKey) {
		query["AttachmentKey"] = request.AttachmentKey
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ParamShrink) {
		query["Param"] = request.ParamShrink
	}

	if !dara.IsNil(request.RelatedUserListShrink) {
		query["RelatedUserList"] = request.RelatedUserListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStructSyncOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStructSyncOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a schema synchronization ticket.
//
// @param request - CreateStructSyncOrderRequest
//
// @return CreateStructSyncOrderResponse
func (client *Client) CreateStructSyncOrder(request *CreateStructSyncOrderRequest) (_result *CreateStructSyncOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStructSyncOrderResponse{}
	_body, _err := client.CreateStructSyncOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task node for a task flow.
//
// @param request - CreateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.GraphParam) {
		query["GraphParam"] = request.GraphParam
	}

	if !dara.IsNil(request.NodeContent) {
		query["NodeContent"] = request.NodeContent
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.NodeOutput) {
		query["NodeOutput"] = request.NodeOutput
	}

	if !dara.IsNil(request.NodeType) {
		query["NodeType"] = request.NodeType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.TimeVariables) {
		query["TimeVariables"] = request.TimeVariables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task node for a task flow.
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task flow.
//
// @param request - CreateTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskFlowResponse
func (client *Client) CreateTaskFlowWithOptions(request *CreateTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *CreateTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagName) {
		query["DagName"] = request.DagName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task flow.
//
// @param request - CreateTaskFlowRequest
//
// @return CreateTaskFlowResponse
func (client *Client) CreateTaskFlow(request *CreateTaskFlowRequest) (_result *CreateTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTaskFlowResponse{}
	_body, _err := client.CreateTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建上传附件任务
//
// @param request - CreateUploadFileJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadFileJobResponse
func (client *Client) CreateUploadFileJobWithOptions(request *CreateUploadFileJobRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadFileJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileSource) {
		query["FileSource"] = request.FileSource
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UploadURL) {
		query["UploadURL"] = request.UploadURL
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUploadFileJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadFileJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建上传附件任务
//
// @param request - CreateUploadFileJobRequest
//
// @return CreateUploadFileJobResponse
func (client *Client) CreateUploadFileJob(request *CreateUploadFileJobRequest) (_result *CreateUploadFileJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUploadFileJobResponse{}
	_body, _err := client.CreateUploadFileJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a task to upload an Object Storage Service (OSS) file and obtain the key of the task.
//
// @param tmpReq - CreateUploadOSSFileJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUploadOSSFileJobResponse
func (client *Client) CreateUploadOSSFileJobWithOptions(tmpReq *CreateUploadOSSFileJobRequest, runtime *dara.RuntimeOptions) (_result *CreateUploadOSSFileJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateUploadOSSFileJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UploadTarget) {
		request.UploadTargetShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UploadTarget, dara.String("UploadTarget"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileSource) {
		query["FileSource"] = request.FileSource
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UploadTargetShrink) {
		query["UploadTarget"] = request.UploadTargetShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUploadOSSFileJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUploadOSSFileJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a task to upload an Object Storage Service (OSS) file and obtain the key of the task.
//
// @param request - CreateUploadOSSFileJobRequest
//
// @return CreateUploadOSSFileJobResponse
func (client *Client) CreateUploadOSSFileJob(request *CreateUploadOSSFileJobRequest) (_result *CreateUploadOSSFileJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUploadOSSFileJobResponse{}
	_body, _err := client.CreateUploadOSSFileJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// Description:
//
// ## [](#)Prerequisites
//
//   - A virtual private cloud (VPC) is created.
//
//   - Log on to the DMS console by using an Alibaba Cloud account or a RAM user that has high permissions, and grant DMS the permissions to access cloud resources (AliyunDMSProcessingDataRolePolicy).
//
// @param request - CreateWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// Description:
//
// ## [](#)Prerequisites
//
//   - A virtual private cloud (VPC) is created.
//
//   - Log on to the DMS console by using an Alibaba Cloud account or a RAM user that has high permissions, and grant DMS the permissions to access cloud resources (AliyunDMSProcessingDataRolePolicy).
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detaches a policy from a user or role.
//
// @param request - DeleteAbacAuthorizationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAbacAuthorizationResponse
func (client *Client) DeleteAbacAuthorizationWithOptions(request *DeleteAbacAuthorizationRequest, runtime *dara.RuntimeOptions) (_result *DeleteAbacAuthorizationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthorizationId) {
		query["AuthorizationId"] = request.AuthorizationId
	}

	if !dara.IsNil(request.IdentityType) {
		query["IdentityType"] = request.IdentityType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAbacAuthorization"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAbacAuthorizationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detaches a policy from a user or role.
//
// @param request - DeleteAbacAuthorizationRequest
//
// @return DeleteAbacAuthorizationResponse
func (client *Client) DeleteAbacAuthorization(request *DeleteAbacAuthorizationRequest) (_result *DeleteAbacAuthorizationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAbacAuthorizationResponse{}
	_body, _err := client.DeleteAbacAuthorizationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a created permission policy.
//
// @param request - DeleteAbacPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAbacPolicyResponse
func (client *Client) DeleteAbacPolicyWithOptions(request *DeleteAbacPolicyRequest, runtime *dara.RuntimeOptions) (_result *DeleteAbacPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbacPolicyId) {
		query["AbacPolicyId"] = request.AbacPolicyId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAbacPolicy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAbacPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a created permission policy.
//
// @param request - DeleteAbacPolicyRequest
//
// @return DeleteAbacPolicyResponse
func (client *Client) DeleteAbacPolicy(request *DeleteAbacPolicyRequest) (_result *DeleteAbacPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAbacPolicyResponse{}
	_body, _err := client.DeleteAbacPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a permission template.
//
// @param request - DeleteAuthorityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAuthorityTemplateResponse
func (client *Client) DeleteAuthorityTemplateWithOptions(request *DeleteAuthorityTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteAuthorityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAuthorityTemplate"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAuthorityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a permission template.
//
// @param request - DeleteAuthorityTemplateRequest
//
// @return DeleteAuthorityTemplateResponse
func (client *Client) DeleteAuthorityTemplate(request *DeleteAuthorityTemplateRequest) (_result *DeleteAuthorityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteAuthorityTemplateResponse{}
	_body, _err := client.DeleteAuthorityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a database from a data lakehouse.
//
// @param request - DeleteDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeDatabaseResponse
func (client *Client) DeleteDataLakeDatabaseWithOptions(request *DeleteDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a database from a data lakehouse.
//
// @param request - DeleteDataLakeDatabaseRequest
//
// @return DeleteDataLakeDatabaseResponse
func (client *Client) DeleteDataLakeDatabase(request *DeleteDataLakeDatabaseRequest) (_result *DeleteDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakeDatabaseResponse{}
	_body, _err := client.DeleteDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function in a data lakehouse.
//
// @param request - DeleteDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeFunctionResponse
func (client *Client) DeleteDataLakeFunctionWithOptions(request *DeleteDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeFunction"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user-defined function in a data lakehouse.
//
// @param request - DeleteDataLakeFunctionRequest
//
// @return DeleteDataLakeFunctionResponse
func (client *Client) DeleteDataLakeFunction(request *DeleteDataLakeFunctionRequest) (_result *DeleteDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakeFunctionResponse{}
	_body, _err := client.DeleteDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a partition from a table in a data lakehouse.
//
// @param tmpReq - DeleteDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakePartitionResponse
func (client *Client) DeleteDataLakePartitionWithOptions(tmpReq *DeleteDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionValues) {
		request.PartitionValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionValues, dara.String("PartitionValues"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.IfExists) {
		query["IfExists"] = request.IfExists
	}

	if !dara.IsNil(request.PartitionValuesShrink) {
		query["PartitionValues"] = request.PartitionValuesShrink
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakePartition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a partition from a table in a data lakehouse.
//
// @param request - DeleteDataLakePartitionRequest
//
// @return DeleteDataLakePartitionResponse
func (client *Client) DeleteDataLakePartition(request *DeleteDataLakePartitionRequest) (_result *DeleteDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakePartitionResponse{}
	_body, _err := client.DeleteDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a table from a data lakehouse.
//
// @param request - DeleteDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataLakeTableResponse
func (client *Client) DeleteDataLakeTableWithOptions(request *DeleteDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataLakeTable"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a table from a data lakehouse.
//
// @param request - DeleteDataLakeTableRequest
//
// @return DeleteDataLakeTableResponse
func (client *Client) DeleteDataLakeTable(request *DeleteDataLakeTableRequest) (_result *DeleteDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataLakeTableResponse{}
	_body, _err := client.DeleteDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a database instance from Data Management (DMS).
//
// Description:
//
// Note: You can call this operation only to remove a database instance from the instance list of DMS. The instance is not deleted or shut down.
//
// @param request - DeleteInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *dara.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a database instance from Data Management (DMS).
//
// Description:
//
// Note: You can call this operation only to remove a database instance from the instance list of DMS. The instance is not deleted or shut down.
//
// @param request - DeleteInstanceRequest
//
// @return DeleteInstanceResponse
func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// @param request - DeleteLakeHouseSpaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLakeHouseSpaceResponse
func (client *Client) DeleteLakeHouseSpaceWithOptions(request *DeleteLakeHouseSpaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteLakeHouseSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceId) {
		query["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLakeHouseSpace"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLakeHouseSpaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// @param request - DeleteLakeHouseSpaceRequest
//
// @return DeleteLakeHouseSpaceResponse
func (client *Client) DeleteLakeHouseSpace(request *DeleteLakeHouseSpaceRequest) (_result *DeleteLakeHouseSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLakeHouseSpaceResponse{}
	_body, _err := client.DeleteLakeHouseSpaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除数仓空间成员
//
// Description:
//
// You must call this operation as a DMS administrator, a database administrator (DBA), or a workspace administrator.
//
// You cannot call this operation to transfer the ownership of a task flow. To transfer the ownership of a task flow, call the [ChangLhDagOwner](https://help.aliyun.com/document_detail/424761.html) operation.
//
// @param tmpReq - DeleteLhMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLhMembersResponse
func (client *Client) DeleteLhMembersWithOptions(tmpReq *DeleteLhMembersRequest, runtime *dara.RuntimeOptions) (_result *DeleteLhMembersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteLhMembersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MemberIds) {
		request.MemberIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MemberIds, dara.String("MemberIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.MemberIdsShrink) {
		query["MemberIds"] = request.MemberIdsShrink
	}

	if !dara.IsNil(request.ObjectId) {
		query["ObjectId"] = request.ObjectId
	}

	if !dara.IsNil(request.ObjectType) {
		query["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLhMembers"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLhMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数仓空间成员
//
// Description:
//
// You must call this operation as a DMS administrator, a database administrator (DBA), or a workspace administrator.
//
// You cannot call this operation to transfer the ownership of a task flow. To transfer the ownership of a task flow, call the [ChangLhDagOwner](https://help.aliyun.com/document_detail/424761.html) operation.
//
// @param request - DeleteLhMembersRequest
//
// @return DeleteLhMembersResponse
func (client *Client) DeleteLhMembers(request *DeleteLhMembersRequest) (_result *DeleteLhMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLhMembersResponse{}
	_body, _err := client.DeleteLhMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a logical database in Database Management (DMS). This operation only deletes the specified logical database but does not delete physical databases.
//
// @param request - DeleteLogicDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogicDatabaseResponse
func (client *Client) DeleteLogicDatabaseWithOptions(request *DeleteLogicDatabaseRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogicDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicDbId) {
		query["LogicDbId"] = request.LogicDbId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogicDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogicDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a logical database in Database Management (DMS). This operation only deletes the specified logical database but does not delete physical databases.
//
// @param request - DeleteLogicDatabaseRequest
//
// @return DeleteLogicDatabaseResponse
func (client *Client) DeleteLogicDatabase(request *DeleteLogicDatabaseRequest) (_result *DeleteLogicDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLogicDatabaseResponse{}
	_body, _err := client.DeleteLogicDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the routing algorithm of a logical table.
//
// @param request - DeleteLogicTableRouteConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLogicTableRouteConfigResponse
func (client *Client) DeleteLogicTableRouteConfigWithOptions(request *DeleteLogicTableRouteConfigRequest, runtime *dara.RuntimeOptions) (_result *DeleteLogicTableRouteConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RouteKey) {
		query["RouteKey"] = request.RouteKey
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLogicTableRouteConfig"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLogicTableRouteConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the routing algorithm of a logical table.
//
// @param request - DeleteLogicTableRouteConfigRequest
//
// @return DeleteLogicTableRouteConfigResponse
func (client *Client) DeleteLogicTableRouteConfig(request *DeleteLogicTableRouteConfigRequest) (_result *DeleteLogicTableRouteConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLogicTableRouteConfigResponse{}
	_body, _err := client.DeleteLogicTableRouteConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an asset category.
//
// @param request - DeleteMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaCategoryResponse
func (client *Client) DeleteMetaCategoryWithOptions(request *DeleteMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetaCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an asset category.
//
// @param request - DeleteMetaCategoryRequest
//
// @return DeleteMetaCategoryResponse
func (client *Client) DeleteMetaCategory(request *DeleteMetaCategoryRequest) (_result *DeleteMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetaCategoryResponse{}
	_body, _err := client.DeleteMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to disable the data security protection proxy of a DB instance.
//
// Description:
//
// After you disable this feature, your DB instance loses the JDBC protocol. All authorization information is recycled.
//
// @param request - DeleteProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProxyResponse
func (client *Client) DeleteProxyWithOptions(request *DeleteProxyRequest, runtime *dara.RuntimeOptions) (_result *DeleteProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProxy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to disable the data security protection proxy of a DB instance.
//
// Description:
//
// After you disable this feature, your DB instance loses the JDBC protocol. All authorization information is recycled.
//
// @param request - DeleteProxyRequest
//
// @return DeleteProxyResponse
func (client *Client) DeleteProxy(request *DeleteProxyRequest) (_result *DeleteProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProxyResponse{}
	_body, _err := client.DeleteProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to DeleteProxyAccess reclaim the data security protection authorization of the target user.
//
// @param request - DeleteProxyAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProxyAccessResponse
func (client *Client) DeleteProxyAccessWithOptions(request *DeleteProxyAccessRequest, runtime *dara.RuntimeOptions) (_result *DeleteProxyAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyAccessId) {
		query["ProxyAccessId"] = request.ProxyAccessId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProxyAccess"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProxyAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to DeleteProxyAccess reclaim the data security protection authorization of the target user.
//
// @param request - DeleteProxyAccessRequest
//
// @return DeleteProxyAccessResponse
func (client *Client) DeleteProxyAccess(request *DeleteProxyAccessRequest) (_result *DeleteProxyAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProxyAccessResponse{}
	_body, _err := client.DeleteProxyAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a business scenario.
//
// Description:
//
// When you call this operation, make sure that no task flow is specified in the business scenario.
//
// @param request - DeleteScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteScenarioResponse
func (client *Client) DeleteScenarioWithOptions(request *DeleteScenarioRequest, runtime *dara.RuntimeOptions) (_result *DeleteScenarioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteScenario"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a business scenario.
//
// Description:
//
// When you call this operation, make sure that no task flow is specified in the business scenario.
//
// @param request - DeleteScenarioRequest
//
// @return DeleteScenarioResponse
func (client *Client) DeleteScenario(request *DeleteScenarioRequest) (_result *DeleteScenarioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteScenarioResponse{}
	_body, _err := client.DeleteScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a security rule set.
//
// Description:
//
// Before you call this API operation, make sure that the security rule set is not associated with an instance.
//
// @param request - DeleteStandardGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStandardGroupResponse
func (client *Client) DeleteStandardGroupWithOptions(request *DeleteStandardGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteStandardGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStandardGroup"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStandardGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a security rule set.
//
// Description:
//
// Before you call this API operation, make sure that the security rule set is not associated with an instance.
//
// @param request - DeleteStandardGroupRequest
//
// @return DeleteStandardGroupResponse
func (client *Client) DeleteStandardGroup(request *DeleteStandardGroupRequest) (_result *DeleteStandardGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStandardGroupResponse{}
	_body, _err := client.DeleteStandardGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task in a specified task flow.
//
// @param request - DeleteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskResponse
func (client *Client) DeleteTaskWithOptions(request *DeleteTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTask"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task in a specified task flow.
//
// @param request - DeleteTaskRequest
//
// @return DeleteTaskResponse
func (client *Client) DeleteTask(request *DeleteTaskRequest) (_result *DeleteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTaskResponse{}
	_body, _err := client.DeleteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a task flow.
//
// @param request - DeleteTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskFlowResponse
func (client *Client) DeleteTaskFlowWithOptions(request *DeleteTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a task flow.
//
// @param request - DeleteTaskFlowRequest
//
// @return DeleteTaskFlowResponse
func (client *Client) DeleteTaskFlow(request *DeleteTaskFlowRequest) (_result *DeleteTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTaskFlowResponse{}
	_body, _err := client.DeleteTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes task flow edges based on multiple conditions.
//
// Description:
//
// This operation is used for multi-condition query. You can call it to delete the edges of a specified task flow that meet all specified conditions.
//
// @param request - DeleteTaskFlowEdgesByConditionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTaskFlowEdgesByConditionResponse
func (client *Client) DeleteTaskFlowEdgesByConditionWithOptions(request *DeleteTaskFlowEdgesByConditionRequest, runtime *dara.RuntimeOptions) (_result *DeleteTaskFlowEdgesByConditionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.NodeEnd) {
		query["NodeEnd"] = request.NodeEnd
	}

	if !dara.IsNil(request.NodeFrom) {
		query["NodeFrom"] = request.NodeFrom
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTaskFlowEdgesByCondition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTaskFlowEdgesByConditionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes task flow edges based on multiple conditions.
//
// Description:
//
// This operation is used for multi-condition query. You can call it to delete the edges of a specified task flow that meet all specified conditions.
//
// @param request - DeleteTaskFlowEdgesByConditionRequest
//
// @return DeleteTaskFlowEdgesByConditionResponse
func (client *Client) DeleteTaskFlowEdgesByCondition(request *DeleteTaskFlowEdgesByConditionRequest) (_result *DeleteTaskFlowEdgesByConditionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTaskFlowEdgesByConditionResponse{}
	_body, _err := client.DeleteTaskFlowEdgesByConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an Alibaba Cloud account that is no longer used.
//
// Description:
//
// The effect of deleting a user by calling this operation is the same as that of deleting a user by choosing System Management > User Management in the DMS Enterprise console. The administrator of DMS Enterprise can call this operation to delete a user that is no longer used from DMS Enterprise. After the user is deleted, the data source permission, data owner configuration, and database administrator (DBA) configuration of the corresponding Alibaba Cloud account or Resource Access Management (RAM) user are revoked and become invalid.
//
// >  This operation only removes the association of the Alibaba Cloud account or RAM user with DMS Enterprise of the enterprise, rather than actually deleting the Alibaba Cloud account or RAM user. After the user is deleted, the Alibaba Cloud account or RAM user cannot log on to DMS Enterprise, unless the user is added to DMS Enterprise again.
//
// @param request - DeleteUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *dara.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an Alibaba Cloud account that is no longer used.
//
// Description:
//
// The effect of deleting a user by calling this operation is the same as that of deleting a user by choosing System Management > User Management in the DMS Enterprise console. The administrator of DMS Enterprise can call this operation to delete a user that is no longer used from DMS Enterprise. After the user is deleted, the data source permission, data owner configuration, and database administrator (DBA) configuration of the corresponding Alibaba Cloud account or Resource Access Management (RAM) user are revoked and become invalid.
//
// >  This operation only removes the association of the Alibaba Cloud account or RAM user with DMS Enterprise of the enterprise, rather than actually deleting the Alibaba Cloud account or RAM user. After the user is deleted, the Alibaba Cloud account or RAM user cannot log on to DMS Enterprise, unless the user is added to DMS Enterprise again.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// @param request - DeleteWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(request *DeleteWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDifyAttributeResponse
func (client *Client) DescribeDifyAttributeWithOptions(request *DescribeDifyAttributeRequest, runtime *dara.RuntimeOptions) (_result *DescribeDifyAttributeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppUuid) {
		query["AppUuid"] = request.AppUuid
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDifyAttribute"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDifyAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyAttributeRequest
//
// @return DescribeDifyAttributeResponse
func (client *Client) DescribeDifyAttribute(request *DescribeDifyAttributeRequest) (_result *DescribeDifyAttributeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDifyAttributeResponse{}
	_body, _err := client.DescribeDifyAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyDefaultVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDifyDefaultVpcResponse
func (client *Client) DescribeDifyDefaultVpcWithOptions(request *DescribeDifyDefaultVpcRequest, runtime *dara.RuntimeOptions) (_result *DescribeDifyDefaultVpcResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDifyDefaultVpc"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDifyDefaultVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyDefaultVpcRequest
//
// @return DescribeDifyDefaultVpcResponse
func (client *Client) DescribeDifyDefaultVpc(request *DescribeDifyDefaultVpcRequest) (_result *DescribeDifyDefaultVpcResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDifyDefaultVpcResponse{}
	_body, _err := client.DescribeDifyDefaultVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于获取Dify当前支持的版本信息
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyEditionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDifyEditionsResponse
func (client *Client) DescribeDifyEditionsWithOptions(request *DescribeDifyEditionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDifyEditionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDifyEditions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDifyEditionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于获取Dify当前支持的版本信息
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyEditionsRequest
//
// @return DescribeDifyEditionsResponse
func (client *Client) DescribeDifyEditions(request *DescribeDifyEditionsRequest) (_result *DescribeDifyEditionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDifyEditionsResponse{}
	_body, _err := client.DescribeDifyEditionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDifyRegionsResponse
func (client *Client) DescribeDifyRegionsWithOptions(request *DescribeDifyRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDifyRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDifyRegions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDifyRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - DescribeDifyRegionsRequest
//
// @return DescribeDifyRegionsResponse
func (client *Client) DescribeDifyRegions(request *DescribeDifyRegionsRequest) (_result *DescribeDifyRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeDifyRegionsResponse{}
	_body, _err := client.DescribeDifyRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to disable a user that is temporarily not used in Data Management (DMS) Enterprise.
//
// Description:
//
// The effect of disabling a user by calling this operation is the same as that of disabling a user by choosing System Management > User Management in the DMS Enterprise console. The administrator of DMS Enterprise can call this operation to disable a user that is temporarily not used in DMS Enterprise. After the user is disabled, the data source permission, data owner configuration, and database administrator (DBA) configuration of the corresponding Alibaba Cloud account or Resource Access Management (RAM) user are revoked and become invalid.
//
// >  This operation only stops the Alibaba Cloud account or RAM user from logging on to DMS Enterprise of the enterprise, rather than actually disabling the Alibaba Cloud account or RAM user. After the user is disabled, the Alibaba Cloud account or RAM user cannot log on to DMS Enterprise, unless the user is enabled again. The disabled user, however, still exists in DMS Enterprise.
//
// @param request - DisableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableUserResponse
func (client *Client) DisableUserWithOptions(request *DisableUserRequest, runtime *dara.RuntimeOptions) (_result *DisableUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DisableUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DisableUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to disable a user that is temporarily not used in Data Management (DMS) Enterprise.
//
// Description:
//
// The effect of disabling a user by calling this operation is the same as that of disabling a user by choosing System Management > User Management in the DMS Enterprise console. The administrator of DMS Enterprise can call this operation to disable a user that is temporarily not used in DMS Enterprise. After the user is disabled, the data source permission, data owner configuration, and database administrator (DBA) configuration of the corresponding Alibaba Cloud account or Resource Access Management (RAM) user are revoked and become invalid.
//
// >  This operation only stops the Alibaba Cloud account or RAM user from logging on to DMS Enterprise of the enterprise, rather than actually disabling the Alibaba Cloud account or RAM user. After the user is disabled, the Alibaba Cloud account or RAM user cannot log on to DMS Enterprise, unless the user is enabled again. The disabled user, however, still exists in DMS Enterprise.
//
// @param request - DisableUserRequest
//
// @return DisableUserResponse
func (client *Client) DisableUser(request *DisableUserRequest) (_result *DisableUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DisableUserResponse{}
	_body, _err := client.DisableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downloads the parsing result of a data tracking task.
//
// @param tmpReq - DownloadDataTrackResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadDataTrackResultResponse
func (client *Client) DownloadDataTrackResultWithOptions(tmpReq *DownloadDataTrackResultRequest, runtime *dara.RuntimeOptions) (_result *DownloadDataTrackResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DownloadDataTrackResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ColumnFilter) {
		request.ColumnFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColumnFilter, dara.String("ColumnFilter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.EventIdList) {
		request.EventIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EventIdList, dara.String("EventIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterTableList) {
		request.FilterTableListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTableList, dara.String("FilterTableList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterTypeList) {
		request.FilterTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTypeList, dara.String("FilterTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnFilterShrink) {
		query["ColumnFilter"] = request.ColumnFilterShrink
	}

	if !dara.IsNil(request.EventIdListShrink) {
		query["EventIdList"] = request.EventIdListShrink
	}

	if !dara.IsNil(request.FilterEndTime) {
		query["FilterEndTime"] = request.FilterEndTime
	}

	if !dara.IsNil(request.FilterStartTime) {
		query["FilterStartTime"] = request.FilterStartTime
	}

	if !dara.IsNil(request.FilterTableListShrink) {
		query["FilterTableList"] = request.FilterTableListShrink
	}

	if !dara.IsNil(request.FilterTypeListShrink) {
		query["FilterTypeList"] = request.FilterTypeListShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RollbackSQLType) {
		query["RollbackSQLType"] = request.RollbackSQLType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DownloadDataTrackResult"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DownloadDataTrackResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads the parsing result of a data tracking task.
//
// @param request - DownloadDataTrackResultRequest
//
// @return DownloadDataTrackResultResponse
func (client *Client) DownloadDataTrackResult(request *DownloadDataTrackResultRequest) (_result *DownloadDataTrackResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DownloadDataTrackResultResponse{}
	_body, _err := client.DownloadDataTrackResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a logical database.
//
// @param tmpReq - EditLogicDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditLogicDatabaseResponse
func (client *Client) EditLogicDatabaseWithOptions(tmpReq *EditLogicDatabaseRequest, runtime *dara.RuntimeOptions) (_result *EditLogicDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EditLogicDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DatabaseIds) {
		request.DatabaseIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DatabaseIds, dara.String("DatabaseIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		query["Alias"] = request.Alias
	}

	if !dara.IsNil(request.DatabaseIdsShrink) {
		query["DatabaseIds"] = request.DatabaseIdsShrink
	}

	if !dara.IsNil(request.LogicDbId) {
		query["LogicDbId"] = request.LogicDbId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditLogicDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditLogicDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a logical database.
//
// @param request - EditLogicDatabaseRequest
//
// @return EditLogicDatabaseResponse
func (client *Client) EditLogicDatabase(request *EditLogicDatabaseRequest) (_result *EditLogicDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditLogicDatabaseResponse{}
	_body, _err := client.EditLogicDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits the metadata business knowledge for specified GUIDs.
//
// Description:
//
// Edits the business knowledge of the metadata represented by the specified GUID.
//
// @param request - EditMetaKnowledgeAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditMetaKnowledgeAssetResponse
func (client *Client) EditMetaKnowledgeAssetWithOptions(request *EditMetaKnowledgeAssetRequest, runtime *dara.RuntimeOptions) (_result *EditMetaKnowledgeAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssetDescription) {
		query["AssetDescription"] = request.AssetDescription
	}

	if !dara.IsNil(request.ColumnName) {
		query["ColumnName"] = request.ColumnName
	}

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableSchemaName) {
		query["TableSchemaName"] = request.TableSchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditMetaKnowledgeAsset"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditMetaKnowledgeAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits the metadata business knowledge for specified GUIDs.
//
// Description:
//
// Edits the business knowledge of the metadata represented by the specified GUID.
//
// @param request - EditMetaKnowledgeAssetRequest
//
// @return EditMetaKnowledgeAssetResponse
func (client *Client) EditMetaKnowledgeAsset(request *EditMetaKnowledgeAssetRequest) (_result *EditMetaKnowledgeAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditMetaKnowledgeAssetResponse{}
	_body, _err := client.EditMetaKnowledgeAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to enable a user that has been disabled in Data Management (DMS) Enterprise.
//
// Description:
//
// The effect of enabling a user by calling this operation is the same as that of enabling a user by choosing System Management > User Management in the DMS Enterprise console. The administrator of DMS Enterprise can call this operation to enable a user that has been disabled in DMS Enterprise. After the user is enabled, the corresponding Alibaba Cloud account or Resource Access Management (RAM) user can continue to log on to DMS Enterprise and perform relevant operations.
//
// >  This operation only enables the Alibaba Cloud account or RAM user to log on to DMS Enterprise of the enterprise and perform relevant operations, rather than granting other permissions to the Alibaba Cloud account or RAM user.
//
// @param request - EnableUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableUserResponse
func (client *Client) EnableUserWithOptions(request *EnableUserRequest, runtime *dara.RuntimeOptions) (_result *EnableUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to enable a user that has been disabled in Data Management (DMS) Enterprise.
//
// Description:
//
// The effect of enabling a user by calling this operation is the same as that of enabling a user by choosing System Management > User Management in the DMS Enterprise console. The administrator of DMS Enterprise can call this operation to enable a user that has been disabled in DMS Enterprise. After the user is enabled, the corresponding Alibaba Cloud account or Resource Access Management (RAM) user can continue to log on to DMS Enterprise and perform relevant operations.
//
// >  This operation only enables the Alibaba Cloud account or RAM user to log on to DMS Enterprise of the enterprise and perform relevant operations, rather than granting other permissions to the Alibaba Cloud account or RAM user.
//
// @param request - EnableUserRequest
//
// @return EnableUserResponse
func (client *Client) EnableUser(request *EnableUserRequest) (_result *EnableUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableUserResponse{}
	_body, _err := client.EnableUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a ticket for data change.
//
// @param tmpReq - ExecuteDataCorrectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteDataCorrectResponse
func (client *Client) ExecuteDataCorrectWithOptions(tmpReq *ExecuteDataCorrectRequest, runtime *dara.RuntimeOptions) (_result *ExecuteDataCorrectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExecuteDataCorrectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionDetail) {
		request.ActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionDetail, dara.String("ActionDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionDetailShrink) {
		query["ActionDetail"] = request.ActionDetailShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteDataCorrect"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteDataCorrectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a ticket for data change.
//
// @param request - ExecuteDataCorrectRequest
//
// @return ExecuteDataCorrectResponse
func (client *Client) ExecuteDataCorrect(request *ExecuteDataCorrectRequest) (_result *ExecuteDataCorrectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteDataCorrectResponse{}
	_body, _err := client.ExecuteDataCorrectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a data export ticket.
//
// @param tmpReq - ExecuteDataExportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteDataExportResponse
func (client *Client) ExecuteDataExportWithOptions(tmpReq *ExecuteDataExportRequest, runtime *dara.RuntimeOptions) (_result *ExecuteDataExportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ExecuteDataExportShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionDetail) {
		request.ActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionDetail, dara.String("ActionDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionDetailShrink) {
		query["ActionDetail"] = request.ActionDetailShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteDataExport"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteDataExportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a data export ticket.
//
// @param request - ExecuteDataExportRequest
//
// @return ExecuteDataExportResponse
func (client *Client) ExecuteDataExport(request *ExecuteDataExportRequest) (_result *ExecuteDataExportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteDataExportResponse{}
	_body, _err := client.ExecuteDataExportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes SQL statements.
//
// Description:
//
// You can call this operation only for instances that are managed in Security Collaboration mode.
//
// @param request - ExecuteScriptRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteScriptResponse
func (client *Client) ExecuteScriptWithOptions(request *ExecuteScriptRequest, runtime *dara.RuntimeOptions) (_result *ExecuteScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Script) {
		query["Script"] = request.Script
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteScript"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes SQL statements.
//
// Description:
//
// You can call this operation only for instances that are managed in Security Collaboration mode.
//
// @param request - ExecuteScriptRequest
//
// @return ExecuteScriptResponse
func (client *Client) ExecuteScript(request *ExecuteScriptRequest) (_result *ExecuteScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteScriptResponse{}
	_body, _err := client.ExecuteScriptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a schema synchronization task.
//
// Description:
//
// If the security rules of an instance indicate that a ticket must be approved before you perform schema synchronization, you can call the [SubmitStructSyncOrderApproval](https://help.aliyun.com/document_detail/206166.html) operation to submit the ticket for approval.
//
// >  You can call the [GetStructSyncJobDetail](https://help.aliyun.com/document_detail/206160.html) operation to query whether you need to submit a ticket for approval.
//
// @param request - ExecuteStructSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteStructSyncResponse
func (client *Client) ExecuteStructSyncWithOptions(request *ExecuteStructSyncRequest, runtime *dara.RuntimeOptions) (_result *ExecuteStructSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteStructSync"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteStructSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a schema synchronization task.
//
// Description:
//
// If the security rules of an instance indicate that a ticket must be approved before you perform schema synchronization, you can call the [SubmitStructSyncOrderApproval](https://help.aliyun.com/document_detail/206166.html) operation to submit the ticket for approval.
//
// >  You can call the [GetStructSyncJobDetail](https://help.aliyun.com/document_detail/206160.html) operation to query whether you need to submit a ticket for approval.
//
// @param request - ExecuteStructSyncRequest
//
// @return ExecuteStructSyncResponse
func (client *Client) ExecuteStructSync(request *ExecuteStructSyncRequest) (_result *ExecuteStructSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteStructSyncResponse{}
	_body, _err := client.ExecuteStructSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses the specified SQL statement, error message, and database ID to analyze and fix the SQL error.
//
// Description:
//
// You can call this operation to analyze and fix SQL errors based on the specified SQL statement, error message, and database ID.
//
// @param request - FixSqlByMetaAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FixSqlByMetaAgentResponse
func (client *Client) FixSqlByMetaAgentWithOptions(request *FixSqlByMetaAgentRequest, runtime *dara.RuntimeOptions) (_result *FixSqlByMetaAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Error) {
		query["Error"] = request.Error
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FixSqlByMetaAgent"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &FixSqlByMetaAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses the specified SQL statement, error message, and database ID to analyze and fix the SQL error.
//
// Description:
//
// You can call this operation to analyze and fix SQL errors based on the specified SQL statement, error message, and database ID.
//
// @param request - FixSqlByMetaAgentRequest
//
// @return FixSqlByMetaAgentResponse
func (client *Client) FixSqlByMetaAgent(request *FixSqlByMetaAgentRequest) (_result *FixSqlByMetaAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &FixSqlByMetaAgentResponse{}
	_body, _err := client.FixSqlByMetaAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 数据库知识库补数据接口
//
// @param request - GenMetaKnowledgeAssetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenMetaKnowledgeAssetResponse
func (client *Client) GenMetaKnowledgeAssetWithOptions(request *GenMetaKnowledgeAssetRequest, runtime *dara.RuntimeOptions) (_result *GenMetaKnowledgeAssetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenMetaKnowledgeAsset"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenMetaKnowledgeAssetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 数据库知识库补数据接口
//
// @param request - GenMetaKnowledgeAssetRequest
//
// @return GenMetaKnowledgeAssetResponse
func (client *Client) GenMetaKnowledgeAsset(request *GenMetaKnowledgeAssetRequest) (_result *GenMetaKnowledgeAssetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenMetaKnowledgeAssetResponse{}
	_body, _err := client.GenMetaKnowledgeAssetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据用户提供的自然语言描述和数据库信息生成对应的SQL语句。
//
// Description:
//
// ## 请求说明
//
// - 该API用于将用户的自然语言问题转换为可执行的SQL查询语句。
//
// - 用户需要提供数据库ID (`dbId`) 和实例ID (`instanceId`)，以及想要查询的问题描述 (`question`)。
//
// - 可选参数包括业务知识 (`knowledge`)、数据库方言 (`dialect`)、历史消息 (`history`)、算法级别 (`level`) 和指定模型名称 (`model`)。
//
// - `dialect` 默认值为 "SQL"，表示生成标准SQL语句；`level` 默认值为 "base"，代表基础级别的处理。
//
// - 如果请求中包含多个表且这些表不属于同一个数据库，则返回的 `dbId` 将为0，表明这是一个跨库查询。
//
// - 返回结果中除了生成的SQL语句外，还包括了求解时用到的表、引用的知识点、展示用的SQL等信息，方便进一步理解和使用。
//
// @param request - GenerateSqlFromNLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateSqlFromNLResponse
func (client *Client) GenerateSqlFromNLWithOptions(request *GenerateSqlFromNLRequest, runtime *dara.RuntimeOptions) (_result *GenerateSqlFromNLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Dialect) {
		query["Dialect"] = request.Dialect
	}

	if !dara.IsNil(request.Knowledge) {
		query["Knowledge"] = request.Knowledge
	}

	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Question) {
		query["Question"] = request.Question
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.TableNames) {
		query["TableNames"] = request.TableNames
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateSqlFromNL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateSqlFromNLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据用户提供的自然语言描述和数据库信息生成对应的SQL语句。
//
// Description:
//
// ## 请求说明
//
// - 该API用于将用户的自然语言问题转换为可执行的SQL查询语句。
//
// - 用户需要提供数据库ID (`dbId`) 和实例ID (`instanceId`)，以及想要查询的问题描述 (`question`)。
//
// - 可选参数包括业务知识 (`knowledge`)、数据库方言 (`dialect`)、历史消息 (`history`)、算法级别 (`level`) 和指定模型名称 (`model`)。
//
// - `dialect` 默认值为 "SQL"，表示生成标准SQL语句；`level` 默认值为 "base"，代表基础级别的处理。
//
// - 如果请求中包含多个表且这些表不属于同一个数据库，则返回的 `dbId` 将为0，表明这是一个跨库查询。
//
// - 返回结果中除了生成的SQL语句外，还包括了求解时用到的表、引用的知识点、展示用的SQL等信息，方便进一步理解和使用。
//
// @param request - GenerateSqlFromNLRequest
//
// @return GenerateSqlFromNLResponse
func (client *Client) GenerateSqlFromNL(request *GenerateSqlFromNLRequest) (_result *GenerateSqlFromNLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateSqlFromNLResponse{}
	_body, _err := client.GenerateSqlFromNLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a policy, including the name, content, and the user who created the policy.
//
// @param request - GetAbacPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAbacPolicyResponse
func (client *Client) GetAbacPolicyWithOptions(request *GetAbacPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetAbacPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbacPolicyId) {
		query["AbacPolicyId"] = request.AbacPolicyId
	}

	if !dara.IsNil(request.AbacPolicyName) {
		query["AbacPolicyName"] = request.AbacPolicyName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAbacPolicy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAbacPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a policy, including the name, content, and the user who created the policy.
//
// @param request - GetAbacPolicyRequest
//
// @return GetAbacPolicyResponse
func (client *Client) GetAbacPolicy(request *GetAbacPolicyRequest) (_result *GetAbacPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAbacPolicyResponse{}
	_body, _err := client.GetAbacPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the approval details of a ticket.
//
// @param request - GetApprovalDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApprovalDetailResponse
func (client *Client) GetApprovalDetailWithOptions(request *GetApprovalDetailRequest, runtime *dara.RuntimeOptions) (_result *GetApprovalDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		query["WorkflowInstanceId"] = request.WorkflowInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApprovalDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApprovalDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the approval details of a ticket.
//
// @param request - GetApprovalDetailRequest
//
// @return GetApprovalDetailResponse
func (client *Client) GetApprovalDetail(request *GetApprovalDetailRequest) (_result *GetApprovalDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetApprovalDetailResponse{}
	_body, _err := client.GetApprovalDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a permission template.
//
// Description:
//
// You must be a Data Management (DMS) administrator or a database administrator (DBA). For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - GetAuthorityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorityTemplateResponse
func (client *Client) GetAuthorityTemplateWithOptions(request *GetAuthorityTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorityTemplate"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a permission template.
//
// Description:
//
// You must be a Data Management (DMS) administrator or a database administrator (DBA). For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - GetAuthorityTemplateRequest
//
// @return GetAuthorityTemplateResponse
func (client *Client) GetAuthorityTemplate(request *GetAuthorityTemplateRequest) (_result *GetAuthorityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuthorityTemplateResponse{}
	_body, _err := client.GetAuthorityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the resources in a permission template.
//
// Description:
//
// You are a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - GetAuthorityTemplateItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorityTemplateItemResponse
func (client *Client) GetAuthorityTemplateItemWithOptions(request *GetAuthorityTemplateItemRequest, runtime *dara.RuntimeOptions) (_result *GetAuthorityTemplateItemResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthorityTemplateItem"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthorityTemplateItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the resources in a permission template.
//
// Description:
//
// You are a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - GetAuthorityTemplateItemRequest
//
// @return GetAuthorityTemplateItemResponse
func (client *Client) GetAuthorityTemplateItem(request *GetAuthorityTemplateItemRequest) (_result *GetAuthorityTemplateItemResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuthorityTemplateItemResponse{}
	_body, _err := client.GetAuthorityTemplateItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the classification and grading template associated with an instance.
//
// @param request - GetClassificationTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClassificationTemplateResponse
func (client *Client) GetClassificationTemplateWithOptions(request *GetClassificationTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetClassificationTemplateResponse, _err error) {
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

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetClassificationTemplate"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetClassificationTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the classification and grading template associated with an instance.
//
// @param request - GetClassificationTemplateRequest
//
// @return GetClassificationTemplateResponse
func (client *Client) GetClassificationTemplate(request *GetClassificationTemplateRequest) (_result *GetClassificationTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetClassificationTemplateResponse{}
	_body, _err := client.GetClassificationTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log that records the scheduling details of an SQL task.
//
// @param request - GetDBTaskSQLJobLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDBTaskSQLJobLogResponse
func (client *Client) GetDBTaskSQLJobLogWithOptions(request *GetDBTaskSQLJobLogRequest, runtime *dara.RuntimeOptions) (_result *GetDBTaskSQLJobLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDBTaskSQLJobLog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDBTaskSQLJobLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log that records the scheduling details of an SQL task.
//
// @param request - GetDBTaskSQLJobLogRequest
//
// @return GetDBTaskSQLJobLogResponse
func (client *Client) GetDBTaskSQLJobLog(request *GetDBTaskSQLJobLogRequest) (_result *GetDBTaskSQLJobLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDBTaskSQLJobLogResponse{}
	_body, _err := client.GetDBTaskSQLJobLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the topologies of a logical database and its physical database shards.
//
// @param request - GetDBTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDBTopologyResponse
func (client *Client) GetDBTopologyWithOptions(request *GetDBTopologyRequest, runtime *dara.RuntimeOptions) (_result *GetDBTopologyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicDbId) {
		query["LogicDbId"] = request.LogicDbId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDBTopology"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDBTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topologies of a logical database and its physical database shards.
//
// @param request - GetDBTopologyRequest
//
// @return GetDBTopologyResponse
func (client *Client) GetDBTopology(request *GetDBTopologyRequest) (_result *GetDBTopologyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDBTopologyResponse{}
	_body, _err := client.GetDBTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of data archiving tickets, including the number of tickets of successful data archiving, failed data archiving,and in-progress data archiving, and the total number of data archiving tickets.
//
// @param request - GetDataArchiveCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataArchiveCountResponse
func (client *Client) GetDataArchiveCountWithOptions(request *GetDataArchiveCountRequest, runtime *dara.RuntimeOptions) (_result *GetDataArchiveCountResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderResultType) {
		query["OrderResultType"] = request.OrderResultType
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.SearchDateType) {
		query["SearchDateType"] = request.SearchDateType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataArchiveCount"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataArchiveCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of data archiving tickets, including the number of tickets of successful data archiving, failed data archiving,and in-progress data archiving, and the total number of data archiving tickets.
//
// @param request - GetDataArchiveCountRequest
//
// @return GetDataArchiveCountResponse
func (client *Client) GetDataArchiveCount(request *GetDataArchiveCountRequest) (_result *GetDataArchiveCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataArchiveCountResponse{}
	_body, _err := client.GetDataArchiveCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data archiving ticket, including the time when the ticket was created, the scheduling information of the data archiving task, the logs of the data archiving task, and the database to which data is archived.
//
// @param request - GetDataArchiveOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataArchiveOrderDetailResponse
func (client *Client) GetDataArchiveOrderDetailWithOptions(request *GetDataArchiveOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDataArchiveOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataArchiveOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataArchiveOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data archiving ticket, including the time when the ticket was created, the scheduling information of the data archiving task, the logs of the data archiving task, and the database to which data is archived.
//
// @param request - GetDataArchiveOrderDetailRequest
//
// @return GetDataArchiveOrderDetailResponse
func (client *Client) GetDataArchiveOrderDetail(request *GetDataArchiveOrderDetailRequest) (_result *GetDataArchiveOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataArchiveOrderDetailResponse{}
	_body, _err := client.GetDataArchiveOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of the backup file for a data change ticket in Data Management (DMS).
//
// @param tmpReq - GetDataCorrectBackupFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCorrectBackupFilesResponse
func (client *Client) GetDataCorrectBackupFilesWithOptions(tmpReq *GetDataCorrectBackupFilesRequest, runtime *dara.RuntimeOptions) (_result *GetDataCorrectBackupFilesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDataCorrectBackupFilesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ActionDetail) {
		request.ActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ActionDetail, dara.String("ActionDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionDetailShrink) {
		query["ActionDetail"] = request.ActionDetailShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCorrectBackupFiles"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCorrectBackupFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of the backup file for a data change ticket in Data Management (DMS).
//
// @param request - GetDataCorrectBackupFilesRequest
//
// @return GetDataCorrectBackupFilesResponse
func (client *Client) GetDataCorrectBackupFiles(request *GetDataCorrectBackupFilesRequest) (_result *GetDataCorrectBackupFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCorrectBackupFilesResponse{}
	_body, _err := client.GetDataCorrectBackupFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a data change ticket in Data Management (DMS).
//
// @param request - GetDataCorrectOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCorrectOrderDetailResponse
func (client *Client) GetDataCorrectOrderDetailWithOptions(request *GetDataCorrectOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDataCorrectOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCorrectOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCorrectOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a data change ticket in Data Management (DMS).
//
// @param request - GetDataCorrectOrderDetailRequest
//
// @return GetDataCorrectOrderDetailResponse
func (client *Client) GetDataCorrectOrderDetail(request *GetDataCorrectOrderDetailRequest) (_result *GetDataCorrectOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCorrectOrderDetailResponse{}
	_body, _err := client.GetDataCorrectOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of the rollback attachment submitted along with a data change ticket.
//
// @param request - GetDataCorrectRollbackFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCorrectRollbackFileResponse
func (client *Client) GetDataCorrectRollbackFileWithOptions(request *GetDataCorrectRollbackFileRequest, runtime *dara.RuntimeOptions) (_result *GetDataCorrectRollbackFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCorrectRollbackFile"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCorrectRollbackFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of the rollback attachment submitted along with a data change ticket.
//
// @param request - GetDataCorrectRollbackFileRequest
//
// @return GetDataCorrectRollbackFileResponse
func (client *Client) GetDataCorrectRollbackFile(request *GetDataCorrectRollbackFileRequest) (_result *GetDataCorrectRollbackFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCorrectRollbackFileResponse{}
	_body, _err := client.GetDataCorrectRollbackFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of the SQL script for a data change task.
//
// Description:
//
// This operation applies to [regular data change](https://help.aliyun.com/document_detail/58419.html) and [batch data import](https://help.aliyun.com/document_detail/144643.html).
//
// @param request - GetDataCorrectSQLFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCorrectSQLFileResponse
func (client *Client) GetDataCorrectSQLFileWithOptions(request *GetDataCorrectSQLFileRequest, runtime *dara.RuntimeOptions) (_result *GetDataCorrectSQLFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCorrectSQLFile"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCorrectSQLFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of the SQL script for a data change task.
//
// Description:
//
// This operation applies to [regular data change](https://help.aliyun.com/document_detail/58419.html) and [batch data import](https://help.aliyun.com/document_detail/144643.html).
//
// @param request - GetDataCorrectSQLFileRequest
//
// @return GetDataCorrectSQLFileResponse
func (client *Client) GetDataCorrectSQLFile(request *GetDataCorrectSQLFileRequest) (_result *GetDataCorrectSQLFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCorrectSQLFileResponse{}
	_body, _err := client.GetDataCorrectSQLFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a data change, lock-free data change, or data import task.
//
// @param request - GetDataCorrectTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCorrectTaskDetailResponse
func (client *Client) GetDataCorrectTaskDetailWithOptions(request *GetDataCorrectTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDataCorrectTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCorrectTaskDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCorrectTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a data change, lock-free data change, or data import task.
//
// @param request - GetDataCorrectTaskDetailRequest
//
// @return GetDataCorrectTaskDetailResponse
func (client *Client) GetDataCorrectTaskDetail(request *GetDataCorrectTaskDetailRequest) (_result *GetDataCorrectTaskDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCorrectTaskDetailResponse{}
	_body, _err := client.GetDataCorrectTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scheduling configuration of a ticket for cleaning up historical data.
//
// @param request - GetDataCronClearConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCronClearConfigResponse
func (client *Client) GetDataCronClearConfigWithOptions(request *GetDataCronClearConfigRequest, runtime *dara.RuntimeOptions) (_result *GetDataCronClearConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCronClearConfig"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCronClearConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scheduling configuration of a ticket for cleaning up historical data.
//
// @param request - GetDataCronClearConfigRequest
//
// @return GetDataCronClearConfigResponse
func (client *Client) GetDataCronClearConfig(request *GetDataCronClearConfigRequest) (_result *GetDataCronClearConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCronClearConfigResponse{}
	_body, _err := client.GetDataCronClearConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a historical data cleansing ticket.
//
// @param request - GetDataCronClearTaskDetailListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataCronClearTaskDetailListResponse
func (client *Client) GetDataCronClearTaskDetailListWithOptions(request *GetDataCronClearTaskDetailListRequest, runtime *dara.RuntimeOptions) (_result *GetDataCronClearTaskDetailListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataCronClearTaskDetailList"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataCronClearTaskDetailListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a historical data cleansing ticket.
//
// @param request - GetDataCronClearTaskDetailListRequest
//
// @return GetDataCronClearTaskDetailListResponse
func (client *Client) GetDataCronClearTaskDetailList(request *GetDataCronClearTaskDetailListRequest) (_result *GetDataCronClearTaskDetailListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataCronClearTaskDetailListResponse{}
	_body, _err := client.GetDataCronClearTaskDetailListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of export results for a data export ticket.
//
// @param request - GetDataExportDownloadURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataExportDownloadURLResponse
func (client *Client) GetDataExportDownloadURLWithOptions(request *GetDataExportDownloadURLRequest, runtime *dara.RuntimeOptions) (_result *GetDataExportDownloadURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataExportDownloadURL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataExportDownloadURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of export results for a data export ticket.
//
// @param request - GetDataExportDownloadURLRequest
//
// @return GetDataExportDownloadURLResponse
func (client *Client) GetDataExportDownloadURL(request *GetDataExportDownloadURLRequest) (_result *GetDataExportDownloadURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataExportDownloadURLResponse{}
	_body, _err := client.GetDataExportDownloadURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a data export ticket.
//
// @param request - GetDataExportOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataExportOrderDetailResponse
func (client *Client) GetDataExportOrderDetailWithOptions(request *GetDataExportOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDataExportOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataExportOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataExportOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a data export ticket.
//
// @param request - GetDataExportOrderDetailRequest
//
// @return GetDataExportOrderDetailResponse
func (client *Client) GetDataExportOrderDetail(request *GetDataExportOrderDetailRequest) (_result *GetDataExportOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataExportOrderDetailResponse{}
	_body, _err := client.GetDataExportOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the precheck details of an SQL result set export ticket.
//
// @param request - GetDataExportPreCheckDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataExportPreCheckDetailResponse
func (client *Client) GetDataExportPreCheckDetailWithOptions(request *GetDataExportPreCheckDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDataExportPreCheckDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataExportPreCheckDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataExportPreCheckDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precheck details of an SQL result set export ticket.
//
// @param request - GetDataExportPreCheckDetailRequest
//
// @return GetDataExportPreCheckDetailResponse
func (client *Client) GetDataExportPreCheckDetail(request *GetDataExportPreCheckDetailRequest) (_result *GetDataExportPreCheckDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataExportPreCheckDetailResponse{}
	_body, _err := client.GetDataExportPreCheckDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SQL statements used for data import in a ticket.
//
// Description:
//
// You can call this operation only if the data is imported in security mode in your data import ticket.
//
// @param request - GetDataImportSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataImportSQLResponse
func (client *Client) GetDataImportSQLWithOptions(request *GetDataImportSQLRequest, runtime *dara.RuntimeOptions) (_result *GetDataImportSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.SqlId) {
		query["SqlId"] = request.SqlId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataImportSQL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataImportSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL statements used for data import in a ticket.
//
// Description:
//
// You can call this operation only if the data is imported in security mode in your data import ticket.
//
// @param request - GetDataImportSQLRequest
//
// @return GetDataImportSQLResponse
func (client *Client) GetDataImportSQL(request *GetDataImportSQLRequest) (_result *GetDataImportSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataImportSQLResponse{}
	_body, _err := client.GetDataImportSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data catalog of the data lake.
//
// @param request - GetDataLakeCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeCatalogResponse
func (client *Client) GetDataLakeCatalogWithOptions(request *GetDataLakeCatalogRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeCatalog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data catalog of the data lake.
//
// @param request - GetDataLakeCatalogRequest
//
// @return GetDataLakeCatalogResponse
func (client *Client) GetDataLakeCatalog(request *GetDataLakeCatalogRequest) (_result *GetDataLakeCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeCatalogResponse{}
	_body, _err := client.GetDataLakeCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the database in the data lake.
//
// @param request - GetDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeDatabaseResponse
func (client *Client) GetDataLakeDatabaseWithOptions(request *GetDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the database in the data lake.
//
// @param request - GetDataLakeDatabaseRequest
//
// @return GetDataLakeDatabaseResponse
func (client *Client) GetDataLakeDatabase(request *GetDataLakeDatabaseRequest) (_result *GetDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeDatabaseResponse{}
	_body, _err := client.GetDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details about a user-defined function in a data lakehouse.
//
// @param request - GetDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeFunctionResponse
func (client *Client) GetDataLakeFunctionWithOptions(request *GetDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeFunction"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details about a user-defined function in a data lakehouse.
//
// @param request - GetDataLakeFunctionRequest
//
// @return GetDataLakeFunctionResponse
func (client *Client) GetDataLakeFunction(request *GetDataLakeFunctionRequest) (_result *GetDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeFunctionResponse{}
	_body, _err := client.GetDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a partition of a table in a data lakehouse.
//
// @param tmpReq - GetDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakePartitionResponse
func (client *Client) GetDataLakePartitionWithOptions(tmpReq *GetDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionValues) {
		request.PartitionValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionValues, dara.String("PartitionValues"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.PartitionValuesShrink) {
		query["PartitionValues"] = request.PartitionValuesShrink
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakePartition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a partition of a table in a data lakehouse.
//
// @param request - GetDataLakePartitionRequest
//
// @return GetDataLakePartitionResponse
func (client *Client) GetDataLakePartition(request *GetDataLakePartitionRequest) (_result *GetDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakePartitionResponse{}
	_body, _err := client.GetDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries basic information about tables in the data lake.
//
// @param request - GetDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataLakeTableResponse
func (client *Client) GetDataLakeTableWithOptions(request *GetDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *GetDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataLakeTable"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries basic information about tables in the data lake.
//
// @param request - GetDataLakeTableRequest
//
// @return GetDataLakeTableResponse
func (client *Client) GetDataLakeTable(request *GetDataLakeTableRequest) (_result *GetDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataLakeTableResponse{}
	_body, _err := client.GetDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress of a data tracking task.
//
// @param request - GetDataTrackJobDegreeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataTrackJobDegreeResponse
func (client *Client) GetDataTrackJobDegreeWithOptions(request *GetDataTrackJobDegreeRequest, runtime *dara.RuntimeOptions) (_result *GetDataTrackJobDegreeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataTrackJobDegree"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataTrackJobDegreeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of a data tracking task.
//
// @param request - GetDataTrackJobDegreeRequest
//
// @return GetDataTrackJobDegreeResponse
func (client *Client) GetDataTrackJobDegree(request *GetDataTrackJobDegreeRequest) (_result *GetDataTrackJobDegreeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataTrackJobDegreeResponse{}
	_body, _err := client.GetDataTrackJobDegreeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of tables involved in a data tracking task.
//
// @param request - GetDataTrackJobTableMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataTrackJobTableMetaResponse
func (client *Client) GetDataTrackJobTableMetaWithOptions(request *GetDataTrackJobTableMetaRequest, runtime *dara.RuntimeOptions) (_result *GetDataTrackJobTableMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataTrackJobTableMeta"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataTrackJobTableMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of tables involved in a data tracking task.
//
// @param request - GetDataTrackJobTableMetaRequest
//
// @return GetDataTrackJobTableMetaResponse
func (client *Client) GetDataTrackJobTableMeta(request *GetDataTrackJobTableMetaRequest) (_result *GetDataTrackJobTableMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataTrackJobTableMetaResponse{}
	_body, _err := client.GetDataTrackJobTableMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a data tracking ticket.
//
// @param request - GetDataTrackOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataTrackOrderDetailResponse
func (client *Client) GetDataTrackOrderDetailWithOptions(request *GetDataTrackOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDataTrackOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataTrackOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataTrackOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a data tracking ticket.
//
// @param request - GetDataTrackOrderDetailRequest
//
// @return GetDataTrackOrderDetailResponse
func (client *Client) GetDataTrackOrderDetail(request *GetDataTrackOrderDetailRequest) (_result *GetDataTrackOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataTrackOrderDetailResponse{}
	_body, _err := client.GetDataTrackOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of a specific database.
//
// @param request - GetDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseResponse
func (client *Client) GetDatabaseWithOptions(request *GetDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details of a specific database.
//
// @param request - GetDatabaseRequest
//
// @return GetDatabaseResponse
func (client *Client) GetDatabase(request *GetDatabaseRequest) (_result *GetDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a database export ticket.
//
// @param request - GetDatabaseExportOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatabaseExportOrderDetailResponse
func (client *Client) GetDatabaseExportOrderDetailWithOptions(request *GetDatabaseExportOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetDatabaseExportOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		body["OrderId"] = request.OrderId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDatabaseExportOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatabaseExportOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a database export ticket.
//
// @param request - GetDatabaseExportOrderDetailRequest
//
// @return GetDatabaseExportOrderDetailResponse
func (client *Client) GetDatabaseExportOrderDetail(request *GetDatabaseExportOrderDetailRequest) (_result *GetDatabaseExportOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDatabaseExportOrderDetailResponse{}
	_body, _err := client.GetDatabaseExportOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of the results of a database export ticket.
//
// @param request - GetDbExportDownloadURLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDbExportDownloadURLResponse
func (client *Client) GetDbExportDownloadURLWithOptions(request *GetDbExportDownloadURLRequest, runtime *dara.RuntimeOptions) (_result *GetDbExportDownloadURLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDbExportDownloadURL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDbExportDownloadURLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of the results of a database export ticket.
//
// @param request - GetDbExportDownloadURLRequest
//
// @return GetDbExportDownloadURLResponse
func (client *Client) GetDbExportDownloadURL(request *GetDbExportDownloadURLRequest) (_result *GetDbExportDownloadURLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDbExportDownloadURLResponse{}
	_body, _err := client.GetDbExportDownloadURLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a database instance.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a database instance.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the minimum scheduling cycle of a task flow when a service level agreement (SLA) timeout rule is configured for the task flow.
//
// Description:
//
// The scheduling cycle of a task flow must be greater than the minimum scheduling cycle configured in the SLA rule for the task flow.
//
// @param request - GetIntervalLimitOfSLARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIntervalLimitOfSLAResponse
func (client *Client) GetIntervalLimitOfSLAWithOptions(request *GetIntervalLimitOfSLARequest, runtime *dara.RuntimeOptions) (_result *GetIntervalLimitOfSLAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIntervalLimitOfSLA"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIntervalLimitOfSLAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the minimum scheduling cycle of a task flow when a service level agreement (SLA) timeout rule is configured for the task flow.
//
// Description:
//
// The scheduling cycle of a task flow must be greater than the minimum scheduling cycle configured in the SLA rule for the task flow.
//
// @param request - GetIntervalLimitOfSLARequest
//
// @return GetIntervalLimitOfSLAResponse
func (client *Client) GetIntervalLimitOfSLA(request *GetIntervalLimitOfSLARequest) (_result *GetIntervalLimitOfSLAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIntervalLimitOfSLAResponse{}
	_body, _err := client.GetIntervalLimitOfSLAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a workspace based on the workspace name in Data Management (DMS).
//
// Description:
//
// You are a DMS administrator or a database administrator (DBA).
//
// @param request - GetLhSpaceByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLhSpaceByNameResponse
func (client *Client) GetLhSpaceByNameWithOptions(request *GetLhSpaceByNameRequest, runtime *dara.RuntimeOptions) (_result *GetLhSpaceByNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceName) {
		query["SpaceName"] = request.SpaceName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLhSpaceByName"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLhSpaceByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workspace based on the workspace name in Data Management (DMS).
//
// Description:
//
// You are a DMS administrator or a database administrator (DBA).
//
// @param request - GetLhSpaceByNameRequest
//
// @return GetLhSpaceByNameResponse
func (client *Client) GetLhSpaceByName(request *GetLhSpaceByNameRequest) (_result *GetLhSpaceByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLhSpaceByNameResponse{}
	_body, _err := client.GetLhSpaceByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a logical database.
//
// @param request - GetLogicDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLogicDatabaseResponse
func (client *Client) GetLogicDatabaseWithOptions(request *GetLogicDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetLogicDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLogicDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLogicDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a logical database.
//
// @param request - GetLogicDatabaseRequest
//
// @return GetLogicDatabaseResponse
func (client *Client) GetLogicDatabase(request *GetLogicDatabaseRequest) (_result *GetLogicDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetLogicDatabaseResponse{}
	_body, _err := client.GetLogicDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of columns in a table.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
//
// @param request - GetMetaTableColumnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableColumnResponse
func (client *Client) GetMetaTableColumnWithOptions(request *GetMetaTableColumnRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableColumnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableColumn"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableColumnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of columns in a table.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
//
// @param request - GetMetaTableColumnRequest
//
// @return GetMetaTableColumnResponse
func (client *Client) GetMetaTableColumn(request *GetMetaTableColumnRequest) (_result *GetMetaTableColumnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableColumnResponse{}
	_body, _err := client.GetMetaTableColumnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of columns and indexes in a table.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
//
// @param request - GetMetaTableDetailInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableDetailInfoResponse
func (client *Client) GetMetaTableDetailInfoWithOptions(request *GetMetaTableDetailInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableDetailInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableDetailInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableDetailInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of columns and indexes in a table.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
//
// @param request - GetMetaTableDetailInfoRequest
//
// @return GetMetaTableDetailInfoResponse
func (client *Client) GetMetaTableDetailInfo(request *GetMetaTableDetailInfoRequest) (_result *GetMetaTableDetailInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableDetailInfoResponse{}
	_body, _err := client.GetMetaTableDetailInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details and progress of an OnlineDDL task.
//
// @param request - GetOnlineDDLProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOnlineDDLProgressResponse
func (client *Client) GetOnlineDDLProgressWithOptions(request *GetOnlineDDLProgressRequest, runtime *dara.RuntimeOptions) (_result *GetOnlineDDLProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobDetailId) {
		query["JobDetailId"] = request.JobDetailId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOnlineDDLProgress"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOnlineDDLProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details and progress of an OnlineDDL task.
//
// @param request - GetOnlineDDLProgressRequest
//
// @return GetOnlineDDLProgressResponse
func (client *Client) GetOnlineDDLProgress(request *GetOnlineDDLProgressRequest) (_result *GetOnlineDDLProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOnlineDDLProgressResponse{}
	_body, _err := client.GetOnlineDDLProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of operation logs that are generated in a specified period of time.
//
// Description:
//
// Prerequisites: You are an administrator of Data Management (DMS) or a security administrator. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain your user role from the RoleIdList parameter that is returned.
//
// @param request - GetOpLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpLogResponse
func (client *Client) GetOpLogWithOptions(request *GetOpLogRequest, runtime *dara.RuntimeOptions) (_result *GetOpLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.Module) {
		query["Module"] = request.Module
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserNick) {
		query["UserNick"] = request.UserNick
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpLog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of operation logs that are generated in a specified period of time.
//
// Description:
//
// Prerequisites: You are an administrator of Data Management (DMS) or a security administrator. You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to obtain your user role from the RoleIdList parameter that is returned.
//
// @param request - GetOpLogRequest
//
// @return GetOpLogResponse
func (client *Client) GetOpLog(request *GetOpLogRequest) (_result *GetOpLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOpLogResponse{}
	_body, _err := client.GetOpLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of the attachment of a ticket.
//
// @param request - GetOrderAttachmentFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderAttachmentFileResponse
func (client *Client) GetOrderAttachmentFileWithOptions(request *GetOrderAttachmentFileRequest, runtime *dara.RuntimeOptions) (_result *GetOrderAttachmentFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrderAttachmentFile"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrderAttachmentFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of the attachment of a ticket.
//
// @param request - GetOrderAttachmentFileRequest
//
// @return GetOrderAttachmentFileResponse
func (client *Client) GetOrderAttachmentFile(request *GetOrderAttachmentFileRequest) (_result *GetOrderAttachmentFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOrderAttachmentFileResponse{}
	_body, _err := client.GetOrderAttachmentFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about a ticket.
//
// @param request - GetOrderBaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrderBaseInfoResponse
func (client *Client) GetOrderBaseInfoWithOptions(request *GetOrderBaseInfoRequest, runtime *dara.RuntimeOptions) (_result *GetOrderBaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrderBaseInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrderBaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about a ticket.
//
// @param request - GetOrderBaseInfoRequest
//
// @return GetOrderBaseInfoResponse
func (client *Client) GetOrderBaseInfo(request *GetOrderBaseInfoRequest) (_result *GetOrderBaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOrderBaseInfoResponse{}
	_body, _err := client.GetOrderBaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the Database-OWNER, Table-OWNER, and Instance-OWNER tickets.
//
// @param request - GetOwnerApplyOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOwnerApplyOrderDetailResponse
func (client *Client) GetOwnerApplyOrderDetailWithOptions(request *GetOwnerApplyOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetOwnerApplyOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOwnerApplyOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOwnerApplyOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the Database-OWNER, Table-OWNER, and Instance-OWNER tickets.
//
// @param request - GetOwnerApplyOrderDetailRequest
//
// @return GetOwnerApplyOrderDetailResponse
func (client *Client) GetOwnerApplyOrderDetail(request *GetOwnerApplyOrderDetailRequest) (_result *GetOwnerApplyOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOwnerApplyOrderDetailResponse{}
	_body, _err := client.GetOwnerApplyOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the task information about a ticket.
//
// @param request - GetPagedInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPagedInstanceResponse
func (client *Client) GetPagedInstanceWithOptions(request *GetPagedInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetPagedInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPagedInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPagedInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task information about a ticket.
//
// @param request - GetPagedInstanceRequest
//
// @return GetPagedInstanceResponse
func (client *Client) GetPagedInstance(request *GetPagedInstanceRequest) (_result *GetPagedInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPagedInstanceResponse{}
	_body, _err := client.GetPagedInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a ticket that applies for permissions.
//
// Description:
//
// You can call this operation to query the information about tickets that apply for permissions on databases, tables, and sensitive columns.
//
// @param request - GetPermApplyOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermApplyOrderDetailResponse
func (client *Client) GetPermApplyOrderDetailWithOptions(request *GetPermApplyOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetPermApplyOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPermApplyOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPermApplyOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a ticket that applies for permissions.
//
// Description:
//
// You can call this operation to query the information about tickets that apply for permissions on databases, tables, and sensitive columns.
//
// @param request - GetPermApplyOrderDetailRequest
//
// @return GetPermApplyOrderDetailResponse
func (client *Client) GetPermApplyOrderDetail(request *GetPermApplyOrderDetailRequest) (_result *GetPermApplyOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPermApplyOrderDetailResponse{}
	_body, _err := client.GetPermApplyOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a physical database.
//
// @param request - GetPhysicalDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPhysicalDatabaseResponse
func (client *Client) GetPhysicalDatabaseWithOptions(request *GetPhysicalDatabaseRequest, runtime *dara.RuntimeOptions) (_result *GetPhysicalDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPhysicalDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPhysicalDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a physical database.
//
// @param request - GetPhysicalDatabaseRequest
//
// @return GetPhysicalDatabaseResponse
func (client *Client) GetPhysicalDatabase(request *GetPhysicalDatabaseRequest) (_result *GetPhysicalDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPhysicalDatabaseResponse{}
	_body, _err := client.GetPhysicalDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a secure access proxy.
//
// @param request - GetProxyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProxyResponse
func (client *Client) GetProxyWithOptions(request *GetProxyRequest, runtime *dara.RuntimeOptions) (_result *GetProxyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProxy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a secure access proxy.
//
// @param request - GetProxyRequest
//
// @return GetProxyResponse
func (client *Client) GetProxy(request *GetProxyRequest) (_result *GetProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProxyResponse{}
	_body, _err := client.GetProxyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the authorization information about the secure access proxy feature.
//
// @param request - GetProxyAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProxyAccessResponse
func (client *Client) GetProxyAccessWithOptions(request *GetProxyAccessRequest, runtime *dara.RuntimeOptions) (_result *GetProxyAccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyAccessId) {
		query["ProxyAccessId"] = request.ProxyAccessId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProxyAccess"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProxyAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the authorization information about the secure access proxy feature.
//
// @param request - GetProxyAccessRequest
//
// @return GetProxyAccessResponse
func (client *Client) GetProxyAccess(request *GetProxyAccessRequest) (_result *GetProxyAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProxyAccessResponse{}
	_body, _err := client.GetProxyAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the maximum number of custom service-level agreement (SLA) rules for task orchestration.
//
// @param request - GetRuleNumLimitOfSLARequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRuleNumLimitOfSLAResponse
func (client *Client) GetRuleNumLimitOfSLAWithOptions(request *GetRuleNumLimitOfSLARequest, runtime *dara.RuntimeOptions) (_result *GetRuleNumLimitOfSLAResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRuleNumLimitOfSLA"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRuleNumLimitOfSLAResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the maximum number of custom service-level agreement (SLA) rules for task orchestration.
//
// @param request - GetRuleNumLimitOfSLARequest
//
// @return GetRuleNumLimitOfSLAResponse
func (client *Client) GetRuleNumLimitOfSLA(request *GetRuleNumLimitOfSLARequest) (_result *GetRuleNumLimitOfSLAResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRuleNumLimitOfSLAResponse{}
	_body, _err := client.GetRuleNumLimitOfSLAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of an SQL review.
//
// Description:
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param request - GetSQLReviewCheckResultStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSQLReviewCheckResultStatusResponse
func (client *Client) GetSQLReviewCheckResultStatusWithOptions(request *GetSQLReviewCheckResultStatusRequest, runtime *dara.RuntimeOptions) (_result *GetSQLReviewCheckResultStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSQLReviewCheckResultStatus"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSQLReviewCheckResultStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of an SQL review.
//
// Description:
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param request - GetSQLReviewCheckResultStatusRequest
//
// @return GetSQLReviewCheckResultStatusResponse
func (client *Client) GetSQLReviewCheckResultStatus(request *GetSQLReviewCheckResultStatusRequest) (_result *GetSQLReviewCheckResultStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSQLReviewCheckResultStatusResponse{}
	_body, _err := client.GetSQLReviewCheckResultStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of optimization suggestions. The operation applies only to the tickets for the data changes and SQL review.
//
// Description:
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param request - GetSQLReviewOptimizeDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSQLReviewOptimizeDetailResponse
func (client *Client) GetSQLReviewOptimizeDetailWithOptions(request *GetSQLReviewOptimizeDetailRequest, runtime *dara.RuntimeOptions) (_result *GetSQLReviewOptimizeDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SQLReviewQueryKey) {
		query["SQLReviewQueryKey"] = request.SQLReviewQueryKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSQLReviewOptimizeDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSQLReviewOptimizeDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of optimization suggestions. The operation applies only to the tickets for the data changes and SQL review.
//
// Description:
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param request - GetSQLReviewOptimizeDetailRequest
//
// @return GetSQLReviewOptimizeDetailResponse
func (client *Client) GetSQLReviewOptimizeDetail(request *GetSQLReviewOptimizeDetailRequest) (_result *GetSQLReviewOptimizeDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSQLReviewOptimizeDetailResponse{}
	_body, _err := client.GetSQLReviewOptimizeDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about security rule sets of a tenant. The information includes the name and description of security rules, and instances associated with security rules.
//
// @param request - GetStandardGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStandardGroupResponse
func (client *Client) GetStandardGroupWithOptions(request *GetStandardGroupRequest, runtime *dara.RuntimeOptions) (_result *GetStandardGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStandardGroup"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStandardGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about security rule sets of a tenant. The information includes the name and description of security rules, and instances associated with security rules.
//
// @param request - GetStandardGroupRequest
//
// @return GetStandardGroupResponse
func (client *Client) GetStandardGroup(request *GetStandardGroupRequest) (_result *GetStandardGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStandardGroupResponse{}
	_body, _err := client.GetStandardGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the SQL statements that are involved in a schema synchronization ticket.
//
// @param request - GetStructSyncExecSqlDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStructSyncExecSqlDetailResponse
func (client *Client) GetStructSyncExecSqlDetailWithOptions(request *GetStructSyncExecSqlDetailRequest, runtime *dara.RuntimeOptions) (_result *GetStructSyncExecSqlDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStructSyncExecSqlDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStructSyncExecSqlDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the SQL statements that are involved in a schema synchronization ticket.
//
// @param request - GetStructSyncExecSqlDetailRequest
//
// @return GetStructSyncExecSqlDetailResponse
func (client *Client) GetStructSyncExecSqlDetail(request *GetStructSyncExecSqlDetailRequest) (_result *GetStructSyncExecSqlDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStructSyncExecSqlDetailResponse{}
	_body, _err := client.GetStructSyncExecSqlDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the analysis result of a schema synchronization task.
//
// @param request - GetStructSyncJobAnalyzeResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStructSyncJobAnalyzeResultResponse
func (client *Client) GetStructSyncJobAnalyzeResultWithOptions(request *GetStructSyncJobAnalyzeResultRequest, runtime *dara.RuntimeOptions) (_result *GetStructSyncJobAnalyzeResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CompareType) {
		query["CompareType"] = request.CompareType
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStructSyncJobAnalyzeResult"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStructSyncJobAnalyzeResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the analysis result of a schema synchronization task.
//
// @param request - GetStructSyncJobAnalyzeResultRequest
//
// @return GetStructSyncJobAnalyzeResultResponse
func (client *Client) GetStructSyncJobAnalyzeResult(request *GetStructSyncJobAnalyzeResultRequest) (_result *GetStructSyncJobAnalyzeResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStructSyncJobAnalyzeResultResponse{}
	_body, _err := client.GetStructSyncJobAnalyzeResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a schema synchronization task.
//
// @param request - GetStructSyncJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStructSyncJobDetailResponse
func (client *Client) GetStructSyncJobDetailWithOptions(request *GetStructSyncJobDetailRequest, runtime *dara.RuntimeOptions) (_result *GetStructSyncJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStructSyncJobDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStructSyncJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a schema synchronization task.
//
// @param request - GetStructSyncJobDetailRequest
//
// @return GetStructSyncJobDetailResponse
func (client *Client) GetStructSyncJobDetail(request *GetStructSyncJobDetailRequest) (_result *GetStructSyncJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStructSyncJobDetailResponse{}
	_body, _err := client.GetStructSyncJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a schema synchronization ticket.
//
// @param request - GetStructSyncOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStructSyncOrderDetailResponse
func (client *Client) GetStructSyncOrderDetailWithOptions(request *GetStructSyncOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetStructSyncOrderDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStructSyncOrderDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStructSyncOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a schema synchronization ticket.
//
// @param request - GetStructSyncOrderDetailRequest
//
// @return GetStructSyncOrderDetailResponse
func (client *Client) GetStructSyncOrderDetail(request *GetStructSyncOrderDetailRequest) (_result *GetStructSyncOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStructSyncOrderDetailResponse{}
	_body, _err := client.GetStructSyncOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the topology of a data table.
//
// @param request - GetTableDBTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableDBTopologyResponse
func (client *Client) GetTableDBTopologyWithOptions(request *GetTableDBTopologyRequest, runtime *dara.RuntimeOptions) (_result *GetTableDBTopologyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableDBTopology"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableDBTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topology of a data table.
//
// @param request - GetTableDBTopologyRequest
//
// @return GetTableDBTopologyResponse
func (client *Client) GetTableDBTopology(request *GetTableDBTopologyRequest) (_result *GetTableDBTopologyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableDBTopologyResponse{}
	_body, _err := client.GetTableDBTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a schema design ticket, such as the current node of the ticket, whether the ticket can be returned to the schema design node, and the publishing strategy.
//
// @param request - GetTableDesignProjectFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableDesignProjectFlowResponse
func (client *Client) GetTableDesignProjectFlowWithOptions(request *GetTableDesignProjectFlowRequest, runtime *dara.RuntimeOptions) (_result *GetTableDesignProjectFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableDesignProjectFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableDesignProjectFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a schema design ticket, such as the current node of the ticket, whether the ticket can be returned to the schema design node, and the publishing strategy.
//
// @param request - GetTableDesignProjectFlowRequest
//
// @return GetTableDesignProjectFlowResponse
func (client *Client) GetTableDesignProjectFlow(request *GetTableDesignProjectFlowRequest) (_result *GetTableDesignProjectFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableDesignProjectFlowResponse{}
	_body, _err := client.GetTableDesignProjectFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the project and change base database of a schema design ticket, such as the state of the ticket, the ID of the user who created the ticket, and the name and ID of the change base database.
//
// @param request - GetTableDesignProjectInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableDesignProjectInfoResponse
func (client *Client) GetTableDesignProjectInfoWithOptions(request *GetTableDesignProjectInfoRequest, runtime *dara.RuntimeOptions) (_result *GetTableDesignProjectInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableDesignProjectInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableDesignProjectInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the project and change base database of a schema design ticket, such as the state of the ticket, the ID of the user who created the ticket, and the name and ID of the change base database.
//
// @param request - GetTableDesignProjectInfoRequest
//
// @return GetTableDesignProjectInfoResponse
func (client *Client) GetTableDesignProjectInfo(request *GetTableDesignProjectInfoRequest) (_result *GetTableDesignProjectInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableDesignProjectInfoResponse{}
	_body, _err := client.GetTableDesignProjectInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets metadata knowledge for a specified GUID.
//
// Description:
//
// This API can be called only for database instances that are enabled for security hosting.
//
// @param request - GetTableKnowledgeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableKnowledgeInfoResponse
func (client *Client) GetTableKnowledgeInfoWithOptions(request *GetTableKnowledgeInfoRequest, runtime *dara.RuntimeOptions) (_result *GetTableKnowledgeInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableSchemaName) {
		query["TableSchemaName"] = request.TableSchemaName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableKnowledgeInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableKnowledgeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets metadata knowledge for a specified GUID.
//
// Description:
//
// This API can be called only for database instances that are enabled for security hosting.
//
// @param request - GetTableKnowledgeInfoRequest
//
// @return GetTableKnowledgeInfoResponse
func (client *Client) GetTableKnowledgeInfo(request *GetTableKnowledgeInfoRequest) (_result *GetTableKnowledgeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableKnowledgeInfoResponse{}
	_body, _err := client.GetTableKnowledgeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the topology of a table.
//
// @param request - GetTableTopologyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableTopologyResponse
func (client *Client) GetTableTopologyWithOptions(request *GetTableTopologyRequest, runtime *dara.RuntimeOptions) (_result *GetTableTopologyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableTopology"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the topology of a table.
//
// @param request - GetTableTopologyRequest
//
// @return GetTableTopologyResponse
func (client *Client) GetTableTopology(request *GetTableTopologyRequest) (_result *GetTableTopologyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTableTopologyResponse{}
	_body, _err := client.GetTableTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations, including time variables, of a specified task node based on the task node ID.
//
// @param request - GetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResponse
func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *dara.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTask"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations, including time variables, of a specified task node based on the task node ID.
//
// @param request - GetTaskRequest
//
// @return GetTaskResponse
func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Directed Acyclic Graph (DAG) of a task flow.
//
// @param request - GetTaskFlowGraphRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskFlowGraphResponse
func (client *Client) GetTaskFlowGraphWithOptions(request *GetTaskFlowGraphRequest, runtime *dara.RuntimeOptions) (_result *GetTaskFlowGraphResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskFlowGraph"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskFlowGraphResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Directed Acyclic Graph (DAG) of a task flow.
//
// @param request - GetTaskFlowGraphRequest
//
// @return GetTaskFlowGraphResponse
func (client *Client) GetTaskFlowGraph(request *GetTaskFlowGraphRequest) (_result *GetTaskFlowGraphResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskFlowGraphResponse{}
	_body, _err := client.GetTaskFlowGraphWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the notification settings for task flows.
//
// @param request - GetTaskFlowNotificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskFlowNotificationResponse
func (client *Client) GetTaskFlowNotificationWithOptions(request *GetTaskFlowNotificationRequest, runtime *dara.RuntimeOptions) (_result *GetTaskFlowNotificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskFlowNotification"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskFlowNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the notification settings for task flows.
//
// @param request - GetTaskFlowNotificationRequest
//
// @return GetTaskFlowNotificationResponse
func (client *Client) GetTaskFlowNotification(request *GetTaskFlowNotificationRequest) (_result *GetTaskFlowNotificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskFlowNotificationResponse{}
	_body, _err := client.GetTaskFlowNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the nodes in an execution record of a task flow.
//
// @param request - GetTaskInstanceRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInstanceRelationResponse
func (client *Client) GetTaskInstanceRelationWithOptions(request *GetTaskInstanceRelationRequest, runtime *dara.RuntimeOptions) (_result *GetTaskInstanceRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagInstanceId) {
		query["DagInstanceId"] = request.DagInstanceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskInstanceRelation"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskInstanceRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the nodes in an execution record of a task flow.
//
// @param request - GetTaskInstanceRelationRequest
//
// @return GetTaskInstanceRelationResponse
func (client *Client) GetTaskInstanceRelation(request *GetTaskInstanceRelationRequest) (_result *GetTaskInstanceRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTaskInstanceRelationResponse{}
	_body, _err := client.GetTaskInstanceRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a user.
//
// @param request - GetUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *dara.RuntimeOptions) (_result *GetUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details of the active tenant.
//
// @param request - GetUserActiveTenantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserActiveTenantResponse
func (client *Client) GetUserActiveTenantWithOptions(request *GetUserActiveTenantRequest, runtime *dara.RuntimeOptions) (_result *GetUserActiveTenantResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserActiveTenant"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserActiveTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details of the active tenant.
//
// @param request - GetUserActiveTenantRequest
//
// @return GetUserActiveTenantResponse
func (client *Client) GetUserActiveTenant(request *GetUserActiveTenantRequest) (_result *GetUserActiveTenantResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserActiveTenantResponse{}
	_body, _err := client.GetUserActiveTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information of a file upload task.
//
// @param request - GetUserUploadFileJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserUploadFileJobResponse
func (client *Client) GetUserUploadFileJobWithOptions(request *GetUserUploadFileJobRequest, runtime *dara.RuntimeOptions) (_result *GetUserUploadFileJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobKey) {
		query["JobKey"] = request.JobKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUserUploadFileJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUserUploadFileJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information of a file upload task.
//
// @param request - GetUserUploadFileJobRequest
//
// @return GetUserUploadFileJobResponse
func (client *Client) GetUserUploadFileJob(request *GetUserUploadFileJobRequest) (_result *GetUserUploadFileJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetUserUploadFileJobResponse{}
	_body, _err := client.GetUserUploadFileJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a workspace.
//
// @param request - GetWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(request *GetWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a workspace.
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on resources to users by using a permission template.
//
// Description:
//
// You must be a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - GrantTemplateAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantTemplateAuthorityResponse
func (client *Client) GrantTemplateAuthorityWithOptions(request *GrantTemplateAuthorityRequest, runtime *dara.RuntimeOptions) (_result *GrantTemplateAuthorityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ExpireDate) {
		query["ExpireDate"] = request.ExpireDate
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantTemplateAuthority"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantTemplateAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on resources to users by using a permission template.
//
// Description:
//
// You must be a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - GrantTemplateAuthorityRequest
//
// @return GrantTemplateAuthorityResponse
func (client *Client) GrantTemplateAuthority(request *GrantTemplateAuthorityRequest) (_result *GrantTemplateAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantTemplateAuthorityResponse{}
	_body, _err := client.GrantTemplateAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions on an instance, a database, or a table to a user.
//
// @param request - GrantUserPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantUserPermissionResponse
func (client *Client) GrantUserPermissionWithOptions(request *GrantUserPermissionRequest, runtime *dara.RuntimeOptions) (_result *GrantUserPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.DsType) {
		query["DsType"] = request.DsType
	}

	if !dara.IsNil(request.ExpireDate) {
		query["ExpireDate"] = request.ExpireDate
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PermTypes) {
		query["PermTypes"] = request.PermTypes
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantUserPermission"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantUserPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions on an instance, a database, or a table to a user.
//
// @param request - GrantUserPermissionRequest
//
// @return GrantUserPermissionResponse
func (client *Client) GrantUserPermission(request *GrantUserPermissionRequest) (_result *GrantUserPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantUserPermissionResponse{}
	_body, _err := client.GrantUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain the authorization password of the security protection agent InspectProxyAccessSecret.
//
// @param request - InspectProxyAccessSecretRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InspectProxyAccessSecretResponse
func (client *Client) InspectProxyAccessSecretWithOptions(request *InspectProxyAccessSecretRequest, runtime *dara.RuntimeOptions) (_result *InspectProxyAccessSecretResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyAccessId) {
		query["ProxyAccessId"] = request.ProxyAccessId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InspectProxyAccessSecret"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &InspectProxyAccessSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to obtain the authorization password of the security protection agent InspectProxyAccessSecret.
//
// @param request - InspectProxyAccessSecretRequest
//
// @return InspectProxyAccessSecretResponse
func (client *Client) InspectProxyAccessSecret(request *InspectProxyAccessSecretRequest) (_result *InspectProxyAccessSecretResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &InspectProxyAccessSecretResponse{}
	_body, _err := client.InspectProxyAccessSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users to which the specified policy is attached and the number of the users.
//
// @param request - ListAbacAuthorizationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAbacAuthorizationsResponse
func (client *Client) ListAbacAuthorizationsWithOptions(request *ListAbacAuthorizationsRequest, runtime *dara.RuntimeOptions) (_result *ListAbacAuthorizationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PolicyId) {
		query["PolicyId"] = request.PolicyId
	}

	if !dara.IsNil(request.PolicySource) {
		query["PolicySource"] = request.PolicySource
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAbacAuthorizations"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAbacAuthorizationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of users to which the specified policy is attached and the number of the users.
//
// @param request - ListAbacAuthorizationsRequest
//
// @return ListAbacAuthorizationsResponse
func (client *Client) ListAbacAuthorizations(request *ListAbacAuthorizationsRequest) (_result *ListAbacAuthorizationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAbacAuthorizationsResponse{}
	_body, _err := client.ListAbacAuthorizationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of policies that are created by a user.
//
// @param request - ListAbacPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAbacPoliciesResponse
func (client *Client) ListAbacPoliciesWithOptions(request *ListAbacPoliciesRequest, runtime *dara.RuntimeOptions) (_result *ListAbacPoliciesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAbacPolicies"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAbacPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of policies that are created by a user.
//
// @param request - ListAbacPoliciesRequest
//
// @return ListAbacPoliciesResponse
func (client *Client) ListAbacPolicies(request *ListAbacPoliciesRequest) (_result *ListAbacPoliciesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAbacPoliciesResponse{}
	_body, _err := client.ListAbacPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about permission templates, including the template creator, template name, and total number of templates.
//
// @param request - ListAuthorityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorityTemplateResponse
func (client *Client) ListAuthorityTemplateWithOptions(request *ListAuthorityTemplateRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorityTemplate"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about permission templates, including the template creator, template name, and total number of templates.
//
// @param request - ListAuthorityTemplateRequest
//
// @return ListAuthorityTemplateResponse
func (client *Client) ListAuthorityTemplate(request *ListAuthorityTemplateRequest) (_result *ListAuthorityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthorityTemplateResponse{}
	_body, _err := client.ListAuthorityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases on which the specified user has permissions.
//
// @param request - ListAuthorizedDatabasesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedDatabasesForUserResponse
func (client *Client) ListAuthorizedDatabasesForUserWithOptions(request *ListAuthorizedDatabasesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedDatabasesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedDatabasesForUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedDatabasesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases on which the specified user has permissions.
//
// @param request - ListAuthorizedDatabasesForUserRequest
//
// @return ListAuthorizedDatabasesForUserResponse
func (client *Client) ListAuthorizedDatabasesForUser(request *ListAuthorizedDatabasesForUserRequest) (_result *ListAuthorizedDatabasesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthorizedDatabasesForUserResponse{}
	_body, _err := client.ListAuthorizedDatabasesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instances on which the specified user has permissions.
//
// @param request - ListAuthorizedInstancesForUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedInstancesForUserResponse
func (client *Client) ListAuthorizedInstancesForUserWithOptions(request *ListAuthorizedInstancesForUserRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedInstancesForUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedInstancesForUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedInstancesForUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instances on which the specified user has permissions.
//
// @param request - ListAuthorizedInstancesForUserRequest
//
// @return ListAuthorizedInstancesForUserResponse
func (client *Client) ListAuthorizedInstancesForUser(request *ListAuthorizedInstancesForUserRequest) (_result *ListAuthorizedInstancesForUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthorizedInstancesForUserResponse{}
	_body, _err := client.ListAuthorizedInstancesForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users that have permissions on the specified database.
//
// @param request - ListAuthorizedUsersForDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedUsersForDatabaseResponse
func (client *Client) ListAuthorizedUsersForDatabaseWithOptions(request *ListAuthorizedUsersForDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedUsersForDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedUsersForDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedUsersForDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of users that have permissions on the specified database.
//
// @param request - ListAuthorizedUsersForDatabaseRequest
//
// @return ListAuthorizedUsersForDatabaseResponse
func (client *Client) ListAuthorizedUsersForDatabase(request *ListAuthorizedUsersForDatabaseRequest) (_result *ListAuthorizedUsersForDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthorizedUsersForDatabaseResponse{}
	_body, _err := client.ListAuthorizedUsersForDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users that have permissions on the specified instance.
//
// @param request - ListAuthorizedUsersForInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAuthorizedUsersForInstanceResponse
func (client *Client) ListAuthorizedUsersForInstanceWithOptions(request *ListAuthorizedUsersForInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListAuthorizedUsersForInstanceResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAuthorizedUsersForInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAuthorizedUsersForInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of users that have permissions on the specified instance.
//
// @param request - ListAuthorizedUsersForInstanceRequest
//
// @return ListAuthorizedUsersForInstanceResponse
func (client *Client) ListAuthorizedUsersForInstance(request *ListAuthorizedUsersForInstanceRequest) (_result *ListAuthorizedUsersForInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAuthorizedUsersForInstanceResponse{}
	_body, _err := client.ListAuthorizedUsersForInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the classification templates.
//
// @param request - ListClassificationTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClassificationTemplatesResponse
func (client *Client) ListClassificationTemplatesWithOptions(request *ListClassificationTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListClassificationTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClassificationTemplates"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClassificationTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the classification templates.
//
// @param request - ListClassificationTemplatesRequest
//
// @return ListClassificationTemplatesResponse
func (client *Client) ListClassificationTemplates(request *ListClassificationTemplatesRequest) (_result *ListClassificationTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClassificationTemplatesResponse{}
	_body, _err := client.ListClassificationTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the column information of a table.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
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
	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListColumns"),
		Version:     dara.String("2018-11-01"),
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
// Returns the column information of a table.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
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
// Queries the published versions of a task flow.
//
// @param request - ListDAGVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDAGVersionsResponse
func (client *Client) ListDAGVersionsWithOptions(request *ListDAGVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListDAGVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDAGVersions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDAGVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the published versions of a task flow.
//
// @param request - ListDAGVersionsRequest
//
// @return ListDAGVersionsResponse
func (client *Client) ListDAGVersions(request *ListDAGVersionsRequest) (_result *ListDAGVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDAGVersionsResponse{}
	_body, _err := client.ListDAGVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of SQL tasks.
//
// @param request - ListDBTaskSQLJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDBTaskSQLJobResponse
func (client *Client) ListDBTaskSQLJobWithOptions(request *ListDBTaskSQLJobRequest, runtime *dara.RuntimeOptions) (_result *ListDBTaskSQLJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DBTaskGroupId) {
		query["DBTaskGroupId"] = request.DBTaskGroupId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDBTaskSQLJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDBTaskSQLJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of SQL tasks.
//
// @param request - ListDBTaskSQLJobRequest
//
// @return ListDBTaskSQLJobResponse
func (client *Client) ListDBTaskSQLJob(request *ListDBTaskSQLJobRequest) (_result *ListDBTaskSQLJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDBTaskSQLJobResponse{}
	_body, _err := client.ListDBTaskSQLJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an SQL task.
//
// @param request - ListDBTaskSQLJobDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDBTaskSQLJobDetailResponse
func (client *Client) ListDBTaskSQLJobDetailWithOptions(request *ListDBTaskSQLJobDetailRequest, runtime *dara.RuntimeOptions) (_result *ListDBTaskSQLJobDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDBTaskSQLJobDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDBTaskSQLJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an SQL task.
//
// @param request - ListDBTaskSQLJobDetailRequest
//
// @return ListDBTaskSQLJobDetailResponse
func (client *Client) ListDBTaskSQLJobDetail(request *ListDBTaskSQLJobDetailRequest) (_result *ListDBTaskSQLJobDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDBTaskSQLJobDetailResponse{}
	_body, _err := client.ListDBTaskSQLJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the publishing details of a schema design ticket.
//
// @param request - ListDDLPublishRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDDLPublishRecordsResponse
func (client *Client) ListDDLPublishRecordsWithOptions(request *ListDDLPublishRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListDDLPublishRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDDLPublishRecords"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDDLPublishRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the publishing details of a schema design ticket.
//
// @param request - ListDDLPublishRecordsRequest
//
// @return ListDDLPublishRecordsResponse
func (client *Client) ListDDLPublishRecords(request *ListDDLPublishRecordsRequest) (_result *ListDDLPublishRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDDLPublishRecordsResponse{}
	_body, _err := client.ListDDLPublishRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases that are involved in the precheck of a data change ticket.
//
// Description:
//
// For more information about the Normal Data Modify feature, see [Change regular data](https://help.aliyun.com/document_detail/58419.html).
//
// @param request - ListDataCorrectPreCheckDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataCorrectPreCheckDBResponse
func (client *Client) ListDataCorrectPreCheckDBWithOptions(request *ListDataCorrectPreCheckDBRequest, runtime *dara.RuntimeOptions) (_result *ListDataCorrectPreCheckDBResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataCorrectPreCheckDB"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataCorrectPreCheckDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases that are involved in the precheck of a data change ticket.
//
// Description:
//
// For more information about the Normal Data Modify feature, see [Change regular data](https://help.aliyun.com/document_detail/58419.html).
//
// @param request - ListDataCorrectPreCheckDBRequest
//
// @return ListDataCorrectPreCheckDBResponse
func (client *Client) ListDataCorrectPreCheckDB(request *ListDataCorrectPreCheckDBRequest) (_result *ListDataCorrectPreCheckDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataCorrectPreCheckDBResponse{}
	_body, _err := client.ListDataCorrectPreCheckDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the precheck information about an SQL statement that is specified in a data change ticket.
//
// Description:
//
// For more information about the Normal Data Modify feature, see [Change regular data](https://help.aliyun.com/document_detail/58419.html).
//
// @param request - ListDataCorrectPreCheckSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataCorrectPreCheckSQLResponse
func (client *Client) ListDataCorrectPreCheckSQLWithOptions(request *ListDataCorrectPreCheckSQLRequest, runtime *dara.RuntimeOptions) (_result *ListDataCorrectPreCheckSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataCorrectPreCheckSQL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataCorrectPreCheckSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precheck information about an SQL statement that is specified in a data change ticket.
//
// Description:
//
// For more information about the Normal Data Modify feature, see [Change regular data](https://help.aliyun.com/document_detail/58419.html).
//
// @param request - ListDataCorrectPreCheckSQLRequest
//
// @return ListDataCorrectPreCheckSQLResponse
func (client *Client) ListDataCorrectPreCheckSQL(request *ListDataCorrectPreCheckSQLRequest) (_result *ListDataCorrectPreCheckSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataCorrectPreCheckSQLResponse{}
	_body, _err := client.ListDataCorrectPreCheckSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the precheck information of SQL statements used for data import in a ticket.
//
// Description:
//
// You can call this operation only if the data is imported in security mode in your data import ticket.
//
// @param request - ListDataImportSQLPreCheckDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataImportSQLPreCheckDetailResponse
func (client *Client) ListDataImportSQLPreCheckDetailWithOptions(request *ListDataImportSQLPreCheckDetailRequest, runtime *dara.RuntimeOptions) (_result *ListDataImportSQLPreCheckDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.PageNumer) {
		query["PageNumer"] = request.PageNumer
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.StatusCode) {
		query["StatusCode"] = request.StatusCode
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataImportSQLPreCheckDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataImportSQLPreCheckDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the precheck information of SQL statements used for data import in a ticket.
//
// Description:
//
// You can call this operation only if the data is imported in security mode in your data import ticket.
//
// @param request - ListDataImportSQLPreCheckDetailRequest
//
// @return ListDataImportSQLPreCheckDetailResponse
func (client *Client) ListDataImportSQLPreCheckDetail(request *ListDataImportSQLPreCheckDetailRequest) (_result *ListDataImportSQLPreCheckDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataImportSQLPreCheckDetailResponse{}
	_body, _err := client.ListDataImportSQLPreCheckDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the types of SQL statements used for data import in a ticket.
//
// Description:
//
// You can call this operation only if the data is imported in security mode in your data import ticket.
//
// @param request - ListDataImportSQLTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataImportSQLTypeResponse
func (client *Client) ListDataImportSQLTypeWithOptions(request *ListDataImportSQLTypeRequest, runtime *dara.RuntimeOptions) (_result *ListDataImportSQLTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataImportSQLType"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataImportSQLTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the types of SQL statements used for data import in a ticket.
//
// Description:
//
// You can call this operation only if the data is imported in security mode in your data import ticket.
//
// @param request - ListDataImportSQLTypeRequest
//
// @return ListDataImportSQLTypeResponse
func (client *Client) ListDataImportSQLType(request *ListDataImportSQLTypeRequest) (_result *ListDataImportSQLTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataImportSQLTypeResponse{}
	_body, _err := client.ListDataImportSQLTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data catalogs in a data lake.
//
// @param request - ListDataLakeCatalogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeCatalogResponse
func (client *Client) ListDataLakeCatalogWithOptions(request *ListDataLakeCatalogRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeCatalogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeCatalog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data catalogs in a data lake.
//
// @param request - ListDataLakeCatalogRequest
//
// @return ListDataLakeCatalogResponse
func (client *Client) ListDataLakeCatalog(request *ListDataLakeCatalogRequest) (_result *ListDataLakeCatalogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeCatalogResponse{}
	_body, _err := client.ListDataLakeCatalogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of databases in the data lake.
//
// @param request - ListDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeDatabaseResponse
func (client *Client) ListDataLakeDatabaseWithOptions(request *ListDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of databases in the data lake.
//
// @param request - ListDataLakeDatabaseRequest
//
// @return ListDataLakeDatabaseResponse
func (client *Client) ListDataLakeDatabase(request *ListDataLakeDatabaseRequest) (_result *ListDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeDatabaseResponse{}
	_body, _err := client.ListDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of functions in a data lake.
//
// @param request - ListDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeFunctionResponse
func (client *Client) ListDataLakeFunctionWithOptions(request *ListDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionNamePattern) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeFunction"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of functions in a data lake.
//
// @param request - ListDataLakeFunctionRequest
//
// @return ListDataLakeFunctionResponse
func (client *Client) ListDataLakeFunction(request *ListDataLakeFunctionRequest) (_result *ListDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeFunctionResponse{}
	_body, _err := client.ListDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets a list of data lake function names.
//
// @param request - ListDataLakeFunctionNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeFunctionNameResponse
func (client *Client) ListDataLakeFunctionNameWithOptions(request *ListDataLakeFunctionNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeFunctionNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionNamePattern) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeFunctionName"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeFunctionNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets a list of data lake function names.
//
// @param request - ListDataLakeFunctionNameRequest
//
// @return ListDataLakeFunctionNameResponse
func (client *Client) ListDataLakeFunctionName(request *ListDataLakeFunctionNameRequest) (_result *ListDataLakeFunctionNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeFunctionNameResponse{}
	_body, _err := client.ListDataLakeFunctionNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of partitions of a table in a data lake.
//
// @param tmpReq - ListDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionResponse
func (client *Client) ListDataLakePartitionWithOptions(tmpReq *ListDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartNames) {
		request.PartNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartNames, dara.String("PartNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartNamesShrink) {
		body["PartNames"] = request.PartNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of partitions of a table in a data lake.
//
// @param request - ListDataLakePartitionRequest
//
// @return ListDataLakePartitionResponse
func (client *Client) ListDataLakePartition(request *ListDataLakePartitionRequest) (_result *ListDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakePartitionResponse{}
	_body, _err := client.ListDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of partitions of a table in a data lake based on filter conditions.
//
// @param request - ListDataLakePartitionByFilterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionByFilterResponse
func (client *Client) ListDataLakePartitionByFilterWithOptions(request *ListDataLakePartitionByFilterRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionByFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartitionByFilter"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionByFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of partitions of a table in a data lake based on filter conditions.
//
// @param request - ListDataLakePartitionByFilterRequest
//
// @return ListDataLakePartitionByFilterResponse
func (client *Client) ListDataLakePartitionByFilter(request *ListDataLakePartitionByFilterRequest) (_result *ListDataLakePartitionByFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakePartitionByFilterResponse{}
	_body, _err := client.ListDataLakePartitionByFilterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of partition names of a table in a data lake.
//
// @param request - ListDataLakePartitionNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakePartitionNameResponse
func (client *Client) ListDataLakePartitionNameWithOptions(request *ListDataLakePartitionNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakePartitionNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakePartitionName"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakePartitionNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of partition names of a table in a data lake.
//
// @param request - ListDataLakePartitionNameRequest
//
// @return ListDataLakePartitionNameResponse
func (client *Client) ListDataLakePartitionName(request *ListDataLakePartitionNameRequest) (_result *ListDataLakePartitionNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakePartitionNameResponse{}
	_body, _err := client.ListDataLakePartitionNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tables in a data lake.
//
// @param request - ListDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTableResponse
func (client *Client) ListDataLakeTableWithOptions(request *ListDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTable"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tables in a data lake.
//
// @param request - ListDataLakeTableRequest
//
// @return ListDataLakeTableResponse
func (client *Client) ListDataLakeTable(request *ListDataLakeTableRequest) (_result *ListDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeTableResponse{}
	_body, _err := client.ListDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of names of tables in a data lake.
//
// @param request - ListDataLakeTableNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTableNameResponse
func (client *Client) ListDataLakeTableNameWithOptions(request *ListDataLakeTableNameRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTableNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TableNamePattern) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !dara.IsNil(request.TableType) {
		query["TableType"] = request.TableType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTableName"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTableNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of names of tables in a data lake.
//
// @param request - ListDataLakeTableNameRequest
//
// @return ListDataLakeTableNameResponse
func (client *Client) ListDataLakeTableName(request *ListDataLakeTableNameRequest) (_result *ListDataLakeTableNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeTableNameResponse{}
	_body, _err := client.ListDataLakeTableNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries basic information about tables in the data lake.
//
// @param request - ListDataLakeTablebaseInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataLakeTablebaseInfoResponse
func (client *Client) ListDataLakeTablebaseInfoWithOptions(request *ListDataLakeTablebaseInfoRequest, runtime *dara.RuntimeOptions) (_result *ListDataLakeTablebaseInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
	}

	if !dara.IsNil(request.Rows) {
		query["Rows"] = request.Rows
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataLakeTablebaseInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataLakeTablebaseInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries basic information about tables in the data lake.
//
// @param request - ListDataLakeTablebaseInfoRequest
//
// @return ListDataLakeTablebaseInfoResponse
func (client *Client) ListDataLakeTablebaseInfo(request *ListDataLakeTablebaseInfoRequest) (_result *ListDataLakeTablebaseInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataLakeTablebaseInfoResponse{}
	_body, _err := client.ListDataLakeTablebaseInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of users on a database.
//
// @param request - ListDatabaseUserPermssionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabaseUserPermssionsResponse
func (client *Client) ListDatabaseUserPermssionsWithOptions(request *ListDatabaseUserPermssionsRequest, runtime *dara.RuntimeOptions) (_result *ListDatabaseUserPermssionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PermType) {
		query["PermType"] = request.PermType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabaseUserPermssions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabaseUserPermssionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of users on a database.
//
// @param request - ListDatabaseUserPermssionsRequest
//
// @return ListDatabaseUserPermssionsResponse
func (client *Client) ListDatabaseUserPermssions(request *ListDatabaseUserPermssionsRequest) (_result *ListDatabaseUserPermssionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabaseUserPermssionsResponse{}
	_body, _err := client.ListDatabaseUserPermssionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the databases in a database instance.
//
// @param request - ListDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatabasesResponse
func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDatabases"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the databases in a database instance.
//
// @param request - ListDatabasesRequest
//
// @return ListDatabasesResponse
func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default service level agreement (SLA) timeout rules.
//
// @param request - ListDefaultSLARulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDefaultSLARulesResponse
func (client *Client) ListDefaultSLARulesWithOptions(request *ListDefaultSLARulesRequest, runtime *dara.RuntimeOptions) (_result *ListDefaultSLARulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDefaultSLARules"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDefaultSLARulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default service level agreement (SLA) timeout rules.
//
// @param request - ListDefaultSLARulesRequest
//
// @return ListDefaultSLARulesResponse
func (client *Client) ListDefaultSLARules(request *ListDefaultSLARulesRequest) (_result *ListDefaultSLARulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDefaultSLARulesResponse{}
	_body, _err := client.ListDefaultSLARulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries masking rules.
//
// @param request - ListDesensitizationRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDesensitizationRuleResponse
func (client *Client) ListDesensitizationRuleWithOptions(request *ListDesensitizationRuleRequest, runtime *dara.RuntimeOptions) (_result *ListDesensitizationRuleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FuncType) {
		query["FuncType"] = request.FuncType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.RuleName) {
		query["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		query["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDesensitizationRule"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDesensitizationRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries masking rules.
//
// @param request - ListDesensitizationRuleRequest
//
// @return ListDesensitizationRuleResponse
func (client *Client) ListDesensitizationRule(request *ListDesensitizationRuleRequest) (_result *ListDesensitizationRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDesensitizationRuleResponse{}
	_body, _err := client.ListDesensitizationRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - ListDifyInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDifyInstancesResponse
func (client *Client) ListDifyInstancesWithOptions(request *ListDifyInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListDifyInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
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
		Action:      dara.String("ListDifyInstances"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDifyInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建DIFY实例及相关资源，支持自定义配置。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - ListDifyInstancesRequest
//
// @return ListDifyInstancesResponse
func (client *Client) ListDifyInstances(request *ListDifyInstancesRequest) (_result *ListDifyInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDifyInstancesResponse{}
	_body, _err := client.ListDifyInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries valid orders.
//
// @param request - ListEffectiveOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEffectiveOrdersResponse
func (client *Client) ListEffectiveOrdersWithOptions(request *ListEffectiveOrdersRequest, runtime *dara.RuntimeOptions) (_result *ListEffectiveOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEffectiveOrders"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEffectiveOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries valid orders.
//
// @param request - ListEffectiveOrdersRequest
//
// @return ListEffectiveOrdersResponse
func (client *Client) ListEffectiveOrders(request *ListEffectiveOrdersRequest) (_result *ListEffectiveOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEffectiveOrdersResponse{}
	_body, _err := client.ListEffectiveOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the indexes of a table.
//
// @param request - ListIndexesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIndexesResponse
func (client *Client) ListIndexesWithOptions(request *ListIndexesRequest, runtime *dara.RuntimeOptions) (_result *ListIndexesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListIndexes"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIndexesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the indexes of a table.
//
// @param request - ListIndexesRequest
//
// @return ListIndexesResponse
func (client *Client) ListIndexes(request *ListIndexesRequest) (_result *ListIndexesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListIndexesResponse{}
	_body, _err := client.ListIndexesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logon records of an instance.
//
// @param request - ListInstanceLoginAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceLoginAuditLogResponse
func (client *Client) ListInstanceLoginAuditLogWithOptions(request *ListInstanceLoginAuditLogRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceLoginAuditLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OpUserName) {
		query["OpUserName"] = request.OpUserName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceLoginAuditLog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceLoginAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logon records of an instance.
//
// @param request - ListInstanceLoginAuditLogRequest
//
// @return ListInstanceLoginAuditLogResponse
func (client *Client) ListInstanceLoginAuditLog(request *ListInstanceLoginAuditLogRequest) (_result *ListInstanceLoginAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceLoginAuditLogResponse{}
	_body, _err := client.ListInstanceLoginAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of a user on a specific instance.
//
// @param request - ListInstanceUserPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceUserPermissionsResponse
func (client *Client) ListInstanceUserPermissionsWithOptions(request *ListInstanceUserPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceUserPermissionsResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserName) {
		query["UserName"] = request.UserName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceUserPermissions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceUserPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of a user on a specific instance.
//
// @param request - ListInstanceUserPermissionsRequest
//
// @return ListInstanceUserPermissionsResponse
func (client *Client) ListInstanceUserPermissions(request *ListInstanceUserPermissionsRequest) (_result *ListInstanceUserPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceUserPermissionsResponse{}
	_body, _err := client.ListInstanceUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about database instances.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.InstanceSource) {
		query["InstanceSource"] = request.InstanceSource
	}

	if !dara.IsNil(request.InstanceState) {
		query["InstanceState"] = request.InstanceState
	}

	if !dara.IsNil(request.NetType) {
		query["NetType"] = request.NetType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about database instances.
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the task flows corresponding to a specific business scenario in a workspace in Data Management (DMS).
//
// Description:
//
//	  Before you call this operation, make sure that you have the access permissions on the workspace. If you do not have the access permissions on the workspace, you can contact a DMS administrator, database administrator (DBA), or workspace administrator to add you as a member of the workspace. The [AddLhMembers](https://help.aliyun.com/document_detail/424759.html) operation can be called to add a workspace member.
//
//		- If you are a DMS administrator or a workspace administrator, you can query the business scenarios and task flows related to a user in a workspace based on the user ID.
//
// @param request - ListLhTaskFlowAndScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLhTaskFlowAndScenarioResponse
func (client *Client) ListLhTaskFlowAndScenarioWithOptions(request *ListLhTaskFlowAndScenarioRequest, runtime *dara.RuntimeOptions) (_result *ListLhTaskFlowAndScenarioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SpaceId) {
		query["SpaceId"] = request.SpaceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLhTaskFlowAndScenario"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLhTaskFlowAndScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the task flows corresponding to a specific business scenario in a workspace in Data Management (DMS).
//
// Description:
//
//	  Before you call this operation, make sure that you have the access permissions on the workspace. If you do not have the access permissions on the workspace, you can contact a DMS administrator, database administrator (DBA), or workspace administrator to add you as a member of the workspace. The [AddLhMembers](https://help.aliyun.com/document_detail/424759.html) operation can be called to add a workspace member.
//
//		- If you are a DMS administrator or a workspace administrator, you can query the business scenarios and task flows related to a user in a workspace based on the user ID.
//
// @param request - ListLhTaskFlowAndScenarioRequest
//
// @return ListLhTaskFlowAndScenarioResponse
func (client *Client) ListLhTaskFlowAndScenario(request *ListLhTaskFlowAndScenarioRequest) (_result *ListLhTaskFlowAndScenarioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLhTaskFlowAndScenarioResponse{}
	_body, _err := client.ListLhTaskFlowAndScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of logical databases.
//
// @param request - ListLogicDatabasesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogicDatabasesResponse
func (client *Client) ListLogicDatabasesWithOptions(request *ListLogicDatabasesRequest, runtime *dara.RuntimeOptions) (_result *ListLogicDatabasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogicDatabases"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogicDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of logical databases.
//
// @param request - ListLogicDatabasesRequest
//
// @return ListLogicDatabasesResponse
func (client *Client) ListLogicDatabases(request *ListLogicDatabasesRequest) (_result *ListLogicDatabasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogicDatabasesResponse{}
	_body, _err := client.ListLogicDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the routing algorithms of a logical table.
//
// @param request - ListLogicTableRouteConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogicTableRouteConfigResponse
func (client *Client) ListLogicTableRouteConfigWithOptions(request *ListLogicTableRouteConfigRequest, runtime *dara.RuntimeOptions) (_result *ListLogicTableRouteConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogicTableRouteConfig"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogicTableRouteConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the routing algorithms of a logical table.
//
// @param request - ListLogicTableRouteConfigRequest
//
// @return ListLogicTableRouteConfigResponse
func (client *Client) ListLogicTableRouteConfig(request *ListLogicTableRouteConfigRequest) (_result *ListLogicTableRouteConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogicTableRouteConfigResponse{}
	_body, _err := client.ListLogicTableRouteConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the logical tables in a logical database.
//
// @param request - ListLogicTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLogicTablesResponse
func (client *Client) ListLogicTablesWithOptions(request *ListLogicTablesRequest, runtime *dara.RuntimeOptions) (_result *ListLogicTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReturnGuid) {
		query["ReturnGuid"] = request.ReturnGuid
	}

	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLogicTables"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLogicTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the logical tables in a logical database.
//
// @param request - ListLogicTablesRequest
//
// @return ListLogicTablesResponse
func (client *Client) ListLogicTables(request *ListLogicTablesRequest) (_result *ListLogicTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLogicTablesResponse{}
	_body, _err := client.ListLogicTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Queries a list of categories
//
// @param request - ListMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaCategoryResponse
func (client *Client) ListMetaCategoryWithOptions(request *ListMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *ListMetaCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentCategoryId) {
		query["ParentCategoryId"] = request.ParentCategoryId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetaCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Queries a list of categories
//
// @param request - ListMetaCategoryRequest
//
// @return ListMetaCategoryResponse
func (client *Client) ListMetaCategory(request *ListMetaCategoryRequest) (_result *ListMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMetaCategoryResponse{}
	_body, _err := client.ListMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tickets in Data Management (DMS).
//
// @param request - ListOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOrdersResponse
func (client *Client) ListOrdersWithOptions(request *ListOrdersRequest, runtime *dara.RuntimeOptions) (_result *ListOrdersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.OrderResultType) {
		query["OrderResultType"] = request.OrderResultType
	}

	if !dara.IsNil(request.OrderStatus) {
		query["OrderStatus"] = request.OrderStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PluginType) {
		query["PluginType"] = request.PluginType
	}

	if !dara.IsNil(request.SearchContent) {
		query["SearchContent"] = request.SearchContent
	}

	if !dara.IsNil(request.SearchDateType) {
		query["SearchDateType"] = request.SearchDateType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOrders"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tickets in Data Management (DMS).
//
// @param request - ListOrdersRequest
//
// @return ListOrdersResponse
func (client *Client) ListOrders(request *ListOrdersRequest) (_result *ListOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListOrdersResponse{}
	_body, _err := client.ListOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the proxies that are generated by the secure access proxy feature.
//
// @param request - ListProxiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProxiesResponse
func (client *Client) ListProxiesWithOptions(request *ListProxiesRequest, runtime *dara.RuntimeOptions) (_result *ListProxiesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProxies"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProxiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the proxies that are generated by the secure access proxy feature.
//
// @param request - ListProxiesRequest
//
// @return ListProxiesResponse
func (client *Client) ListProxies(request *ListProxiesRequest) (_result *ListProxiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProxiesResponse{}
	_body, _err := client.ListProxiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about users that are authorized to access a database instance by using the secure access proxy feature.
//
// @param request - ListProxyAccessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProxyAccessesResponse
func (client *Client) ListProxyAccessesWithOptions(request *ListProxyAccessesRequest, runtime *dara.RuntimeOptions) (_result *ListProxyAccessesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProxyId) {
		query["ProxyId"] = request.ProxyId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProxyAccesses"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProxyAccessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about users that are authorized to access a database instance by using the secure access proxy feature.
//
// @param request - ListProxyAccessesRequest
//
// @return ListProxyAccessesResponse
func (client *Client) ListProxyAccesses(request *ListProxyAccessesRequest) (_result *ListProxyAccessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProxyAccessesResponse{}
	_body, _err := client.ListProxyAccessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 操作审计-数据安全代理SQL执行列表
//
// @param request - ListProxySQLExecAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProxySQLExecAuditLogResponse
func (client *Client) ListProxySQLExecAuditLogWithOptions(request *ListProxySQLExecAuditLogRequest, runtime *dara.RuntimeOptions) (_result *ListProxySQLExecAuditLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecState) {
		query["ExecState"] = request.ExecState
	}

	if !dara.IsNil(request.OpUserName) {
		query["OpUserName"] = request.OpUserName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SQLType) {
		query["SQLType"] = request.SQLType
	}

	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProxySQLExecAuditLog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProxySQLExecAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作审计-数据安全代理SQL执行列表
//
// @param request - ListProxySQLExecAuditLogRequest
//
// @return ListProxySQLExecAuditLogResponse
func (client *Client) ListProxySQLExecAuditLog(request *ListProxySQLExecAuditLogRequest) (_result *ListProxySQLExecAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProxySQLExecAuditLogResponse{}
	_body, _err := client.ListProxySQLExecAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the custom service level agreement (SLA) rules.
//
// @param request - ListSLARulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSLARulesResponse
func (client *Client) ListSLARulesWithOptions(request *ListSLARulesRequest, runtime *dara.RuntimeOptions) (_result *ListSLARulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSLARules"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSLARulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the custom service level agreement (SLA) rules.
//
// @param request - ListSLARulesRequest
//
// @return ListSLARulesResponse
func (client *Client) ListSLARules(request *ListSLARulesRequest) (_result *ListSLARulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSLARulesResponse{}
	_body, _err := client.ListSLARulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries SQL statements that were written on the SQLConsole tab.
//
// @param request - ListSQLExecAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSQLExecAuditLogResponse
func (client *Client) ListSQLExecAuditLogWithOptions(request *ListSQLExecAuditLogRequest, runtime *dara.RuntimeOptions) (_result *ListSQLExecAuditLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ExecState) {
		query["ExecState"] = request.ExecState
	}

	if !dara.IsNil(request.OpUserName) {
		query["OpUserName"] = request.OpUserName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.SqlType) {
		query["SqlType"] = request.SqlType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSQLExecAuditLog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSQLExecAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries SQL statements that were written on the SQLConsole tab.
//
// @param request - ListSQLExecAuditLogRequest
//
// @return ListSQLExecAuditLogResponse
func (client *Client) ListSQLExecAuditLog(request *ListSQLExecAuditLogRequest) (_result *ListSQLExecAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSQLExecAuditLogResponse{}
	_body, _err := client.ListSQLExecAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the SQL statements that are involved in an SQL review ticket.
//
// Description:
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param tmpReq - ListSQLReviewOriginSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSQLReviewOriginSQLResponse
func (client *Client) ListSQLReviewOriginSQLWithOptions(tmpReq *ListSQLReviewOriginSQLRequest, runtime *dara.RuntimeOptions) (_result *ListSQLReviewOriginSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListSQLReviewOriginSQLShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OrderActionDetail) {
		request.OrderActionDetailShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderActionDetail, dara.String("OrderActionDetail"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderActionDetailShrink) {
		query["OrderActionDetail"] = request.OrderActionDetailShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSQLReviewOriginSQL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSQLReviewOriginSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the SQL statements that are involved in an SQL review ticket.
//
// Description:
//
// For more information about the SQL review feature, see [SQL review](https://help.aliyun.com/document_detail/60374.html).
//
// @param request - ListSQLReviewOriginSQLRequest
//
// @return ListSQLReviewOriginSQLResponse
func (client *Client) ListSQLReviewOriginSQL(request *ListSQLReviewOriginSQLRequest) (_result *ListSQLReviewOriginSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSQLReviewOriginSQLResponse{}
	_body, _err := client.ListSQLReviewOriginSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries business scenarios.
//
// @param request - ListScenariosRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScenariosResponse
func (client *Client) ListScenariosWithOptions(request *ListScenariosRequest, runtime *dara.RuntimeOptions) (_result *ListScenariosResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListScenarios"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListScenariosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries business scenarios.
//
// @param request - ListScenariosRequest
//
// @return ListScenariosResponse
func (client *Client) ListScenarios(request *ListScenariosRequest) (_result *ListScenariosResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListScenariosResponse{}
	_body, _err := client.ListScenariosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a sensitive field. The information includes the masking algorithm ID, masking algorithm name, and field security level.
//
// @param request - ListSensitiveColumnInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSensitiveColumnInfoResponse
func (client *Client) ListSensitiveColumnInfoWithOptions(request *ListSensitiveColumnInfoRequest, runtime *dara.RuntimeOptions) (_result *ListSensitiveColumnInfoResponse, _err error) {
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

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSensitiveColumnInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSensitiveColumnInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a sensitive field. The information includes the masking algorithm ID, masking algorithm name, and field security level.
//
// @param request - ListSensitiveColumnInfoRequest
//
// @return ListSensitiveColumnInfoResponse
func (client *Client) ListSensitiveColumnInfo(request *ListSensitiveColumnInfoRequest) (_result *ListSensitiveColumnInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSensitiveColumnInfoResponse{}
	_body, _err := client.ListSensitiveColumnInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries sensitive fields in a table of a database.
//
// @param request - ListSensitiveColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSensitiveColumnsResponse
func (client *Client) ListSensitiveColumnsWithOptions(request *ListSensitiveColumnsRequest, runtime *dara.RuntimeOptions) (_result *ListSensitiveColumnsResponse, _err error) {
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

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.SecurityLevel) {
		query["SecurityLevel"] = request.SecurityLevel
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSensitiveColumns"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSensitiveColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sensitive fields in a table of a database.
//
// @param request - ListSensitiveColumnsRequest
//
// @return ListSensitiveColumnsResponse
func (client *Client) ListSensitiveColumns(request *ListSensitiveColumnsRequest) (_result *ListSensitiveColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSensitiveColumnsResponse{}
	_body, _err := client.ListSensitiveColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a sensitive field.
//
// @param request - ListSensitiveColumnsDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSensitiveColumnsDetailResponse
func (client *Client) ListSensitiveColumnsDetailWithOptions(request *ListSensitiveColumnsDetailRequest, runtime *dara.RuntimeOptions) (_result *ListSensitiveColumnsDetailResponse, _err error) {
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

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSensitiveColumnsDetail"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSensitiveColumnsDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a sensitive field.
//
// @param request - ListSensitiveColumnsDetailRequest
//
// @return ListSensitiveColumnsDetailResponse
func (client *Client) ListSensitiveColumnsDetail(request *ListSensitiveColumnsDetailRequest) (_result *ListSensitiveColumnsDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSensitiveColumnsDetailResponse{}
	_body, _err := client.ListSensitiveColumnsDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the audit logs for sensitive information.
//
// @param request - ListSensitiveDataAuditLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSensitiveDataAuditLogResponse
func (client *Client) ListSensitiveDataAuditLogWithOptions(request *ListSensitiveDataAuditLogRequest, runtime *dara.RuntimeOptions) (_result *ListSensitiveDataAuditLogResponse, _err error) {
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

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ModuleName) {
		query["ModuleName"] = request.ModuleName
	}

	if !dara.IsNil(request.OpUserName) {
		query["OpUserName"] = request.OpUserName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSensitiveDataAuditLog"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSensitiveDataAuditLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the audit logs for sensitive information.
//
// @param request - ListSensitiveDataAuditLogRequest
//
// @return ListSensitiveDataAuditLogResponse
func (client *Client) ListSensitiveDataAuditLog(request *ListSensitiveDataAuditLogRequest) (_result *ListSensitiveDataAuditLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSensitiveDataAuditLogResponse{}
	_body, _err := client.ListSensitiveDataAuditLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the sensitivity levels of a classification template.
//
// @param request - ListSensitivityLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSensitivityLevelResponse
func (client *Client) ListSensitivityLevelWithOptions(request *ListSensitivityLevelRequest, runtime *dara.RuntimeOptions) (_result *ListSensitivityLevelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSensitivityLevel"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSensitivityLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sensitivity levels of a classification template.
//
// @param request - ListSensitivityLevelRequest
//
// @return ListSensitivityLevelResponse
func (client *Client) ListSensitivityLevel(request *ListSensitivityLevelRequest) (_result *ListSensitivityLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSensitivityLevelResponse{}
	_body, _err := client.ListSensitivityLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries security rule sets.
//
// @param request - ListStandardGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStandardGroupsResponse
func (client *Client) ListStandardGroupsWithOptions(request *ListStandardGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListStandardGroupsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStandardGroups"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStandardGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries security rule sets.
//
// @param request - ListStandardGroupsRequest
//
// @return ListStandardGroupsResponse
func (client *Client) ListStandardGroups(request *ListStandardGroupsRequest) (_result *ListStandardGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStandardGroupsResponse{}
	_body, _err := client.ListStandardGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details about fields in a table.
//
// @param request - ListTableColumnsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableColumnsResponse
func (client *Client) ListTableColumnsWithOptions(request *ListTableColumnsRequest, runtime *dara.RuntimeOptions) (_result *ListTableColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableSchemaName) {
		query["TableSchemaName"] = request.TableSchemaName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableColumns"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about fields in a table.
//
// @param request - ListTableColumnsRequest
//
// @return ListTableColumnsResponse
func (client *Client) ListTableColumns(request *ListTableColumnsRequest) (_result *ListTableColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTableColumnsResponse{}
	_body, _err := client.ListTableColumnsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tables in a database.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
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
	if !dara.IsNil(request.DatabaseId) {
		query["DatabaseId"] = request.DatabaseId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.ReturnGuid) {
		query["ReturnGuid"] = request.ReturnGuid
	}

	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2018-11-01"),
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
// Queries the tables in a database.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
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
// Obtains the list of tables in a category.
//
// @param request - ListTablesInCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesInCategoryResponse
func (client *Client) ListTablesInCategoryWithOptions(request *ListTablesInCategoryRequest, runtime *dara.RuntimeOptions) (_result *ListTablesInCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTablesInCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTablesInCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of tables in a category.
//
// @param request - ListTablesInCategoryRequest
//
// @return ListTablesInCategoryResponse
func (client *Client) ListTablesInCategory(request *ListTablesInCategoryRequest) (_result *ListTablesInCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTablesInCategoryResponse{}
	_body, _err := client.ListTablesInCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 任务编排获取任务流列表
//
// @param request - ListTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowResponse
func (client *Client) ListTaskFlowWithOptions(request *ListTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务编排获取任务流列表
//
// @param request - ListTaskFlowRequest
//
// @return ListTaskFlowResponse
func (client *Client) ListTaskFlow(request *ListTaskFlowRequest) (_result *ListTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowResponse{}
	_body, _err := client.ListTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the constants for a task flow.
//
// @param request - ListTaskFlowConstantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowConstantsResponse
func (client *Client) ListTaskFlowConstantsWithOptions(request *ListTaskFlowConstantsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowConstantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlowConstants"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowConstantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the constants for a task flow.
//
// @param request - ListTaskFlowConstantsRequest
//
// @return ListTaskFlowConstantsResponse
func (client *Client) ListTaskFlowConstants(request *ListTaskFlowConstantsRequest) (_result *ListTaskFlowConstantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowConstantsResponse{}
	_body, _err := client.ListTaskFlowConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the users that are involved in a specified task flow.
//
// @param request - ListTaskFlowCooperatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowCooperatorsResponse
func (client *Client) ListTaskFlowCooperatorsWithOptions(request *ListTaskFlowCooperatorsRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowCooperatorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlowCooperators"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowCooperatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the users that are involved in a specified task flow.
//
// @param request - ListTaskFlowCooperatorsRequest
//
// @return ListTaskFlowCooperatorsResponse
func (client *Client) ListTaskFlowCooperators(request *ListTaskFlowCooperatorsRequest) (_result *ListTaskFlowCooperatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowCooperatorsResponse{}
	_body, _err := client.ListTaskFlowCooperatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the edges of the directed acyclic graph (DAG) for a specified task flow based on multiple conditions.
//
// Description:
//
// This operation is used for multi-condition query. You can call this operation to query the edges of a specified task flow that meet all specified conditions.
//
// @param request - ListTaskFlowEdgesByConditionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowEdgesByConditionResponse
func (client *Client) ListTaskFlowEdgesByConditionWithOptions(request *ListTaskFlowEdgesByConditionRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowEdgesByConditionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.NodeEnd) {
		query["NodeEnd"] = request.NodeEnd
	}

	if !dara.IsNil(request.NodeFrom) {
		query["NodeFrom"] = request.NodeFrom
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlowEdgesByCondition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowEdgesByConditionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the edges of the directed acyclic graph (DAG) for a specified task flow based on multiple conditions.
//
// Description:
//
// This operation is used for multi-condition query. You can call this operation to query the edges of a specified task flow that meet all specified conditions.
//
// @param request - ListTaskFlowEdgesByConditionRequest
//
// @return ListTaskFlowEdgesByConditionResponse
func (client *Client) ListTaskFlowEdgesByCondition(request *ListTaskFlowEdgesByConditionRequest) (_result *ListTaskFlowEdgesByConditionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowEdgesByConditionResponse{}
	_body, _err := client.ListTaskFlowEdgesByConditionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution records of a task flow.
//
// @param request - ListTaskFlowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowInstanceResponse
func (client *Client) ListTaskFlowInstanceWithOptions(request *ListTaskFlowInstanceRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTimeBegin) {
		query["StartTimeBegin"] = request.StartTimeBegin
	}

	if !dara.IsNil(request.StartTimeEnd) {
		query["StartTimeEnd"] = request.StartTimeEnd
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	if !dara.IsNil(request.UseBizDate) {
		query["UseBizDate"] = request.UseBizDate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlowInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution records of a task flow.
//
// @param request - ListTaskFlowInstanceRequest
//
// @return ListTaskFlowInstanceResponse
func (client *Client) ListTaskFlowInstance(request *ListTaskFlowInstanceRequest) (_result *ListTaskFlowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowInstanceResponse{}
	_body, _err := client.ListTaskFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the time variables for a task flow.
//
// @param request - ListTaskFlowTimeVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowTimeVariablesResponse
func (client *Client) ListTaskFlowTimeVariablesWithOptions(request *ListTaskFlowTimeVariablesRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowTimeVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlowTimeVariables"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowTimeVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the time variables for a task flow.
//
// @param request - ListTaskFlowTimeVariablesRequest
//
// @return ListTaskFlowTimeVariablesResponse
func (client *Client) ListTaskFlowTimeVariables(request *ListTaskFlowTimeVariablesRequest) (_result *ListTaskFlowTimeVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowTimeVariablesResponse{}
	_body, _err := client.ListTaskFlowTimeVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries task flows by page.
//
// @param tmpReq - ListTaskFlowsByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskFlowsByPageResponse
func (client *Client) ListTaskFlowsByPageWithOptions(tmpReq *ListTaskFlowsByPageRequest, runtime *dara.RuntimeOptions) (_result *ListTaskFlowsByPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListTaskFlowsByPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DagIdList) {
		request.DagIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DagIdList, dara.String("DagIdList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DagIdListShrink) {
		query["DagIdList"] = request.DagIdListShrink
	}

	if !dara.IsNil(request.PageIndex) {
		query["PageIndex"] = request.PageIndex
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTaskFlowsByPage"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTaskFlowsByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries task flows by page.
//
// @param request - ListTaskFlowsByPageRequest
//
// @return ListTaskFlowsByPageResponse
func (client *Client) ListTaskFlowsByPage(request *ListTaskFlowsByPageRequest) (_result *ListTaskFlowsByPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTaskFlowsByPageResponse{}
	_body, _err := client.ListTaskFlowsByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tasks in a specified task flow.
//
// @param request - ListTasksInTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTasksInTaskFlowResponse
func (client *Client) ListTasksInTaskFlowWithOptions(request *ListTasksInTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *ListTasksInTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTasksInTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTasksInTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tasks in a specified task flow.
//
// @param request - ListTasksInTaskFlowRequest
//
// @return ListTasksInTaskFlowResponse
func (client *Client) ListTasksInTaskFlow(request *ListTasksInTaskFlowRequest) (_result *ListTasksInTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTasksInTaskFlowResponse{}
	_body, _err := client.ListTasksInTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Obtains
//
// @param request - ListUserOwnedResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserOwnedResourcesResponse
func (client *Client) ListUserOwnedResourcesWithOptions(request *ListUserOwnedResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListUserOwnedResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.OwnerType) {
		query["OwnerType"] = request.OwnerType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserOwnedResources"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserOwnedResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Obtains
//
// @param request - ListUserOwnedResourcesRequest
//
// @return ListUserOwnedResourcesResponse
func (client *Client) ListUserOwnedResources(request *ListUserOwnedResourcesRequest) (_result *ListUserOwnedResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserOwnedResourcesResponse{}
	_body, _err := client.ListUserOwnedResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the permissions of a specific user on a database or a table.
//
// @param request - ListUserPermissionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPermissionsResponse
func (client *Client) ListUserPermissionsWithOptions(request *ListUserPermissionsRequest, runtime *dara.RuntimeOptions) (_result *ListUserPermissionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PermType) {
		query["PermType"] = request.PermType
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserPermissions"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the permissions of a specific user on a database or a table.
//
// @param request - ListUserPermissionsRequest
//
// @return ListUserPermissionsResponse
func (client *Client) ListUserPermissions(request *ListUserPermissionsRequest) (_result *ListUserPermissionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserPermissionsResponse{}
	_body, _err := client.ListUserPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户租户列表
//
// @param request - ListUserTenantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserTenantsResponse
func (client *Client) ListUserTenantsWithOptions(request *ListUserTenantsRequest, runtime *dara.RuntimeOptions) (_result *ListUserTenantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserTenants"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUserTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户租户列表
//
// @param request - ListUserTenantsRequest
//
// @return ListUserTenantsResponse
func (client *Client) ListUserTenants(request *ListUserTenantsRequest) (_result *ListUserTenantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUserTenantsResponse{}
	_body, _err := client.ListUserTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of users.
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *dara.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserState) {
		query["UserState"] = request.UserState
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListUsers"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of users.
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries approval nodes.
//
// @param request - ListWorkFlowNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkFlowNodesResponse
func (client *Client) ListWorkFlowNodesWithOptions(request *ListWorkFlowNodesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkFlowNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkFlowNodes"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkFlowNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries approval nodes.
//
// @param request - ListWorkFlowNodesRequest
//
// @return ListWorkFlowNodesResponse
func (client *Client) ListWorkFlowNodes(request *ListWorkFlowNodesRequest) (_result *ListWorkFlowNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkFlowNodesResponse{}
	_body, _err := client.ListWorkFlowNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries approval templates.
//
// @param request - ListWorkFlowTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkFlowTemplatesResponse
func (client *Client) ListWorkFlowTemplatesWithOptions(request *ListWorkFlowTemplatesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkFlowTemplatesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SearchName) {
		query["SearchName"] = request.SearchName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkFlowTemplates"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkFlowTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries approval templates.
//
// @param request - ListWorkFlowTemplatesRequest
//
// @return ListWorkFlowTemplatesResponse
func (client *Client) ListWorkFlowTemplates(request *ListWorkFlowTemplatesRequest) (_result *ListWorkFlowTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkFlowTemplatesResponse{}
	_body, _err := client.ListWorkFlowTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workspaces of the tenant.
//
// @param request - ListWorkspacesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AlreadyJoined) {
		query["AlreadyJoined"] = request.AlreadyJoined
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Region) {
		query["Region"] = request.Region
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.ServiceAccountId) {
		query["ServiceAccountId"] = request.ServiceAccountId
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workspaces of the tenant.
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the status of a task flow to "Successful".
//
// @param request - MakeTaskFlowInstanceSuccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MakeTaskFlowInstanceSuccessResponse
func (client *Client) MakeTaskFlowInstanceSuccessWithOptions(request *MakeTaskFlowInstanceSuccessRequest, runtime *dara.RuntimeOptions) (_result *MakeTaskFlowInstanceSuccessResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagInstanceId) {
		query["DagInstanceId"] = request.DagInstanceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MakeTaskFlowInstanceSuccess"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MakeTaskFlowInstanceSuccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the status of a task flow to "Successful".
//
// @param request - MakeTaskFlowInstanceSuccessRequest
//
// @return MakeTaskFlowInstanceSuccessResponse
func (client *Client) MakeTaskFlowInstanceSuccess(request *MakeTaskFlowInstanceSuccessRequest) (_result *MakeTaskFlowInstanceSuccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MakeTaskFlowInstanceSuccessResponse{}
	_body, _err := client.MakeTaskFlowInstanceSuccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the SQL script that is submitted by using a data change ticket.
//
// @param request - ModifyDataCorrectExecSQLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDataCorrectExecSQLResponse
func (client *Client) ModifyDataCorrectExecSQLWithOptions(request *ModifyDataCorrectExecSQLRequest, runtime *dara.RuntimeOptions) (_result *ModifyDataCorrectExecSQLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExecSQL) {
		query["ExecSQL"] = request.ExecSQL
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDataCorrectExecSQL"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDataCorrectExecSQLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the SQL script that is submitted by using a data change ticket.
//
// @param request - ModifyDataCorrectExecSQLRequest
//
// @return ModifyDataCorrectExecSQLResponse
func (client *Client) ModifyDataCorrectExecSQL(request *ModifyDataCorrectExecSQLRequest) (_result *ModifyDataCorrectExecSQLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDataCorrectExecSQLResponse{}
	_body, _err := client.ModifyDataCorrectExecSQLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the masking rule bound to a specified field.
//
// @param request - ModifyDesensitizationStrategyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDesensitizationStrategyResponse
func (client *Client) ModifyDesensitizationStrategyWithOptions(request *ModifyDesensitizationStrategyRequest, runtime *dara.RuntimeOptions) (_result *ModifyDesensitizationStrategyResponse, _err error) {
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

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.IsLogic) {
		query["IsLogic"] = request.IsLogic
	}

	if !dara.IsNil(request.IsReset) {
		query["IsReset"] = request.IsReset
	}

	if !dara.IsNil(request.RuleId) {
		query["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyDesensitizationStrategy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyDesensitizationStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the masking rule bound to a specified field.
//
// @param request - ModifyDesensitizationStrategyRequest
//
// @return ModifyDesensitizationStrategyResponse
func (client *Client) ModifyDesensitizationStrategy(request *ModifyDesensitizationStrategyRequest) (_result *ModifyDesensitizationStrategyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyDesensitizationStrategyResponse{}
	_body, _err := client.ModifyDesensitizationStrategyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic and advanced information about a database instance in Data Management (DMS). The basic information includes the hosting mode of the instance, the instance account username and password, and features that are enabled for the instance. The advanced information includes the environment type, the instance name in DMS, and the timeout period for querying and exporting data.
//
// @param request - ModifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstanceWithOptions(request *ModifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *ModifyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataLinkName) {
		query["DataLinkName"] = request.DataLinkName
	}

	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.DatabaseUser) {
		query["DatabaseUser"] = request.DatabaseUser
	}

	if !dara.IsNil(request.DbaId) {
		query["DbaId"] = request.DbaId
	}

	if !dara.IsNil(request.DdlOnline) {
		query["DdlOnline"] = request.DdlOnline
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.EcsRegion) {
		query["EcsRegion"] = request.EcsRegion
	}

	if !dara.IsNil(request.EnableSellCommon) {
		query["EnableSellCommon"] = request.EnableSellCommon
	}

	if !dara.IsNil(request.EnableSellSitd) {
		query["EnableSellSitd"] = request.EnableSellSitd
	}

	if !dara.IsNil(request.EnableSellStable) {
		query["EnableSellStable"] = request.EnableSellStable
	}

	if !dara.IsNil(request.EnableSellTrust) {
		query["EnableSellTrust"] = request.EnableSellTrust
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ExportTimeout) {
		query["ExportTimeout"] = request.ExportTimeout
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceSource) {
		query["InstanceSource"] = request.InstanceSource
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.QueryTimeout) {
		query["QueryTimeout"] = request.QueryTimeout
	}

	if !dara.IsNil(request.SafeRule) {
		query["SafeRule"] = request.SafeRule
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.SkipTest) {
		query["SkipTest"] = request.SkipTest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UseDsql) {
		query["UseDsql"] = request.UseDsql
	}

	if !dara.IsNil(request.UseSsl) {
		query["UseSsl"] = request.UseSsl
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic and advanced information about a database instance in Data Management (DMS). The basic information includes the hosting mode of the instance, the instance account username and password, and features that are enabled for the instance. The advanced information includes the environment type, the instance name in DMS, and the timeout period for querying and exporting data.
//
// @param request - ModifyInstanceRequest
//
// @return ModifyInstanceResponse
func (client *Client) ModifyInstance(request *ModifyInstanceRequest) (_result *ModifyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ModifyInstanceResponse{}
	_body, _err := client.ModifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Migrates a task flow to a specified business scenario.
//
// @param request - MoveTaskFlowToScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveTaskFlowToScenarioResponse
func (client *Client) MoveTaskFlowToScenarioWithOptions(request *MoveTaskFlowToScenarioRequest, runtime *dara.RuntimeOptions) (_result *MoveTaskFlowToScenarioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveTaskFlowToScenario"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveTaskFlowToScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Migrates a task flow to a specified business scenario.
//
// @param request - MoveTaskFlowToScenarioRequest
//
// @return MoveTaskFlowToScenarioResponse
func (client *Client) MoveTaskFlowToScenario(request *MoveTaskFlowToScenarioRequest) (_result *MoveTaskFlowToScenarioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveTaskFlowToScenarioResponse{}
	_body, _err := client.MoveTaskFlowToScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unpublishes a published task flow.
//
// @param request - OfflineTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineTaskFlowResponse
func (client *Client) OfflineTaskFlowWithOptions(request *OfflineTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *OfflineTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes a published task flow.
//
// @param request - OfflineTaskFlowRequest
//
// @return OfflineTaskFlowResponse
func (client *Client) OfflineTaskFlow(request *OfflineTaskFlowRequest) (_result *OfflineTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineTaskFlowResponse{}
	_body, _err := client.OfflineTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Analyzes and optimizes the SQL statements in the specified database.
//
// Description:
//
// You can call this operation to analyze and optimize the SQL statements in the specified database.
//
// @param request - OptimizeSqlByMetaAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OptimizeSqlByMetaAgentResponse
func (client *Client) OptimizeSqlByMetaAgentWithOptions(request *OptimizeSqlByMetaAgentRequest, runtime *dara.RuntimeOptions) (_result *OptimizeSqlByMetaAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Model) {
		query["Model"] = request.Model
	}

	if !dara.IsNil(request.Query) {
		query["Query"] = request.Query
	}

	if !dara.IsNil(request.Sql) {
		query["Sql"] = request.Sql
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OptimizeSqlByMetaAgent"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OptimizeSqlByMetaAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Analyzes and optimizes the SQL statements in the specified database.
//
// Description:
//
// You can call this operation to analyze and optimize the SQL statements in the specified database.
//
// @param request - OptimizeSqlByMetaAgentRequest
//
// @return OptimizeSqlByMetaAgentResponse
func (client *Client) OptimizeSqlByMetaAgent(request *OptimizeSqlByMetaAgentRequest) (_result *OptimizeSqlByMetaAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OptimizeSqlByMetaAgentResponse{}
	_body, _err := client.OptimizeSqlByMetaAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses a SQL task for data change.
//
// Description:
//
// You can call this operation only for database instances that are managed in Security Collaboration mode.
//
// @param request - PauseDataCorrectSQLJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseDataCorrectSQLJobResponse
func (client *Client) PauseDataCorrectSQLJobWithOptions(request *PauseDataCorrectSQLJobRequest, runtime *dara.RuntimeOptions) (_result *PauseDataCorrectSQLJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseDataCorrectSQLJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseDataCorrectSQLJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a SQL task for data change.
//
// Description:
//
// You can call this operation only for database instances that are managed in Security Collaboration mode.
//
// @param request - PauseDataCorrectSQLJobRequest
//
// @return PauseDataCorrectSQLJobResponse
func (client *Client) PauseDataCorrectSQLJob(request *PauseDataCorrectSQLJobRequest) (_result *PauseDataCorrectSQLJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseDataCorrectSQLJobResponse{}
	_body, _err := client.PauseDataCorrectSQLJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a SQL result set export task.
//
// @param request - PauseDataExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseDataExportJobResponse
func (client *Client) PauseDataExportJobWithOptions(request *PauseDataExportJobRequest, runtime *dara.RuntimeOptions) (_result *PauseDataExportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseDataExportJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseDataExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a SQL result set export task.
//
// @param request - PauseDataExportJobRequest
//
// @return PauseDataExportJobResponse
func (client *Client) PauseDataExportJob(request *PauseDataExportJobRequest) (_result *PauseDataExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseDataExportJobResponse{}
	_body, _err := client.PauseDataExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Previews the approval process of a ticket before the ticket is submitted.
//
// @param request - PreviewWorkflowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewWorkflowResponse
func (client *Client) PreviewWorkflowWithOptions(request *PreviewWorkflowRequest, runtime *dara.RuntimeOptions) (_result *PreviewWorkflowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewWorkflow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Previews the approval process of a ticket before the ticket is submitted.
//
// @param request - PreviewWorkflowRequest
//
// @return PreviewWorkflowResponse
func (client *Client) PreviewWorkflow(request *PreviewWorkflowRequest) (_result *PreviewWorkflowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreviewWorkflowResponse{}
	_body, _err := client.PreviewWorkflowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes and deploys a task flow.
//
// @param request - PublishAndDeployTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishAndDeployTaskFlowResponse
func (client *Client) PublishAndDeployTaskFlowWithOptions(request *PublishAndDeployTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *PublishAndDeployTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.VersionComments) {
		query["VersionComments"] = request.VersionComments
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishAndDeployTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishAndDeployTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes and deploys a task flow.
//
// @param request - PublishAndDeployTaskFlowRequest
//
// @return PublishAndDeployTaskFlowResponse
func (client *Client) PublishAndDeployTaskFlow(request *PublishAndDeployTaskFlowRequest) (_result *PublishAndDeployTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishAndDeployTaskFlowResponse{}
	_body, _err := client.PublishAndDeployTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download and parsing progress of data tracking logs.
//
// @param request - QueryDataTrackResultDownloadStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDataTrackResultDownloadStatusResponse
func (client *Client) QueryDataTrackResultDownloadStatusWithOptions(request *QueryDataTrackResultDownloadStatusRequest, runtime *dara.RuntimeOptions) (_result *QueryDataTrackResultDownloadStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DownloadKeyId) {
		query["DownloadKeyId"] = request.DownloadKeyId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDataTrackResultDownloadStatus"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDataTrackResultDownloadStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download and parsing progress of data tracking logs.
//
// @param request - QueryDataTrackResultDownloadStatusRequest
//
// @return QueryDataTrackResultDownloadStatusResponse
func (client *Client) QueryDataTrackResultDownloadStatus(request *QueryDataTrackResultDownloadStatusRequest) (_result *QueryDataTrackResultDownloadStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDataTrackResultDownloadStatusResponse{}
	_body, _err := client.QueryDataTrackResultDownloadStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 部署任务流的历史版本
//
// @param request - ReDeployLhDagVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReDeployLhDagVersionResponse
func (client *Client) ReDeployLhDagVersionWithOptions(request *ReDeployLhDagVersionRequest, runtime *dara.RuntimeOptions) (_result *ReDeployLhDagVersionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagVersion) {
		query["DagVersion"] = request.DagVersion
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReDeployLhDagVersion"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReDeployLhDagVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 部署任务流的历史版本
//
// @param request - ReDeployLhDagVersionRequest
//
// @return ReDeployLhDagVersionResponse
func (client *Client) ReDeployLhDagVersion(request *ReDeployLhDagVersionRequest) (_result *ReDeployLhDagVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReDeployLhDagVersionResponse{}
	_body, _err := client.ReDeployLhDagVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns a task flow instance.
//
// @param request - ReRunTaskFlowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReRunTaskFlowInstanceResponse
func (client *Client) ReRunTaskFlowInstanceWithOptions(request *ReRunTaskFlowInstanceRequest, runtime *dara.RuntimeOptions) (_result *ReRunTaskFlowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagInstanceId) {
		query["DagInstanceId"] = request.DagInstanceId
	}

	if !dara.IsNil(request.DagVersion) {
		query["DagVersion"] = request.DagVersion
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReRunTaskFlowInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ReRunTaskFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns a task flow instance.
//
// @param request - ReRunTaskFlowInstanceRequest
//
// @return ReRunTaskFlowInstanceResponse
func (client *Client) ReRunTaskFlowInstance(request *ReRunTaskFlowInstanceRequest) (_result *ReRunTaskFlowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ReRunTaskFlowInstanceResponse{}
	_body, _err := client.ReRunTaskFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用于重新部署DIFY实例。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - RedeployDifyInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RedeployDifyInstanceResponse
func (client *Client) RedeployDifyInstanceWithOptions(request *RedeployDifyInstanceRequest, runtime *dara.RuntimeOptions) (_result *RedeployDifyInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RedeployDifyInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RedeployDifyInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于重新部署DIFY实例。
//
// Description:
//
// ## 请求说明
//
// - `workspaceOption` 参数指示是否新建工作空间，默认使用已有工作空间。
//
// - 如果选择新建工作空间 (`CreateNewInstance`)，则必须提供 `workspaceName` 和 `workspaceDescription`。
//
// - `vpcId`, `VSwitchID`, `zoneId`, `regionCode`, `ResourceQuota`, `Replicas`, `storageType`, `dbInstanceClass`, `dbEngineVersion`, `kvstoreEngineVersion` 是必填项。
//
// - 当 `storageType` 为 `oss` 时，需要指定 `ossBucketResourceId` 和 `ossPath`。
//
// - 如果需要新建数据库实例，则必须提供 `instanceAccount` 和 `instancePassword`。
//
// - 预付费模式下，`PayPeriodType` 和 `PayPeriod` 必须填写。
//
// - 可以通过设置 `dryRun` 为 `true` 来执行预检查而不实际创建实例。
//
// @param request - RedeployDifyInstanceRequest
//
// @return RedeployDifyInstanceResponse
func (client *Client) RedeployDifyInstance(request *RedeployDifyInstanceRequest) (_result *RedeployDifyInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RedeployDifyInstanceResponse{}
	_body, _err := client.RedeployDifyInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unsubscribes from a pay-as-you-go Data Management (DMS) resource.
//
// @param request - RefundPayAsYouGoOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RefundPayAsYouGoOrderResponse
func (client *Client) RefundPayAsYouGoOrderWithOptions(request *RefundPayAsYouGoOrderRequest, runtime *dara.RuntimeOptions) (_result *RefundPayAsYouGoOrderResponse, _err error) {
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

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RefundPayAsYouGoOrder"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RefundPayAsYouGoOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unsubscribes from a pay-as-you-go Data Management (DMS) resource.
//
// @param request - RefundPayAsYouGoOrderRequest
//
// @return RefundPayAsYouGoOrderResponse
func (client *Client) RefundPayAsYouGoOrder(request *RefundPayAsYouGoOrderRequest) (_result *RefundPayAsYouGoOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RefundPayAsYouGoOrderResponse{}
	_body, _err := client.RefundPayAsYouGoOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a database instance in Data Management (DMS).
//
// Description:
//
// Prerequisites: You are a DMS administrator or a database administrator (DBA). You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to query your user role from the RoleIdList parameter that is returned.
//
// @param request - RegisterInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterInstanceResponse
func (client *Client) RegisterInstanceWithOptions(request *RegisterInstanceRequest, runtime *dara.RuntimeOptions) (_result *RegisterInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataLinkName) {
		query["DataLinkName"] = request.DataLinkName
	}

	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.DatabaseUser) {
		query["DatabaseUser"] = request.DatabaseUser
	}

	if !dara.IsNil(request.DbaUid) {
		query["DbaUid"] = request.DbaUid
	}

	if !dara.IsNil(request.DbaUidByString) {
		query["DbaUidByString"] = request.DbaUidByString
	}

	if !dara.IsNil(request.DdlOnline) {
		query["DdlOnline"] = request.DdlOnline
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.EcsRegion) {
		query["EcsRegion"] = request.EcsRegion
	}

	if !dara.IsNil(request.EnableSellSitd) {
		query["EnableSellSitd"] = request.EnableSellSitd
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ExportTimeout) {
		query["ExportTimeout"] = request.ExportTimeout
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceSource) {
		query["InstanceSource"] = request.InstanceSource
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.NetworkType) {
		query["NetworkType"] = request.NetworkType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.QueryTimeout) {
		query["QueryTimeout"] = request.QueryTimeout
	}

	if !dara.IsNil(request.ResourceGroup) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	if !dara.IsNil(request.SafeRule) {
		query["SafeRule"] = request.SafeRule
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.SkipTest) {
		query["SkipTest"] = request.SkipTest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UseDsql) {
		query["UseDsql"] = request.UseDsql
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a database instance in Data Management (DMS).
//
// Description:
//
// Prerequisites: You are a DMS administrator or a database administrator (DBA). You can call the [ListUsers](https://help.aliyun.com/document_detail/141938.html) or [GetUser](https://help.aliyun.com/document_detail/147098.html) operation to query your user role from the RoleIdList parameter that is returned.
//
// @param request - RegisterInstanceRequest
//
// @return RegisterInstanceResponse
func (client *Client) RegisterInstance(request *RegisterInstanceRequest) (_result *RegisterInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterInstanceResponse{}
	_body, _err := client.RegisterInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers a user for your enterprise.
//
// Description:
//
// If you are an **administrator*	- in Data Management (DMS), you can call this operation to register a user for your enterprise. To view users that are assigned the administrator role, perform the following steps: Log on to the DMS console. In the top navigation bar, click O&M. In the left-side navigation pane, click User.
//
// @param request - RegisterUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterUserResponse
func (client *Client) RegisterUserWithOptions(request *RegisterUserRequest, runtime *dara.RuntimeOptions) (_result *RegisterUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.RoleNames) {
		query["RoleNames"] = request.RoleNames
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.UserNick) {
		query["UserNick"] = request.UserNick
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers a user for your enterprise.
//
// Description:
//
// If you are an **administrator*	- in Data Management (DMS), you can call this operation to register a user for your enterprise. To view users that are assigned the administrator role, perform the following steps: Log on to the DMS console. In the top navigation bar, click O&M. In the left-side navigation pane, click User.
//
// @param request - RegisterUserRequest
//
// @return RegisterUserResponse
func (client *Client) RegisterUser(request *RegisterUserRequest) (_result *RegisterUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterUserResponse{}
	_body, _err := client.RegisterUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data export task.
//
// @param request - RemoveDataExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataExportJobResponse
func (client *Client) RemoveDataExportJobWithOptions(request *RemoveDataExportJobRequest, runtime *dara.RuntimeOptions) (_result *RemoveDataExportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveDataExportJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveDataExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data export task.
//
// @param request - RemoveDataExportJobRequest
//
// @return RemoveDataExportJobResponse
func (client *Client) RemoveDataExportJob(request *RemoveDataExportJobRequest) (_result *RemoveDataExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveDataExportJobResponse{}
	_body, _err := client.RemoveDataExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes table from the asset category.
//
// @param request - RemoveTableFromCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTableFromCategoryResponse
func (client *Client) RemoveTableFromCategoryWithOptions(request *RemoveTableFromCategoryRequest, runtime *dara.RuntimeOptions) (_result *RemoveTableFromCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TableSchemaName) {
		query["TableSchemaName"] = request.TableSchemaName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveTableFromCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveTableFromCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes table from the asset category.
//
// @param request - RemoveTableFromCategoryRequest
//
// @return RemoveTableFromCategoryResponse
func (client *Client) RemoveTableFromCategory(request *RemoveTableFromCategoryRequest) (_result *RemoveTableFromCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveTableFromCategoryResponse{}
	_body, _err := client.RemoveTableFromCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns a failed SQL task for data change.
//
// @param request - RestartDataCorrectSQLJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDataCorrectSQLJobResponse
func (client *Client) RestartDataCorrectSQLJobWithOptions(request *RestartDataCorrectSQLJobRequest, runtime *dara.RuntimeOptions) (_result *RestartDataCorrectSQLJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDataCorrectSQLJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDataCorrectSQLJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns a failed SQL task for data change.
//
// @param request - RestartDataCorrectSQLJobRequest
//
// @return RestartDataCorrectSQLJobResponse
func (client *Client) RestartDataCorrectSQLJob(request *RestartDataCorrectSQLJobRequest) (_result *RestartDataCorrectSQLJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDataCorrectSQLJobResponse{}
	_body, _err := client.RestartDataCorrectSQLJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a SQL result set export task.
//
// @param request - RestartDataExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDataExportJobResponse
func (client *Client) RestartDataExportJobWithOptions(request *RestartDataExportJobRequest, runtime *dara.RuntimeOptions) (_result *RestartDataExportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartDataExportJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartDataExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a SQL result set export task.
//
// @param request - RestartDataExportJobRequest
//
// @return RestartDataExportJobResponse
func (client *Client) RestartDataExportJob(request *RestartDataExportJobRequest) (_result *RestartDataExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartDataExportJobResponse{}
	_body, _err := client.RestartDataExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a suspended task flow.
//
// Description:
//
// You can call this operation only for task flows that are suspended.
//
// @param request - ResumeTaskFlowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeTaskFlowInstanceResponse
func (client *Client) ResumeTaskFlowInstanceWithOptions(request *ResumeTaskFlowInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResumeTaskFlowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagInstanceId) {
		query["DagInstanceId"] = request.DagInstanceId
	}

	if !dara.IsNil(request.DagVersion) {
		query["DagVersion"] = request.DagVersion
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeTaskFlowInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeTaskFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a suspended task flow.
//
// Description:
//
// You can call this operation only for task flows that are suspended.
//
// @param request - ResumeTaskFlowInstanceRequest
//
// @return ResumeTaskFlowInstanceResponse
func (client *Client) ResumeTaskFlowInstance(request *ResumeTaskFlowInstanceRequest) (_result *ResumeTaskFlowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeTaskFlowInstanceResponse{}
	_body, _err := client.ResumeTaskFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Reruns the precheck for a data change ticket.
//
// @param request - RetryDataCorrectPreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryDataCorrectPreCheckResponse
func (client *Client) RetryDataCorrectPreCheckWithOptions(request *RetryDataCorrectPreCheckRequest, runtime *dara.RuntimeOptions) (_result *RetryDataCorrectPreCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetryDataCorrectPreCheck"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryDataCorrectPreCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reruns the precheck for a data change ticket.
//
// @param request - RetryDataCorrectPreCheckRequest
//
// @return RetryDataCorrectPreCheckResponse
func (client *Client) RetryDataCorrectPreCheck(request *RetryDataCorrectPreCheckRequest) (_result *RetryDataCorrectPreCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RetryDataCorrectPreCheckResponse{}
	_body, _err := client.RetryDataCorrectPreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on resources from Data Management (DMS) users by using a permission template.
//
// Description:
//
// You must be a database administrator (DBA) or a DMS administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - RevokeTemplateAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTemplateAuthorityResponse
func (client *Client) RevokeTemplateAuthorityWithOptions(request *RevokeTemplateAuthorityRequest, runtime *dara.RuntimeOptions) (_result *RevokeTemplateAuthorityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserIds) {
		query["UserIds"] = request.UserIds
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeTemplateAuthority"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeTemplateAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on resources from Data Management (DMS) users by using a permission template.
//
// Description:
//
// You must be a database administrator (DBA) or a DMS administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - RevokeTemplateAuthorityRequest
//
// @return RevokeTemplateAuthorityResponse
func (client *Client) RevokeTemplateAuthority(request *RevokeTemplateAuthorityRequest) (_result *RevokeTemplateAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeTemplateAuthorityResponse{}
	_body, _err := client.RevokeTemplateAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the permissions on instances, databases, and tables from a user.
//
// @param request - RevokeUserPermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeUserPermissionResponse
func (client *Client) RevokeUserPermissionWithOptions(request *RevokeUserPermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeUserPermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.DsType) {
		query["DsType"] = request.DsType
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.PermTypes) {
		query["PermTypes"] = request.PermTypes
	}

	if !dara.IsNil(request.TableId) {
		query["TableId"] = request.TableId
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UserAccessId) {
		query["UserAccessId"] = request.UserAccessId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeUserPermission"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeUserPermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the permissions on instances, databases, and tables from a user.
//
// @param request - RevokeUserPermissionRequest
//
// @return RevokeUserPermissionResponse
func (client *Client) RevokeUserPermission(request *RevokeUserPermissionRequest) (_result *RevokeUserPermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeUserPermissionResponse{}
	_body, _err := client.RevokeUserPermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for the parsing result of a data tracking task.
//
// @param tmpReq - SearchDataTrackResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDataTrackResultResponse
func (client *Client) SearchDataTrackResultWithOptions(tmpReq *SearchDataTrackResultRequest, runtime *dara.RuntimeOptions) (_result *SearchDataTrackResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SearchDataTrackResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ColumnFilter) {
		request.ColumnFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColumnFilter, dara.String("ColumnFilter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterTableList) {
		request.FilterTableListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTableList, dara.String("FilterTableList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.FilterTypeList) {
		request.FilterTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterTypeList, dara.String("FilterTypeList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ColumnFilterShrink) {
		query["ColumnFilter"] = request.ColumnFilterShrink
	}

	if !dara.IsNil(request.FilterEndTime) {
		query["FilterEndTime"] = request.FilterEndTime
	}

	if !dara.IsNil(request.FilterStartTime) {
		query["FilterStartTime"] = request.FilterStartTime
	}

	if !dara.IsNil(request.FilterTableListShrink) {
		query["FilterTableList"] = request.FilterTableListShrink
	}

	if !dara.IsNil(request.FilterTypeListShrink) {
		query["FilterTypeList"] = request.FilterTypeListShrink
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchDataTrackResult"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchDataTrackResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for the parsing result of a data tracking task.
//
// @param request - SearchDataTrackResultRequest
//
// @return SearchDataTrackResultResponse
func (client *Client) SearchDataTrackResult(request *SearchDataTrackResultRequest) (_result *SearchDataTrackResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchDataTrackResultResponse{}
	_body, _err := client.SearchDataTrackResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of the databases.
//
// @param request - SearchDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDatabaseResponse
func (client *Client) SearchDatabaseWithOptions(request *SearchDatabaseRequest, runtime *dara.RuntimeOptions) (_result *SearchDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SearchRange) {
		query["SearchRange"] = request.SearchRange
	}

	if !dara.IsNil(request.SearchTarget) {
		query["SearchTarget"] = request.SearchTarget
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of the databases.
//
// @param request - SearchDatabaseRequest
//
// @return SearchDatabaseResponse
func (client *Client) SearchDatabase(request *SearchDatabaseRequest) (_result *SearchDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchDatabaseResponse{}
	_body, _err := client.SearchDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries detailed information about tables.
//
// Description:
//
// You can call this operation only for database instances that are managed in Security Collaboration mode.
//
// @param request - SearchTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchTableResponse
func (client *Client) SearchTableWithOptions(request *SearchTableRequest, runtime *dara.RuntimeOptions) (_result *SearchTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbType) {
		query["DbType"] = request.DbType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ReturnGuid) {
		query["ReturnGuid"] = request.ReturnGuid
	}

	if !dara.IsNil(request.SearchKey) {
		query["SearchKey"] = request.SearchKey
	}

	if !dara.IsNil(request.SearchRange) {
		query["SearchRange"] = request.SearchRange
	}

	if !dara.IsNil(request.SearchTarget) {
		query["SearchTarget"] = request.SearchTarget
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchTable"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries detailed information about tables.
//
// Description:
//
// You can call this operation only for database instances that are managed in Security Collaboration mode.
//
// @param request - SearchTableRequest
//
// @return SearchTableResponse
func (client *Client) SearchTable(request *SearchTableRequest) (_result *SearchTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchTableResponse{}
	_body, _err := client.SearchTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures the owner of an instance, a database, or a table.
//
// @param request - SetOwnersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetOwnersResponse
func (client *Client) SetOwnersWithOptions(request *SetOwnersRequest, runtime *dara.RuntimeOptions) (_result *SetOwnersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerIds) {
		query["OwnerIds"] = request.OwnerIds
	}

	if !dara.IsNil(request.OwnerType) {
		query["OwnerType"] = request.OwnerType
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetOwners"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetOwnersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures the owner of an instance, a database, or a table.
//
// @param request - SetOwnersRequest
//
// @return SetOwnersResponse
func (client *Client) SetOwners(request *SetOwnersRequest) (_result *SetOwnersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetOwnersResponse{}
	_body, _err := client.SetOwnersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures whether the operation buttons, such as Agree, Revoke, and Sign, are displayed in the ticket approval section of the DMS console.
//
// @param request - SetWorkflowExtraInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetWorkflowExtraInfoResponse
func (client *Client) SetWorkflowExtraInfoWithOptions(request *SetWorkflowExtraInfoRequest, runtime *dara.RuntimeOptions) (_result *SetWorkflowExtraInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RenderAddApprovalNode) {
		query["RenderAddApprovalNode"] = request.RenderAddApprovalNode
	}

	if !dara.IsNil(request.RenderAgree) {
		query["RenderAgree"] = request.RenderAgree
	}

	if !dara.IsNil(request.RenderCancel) {
		query["RenderCancel"] = request.RenderCancel
	}

	if !dara.IsNil(request.RenderReject) {
		query["RenderReject"] = request.RenderReject
	}

	if !dara.IsNil(request.RenderTransfer) {
		query["RenderTransfer"] = request.RenderTransfer
	}

	if !dara.IsNil(request.ThirdpartyWorkflowComment) {
		query["ThirdpartyWorkflowComment"] = request.ThirdpartyWorkflowComment
	}

	if !dara.IsNil(request.ThirdpartyWorkflowUrl) {
		query["ThirdpartyWorkflowUrl"] = request.ThirdpartyWorkflowUrl
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkflowInstanceId) {
		query["WorkflowInstanceId"] = request.WorkflowInstanceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetWorkflowExtraInfo"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetWorkflowExtraInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures whether the operation buttons, such as Agree, Revoke, and Sign, are displayed in the ticket approval section of the DMS console.
//
// @param request - SetWorkflowExtraInfoRequest
//
// @return SetWorkflowExtraInfoResponse
func (client *Client) SetWorkflowExtraInfo(request *SetWorkflowExtraInfoRequest) (_result *SetWorkflowExtraInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetWorkflowExtraInfoResponse{}
	_body, _err := client.SetWorkflowExtraInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加实例
//
// @param request - SimplyAddInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SimplyAddInstanceResponse
func (client *Client) SimplyAddInstanceWithOptions(request *SimplyAddInstanceRequest, runtime *dara.RuntimeOptions) (_result *SimplyAddInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.DatabaseUser) {
		query["DatabaseUser"] = request.DatabaseUser
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceRegion) {
		query["InstanceRegion"] = request.InstanceRegion
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SimplyAddInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SimplyAddInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加实例
//
// @param request - SimplyAddInstanceRequest
//
// @return SimplyAddInstanceResponse
func (client *Client) SimplyAddInstance(request *SimplyAddInstanceRequest) (_result *SimplyAddInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SimplyAddInstanceResponse{}
	_body, _err := client.SimplyAddInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Skips the verification on the number of rows in the precheck for data change.
//
// @param request - SkipDataCorrectRowCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipDataCorrectRowCheckResponse
func (client *Client) SkipDataCorrectRowCheckWithOptions(request *SkipDataCorrectRowCheckRequest, runtime *dara.RuntimeOptions) (_result *SkipDataCorrectRowCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Reason) {
		query["Reason"] = request.Reason
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipDataCorrectRowCheck"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipDataCorrectRowCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Skips the verification on the number of rows in the precheck for data change.
//
// @param request - SkipDataCorrectRowCheckRequest
//
// @return SkipDataCorrectRowCheckResponse
func (client *Client) SkipDataCorrectRowCheck(request *SkipDataCorrectRowCheckRequest) (_result *SkipDataCorrectRowCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SkipDataCorrectRowCheckResponse{}
	_body, _err := client.SkipDataCorrectRowCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a task flow instance.
//
// @param request - StopTaskFlowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTaskFlowInstanceResponse
func (client *Client) StopTaskFlowInstanceWithOptions(request *StopTaskFlowInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopTaskFlowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagInstanceId) {
		query["DagInstanceId"] = request.DagInstanceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTaskFlowInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTaskFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a task flow instance.
//
// @param request - StopTaskFlowInstanceRequest
//
// @return StopTaskFlowInstanceResponse
func (client *Client) StopTaskFlowInstance(request *StopTaskFlowInstanceRequest) (_result *StopTaskFlowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopTaskFlowInstanceResponse{}
	_body, _err := client.StopTaskFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a ticket for approval.
//
// @param request - SubmitOrderApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitOrderApprovalResponse
func (client *Client) SubmitOrderApprovalWithOptions(request *SubmitOrderApprovalRequest, runtime *dara.RuntimeOptions) (_result *SubmitOrderApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.RealLoginUserUid) {
		query["RealLoginUserUid"] = request.RealLoginUserUid
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitOrderApproval"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitOrderApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a ticket for approval.
//
// @param request - SubmitOrderApprovalRequest
//
// @return SubmitOrderApprovalResponse
func (client *Client) SubmitOrderApproval(request *SubmitOrderApprovalRequest) (_result *SubmitOrderApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitOrderApprovalResponse{}
	_body, _err := client.SubmitOrderApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a schema synchronization ticket for approval.
//
// @param request - SubmitStructSyncOrderApprovalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitStructSyncOrderApprovalResponse
func (client *Client) SubmitStructSyncOrderApprovalWithOptions(request *SubmitStructSyncOrderApprovalRequest, runtime *dara.RuntimeOptions) (_result *SubmitStructSyncOrderApprovalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitStructSyncOrderApproval"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitStructSyncOrderApprovalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a schema synchronization ticket for approval.
//
// @param request - SubmitStructSyncOrderApprovalRequest
//
// @return SubmitStructSyncOrderApprovalResponse
func (client *Client) SubmitStructSyncOrderApproval(request *SubmitStructSyncOrderApprovalRequest) (_result *SubmitStructSyncOrderApprovalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitStructSyncOrderApprovalResponse{}
	_body, _err := client.SubmitStructSyncOrderApprovalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends an ongoing or scheduled SQL result set export task.
//
// @param request - SuspendDataExportJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendDataExportJobResponse
func (client *Client) SuspendDataExportJobWithOptions(request *SuspendDataExportJobRequest, runtime *dara.RuntimeOptions) (_result *SuspendDataExportJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		query["JobId"] = request.JobId
	}

	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendDataExportJob"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendDataExportJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends an ongoing or scheduled SQL result set export task.
//
// @param request - SuspendDataExportJobRequest
//
// @return SuspendDataExportJobResponse
func (client *Client) SuspendDataExportJob(request *SuspendDataExportJobRequest) (_result *SuspendDataExportJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendDataExportJobResponse{}
	_body, _err := client.SuspendDataExportJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends a task flow instance.
//
// @param request - SuspendTaskFlowInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendTaskFlowInstanceResponse
func (client *Client) SuspendTaskFlowInstanceWithOptions(request *SuspendTaskFlowInstanceRequest, runtime *dara.RuntimeOptions) (_result *SuspendTaskFlowInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagInstanceId) {
		query["DagInstanceId"] = request.DagInstanceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendTaskFlowInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendTaskFlowInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a task flow instance.
//
// @param request - SuspendTaskFlowInstanceRequest
//
// @return SuspendTaskFlowInstanceResponse
func (client *Client) SuspendTaskFlowInstance(request *SuspendTaskFlowInstanceRequest) (_result *SuspendTaskFlowInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendTaskFlowInstanceResponse{}
	_body, _err := client.SuspendTaskFlowInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the metadata of a database.
//
// @param request - SyncDatabaseMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncDatabaseMetaResponse
func (client *Client) SyncDatabaseMetaWithOptions(request *SyncDatabaseMetaRequest, runtime *dara.RuntimeOptions) (_result *SyncDatabaseMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DbId) {
		query["DbId"] = request.DbId
	}

	if !dara.IsNil(request.Logic) {
		query["Logic"] = request.Logic
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncDatabaseMeta"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncDatabaseMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the metadata of a database.
//
// @param request - SyncDatabaseMetaRequest
//
// @return SyncDatabaseMetaResponse
func (client *Client) SyncDatabaseMeta(request *SyncDatabaseMetaRequest) (_result *SyncDatabaseMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncDatabaseMetaResponse{}
	_body, _err := client.SyncDatabaseMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Synchronizes the metadata of all databases in a database instance.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
//
// @param request - SyncInstanceMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncInstanceMetaResponse
func (client *Client) SyncInstanceMetaWithOptions(request *SyncInstanceMetaRequest, runtime *dara.RuntimeOptions) (_result *SyncInstanceMetaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IgnoreTable) {
		query["IgnoreTable"] = request.IgnoreTable
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SyncInstanceMeta"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SyncInstanceMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Synchronizes the metadata of all databases in a database instance.
//
// Description:
//
// You can call this operation only for database instances whose control mode is Security Collaboration.
//
// @param request - SyncInstanceMetaRequest
//
// @return SyncInstanceMetaResponse
func (client *Client) SyncInstanceMeta(request *SyncInstanceMetaRequest) (_result *SyncInstanceMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SyncInstanceMetaResponse{}
	_body, _err := client.SyncInstanceMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the task flow of task orchestration.
//
// Description:
//
// ### [](#)Usage notes
//
// An edge can be created only when the following conditions are met:
//
// 1.  The start and end nodes of the edge exist in the directed acyclic graph (DAG) of the task flow specified by DagId.
//
// 2.  After the edge is created, the DAG does not contain loops.
//
// @param request - TryRunTaskFlowRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TryRunTaskFlowResponse
func (client *Client) TryRunTaskFlowWithOptions(request *TryRunTaskFlowRequest, runtime *dara.RuntimeOptions) (_result *TryRunTaskFlowResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TryRunTaskFlow"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TryRunTaskFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the task flow of task orchestration.
//
// Description:
//
// ### [](#)Usage notes
//
// An edge can be created only when the following conditions are met:
//
// 1.  The start and end nodes of the edge exist in the directed acyclic graph (DAG) of the task flow specified by DagId.
//
// 2.  After the edge is created, the DAG does not contain loops.
//
// @param request - TryRunTaskFlowRequest
//
// @return TryRunTaskFlowResponse
func (client *Client) TryRunTaskFlow(request *TryRunTaskFlowRequest) (_result *TryRunTaskFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TryRunTaskFlowResponse{}
	_body, _err := client.TryRunTaskFlowWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name, description, or content of a policy.
//
// @param request - UpdateAbacPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAbacPolicyResponse
func (client *Client) UpdateAbacPolicyWithOptions(request *UpdateAbacPolicyRequest, runtime *dara.RuntimeOptions) (_result *UpdateAbacPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AbacPolicyContent) {
		query["AbacPolicyContent"] = request.AbacPolicyContent
	}

	if !dara.IsNil(request.AbacPolicyDesc) {
		query["AbacPolicyDesc"] = request.AbacPolicyDesc
	}

	if !dara.IsNil(request.AbacPolicyId) {
		query["AbacPolicyId"] = request.AbacPolicyId
	}

	if !dara.IsNil(request.AbacPolicyName) {
		query["AbacPolicyName"] = request.AbacPolicyName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAbacPolicy"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAbacPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name, description, or content of a policy.
//
// @param request - UpdateAbacPolicyRequest
//
// @return UpdateAbacPolicyResponse
func (client *Client) UpdateAbacPolicy(request *UpdateAbacPolicyRequest) (_result *UpdateAbacPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAbacPolicyResponse{}
	_body, _err := client.UpdateAbacPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a permission template.
//
// Description:
//
// You are a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - UpdateAuthorityTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAuthorityTemplateResponse
func (client *Client) UpdateAuthorityTemplateWithOptions(request *UpdateAuthorityTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateAuthorityTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAuthorityTemplate"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAuthorityTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a permission template.
//
// Description:
//
// You are a database administrator (DBA) or a Data Management (DMS) administrator. For more information about how to view system roles, see [View system roles](https://help.aliyun.com/document_detail/324212.html).
//
// @param request - UpdateAuthorityTemplateRequest
//
// @return UpdateAuthorityTemplateResponse
func (client *Client) UpdateAuthorityTemplate(request *UpdateAuthorityTemplateRequest) (_result *UpdateAuthorityTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAuthorityTemplateResponse{}
	_body, _err := client.UpdateAuthorityTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a database in a data lakehouse.
//
// @param tmpReq - UpdateDataLakeDatabaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeDatabaseResponse
func (client *Client) UpdateDataLakeDatabaseWithOptions(tmpReq *UpdateDataLakeDatabaseRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeDatabaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeDatabaseShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.ParametersShrink) {
		query["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeDatabase"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a database in a data lakehouse.
//
// @param request - UpdateDataLakeDatabaseRequest
//
// @return UpdateDataLakeDatabaseResponse
func (client *Client) UpdateDataLakeDatabase(request *UpdateDataLakeDatabaseRequest) (_result *UpdateDataLakeDatabaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakeDatabaseResponse{}
	_body, _err := client.UpdateDataLakeDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a user-defined function in a data lakehouse.
//
// @param tmpReq - UpdateDataLakeFunctionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeFunctionResponse
func (client *Client) UpdateDataLakeFunctionWithOptions(tmpReq *UpdateDataLakeFunctionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeFunctionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeFunctionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.FunctionInput) {
		request.FunctionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FunctionInput, dara.String("FunctionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.FunctionName) {
		query["FunctionName"] = request.FunctionName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.FunctionInputShrink) {
		body["FunctionInput"] = request.FunctionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeFunction"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a user-defined function in a data lakehouse.
//
// @param request - UpdateDataLakeFunctionRequest
//
// @return UpdateDataLakeFunctionResponse
func (client *Client) UpdateDataLakeFunction(request *UpdateDataLakeFunctionRequest) (_result *UpdateDataLakeFunctionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakeFunctionResponse{}
	_body, _err := client.UpdateDataLakeFunctionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the partition information of a table in a data lakehouse.
//
// @param tmpReq - UpdateDataLakePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakePartitionResponse
func (client *Client) UpdateDataLakePartitionWithOptions(tmpReq *UpdateDataLakePartitionRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakePartitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PartitionInput) {
		request.PartitionInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionInput, dara.String("PartitionInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.PartitionInputShrink) {
		body["PartitionInput"] = request.PartitionInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakePartition"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the partition information of a table in a data lakehouse.
//
// @param request - UpdateDataLakePartitionRequest
//
// @return UpdateDataLakePartitionResponse
func (client *Client) UpdateDataLakePartition(request *UpdateDataLakePartitionRequest) (_result *UpdateDataLakePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakePartitionResponse{}
	_body, _err := client.UpdateDataLakePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a table in a data lakehouse.
//
// @param tmpReq - UpdateDataLakeTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataLakeTableResponse
func (client *Client) UpdateDataLakeTableWithOptions(tmpReq *UpdateDataLakeTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataLakeTableResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateDataLakeTableShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TableInput) {
		request.TableInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableInput, dara.String("TableInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.DataRegion) {
		query["DataRegion"] = request.DataRegion
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TableInputShrink) {
		body["TableInput"] = request.TableInputShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataLakeTable"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataLakeTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a table in a data lakehouse.
//
// @param request - UpdateDataLakeTableRequest
//
// @return UpdateDataLakeTableResponse
func (client *Client) UpdateDataLakeTable(request *UpdateDataLakeTableRequest) (_result *UpdateDataLakeTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataLakeTableResponse{}
	_body, _err := client.UpdateDataLakeTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a database instance and checks the connectivity of the database instance.
//
// Description:
//
// Before you call the UpdateInstance operation, call the [GetInstance](https://help.aliyun.com/document_detail/141567.html) or [ListInstances](https://help.aliyun.com/document_detail/141936.html) operation to obtain the complete information about the instance.
//
// @param request - UpdateInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstanceWithOptions(request *UpdateInstanceRequest, runtime *dara.RuntimeOptions) (_result *UpdateInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataLinkName) {
		query["DataLinkName"] = request.DataLinkName
	}

	if !dara.IsNil(request.DatabasePassword) {
		query["DatabasePassword"] = request.DatabasePassword
	}

	if !dara.IsNil(request.DatabaseUser) {
		query["DatabaseUser"] = request.DatabaseUser
	}

	if !dara.IsNil(request.DbaId) {
		query["DbaId"] = request.DbaId
	}

	if !dara.IsNil(request.DdlOnline) {
		query["DdlOnline"] = request.DdlOnline
	}

	if !dara.IsNil(request.EcsInstanceId) {
		query["EcsInstanceId"] = request.EcsInstanceId
	}

	if !dara.IsNil(request.EcsRegion) {
		query["EcsRegion"] = request.EcsRegion
	}

	if !dara.IsNil(request.EnableSellSitd) {
		query["EnableSellSitd"] = request.EnableSellSitd
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ExportTimeout) {
		query["ExportTimeout"] = request.ExportTimeout
	}

	if !dara.IsNil(request.Host) {
		query["Host"] = request.Host
	}

	if !dara.IsNil(request.InstanceAlias) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !dara.IsNil(request.InstanceId) {
		query["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.InstanceSource) {
		query["InstanceSource"] = request.InstanceSource
	}

	if !dara.IsNil(request.InstanceType) {
		query["InstanceType"] = request.InstanceType
	}

	if !dara.IsNil(request.Port) {
		query["Port"] = request.Port
	}

	if !dara.IsNil(request.QueryTimeout) {
		query["QueryTimeout"] = request.QueryTimeout
	}

	if !dara.IsNil(request.SafeRuleId) {
		query["SafeRuleId"] = request.SafeRuleId
	}

	if !dara.IsNil(request.Sid) {
		query["Sid"] = request.Sid
	}

	if !dara.IsNil(request.SkipTest) {
		query["SkipTest"] = request.SkipTest
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.UseDsql) {
		query["UseDsql"] = request.UseDsql
	}

	if !dara.IsNil(request.VpcId) {
		query["VpcId"] = request.VpcId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateInstance"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a database instance and checks the connectivity of the database instance.
//
// Description:
//
// Before you call the UpdateInstance operation, call the [GetInstance](https://help.aliyun.com/document_detail/141567.html) or [ListInstances](https://help.aliyun.com/document_detail/141936.html) operation to obtain the complete information about the instance.
//
// @param request - UpdateInstanceRequest
//
// @return UpdateInstanceResponse
func (client *Client) UpdateInstance(request *UpdateInstanceRequest) (_result *UpdateInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateInstanceResponse{}
	_body, _err := client.UpdateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates asset category information.
//
// @param request - UpdateMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaCategoryResponse
func (client *Client) UpdateMetaCategoryWithOptions(request *UpdateMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaCategoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaCategory"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates asset category information.
//
// @param request - UpdateMetaCategoryRequest
//
// @return UpdateMetaCategoryResponse
func (client *Client) UpdateMetaCategory(request *UpdateMetaCategoryRequest) (_result *UpdateMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMetaCategoryResponse{}
	_body, _err := client.UpdateMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the service level agreement (SLA) timeout reminder for a task flow.
//
// Description:
//
// SLA rules take effect after task flows are deployed and published.
//
// @param tmpReq - UpdateSLARulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSLARulesResponse
func (client *Client) UpdateSLARulesWithOptions(tmpReq *UpdateSLARulesRequest, runtime *dara.RuntimeOptions) (_result *UpdateSLARulesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateSLARulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SlaRuleList) {
		request.SlaRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SlaRuleList, dara.String("SlaRuleList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.SlaRuleListShrink) {
		query["SlaRuleList"] = request.SlaRuleListShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateSLARules"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateSLARulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the service level agreement (SLA) timeout reminder for a task flow.
//
// Description:
//
// SLA rules take effect after task flows are deployed and published.
//
// @param request - UpdateSLARulesRequest
//
// @return UpdateSLARulesResponse
func (client *Client) UpdateSLARules(request *UpdateSLARulesRequest) (_result *UpdateSLARulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateSLARulesResponse{}
	_body, _err := client.UpdateSLARulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name and description of the business scenario for a specified task flow.
//
// @param request - UpdateScenarioRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateScenarioResponse
func (client *Client) UpdateScenarioWithOptions(request *UpdateScenarioRequest, runtime *dara.RuntimeOptions) (_result *UpdateScenarioResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ScenarioId) {
		query["ScenarioId"] = request.ScenarioId
	}

	if !dara.IsNil(request.ScenarioName) {
		query["ScenarioName"] = request.ScenarioName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateScenario"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateScenarioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name and description of the business scenario for a specified task flow.
//
// @param request - UpdateScenarioRequest
//
// @return UpdateScenarioResponse
func (client *Client) UpdateScenario(request *UpdateScenarioRequest) (_result *UpdateScenarioResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateScenarioResponse{}
	_body, _err := client.UpdateScenarioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a specific security rule set.
//
// @param request - UpdateStandardGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStandardGroupResponse
func (client *Client) UpdateStandardGroupWithOptions(request *UpdateStandardGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateStandardGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupId) {
		query["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStandardGroup"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStandardGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the name and description of a specific security rule set.
//
// @param request - UpdateStandardGroupRequest
//
// @return UpdateStandardGroupResponse
func (client *Client) UpdateStandardGroup(request *UpdateStandardGroupRequest) (_result *UpdateStandardGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStandardGroupResponse{}
	_body, _err := client.UpdateStandardGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the advanced configuration of a task node.
//
// Description:
//
// You can call this operation to configure a failed task or rerun a task.
//
// @param request - UpdateTaskConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskConfigResponse
func (client *Client) UpdateTaskConfigWithOptions(request *UpdateTaskConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeConfig) {
		query["NodeConfig"] = request.NodeConfig
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskConfig"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the advanced configuration of a task node.
//
// Description:
//
// You can call this operation to configure a failed task or rerun a task.
//
// @param request - UpdateTaskConfigRequest
//
// @return UpdateTaskConfigResponse
func (client *Client) UpdateTaskConfig(request *UpdateTaskConfigRequest) (_result *UpdateTaskConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskConfigResponse{}
	_body, _err := client.UpdateTaskConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates tasks in a task flow.
//
// Description:
//
// You can call this operation to modify node configurations.
//
// @param request - UpdateTaskContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskContentResponse
func (client *Client) UpdateTaskContentWithOptions(request *UpdateTaskContentRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeContent) {
		query["NodeContent"] = request.NodeContent
	}

	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskContent"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates tasks in a task flow.
//
// Description:
//
// You can call this operation to modify node configurations.
//
// @param request - UpdateTaskContentRequest
//
// @return UpdateTaskContentResponse
func (client *Client) UpdateTaskContent(request *UpdateTaskContentRequest) (_result *UpdateTaskContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskContentResponse{}
	_body, _err := client.UpdateTaskContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the node content in large SQL files.
//
// Description:
//
// You can call this operation to modify node configurations.
//
// @param request - UpdateTaskContentV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskContentV2Response
func (client *Client) UpdateTaskContentV2WithOptions(request *UpdateTaskContentV2Request, runtime *dara.RuntimeOptions) (_result *UpdateTaskContentV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeContent) {
		body["NodeContent"] = request.NodeContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskContentV2"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskContentV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the node content in large SQL files.
//
// Description:
//
// You can call this operation to modify node configurations.
//
// @param request - UpdateTaskContentV2Request
//
// @return UpdateTaskContentV2Response
func (client *Client) UpdateTaskContentV2(request *UpdateTaskContentV2Request) (_result *UpdateTaskContentV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskContentV2Response{}
	_body, _err := client.UpdateTaskContentV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the constants for a specified task flow.
//
// @param tmpReq - UpdateTaskFlowConstantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowConstantsResponse
func (client *Client) UpdateTaskFlowConstantsWithOptions(tmpReq *UpdateTaskFlowConstantsRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowConstantsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskFlowConstantsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DagConstants) {
		request.DagConstantsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DagConstants, dara.String("DagConstants"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DagConstantsShrink) {
		query["DagConstants"] = request.DagConstantsShrink
	}

	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowConstants"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowConstantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the constants for a specified task flow.
//
// @param request - UpdateTaskFlowConstantsRequest
//
// @return UpdateTaskFlowConstantsResponse
func (client *Client) UpdateTaskFlowConstants(request *UpdateTaskFlowConstantsRequest) (_result *UpdateTaskFlowConstantsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowConstantsResponse{}
	_body, _err := client.UpdateTaskFlowConstantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the IDs of the users who are involved in the task flow.
//
// @param tmpReq - UpdateTaskFlowCooperatorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowCooperatorsResponse
func (client *Client) UpdateTaskFlowCooperatorsWithOptions(tmpReq *UpdateTaskFlowCooperatorsRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowCooperatorsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskFlowCooperatorsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CooperatorIds) {
		request.CooperatorIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CooperatorIds, dara.String("CooperatorIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CooperatorIdsShrink) {
		query["CooperatorIds"] = request.CooperatorIdsShrink
	}

	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowCooperators"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowCooperatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the IDs of the users who are involved in the task flow.
//
// @param request - UpdateTaskFlowCooperatorsRequest
//
// @return UpdateTaskFlowCooperatorsResponse
func (client *Client) UpdateTaskFlowCooperators(request *UpdateTaskFlowCooperatorsRequest) (_result *UpdateTaskFlowCooperatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowCooperatorsResponse{}
	_body, _err := client.UpdateTaskFlowCooperatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the start node and end node of multiple edges at a time for a task flow.
//
// Description:
//
// ###
//
// The edges can be updated only when the following conditions are met:
//
// 1.  The specified edge exists in the directed acyclic graph (DAG) of the task flow specified by DagId.
//
// 2.  The specified edge nodes exist in the DAG of the task flow specified by DagId.
//
// 3.  After the update, rings do not exist in the DAG.
//
// @param tmpReq - UpdateTaskFlowEdgesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowEdgesResponse
func (client *Client) UpdateTaskFlowEdgesWithOptions(tmpReq *UpdateTaskFlowEdgesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowEdgesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskFlowEdgesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Edges) {
		request.EdgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Edges, dara.String("Edges"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.EdgesShrink) {
		query["Edges"] = request.EdgesShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowEdges"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowEdgesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the start node and end node of multiple edges at a time for a task flow.
//
// Description:
//
// ###
//
// The edges can be updated only when the following conditions are met:
//
// 1.  The specified edge exists in the directed acyclic graph (DAG) of the task flow specified by DagId.
//
// 2.  The specified edge nodes exist in the DAG of the task flow specified by DagId.
//
// 3.  After the update, rings do not exist in the DAG.
//
// @param request - UpdateTaskFlowEdgesRequest
//
// @return UpdateTaskFlowEdgesResponse
func (client *Client) UpdateTaskFlowEdges(request *UpdateTaskFlowEdgesRequest) (_result *UpdateTaskFlowEdgesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowEdgesResponse{}
	_body, _err := client.UpdateTaskFlowEdgesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name and description of a task flow.
//
// @param request - UpdateTaskFlowNameAndDescRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowNameAndDescResponse
func (client *Client) UpdateTaskFlowNameAndDescWithOptions(request *UpdateTaskFlowNameAndDescRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowNameAndDescResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagName) {
		query["DagName"] = request.DagName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowNameAndDesc"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowNameAndDescResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name and description of a task flow.
//
// @param request - UpdateTaskFlowNameAndDescRequest
//
// @return UpdateTaskFlowNameAndDescResponse
func (client *Client) UpdateTaskFlowNameAndDesc(request *UpdateTaskFlowNameAndDescRequest) (_result *UpdateTaskFlowNameAndDescResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowNameAndDescResponse{}
	_body, _err := client.UpdateTaskFlowNameAndDescWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the notification settings for task flows.
//
// @param request - UpdateTaskFlowNotificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowNotificationResponse
func (client *Client) UpdateTaskFlowNotificationWithOptions(request *UpdateTaskFlowNotificationRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowNotificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.DagNotificationFail) {
		query["DagNotificationFail"] = request.DagNotificationFail
	}

	if !dara.IsNil(request.DagNotificationSla) {
		query["DagNotificationSla"] = request.DagNotificationSla
	}

	if !dara.IsNil(request.DagNotificationSuccess) {
		query["DagNotificationSuccess"] = request.DagNotificationSuccess
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowNotification"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowNotificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the notification settings for task flows.
//
// @param request - UpdateTaskFlowNotificationRequest
//
// @return UpdateTaskFlowNotificationResponse
func (client *Client) UpdateTaskFlowNotification(request *UpdateTaskFlowNotificationRequest) (_result *UpdateTaskFlowNotificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowNotificationResponse{}
	_body, _err := client.UpdateTaskFlowNotificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the owner of a task flow.
//
// Description:
//
// Note: The new owner of the task flow must belong to the same tenant as the previous owner.
//
// @param request - UpdateTaskFlowOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowOwnerResponse
func (client *Client) UpdateTaskFlowOwnerWithOptions(request *UpdateTaskFlowOwnerRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowOwnerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.NewOwnerId) {
		query["NewOwnerId"] = request.NewOwnerId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowOwner"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the owner of a task flow.
//
// Description:
//
// Note: The new owner of the task flow must belong to the same tenant as the previous owner.
//
// @param request - UpdateTaskFlowOwnerRequest
//
// @return UpdateTaskFlowOwnerResponse
func (client *Client) UpdateTaskFlowOwner(request *UpdateTaskFlowOwnerRequest) (_result *UpdateTaskFlowOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowOwnerResponse{}
	_body, _err := client.UpdateTaskFlowOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Fully updates the edges of a task flow.
//
// Description:
//
// You can call this operation to perform a full update. For incremental updates, see AddTaskFlowEdges, UpdateTaskFlowEdges, and DeleteTaskFlowEdgesByMultiCondition.
//
// @param tmpReq - UpdateTaskFlowRelationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowRelationsResponse
func (client *Client) UpdateTaskFlowRelationsWithOptions(tmpReq *UpdateTaskFlowRelationsRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowRelationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskFlowRelationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Edges) {
		request.EdgesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Edges, dara.String("Edges"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.EdgesShrink) {
		query["Edges"] = request.EdgesShrink
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowRelations"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowRelationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Fully updates the edges of a task flow.
//
// Description:
//
// You can call this operation to perform a full update. For incremental updates, see AddTaskFlowEdges, UpdateTaskFlowEdges, and DeleteTaskFlowEdgesByMultiCondition.
//
// @param request - UpdateTaskFlowRelationsRequest
//
// @return UpdateTaskFlowRelationsResponse
func (client *Client) UpdateTaskFlowRelations(request *UpdateTaskFlowRelationsRequest) (_result *UpdateTaskFlowRelationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowRelationsResponse{}
	_body, _err := client.UpdateTaskFlowRelationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the scheduling properties for a task flow.
//
// Description:
//
// You can call this operation to update the scheduling properties for a task flow in the editing state. You can configure a **timed scheduling*	- task flow or an **event scheduling*	- task flow. When you configure a **timed scheduling*	- task flow, you can choose from one-time scheduling or periodic scheduling. When you configure an **event scheduling*	- task flow, you can subscribe to task flows or task flow nodes.****\\
//
// After you update the scheduling properties, you need to publish and deploy the task flow again. The new task flow instance will run based on the updated scheduling properties.
//
// @param request - UpdateTaskFlowScheduleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowScheduleResponse
func (client *Client) UpdateTaskFlowScheduleWithOptions(request *UpdateTaskFlowScheduleRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowScheduleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CronBeginDate) {
		query["CronBeginDate"] = request.CronBeginDate
	}

	if !dara.IsNil(request.CronEndDate) {
		query["CronEndDate"] = request.CronEndDate
	}

	if !dara.IsNil(request.CronStr) {
		query["CronStr"] = request.CronStr
	}

	if !dara.IsNil(request.CronType) {
		query["CronType"] = request.CronType
	}

	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.ScheduleParam) {
		query["ScheduleParam"] = request.ScheduleParam
	}

	if !dara.IsNil(request.ScheduleSwitch) {
		query["ScheduleSwitch"] = request.ScheduleSwitch
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.TimeZoneId) {
		query["TimeZoneId"] = request.TimeZoneId
	}

	if !dara.IsNil(request.TriggerType) {
		query["TriggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowSchedule"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowScheduleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the scheduling properties for a task flow.
//
// Description:
//
// You can call this operation to update the scheduling properties for a task flow in the editing state. You can configure a **timed scheduling*	- task flow or an **event scheduling*	- task flow. When you configure a **timed scheduling*	- task flow, you can choose from one-time scheduling or periodic scheduling. When you configure an **event scheduling*	- task flow, you can subscribe to task flows or task flow nodes.****\\
//
// After you update the scheduling properties, you need to publish and deploy the task flow again. The new task flow instance will run based on the updated scheduling properties.
//
// @param request - UpdateTaskFlowScheduleRequest
//
// @return UpdateTaskFlowScheduleResponse
func (client *Client) UpdateTaskFlowSchedule(request *UpdateTaskFlowScheduleRequest) (_result *UpdateTaskFlowScheduleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowScheduleResponse{}
	_body, _err := client.UpdateTaskFlowScheduleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the time variables for a task flow.
//
// @param request - UpdateTaskFlowTimeVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskFlowTimeVariablesResponse
func (client *Client) UpdateTaskFlowTimeVariablesWithOptions(request *UpdateTaskFlowTimeVariablesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskFlowTimeVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		query["DagId"] = request.DagId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.TimeVariables) {
		query["TimeVariables"] = request.TimeVariables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskFlowTimeVariables"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskFlowTimeVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the time variables for a task flow.
//
// @param request - UpdateTaskFlowTimeVariablesRequest
//
// @return UpdateTaskFlowTimeVariablesResponse
func (client *Client) UpdateTaskFlowTimeVariables(request *UpdateTaskFlowTimeVariablesRequest) (_result *UpdateTaskFlowTimeVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskFlowTimeVariablesResponse{}
	_body, _err := client.UpdateTaskFlowTimeVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the name of a specified task.
//
// @param request - UpdateTaskNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskNameResponse
func (client *Client) UpdateTaskNameWithOptions(request *UpdateTaskNameRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeName) {
		query["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskName"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name of a specified task.
//
// @param request - UpdateTaskNameRequest
//
// @return UpdateTaskNameResponse
func (client *Client) UpdateTaskName(request *UpdateTaskNameRequest) (_result *UpdateTaskNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskNameResponse{}
	_body, _err := client.UpdateTaskNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the output variables for a specified task node.
//
// Description:
//
// Only nodes of single-instance SQL assignment, script code, and ECS remote command have output variables.
//
// @param request - UpdateTaskOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskOutputResponse
func (client *Client) UpdateTaskOutputWithOptions(request *UpdateTaskOutputRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskOutputResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeOutput) {
		query["NodeOutput"] = request.NodeOutput
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskOutput"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the output variables for a specified task node.
//
// Description:
//
// Only nodes of single-instance SQL assignment, script code, and ECS remote command have output variables.
//
// @param request - UpdateTaskOutputRequest
//
// @return UpdateTaskOutputResponse
func (client *Client) UpdateTaskOutput(request *UpdateTaskOutputRequest) (_result *UpdateTaskOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskOutputResponse{}
	_body, _err := client.UpdateTaskOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates time variables for a task.
//
// @param request - UpdateTaskTimeVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskTimeVariablesResponse
func (client *Client) UpdateTaskTimeVariablesWithOptions(request *UpdateTaskTimeVariablesRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskTimeVariablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		query["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.TimeVariables) {
		query["TimeVariables"] = request.TimeVariables
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTaskTimeVariables"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskTimeVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates time variables for a task.
//
// @param request - UpdateTaskTimeVariablesRequest
//
// @return UpdateTaskTimeVariablesResponse
func (client *Client) UpdateTaskTimeVariables(request *UpdateTaskTimeVariablesRequest) (_result *UpdateTaskTimeVariablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskTimeVariablesResponse{}
	_body, _err := client.UpdateTaskTimeVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates user information.
//
// @param request - UpdateUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *dara.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxExecuteCount) {
		query["MaxExecuteCount"] = request.MaxExecuteCount
	}

	if !dara.IsNil(request.MaxResultCount) {
		query["MaxResultCount"] = request.MaxResultCount
	}

	if !dara.IsNil(request.Mobile) {
		query["Mobile"] = request.Mobile
	}

	if !dara.IsNil(request.RoleNames) {
		query["RoleNames"] = request.RoleNames
	}

	if !dara.IsNil(request.Tid) {
		query["Tid"] = request.Tid
	}

	if !dara.IsNil(request.Uid) {
		query["Uid"] = request.Uid
	}

	if !dara.IsNil(request.UidString) {
		query["UidString"] = request.UidString
	}

	if !dara.IsNil(request.UserNick) {
		query["UserNick"] = request.UserNick
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUser"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates user information.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspaceWithOptions(request *UpdateWorkspaceRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceName) {
		query["WorkspaceName"] = request.WorkspaceName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspace"),
		Version:     dara.String("2018-11-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspace(request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
