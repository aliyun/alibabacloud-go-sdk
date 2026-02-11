// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// # Apply for Bastion Host Account
//
// Description:
//
// ## Request Description
//
// - This interface is used for customers to create a modeling project for the first time.
//
// - `projectName` is a required field, with a maximum length of 50 characters.
//
// - `remark` and `instanceSpec` are optional, where `remark` has a maximum length of 200 characters.
//
// - The available values for `instanceSpec` include `SECURE_ENV_LITE` and `SECURE_ENV_PRO`.
//
// @param request - ApplyBastionAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyBastionAccountResponse
func (client *Client) ApplyBastionAccountWithContext(ctx context.Context, request *ApplyBastionAccountRequest, runtime *dara.RuntimeOptions) (_result *ApplyBastionAccountResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApplyBastionAccount"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ApplyBastionAccountResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Associate Retrospective Task
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - AssociatePocTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssociatePocTaskResponse
func (client *Client) AssociatePocTaskWithContext(ctx context.Context, request *AssociatePocTaskRequest, runtime *dara.RuntimeOptions) (_result *AssociatePocTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AssociatePocTask"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AssociatePocTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Complete project and release resources.
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - CompleteModelingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteModelingProjectResponse
func (client *Client) CompleteModelingProjectWithContext(ctx context.Context, request *CompleteModelingProjectRequest, runtime *dara.RuntimeOptions) (_result *CompleteModelingProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CompleteModelingProject"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CompleteModelingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create basic model information
//
// @param request - CreateCustomerModuleBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerModuleBasicInfoResponse
func (client *Client) CreateCustomerModuleBasicInfoWithContext(ctx context.Context, request *CreateCustomerModuleBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerModuleBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleName) {
		query["CustomerModuleName"] = request.CustomerModuleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ModuleType) {
		query["ModuleType"] = request.ModuleType
	}

	if !dara.IsNil(request.StorePath) {
		query["StorePath"] = request.StorePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomerModuleBasicInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerModuleBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Save Model Feature Configuration
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each option including a label (label) and value (value).
//
// @param request - CreateCustomerModuleMetaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerModuleMetaInfoResponse
func (client *Client) CreateCustomerModuleMetaInfoWithContext(ctx context.Context, request *CreateCustomerModuleMetaInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerModuleMetaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FeatureString) {
		query["FeatureString"] = request.FeatureString
	}

	if !dara.IsNil(request.FeatureTemplate) {
		query["FeatureTemplate"] = request.FeatureTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomerModuleMetaInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerModuleMetaInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Model Output Information
//
// @param request - CreateCustomerModuleOutputInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomerModuleOutputInfoResponse
func (client *Client) CreateCustomerModuleOutputInfoWithContext(ctx context.Context, request *CreateCustomerModuleOutputInfoRequest, runtime *dara.RuntimeOptions) (_result *CreateCustomerModuleOutputInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FinalScoreFormat) {
		query["FinalScoreFormat"] = request.FinalScoreFormat
	}

	if !dara.IsNil(request.ProcessExpression) {
		query["ProcessExpression"] = request.ProcessExpression
	}

	if !dara.IsNil(request.TestFileUrl) {
		query["TestFileUrl"] = request.TestFileUrl
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCustomerModuleOutputInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomerModuleOutputInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initialize a modeling project.
//
// Description:
//
// ## Request Description
//
// - This interface is used for customers to create a modeling project for the first time.
//
// - `projectName` is a required field, with a maximum length of 50 characters.
//
// - `remark` and `instanceSpec` are optional, where `remark` has a maximum length of 200 characters.
//
// - The available values for `instanceSpec` include `SECURE_ENV_LITE` and `SECURE_ENV_PRO`.
//
// @param request - CreateModelingProjectRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelingProjectResponse
func (client *Client) CreateModelingProjectWithContext(ctx context.Context, request *CreateModelingProjectRequest, runtime *dara.RuntimeOptions) (_result *CreateModelingProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceSpec) {
		query["InstanceSpec"] = request.InstanceSpec
	}

	if !dara.IsNil(request.ProjectName) {
		query["ProjectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateModelingProject"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelingProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified customer model based on the provided customer model ID.
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: Deletion is irreversible, please use with caution.
//
// @param request - DeleteModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelResponse
func (client *Client) DeleteModelWithContext(ctx context.Context, request *DeleteModelRequest, runtime *dara.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Deploy Model File
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - DeployModelFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployModelFileResponse
func (client *Client) DeployModelFileWithContext(ctx context.Context, request *DeployModelFileRequest, runtime *dara.RuntimeOptions) (_result *DeployModelFileResponse, _err error) {
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

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeployModelFile"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeployModelFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query basic model information based on the customer model ID
//
// Description:
//
// ## Request Description
//
// By providing the `customerModuleId` parameter, you can query the current status of a specified customer model. The status values may include but are not limited to "EDIT", "ONLINE", etc.
//
// @param request - DescribeCustomerModuleBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomerModuleBasicInfoResponse
func (client *Client) DescribeCustomerModuleBasicInfoWithContext(ctx context.Context, request *DescribeCustomerModuleBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomerModuleBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomerModuleBasicInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomerModuleBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query model feature information based on the customer model ID
//
// Description:
//
// ## Request Description
//
// By providing the `customerModuleId` parameter, you can query the current status of a specified customer model. The status values may include, but are not limited to, "EDIT", "ONLINE", etc.
//
// @param request - DescribeCustomerModuleMetaInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomerModuleMetaInfoResponse
func (client *Client) DescribeCustomerModuleMetaInfoWithContext(ctx context.Context, request *DescribeCustomerModuleMetaInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomerModuleMetaInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomerModuleMetaInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomerModuleMetaInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query model output information based on the customer model ID
//
// Description:
//
// ## Request Description
//
// By providing the `customerModuleId` parameter, you can query the current status of a specified customer model. The status values may include but are not limited to "EDIT", "ONLINE", etc.
//
// @param request - DescribeCustomerModuleOutputInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCustomerModuleOutputInfoResponse
func (client *Client) DescribeCustomerModuleOutputInfoWithContext(ctx context.Context, request *DescribeCustomerModuleOutputInfoRequest, runtime *dara.RuntimeOptions) (_result *DescribeCustomerModuleOutputInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCustomerModuleOutputInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCustomerModuleOutputInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Return specific feature options based on the feature template
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each of which includes a label (label) and a value (value).
//
// @param request - DescribeFeatureOptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFeatureOptionResponse
func (client *Client) DescribeFeatureOptionWithContext(ctx context.Context, request *DescribeFeatureOptionRequest, runtime *dara.RuntimeOptions) (_result *DescribeFeatureOptionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FeatureTemplate) {
		query["FeatureTemplate"] = request.FeatureTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFeatureOption"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFeatureOptionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get File Download Link
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - DescribeFileDownloadUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFileDownloadUrlResponse
func (client *Client) DescribeFileDownloadUrlWithContext(ctx context.Context, request *DescribeFileDownloadUrlRequest, runtime *dara.RuntimeOptions) (_result *DescribeFileDownloadUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeFileDownloadUrl"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeFileDownloadUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Render the feature configuration of the model
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each including a label (label) and value (value).
//
// @param request - DescribeModelFeatureRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelFeatureResponse
func (client *Client) DescribeModelFeatureWithContext(ctx context.Context, request *DescribeModelFeatureRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelFeatureResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FeatureTemplate) {
		query["FeatureTemplate"] = request.FeatureTemplate
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelFeature"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelFeatureResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get detailed project data
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - DescribeModelingProjectDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelingProjectDetailResponse
func (client *Client) DescribeModelingProjectDetailWithContext(ctx context.Context, request *DescribeModelingProjectDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelingProjectDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelingProjectDetail"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelingProjectDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Paginated query for the list of modeling projects under the current user.
//
// Description:
//
// ## 请求说明
//
// - 该API用于获取指定租户下的所有建模项目的概览信息。
//
// - 支持通过`pageSize`和`currentPage`参数进行分页查询，默认每页显示10条记录。
//
// - 可选地，使用`status`参数来过滤特定状态（如`active`, `released`等）的项目。
//
// - 返回结果中包含每个项目的ID、名称、环境状态、建模状态、开始时间及结束时间（如果有的话），以及创建该项目的登录账号。
//
// @param request - DescribeModelingProjectListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModelingProjectListResponse
func (client *Client) DescribeModelingProjectListWithContext(ctx context.Context, request *DescribeModelingProjectListRequest, runtime *dara.RuntimeOptions) (_result *DescribeModelingProjectListResponse, _err error) {
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

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModelingProjectList"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModelingProjectListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Whether the Service Corresponding to a Specific Model Exists Based on Customer Model ID
//
// Description:
//
// ## Request Description
//
// This interface is used to check whether a specific model service exists by providing the `customerModuleId`. If it exists, it returns `true`; otherwise, it returns `false`.
//
// ### Notes
//
// - `customerModuleId` is a required parameter and must be of string type.
//
// - This API is mainly used for front-end page display or logical judgment to confirm whether the service corresponding to the user\\"s selected model has been created.
//
// @param request - DescribeModuleServiceExistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModuleServiceExistResponse
func (client *Client) DescribeModuleServiceExistWithContext(ctx context.Context, request *DescribeModuleServiceExistRequest, runtime *dara.RuntimeOptions) (_result *DescribeModuleServiceExistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModuleServiceExist"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModuleServiceExistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the current status of a model based on the customer model ID
//
// Description:
//
// ## Request Description
//
// By providing the `customerModuleId` parameter, you can query the current status of a specified customer model. The status values may include, but are not limited to, "EDIT", "ONLINE", etc.
//
// @param request - DescribeModuleStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeModuleStatusResponse
func (client *Client) DescribeModuleStatusWithContext(ctx context.Context, request *DescribeModuleStatusRequest, runtime *dara.RuntimeOptions) (_result *DescribeModuleStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeModuleStatus"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeModuleStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Scene and Service
//
// @param request - DescribeServiceAndSceneRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeServiceAndSceneResponse
func (client *Client) DescribeServiceAndSceneWithContext(ctx context.Context, request *DescribeServiceAndSceneRequest, runtime *dara.RuntimeOptions) (_result *DescribeServiceAndSceneResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeServiceAndScene"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeServiceAndSceneResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get all model information for a specified user, supporting pagination
//
// Description:
//
// ## Request Description
//
// This API is used to query all model information under a specific user and supports pagination through page parameters. Fuzzy search can be performed using the `name` parameter.
//
// - `regId`: Region identifier, required.
//
// - `pageSize`: Number of items per page, required.
//
// - `currentPage`: Current page number, starting from 1, required.
//
// - `userId`: User ID, required.
//
// @param request - DescribeUserModelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserModelListResponse
func (client *Client) DescribeUserModelListWithContext(ctx context.Context, request *DescribeUserModelListRequest, runtime *dara.RuntimeOptions) (_result *DescribeUserModelListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuthType) {
		query["AuthType"] = request.AuthType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeUserModelList"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeUserModelListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Iterate Model
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: The deletion operation is irreversible, please use with caution.
//
// @param request - DuplicateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DuplicateModelResponse
func (client *Client) DuplicateModelWithContext(ctx context.Context, request *DuplicateModelRequest, runtime *dara.RuntimeOptions) (_result *DuplicateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DuplicateModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DuplicateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Edit Model
//
// @param request - EditModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditModelResponse
func (client *Client) EditModelWithContext(ctx context.Context, request *EditModelRequest, runtime *dara.RuntimeOptions) (_result *EditModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # View Bastion Host Initial Password
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - GetBastionHostCertificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBastionHostCertificationResponse
func (client *Client) GetBastionHostCertificationWithContext(ctx context.Context, request *GetBastionHostCertificationRequest, runtime *dara.RuntimeOptions) (_result *GetBastionHostCertificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBastionHostCertification"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBastionHostCertificationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Iterate Model
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: The deletion operation is irreversible, please use with caution.
//
// @param request - IterateModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IterateModelResponse
func (client *Client) IterateModelWithContext(ctx context.Context, request *IterateModelRequest, runtime *dara.RuntimeOptions) (_result *IterateModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IterateModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IterateModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Offline Model
//
// @param request - OfflineModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineModelResponse
func (client *Client) OfflineModelWithContext(ctx context.Context, request *OfflineModelRequest, runtime *dara.RuntimeOptions) (_result *OfflineModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Online Model
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: Deletion is irreversible, please use with caution.
//
// @param request - OnlineModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OnlineModelResponse
func (client *Client) OnlineModelWithContext(ctx context.Context, request *OnlineModelRequest, runtime *dara.RuntimeOptions) (_result *OnlineModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OnlineModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &OnlineModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Parse Expression Parameters
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each of which includes a label (label) and a value (value).
//
// @param request - ParseExpressionParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ParseExpressionParametersResponse
func (client *Client) ParseExpressionParametersWithContext(ctx context.Context, request *ParseExpressionParametersRequest, runtime *dara.RuntimeOptions) (_result *ParseExpressionParametersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ParseExpressionParameters"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ParseExpressionParametersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pre-release Model
//
// @param request - PrepublishModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PrepublishModelResponse
func (client *Client) PrepublishModelWithContext(ctx context.Context, request *PrepublishModelRequest, runtime *dara.RuntimeOptions) (_result *PrepublishModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PrepublishModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PrepublishModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Request to Sync Model Files
//
// Description:
//
// ## 请求说明
//
// - 本接口用于客户首次创建建模项目。
//
// - `projectName` 是必填项，长度不超过50个字符。
//
// - `remark` 和 `instanceSpec` 为可选项，其中 `remark` 长度不超过200个字符。
//
// - `instanceSpec` 可选值包括 `SECURE_ENV_LITE` 和 `SECURE_ENV_PRO`。
//
// @param request - RequestModelFileSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequestModelFileSyncResponse
func (client *Client) RequestModelFileSyncWithContext(ctx context.Context, request *RequestModelFileSyncRequest, runtime *dara.RuntimeOptions) (_result *RequestModelFileSyncResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["ProjectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RequestModelFileSync"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RequestModelFileSyncResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Roll back the specified customer model based on the provided customer model ID.
//
// Description:
//
// ## Request Description
//
// This API is used to delete a specified customer model from the system. When calling, you must provide the `customerModuleId` parameter, which identifies the specific model to be deleted.
//
// - **Note**: The deletion operation is irreversible, please use with caution.
//
// @param request - RollbackModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RollbackModelResponse
func (client *Client) RollbackModelWithContext(ctx context.Context, request *RollbackModelRequest, runtime *dara.RuntimeOptions) (_result *RollbackModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RollbackModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RollbackModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pre-release Model Testing
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each of which includes a label (label) and a value (value).
//
// @param request - TestModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestModelResponse
func (client *Client) TestModelWithContext(ctx context.Context, request *TestModelRequest, runtime *dara.RuntimeOptions) (_result *TestModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Pre-release Test Model
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each of which includes a label (label) and a value (value).
//
// @param request - TestPreModelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestPreModelResponse
func (client *Client) TestPreModelWithContext(ctx context.Context, request *TestPreModelRequest, runtime *dara.RuntimeOptions) (_result *TestPreModelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestPreModel"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestPreModelResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Test Expression
//
// Description:
//
// ## Request Description
//
// - This interface is used to query all available feature templates in the system.
//
// - The request method is GET, and no additional parameters are required.
//
// - The returned result includes multiple feature template options, each of which includes a label (label) and a value (value).
//
// @param request - TestProcessExpressionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TestProcessExpressionResponse
func (client *Client) TestProcessExpressionWithContext(ctx context.Context, request *TestProcessExpressionRequest, runtime *dara.RuntimeOptions) (_result *TestProcessExpressionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Expression) {
		query["Expression"] = request.Expression
	}

	if !dara.IsNil(request.Params) {
		query["Params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TestProcessExpression"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TestProcessExpressionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update basic model information
//
// @param request - UpdateModuleBasicInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModuleBasicInfoResponse
func (client *Client) UpdateModuleBasicInfoWithContext(ctx context.Context, request *UpdateModuleBasicInfoRequest, runtime *dara.RuntimeOptions) (_result *UpdateModuleBasicInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.CustomerModuleName) {
		query["CustomerModuleName"] = request.CustomerModuleName
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.ModuleType) {
		query["ModuleType"] = request.ModuleType
	}

	if !dara.IsNil(request.StorePath) {
		query["StorePath"] = request.StorePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateModuleBasicInfo"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModuleBasicInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Validate model file
//
// @param request - ValidateModelFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateModelFileResponse
func (client *Client) ValidateModelFileWithContext(ctx context.Context, request *ValidateModelFileRequest, runtime *dara.RuntimeOptions) (_result *ValidateModelFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateModelFile"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateModelFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Validate Test File
//
// @param request - ValidateTestFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTestFileResponse
func (client *Client) ValidateTestFileWithContext(ctx context.Context, request *ValidateTestFileRequest, runtime *dara.RuntimeOptions) (_result *ValidateTestFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomerModuleId) {
		query["CustomerModuleId"] = request.CustomerModuleId
	}

	if !dara.IsNil(request.FilePath) {
		query["FilePath"] = request.FilePath
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateTestFile"),
		Version:     dara.String("2025-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateTestFileResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
