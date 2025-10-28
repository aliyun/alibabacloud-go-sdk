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
	client.Endpoint, _err = client.GetEndpoint(dara.String("ros"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels operations on a stack.
//
// @param request - CancelStackOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelStackOperationResponse
func (client *Client) CancelStackOperationWithOptions(request *CancelStackOperationRequest, runtime *dara.RuntimeOptions) (_result *CancelStackOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AllowedStackOperations) {
		query["AllowedStackOperations"] = request.AllowedStackOperations
	}

	if !dara.IsNil(request.CancelType) {
		query["CancelType"] = request.CancelType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelStackOperation"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelStackOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels operations on a stack.
//
// @param request - CancelStackOperationRequest
//
// @return CancelStackOperationResponse
func (client *Client) CancelStackOperation(request *CancelStackOperationRequest) (_result *CancelStackOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelStackOperationResponse{}
	_body, _err := client.CancelStackOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an update operation on a stack. You can call this operation to cancel an update operation on a stack when the stack is being updated or created.
//
// @param request - CancelUpdateStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelUpdateStackResponse
func (client *Client) CancelUpdateStackWithOptions(request *CancelUpdateStackRequest, runtime *dara.RuntimeOptions) (_result *CancelUpdateStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CancelType) {
		query["CancelType"] = request.CancelType
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelUpdateStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelUpdateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an update operation on a stack. You can call this operation to cancel an update operation on a stack when the stack is being updated or created.
//
// @param request - CancelUpdateStackRequest
//
// @return CancelUpdateStackResponse
func (client *Client) CancelUpdateStack(request *CancelUpdateStackRequest) (_result *CancelUpdateStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelUpdateStackResponse{}
	_body, _err := client.CancelUpdateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Continues to create a stack after the stack fails to be created.
//
// Description:
//
// This topic provides an example on how to continue to create a stack after the stack fails to be created. In this example, the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is created in the China (Hangzhou) region.
//
// @param request - ContinueCreateStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueCreateStackResponse
func (client *Client) ContinueCreateStackWithOptions(request *ContinueCreateStackRequest, runtime *dara.RuntimeOptions) (_result *ContinueCreateStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.Mode) {
		query["Mode"] = request.Mode
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.RecreatingOptions) {
		query["RecreatingOptions"] = request.RecreatingOptions
	}

	if !dara.IsNil(request.RecreatingResources) {
		query["RecreatingResources"] = request.RecreatingResources
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateBody) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ContinueCreateStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ContinueCreateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Continues to create a stack after the stack fails to be created.
//
// Description:
//
// This topic provides an example on how to continue to create a stack after the stack fails to be created. In this example, the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is created in the China (Hangzhou) region.
//
// @param request - ContinueCreateStackRequest
//
// @return ContinueCreateStackResponse
func (client *Client) ContinueCreateStack(request *ContinueCreateStackRequest) (_result *ContinueCreateStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ContinueCreateStackResponse{}
	_body, _err := client.ContinueCreateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create AI Task
//
// Description:
//
// This API allows users to create an AI task based on the specified task type, covering a range of capabilities from natural language understanding to resource stack deployment. Users need to provide the task type and any required parameters, and the API will return a unique TaskId for tracking the status and results of the task.
//
// @param request - CreateAITaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAITaskResponse
func (client *Client) CreateAITaskWithOptions(request *CreateAITaskRequest, runtime *dara.RuntimeOptions) (_result *CreateAITaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Prompt) {
		query["Prompt"] = request.Prompt
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Template) {
		body["Template"] = request.Template
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAITask"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAITaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create AI Task
//
// Description:
//
// This API allows users to create an AI task based on the specified task type, covering a range of capabilities from natural language understanding to resource stack deployment. Users need to provide the task type and any required parameters, and the API will return a unique TaskId for tracking the status and results of the task.
//
// @param request - CreateAITaskRequest
//
// @return CreateAITaskResponse
func (client *Client) CreateAITask(request *CreateAITaskRequest) (_result *CreateAITaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAITaskResponse{}
	_body, _err := client.CreateAITaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a change set for a stack. You can view proposed changes before you execute the change set.
//
// Description:
//
// ### [](#)Scenarios
//
// #### [](#)Use a change set to create a stack
//
// If you want to manage a large number of cloud resources and preview the creation effect of the resources before a stack that contains the resources is created, you can create the stack by using a change set. In this case, you must set `ChangeSetType` to CREATE and configure the relevant parameters. For more information about change sets, see [Change set](https://help.aliyun.com/document_detail/155649.html).
//
// #### [](#)Use a change set to update a stack
//
// If you want to preview the impacts of changes to an existing stack before you update the stack resources, you can create a change set for the stack. In this case, you must set ChangeSetType to UPDATE and configure the relevant parameters. For more information about change sets, see [Change set](https://help.aliyun.com/document_detail/155649.html).
//
// #### [](#)Use a change set and existing resources to create a stack
//
// If you want to add existing cloud resources to a new stack for centralized management, you can use a change set to create a stack and import the resources to the stack. In this case, you must set ChangeSetType to IMPORT and configure the relevant parameters. For more information about the resource import feature, see [Overview](https://help.aliyun.com/document_detail/193454.html).
//
// #### [](#)Use a change set and existing resources to update a stack
//
// If you want to import existing resources to an existing stack for centralized management, you can use a change set to update the stack. In this case, you must set ChangeSetType to IMPORT and configure the relevant parameters. For more information about the resource import feature, see [Overview](https://help.aliyun.com/document_detail/193454.html).
//
// ### [](#)Limits
//
//   - You can use change sets to update only stacks that are in specific states. For more information, see [Use a change set to update a stack](https://help.aliyun.com/document_detail/155873.html).
//
//   - A stack can have up to 20 change sets.
//
//   - Change sets reflect only the changes to stacks. Change sets do not reflect whether stacks can be successfully updated.
//
//   - A change set does not check if you exceed an account limit, if you update resources that cannot be updated, or if you have insufficient permissions to modify resources, all of which can cause a stack update to fail. If a stack update fails, Resource Orchestration Service (ROS) attempts to roll back your resources to their original status.
//
// This topic provides an example on how to use a change set to update a stack. In this example, a change set named `MyChangeSet` is created in the `China (Hangzhou)` region. The template of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is updated to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
//
// @param request - CreateChangeSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChangeSetResponse
func (client *Client) CreateChangeSetWithOptions(request *CreateChangeSetRequest, runtime *dara.RuntimeOptions) (_result *CreateChangeSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetName) {
		query["ChangeSetName"] = request.ChangeSetName
	}

	if !dara.IsNil(request.ChangeSetType) {
		query["ChangeSetType"] = request.ChangeSetType
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.DisableRollback) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.NotificationURLs) {
		query["NotificationURLs"] = request.NotificationURLs
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplacementOption) {
		query["ReplacementOption"] = request.ReplacementOption
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourcesToImport) {
		query["ResourcesToImport"] = request.ResourcesToImport
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.StackName) {
		query["StackName"] = request.StackName
	}

	if !dara.IsNil(request.StackPolicyBody) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !dara.IsNil(request.StackPolicyDuringUpdateBody) {
		query["StackPolicyDuringUpdateBody"] = request.StackPolicyDuringUpdateBody
	}

	if !dara.IsNil(request.StackPolicyDuringUpdateURL) {
		query["StackPolicyDuringUpdateURL"] = request.StackPolicyDuringUpdateURL
	}

	if !dara.IsNil(request.StackPolicyURL) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TaintResources) {
		query["TaintResources"] = request.TaintResources
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.TimeoutInMinutes) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	if !dara.IsNil(request.UsePreviousParameters) {
		query["UsePreviousParameters"] = request.UsePreviousParameters
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChangeSet"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a change set for a stack. You can view proposed changes before you execute the change set.
//
// Description:
//
// ### [](#)Scenarios
//
// #### [](#)Use a change set to create a stack
//
// If you want to manage a large number of cloud resources and preview the creation effect of the resources before a stack that contains the resources is created, you can create the stack by using a change set. In this case, you must set `ChangeSetType` to CREATE and configure the relevant parameters. For more information about change sets, see [Change set](https://help.aliyun.com/document_detail/155649.html).
//
// #### [](#)Use a change set to update a stack
//
// If you want to preview the impacts of changes to an existing stack before you update the stack resources, you can create a change set for the stack. In this case, you must set ChangeSetType to UPDATE and configure the relevant parameters. For more information about change sets, see [Change set](https://help.aliyun.com/document_detail/155649.html).
//
// #### [](#)Use a change set and existing resources to create a stack
//
// If you want to add existing cloud resources to a new stack for centralized management, you can use a change set to create a stack and import the resources to the stack. In this case, you must set ChangeSetType to IMPORT and configure the relevant parameters. For more information about the resource import feature, see [Overview](https://help.aliyun.com/document_detail/193454.html).
//
// #### [](#)Use a change set and existing resources to update a stack
//
// If you want to import existing resources to an existing stack for centralized management, you can use a change set to update the stack. In this case, you must set ChangeSetType to IMPORT and configure the relevant parameters. For more information about the resource import feature, see [Overview](https://help.aliyun.com/document_detail/193454.html).
//
// ### [](#)Limits
//
//   - You can use change sets to update only stacks that are in specific states. For more information, see [Use a change set to update a stack](https://help.aliyun.com/document_detail/155873.html).
//
//   - A stack can have up to 20 change sets.
//
//   - Change sets reflect only the changes to stacks. Change sets do not reflect whether stacks can be successfully updated.
//
//   - A change set does not check if you exceed an account limit, if you update resources that cannot be updated, or if you have insufficient permissions to modify resources, all of which can cause a stack update to fail. If a stack update fails, Resource Orchestration Service (ROS) attempts to roll back your resources to their original status.
//
// This topic provides an example on how to use a change set to update a stack. In this example, a change set named `MyChangeSet` is created in the `China (Hangzhou)` region. The template of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is updated to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
//
// @param request - CreateChangeSetRequest
//
// @return CreateChangeSetResponse
func (client *Client) CreateChangeSet(request *CreateChangeSetRequest) (_result *CreateChangeSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateChangeSetResponse{}
	_body, _err := client.CreateChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a dignosis task.
//
// @param request - CreateDiagnosticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnosticResponse
func (client *Client) CreateDiagnosticWithOptions(request *CreateDiagnosticRequest, runtime *dara.RuntimeOptions) (_result *CreateDiagnosticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiagnosticKey) {
		query["DiagnosticKey"] = request.DiagnosticKey
	}

	if !dara.IsNil(request.DiagnosticType) {
		query["DiagnosticType"] = request.DiagnosticType
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Product) {
		query["Product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiagnostic"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDiagnosticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a dignosis task.
//
// @param request - CreateDiagnosticRequest
//
// @return CreateDiagnosticResponse
func (client *Client) CreateDiagnostic(request *CreateDiagnosticRequest) (_result *CreateDiagnosticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateDiagnosticResponse{}
	_body, _err := client.CreateDiagnosticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a stack that contains a collection of resources by using a Resource Orchestration Service (ROS) template.
//
// Description:
//
// A stack is a collection of ROS resources that you can manage as a single unit. To create a collection of resources, you can create a stack. For more information about stacks, see [Overview](https://help.aliyun.com/document_detail/172973.html).\\
//
// When you call the operation, take note of the following limits:
//
//   - You can create up to 200 stacks within an Alibaba Cloud account.
//
//   - You can create up to 200 resources in a stack.
//
// This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template. In this example, `TemplateBody` is set to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
//
// @param request - CreateStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStackResponse
func (client *Client) CreateStackWithOptions(request *CreateStackRequest, runtime *dara.RuntimeOptions) (_result *CreateStackResponse, _err error) {
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

	if !dara.IsNil(request.CreateOption) {
		query["CreateOption"] = request.CreateOption
	}

	if !dara.IsNil(request.CreateOptions) {
		query["CreateOptions"] = request.CreateOptions
	}

	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.DisableRollback) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.NotificationURLs) {
		query["NotificationURLs"] = request.NotificationURLs
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StackName) {
		query["StackName"] = request.StackName
	}

	if !dara.IsNil(request.StackPolicyBody) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !dara.IsNil(request.StackPolicyURL) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !dara.IsNil(request.TemplateScratchRegionId) {
		query["TemplateScratchRegionId"] = request.TemplateScratchRegionId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.TimeoutInMinutes) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a stack that contains a collection of resources by using a Resource Orchestration Service (ROS) template.
//
// Description:
//
// A stack is a collection of ROS resources that you can manage as a single unit. To create a collection of resources, you can create a stack. For more information about stacks, see [Overview](https://help.aliyun.com/document_detail/172973.html).\\
//
// When you call the operation, take note of the following limits:
//
//   - You can create up to 200 stacks within an Alibaba Cloud account.
//
//   - You can create up to 200 resources in a stack.
//
// This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template. In this example, `TemplateBody` is set to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
//
// @param request - CreateStackRequest
//
// @return CreateStackResponse
func (client *Client) CreateStack(request *CreateStackRequest) (_result *CreateStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStackResponse{}
	_body, _err := client.CreateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates stack groups based on Resource Orchestration Service (ROS) templates. Stack groups allow you to create stacks within multiple Alibaba Cloud accounts across regions.
//
// Description:
//
// A stack group is a collection of ROS stacks that you can manage as a unit. You can use an ROS template of a stack group to create stacks within Alibaba Cloud accounts across regions.
//
// You can create a stack group that is granted self-managed or service-managed permissions:
//
//   - If you use an Alibaba Cloud account to create a self-managed stack group, the administrator account and the execution account are Alibaba Cloud accounts.
//
//   - If you enable a resource directory and use the management account or a delegated administrator account of the resource directory to create a service-managed stack group, the administrator account is the management account or delegated administrator account, and the execution account is a member account of the resource directory.
//
// For more information about stack groups, see [Overview](https://help.aliyun.com/document_detail/154578.html).
//
// In this topic, a stack group named `MyStackGroup` is created in the `China (Hangzhou)` region and granted the self-managed permissions. In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is used.
//
// @param tmpReq - CreateStackGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStackGroupResponse
func (client *Client) CreateStackGroupWithOptions(tmpReq *CreateStackGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateStackGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStackGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AutoDeployment) {
		request.AutoDeploymentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoDeployment, dara.String("AutoDeployment"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AdministrationRoleName) {
		query["AdministrationRoleName"] = request.AdministrationRoleName
	}

	if !dara.IsNil(request.AutoDeploymentShrink) {
		query["AutoDeployment"] = request.AutoDeploymentShrink
	}

	if !dara.IsNil(request.Capabilities) {
		query["Capabilities"] = request.Capabilities
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionRoleName) {
		query["ExecutionRoleName"] = request.ExecutionRoleName
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.PermissionModel) {
		query["PermissionModel"] = request.PermissionModel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StackArn) {
		query["StackArn"] = request.StackArn
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStackGroup"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates stack groups based on Resource Orchestration Service (ROS) templates. Stack groups allow you to create stacks within multiple Alibaba Cloud accounts across regions.
//
// Description:
//
// A stack group is a collection of ROS stacks that you can manage as a unit. You can use an ROS template of a stack group to create stacks within Alibaba Cloud accounts across regions.
//
// You can create a stack group that is granted self-managed or service-managed permissions:
//
//   - If you use an Alibaba Cloud account to create a self-managed stack group, the administrator account and the execution account are Alibaba Cloud accounts.
//
//   - If you enable a resource directory and use the management account or a delegated administrator account of the resource directory to create a service-managed stack group, the administrator account is the management account or delegated administrator account, and the execution account is a member account of the resource directory.
//
// For more information about stack groups, see [Overview](https://help.aliyun.com/document_detail/154578.html).
//
// In this topic, a stack group named `MyStackGroup` is created in the `China (Hangzhou)` region and granted the self-managed permissions. In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is used.
//
// @param request - CreateStackGroupRequest
//
// @return CreateStackGroupResponse
func (client *Client) CreateStackGroup(request *CreateStackGroupRequest) (_result *CreateStackGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStackGroupResponse{}
	_body, _err := client.CreateStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates stack instances in the specified accounts and regions.
//
// Description:
//
// Before you call this operation, make sure that a stack group is created. For more information, see [CreateStackGroup](https://help.aliyun.com/document_detail/151333.html).
//
// In this topic, the stack group named `MyStackGroup` is used. The stack group is created in the China (Hangzhou) region and granted the self-managed permissions. In this example, stacks are created by using Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
//
// @param tmpReq - CreateStackInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStackInstancesResponse
func (client *Client) CreateStackInstancesWithOptions(tmpReq *CreateStackInstancesRequest, runtime *dara.RuntimeOptions) (_result *CreateStackInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("AccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeploymentTargets) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, dara.String("DeploymentTargets"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OperationPreferences) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, dara.String("OperationPreferences"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeploymentOptions) {
		query["DeploymentOptions"] = request.DeploymentOptions
	}

	if !dara.IsNil(request.DeploymentTargetsShrink) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !dara.IsNil(request.DisableRollback) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.OperationDescription) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !dara.IsNil(request.OperationPreferencesShrink) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !dara.IsNil(request.ParameterOverrides) {
		query["ParameterOverrides"] = request.ParameterOverrides
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.TimeoutInMinutes) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateStackInstances"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates stack instances in the specified accounts and regions.
//
// Description:
//
// Before you call this operation, make sure that a stack group is created. For more information, see [CreateStackGroup](https://help.aliyun.com/document_detail/151333.html).
//
// In this topic, the stack group named `MyStackGroup` is used. The stack group is created in the China (Hangzhou) region and granted the self-managed permissions. In this example, stacks are created by using Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
//
// @param request - CreateStackInstancesRequest
//
// @return CreateStackInstancesResponse
func (client *Client) CreateStackInstances(request *CreateStackInstancesRequest) (_result *CreateStackInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateStackInstancesResponse{}
	_body, _err := client.CreateStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom template.
//
// Description:
//
// In this topic, a custom template named `MyTemplate` is created in the `cn-hangzhou` region. The `TemplateBody` parameter of the template is set to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
//
// @param request - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.ValidationOptions) {
		query["ValidationOptions"] = request.ValidationOptions
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom template.
//
// Description:
//
// In this topic, a custom template named `MyTemplate` is created in the `cn-hangzhou` region. The `TemplateBody` parameter of the template is set to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a resource scenario.
//
// Description:
//
// ### [](#)Limits
//
// Only specific resource types support the resource scenario feature. For more information, see [Resource types that support the scenario feature](https://help.aliyun.com/document_detail/353175.htmll).
//
// ### [](#)Description
//
// Resource Orchestration Service (ROS) provides the resource scenario feature that allows you to specify the scope of a collection of resources on a user interface (UI) and perform operations, such as replication and management, on the resources. This helps you manage resources in a simplified manner. For more information about resource scenarios, see [Overview](https://help.aliyun.com/document_detail/352074.html).
//
// #### [](#)Resource replication scenario
//
// If you want to replicate a collection of resources and dependencies between the resources, you can create a resource replication scenario. This type of resource scenario allows you to replicate all existing resources within the specified scope and generate a collection of resources that have the same architecture as the existing resources. For more information, see [Resource replication scenario](https://help.aliyun.com/document_detail/353133.html).
//
// #### [](#)Resource detection scenario
//
// If the relationships between resources that you want to create are complicated, you can create a resource detection scenario to preview the overall resource architecture or the architecture of a specific resource. This facilitates resource management. For more information, see [Resource detection scenario](https://help.aliyun.com/document_detail/2591901.html).
//
// #### [](#)Resource management scenario
//
// If you want to import a collection of existing resources to a stack and manage the resources in a centralized manner, you can create a resource management scenario. For more information, see [Resource management scenario](https://help.aliyun.com/document_detail/353163.html).
//
// #### [](#)Resource migration scenario
//
// If you want to migrate a collection of resources and dependencies between the resources, you can create a resource migration scenario. When you migrate the resources, ROS generates a stack. You can view the migration progress on the Stacks tab of the scenario details page. After you migrate the resources, you can delete source resources. For more information, see [Resource migration scenario](https://help.aliyun.com/document_detail/379902.html).
//
// This topic provides an example on how to create a resource replication scenario in the China (Hangzhou) region to replicate a resource. In this example, a virtual private cloud (VPC) whose ID is `vpc-bp1m6fww66xbntjyc****` is replicated.
//
// @param tmpReq - CreateTemplateScratchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateScratchResponse
func (client *Client) CreateTemplateScratchWithOptions(tmpReq *CreateTemplateScratchRequest, runtime *dara.RuntimeOptions) (_result *CreateTemplateScratchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateTemplateScratchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PreferenceParameters) {
		request.PreferenceParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreferenceParameters, dara.String("PreferenceParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceResourceGroup) {
		request.SourceResourceGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResourceGroup, dara.String("SourceResourceGroup"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceResources) {
		request.SourceResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResources, dara.String("SourceResources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceTag) {
		request.SourceTagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTag, dara.String("SourceTag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionMode) {
		query["ExecutionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.LogicalIdStrategy) {
		query["LogicalIdStrategy"] = request.LogicalIdStrategy
	}

	if !dara.IsNil(request.PreferenceParametersShrink) {
		query["PreferenceParameters"] = request.PreferenceParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceResourceGroupShrink) {
		query["SourceResourceGroup"] = request.SourceResourceGroupShrink
	}

	if !dara.IsNil(request.SourceResourcesShrink) {
		query["SourceResources"] = request.SourceResourcesShrink
	}

	if !dara.IsNil(request.SourceTagShrink) {
		query["SourceTag"] = request.SourceTagShrink
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplateScratchType) {
		query["TemplateScratchType"] = request.TemplateScratchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTemplateScratch"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a resource scenario.
//
// Description:
//
// ### [](#)Limits
//
// Only specific resource types support the resource scenario feature. For more information, see [Resource types that support the scenario feature](https://help.aliyun.com/document_detail/353175.htmll).
//
// ### [](#)Description
//
// Resource Orchestration Service (ROS) provides the resource scenario feature that allows you to specify the scope of a collection of resources on a user interface (UI) and perform operations, such as replication and management, on the resources. This helps you manage resources in a simplified manner. For more information about resource scenarios, see [Overview](https://help.aliyun.com/document_detail/352074.html).
//
// #### [](#)Resource replication scenario
//
// If you want to replicate a collection of resources and dependencies between the resources, you can create a resource replication scenario. This type of resource scenario allows you to replicate all existing resources within the specified scope and generate a collection of resources that have the same architecture as the existing resources. For more information, see [Resource replication scenario](https://help.aliyun.com/document_detail/353133.html).
//
// #### [](#)Resource detection scenario
//
// If the relationships between resources that you want to create are complicated, you can create a resource detection scenario to preview the overall resource architecture or the architecture of a specific resource. This facilitates resource management. For more information, see [Resource detection scenario](https://help.aliyun.com/document_detail/2591901.html).
//
// #### [](#)Resource management scenario
//
// If you want to import a collection of existing resources to a stack and manage the resources in a centralized manner, you can create a resource management scenario. For more information, see [Resource management scenario](https://help.aliyun.com/document_detail/353163.html).
//
// #### [](#)Resource migration scenario
//
// If you want to migrate a collection of resources and dependencies between the resources, you can create a resource migration scenario. When you migrate the resources, ROS generates a stack. You can view the migration progress on the Stacks tab of the scenario details page. After you migrate the resources, you can delete source resources. For more information, see [Resource migration scenario](https://help.aliyun.com/document_detail/379902.html).
//
// This topic provides an example on how to create a resource replication scenario in the China (Hangzhou) region to replicate a resource. In this example, a virtual private cloud (VPC) whose ID is `vpc-bp1m6fww66xbntjyc****` is replicated.
//
// @param request - CreateTemplateScratchRequest
//
// @return CreateTemplateScratchResponse
func (client *Client) CreateTemplateScratch(request *CreateTemplateScratchRequest) (_result *CreateTemplateScratchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateTemplateScratchResponse{}
	_body, _err := client.CreateTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes change sets.
//
// Description:
//
//	  Before you call this operation, make sure that the following requirements are met:
//
//	    	- The status of the change set is CREATE_COMPLETE, CREATE_FAILED, or DELETE_FAILED.
//
//	    	- The execution status is UNAVAILABLE or AVAILABLE.
//
//		- After a change set is executed, other change sets associated with the same stack as this change set are also deleted.
//
//		- After a stack is deleted, change sets associated with the stack are deleted.
//
//		- If a change set of the CREATE type is deleted, you must delete stacks associated with the change set.
//
// In this example, a change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is deleted. The change set is created in the China (Hangzhou) region.
//
// @param request - DeleteChangeSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChangeSetResponse
func (client *Client) DeleteChangeSetWithOptions(request *DeleteChangeSetRequest, runtime *dara.RuntimeOptions) (_result *DeleteChangeSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetId) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteChangeSet"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes change sets.
//
// Description:
//
//	  Before you call this operation, make sure that the following requirements are met:
//
//	    	- The status of the change set is CREATE_COMPLETE, CREATE_FAILED, or DELETE_FAILED.
//
//	    	- The execution status is UNAVAILABLE or AVAILABLE.
//
//		- After a change set is executed, other change sets associated with the same stack as this change set are also deleted.
//
//		- After a stack is deleted, change sets associated with the stack are deleted.
//
//		- If a change set of the CREATE type is deleted, you must delete stacks associated with the change set.
//
// In this example, a change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is deleted. The change set is created in the China (Hangzhou) region.
//
// @param request - DeleteChangeSetRequest
//
// @return DeleteChangeSetResponse
func (client *Client) DeleteChangeSet(request *DeleteChangeSetRequest) (_result *DeleteChangeSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteChangeSetResponse{}
	_body, _err := client.DeleteChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a diagnostic record.
//
// @param request - DeleteDiagnosticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDiagnosticResponse
func (client *Client) DeleteDiagnosticWithOptions(request *DeleteDiagnosticRequest, runtime *dara.RuntimeOptions) (_result *DeleteDiagnosticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDiagnostic"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDiagnosticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a diagnostic record.
//
// @param request - DeleteDiagnosticRequest
//
// @return DeleteDiagnosticResponse
func (client *Client) DeleteDiagnostic(request *DeleteDiagnosticRequest) (_result *DeleteDiagnosticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteDiagnosticResponse{}
	_body, _err := client.DeleteDiagnosticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a stack. You can specify whether to retain resources.
//
// @param request - DeleteStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStackResponse
func (client *Client) DeleteStackWithOptions(request *DeleteStackRequest, runtime *dara.RuntimeOptions) (_result *DeleteStackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteOptions) {
		query["DeleteOptions"] = request.DeleteOptions
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetainAllResources) {
		query["RetainAllResources"] = request.RetainAllResources
	}

	if !dara.IsNil(request.RetainResources) {
		query["RetainResources"] = request.RetainResources
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stack. You can specify whether to retain resources.
//
// @param request - DeleteStackRequest
//
// @return DeleteStackResponse
func (client *Client) DeleteStack(request *DeleteStackRequest) (_result *DeleteStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStackResponse{}
	_body, _err := client.DeleteStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a stack group.
//
// Description:
//
// A stack group can be deleted only when the stack group does not contain stacks. You can call the [DeleteStackInstances](https://help.aliyun.com/document_detail/151715.html) operation to delete stacks.
//
// This topic provides an example on how to delete a stack group. In this example, a stack group named `MyStackGroup` in the China (Hangzhou) region is deleted.
//
// @param request - DeleteStackGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStackGroupResponse
func (client *Client) DeleteStackGroupWithOptions(request *DeleteStackGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteStackGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStackGroup"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a stack group.
//
// Description:
//
// A stack group can be deleted only when the stack group does not contain stacks. You can call the [DeleteStackInstances](https://help.aliyun.com/document_detail/151715.html) operation to delete stacks.
//
// This topic provides an example on how to delete a stack group. In this example, a stack group named `MyStackGroup` in the China (Hangzhou) region is deleted.
//
// @param request - DeleteStackGroupRequest
//
// @return DeleteStackGroupResponse
func (client *Client) DeleteStackGroup(request *DeleteStackGroupRequest) (_result *DeleteStackGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStackGroupResponse{}
	_body, _err := client.DeleteStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes stack instances in the specified accounts and regions. You can retain specific resources based on your business requirements when you call this operation.
//
// Description:
//
// In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. In this example, the stacks of the stack group that are deployed in the China (Beijing) region by using the Alibaba Cloud account whose ID is `151266687691****` are deleted.
//
// @param tmpReq - DeleteStackInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStackInstancesResponse
func (client *Client) DeleteStackInstancesWithOptions(tmpReq *DeleteStackInstancesRequest, runtime *dara.RuntimeOptions) (_result *DeleteStackInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("AccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeploymentTargets) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, dara.String("DeploymentTargets"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OperationPreferences) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, dara.String("OperationPreferences"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeploymentTargetsShrink) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !dara.IsNil(request.OperationDescription) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !dara.IsNil(request.OperationPreferencesShrink) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.RetainStacks) {
		query["RetainStacks"] = request.RetainStacks
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteStackInstances"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes stack instances in the specified accounts and regions. You can retain specific resources based on your business requirements when you call this operation.
//
// Description:
//
// In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. In this example, the stacks of the stack group that are deployed in the China (Beijing) region by using the Alibaba Cloud account whose ID is `151266687691****` are deleted.
//
// @param request - DeleteStackInstancesRequest
//
// @return DeleteStackInstancesResponse
func (client *Client) DeleteStackInstances(request *DeleteStackInstancesRequest) (_result *DeleteStackInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteStackInstancesResponse{}
	_body, _err := client.DeleteStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// Description:
//
// If a template is shared with other Alibaba Cloud accounts, you must unshare the template before you delete it.
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// Description:
//
// If a template is shared with other Alibaba Cloud accounts, you must unshare the template before you delete it.
//
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a scenario.
//
// Description:
//
// In this topic, a scenario whose ID is `ts-4f83704400994409****` is deleted in the China (Hangzhou) region.
//
// @param request - DeleteTemplateScratchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateScratchResponse
func (client *Client) DeleteTemplateScratchWithOptions(request *DeleteTemplateScratchRequest, runtime *dara.RuntimeOptions) (_result *DeleteTemplateScratchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplateScratch"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a scenario.
//
// Description:
//
// In this topic, a scenario whose ID is `ts-4f83704400994409****` is deleted in the China (Hangzhou) region.
//
// @param request - DeleteTemplateScratchRequest
//
// @return DeleteTemplateScratchResponse
func (client *Client) DeleteTemplateScratch(request *DeleteTemplateScratchRequest) (_result *DeleteTemplateScratchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteTemplateScratchResponse{}
	_body, _err := client.DeleteTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource type or a version of a resource type.
//
// Description:
//
//	  If you delete a resource type, you can no longer use the resource type in Resource Orchestration Service (ROS).
//
//		- If you delete a version of a resource type, you can no longer use the version in ROS.
//
//		- If a resource type has only one version, you can delete the version by calling the operation. If a resource type has more than one version, you must manually delete the remaining versions.
//
//		- When a resource type has more than one version, you cannot delete the default version by calling the operation.
//
//		- When a resource type has only one version, you can delete the resource type and the version by calling the operation.
//
// @param request - DeregisterResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeregisterResourceTypeResponse
func (client *Client) DeregisterResourceTypeWithOptions(request *DeregisterResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *DeregisterResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeregisterResourceType"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeregisterResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource type or a version of a resource type.
//
// Description:
//
//	  If you delete a resource type, you can no longer use the resource type in Resource Orchestration Service (ROS).
//
//		- If you delete a version of a resource type, you can no longer use the version in ROS.
//
//		- If a resource type has only one version, you can delete the version by calling the operation. If a resource type has more than one version, you must manually delete the remaining versions.
//
//		- When a resource type has more than one version, you cannot delete the default version by calling the operation.
//
//		- When a resource type has only one version, you can delete the resource type and the version by calling the operation.
//
// @param request - DeregisterResourceTypeRequest
//
// @return DeregisterResourceTypeResponse
func (client *Client) DeregisterResourceType(request *DeregisterResourceTypeRequest) (_result *DeregisterResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeregisterResourceTypeResponse{}
	_body, _err := client.DeregisterResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of available regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of available regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to detect drift on a stack.
//
// @param request - DetectStackDriftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectStackDriftResponse
func (client *Client) DetectStackDriftWithOptions(request *DetectStackDriftRequest, runtime *dara.RuntimeOptions) (_result *DetectStackDriftResponse, _err error) {
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

	if !dara.IsNil(request.LogicalResourceId) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectStackDrift"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectStackDriftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to detect drift on a stack.
//
// @param request - DetectStackDriftRequest
//
// @return DetectStackDriftResponse
func (client *Client) DetectStackDrift(request *DetectStackDriftRequest) (_result *DetectStackDriftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectStackDriftResponse{}
	_body, _err := client.DetectStackDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 对资源栈组进行偏差检测
//
// @param tmpReq - DetectStackGroupDriftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectStackGroupDriftResponse
func (client *Client) DetectStackGroupDriftWithOptions(tmpReq *DetectStackGroupDriftRequest, runtime *dara.RuntimeOptions) (_result *DetectStackGroupDriftResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetectStackGroupDriftShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationPreferences) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, dara.String("OperationPreferences"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OperationPreferencesShrink) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectStackGroupDrift"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectStackGroupDriftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 对资源栈组进行偏差检测
//
// @param request - DetectStackGroupDriftRequest
//
// @return DetectStackGroupDriftResponse
func (client *Client) DetectStackGroupDrift(request *DetectStackGroupDriftRequest) (_result *DetectStackGroupDriftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectStackGroupDriftResponse{}
	_body, _err := client.DetectStackGroupDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Performs drift detection on resources in a stack to determine whether the resources have drifted from the expected configurations.
//
// @param request - DetectStackResourceDriftRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectStackResourceDriftResponse
func (client *Client) DetectStackResourceDriftWithOptions(request *DetectStackResourceDriftRequest, runtime *dara.RuntimeOptions) (_result *DetectStackResourceDriftResponse, _err error) {
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

	if !dara.IsNil(request.LogicalResourceId) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DetectStackResourceDrift"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetectStackResourceDriftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Performs drift detection on resources in a stack to determine whether the resources have drifted from the expected configurations.
//
// @param request - DetectStackResourceDriftRequest
//
// @return DetectStackResourceDriftResponse
func (client *Client) DetectStackResourceDrift(request *DetectStackResourceDriftRequest) (_result *DetectStackResourceDriftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetectStackResourceDriftResponse{}
	_body, _err := client.DetectStackResourceDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 开启可信服务访问
//
// @param request - EnableServiceAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableServiceAccessResponse
func (client *Client) EnableServiceAccessWithOptions(runtime *dara.RuntimeOptions) (_result *EnableServiceAccessResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("EnableServiceAccess"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableServiceAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 开启可信服务访问
//
// @return EnableServiceAccessResponse
func (client *Client) EnableServiceAccess() (_result *EnableServiceAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableServiceAccessResponse{}
	_body, _err := client.EnableServiceAccessWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量开通
//
// @param tmpReq - EnableServicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableServicesResponse
func (client *Client) EnableServicesWithOptions(tmpReq *EnableServicesRequest, runtime *dara.RuntimeOptions) (_result *EnableServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EnableServicesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ServiceNames) {
		request.ServiceNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServiceNames, dara.String("ServiceNames"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ServiceNamesShrink) {
		query["ServiceNames"] = request.ServiceNamesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EnableServices"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EnableServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量开通
//
// @param request - EnableServicesRequest
//
// @return EnableServicesResponse
func (client *Client) EnableServices(request *EnableServicesRequest) (_result *EnableServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EnableServicesResponse{}
	_body, _err := client.EnableServicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes change sets.
//
// Description:
//
// In this example, the change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is executed. The change set is created in the `China (Hangzhou)` region.
//
// @param request - ExecuteChangeSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteChangeSetResponse
func (client *Client) ExecuteChangeSetWithOptions(request *ExecuteChangeSetRequest, runtime *dara.RuntimeOptions) (_result *ExecuteChangeSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetId) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteChangeSet"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes change sets.
//
// Description:
//
// In this example, the change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is executed. The change set is created in the `China (Hangzhou)` region.
//
// @param request - ExecuteChangeSetRequest
//
// @return ExecuteChangeSetResponse
func (client *Client) ExecuteChangeSet(request *ExecuteChangeSetRequest) (_result *ExecuteChangeSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteChangeSetResponse{}
	_body, _err := client.ExecuteChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a template for a resource scenario.
//
// Description:
//
// In this example, a template is generated for a resource management scenario that resides in the China (Hangzhou) region. The ID of the resource scenario is `ts-aa9c62feab844a6b****`.
//
// >  You cannot generate a template for a resource detection scenario.
//
// @param request - GenerateTemplateByScratchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTemplateByScratchResponse
func (client *Client) GenerateTemplateByScratchWithOptions(request *GenerateTemplateByScratchRequest, runtime *dara.RuntimeOptions) (_result *GenerateTemplateByScratchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProvisionRegionId) {
		query["ProvisionRegionId"] = request.ProvisionRegionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTemplateByScratch"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTemplateByScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a template for a resource scenario.
//
// Description:
//
// In this example, a template is generated for a resource management scenario that resides in the China (Hangzhou) region. The ID of the resource scenario is `ts-aa9c62feab844a6b****`.
//
// >  You cannot generate a template for a resource detection scenario.
//
// @param request - GenerateTemplateByScratchRequest
//
// @return GenerateTemplateByScratchResponse
func (client *Client) GenerateTemplateByScratch(request *GenerateTemplateByScratchRequest) (_result *GenerateTemplateByScratchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateTemplateByScratchResponse{}
	_body, _err := client.GenerateTemplateByScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates the information about a policy that is required by a template.
//
// Description:
//
// If the policy information is related to Enterprise Distributed Application Service (EDAS), you must log on to your Alibaba Cloud account and grant the required permissions to the relevant RAM users.
//
// In this example, a policy is generated for a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****`.
//
// @param request - GenerateTemplatePolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateTemplatePolicyResponse
func (client *Client) GenerateTemplatePolicyWithOptions(request *GenerateTemplatePolicyRequest, runtime *dara.RuntimeOptions) (_result *GenerateTemplatePolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationTypes) {
		query["OperationTypes"] = request.OperationTypes
	}

	if !dara.IsNil(request.TemplateBody) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateTemplatePolicy"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateTemplatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates the information about a policy that is required by a template.
//
// Description:
//
// If the policy information is related to Enterprise Distributed Application Service (EDAS), you must log on to your Alibaba Cloud account and grant the required permissions to the relevant RAM users.
//
// In this example, a policy is generated for a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****`.
//
// @param request - GenerateTemplatePolicyRequest
//
// @return GenerateTemplatePolicyResponse
func (client *Client) GenerateTemplatePolicy(request *GenerateTemplatePolicyRequest) (_result *GenerateTemplatePolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GenerateTemplatePolicyResponse{}
	_body, _err := client.GenerateTemplatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an AI task by task ID.
//
// @param request - GetAITaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAITaskResponse
func (client *Client) GetAITaskWithOptions(request *GetAITaskRequest, runtime *dara.RuntimeOptions) (_result *GetAITaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutputOption) {
		query["OutputOption"] = request.OutputOption
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAITask"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAITaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an AI task by task ID.
//
// @param request - GetAITaskRequest
//
// @return GetAITaskResponse
func (client *Client) GetAITask(request *GetAITaskRequest) (_result *GetAITaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAITaskResponse{}
	_body, _err := client.GetAITaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries change sets. You can determine whether to query the templates of change sets.
//
// Description:
//
// In this example, the details of a change set whose ID is `4c11658d-bd47-4dd0-ba64-727edc62****` is queried. The change set is created in the China (Hangzhou) region.
//
// @param request - GetChangeSetRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChangeSetResponse
func (client *Client) GetChangeSetWithOptions(request *GetChangeSetRequest, runtime *dara.RuntimeOptions) (_result *GetChangeSetResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetId) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowTemplate) {
		query["ShowTemplate"] = request.ShowTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChangeSet"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries change sets. You can determine whether to query the templates of change sets.
//
// Description:
//
// In this example, the details of a change set whose ID is `4c11658d-bd47-4dd0-ba64-727edc62****` is queried. The change set is created in the China (Hangzhou) region.
//
// @param request - GetChangeSetRequest
//
// @return GetChangeSetResponse
func (client *Client) GetChangeSet(request *GetChangeSetRequest) (_result *GetChangeSetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetChangeSetResponse{}
	_body, _err := client.GetChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the diagnosis details based on a specified diagnostic report ID.
//
// @param request - GetDiagnosticRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDiagnosticResponse
func (client *Client) GetDiagnosticWithOptions(request *GetDiagnosticRequest, runtime *dara.RuntimeOptions) (_result *GetDiagnosticResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ReportId) {
		query["ReportId"] = request.ReportId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDiagnostic"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDiagnosticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the diagnosis details based on a specified diagnostic report ID.
//
// @param request - GetDiagnosticRequest
//
// @return GetDiagnosticResponse
func (client *Client) GetDiagnostic(request *GetDiagnosticRequest) (_result *GetDiagnosticResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDiagnosticResponse{}
	_body, _err := client.GetDiagnosticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of features that are supported by Resource Orchestration Service (ROS).
//
// Description:
//
// You can call this operation to query the Terraform hosting, resource cleaner, and scenario features.
//
// This topic provides an example on how to query the details of features supported by ROS in the China (Hangzhou) region. The details include Terraform versions, provider versions, and supported resource types.
//
// >  In the Examples section, only part of the sample code is provided.
//
// @param request - GetFeatureDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFeatureDetailsResponse
func (client *Client) GetFeatureDetailsWithOptions(request *GetFeatureDetailsRequest, runtime *dara.RuntimeOptions) (_result *GetFeatureDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Feature) {
		query["Feature"] = request.Feature
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFeatureDetails"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFeatureDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of features that are supported by Resource Orchestration Service (ROS).
//
// Description:
//
// You can call this operation to query the Terraform hosting, resource cleaner, and scenario features.
//
// This topic provides an example on how to query the details of features supported by ROS in the China (Hangzhou) region. The details include Terraform versions, provider versions, and supported resource types.
//
// >  In the Examples section, only part of the sample code is provided.
//
// @param request - GetFeatureDetailsRequest
//
// @return GetFeatureDetailsResponse
func (client *Client) GetFeatureDetails(request *GetFeatureDetailsRequest) (_result *GetFeatureDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetFeatureDetailsResponse{}
	_body, _err := client.GetFeatureDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the details of `ALIYUN::ROS::WaitConditionHandle`.
//
// Description:
//
// For more information about common request parameters, see [Common parameters](https://help.aliyun.com/document_detail/131957.html).
//
// @param request - GetResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceTypeWithOptions(request *GetResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceType"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the details of `ALIYUN::ROS::WaitConditionHandle`.
//
// Description:
//
// For more information about common request parameters, see [Common parameters](https://help.aliyun.com/document_detail/131957.html).
//
// @param request - GetResourceTypeRequest
//
// @return GetResourceTypeResponse
func (client *Client) GetResourceType(request *GetResourceTypeRequest) (_result *GetResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.GetResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates a sample template based on a resource type.
//
// @param request - GetResourceTypeTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceTypeTemplateResponse
func (client *Client) GetResourceTypeTemplateWithOptions(request *GetResourceTypeTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetResourceTypeTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetResourceTypeTemplate"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetResourceTypeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates a sample template based on a resource type.
//
// @param request - GetResourceTypeTemplateRequest
//
// @return GetResourceTypeTemplateResponse
func (client *Client) GetResourceTypeTemplate(request *GetResourceTypeTemplateRequest) (_result *GetResourceTypeTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetResourceTypeTemplateResponse{}
	_body, _err := client.GetResourceTypeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询可信服务
//
// @param request - GetServiceAccessRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceAccessResponse
func (client *Client) GetServiceAccessWithOptions(runtime *dara.RuntimeOptions) (_result *GetServiceAccessResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceAccess"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询可信服务
//
// @return GetServiceAccessResponse
func (client *Client) GetServiceAccess() (_result *GetServiceAccessResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceAccessResponse{}
	_body, _err := client.GetServiceAccessWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the activation status and the RAM roles of an Alibaba Cloud service.
//
// Description:
//
// ### Description
//
// This topic describes how to query the activation status and the RAM roles of an Alibaba Cloud service. In this example, the Elastic High Performance Computing (E-HPC) service that is deployed in the China (Hangzhou) region is queried.
//
// > Make sure that you have the permissions to call the [GetRole](https://help.aliyun.com/document_detail/28711.html) operation.
//
// @param request - GetServiceProvisionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisionsWithOptions(request *GetServiceProvisionsRequest, runtime *dara.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.Services) {
		query["Services"] = request.Services
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetServiceProvisions"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the activation status and the RAM roles of an Alibaba Cloud service.
//
// Description:
//
// ### Description
//
// This topic describes how to query the activation status and the RAM roles of an Alibaba Cloud service. In this example, the Elastic High Performance Computing (E-HPC) service that is deployed in the China (Hangzhou) region is queried.
//
// > Make sure that you have the permissions to call the [GetRole](https://help.aliyun.com/document_detail/28711.html) operation.
//
// @param request - GetServiceProvisionsRequest
//
// @return GetServiceProvisionsResponse
func (client *Client) GetServiceProvisions(request *GetServiceProvisionsRequest) (_result *GetServiceProvisionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.GetServiceProvisionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a stack in Resource Orchestration Service (ROS).
//
// Description:
//
// In this example, the information about a stack whose ID is `c754d2a4-28f1-46df-b557-9586173a****` in the China (Hangzhou) region is queried.
//
// @param request - GetStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackResponse
func (client *Client) GetStackWithOptions(request *GetStackRequest, runtime *dara.RuntimeOptions) (_result *GetStackResponse, _err error) {
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

	if !dara.IsNil(request.LogOption) {
		query["LogOption"] = request.LogOption
	}

	if !dara.IsNil(request.OutputOption) {
		query["OutputOption"] = request.OutputOption
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowResourceProgress) {
		query["ShowResourceProgress"] = request.ShowResourceProgress
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a stack in Resource Orchestration Service (ROS).
//
// Description:
//
// In this example, the information about a stack whose ID is `c754d2a4-28f1-46df-b557-9586173a****` in the China (Hangzhou) region is queried.
//
// @param request - GetStackRequest
//
// @return GetStackResponse
func (client *Client) GetStack(request *GetStackRequest) (_result *GetStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackResponse{}
	_body, _err := client.GetStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the drift detection status of a stack.
//
// Description:
//
// In this topic, the status of a drift detection operation whose ID is `a7044f0d-6f2e-4128-a307-4524ef88****` is queried. The operation is performed in the China (Hangzhou) region.
//
// @param request - GetStackDriftDetectionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackDriftDetectionStatusResponse
func (client *Client) GetStackDriftDetectionStatusWithOptions(request *GetStackDriftDetectionStatusRequest, runtime *dara.RuntimeOptions) (_result *GetStackDriftDetectionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DriftDetectionId) {
		query["DriftDetectionId"] = request.DriftDetectionId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackDriftDetectionStatus"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackDriftDetectionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the drift detection status of a stack.
//
// Description:
//
// In this topic, the status of a drift detection operation whose ID is `a7044f0d-6f2e-4128-a307-4524ef88****` is queried. The operation is performed in the China (Hangzhou) region.
//
// @param request - GetStackDriftDetectionStatusRequest
//
// @return GetStackDriftDetectionStatusResponse
func (client *Client) GetStackDriftDetectionStatus(request *GetStackDriftDetectionStatusRequest) (_result *GetStackDriftDetectionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackDriftDetectionStatusResponse{}
	_body, _err := client.GetStackDriftDetectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// In this example, the information about a stack group named \\`MyStackGroup\\` is queried. The stack group is granted self-managed permissions and created in the China (Hangzhou) region.
//
// Description:
//
// For more information about common request parameters, see [Common parameters](https://help.aliyun.com/document_detail/131957.html).
//
// @param request - GetStackGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackGroupResponse
func (client *Client) GetStackGroupWithOptions(request *GetStackGroupRequest, runtime *dara.RuntimeOptions) (_result *GetStackGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupId) {
		query["StackGroupId"] = request.StackGroupId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackGroup"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// In this example, the information about a stack group named \\`MyStackGroup\\` is queried. The stack group is granted self-managed permissions and created in the China (Hangzhou) region.
//
// Description:
//
// For more information about common request parameters, see [Common parameters](https://help.aliyun.com/document_detail/131957.html).
//
// @param request - GetStackGroupRequest
//
// @return GetStackGroupResponse
func (client *Client) GetStackGroup(request *GetStackGroupRequest) (_result *GetStackGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackGroupResponse{}
	_body, _err := client.GetStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a stack group operation in an Alibaba Cloud region.
//
// Description:
//
// In this example, the information about the stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863d****` is queried. The stack group named `MyStackGroup` is granted self-managed permissions and deployed in the China (Hangzhou) region.
//
// @param request - GetStackGroupOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackGroupOperationResponse
func (client *Client) GetStackGroupOperationWithOptions(request *GetStackGroupOperationRequest, runtime *dara.RuntimeOptions) (_result *GetStackGroupOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackGroupOperation"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackGroupOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a stack group operation in an Alibaba Cloud region.
//
// Description:
//
// In this example, the information about the stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863d****` is queried. The stack group named `MyStackGroup` is granted self-managed permissions and deployed in the China (Hangzhou) region.
//
// @param request - GetStackGroupOperationRequest
//
// @return GetStackGroupOperationResponse
func (client *Client) GetStackGroupOperation(request *GetStackGroupOperationRequest) (_result *GetStackGroupOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackGroupOperationResponse{}
	_body, _err := client.GetStackGroupOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a stack instance that is associated with a stack group.
//
// Description:
//
// In this example, the information about a stack instance associated with a stack group named `MyStackGroup` is queried. The stack instance is deployed in the China (Beijing) region within the `151266687691****` Alibaba Cloud account. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
//
// @param request - GetStackInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackInstanceResponse
func (client *Client) GetStackInstanceWithOptions(request *GetStackInstanceRequest, runtime *dara.RuntimeOptions) (_result *GetStackInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OutputOption) {
		query["OutputOption"] = request.OutputOption
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.StackInstanceAccountId) {
		query["StackInstanceAccountId"] = request.StackInstanceAccountId
	}

	if !dara.IsNil(request.StackInstanceRegionId) {
		query["StackInstanceRegionId"] = request.StackInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackInstance"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a stack instance that is associated with a stack group.
//
// Description:
//
// In this example, the information about a stack instance associated with a stack group named `MyStackGroup` is queried. The stack instance is deployed in the China (Beijing) region within the `151266687691****` Alibaba Cloud account. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
//
// @param request - GetStackInstanceRequest
//
// @return GetStackInstanceResponse
func (client *Client) GetStackInstance(request *GetStackInstanceRequest) (_result *GetStackInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackInstanceResponse{}
	_body, _err := client.GetStackInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query information about a stack policy.
//
// Description:
//
// In this example, the stack policy of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is queried. The stack is deployed in the China (Hangzhou) region.
//
// @param request - GetStackPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackPolicyResponse
func (client *Client) GetStackPolicyWithOptions(request *GetStackPolicyRequest, runtime *dara.RuntimeOptions) (_result *GetStackPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackPolicy"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query information about a stack policy.
//
// Description:
//
// In this example, the stack policy of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is queried. The stack is deployed in the China (Hangzhou) region.
//
// @param request - GetStackPolicyRequest
//
// @return GetStackPolicyResponse
func (client *Client) GetStackPolicy(request *GetStackPolicyRequest) (_result *GetStackPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackPolicyResponse{}
	_body, _err := client.GetStackPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// For more information about common request parameters, see [Common parameters]\\(~~131957~~).
//
// Description:
//
// | Http status code | Error code | Error message | Description |
//
// | ---------------- | ---------- | ------------- | ----------- |
//
// | 404 | ResourceNotFound | The Resource ({name}) could not be found in Stack {stack}. | The error message returned because the specified resource does not exist in the stack. name indicates the resource name. stack indicates the stack name or ID. |
//
// | 404 | StackNotFound | The Stack ({name}) could not be found. | The error message returned because the stack does not exist. name indicates the name or ID of the stack. |
//
// @param request - GetStackResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStackResourceResponse
func (client *Client) GetStackResourceWithOptions(request *GetStackResourceRequest, runtime *dara.RuntimeOptions) (_result *GetStackResourceResponse, _err error) {
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

	if !dara.IsNil(request.LogicalResourceId) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceAttributes) {
		query["ResourceAttributes"] = request.ResourceAttributes
	}

	if !dara.IsNil(request.ShowResourceAttributes) {
		query["ShowResourceAttributes"] = request.ShowResourceAttributes
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStackResource"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStackResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For more information about common request parameters, see [Common parameters]\\(~~131957~~).
//
// Description:
//
// | Http status code | Error code | Error message | Description |
//
// | ---------------- | ---------- | ------------- | ----------- |
//
// | 404 | ResourceNotFound | The Resource ({name}) could not be found in Stack {stack}. | The error message returned because the specified resource does not exist in the stack. name indicates the resource name. stack indicates the stack name or ID. |
//
// | 404 | StackNotFound | The Stack ({name}) could not be found. | The error message returned because the stack does not exist. name indicates the name or ID of the stack. |
//
// @param request - GetStackResourceRequest
//
// @return GetStackResourceResponse
func (client *Client) GetStackResource(request *GetStackResourceRequest) (_result *GetStackResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStackResourceResponse{}
	_body, _err := client.GetStackResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a template based on stacks, stack groups, change sets, or any custom template information.
//
// Description:
//
// In this example, the details of a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is queried. The region ID of the template is `cn-hangzhou`.
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetId) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !dara.IsNil(request.IncludePermission) {
		query["IncludePermission"] = request.IncludePermission
	}

	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateStage) {
		query["TemplateStage"] = request.TemplateStage
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a template based on stacks, stack groups, change sets, or any custom template information.
//
// Description:
//
// In this example, the details of a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is queried. The region ID of the template is `cn-hangzhou`.
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the estimated prices of the resources in a template.
//
// Description:
//
// ###
//
//   - For more information about the resources that support price inquiry in Resource Orchestration Service (ROS) templates, see the "**Resource types that support price inquiry**" section of the [Estimate resource prices](https://help.aliyun.com/document_detail/203165.html) topic.
//
//   - For more information about the resources that support price inquiry in Terraform templates, see the "**ROS resources supported by Terraform**" section of the [ROS features and resources supported by Terraform](https://help.aliyun.com/document_detail/184389.html) topic.
//
// The following sample code provides an example on how to query the estimated price of an elastic IP address (EIP) that you want to create based on a template. In this example, the following template is used:
//
//	{
//
//	  "ROSTemplateFormatVersion": "2015-09-01",
//
//	  "Parameters": {
//
//	    "Isp": {
//
//	      "Type": "String",
//
//	      "Default": "BGP"
//
//	    },
//
//	    "Name": {
//
//	      "Type": "String",
//
//	      "Default": "test"
//
//	    },
//
//	    "Netmode": {
//
//	      "Type": "String",
//
//	      "Default": "public"
//
//	    },
//
//	    "Bandwidth": {
//
//	      "Type": "Number",
//
//	      "Default": 5
//
//	    }
//
//	  },
//
//	  "Resources": {
//
//	    "NewEip": {
//
//	      "Type": "ALIYUN::VPC::EIP",
//
//	      "Properties": {
//
//	        "InstanceChargeType": "Prepaid",
//
//	        "PricingCycle": "Month",
//
//	        "Isp": {
//
//	          "Ref": "Isp"
//
//	        },
//
//	        "Period": 1,
//
//	        "DeletionProtection": false,
//
//	        "AutoPay": false,
//
//	        "Name": {
//
//	          "Ref": "Name"
//
//	        },
//
//	        "InternetChargeType": "PayByTraffic",
//
//	        "Netmode": {
//
//	          "Ref": "Netmode"
//
//	        },
//
//	        "Bandwidth": {
//
//	          "Ref": "Bandwidth"
//
//	        }
//
//	      }
//
//	    }
//
//	  }
//
//	}
//
// @param request - GetTemplateEstimateCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateEstimateCostResponse
func (client *Client) GetTemplateEstimateCostWithOptions(request *GetTemplateEstimateCostRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateEstimateCostResponse, _err error) {
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

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !dara.IsNil(request.TemplateScratchRegionId) {
		query["TemplateScratchRegionId"] = request.TemplateScratchRegionId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateEstimateCost"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the estimated prices of the resources in a template.
//
// Description:
//
// ###
//
//   - For more information about the resources that support price inquiry in Resource Orchestration Service (ROS) templates, see the "**Resource types that support price inquiry**" section of the [Estimate resource prices](https://help.aliyun.com/document_detail/203165.html) topic.
//
//   - For more information about the resources that support price inquiry in Terraform templates, see the "**ROS resources supported by Terraform**" section of the [ROS features and resources supported by Terraform](https://help.aliyun.com/document_detail/184389.html) topic.
//
// The following sample code provides an example on how to query the estimated price of an elastic IP address (EIP) that you want to create based on a template. In this example, the following template is used:
//
//	{
//
//	  "ROSTemplateFormatVersion": "2015-09-01",
//
//	  "Parameters": {
//
//	    "Isp": {
//
//	      "Type": "String",
//
//	      "Default": "BGP"
//
//	    },
//
//	    "Name": {
//
//	      "Type": "String",
//
//	      "Default": "test"
//
//	    },
//
//	    "Netmode": {
//
//	      "Type": "String",
//
//	      "Default": "public"
//
//	    },
//
//	    "Bandwidth": {
//
//	      "Type": "Number",
//
//	      "Default": 5
//
//	    }
//
//	  },
//
//	  "Resources": {
//
//	    "NewEip": {
//
//	      "Type": "ALIYUN::VPC::EIP",
//
//	      "Properties": {
//
//	        "InstanceChargeType": "Prepaid",
//
//	        "PricingCycle": "Month",
//
//	        "Isp": {
//
//	          "Ref": "Isp"
//
//	        },
//
//	        "Period": 1,
//
//	        "DeletionProtection": false,
//
//	        "AutoPay": false,
//
//	        "Name": {
//
//	          "Ref": "Name"
//
//	        },
//
//	        "InternetChargeType": "PayByTraffic",
//
//	        "Netmode": {
//
//	          "Ref": "Netmode"
//
//	        },
//
//	        "Bandwidth": {
//
//	          "Ref": "Bandwidth"
//
//	        }
//
//	      }
//
//	    }
//
//	  }
//
//	}
//
// @param request - GetTemplateEstimateCostRequest
//
// @return GetTemplateEstimateCostResponse
func (client *Client) GetTemplateEstimateCost(request *GetTemplateEstimateCostRequest) (_result *GetTemplateEstimateCostResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateEstimateCostResponse{}
	_body, _err := client.GetTemplateEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the values of one or more parameters in a template.
//
// Description:
//
// This topic provides an example on how to query the values of a parameter. In this example, the values of the `ZoneInfo` parameter in a template that is created in the China (Hangzhou) region are queried. The template body is `{"Parameters":{"ZoneInfo":{"Type": "String"},"InstanceType": {"Type": "String"}},"ROSTemplateFormatVersion": "2015-09-01","Resources":{"ECS":{"Properties":{"ZoneId":{"Ref": "ZoneInfo"},"InstanceType": {"Ref": "InstanceType"}},"Type": "ALIYUN::ECS::Instance"}}}`.
//
// For more information about the template parameters whose values you can query by calling this operation and the sample code of the template, see [Query the constraints of parameters](https://help.aliyun.com/document_detail/432820.html).
//
// @param tmpReq - GetTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateParameterConstraintsResponse
func (client *Client) GetTemplateParameterConstraintsWithOptions(tmpReq *GetTemplateParameterConstraintsRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateParameterConstraintsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GetTemplateParameterConstraintsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ParametersKeyFilter) {
		request.ParametersKeyFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParametersKeyFilter, dara.String("ParametersKeyFilter"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ParametersOrder) {
		request.ParametersOrderShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParametersOrder, dara.String("ParametersOrder"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.ParametersKeyFilterShrink) {
		query["ParametersKeyFilter"] = request.ParametersKeyFilterShrink
	}

	if !dara.IsNil(request.ParametersOrderShrink) {
		query["ParametersOrder"] = request.ParametersOrderShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateParameterConstraints"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the values of one or more parameters in a template.
//
// Description:
//
// This topic provides an example on how to query the values of a parameter. In this example, the values of the `ZoneInfo` parameter in a template that is created in the China (Hangzhou) region are queried. The template body is `{"Parameters":{"ZoneInfo":{"Type": "String"},"InstanceType": {"Type": "String"}},"ROSTemplateFormatVersion": "2015-09-01","Resources":{"ECS":{"Properties":{"ZoneId":{"Ref": "ZoneInfo"},"InstanceType": {"Ref": "InstanceType"}},"Type": "ALIYUN::ECS::Instance"}}}`.
//
// For more information about the template parameters whose values you can query by calling this operation and the sample code of the template, see [Query the constraints of parameters](https://help.aliyun.com/document_detail/432820.html).
//
// @param request - GetTemplateParameterConstraintsRequest
//
// @return GetTemplateParameterConstraintsResponse
func (client *Client) GetTemplateParameterConstraints(request *GetTemplateParameterConstraintsRequest) (_result *GetTemplateParameterConstraintsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateParameterConstraintsResponse{}
	_body, _err := client.GetTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 推荐参数
//
// @param request - GetTemplateRecommendParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateRecommendParametersResponse
func (client *Client) GetTemplateRecommendParametersWithOptions(request *GetTemplateRecommendParametersRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateRecommendParametersResponse, _err error) {
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

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateBody) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateRecommendParameters"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateRecommendParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 推荐参数
//
// @param request - GetTemplateRecommendParametersRequest
//
// @return GetTemplateRecommendParametersResponse
func (client *Client) GetTemplateRecommendParameters(request *GetTemplateRecommendParametersRequest) (_result *GetTemplateRecommendParametersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateRecommendParametersResponse{}
	_body, _err := client.GetTemplateRecommendParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a resource scenario.
//
// Description:
//
// In this example, the details of the resource scenario whose ID is `ts-7f7a704cf71c49a6****` is queried. In the response, the source node data is displayed.
//
// @param request - GetTemplateScratchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateScratchResponse
func (client *Client) GetTemplateScratchWithOptions(request *GetTemplateScratchRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateScratchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ShowDataOption) {
		query["ShowDataOption"] = request.ShowDataOption
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateScratch"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a resource scenario.
//
// Description:
//
// In this example, the details of the resource scenario whose ID is `ts-7f7a704cf71c49a6****` is queried. In the response, the source node data is displayed.
//
// @param request - GetTemplateScratchRequest
//
// @return GetTemplateScratchResponse
func (client *Client) GetTemplateScratch(request *GetTemplateScratchRequest) (_result *GetTemplateScratchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateScratchResponse{}
	_body, _err := client.GetTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a template resource by using the relevant template, stack, stack group, or change set.
//
// @param request - GetTemplateSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateSummaryResponse
func (client *Client) GetTemplateSummaryWithOptions(request *GetTemplateSummaryRequest, runtime *dara.RuntimeOptions) (_result *GetTemplateSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetId) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateBody) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplateSummary"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a template resource by using the relevant template, stack, stack group, or change set.
//
// @param request - GetTemplateSummaryRequest
//
// @return GetTemplateSummaryResponse
func (client *Client) GetTemplateSummary(request *GetTemplateSummaryRequest) (_result *GetTemplateSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemplateSummaryResponse{}
	_body, _err := client.GetTemplateSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Import stacks from multiple different accounts into a stack group.
//
// @param tmpReq - ImportStacksToStackGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportStacksToStackGroupResponse
func (client *Client) ImportStacksToStackGroupWithOptions(tmpReq *ImportStacksToStackGroupRequest, runtime *dara.RuntimeOptions) (_result *ImportStacksToStackGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportStacksToStackGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationPreferences) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, dara.String("OperationPreferences"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ResourceDirectoryFolderIds) {
		request.ResourceDirectoryFolderIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceDirectoryFolderIds, dara.String("ResourceDirectoryFolderIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StackArns) {
		request.StackArnsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StackArns, dara.String("StackArns"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.OperationDescription) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !dara.IsNil(request.OperationPreferencesShrink) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceDirectoryFolderIdsShrink) {
		query["ResourceDirectoryFolderIds"] = request.ResourceDirectoryFolderIdsShrink
	}

	if !dara.IsNil(request.StackArnsShrink) {
		query["StackArns"] = request.StackArnsShrink
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportStacksToStackGroup"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportStacksToStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Import stacks from multiple different accounts into a stack group.
//
// @param request - ImportStacksToStackGroupRequest
//
// @return ImportStacksToStackGroupResponse
func (client *Client) ImportStacksToStackGroup(request *ImportStacksToStackGroupRequest) (_result *ImportStacksToStackGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportStacksToStackGroupResponse{}
	_body, _err := client.ImportStacksToStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the events of an AI task.
//
// @param request - ListAITaskEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAITaskEventsResponse
func (client *Client) ListAITaskEventsWithOptions(request *ListAITaskEventsRequest, runtime *dara.RuntimeOptions) (_result *ListAITaskEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAITaskEvents"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAITaskEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the events of an AI task.
//
// @param request - ListAITaskEventsRequest
//
// @return ListAITaskEventsResponse
func (client *Client) ListAITaskEvents(request *ListAITaskEventsRequest) (_result *ListAITaskEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAITaskEventsResponse{}
	_body, _err := client.ListAITaskEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AI tasks.
//
// @param request - ListAITasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAITasksResponse
func (client *Client) ListAITasksWithOptions(request *ListAITasksRequest, runtime *dara.RuntimeOptions) (_result *ListAITasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAITasks"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAITasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of AI tasks.
//
// @param request - ListAITasksRequest
//
// @return ListAITasksResponse
func (client *Client) ListAITasks(request *ListAITasksRequest) (_result *ListAITasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAITasksResponse{}
	_body, _err := client.ListAITasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries change sets.
//
// @param request - ListChangeSetsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChangeSetsResponse
func (client *Client) ListChangeSetsWithOptions(request *ListChangeSetsRequest, runtime *dara.RuntimeOptions) (_result *ListChangeSetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChangeSetId) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !dara.IsNil(request.ChangeSetName) {
		query["ChangeSetName"] = request.ChangeSetName
	}

	if !dara.IsNil(request.ExecutionStatus) {
		query["ExecutionStatus"] = request.ExecutionStatus
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListChangeSets"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListChangeSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries change sets.
//
// @param request - ListChangeSetsRequest
//
// @return ListChangeSetsResponse
func (client *Client) ListChangeSets(request *ListChangeSetsRequest) (_result *ListChangeSetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListChangeSetsResponse{}
	_body, _err := client.ListChangeSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a diagnostic report.
//
// @param request - ListDiagnosticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosticsResponse
func (client *Client) ListDiagnosticsWithOptions(request *ListDiagnosticsRequest, runtime *dara.RuntimeOptions) (_result *ListDiagnosticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DiagnosticKey) {
		query["DiagnosticKey"] = request.DiagnosticKey
	}

	if !dara.IsNil(request.DiagnosticProduct) {
		query["DiagnosticProduct"] = request.DiagnosticProduct
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnostics"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDiagnosticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a diagnostic report.
//
// @param request - ListDiagnosticsRequest
//
// @return ListDiagnosticsResponse
func (client *Client) ListDiagnostics(request *ListDiagnosticsRequest) (_result *ListDiagnosticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListDiagnosticsResponse{}
	_body, _err := client.ListDiagnosticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the registration records of a resource.
//
// @param request - ListResourceTypeRegistrationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypeRegistrationsResponse
func (client *Client) ListResourceTypeRegistrationsWithOptions(request *ListResourceTypeRegistrationsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTypeRegistrationsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegistrationId) {
		query["RegistrationId"] = request.RegistrationId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypeRegistrations"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypeRegistrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the registration records of a resource.
//
// @param request - ListResourceTypeRegistrationsRequest
//
// @return ListResourceTypeRegistrationsResponse
func (client *Client) ListResourceTypeRegistrations(request *ListResourceTypeRegistrationsRequest) (_result *ListResourceTypeRegistrationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceTypeRegistrationsResponse{}
	_body, _err := client.ListResourceTypeRegistrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of resource types, including the resource types created by you and provided by Resource Orchestration Service (ROS).
//
// @param request - ListResourceTypeVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypeVersionsResponse
func (client *Client) ListResourceTypeVersionsWithOptions(request *ListResourceTypeVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTypeVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypeVersions"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypeVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of resource types, including the resource types created by you and provided by Resource Orchestration Service (ROS).
//
// @param request - ListResourceTypeVersionsRequest
//
// @return ListResourceTypeVersionsResponse
func (client *Client) ListResourceTypeVersions(request *ListResourceTypeVersionsRequest) (_result *ListResourceTypeVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceTypeVersionsResponse{}
	_body, _err := client.ListResourceTypeVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the list of resource types supported by Resource Orchestration Service (ROS).
//
// Description:
//
// For more information about errors common to all operations, see [Common error codes](/help/en/resource-orchestration-service/latest/common-error-codes).
//
// @param request - ListResourceTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypesWithOptions(request *ListResourceTypesRequest, runtime *dara.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.Provider) {
		query["Provider"] = request.Provider
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListResourceTypes"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the list of resource types supported by Resource Orchestration Service (ROS).
//
// Description:
//
// For more information about errors common to all operations, see [Common error codes](/help/en/resource-orchestration-service/latest/common-error-codes).
//
// @param request - ListResourceTypesRequest
//
// @return ListResourceTypesResponse
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a stack and the resource events of the stack.
//
// @param request - ListStackEventsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackEventsResponse
func (client *Client) ListStackEventsWithOptions(request *ListStackEventsRequest, runtime *dara.RuntimeOptions) (_result *ListStackEventsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogicalResourceId) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackEvents"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a stack and the resource events of the stack.
//
// @param request - ListStackEventsRequest
//
// @return ListStackEventsResponse
func (client *Client) ListStackEvents(request *ListStackEventsRequest) (_result *ListStackEventsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackEventsResponse{}
	_body, _err := client.ListStackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of an operation on a stack group.
//
// Description:
//
// In this example, the operation ID `6da106ca-1784-4a6f-a7e1-e723863d∗∗∗∗` is set to query the results of an operation on a stack group named `MyStackGroup`. The stack group is granted self-managed permissions and created in the China (Hangzhou) region.
//
// @param request - ListStackGroupOperationResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackGroupOperationResultsResponse
func (client *Client) ListStackGroupOperationResultsWithOptions(request *ListStackGroupOperationResultsRequest, runtime *dara.RuntimeOptions) (_result *ListStackGroupOperationResultsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackGroupOperationResults"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackGroupOperationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of an operation on a stack group.
//
// Description:
//
// In this example, the operation ID `6da106ca-1784-4a6f-a7e1-e723863d∗∗∗∗` is set to query the results of an operation on a stack group named `MyStackGroup`. The stack group is granted self-managed permissions and created in the China (Hangzhou) region.
//
// @param request - ListStackGroupOperationResultsRequest
//
// @return ListStackGroupOperationResultsResponse
func (client *Client) ListStackGroupOperationResults(request *ListStackGroupOperationResultsRequest) (_result *ListStackGroupOperationResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackGroupOperationResultsResponse{}
	_body, _err := client.ListStackGroupOperationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about stack group operations in an Alibaba Cloud region.
//
// @param request - ListStackGroupOperationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackGroupOperationsResponse
func (client *Client) ListStackGroupOperationsWithOptions(request *ListStackGroupOperationsRequest, runtime *dara.RuntimeOptions) (_result *ListStackGroupOperationsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackGroupOperations"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackGroupOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about stack group operations in an Alibaba Cloud region.
//
// @param request - ListStackGroupOperationsRequest
//
// @return ListStackGroupOperationsResponse
func (client *Client) ListStackGroupOperations(request *ListStackGroupOperationsRequest) (_result *ListStackGroupOperationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackGroupOperationsResponse{}
	_body, _err := client.ListStackGroupOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of stack groups in an Alibaba Cloud region.
//
// Description:
//
// In this example, the list of stack groups that are in the ACTIVE state and deployed in the China (Hangzhou) region is queried.
//
// @param request - ListStackGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackGroupsResponse
func (client *Client) ListStackGroupsWithOptions(request *ListStackGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListStackGroupsResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackGroups"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of stack groups in an Alibaba Cloud region.
//
// Description:
//
// In this example, the list of stack groups that are in the ACTIVE state and deployed in the China (Hangzhou) region is queried.
//
// @param request - ListStackGroupsRequest
//
// @return ListStackGroupsResponse
func (client *Client) ListStackGroups(request *ListStackGroupsRequest) (_result *ListStackGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackGroupsResponse{}
	_body, _err := client.ListStackGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of stack instances that are associated with a stack group in an Alibaba Cloud region.
//
// Description:
//
// In this example, the list of stack instances that are associated with a stack group named `MyStackGroup` is queried. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
//
// @param request - ListStackInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackInstancesResponse
func (client *Client) ListStackInstancesWithOptions(request *ListStackInstancesRequest, runtime *dara.RuntimeOptions) (_result *ListStackInstancesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.StackInstanceAccountId) {
		query["StackInstanceAccountId"] = request.StackInstanceAccountId
	}

	if !dara.IsNil(request.StackInstanceRegionId) {
		query["StackInstanceRegionId"] = request.StackInstanceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackInstances"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of stack instances that are associated with a stack group in an Alibaba Cloud region.
//
// Description:
//
// In this example, the list of stack instances that are associated with a stack group named `MyStackGroup` is queried. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
//
// @param request - ListStackInstancesRequest
//
// @return ListStackInstancesResponse
func (client *Client) ListStackInstances(request *ListStackInstancesRequest) (_result *ListStackInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackInstancesResponse{}
	_body, _err := client.ListStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects stack-related operation risks and returns missing permissions and the causes of the risks.
//
// Description:
//
// The ListStackOperationRisks operation is suitable for the following scenarios:
//
//   - You want to detect high risks that may arise in resources when you delete a stack that contains the resources, and query the cause of each risk in a resource.
//
//   - When you create a stack, the creation may fail. In this case, you can call this operation to check which types of permissions that are required to create stacks are missing.
//
// @param request - ListStackOperationRisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackOperationRisksResponse
func (client *Client) ListStackOperationRisksWithOptions(request *ListStackOperationRisksRequest, runtime *dara.RuntimeOptions) (_result *ListStackOperationRisksResponse, _err error) {
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

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RetainAllResources) {
		query["RetainAllResources"] = request.RetainAllResources
	}

	if !dara.IsNil(request.RetainResources) {
		query["RetainResources"] = request.RetainResources
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackOperationRisks"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackOperationRisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects stack-related operation risks and returns missing permissions and the causes of the risks.
//
// Description:
//
// The ListStackOperationRisks operation is suitable for the following scenarios:
//
//   - You want to detect high risks that may arise in resources when you delete a stack that contains the resources, and query the cause of each risk in a resource.
//
//   - When you create a stack, the creation may fail. In this case, you can call this operation to check which types of permissions that are required to create stacks are missing.
//
// @param request - ListStackOperationRisksRequest
//
// @return ListStackOperationRisksResponse
func (client *Client) ListStackOperationRisks(request *ListStackOperationRisksRequest) (_result *ListStackOperationRisksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackOperationRisksResponse{}
	_body, _err := client.ListStackOperationRisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The query token. Set this parameter to the NextToken value returned in the last API call.
//
// @param request - ListStackResourceDriftsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackResourceDriftsResponse
func (client *Client) ListStackResourceDriftsWithOptions(request *ListStackResourceDriftsRequest, runtime *dara.RuntimeOptions) (_result *ListStackResourceDriftsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceDriftStatus) {
		query["ResourceDriftStatus"] = request.ResourceDriftStatus
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackResourceDrifts"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackResourceDriftsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The query token. Set this parameter to the NextToken value returned in the last API call.
//
// @param request - ListStackResourceDriftsRequest
//
// @return ListStackResourceDriftsResponse
func (client *Client) ListStackResourceDrifts(request *ListStackResourceDriftsRequest) (_result *ListStackResourceDriftsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackResourceDriftsResponse{}
	_body, _err := client.ListStackResourceDriftsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the resources in a specified stack. In this example, the resources in the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` in the China (Hangzhou) region are queried.
//
// Description:
//
// For more information about common request parameters, see [Common parameters](https://help.aliyun.com/document_detail/131957.html).
//
// @param request - ListStackResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStackResourcesResponse
func (client *Client) ListStackResourcesWithOptions(request *ListStackResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListStackResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStackResources"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStackResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This topic provides an example on how to query the resources in a specified stack. In this example, the resources in the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` in the China (Hangzhou) region are queried.
//
// Description:
//
// For more information about common request parameters, see [Common parameters](https://help.aliyun.com/document_detail/131957.html).
//
// @param request - ListStackResourcesRequest
//
// @return ListStackResourcesResponse
func (client *Client) ListStackResources(request *ListStackResourcesRequest) (_result *ListStackResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStackResourcesResponse{}
	_body, _err := client.ListStackResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of stacks.
//
// Description:
//
// ###
//
// This topic provides an example on how to query a list of stacks. In this example, the stacks that are deployed in the China (Hangzhou) region are queried.
//
// @param request - ListStacksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStacksResponse
func (client *Client) ListStacksWithOptions(request *ListStacksRequest, runtime *dara.RuntimeOptions) (_result *ListStacksResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ParentStackId) {
		query["ParentStackId"] = request.ParentStackId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShowNestedStack) {
		query["ShowNestedStack"] = request.ShowNestedStack
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.StackIds) {
		query["StackIds"] = request.StackIds
	}

	if !dara.IsNil(request.StackName) {
		query["StackName"] = request.StackName
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListStacks"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListStacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of stacks.
//
// Description:
//
// ###
//
// This topic provides an example on how to query a list of stacks. In this example, the stacks that are deployed in the China (Hangzhou) region are queried.
//
// @param request - ListStacksRequest
//
// @return ListStacksResponse
func (client *Client) ListStacks(request *ListStacksRequest) (_result *ListStacksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListStacksResponse{}
	_body, _err := client.ListStacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询总览
//
// @param request - ListSummariesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSummariesResponse
func (client *Client) ListSummariesWithOptions(request *ListSummariesRequest, runtime *dara.RuntimeOptions) (_result *ListSummariesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Option) {
		query["Option"] = request.Option
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSummaries"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSummariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询总览
//
// @param request - ListSummariesRequest
//
// @return ListSummariesResponse
func (client *Client) ListSummaries(request *ListSummariesRequest) (_result *ListSummariesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSummariesResponse{}
	_body, _err := client.ListSummariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag keys that are added to resources in a template or stack in an Alibaba Cloud region.
//
// Description:
//
// In this example, the tag keys that are added to a stack in the China (Hangzhou) region are queried.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *dara.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagKeys"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag keys that are added to resources in a template or stack in an Alibaba Cloud region.
//
// Description:
//
// In this example, the tag keys that are added to a stack in the China (Hangzhou) region are queried.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources in a template or stack in an Alibaba Cloud region.
//
// Description:
//
// ###
//
//   - To specify the query object, specify ResourceId or Tag in the request. Tag consists of Key and Value.
//
//   - If you specify Tag and ResourceId, ROS resources that match both the parameters are returned.
//
// This topic provides an example on how to query the tags that are added to a stack. In this example, the stack ID is `6bc589b5-9c02-4944-8fc3-f3624234****`. The stack is deployed in the China (Hangzhou) region.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to resources in a template or stack in an Alibaba Cloud region.
//
// Description:
//
// ###
//
//   - To specify the query object, specify ResourceId or Tag in the request. Tag consists of Key and Value.
//
//   - If you specify Tag and ResourceId, ROS resources that match both the parameters are returned.
//
// This topic provides an example on how to query the tags that are added to a stack. In this example, the stack ID is `6bc589b5-9c02-4944-8fc3-f3624234****`. The stack is deployed in the China (Hangzhou) region.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tag values that are added to resources in a template or stack in an Alibaba Cloud region.
//
// Description:
//
// In this example, the tag values of `TagKey1` that is added to a stack in the China (Hangzhou) region are queried.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *dara.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Key) {
		query["Key"] = request.Key
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagValues"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tag values that are added to resources in a template or stack in an Alibaba Cloud region.
//
// Description:
//
// In this example, the tag values of `TagKey1` that is added to a stack in the China (Hangzhou) region are queried.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries scenarios.
//
// Description:
//
// In this example, the scenarios that are created in the China (Hangzhou) region are queried. In the response, a scenario of the Resource Management and a scenario of the Resource Replication type are returned.
//
// @param request - ListTemplateScratchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateScratchesResponse
func (client *Client) ListTemplateScratchesWithOptions(request *ListTemplateScratchesRequest, runtime *dara.RuntimeOptions) (_result *ListTemplateScratchesResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !dara.IsNil(request.TemplateScratchType) {
		query["TemplateScratchType"] = request.TemplateScratchType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplateScratches"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateScratchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries scenarios.
//
// Description:
//
// In this example, the scenarios that are created in the China (Hangzhou) region are queried. In the response, a scenario of the Resource Management and a scenario of the Resource Replication type are returned.
//
// @param request - ListTemplateScratchesRequest
//
// @return ListTemplateScratchesResponse
func (client *Client) ListTemplateScratches(request *ListTemplateScratchesRequest) (_result *ListTemplateScratchesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplateScratchesResponse{}
	_body, _err := client.ListTemplateScratchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of versions of a template.
//
// @param request - ListTemplateVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateVersionsResponse
func (client *Client) ListTemplateVersionsWithOptions(request *ListTemplateVersionsRequest, runtime *dara.RuntimeOptions) (_result *ListTemplateVersionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplateVersions"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of versions of a template.
//
// @param request - ListTemplateVersionsRequest
//
// @return ListTemplateVersionsResponse
func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (_result *ListTemplateVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.ListTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Templates
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
	if !dara.IsNil(request.Filters) {
		query["Filters"] = request.Filters
	}

	if !dara.IsNil(request.IncludeTags) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ShareType) {
		query["ShareType"] = request.ShareType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2019-09-10"),
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
// # List Templates
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
// Moves a resource to a specific resource group.
//
// Description:
//
// In this example, a stack deployed in the `China (Hangzhou)` region is moved to a specific resource group. The ID of the stack is `4e8611cb-251e-42b7-b9cb-3496362c****` and the ID of the resource group is `rg-acfm3peow3k****`.
//
// @param request - MoveResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NewResourceGroupId) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("MoveResourceGroup"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a resource to a specific resource group.
//
// Description:
//
// In this example, a stack deployed in the `China (Hangzhou)` region is moved to a specific resource group. The ID of the stack is `4e8611cb-251e-42b7-b9cb-3496362c****` and the ID of the resource group is `rg-acfm3peow3k****`.
//
// @param request - MoveResourceGroupRequest
//
// @return MoveResourceGroupResponse
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Previews the information about a stack that you want to create based on a template. You can call this operation to verify whether the template resources are valid.
//
// Description:
//
// This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template and preview the information about the stack. In this example, the `template body` is `{"ROSTemplateFormatVersion":"2015-09-01"}`.
//
// @param request - PreviewStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewStackResponse
func (client *Client) PreviewStackWithOptions(request *PreviewStackRequest, runtime *dara.RuntimeOptions) (_result *PreviewStackResponse, _err error) {
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

	if !dara.IsNil(request.DisableRollback) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.EnablePreConfig) {
		query["EnablePreConfig"] = request.EnablePreConfig
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.StackName) {
		query["StackName"] = request.StackName
	}

	if !dara.IsNil(request.StackPolicyBody) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !dara.IsNil(request.StackPolicyURL) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !dara.IsNil(request.TaintResources) {
		query["TaintResources"] = request.TaintResources
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !dara.IsNil(request.TemplateScratchRegionId) {
		query["TemplateScratchRegionId"] = request.TemplateScratchRegionId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.TimeoutInMinutes) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	if !dara.IsNil(request.UsePreviousParameters) {
		query["UsePreviousParameters"] = request.UsePreviousParameters
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Previews the information about a stack that you want to create based on a template. You can call this operation to verify whether the template resources are valid.
//
// Description:
//
// This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template and preview the information about the stack. In this example, the `template body` is `{"ROSTemplateFormatVersion":"2015-09-01"}`.
//
// @param request - PreviewStackRequest
//
// @return PreviewStackResponse
func (client *Client) PreviewStack(request *PreviewStackRequest) (_result *PreviewStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PreviewStackResponse{}
	_body, _err := client.PreviewStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a new resource type, or creates a new version for an existing resource type.
//
// Description:
//
//	  Versions increase from v1.
//
//		- If you create a new resource type, v1 is used as the default version of the resource type. You can call the SetResourceType operation to change the default version of a resource type.
//
// @param request - RegisterResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterResourceTypeResponse
func (client *Client) RegisterResourceTypeWithOptions(request *RegisterResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *RegisterResourceTypeResponse, _err error) {
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

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EntityType) {
		query["EntityType"] = request.EntityType
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RegisterResourceType"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RegisterResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new resource type, or creates a new version for an existing resource type.
//
// Description:
//
//	  Versions increase from v1.
//
//		- If you create a new resource type, v1 is used as the default version of the resource type. You can call the SetResourceType operation to change the default version of a resource type.
//
// @param request - RegisterResourceTypeRequest
//
// @return RegisterResourceTypeResponse
func (client *Client) RegisterResourceType(request *RegisterResourceTypeRequest) (_result *RegisterResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RegisterResourceTypeResponse{}
	_body, _err := client.RegisterResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改资源栈的删除保护属性
//
// @param request - SetDeletionProtectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtectionWithOptions(request *SetDeletionProtectionRequest, runtime *dara.RuntimeOptions) (_result *SetDeletionProtectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeletionProtection) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetDeletionProtection"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改资源栈的删除保护属性
//
// @param request - SetDeletionProtectionRequest
//
// @return SetDeletionProtectionResponse
func (client *Client) SetDeletionProtection(request *SetDeletionProtectionRequest) (_result *SetDeletionProtectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.SetDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a resource type or a version of a resource type.
//
// @param request - SetResourceTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetResourceTypeResponse
func (client *Client) SetResourceTypeWithOptions(request *SetResourceTypeRequest, runtime *dara.RuntimeOptions) (_result *SetResourceTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DefaultVersionId) {
		query["DefaultVersionId"] = request.DefaultVersionId
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.VersionId) {
		query["VersionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetResourceType"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a resource type or a version of a resource type.
//
// @param request - SetResourceTypeRequest
//
// @return SetResourceTypeResponse
func (client *Client) SetResourceType(request *SetResourceTypeRequest) (_result *SetResourceTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetResourceTypeResponse{}
	_body, _err := client.SetResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to configure a stack policy.
//
// Description:
//
// In this example, a stack policy is configured for a stack deployed in the `China (Hangzhou)` region whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The URL to the stack policy body is `oss://ros/stack-policy/demo`.
//
// @param request - SetStackPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetStackPolicyResponse
func (client *Client) SetStackPolicyWithOptions(request *SetStackPolicyRequest, runtime *dara.RuntimeOptions) (_result *SetStackPolicyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.StackPolicyBody) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !dara.IsNil(request.StackPolicyURL) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetStackPolicy"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetStackPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to configure a stack policy.
//
// Description:
//
// In this example, a stack policy is configured for a stack deployed in the `China (Hangzhou)` region whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The URL to the stack policy body is `oss://ros/stack-policy/demo`.
//
// @param request - SetStackPolicyRequest
//
// @return SetStackPolicyResponse
func (client *Client) SetStackPolicy(request *SetStackPolicyRequest) (_result *SetStackPolicyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetStackPolicyResponse{}
	_body, _err := client.SetStackPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Shares or unshares a template.
//
// Description:
//
// In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is shared with an Alibaba Cloud account. The ID of the Alibaba Cloud account is `151266687691****`.
//
// > The recipient Alibaba Cloud account (ID: `151266687691****`) can authorize RAM users to use the shared template.
//
// @param request - SetTemplatePermissionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTemplatePermissionResponse
func (client *Client) SetTemplatePermissionWithOptions(request *SetTemplatePermissionRequest, runtime *dara.RuntimeOptions) (_result *SetTemplatePermissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIds) {
		query["AccountIds"] = request.AccountIds
	}

	if !dara.IsNil(request.ShareOption) {
		query["ShareOption"] = request.ShareOption
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.VersionOption) {
		query["VersionOption"] = request.VersionOption
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTemplatePermission"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTemplatePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shares or unshares a template.
//
// Description:
//
// In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is shared with an Alibaba Cloud account. The ID of the Alibaba Cloud account is `151266687691****`.
//
// > The recipient Alibaba Cloud account (ID: `151266687691****`) can authorize RAM users to use the shared template.
//
// @param request - SetTemplatePermissionRequest
//
// @return SetTemplatePermissionResponse
func (client *Client) SetTemplatePermission(request *SetTemplatePermissionRequest) (_result *SetTemplatePermissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetTemplatePermissionResponse{}
	_body, _err := client.SetTemplatePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a signal to a resource in a stack.
//
// @param request - SignalResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SignalResourceResponse
func (client *Client) SignalResourceWithOptions(request *SignalResourceRequest, runtime *dara.RuntimeOptions) (_result *SignalResourceResponse, _err error) {
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

	if !dara.IsNil(request.LogicalResourceId) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SignalResource"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SignalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a signal to a resource in a stack.
//
// @param request - SignalResourceRequest
//
// @return SignalResourceResponse
func (client *Client) SignalResource(request *SignalResourceRequest) (_result *SignalResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SignalResourceResponse{}
	_body, _err := client.SignalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a stack group operation.
//
// Description:
//
// This topic provides an example on how to stop a stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863****` in the China (Hangzhou) region.
//
// @param request - StopStackGroupOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStackGroupOperationResponse
func (client *Client) StopStackGroupOperationWithOptions(request *StopStackGroupOperationRequest, runtime *dara.RuntimeOptions) (_result *StopStackGroupOperationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationId) {
		query["OperationId"] = request.OperationId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopStackGroupOperation"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopStackGroupOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a stack group operation.
//
// Description:
//
// This topic provides an example on how to stop a stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863****` in the China (Hangzhou) region.
//
// @param request - StopStackGroupOperationRequest
//
// @return StopStackGroupOperationResponse
func (client *Client) StopStackGroupOperation(request *StopStackGroupOperationRequest) (_result *StopStackGroupOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopStackGroupOperationResponse{}
	_body, _err := client.StopStackGroupOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates and adds tags to resources.
//
// Description:
//
// This topic provides an example on how to create a tag and add the tag to a stack. In this example, the stack ID is `7fee80e1-8c48-4c2f-8300-0f6dc40b****`, the tag key is `FinanceDept`, and the tag value is `FinanceJoshua`.
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates and adds tags to resources.
//
// Description:
//
// This topic provides an example on how to create a tag and add the tag to a stack. In this example, the stack ID is `7fee80e1-8c48-4c2f-8300-0f6dc40b****`, the tag key is `FinanceDept`, and the tag value is `FinanceJoshua`.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from resources and then deletes the tags.
//
// Description:
//
// This topic provides an example on how to remove all tags from a stack that is deployed in the China (Hangzhou) region. In this example, the stack ID is `46ec7b78-9d5e-4b21-aefd-448c90aa****`.
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from resources and then deletes the tags.
//
// Description:
//
// This topic provides an example on how to remove all tags from a stack that is deployed in the China (Hangzhou) region. In this example, the stack ID is `46ec7b78-9d5e-4b21-aefd-448c90aa****`.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a stack.
//
// Description:
//
// The values of parameters in the Parameters section vary based on the value that you specify for the UsePreviousParameters parameter in the request. If you do not add the parameters that are defined in the template to the Parameters section, take note of the following items:
//
//   - UsePreviousParameters is set to false: If the template parameters have default values, the default values are used. Otherwise, you must specify values for the template parameters in the Parameters section.
//
//   - UsePreviousParameters is set to true: If you specify values for the template parameters when you create a stack, the values are used. If you leave the template parameters empty when you create a stack but the template parameters have default values, the default values are used.
//
// This topic describes how to update a stack. In this example, the template body of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` in the China (Beijing) region is updated to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
//
// @param request - UpdateStackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStackResponse
func (client *Client) UpdateStackWithOptions(request *UpdateStackRequest, runtime *dara.RuntimeOptions) (_result *UpdateStackResponse, _err error) {
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

	if !dara.IsNil(request.DisableRollback) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.DryRunOptions) {
		query["DryRunOptions"] = request.DryRunOptions
	}

	if !dara.IsNil(request.Parallelism) {
		query["Parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.RamRoleName) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ReplacementOption) {
		query["ReplacementOption"] = request.ReplacementOption
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.StackPolicyBody) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !dara.IsNil(request.StackPolicyDuringUpdateBody) {
		query["StackPolicyDuringUpdateBody"] = request.StackPolicyDuringUpdateBody
	}

	if !dara.IsNil(request.StackPolicyDuringUpdateURL) {
		query["StackPolicyDuringUpdateURL"] = request.StackPolicyDuringUpdateURL
	}

	if !dara.IsNil(request.StackPolicyURL) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	if !dara.IsNil(request.TaintResources) {
		query["TaintResources"] = request.TaintResources
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !dara.IsNil(request.TimeoutInMinutes) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	if !dara.IsNil(request.UsePreviousParameters) {
		query["UsePreviousParameters"] = request.UsePreviousParameters
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStack"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a stack.
//
// Description:
//
// The values of parameters in the Parameters section vary based on the value that you specify for the UsePreviousParameters parameter in the request. If you do not add the parameters that are defined in the template to the Parameters section, take note of the following items:
//
//   - UsePreviousParameters is set to false: If the template parameters have default values, the default values are used. Otherwise, you must specify values for the template parameters in the Parameters section.
//
//   - UsePreviousParameters is set to true: If you specify values for the template parameters when you create a stack, the values are used. If you leave the template parameters empty when you create a stack but the template parameters have default values, the default values are used.
//
// This topic describes how to update a stack. In this example, the template body of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` in the China (Beijing) region is updated to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
//
// @param request - UpdateStackRequest
//
// @return UpdateStackResponse
func (client *Client) UpdateStack(request *UpdateStackRequest) (_result *UpdateStackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStackResponse{}
	_body, _err := client.UpdateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The region ID of the stack group. You can call the [DescribeRegions]\\(~~131035~~) operation to query the latest list of Alibaba Cloud regions.
//
// Description:
//
// The name of the stack group. The name must be unique within a region.
//
// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or a letter.
//
// @param tmpReq - UpdateStackGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStackGroupResponse
func (client *Client) UpdateStackGroupWithOptions(tmpReq *UpdateStackGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateStackGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStackGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("AccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AutoDeployment) {
		request.AutoDeploymentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoDeployment, dara.String("AutoDeployment"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeploymentTargets) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, dara.String("DeploymentTargets"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OperationPreferences) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, dara.String("OperationPreferences"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.AdministrationRoleName) {
		query["AdministrationRoleName"] = request.AdministrationRoleName
	}

	if !dara.IsNil(request.AutoDeploymentShrink) {
		query["AutoDeployment"] = request.AutoDeploymentShrink
	}

	if !dara.IsNil(request.Capabilities) {
		query["Capabilities"] = request.Capabilities
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeploymentOptions) {
		query["DeploymentOptions"] = request.DeploymentOptions
	}

	if !dara.IsNil(request.DeploymentTargetsShrink) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionRoleName) {
		query["ExecutionRoleName"] = request.ExecutionRoleName
	}

	if !dara.IsNil(request.OperationDescription) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !dara.IsNil(request.OperationPreferencesShrink) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !dara.IsNil(request.Parameters) {
		query["Parameters"] = request.Parameters
	}

	if !dara.IsNil(request.PermissionModel) {
		query["PermissionModel"] = request.PermissionModel
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.TemplateVersion) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStackGroup"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The region ID of the stack group. You can call the [DescribeRegions]\\(~~131035~~) operation to query the latest list of Alibaba Cloud regions.
//
// Description:
//
// The name of the stack group. The name must be unique within a region.
//
// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or a letter.
//
// @param request - UpdateStackGroupRequest
//
// @return UpdateStackGroupResponse
func (client *Client) UpdateStackGroup(request *UpdateStackGroupRequest) (_result *UpdateStackGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStackGroupResponse{}
	_body, _err := client.UpdateStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates stack instances in the specified accounts and regions.
//
// Description:
//
// In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. The stack group is granted the self-managed permissions. In this example, stacks of the stack group are updated by using the Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
//
// @param tmpReq - UpdateStackInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStackInstancesResponse
func (client *Client) UpdateStackInstancesWithOptions(tmpReq *UpdateStackInstancesRequest, runtime *dara.RuntimeOptions) (_result *UpdateStackInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AccountIds) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, dara.String("AccountIds"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DeploymentTargets) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, dara.String("DeploymentTargets"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.OperationPreferences) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, dara.String("OperationPreferences"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RegionIds) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, dara.String("RegionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AccountIdsShrink) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.DeploymentTargetsShrink) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !dara.IsNil(request.OperationDescription) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !dara.IsNil(request.OperationPreferencesShrink) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !dara.IsNil(request.ParameterOverrides) {
		query["ParameterOverrides"] = request.ParameterOverrides
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.RegionIdsShrink) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !dara.IsNil(request.StackGroupName) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !dara.IsNil(request.TimeoutInMinutes) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStackInstances"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates stack instances in the specified accounts and regions.
//
// Description:
//
// In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. The stack group is granted the self-managed permissions. In this example, stacks of the stack group are updated by using the Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
//
// @param request - UpdateStackInstancesRequest
//
// @return UpdateStackInstancesResponse
func (client *Client) UpdateStackInstances(request *UpdateStackInstancesRequest) (_result *UpdateStackInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStackInstancesResponse{}
	_body, _err := client.UpdateStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Corrects a template to eliminate stack drift.
//
// Description:
//
// Limits: You can eliminate only drift on stacks that have drifted. You must call the [DetectStackDrift](https://help.aliyun.com/document_detail/155094.html) operation to perform drift detection on a stack, call the [GetStackDriftDetectionStatus](https://help.aliyun.com/document_detail/155097.html) operation to query the drift status of the stack to make sure that the stack has drifted, and then call the UpdateStackTemplateByResources operation to eliminate drift.
//
// In this topic, drift is eliminated for a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The stack is deployed in the China (Hangzhou) region.
//
// @param request - UpdateStackTemplateByResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStackTemplateByResourcesResponse
func (client *Client) UpdateStackTemplateByResourcesWithOptions(request *UpdateStackTemplateByResourcesRequest, runtime *dara.RuntimeOptions) (_result *UpdateStackTemplateByResourcesResponse, _err error) {
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

	if !dara.IsNil(request.DryRun) {
		query["DryRun"] = request.DryRun
	}

	if !dara.IsNil(request.LogicalResourceId) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.StackId) {
		query["StackId"] = request.StackId
	}

	if !dara.IsNil(request.TemplateFormat) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateStackTemplateByResources"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateStackTemplateByResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Corrects a template to eliminate stack drift.
//
// Description:
//
// Limits: You can eliminate only drift on stacks that have drifted. You must call the [DetectStackDrift](https://help.aliyun.com/document_detail/155094.html) operation to perform drift detection on a stack, call the [GetStackDriftDetectionStatus](https://help.aliyun.com/document_detail/155097.html) operation to query the drift status of the stack to make sure that the stack has drifted, and then call the UpdateStackTemplateByResources operation to eliminate drift.
//
// In this topic, drift is eliminated for a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The stack is deployed in the China (Hangzhou) region.
//
// @param request - UpdateStackTemplateByResourcesRequest
//
// @return UpdateStackTemplateByResourcesResponse
func (client *Client) UpdateStackTemplateByResources(request *UpdateStackTemplateByResourcesRequest) (_result *UpdateStackTemplateByResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateStackTemplateByResourcesResponse{}
	_body, _err := client.UpdateStackTemplateByResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Template
//
// Description:
//
// When updating a template, please note:
//
// - If you specify `TemplateBody` or `TemplateURL`, the template version will be incremented by 1 after a successful update. For example, the version changes from v1 to v2.
//
// - If neither `TemplateBody` nor `TemplateURL` is specified, the template version remains unchanged.
//
// - A template can have up to 100 versions. If the version limit is reached, the template update will fail, and you need to recreate the template.
//
// @param request - UpdateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(request *UpdateTemplateRequest, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
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

	if !dara.IsNil(request.IsDraft) {
		query["IsDraft"] = request.IsDraft
	}

	if !dara.IsNil(request.RotateStrategy) {
		query["RotateStrategy"] = request.RotateStrategy
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.ValidationOptions) {
		query["ValidationOptions"] = request.ValidationOptions
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Template
//
// Description:
//
// When updating a template, please note:
//
// - If you specify `TemplateBody` or `TemplateURL`, the template version will be incremented by 1 after a successful update. For example, the version changes from v1 to v2.
//
// - If neither `TemplateBody` nor `TemplateURL` is specified, the template version remains unchanged.
//
// - A template can have up to 100 versions. If the version limit is reached, the template update will fail, and you need to recreate the template.
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a scenario.
//
// Description:
//
// ### [](#)Resource replication scenario
//
// Resource Orchestration Service (ROS) allows you to update a resource replication scenario. The updates that you make to a resource replication scenario do not affect the stack that is generated by using the resource scenario. You can call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
// ### [](#)Resource migration scenario
//
//   - If you want to update a resource migration scenario in which the migrated source resources are retained, you can delete the source resources to manage the updated resource migration scenario. You can also call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
//     **
//
//     **Note*	- Make sure that the source resources that you want to delete from a resource migration scenario are associated only with the resource scenario. Otherwise, the source resources fail to be deleted.
//
//   - If you want to update a resource migration scenario in which the migrated source resources are deleted, you can only call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
// ### [](#)Resource management scenario
//
// If you want to update a resource management scenario after you use the resource scenario to manage resources, you can only call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
// ### [](#)Resource detection scenario
//
// After you update a resource detection scenario, ROS obtains the most recent data from Resource Center and renders the architecture diagram.
//
// This topic provides an example on how to update a resource scenario. In this example, the ID of a virtual private cloud (VPC) in a resource scenario whose ID is `ts-7f7a704cf71c49a6****` is updated to `vpc-bp1m6fww66xbntjyc****`.
//
// @param tmpReq - UpdateTemplateScratchRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateScratchResponse
func (client *Client) UpdateTemplateScratchWithOptions(tmpReq *UpdateTemplateScratchRequest, runtime *dara.RuntimeOptions) (_result *UpdateTemplateScratchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTemplateScratchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.PreferenceParameters) {
		request.PreferenceParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreferenceParameters, dara.String("PreferenceParameters"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceResourceGroup) {
		request.SourceResourceGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResourceGroup, dara.String("SourceResourceGroup"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceResources) {
		request.SourceResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResources, dara.String("SourceResources"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SourceTag) {
		request.SourceTagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTag, dara.String("SourceTag"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ExecutionMode) {
		query["ExecutionMode"] = request.ExecutionMode
	}

	if !dara.IsNil(request.LogicalIdStrategy) {
		query["LogicalIdStrategy"] = request.LogicalIdStrategy
	}

	if !dara.IsNil(request.PreferenceParametersShrink) {
		query["PreferenceParameters"] = request.PreferenceParametersShrink
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SourceResourceGroupShrink) {
		query["SourceResourceGroup"] = request.SourceResourceGroupShrink
	}

	if !dara.IsNil(request.SourceResourcesShrink) {
		query["SourceResources"] = request.SourceResourcesShrink
	}

	if !dara.IsNil(request.SourceTagShrink) {
		query["SourceTag"] = request.SourceTagShrink
	}

	if !dara.IsNil(request.TemplateScratchId) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplateScratch"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a scenario.
//
// Description:
//
// ### [](#)Resource replication scenario
//
// Resource Orchestration Service (ROS) allows you to update a resource replication scenario. The updates that you make to a resource replication scenario do not affect the stack that is generated by using the resource scenario. You can call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
// ### [](#)Resource migration scenario
//
//   - If you want to update a resource migration scenario in which the migrated source resources are retained, you can delete the source resources to manage the updated resource migration scenario. You can also call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
//     **
//
//     **Note*	- Make sure that the source resources that you want to delete from a resource migration scenario are associated only with the resource scenario. Otherwise, the source resources fail to be deleted.
//
//   - If you want to update a resource migration scenario in which the migrated source resources are deleted, you can only call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
// ### [](#)Resource management scenario
//
// If you want to update a resource management scenario after you use the resource scenario to manage resources, you can only call the [GenerateTemplateByScratch](https://help.aliyun.com/document_detail/610829.html) operation to generate a template for the resource scenario.
//
// ### [](#)Resource detection scenario
//
// After you update a resource detection scenario, ROS obtains the most recent data from Resource Center and renders the architecture diagram.
//
// This topic provides an example on how to update a resource scenario. In this example, the ID of a virtual private cloud (VPC) in a resource scenario whose ID is `ts-7f7a704cf71c49a6****` is updated to `vpc-bp1m6fww66xbntjyc****`.
//
// @param request - UpdateTemplateScratchRequest
//
// @return UpdateTemplateScratchResponse
func (client *Client) UpdateTemplateScratch(request *UpdateTemplateScratchRequest) (_result *UpdateTemplateScratchResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTemplateScratchResponse{}
	_body, _err := client.UpdateTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Validates a template by using a template URL or template body. The template is used to create a stack.
//
// Description:
//
// In this example, a template that you want to use to create a stack is validated. `TemplateURL` is set to `oss://ros/template/demo`.
//
// @param request - ValidateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTemplateResponse
func (client *Client) ValidateTemplateWithOptions(request *ValidateTemplateRequest, runtime *dara.RuntimeOptions) (_result *ValidateTemplateResponse, _err error) {
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

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TemplateURL) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !dara.IsNil(request.UpdateInfoOptions) {
		query["UpdateInfoOptions"] = request.UpdateInfoOptions
	}

	if !dara.IsNil(request.ValidationOption) {
		query["ValidationOption"] = request.ValidationOption
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.TemplateBody) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateTemplate"),
		Version:     dara.String("2019-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates a template by using a template URL or template body. The template is used to create a stack.
//
// Description:
//
// In this example, a template that you want to use to create a stack is validated. `TemplateURL` is set to `oss://ros/template/demo`.
//
// @param request - ValidateTemplateRequest
//
// @return ValidateTemplateResponse
func (client *Client) ValidateTemplate(request *ValidateTemplateRequest) (_result *ValidateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ValidateTemplateResponse{}
	_body, _err := client.ValidateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
