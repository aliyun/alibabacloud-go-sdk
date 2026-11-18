// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Submits asynchronous retests in bulk for specified vulnerabilities under the current account and returns the acceptance results for each item.
//
// Description:
//
// Submits retests for existing vulnerabilities that the current account has access to. OperationInput accepts only VulnIds. The original array must contain 1 to 100 elements, and duplicate elements are counted. The original length is validated before deduplication. If the limit is exceeded, the entire request is rejected and no verification records are created. Valid vulnerability IDs are processed item by item. A single item failure does not block the remaining items or roll back accepted items. SuccessCount and Results.Success only indicate that the verification records are created and accepted, and do not mean that the vulnerabilities have been retested or passed verification. Retests are executed asynchronously and may consume credits.
//
// @param tmpReq - BatchVerifyPentestVulnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchVerifyPentestVulnResponse
func (client *Client) BatchVerifyPentestVulnWithContext(ctx context.Context, tmpReq *BatchVerifyPentestVulnRequest, runtime *dara.RuntimeOptions) (_result *BatchVerifyPentestVulnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchVerifyPentestVulnShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchVerifyPentestVuln"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchVerifyPentestVulnResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates penetration testing tasks for authorized assets of the current account and returns the task IDs. You can specify up to five targets at a time, and one task is created for each target.
//
// Description:
//
// Creates penetration testing tasks for authorized assets of the current account and returns the task IDs. You can specify up to five targets at a time, and one task is created for each target.
//
// Before you call this operation, you must activate the AgenticBAS service or meet the validity period and quota conditions of the trial service. The target assets must meet the ownership verification requirements of the selected asset source. Business parameters are passed in OperationInput. VulnCheckItems and PentestRequirement cannot be empty at the same time. Currently, only IMMEDIATE execution is supported, and scheduled execution is not supported.
//
// A successful creation response only indicates that the tasks are created and does not wait for the scan to complete. Save the TaskId in TaskList, query the status and usage by calling DescribePentestTaskList, and then use the same TaskId to query vulnerabilities or reports. Tasks may enter a queue. If a report is not generated, the query response may not contain a download link.
//
// Credentials use an object structure of "identity name -> credential configuration". The following are synthetic examples. Replace them with authorized test identities in actual calls.
//
// ```json
//
// {
//
//	"Credentials": {
//
//	  "accountUser": {
//
//	    "account": {
//
//	      "username": "example-user",
//
//	      "password": "example-password"
//
//	    }
//
//	  },
//
//	  "cookieUser": {
//
//	    "cookies": [
//
//	      {
//
//	        "name": "session",
//
//	        "value": "example-session",
//
//	        "domain": "example.com",
//
//	        "path": "/",
//
//	        "secure": true,
//
//	        "httpOnly": true,
//
//	        "sameSite": "Lax"
//
//	      }
//
//	    ]
//
//	  },
//
//	  "localUser": {
//
//	    "localStorage": {
//
//	      "access_token": "example-local-token"
//
//	    }
//
//	  },
//
//	  "sessionUser": {
//
//	    "sessionStorage": {
//
//	      "session_token": "example-session-token"
//
//	    }
//
//	  },
//
//	  "headerUser": {
//
//	    "extraHeaders": {
//
//	      "Authorization": "Bearer example-token"
//
//	    }
//
//	  }
//
//	}
//
// }
//
// ```
//
// @param tmpReq - CreatePentestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePentestTaskResponse
func (client *Client) CreatePentestTaskWithContext(ctx context.Context, tmpReq *CreatePentestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatePentestTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreatePentestTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePentestTask"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePentestTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PDF report of a specified penetration test task for the current account, and returns a temporary Object Storage Service (OSS) download URL and the expiration time without returning the report content.
//
// @param tmpReq - DescribePentestReportContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestReportContentResponse
func (client *Client) DescribePentestReportContentWithContext(ctx context.Context, tmpReq *DescribePentestReportContentRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestReportContentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestReportContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestReportContent"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestReportContentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the penetration testing nodes of the current account by paging and returns the node status, execution duration, Credits usage, time, and vulnerability statistics.
//
// Description:
//
// Queries the penetration testing nodes of the current account by paging and returns the node status, execution duration, Credits usage, time, and vulnerability statistics.
//
// Use OperationInput to specify query conditions. Paging is implemented by using CurrentPage and PageSize. By default, the first page is returned with 10 entries per page. You can filter nodes by node name keyword and node status.
//
// The response includes TaskList and PageInfo. Duration indicates the execution duration in seconds. CreditsUsed indicates the number of Credits used. Nodes that have not started may not have a start time or duration, and nodes that have not finished may not have a completion time. All UNIX timestamps are in UNIX seconds. You can transform them to the time zone of the invoker for display.
//
// @param tmpReq - DescribePentestTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestTaskListResponse
func (client *Client) DescribePentestTaskListWithContext(ctx context.Context, tmpReq *DescribePentestTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestTaskListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestTaskList"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestTaskListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a single vulnerability under the current account without triggering a scan or vulnerability verification.
//
// Description:
//
// Queries a specified single vulnerability under the current account by using OperationInput.VulnId and returns a VulnDetail object. This operation only reads existing vulnerability data, does not create scan or verification tasks, and does not support batch queries by using TaskId or VulnIds. You can call DescribePentestVulnList to obtain vulnerability IDs and then query the details of each vulnerability. The caller identity is determined by the gateway, and you cannot specify other accounts through request parameters.
//
// @param tmpReq - DescribePentestVulnDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestVulnDetailResponse
func (client *Client) DescribePentestVulnDetailWithContext(ctx context.Context, tmpReq *DescribePentestVulnDetailRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestVulnDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestVulnDetailShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestVulnDetail"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestVulnDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the penetration testing vulnerabilities of the current account by task ID with pagination. Returns information such as the vulnerability title, severity, authentication status, affected URLs, and discovery time, and supports filtering by severity.
//
// @param tmpReq - DescribePentestVulnListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestVulnListResponse
func (client *Client) DescribePentestVulnListWithContext(ctx context.Context, tmpReq *DescribePentestVulnListRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestVulnListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DescribePentestVulnListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.OperationInput) {
		request.OperationInputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationInput, dara.String("OperationInput"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.OperationInputShrink) {
		query["OperationInput"] = request.OperationInputShrink
	}

	if !dara.IsNil(request.OperationType) {
		query["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribePentestVulnList"),
		Version:     dara.String("2026-06-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribePentestVulnListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
