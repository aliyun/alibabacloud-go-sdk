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
		"ap-northeast-1":        dara.String("dataworks.ap-northeast-1.aliyuncs.com"),
		"ap-south-1":            dara.String("dataworks.ap-south-1.aliyuncs.com"),
		"ap-southeast-1":        dara.String("dataworks.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        dara.String("dataworks.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        dara.String("dataworks.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        dara.String("dataworks.ap-southeast-5.aliyuncs.com"),
		"cn-beijing":            dara.String("dataworks.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            dara.String("dataworks.cn-chengdu.aliyuncs.com"),
		"cn-hangzhou":           dara.String("dataworks.cn-hangzhou.aliyuncs.com"),
		"cn-hongkong":           dara.String("dataworks.cn-hongkong.aliyuncs.com"),
		"cn-huhehaote":          dara.String("dataworks.aliyuncs.com"),
		"cn-qingdao":            dara.String("dataworks.aliyuncs.com"),
		"cn-shanghai":           dara.String("dataworks.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           dara.String("dataworks.cn-shenzhen.aliyuncs.com"),
		"cn-zhangjiakou":        dara.String("dataworks.aliyuncs.com"),
		"eu-central-1":          dara.String("dataworks.eu-central-1.aliyuncs.com"),
		"eu-west-1":             dara.String("dataworks.eu-west-1.aliyuncs.com"),
		"me-east-1":             dara.String("dataworks.me-east-1.aliyuncs.com"),
		"us-east-1":             dara.String("dataworks.us-east-1.aliyuncs.com"),
		"us-west-1":             dara.String("dataworks.us-west-1.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("dataworks.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("dataworks.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("dataworks.aliyuncs.com"),
		"cn-north-2-gov-1":      dara.String("dataworks.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dataworks-public"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// Unpublishes a DataService Studio API.
//
// @param request - AbolishDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AbolishDataServiceApiResponse
func (client *Client) AbolishDataServiceApiWithOptions(request *AbolishDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *AbolishDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AbolishDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AbolishDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublishes a DataService Studio API.
//
// @param request - AbolishDataServiceApiRequest
//
// @return AbolishDataServiceApiResponse
func (client *Client) AbolishDataServiceApi(request *AbolishDataServiceApiRequest) (_result *AbolishDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AbolishDataServiceApiResponse{}
	_body, _err := client.AbolishDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds an entity to a collection.
//
// @param request - AddMetaCollectionEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMetaCollectionEntityResponse
func (client *Client) AddMetaCollectionEntityWithOptions(request *AddMetaCollectionEntityRequest, runtime *dara.RuntimeOptions) (_result *AddMetaCollectionEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CollectionQualifiedName) {
		query["CollectionQualifiedName"] = request.CollectionQualifiedName
	}

	if !dara.IsNil(request.EntityQualifiedName) {
		query["EntityQualifiedName"] = request.EntityQualifiedName
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMetaCollectionEntity"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMetaCollectionEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds an entity to a collection.
//
// @param request - AddMetaCollectionEntityRequest
//
// @return AddMetaCollectionEntityResponse
func (client *Client) AddMetaCollectionEntity(request *AddMetaCollectionEntityRequest) (_result *AddMetaCollectionEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddMetaCollectionEntityResponse{}
	_body, _err := client.AddMetaCollectionEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns a role to a member of a DataWorks workspace. Before you call this operation, you must add your account to a DataWorks workspace as a member.
//
// Description:
//
//	  For information about how to add an account to a DataWorks workspace as a member, see [Add workspace members and assign roles to them](https://help.aliyun.com/document_detail/136941.html).
//
//		- If you assign a built-in workspace-level role to a member of a DataWorks workspace, the member is automatically granted the permissions of the mapped role of the MaxCompute compute engine in the development environment. For more information, see [Appendix: Mappings between the built-in workspace-level roles of DataWorks and the roles of MaxCompute](https://help.aliyun.com/document_detail/449397.html).
//
// @param request - AddProjectMemberToRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddProjectMemberToRoleResponse
func (client *Client) AddProjectMemberToRoleWithOptions(request *AddProjectMemberToRoleRequest, runtime *dara.RuntimeOptions) (_result *AddProjectMemberToRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCode) {
		query["RoleCode"] = request.RoleCode
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddProjectMemberToRole"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddProjectMemberToRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns a role to a member of a DataWorks workspace. Before you call this operation, you must add your account to a DataWorks workspace as a member.
//
// Description:
//
//	  For information about how to add an account to a DataWorks workspace as a member, see [Add workspace members and assign roles to them](https://help.aliyun.com/document_detail/136941.html).
//
//		- If you assign a built-in workspace-level role to a member of a DataWorks workspace, the member is automatically granted the permissions of the mapped role of the MaxCompute compute engine in the development environment. For more information, see [Appendix: Mappings between the built-in workspace-level roles of DataWorks and the roles of MaxCompute](https://help.aliyun.com/document_detail/449397.html).
//
// @param request - AddProjectMemberToRoleRequest
//
// @return AddProjectMemberToRoleResponse
func (client *Client) AddProjectMemberToRole(request *AddProjectMemberToRoleRequest) (_result *AddProjectMemberToRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddProjectMemberToRoleResponse{}
	_body, _err := client.AddProjectMemberToRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a sensitive field that is defined based on the category and sensitivity level of data in Data Security Guard.
//
// @param request - AddRecognizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecognizeRuleResponse
func (client *Client) AddRecognizeRuleWithOptions(request *AddRecognizeRuleRequest, runtime *dara.RuntimeOptions) (_result *AddRecognizeRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ColExclude) {
		body["ColExclude"] = request.ColExclude
	}

	if !dara.IsNil(request.ColScan) {
		body["ColScan"] = request.ColScan
	}

	if !dara.IsNil(request.CommentScan) {
		body["CommentScan"] = request.CommentScan
	}

	if !dara.IsNil(request.ContentScan) {
		body["ContentScan"] = request.ContentScan
	}

	if !dara.IsNil(request.HitThreshold) {
		body["HitThreshold"] = request.HitThreshold
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.LevelName) {
		body["LevelName"] = request.LevelName
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeParent) {
		body["NodeParent"] = request.NodeParent
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.RecognizeRules) {
		body["RecognizeRules"] = request.RecognizeRules
	}

	if !dara.IsNil(request.RecognizeRulesType) {
		body["RecognizeRulesType"] = request.RecognizeRulesType
	}

	if !dara.IsNil(request.SensitiveDescription) {
		body["SensitiveDescription"] = request.SensitiveDescription
	}

	if !dara.IsNil(request.SensitiveName) {
		body["SensitiveName"] = request.SensitiveName
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddRecognizeRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddRecognizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a sensitive field that is defined based on the category and sensitivity level of data in Data Security Guard.
//
// @param request - AddRecognizeRuleRequest
//
// @return AddRecognizeRuleResponse
func (client *Client) AddRecognizeRule(request *AddRecognizeRuleRequest) (_result *AddRecognizeRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddRecognizeRuleResponse{}
	_body, _err := client.AddRecognizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a metatable to a specified category.
//
// @param request - AddToMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddToMetaCategoryResponse
func (client *Client) AddToMetaCategoryWithOptions(request *AddToMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *AddToMetaCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddToMetaCategory"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddToMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a metatable to a specified category.
//
// @param request - AddToMetaCategoryRequest
//
// @return AddToMetaCategoryResponse
func (client *Client) AddToMetaCategory(request *AddToMetaCategoryRequest) (_result *AddToMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddToMetaCategoryResponse{}
	_body, _err := client.AddToMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Processes a permission request order.
//
// @param request - ApprovePermissionApplyOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApprovePermissionApplyOrderResponse
func (client *Client) ApprovePermissionApplyOrderWithOptions(request *ApprovePermissionApplyOrderRequest, runtime *dara.RuntimeOptions) (_result *ApprovePermissionApplyOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApproveAction) {
		query["ApproveAction"] = request.ApproveAction
	}

	if !dara.IsNil(request.ApproveComment) {
		query["ApproveComment"] = request.ApproveComment
	}

	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApprovePermissionApplyOrder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApprovePermissionApplyOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Processes a permission request order.
//
// @param request - ApprovePermissionApplyOrderRequest
//
// @return ApprovePermissionApplyOrderResponse
func (client *Client) ApprovePermissionApplyOrder(request *ApprovePermissionApplyOrderRequest) (_result *ApprovePermissionApplyOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ApprovePermissionApplyOrderResponse{}
	_body, _err := client.ApprovePermissionApplyOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends the processing result of an extension point event by an extension to DataWorks.
//
// @param request - CallbackExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallbackExtensionResponse
func (client *Client) CallbackExtensionWithOptions(request *CallbackExtensionRequest, runtime *dara.RuntimeOptions) (_result *CallbackExtensionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckMessage) {
		body["CheckMessage"] = request.CheckMessage
	}

	if !dara.IsNil(request.CheckResult) {
		body["CheckResult"] = request.CheckResult
	}

	if !dara.IsNil(request.ExtensionCode) {
		body["ExtensionCode"] = request.ExtensionCode
	}

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CallbackExtension"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CallbackExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends the processing result of an extension point event by an extension to DataWorks.
//
// @param request - CallbackExtensionRequest
//
// @return CallbackExtensionResponse
func (client *Client) CallbackExtension(request *CallbackExtensionRequest) (_result *CallbackExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CallbackExtensionResponse{}
	_body, _err := client.CallbackExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a resource belongs.
//
// @param request - ChangeResourceManagerResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceManagerResourceGroupResponse
func (client *Client) ChangeResourceManagerResourceGroupWithOptions(request *ChangeResourceManagerResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceManagerResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceManagerResourceGroup"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeResourceManagerResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a resource belongs.
//
// @param request - ChangeResourceManagerResourceGroupRequest
//
// @return ChangeResourceManagerResourceGroupResponse
func (client *Client) ChangeResourceManagerResourceGroup(request *ChangeResourceManagerResourceGroupRequest) (_result *ChangeResourceManagerResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeResourceManagerResourceGroupResponse{}
	_body, _err := client.ChangeResourceManagerResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the check events of a file. After you commit your file that is created on the DataStudio page, the system checks the file and returns check events before the system deploys the file. You must determine whether the check can be continued based on the events. You can call this operation to return the check events for the file that you want to deploy to DataWorks.
//
// @param request - CheckFileDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckFileDeploymentResponse
func (client *Client) CheckFileDeploymentWithOptions(request *CheckFileDeploymentRequest, runtime *dara.RuntimeOptions) (_result *CheckFileDeploymentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckDetailUrl) {
		body["CheckDetailUrl"] = request.CheckDetailUrl
	}

	if !dara.IsNil(request.CheckerInstanceId) {
		body["CheckerInstanceId"] = request.CheckerInstanceId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckFileDeployment"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckFileDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the check events of a file. After you commit your file that is created on the DataStudio page, the system checks the file and returns check events before the system deploys the file. You must determine whether the check can be continued based on the events. You can call this operation to return the check events for the file that you want to deploy to DataWorks.
//
// @param request - CheckFileDeploymentRequest
//
// @return CheckFileDeploymentResponse
func (client *Client) CheckFileDeployment(request *CheckFileDeploymentRequest) (_result *CheckFileDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckFileDeploymentResponse{}
	_body, _err := client.CheckFileDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a partition in a MaxCompute metatable exists.
//
// @param request - CheckMetaPartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMetaPartitionResponse
func (client *Client) CheckMetaPartitionWithOptions(request *CheckMetaPartitionRequest, runtime *dara.RuntimeOptions) (_result *CheckMetaPartitionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Partition) {
		query["Partition"] = request.Partition
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckMetaPartition"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckMetaPartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a partition in a MaxCompute metatable exists.
//
// @param request - CheckMetaPartitionRequest
//
// @return CheckMetaPartitionResponse
func (client *Client) CheckMetaPartition(request *CheckMetaPartitionRequest) (_result *CheckMetaPartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckMetaPartitionResponse{}
	_body, _err := client.CheckMetaPartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a metatable exists.
//
// @param request - CheckMetaTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckMetaTableResponse
func (client *Client) CheckMetaTableWithOptions(request *CheckMetaTableRequest, runtime *dara.RuntimeOptions) (_result *CheckMetaTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckMetaTable"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckMetaTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a metatable exists.
//
// @param request - CheckMetaTableRequest
//
// @return CheckMetaTableResponse
func (client *Client) CheckMetaTable(request *CheckMetaTableRequest) (_result *CheckMetaTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CheckMetaTableResponse{}
	_body, _err := client.CheckMetaTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a baseline.
//
// @param request - CreateBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBaselineResponse
func (client *Client) CreateBaselineWithOptions(request *CreateBaselineRequest, runtime *dara.RuntimeOptions) (_result *CreateBaselineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertMarginThreshold) {
		body["AlertMarginThreshold"] = request.AlertMarginThreshold
	}

	if !dara.IsNil(request.BaselineName) {
		body["BaselineName"] = request.BaselineName
	}

	if !dara.IsNil(request.BaselineType) {
		body["BaselineType"] = request.BaselineType
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.OvertimeSettings) {
		body["OvertimeSettings"] = request.OvertimeSettings
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBaseline"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a baseline.
//
// @param request - CreateBaselineRequest
//
// @return CreateBaselineResponse
func (client *Client) CreateBaseline(request *CreateBaselineRequest) (_result *CreateBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBaselineResponse{}
	_body, _err := client.CreateBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow in DataStudio.
//
// @param request - CreateBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBusinessResponse
func (client *Client) CreateBusinessWithOptions(request *CreateBusinessRequest, runtime *dara.RuntimeOptions) (_result *CreateBusinessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.UseType) {
		body["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateBusiness"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow in DataStudio.
//
// @param request - CreateBusinessRequest
//
// @return CreateBusinessResponse
func (client *Client) CreateBusiness(request *CreateBusinessRequest) (_result *CreateBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateBusinessResponse{}
	_body, _err := client.CreateBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateConnection is deprecated, please use dataworks-public::2020-05-18::CreateDataSource instead.
//
// Summary:
//
// Adds a data source.
//
// @param request - CreateConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConnectionResponse
func (client *Client) CreateConnectionWithOptions(request *CreateConnectionRequest, runtime *dara.RuntimeOptions) (_result *CreateConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionType) {
		query["ConnectionType"] = request.ConnectionType
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubType) {
		query["SubType"] = request.SubType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConnection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateConnection is deprecated, please use dataworks-public::2020-05-18::CreateDataSource instead.
//
// Summary:
//
// Adds a data source.
//
// @param request - CreateConnectionRequest
//
// @return CreateConnectionResponse
// Deprecated
func (client *Client) CreateConnection(request *CreateConnectionRequest) (_result *CreateConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateConnectionResponse{}
	_body, _err := client.CreateConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an alert rule for a Data Integration task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization.
//
// @param tmpReq - CreateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRuleWithOptions(tmpReq *CreateDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateDIAlarmRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationSettings) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, dara.String("NotificationSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConditions) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, dara.String("TriggerConditions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.MetricType) {
		body["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.NotificationSettingsShrink) {
		body["NotificationSettings"] = request.NotificationSettingsShrink
	}

	if !dara.IsNil(request.TriggerConditionsShrink) {
		body["TriggerConditions"] = request.TriggerConditionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDIAlarmRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an alert rule for a Data Integration task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization.
//
// @param request - CreateDIAlarmRuleRequest
//
// @return CreateDIAlarmRuleResponse
func (client *Client) CreateDIAlarmRule(request *CreateDIAlarmRuleRequest) (_result *CreateDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDIAlarmRuleResponse{}
	_body, _err := client.CreateDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new-version synchronization task. The following types of synchronization tasks are supported: real-time synchronization of all data in a MySQL database to Hologres and batch synchronization of all data in a MySQL database to Hive.
//
// @param tmpReq - CreateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJobWithOptions(tmpReq *CreateDIJobRequest, runtime *dara.RuntimeOptions) (_result *CreateDIJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DestinationDataSourceSettings) {
		request.DestinationDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestinationDataSourceSettings, dara.String("DestinationDataSourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.JobSettings) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, dara.String("JobSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSettings) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, dara.String("ResourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceDataSourceSettings) {
		request.SourceDataSourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceDataSourceSettings, dara.String("SourceDataSourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableMappings) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, dara.String("TableMappings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransformationRules) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, dara.String("TransformationRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SystemDebug) {
		query["SystemDebug"] = request.SystemDebug
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DestinationDataSourceSettingsShrink) {
		body["DestinationDataSourceSettings"] = request.DestinationDataSourceSettingsShrink
	}

	if !dara.IsNil(request.DestinationDataSourceType) {
		body["DestinationDataSourceType"] = request.DestinationDataSourceType
	}

	if !dara.IsNil(request.JobName) {
		body["JobName"] = request.JobName
	}

	if !dara.IsNil(request.JobSettingsShrink) {
		body["JobSettings"] = request.JobSettingsShrink
	}

	if !dara.IsNil(request.MigrationType) {
		body["MigrationType"] = request.MigrationType
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceSettingsShrink) {
		body["ResourceSettings"] = request.ResourceSettingsShrink
	}

	if !dara.IsNil(request.SourceDataSourceSettingsShrink) {
		body["SourceDataSourceSettings"] = request.SourceDataSourceSettingsShrink
	}

	if !dara.IsNil(request.SourceDataSourceType) {
		body["SourceDataSourceType"] = request.SourceDataSourceType
	}

	if !dara.IsNil(request.TableMappingsShrink) {
		body["TableMappings"] = request.TableMappingsShrink
	}

	if !dara.IsNil(request.TransformationRulesShrink) {
		body["TransformationRules"] = request.TransformationRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDIJob"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new-version synchronization task. The following types of synchronization tasks are supported: real-time synchronization of all data in a MySQL database to Hologres and batch synchronization of all data in a MySQL database to Hive.
//
// @param request - CreateDIJobRequest
//
// @return CreateDIJobResponse
func (client *Client) CreateDIJob(request *CreateDIJobRequest) (_result *CreateDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDIJobResponse{}
	_body, _err := client.CreateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a data synchronization task.
//
// Description:
//
// You cannot configure scheduling properties for a task by calling this operation. If you want to configure scheduling properties for a task, you can call the UpdateFile operation.[](~~2780137~~)
//
// @param request - CreateDISyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDISyncTaskResponse
func (client *Client) CreateDISyncTaskWithOptions(request *CreateDISyncTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDISyncTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskParam) {
		query["TaskParam"] = request.TaskParam
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskContent) {
		body["TaskContent"] = request.TaskContent
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDISyncTask"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDISyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a data synchronization task.
//
// Description:
//
// You cannot configure scheduling properties for a task by calling this operation. If you want to configure scheduling properties for a task, you can call the UpdateFile operation.[](~~2780137~~)
//
// @param request - CreateDISyncTaskRequest
//
// @return CreateDISyncTaskResponse
func (client *Client) CreateDISyncTask(request *CreateDISyncTaskRequest) (_result *CreateDISyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDISyncTaskResponse{}
	_body, _err := client.CreateDISyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateDagComplement is deprecated
//
// @param request - CreateDagComplementRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDagComplementResponse
func (client *Client) CreateDagComplementWithOptions(request *CreateDagComplementRequest, runtime *dara.RuntimeOptions) (_result *CreateDagComplementResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizBeginTime) {
		body["BizBeginTime"] = request.BizBeginTime
	}

	if !dara.IsNil(request.BizEndTime) {
		body["BizEndTime"] = request.BizEndTime
	}

	if !dara.IsNil(request.EndBizDate) {
		body["EndBizDate"] = request.EndBizDate
	}

	if !dara.IsNil(request.ExcludeNodeIds) {
		body["ExcludeNodeIds"] = request.ExcludeNodeIds
	}

	if !dara.IsNil(request.IncludeNodeIds) {
		body["IncludeNodeIds"] = request.IncludeNodeIds
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeParams) {
		body["NodeParams"] = request.NodeParams
	}

	if !dara.IsNil(request.Parallelism) {
		body["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.RootNodeId) {
		body["RootNodeId"] = request.RootNodeId
	}

	if !dara.IsNil(request.StartBizDate) {
		body["StartBizDate"] = request.StartBizDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDagComplement"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDagComplementResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateDagComplement is deprecated
//
// @param request - CreateDagComplementRequest
//
// @return CreateDagComplementResponse
// Deprecated
func (client *Client) CreateDagComplement(request *CreateDagComplementRequest) (_result *CreateDagComplementResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDagComplementResponse{}
	_body, _err := client.CreateDagComplementWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI CreateDagTest is deprecated
//
// @param request - CreateDagTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDagTestResponse
func (client *Client) CreateDagTestWithOptions(request *CreateDagTestRequest, runtime *dara.RuntimeOptions) (_result *CreateDagTestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeParams) {
		body["NodeParams"] = request.NodeParams
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDagTest"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDagTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateDagTest is deprecated
//
// @param request - CreateDagTestRequest
//
// @return CreateDagTestResponse
// Deprecated
func (client *Client) CreateDagTest(request *CreateDagTestRequest) (_result *CreateDagTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDagTestResponse{}
	_body, _err := client.CreateDagTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an API.
//
// @param request - CreateDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApiWithOptions(request *CreateDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiDescription) {
		body["ApiDescription"] = request.ApiDescription
	}

	if !dara.IsNil(request.ApiMode) {
		body["ApiMode"] = request.ApiMode
	}

	if !dara.IsNil(request.ApiName) {
		body["ApiName"] = request.ApiName
	}

	if !dara.IsNil(request.ApiPath) {
		body["ApiPath"] = request.ApiPath
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Protocols) {
		body["Protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RegistrationDetails) {
		body["RegistrationDetails"] = request.RegistrationDetails
	}

	if !dara.IsNil(request.RequestContentType) {
		body["RequestContentType"] = request.RequestContentType
	}

	if !dara.IsNil(request.RequestMethod) {
		body["RequestMethod"] = request.RequestMethod
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResponseContentType) {
		body["ResponseContentType"] = request.ResponseContentType
	}

	if !dara.IsNil(request.ScriptDetails) {
		body["ScriptDetails"] = request.ScriptDetails
	}

	if !dara.IsNil(request.SqlMode) {
		body["SqlMode"] = request.SqlMode
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.VisibleRange) {
		body["VisibleRange"] = request.VisibleRange
	}

	if !dara.IsNil(request.WizardDetails) {
		body["WizardDetails"] = request.WizardDetails
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an API.
//
// @param request - CreateDataServiceApiRequest
//
// @return CreateDataServiceApiResponse
func (client *Client) CreateDataServiceApi(request *CreateDataServiceApiRequest) (_result *CreateDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceApiResponse{}
	_body, _err := client.CreateDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an API in DataService Studio.
//
// @param request - CreateDataServiceApiAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceApiAuthorityResponse
func (client *Client) CreateDataServiceApiAuthorityWithOptions(request *CreateDataServiceApiAuthorityRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceApiAuthorityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AuthorizedProjectId) {
		body["AuthorizedProjectId"] = request.AuthorizedProjectId
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceApiAuthority"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceApiAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants the access permissions on an API in DataService Studio.
//
// @param request - CreateDataServiceApiAuthorityRequest
//
// @return CreateDataServiceApiAuthorityResponse
func (client *Client) CreateDataServiceApiAuthority(request *CreateDataServiceApiAuthorityRequest) (_result *CreateDataServiceApiAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceApiAuthorityResponse{}
	_body, _err := client.CreateDataServiceApiAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a folder in DataService Studio.
//
// @param request - CreateDataServiceFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceFolderResponse
func (client *Client) CreateDataServiceFolderWithOptions(request *CreateDataServiceFolderRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderName) {
		body["FolderName"] = request.FolderName
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ParentId) {
		body["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceFolder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a folder in DataService Studio.
//
// @param request - CreateDataServiceFolderRequest
//
// @return CreateDataServiceFolderResponse
func (client *Client) CreateDataServiceFolder(request *CreateDataServiceFolderRequest) (_result *CreateDataServiceFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceFolderResponse{}
	_body, _err := client.CreateDataServiceFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a business process.
//
// @param request - CreateDataServiceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataServiceGroupResponse
func (client *Client) CreateDataServiceGroupWithOptions(request *CreateDataServiceGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateDataServiceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiGatewayGroupId) {
		body["ApiGatewayGroupId"] = request.ApiGatewayGroupId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.GroupName) {
		body["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataServiceGroup"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a business process.
//
// @param request - CreateDataServiceGroupRequest
//
// @return CreateDataServiceGroupResponse
func (client *Client) CreateDataServiceGroup(request *CreateDataServiceGroupRequest) (_result *CreateDataServiceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataServiceGroupResponse{}
	_body, _err := client.CreateDataServiceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a data source to DataWorks.
//
// @param request - CreateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SubType) {
		query["SubType"] = request.SubType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataSource"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a data source to DataWorks.
//
// @param request - CreateDataSourceRequest
//
// @return CreateDataSourceResponse
func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an export task. You can use this operation to create an export task but cannot use this operation to start the created export task.
//
// @param request - CreateExportMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExportMigrationResponse
func (client *Client) CreateExportMigrationWithOptions(request *CreateExportMigrationRequest, runtime *dara.RuntimeOptions) (_result *CreateExportMigrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.ExportMode) {
		body["ExportMode"] = request.ExportMode
	}

	if !dara.IsNil(request.ExportObjectStatus) {
		body["ExportObjectStatus"] = request.ExportObjectStatus
	}

	if !dara.IsNil(request.IncrementalSince) {
		body["IncrementalSince"] = request.IncrementalSince
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExportMigration"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExportMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an export task. You can use this operation to create an export task but cannot use this operation to start the created export task.
//
// @param request - CreateExportMigrationRequest
//
// @return CreateExportMigrationResponse
func (client *Client) CreateExportMigration(request *CreateExportMigrationRequest) (_result *CreateExportMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateExportMigrationResponse{}
	_body, _err := client.CreateExportMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file in DataStudio. You cannot call this operation to create files for Data Integration nodes.
//
// @param request - CreateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithOptions(request *CreateFileRequest, runtime *dara.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedSettings) {
		body["AdvancedSettings"] = request.AdvancedSettings
	}

	if !dara.IsNil(request.ApplyScheduleImmediately) {
		body["ApplyScheduleImmediately"] = request.ApplyScheduleImmediately
	}

	if !dara.IsNil(request.AutoParsing) {
		body["AutoParsing"] = request.AutoParsing
	}

	if !dara.IsNil(request.AutoRerunIntervalMillis) {
		body["AutoRerunIntervalMillis"] = request.AutoRerunIntervalMillis
	}

	if !dara.IsNil(request.AutoRerunTimes) {
		body["AutoRerunTimes"] = request.AutoRerunTimes
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CreateFolderIfNotExists) {
		body["CreateFolderIfNotExists"] = request.CreateFolderIfNotExists
	}

	if !dara.IsNil(request.CronExpress) {
		body["CronExpress"] = request.CronExpress
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.DependentNodeIdList) {
		body["DependentNodeIdList"] = request.DependentNodeIdList
	}

	if !dara.IsNil(request.DependentType) {
		body["DependentType"] = request.DependentType
	}

	if !dara.IsNil(request.EndEffectDate) {
		body["EndEffectDate"] = request.EndEffectDate
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.IgnoreParentSkipRunningProperty) {
		body["IgnoreParentSkipRunningProperty"] = request.IgnoreParentSkipRunningProperty
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InputList) {
		body["InputList"] = request.InputList
	}

	if !dara.IsNil(request.InputParameters) {
		body["InputParameters"] = request.InputParameters
	}

	if !dara.IsNil(request.OutputParameters) {
		body["OutputParameters"] = request.OutputParameters
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParaValue) {
		body["ParaValue"] = request.ParaValue
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceGroupIdentifier) {
		body["ResourceGroupIdentifier"] = request.ResourceGroupIdentifier
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	if !dara.IsNil(request.StartEffectDate) {
		body["StartEffectDate"] = request.StartEffectDate
	}

	if !dara.IsNil(request.StartImmediately) {
		body["StartImmediately"] = request.StartImmediately
	}

	if !dara.IsNil(request.Stop) {
		body["Stop"] = request.Stop
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file in DataStudio. You cannot call this operation to create files for Data Integration nodes.
//
// @param request - CreateFileRequest
//
// @return CreateFileResponse
func (client *Client) CreateFile(request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The operation that you want to perform. Set the value to \\*\\*CreateFolder\\*\\*.
//
// @param request - CreateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFolderResponse
func (client *Client) CreateFolderWithOptions(request *CreateFolderRequest, runtime *dara.RuntimeOptions) (_result *CreateFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderPath) {
		body["FolderPath"] = request.FolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFolder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The operation that you want to perform. Set the value to \\*\\*CreateFolder\\*\\*.
//
// @param request - CreateFolderRequest
//
// @return CreateFolderResponse
func (client *Client) CreateFolder(request *CreateFolderRequest) (_result *CreateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateFolderResponse{}
	_body, _err := client.CreateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an import task. The import task contains the import packages of data sources, nodes, and tables.
//
// Description:
//
// The import package must be uploaded. Example of the upload method:
//
//	Config config = new Config();
//
//	config.setAccessKeyId(accessId);
//
//	config.setAccessKeySecret(accessKey);
//
//	config.setEndpoint(popEndpoint);
//
//	config.setRegionId(regionId);
//
//
//
//	Client client = new Client(config);
//
//	CreateImportMigrationAdvanceRequest request = new CreateImportMigrationAdvanceRequest();
//
//	request.setName("test_migration_api_" + System.currentTimeMillis());
//
//	request.setProjectId(123456L);
//
//	request.setPackageType("DATAWORKS_MODEL");
//
//	request.setPackageFileObject(new FileInputStream("/home/admin/Downloads/test.zip"));
//
//	RuntimeOptions runtime = new RuntimeOptions();
//
//	CreateImportMigrationResponse response = client.createImportMigrationAdvance(request, runtime);
//
//	...
//
// @param request - CreateImportMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImportMigrationResponse
func (client *Client) CreateImportMigrationWithOptions(request *CreateImportMigrationRequest, runtime *dara.RuntimeOptions) (_result *CreateImportMigrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CalculateEngineMap) {
		body["CalculateEngineMap"] = request.CalculateEngineMap
	}

	if !dara.IsNil(request.CommitRule) {
		body["CommitRule"] = request.CommitRule
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.PackageFile) {
		body["PackageFile"] = request.PackageFile
	}

	if !dara.IsNil(request.PackageType) {
		body["PackageType"] = request.PackageType
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroupMap) {
		body["ResourceGroupMap"] = request.ResourceGroupMap
	}

	if !dara.IsNil(request.WorkspaceMap) {
		body["WorkspaceMap"] = request.WorkspaceMap
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImportMigration"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImportMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an import task. The import task contains the import packages of data sources, nodes, and tables.
//
// Description:
//
// The import package must be uploaded. Example of the upload method:
//
//	Config config = new Config();
//
//	config.setAccessKeyId(accessId);
//
//	config.setAccessKeySecret(accessKey);
//
//	config.setEndpoint(popEndpoint);
//
//	config.setRegionId(regionId);
//
//
//
//	Client client = new Client(config);
//
//	CreateImportMigrationAdvanceRequest request = new CreateImportMigrationAdvanceRequest();
//
//	request.setName("test_migration_api_" + System.currentTimeMillis());
//
//	request.setProjectId(123456L);
//
//	request.setPackageType("DATAWORKS_MODEL");
//
//	request.setPackageFileObject(new FileInputStream("/home/admin/Downloads/test.zip"));
//
//	RuntimeOptions runtime = new RuntimeOptions();
//
//	CreateImportMigrationResponse response = client.createImportMigrationAdvance(request, runtime);
//
//	...
//
// @param request - CreateImportMigrationRequest
//
// @return CreateImportMigrationResponse
func (client *Client) CreateImportMigration(request *CreateImportMigrationRequest) (_result *CreateImportMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateImportMigrationResponse{}
	_body, _err := client.CreateImportMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImportMigrationAdvance(request *CreateImportMigrationAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CreateImportMigrationResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("dataworks-public"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	createImportMigrationReq := &CreateImportMigrationRequest{}
	openapiutil.Convert(request, createImportMigrationReq)
	if !dara.IsNil(request.PackageFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.PackageFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		createImportMigrationReq.PackageFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createImportMigrationResp, _err := client.CreateImportMigrationWithOptions(createImportMigrationReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createImportMigrationResp
	return _result, _err
}

// Deprecated: OpenAPI CreateManualDag is deprecated
//
// Summary:
//
// Triggers a manually triggered workflow to run. Before you call this operation, make sure that the manually triggered workflow is committed and deployed. You can find the manually triggered workflow on the Operation Center page only after the manually triggered workflow is committed and deployed.
//
// @param request - CreateManualDagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateManualDagResponse
func (client *Client) CreateManualDagWithOptions(request *CreateManualDagRequest, runtime *dara.RuntimeOptions) (_result *CreateManualDagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.DagParameters) {
		body["DagParameters"] = request.DagParameters
	}

	if !dara.IsNil(request.ExcludeNodeIds) {
		body["ExcludeNodeIds"] = request.ExcludeNodeIds
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.IncludeNodeIds) {
		body["IncludeNodeIds"] = request.IncludeNodeIds
	}

	if !dara.IsNil(request.NodeParameters) {
		body["NodeParameters"] = request.NodeParameters
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateManualDag"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateManualDagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI CreateManualDag is deprecated
//
// Summary:
//
// Triggers a manually triggered workflow to run. Before you call this operation, make sure that the manually triggered workflow is committed and deployed. You can find the manually triggered workflow on the Operation Center page only after the manually triggered workflow is committed and deployed.
//
// @param request - CreateManualDagRequest
//
// @return CreateManualDagResponse
// Deprecated
func (client *Client) CreateManualDag(request *CreateManualDagRequest) (_result *CreateManualDagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateManualDagResponse{}
	_body, _err := client.CreateManualDagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a category.
//
// @param request - CreateMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetaCategoryResponse
func (client *Client) CreateMetaCategoryWithOptions(request *CreateMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *CreateMetaCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		body["ParentId"] = request.ParentId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetaCategory"),
		Version:     dara.String("2020-05-18"),
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
// Creates a category.
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
// Creates a collection.
//
// Description:
//
// Collections are classified into various types. The names of collections of the same type must be different.
//
// @param request - CreateMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMetaCollectionResponse
func (client *Client) CreateMetaCollectionWithOptions(request *CreateMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *CreateMetaCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CollectionType) {
		query["CollectionType"] = request.CollectionType
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentQualifiedName) {
		query["ParentQualifiedName"] = request.ParentQualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMetaCollection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a collection.
//
// Description:
//
// Collections are classified into various types. The names of collections of the same type must be different.
//
// @param request - CreateMetaCollectionRequest
//
// @return CreateMetaCollectionResponse
func (client *Client) CreateMetaCollection(request *CreateMetaCollectionRequest) (_result *CreateMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateMetaCollectionResponse{}
	_body, _err := client.CreateMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a permission request order.
//
// @param request - CreatePermissionApplyOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePermissionApplyOrderResponse
func (client *Client) CreatePermissionApplyOrderWithOptions(request *CreatePermissionApplyOrderRequest, runtime *dara.RuntimeOptions) (_result *CreatePermissionApplyOrderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyObject) {
		query["ApplyObject"] = request.ApplyObject
	}

	if !dara.IsNil(request.ApplyReason) {
		query["ApplyReason"] = request.ApplyReason
	}

	if !dara.IsNil(request.ApplyType) {
		query["ApplyType"] = request.ApplyType
	}

	if !dara.IsNil(request.ApplyUserIds) {
		query["ApplyUserIds"] = request.ApplyUserIds
	}

	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.Deadline) {
		query["Deadline"] = request.Deadline
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.MaxComputeProjectName) {
		query["MaxComputeProjectName"] = request.MaxComputeProjectName
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePermissionApplyOrder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePermissionApplyOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a permission request order.
//
// @param request - CreatePermissionApplyOrderRequest
//
// @return CreatePermissionApplyOrderResponse
func (client *Client) CreatePermissionApplyOrder(request *CreatePermissionApplyOrderRequest) (_result *CreatePermissionApplyOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePermissionApplyOrderResponse{}
	_body, _err := client.CreatePermissionApplyOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a DataWorks workspace.
//
// @param tmpReq - CreateProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectResponse
func (client *Client) CreateProjectWithOptions(tmpReq *CreateProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DisableDevelopment) {
		query["DisableDevelopment"] = request.DisableDevelopment
	}

	if !dara.IsNil(request.IsAllowDownload) {
		query["IsAllowDownload"] = request.IsAllowDownload
	}

	if !dara.IsNil(request.ProjectDescription) {
		query["ProjectDescription"] = request.ProjectDescription
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		query["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.ProjectMode) {
		query["ProjectMode"] = request.ProjectMode
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProject"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a DataWorks workspace.
//
// @param request - CreateProjectRequest
//
// @return CreateProjectResponse
func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a user to a DataWorks workspace.
//
// @param request - CreateProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMemberWithOptions(request *CreateProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *CreateProjectMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCode) {
		query["RoleCode"] = request.RoleCode
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProjectMember"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a user to a DataWorks workspace.
//
// @param request - CreateProjectMemberRequest
//
// @return CreateProjectMemberResponse
func (client *Client) CreateProjectMember(request *CreateProjectMemberRequest) (_result *CreateProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateProjectMemberResponse{}
	_body, _err := client.CreateProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a partition filter expression.
//
// @param request - CreateQualityEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityEntityResponse
func (client *Client) CreateQualityEntityWithOptions(request *CreateQualityEntityRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityLevel) {
		body["EntityLevel"] = request.EntityLevel
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.MatchExpression) {
		body["MatchExpression"] = request.MatchExpression
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityEntity"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a partition filter expression.
//
// @param request - CreateQualityEntityRequest
//
// @return CreateQualityEntityResponse
func (client *Client) CreateQualityEntity(request *CreateQualityEntityRequest) (_result *CreateQualityEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQualityEntityResponse{}
	_body, _err := client.CreateQualityEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a subscriber for a partition filter expression.
//
// @param request - CreateQualityFollowerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityFollowerResponse
func (client *Client) CreateQualityFollowerWithOptions(request *CreateQualityFollowerRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityFollowerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmMode) {
		body["AlarmMode"] = request.AlarmMode
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.Follower) {
		body["Follower"] = request.Follower
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityFollower"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityFollowerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a subscriber for a partition filter expression.
//
// @param request - CreateQualityFollowerRequest
//
// @return CreateQualityFollowerResponse
func (client *Client) CreateQualityFollower(request *CreateQualityFollowerRequest) (_result *CreateQualityFollowerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQualityFollowerResponse{}
	_body, _err := client.CreateQualityFollowerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a node with a partition filter expression.
//
// @param request - CreateQualityRelativeNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityRelativeNodeResponse
func (client *Client) CreateQualityRelativeNodeWithOptions(request *CreateQualityRelativeNodeRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityRelativeNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.MatchExpression) {
		body["MatchExpression"] = request.MatchExpression
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TargetNodeProjectId) {
		body["TargetNodeProjectId"] = request.TargetNodeProjectId
	}

	if !dara.IsNil(request.TargetNodeProjectName) {
		body["TargetNodeProjectName"] = request.TargetNodeProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityRelativeNode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityRelativeNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a node with a partition filter expression.
//
// @param request - CreateQualityRelativeNodeRequest
//
// @return CreateQualityRelativeNodeResponse
func (client *Client) CreateQualityRelativeNode(request *CreateQualityRelativeNodeRequest) (_result *CreateQualityRelativeNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQualityRelativeNodeResponse{}
	_body, _err := client.CreateQualityRelativeNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a monitoring rule.
//
// @param request - CreateQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityRuleResponse
func (client *Client) CreateQualityRuleWithOptions(request *CreateQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *CreateQualityRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockType) {
		body["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.Checker) {
		body["Checker"] = request.Checker
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CriticalThreshold) {
		body["CriticalThreshold"] = request.CriticalThreshold
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ExpectValue) {
		body["ExpectValue"] = request.ExpectValue
	}

	if !dara.IsNil(request.MethodName) {
		body["MethodName"] = request.MethodName
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.PredictType) {
		body["PredictType"] = request.PredictType
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Property) {
		body["Property"] = request.Property
	}

	if !dara.IsNil(request.PropertyType) {
		body["PropertyType"] = request.PropertyType
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.TaskSetting) {
		body["TaskSetting"] = request.TaskSetting
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Trend) {
		body["Trend"] = request.Trend
	}

	if !dara.IsNil(request.WarningThreshold) {
		body["WarningThreshold"] = request.WarningThreshold
	}

	if !dara.IsNil(request.WhereCondition) {
		body["WhereCondition"] = request.WhereCondition
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a monitoring rule.
//
// @param request - CreateQualityRuleRequest
//
// @return CreateQualityRuleResponse
func (client *Client) CreateQualityRule(request *CreateQualityRuleRequest) (_result *CreateQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateQualityRuleResponse{}
	_body, _err := client.CreateQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom alert rule.
//
// @param request - CreateRemindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRemindResponse
func (client *Client) CreateRemindWithOptions(request *CreateRemindRequest, runtime *dara.RuntimeOptions) (_result *CreateRemindResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertInterval) {
		body["AlertInterval"] = request.AlertInterval
	}

	if !dara.IsNil(request.AlertMethods) {
		body["AlertMethods"] = request.AlertMethods
	}

	if !dara.IsNil(request.AlertTargets) {
		body["AlertTargets"] = request.AlertTargets
	}

	if !dara.IsNil(request.AlertUnit) {
		body["AlertUnit"] = request.AlertUnit
	}

	if !dara.IsNil(request.BaselineIds) {
		body["BaselineIds"] = request.BaselineIds
	}

	if !dara.IsNil(request.BizProcessIds) {
		body["BizProcessIds"] = request.BizProcessIds
	}

	if !dara.IsNil(request.Detail) {
		body["Detail"] = request.Detail
	}

	if !dara.IsNil(request.DndEnd) {
		body["DndEnd"] = request.DndEnd
	}

	if !dara.IsNil(request.MaxAlertTimes) {
		body["MaxAlertTimes"] = request.MaxAlertTimes
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RemindName) {
		body["RemindName"] = request.RemindName
	}

	if !dara.IsNil(request.RemindType) {
		body["RemindType"] = request.RemindType
	}

	if !dara.IsNil(request.RemindUnit) {
		body["RemindUnit"] = request.RemindUnit
	}

	if !dara.IsNil(request.RobotUrls) {
		body["RobotUrls"] = request.RobotUrls
	}

	if !dara.IsNil(request.Webhooks) {
		body["Webhooks"] = request.Webhooks
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRemind"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRemindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom alert rule.
//
// @param request - CreateRemindRequest
//
// @return CreateRemindResponse
func (client *Client) CreateRemind(request *CreateRemindRequest) (_result *CreateRemindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRemindResponse{}
	_body, _err := client.CreateRemindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates or uploads a resource file in DataStudio. The feature that is implemented by calling this operation is the same as the resource creation feature provided in the integrated development environment (IDE).
//
// @param request - CreateResourceFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceFileResponse
func (client *Client) CreateResourceFileWithOptions(request *CreateResourceFileRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.OriginResourceName) {
		body["OriginResourceName"] = request.OriginResourceName
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RegisterToCalcEngine) {
		body["RegisterToCalcEngine"] = request.RegisterToCalcEngine
	}

	if !dara.IsNil(request.ResourceFile) {
		body["ResourceFile"] = request.ResourceFile
	}

	if !dara.IsNil(request.StorageURL) {
		body["StorageURL"] = request.StorageURL
	}

	if !dara.IsNil(request.UploadMode) {
		body["UploadMode"] = request.UploadMode
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateResourceFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateResourceFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates or uploads a resource file in DataStudio. The feature that is implemented by calling this operation is the same as the resource creation feature provided in the integrated development environment (IDE).
//
// @param request - CreateResourceFileRequest
//
// @return CreateResourceFileResponse
func (client *Client) CreateResourceFile(request *CreateResourceFileRequest) (_result *CreateResourceFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateResourceFileResponse{}
	_body, _err := client.CreateResourceFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceFileAdvance(request *CreateResourceFileAdvanceRequest, runtime *dara.RuntimeOptions) (_result *CreateResourceFileResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("dataworks-public"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	createResourceFileReq := &CreateResourceFileRequest{}
	openapiutil.Convert(request, createResourceFileReq)
	if !dara.IsNil(request.ResourceFileObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.ResourceFileObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		createResourceFileReq.ResourceFile = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	createResourceFileResp, _err := client.CreateResourceFileWithOptions(createResourceFileReq, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createResourceFileResp
	return _result, _err
}

// Summary:
//
// Creates a MaxCompute table or view.
//
// @param request - CreateTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableResponse
func (client *Client) CreateTableWithOptions(request *CreateTableRequest, runtime *dara.RuntimeOptions) (_result *CreateTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGuid) {
		query["AppGuid"] = request.AppGuid
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.ExternalTableType) {
		query["ExternalTableType"] = request.ExternalTableType
	}

	if !dara.IsNil(request.HasPart) {
		query["HasPart"] = request.HasPart
	}

	if !dara.IsNil(request.IsView) {
		query["IsView"] = request.IsView
	}

	if !dara.IsNil(request.LifeCycle) {
		query["LifeCycle"] = request.LifeCycle
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.LogicalLevelId) {
		query["LogicalLevelId"] = request.LogicalLevelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhysicsLevelId) {
		query["PhysicsLevelId"] = request.PhysicsLevelId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Columns) {
		body["Columns"] = request.Columns
	}

	if !dara.IsNil(request.Endpoint) {
		body["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Themes) {
		body["Themes"] = request.Themes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTable"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a MaxCompute table or view.
//
// @param request - CreateTableRequest
//
// @return CreateTableResponse
func (client *Client) CreateTable(request *CreateTableRequest) (_result *CreateTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTableResponse{}
	_body, _err := client.CreateTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a table level. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - CreateTableLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableLevelResponse
func (client *Client) CreateTableLevelWithOptions(request *CreateTableLevelRequest, runtime *dara.RuntimeOptions) (_result *CreateTableLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.LevelType) {
		query["LevelType"] = request.LevelType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTableLevel"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTableLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a table level. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - CreateTableLevelRequest
//
// @return CreateTableLevelResponse
func (client *Client) CreateTableLevel(request *CreateTableLevelRequest) (_result *CreateTableLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTableLevelResponse{}
	_body, _err := client.CreateTableLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a table theme. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - CreateTableThemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTableThemeResponse
func (client *Client) CreateTableThemeWithOptions(request *CreateTableThemeRequest, runtime *dara.RuntimeOptions) (_result *CreateTableThemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Level) {
		query["Level"] = request.Level
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ParentId) {
		query["ParentId"] = request.ParentId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTableTheme"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTableThemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a table theme. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - CreateTableThemeRequest
//
// @return CreateTableThemeResponse
func (client *Client) CreateTableTheme(request *CreateTableThemeRequest) (_result *CreateTableThemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTableThemeResponse{}
	_body, _err := client.CreateTableThemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file for a function in DataStudio.
//
// @param request - CreateUdfFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUdfFileResponse
func (client *Client) CreateUdfFileWithOptions(request *CreateUdfFileRequest, runtime *dara.RuntimeOptions) (_result *CreateUdfFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.CmdDescription) {
		body["CmdDescription"] = request.CmdDescription
	}

	if !dara.IsNil(request.CreateFolderIfNotExists) {
		body["CreateFolderIfNotExists"] = request.CreateFolderIfNotExists
	}

	if !dara.IsNil(request.Example) {
		body["Example"] = request.Example
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ParameterDescription) {
		body["ParameterDescription"] = request.ParameterDescription
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ReturnValue) {
		body["ReturnValue"] = request.ReturnValue
	}

	if !dara.IsNil(request.UdfDescription) {
		body["UdfDescription"] = request.UdfDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateUdfFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateUdfFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file for a function in DataStudio.
//
// @param request - CreateUdfFileRequest
//
// @return CreateUdfFileResponse
func (client *Client) CreateUdfFile(request *CreateUdfFileRequest) (_result *CreateUdfFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateUdfFileResponse{}
	_body, _err := client.CreateUdfFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a baseline based on its ID. You can delete a baseline only if the nodes in the baseline does not have ancestor nodes. You can call the UpdateBaseline operation to delete the relationships between the nodes and their ancestor nodes.
//
// @param request - DeleteBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBaselineResponse
func (client *Client) DeleteBaselineWithOptions(request *DeleteBaselineRequest, runtime *dara.RuntimeOptions) (_result *DeleteBaselineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBaseline"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a baseline based on its ID. You can delete a baseline only if the nodes in the baseline does not have ancestor nodes. You can call the UpdateBaseline operation to delete the relationships between the nodes and their ancestor nodes.
//
// @param request - DeleteBaselineRequest
//
// @return DeleteBaselineResponse
func (client *Client) DeleteBaseline(request *DeleteBaselineRequest) (_result *DeleteBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBaselineResponse{}
	_body, _err := client.DeleteBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBusinessResponse
func (client *Client) DeleteBusinessWithOptions(request *DeleteBusinessRequest, runtime *dara.RuntimeOptions) (_result *DeleteBusinessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBusiness"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteBusinessRequest
//
// @return DeleteBusinessResponse
func (client *Client) DeleteBusiness(request *DeleteBusinessRequest) (_result *DeleteBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteBusinessResponse{}
	_body, _err := client.DeleteBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI DeleteConnection is deprecated
//
// Summary:
//
// Removes a data source.
//
// @param request - DeleteConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConnectionResponse
func (client *Client) DeleteConnectionWithOptions(request *DeleteConnectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteConnection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI DeleteConnection is deprecated
//
// Summary:
//
// Removes a data source.
//
// @param request - DeleteConnectionRequest
//
// @return DeleteConnectionResponse
// Deprecated
func (client *Client) DeleteConnection(request *DeleteConnectionRequest) (_result *DeleteConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteConnectionResponse{}
	_body, _err := client.DeleteConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an alert rule for a Data Integration task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks whose MigrationType is set to RealtimeIncremental.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRuleWithOptions(request *DeleteDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteDIAlarmRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DIAlarmRuleId) {
		body["DIAlarmRuleId"] = request.DIAlarmRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDIAlarmRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an alert rule for a Data Integration task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks whose MigrationType is set to RealtimeIncremental.
//
// @param request - DeleteDIAlarmRuleRequest
//
// @return DeleteDIAlarmRuleResponse
func (client *Client) DeleteDIAlarmRule(request *DeleteDIAlarmRuleRequest) (_result *DeleteDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDIAlarmRuleResponse{}
	_body, _err := client.DeleteDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Data Integration task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// @param request - DeleteDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJobWithOptions(request *DeleteDIJobRequest, runtime *dara.RuntimeOptions) (_result *DeleteDIJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDIJob"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Data Integration task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// @param request - DeleteDIJobRequest
//
// @return DeleteDIJobResponse
func (client *Client) DeleteDIJob(request *DeleteDIJobRequest) (_result *DeleteDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDIJobResponse{}
	_body, _err := client.DeleteDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a synchronization task. You can call this operation to delete only a real-time synchronization task.
//
// Description:
//
// If you want to delete a batch synchronization task, call the DeleteFile operation. For more information, see [Delete a synchronization task](https://help.aliyun.com/document_detail/321443.html).
//
// @param request - DeleteDISyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDISyncTaskResponse
func (client *Client) DeleteDISyncTaskWithOptions(request *DeleteDISyncTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteDISyncTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDISyncTask"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDISyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a synchronization task. You can call this operation to delete only a real-time synchronization task.
//
// Description:
//
// If you want to delete a batch synchronization task, call the DeleteFile operation. For more information, see [Delete a synchronization task](https://help.aliyun.com/document_detail/321443.html).
//
// @param request - DeleteDISyncTaskRequest
//
// @return DeleteDISyncTaskResponse
func (client *Client) DeleteDISyncTask(request *DeleteDISyncTaskRequest) (_result *DeleteDISyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDISyncTaskResponse{}
	_body, _err := client.DeleteDISyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an API in DataService Studio.
//
// @param request - DeleteDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataServiceApiResponse
func (client *Client) DeleteDataServiceApiWithOptions(request *DeleteDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an API in DataService Studio.
//
// @param request - DeleteDataServiceApiRequest
//
// @return DeleteDataServiceApiResponse
func (client *Client) DeleteDataServiceApi(request *DeleteDataServiceApiRequest) (_result *DeleteDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataServiceApiResponse{}
	_body, _err := client.DeleteDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on an API.
//
// @param request - DeleteDataServiceApiAuthorityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataServiceApiAuthorityResponse
func (client *Client) DeleteDataServiceApiAuthorityWithOptions(request *DeleteDataServiceApiAuthorityRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataServiceApiAuthorityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AuthorizedProjectId) {
		body["AuthorizedProjectId"] = request.AuthorizedProjectId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataServiceApiAuthority"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataServiceApiAuthorityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes the access permissions on an API.
//
// @param request - DeleteDataServiceApiAuthorityRequest
//
// @return DeleteDataServiceApiAuthorityResponse
func (client *Client) DeleteDataServiceApiAuthority(request *DeleteDataServiceApiAuthorityRequest) (_result *DeleteDataServiceApiAuthorityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataServiceApiAuthorityResponse{}
	_body, _err := client.DeleteDataServiceApiAuthorityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a data source.
//
// @param request - DeleteDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSourceWithOptions(request *DeleteDataSourceRequest, runtime *dara.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDataSource"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a data source.
//
// @param request - DeleteDataSourceRequest
//
// @return DeleteDataSourceResponse
func (client *Client) DeleteDataSource(request *DeleteDataSourceRequest) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file from DataStudio. If the file has been committed, an asynchronous process is triggered to delete the file in the scheduling system. The value of the DeploymentId parameter returned is used to call the GetDeployment operation to poll the status of the asynchronous process.
//
// @param request - DeleteFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, runtime *dara.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file from DataStudio. If the file has been committed, an asynchronous process is triggered to delete the file in the scheduling system. The value of the DeploymentId parameter returned is used to call the GetDeployment operation to poll the status of the asynchronous process.
//
// @param request - DeleteFileRequest
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteFolderRequest
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(request *DeleteFolderRequest) (_result *DeleteFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a table from a specified category.
//
// @param request - DeleteFromMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFromMetaCategoryResponse
func (client *Client) DeleteFromMetaCategoryWithOptions(request *DeleteFromMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteFromMetaCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFromMetaCategory"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFromMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a table from a specified category.
//
// @param request - DeleteFromMetaCategoryRequest
//
// @return DeleteFromMetaCategoryResponse
func (client *Client) DeleteFromMetaCategory(request *DeleteFromMetaCategoryRequest) (_result *DeleteFromMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteFromMetaCategoryResponse{}
	_body, _err := client.DeleteFromMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete lineage, supports deleting user-defined lineage relationships
//
// Description:
//
// This API is currently in the trial phase. Users who wish to experience it can apply, and after the administrator adds them to the trial list, they can call this API.
//
// @param request - DeleteLineageRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLineageRelationResponse
func (client *Client) DeleteLineageRelationWithOptions(request *DeleteLineageRelationRequest, runtime *dara.RuntimeOptions) (_result *DeleteLineageRelationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestEntityQualifiedName) {
		query["DestEntityQualifiedName"] = request.DestEntityQualifiedName
	}

	if !dara.IsNil(request.RelationshipGuid) {
		query["RelationshipGuid"] = request.RelationshipGuid
	}

	if !dara.IsNil(request.RelationshipType) {
		query["RelationshipType"] = request.RelationshipType
	}

	if !dara.IsNil(request.SrcEntityQualifiedName) {
		query["SrcEntityQualifiedName"] = request.SrcEntityQualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLineageRelation"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLineageRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete lineage, supports deleting user-defined lineage relationships
//
// Description:
//
// This API is currently in the trial phase. Users who wish to experience it can apply, and after the administrator adds them to the trial list, they can call this API.
//
// @param request - DeleteLineageRelationRequest
//
// @return DeleteLineageRelationResponse
func (client *Client) DeleteLineageRelation(request *DeleteLineageRelationRequest) (_result *DeleteLineageRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteLineageRelationResponse{}
	_body, _err := client.DeleteLineageRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a category.
//
// @param request - DeleteMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaCategoryResponse
func (client *Client) DeleteMetaCategoryWithOptions(request *DeleteMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetaCategory"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Deletes a category.
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
// Deletes a collection.
//
// @param request - DeleteMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaCollectionResponse
func (client *Client) DeleteMetaCollectionWithOptions(request *DeleteMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QualifiedName) {
		query["QualifiedName"] = request.QualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetaCollection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a collection.
//
// @param request - DeleteMetaCollectionRequest
//
// @return DeleteMetaCollectionResponse
func (client *Client) DeleteMetaCollection(request *DeleteMetaCollectionRequest) (_result *DeleteMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetaCollectionResponse{}
	_body, _err := client.DeleteMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an entity from a collection.
//
// @param request - DeleteMetaCollectionEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMetaCollectionEntityResponse
func (client *Client) DeleteMetaCollectionEntityWithOptions(request *DeleteMetaCollectionEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteMetaCollectionEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CollectionQualifiedName) {
		query["CollectionQualifiedName"] = request.CollectionQualifiedName
	}

	if !dara.IsNil(request.EntityQualifiedName) {
		query["EntityQualifiedName"] = request.EntityQualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMetaCollectionEntity"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMetaCollectionEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an entity from a collection.
//
// @param request - DeleteMetaCollectionEntityRequest
//
// @return DeleteMetaCollectionEntityResponse
func (client *Client) DeleteMetaCollectionEntity(request *DeleteMetaCollectionEntityRequest) (_result *DeleteMetaCollectionEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteMetaCollectionEntityResponse{}
	_body, _err := client.DeleteMetaCollectionEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a user from a DataWorks workspace.
//
// @param request - DeleteProjectMemberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMemberWithOptions(request *DeleteProjectMemberRequest, runtime *dara.RuntimeOptions) (_result *DeleteProjectMemberResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteProjectMember"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a user from a DataWorks workspace.
//
// @param request - DeleteProjectMemberRequest
//
// @return DeleteProjectMemberResponse
func (client *Client) DeleteProjectMember(request *DeleteProjectMemberRequest) (_result *DeleteProjectMemberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteProjectMemberResponse{}
	_body, _err := client.DeleteProjectMemberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a partition filter expression.
//
// @param request - DeleteQualityEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityEntityResponse
func (client *Client) DeleteQualityEntityWithOptions(request *DeleteQualityEntityRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityEntity"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a partition filter expression.
//
// @param request - DeleteQualityEntityRequest
//
// @return DeleteQualityEntityResponse
func (client *Client) DeleteQualityEntity(request *DeleteQualityEntityRequest) (_result *DeleteQualityEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityEntityResponse{}
	_body, _err := client.DeleteQualityEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a subscriber of a partition filter expression.
//
// Description:
//
// In Data Quality, you must configure monitoring rules based on a partition filter expression. Data Quality uses these rules to detect changes in source data and dirty data generated during the process of extract, transform, and load (ETL). This way, you can prevent tasks from producing unexpected dirty data that affects the smooth running of tasks and business decision-making. You can go to the Manage Subscriptions page to add subscribers for a partition filter expression. When the monitoring rule that is created based on the partition filter expression is triggered, the subscribers can receive notifications and troubleshoot errors at the earliest opportunity. For more information, see [Configure monitoring rules](https://help.aliyun.com/document_detail/73690.html).
//
// @param request - DeleteQualityFollowerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityFollowerResponse
func (client *Client) DeleteQualityFollowerWithOptions(request *DeleteQualityFollowerRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityFollowerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FollowerId) {
		body["FollowerId"] = request.FollowerId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityFollower"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityFollowerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a subscriber of a partition filter expression.
//
// Description:
//
// In Data Quality, you must configure monitoring rules based on a partition filter expression. Data Quality uses these rules to detect changes in source data and dirty data generated during the process of extract, transform, and load (ETL). This way, you can prevent tasks from producing unexpected dirty data that affects the smooth running of tasks and business decision-making. You can go to the Manage Subscriptions page to add subscribers for a partition filter expression. When the monitoring rule that is created based on the partition filter expression is triggered, the subscribers can receive notifications and troubleshoot errors at the earliest opportunity. For more information, see [Configure monitoring rules](https://help.aliyun.com/document_detail/73690.html).
//
// @param request - DeleteQualityFollowerRequest
//
// @return DeleteQualityFollowerResponse
func (client *Client) DeleteQualityFollower(request *DeleteQualityFollowerRequest) (_result *DeleteQualityFollowerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityFollowerResponse{}
	_body, _err := client.DeleteQualityFollowerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteQualityRelativeNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityRelativeNodeResponse
func (client *Client) DeleteQualityRelativeNodeWithOptions(request *DeleteQualityRelativeNodeRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityRelativeNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.MatchExpression) {
		body["MatchExpression"] = request.MatchExpression
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	if !dara.IsNil(request.TargetNodeProjectId) {
		body["TargetNodeProjectId"] = request.TargetNodeProjectId
	}

	if !dara.IsNil(request.TargetNodeProjectName) {
		body["TargetNodeProjectName"] = request.TargetNodeProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityRelativeNode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityRelativeNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteQualityRelativeNodeRequest
//
// @return DeleteQualityRelativeNodeResponse
func (client *Client) DeleteQualityRelativeNode(request *DeleteQualityRelativeNodeRequest) (_result *DeleteQualityRelativeNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityRelativeNodeResponse{}
	_body, _err := client.DeleteQualityRelativeNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a monitoring rule.
//
// @param request - DeleteQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQualityRuleResponse
func (client *Client) DeleteQualityRuleWithOptions(request *DeleteQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteQualityRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQualityRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a monitoring rule.
//
// @param request - DeleteQualityRuleRequest
//
// @return DeleteQualityRuleResponse
func (client *Client) DeleteQualityRule(request *DeleteQualityRuleRequest) (_result *DeleteQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteQualityRuleResponse{}
	_body, _err := client.DeleteQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes sensitive field types.
//
// @param request - DeleteRecognizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecognizeRuleResponse
func (client *Client) DeleteRecognizeRuleWithOptions(request *DeleteRecognizeRuleRequest, runtime *dara.RuntimeOptions) (_result *DeleteRecognizeRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SensitiveId) {
		body["SensitiveId"] = request.SensitiveId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRecognizeRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRecognizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes sensitive field types.
//
// @param request - DeleteRecognizeRuleRequest
//
// @return DeleteRecognizeRuleResponse
func (client *Client) DeleteRecognizeRule(request *DeleteRecognizeRuleRequest) (_result *DeleteRecognizeRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRecognizeRuleResponse{}
	_body, _err := client.DeleteRecognizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom alert rule.
//
// @param request - DeleteRemindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRemindResponse
func (client *Client) DeleteRemindWithOptions(request *DeleteRemindRequest, runtime *dara.RuntimeOptions) (_result *DeleteRemindResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RemindId) {
		body["RemindId"] = request.RemindId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRemind"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRemindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom alert rule.
//
// @param request - DeleteRemindRequest
//
// @return DeleteRemindResponse
func (client *Client) DeleteRemind(request *DeleteRemindRequest) (_result *DeleteRemindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRemindResponse{}
	_body, _err := client.DeleteRemindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableResponse
func (client *Client) DeleteTableWithOptions(request *DeleteTableRequest, runtime *dara.RuntimeOptions) (_result *DeleteTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGuid) {
		query["AppGuid"] = request.AppGuid
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTable"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteTableRequest
//
// @return DeleteTableResponse
func (client *Client) DeleteTable(request *DeleteTableRequest) (_result *DeleteTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTableResponse{}
	_body, _err := client.DeleteTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a table level. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - DeleteTableLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableLevelResponse
func (client *Client) DeleteTableLevelWithOptions(request *DeleteTableLevelRequest, runtime *dara.RuntimeOptions) (_result *DeleteTableLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LevelId) {
		query["LevelId"] = request.LevelId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTableLevel"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTableLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a table level. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - DeleteTableLevelRequest
//
// @return DeleteTableLevelResponse
func (client *Client) DeleteTableLevel(request *DeleteTableLevelRequest) (_result *DeleteTableLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTableLevelResponse{}
	_body, _err := client.DeleteTableLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a table theme. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - DeleteTableThemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTableThemeResponse
func (client *Client) DeleteTableThemeWithOptions(request *DeleteTableThemeRequest, runtime *dara.RuntimeOptions) (_result *DeleteTableThemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ThemeId) {
		query["ThemeId"] = request.ThemeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTableTheme"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTableThemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a table theme. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - DeleteTableThemeRequest
//
// @return DeleteTableThemeResponse
func (client *Client) DeleteTableTheme(request *DeleteTableThemeRequest) (_result *DeleteTableThemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTableThemeResponse{}
	_body, _err := client.DeleteTableThemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys a real-time synchronization task.
//
// @param request - DeployDISyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployDISyncTaskResponse
func (client *Client) DeployDISyncTaskWithOptions(request *DeployDISyncTaskRequest, runtime *dara.RuntimeOptions) (_result *DeployDISyncTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployDISyncTask"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployDISyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a real-time synchronization task.
//
// @param request - DeployDISyncTaskRequest
//
// @return DeployDISyncTaskResponse
func (client *Client) DeployDISyncTask(request *DeployDISyncTaskRequest) (_result *DeployDISyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployDISyncTaskResponse{}
	_body, _err := client.DeployDISyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys a file to the production environment.
//
// @param request - DeployFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployFileResponse
func (client *Client) DeployFileWithOptions(request *DeployFileRequest, runtime *dara.RuntimeOptions) (_result *DeployFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys a file to the production environment.
//
// @param request - DeployFileRequest
//
// @return DeployFileResponse
func (client *Client) DeployFile(request *DeployFileRequest) (_result *DeployFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeployFileResponse{}
	_body, _err := client.DeployFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Masks data.
//
// @param request - DesensitizeDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DesensitizeDataResponse
func (client *Client) DesensitizeDataWithOptions(request *DesensitizeDataRequest, runtime *dara.RuntimeOptions) (_result *DesensitizeDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DesensitizeData"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DesensitizeDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Masks data.
//
// @param request - DesensitizeDataRequest
//
// @return DesensitizeDataResponse
func (client *Client) DesensitizeData(request *DesensitizeDataRequest) (_result *DesensitizeDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DesensitizeDataResponse{}
	_body, _err := client.DesensitizeDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies a data masking rule.
//
// @param tmpReq - DsgDesensPlanAddOrUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgDesensPlanAddOrUpdateResponse
func (client *Client) DsgDesensPlanAddOrUpdateWithOptions(tmpReq *DsgDesensPlanAddOrUpdateRequest, runtime *dara.RuntimeOptions) (_result *DsgDesensPlanAddOrUpdateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgDesensPlanAddOrUpdateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.DesensRules) {
		request.DesensRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DesensRules, dara.String("DesensRules"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.DesensRulesShrink) {
		query["DesensRules"] = request.DesensRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgDesensPlanAddOrUpdate"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgDesensPlanAddOrUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies a data masking rule.
//
// @param request - DsgDesensPlanAddOrUpdateRequest
//
// @return DsgDesensPlanAddOrUpdateResponse
func (client *Client) DsgDesensPlanAddOrUpdate(request *DsgDesensPlanAddOrUpdateRequest) (_result *DsgDesensPlanAddOrUpdateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgDesensPlanAddOrUpdateResponse{}
	_body, _err := client.DsgDesensPlanAddOrUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data masking rule created in Data Security Guard.
//
// @param tmpReq - DsgDesensPlanDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgDesensPlanDeleteResponse
func (client *Client) DsgDesensPlanDeleteWithOptions(tmpReq *DsgDesensPlanDeleteRequest, runtime *dara.RuntimeOptions) (_result *DsgDesensPlanDeleteResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgDesensPlanDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgDesensPlanDelete"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgDesensPlanDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data masking rule created in Data Security Guard.
//
// @param request - DsgDesensPlanDeleteRequest
//
// @return DsgDesensPlanDeleteResponse
func (client *Client) DsgDesensPlanDelete(request *DsgDesensPlanDeleteRequest) (_result *DsgDesensPlanDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgDesensPlanDeleteResponse{}
	_body, _err := client.DsgDesensPlanDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data masking rules.
//
// @param request - DsgDesensPlanQueryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgDesensPlanQueryListResponse
func (client *Client) DsgDesensPlanQueryListWithOptions(request *DsgDesensPlanQueryListRequest, runtime *dara.RuntimeOptions) (_result *DsgDesensPlanQueryListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgDesensPlanQueryList"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgDesensPlanQueryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data masking rules.
//
// @param request - DsgDesensPlanQueryListRequest
//
// @return DsgDesensPlanQueryListResponse
func (client *Client) DsgDesensPlanQueryList(request *DsgDesensPlanQueryListRequest) (_result *DsgDesensPlanQueryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgDesensPlanQueryListResponse{}
	_body, _err := client.DsgDesensPlanQueryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the status of a data masking rule.
//
// @param tmpReq - DsgDesensPlanUpdateStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgDesensPlanUpdateStatusResponse
func (client *Client) DsgDesensPlanUpdateStatusWithOptions(tmpReq *DsgDesensPlanUpdateStatusRequest, runtime *dara.RuntimeOptions) (_result *DsgDesensPlanUpdateStatusResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgDesensPlanUpdateStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	if !dara.IsNil(request.SceneCode) {
		query["SceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgDesensPlanUpdateStatus"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgDesensPlanUpdateStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the status of a data masking rule.
//
// @param request - DsgDesensPlanUpdateStatusRequest
//
// @return DsgDesensPlanUpdateStatusResponse
func (client *Client) DsgDesensPlanUpdateStatus(request *DsgDesensPlanUpdateStatusRequest) (_result *DsgDesensPlanUpdateStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgDesensPlanUpdateStatusResponse{}
	_body, _err := client.DsgDesensPlanUpdateStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of compute engines of different types in the current tenant.
//
// @param request - DsgPlatformQueryProjectsAndSchemaFromMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgPlatformQueryProjectsAndSchemaFromMetaResponse
func (client *Client) DsgPlatformQueryProjectsAndSchemaFromMetaWithOptions(request *DsgPlatformQueryProjectsAndSchemaFromMetaRequest, runtime *dara.RuntimeOptions) (_result *DsgPlatformQueryProjectsAndSchemaFromMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgPlatformQueryProjectsAndSchemaFromMeta"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgPlatformQueryProjectsAndSchemaFromMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of compute engines of different types in the current tenant.
//
// @param request - DsgPlatformQueryProjectsAndSchemaFromMetaRequest
//
// @return DsgPlatformQueryProjectsAndSchemaFromMetaResponse
func (client *Client) DsgPlatformQueryProjectsAndSchemaFromMeta(request *DsgPlatformQueryProjectsAndSchemaFromMetaRequest) (_result *DsgPlatformQueryProjectsAndSchemaFromMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgPlatformQueryProjectsAndSchemaFromMetaResponse{}
	_body, _err := client.DsgPlatformQueryProjectsAndSchemaFromMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of available sensitive field type templates and the data masking rules supported by the templates. You can refer to the response parameters of this operation to configure a data masking rule.
//
// @param request - DsgQueryDefaultTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgQueryDefaultTemplatesResponse
func (client *Client) DsgQueryDefaultTemplatesWithOptions(request *DsgQueryDefaultTemplatesRequest, runtime *dara.RuntimeOptions) (_result *DsgQueryDefaultTemplatesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgQueryDefaultTemplates"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgQueryDefaultTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of available sensitive field type templates and the data masking rules supported by the templates. You can refer to the response parameters of this operation to configure a data masking rule.
//
// @param request - DsgQueryDefaultTemplatesRequest
//
// @return DsgQueryDefaultTemplatesResponse
func (client *Client) DsgQueryDefaultTemplates(request *DsgQueryDefaultTemplatesRequest) (_result *DsgQueryDefaultTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgQueryDefaultTemplatesResponse{}
	_body, _err := client.DsgQueryDefaultTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the identification results of sensitive data.
//
// Description:
//
// The query capability of the API operation is similar to the query feature in Data Security Guard in the DataWorks console. The API operation can be used to query the identification results of sensitive data of a tenant based on the association with the tenant ID.
//
//   - You can search for a specific identification result based on filter conditions such as data source type and workspace.
//
//   - You can sort the identification results of sensitive data of a tenant based on the values of a field in ascending or descending order.
//
//   - This operation supports paged query.
//
// @param request - DsgQuerySensResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgQuerySensResultResponse
func (client *Client) DsgQuerySensResultWithOptions(request *DsgQuerySensResultRequest, runtime *dara.RuntimeOptions) (_result *DsgQuerySensResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Col) {
		body["Col"] = request.Col
	}

	if !dara.IsNil(request.DbType) {
		body["DbType"] = request.DbType
	}

	if !dara.IsNil(request.Level) {
		body["Level"] = request.Level
	}

	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.Order) {
		body["Order"] = request.Order
	}

	if !dara.IsNil(request.OrderField) {
		body["OrderField"] = request.OrderField
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.SchemaName) {
		body["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.SensStatus) {
		body["SensStatus"] = request.SensStatus
	}

	if !dara.IsNil(request.SensitiveId) {
		body["SensitiveId"] = request.SensitiveId
	}

	if !dara.IsNil(request.SensitiveName) {
		body["SensitiveName"] = request.SensitiveName
	}

	if !dara.IsNil(request.Table) {
		body["Table"] = request.Table
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgQuerySensResult"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgQuerySensResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the identification results of sensitive data.
//
// Description:
//
// The query capability of the API operation is similar to the query feature in Data Security Guard in the DataWorks console. The API operation can be used to query the identification results of sensitive data of a tenant based on the association with the tenant ID.
//
//   - You can search for a specific identification result based on filter conditions such as data source type and workspace.
//
//   - You can sort the identification results of sensitive data of a tenant based on the values of a field in ascending or descending order.
//
//   - This operation supports paged query.
//
// @param request - DsgQuerySensResultRequest
//
// @return DsgQuerySensResultResponse
func (client *Client) DsgQuerySensResult(request *DsgQuerySensResultRequest) (_result *DsgQuerySensResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgQuerySensResultResponse{}
	_body, _err := client.DsgQuerySensResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a sensitive data identification task in Data Security Guard.
//
// @param tmpReq - DsgRunSensIdentifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgRunSensIdentifyResponse
func (client *Client) DsgRunSensIdentifyWithOptions(tmpReq *DsgRunSensIdentifyRequest, runtime *dara.RuntimeOptions) (_result *DsgRunSensIdentifyResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgRunSensIdentifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.EsMetaParams) {
		request.EsMetaParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EsMetaParams, dara.String("EsMetaParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.EsMetaParamsShrink) {
		body["EsMetaParams"] = request.EsMetaParamsShrink
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgRunSensIdentify"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgRunSensIdentifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a sensitive data identification task in Data Security Guard.
//
// @param request - DsgRunSensIdentifyRequest
//
// @return DsgRunSensIdentifyResponse
func (client *Client) DsgRunSensIdentify(request *DsgRunSensIdentifyRequest) (_result *DsgRunSensIdentifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgRunSensIdentifyResponse{}
	_body, _err := client.DsgRunSensIdentifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies a level-2 data masking scenario.
//
// @param tmpReq - DsgSceneAddOrUpdateSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgSceneAddOrUpdateSceneResponse
func (client *Client) DsgSceneAddOrUpdateSceneWithOptions(tmpReq *DsgSceneAddOrUpdateSceneRequest, runtime *dara.RuntimeOptions) (_result *DsgSceneAddOrUpdateSceneResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgSceneAddOrUpdateSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Scenes) {
		request.ScenesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scenes, dara.String("scenes"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ScenesShrink) {
		query["scenes"] = request.ScenesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgSceneAddOrUpdateScene"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgSceneAddOrUpdateSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies a level-2 data masking scenario.
//
// @param request - DsgSceneAddOrUpdateSceneRequest
//
// @return DsgSceneAddOrUpdateSceneResponse
func (client *Client) DsgSceneAddOrUpdateScene(request *DsgSceneAddOrUpdateSceneRequest) (_result *DsgSceneAddOrUpdateSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgSceneAddOrUpdateSceneResponse{}
	_body, _err := client.DsgSceneAddOrUpdateSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of data masking scenarios.
//
// @param request - DsgSceneQuerySceneListByNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgSceneQuerySceneListByNameResponse
func (client *Client) DsgSceneQuerySceneListByNameWithOptions(request *DsgSceneQuerySceneListByNameRequest, runtime *dara.RuntimeOptions) (_result *DsgSceneQuerySceneListByNameResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgSceneQuerySceneListByName"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgSceneQuerySceneListByNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of data masking scenarios.
//
// @param request - DsgSceneQuerySceneListByNameRequest
//
// @return DsgSceneQuerySceneListByNameResponse
func (client *Client) DsgSceneQuerySceneListByName(request *DsgSceneQuerySceneListByNameRequest) (_result *DsgSceneQuerySceneListByNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgSceneQuerySceneListByNameResponse{}
	_body, _err := client.DsgSceneQuerySceneListByNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a level-2 data masking scenario created in Data Security Guard.
//
// @param tmpReq - DsgScenedDeleteSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgScenedDeleteSceneResponse
func (client *Client) DsgScenedDeleteSceneWithOptions(tmpReq *DsgScenedDeleteSceneRequest, runtime *dara.RuntimeOptions) (_result *DsgScenedDeleteSceneResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgScenedDeleteSceneShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgScenedDeleteScene"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgScenedDeleteSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a level-2 data masking scenario created in Data Security Guard.
//
// @param request - DsgScenedDeleteSceneRequest
//
// @return DsgScenedDeleteSceneResponse
func (client *Client) DsgScenedDeleteScene(request *DsgScenedDeleteSceneRequest) (_result *DsgScenedDeleteSceneResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgScenedDeleteSceneResponse{}
	_body, _err := client.DsgScenedDeleteSceneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a sensitive data identification task.
//
// @param request - DsgStopSensIdentifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgStopSensIdentifyResponse
func (client *Client) DsgStopSensIdentifyWithOptions(request *DsgStopSensIdentifyRequest, runtime *dara.RuntimeOptions) (_result *DsgStopSensIdentifyResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["JobId"] = request.JobId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgStopSensIdentify"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgStopSensIdentifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a sensitive data identification task.
//
// @param request - DsgStopSensIdentifyRequest
//
// @return DsgStopSensIdentifyResponse
func (client *Client) DsgStopSensIdentify(request *DsgStopSensIdentifyRequest) (_result *DsgStopSensIdentifyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgStopSensIdentifyResponse{}
	_body, _err := client.DsgStopSensIdentifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies a user group.
//
// @param tmpReq - DsgUserGroupAddOrUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgUserGroupAddOrUpdateResponse
func (client *Client) DsgUserGroupAddOrUpdateWithOptions(tmpReq *DsgUserGroupAddOrUpdateRequest, runtime *dara.RuntimeOptions) (_result *DsgUserGroupAddOrUpdateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgUserGroupAddOrUpdateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.UserGroups) {
		request.UserGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroups, dara.String("UserGroups"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.UserGroupsShrink) {
		query["UserGroups"] = request.UserGroupsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgUserGroupAddOrUpdate"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgUserGroupAddOrUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies a user group.
//
// @param request - DsgUserGroupAddOrUpdateRequest
//
// @return DsgUserGroupAddOrUpdateResponse
func (client *Client) DsgUserGroupAddOrUpdate(request *DsgUserGroupAddOrUpdateRequest) (_result *DsgUserGroupAddOrUpdateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgUserGroupAddOrUpdateResponse{}
	_body, _err := client.DsgUserGroupAddOrUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user group configured in Data Security Guard.
//
// @param tmpReq - DsgUserGroupDeleteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgUserGroupDeleteResponse
func (client *Client) DsgUserGroupDeleteWithOptions(tmpReq *DsgUserGroupDeleteRequest, runtime *dara.RuntimeOptions) (_result *DsgUserGroupDeleteResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgUserGroupDeleteShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgUserGroupDelete"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgUserGroupDeleteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user group configured in Data Security Guard.
//
// @param request - DsgUserGroupDeleteRequest
//
// @return DsgUserGroupDeleteResponse
func (client *Client) DsgUserGroupDelete(request *DsgUserGroupDeleteRequest) (_result *DsgUserGroupDeleteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgUserGroupDeleteResponse{}
	_body, _err := client.DsgUserGroupDeleteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of MaxCompute roles that can be selected by the members of a user group when the user group is created or modified by the tenant in Data Security Guard.
//
// @param request - DsgUserGroupGetOdpsRoleGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgUserGroupGetOdpsRoleGroupsResponse
func (client *Client) DsgUserGroupGetOdpsRoleGroupsWithOptions(request *DsgUserGroupGetOdpsRoleGroupsRequest, runtime *dara.RuntimeOptions) (_result *DsgUserGroupGetOdpsRoleGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgUserGroupGetOdpsRoleGroups"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgUserGroupGetOdpsRoleGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of MaxCompute roles that can be selected by the members of a user group when the user group is created or modified by the tenant in Data Security Guard.
//
// @param request - DsgUserGroupGetOdpsRoleGroupsRequest
//
// @return DsgUserGroupGetOdpsRoleGroupsResponse
func (client *Client) DsgUserGroupGetOdpsRoleGroups(request *DsgUserGroupGetOdpsRoleGroupsRequest) (_result *DsgUserGroupGetOdpsRoleGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgUserGroupGetOdpsRoleGroupsResponse{}
	_body, _err := client.DsgUserGroupGetOdpsRoleGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of user groups in Data Security Guard.
//
// @param request - DsgUserGroupQueryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgUserGroupQueryListResponse
func (client *Client) DsgUserGroupQueryListWithOptions(request *DsgUserGroupQueryListRequest, runtime *dara.RuntimeOptions) (_result *DsgUserGroupQueryListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgUserGroupQueryList"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgUserGroupQueryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of user groups in Data Security Guard.
//
// @param request - DsgUserGroupQueryListRequest
//
// @return DsgUserGroupQueryListResponse
func (client *Client) DsgUserGroupQueryList(request *DsgUserGroupQueryListRequest) (_result *DsgUserGroupQueryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgUserGroupQueryListResponse{}
	_body, _err := client.DsgUserGroupQueryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of users or roles of the current tenant.
//
// @param request - DsgUserGroupQueryUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgUserGroupQueryUserListResponse
func (client *Client) DsgUserGroupQueryUserListWithOptions(runtime *dara.RuntimeOptions) (_result *DsgUserGroupQueryUserListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("DsgUserGroupQueryUserList"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgUserGroupQueryUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of users or roles of the current tenant.
//
// @return DsgUserGroupQueryUserListResponse
func (client *Client) DsgUserGroupQueryUserList() (_result *DsgUserGroupQueryUserListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgUserGroupQueryUserListResponse{}
	_body, _err := client.DsgUserGroupQueryUserListWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds or modifies a data masking whitelist.
//
// @param tmpReq - DsgWhiteListAddOrUpdateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgWhiteListAddOrUpdateResponse
func (client *Client) DsgWhiteListAddOrUpdateWithOptions(tmpReq *DsgWhiteListAddOrUpdateRequest, runtime *dara.RuntimeOptions) (_result *DsgWhiteListAddOrUpdateResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgWhiteListAddOrUpdateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.WhiteLists) {
		request.WhiteListsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.WhiteLists, dara.String("WhiteLists"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.WhiteListsShrink) {
		query["WhiteLists"] = request.WhiteListsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgWhiteListAddOrUpdate"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgWhiteListAddOrUpdateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds or modifies a data masking whitelist.
//
// @param request - DsgWhiteListAddOrUpdateRequest
//
// @return DsgWhiteListAddOrUpdateResponse
func (client *Client) DsgWhiteListAddOrUpdate(request *DsgWhiteListAddOrUpdateRequest) (_result *DsgWhiteListAddOrUpdateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgWhiteListAddOrUpdateResponse{}
	_body, _err := client.DsgWhiteListAddOrUpdateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a data masking whitelist configured in Data Security Guard.
//
// @param tmpReq - DsgWhiteListDeleteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgWhiteListDeleteListResponse
func (client *Client) DsgWhiteListDeleteListWithOptions(tmpReq *DsgWhiteListDeleteListRequest, runtime *dara.RuntimeOptions) (_result *DsgWhiteListDeleteListResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &DsgWhiteListDeleteListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Ids) {
		request.IdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Ids, dara.String("Ids"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdsShrink) {
		query["Ids"] = request.IdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgWhiteListDeleteList"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgWhiteListDeleteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data masking whitelist configured in Data Security Guard.
//
// @param request - DsgWhiteListDeleteListRequest
//
// @return DsgWhiteListDeleteListResponse
func (client *Client) DsgWhiteListDeleteList(request *DsgWhiteListDeleteListRequest) (_result *DsgWhiteListDeleteListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgWhiteListDeleteListResponse{}
	_body, _err := client.DsgWhiteListDeleteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a data masking whitelist.
//
// @param request - DsgWhiteListQueryListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DsgWhiteListQueryListResponse
func (client *Client) DsgWhiteListQueryListWithOptions(request *DsgWhiteListQueryListRequest, runtime *dara.RuntimeOptions) (_result *DsgWhiteListQueryListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DsgWhiteListQueryList"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DsgWhiteListQueryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a data masking whitelist.
//
// @param request - DsgWhiteListQueryListRequest
//
// @return DsgWhiteListQueryListResponse
func (client *Client) DsgWhiteListQueryList(request *DsgWhiteListQueryListRequest) (_result *DsgWhiteListQueryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DsgWhiteListQueryListResponse{}
	_body, _err := client.DsgWhiteListQueryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Edits a sensitive field that is defined based on the category and sensitivity level of data in Data Security Guard.
//
// @param request - EditRecognizeRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditRecognizeRuleResponse
func (client *Client) EditRecognizeRuleWithOptions(request *EditRecognizeRuleRequest, runtime *dara.RuntimeOptions) (_result *EditRecognizeRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AccountName) {
		body["AccountName"] = request.AccountName
	}

	if !dara.IsNil(request.ColExclude) {
		body["ColExclude"] = request.ColExclude
	}

	if !dara.IsNil(request.ColScan) {
		body["ColScan"] = request.ColScan
	}

	if !dara.IsNil(request.CommentScan) {
		body["CommentScan"] = request.CommentScan
	}

	if !dara.IsNil(request.ContentScan) {
		body["ContentScan"] = request.ContentScan
	}

	if !dara.IsNil(request.HitThreshold) {
		body["HitThreshold"] = request.HitThreshold
	}

	if !dara.IsNil(request.LevelName) {
		body["LevelName"] = request.LevelName
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeParent) {
		body["NodeParent"] = request.NodeParent
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.RecognizeRules) {
		body["RecognizeRules"] = request.RecognizeRules
	}

	if !dara.IsNil(request.RecognizeRulesType) {
		body["RecognizeRulesType"] = request.RecognizeRulesType
	}

	if !dara.IsNil(request.SensitiveDescription) {
		body["SensitiveDescription"] = request.SensitiveDescription
	}

	if !dara.IsNil(request.SensitiveId) {
		body["SensitiveId"] = request.SensitiveId
	}

	if !dara.IsNil(request.SensitiveName) {
		body["SensitiveName"] = request.SensitiveName
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Level) {
		body["level"] = request.Level
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditRecognizeRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditRecognizeRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Edits a sensitive field that is defined based on the category and sensitivity level of data in Data Security Guard.
//
// @param request - EditRecognizeRuleRequest
//
// @return EditRecognizeRuleResponse
func (client *Client) EditRecognizeRule(request *EditRecognizeRuleRequest) (_result *EditRecognizeRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditRecognizeRuleResponse{}
	_body, _err := client.EditRecognizeRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EstablishRelationTableToBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EstablishRelationTableToBusinessResponse
func (client *Client) EstablishRelationTableToBusinessWithOptions(request *EstablishRelationTableToBusinessRequest, runtime *dara.RuntimeOptions) (_result *EstablishRelationTableToBusinessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.TableGuid) {
		body["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EstablishRelationTableToBusiness"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EstablishRelationTableToBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - EstablishRelationTableToBusinessRequest
//
// @return EstablishRelationTableToBusinessResponse
func (client *Client) EstablishRelationTableToBusiness(request *EstablishRelationTableToBusinessRequest) (_result *EstablishRelationTableToBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EstablishRelationTableToBusinessResponse{}
	_body, _err := client.EstablishRelationTableToBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Exports a list of data sources.
//
// @param request - ExportDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportDataSourcesResponse
func (client *Client) ExportDataSourcesWithOptions(request *ExportDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ExportDataSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExportDataSources"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExportDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Exports a list of data sources.
//
// @param request - ExportDataSourcesRequest
//
// @return ExportDataSourcesResponse
func (client *Client) ExportDataSources(request *ExportDataSourcesRequest) (_result *ExportDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExportDataSourcesResponse{}
	_body, _err := client.ExportDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an ID for an asynchronous thread that is used to create a synchronization task in Data Integration.
//
// Description:
//
// DataWorks allows you to use the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) operation to directly create a batch synchronization task in Data Integration. To create a real-time synchronization task or another type of synchronization task, you must first call the [GenerateDISyncTaskConfigForCreating](https://help.aliyun.com/document_detail/383463.html) operation to generate the ID of an asynchronous thread and call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. Then, you can use the parameters as request parameters of [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) and call the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) operation to create a real-time synchronization task or another type of synchronization task. DataWorks allows you to create real-time synchronization tasks and other types of synchronization tasks in Data Integration only in asynchronous mode.
//
// @param request - GenerateDISyncTaskConfigForCreatingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDISyncTaskConfigForCreatingResponse
func (client *Client) GenerateDISyncTaskConfigForCreatingWithOptions(request *GenerateDISyncTaskConfigForCreatingRequest, runtime *dara.RuntimeOptions) (_result *GenerateDISyncTaskConfigForCreatingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskParam) {
		query["TaskParam"] = request.TaskParam
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateDISyncTaskConfigForCreating"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDISyncTaskConfigForCreatingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an ID for an asynchronous thread that is used to create a synchronization task in Data Integration.
//
// Description:
//
// DataWorks allows you to use the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) operation to directly create a batch synchronization task in Data Integration. To create a real-time synchronization task or another type of synchronization task, you must first call the [GenerateDISyncTaskConfigForCreating](https://help.aliyun.com/document_detail/383463.html) operation to generate the ID of an asynchronous thread and call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. Then, you can use the parameters as request parameters of [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) and call the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) operation to create a real-time synchronization task or another type of synchronization task. DataWorks allows you to create real-time synchronization tasks and other types of synchronization tasks in Data Integration only in asynchronous mode.
//
// @param request - GenerateDISyncTaskConfigForCreatingRequest
//
// @return GenerateDISyncTaskConfigForCreatingResponse
func (client *Client) GenerateDISyncTaskConfigForCreating(request *GenerateDISyncTaskConfigForCreatingRequest) (_result *GenerateDISyncTaskConfigForCreatingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateDISyncTaskConfigForCreatingResponse{}
	_body, _err := client.GenerateDISyncTaskConfigForCreatingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the JSON for an asynchronous thread that is used to update a real-time synchronization task in Data Integration.
//
// Description:
//
// DataWorks allows you to use only the [UpdateDISyncTask](https://help.aliyun.com/document_detail/289109.html) operation to update a batch synchronization task in Data Integration. To update a real-time synchronization task, you must first call the GenerateDISyncTaskConfigForUpdating operation to generate the ID of an asynchronous thread and call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. Then, you can call the UpdateDISyncTask operation and use the parameters as request parameters to update a real-time synchronization task in Data Integration. DataWorks allows you to create or update real-time synchronization tasks in Data Integration only in asynchronous mode.
//
// @param request - GenerateDISyncTaskConfigForUpdatingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateDISyncTaskConfigForUpdatingResponse
func (client *Client) GenerateDISyncTaskConfigForUpdatingWithOptions(request *GenerateDISyncTaskConfigForUpdatingRequest, runtime *dara.RuntimeOptions) (_result *GenerateDISyncTaskConfigForUpdatingResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskParam) {
		query["TaskParam"] = request.TaskParam
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateDISyncTaskConfigForUpdating"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateDISyncTaskConfigForUpdatingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the JSON for an asynchronous thread that is used to update a real-time synchronization task in Data Integration.
//
// Description:
//
// DataWorks allows you to use only the [UpdateDISyncTask](https://help.aliyun.com/document_detail/289109.html) operation to update a batch synchronization task in Data Integration. To update a real-time synchronization task, you must first call the GenerateDISyncTaskConfigForUpdating operation to generate the ID of an asynchronous thread and call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. Then, you can call the UpdateDISyncTask operation and use the parameters as request parameters to update a real-time synchronization task in Data Integration. DataWorks allows you to create or update real-time synchronization tasks in Data Integration only in asynchronous mode.
//
// @param request - GenerateDISyncTaskConfigForUpdatingRequest
//
// @return GenerateDISyncTaskConfigForUpdatingResponse
func (client *Client) GenerateDISyncTaskConfigForUpdating(request *GenerateDISyncTaskConfigForUpdatingRequest) (_result *GenerateDISyncTaskConfigForUpdatingResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateDISyncTaskConfigForUpdatingResponse{}
	_body, _err := client.GenerateDISyncTaskConfigForUpdatingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries alert information based on the alert ID that is specified by the AlertId parameter.
//
// @param request - GetAlertMessageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAlertMessageResponse
func (client *Client) GetAlertMessageWithOptions(request *GetAlertMessageRequest, runtime *dara.RuntimeOptions) (_result *GetAlertMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertId) {
		body["AlertId"] = request.AlertId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAlertMessage"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAlertMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries alert information based on the alert ID that is specified by the AlertId parameter.
//
// @param request - GetAlertMessageRequest
//
// @return GetAlertMessageResponse
func (client *Client) GetAlertMessage(request *GetAlertMessageRequest) (_result *GetAlertMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAlertMessageResponse{}
	_body, _err := client.GetAlertMessageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a baseline based on its ID.
//
// @param request - GetBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineResponse
func (client *Client) GetBaselineWithOptions(request *GetBaselineRequest, runtime *dara.RuntimeOptions) (_result *GetBaselineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBaseline"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a baseline based on its ID.
//
// @param request - GetBaselineRequest
//
// @return GetBaselineResponse
func (client *Client) GetBaseline(request *GetBaselineRequest) (_result *GetBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBaselineResponse{}
	_body, _err := client.GetBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of a baseline.
//
// @param request - GetBaselineConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineConfigResponse
func (client *Client) GetBaselineConfigWithOptions(request *GetBaselineConfigRequest, runtime *dara.RuntimeOptions) (_result *GetBaselineConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBaselineConfig"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBaselineConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of a baseline.
//
// @param request - GetBaselineConfigRequest
//
// @return GetBaselineConfigResponse
func (client *Client) GetBaselineConfig(request *GetBaselineConfigRequest) (_result *GetBaselineConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBaselineConfigResponse{}
	_body, _err := client.GetBaselineConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The information about the events that are associated with the instance.
//
// @param request - GetBaselineKeyPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineKeyPathResponse
func (client *Client) GetBaselineKeyPathWithOptions(request *GetBaselineKeyPathRequest, runtime *dara.RuntimeOptions) (_result *GetBaselineKeyPathResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.InGroupId) {
		body["InGroupId"] = request.InGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBaselineKeyPath"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBaselineKeyPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The information about the events that are associated with the instance.
//
// @param request - GetBaselineKeyPathRequest
//
// @return GetBaselineKeyPathResponse
func (client *Client) GetBaselineKeyPath(request *GetBaselineKeyPathRequest) (_result *GetBaselineKeyPathResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBaselineKeyPathResponse{}
	_body, _err := client.GetBaselineKeyPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a baseline instance.
//
// @param request - GetBaselineStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBaselineStatusResponse
func (client *Client) GetBaselineStatusWithOptions(request *GetBaselineStatusRequest, runtime *dara.RuntimeOptions) (_result *GetBaselineStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.InGroupId) {
		body["InGroupId"] = request.InGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBaselineStatus"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBaselineStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a baseline instance.
//
// @param request - GetBaselineStatusRequest
//
// @return GetBaselineStatusResponse
func (client *Client) GetBaselineStatus(request *GetBaselineStatusRequest) (_result *GetBaselineStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBaselineStatusResponse{}
	_body, _err := client.GetBaselineStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// @param request - GetBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBusinessResponse
func (client *Client) GetBusinessWithOptions(request *GetBusinessRequest, runtime *dara.RuntimeOptions) (_result *GetBusinessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBusiness"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a workflow.
//
// @param request - GetBusinessRequest
//
// @return GetBusinessResponse
func (client *Client) GetBusiness(request *GetBusinessRequest) (_result *GetBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBusinessResponse{}
	_body, _err := client.GetBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a table creation, update, or deletion task.
//
// @param request - GetDDLJobStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDDLJobStatusResponse
func (client *Client) GetDDLJobStatusWithOptions(request *GetDDLJobStatusRequest, runtime *dara.RuntimeOptions) (_result *GetDDLJobStatusResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDDLJobStatus"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDDLJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a table creation, update, or deletion task.
//
// @param request - GetDDLJobStatusRequest
//
// @return GetDDLJobStatusResponse
func (client *Client) GetDDLJobStatus(request *GetDDLJobStatusRequest) (_result *GetDDLJobStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDDLJobStatusResponse{}
	_body, _err := client.GetDDLJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an alert rule configured for a new-version synchronization task. Only the following types of tasks are supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization.
//
// @param request - GetDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIAlarmRuleResponse
func (client *Client) GetDIAlarmRuleWithOptions(request *GetDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *GetDIAlarmRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DIAlarmRuleId) {
		body["DIAlarmRuleId"] = request.DIAlarmRuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDIAlarmRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an alert rule configured for a new-version synchronization task. Only the following types of tasks are supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization.
//
// @param request - GetDIAlarmRuleRequest
//
// @return GetDIAlarmRuleResponse
func (client *Client) GetDIAlarmRule(request *GetDIAlarmRuleRequest) (_result *GetDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDIAlarmRuleResponse{}
	_body, _err := client.GetDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a new-version synchronization task created in Data Integration. The following types of synchronization tasks are supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// @param request - GetDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDIJobResponse
func (client *Client) GetDIJobWithOptions(request *GetDIJobRequest, runtime *dara.RuntimeOptions) (_result *GetDIJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.WithDetails) {
		body["WithDetails"] = request.WithDetails
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDIJob"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a new-version synchronization task created in Data Integration. The following types of synchronization tasks are supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// @param request - GetDIJobRequest
//
// @return GetDIJobResponse
func (client *Client) GetDIJob(request *GetDIJobRequest) (_result *GetDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDIJobResponse{}
	_body, _err := client.GetDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of a real-time synchronization task or a data synchronization solution.
//
// @param request - GetDISyncInstanceInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDISyncInstanceInfoResponse
func (client *Client) GetDISyncInstanceInfoWithOptions(request *GetDISyncInstanceInfoRequest, runtime *dara.RuntimeOptions) (_result *GetDISyncInstanceInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDISyncInstanceInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDISyncInstanceInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of a real-time synchronization task or a data synchronization solution.
//
// @param request - GetDISyncInstanceInfoRequest
//
// @return GetDISyncInstanceInfoResponse
func (client *Client) GetDISyncInstanceInfo(request *GetDISyncInstanceInfoRequest) (_result *GetDISyncInstanceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDISyncInstanceInfoResponse{}
	_body, _err := client.GetDISyncInstanceInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a real-time synchronization task or a data synchronization solution.
//
// @param request - GetDISyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDISyncTaskResponse
func (client *Client) GetDISyncTaskWithOptions(request *GetDISyncTaskRequest, runtime *dara.RuntimeOptions) (_result *GetDISyncTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDISyncTask"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDISyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a real-time synchronization task or a data synchronization solution.
//
// @param request - GetDISyncTaskRequest
//
// @return GetDISyncTaskResponse
func (client *Client) GetDISyncTask(request *GetDISyncTaskRequest) (_result *GetDISyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDISyncTaskResponse{}
	_body, _err := client.GetDISyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a directed acyclic graph (DAG). You can call the GetDag operation to query the information about the DAG for a manually triggered workflow, a manually triggered node, or a data backfill instance. However, you cannot query the information about the DAG for an auto triggered node or an auto triggered workflow.
//
// Description:
//
// Supported DAG types:
//
//   - MANUAL: DAG for a manually triggered workflow
//
//   - SMOKE_TEST: DAG for a smoke testing workflow
//
//   - SUPPLY_DATA: DAG for a data backfill instance
//
//   - BUSINESS_PROCESS_DAG: DAG for a one-time workflow
//
// Supported DAG states:
//
//   - CREATED
//
//   - RUNNING
//
//   - FAILURE
//
//   - SUCCESS
//
// @param request - GetDagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDagResponse
func (client *Client) GetDagWithOptions(request *GetDagRequest, runtime *dara.RuntimeOptions) (_result *GetDagResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		body["DagId"] = request.DagId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDag"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a directed acyclic graph (DAG). You can call the GetDag operation to query the information about the DAG for a manually triggered workflow, a manually triggered node, or a data backfill instance. However, you cannot query the information about the DAG for an auto triggered node or an auto triggered workflow.
//
// Description:
//
// Supported DAG types:
//
//   - MANUAL: DAG for a manually triggered workflow
//
//   - SMOKE_TEST: DAG for a smoke testing workflow
//
//   - SUPPLY_DATA: DAG for a data backfill instance
//
//   - BUSINESS_PROCESS_DAG: DAG for a one-time workflow
//
// Supported DAG states:
//
//   - CREATED
//
//   - RUNNING
//
//   - FAILURE
//
//   - SUCCESS
//
// @param request - GetDagRequest
//
// @return GetDagResponse
func (client *Client) GetDag(request *GetDagRequest) (_result *GetDagResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDagResponse{}
	_body, _err := client.GetDagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a DataService Studio API in the development state.
//
// @param request - GetDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiResponse
func (client *Client) GetDataServiceApiWithOptions(request *GetDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a DataService Studio API in the development state.
//
// @param request - GetDataServiceApiRequest
//
// @return GetDataServiceApiResponse
func (client *Client) GetDataServiceApi(request *GetDataServiceApiRequest) (_result *GetDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiResponse{}
	_body, _err := client.GetDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the test results of an API in DataService Studio.
//
// @param request - GetDataServiceApiTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApiTestResponse
func (client *Client) GetDataServiceApiTestWithOptions(request *GetDataServiceApiTestRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApiTestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApiTest"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApiTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the test results of an API in DataService Studio.
//
// @param request - GetDataServiceApiTestRequest
//
// @return GetDataServiceApiTestResponse
func (client *Client) GetDataServiceApiTest(request *GetDataServiceApiTestRequest) (_result *GetDataServiceApiTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApiTestResponse{}
	_body, _err := client.GetDataServiceApiTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an application.
//
// @param request - GetDataServiceApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceApplicationResponse
func (client *Client) GetDataServiceApplicationWithOptions(request *GetDataServiceApplicationRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceApplicationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApplicationId) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceApplication"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an application.
//
// @param request - GetDataServiceApplicationRequest
//
// @return GetDataServiceApplicationResponse
func (client *Client) GetDataServiceApplication(request *GetDataServiceApplicationRequest) (_result *GetDataServiceApplicationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceApplicationResponse{}
	_body, _err := client.GetDataServiceApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a folder.
//
// @param request - GetDataServiceFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceFolderResponse
func (client *Client) GetDataServiceFolderWithOptions(request *GetDataServiceFolderRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceFolder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a folder.
//
// @param request - GetDataServiceFolderRequest
//
// @return GetDataServiceFolderResponse
func (client *Client) GetDataServiceFolder(request *GetDataServiceFolderRequest) (_result *GetDataServiceFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceFolderResponse{}
	_body, _err := client.GetDataServiceFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a business process.
//
// @param request - GetDataServiceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServiceGroupResponse
func (client *Client) GetDataServiceGroupWithOptions(request *GetDataServiceGroupRequest, runtime *dara.RuntimeOptions) (_result *GetDataServiceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServiceGroup"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServiceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a business process.
//
// @param request - GetDataServiceGroupRequest
//
// @return GetDataServiceGroupResponse
func (client *Client) GetDataServiceGroup(request *GetDataServiceGroupRequest) (_result *GetDataServiceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServiceGroupResponse{}
	_body, _err := client.GetDataServiceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a DataService Studio API in the published state.
//
// @param request - GetDataServicePublishedApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataServicePublishedApiResponse
func (client *Client) GetDataServicePublishedApiWithOptions(request *GetDataServicePublishedApiRequest, runtime *dara.RuntimeOptions) (_result *GetDataServicePublishedApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataServicePublishedApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataServicePublishedApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a DataService Studio API in the published state.
//
// @param request - GetDataServicePublishedApiRequest
//
// @return GetDataServicePublishedApiResponse
func (client *Client) GetDataServicePublishedApi(request *GetDataServicePublishedApiRequest) (_result *GetDataServicePublishedApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataServicePublishedApiResponse{}
	_body, _err := client.GetDataServicePublishedApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specified data source.
//
// @param request - GetDataSourceMetaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataSourceMetaResponse
func (client *Client) GetDataSourceMetaWithOptions(request *GetDataSourceMetaRequest, runtime *dara.RuntimeOptions) (_result *GetDataSourceMetaResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasourceName) {
		query["DatasourceName"] = request.DatasourceName
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataSourceMeta"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataSourceMetaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the metadata of a specified data source.
//
// @param request - GetDataSourceMetaRequest
//
// @return GetDataSourceMetaResponse
func (client *Client) GetDataSourceMeta(request *GetDataSourceMetaRequest) (_result *GetDataSourceMetaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDataSourceMetaResponse{}
	_body, _err := client.GetDataSourceMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a deployment package.
//
// @param request - GetDeploymentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeploymentResponse
func (client *Client) GetDeploymentWithOptions(request *GetDeploymentRequest, runtime *dara.RuntimeOptions) (_result *GetDeploymentResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeploymentId) {
		body["DeploymentId"] = request.DeploymentId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeployment"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeploymentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a deployment package.
//
// @param request - GetDeploymentRequest
//
// @return GetDeploymentResponse
func (client *Client) GetDeployment(request *GetDeploymentRequest) (_result *GetDeploymentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDeploymentResponse{}
	_body, _err := client.GetDeploymentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an extension.
//
// @param request - GetExtensionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExtensionResponse
func (client *Client) GetExtensionWithOptions(request *GetExtensionRequest, runtime *dara.RuntimeOptions) (_result *GetExtensionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtensionCode) {
		query["ExtensionCode"] = request.ExtensionCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExtension"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExtensionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an extension.
//
// @param request - GetExtensionRequest
//
// @return GetExtensionResponse
func (client *Client) GetExtension(request *GetExtensionRequest) (_result *GetExtensionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetExtensionResponse{}
	_body, _err := client.GetExtensionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithOptions(request *GetFileRequest, runtime *dara.RuntimeOptions) (_result *GetFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @return GetFileResponse
func (client *Client) GetFile(request *GetFileRequest) (_result *GetFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the distribution of node types.
//
// @param request - GetFileTypeStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileTypeStatisticResponse
func (client *Client) GetFileTypeStatisticWithOptions(request *GetFileTypeStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetFileTypeStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileTypeStatistic"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileTypeStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the distribution of node types.
//
// @param request - GetFileTypeStatisticRequest
//
// @return GetFileTypeStatisticResponse
func (client *Client) GetFileTypeStatistic(request *GetFileTypeStatisticRequest) (_result *GetFileTypeStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileTypeStatisticResponse{}
	_body, _err := client.GetFileTypeStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file version.
//
// @param request - GetFileVersionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileVersionResponse
func (client *Client) GetFileVersionWithOptions(request *GetFileVersionRequest, runtime *dara.RuntimeOptions) (_result *GetFileVersionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileVersion) {
		body["FileVersion"] = request.FileVersion
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFileVersion"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFileVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file version.
//
// @param request - GetFileVersionRequest
//
// @return GetFileVersionResponse
func (client *Client) GetFileVersion(request *GetFileVersionRequest) (_result *GetFileVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFileVersionResponse{}
	_body, _err := client.GetFileVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a folder.
//
// @param request - GetFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFolderResponse
func (client *Client) GetFolderWithOptions(request *GetFolderRequest, runtime *dara.RuntimeOptions) (_result *GetFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.FolderPath) {
		body["FolderPath"] = request.FolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFolder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a folder.
//
// @param request - GetFolderRequest
//
// @return GetFolderResponse
func (client *Client) GetFolder(request *GetFolderRequest) (_result *GetFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFolderResponse{}
	_body, _err := client.GetFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data snapshot of an extension point based on the ID of a message in DataWorks OpenEvent when the related extension point event is triggered.
//
// @param request - GetIDEEventDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIDEEventDetailResponse
func (client *Client) GetIDEEventDetailWithOptions(request *GetIDEEventDetailRequest, runtime *dara.RuntimeOptions) (_result *GetIDEEventDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIDEEventDetail"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIDEEventDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the data snapshot of an extension point based on the ID of a message in DataWorks OpenEvent when the related extension point event is triggered.
//
// @param request - GetIDEEventDetailRequest
//
// @return GetIDEEventDetailResponse
func (client *Client) GetIDEEventDetail(request *GetIDEEventDetailRequest) (_result *GetIDEEventDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetIDEEventDetailResponse{}
	_body, _err := client.GetIDEEventDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an instance.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstance"),
		Version:     dara.String("2020-05-18"),
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
// Queries the information about an instance.
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

// Deprecated: OpenAPI GetInstanceConsumeTimeRank is deprecated
//
// Summary:
//
// Queries the ranking of the running durations of instances.
//
// @param request - GetInstanceConsumeTimeRankRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceConsumeTimeRankResponse
func (client *Client) GetInstanceConsumeTimeRankWithOptions(request *GetInstanceConsumeTimeRankRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceConsumeTimeRankResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceConsumeTimeRank"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceConsumeTimeRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceConsumeTimeRank is deprecated
//
// Summary:
//
// Queries the ranking of the running durations of instances.
//
// @param request - GetInstanceConsumeTimeRankRequest
//
// @return GetInstanceConsumeTimeRankResponse
// Deprecated
func (client *Client) GetInstanceConsumeTimeRank(request *GetInstanceConsumeTimeRankRequest) (_result *GetInstanceConsumeTimeRankResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceConsumeTimeRankResponse{}
	_body, _err := client.GetInstanceConsumeTimeRankWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceCountTrend is deprecated
//
// Summary:
//
// Queries the quantity trend of auto triggered instances.
//
// @param request - GetInstanceCountTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceCountTrendResponse
func (client *Client) GetInstanceCountTrendWithOptions(request *GetInstanceCountTrendRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceCountTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginDate) {
		body["BeginDate"] = request.BeginDate
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceCountTrend"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceCountTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceCountTrend is deprecated
//
// Summary:
//
// Queries the quantity trend of auto triggered instances.
//
// @param request - GetInstanceCountTrendRequest
//
// @return GetInstanceCountTrendResponse
// Deprecated
func (client *Client) GetInstanceCountTrend(request *GetInstanceCountTrendRequest) (_result *GetInstanceCountTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceCountTrendResponse{}
	_body, _err := client.GetInstanceCountTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceErrorRank is deprecated
//
// Summary:
//
// Queries the ranking of nodes on which errors occur within the last month.
//
// @param request - GetInstanceErrorRankRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceErrorRankResponse
func (client *Client) GetInstanceErrorRankWithOptions(request *GetInstanceErrorRankRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceErrorRankResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceErrorRank"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceErrorRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceErrorRank is deprecated
//
// Summary:
//
// Queries the ranking of nodes on which errors occur within the last month.
//
// @param request - GetInstanceErrorRankRequest
//
// @return GetInstanceErrorRankResponse
// Deprecated
func (client *Client) GetInstanceErrorRank(request *GetInstanceErrorRankRequest) (_result *GetInstanceErrorRankResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceErrorRankResponse{}
	_body, _err := client.GetInstanceErrorRankWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of an instance.
//
// Description:
//
// You may not obtain the instance logs that were generated more than seven days ago.
//
// @param request - GetInstanceLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceLogResponse
func (client *Client) GetInstanceLogWithOptions(request *GetInstanceLogRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceHistoryId) {
		body["InstanceHistoryId"] = request.InstanceHistoryId
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceLog"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of an instance.
//
// Description:
//
// You may not obtain the instance logs that were generated more than seven days ago.
//
// @param request - GetInstanceLogRequest
//
// @return GetInstanceLogResponse
func (client *Client) GetInstanceLog(request *GetInstanceLogRequest) (_result *GetInstanceLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceLogResponse{}
	_body, _err := client.GetInstanceLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceStatusCount is deprecated
//
// Summary:
//
// Queries the statistics of instances in different states.
//
// @param request - GetInstanceStatusCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceStatusCountResponse
func (client *Client) GetInstanceStatusCountWithOptions(request *GetInstanceStatusCountRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceStatusCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceStatusCount"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceStatusCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetInstanceStatusCount is deprecated
//
// Summary:
//
// Queries the statistics of instances in different states.
//
// @param request - GetInstanceStatusCountRequest
//
// @return GetInstanceStatusCountResponse
// Deprecated
func (client *Client) GetInstanceStatusCount(request *GetInstanceStatusCountRequest) (_result *GetInstanceStatusCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceStatusCountResponse{}
	_body, _err := client.GetInstanceStatusCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of instances that are in each state.
//
// @param request - GetInstanceStatusStatisticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceStatusStatisticResponse
func (client *Client) GetInstanceStatusStatisticWithOptions(request *GetInstanceStatusStatisticRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceStatusStatisticResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.DagType) {
		body["DagType"] = request.DagType
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SchedulerPeriod) {
		body["SchedulerPeriod"] = request.SchedulerPeriod
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceStatusStatistic"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceStatusStatisticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of instances that are in each state.
//
// @param request - GetInstanceStatusStatisticRequest
//
// @return GetInstanceStatusStatisticResponse
func (client *Client) GetInstanceStatusStatistic(request *GetInstanceStatusStatisticRequest) (_result *GetInstanceStatusStatisticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceStatusStatisticResponse{}
	_body, _err := client.GetInstanceStatusStatisticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetManualDagInstances is deprecated
//
// Summary:
//
// Queries the information about instances in a manually triggered workflow.
//
// @param request - GetManualDagInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManualDagInstancesResponse
func (client *Client) GetManualDagInstancesWithOptions(request *GetManualDagInstancesRequest, runtime *dara.RuntimeOptions) (_result *GetManualDagInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		body["DagId"] = request.DagId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetManualDagInstances"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetManualDagInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetManualDagInstances is deprecated
//
// Summary:
//
// Queries the information about instances in a manually triggered workflow.
//
// @param request - GetManualDagInstancesRequest
//
// @return GetManualDagInstancesResponse
// Deprecated
func (client *Client) GetManualDagInstances(request *GetManualDagInstancesRequest) (_result *GetManualDagInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetManualDagInstancesResponse{}
	_body, _err := client.GetManualDagInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a category tree.
//
// @param request - GetMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaCategoryResponse
func (client *Client) GetMetaCategoryWithOptions(request *GetMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetMetaCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentCategoryId) {
		query["ParentCategoryId"] = request.ParentCategoryId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaCategory"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a category tree.
//
// @param request - GetMetaCategoryRequest
//
// @return GetMetaCategoryResponse
func (client *Client) GetMetaCategory(request *GetMetaCategoryRequest) (_result *GetMetaCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaCategoryResponse{}
	_body, _err := client.GetMetaCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a collection.
//
// @param request - GetMetaCollectionDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaCollectionDetailResponse
func (client *Client) GetMetaCollectionDetailWithOptions(request *GetMetaCollectionDetailRequest, runtime *dara.RuntimeOptions) (_result *GetMetaCollectionDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.QualifiedName) {
		query["QualifiedName"] = request.QualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaCollectionDetail"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaCollectionDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a collection.
//
// @param request - GetMetaCollectionDetailRequest
//
// @return GetMetaCollectionDetailResponse
func (client *Client) GetMetaCollectionDetail(request *GetMetaCollectionDetailRequest) (_result *GetMetaCollectionDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaCollectionDetailResponse{}
	_body, _err := client.GetMetaCollectionDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the lineage of a field in a metatable.
//
// @param request - GetMetaColumnLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaColumnLineageResponse
func (client *Client) GetMetaColumnLineageWithOptions(request *GetMetaColumnLineageRequest, runtime *dara.RuntimeOptions) (_result *GetMetaColumnLineageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ColumnGuid) {
		query["ColumnGuid"] = request.ColumnGuid
	}

	if !dara.IsNil(request.ColumnName) {
		query["ColumnName"] = request.ColumnName
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaColumnLineage"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaColumnLineageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lineage of a field in a metatable.
//
// @param request - GetMetaColumnLineageRequest
//
// @return GetMetaColumnLineageResponse
func (client *Client) GetMetaColumnLineage(request *GetMetaColumnLineageRequest) (_result *GetMetaColumnLineageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaColumnLineageResponse{}
	_body, _err := client.GetMetaColumnLineageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic metadata information about a compute engine instance.
//
// Description:
//
// The ID of the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
//
// You can log on to the [EMR console](https://emr.console.aliyun.com/?spm=a2c4g.11186623.0.0.965cc5c2GeiHet#/cn-hangzhou) to obtain the ID of the EMR cluster.
//
// @param request - GetMetaDBInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaDBInfoResponse
func (client *Client) GetMetaDBInfoWithOptions(request *GetMetaDBInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMetaDBInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaDBInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaDBInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic metadata information about a compute engine instance.
//
// Description:
//
// The ID of the EMR cluster. This parameter is required only if you set the DataSourceType parameter to emr.
//
// You can log on to the [EMR console](https://emr.console.aliyun.com/?spm=a2c4g.11186623.0.0.965cc5c2GeiHet#/cn-hangzhou) to obtain the ID of the EMR cluster.
//
// @param request - GetMetaDBInfoRequest
//
// @return GetMetaDBInfoResponse
func (client *Client) GetMetaDBInfo(request *GetMetaDBInfoRequest) (_result *GetMetaDBInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaDBInfoResponse{}
	_body, _err := client.GetMetaDBInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metatables in a compute engine instance.
//
// @param request - GetMetaDBTableListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaDBTableListResponse
func (client *Client) GetMetaDBTableListWithOptions(request *GetMetaDBTableListRequest, runtime *dara.RuntimeOptions) (_result *GetMetaDBTableListResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGuid) {
		query["AppGuid"] = request.AppGuid
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaDBTableList"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaDBTableListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metatables in a compute engine instance.
//
// @param request - GetMetaDBTableListRequest
//
// @return GetMetaDBTableListResponse
func (client *Client) GetMetaDBTableList(request *GetMetaDBTableListRequest) (_result *GetMetaDBTableListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaDBTableListResponse{}
	_body, _err := client.GetMetaDBTableListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information about a metatable.
//
// @param request - GetMetaTableBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableBasicInfoResponse
func (client *Client) GetMetaTableBasicInfoWithOptions(request *GetMetaTableBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableBasicInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableBasicInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableBasicInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information about a metatable.
//
// @param request - GetMetaTableBasicInfoRequest
//
// @return GetMetaTableBasicInfoResponse
func (client *Client) GetMetaTableBasicInfo(request *GetMetaTableBasicInfoRequest) (_result *GetMetaTableBasicInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableBasicInfoResponse{}
	_body, _err := client.GetMetaTableBasicInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the change logs of a metatable.
//
// Description:
//
// > This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - GetMetaTableChangeLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableChangeLogResponse
func (client *Client) GetMetaTableChangeLogWithOptions(request *GetMetaTableChangeLogRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableChangeLogResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChangeType) {
		body["ChangeType"] = request.ChangeType
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ObjectType) {
		body["ObjectType"] = request.ObjectType
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TableGuid) {
		body["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableChangeLog"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableChangeLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the change logs of a metatable.
//
// Description:
//
// > This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - GetMetaTableChangeLogRequest
//
// @return GetMetaTableChangeLogResponse
func (client *Client) GetMetaTableChangeLog(request *GetMetaTableChangeLogRequest) (_result *GetMetaTableChangeLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableChangeLogResponse{}
	_body, _err := client.GetMetaTableChangeLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the field information of a metatable.
//
// @param request - GetMetaTableColumnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableColumnResponse
func (client *Client) GetMetaTableColumnWithOptions(request *GetMetaTableColumnRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableColumnResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableColumn"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Queries the field information of a metatable.
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
// Queries the complete information about a table, including information about fields in the table.
//
// Description:
//
// You can call this operation to query only the information about a table of the E-MapReduce (EMR) compute engine type.
//
// @param request - GetMetaTableFullInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableFullInfoResponse
func (client *Client) GetMetaTableFullInfoWithOptions(request *GetMetaTableFullInfoRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableFullInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableFullInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableFullInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the complete information about a table, including information about fields in the table.
//
// Description:
//
// You can call this operation to query only the information about a table of the E-MapReduce (EMR) compute engine type.
//
// @param request - GetMetaTableFullInfoRequest
//
// @return GetMetaTableFullInfoResponse
func (client *Client) GetMetaTableFullInfo(request *GetMetaTableFullInfoRequest) (_result *GetMetaTableFullInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableFullInfoResponse{}
	_body, _err := client.GetMetaTableFullInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instructions on how to use a table.
//
// @param request - GetMetaTableIntroWikiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableIntroWikiResponse
func (client *Client) GetMetaTableIntroWikiWithOptions(request *GetMetaTableIntroWikiRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableIntroWikiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.WikiVersion) {
		query["WikiVersion"] = request.WikiVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableIntroWiki"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableIntroWikiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the instructions on how to use a table.
//
// @param request - GetMetaTableIntroWikiRequest
//
// @return GetMetaTableIntroWikiResponse
func (client *Client) GetMetaTableIntroWiki(request *GetMetaTableIntroWikiRequest) (_result *GetMetaTableIntroWikiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableIntroWikiResponse{}
	_body, _err := client.GetMetaTableIntroWikiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the lineage of a metatable.
//
// @param request - GetMetaTableLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableLineageResponse
func (client *Client) GetMetaTableLineageWithOptions(request *GetMetaTableLineageRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableLineageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.NextPrimaryKey) {
		query["NextPrimaryKey"] = request.NextPrimaryKey
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableLineage"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableLineageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lineage of a metatable.
//
// @param request - GetMetaTableLineageRequest
//
// @return GetMetaTableLineageResponse
func (client *Client) GetMetaTableLineage(request *GetMetaTableLineageRequest) (_result *GetMetaTableLineageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableLineageResponse{}
	_body, _err := client.GetMetaTableLineageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metatables in a specified category.
//
// @param request - GetMetaTableListByCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableListByCategoryResponse
func (client *Client) GetMetaTableListByCategoryWithOptions(request *GetMetaTableListByCategoryRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableListByCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableListByCategory"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableListByCategoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metatables in a specified category.
//
// @param request - GetMetaTableListByCategoryRequest
//
// @return GetMetaTableListByCategoryResponse
func (client *Client) GetMetaTableListByCategory(request *GetMetaTableListByCategoryRequest) (_result *GetMetaTableListByCategoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableListByCategoryResponse{}
	_body, _err := client.GetMetaTableListByCategoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the output information of a metatable.
//
// @param request - GetMetaTableOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableOutputResponse
func (client *Client) GetMetaTableOutputWithOptions(request *GetMetaTableOutputRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		query["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartDate) {
		query["StartDate"] = request.StartDate
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableOutput"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the output information of a metatable.
//
// @param request - GetMetaTableOutputRequest
//
// @return GetMetaTableOutputResponse
func (client *Client) GetMetaTableOutput(request *GetMetaTableOutputRequest) (_result *GetMetaTableOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableOutputResponse{}
	_body, _err := client.GetMetaTableOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of partitions in a metatable.
//
// Description:
//
// You can call this operation to query only the partitions of a metatable in a MaxCompute or E-MapReduce (EMR) compute engine. If you query partitions of a metatable in an EMR compute engine, only DataLake clusters that use Data Lake Formation (DLF) to manage metadata and Hadoop clusters whose cluster version is earlier than 3.41.0 or 5.7.0 are supported.
//
// @param tmpReq - GetMetaTablePartitionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTablePartitionResponse
func (client *Client) GetMetaTablePartitionWithOptions(tmpReq *GetMetaTablePartitionRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTablePartitionResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetMetaTablePartitionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SortCriterion) {
		request.SortCriterionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SortCriterion, dara.String("SortCriterion"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DatabaseName) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortCriterionShrink) {
		query["SortCriterion"] = request.SortCriterionShrink
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTablePartition"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTablePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of partitions in a metatable.
//
// Description:
//
// You can call this operation to query only the partitions of a metatable in a MaxCompute or E-MapReduce (EMR) compute engine. If you query partitions of a metatable in an EMR compute engine, only DataLake clusters that use Data Lake Formation (DLF) to manage metadata and Hadoop clusters whose cluster version is earlier than 3.41.0 or 5.7.0 are supported.
//
// @param request - GetMetaTablePartitionRequest
//
// @return GetMetaTablePartitionResponse
func (client *Client) GetMetaTablePartition(request *GetMetaTablePartitionRequest) (_result *GetMetaTablePartitionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTablePartitionResponse{}
	_body, _err := client.GetMetaTablePartitionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the output tasks of a metatable.
//
// @param request - GetMetaTableProducingTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableProducingTasksResponse
func (client *Client) GetMetaTableProducingTasksWithOptions(request *GetMetaTableProducingTasksRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableProducingTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.DbName) {
		query["DbName"] = request.DbName
	}

	if !dara.IsNil(request.SchemaName) {
		query["SchemaName"] = request.SchemaName
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableProducingTasks"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableProducingTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the output tasks of a metatable.
//
// @param request - GetMetaTableProducingTasksRequest
//
// @return GetMetaTableProducingTasksResponse
func (client *Client) GetMetaTableProducingTasks(request *GetMetaTableProducingTasksRequest) (_result *GetMetaTableProducingTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableProducingTasksResponse{}
	_body, _err := client.GetMetaTableProducingTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the themes and levels of a metatable.
//
// @param request - GetMetaTableThemeLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMetaTableThemeLevelResponse
func (client *Client) GetMetaTableThemeLevelWithOptions(request *GetMetaTableThemeLevelRequest, runtime *dara.RuntimeOptions) (_result *GetMetaTableThemeLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMetaTableThemeLevel"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMetaTableThemeLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the themes and levels of a metatable.
//
// @param request - GetMetaTableThemeLevelRequest
//
// @return GetMetaTableThemeLevelResponse
func (client *Client) GetMetaTableThemeLevel(request *GetMetaTableThemeLevelRequest) (_result *GetMetaTableThemeLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMetaTableThemeLevelResponse{}
	_body, _err := client.GetMetaTableThemeLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the progress of a migration task.
//
// @param request - GetMigrationProcessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMigrationProcessResponse
func (client *Client) GetMigrationProcessWithOptions(request *GetMigrationProcessRequest, runtime *dara.RuntimeOptions) (_result *GetMigrationProcessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MigrationId) {
		body["MigrationId"] = request.MigrationId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMigrationProcess"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMigrationProcessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the progress of a migration task.
//
// @param request - GetMigrationProcessRequest
//
// @return GetMigrationProcessResponse
func (client *Client) GetMigrationProcess(request *GetMigrationProcessRequest) (_result *GetMigrationProcessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMigrationProcessResponse{}
	_body, _err := client.GetMigrationProcessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a migration task.
//
// @param request - GetMigrationSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMigrationSummaryResponse
func (client *Client) GetMigrationSummaryWithOptions(request *GetMigrationSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetMigrationSummaryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MigrationId) {
		body["MigrationId"] = request.MigrationId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMigrationSummary"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMigrationSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a migration task.
//
// @param request - GetMigrationSummaryRequest
//
// @return GetMigrationSummaryResponse
func (client *Client) GetMigrationSummary(request *GetMigrationSummaryRequest) (_result *GetMigrationSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetMigrationSummaryResponse{}
	_body, _err := client.GetMigrationSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Indicates whether the request is successful.
//
// @param request - GetNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeResponse
func (client *Client) GetNodeWithOptions(request *GetNodeRequest, runtime *dara.RuntimeOptions) (_result *GetNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Indicates whether the request is successful.
//
// @param request - GetNodeRequest
//
// @return GetNodeResponse
func (client *Client) GetNode(request *GetNodeRequest) (_result *GetNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeResponse{}
	_body, _err := client.GetNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - GetNodeChildrenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeChildrenResponse
func (client *Client) GetNodeChildrenWithOptions(request *GetNodeChildrenRequest, runtime *dara.RuntimeOptions) (_result *GetNodeChildrenResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeChildren"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeChildrenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - GetNodeChildrenRequest
//
// @return GetNodeChildrenResponse
func (client *Client) GetNodeChildren(request *GetNodeChildrenRequest) (_result *GetNodeChildrenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeChildrenResponse{}
	_body, _err := client.GetNodeChildrenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the code of a node.
//
// @param request - GetNodeCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeCodeResponse
func (client *Client) GetNodeCodeWithOptions(request *GetNodeCodeRequest, runtime *dara.RuntimeOptions) (_result *GetNodeCodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeCode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the code of a node.
//
// @param request - GetNodeCodeRequest
//
// @return GetNodeCodeResponse
func (client *Client) GetNodeCode(request *GetNodeCodeRequest) (_result *GetNodeCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeCodeResponse{}
	_body, _err := client.GetNodeCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetNodeOnBaseline is deprecated
//
// Summary:
//
// Queries the nodes associated with a baseline.
//
// @param request - GetNodeOnBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeOnBaselineResponse
func (client *Client) GetNodeOnBaselineWithOptions(request *GetNodeOnBaselineRequest, runtime *dara.RuntimeOptions) (_result *GetNodeOnBaselineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeOnBaseline"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeOnBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetNodeOnBaseline is deprecated
//
// Summary:
//
// Queries the nodes associated with a baseline.
//
// @param request - GetNodeOnBaselineRequest
//
// @return GetNodeOnBaselineResponse
// Deprecated
func (client *Client) GetNodeOnBaseline(request *GetNodeOnBaselineRequest) (_result *GetNodeOnBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeOnBaselineResponse{}
	_body, _err := client.GetNodeOnBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor nodes of a node.
//
// @param request - GetNodeParentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeParentsResponse
func (client *Client) GetNodeParentsWithOptions(request *GetNodeParentsRequest, runtime *dara.RuntimeOptions) (_result *GetNodeParentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeParents"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeParentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of ancestor nodes of a node.
//
// @param request - GetNodeParentsRequest
//
// @return GetNodeParentsResponse
func (client *Client) GetNodeParents(request *GetNodeParentsRequest) (_result *GetNodeParentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeParentsResponse{}
	_body, _err := client.GetNodeParentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetNodeTypeListInfo is deprecated
//
// Summary:
//
// Queries the information about node types, including the code and name of a node type.
//
// @param request - GetNodeTypeListInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNodeTypeListInfoResponse
func (client *Client) GetNodeTypeListInfoWithOptions(request *GetNodeTypeListInfoRequest, runtime *dara.RuntimeOptions) (_result *GetNodeTypeListInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Locale) {
		body["Locale"] = request.Locale
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetNodeTypeListInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetNodeTypeListInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetNodeTypeListInfo is deprecated
//
// Summary:
//
// Queries the information about node types, including the code and name of a node type.
//
// @param request - GetNodeTypeListInfoRequest
//
// @return GetNodeTypeListInfoResponse
// Deprecated
func (client *Client) GetNodeTypeListInfo(request *GetNodeTypeListInfoRequest) (_result *GetNodeTypeListInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetNodeTypeListInfoResponse{}
	_body, _err := client.GetNodeTypeListInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the records that are generated on a specified date for access to the high-risk sensitive data in all the DataWorks workspaces of a tenant.
//
// @param request - GetOpRiskDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpRiskDataResponse
func (client *Client) GetOpRiskDataWithOptions(request *GetOpRiskDataRequest, runtime *dara.RuntimeOptions) (_result *GetOpRiskDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpRiskData"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpRiskDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the records that are generated on a specified date for access to the high-risk sensitive data in all the DataWorks workspaces of a tenant.
//
// @param request - GetOpRiskDataRequest
//
// @return GetOpRiskDataResponse
func (client *Client) GetOpRiskData(request *GetOpRiskDataRequest) (_result *GetOpRiskDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOpRiskDataResponse{}
	_body, _err := client.GetOpRiskDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the records that are generated on a specified date for access to sensitive data in all the DataWorks workspaces of a tenant.
//
// @param request - GetOpSensitiveDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpSensitiveDataResponse
func (client *Client) GetOpSensitiveDataWithOptions(request *GetOpSensitiveDataRequest, runtime *dara.RuntimeOptions) (_result *GetOpSensitiveDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOpSensitiveData"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOpSensitiveDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the records that are generated on a specified date for access to sensitive data in all the DataWorks workspaces of a tenant.
//
// @param request - GetOpSensitiveDataRequest
//
// @return GetOpSensitiveDataResponse
func (client *Client) GetOpSensitiveData(request *GetOpSensitiveDataRequest) (_result *GetOpSensitiveDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOpSensitiveDataResponse{}
	_body, _err := client.GetOpSensitiveDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the option settings of an extension in a workspace.
//
// @param request - GetOptionValueForProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOptionValueForProjectResponse
func (client *Client) GetOptionValueForProjectWithOptions(request *GetOptionValueForProjectRequest, runtime *dara.RuntimeOptions) (_result *GetOptionValueForProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExtensionCode) {
		body["ExtensionCode"] = request.ExtensionCode
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOptionValueForProject"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOptionValueForProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the option settings of an extension in a workspace.
//
// @param request - GetOptionValueForProjectRequest
//
// @return GetOptionValueForProjectResponse
func (client *Client) GetOptionValueForProject(request *GetOptionValueForProjectRequest) (_result *GetOptionValueForProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOptionValueForProjectResponse{}
	_body, _err := client.GetOptionValueForProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a permission request order.
//
// @param request - GetPermissionApplyOrderDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPermissionApplyOrderDetailResponse
func (client *Client) GetPermissionApplyOrderDetailWithOptions(request *GetPermissionApplyOrderDetailRequest, runtime *dara.RuntimeOptions) (_result *GetPermissionApplyOrderDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowId) {
		query["FlowId"] = request.FlowId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPermissionApplyOrderDetail"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPermissionApplyOrderDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a permission request order.
//
// @param request - GetPermissionApplyOrderDetailRequest
//
// @return GetPermissionApplyOrderDetailResponse
func (client *Client) GetPermissionApplyOrderDetail(request *GetPermissionApplyOrderDetailRequest) (_result *GetPermissionApplyOrderDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetPermissionApplyOrderDetailResponse{}
	_body, _err := client.GetPermissionApplyOrderDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a DataWorks workspace.
//
// @param request - GetProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectResponse
func (client *Client) GetProjectWithOptions(request *GetProjectRequest, runtime *dara.RuntimeOptions) (_result *GetProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		query["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProject"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a DataWorks workspace.
//
// @param request - GetProjectRequest
//
// @return GetProjectResponse
func (client *Client) GetProject(request *GetProjectRequest) (_result *GetProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectResponse{}
	_body, _err := client.GetProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetProjectDetail is deprecated
//
// Summary:
//
// Queries the information about a DataWorks workspace.
//
// @param request - GetProjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectDetailResponse
func (client *Client) GetProjectDetailWithOptions(request *GetProjectDetailRequest, runtime *dara.RuntimeOptions) (_result *GetProjectDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectDetail"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetProjectDetail is deprecated
//
// Summary:
//
// Queries the information about a DataWorks workspace.
//
// @param request - GetProjectDetailRequest
//
// @return GetProjectDetailResponse
// Deprecated
func (client *Client) GetProjectDetail(request *GetProjectDetailRequest) (_result *GetProjectDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetProjectDetailResponse{}
	_body, _err := client.GetProjectDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetQualityEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityEntityResponse
func (client *Client) GetQualityEntityWithOptions(request *GetQualityEntityRequest, runtime *dara.RuntimeOptions) (_result *GetQualityEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.MatchExpression) {
		body["MatchExpression"] = request.MatchExpression
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.TableName) {
		body["TableName"] = request.TableName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityEntity"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetQualityEntityRequest
//
// @return GetQualityEntityResponse
func (client *Client) GetQualityEntity(request *GetQualityEntityRequest) (_result *GetQualityEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityEntityResponse{}
	_body, _err := client.GetQualityEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the subscribers of a partition filter expression.
//
// @param request - GetQualityFollowerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityFollowerResponse
func (client *Client) GetQualityFollowerWithOptions(request *GetQualityFollowerRequest, runtime *dara.RuntimeOptions) (_result *GetQualityFollowerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityFollower"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityFollowerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the subscribers of a partition filter expression.
//
// @param request - GetQualityFollowerRequest
//
// @return GetQualityFollowerResponse
func (client *Client) GetQualityFollower(request *GetQualityFollowerRequest) (_result *GetQualityFollowerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityFollowerResponse{}
	_body, _err := client.GetQualityFollowerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a monitoring rule.
//
// @param request - GetQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityRuleResponse
func (client *Client) GetQualityRuleWithOptions(request *GetQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *GetQualityRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQualityRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a monitoring rule.
//
// @param request - GetQualityRuleRequest
//
// @return GetQualityRuleResponse
func (client *Client) GetQualityRule(request *GetQualityRuleRequest) (_result *GetQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetQualityRuleResponse{}
	_body, _err := client.GetQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a custom alert rule.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=dataworks-public\\&api=GetRemind\\&type=RPC\\&version=2020-05-18)
//
// @param request - GetRemindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRemindResponse
func (client *Client) GetRemindWithOptions(request *GetRemindRequest, runtime *dara.RuntimeOptions) (_result *GetRemindResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RemindId) {
		body["RemindId"] = request.RemindId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRemind"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRemindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a custom alert rule.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=dataworks-public\\&api=GetRemind\\&type=RPC\\&version=2020-05-18)
//
// @param request - GetRemindRequest
//
// @return GetRemindResponse
func (client *Client) GetRemind(request *GetRemindRequest) (_result *GetRemindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetRemindResponse{}
	_body, _err := client.GetRemindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the latest sensitive data in all the DataWorks workspaces of a tenant.
//
// @param request - GetSensitiveDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSensitiveDataResponse
func (client *Client) GetSensitiveDataWithOptions(request *GetSensitiveDataRequest, runtime *dara.RuntimeOptions) (_result *GetSensitiveDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSensitiveData"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSensitiveDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the latest sensitive data in all the DataWorks workspaces of a tenant.
//
// @param request - GetSensitiveDataRequest
//
// @return GetSensitiveDataResponse
func (client *Client) GetSensitiveData(request *GetSensitiveDataRequest) (_result *GetSensitiveDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSensitiveDataResponse{}
	_body, _err := client.GetSensitiveDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI GetSuccessInstanceTrend is deprecated
//
// Summary:
//
// Queries the statistics of instances in different periods of a day.
//
// @param request - GetSuccessInstanceTrendRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSuccessInstanceTrendResponse
func (client *Client) GetSuccessInstanceTrendWithOptions(request *GetSuccessInstanceTrendRequest, runtime *dara.RuntimeOptions) (_result *GetSuccessInstanceTrendResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSuccessInstanceTrend"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSuccessInstanceTrendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI GetSuccessInstanceTrend is deprecated
//
// Summary:
//
// Queries the statistics of instances in different periods of a day.
//
// @param request - GetSuccessInstanceTrendRequest
//
// @return GetSuccessInstanceTrendResponse
// Deprecated
func (client *Client) GetSuccessInstanceTrend(request *GetSuccessInstanceTrendRequest) (_result *GetSuccessInstanceTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetSuccessInstanceTrendResponse{}
	_body, _err := client.GetSuccessInstanceTrendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an event.
//
// Description:
//
// ***
//
// @param request - GetTopicRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicResponse
func (client *Client) GetTopicWithOptions(request *GetTopicRequest, runtime *dara.RuntimeOptions) (_result *GetTopicResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TopicId) {
		body["TopicId"] = request.TopicId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopic"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an event.
//
// Description:
//
// ***
//
// @param request - GetTopicRequest
//
// @return GetTopicResponse
func (client *Client) GetTopic(request *GetTopicRequest) (_result *GetTopicResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTopicResponse{}
	_body, _err := client.GetTopicWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries baseline instances affected by an event.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=dataworks-public\\&api=GetTopicInfluence\\&type=RPC\\&version=2020-05-18)
//
// @param request - GetTopicInfluenceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTopicInfluenceResponse
func (client *Client) GetTopicInfluenceWithOptions(request *GetTopicInfluenceRequest, runtime *dara.RuntimeOptions) (_result *GetTopicInfluenceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TopicId) {
		body["TopicId"] = request.TopicId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTopicInfluence"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTopicInfluenceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries baseline instances affected by an event.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=dataworks-public\\&api=GetTopicInfluence\\&type=RPC\\&version=2020-05-18)
//
// @param request - GetTopicInfluenceRequest
//
// @return GetTopicInfluenceResponse
func (client *Client) GetTopicInfluence(request *GetTopicInfluenceRequest) (_result *GetTopicInfluenceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTopicInfluenceResponse{}
	_body, _err := client.GetTopicInfluenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports data sources from your on-premises machine to a specific DataWorks workspace.
//
// Description:
//
// You can import self-managed data sources or data sources that are exported from other DataWorks workspaces to a specific DataWorks workspace.
//
//   - To import a self-managed data source to a DataWorks workspace, the data source type must be supported by DataWorks. For more information about the types of data sources supported by DataWorks, see [Supported data stores](https://help.aliyun.com/document_detail/181656.html).
//
//   - For more information about how to export data sources from DataWorks workspaces to your on-premises machine, see [ExportDataSources](https://help.aliyun.com/document_detail/279570.html).
//
// @param request - ImportDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportDataSourcesResponse
func (client *Client) ImportDataSourcesWithOptions(request *ImportDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ImportDataSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSources) {
		query["DataSources"] = request.DataSources
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportDataSources"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports data sources from your on-premises machine to a specific DataWorks workspace.
//
// Description:
//
// You can import self-managed data sources or data sources that are exported from other DataWorks workspaces to a specific DataWorks workspace.
//
//   - To import a self-managed data source to a DataWorks workspace, the data source type must be supported by DataWorks. For more information about the types of data sources supported by DataWorks, see [Supported data stores](https://help.aliyun.com/document_detail/181656.html).
//
//   - For more information about how to export data sources from DataWorks workspaces to your on-premises machine, see [ExportDataSources](https://help.aliyun.com/document_detail/279570.html).
//
// @param request - ImportDataSourcesRequest
//
// @return ImportDataSourcesResponse
func (client *Client) ImportDataSources(request *ImportDataSourcesRequest) (_result *ImportDataSourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportDataSourcesResponse{}
	_body, _err := client.ImportDataSourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of alerts.
//
// @param request - ListAlertMessagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAlertMessagesResponse
func (client *Client) ListAlertMessagesWithOptions(request *ListAlertMessagesRequest, runtime *dara.RuntimeOptions) (_result *ListAlertMessagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertMethods) {
		body["AlertMethods"] = request.AlertMethods
	}

	if !dara.IsNil(request.AlertRuleTypes) {
		body["AlertRuleTypes"] = request.AlertRuleTypes
	}

	if !dara.IsNil(request.AlertUser) {
		body["AlertUser"] = request.AlertUser
	}

	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.BeginTime) {
		body["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RemindId) {
		body["RemindId"] = request.RemindId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAlertMessages"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAlertMessagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of alerts.
//
// @param request - ListAlertMessagesRequest
//
// @return ListAlertMessagesResponse
func (client *Client) ListAlertMessages(request *ListAlertMessagesRequest) (_result *ListAlertMessagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAlertMessagesResponse{}
	_body, _err := client.ListAlertMessagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of baselines.
//
// @param request - ListBaselineConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaselineConfigsResponse
func (client *Client) ListBaselineConfigsWithOptions(request *ListBaselineConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListBaselineConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineTypes) {
		body["BaselineTypes"] = request.BaselineTypes
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SearchText) {
		body["SearchText"] = request.SearchText
	}

	if !dara.IsNil(request.Useflag) {
		body["Useflag"] = request.Useflag
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaselineConfigs"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaselineConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of baselines.
//
// @param request - ListBaselineConfigsRequest
//
// @return ListBaselineConfigsResponse
func (client *Client) ListBaselineConfigs(request *ListBaselineConfigsRequest) (_result *ListBaselineConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaselineConfigsResponse{}
	_body, _err := client.ListBaselineConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of baseline instances.
//
// @param request - ListBaselineStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaselineStatusesResponse
func (client *Client) ListBaselineStatusesWithOptions(request *ListBaselineStatusesRequest, runtime *dara.RuntimeOptions) (_result *ListBaselineStatusesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineTypes) {
		body["BaselineTypes"] = request.BaselineTypes
	}

	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.FinishStatus) {
		body["FinishStatus"] = request.FinishStatus
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.SearchText) {
		body["SearchText"] = request.SearchText
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	if !dara.IsNil(request.TopicId) {
		body["TopicId"] = request.TopicId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaselineStatuses"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaselineStatusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of baseline instances.
//
// @param request - ListBaselineStatusesRequest
//
// @return ListBaselineStatusesResponse
func (client *Client) ListBaselineStatuses(request *ListBaselineStatusesRequest) (_result *ListBaselineStatusesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaselineStatusesResponse{}
	_body, _err := client.ListBaselineStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a list of baselines.
//
// @param request - ListBaselinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBaselinesResponse
func (client *Client) ListBaselinesWithOptions(request *ListBaselinesRequest, runtime *dara.RuntimeOptions) (_result *ListBaselinesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineTypes) {
		body["BaselineTypes"] = request.BaselineTypes
	}

	if !dara.IsNil(request.Enable) {
		body["Enable"] = request.Enable
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SearchText) {
		body["SearchText"] = request.SearchText
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBaselines"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBaselinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a list of baselines.
//
// @param request - ListBaselinesRequest
//
// @return ListBaselinesResponse
func (client *Client) ListBaselines(request *ListBaselinesRequest) (_result *ListBaselinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBaselinesResponse{}
	_body, _err := client.ListBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of workflows.
//
// @param request - ListBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBusinessResponse
func (client *Client) ListBusinessWithOptions(request *ListBusinessRequest, runtime *dara.RuntimeOptions) (_result *ListBusinessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBusiness"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of workflows.
//
// @param request - ListBusinessRequest
//
// @return ListBusinessResponse
func (client *Client) ListBusiness(request *ListBusinessRequest) (_result *ListBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListBusinessResponse{}
	_body, _err := client.ListBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of compute engines that are associated with a DataWorks workspace.
//
// @param request - ListCalcEnginesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCalcEnginesResponse
func (client *Client) ListCalcEnginesWithOptions(request *ListCalcEnginesRequest, runtime *dara.RuntimeOptions) (_result *ListCalcEnginesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalcEngineType) {
		query["CalcEngineType"] = request.CalcEngineType
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCalcEngines"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCalcEnginesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of compute engines that are associated with a DataWorks workspace.
//
// @param request - ListCalcEnginesRequest
//
// @return ListCalcEnginesResponse
func (client *Client) ListCalcEngines(request *ListCalcEnginesRequest) (_result *ListCalcEnginesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCalcEnginesResponse{}
	_body, _err := client.ListCalcEnginesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the check results of extension point events.
//
// @param request - ListCheckProcessesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCheckProcessesResponse
func (client *Client) ListCheckProcessesWithOptions(request *ListCheckProcessesRequest, runtime *dara.RuntimeOptions) (_result *ListCheckProcessesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventCode) {
		body["EventCode"] = request.EventCode
	}

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCheckProcesses"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCheckProcessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the check results of extension point events.
//
// @param request - ListCheckProcessesRequest
//
// @return ListCheckProcessesResponse
func (client *Client) ListCheckProcesses(request *ListCheckProcessesRequest) (_result *ListCheckProcessesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCheckProcessesResponse{}
	_body, _err := client.ListCheckProcessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of submodules in a workspace. You can query information about SPARK parameters.
//
// @param request - ListClusterConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterConfigsResponse
func (client *Client) ListClusterConfigsWithOptions(request *ListClusterConfigsRequest, runtime *dara.RuntimeOptions) (_result *ListClusterConfigsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterConfigs"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClusterConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of submodules in a workspace. You can query information about SPARK parameters.
//
// @param request - ListClusterConfigsRequest
//
// @return ListClusterConfigsResponse
func (client *Client) ListClusterConfigs(request *ListClusterConfigsRequest) (_result *ListClusterConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClusterConfigsResponse{}
	_body, _err := client.ListClusterConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries clusters that are registered in DataWorks. E-MapReduce (EMR) clusters and Cloudera\\"s Distribution Including Apache Hadoop (CDH) clusters are supported.
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries clusters that are registered in DataWorks. E-MapReduce (EMR) clusters and Cloudera\\"s Distribution Including Apache Hadoop (CDH) clusters are supported.
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListConnections is deprecated
//
// Summary:
//
// Queries a list of data sources.
//
// @param request - ListConnectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectionsResponse
func (client *Client) ListConnectionsWithOptions(request *ListConnectionsRequest, runtime *dara.RuntimeOptions) (_result *ListConnectionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConnections"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConnectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListConnections is deprecated
//
// Summary:
//
// Queries a list of data sources.
//
// @param request - ListConnectionsRequest
//
// @return ListConnectionsResponse
// Deprecated
func (client *Client) ListConnections(request *ListConnectionsRequest) (_result *ListConnectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListConnectionsResponse{}
	_body, _err := client.ListConnectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of alert rules configured for a new-version synchronization task. The following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization.
//
// @param request - ListDIAlarmRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRulesWithOptions(request *ListDIAlarmRulesRequest, runtime *dara.RuntimeOptions) (_result *ListDIAlarmRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDIAlarmRules"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of alert rules configured for a new-version synchronization task. The following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization.
//
// @param request - ListDIAlarmRulesRequest
//
// @return ListDIAlarmRulesResponse
func (client *Client) ListDIAlarmRules(request *ListDIAlarmRulesRequest) (_result *ListDIAlarmRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIAlarmRulesResponse{}
	_body, _err := client.ListDIAlarmRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of new-version synchronization tasks. The following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can call this operation to obtain only the basic information about the tasks. If you want to obtain the details of a task, call the GetDIJob operation.
//
// @param request - ListDIJobsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobsWithOptions(request *ListDIJobsRequest, runtime *dara.RuntimeOptions) (_result *ListDIJobsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DestinationDataSourceType) {
		body["DestinationDataSourceType"] = request.DestinationDataSourceType
	}

	if !dara.IsNil(request.JobName) {
		body["JobName"] = request.JobName
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SourceDataSourceType) {
		body["SourceDataSourceType"] = request.SourceDataSourceType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDIJobs"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of new-version synchronization tasks. The following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// Description:
//
// You can call this operation to obtain only the basic information about the tasks. If you want to obtain the details of a task, call the GetDIJob operation.
//
// @param request - ListDIJobsRequest
//
// @return ListDIJobsResponse
func (client *Client) ListDIJobs(request *ListDIJobsRequest) (_result *ListDIJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIJobsResponse{}
	_body, _err := client.ListDIJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default global configurations of synchronization solutions in a specified DataWorks workspace.
//
// Description:
//
// DataWorks allows you to specify a default global configuration only for the processing rules of DDL messages in synchronization solutions. After you configure the **processing rules of DDL messages*	- in synchronization solutions, the configuration is used as the default global configuration and applies to all real-time synchronization tasks in the solutions. You can modify the **processing rules of DDL messages*	- based on your business requirements. For more information about how to configure a synchronization solution, see [Synchronization solutions](https://help.aliyun.com/document_detail/199008.html).
//
// @param request - ListDIProjectConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDIProjectConfigResponse
func (client *Client) ListDIProjectConfigWithOptions(request *ListDIProjectConfigRequest, runtime *dara.RuntimeOptions) (_result *ListDIProjectConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDIProjectConfig"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDIProjectConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default global configurations of synchronization solutions in a specified DataWorks workspace.
//
// Description:
//
// DataWorks allows you to specify a default global configuration only for the processing rules of DDL messages in synchronization solutions. After you configure the **processing rules of DDL messages*	- in synchronization solutions, the configuration is used as the default global configuration and applies to all real-time synchronization tasks in the solutions. You can modify the **processing rules of DDL messages*	- based on your business requirements. For more information about how to configure a synchronization solution, see [Synchronization solutions](https://help.aliyun.com/document_detail/199008.html).
//
// @param request - ListDIProjectConfigRequest
//
// @return ListDIProjectConfigResponse
func (client *Client) ListDIProjectConfig(request *ListDIProjectConfigRequest) (_result *ListDIProjectConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDIProjectConfigResponse{}
	_body, _err := client.ListDIProjectConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of directed acyclic graphs (DAGs) for a single data backfill instance based on OpSeq.
//
// Description:
//
// Supported DAG types:
//
//   - MANUAL: DAG for a manually triggered workflow
//
//   - SMOKE_TEST: DAG for a smoke testing workflow
//
//   - SUPPLY_DATA: DAG for a data backfill instance
//
//   - BUSINESS_PROCESS_DAG: DAG for a one-time workflow
//
// Supported DAG states:
//
//   - CREATED: The DAG is created.
//
//   - RUNNING: The DAG is running.
//
//   - FAILURE: The DAG fails to run.
//
//   - SUCCESS: The DAG is successfully run.
//
// @param request - ListDagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDagsResponse
func (client *Client) ListDagsWithOptions(request *ListDagsRequest, runtime *dara.RuntimeOptions) (_result *ListDagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OpSeq) {
		body["OpSeq"] = request.OpSeq
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDags"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of directed acyclic graphs (DAGs) for a single data backfill instance based on OpSeq.
//
// Description:
//
// Supported DAG types:
//
//   - MANUAL: DAG for a manually triggered workflow
//
//   - SMOKE_TEST: DAG for a smoke testing workflow
//
//   - SUPPLY_DATA: DAG for a data backfill instance
//
//   - BUSINESS_PROCESS_DAG: DAG for a one-time workflow
//
// Supported DAG states:
//
//   - CREATED: The DAG is created.
//
//   - RUNNING: The DAG is running.
//
//   - FAILURE: The DAG fails to run.
//
//   - SUCCESS: The DAG is successfully run.
//
// @param request - ListDagsRequest
//
// @return ListDagsResponse
func (client *Client) ListDags(request *ListDagsRequest) (_result *ListDagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDagsResponse{}
	_body, _err := client.ListDagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs on which other users are granted the access permissions.
//
// @param request - ListDataServiceApiAuthoritiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiAuthoritiesResponse
func (client *Client) ListDataServiceApiAuthoritiesWithOptions(request *ListDataServiceApiAuthoritiesRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiAuthoritiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameKeyword) {
		body["ApiNameKeyword"] = request.ApiNameKeyword
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiAuthorities"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs on which other users are granted the access permissions.
//
// @param request - ListDataServiceApiAuthoritiesRequest
//
// @return ListDataServiceApiAuthoritiesResponse
func (client *Client) ListDataServiceApiAuthorities(request *ListDataServiceApiAuthoritiesRequest) (_result *ListDataServiceApiAuthoritiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApiAuthoritiesResponse{}
	_body, _err := client.ListDataServiceApiAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the test records of a DataService Studio API. This API operation allows you to query only the test records that are generated within the previous month.
//
// @param request - ListDataServiceApiTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApiTestResponse
func (client *Client) ListDataServiceApiTestWithOptions(request *ListDataServiceApiTestRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApiTestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApiTest"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApiTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the test records of a DataService Studio API. This API operation allows you to query only the test records that are generated within the previous month.
//
// @param request - ListDataServiceApiTestRequest
//
// @return ListDataServiceApiTestResponse
func (client *Client) ListDataServiceApiTest(request *ListDataServiceApiTestRequest) (_result *ListDataServiceApiTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApiTestResponse{}
	_body, _err := client.ListDataServiceApiTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of APIs in the development state.
//
// @param request - ListDataServiceApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApisResponse
func (client *Client) ListDataServiceApisWithOptions(request *ListDataServiceApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameKeyword) {
		body["ApiNameKeyword"] = request.ApiNameKeyword
	}

	if !dara.IsNil(request.ApiPathKeyword) {
		body["ApiPathKeyword"] = request.ApiPathKeyword
	}

	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApis"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of APIs in the development state.
//
// @param request - ListDataServiceApisRequest
//
// @return ListDataServiceApisResponse
func (client *Client) ListDataServiceApis(request *ListDataServiceApisRequest) (_result *ListDataServiceApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApisResponse{}
	_body, _err := client.ListDataServiceApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the basic information of applications.
//
// @param request - ListDataServiceApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceApplicationsResponse
func (client *Client) ListDataServiceApplicationsWithOptions(request *ListDataServiceApplicationsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceApplicationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectIdList) {
		body["ProjectIdList"] = request.ProjectIdList
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceApplications"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the basic information of applications.
//
// @param request - ListDataServiceApplicationsRequest
//
// @return ListDataServiceApplicationsResponse
func (client *Client) ListDataServiceApplications(request *ListDataServiceApplicationsRequest) (_result *ListDataServiceApplicationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceApplicationsResponse{}
	_body, _err := client.ListDataServiceApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the APIs that you are authorized to access.
//
// @param request - ListDataServiceAuthorizedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceAuthorizedApisResponse
func (client *Client) ListDataServiceAuthorizedApisWithOptions(request *ListDataServiceAuthorizedApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceAuthorizedApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameKeyword) {
		body["ApiNameKeyword"] = request.ApiNameKeyword
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceAuthorizedApis"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceAuthorizedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the APIs that you are authorized to access.
//
// @param request - ListDataServiceAuthorizedApisRequest
//
// @return ListDataServiceAuthorizedApisResponse
func (client *Client) ListDataServiceAuthorizedApis(request *ListDataServiceAuthorizedApisRequest) (_result *ListDataServiceAuthorizedApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceAuthorizedApisResponse{}
	_body, _err := client.ListDataServiceAuthorizedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries folders.
//
// @param request - ListDataServiceFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceFoldersResponse
func (client *Client) ListDataServiceFoldersWithOptions(request *ListDataServiceFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceFoldersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderNameKeyword) {
		body["FolderNameKeyword"] = request.FolderNameKeyword
	}

	if !dara.IsNil(request.GroupId) {
		body["GroupId"] = request.GroupId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceFolders"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceFoldersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries folders.
//
// @param request - ListDataServiceFoldersRequest
//
// @return ListDataServiceFoldersResponse
func (client *Client) ListDataServiceFolders(request *ListDataServiceFoldersRequest) (_result *ListDataServiceFoldersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceFoldersResponse{}
	_body, _err := client.ListDataServiceFoldersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries business processes.
//
// @param request - ListDataServiceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServiceGroupsResponse
func (client *Client) ListDataServiceGroupsWithOptions(request *ListDataServiceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListDataServiceGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GroupNameKeyword) {
		body["GroupNameKeyword"] = request.GroupNameKeyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServiceGroups"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServiceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries business processes.
//
// @param request - ListDataServiceGroupsRequest
//
// @return ListDataServiceGroupsResponse
func (client *Client) ListDataServiceGroups(request *ListDataServiceGroupsRequest) (_result *ListDataServiceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServiceGroupsResponse{}
	_body, _err := client.ListDataServiceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of APIs in the published state.
//
// @param request - ListDataServicePublishedApisRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApisWithOptions(request *ListDataServicePublishedApisRequest, runtime *dara.RuntimeOptions) (_result *ListDataServicePublishedApisResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiNameKeyword) {
		body["ApiNameKeyword"] = request.ApiNameKeyword
	}

	if !dara.IsNil(request.ApiPathKeyword) {
		body["ApiPathKeyword"] = request.ApiPathKeyword
	}

	if !dara.IsNil(request.CreatorId) {
		body["CreatorId"] = request.CreatorId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataServicePublishedApis"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDataServicePublishedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of APIs in the published state.
//
// @param request - ListDataServicePublishedApisRequest
//
// @return ListDataServicePublishedApisResponse
func (client *Client) ListDataServicePublishedApis(request *ListDataServicePublishedApisRequest) (_result *ListDataServicePublishedApisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDataServicePublishedApisResponse{}
	_body, _err := client.ListDataServicePublishedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the data sources added to a DataWorks workspace.
//
// @param request - ListDataSourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourcesResponse
func (client *Client) ListDataSourcesWithOptions(request *ListDataSourcesRequest, runtime *dara.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDataSources"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
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
// Queries the data sources added to a DataWorks workspace.
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
// Queries a list of deployment packages. This operation is equivalent to viewing a list of deployment packages on the Deployment Packages page of the DataWorks console.
//
// @param request - ListDeploymentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeploymentsResponse
func (client *Client) ListDeploymentsWithOptions(request *ListDeploymentsRequest, runtime *dara.RuntimeOptions) (_result *ListDeploymentsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		body["Creator"] = request.Creator
	}

	if !dara.IsNil(request.EndCreateTime) {
		body["EndCreateTime"] = request.EndCreateTime
	}

	if !dara.IsNil(request.EndExecuteTime) {
		body["EndExecuteTime"] = request.EndExecuteTime
	}

	if !dara.IsNil(request.Executor) {
		body["Executor"] = request.Executor
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDeployments"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of deployment packages. This operation is equivalent to viewing a list of deployment packages on the Deployment Packages page of the DataWorks console.
//
// @param request - ListDeploymentsRequest
//
// @return ListDeploymentsResponse
func (client *Client) ListDeployments(request *ListDeploymentsRequest) (_result *ListDeploymentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDeploymentsResponse{}
	_body, _err := client.ListDeploymentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of built-in and custom extensions that are enabled in a workspace.
//
// Description:
//
// For information about codes of extension point events, see [Development references: Extension point event codes](https://help.aliyun.com/document_detail/463357.html).
//
// @param request - ListEnabledExtensionsForProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEnabledExtensionsForProjectResponse
func (client *Client) ListEnabledExtensionsForProjectWithOptions(request *ListEnabledExtensionsForProjectRequest, runtime *dara.RuntimeOptions) (_result *ListEnabledExtensionsForProjectResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EventCode) {
		body["EventCode"] = request.EventCode
	}

	if !dara.IsNil(request.FileType) {
		body["FileType"] = request.FileType
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEnabledExtensionsForProject"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEnabledExtensionsForProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of built-in and custom extensions that are enabled in a workspace.
//
// Description:
//
// For information about codes of extension point events, see [Development references: Extension point event codes](https://help.aliyun.com/document_detail/463357.html).
//
// @param request - ListEnabledExtensionsForProjectRequest
//
// @return ListEnabledExtensionsForProjectResponse
func (client *Client) ListEnabledExtensionsForProject(request *ListEnabledExtensionsForProjectRequest) (_result *ListEnabledExtensionsForProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEnabledExtensionsForProjectResponse{}
	_body, _err := client.ListEnabledExtensionsForProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of entities by tag. Only entities of the maxcompute-table type are supported.
//
// @param tmpReq - ListEntitiesByTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntitiesByTagsResponse
func (client *Client) ListEntitiesByTagsWithOptions(tmpReq *ListEntitiesByTagsRequest, runtime *dara.RuntimeOptions) (_result *ListEntitiesByTagsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListEntitiesByTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEntitiesByTags"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEntitiesByTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of entities by tag. Only entities of the maxcompute-table type are supported.
//
// @param request - ListEntitiesByTagsRequest
//
// @return ListEntitiesByTagsResponse
func (client *Client) ListEntitiesByTags(request *ListEntitiesByTagsRequest) (_result *ListEntitiesByTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEntitiesByTagsResponse{}
	_body, _err := client.ListEntitiesByTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of tags of an entity. Only entities of the maxcompute-table type are supported.
//
// @param request - ListEntityTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEntityTagsResponse
func (client *Client) ListEntityTagsWithOptions(request *ListEntityTagsRequest, runtime *dara.RuntimeOptions) (_result *ListEntityTagsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEntityTags"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEntityTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of tags of an entity. Only entities of the maxcompute-table type are supported.
//
// @param request - ListEntityTagsRequest
//
// @return ListEntityTagsResponse
func (client *Client) ListEntityTags(request *ListEntityTagsRequest) (_result *ListEntityTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListEntityTagsResponse{}
	_body, _err := client.ListEntityTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of extensions.
//
// @param request - ListExtensionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExtensionsResponse
func (client *Client) ListExtensionsWithOptions(request *ListExtensionsRequest, runtime *dara.RuntimeOptions) (_result *ListExtensionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExtensions"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExtensionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of extensions.
//
// @param request - ListExtensionsRequest
//
// @return ListExtensionsResponse
func (client *Client) ListExtensions(request *ListExtensionsRequest) (_result *ListExtensionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListExtensionsResponse{}
	_body, _err := client.ListExtensionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about node types, such as the code and name.
//
// @param request - ListFileTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileTypeResponse
func (client *Client) ListFileTypeWithOptions(request *ListFileTypeRequest, runtime *dara.RuntimeOptions) (_result *ListFileTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Locale) {
		body["Locale"] = request.Locale
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFileType"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about node types, such as the code and name.
//
// @param request - ListFileTypeRequest
//
// @return ListFileTypeResponse
func (client *Client) ListFileType(request *ListFileTypeRequest) (_result *ListFileTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFileTypeResponse{}
	_body, _err := client.ListFileTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of file versions.
//
// @param request - ListFileVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileVersionsResponse
func (client *Client) ListFileVersionsWithOptions(request *ListFileVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListFileVersionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFileVersions"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFileVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of file versions.
//
// @param request - ListFileVersionsRequest
//
// @return ListFileVersionsResponse
func (client *Client) ListFileVersions(request *ListFileVersionsRequest) (_result *ListFileVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFileVersionsResponse{}
	_body, _err := client.ListFileVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of files.
//
// @param request - ListFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFilesResponse
func (client *Client) ListFilesWithOptions(request *ListFilesRequest, runtime *dara.RuntimeOptions) (_result *ListFilesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExactFileName) {
		body["ExactFileName"] = request.ExactFileName
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileIdIn) {
		body["FileIdIn"] = request.FileIdIn
	}

	if !dara.IsNil(request.FileTypes) {
		body["FileTypes"] = request.FileTypes
	}

	if !dara.IsNil(request.Keyword) {
		body["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.LastEditUser) {
		body["LastEditUser"] = request.LastEditUser
	}

	if !dara.IsNil(request.NeedAbsoluteFolderPath) {
		body["NeedAbsoluteFolderPath"] = request.NeedAbsoluteFolderPath
	}

	if !dara.IsNil(request.NeedContent) {
		body["NeedContent"] = request.NeedContent
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.UseType) {
		body["UseType"] = request.UseType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFiles"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files.
//
// @param request - ListFilesRequest
//
// @return ListFilesResponse
func (client *Client) ListFiles(request *ListFilesRequest) (_result *ListFilesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFilesResponse{}
	_body, _err := client.ListFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of folders.
//
// @param request - ListFoldersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoldersResponse
func (client *Client) ListFoldersWithOptions(request *ListFoldersRequest, runtime *dara.RuntimeOptions) (_result *ListFoldersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentFolderPath) {
		body["ParentFolderPath"] = request.ParentFolderPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFolders"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFoldersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of folders.
//
// @param request - ListFoldersRequest
//
// @return ListFoldersResponse
func (client *Client) ListFolders(request *ListFoldersRequest) (_result *ListFoldersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListFoldersResponse{}
	_body, _err := client.ListFoldersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about inner nodes. For example, you can call this operation to query the inner nodes of a node group or a do-while node. You cannot call this operation to query the inner nodes of a PAI node.
//
// @param request - ListInnerNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInnerNodesResponse
func (client *Client) ListInnerNodesWithOptions(request *ListInnerNodesRequest, runtime *dara.RuntimeOptions) (_result *ListInnerNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.OuterNodeId) {
		body["OuterNodeId"] = request.OuterNodeId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProgramType) {
		body["ProgramType"] = request.ProgramType
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInnerNodes"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInnerNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about inner nodes. For example, you can call this operation to query the inner nodes of a node group or a do-while node. You cannot call this operation to query the inner nodes of a PAI node.
//
// @param request - ListInnerNodesRequest
//
// @return ListInnerNodesResponse
func (client *Client) ListInnerNodes(request *ListInnerNodesRequest) (_result *ListInnerNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInnerNodesResponse{}
	_body, _err := client.ListInnerNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trend of the number of auto triggered node instances within a specified period of time.
//
// @param request - ListInstanceAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceAmountResponse
func (client *Client) ListInstanceAmountWithOptions(request *ListInstanceAmountRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginDate) {
		body["BeginDate"] = request.BeginDate
	}

	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceAmount"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of the number of auto triggered node instances within a specified period of time.
//
// @param request - ListInstanceAmountRequest
//
// @return ListInstanceAmountResponse
func (client *Client) ListInstanceAmount(request *ListInstanceAmountRequest) (_result *ListInstanceAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceAmountResponse{}
	_body, _err := client.ListInstanceAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the historical records of all instances. One historical record is generated if an instance is rerun once.
//
// @param request - ListInstanceHistoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceHistoryResponse
func (client *Client) ListInstanceHistoryWithOptions(request *ListInstanceHistoryRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceHistoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceHistory"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the historical records of all instances. One historical record is generated if an instance is rerun once.
//
// @param request - ListInstanceHistoryRequest
//
// @return ListInstanceHistoryResponse
func (client *Client) ListInstanceHistory(request *ListInstanceHistoryRequest) (_result *ListInstanceHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceHistoryResponse{}
	_body, _err := client.ListInstanceHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of instances.
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginBizdate) {
		body["BeginBizdate"] = request.BeginBizdate
	}

	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.DagId) {
		body["DagId"] = request.DagId
	}

	if !dara.IsNil(request.EndBizdate) {
		body["EndBizdate"] = request.EndBizdate
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.OrderBy) {
		body["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProgramType) {
		body["ProgramType"] = request.ProgramType
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Status) {
		body["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstances"),
		Version:     dara.String("2020-05-18"),
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
// Queries a list of instances.
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
// Queries the ancestor or descendant lineage of an entity.
//
// @param request - ListLineageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLineageResponse
func (client *Client) ListLineageWithOptions(request *ListLineageRequest, runtime *dara.RuntimeOptions) (_result *ListLineageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EntityQualifiedName) {
		query["EntityQualifiedName"] = request.EntityQualifiedName
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListLineage"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListLineageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ancestor or descendant lineage of an entity.
//
// @param request - ListLineageRequest
//
// @return ListLineageResponse
func (client *Client) ListLineage(request *ListLineageRequest) (_result *ListLineageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListLineageResponse{}
	_body, _err := client.ListLineageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about instances in a manually triggered workflow.
//
// @param request - ListManualDagInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListManualDagInstancesResponse
func (client *Client) ListManualDagInstancesWithOptions(request *ListManualDagInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListManualDagInstancesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DagId) {
		body["DagId"] = request.DagId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListManualDagInstances"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListManualDagInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about instances in a manually triggered workflow.
//
// @param request - ListManualDagInstancesRequest
//
// @return ListManualDagInstancesResponse
func (client *Client) ListManualDagInstances(request *ListManualDagInstancesRequest) (_result *ListManualDagInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListManualDagInstancesResponse{}
	_body, _err := client.ListManualDagInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of phone call-based alerts or text message-based alerts reported within the tenant to which your account belongs during the previous 30 days.
//
// @param request - ListMeasureDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMeasureDataResponse
func (client *Client) ListMeasureDataWithOptions(request *ListMeasureDataRequest, runtime *dara.RuntimeOptions) (_result *ListMeasureDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ComponentCode) {
		query["ComponentCode"] = request.ComponentCode
	}

	if !dara.IsNil(request.DomainCode) {
		query["DomainCode"] = request.DomainCode
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMeasureData"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMeasureDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the statistics on the number of phone call-based alerts or text message-based alerts reported within the tenant to which your account belongs during the previous 30 days.
//
// @param request - ListMeasureDataRequest
//
// @return ListMeasureDataResponse
func (client *Client) ListMeasureData(request *ListMeasureDataRequest) (_result *ListMeasureDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMeasureDataResponse{}
	_body, _err := client.ListMeasureDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the entities in a collection.
//
// @param request - ListMetaCollectionEntitiesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaCollectionEntitiesResponse
func (client *Client) ListMetaCollectionEntitiesWithOptions(request *ListMetaCollectionEntitiesRequest, runtime *dara.RuntimeOptions) (_result *ListMetaCollectionEntitiesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CollectionQualifiedName) {
		query["CollectionQualifiedName"] = request.CollectionQualifiedName
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetaCollectionEntities"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaCollectionEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the entities in a collection.
//
// @param request - ListMetaCollectionEntitiesRequest
//
// @return ListMetaCollectionEntitiesResponse
func (client *Client) ListMetaCollectionEntities(request *ListMetaCollectionEntitiesRequest) (_result *ListMetaCollectionEntitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMetaCollectionEntitiesResponse{}
	_body, _err := client.ListMetaCollectionEntitiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about collections. Collections include data albums that are displayed on the Data Map page and categories that are created in the data albums. You can call this API operation to query collections by type.
//
// Description:
//
// The type can be ALBUM or ALBUM_CATEGORY. ALBUM indicates data albums. ALBUM_CATEGORY indicates categories.
//
// @param request - ListMetaCollectionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaCollectionsResponse
func (client *Client) ListMetaCollectionsWithOptions(request *ListMetaCollectionsRequest, runtime *dara.RuntimeOptions) (_result *ListMetaCollectionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Administrator) {
		query["Administrator"] = request.Administrator
	}

	if !dara.IsNil(request.CollectionType) {
		query["CollectionType"] = request.CollectionType
	}

	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.Follower) {
		query["Follower"] = request.Follower
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OrderBy) {
		query["OrderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentQualifiedName) {
		query["ParentQualifiedName"] = request.ParentQualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetaCollections"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about collections. Collections include data albums that are displayed on the Data Map page and categories that are created in the data albums. You can call this API operation to query collections by type.
//
// Description:
//
// The type can be ALBUM or ALBUM_CATEGORY. ALBUM indicates data albums. ALBUM_CATEGORY indicates categories.
//
// @param request - ListMetaCollectionsRequest
//
// @return ListMetaCollectionsResponse
func (client *Client) ListMetaCollections(request *ListMetaCollectionsRequest) (_result *ListMetaCollectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMetaCollectionsResponse{}
	_body, _err := client.ListMetaCollectionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of metadatabases.
//
// @param request - ListMetaDBRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMetaDBResponse
func (client *Client) ListMetaDBWithOptions(request *ListMetaDBRequest, runtime *dara.RuntimeOptions) (_result *ListMetaDBResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMetaDB"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMetaDBResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of metadatabases.
//
// @param request - ListMetaDBRequest
//
// @return ListMetaDBResponse
func (client *Client) ListMetaDB(request *ListMetaDBRequest) (_result *ListMetaDBResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMetaDBResponse{}
	_body, _err := client.ListMetaDBWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of migration tasks.
//
// @param request - ListMigrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMigrationsResponse
func (client *Client) ListMigrationsWithOptions(request *ListMigrationsRequest, runtime *dara.RuntimeOptions) (_result *ListMigrationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MigrationType) {
		body["MigrationType"] = request.MigrationType
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMigrations"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMigrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of migration tasks.
//
// @param request - ListMigrationsRequest
//
// @return ListMigrationsResponse
func (client *Client) ListMigrations(request *ListMigrationsRequest) (_result *ListMigrationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListMigrationsResponse{}
	_body, _err := client.ListMigrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListNodeIO is deprecated
//
// Summary:
//
// Queries the information about one level of ancestor or descendant nodes of a node.
//
// @param request - ListNodeIORequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeIOResponse
func (client *Client) ListNodeIOWithOptions(request *ListNodeIORequest, runtime *dara.RuntimeOptions) (_result *ListNodeIOResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IoType) {
		body["IoType"] = request.IoType
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeIO"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeIOResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListNodeIO is deprecated
//
// Summary:
//
// Queries the information about one level of ancestor or descendant nodes of a node.
//
// @param request - ListNodeIORequest
//
// @return ListNodeIOResponse
// Deprecated
func (client *Client) ListNodeIO(request *ListNodeIORequest) (_result *ListNodeIOResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodeIOResponse{}
	_body, _err := client.ListNodeIOWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the input and output information about a node. Only the ancestor or descendant nodes at the nearest level can be queried each time.
//
// @param request - ListNodeInputOrOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeInputOrOutputResponse
func (client *Client) ListNodeInputOrOutputWithOptions(request *ListNodeInputOrOutputRequest, runtime *dara.RuntimeOptions) (_result *ListNodeInputOrOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.IoType) {
		body["IoType"] = request.IoType
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeInputOrOutput"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodeInputOrOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the input and output information about a node. Only the ancestor or descendant nodes at the nearest level can be queried each time.
//
// @param request - ListNodeInputOrOutputRequest
//
// @return ListNodeInputOrOutputResponse
func (client *Client) ListNodeInputOrOutput(request *ListNodeInputOrOutputRequest) (_result *ListNodeInputOrOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodeInputOrOutputResponse{}
	_body, _err := client.ListNodeInputOrOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the workspace.
//
// @param request - ListNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesResponse
func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *dara.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizName) {
		body["BizName"] = request.BizName
	}

	if !dara.IsNil(request.NodeName) {
		body["NodeName"] = request.NodeName
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProgramType) {
		body["ProgramType"] = request.ProgramType
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodes"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the workspace.
//
// @param request - ListNodesRequest
//
// @return ListNodesResponse
func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries nodes in a baseline.
//
// @param request - ListNodesByBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesByBaselineResponse
func (client *Client) ListNodesByBaselineWithOptions(request *ListNodesByBaselineRequest, runtime *dara.RuntimeOptions) (_result *ListNodesByBaselineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodesByBaseline"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesByBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries nodes in a baseline.
//
// @param request - ListNodesByBaselineRequest
//
// @return ListNodesByBaselineResponse
func (client *Client) ListNodesByBaseline(request *ListNodesByBaselineRequest) (_result *ListNodesByBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesByBaselineResponse{}
	_body, _err := client.ListNodesByBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries nodes based on the output of the nodes.
//
// @param request - ListNodesByOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodesByOutputResponse
func (client *Client) ListNodesByOutputWithOptions(request *ListNodesByOutputRequest, runtime *dara.RuntimeOptions) (_result *ListNodesByOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Outputs) {
		body["Outputs"] = request.Outputs
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodesByOutput"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListNodesByOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries nodes based on the output of the nodes.
//
// @param request - ListNodesByOutputRequest
//
// @return ListNodesByOutputResponse
func (client *Client) ListNodesByOutput(request *ListNodesByOutputRequest) (_result *ListNodesByOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListNodesByOutputResponse{}
	_body, _err := client.ListNodesByOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of permission request orders.
//
// @param request - ListPermissionApplyOrdersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPermissionApplyOrdersResponse
func (client *Client) ListPermissionApplyOrdersWithOptions(request *ListPermissionApplyOrdersRequest, runtime *dara.RuntimeOptions) (_result *ListPermissionApplyOrdersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApplyType) {
		query["ApplyType"] = request.ApplyType
	}

	if !dara.IsNil(request.CatalogName) {
		query["CatalogName"] = request.CatalogName
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EngineType) {
		query["EngineType"] = request.EngineType
	}

	if !dara.IsNil(request.FlowStatus) {
		query["FlowStatus"] = request.FlowStatus
	}

	if !dara.IsNil(request.MaxComputeProjectName) {
		query["MaxComputeProjectName"] = request.MaxComputeProjectName
	}

	if !dara.IsNil(request.OrderType) {
		query["OrderType"] = request.OrderType
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QueryType) {
		query["QueryType"] = request.QueryType
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPermissionApplyOrders"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPermissionApplyOrdersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of permission request orders.
//
// @param request - ListPermissionApplyOrdersRequest
//
// @return ListPermissionApplyOrdersResponse
func (client *Client) ListPermissionApplyOrders(request *ListPermissionApplyOrdersRequest) (_result *ListPermissionApplyOrdersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListPermissionApplyOrdersResponse{}
	_body, _err := client.ListPermissionApplyOrdersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI ListProgramTypeCount is deprecated
//
// Summary:
//
// Queries the distribution of different types of nodes.
//
// @param request - ListProgramTypeCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProgramTypeCountResponse
func (client *Client) ListProgramTypeCountWithOptions(request *ListProgramTypeCountRequest, runtime *dara.RuntimeOptions) (_result *ListProgramTypeCountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProgramTypeCount"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProgramTypeCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI ListProgramTypeCount is deprecated
//
// Summary:
//
// Queries the distribution of different types of nodes.
//
// @param request - ListProgramTypeCountRequest
//
// @return ListProgramTypeCountResponse
// Deprecated
func (client *Client) ListProgramTypeCount(request *ListProgramTypeCountRequest) (_result *ListProgramTypeCountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProgramTypeCountResponse{}
	_body, _err := client.ListProgramTypeCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the IDs of the workspaces on which a specific Alibaba Cloud account or RAM user has permissions in a specific region.
//
// Description:
//
// An Alibaba Cloud account can assume a role such as the developer, O\\&M engineer, or workspace administrator role in a workspace. For more information, see [Manage members and roles](https://help.aliyun.com/document_detail/136941.html).
//
// @param request - ListProjectIdsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectIdsResponse
func (client *Client) ListProjectIdsWithOptions(request *ListProjectIdsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectIdsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectIds"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectIdsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the IDs of the workspaces on which a specific Alibaba Cloud account or RAM user has permissions in a specific region.
//
// Description:
//
// An Alibaba Cloud account can assume a role such as the developer, O\\&M engineer, or workspace administrator role in a workspace. For more information, see [Manage members and roles](https://help.aliyun.com/document_detail/136941.html).
//
// @param request - ListProjectIdsRequest
//
// @return ListProjectIdsResponse
func (client *Client) ListProjectIds(request *ListProjectIdsRequest) (_result *ListProjectIdsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectIdsResponse{}
	_body, _err := client.ListProjectIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of existing members in a DataWorks workspace.
//
// @param request - ListProjectMembersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembersWithOptions(request *ListProjectMembersRequest, runtime *dara.RuntimeOptions) (_result *ListProjectMembersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectMembers"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of existing members in a DataWorks workspace.
//
// @param request - ListProjectMembersRequest
//
// @return ListProjectMembersResponse
func (client *Client) ListProjectMembers(request *ListProjectMembersRequest) (_result *ListProjectMembersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectMembersResponse{}
	_body, _err := client.ListProjectMembersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of roles in a DataWorks workspace.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=dataworks-public\\&api=ListProjectRoles\\&type=RPC\\&version=2020-05-18)
//
// @param request - ListProjectRolesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRolesWithOptions(request *ListProjectRolesRequest, runtime *dara.RuntimeOptions) (_result *ListProjectRolesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjectRoles"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of roles in a DataWorks workspace.
//
// Description:
//
// ## Debugging
//
// [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=dataworks-public\\&api=ListProjectRoles\\&type=RPC\\&version=2020-05-18)
//
// @param request - ListProjectRolesRequest
//
// @return ListProjectRolesResponse
func (client *Client) ListProjectRoles(request *ListProjectRolesRequest) (_result *ListProjectRolesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectRolesResponse{}
	_body, _err := client.ListProjectRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of DataWorks workspaces of the tenant to which a user belongs.
//
// @param tmpReq - ListProjectsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProjectsResponse
func (client *Client) ListProjectsWithOptions(tmpReq *ListProjectsRequest, runtime *dara.RuntimeOptions) (_result *ListProjectsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListProjects"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of DataWorks workspaces of the tenant to which a user belongs.
//
// @param request - ListProjectsRequest
//
// @return ListProjectsResponse
func (client *Client) ListProjects(request *ListProjectsRequest) (_result *ListProjectsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListProjectsResponse{}
	_body, _err := client.ListProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of historical check results based on a partition filter expression.
//
// Description:
//
// ***
//
// @param request - ListQualityResultsByEntityRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityResultsByEntityResponse
func (client *Client) ListQualityResultsByEntityWithOptions(request *ListQualityResultsByEntityRequest, runtime *dara.RuntimeOptions) (_result *ListQualityResultsByEntityResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityResultsByEntity"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityResultsByEntityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of historical check results based on a partition filter expression.
//
// Description:
//
// ***
//
// @param request - ListQualityResultsByEntityRequest
//
// @return ListQualityResultsByEntityResponse
func (client *Client) ListQualityResultsByEntity(request *ListQualityResultsByEntityRequest) (_result *ListQualityResultsByEntityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityResultsByEntityResponse{}
	_body, _err := client.ListQualityResultsByEntityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries monitoring results after the data quality of a data source or a compute engine is monitored based on monitoring rules.
//
// @param request - ListQualityResultsByRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityResultsByRuleResponse
func (client *Client) ListQualityResultsByRuleWithOptions(request *ListQualityResultsByRuleRequest, runtime *dara.RuntimeOptions) (_result *ListQualityResultsByRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndDate) {
		body["EndDate"] = request.EndDate
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.RuleId) {
		body["RuleId"] = request.RuleId
	}

	if !dara.IsNil(request.StartDate) {
		body["StartDate"] = request.StartDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityResultsByRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityResultsByRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries monitoring results after the data quality of a data source or a compute engine is monitored based on monitoring rules.
//
// @param request - ListQualityResultsByRuleRequest
//
// @return ListQualityResultsByRuleResponse
func (client *Client) ListQualityResultsByRule(request *ListQualityResultsByRuleRequest) (_result *ListQualityResultsByRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityResultsByRuleResponse{}
	_body, _err := client.ListQualityResultsByRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries monitoring rules based on a partition filter expression.
//
// @param request - ListQualityRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQualityRulesResponse
func (client *Client) ListQualityRulesWithOptions(request *ListQualityRulesRequest, runtime *dara.RuntimeOptions) (_result *ListQualityRulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQualityRules"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQualityRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries monitoring rules based on a partition filter expression.
//
// @param request - ListQualityRulesRequest
//
// @return ListQualityRulesResponse
func (client *Client) ListQualityRules(request *ListQualityRulesRequest) (_result *ListQualityRulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListQualityRulesResponse{}
	_body, _err := client.ListQualityRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries synchronization tasks in Data Integration that use a specific data source.
//
// @param request - ListRefDISyncTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRefDISyncTasksResponse
func (client *Client) ListRefDISyncTasksWithOptions(request *ListRefDISyncTasksRequest, runtime *dara.RuntimeOptions) (_result *ListRefDISyncTasksResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasourceName) {
		query["DatasourceName"] = request.DatasourceName
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RefType) {
		query["RefType"] = request.RefType
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRefDISyncTasks"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRefDISyncTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries synchronization tasks in Data Integration that use a specific data source.
//
// @param request - ListRefDISyncTasksRequest
//
// @return ListRefDISyncTasksResponse
func (client *Client) ListRefDISyncTasks(request *ListRefDISyncTasksRequest) (_result *ListRefDISyncTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRefDISyncTasksResponse{}
	_body, _err := client.ListRefDISyncTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of custom alert rules.
//
// @param request - ListRemindsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRemindsResponse
func (client *Client) ListRemindsWithOptions(request *ListRemindsRequest, runtime *dara.RuntimeOptions) (_result *ListRemindsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertTarget) {
		body["AlertTarget"] = request.AlertTarget
	}

	if !dara.IsNil(request.Founder) {
		body["Founder"] = request.Founder
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RemindTypes) {
		body["RemindTypes"] = request.RemindTypes
	}

	if !dara.IsNil(request.SearchText) {
		body["SearchText"] = request.SearchText
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListReminds"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRemindsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of custom alert rules.
//
// @param request - ListRemindsRequest
//
// @return ListRemindsResponse
func (client *Client) ListReminds(request *ListRemindsRequest) (_result *ListRemindsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListRemindsResponse{}
	_body, _err := client.ListRemindsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of resource groups of a specific type.
//
// @param tmpReq - ListResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroupsWithOptions(tmpReq *ListResourceGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceGroupsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListResourceGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BizExtKey) {
		query["BizExtKey"] = request.BizExtKey
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.ResourceGroupType) {
		query["ResourceGroupType"] = request.ResourceGroupType
	}

	if !dara.IsNil(request.ResourceManagerResourceGroupId) {
		query["ResourceManagerResourceGroupId"] = request.ResourceManagerResourceGroupId
	}

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceGroups"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of resource groups of a specific type.
//
// @param request - ListResourceGroupsRequest
//
// @return ListResourceGroupsResponse
func (client *Client) ListResourceGroups(request *ListResourceGroupsRequest) (_result *ListResourceGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceGroupsResponse{}
	_body, _err := client.ListResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of on-duty engineers in a shift schedule.
//
// @param request - ListShiftPersonnelsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShiftPersonnelsResponse
func (client *Client) ListShiftPersonnelsWithOptions(request *ListShiftPersonnelsRequest, runtime *dara.RuntimeOptions) (_result *ListShiftPersonnelsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		body["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.ShiftPersonUID) {
		body["ShiftPersonUID"] = request.ShiftPersonUID
	}

	if !dara.IsNil(request.ShiftScheduleIdentifier) {
		body["ShiftScheduleIdentifier"] = request.ShiftScheduleIdentifier
	}

	if !dara.IsNil(request.UserType) {
		body["UserType"] = request.UserType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListShiftPersonnels"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListShiftPersonnelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of on-duty engineers in a shift schedule.
//
// @param request - ListShiftPersonnelsRequest
//
// @return ListShiftPersonnelsResponse
func (client *Client) ListShiftPersonnels(request *ListShiftPersonnelsRequest) (_result *ListShiftPersonnelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListShiftPersonnelsResponse{}
	_body, _err := client.ListShiftPersonnelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of shift schedules in Operation Center.
//
// @param request - ListShiftSchedulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShiftSchedulesResponse
func (client *Client) ListShiftSchedulesWithOptions(request *ListShiftSchedulesRequest, runtime *dara.RuntimeOptions) (_result *ListShiftSchedulesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ShiftScheduleName) {
		body["ShiftScheduleName"] = request.ShiftScheduleName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListShiftSchedules"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListShiftSchedulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of shift schedules in Operation Center.
//
// @param request - ListShiftSchedulesRequest
//
// @return ListShiftSchedulesResponse
func (client *Client) ListShiftSchedules(request *ListShiftSchedulesRequest) (_result *ListShiftSchedulesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListShiftSchedulesResponse{}
	_body, _err := client.ListShiftSchedulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the trend of the number of auto triggered node instances that are successfully run every hour on the hour of the current day.
//
// @param request - ListSuccessInstanceAmountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSuccessInstanceAmountResponse
func (client *Client) ListSuccessInstanceAmountWithOptions(request *ListSuccessInstanceAmountRequest, runtime *dara.RuntimeOptions) (_result *ListSuccessInstanceAmountResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSuccessInstanceAmount"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSuccessInstanceAmountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the trend of the number of auto triggered node instances that are successfully run every hour on the hour of the current day.
//
// @param request - ListSuccessInstanceAmountRequest
//
// @return ListSuccessInstanceAmountResponse
func (client *Client) ListSuccessInstanceAmount(request *ListSuccessInstanceAmountRequest) (_result *ListSuccessInstanceAmountResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSuccessInstanceAmountResponse{}
	_body, _err := client.ListSuccessInstanceAmountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of table levels. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - ListTableLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableLevelResponse
func (client *Client) ListTableLevelWithOptions(request *ListTableLevelRequest, runtime *dara.RuntimeOptions) (_result *ListTableLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableLevel"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of table levels. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - ListTableLevelRequest
//
// @return ListTableLevelResponse
func (client *Client) ListTableLevel(request *ListTableLevelRequest) (_result *ListTableLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTableLevelResponse{}
	_body, _err := client.ListTableLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of table themes. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - ListTableThemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTableThemeResponse
func (client *Client) ListTableThemeWithOptions(request *ListTableThemeRequest, runtime *dara.RuntimeOptions) (_result *ListTableThemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTableTheme"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTableThemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of table themes. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - ListTableThemeRequest
//
// @return ListTableThemeResponse
func (client *Client) ListTableTheme(request *ListTableThemeRequest) (_result *ListTableThemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTableThemeResponse{}
	_body, _err := client.ListTableThemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains tables of different data source types within a tenant by page.
//
// @param request - ListTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTablesResponse
func (client *Client) ListTablesWithOptions(request *ListTablesRequest, runtime *dara.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTables"),
		Version:     dara.String("2020-05-18"),
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
// Obtains tables of different data source types within a tenant by page.
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
// Queries events.
//
// @param request - ListTopicsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTopicsResponse
func (client *Client) ListTopicsWithOptions(request *ListTopicsRequest, runtime *dara.RuntimeOptions) (_result *ListTopicsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BeginTime) {
		body["BeginTime"] = request.BeginTime
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.TopicStatuses) {
		body["TopicStatuses"] = request.TopicStatuses
	}

	if !dara.IsNil(request.TopicTypes) {
		body["TopicTypes"] = request.TopicTypes
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTopics"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTopicsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries events.
//
// @param request - ListTopicsRequest
//
// @return ListTopicsResponse
func (client *Client) ListTopics(request *ListTopicsRequest) (_result *ListTopicsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTopicsResponse{}
	_body, _err := client.ListTopicsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Undeploys a node.
//
// @param request - OfflineNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineNodeResponse
func (client *Client) OfflineNodeWithOptions(request *OfflineNodeRequest, runtime *dara.RuntimeOptions) (_result *OfflineNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineNode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Undeploys a node.
//
// @param request - OfflineNodeRequest
//
// @return OfflineNodeResponse
func (client *Client) OfflineNode(request *OfflineNodeRequest) (_result *OfflineNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &OfflineNodeResponse{}
	_body, _err := client.OfflineNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes an API.
//
// @param request - PublishDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApiWithOptions(request *PublishDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *PublishDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PublishDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes an API.
//
// @param request - PublishDataServiceApiRequest
//
// @return PublishDataServiceApiResponse
func (client *Client) PublishDataServiceApi(request *PublishDataServiceApiRequest) (_result *PublishDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PublishDataServiceApiResponse{}
	_body, _err := client.PublishDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution results of an asynchronous task.
//
// Description:
//
// DataWorks allows you to call only the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) operation to create a batch synchronization task or the [UpdateDISyncTask](https://help.aliyun.com/document_detail/289109.html) operation to update a batch synchronization task in Data Integration. To create or update a real-time synchronization task, you must first call the [GenerateDISyncTaskConfigForCreating](https://help.aliyun.com/document_detail/383463.html) or [GenerateDISyncTaskConfigForUpdating](https://help.aliyun.com/document_detail/383464.html) operation to obtain the ID of an asynchronous thread and call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. Then, you can call the CreateDISyncTask or UpdateDISyncTask operation and use the parameters as request parameters to create or update a real-time synchronization task. DataWorks allows you to create or update real-time synchronization tasks in Data Integration only in asynchronous mode.
//
// @param request - QueryDISyncTaskConfigProcessResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDISyncTaskConfigProcessResultResponse
func (client *Client) QueryDISyncTaskConfigProcessResultWithOptions(request *QueryDISyncTaskConfigProcessResultRequest, runtime *dara.RuntimeOptions) (_result *QueryDISyncTaskConfigProcessResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsyncProcessId) {
		query["AsyncProcessId"] = request.AsyncProcessId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDISyncTaskConfigProcessResult"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDISyncTaskConfigProcessResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution results of an asynchronous task.
//
// Description:
//
// DataWorks allows you to call only the [CreateDISyncTask](https://help.aliyun.com/document_detail/278725.html) operation to create a batch synchronization task or the [UpdateDISyncTask](https://help.aliyun.com/document_detail/289109.html) operation to update a batch synchronization task in Data Integration. To create or update a real-time synchronization task, you must first call the [GenerateDISyncTaskConfigForCreating](https://help.aliyun.com/document_detail/383463.html) or [GenerateDISyncTaskConfigForUpdating](https://help.aliyun.com/document_detail/383464.html) operation to obtain the ID of an asynchronous thread and call the [QueryDISyncTaskConfigProcessResult](https://help.aliyun.com/document_detail/383465.html) operation to obtain the asynchronously generated parameters based on the ID. Then, you can call the CreateDISyncTask or UpdateDISyncTask operation and use the parameters as request parameters to create or update a real-time synchronization task. DataWorks allows you to create or update real-time synchronization tasks in Data Integration only in asynchronous mode.
//
// @param request - QueryDISyncTaskConfigProcessResultRequest
//
// @return QueryDISyncTaskConfigProcessResultResponse
func (client *Client) QueryDISyncTaskConfigProcessResult(request *QueryDISyncTaskConfigProcessResultRequest) (_result *QueryDISyncTaskConfigProcessResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDISyncTaskConfigProcessResultResponse{}
	_body, _err := client.QueryDISyncTaskConfigProcessResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default data category and data sensitivity level template defined by Data Security Guard.
//
// @param request - QueryDefaultTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryDefaultTemplateResponse
func (client *Client) QueryDefaultTemplateWithOptions(request *QueryDefaultTemplateRequest, runtime *dara.RuntimeOptions) (_result *QueryDefaultTemplateResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryDefaultTemplate"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryDefaultTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default data category and data sensitivity level template defined by Data Security Guard.
//
// @param request - QueryDefaultTemplateRequest
//
// @return QueryDefaultTemplateResponse
func (client *Client) QueryDefaultTemplate(request *QueryDefaultTemplateRequest) (_result *QueryDefaultTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryDefaultTemplateResponse{}
	_body, _err := client.QueryDefaultTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about objects that are created in Data Modeling by using fast modeling language (FML) statements.
//
// Description:
//
//	  Each time you call this API operation, you must use FML statements to query information about objects that are created in Data Modeling.
//
//		- The information about the objects can be queried by page, except for data layers, business processes, and data domains. You can add an offset to the end of an FML statement. The num LIMIT num statement specifies the offset when the information about the objects is queried, and the number of pages to return each time. The offset value must be a multiple of the number of pages.
//
//		- A maximum of 1,000 entries can be returned each time you call this API operation.
//
// @param request - QueryPublicModelEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryPublicModelEngineResponse
func (client *Client) QueryPublicModelEngineWithOptions(request *QueryPublicModelEngineRequest, runtime *dara.RuntimeOptions) (_result *QueryPublicModelEngineResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Text) {
		body["Text"] = request.Text
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryPublicModelEngine"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryPublicModelEngineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about objects that are created in Data Modeling by using fast modeling language (FML) statements.
//
// Description:
//
//	  Each time you call this API operation, you must use FML statements to query information about objects that are created in Data Modeling.
//
//		- The information about the objects can be queried by page, except for data layers, business processes, and data domains. You can add an offset to the end of an FML statement. The num LIMIT num statement specifies the offset when the information about the objects is queried, and the number of pages to return each time. The offset value must be a multiple of the number of pages.
//
//		- A maximum of 1,000 entries can be returned each time you call this API operation.
//
// @param request - QueryPublicModelEngineRequest
//
// @return QueryPublicModelEngineResponse
func (client *Client) QueryPublicModelEngine(request *QueryPublicModelEngineRequest) (_result *QueryPublicModelEngineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryPublicModelEngineResponse{}
	_body, _err := client.QueryPublicModelEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the type of a sensitive data identification rule.
//
// @param request - QueryRecognizeDataByRuleTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecognizeDataByRuleTypeResponse
func (client *Client) QueryRecognizeDataByRuleTypeWithOptions(request *QueryRecognizeDataByRuleTypeRequest, runtime *dara.RuntimeOptions) (_result *QueryRecognizeDataByRuleTypeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RecognizeRulesType) {
		body["RecognizeRulesType"] = request.RecognizeRulesType
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecognizeDataByRuleType"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecognizeDataByRuleTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the type of a sensitive data identification rule.
//
// @param request - QueryRecognizeDataByRuleTypeRequest
//
// @return QueryRecognizeDataByRuleTypeResponse
func (client *Client) QueryRecognizeDataByRuleType(request *QueryRecognizeDataByRuleTypeRequest) (_result *QueryRecognizeDataByRuleTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecognizeDataByRuleTypeResponse{}
	_body, _err := client.QueryRecognizeDataByRuleTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specified sensitive field in Data Security Guard.
//
// @param request - QueryRecognizeRuleDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecognizeRuleDetailResponse
func (client *Client) QueryRecognizeRuleDetailWithOptions(request *QueryRecognizeRuleDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryRecognizeRuleDetailResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SensitiveName) {
		body["SensitiveName"] = request.SensitiveName
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecognizeRuleDetail"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecognizeRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specified sensitive field in Data Security Guard.
//
// @param request - QueryRecognizeRuleDetailRequest
//
// @return QueryRecognizeRuleDetailResponse
func (client *Client) QueryRecognizeRuleDetail(request *QueryRecognizeRuleDetailRequest) (_result *QueryRecognizeRuleDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecognizeRuleDetailResponse{}
	_body, _err := client.QueryRecognizeRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the built-in sensitive data identification rule that is used to configure a sensitive field.
//
// @param request - QueryRecognizeRulesTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRecognizeRulesTypeResponse
func (client *Client) QueryRecognizeRulesTypeWithOptions(runtime *dara.RuntimeOptions) (_result *QueryRecognizeRulesTypeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRecognizeRulesType"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRecognizeRulesTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the built-in sensitive data identification rule that is used to configure a sensitive field.
//
// @return QueryRecognizeRulesTypeResponse
func (client *Client) QueryRecognizeRulesType() (_result *QueryRecognizeRulesTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRecognizeRulesTypeResponse{}
	_body, _err := client.QueryRecognizeRulesTypeWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data categories.
//
// @param request - QuerySensClassificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySensClassificationResponse
func (client *Client) QuerySensClassificationWithOptions(request *QuerySensClassificationRequest, runtime *dara.RuntimeOptions) (_result *QuerySensClassificationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySensClassification"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySensClassificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data categories.
//
// @param request - QuerySensClassificationRequest
//
// @return QuerySensClassificationResponse
func (client *Client) QuerySensClassification(request *QuerySensClassificationRequest) (_result *QuerySensClassificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySensClassificationResponse{}
	_body, _err := client.QuerySensClassificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries data sensitivity levels in Data Security Guard.
//
// @param request - QuerySensLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySensLevelResponse
func (client *Client) QuerySensLevelWithOptions(request *QuerySensLevelRequest, runtime *dara.RuntimeOptions) (_result *QuerySensLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["tenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySensLevel"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySensLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries data sensitivity levels in Data Security Guard.
//
// @param request - QuerySensLevelRequest
//
// @return QuerySensLevelResponse
func (client *Client) QuerySensLevel(request *QuerySensLevelRequest) (_result *QuerySensLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySensLevelResponse{}
	_body, _err := client.QuerySensLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries sensitive data identification rules.
//
// @param request - QuerySensNodeInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySensNodeInfoResponse
func (client *Client) QuerySensNodeInfoWithOptions(request *QuerySensNodeInfoRequest, runtime *dara.RuntimeOptions) (_result *QuerySensNodeInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.PageNo) {
		body["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SensitiveName) {
		body["SensitiveName"] = request.SensitiveName
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySensNodeInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySensNodeInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries sensitive data identification rules.
//
// @param request - QuerySensNodeInfoRequest
//
// @return QuerySensNodeInfoResponse
func (client *Client) QuerySensNodeInfo(request *QuerySensNodeInfoRequest) (_result *QuerySensNodeInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuerySensNodeInfoResponse{}
	_body, _err := client.QuerySensNodeInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers the lineage between self-managed entities to DataWorks.
//
// Description:
//
// This operation is in the trial phase. Users who need to call this operation can apply for it. The users can call this operation after the administrator adds the users to the trial list.
//
// @param tmpReq - RegisterLineageRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterLineageRelationResponse
func (client *Client) RegisterLineageRelationWithOptions(tmpReq *RegisterLineageRelationRequest, runtime *dara.RuntimeOptions) (_result *RegisterLineageRelationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RegisterLineageRelationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.LineageRelationRegisterVO) {
		request.LineageRelationRegisterVOShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LineageRelationRegisterVO, dara.String("LineageRelationRegisterVO"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.LineageRelationRegisterVOShrink) {
		body["LineageRelationRegisterVO"] = request.LineageRelationRegisterVOShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterLineageRelation"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterLineageRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers the lineage between self-managed entities to DataWorks.
//
// Description:
//
// This operation is in the trial phase. Users who need to call this operation can apply for it. The users can call this operation after the administrator adds the users to the trial list.
//
// @param request - RegisterLineageRelationRequest
//
// @return RegisterLineageRelationResponse
func (client *Client) RegisterLineageRelation(request *RegisterLineageRelationRequest) (_result *RegisterLineageRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterLineageRelationResponse{}
	_body, _err := client.RegisterLineageRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from an entity. Only entities of the maxcompute-table type are supported.
//
// @param tmpReq - RemoveEntityTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveEntityTagsResponse
func (client *Client) RemoveEntityTagsWithOptions(tmpReq *RemoveEntityTagsRequest, runtime *dara.RuntimeOptions) (_result *RemoveEntityTagsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RemoveEntityTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TagKeys) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, dara.String("TagKeys"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.QualifiedName) {
		query["QualifiedName"] = request.QualifiedName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TagKeysShrink) {
		body["TagKeys"] = request.TagKeysShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveEntityTags"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveEntityTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from an entity. Only entities of the maxcompute-table type are supported.
//
// @param request - RemoveEntityTagsRequest
//
// @return RemoveEntityTagsResponse
func (client *Client) RemoveEntityTags(request *RemoveEntityTagsRequest) (_result *RemoveEntityTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveEntityTagsResponse{}
	_body, _err := client.RemoveEntityTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a role from a user in a DataWorks workspace.
//
// @param request - RemoveProjectMemberFromRoleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveProjectMemberFromRoleResponse
func (client *Client) RemoveProjectMemberFromRoleWithOptions(request *RemoveProjectMemberFromRoleRequest, runtime *dara.RuntimeOptions) (_result *RemoveProjectMemberFromRoleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RoleCode) {
		query["RoleCode"] = request.RoleCode
	}

	if !dara.IsNil(request.UserId) {
		query["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RemoveProjectMemberFromRole"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RemoveProjectMemberFromRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a role from a user in a DataWorks workspace.
//
// @param request - RemoveProjectMemberFromRoleRequest
//
// @return RemoveProjectMemberFromRoleResponse
func (client *Client) RemoveProjectMemberFromRole(request *RemoveProjectMemberFromRoleRequest) (_result *RemoveProjectMemberFromRoleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RemoveProjectMemberFromRoleResponse{}
	_body, _err := client.RemoveProjectMemberFromRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts an instance.
//
// @param request - RestartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartInstanceResponse
func (client *Client) RestartInstanceWithOptions(request *RestartInstanceRequest, runtime *dara.RuntimeOptions) (_result *RestartInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RestartInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RestartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts an instance.
//
// @param request - RestartInstanceRequest
//
// @return RestartInstanceResponse
func (client *Client) RestartInstance(request *RestartInstanceRequest) (_result *RestartInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RestartInstanceResponse{}
	_body, _err := client.RestartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a suspended instance.
//
// @param request - ResumeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstanceWithOptions(request *ResumeInstanceRequest, runtime *dara.RuntimeOptions) (_result *ResumeInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a suspended instance.
//
// @param request - ResumeInstanceRequest
//
// @return ResumeInstanceResponse
func (client *Client) ResumeInstance(request *ResumeInstanceRequest) (_result *ResumeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeInstanceResponse{}
	_body, _err := client.ResumeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Revokes permissions on a table from a user.
//
// @param request - RevokeTablePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeTablePermissionResponse
func (client *Client) RevokeTablePermissionWithOptions(request *RevokeTablePermissionRequest, runtime *dara.RuntimeOptions) (_result *RevokeTablePermissionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Actions) {
		query["Actions"] = request.Actions
	}

	if !dara.IsNil(request.MaxComputeProjectName) {
		query["MaxComputeProjectName"] = request.MaxComputeProjectName
	}

	if !dara.IsNil(request.RevokeUserId) {
		query["RevokeUserId"] = request.RevokeUserId
	}

	if !dara.IsNil(request.RevokeUserName) {
		query["RevokeUserName"] = request.RevokeUserName
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RevokeTablePermission"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RevokeTablePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Revokes permissions on a table from a user.
//
// @param request - RevokeTablePermissionRequest
//
// @return RevokeTablePermissionResponse
func (client *Client) RevokeTablePermission(request *RevokeTablePermissionRequest) (_result *RevokeTablePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RevokeTablePermissionResponse{}
	_body, _err := client.RevokeTablePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow to backfill data.
//
// Description:
//
// For more information about data backfill, see [Backfill data](https://help.aliyun.com/document_detail/137937.html).
//
// @param request - RunCycleDagNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCycleDagNodesResponse
func (client *Client) RunCycleDagNodesWithOptions(request *RunCycleDagNodesRequest, runtime *dara.RuntimeOptions) (_result *RunCycleDagNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertNoticeType) {
		body["AlertNoticeType"] = request.AlertNoticeType
	}

	if !dara.IsNil(request.AlertType) {
		body["AlertType"] = request.AlertType
	}

	if !dara.IsNil(request.BizBeginTime) {
		body["BizBeginTime"] = request.BizBeginTime
	}

	if !dara.IsNil(request.BizEndTime) {
		body["BizEndTime"] = request.BizEndTime
	}

	if !dara.IsNil(request.ConcurrentRuns) {
		body["ConcurrentRuns"] = request.ConcurrentRuns
	}

	if !dara.IsNil(request.EndBizDate) {
		body["EndBizDate"] = request.EndBizDate
	}

	if !dara.IsNil(request.ExcludeNodeIds) {
		body["ExcludeNodeIds"] = request.ExcludeNodeIds
	}

	if !dara.IsNil(request.IncludeNodeIds) {
		body["IncludeNodeIds"] = request.IncludeNodeIds
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeParams) {
		body["NodeParams"] = request.NodeParams
	}

	if !dara.IsNil(request.Parallelism) {
		body["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.RootNodeId) {
		body["RootNodeId"] = request.RootNodeId
	}

	if !dara.IsNil(request.StartBizDate) {
		body["StartBizDate"] = request.StartBizDate
	}

	if !dara.IsNil(request.StartFutureInstanceImmediately) {
		body["StartFutureInstanceImmediately"] = request.StartFutureInstanceImmediately
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCycleDagNodes"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCycleDagNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow to backfill data.
//
// Description:
//
// For more information about data backfill, see [Backfill data](https://help.aliyun.com/document_detail/137937.html).
//
// @param request - RunCycleDagNodesRequest
//
// @return RunCycleDagNodesResponse
func (client *Client) RunCycleDagNodes(request *RunCycleDagNodesRequest) (_result *RunCycleDagNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunCycleDagNodesResponse{}
	_body, _err := client.RunCycleDagNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs nodes in a manually triggered workflow. Before you call this operation, make sure that the manually triggered workflow is committed and deployed. You can find a manually triggered workflow in Operation Center only after the manually triggered workflow is committed and deployed.
//
// @param request - RunManualDagNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunManualDagNodesResponse
func (client *Client) RunManualDagNodesWithOptions(request *RunManualDagNodesRequest, runtime *dara.RuntimeOptions) (_result *RunManualDagNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.DagParameters) {
		body["DagParameters"] = request.DagParameters
	}

	if !dara.IsNil(request.EndBizDate) {
		body["EndBizDate"] = request.EndBizDate
	}

	if !dara.IsNil(request.ExcludeNodeIds) {
		body["ExcludeNodeIds"] = request.ExcludeNodeIds
	}

	if !dara.IsNil(request.FlowName) {
		body["FlowName"] = request.FlowName
	}

	if !dara.IsNil(request.IncludeNodeIds) {
		body["IncludeNodeIds"] = request.IncludeNodeIds
	}

	if !dara.IsNil(request.NodeParameters) {
		body["NodeParameters"] = request.NodeParameters
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.StartBizDate) {
		body["StartBizDate"] = request.StartBizDate
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunManualDagNodes"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunManualDagNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs nodes in a manually triggered workflow. Before you call this operation, make sure that the manually triggered workflow is committed and deployed. You can find a manually triggered workflow in Operation Center only after the manually triggered workflow is committed and deployed.
//
// @param request - RunManualDagNodesRequest
//
// @return RunManualDagNodesResponse
func (client *Client) RunManualDagNodes(request *RunManualDagNodesRequest) (_result *RunManualDagNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunManualDagNodesResponse{}
	_body, _err := client.RunManualDagNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow to perform smoke testing.
//
// @param request - RunSmokeTestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunSmokeTestResponse
func (client *Client) RunSmokeTestWithOptions(request *RunSmokeTestRequest, runtime *dara.RuntimeOptions) (_result *RunSmokeTestResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Bizdate) {
		body["Bizdate"] = request.Bizdate
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.NodeParams) {
		body["NodeParams"] = request.NodeParams
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunSmokeTest"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunSmokeTestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow to perform smoke testing.
//
// @param request - RunSmokeTestRequest
//
// @return RunSmokeTestResponse
func (client *Client) RunSmokeTest(request *RunSmokeTestRequest) (_result *RunSmokeTestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunSmokeTestResponse{}
	_body, _err := client.RunSmokeTestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs a manually triggered node.
//
// @param request - RunTriggerNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunTriggerNodeResponse
func (client *Client) RunTriggerNodeWithOptions(request *RunTriggerNodeRequest, runtime *dara.RuntimeOptions) (_result *RunTriggerNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["AppId"] = request.AppId
	}

	if !dara.IsNil(request.BizDate) {
		body["BizDate"] = request.BizDate
	}

	if !dara.IsNil(request.CycleTime) {
		body["CycleTime"] = request.CycleTime
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunTriggerNode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RunTriggerNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a manually triggered node.
//
// @param request - RunTriggerNodeRequest
//
// @return RunTriggerNodeResponse
func (client *Client) RunTriggerNode(request *RunTriggerNodeRequest) (_result *RunTriggerNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RunTriggerNodeResponse{}
	_body, _err := client.RunTriggerNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Saves the test results of an API.
//
// @param request - SaveDataServiceApiTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveDataServiceApiTestResultResponse
func (client *Client) SaveDataServiceApiTestResultWithOptions(request *SaveDataServiceApiTestResultRequest, runtime *dara.RuntimeOptions) (_result *SaveDataServiceApiTestResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.AutoGenerate) {
		body["AutoGenerate"] = request.AutoGenerate
	}

	if !dara.IsNil(request.FailResultSample) {
		body["FailResultSample"] = request.FailResultSample
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResultSample) {
		body["ResultSample"] = request.ResultSample
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveDataServiceApiTestResult"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveDataServiceApiTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Saves the test results of an API.
//
// @param request - SaveDataServiceApiTestResultRequest
//
// @return SaveDataServiceApiTestResultResponse
func (client *Client) SaveDataServiceApiTestResult(request *SaveDataServiceApiTestResultRequest) (_result *SaveDataServiceApiTestResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SaveDataServiceApiTestResultResponse{}
	_body, _err := client.SaveDataServiceApiTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether input data contains sensitive data.
//
// @param request - ScanSensitiveDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanSensitiveDataResponse
func (client *Client) ScanSensitiveDataWithOptions(request *ScanSensitiveDataRequest, runtime *dara.RuntimeOptions) (_result *ScanSensitiveDataResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ScanSensitiveData"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ScanSensitiveDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether input data contains sensitive data.
//
// @param request - ScanSensitiveDataRequest
//
// @return ScanSensitiveDataResponse
func (client *Client) ScanSensitiveData(request *ScanSensitiveDataRequest) (_result *ScanSensitiveDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ScanSensitiveDataResponse{}
	_body, _err := client.ScanSensitiveDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries metatables based on specific conditions.
//
// Description:
//
// You can call this operation to query only metatables in a MaxCompute or E-MapReduce (EMR) compute engine.
//
// @param request - SearchMetaTablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMetaTablesResponse
func (client *Client) SearchMetaTablesWithOptions(request *SearchMetaTablesRequest, runtime *dara.RuntimeOptions) (_result *SearchMetaTablesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGuid) {
		query["AppGuid"] = request.AppGuid
	}

	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DataSourceType) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Keyword) {
		query["Keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMetaTables"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchMetaTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries metatables based on specific conditions.
//
// Description:
//
// You can call this operation to query only metatables in a MaxCompute or E-MapReduce (EMR) compute engine.
//
// @param request - SearchMetaTablesRequest
//
// @return SearchMetaTablesResponse
func (client *Client) SearchMetaTables(request *SearchMetaTablesRequest) (_result *SearchMetaTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchMetaTablesResponse{}
	_body, _err := client.SearchMetaTablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI SearchNodesByOutput is deprecated
//
// Summary:
//
// Queries a node based on the output.
//
// @param request - SearchNodesByOutputRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchNodesByOutputResponse
func (client *Client) SearchNodesByOutputWithOptions(request *SearchNodesByOutputRequest, runtime *dara.RuntimeOptions) (_result *SearchNodesByOutputResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Outputs) {
		body["Outputs"] = request.Outputs
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchNodesByOutput"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchNodesByOutputResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI SearchNodesByOutput is deprecated
//
// Summary:
//
// Queries a node based on the output.
//
// @param request - SearchNodesByOutputRequest
//
// @return SearchNodesByOutputResponse
// Deprecated
func (client *Client) SearchNodesByOutput(request *SearchNodesByOutputRequest) (_result *SearchNodesByOutputResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SearchNodesByOutputResponse{}
	_body, _err := client.SearchNodesByOutputWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI SetDataSourceShare is deprecated
//
// Summary:
//
// Shares a data source to a specific DataWorks workspace or a specific user.
//
// @param request - SetDataSourceShareRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDataSourceShareResponse
func (client *Client) SetDataSourceShareWithOptions(request *SetDataSourceShareRequest, runtime *dara.RuntimeOptions) (_result *SetDataSourceShareResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasourceName) {
		query["DatasourceName"] = request.DatasourceName
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectPermissions) {
		query["ProjectPermissions"] = request.ProjectPermissions
	}

	if !dara.IsNil(request.UserPermissions) {
		query["UserPermissions"] = request.UserPermissions
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDataSourceShare"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDataSourceShareResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI SetDataSourceShare is deprecated
//
// Summary:
//
// Shares a data source to a specific DataWorks workspace or a specific user.
//
// @param request - SetDataSourceShareRequest
//
// @return SetDataSourceShareResponse
// Deprecated
func (client *Client) SetDataSourceShare(request *SetDataSourceShareRequest) (_result *SetDataSourceShareResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDataSourceShareResponse{}
	_body, _err := client.SetDataSourceShareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Configures tags for an entity. Only entities of the maxcompute-table type are supported.
//
// @param tmpReq - SetEntityTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetEntityTagsResponse
func (client *Client) SetEntityTagsWithOptions(tmpReq *SetEntityTagsRequest, runtime *dara.RuntimeOptions) (_result *SetEntityTagsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SetEntityTagsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.QualifiedName) {
		query["QualifiedName"] = request.QualifiedName
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TagsShrink) {
		body["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetEntityTags"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetEntityTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Configures tags for an entity. Only entities of the maxcompute-table type are supported.
//
// @param request - SetEntityTagsRequest
//
// @return SetEntityTagsResponse
func (client *Client) SetEntityTags(request *SetEntityTagsRequest) (_result *SetEntityTagsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetEntityTagsResponse{}
	_body, _err := client.SetEntityTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the state of a failed instance to successful.
//
// @param request - SetSuccessInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetSuccessInstanceResponse
func (client *Client) SetSuccessInstanceWithOptions(request *SetSuccessInstanceRequest, runtime *dara.RuntimeOptions) (_result *SetSuccessInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetSuccessInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetSuccessInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the state of a failed instance to successful.
//
// @param request - SetSuccessInstanceRequest
//
// @return SetSuccessInstanceResponse
func (client *Client) SetSuccessInstance(request *SetSuccessInstanceRequest) (_result *SetSuccessInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetSuccessInstanceResponse{}
	_body, _err := client.SetSuccessInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a synchronization task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// @param tmpReq - StartDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDIJobResponse
func (client *Client) StartDIJobWithOptions(tmpReq *StartDIJobRequest, runtime *dara.RuntimeOptions) (_result *StartDIJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StartDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RealtimeStartSettings) {
		request.RealtimeStartSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RealtimeStartSettings, dara.String("RealtimeStartSettings"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.ForceToRerun) {
		body["ForceToRerun"] = request.ForceToRerun
	}

	if !dara.IsNil(request.RealtimeStartSettingsShrink) {
		body["RealtimeStartSettings"] = request.RealtimeStartSettingsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDIJob"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a synchronization task of a new version. Only the following type of task is supported: real-time data synchronization from a MySQL database to Hologres.
//
// @param request - StartDIJobRequest
//
// @return StartDIJobResponse
func (client *Client) StartDIJob(request *StartDIJobRequest) (_result *StartDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDIJobResponse{}
	_body, _err := client.StartDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a real-time synchronization task or a synchronization solution.
//
// @param request - StartDISyncInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDISyncInstanceResponse
func (client *Client) StartDISyncInstanceWithOptions(request *StartDISyncInstanceRequest, runtime *dara.RuntimeOptions) (_result *StartDISyncInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StartParam) {
		query["StartParam"] = request.StartParam
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartDISyncInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartDISyncInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a real-time synchronization task or a synchronization solution.
//
// @param request - StartDISyncInstanceRequest
//
// @return StartDISyncInstanceResponse
func (client *Client) StartDISyncInstance(request *StartDISyncInstanceRequest) (_result *StartDISyncInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartDISyncInstanceResponse{}
	_body, _err := client.StartDISyncInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a migration task.
//
// @param request - StartMigrationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMigrationResponse
func (client *Client) StartMigrationWithOptions(request *StartMigrationRequest, runtime *dara.RuntimeOptions) (_result *StartMigrationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MigrationId) {
		body["MigrationId"] = request.MigrationId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartMigration"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a migration task.
//
// @param request - StartMigrationRequest
//
// @return StartMigrationResponse
func (client *Client) StartMigration(request *StartMigrationRequest) (_result *StartMigrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartMigrationResponse{}
	_body, _err := client.StartMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a new-version synchronization task. The following type of synchronization task is supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// @param request - StopDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDIJobResponse
func (client *Client) StopDIJobWithOptions(request *StopDIJobRequest, runtime *dara.RuntimeOptions) (_result *StopDIJobResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDIJob"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a new-version synchronization task. The following type of synchronization task is supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// @param request - StopDIJobRequest
//
// @return StopDIJobResponse
func (client *Client) StopDIJob(request *StopDIJobRequest) (_result *StopDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDIJobResponse{}
	_body, _err := client.StopDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a real-time synchronization task.
//
// @param request - StopDISyncInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDISyncInstanceResponse
func (client *Client) StopDISyncInstanceWithOptions(request *StopDISyncInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopDISyncInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopDISyncInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopDISyncInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a real-time synchronization task.
//
// @param request - StopDISyncInstanceRequest
//
// @return StopDISyncInstanceResponse
func (client *Client) StopDISyncInstance(request *StopDISyncInstanceRequest) (_result *StopDISyncInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopDISyncInstanceResponse{}
	_body, _err := client.StopDISyncInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates an instance.
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *dara.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates an instance.
//
// @param request - StopInstanceRequest
//
// @return StopInstanceResponse
func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopInstanceResponse{}
	_body, _err := client.StopInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits an API in DataService Studio.
//
// @param request - SubmitDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDataServiceApiResponse
func (client *Client) SubmitDataServiceApiWithOptions(request *SubmitDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *SubmitDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits an API in DataService Studio.
//
// @param request - SubmitDataServiceApiRequest
//
// @return SubmitDataServiceApiResponse
func (client *Client) SubmitDataServiceApi(request *SubmitDataServiceApiRequest) (_result *SubmitDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitDataServiceApiResponse{}
	_body, _err := client.SubmitDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Commits a file to the development environment of the scheduling system to generate a task.
//
// @param request - SubmitFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitFileResponse
func (client *Client) SubmitFileWithOptions(request *SubmitFileRequest, runtime *dara.RuntimeOptions) (_result *SubmitFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.SkipAllDeployFileExtensions) {
		body["SkipAllDeployFileExtensions"] = request.SkipAllDeployFileExtensions
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Commits a file to the development environment of the scheduling system to generate a task.
//
// @param request - SubmitFileRequest
//
// @return SubmitFileResponse
func (client *Client) SubmitFile(request *SubmitFileRequest) (_result *SubmitFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitFileResponse{}
	_body, _err := client.SubmitFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends an instance.
//
// @param request - SuspendInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SuspendInstanceResponse
func (client *Client) SuspendInstanceWithOptions(request *SuspendInstanceRequest, runtime *dara.RuntimeOptions) (_result *SuspendInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SuspendInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SuspendInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends an instance.
//
// @param request - SuspendInstanceRequest
//
// @return SuspendInstanceResponse
func (client *Client) SuspendInstance(request *SuspendInstanceRequest) (_result *SuspendInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SuspendInstanceResponse{}
	_body, _err := client.SuspendInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Undeploys a real-time synchronization task.
//
// @param request - TerminateDISyncInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminateDISyncInstanceResponse
func (client *Client) TerminateDISyncInstanceWithOptions(request *TerminateDISyncInstanceRequest, runtime *dara.RuntimeOptions) (_result *TerminateDISyncInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminateDISyncInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminateDISyncInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Undeploys a real-time synchronization task.
//
// @param request - TerminateDISyncInstanceRequest
//
// @return TerminateDISyncInstanceResponse
func (client *Client) TerminateDISyncInstance(request *TerminateDISyncInstanceRequest) (_result *TerminateDISyncInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TerminateDISyncInstanceResponse{}
	_body, _err := client.TerminateDISyncInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests a DataService Studio API in asynchronous mode. You can call the GetDataServiceApiTest operation to query the test result.
//
// @param request - TestDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestDataServiceApiResponse
func (client *Client) TestDataServiceApiWithOptions(request *TestDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *TestDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiId) {
		query["ApiId"] = request.ApiId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.BodyContent) {
		body["BodyContent"] = request.BodyContent
	}

	if !dara.IsNil(request.BodyParams) {
		body["BodyParams"] = request.BodyParams
	}

	if !dara.IsNil(request.HeadParams) {
		body["HeadParams"] = request.HeadParams
	}

	if !dara.IsNil(request.PathParams) {
		body["PathParams"] = request.PathParams
	}

	if !dara.IsNil(request.QueryParam) {
		body["QueryParam"] = request.QueryParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests a DataService Studio API in asynchronous mode. You can call the GetDataServiceApiTest operation to query the test result.
//
// @param request - TestDataServiceApiRequest
//
// @return TestDataServiceApiResponse
func (client *Client) TestDataServiceApi(request *TestDataServiceApiRequest) (_result *TestDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestDataServiceApiResponse{}
	_body, _err := client.TestDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Tests the network connectivity between a data source and a resource group.
//
// @param request - TestNetworkConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestNetworkConnectionResponse
func (client *Client) TestNetworkConnectionWithOptions(request *TestNetworkConnectionRequest, runtime *dara.RuntimeOptions) (_result *TestNetworkConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasourceName) {
		query["DatasourceName"] = request.DatasourceName
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResourceGroup) {
		query["ResourceGroup"] = request.ResourceGroup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestNetworkConnection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestNetworkConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Tests the network connectivity between a data source and a resource group.
//
// @param request - TestNetworkConnectionRequest
//
// @return TestNetworkConnectionResponse
func (client *Client) TestNetworkConnection(request *TestNetworkConnectionRequest) (_result *TestNetworkConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TestNetworkConnectionResponse{}
	_body, _err := client.TestNetworkConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ranking of the running durations of instances.
//
// @param request - TopTenElapsedTimeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TopTenElapsedTimeInstanceResponse
func (client *Client) TopTenElapsedTimeInstanceWithOptions(request *TopTenElapsedTimeInstanceRequest, runtime *dara.RuntimeOptions) (_result *TopTenElapsedTimeInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TopTenElapsedTimeInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TopTenElapsedTimeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ranking of the running durations of instances.
//
// @param request - TopTenElapsedTimeInstanceRequest
//
// @return TopTenElapsedTimeInstanceResponse
func (client *Client) TopTenElapsedTimeInstance(request *TopTenElapsedTimeInstanceRequest) (_result *TopTenElapsedTimeInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TopTenElapsedTimeInstanceResponse{}
	_body, _err := client.TopTenElapsedTimeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the ranking of nodes on which errors occur within the previous month.
//
// @param request - TopTenErrorTimesInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TopTenErrorTimesInstanceResponse
func (client *Client) TopTenErrorTimesInstanceWithOptions(request *TopTenErrorTimesInstanceRequest, runtime *dara.RuntimeOptions) (_result *TopTenErrorTimesInstanceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TopTenErrorTimesInstance"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TopTenErrorTimesInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the ranking of nodes on which errors occur within the previous month.
//
// @param request - TopTenErrorTimesInstanceRequest
//
// @return TopTenErrorTimesInstanceResponse
func (client *Client) TopTenErrorTimesInstance(request *TopTenErrorTimesInstanceRequest) (_result *TopTenErrorTimesInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TopTenErrorTimesInstanceResponse{}
	_body, _err := client.TopTenErrorTimesInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a baseline.
//
// @param tmpReq - UpdateBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBaselineResponse
func (client *Client) UpdateBaselineWithOptions(tmpReq *UpdateBaselineRequest, runtime *dara.RuntimeOptions) (_result *UpdateBaselineResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateBaselineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AlertSettings) {
		request.AlertSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlertSettings, dara.String("AlertSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OvertimeSettings) {
		request.OvertimeSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OvertimeSettings, dara.String("OvertimeSettings"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertEnabled) {
		body["AlertEnabled"] = request.AlertEnabled
	}

	if !dara.IsNil(request.AlertMarginThreshold) {
		body["AlertMarginThreshold"] = request.AlertMarginThreshold
	}

	if !dara.IsNil(request.AlertSettingsShrink) {
		body["AlertSettings"] = request.AlertSettingsShrink
	}

	if !dara.IsNil(request.BaselineId) {
		body["BaselineId"] = request.BaselineId
	}

	if !dara.IsNil(request.BaselineName) {
		body["BaselineName"] = request.BaselineName
	}

	if !dara.IsNil(request.BaselineType) {
		body["BaselineType"] = request.BaselineType
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.OvertimeSettingsShrink) {
		body["OvertimeSettings"] = request.OvertimeSettingsShrink
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.Priority) {
		body["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RemoveNodeIds) {
		body["RemoveNodeIds"] = request.RemoveNodeIds
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBaseline"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a baseline.
//
// @param request - UpdateBaselineRequest
//
// @return UpdateBaselineResponse
func (client *Client) UpdateBaseline(request *UpdateBaselineRequest) (_result *UpdateBaselineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBaselineResponse{}
	_body, _err := client.UpdateBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// @param request - UpdateBusinessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateBusinessResponse
func (client *Client) UpdateBusinessWithOptions(request *UpdateBusinessRequest, runtime *dara.RuntimeOptions) (_result *UpdateBusinessResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BusinessId) {
		body["BusinessId"] = request.BusinessId
	}

	if !dara.IsNil(request.BusinessName) {
		body["BusinessName"] = request.BusinessName
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateBusiness"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateBusinessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// @param request - UpdateBusinessRequest
//
// @return UpdateBusinessResponse
func (client *Client) UpdateBusiness(request *UpdateBusinessRequest) (_result *UpdateBusinessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateBusinessResponse{}
	_body, _err := client.UpdateBusinessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configurations of submodules in a workspace. You can configure SPARK parameters.
//
// @param tmpReq - UpdateClusterConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClusterConfigsResponse
func (client *Client) UpdateClusterConfigsWithOptions(tmpReq *UpdateClusterConfigsRequest, runtime *dara.RuntimeOptions) (_result *UpdateClusterConfigsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateClusterConfigsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ConfigValues) {
		request.ConfigValuesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigValues, dara.String("ConfigValues"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		query["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ConfigType) {
		query["ConfigType"] = request.ConfigType
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigValuesShrink) {
		body["ConfigValues"] = request.ConfigValuesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateClusterConfigs"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateClusterConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configurations of submodules in a workspace. You can configure SPARK parameters.
//
// @param request - UpdateClusterConfigsRequest
//
// @return UpdateClusterConfigsResponse
func (client *Client) UpdateClusterConfigs(request *UpdateClusterConfigsRequest) (_result *UpdateClusterConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateClusterConfigsResponse{}
	_body, _err := client.UpdateClusterConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Deprecated: OpenAPI UpdateConnection is deprecated, please use dataworks-public::2020-05-18::UpdateDataSource instead.
//
// Summary:
//
// Updates a data source.
//
// @param request - UpdateConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConnectionResponse
func (client *Client) UpdateConnectionWithOptions(request *UpdateConnectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateConnectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ConnectionId) {
		query["ConnectionId"] = request.ConnectionId
	}

	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConnection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Deprecated: OpenAPI UpdateConnection is deprecated, please use dataworks-public::2020-05-18::UpdateDataSource instead.
//
// Summary:
//
// Updates a data source.
//
// @param request - UpdateConnectionRequest
//
// @return UpdateConnectionResponse
// Deprecated
func (client *Client) UpdateConnection(request *UpdateConnectionRequest) (_result *UpdateConnectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateConnectionResponse{}
	_body, _err := client.UpdateConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an alert rule for a new-version synchronization task. The following type of task is supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization. You must update all fields in the alert rule.
//
// @param tmpReq - UpdateDIAlarmRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRuleWithOptions(tmpReq *UpdateDIAlarmRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIAlarmRuleResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDIAlarmRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NotificationSettings) {
		request.NotificationSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NotificationSettings, dara.String("NotificationSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TriggerConditions) {
		request.TriggerConditionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TriggerConditions, dara.String("TriggerConditions"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DIAlarmRuleId) {
		body["DIAlarmRuleId"] = request.DIAlarmRuleId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.MetricType) {
		body["MetricType"] = request.MetricType
	}

	if !dara.IsNil(request.NotificationSettingsShrink) {
		body["NotificationSettings"] = request.NotificationSettingsShrink
	}

	if !dara.IsNil(request.TriggerConditionsShrink) {
		body["TriggerConditions"] = request.TriggerConditionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIAlarmRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an alert rule for a new-version synchronization task. The following type of task is supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// Description:
//
// You can configure alert rules only for tasks that can be used for real-time data synchronization. You must update all fields in the alert rule.
//
// @param request - UpdateDIAlarmRuleRequest
//
// @return UpdateDIAlarmRuleResponse
func (client *Client) UpdateDIAlarmRule(request *UpdateDIAlarmRuleRequest) (_result *UpdateDIAlarmRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDIAlarmRuleResponse{}
	_body, _err := client.UpdateDIAlarmRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a new-version synchronization task. The following type of task is supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// @param tmpReq - UpdateDIJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJobWithOptions(tmpReq *UpdateDIJobRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIJobResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &UpdateDIJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.JobSettings) {
		request.JobSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.JobSettings, dara.String("JobSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceSettings) {
		request.ResourceSettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceSettings, dara.String("ResourceSettings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TableMappings) {
		request.TableMappingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TableMappings, dara.String("TableMappings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TransformationRules) {
		request.TransformationRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TransformationRules, dara.String("TransformationRules"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DIJobId) {
		body["DIJobId"] = request.DIJobId
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.JobSettingsShrink) {
		body["JobSettings"] = request.JobSettingsShrink
	}

	if !dara.IsNil(request.ResourceSettingsShrink) {
		body["ResourceSettings"] = request.ResourceSettingsShrink
	}

	if !dara.IsNil(request.TableMappingsShrink) {
		body["TableMappings"] = request.TableMappingsShrink
	}

	if !dara.IsNil(request.TransformationRulesShrink) {
		body["TransformationRules"] = request.TransformationRulesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIJob"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a new-version synchronization task. The following type of task is supported: real-time synchronization of all data in a MySQL database to Hologres.
//
// @param request - UpdateDIJobRequest
//
// @return UpdateDIJobResponse
func (client *Client) UpdateDIJob(request *UpdateDIJobRequest) (_result *UpdateDIJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDIJobResponse{}
	_body, _err := client.UpdateDIJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the default global configuration of synchronization solutions in a DataWorks workspace.
//
// Description:
//
// DataWorks allows you to specify a default global configuration only for the processing rules of DDL messages in synchronization solutions. After you configure the **processing rules of DDL messages*	- in synchronization solutions, the configuration is used as the default global configuration and applies to all real-time synchronization tasks in the solutions. You can modify the **processing rules of DDL messages*	- based on your business requirements. For more information about how to configure a synchronization solution, see [Synchronization solutions](https://help.aliyun.com/document_detail/199008.html).
//
// @param request - UpdateDIProjectConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDIProjectConfigResponse
func (client *Client) UpdateDIProjectConfigWithOptions(request *UpdateDIProjectConfigRequest, runtime *dara.RuntimeOptions) (_result *UpdateDIProjectConfigResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DestinationType) {
		query["DestinationType"] = request.DestinationType
	}

	if !dara.IsNil(request.ProjectConfig) {
		query["ProjectConfig"] = request.ProjectConfig
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SourceType) {
		query["SourceType"] = request.SourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDIProjectConfig"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDIProjectConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the default global configuration of synchronization solutions in a DataWorks workspace.
//
// Description:
//
// DataWorks allows you to specify a default global configuration only for the processing rules of DDL messages in synchronization solutions. After you configure the **processing rules of DDL messages*	- in synchronization solutions, the configuration is used as the default global configuration and applies to all real-time synchronization tasks in the solutions. You can modify the **processing rules of DDL messages*	- based on your business requirements. For more information about how to configure a synchronization solution, see [Synchronization solutions](https://help.aliyun.com/document_detail/199008.html).
//
// @param request - UpdateDIProjectConfigRequest
//
// @return UpdateDIProjectConfigResponse
func (client *Client) UpdateDIProjectConfig(request *UpdateDIProjectConfigRequest) (_result *UpdateDIProjectConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDIProjectConfigResponse{}
	_body, _err := client.UpdateDIProjectConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data synchronization task.
//
// @param request - UpdateDISyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDISyncTaskResponse
func (client *Client) UpdateDISyncTaskWithOptions(request *UpdateDISyncTaskRequest, runtime *dara.RuntimeOptions) (_result *UpdateDISyncTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskContent) {
		query["TaskContent"] = request.TaskContent
	}

	if !dara.IsNil(request.TaskParam) {
		query["TaskParam"] = request.TaskParam
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDISyncTask"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDISyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data synchronization task.
//
// @param request - UpdateDISyncTaskRequest
//
// @return UpdateDISyncTaskResponse
func (client *Client) UpdateDISyncTask(request *UpdateDISyncTaskRequest) (_result *UpdateDISyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDISyncTaskResponse{}
	_body, _err := client.UpdateDISyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an API in the development state in DataService Studio.
//
// @param request - UpdateDataServiceApiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataServiceApiResponse
func (client *Client) UpdateDataServiceApiWithOptions(request *UpdateDataServiceApiRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataServiceApiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ApiDescription) {
		body["ApiDescription"] = request.ApiDescription
	}

	if !dara.IsNil(request.ApiId) {
		body["ApiId"] = request.ApiId
	}

	if !dara.IsNil(request.ApiPath) {
		body["ApiPath"] = request.ApiPath
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Protocols) {
		body["Protocols"] = request.Protocols
	}

	if !dara.IsNil(request.RegistrationDetails) {
		body["RegistrationDetails"] = request.RegistrationDetails
	}

	if !dara.IsNil(request.RequestMethod) {
		body["RequestMethod"] = request.RequestMethod
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResponseContentType) {
		body["ResponseContentType"] = request.ResponseContentType
	}

	if !dara.IsNil(request.ScriptDetails) {
		body["ScriptDetails"] = request.ScriptDetails
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.VisibleRange) {
		body["VisibleRange"] = request.VisibleRange
	}

	if !dara.IsNil(request.WizardDetails) {
		body["WizardDetails"] = request.WizardDetails
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataServiceApi"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataServiceApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an API in the development state in DataService Studio.
//
// @param request - UpdateDataServiceApiRequest
//
// @return UpdateDataServiceApiResponse
func (client *Client) UpdateDataServiceApi(request *UpdateDataServiceApiRequest) (_result *UpdateDataServiceApiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataServiceApiResponse{}
	_body, _err := client.UpdateDataServiceApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a data source.
//
// @param request - UpdateDataSourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSourceWithOptions(request *UpdateDataSourceRequest, runtime *dara.RuntimeOptions) (_result *UpdateDataSourceResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
	}

	if !dara.IsNil(request.DataSourceId) {
		query["DataSourceId"] = request.DataSourceId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataSource"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data source.
//
// @param request - UpdateDataSourceRequest
//
// @return UpdateDataSourceResponse
func (client *Client) UpdateDataSource(request *UpdateDataSourceRequest) (_result *UpdateDataSourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateDataSourceResponse{}
	_body, _err := client.UpdateDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a file.
//
// Description:
//
// When you debug or call this operation, you must specify new values for the specified parameters to ensure that the values are different from the original configurations of the file. For example, if the original value of a parameter is A, you must change the value of this parameter to B before you commit the node. If you set the parameter to A, an exception that indicates invalid data occurs.
//
// @param request - UpdateFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileResponse
func (client *Client) UpdateFileWithOptions(request *UpdateFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdvancedSettings) {
		body["AdvancedSettings"] = request.AdvancedSettings
	}

	if !dara.IsNil(request.ApplyScheduleImmediately) {
		body["ApplyScheduleImmediately"] = request.ApplyScheduleImmediately
	}

	if !dara.IsNil(request.AutoParsing) {
		body["AutoParsing"] = request.AutoParsing
	}

	if !dara.IsNil(request.AutoRerunIntervalMillis) {
		body["AutoRerunIntervalMillis"] = request.AutoRerunIntervalMillis
	}

	if !dara.IsNil(request.AutoRerunTimes) {
		body["AutoRerunTimes"] = request.AutoRerunTimes
	}

	if !dara.IsNil(request.ConnectionName) {
		body["ConnectionName"] = request.ConnectionName
	}

	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.CronExpress) {
		body["CronExpress"] = request.CronExpress
	}

	if !dara.IsNil(request.CycleType) {
		body["CycleType"] = request.CycleType
	}

	if !dara.IsNil(request.DependentNodeIdList) {
		body["DependentNodeIdList"] = request.DependentNodeIdList
	}

	if !dara.IsNil(request.DependentType) {
		body["DependentType"] = request.DependentType
	}

	if !dara.IsNil(request.EndEffectDate) {
		body["EndEffectDate"] = request.EndEffectDate
	}

	if !dara.IsNil(request.FileDescription) {
		body["FileDescription"] = request.FileDescription
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		body["FileName"] = request.FileName
	}

	if !dara.IsNil(request.IgnoreParentSkipRunningProperty) {
		body["IgnoreParentSkipRunningProperty"] = request.IgnoreParentSkipRunningProperty
	}

	if !dara.IsNil(request.ImageId) {
		body["ImageId"] = request.ImageId
	}

	if !dara.IsNil(request.InputList) {
		body["InputList"] = request.InputList
	}

	if !dara.IsNil(request.InputParameters) {
		body["InputParameters"] = request.InputParameters
	}

	if !dara.IsNil(request.OutputList) {
		body["OutputList"] = request.OutputList
	}

	if !dara.IsNil(request.OutputParameters) {
		body["OutputParameters"] = request.OutputParameters
	}

	if !dara.IsNil(request.Owner) {
		body["Owner"] = request.Owner
	}

	if !dara.IsNil(request.ParaValue) {
		body["ParaValue"] = request.ParaValue
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.RerunMode) {
		body["RerunMode"] = request.RerunMode
	}

	if !dara.IsNil(request.ResourceGroupIdentifier) {
		body["ResourceGroupIdentifier"] = request.ResourceGroupIdentifier
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	if !dara.IsNil(request.StartEffectDate) {
		body["StartEffectDate"] = request.StartEffectDate
	}

	if !dara.IsNil(request.StartImmediately) {
		body["StartImmediately"] = request.StartImmediately
	}

	if !dara.IsNil(request.Stop) {
		body["Stop"] = request.Stop
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a file.
//
// Description:
//
// When you debug or call this operation, you must specify new values for the specified parameters to ensure that the values are different from the original configurations of the file. For example, if the original value of a parameter is A, you must change the value of this parameter to B before you commit the node. If you set the parameter to A, an exception that indicates invalid data occurs.
//
// @param request - UpdateFileRequest
//
// @return UpdateFileResponse
func (client *Client) UpdateFile(request *UpdateFileRequest) (_result *UpdateFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFileResponse{}
	_body, _err := client.UpdateFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a folder.
//
// @param request - UpdateFolderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolderWithOptions(request *UpdateFolderRequest, runtime *dara.RuntimeOptions) (_result *UpdateFolderResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["FolderId"] = request.FolderId
	}

	if !dara.IsNil(request.FolderName) {
		body["FolderName"] = request.FolderName
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFolder"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFolderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a folder.
//
// @param request - UpdateFolderRequest
//
// @return UpdateFolderResponse
func (client *Client) UpdateFolder(request *UpdateFolderRequest) (_result *UpdateFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateFolderResponse{}
	_body, _err := client.UpdateFolderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the check result of an extension point event to DataStudio after the extension point event is triggered during data development and checked by an extension.
//
// @param request - UpdateIDEEventResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIDEEventResultResponse
func (client *Client) UpdateIDEEventResultWithOptions(request *UpdateIDEEventResultRequest, runtime *dara.RuntimeOptions) (_result *UpdateIDEEventResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CheckResult) {
		body["CheckResult"] = request.CheckResult
	}

	if !dara.IsNil(request.CheckResultTip) {
		body["CheckResultTip"] = request.CheckResultTip
	}

	if !dara.IsNil(request.ExtensionCode) {
		body["ExtensionCode"] = request.ExtensionCode
	}

	if !dara.IsNil(request.MessageId) {
		body["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateIDEEventResult"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIDEEventResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the check result of an extension point event to DataStudio after the extension point event is triggered during data development and checked by an extension.
//
// @param request - UpdateIDEEventResultRequest
//
// @return UpdateIDEEventResultResponse
func (client *Client) UpdateIDEEventResult(request *UpdateIDEEventResultRequest) (_result *UpdateIDEEventResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateIDEEventResultResponse{}
	_body, _err := client.UpdateIDEEventResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a category.
//
// @param request - UpdateMetaCategoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaCategoryResponse
func (client *Client) UpdateMetaCategoryWithOptions(request *UpdateMetaCategoryRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaCategoryResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryId) {
		body["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaCategory"),
		Version:     dara.String("2020-05-18"),
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
// Updates a category.
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
// Updates the name and comment of a collection.
//
// Description:
//
// Only the name and comment of a collection can be updated.
//
// @param request - UpdateMetaCollectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaCollectionResponse
func (client *Client) UpdateMetaCollectionWithOptions(request *UpdateMetaCollectionRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaCollectionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.QualifiedName) {
		query["QualifiedName"] = request.QualifiedName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaCollection"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the name and comment of a collection.
//
// Description:
//
// Only the name and comment of a collection can be updated.
//
// @param request - UpdateMetaCollectionRequest
//
// @return UpdateMetaCollectionResponse
func (client *Client) UpdateMetaCollection(request *UpdateMetaCollectionRequest) (_result *UpdateMetaCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMetaCollectionResponse{}
	_body, _err := client.UpdateMetaCollectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the metadata information about a table. Only MaxCompute tables are supported.
//
// @param request - UpdateMetaTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaTableResponse
func (client *Client) UpdateMetaTableWithOptions(request *UpdateMetaTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caption) {
		query["Caption"] = request.Caption
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.EnvType) {
		query["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.NewOwnerId) {
		query["NewOwnerId"] = request.NewOwnerId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AddedLabels) {
		body["AddedLabels"] = request.AddedLabels
	}

	if !dara.IsNil(request.RemovedLabels) {
		body["RemovedLabels"] = request.RemovedLabels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaTable"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the metadata information about a table. Only MaxCompute tables are supported.
//
// @param request - UpdateMetaTableRequest
//
// @return UpdateMetaTableResponse
func (client *Client) UpdateMetaTable(request *UpdateMetaTableRequest) (_result *UpdateMetaTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMetaTableResponse{}
	_body, _err := client.UpdateMetaTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the instructions on how to use a table. If no instruction on how to use the table is available, the instructions that are configured by calling this operation are added.
//
// @param request - UpdateMetaTableIntroWikiRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMetaTableIntroWikiResponse
func (client *Client) UpdateMetaTableIntroWikiWithOptions(request *UpdateMetaTableIntroWikiRequest, runtime *dara.RuntimeOptions) (_result *UpdateMetaTableIntroWikiResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMetaTableIntroWiki"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMetaTableIntroWikiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the instructions on how to use a table. If no instruction on how to use the table is available, the instructions that are configured by calling this operation are added.
//
// @param request - UpdateMetaTableIntroWikiRequest
//
// @return UpdateMetaTableIntroWikiResponse
func (client *Client) UpdateMetaTableIntroWiki(request *UpdateMetaTableIntroWikiRequest) (_result *UpdateMetaTableIntroWikiResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateMetaTableIntroWikiResponse{}
	_body, _err := client.UpdateMetaTableIntroWikiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the owner of a node.
//
// @param request - UpdateNodeOwnerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeOwnerResponse
func (client *Client) UpdateNodeOwnerWithOptions(request *UpdateNodeOwnerRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodeOwnerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.UserId) {
		body["UserId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodeOwner"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodeOwnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the owner of a node.
//
// @param request - UpdateNodeOwnerRequest
//
// @return UpdateNodeOwnerResponse
func (client *Client) UpdateNodeOwner(request *UpdateNodeOwnerRequest) (_result *UpdateNodeOwnerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNodeOwnerResponse{}
	_body, _err := client.UpdateNodeOwnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Freezes or unfreezes a node.
//
// @param request - UpdateNodeRunModeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeRunModeResponse
func (client *Client) UpdateNodeRunModeWithOptions(request *UpdateNodeRunModeRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodeRunModeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ProjectEnv) {
		body["ProjectEnv"] = request.ProjectEnv
	}

	if !dara.IsNil(request.SchedulerType) {
		body["SchedulerType"] = request.SchedulerType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodeRunMode"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateNodeRunModeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Freezes or unfreezes a node.
//
// @param request - UpdateNodeRunModeRequest
//
// @return UpdateNodeRunModeResponse
func (client *Client) UpdateNodeRunMode(request *UpdateNodeRunModeRequest) (_result *UpdateNodeRunModeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateNodeRunModeResponse{}
	_body, _err := client.UpdateNodeRunModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a subscription relationship.
//
// @param request - UpdateQualityFollowerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityFollowerResponse
func (client *Client) UpdateQualityFollowerWithOptions(request *UpdateQualityFollowerRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityFollowerResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlarmMode) {
		body["AlarmMode"] = request.AlarmMode
	}

	if !dara.IsNil(request.Follower) {
		body["Follower"] = request.Follower
	}

	if !dara.IsNil(request.FollowerId) {
		body["FollowerId"] = request.FollowerId
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityFollower"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityFollowerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a subscription relationship.
//
// @param request - UpdateQualityFollowerRequest
//
// @return UpdateQualityFollowerResponse
func (client *Client) UpdateQualityFollower(request *UpdateQualityFollowerRequest) (_result *UpdateQualityFollowerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityFollowerResponse{}
	_body, _err := client.UpdateQualityFollowerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a monitoring rule.
//
// @param request - UpdateQualityRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQualityRuleResponse
func (client *Client) UpdateQualityRuleWithOptions(request *UpdateQualityRuleRequest, runtime *dara.RuntimeOptions) (_result *UpdateQualityRuleResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BlockType) {
		body["BlockType"] = request.BlockType
	}

	if !dara.IsNil(request.Checker) {
		body["Checker"] = request.Checker
	}

	if !dara.IsNil(request.Comment) {
		body["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CriticalThreshold) {
		body["CriticalThreshold"] = request.CriticalThreshold
	}

	if !dara.IsNil(request.EntityId) {
		body["EntityId"] = request.EntityId
	}

	if !dara.IsNil(request.ExpectValue) {
		body["ExpectValue"] = request.ExpectValue
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.MethodName) {
		body["MethodName"] = request.MethodName
	}

	if !dara.IsNil(request.OpenSwitch) {
		body["OpenSwitch"] = request.OpenSwitch
	}

	if !dara.IsNil(request.Operator) {
		body["Operator"] = request.Operator
	}

	if !dara.IsNil(request.PredictType) {
		body["PredictType"] = request.PredictType
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Property) {
		body["Property"] = request.Property
	}

	if !dara.IsNil(request.PropertyType) {
		body["PropertyType"] = request.PropertyType
	}

	if !dara.IsNil(request.RuleName) {
		body["RuleName"] = request.RuleName
	}

	if !dara.IsNil(request.RuleType) {
		body["RuleType"] = request.RuleType
	}

	if !dara.IsNil(request.TaskSetting) {
		body["TaskSetting"] = request.TaskSetting
	}

	if !dara.IsNil(request.TemplateId) {
		body["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.Trend) {
		body["Trend"] = request.Trend
	}

	if !dara.IsNil(request.WarningThreshold) {
		body["WarningThreshold"] = request.WarningThreshold
	}

	if !dara.IsNil(request.WhereCondition) {
		body["WhereCondition"] = request.WhereCondition
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQualityRule"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQualityRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a monitoring rule.
//
// @param request - UpdateQualityRuleRequest
//
// @return UpdateQualityRuleResponse
func (client *Client) UpdateQualityRule(request *UpdateQualityRuleRequest) (_result *UpdateQualityRuleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateQualityRuleResponse{}
	_body, _err := client.UpdateQualityRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a custom alert rule.
//
// @param request - UpdateRemindRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRemindResponse
func (client *Client) UpdateRemindWithOptions(request *UpdateRemindRequest, runtime *dara.RuntimeOptions) (_result *UpdateRemindResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AlertInterval) {
		body["AlertInterval"] = request.AlertInterval
	}

	if !dara.IsNil(request.AlertMethods) {
		body["AlertMethods"] = request.AlertMethods
	}

	if !dara.IsNil(request.AlertTargets) {
		body["AlertTargets"] = request.AlertTargets
	}

	if !dara.IsNil(request.AlertUnit) {
		body["AlertUnit"] = request.AlertUnit
	}

	if !dara.IsNil(request.BaselineIds) {
		body["BaselineIds"] = request.BaselineIds
	}

	if !dara.IsNil(request.BizProcessIds) {
		body["BizProcessIds"] = request.BizProcessIds
	}

	if !dara.IsNil(request.Detail) {
		body["Detail"] = request.Detail
	}

	if !dara.IsNil(request.DndEnd) {
		body["DndEnd"] = request.DndEnd
	}

	if !dara.IsNil(request.MaxAlertTimes) {
		body["MaxAlertTimes"] = request.MaxAlertTimes
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RemindId) {
		body["RemindId"] = request.RemindId
	}

	if !dara.IsNil(request.RemindName) {
		body["RemindName"] = request.RemindName
	}

	if !dara.IsNil(request.RemindType) {
		body["RemindType"] = request.RemindType
	}

	if !dara.IsNil(request.RemindUnit) {
		body["RemindUnit"] = request.RemindUnit
	}

	if !dara.IsNil(request.RobotUrls) {
		body["RobotUrls"] = request.RobotUrls
	}

	if !dara.IsNil(request.UseFlag) {
		body["UseFlag"] = request.UseFlag
	}

	if !dara.IsNil(request.Webhooks) {
		body["Webhooks"] = request.Webhooks
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateRemind"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateRemindResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a custom alert rule.
//
// @param request - UpdateRemindRequest
//
// @return UpdateRemindResponse
func (client *Client) UpdateRemind(request *UpdateRemindRequest) (_result *UpdateRemindResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateRemindResponse{}
	_body, _err := client.UpdateRemindWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a MaxCompute table.
//
// @param request - UpdateTableRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableResponse
func (client *Client) UpdateTableWithOptions(request *UpdateTableRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppGuid) {
		query["AppGuid"] = request.AppGuid
	}

	if !dara.IsNil(request.CategoryId) {
		query["CategoryId"] = request.CategoryId
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.CreateIfNotExists) {
		query["CreateIfNotExists"] = request.CreateIfNotExists
	}

	if !dara.IsNil(request.ExternalTableType) {
		query["ExternalTableType"] = request.ExternalTableType
	}

	if !dara.IsNil(request.HasPart) {
		query["HasPart"] = request.HasPart
	}

	if !dara.IsNil(request.IsView) {
		query["IsView"] = request.IsView
	}

	if !dara.IsNil(request.LifeCycle) {
		query["LifeCycle"] = request.LifeCycle
	}

	if !dara.IsNil(request.Location) {
		query["Location"] = request.Location
	}

	if !dara.IsNil(request.LogicalLevelId) {
		query["LogicalLevelId"] = request.LogicalLevelId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhysicsLevelId) {
		query["PhysicsLevelId"] = request.PhysicsLevelId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.Schema) {
		query["Schema"] = request.Schema
	}

	if !dara.IsNil(request.TableName) {
		query["TableName"] = request.TableName
	}

	if !dara.IsNil(request.Visibility) {
		query["Visibility"] = request.Visibility
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Columns) {
		body["Columns"] = request.Columns
	}

	if !dara.IsNil(request.Endpoint) {
		body["Endpoint"] = request.Endpoint
	}

	if !dara.IsNil(request.EnvType) {
		body["EnvType"] = request.EnvType
	}

	if !dara.IsNil(request.Themes) {
		body["Themes"] = request.Themes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTable"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a MaxCompute table.
//
// @param request - UpdateTableRequest
//
// @return UpdateTableResponse
func (client *Client) UpdateTable(request *UpdateTableRequest) (_result *UpdateTableResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTableResponse{}
	_body, _err := client.UpdateTableWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the fields in a MaxCompute table.
//
// @param request - UpdateTableAddColumnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableAddColumnResponse
func (client *Client) UpdateTableAddColumnWithOptions(request *UpdateTableAddColumnRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableAddColumnResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Column) {
		body["Column"] = request.Column
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableAddColumn"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableAddColumnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the fields in a MaxCompute table.
//
// @param request - UpdateTableAddColumnRequest
//
// @return UpdateTableAddColumnResponse
func (client *Client) UpdateTableAddColumn(request *UpdateTableAddColumnRequest) (_result *UpdateTableAddColumnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTableAddColumnResponse{}
	_body, _err := client.UpdateTableAddColumnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a table level. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - UpdateTableLevelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableLevelResponse
func (client *Client) UpdateTableLevelWithOptions(request *UpdateTableLevelRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableLevelResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.LevelId) {
		query["LevelId"] = request.LevelId
	}

	if !dara.IsNil(request.LevelType) {
		query["LevelType"] = request.LevelType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableLevel"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a table level. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - UpdateTableLevelRequest
//
// @return UpdateTableLevelResponse
func (client *Client) UpdateTableLevel(request *UpdateTableLevelRequest) (_result *UpdateTableLevelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTableLevelResponse{}
	_body, _err := client.UpdateTableLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a table, such as the table folder, level, and category.
//
// @param request - UpdateTableModelInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableModelInfoResponse
func (client *Client) UpdateTableModelInfoWithOptions(request *UpdateTableModelInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableModelInfoResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FirstLevelThemeId) {
		query["FirstLevelThemeId"] = request.FirstLevelThemeId
	}

	if !dara.IsNil(request.LevelId) {
		query["LevelId"] = request.LevelId
	}

	if !dara.IsNil(request.LevelType) {
		query["LevelType"] = request.LevelType
	}

	if !dara.IsNil(request.SecondLevelThemeId) {
		query["SecondLevelThemeId"] = request.SecondLevelThemeId
	}

	if !dara.IsNil(request.TableGuid) {
		query["TableGuid"] = request.TableGuid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableModelInfo"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableModelInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a table, such as the table folder, level, and category.
//
// @param request - UpdateTableModelInfoRequest
//
// @return UpdateTableModelInfoResponse
func (client *Client) UpdateTableModelInfo(request *UpdateTableModelInfoRequest) (_result *UpdateTableModelInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTableModelInfoResponse{}
	_body, _err := client.UpdateTableModelInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a table theme. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - UpdateTableThemeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTableThemeResponse
func (client *Client) UpdateTableThemeWithOptions(request *UpdateTableThemeRequest, runtime *dara.RuntimeOptions) (_result *UpdateTableThemeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ThemeId) {
		query["ThemeId"] = request.ThemeId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTableTheme"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTableThemeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a table theme. This operation will be replaced soon. We recommend that you do not call this operation.
//
// @param request - UpdateTableThemeRequest
//
// @return UpdateTableThemeResponse
func (client *Client) UpdateTableTheme(request *UpdateTableThemeRequest) (_result *UpdateTableThemeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTableThemeResponse{}
	_body, _err := client.UpdateTableThemeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the file information about a function.
//
// @param request - UpdateUdfFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUdfFileResponse
func (client *Client) UpdateUdfFileWithOptions(request *UpdateUdfFileRequest, runtime *dara.RuntimeOptions) (_result *UpdateUdfFileResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClassName) {
		body["ClassName"] = request.ClassName
	}

	if !dara.IsNil(request.CmdDescription) {
		body["CmdDescription"] = request.CmdDescription
	}

	if !dara.IsNil(request.Example) {
		body["Example"] = request.Example
	}

	if !dara.IsNil(request.FileFolderPath) {
		body["FileFolderPath"] = request.FileFolderPath
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FunctionType) {
		body["FunctionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ParameterDescription) {
		body["ParameterDescription"] = request.ParameterDescription
	}

	if !dara.IsNil(request.ProjectId) {
		body["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectIdentifier) {
		body["ProjectIdentifier"] = request.ProjectIdentifier
	}

	if !dara.IsNil(request.Resources) {
		body["Resources"] = request.Resources
	}

	if !dara.IsNil(request.ReturnValue) {
		body["ReturnValue"] = request.ReturnValue
	}

	if !dara.IsNil(request.UdfDescription) {
		body["UdfDescription"] = request.UdfDescription
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateUdfFile"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateUdfFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the file information about a function.
//
// @param request - UpdateUdfFileRequest
//
// @return UpdateUdfFileResponse
func (client *Client) UpdateUdfFile(request *UpdateUdfFileRequest) (_result *UpdateUdfFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateUdfFileResponse{}
	_body, _err := client.UpdateUdfFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns the processing result sent by an extension after a process in Operation Center is blocked by the extension.
//
// @param request - UpdateWorkbenchEventResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkbenchEventResultResponse
func (client *Client) UpdateWorkbenchEventResultWithOptions(request *UpdateWorkbenchEventResultRequest, runtime *dara.RuntimeOptions) (_result *UpdateWorkbenchEventResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckResult) {
		query["CheckResult"] = request.CheckResult
	}

	if !dara.IsNil(request.CheckResultTip) {
		query["CheckResultTip"] = request.CheckResultTip
	}

	if !dara.IsNil(request.ExtensionCode) {
		query["ExtensionCode"] = request.ExtensionCode
	}

	if !dara.IsNil(request.MessageId) {
		query["MessageId"] = request.MessageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkbenchEventResult"),
		Version:     dara.String("2020-05-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkbenchEventResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns the processing result sent by an extension after a process in Operation Center is blocked by the extension.
//
// @param request - UpdateWorkbenchEventResultRequest
//
// @return UpdateWorkbenchEventResultResponse
func (client *Client) UpdateWorkbenchEventResult(request *UpdateWorkbenchEventResultRequest) (_result *UpdateWorkbenchEventResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateWorkbenchEventResultResponse{}
	_body, _err := client.UpdateWorkbenchEventResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}
