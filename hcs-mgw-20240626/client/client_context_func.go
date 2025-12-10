// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Creates a data address.
//
// Description:
//
//	  To create a data address, you must have the permission on mgw:CreateImportAddress.
//
//		- If you want to use an agent to migrate data, you must create an agent first and then associate the agent with a data address when you create the data address.
//
// @param request - CreateAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAddressResponse
func (client *Client) CreateAddressWithContext(ctx context.Context, userid *string, request *CreateAddressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportAddress) {
		body["ImportAddress"] = request.ImportAddress
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAddress"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/address"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &CreateAddressResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The request boy for creating the agent.
//
// Description:
//
//	  To create an agent, you must have the permission on mgw:CreateImportAgent.
//
//		- If you want to migrate data to Alibaba Cloud over an Express Connect circuit or a VPN gateway, or migrate data from a self-managed storage space to Alibaba Cloud, you can deploy an agent.
//
//		- Before you create an agent, you must create a tunnel. An agent must be associated with a tunnel.
//
// @param request - CreateAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithContext(ctx context.Context, userid *string, request *CreateAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportAgent) {
		body["ImportAgent"] = request.ImportAgent
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgent"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a migration task.
//
// Description:
//
//	  To create a migration task, you must have the permission on mgw:CreateImportJob.
//
//		- Before you create a migration task, you must create data addresses.
//
//		- A migration task can run multiple rounds. Each round has an execution ID.
//
// @param request - CreateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateJobResponse
func (client *Client) CreateJobWithContext(ctx context.Context, userid *string, request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportJob) {
		body["ImportJob"] = request.ImportJob
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateJob"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/job"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a migration report.
//
// Description:
//
//	  To create a migration report, you must have the permission on mgw:CreateImportReport.
//
//		- If you specify that a migration report is to be generated when you create a migration task, you do not need to call this operation. If you do not specify that a migration report is to be generated when you create a migration task, you can call this operation to create a migration report for an execution with the specified ID.
//
// @param request - CreateReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateReportResponse
func (client *Client) CreateReportWithContext(ctx context.Context, userid *string, request *CreateReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateReport) {
		body["CreateReport"] = request.CreateReport
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateReport"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/report"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &CreateReportResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tunnel.
//
// Description:
//
//	  To create a tunnel, you must have the permission on mgw:CreateImportTunnel.
//
//		- When you use an agent to migrate data, the agent must be associated with a tunnel.
//
//		- A tunnel can be associated with multiple agents. You can throttle the traffic of the agents that are associated with the same tunnel by setting the bandwidth and the number of requests per second for the tunnel.
//
// @param request - CreateTunnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTunnelResponse
func (client *Client) CreateTunnelWithContext(ctx context.Context, userid *string, request *CreateTunnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTunnelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportTunnel) {
		body["ImportTunnel"] = request.ImportTunnel
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTunnel"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tunnel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &CreateTunnelResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a data address.
//
// Description:
//
//	To delete a data address, you must have the permission on mgw:DeleteImportAddress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAddressResponse
func (client *Client) DeleteAddressWithContext(ctx context.Context, userid *string, addressName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAddressResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAddress"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/address/" + dara.StringValue(addressName)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &DeleteAddressResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// Description:
//
//	To delete an agent, you must have the permission on mgw:DeleteImportAgent.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithContext(ctx context.Context, userid *string, agentName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agent/" + dara.StringValue(agentName)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a migration task.
//
// Description:
//
//	  To delete a migration task, you must have the permission on mgw:DeleteImportJob.
//
//		- The operation to delete a migration task is asynchronous. The migration task remains in the Deleting state until it is deleted.
//
// @param request - DeleteJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteJobResponse
func (client *Client) DeleteJobWithContext(ctx context.Context, userid *string, jobName *string, request *DeleteJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.ForceDelete) {
		query["forceDelete"] = request.ForceDelete
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteJob"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/job/" + dara.StringValue(jobName)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a tunnel.
//
// Description:
//
//	To delete a tunnel, you must have the permission on mgw:DeleteImportTunnel.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTunnelResponse
func (client *Client) DeleteTunnelWithContext(ctx context.Context, userid *string, tunnelId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTunnelResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTunnel"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tunnel/" + dara.StringValue(tunnelId)),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &DeleteTunnelResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a data address.
//
// Description:
//
//	To query the information about a data address, you must have the permission on mgw:GetImportAddress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddressResponse
func (client *Client) GetAddressWithContext(ctx context.Context, userid *string, addressName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddressResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAddress"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/address/" + dara.StringValue(addressName)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetAddressResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of an agent.
//
// Description:
//
//	To query the information about an agent, you must have the permission on mgw:GetImportAgent.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentResponse
func (client *Client) GetAgentWithContext(ctx context.Context, userid *string, agentName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgent"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agent/" + dara.StringValue(agentName)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetAgentResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the running status of an agent.
//
// Description:
//
//	To query the status of an agent, you must have the permission on mgw:GetImportAgent.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentStatusResponse
func (client *Client) GetAgentStatusWithContext(ctx context.Context, userid *string, agentName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentStatusResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentStatus"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agent/" + dara.StringValue(agentName) + "?status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetAgentStatusResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a migration task.
//
// Description:
//
//	To query the information about a migration task, you must have the permission on mgw:GetImportJob.
//
// @param request - GetJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResponse
func (client *Client) GetJobWithContext(ctx context.Context, userid *string, jobName *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.ByVersion) {
		query["byVersion"] = request.ByVersion
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJob"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/job/" + dara.StringValue(jobName)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the list of files that fail to be migrated when files fail to be migrated during a migration task.
//
// Description:
//
//	  To query the retry information about a migration task, you must have the permission on mgw:GetImportJobResult.
//
//		- If files fail to be migrated during a migration task, a list of files that fail to be migrated is generated. You can call this operation to query this list. You can create a data address based on this list and create a subtask. This way, you can migrate these files again.
//
// @param request - GetJobResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetJobResultResponse
func (client *Client) GetJobResultWithContext(ctx context.Context, userid *string, jobName *string, request *GetJobResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuntimeId) {
		query["runtimeId"] = request.RuntimeId
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetJobResult"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/job/" + dara.StringValue(jobName) + "?jobResult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetJobResultResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a migration report.
//
// Description:
//
//	  To query the information about a migration report, you must have the permission on mgw:GetImportReport.
//
//		- The migration report is pushed to the destination data address. For more information, see the "View a migration report" section of the "Subsequent operations" topic in migration tutorials.
//
// @param request - GetReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportResponse
func (client *Client) GetReportWithContext(ctx context.Context, userid *string, request *GetReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.RuntimeId) {
		query["runtimeId"] = request.RuntimeId
	}

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReport"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/report"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetReportResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the details of a tunnel.
//
// Description:
//
//	To query the information about a tunnel, you must have the permission on mgw:GetImportTunnel.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTunnelResponse
func (client *Client) GetTunnelWithContext(ctx context.Context, userid *string, tunnelId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTunnelResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTunnel"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tunnel/" + dara.StringValue(tunnelId)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &GetTunnelResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the data addresses created by a user in the specific region.
//
// Description:
//
//	To query a list of data addresses, you must have the permission on mgw:ListImportAddress.
//
// @param request - ListAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddressResponse
func (client *Client) ListAddressWithContext(ctx context.Context, userid *string, request *ListAddressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAddress"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/addresslist"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &ListAddressResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the agents created by a user in the specific region.
//
// Description:
//
//	To query a list of agents, you must have the permission on mgw:ListImportAgent.
//
// @param request - ListAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentResponse
func (client *Client) ListAgentWithContext(ctx context.Context, userid *string, request *ListAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgent"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentlist"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &ListAgentResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the migration tasks created by a user in the specific region.
//
// Description:
//
//	To query a list of migration tasks, you must have the permission on mgw:ListImportJob.
//
// @param request - ListJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobResponse
func (client *Client) ListJobWithContext(ctx context.Context, userid *string, request *ListJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["all"] = request.All
	}

	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.ParentName) {
		query["parentName"] = request.ParentName
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJob"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/joblist"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &ListJobResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the running history of a migration task.
//
// Description:
//
//	  To query the execution history of a migration task, you must have the permission on mgw:ListImportJobHistory.
//
//		- A migration task can run multiple rounds. A unique execution ID is generated for each round.
//
//		- The execution history of a migration task records the change history of the task status.
//
// @param request - ListJobHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobHistoryResponse
func (client *Client) ListJobHistoryWithContext(ctx context.Context, userid *string, jobName *string, request *ListJobHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	if !dara.IsNil(request.RuntimeId) {
		query["runtimeId"] = request.RuntimeId
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListJobHistory"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/jobhistory/" + dara.StringValue(jobName)),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &ListJobHistoryResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists tunnels.
//
// Description:
//
//	To query a list of tunnels, you must have the permission on mgw:ListImportTunnel.
//
// @param request - ListTunnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTunnelResponse
func (client *Client) ListTunnelWithContext(ctx context.Context, userid *string, request *ListTunnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTunnelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	query := map[string]interface{}{}
	if !dara.IsNil(request.Count) {
		query["count"] = request.Count
	}

	if !dara.IsNil(request.Marker) {
		query["marker"] = request.Marker
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTunnel"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tunnellist"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &ListTunnelResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a data address.
//
// Description:
//
//	  To update a data address, you must have the permission on mgw:UpdateImportAddress.
//
//		- If the data address is associated with an agent, you can scale up or down the agent.
//
// @param request - UpdateAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAddressResponse
func (client *Client) UpdateAddressWithContext(ctx context.Context, userid *string, addressName *string, request *UpdateAddressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAddressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportAddress) {
		body["ImportAddress"] = request.ImportAddress
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAddress"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/address/" + dara.StringValue(addressName)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &UpdateAddressResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status or throttling of a task.
//
// Description:
//
//	  To update a migration task, you must have the permission on mgw:UpdateImportJob.
//
//		- You can update only the status or throttling settings of a task in a single request.
//
// @param request - UpdateJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateJobResponse
func (client *Client) UpdateJobWithContext(ctx context.Context, userid *string, jobName *string, request *UpdateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportJob) {
		body["ImportJob"] = request.ImportJob
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateJob"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/job/" + dara.StringValue(jobName)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tunnel.
//
// Description:
//
//	To update a tunnel, you must have the permission on mgw:UpdateImportTunnel.
//
// @param request - UpdateTunnelRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTunnelResponse
func (client *Client) UpdateTunnelWithContext(ctx context.Context, userid *string, tunnelId *string, request *UpdateTunnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTunnelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImportTunnel) {
		body["ImportTunnel"] = request.ImportTunnel
	}

	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTunnel"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/tunnel/" + dara.StringValue(tunnelId)),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &UpdateTunnelResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Verifies whether a data address is available.
//
// Description:
//
//	  To verify a data address, you must have the permission on mgw:VerifyImportAddress.
//
//		- A data address may not be available even if the data address passes the availability verification. The data migration results prevail.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyAddressResponse
func (client *Client) VerifyAddressWithContext(ctx context.Context, userid *string, addressName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyAddressResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["userid"] = userid
	req := &openapiutil.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyAddress"),
		Version:     dara.String("2024-06-26"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/address/" + dara.StringValue(addressName) + "?verify"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("xml"),
		BodyType:    dara.String("xml"),
	}
	_result = &VerifyAddressResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
