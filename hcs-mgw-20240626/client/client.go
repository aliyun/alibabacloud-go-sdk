// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-oss/client"
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
	client.ProductId = dara.String("hcs-mgw")
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":     dara.String("cn-beijing.mgw.aliyuncs.com"),
		"cn-wulanchabu":  dara.String("cn-wulanchabu.mgw.aliyuncs.com"),
		"cn-hangzhou":    dara.String("cn-hangzhou.mgw.aliyuncs.com"),
		"cn-shanghai":    dara.String("cn-shanghai.mgw.aliyuncs.com"),
		"cn-shenzhen":    dara.String("cn-shenzhen.mgw.aliyuncs.com"),
		"cn-chengdu":     dara.String("cn-chengdu.mgw.aliyuncs.com"),
		"cn-hongkong":    dara.String("cn-hongkong.mgw.aliyuncs.com"),
		"ap-northeast-1": dara.String("ap-northeast-1.mgw.aliyuncs.com"),
		"ap-southeast-1": dara.String("ap-southeast-1.mgw.aliyuncs.com"),
		"eu-central-1":   dara.String("eu-central-1.mgw.aliyuncs.com"),
		"us-east-1":      dara.String("us-east-1.mgw.aliyuncs.com"),
		"us-west-1":      dara.String("us-west-1.mgw.aliyuncs.com"),
	}
	return nil
}

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
func (client *Client) CreateAddressWithOptions(userid *string, request *CreateAddressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAddressResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

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
// @return CreateAddressResponse
func (client *Client) CreateAddress(userid *string, request *CreateAddressRequest) (_result *CreateAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAddressResponse{}
	_body, _err := client.CreateAddressWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAgentWithOptions(userid *string, request *CreateAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return CreateAgentResponse
func (client *Client) CreateAgent(userid *string, request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateJobWithOptions(userid *string, request *CreateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateJobResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return CreateJobResponse
func (client *Client) CreateJob(userid *string, request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateReportWithOptions(userid *string, request *CreateReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateReportResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return CreateReportResponse
func (client *Client) CreateReport(userid *string, request *CreateReportRequest) (_result *CreateReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateReportResponse{}
	_body, _err := client.CreateReportWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTunnelWithOptions(userid *string, request *CreateTunnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTunnelResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return CreateTunnelResponse
func (client *Client) CreateTunnel(userid *string, request *CreateTunnelRequest) (_result *CreateTunnelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTunnelResponse{}
	_body, _err := client.CreateTunnelWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAddressWithOptions(userid *string, addressName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAddressResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return DeleteAddressResponse
func (client *Client) DeleteAddress(userid *string, addressName *string) (_result *DeleteAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAddressResponse{}
	_body, _err := client.DeleteAddressWithOptions(userid, addressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAgentWithOptions(userid *string, agentName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(userid *string, agentName *string) (_result *DeleteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(userid, agentName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteJobWithOptions(userid *string, jobName *string, request *DeleteJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return DeleteJobResponse
func (client *Client) DeleteJob(userid *string, jobName *string, request *DeleteJobRequest) (_result *DeleteJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteTunnelWithOptions(userid *string, tunnelId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTunnelResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return DeleteTunnelResponse
func (client *Client) DeleteTunnel(userid *string, tunnelId *string) (_result *DeleteTunnelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTunnelResponse{}
	_body, _err := client.DeleteTunnelWithOptions(userid, tunnelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the details of a data address.
//
// Description:
//
// - To query the information about a data address, you must have the permission on mgw:GetImportAddress.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAddressResponse
func (client *Client) GetAddressWithOptions(userid *string, addressName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAddressResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// - To query the information about a data address, you must have the permission on mgw:GetImportAddress.
//
// @return GetAddressResponse
func (client *Client) GetAddress(userid *string, addressName *string) (_result *GetAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddressResponse{}
	_body, _err := client.GetAddressWithOptions(userid, addressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAgentWithOptions(userid *string, agentName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return GetAgentResponse
func (client *Client) GetAgent(userid *string, agentName *string) (_result *GetAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentResponse{}
	_body, _err := client.GetAgentWithOptions(userid, agentName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAgentStatusWithOptions(userid *string, agentName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentStatusResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return GetAgentStatusResponse
func (client *Client) GetAgentStatus(userid *string, agentName *string) (_result *GetAgentStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentStatusResponse{}
	_body, _err := client.GetAgentStatusWithOptions(userid, agentName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetJobWithOptions(userid *string, jobName *string, request *GetJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return GetJobResponse
func (client *Client) GetJob(userid *string, jobName *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetJobResultWithOptions(userid *string, jobName *string, request *GetJobResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetJobResultResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return GetJobResultResponse
func (client *Client) GetJobResult(userid *string, jobName *string, request *GetJobResultRequest) (_result *GetJobResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResultResponse{}
	_body, _err := client.GetJobResultWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetReportWithOptions(userid *string, request *GetReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReportResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return GetReportResponse
func (client *Client) GetReport(userid *string, request *GetReportRequest) (_result *GetReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReportResponse{}
	_body, _err := client.GetReportWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTunnelWithOptions(userid *string, tunnelId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTunnelResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return GetTunnelResponse
func (client *Client) GetTunnel(userid *string, tunnelId *string) (_result *GetTunnelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTunnelResponse{}
	_body, _err := client.GetTunnelWithOptions(userid, tunnelId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the data addresses created by a user in the specific region.
//
// Description:
//
// - To query a list of data addresses, you must have the permission on mgw:ListImportAddress.
//
// @param request - ListAddressRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddressResponse
func (client *Client) ListAddressWithOptions(userid *string, request *ListAddressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAddressResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// - To query a list of data addresses, you must have the permission on mgw:ListImportAddress.
//
// @param request - ListAddressRequest
//
// @return ListAddressResponse
func (client *Client) ListAddress(userid *string, request *ListAddressRequest) (_result *ListAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAddressResponse{}
	_body, _err := client.ListAddressWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAgentWithOptions(userid *string, request *ListAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return ListAgentResponse
func (client *Client) ListAgent(userid *string, request *ListAgentRequest) (_result *ListAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentResponse{}
	_body, _err := client.ListAgentWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListJobWithOptions(userid *string, request *ListJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return ListJobResponse
func (client *Client) ListJob(userid *string, request *ListJobRequest) (_result *ListJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobResponse{}
	_body, _err := client.ListJobWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the run history of a migration task.
//
// Description:
//
// - Call this operation to retrieve the run history of a specified migration task. A migration task can have multiple runs, each identified by a unique execution ID. The run history records status changes that occur during each execution.
//
// - Required permission: mgw:ListImportJobHistory
//
// @param request - ListJobHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListJobHistoryResponse
func (client *Client) ListJobHistoryWithOptions(userid *string, jobName *string, request *ListJobHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListJobHistoryResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the run history of a migration task.
//
// Description:
//
// - Call this operation to retrieve the run history of a specified migration task. A migration task can have multiple runs, each identified by a unique execution ID. The run history records status changes that occur during each execution.
//
// - Required permission: mgw:ListImportJobHistory
//
// @param request - ListJobHistoryRequest
//
// @return ListJobHistoryResponse
func (client *Client) ListJobHistory(userid *string, jobName *string, request *ListJobHistoryRequest) (_result *ListJobHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobHistoryResponse{}
	_body, _err := client.ListJobHistoryWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTunnelWithOptions(userid *string, request *ListTunnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTunnelResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return ListTunnelResponse
func (client *Client) ListTunnel(userid *string, request *ListTunnelRequest) (_result *ListTunnelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTunnelResponse{}
	_body, _err := client.ListTunnelWithOptions(userid, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAddressWithOptions(userid *string, addressName *string, request *UpdateAddressRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAddressResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return UpdateAddressResponse
func (client *Client) UpdateAddress(userid *string, addressName *string, request *UpdateAddressRequest) (_result *UpdateAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAddressResponse{}
	_body, _err := client.UpdateAddressWithOptions(userid, addressName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateJobWithOptions(userid *string, jobName *string, request *UpdateJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return UpdateJobResponse
func (client *Client) UpdateJob(userid *string, jobName *string, request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(userid, jobName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateTunnelWithOptions(userid *string, tunnelId *string, request *UpdateTunnelRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTunnelResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return UpdateTunnelResponse
func (client *Client) UpdateTunnel(userid *string, tunnelId *string, request *UpdateTunnelRequest) (_result *UpdateTunnelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTunnelResponse{}
	_body, _err := client.UpdateTunnelWithOptions(userid, tunnelId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VerifyAddressWithOptions(userid *string, addressName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *VerifyAddressResponse, _err error) {
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
	_body, _err := client.Execute(params, req, runtime)
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
// @return VerifyAddressResponse
func (client *Client) VerifyAddress(userid *string, addressName *string) (_result *VerifyAddressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyAddressResponse{}
	_body, _err := client.VerifyAddressWithOptions(userid, addressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
