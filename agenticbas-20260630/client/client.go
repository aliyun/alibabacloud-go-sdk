// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pop/client"
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
	client.ProductId = dara.String("AgenticBAS")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("agenticbas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
func (client *Client) CreatePentestTaskWithOptions(tmpReq *CreatePentestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreatePentestTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - CreatePentestTaskRequest
//
// @return CreatePentestTaskResponse
func (client *Client) CreatePentestTask(request *CreatePentestTaskRequest) (_result *CreatePentestTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreatePentestTaskResponse{}
	_body, _err := client.CreatePentestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the PDF report of a specified penetration test task under the current account and returns a temporary OSS download URL and its expiration time. The report content is not returned.
//
// @param tmpReq - DescribePentestReportContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestReportContentResponse
func (client *Client) DescribePentestReportContentWithOptions(tmpReq *DescribePentestReportContentRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestReportContentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the PDF report of a specified penetration test task under the current account and returns a temporary OSS download URL and its expiration time. The report content is not returned.
//
// @param request - DescribePentestReportContentRequest
//
// @return DescribePentestReportContentResponse
func (client *Client) DescribePentestReportContent(request *DescribePentestReportContentRequest) (_result *DescribePentestReportContentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePentestReportContentResponse{}
	_body, _err := client.DescribePentestReportContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the penetration testing tasks of the current account by using paging and returns the task status, execution duration, Credits usage, time, and vulnerability statistics.
//
// Description:
//
// Queries the penetration testing tasks of the current account by using paging and returns the task status, execution duration, Credits usage, time, and vulnerability statistics.
//
// Use OperationInput to specify query conditions. Paging is implemented by using CurrentPage and PageSize. By default, the first page is queried and each page contains 10 entries. You can filter tasks by task name keyword and task status.
//
// The response includes TaskList and PageInfo. Duration indicates the execution duration in seconds. CreditsUsed indicates the used Credits. Tasks that have not started may not have a start time or execution duration, and unfinished tasks may not have a completion time. UNIX timestamps uniformly use UNIX seconds. When displaying, you can convert them to the time zone of the caller.
//
// @param tmpReq - DescribePentestTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestTaskListResponse
func (client *Client) DescribePentestTaskListWithOptions(tmpReq *DescribePentestTaskListRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestTaskListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the penetration testing tasks of the current account by using paging and returns the task status, execution duration, Credits usage, time, and vulnerability statistics.
//
// Description:
//
// Queries the penetration testing tasks of the current account by using paging and returns the task status, execution duration, Credits usage, time, and vulnerability statistics.
//
// Use OperationInput to specify query conditions. Paging is implemented by using CurrentPage and PageSize. By default, the first page is queried and each page contains 10 entries. You can filter tasks by task name keyword and task status.
//
// The response includes TaskList and PageInfo. Duration indicates the execution duration in seconds. CreditsUsed indicates the used Credits. Tasks that have not started may not have a start time or execution duration, and unfinished tasks may not have a completion time. UNIX timestamps uniformly use UNIX seconds. When displaying, you can convert them to the time zone of the caller.
//
// @param request - DescribePentestTaskListRequest
//
// @return DescribePentestTaskListResponse
func (client *Client) DescribePentestTaskList(request *DescribePentestTaskListRequest) (_result *DescribePentestTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePentestTaskListResponse{}
	_body, _err := client.DescribePentestTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries penetration test vulnerabilities for the current account by task ID with paging. Returns vulnerability titles, severity levels, authentication statuses, affected URLs, and discovery times. Supports filtering by severity level.
//
// @param tmpReq - DescribePentestVulnListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePentestVulnListResponse
func (client *Client) DescribePentestVulnListWithOptions(tmpReq *DescribePentestVulnListRequest, runtime *dara.RuntimeOptions) (_result *DescribePentestVulnListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries penetration test vulnerabilities for the current account by task ID with paging. Returns vulnerability titles, severity levels, authentication statuses, affected URLs, and discovery times. Supports filtering by severity level.
//
// @param request - DescribePentestVulnListRequest
//
// @return DescribePentestVulnListResponse
func (client *Client) DescribePentestVulnList(request *DescribePentestVulnListRequest) (_result *DescribePentestVulnListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribePentestVulnListResponse{}
	_body, _err := client.DescribePentestVulnListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
