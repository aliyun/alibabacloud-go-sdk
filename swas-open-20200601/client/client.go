// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddCustomImageShareAccountRequest struct {
	// This parameter is required.
	Account []*int64 `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-saacssasc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddCustomImageShareAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s AddCustomImageShareAccountRequest) GoString() string {
	return s.String()
}

func (s *AddCustomImageShareAccountRequest) SetAccount(v []*int64) *AddCustomImageShareAccountRequest {
	s.Account = v
	return s
}

func (s *AddCustomImageShareAccountRequest) SetClientToken(v string) *AddCustomImageShareAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCustomImageShareAccountRequest) SetImageId(v string) *AddCustomImageShareAccountRequest {
	s.ImageId = &v
	return s
}

func (s *AddCustomImageShareAccountRequest) SetRegionId(v string) *AddCustomImageShareAccountRequest {
	s.RegionId = &v
	return s
}

type AddCustomImageShareAccountResponseBody struct {
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCustomImageShareAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddCustomImageShareAccountResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomImageShareAccountResponseBody) SetRequestId(v string) *AddCustomImageShareAccountResponseBody {
	s.RequestId = &v
	return s
}

type AddCustomImageShareAccountResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomImageShareAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomImageShareAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s AddCustomImageShareAccountResponse) GoString() string {
	return s.String()
}

func (s *AddCustomImageShareAccountResponse) SetHeaders(v map[string]*string) *AddCustomImageShareAccountResponse {
	s.Headers = v
	return s
}

func (s *AddCustomImageShareAccountResponse) SetStatusCode(v int32) *AddCustomImageShareAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomImageShareAccountResponse) SetBody(v *AddCustomImageShareAccountResponseBody) *AddCustomImageShareAccountResponse {
	s.Body = v
	return s
}

type AllocatePublicConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AllocatePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionRequest) SetClientToken(v string) *AllocatePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *AllocatePublicConnectionRequest) SetDatabaseInstanceId(v string) *AllocatePublicConnectionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *AllocatePublicConnectionRequest) SetRegionId(v string) *AllocatePublicConnectionRequest {
	s.RegionId = &v
	return s
}

type AllocatePublicConnectionResponseBody struct {
	// The public endpoint that is assigned to the Simple Database Service instance.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****.mysql.rds.aliyuncs.com
	PublicConnection *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocatePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionResponseBody) SetPublicConnection(v string) *AllocatePublicConnectionResponseBody {
	s.PublicConnection = &v
	return s
}

func (s *AllocatePublicConnectionResponseBody) SetRequestId(v string) *AllocatePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type AllocatePublicConnectionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocatePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocatePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocatePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocatePublicConnectionResponse) SetHeaders(v map[string]*string) *AllocatePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *AllocatePublicConnectionResponse) SetStatusCode(v int32) *AllocatePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocatePublicConnectionResponse) SetBody(v *AllocatePublicConnectionResponseBody) *AllocatePublicConnectionResponse {
	s.Body = v
	return s
}

type ApplyFirewallTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ApplyFirewallTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyFirewallTemplateRequest) GoString() string {
	return s.String()
}

func (s *ApplyFirewallTemplateRequest) SetClientToken(v string) *ApplyFirewallTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetFirewallTemplateId(v string) *ApplyFirewallTemplateRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetInstanceId(v string) *ApplyFirewallTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetInstanceIds(v []*string) *ApplyFirewallTemplateRequest {
	s.InstanceIds = v
	return s
}

func (s *ApplyFirewallTemplateRequest) SetRegionId(v string) *ApplyFirewallTemplateRequest {
	s.RegionId = &v
	return s
}

type ApplyFirewallTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the execution to apply the template.
	//
	// example:
	//
	// aft-ikgt0bynitvs3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ApplyFirewallTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyFirewallTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyFirewallTemplateResponseBody) SetRequestId(v string) *ApplyFirewallTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyFirewallTemplateResponseBody) SetTaskId(v string) *ApplyFirewallTemplateResponseBody {
	s.TaskId = &v
	return s
}

type ApplyFirewallTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFirewallTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFirewallTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyFirewallTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApplyFirewallTemplateResponse) SetHeaders(v map[string]*string) *ApplyFirewallTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApplyFirewallTemplateResponse) SetStatusCode(v int32) *ApplyFirewallTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFirewallTemplateResponse) SetBody(v *ApplyFirewallTemplateResponseBody) *ApplyFirewallTemplateResponse {
	s.Body = v
	return s
}

type AttachKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of simple application servers. You can specify a maximum of 50 IDs of simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the key pair that you want to bind to the simple application server. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_gin
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *AttachKeyPairRequest) SetClientToken(v string) *AttachKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachKeyPairRequest) SetInstanceIds(v []*string) *AttachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *AttachKeyPairRequest) SetKeyPairName(v string) *AttachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *AttachKeyPairRequest) SetRegionId(v string) *AttachKeyPairRequest {
	s.RegionId = &v
	return s
}

type AttachKeyPairResponseBody struct {
	// The total number of simple application servers to which the key pair failed to be bound.
	//
	// example:
	//
	// 2
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results.
	Results []*AttachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of simple application servers to which the key pair is bound.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AttachKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBody) SetFailCount(v int32) *AttachKeyPairResponseBody {
	s.FailCount = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetRequestId(v string) *AttachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachKeyPairResponseBody) SetResults(v []*AttachKeyPairResponseBodyResults) *AttachKeyPairResponseBody {
	s.Results = v
	return s
}

func (s *AttachKeyPairResponseBody) SetTotalCount(v int32) *AttachKeyPairResponseBody {
	s.TotalCount = &v
	return s
}

type AttachKeyPairResponseBodyResults struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The simple application server ID.
	//
	// example:
	//
	// aa6e71ddb35c46679bc4753d6219d604
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the key pair is bound to the simple application server successfully. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachKeyPairResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyResults) SetCode(v string) *AttachKeyPairResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) SetInstanceId(v string) *AttachKeyPairResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) SetMessage(v string) *AttachKeyPairResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AttachKeyPairResponseBodyResults) SetSuccess(v string) *AttachKeyPairResponseBodyResults {
	s.Success = &v
	return s
}

type AttachKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachKeyPairResponse) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponse) SetHeaders(v map[string]*string) *AttachKeyPairResponse {
	s.Headers = v
	return s
}

func (s *AttachKeyPairResponse) SetStatusCode(v int32) *AttachKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachKeyPairResponse) SetBody(v *AttachKeyPairResponseBody) *AttachKeyPairResponse {
	s.Body = v
	return s
}

type CreateCommandRequest struct {
	// The command content. When you specify this parameter, take note of the following items:
	//
	// 	- When `EnableParameter` is set to true, the custom parameter feature is enabled, and you can configure custom parameters in the command based on the following rules:
	//
	// 	- Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	// 	- You can specify up to 20 custom parameters.
	//
	// 	- The name of a custom parameter can contain only letters, digits, underscores (_), and hyphens (-). The name is case-insensitive.
	//
	// 	- The name of a custom parameter cannot exceed 64 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// ifconfig -s
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The description of the command. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use custom parameters in the command.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The name of the command. The name supports all character sets and can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that you want to add to the command. You can specify up to 20 tags.
	Tag []*CreateCommandRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The timeout period for the command execution on the instance.
	//
	// If a command execution task times out, Command Assistant forcefully terminates the task process. Valid values: 10 to 86400. Unit: seconds. The period of 86400 seconds is equal to 24 hours.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances
	//
	// 	- RunShellScript: shell command, applicable to Linux instances
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The working directory of the command on the ECS instance.
	//
	// Default values:
	//
	// 	- For a Linux instance, the default value is the home directory of the root user, which is the `/root` directory.
	//
	// 	- For a Windows instance, the default value is the directory where the Cloud Assistant client process resides. Example: `C:\\Windows\\System32`.
	//
	// example:
	//
	// /root/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateCommandRequest) SetCommandContent(v string) *CreateCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *CreateCommandRequest) SetDescription(v string) *CreateCommandRequest {
	s.Description = &v
	return s
}

func (s *CreateCommandRequest) SetEnableParameter(v bool) *CreateCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *CreateCommandRequest) SetName(v string) *CreateCommandRequest {
	s.Name = &v
	return s
}

func (s *CreateCommandRequest) SetRegionId(v string) *CreateCommandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCommandRequest) SetResourceGroupId(v string) *CreateCommandRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCommandRequest) SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest {
	s.Tag = v
	return s
}

func (s *CreateCommandRequest) SetTimeout(v int64) *CreateCommandRequest {
	s.Timeout = &v
	return s
}

func (s *CreateCommandRequest) SetType(v string) *CreateCommandRequest {
	s.Type = &v
	return s
}

func (s *CreateCommandRequest) SetWorkingDir(v string) *CreateCommandRequest {
	s.WorkingDir = &v
	return s
}

type CreateCommandRequestTag struct {
	// The key of tag N that you want to add to the command. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the command. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCommandRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateCommandRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCommandRequestTag) SetKey(v string) *CreateCommandRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCommandRequestTag) SetValue(v string) *CreateCommandRequestTag {
	s.Value = &v
	return s
}

type CreateCommandResponseBody struct {
	// The command ID.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCommandResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCommandResponseBody) SetCommandId(v string) *CreateCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *CreateCommandResponseBody) SetRequestId(v string) *CreateCommandResponseBody {
	s.RequestId = &v
	return s
}

type CreateCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCommandResponse) GoString() string {
	return s.String()
}

func (s *CreateCommandResponse) SetHeaders(v map[string]*string) *CreateCommandResponse {
	s.Headers = v
	return s
}

func (s *CreateCommandResponse) SetStatusCode(v int32) *CreateCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCommandResponse) SetBody(v *CreateCommandResponseBody) *CreateCommandResponse {
	s.Body = v
	return s
}

type CreateCustomImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data disk snapshot.
	//
	// example:
	//
	// s-acscasca****
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The description of the custom image.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter or a digit. This parameter is empty by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// customImage-test
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the database. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResoureGroupId *string `json:"ResoureGroupId,omitempty" xml:"ResoureGroupId,omitempty"`
	// The ID of the system disk snapshot.
	//
	// example:
	//
	// s-acscasca****
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// The tags that you want to add to the custom image. You can specify up to 20 tags.
	Tag []*CreateCustomImageRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCustomImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequest) SetClientToken(v string) *CreateCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCustomImageRequest) SetDataSnapshotId(v string) *CreateCustomImageRequest {
	s.DataSnapshotId = &v
	return s
}

func (s *CreateCustomImageRequest) SetDescription(v string) *CreateCustomImageRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomImageRequest) SetImageName(v string) *CreateCustomImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateCustomImageRequest) SetInstanceId(v string) *CreateCustomImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomImageRequest) SetRegionId(v string) *CreateCustomImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCustomImageRequest) SetResoureGroupId(v string) *CreateCustomImageRequest {
	s.ResoureGroupId = &v
	return s
}

func (s *CreateCustomImageRequest) SetSystemSnapshotId(v string) *CreateCustomImageRequest {
	s.SystemSnapshotId = &v
	return s
}

func (s *CreateCustomImageRequest) SetTag(v []*CreateCustomImageRequestTag) *CreateCustomImageRequest {
	s.Tag = v
	return s
}

type CreateCustomImageRequestTag struct {
	// The key of tag N that you want to add to the custom image. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the custom image. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCustomImageRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCustomImageRequestTag) SetKey(v string) *CreateCustomImageRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCustomImageRequestTag) SetValue(v string) *CreateCustomImageRequestTag {
	s.Value = &v
	return s
}

type CreateCustomImageResponseBody struct {
	// The custom image ID.
	//
	// example:
	//
	// m-csaascsaccscs****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponseBody) SetImageId(v string) *CreateCustomImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateCustomImageResponseBody) SetRequestId(v string) *CreateCustomImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomImageResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomImageResponse) SetHeaders(v map[string]*string) *CreateCustomImageResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomImageResponse) SetStatusCode(v int32) *CreateCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomImageResponse) SetBody(v *CreateCustomImageResponseBody) *CreateCustomImageResponse {
	s.Body = v
	return s
}

type CreateFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range.
	//
	// 	- When the transport layer protocol is TCP and/or UDP, the port range is 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: 1024/1055, which specifies the port range from 1024 to 1055.
	//
	// 	- When the transport layer protocol is ICMP, the port range is -1/-1, which indicates all ports.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// TEST
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
}

func (s CreateFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleRequest) SetClientToken(v string) *CreateFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetInstanceId(v string) *CreateFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetPort(v string) *CreateFirewallRuleRequest {
	s.Port = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRegionId(v string) *CreateFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRemark(v string) *CreateFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *CreateFirewallRuleRequest) SetRuleProtocol(v string) *CreateFirewallRuleRequest {
	s.RuleProtocol = &v
	return s
}

type CreateFirewallRuleResponseBody struct {
	// The ID of the firewall rule.
	//
	// example:
	//
	// 8007e18c61024aafbd776d52d0****
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleResponseBody) SetFirewallId(v string) *CreateFirewallRuleResponseBody {
	s.FirewallId = &v
	return s
}

func (s *CreateFirewallRuleResponseBody) SetRequestId(v string) *CreateFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallRuleResponse) SetHeaders(v map[string]*string) *CreateFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallRuleResponse) SetStatusCode(v int32) *CreateFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallRuleResponse) SetBody(v *CreateFirewallRuleResponseBody) *CreateFirewallRuleResponse {
	s.Body = v
	return s
}

type CreateFirewallRulesRequest struct {
	// The client token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Details about the firewall rules.
	FirewallRules []*CreateFirewallRulesRequestFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags that you want to add to the firewall. You can specify up to 20 tags.
	Tag []*CreateFirewallRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateFirewallRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesRequest) SetClientToken(v string) *CreateFirewallRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRulesRequest) SetFirewallRules(v []*CreateFirewallRulesRequestFirewallRules) *CreateFirewallRulesRequest {
	s.FirewallRules = v
	return s
}

func (s *CreateFirewallRulesRequest) SetInstanceId(v string) *CreateFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRulesRequest) SetRegionId(v string) *CreateFirewallRulesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRulesRequest) SetTag(v []*CreateFirewallRulesRequestTag) *CreateFirewallRulesRequest {
	s.Tag = v
	return s
}

type CreateFirewallRulesRequestFirewallRules struct {
	// The port number.
	//
	// 	- When the transport layer protocol is TCP and/or UDP, the port number range is 1 to 65535. Specify a port range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- When the transport layer protocol is ICMP, the port number range is -1/-1, which indicates all ports.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The description of the firewall rule.
	//
	// example:
	//
	// TEST
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The IP address or CIDR block that is allowed in the firewall rule.
	//
	// example:
	//
	// 47.101.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallRulesRequestFirewallRules) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesRequestFirewallRules) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesRequestFirewallRules) SetPort(v string) *CreateFirewallRulesRequestFirewallRules {
	s.Port = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) SetRemark(v string) *CreateFirewallRulesRequestFirewallRules {
	s.Remark = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) SetRuleProtocol(v string) *CreateFirewallRulesRequestFirewallRules {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallRulesRequestFirewallRules) SetSourceCidrIp(v string) *CreateFirewallRulesRequestFirewallRules {
	s.SourceCidrIp = &v
	return s
}

type CreateFirewallRulesRequestTag struct {
	// The tag key. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFirewallRulesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesRequestTag) SetKey(v string) *CreateFirewallRulesRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFirewallRulesRequestTag) SetValue(v string) *CreateFirewallRulesRequestTag {
	s.Value = &v
	return s
}

type CreateFirewallRulesShrinkRequest struct {
	// The client token.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Details about the firewall rules.
	FirewallRulesShrink *string `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags that you want to add to the firewall. You can specify up to 20 tags.
	Tag []*CreateFirewallRulesShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateFirewallRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesShrinkRequest) SetClientToken(v string) *CreateFirewallRulesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetFirewallRulesShrink(v string) *CreateFirewallRulesShrinkRequest {
	s.FirewallRulesShrink = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetInstanceId(v string) *CreateFirewallRulesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetRegionId(v string) *CreateFirewallRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequest) SetTag(v []*CreateFirewallRulesShrinkRequestTag) *CreateFirewallRulesShrinkRequest {
	s.Tag = v
	return s
}

type CreateFirewallRulesShrinkRequestTag struct {
	// The tag key. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFirewallRulesShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesShrinkRequestTag) SetKey(v string) *CreateFirewallRulesShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFirewallRulesShrinkRequestTag) SetValue(v string) *CreateFirewallRulesShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateFirewallRulesResponseBody struct {
	// The IDs of the firewall rules that you created.
	FirewallRuleIds []*string `json:"FirewallRuleIds,omitempty" xml:"FirewallRuleIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesResponseBody) SetFirewallRuleIds(v []*string) *CreateFirewallRulesResponseBody {
	s.FirewallRuleIds = v
	return s
}

func (s *CreateFirewallRulesResponseBody) SetRequestId(v string) *CreateFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

type CreateFirewallRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesResponse) SetHeaders(v map[string]*string) *CreateFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallRulesResponse) SetStatusCode(v int32) *CreateFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallRulesResponse) SetBody(v *CreateFirewallRulesResponseBody) *CreateFirewallRulesResponse {
	s.Body = v
	return s
}

type CreateFirewallTemplateRequest struct {
	// The description of the firewall template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The details of the firewall rule.
	FirewallRule []*CreateFirewallTemplateRequestFirewallRule `json:"FirewallRule,omitempty" xml:"FirewallRule,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the simple application server to which the firewall template belongs. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateFirewallTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRequest) SetDescription(v string) *CreateFirewallTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateFirewallTemplateRequest) SetFirewallRule(v []*CreateFirewallTemplateRequestFirewallRule) *CreateFirewallTemplateRequest {
	s.FirewallRule = v
	return s
}

func (s *CreateFirewallTemplateRequest) SetName(v string) *CreateFirewallTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateFirewallTemplateRequest) SetRegionId(v string) *CreateFirewallTemplateRequest {
	s.RegionId = &v
	return s
}

type CreateFirewallTemplateRequestFirewallRule struct {
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 223.166.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallTemplateRequestFirewallRule) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRequestFirewallRule) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetPort(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.Port = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetRemark(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.Remark = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetRuleProtocol(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallTemplateRequestFirewallRule) SetSourceCidrIp(v string) *CreateFirewallTemplateRequestFirewallRule {
	s.SourceCidrIp = &v
	return s
}

type CreateFirewallTemplateResponseBody struct {
	// The ID of the firewall template.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateResponseBody) SetFirewallTemplateId(v string) *CreateFirewallTemplateResponseBody {
	s.FirewallTemplateId = &v
	return s
}

func (s *CreateFirewallTemplateResponseBody) SetRequestId(v string) *CreateFirewallTemplateResponseBody {
	s.RequestId = &v
	return s
}

type CreateFirewallTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateResponse) SetHeaders(v map[string]*string) *CreateFirewallTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallTemplateResponse) SetStatusCode(v int32) *CreateFirewallTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallTemplateResponse) SetBody(v *CreateFirewallTemplateResponseBody) *CreateFirewallTemplateResponse {
	s.Body = v
	return s
}

type CreateFirewallTemplateRulesRequest struct {
	// The details of the firewall rule.
	//
	// This parameter is required.
	FirewallRule []*CreateFirewallTemplateRulesRequestFirewallRule `json:"FirewallRule,omitempty" xml:"FirewallRule,omitempty" type:"Repeated"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The region ID of the simple application server to which the firewall template is applied. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateFirewallTemplateRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesRequest) SetFirewallRule(v []*CreateFirewallTemplateRulesRequestFirewallRule) *CreateFirewallTemplateRulesRequest {
	s.FirewallRule = v
	return s
}

func (s *CreateFirewallTemplateRulesRequest) SetFirewallTemplateId(v string) *CreateFirewallTemplateRulesRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequest) SetRegionId(v string) *CreateFirewallTemplateRulesRequest {
	s.RegionId = &v
	return s
}

type CreateFirewallTemplateRulesRequestFirewallRule struct {
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 222.70.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallTemplateRulesRequestFirewallRule) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRulesRequestFirewallRule) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetPort(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.Port = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetRemark(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.Remark = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetRuleProtocol(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallTemplateRulesRequestFirewallRule) SetSourceCidrIp(v string) *CreateFirewallTemplateRulesRequestFirewallRule {
	s.SourceCidrIp = &v
	return s
}

type CreateFirewallTemplateRulesResponseBody struct {
	// The firewall template rules.
	FirewallTemplateRules []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules `json:"FirewallTemplateRules,omitempty" xml:"FirewallTemplateRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallTemplateRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesResponseBody) SetFirewallTemplateRules(v []*CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) *CreateFirewallTemplateRulesResponseBody {
	s.FirewallTemplateRules = v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBody) SetRequestId(v string) *CreateFirewallTemplateRulesResponseBody {
	s.RequestId = &v
	return s
}

type CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules struct {
	// The ID of the firewall template rule.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateRuleId *string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty"`
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP: the ICMP protocol
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 171.233.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetFirewallTemplateRuleId(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.FirewallTemplateRuleId = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetPort(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.Port = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetRemark(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.Remark = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetRuleProtocol(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.RuleProtocol = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules) SetSourceCidrIp(v string) *CreateFirewallTemplateRulesResponseBodyFirewallTemplateRules {
	s.SourceCidrIp = &v
	return s
}

type CreateFirewallTemplateRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallTemplateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallTemplateRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallTemplateRulesResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateRulesResponse) SetHeaders(v map[string]*string) *CreateFirewallTemplateRulesResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallTemplateRulesResponse) SetStatusCode(v int32) *CreateFirewallTemplateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallTemplateRulesResponse) SetBody(v *CreateFirewallTemplateRulesResponseBody) *CreateFirewallTemplateRulesResponse {
	s.Body = v
	return s
}

type CreateInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// ceshi
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstanceKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceKeyPairRequest) SetClientToken(v string) *CreateInstanceKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceKeyPairRequest) SetInstanceId(v string) *CreateInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceKeyPairRequest) SetKeyPairName(v string) *CreateInstanceKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceKeyPairRequest) SetRegionId(v string) *CreateInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

type CreateInstanceKeyPairResponseBody struct {
	// The fingerprint of the key pair.
	//
	// example:
	//
	// If2K1ItazA4GlKkWCEhdRj8Wd6czAvK9*****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// ceshi
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key.
	//
	// example:
	//
	// ***
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceKeyPairResponseBody) SetFingerprint(v string) *CreateInstanceKeyPairResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) SetKeyPairName(v string) *CreateInstanceKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) SetPrivateKey(v string) *CreateInstanceKeyPairResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *CreateInstanceKeyPairResponseBody) SetRequestId(v string) *CreateInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceKeyPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceKeyPairResponse) SetHeaders(v map[string]*string) *CreateInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceKeyPairResponse) SetStatusCode(v int32) *CreateInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceKeyPairResponse) SetBody(v *CreateInstanceKeyPairResponseBody) *CreateInstanceKeyPairResponse {
	s.Body = v
	return s
}

type CreateInstancesRequest struct {
	// The number of simple application servers that you want to create. Valid values: 1 to 20.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. This parameter is required only when you set `AutoRenew` to true. Unit: months. Valid values: 1, 3, 6, 12, 24, and 36.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The billing method of the simple application servers. Set the value to PrePaid, which indicates the subscription billing method.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The size of the data disk that is attached to the server. Unit: GB. Valid values: 0 to 16380. The value must be an integral multiple of 20.
	//
	// 	- A value of 0 indicates that no data disk is attached.
	//
	// 	- If the disk included in the specified plan is a standard SSD, the data disk must be 20 GB or larger in size.
	//
	// Default value: 0.
	//
	// example:
	//
	// 20
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The image ID. You can call the [ListImages](https://help.aliyun.com/document_detail/189313.html) operation to query the available images in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// e2c9c365024a44369c9b955a998a****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The subscription period of the servers. Unit: months. Valid values: 1, 3, 6, 12, 24, and 36.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The plan ID. You can call the [ListPlans](https://help.aliyun.com/document_detail/189314.html) operation to query all plans provided by Simple Application Server in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// swas.s1.c1m1s40b3t05
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateInstancesRequest) SetAmount(v int32) *CreateInstancesRequest {
	s.Amount = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenew(v bool) *CreateInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstancesRequest) SetAutoRenewPeriod(v int32) *CreateInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstancesRequest) SetChargeType(v string) *CreateInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstancesRequest) SetClientToken(v string) *CreateInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstancesRequest) SetDataDiskSize(v int64) *CreateInstancesRequest {
	s.DataDiskSize = &v
	return s
}

func (s *CreateInstancesRequest) SetImageId(v string) *CreateInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstancesRequest) SetPeriod(v int32) *CreateInstancesRequest {
	s.Period = &v
	return s
}

func (s *CreateInstancesRequest) SetPlanId(v string) *CreateInstancesRequest {
	s.PlanId = &v
	return s
}

func (s *CreateInstancesRequest) SetRegionId(v string) *CreateInstancesRequest {
	s.RegionId = &v
	return s
}

type CreateInstancesResponseBody struct {
	// The IDs of the simple application servers.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E1FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponseBody) SetInstanceIds(v []*string) *CreateInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateInstancesResponseBody) SetRequestId(v string) *CreateInstancesResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateInstancesResponse) SetHeaders(v map[string]*string) *CreateInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateInstancesResponse) SetStatusCode(v int32) *CreateInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstancesResponse) SetBody(v *CreateInstancesResponseBody) *CreateInstancesResponse {
	s.Body = v
	return s
}

type CreateKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the key pair. The name must be 2 to 64 characters in length and can contain letters, digits, colons (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) SetClientToken(v string) *CreateKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairRequest) SetRegionId(v string) *CreateKeyPairRequest {
	s.RegionId = &v
	return s
}

type CreateKeyPairResponseBody struct {
	// The name of the key pair. The name must be 2 to 64 characters in length and can contain letters, digits, colons (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with http:// or https://.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key of the key pair. The PEM-encoded private key is in PKCS#8 format.
	//
	// example:
	//
	// MIIEpAIBAAKCAQEAtReyMzLIcBH78EV2zj****
	PrivateKeyBody *string `json:"PrivateKeyBody,omitempty" xml:"PrivateKeyBody,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponseBody) SetKeyPairName(v string) *CreateKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetPrivateKeyBody(v string) *CreateKeyPairResponseBody {
	s.PrivateKeyBody = &v
	return s
}

func (s *CreateKeyPairResponseBody) SetRequestId(v string) *CreateKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type CreateKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyPairResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyPairResponse) SetHeaders(v map[string]*string) *CreateKeyPairResponse {
	s.Headers = v
	return s
}

func (s *CreateKeyPairResponse) SetStatusCode(v int32) *CreateKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKeyPairResponse) SetBody(v *CreateKeyPairResponseBody) *CreateKeyPairResponse {
	s.Body = v
	return s
}

type CreateSnapshotRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp18kjxg9ebrhsgi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server to which the disk is attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot name. The name must be 2 to 50 characters in length. It must start with a letter but cannot start with `http://` or `https://`. The name can only contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-SnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The tags that you want to add to the snapshot. You can specify up to 20 tags.
	Tag []*CreateSnapshotRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) SetClientToken(v string) *CreateSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSnapshotRequest) SetDiskId(v string) *CreateSnapshotRequest {
	s.DiskId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRegionId(v string) *CreateSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnapshotRequest) SetResourceGroupId(v string) *CreateSnapshotRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) SetTag(v []*CreateSnapshotRequestTag) *CreateSnapshotRequest {
	s.Tag = v
	return s
}

type CreateSnapshotRequestTag struct {
	// The key of the tag to add to the snapshot. Valid values of N: 1 to 20.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the snapshot. Valid values of N: 1 to 20.
	//
	// The tag value can be an empty string. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSnapshotRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotRequestTag) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequestTag) SetKey(v string) *CreateSnapshotRequestTag {
	s.Key = &v
	return s
}

func (s *CreateSnapshotRequestTag) SetValue(v string) *CreateSnapshotRequestTag {
	s.Value = &v
	return s
}

type CreateSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponseBody) SetRequestId(v string) *CreateSnapshotResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSnapshotResponseBody) SetSnapshotId(v string) *CreateSnapshotResponseBody {
	s.SnapshotId = &v
	return s
}

type CreateSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSnapshotResponse) GoString() string {
	return s.String()
}

func (s *CreateSnapshotResponse) SetHeaders(v map[string]*string) *CreateSnapshotResponse {
	s.Headers = v
	return s
}

func (s *CreateSnapshotResponse) SetStatusCode(v int32) *CreateSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSnapshotResponse) SetBody(v *CreateSnapshotResponseBody) *CreateSnapshotResponse {
	s.Body = v
	return s
}

type DeleteCommandRequest struct {
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommandRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommandRequest) SetCommandId(v string) *DeleteCommandRequest {
	s.CommandId = &v
	return s
}

func (s *DeleteCommandRequest) SetRegionId(v string) *DeleteCommandRequest {
	s.RegionId = &v
	return s
}

type DeleteCommandResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommandResponseBody) SetRequestId(v string) *DeleteCommandResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCommandResponse) GoString() string {
	return s.String()
}

func (s *DeleteCommandResponse) SetHeaders(v map[string]*string) *DeleteCommandResponse {
	s.Headers = v
	return s
}

func (s *DeleteCommandResponse) SetStatusCode(v int32) *DeleteCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCommandResponse) SetBody(v *DeleteCommandResponseBody) *DeleteCommandResponse {
	s.Body = v
	return s
}

type DeleteCustomImageRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The custom image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-2zehv38jjmwva1ee****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID of the custom image. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageRequest) SetClientToken(v string) *DeleteCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomImageRequest) SetImageId(v string) *DeleteCustomImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteCustomImageRequest) SetRegionId(v string) *DeleteCustomImageRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomImageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageResponseBody) SetRequestId(v string) *DeleteCustomImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageResponse) SetHeaders(v map[string]*string) *DeleteCustomImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomImageResponse) SetStatusCode(v int32) *DeleteCustomImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomImageResponse) SetBody(v *DeleteCustomImageResponseBody) *DeleteCustomImageResponse {
	s.Body = v
	return s
}

type DeleteCustomImagesRequest struct {
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the custom image. The value can be a JSON array that consists of up to 15 image IDs. Separate multiple image IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["m-2zehv38jjmwva1ee****", "m-bp1hj0zhmheyq2kz****"]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomImagesRequest) SetClientToken(v string) *DeleteCustomImagesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomImagesRequest) SetImageIds(v string) *DeleteCustomImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *DeleteCustomImagesRequest) SetRegionId(v string) *DeleteCustomImagesRequest {
	s.RegionId = &v
	return s
}

type DeleteCustomImagesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomImagesResponseBody) SetRequestId(v string) *DeleteCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomImagesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomImagesResponse) SetHeaders(v map[string]*string) *DeleteCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomImagesResponse) SetStatusCode(v int32) *DeleteCustomImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomImagesResponse) SetBody(v *DeleteCustomImagesResponseBody) *DeleteCustomImagesResponse {
	s.Body = v
	return s
}

type DeleteFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the firewall rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleRequest) SetClientToken(v string) *DeleteFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetInstanceId(v string) *DeleteFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetRegionId(v string) *DeleteFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallRuleRequest) SetRuleId(v string) *DeleteFirewallRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleResponseBody) SetRequestId(v string) *DeleteFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRuleResponse) SetHeaders(v map[string]*string) *DeleteFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallRuleResponse) SetStatusCode(v int32) *DeleteFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallRuleResponse) SetBody(v *DeleteFirewallRuleResponseBody) *DeleteFirewallRuleResponse {
	s.Body = v
	return s
}

type DeleteFirewallRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the firewall rules that you want to delete.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteFirewallRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesRequest) SetClientToken(v string) *DeleteFirewallRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetInstanceId(v string) *DeleteFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetRegionId(v string) *DeleteFirewallRulesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallRulesRequest) SetRuleIds(v []*string) *DeleteFirewallRulesRequest {
	s.RuleIds = v
	return s
}

type DeleteFirewallRulesShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the firewall rules that you want to delete.
	RuleIdsShrink *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
}

func (s DeleteFirewallRulesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesShrinkRequest) SetClientToken(v string) *DeleteFirewallRulesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallRulesShrinkRequest) SetInstanceId(v string) *DeleteFirewallRulesShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallRulesShrinkRequest) SetRegionId(v string) *DeleteFirewallRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFirewallRulesShrinkRequest) SetRuleIdsShrink(v string) *DeleteFirewallRulesShrinkRequest {
	s.RuleIdsShrink = &v
	return s
}

type DeleteFirewallRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesResponseBody) SetRequestId(v string) *DeleteFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFirewallRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallRulesResponse) SetHeaders(v map[string]*string) *DeleteFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallRulesResponse) SetStatusCode(v int32) *DeleteFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallRulesResponse) SetBody(v *DeleteFirewallRulesResponseBody) *DeleteFirewallRulesResponse {
	s.Body = v
	return s
}

type DeleteFirewallTemplateRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The IDs of the firewall template rules.
	//
	// This parameter is required.
	FirewallTemplateRuleId []*string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty" type:"Repeated"`
	// The ID of the simple application server to which the firewall template is applied.
	//
	// example:
	//
	// 9ae7106e68eb4402b0dcbd48a9de****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFirewallTemplateRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallTemplateRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplateRulesRequest) SetClientToken(v string) *DeleteFirewallTemplateRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetFirewallTemplateId(v string) *DeleteFirewallTemplateRulesRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetFirewallTemplateRuleId(v []*string) *DeleteFirewallTemplateRulesRequest {
	s.FirewallTemplateRuleId = v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetInstanceId(v string) *DeleteFirewallTemplateRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallTemplateRulesRequest) SetRegionId(v string) *DeleteFirewallTemplateRulesRequest {
	s.RegionId = &v
	return s
}

type DeleteFirewallTemplateRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallTemplateRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallTemplateRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplateRulesResponseBody) SetRequestId(v string) *DeleteFirewallTemplateRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFirewallTemplateRulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallTemplateRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallTemplateRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallTemplateRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplateRulesResponse) SetHeaders(v map[string]*string) *DeleteFirewallTemplateRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallTemplateRulesResponse) SetStatusCode(v int32) *DeleteFirewallTemplateRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallTemplateRulesResponse) SetBody(v *DeleteFirewallTemplateRulesResponseBody) *DeleteFirewallTemplateRulesResponse {
	s.Body = v
	return s
}

type DeleteFirewallTemplatesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the firewall templates.
	//
	// This parameter is required.
	FirewallTemplateId []*string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty" type:"Repeated"`
	// The ID of the simple application server to which the firewall templates belong.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFirewallTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplatesRequest) SetClientToken(v string) *DeleteFirewallTemplatesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFirewallTemplatesRequest) SetFirewallTemplateId(v []*string) *DeleteFirewallTemplatesRequest {
	s.FirewallTemplateId = v
	return s
}

func (s *DeleteFirewallTemplatesRequest) SetInstanceId(v string) *DeleteFirewallTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteFirewallTemplatesRequest) SetRegionId(v string) *DeleteFirewallTemplatesRequest {
	s.RegionId = &v
	return s
}

type DeleteFirewallTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFirewallTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplatesResponseBody) SetRequestId(v string) *DeleteFirewallTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFirewallTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFirewallTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplatesResponse) SetHeaders(v map[string]*string) *DeleteFirewallTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallTemplatesResponse) SetStatusCode(v int32) *DeleteFirewallTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallTemplatesResponse) SetBody(v *DeleteFirewallTemplatesResponseBody) *DeleteFirewallTemplatesResponse {
	s.Body = v
	return s
}

type DeleteInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteInstanceKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceKeyPairRequest) SetClientToken(v string) *DeleteInstanceKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteInstanceKeyPairRequest) SetInstanceId(v string) *DeleteInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceKeyPairRequest) SetRegionId(v string) *DeleteInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

type DeleteInstanceKeyPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceKeyPairResponseBody) SetRequestId(v string) *DeleteInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceKeyPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceKeyPairResponse) SetHeaders(v map[string]*string) *DeleteInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceKeyPairResponse) SetStatusCode(v int32) *DeleteInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceKeyPairResponse) SetBody(v *DeleteInstanceKeyPairResponseBody) *DeleteInstanceKeyPairResponse {
	s.Body = v
	return s
}

type DeleteKeyPairsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The names of the SSH key pairs. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with http:// or https://. You can specify the names of a maximum of 50 SSH key pairs.
	//
	// This parameter is required.
	KeyPairNames []*string `json:"KeyPairNames,omitempty" xml:"KeyPairNames,omitempty" type:"Repeated"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) SetClientToken(v string) *DeleteKeyPairsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetKeyPairNames(v []*string) *DeleteKeyPairsRequest {
	s.KeyPairNames = v
	return s
}

func (s *DeleteKeyPairsRequest) SetRegionId(v string) *DeleteKeyPairsRequest {
	s.RegionId = &v
	return s
}

type DeleteKeyPairsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponseBody) SetRequestId(v string) *DeleteKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeyPairsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponse) SetHeaders(v map[string]*string) *DeleteKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyPairsResponse) SetStatusCode(v int32) *DeleteKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyPairsResponse) SetBody(v *DeleteKeyPairsResponseBody) *DeleteKeyPairsResponse {
	s.Body = v
	return s
}

type DeleteSnapshotRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) SetClientToken(v string) *DeleteSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnapshotRequest) SetRegionId(v string) *DeleteSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

type DeleteSnapshotResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponseBody) SetRequestId(v string) *DeleteSnapshotResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotResponse) SetHeaders(v map[string]*string) *DeleteSnapshotResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotResponse) SetStatusCode(v int32) *DeleteSnapshotResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotResponse) SetBody(v *DeleteSnapshotResponseBody) *DeleteSnapshotResponse {
	s.Body = v
	return s
}

type DeleteSnapshotsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple snapshot IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["s-bp16oazlsold4dks****", "s-bp16oazlsold4abc****"]
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
}

func (s DeleteSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotsRequest) SetClientToken(v string) *DeleteSnapshotsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteSnapshotsRequest) SetRegionId(v string) *DeleteSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnapshotsRequest) SetSnapshotIds(v string) *DeleteSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

type DeleteSnapshotsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C2DE174B-7196-5778-A00D-6EA2601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotsResponseBody) SetRequestId(v string) *DeleteSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSnapshotsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotsResponse) SetHeaders(v map[string]*string) *DeleteSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *DeleteSnapshotsResponse) SetStatusCode(v int32) *DeleteSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSnapshotsResponse) SetBody(v *DeleteSnapshotsResponseBody) *DeleteSnapshotsResponse {
	s.Body = v
	return s
}

type DescribeCloudAssistantAttributesRequest struct {
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the specified simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesRequest) SetInstanceIds(v []*string) *DescribeCloudAssistantAttributesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) SetPageNumber(v int32) *DescribeCloudAssistantAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) SetPageSize(v int32) *DescribeCloudAssistantAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantAttributesRequest) SetRegionId(v string) *DescribeCloudAssistantAttributesRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudAssistantAttributesShrinkRequest struct {
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the specified simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantAttributesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantAttributesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetInstanceIdsShrink(v string) *DescribeCloudAssistantAttributesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetPageNumber(v int32) *DescribeCloudAssistantAttributesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetPageSize(v int32) *DescribeCloudAssistantAttributesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantAttributesShrinkRequest) SetRegionId(v string) *DescribeCloudAssistantAttributesShrinkRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudAssistantAttributesResponseBody struct {
	// The Command Assistant information.
	CloudAssistant []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant `json:"CloudAssistant,omitempty" xml:"CloudAssistant,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudAssistantAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetCloudAssistant(v []*DescribeCloudAssistantAttributesResponseBodyCloudAssistant) *DescribeCloudAssistantAttributesResponseBody {
	s.CloudAssistant = v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetPageNumber(v int32) *DescribeCloudAssistantAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetPageSize(v int32) *DescribeCloudAssistantAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetRequestId(v string) *DescribeCloudAssistantAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBody) SetTotalCount(v int32) *DescribeCloudAssistantAttributesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCloudAssistantAttributesResponseBodyCloudAssistant struct {
	// The number of active tasks in Command Assistant.
	//
	// example:
	//
	// 0
	ActiveTaskCount *int64 `json:"ActiveTaskCount,omitempty" xml:"ActiveTaskCount,omitempty"`
	// Indicates whether Command Assistant is running. Valid values:
	//
	// true: Heartbeats are detected in the last 2 minutes.
	//
	// false: Heartbeats are not detected in the last 2 minutes.
	//
	// example:
	//
	// true
	CloudAssistantStatus *string `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty"`
	// The version number of the Command Assistant agent. Null is returned if the Command Assistant agent is not installed or is not running.
	//
	// example:
	//
	// 2.2.0.106
	CloudAssistantVersion *string `json:"CloudAssistantVersion,omitempty" xml:"CloudAssistantVersion,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 85dbe3e7cc7b49e1a3df4af3bfa4ebbf
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of completed tasks in Command Assistant.
	//
	// example:
	//
	// 4
	InvocationCount *int64 `json:"InvocationCount,omitempty" xml:"InvocationCount,omitempty"`
	// The time when the last heartbeat of Command Assistant was detected. The value is updated every minute on average. The interval can be 55, 60, or 65 seconds.
	//
	// example:
	//
	// 2021-03-15T09:00:00Z
	LastHeartbeatTime *string `json:"LastHeartbeatTime,omitempty" xml:"LastHeartbeatTime,omitempty"`
	// The time when commands were last run.
	//
	// example:
	//
	// 2021-03-15T08:00:00Z
	LastInvokedTime *string `json:"LastInvokedTime,omitempty" xml:"LastInvokedTime,omitempty"`
	// The OS type of the simple application server. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// 	- FreeBSD
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// Indicates whether Command Assistant supports session management. If Command Assistant does not support session management, the version of the Command Assistant agent is too earlier. We recommend that you update your Command Assistant agent to the latest version.
	//
	// To use the session management feature, you must make sure that the version of your Command Assistant agent meets one of the following requirements:
	//
	// If your simple application server runs Linux, the version of the Command Assistant agent on the server must be 2.2.3.189 or later. If your simple application server runs Windows, the version of the Command Assistant agent on the server must be 2.1.3.189 or later.
	//
	// example:
	//
	// true
	SupportSessionManager *bool `json:"SupportSessionManager,omitempty" xml:"SupportSessionManager,omitempty"`
}

func (s DescribeCloudAssistantAttributesResponseBodyCloudAssistant) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantAttributesResponseBodyCloudAssistant) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetActiveTaskCount(v int64) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.ActiveTaskCount = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetCloudAssistantStatus(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.CloudAssistantStatus = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetCloudAssistantVersion(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.CloudAssistantVersion = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetInstanceId(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetInvocationCount(v int64) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.InvocationCount = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetLastHeartbeatTime(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.LastHeartbeatTime = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetLastInvokedTime(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.LastInvokedTime = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetOSType(v string) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.OSType = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponseBodyCloudAssistant) SetSupportSessionManager(v bool) *DescribeCloudAssistantAttributesResponseBodyCloudAssistant {
	s.SupportSessionManager = &v
	return s
}

type DescribeCloudAssistantAttributesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudAssistantAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudAssistantAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantAttributesResponse) SetHeaders(v map[string]*string) *DescribeCloudAssistantAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudAssistantAttributesResponse) SetStatusCode(v int32) *DescribeCloudAssistantAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudAssistantAttributesResponse) SetBody(v *DescribeCloudAssistantAttributesResponseBody) *DescribeCloudAssistantAttributesResponse {
	s.Body = v
	return s
}

type DescribeCloudAssistantStatusRequest struct {
	// The IDs of the simple application servers.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusRequest) SetInstanceIds(v []*string) *DescribeCloudAssistantStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageNumber(v int32) *DescribeCloudAssistantStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageSize(v int32) *DescribeCloudAssistantStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetRegionId(v string) *DescribeCloudAssistantStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudAssistantStatusShrinkRequest struct {
	// The IDs of the simple application servers.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudAssistantStatusShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetInstanceIdsShrink(v string) *DescribeCloudAssistantStatusShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetPageNumber(v int32) *DescribeCloudAssistantStatusShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetPageSize(v int32) *DescribeCloudAssistantStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusShrinkRequest) SetRegionId(v string) *DescribeCloudAssistantStatusShrinkRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudAssistantStatusResponseBody struct {
	// Indicates whether the Cloud Assistant client is installed on the server.
	CloudAssistantStatus []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus `json:"CloudAssistantStatus,omitempty" xml:"CloudAssistantStatus,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBody) SetCloudAssistantStatus(v []*DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) *DescribeCloudAssistantStatusResponseBody {
	s.CloudAssistantStatus = v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageNumber(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetPageSize(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetRequestId(v string) *DescribeCloudAssistantStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBody) SetTotalCount(v int32) *DescribeCloudAssistantStatusResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the Cloud Assistant client is installed on the server.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) SetInstanceId(v string) *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus) SetStatus(v bool) *DescribeCloudAssistantStatusResponseBodyCloudAssistantStatus {
	s.Status = &v
	return s
}

type DescribeCloudAssistantStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudAssistantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudAssistantStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudAssistantStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusResponse) SetHeaders(v map[string]*string) *DescribeCloudAssistantStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) SetStatusCode(v int32) *DescribeCloudAssistantStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudAssistantStatusResponse) SetBody(v *DescribeCloudAssistantStatusResponseBody) *DescribeCloudAssistantStatusResponse {
	s.Body = v
	return s
}

type DescribeCloudMonitorAgentStatusesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc******","2ad1ae67295445f598017499dc******"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesRequest) SetClientToken(v string) *DescribeCloudMonitorAgentStatusesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesRequest) SetInstanceIds(v string) *DescribeCloudMonitorAgentStatusesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesRequest) SetRegionId(v string) *DescribeCloudMonitorAgentStatusesRequest {
	s.RegionId = &v
	return s
}

type DescribeCloudMonitorAgentStatusesResponseBody struct {
	// Indicates whether the Cloud Monitor agent was automatically installed on the simple application server.
	InstanceStatusList []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E1FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) SetInstanceStatusList(v []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) *DescribeCloudMonitorAgentStatusesResponseBody {
	s.InstanceStatusList = v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBody) SetRequestId(v string) *DescribeCloudMonitorAgentStatusesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList struct {
	// Indicates whether the Cloud Monitor agent was automatically installed on the simple application server. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// c854dc6f07e74953830bb5808d0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The running status of the Cloud Monitoring plug-in. Possible values are:
	//
	// - running: Cloud Monitoring plug-in running.
	//
	// - stopped: Cloud Monitoring plug-in stopped.
	//
	// - installing: Cloud Monitoring plug-in installing.
	//
	// - install_faild: Cloud Monitoring plug-in installation failed.
	//
	// - abnormal: Cloud Monitoring plug-in installation abnormal.
	//
	// - not_installed: Cloud Monitoring plug-in not installed.
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) SetAutoInstall(v bool) *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	s.AutoInstall = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) SetInstanceId(v string) *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList) SetStatus(v string) *DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList {
	s.Status = &v
	return s
}

type DescribeCloudMonitorAgentStatusesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudMonitorAgentStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudMonitorAgentStatusesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCloudMonitorAgentStatusesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudMonitorAgentStatusesResponse) SetHeaders(v map[string]*string) *DescribeCloudMonitorAgentStatusesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponse) SetStatusCode(v int32) *DescribeCloudMonitorAgentStatusesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudMonitorAgentStatusesResponse) SetBody(v *DescribeCloudMonitorAgentStatusesResponseBody) *DescribeCloudMonitorAgentStatusesResponse {
	s.Body = v
	return s
}

type DescribeCommandInvocationsRequest struct {
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name. If both CommandName and InstanceId are specified, CommandName does not take effect.
	//
	// example:
	//
	// testName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The command type. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances
	//
	// 	- RunShellScript: shell command, applicable to Linux instances
	//
	// example:
	//
	// RunPowerShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The overall execution state of the command. The value of this parameter depends on the execution status of the command on all the involved instances. Valid values:
	//
	// 	- Pending: The command is being verified or sent. When the execution state on at least one instance is Pending, the overall execution state is Pending.
	//
	// 	- Running: The command is being run on the instances. When the execution state on at least one instance is Running, the overall execution state is Running.
	//
	// 	- Success: When the execution state on at least one instance is Success and the execution state on other instances is Stopped or Success, the overall execution state is Success.
	//
	//     	- Command that is set to run immediately: The command execution is complete, and the exit code is 0.
	//
	// 	- Failed: When the execution state on all instances is Stopped or Failed, the overall execution state is Failed. When the execution state on an instance is one of the following values, Failed is returned as the overall execution state:
	//
	//     	- Invalid: The command is invalid.
	//
	//     	- Aborted: The command fails to be sent.
	//
	//     	- Failed: The command execution is complete, and the exit code is not 0.
	//
	//     	- Timeout: The command execution times out.
	//
	//     	- Error: An error occurs when the command is being run.
	//
	// 	- Stopping: The command task is being stopped. When the execution state on at least one instance is Stopping, the overall execution state is Stopping.
	//
	// 	- Stopped: The command task is stopped. When the execution state on all instances is Stopped, the overall execution state is Stopped. When the execution state on an instance is one of the following values, Stopped is returned as the overall execution state:
	//
	//     	- Cancelled: The command task is canceled.
	//
	//     	- Terminated: The command task is terminated.
	//
	// 	- PartialFailed: The command execution succeeds on some instances and fails on other instances. When the execution state on some instances is Success and the execution state on other instances is Failed or Stopped, the overall execution state is PartialFailed.
	//
	// >  The value of the `InvokeStatus` response parameter is similar to the value of InvocationStatus. We recommend that you ignore InvokeStatus and check the value of InvocationStatus.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The execution ID of the command.
	//
	// example:
	//
	// t-hz02p9545t6****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCommandInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsRequest) SetCommandId(v string) *DescribeCommandInvocationsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetCommandName(v string) *DescribeCommandInvocationsRequest {
	s.CommandName = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetCommandType(v string) *DescribeCommandInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetInstanceId(v string) *DescribeCommandInvocationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetInvocationStatus(v string) *DescribeCommandInvocationsRequest {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetInvokeId(v string) *DescribeCommandInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetPageNumber(v string) *DescribeCommandInvocationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetPageSize(v string) *DescribeCommandInvocationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetRegionId(v string) *DescribeCommandInvocationsRequest {
	s.RegionId = &v
	return s
}

type DescribeCommandInvocationsResponseBody struct {
	// The command executions.
	CommandInvocations []*DescribeCommandInvocationsResponseBodyCommandInvocations `json:"CommandInvocations,omitempty" xml:"CommandInvocations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommandInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponseBody) SetCommandInvocations(v []*DescribeCommandInvocationsResponseBodyCommandInvocations) *DescribeCommandInvocationsResponseBody {
	s.CommandInvocations = v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetPageNumber(v int32) *DescribeCommandInvocationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetPageSize(v int32) *DescribeCommandInvocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetRequestId(v string) *DescribeCommandInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetTotalCount(v int32) *DescribeCommandInvocationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCommandInvocationsResponseBodyCommandInvocations struct {
	// The content of the command.
	//
	// example:
	//
	// echo 123
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The description of the command.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The command ID.
	//
	// example:
	//
	// c-hy0338xh28r****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name.
	//
	// example:
	//
	// testCommandName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The command type.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The time when the command was created.
	//
	// example:
	//
	// 2023-04-27T10:11:58
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The overall execution state of the command. Valid values:
	//
	// 	- Pending: The command is being verified or sent.
	//
	// 	- Invalid: The specified command type or parameter is invalid.
	//
	// 	- Aborted: The command failed to be sent to the instances. To send a command to an instance, make sure that the instance is in the Running state and the command is sent to the instance within 1 minute.
	//
	// 	- Running: The command is being run on the instances.
	//
	// 	- Success: The command execution is complete, and the exit code is 0.
	//
	// 	- Failed: The command execution is complete, and the exit code is not 0.
	//
	// 	- Error: The command execution cannot proceed due to an exception.
	//
	// 	- Timeout: The command execution timed out.
	//
	// 	- Cancelled: The command execution is canceled, and the command is not started.
	//
	// 	- Stopping: The command in the Running state is being stopped.
	//
	// 	- Terminated: The command is terminated when it is being run.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The execution ID of the command.
	//
	// example:
	//
	// t-hz0373jyzxt****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The instances on which the command is run.
	InvokeInstances []*DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Repeated"`
	// The custom parameters in the command. If no custom parameter exists in the command, the default value is {}.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The timeout period. Unit: seconds.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username that is used to run the command.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The working directory of the command.
	//
	// example:
	//
	// c:\\wwwroot
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocations) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandContent(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandDescription(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandDescription = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandId(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandName(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandName = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandType(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCreationTime(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetInvocationStatus(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetInvokeId(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetInvokeInstances(v []*DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.InvokeInstances = v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetParameters(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.Parameters = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetTimeout(v int64) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.Timeout = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetUsername(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.Username = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetWorkingDir(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.WorkingDir = &v
	return s
}

type DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances struct {
	// The error code returned if the command failed to be sent or run.
	//
	// 	- null: The command is run as expected.
	//
	// 	- InstanceNotExists: The specified instance does not exist or is released.
	//
	// 	- InstanceReleased: The instance is released when the command is being run.
	//
	// 	- InstanceNotRunning: The instance is not in the Running state when the command is being run.
	//
	// 	- CommandNotApplicable: The command is not applicable to the specified instance.
	//
	// 	- AccountNotExists: The specified account does not exist.
	//
	// 	- DirectoryNotExists: The specified directory does not exist.
	//
	// 	- BadCronExpression: The specified CRON expression for the execution schedule is invalid.
	//
	// 	- ClientNotRunning: Cloud Assistant Agent is not running.
	//
	// 	- ClientNotResponse: Cloud Assistant Agent does not respond to your request.
	//
	// 	- ClientIsUpgrading: Cloud Assistant Agent is being updated.
	//
	// 	- ClientNeedUpgrade: Cloud Assistant Agent needs to be updated.
	//
	// 	- DeliveryTimeout: The request to send the command timed out.
	//
	// 	- ExecutionTimeout: The command execution timed out.
	//
	// 	- ExecutionException: An exception occurred when the command was being run.
	//
	// 	- ExecutionInterrupted: The command execution is interrupted.
	//
	// 	- ExitCodeNonzero: The command execution is complete, and the exit code is not 0.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the command failed to be sent or run. Valid values:
	//
	// 	- null: The command is run as expected.
	//
	// 	- the specified instance does not exists: The specified instance does not exist or is released.
	//
	// 	- the instance has released when create task: The instance is released when the command is being run.
	//
	// 	- the instance is not running when create task: The instance is not in the Running state when the command is being run.
	//
	// 	- the command is not applicable: The command is not applicable to the specified instance.
	//
	// 	- the specified account does not exists: The specified account does not exist.
	//
	// 	- the specified directory does not exists: The specified directory does not exist.
	//
	// 	- the cron job expression is invalid: The specified CRON expression for the execution schedule is invalid.
	//
	// 	- the aliyun service is not running on the instance: Cloud Assistant Agent is not running.
	//
	// 	- the aliyun service in the instance does not response: Cloud Assistant Agent does not respond to your request.
	//
	// 	- the aliyun service in the instance is upgrading now: Cloud Assistant Agent is being updated.
	//
	// 	- the aliyun service in the instance need upgrade: Cloud Assistant Agent needs to be updated.
	//
	// 	- the command delivery has been timeout: The request to send the command timed out.
	//
	// 	- the command execution has been timeout: The command execution timed out.
	//
	// 	- the command execution got an exception: An exception occurred when the command was being run.
	//
	// 	- the command execution has been interrupted: The command execution is interrupted.
	//
	// 	- the command execution exit code is not zero: The command execution is complete, and the exit code is not 0.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command.
	//
	// 	- For Linux instances, the exit code is the exit code of the shell command.
	//
	// 	- For Windows instances, the exit code is the exit code of the batch or PowerShell command.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The end of the time range during which the command is run on the instance.
	//
	// example:
	//
	// 2023-04-03T02:42:29Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 2445f4aecdac4b71ba2c7e3a7ccf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution state of the command on a single instance. Valid values:
	//
	// 	- Pending: The command is being verified or sent.
	//
	// 	- Invalid: The specified command type or parameter is invalid.
	//
	// 	- Aborted: The command failed to be sent to the instance. To send a command to an instance, make sure that the instance is in the Running state and the command is sent to the instance within 1 minute.
	//
	// 	- Running: The command is being run on the instance.
	//
	// 	- Success:
	//
	//     	- Command that is set to run only once: The command execution is complete, and the exit code is 0.
	//
	//     	- Command that is set to run on a schedule: The previous command execution is complete, the exit code is 0, and the specified cycle ends.
	//
	// 	- Failed:
	//
	//     	- Command that is set to run only once: The command execution is complete, and the exit code is not 0.
	//
	//     	- Command that is set to run on a schedule: The previous command execution is complete, the exit code is not 0, and the specified cycle is about to end.
	//
	// 	- Error: The command execution cannot proceed due to an exception.
	//
	// 	- Timeout: The command execution timed out.
	//
	// 	- Cancelled: The command execution is canceled, and the command is not started.
	//
	// 	- Stopping: The command task is being stopped.
	//
	// 	- Terminated: The command is terminated when it is being run.
	//
	// example:
	//
	// Running
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The command output.
	//
	// example:
	//
	// OutputMsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The beginning of the time range during which the command is run on the instance.
	//
	// example:
	//
	// 2023-05-09T03:32:24Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetErrorCode(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetErrorInfo(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetExitCode(v int64) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.ExitCode = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetFinishTime(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.FinishTime = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetInstanceId(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetInvocationStatus(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetOutput(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.Output = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetStartTime(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.StartTime = &v
	return s
}

type DescribeCommandInvocationsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommandInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommandInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponse) SetHeaders(v map[string]*string) *DescribeCommandInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommandInvocationsResponse) SetStatusCode(v int32) *DescribeCommandInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommandInvocationsResponse) SetBody(v *DescribeCommandInvocationsResponseBody) *DescribeCommandInvocationsResponse {
	s.Body = v
	return s
}

type DescribeCommandsRequest struct {
	// The command ID.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name. Fuzzy match is not supported.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// Pages start from 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The provider of the common command. Take note of the following items:
	//
	// 	- If you set this parameter to `AlibabaCloud`, all the common commands provided by Alibaba Cloud are queried.
	//
	// 	- If you set this parameter to `User`, all the custom commands created by you are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// AlibabaCloud
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are bound to the command.
	Tag []*DescribeCommandsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The language type of the command. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances
	//
	// 	- RunShellScript: shell command, applicable to Linux instances
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommandsRequest) SetCommandId(v string) *DescribeCommandsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandsRequest) SetName(v string) *DescribeCommandsRequest {
	s.Name = &v
	return s
}

func (s *DescribeCommandsRequest) SetPageNumber(v string) *DescribeCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandsRequest) SetPageSize(v string) *DescribeCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandsRequest) SetProvider(v string) *DescribeCommandsRequest {
	s.Provider = &v
	return s
}

func (s *DescribeCommandsRequest) SetRegionId(v string) *DescribeCommandsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommandsRequest) SetResourceGroupId(v string) *DescribeCommandsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommandsRequest) SetTag(v []*DescribeCommandsRequestTag) *DescribeCommandsRequest {
	s.Tag = v
	return s
}

func (s *DescribeCommandsRequest) SetType(v string) *DescribeCommandsRequest {
	s.Type = &v
	return s
}

type DescribeCommandsRequestTag struct {
	// The tag key of the command. A tag key can be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command. A tag value can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommandsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeCommandsRequestTag) SetKey(v string) *DescribeCommandsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeCommandsRequestTag) SetValue(v string) *DescribeCommandsRequestTag {
	s.Value = &v
	return s
}

type DescribeCommandsResponseBody struct {
	// The queried commands.
	Commands []*DescribeCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of commands.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBody) SetCommands(v []*DescribeCommandsResponseBodyCommands) *DescribeCommandsResponseBody {
	s.Commands = v
	return s
}

func (s *DescribeCommandsResponseBody) SetPageNumber(v int32) *DescribeCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetPageSize(v int32) *DescribeCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetRequestId(v string) *DescribeCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommandsResponseBody) SetTotalCount(v int32) *DescribeCommandsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeCommandsResponseBodyCommands struct {
	// The content of the command.
	//
	// example:
	//
	// cat /etc/ssh/sshd_config
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command ID.
	//
	// example:
	//
	// c-gov1k1tqwi9****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The time when the command was created.
	//
	// example:
	//
	// 2023-01-05T06:38:53Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the command.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the custom parameter feature is enabled for the command.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The name of the command.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Details of the custom parameter.
	ParameterDefinitions []*DescribeCommandsResponseBodyCommandsParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Repeated"`
	// The custom parameter names that are parsed from the command content specified when the command was created. The custom parameter names are returned in the list format. If the custom parameter feature is disabled, an empty list is returned.
	ParameterNames []*string `json:"ParameterNames,omitempty" xml:"ParameterNames,omitempty" type:"Repeated"`
	// The provider of the command.
	//
	// example:
	//
	// User
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are bound to the command.
	Tags []*DescribeCommandsResponseBodyCommandsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period of the command.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the command.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The execution path of the command.
	//
	// example:
	//
	// /home
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeCommandsResponseBodyCommands) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommands) SetCommandContent(v string) *DescribeCommandsResponseBodyCommands {
	s.CommandContent = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetCommandId(v string) *DescribeCommandsResponseBodyCommands {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetCreationTime(v string) *DescribeCommandsResponseBodyCommands {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetDescription(v string) *DescribeCommandsResponseBodyCommands {
	s.Description = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetEnableParameter(v bool) *DescribeCommandsResponseBodyCommands {
	s.EnableParameter = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetName(v string) *DescribeCommandsResponseBodyCommands {
	s.Name = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetParameterDefinitions(v []*DescribeCommandsResponseBodyCommandsParameterDefinitions) *DescribeCommandsResponseBodyCommands {
	s.ParameterDefinitions = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetParameterNames(v []*string) *DescribeCommandsResponseBodyCommands {
	s.ParameterNames = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetProvider(v string) *DescribeCommandsResponseBodyCommands {
	s.Provider = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetResourceGroupId(v string) *DescribeCommandsResponseBodyCommands {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetTags(v []*DescribeCommandsResponseBodyCommandsTags) *DescribeCommandsResponseBodyCommands {
	s.Tags = v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetTimeout(v int64) *DescribeCommandsResponseBodyCommands {
	s.Timeout = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetType(v string) *DescribeCommandsResponseBodyCommands {
	s.Type = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommands) SetWorkingDir(v string) *DescribeCommandsResponseBodyCommands {
	s.WorkingDir = &v
	return s
}

type DescribeCommandsResponseBodyCommandsParameterDefinitions struct {
	// The default value of the custom parameter.
	//
	// example:
	//
	// https://aliyun-client-assist.oss-accelerate.aliyuncs.com/linux/aliyun_assist_latest.rpm
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The description of the custom parameter.
	//
	// example:
	//
	// Command Assistant Agent Installation Package Path
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom parameter.
	//
	// example:
	//
	// DownloadUrl
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The valid values of the custom parameter.
	PossibleValues []*string `json:"PossibleValues,omitempty" xml:"PossibleValues,omitempty" type:"Repeated"`
	// Indicates whether the custom parameter is required. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsParameterDefinitions) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsParameterDefinitions) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetDefaultValue(v string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.DefaultValue = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetDescription(v string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.Description = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetParameterName(v string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.ParameterName = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetPossibleValues(v []*string) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.PossibleValues = v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsParameterDefinitions) SetRequired(v bool) *DescribeCommandsResponseBodyCommandsParameterDefinitions {
	s.Required = &v
	return s
}

type DescribeCommandsResponseBodyCommandsTags struct {
	// The tag key of the command.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCommandsResponseBodyCommandsTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsResponseBodyCommandsTags) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponseBodyCommandsTags) SetKey(v string) *DescribeCommandsResponseBodyCommandsTags {
	s.Key = &v
	return s
}

func (s *DescribeCommandsResponseBodyCommandsTags) SetValue(v string) *DescribeCommandsResponseBodyCommandsTags {
	s.Value = &v
	return s
}

type DescribeCommandsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCommandsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommandsResponse) SetHeaders(v map[string]*string) *DescribeCommandsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommandsResponse) SetStatusCode(v int32) *DescribeCommandsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommandsResponse) SetBody(v *DescribeCommandsResponseBody) *DescribeCommandsResponse {
	s.Body = v
	return s
}

type DescribeDatabaseErrorLogsRequest struct {
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-08T04:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Simple Database Service instance.
	//
	// You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseErrorLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseErrorLogsRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetEndTime(v string) *DescribeDatabaseErrorLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetPageNumber(v int32) *DescribeDatabaseErrorLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetPageSize(v int32) *DescribeDatabaseErrorLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetRegionId(v string) *DescribeDatabaseErrorLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsRequest) SetStartTime(v string) *DescribeDatabaseErrorLogsRequest {
	s.StartTime = &v
	return s
}

type DescribeDatabaseErrorLogsResponseBody struct {
	// The time when the error log entry was generated. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	ErrorLogs []*DescribeDatabaseErrorLogsResponseBodyErrorLogs `json:"ErrorLogs,omitempty" xml:"ErrorLogs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetErrorLogs(v []*DescribeDatabaseErrorLogsResponseBodyErrorLogs) *DescribeDatabaseErrorLogsResponseBody {
	s.ErrorLogs = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetPageNumber(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetPageSize(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetRequestId(v string) *DescribeDatabaseErrorLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBody) SetTotalCount(v int32) *DescribeDatabaseErrorLogsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabaseErrorLogsResponseBodyErrorLogs struct {
	// The time when the resource was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-08T12:11:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// spid52 DBCC TRACEON 3499, server process ID (SPID) 52. This is an informational message only; no user action is required
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponseBodyErrorLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponseBodyErrorLogs) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) SetCreateTime(v string) *DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	s.CreateTime = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponseBodyErrorLogs) SetErrorInfo(v string) *DescribeDatabaseErrorLogsResponseBodyErrorLogs {
	s.ErrorInfo = &v
	return s
}

type DescribeDatabaseErrorLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseErrorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseErrorLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseErrorLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseErrorLogsResponse) SetHeaders(v map[string]*string) *DescribeDatabaseErrorLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) SetStatusCode(v int32) *DescribeDatabaseErrorLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseErrorLogsResponse) SetBody(v *DescribeDatabaseErrorLogsResponseBody) *DescribeDatabaseErrorLogsResponse {
	s.Body = v
	return s
}

type DescribeDatabaseInstanceMetricDataRequest struct {
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the metric. Valid values:
	//
	// 	- MySQL_MemCpuUsage: The CPU utilization and memory usage of the instance within the entire operating system.
	//
	// 	- MySQL_DetailedSpaceUsage: The total space usage, data space, log space, temporary space, and system space of the instance.
	//
	// 	- MySQL_Sessions : The total number of active connections.
	//
	// 	- MySQL_IOPS: The IOPS of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL_MemCpuUsage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-06T04:04Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetEndTime(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetMetricName(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetRegionId(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataRequest) SetStartTime(v string) *DescribeDatabaseInstanceMetricDataRequest {
	s.StartTime = &v
	return s
}

type DescribeDatabaseInstanceMetricDataResponseBody struct {
	// The data format. Valid values:
	//
	// 	- cpuusage\\&memusage
	//
	// 	- active_session\\&total_session
	//
	// 	- ins_size\\&data_size\\&log_size\\&tmp_size\\&other_size
	//
	// 	- io
	//
	// example:
	//
	// cpuusage&memusage
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	// The monitoring data.
	//
	// example:
	//
	// [  {     \\"date"\\: " 2022-09-06T04:04:00Z",\\"value\\":\\"0.77&3.69\\"  } ]
	MetricData *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	// The name of the metric. Valid values:
	//
	// 	- MySQL_MemCpuUsage: The CPU utilization and memory usage of the instance within the entire operating system.
	//
	// 	- MySQL_DetailedSpaceUsage: The total space usage, data space, log space, temporary space, and system space of the instance.
	//
	// 	- MySQL_Sessions : The total number of active connections.
	//
	// 	- MySQL_IOPS: The IOPS of the instance.
	//
	// example:
	//
	// MySQL_MemCpuUsage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unit of the monitoring metric.
	//
	// 	- %
	//
	// 	- int
	//
	// 	- MB
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetDataFormat(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.DataFormat = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetMetricData(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.MetricData = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetMetricName(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.MetricName = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetRequestId(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetUnit(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.Unit = &v
	return s
}

type DescribeDatabaseInstanceMetricDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseInstanceMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstanceMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetStatusCode(v int32) *DescribeDatabaseInstanceMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponse) SetBody(v *DescribeDatabaseInstanceMetricDataResponseBody) *DescribeDatabaseInstanceMetricDataResponse {
	s.Body = v
	return s
}

type DescribeDatabaseInstanceParametersRequest struct {
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDatabaseInstanceParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseInstanceParametersRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersRequest) SetRegionId(v string) *DescribeDatabaseInstanceParametersRequest {
	s.RegionId = &v
	return s
}

type DescribeDatabaseInstanceParametersResponseBody struct {
	// The range of ParameterValue.
	//
	// > The value of CheckingCode varies based on the value of ParameterName.
	ConfigParameters []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters `json:"ConfigParameters,omitempty" xml:"ConfigParameters,omitempty" type:"Repeated"`
	// The database engine that the instance runs. The value must be MySQL.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine. Valid values:
	//
	// 	- 5.7: MySQL 5.7.
	//
	// 	- 8.0: MySQL 8.0.
	//
	// example:
	//
	// 5.5
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The range of ParameterValue.
	//
	// > The value of CheckingCode varies based on the value of ParameterName.
	RunningParameters []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeDatabaseInstanceParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetConfigParameters(v []*DescribeDatabaseInstanceParametersResponseBodyConfigParameters) *DescribeDatabaseInstanceParametersResponseBody {
	s.ConfigParameters = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetEngine(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetEngineVersion(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetRequestId(v string) *DescribeDatabaseInstanceParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBody) SetRunningParameters(v []*DescribeDatabaseInstanceParametersResponseBodyRunningParameters) *DescribeDatabaseInstanceParametersResponseBody {
	s.RunningParameters = v
	return s
}

type DescribeDatabaseInstanceParametersResponseBodyConfigParameters struct {
	// The check code that indicates the valid values of the parameter.
	//
	// example:
	//
	// [1-65535]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Does it support modifying parameter values. Possible values:
	//
	// - true:Support modifying parameter values.
	//
	// - false:Modifying parameter values is not supported.
	//
	// example:
	//
	// true
	ForceModify *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	// Specifies whether to forcibly restart the instance after parameters are modified. Valid values:
	//
	// 	- true: forcibly restarts the instance. If a new parameter value takes effect only after the instance restarts, you must set this parameter to true. Otherwise, the new parameter value cannot take effect.
	//
	// 	- false: does not forcibly restart the instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// Auto-increment columns are incremented by this
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// auto_increment_increment
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponseBodyConfigParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBodyConfigParameters) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetCheckingCode(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetForceModify(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ForceModify = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetForceRestart(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterDescription(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterName(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyConfigParameters) SetParameterValue(v string) *DescribeDatabaseInstanceParametersResponseBodyConfigParameters {
	s.ParameterValue = &v
	return s
}

type DescribeDatabaseInstanceParametersResponseBodyRunningParameters struct {
	// The check code that indicates the valid values of the parameter.
	//
	// example:
	//
	// [ON|OFF]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// Does it support modifying parameter values. Possible values:
	//
	// - true:Support modifying parameter values.
	//
	// - false:Modifying parameter values is not supported.
	//
	// example:
	//
	// true
	ForceModify *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	// Specifies whether to forcibly restart the instance after parameters are modified. Valid values:
	//
	// 	- true: forcibly restarts the instance. If a new parameter value takes effect only after the instance restarts, you must set this parameter to true. Otherwise, the new parameter value cannot take effect.
	//
	// 	- false: does not forcibly restart the instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceRestart *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// When this variable has a value of 1 (the default), the server automatically grants the EXECUTE and ALTER ROUTINE privileges to the creator of a stored routine, if the user cannot already execute and alter or drop the routine. (The ALTER ROUTINE privilege is required to drop the routine.) The server also automatically drops those privileges from the creator when the routine is dropped. If automatic_sp_privileges is 0, the server does not automatically add or drop these privileges.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// autocommit
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// ON
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponseBodyRunningParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponseBodyRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetCheckingCode(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetForceModify(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ForceModify = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetForceRestart(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterDescription(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterName(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponseBodyRunningParameters) SetParameterValue(v string) *DescribeDatabaseInstanceParametersResponseBodyRunningParameters {
	s.ParameterValue = &v
	return s
}

type DescribeDatabaseInstanceParametersResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseInstanceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseInstanceParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstanceParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceParametersResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstanceParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) SetStatusCode(v int32) *DescribeDatabaseInstanceParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstanceParametersResponse) SetBody(v *DescribeDatabaseInstanceParametersResponseBody) *DescribeDatabaseInstanceParametersResponse {
	s.Body = v
	return s
}

type DescribeDatabaseInstancesRequest struct {
	// The IDs of the Simple Database Service instances. The value can be a JSON array that consists of up to 100 Simple Database Service instance IDs. Separate multiple instance IDs with commas (,).
	//
	// example:
	//
	// ["swasdb-xxx******","swasdb-yyy******"]
	DatabaseInstanceIds *string `json:"DatabaseInstanceIds,omitempty" xml:"DatabaseInstanceIds,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Simple Database Service instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDatabaseInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesRequest) SetDatabaseInstanceIds(v string) *DescribeDatabaseInstancesRequest {
	s.DatabaseInstanceIds = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetPageNumber(v int32) *DescribeDatabaseInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetPageSize(v int32) *DescribeDatabaseInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseInstancesRequest) SetRegionId(v string) *DescribeDatabaseInstancesRequest {
	s.RegionId = &v
	return s
}

type DescribeDatabaseInstancesResponseBody struct {
	// The information about the Simple Database Service instances.
	DatabaseInstances []*DescribeDatabaseInstancesResponseBodyDatabaseInstances `json:"DatabaseInstances,omitempty" xml:"DatabaseInstances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponseBody) SetDatabaseInstances(v []*DescribeDatabaseInstancesResponseBodyDatabaseInstances) *DescribeDatabaseInstancesResponseBody {
	s.DatabaseInstances = v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetPageNumber(v int32) *DescribeDatabaseInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetPageSize(v int32) *DescribeDatabaseInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetRequestId(v string) *DescribeDatabaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBody) SetTotalCount(v int32) *DescribeDatabaseInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabaseInstancesResponseBodyDatabaseInstances struct {
	// The business status of the instance. Valid values:
	//
	// 	- normal
	//
	// 	- expired
	//
	// 	- overdue
	//
	// example:
	//
	// normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the Simple Database Service instance. Set the value to PrePaid. This value indicates the subscription billing method.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 1
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the instance was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-01T02:39:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The plan edition ID of the Simple Database Service instance. Valid values:
	//
	// 	- swas.db.c1m1s25: CNY 35/month.
	//
	// 	- swas.db.c1m2s80: CNY 100/month.
	//
	// 	- swas.db.c2m4s120: CNY 200/month.
	//
	// 	- swas.db.c2m8s240: CNY 300/month.
	//
	// example:
	//
	// swas.db.c1m1s25
	DatabaseInstanceEdition *string `json:"DatabaseInstanceEdition,omitempty" xml:"DatabaseInstanceEdition,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The name of the Simple Database Service instance.
	DatabaseInstanceName *string `json:"DatabaseInstanceName,omitempty" xml:"DatabaseInstanceName,omitempty"`
	// The status of the Simple Database Service instance. Valid values:
	//
	// 	- Pending: The instance is being created.
	//
	// 	- Restarting: The instance is being restarted.
	//
	// 	- Running: The instance is running.
	//
	// 	- Stopping: The instance is being stopped.
	//
	// 	- Stopped: The instance is stopped.
	//
	// 	- UPGRADING: The instance is being upgraded.
	//
	// 	- DISABLED: The instance is disabled.
	//
	// example:
	//
	// Running
	DatabaseInstanceStatus *string `json:"DatabaseInstanceStatus,omitempty" xml:"DatabaseInstanceStatus,omitempty"`
	// The database engine version of the instance. Valid values:
	//
	// 	- 5.7: MySQL 5.7.
	//
	// 	- 8.0: MySQL 8.0.
	//
	// example:
	//
	// 5.7
	DatabaseVersion *string `json:"DatabaseVersion,omitempty" xml:"DatabaseVersion,omitempty"`
	// The time when the instance expires. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// example:
	//
	// 2022-10-01T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The memory size of the instance. Unit: GB.
	//
	// example:
	//
	// 1
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The private endpoint.
	//
	// example:
	//
	// rm-bp1d39opj7906****.mysql.rds.aliyuncs.com
	PrivateConnection *string `json:"PrivateConnection,omitempty" xml:"PrivateConnection,omitempty"`
	// The public endpoint.
	//
	// >  This parameter is displayed only after you apply for a public endpoint for the instance and a public endpoint is assigned to the instance. You can call [AllocatePublicConnection](https://help.aliyun.com/document_detail/451413.html) to apply for a public endpoint for the instance.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****.mysql.rds.aliyuncs.com
	PublicConnection *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	// The region ID of the Simple Database Service instances.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The size of the enhanced SSD (ESSD). Unit: GB.
	//
	// example:
	//
	// 25
	Storage *int32 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// The name of the super administrator account of the Simple Database Service instance.
	//
	// example:
	//
	// administrator
	SuperAccountName *string `json:"SuperAccountName,omitempty" xml:"SuperAccountName,omitempty"`
}

func (s DescribeDatabaseInstancesResponseBodyDatabaseInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesResponseBodyDatabaseInstances) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetBusinessStatus(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetChargeType(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.ChargeType = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetCpu(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetCreationTime(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceEdition(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceEdition = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceId(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceName(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceName = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseInstanceStatus(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseInstanceStatus = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetDatabaseVersion(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.DatabaseVersion = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetExpiredTime(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetMemory(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Memory = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetPrivateConnection(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.PrivateConnection = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetPublicConnection(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.PublicConnection = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetRegionId(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetStorage(v int32) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.Storage = &v
	return s
}

func (s *DescribeDatabaseInstancesResponseBodyDatabaseInstances) SetSuperAccountName(v string) *DescribeDatabaseInstancesResponseBodyDatabaseInstances {
	s.SuperAccountName = &v
	return s
}

type DescribeDatabaseInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstancesResponse) SetHeaders(v map[string]*string) *DescribeDatabaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseInstancesResponse) SetStatusCode(v int32) *DescribeDatabaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseInstancesResponse) SetBody(v *DescribeDatabaseInstancesResponseBody) *DescribeDatabaseInstancesResponse {
	s.Body = v
	return s
}

type DescribeDatabaseSlowLogRecordsRequest struct {
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The interval between the start time and the end time must be less than 7 days.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-08T04:04:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30 to 100.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetDatabaseInstanceId(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetEndTime(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetRegionId(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetStartTime(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

type DescribeDatabaseSlowLogRecordsResponseBody struct {
	// The database engine that the instance runs.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30 to 100.
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of logical reads.
	//
	// example:
	//
	// 200
	PhysicalIORead *int64 `json:"PhysicalIORead,omitempty" xml:"PhysicalIORead,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The slow query logs returned.
	SlowLogs []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs `json:"SlowLogs,omitempty" xml:"SlowLogs,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetEngine(v string) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPageNumber(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPageSize(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetPhysicalIORead(v int64) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.PhysicalIORead = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetRequestId(v string) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetSlowLogs(v []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.SlowLogs = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBody) SetTotalCount(v int32) *DescribeDatabaseSlowLogRecordsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabaseSlowLogRecordsResponseBodySlowLogs struct {
	// The name of the database.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The time when the execution of the SQL statement started. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// example:
	//
	// 2022-09-08T01:40:44Z
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	// The name and IP address of the client that is connected to the database.
	//
	// example:
	//
	// xxx[xxx] @ [1xx.xxx.xxx.xx]
	HostAddress *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	// The lock duration of the SQL statement. Unit: seconds.
	//
	// example:
	//
	// 0
	LockTimes *int64 `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	// The number of rows parsed by the SQL statement.
	//
	// example:
	//
	// 1
	ParseRowCounts *int64 `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	// The execution duration of the slow query. Unit: millisecond.
	//
	// example:
	//
	// 2001
	QueryTimeMS *int64 `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	// The execution duration of the slow query. Unit: seconds.
	//
	// example:
	//
	// 2
	QueryTimes *int64 `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	// The number of rows returned by the SQL statement.
	//
	// example:
	//
	// 1
	ReturnRowCounts *int64 `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	// The details of the SQL statement.
	//
	// example:
	//
	// select sleep(2)
	SQLText *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetDBName(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.DBName = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetExecutionStartTime(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ExecutionStartTime = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetHostAddress(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.HostAddress = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetLockTimes(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.LockTimes = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetParseRowCounts(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetQueryTimeMS(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.QueryTimeMS = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetQueryTimes(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.QueryTimes = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetReturnRowCounts(v int64) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs) SetSQLText(v string) *DescribeDatabaseSlowLogRecordsResponseBodySlowLogs {
	s.SQLText = &v
	return s
}

type DescribeDatabaseSlowLogRecordsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDatabaseSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetHeaders(v map[string]*string) *DescribeDatabaseSlowLogRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetStatusCode(v int32) *DescribeDatabaseSlowLogRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabaseSlowLogRecordsResponse) SetBody(v *DescribeDatabaseSlowLogRecordsResponseBody) *DescribeDatabaseSlowLogRecordsResponse {
	s.Body = v
	return s
}

type DescribeFirewallTemplateApplyResultsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The ID of the simple application server to which the template belongs.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the execution to apply the template.
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
}

func (s DescribeFirewallTemplateApplyResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetClientToken(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetFirewallTemplateId(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetInstanceId(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetPageNumber(v int32) *DescribeFirewallTemplateApplyResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetPageSize(v int32) *DescribeFirewallTemplateApplyResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetRegionId(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetTaskId(v []*string) *DescribeFirewallTemplateApplyResultsRequest {
	s.TaskId = v
	return s
}

type DescribeFirewallTemplateApplyResultsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The returned results.
	Data []*DescribeFirewallTemplateApplyResultsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeFirewallTemplateApplyResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetPageNumber(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetPageSize(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetRequestId(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetTotalCount(v string) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBody) SetData(v []*DescribeFirewallTemplateApplyResultsResponseBodyData) *DescribeFirewallTemplateApplyResultsResponseBody {
	s.Data = v
	return s
}

type DescribeFirewallTemplateApplyResultsResponseBodyData struct {
	// The time when the firewall template was applied to the simple application servers.
	//
	// example:
	//
	// Tue May 14 11:53:07 CST 2024
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The total number of simple application servers to which the firewall template fails to apply.
	//
	// example:
	//
	// 0
	FailedCount *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The ID of the firewall template.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The result of applying the firewall template to the simple application servers.
	InstanceApplyResults []*DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults `json:"InstanceApplyResults,omitempty" xml:"InstanceApplyResults,omitempty" type:"Repeated"`
	// The status of applying the template to all simple application servers. Valid values:
	//
	// 	- Running: The firewall template is being applied to simple application servers.
	//
	// 	- Failed: The firewall template is applied to none of simple application servers.
	//
	// 	- Success: The firewall template is applied to all simple application servers.
	//
	// 	- PartFailed: The firewall template fails to be applied to some simple application servers.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the execution to apply the template.
	//
	// example:
	//
	// aft-ikgt0bynitvs3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetCreateTime(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetFailedCount(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.FailedCount = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetFirewallTemplateId(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetInstanceApplyResults(v []*DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.InstanceApplyResults = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetStatus(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetTaskId(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyData) SetTotalCount(v string) *DescribeFirewallTemplateApplyResultsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// 33774babccc84003a2b1ad47e8001233
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of applying the firewall template to the simple application servers. Valid values:
	//
	// 	- Running: The firewall template is being applied to the simple application servers.
	//
	// 	- Failed: The firewall template is applied to none of the simple application servers.
	//
	// 	- Success: The firewall template is applied to all the simple application servers.
	//
	// 	- PartFailed: The firewall template fails to be applied to some simple application servers.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) SetInstanceId(v string) *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults {
	s.InstanceId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults) SetStatus(v string) *DescribeFirewallTemplateApplyResultsResponseBodyDataInstanceApplyResults {
	s.Status = &v
	return s
}

type DescribeFirewallTemplateApplyResultsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTemplateApplyResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTemplateApplyResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsResponse) SetHeaders(v map[string]*string) *DescribeFirewallTemplateApplyResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponse) SetStatusCode(v int32) *DescribeFirewallTemplateApplyResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsResponse) SetBody(v *DescribeFirewallTemplateApplyResultsResponseBody) *DescribeFirewallTemplateApplyResultsResponse {
	s.Body = v
	return s
}

type DescribeFirewallTemplateRulesApplyResultRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the execution to apply the template rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// aft-ikgt0bynitvs3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeFirewallTemplateRulesApplyResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetClientToken(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetFirewallTemplateId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetInstanceId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetRegionId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetTaskId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.TaskId = &v
	return s
}

type DescribeFirewallTemplateRulesApplyResultResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Data []*DescribeFirewallTemplateRulesApplyResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) SetRequestId(v string) *DescribeFirewallTemplateRulesApplyResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBody) SetData(v []*DescribeFirewallTemplateRulesApplyResultResponseBodyData) *DescribeFirewallTemplateRulesApplyResultResponseBody {
	s.Data = v
	return s
}

type DescribeFirewallTemplateRulesApplyResultResponseBodyData struct {
	// The error codes when the firewall template rule fails to be applied.
	//
	// example:
	//
	// 500
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is displayed when the firewall template rule fails to be applied.
	//
	// example:
	//
	// An error occurred while processing your request.
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 119.145.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The status of applying the firewall template rule to the simple application servers.
	//
	// 	- Pending: The template rule does not start to be applied to the simple application servers.
	//
	// 	- Applying: The template rule is being applied to the simple application servers.
	//
	// 	- Success: The template rule is successfully applied to the simple application servers.
	//
	// 	- PartFailed: The template rule fails to be applied to some simple application servers.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetErrorCode(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetErrorInfo(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetPort(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.Port = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetRemark(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.Remark = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetRuleProtocol(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.RuleProtocol = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetSourceCidrIp(v string) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponseBodyData) SetSuccess(v bool) *DescribeFirewallTemplateRulesApplyResultResponseBodyData {
	s.Success = &v
	return s
}

type DescribeFirewallTemplateRulesApplyResultResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTemplateRulesApplyResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTemplateRulesApplyResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) SetHeaders(v map[string]*string) *DescribeFirewallTemplateRulesApplyResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) SetStatusCode(v int32) *DescribeFirewallTemplateRulesApplyResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultResponse) SetBody(v *DescribeFirewallTemplateRulesApplyResultResponseBody) *DescribeFirewallTemplateRulesApplyResultResponse {
	s.Body = v
	return s
}

type DescribeFirewallTemplatesRequest struct {
	// The IDs of the firewall templates.
	FirewallTemplateId []*string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFirewallTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesRequest) SetFirewallTemplateId(v []*string) *DescribeFirewallTemplatesRequest {
	s.FirewallTemplateId = v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetName(v string) *DescribeFirewallTemplatesRequest {
	s.Name = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetPageNumber(v int32) *DescribeFirewallTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetPageSize(v int32) *DescribeFirewallTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplatesRequest) SetRegionId(v string) *DescribeFirewallTemplatesRequest {
	s.RegionId = &v
	return s
}

type DescribeFirewallTemplatesResponseBody struct {
	// The information about the queried firewall templates.
	FirewallTemplates []*DescribeFirewallTemplatesResponseBodyFirewallTemplates `json:"FirewallTemplates,omitempty" xml:"FirewallTemplates,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFirewallTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponseBody) SetFirewallTemplates(v []*DescribeFirewallTemplatesResponseBodyFirewallTemplates) *DescribeFirewallTemplatesResponseBody {
	s.FirewallTemplates = v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetPageNumber(v int32) *DescribeFirewallTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetPageSize(v int32) *DescribeFirewallTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetRequestId(v string) *DescribeFirewallTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBody) SetTotalCount(v int32) *DescribeFirewallTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeFirewallTemplatesResponseBodyFirewallTemplates struct {
	// The time when the firewall template was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// >  The time displayed in the Simple Application Server console is in the format of UTC+8.
	//
	// example:
	//
	// 2024-04-29T02:01:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the firewall template was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-04-10T02:10:14Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the firewall template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the firewall template.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The details of the firewall template rules.
	FirewallTemplateRules []*DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules `json:"FirewallTemplateRules,omitempty" xml:"FirewallTemplateRules,omitempty" type:"Repeated"`
	// The name of the firewall template.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplates) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetCreateTime(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetCreationTime(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.CreationTime = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetDescription(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.Description = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetFirewallTemplateId(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetFirewallTemplateRules(v []*DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.FirewallTemplateRules = v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplates) SetName(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplates {
	s.Name = &v
	return s
}

type DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules struct {
	// The ID of the firewall template rule.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	FirewallTemplateRuleId *string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty"`
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall template rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 119.145.XX.XX
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetFirewallTemplateRuleId(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.FirewallTemplateRuleId = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetPort(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.Port = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetRemark(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.Remark = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetRuleProtocol(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.RuleProtocol = &v
	return s
}

func (s *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules) SetSourceCidrIp(v string) *DescribeFirewallTemplatesResponseBodyFirewallTemplatesFirewallTemplateRules {
	s.SourceCidrIp = &v
	return s
}

type DescribeFirewallTemplatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirewallTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFirewallTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirewallTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplatesResponse) SetHeaders(v map[string]*string) *DescribeFirewallTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirewallTemplatesResponse) SetStatusCode(v int32) *DescribeFirewallTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirewallTemplatesResponse) SetBody(v *DescribeFirewallTemplatesResponseBody) *DescribeFirewallTemplatesResponse {
	s.Body = v
	return s
}

type DescribeInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeyPairRequest) SetClientToken(v string) *DescribeInstanceKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceKeyPairRequest) SetInstanceId(v string) *DescribeInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceKeyPairRequest) SetRegionId(v string) *DescribeInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceKeyPairResponseBody struct {
	// The fingerprint of the key pair.
	//
	// example:
	//
	// 4f:70:62:e9:0c:72:f7:ee:74:ce:e3:bf:e0:82:**:**
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeyPairResponseBody) SetFingerprint(v string) *DescribeInstanceKeyPairResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *DescribeInstanceKeyPairResponseBody) SetKeyPairName(v string) *DescribeInstanceKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *DescribeInstanceKeyPairResponseBody) SetRequestId(v string) *DescribeInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceKeyPairResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceKeyPairResponse) SetHeaders(v map[string]*string) *DescribeInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceKeyPairResponse) SetStatusCode(v int32) *DescribeInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceKeyPairResponse) SetBody(v *DescribeInstanceKeyPairResponseBody) *DescribeInstanceKeyPairResponse {
	s.Body = v
	return s
}

type DescribeInstancePasswordsSettingRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstancePasswordsSettingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePasswordsSettingRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancePasswordsSettingRequest) SetClientToken(v string) *DescribeInstancePasswordsSettingRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstancePasswordsSettingRequest) SetInstanceId(v string) *DescribeInstancePasswordsSettingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancePasswordsSettingRequest) SetRegionId(v string) *DescribeInstancePasswordsSettingRequest {
	s.RegionId = &v
	return s
}

type DescribeInstancePasswordsSettingResponseBody struct {
	// Indicates whether a logon password is set for the simple application server.
	//
	// example:
	//
	// true
	InstancePasswordSetting *bool `json:"InstancePasswordSetting,omitempty" xml:"InstancePasswordSetting,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether a VNC connection password is set.
	//
	// example:
	//
	// true
	VncPasswordSetting *bool `json:"VncPasswordSetting,omitempty" xml:"VncPasswordSetting,omitempty"`
}

func (s DescribeInstancePasswordsSettingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePasswordsSettingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancePasswordsSettingResponseBody) SetInstancePasswordSetting(v bool) *DescribeInstancePasswordsSettingResponseBody {
	s.InstancePasswordSetting = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponseBody) SetRequestId(v string) *DescribeInstancePasswordsSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponseBody) SetVncPasswordSetting(v bool) *DescribeInstancePasswordsSettingResponseBody {
	s.VncPasswordSetting = &v
	return s
}

type DescribeInstancePasswordsSettingResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstancePasswordsSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstancePasswordsSettingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancePasswordsSettingResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancePasswordsSettingResponse) SetHeaders(v map[string]*string) *DescribeInstancePasswordsSettingResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancePasswordsSettingResponse) SetStatusCode(v int32) *DescribeInstancePasswordsSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancePasswordsSettingResponse) SetBody(v *DescribeInstancePasswordsSettingResponseBody) *DescribeInstancePasswordsSettingResponse {
	s.Body = v
	return s
}

type DescribeInstanceVncUrlRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceVncUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlRequest) SetClientToken(v string) *DescribeInstanceVncUrlRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetInstanceId(v string) *DescribeInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceVncUrlRequest) SetRegionId(v string) *DescribeInstanceVncUrlRequest {
	s.RegionId = &v
	return s
}

type DescribeInstanceVncUrlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VNC connection address of the server.
	//
	// example:
	//
	// wss%3A%2F%2Fhz01-vncproxy.aliyun.com%2Fwebsockify%2F%3Fs%3Dwz3L8wEMO6KMt7%252FXInEMtKVubBB%252F7rv055kOm8eUOD%252*****YlmsKjOfz6
	VncUrl *string `json:"VncUrl,omitempty" xml:"VncUrl,omitempty"`
}

func (s DescribeInstanceVncUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponseBody) SetRequestId(v string) *DescribeInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceVncUrlResponseBody) SetVncUrl(v string) *DescribeInstanceVncUrlResponseBody {
	s.VncUrl = &v
	return s
}

type DescribeInstanceVncUrlResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceVncUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceVncUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceVncUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceVncUrlResponse) SetHeaders(v map[string]*string) *DescribeInstanceVncUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceVncUrlResponse) SetStatusCode(v int32) *DescribeInstanceVncUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceVncUrlResponse) SetBody(v *DescribeInstanceVncUrlResponseBody) *DescribeInstanceVncUrlResponse {
	s.Body = v
	return s
}

type DescribeInvocationResultRequest struct {
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution ID. You can call the [DescribeInvocations](https://help.aliyun.com/document_detail/439368.html) operation to query execution IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-bj02prjhw1n****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The region ID of the simple application server. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultRequest) SetInstanceId(v string) *DescribeInvocationResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultRequest) SetInvokeId(v string) *DescribeInvocationResultRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultRequest) SetRegionId(v string) *DescribeInvocationResultRequest {
	s.RegionId = &v
	return s
}

type DescribeInvocationResultResponseBody struct {
	// The execution results.
	InvocationResult *DescribeInvocationResultResponseBodyInvocationResult `json:"InvocationResult,omitempty" xml:"InvocationResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponseBody) SetInvocationResult(v *DescribeInvocationResultResponseBodyInvocationResult) *DescribeInvocationResultResponseBody {
	s.InvocationResult = v
	return s
}

func (s *DescribeInvocationResultResponseBody) SetRequestId(v string) *DescribeInvocationResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvocationResultResponseBodyInvocationResult struct {
	// The error code that is returned if the command failed to be sent or executed.
	//
	// 	- If this parameter is empty, the command is executed normally.
	//
	// 	- InstanceNotExists: The specified server does not exist or is released.
	//
	// 	- InstanceReleased: The server was released while the command was being executed on the server.
	//
	// 	- InstanceNotRunning: The server is not in the Running state while the command is being executed.
	//
	// 	- CommandNotApplicable: The command is not applicable to the specified server.
	//
	// 	- AccountNotExists: The specified account does not exist.
	//
	// 	- DirectoryNotExists: The specified directory does not exist.
	//
	// 	- BadCronExpression: The specified cron expression for the execution schedule is invalid.
	//
	// 	- ClientNotRunning: The Cloud Assistant client is not running.
	//
	// 	- ClientNotResponse: The Cloud Assistant client does not respond.
	//
	// 	- ClientIsUpgrading: The Cloud Assistant client is being upgraded.
	//
	// 	- ClientNeedUpgrade: The Cloud Assistant client needs to be upgraded.
	//
	// 	- DeliveryTimeout: Command sending times out.
	//
	// 	- ExecutionTimeout: The execution times out.
	//
	// 	- ExecutionException: An exception occurs while the command is being executed.
	//
	// 	- ExecutionInterrupted: The execution is interrupted.
	//
	// 	- ExitCodeNonzero: The execution is complete, but the exit code is not 0.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the command is not successfully sent or executed. Valid values:
	//
	// 	- If this parameter is empty, the command is executed normally.
	//
	// 	- the specified instance does not exists: The specified server does not exist or is released.
	//
	// 	- the instance has released when create task: The server was released while the command was being executed on the server.
	//
	// 	- the instance is not running when create task: The server is not in the Running state while the command is being executed.
	//
	// 	- the command is not applicable: The command is not applicable to the specified server.
	//
	// 	- the specified account does not exists: The specified account does not exist.
	//
	// 	- the specified directory does not exists: The specified directory does not exist.
	//
	// 	- the cron job expression is invalid: The specified cron expression is invalid.
	//
	// 	- the aliyun service is not running on the instance: The Cloud Assistance client is not running.
	//
	// 	- the aliyun service in the instance does not response: The Cloud Assistant client does not respond to your request.
	//
	// 	- the aliyun service in the instance is upgrading now: The Cloud Assistant client is being upgraded.
	//
	// 	- the aliyun service in the instance need upgrade: The Cloud Assistant client needs to be upgraded.
	//
	// 	- the command delivery has been timeout: Command sending times out.
	//
	// 	- the command execution has been timeout: The execution times out.
	//
	// 	- the command execution got an exception: An exception occurs while the command is being executed.
	//
	// 	- the command execution has been interrupted: The execution is interrupted.
	//
	// 	- the command execution exit code is not zero: The execution is complete, and the exit code is not 0.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command.
	//
	// 	- For Linux instances, the exit code is the exit code of the shell command.
	//
	// 	- For Windows instances, the exit code is the exit code of the batch or PowerShell command.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the execution ended.
	//
	// example:
	//
	// 2022-07-11T06:37:17Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the execution progress. Valid values:
	//
	// 	- Pending: The command is being verified or sent.
	//
	// 	- Invalid: The specified command type or parameter is invalid.
	//
	// 	- Aborted: The command fails to be sent to the server. To send a command to a server, make sure that the server is in the Running state and the command can be sent within 1 minute.
	//
	// 	- Running: The command is being executed on the server.
	//
	// 	- Success: The execution is completed, and the exit code is 0.
	//
	// 	- Failed: The execution is completed, and the exit code is not 0.
	//
	// 	- Error: The execution cannot proceed due to an exception.
	//
	// 	- Timeout: The execution times out.
	//
	// 	- Cancelled: The execution is canceled, and the command is not executed.
	//
	// 	- Stopping: The command in the Running state is being stopped.
	//
	// 	- Terminated: The command is terminated while it is being executed.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The execution ID.
	//
	// example:
	//
	// t-bj02prjje65****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The status of the execution. Valid values:
	//
	// 	- Running
	//
	// 	- Finished
	//
	// 	- Failed
	//
	// 	- Stopped
	//
	// example:
	//
	// Finished
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The username who executes the command on the simple application server.
	//
	// example:
	//
	// root
	InvokeUser *string `json:"InvokeUser,omitempty" xml:"InvokeUser,omitempty"`
	// The command output.
	//
	// example:
	//
	// YWRtaW4K
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2022-07-11T06:37:16Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvocationResultResponseBodyInvocationResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultResponseBodyInvocationResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetErrorCode(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetErrorInfo(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetExitCode(v int64) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetFinishedTime(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.FinishedTime = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInstanceId(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvocationStatus(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeId(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeRecordStatus(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeUser(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeUser = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetOutput(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.Output = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetStartTime(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.StartTime = &v
	return s
}

type DescribeInvocationResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponse) SetHeaders(v map[string]*string) *DescribeInvocationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationResultResponse) SetStatusCode(v int32) *DescribeInvocationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationResultResponse) SetBody(v *DescribeInvocationResultResponseBody) *DescribeInvocationResultResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3a658ca270df4df39f22e289b338****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the command execution. Valid values:
	//
	// 	- Running: The command is being executed.
	//
	// 	- Finished: The execution is complete.
	//
	// 	- Failed: The execution fails.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetInstanceId(v string) *DescribeInvocationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageNumber(v int32) *DescribeInvocationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageSize(v int32) *DescribeInvocationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	// The command name.
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageNumber(v int32) *DescribeInvocationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageSize(v int32) *DescribeInvocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetTotalCount(v int32) *DescribeInvocationsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInvocationsResponseBodyInvocations struct {
	// The content of the command, which is Base64-encoded.
	//
	// example:
	//
	// bHM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The name of the command.
	//
	// example:
	//
	// testname
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The type of the command. Valid values:
	//
	// 	- RunBatScript: batch command (applicable to Windows instances).
	//
	// 	- RunPowerShellScript: PowerShell command (applicable to Windows instances).
	//
	// 	- RunShellScript: shell command (applicable to Linux instances).
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The time when the command was created.
	//
	// example:
	//
	// 2022-07-11T06:37:16Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The status of the command. Valid values:
	//
	// 	- Pending: The command is being verified or sent.
	//
	// 	- Invalid: The specified command type or parameter is invalid.
	//
	// 	- Aborted: The command failed to be sent. To send a command to an instance, make sure that the instance is in the Running state and the command is sent to the instance within 1 minute.
	//
	// 	- Running: The command is being run on the instance.
	//
	// 	- Success: The command finishes running, and the exit code is 0.
	//
	// 	- Failed: The command finishes running, but the exit code is not 0.
	//
	// 	- Error: The running of the command cannot proceed due to an exception.
	//
	// 	- Timeout: The running of the command times out.
	//
	// 	- Cancelled: The running is canceled, and the command is not run.
	//
	// 	- Stopping: The command that is running is being stopped.
	//
	// 	- Terminated: The command is terminated while it is being run.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The ID of the command task.
	//
	// example:
	//
	// t-hz02p9545t6****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The status of the command. Valid values:
	//
	// 	- Running: The command is running.
	//
	// 	- Finished: The command finishes running.
	//
	// 	- Failed: The running of the command failed.
	//
	// 	- Stopped: The running is stopped.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The custom parameters in the command. If no custom parameter exists in the command, the default value is {}.
	//
	// example:
	//
	// {}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandType(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetParameters(v map[string]interface{}) *DescribeInvocationsResponseBodyInvocations {
	s.Parameters = v
	return s
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetStatusCode(v int32) *DescribeInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type DescribeMonitorDataRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 January 1, 1970.
	//
	// 	- Time format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > The interval between the start time and the end time is less than or equal to 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-08T08:04:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Valid values: 1 to 1440.
	//
	// example:
	//
	// 100
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the metric. Valid values:
	//
	// 	- MEMORY_ACTUALUSEDSPACE: the memory usage. Unit: bytes.
	//
	// 	- DISKUSAGE_USED: the disk usage. Unit: bytes.
	//
	// 	- CPU_UTILIZATION: the CPU usage, in percentage.
	//
	// 	- VPC_PUBLICIP_INTERNETOUT_RATE: the outbound bandwidth. Unit: bits/s.
	//
	// 	- VPC_PUBLICIP_INTERNETIN_RATE: the inbound bandwidth. Unit: bits/s.
	//
	// 	- DISK_READ_IOPS: the read IOPS of the disk. Unit: count/s.
	//
	// 	- DISK_WRITE_IOPS: the write IOPS of the disk. Unit: count/s.
	//
	// 	- FLOW_USED: the traffic usage. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISKUSAGE_USED
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nOc1nj4M9UaAZ/I8db***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The interval at which the monitoring data is queried. Valid values: 60, 300, and 900. Unit: seconds.
	//
	// >
	//
	// If MetricName is set to FLOW_USED, Period is set to 3600 (one hour). In other cases, set Period based on your business requirements.
	//
	// **
	//
	// ****
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 January 1, 1970.
	//
	// 	- Time format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > The specified time range includes the end time and excludes the start time. The start time must be earlier than the end time.
	//
	// The interval between the start time and the end time is less than or equal to 31 days.
	//
	// **
	//
	// ****
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMonitorDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataRequest) SetClientToken(v string) *DescribeMonitorDataRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetEndTime(v string) *DescribeMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetInstanceId(v string) *DescribeMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetLength(v string) *DescribeMonitorDataRequest {
	s.Length = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetMetricName(v string) *DescribeMonitorDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetNextToken(v string) *DescribeMonitorDataRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetPeriod(v string) *DescribeMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetRegionId(v string) *DescribeMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetStartTime(v string) *DescribeMonitorDataRequest {
	s.StartTime = &v
	return s
}

type DescribeMonitorDataResponseBody struct {
	// The monitoring data.
	//
	// example:
	//
	// []
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nOc1nj4M9UaAZ/I8db***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The interval at which the monitoring data is queried. Valid values: 60, 300, and 900. Unit: seconds.
	//
	// >
	//
	// If MetricName is set to FLOW_USED, the value of Period is 3600 (one hour).
	//
	// **
	//
	// ****
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMonitorDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponseBody) SetDatapoints(v string) *DescribeMonitorDataResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetNextToken(v string) *DescribeMonitorDataResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetPeriod(v string) *DescribeMonitorDataResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMonitorDataResponseBody) SetRequestId(v string) *DescribeMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMonitorDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMonitorDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataResponse) SetHeaders(v map[string]*string) *DescribeMonitorDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorDataResponse) SetStatusCode(v int32) *DescribeMonitorDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorDataResponse) SetBody(v *DescribeMonitorDataResponseBody) *DescribeMonitorDataResponse {
	s.Body = v
	return s
}

type DescribeSecurityAgentStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityAgentStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityAgentStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityAgentStatusRequest) SetClientToken(v string) *DescribeSecurityAgentStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSecurityAgentStatusRequest) SetInstanceId(v string) *DescribeSecurityAgentStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityAgentStatusRequest) SetRegionId(v string) *DescribeSecurityAgentStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeSecurityAgentStatusResponseBody struct {
	// The status of the Security Center agent. Valid values:
	//
	// 	- pause: The Security Center agent suspends protection for your server.
	//
	// 	- online: The Security Center agent is protecting your server.
	//
	// 	- offline: The Security Center agent does not protect your server.
	//
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityAgentStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityAgentStatusResponseBody) SetClientStatus(v string) *DescribeSecurityAgentStatusResponseBody {
	s.ClientStatus = &v
	return s
}

func (s *DescribeSecurityAgentStatusResponseBody) SetRequestId(v string) *DescribeSecurityAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

type DescribeSecurityAgentStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityAgentStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityAgentStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityAgentStatusResponse) SetHeaders(v map[string]*string) *DescribeSecurityAgentStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityAgentStatusResponse) SetStatusCode(v int32) *DescribeSecurityAgentStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityAgentStatusResponse) SetBody(v *DescribeSecurityAgentStatusResponseBody) *DescribeSecurityAgentStatusResponse {
	s.Body = v
	return s
}

type DetachKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers from which you want to unbind SSH key pairs. You can specify a maximum of 50 IDs of simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the key pair. The name must be globally unique. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter but cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairRequest) GoString() string {
	return s.String()
}

func (s *DetachKeyPairRequest) SetClientToken(v string) *DetachKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachKeyPairRequest) SetInstanceIds(v []*string) *DetachKeyPairRequest {
	s.InstanceIds = v
	return s
}

func (s *DetachKeyPairRequest) SetKeyPairName(v string) *DetachKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *DetachKeyPairRequest) SetRegionId(v string) *DetachKeyPairRequest {
	s.RegionId = &v
	return s
}

type DetachKeyPairResponseBody struct {
	// The total number of simple application servers from which you fail to unbind key pairs.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request results.
	Results []*DetachKeyPairResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The total number of simple application servers from which the SSH key pair is unbound.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DetachKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBody) SetFailCount(v int32) *DetachKeyPairResponseBody {
	s.FailCount = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetRequestId(v string) *DetachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachKeyPairResponseBody) SetResults(v []*DetachKeyPairResponseBodyResults) *DetachKeyPairResponseBody {
	s.Results = v
	return s
}

func (s *DetachKeyPairResponseBody) SetTotalCount(v int32) *DetachKeyPairResponseBody {
	s.TotalCount = &v
	return s
}

type DetachKeyPairResponseBodyResults struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// aa6e71ddb35c46679bc4753d6219d604
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the key pair is unbound from the simple application server successfully. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachKeyPairResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyResults) SetCode(v string) *DetachKeyPairResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) SetInstanceId(v string) *DetachKeyPairResponseBodyResults {
	s.InstanceId = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) SetMessage(v string) *DetachKeyPairResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DetachKeyPairResponseBodyResults) SetSuccess(v string) *DetachKeyPairResponseBodyResults {
	s.Success = &v
	return s
}

type DetachKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachKeyPairResponse) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponse) SetHeaders(v map[string]*string) *DetachKeyPairResponse {
	s.Headers = v
	return s
}

func (s *DetachKeyPairResponse) SetStatusCode(v int32) *DetachKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachKeyPairResponse) SetBody(v *DetachKeyPairResponseBody) *DetachKeyPairResponse {
	s.Body = v
	return s
}

type DisableFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// custom
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule. You can call the ListFirewallRules operation to query the ID of the firewall rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableFirewallRuleRequest) SetClientToken(v string) *DisableFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetInstanceId(v string) *DisableFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetRegionId(v string) *DisableFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetRemark(v string) *DisableFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *DisableFirewallRuleRequest) SetRuleId(v string) *DisableFirewallRuleRequest {
	s.RuleId = &v
	return s
}

type DisableFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableFirewallRuleResponseBody) SetRequestId(v string) *DisableFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type DisableFirewallRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableFirewallRuleResponse) SetHeaders(v map[string]*string) *DisableFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableFirewallRuleResponse) SetStatusCode(v int32) *DisableFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableFirewallRuleResponse) SetBody(v *DisableFirewallRuleResponseBody) *DisableFirewallRuleResponse {
	s.Body = v
	return s
}

type EnableFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// custom
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The IP address or CIDR block that is allowed in the firewall policy.
	//
	// example:
	//
	// 10.147.33.**
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s EnableFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableFirewallRuleRequest) SetClientToken(v string) *EnableFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *EnableFirewallRuleRequest) SetInstanceId(v string) *EnableFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *EnableFirewallRuleRequest) SetRegionId(v string) *EnableFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *EnableFirewallRuleRequest) SetRemark(v string) *EnableFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *EnableFirewallRuleRequest) SetRuleId(v string) *EnableFirewallRuleRequest {
	s.RuleId = &v
	return s
}

func (s *EnableFirewallRuleRequest) SetSourceCidrIp(v string) *EnableFirewallRuleRequest {
	s.SourceCidrIp = &v
	return s
}

type EnableFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableFirewallRuleResponseBody) SetRequestId(v string) *EnableFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableFirewallRuleResponse) SetHeaders(v map[string]*string) *EnableFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableFirewallRuleResponse) SetStatusCode(v int32) *EnableFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableFirewallRuleResponse) SetBody(v *EnableFirewallRuleResponseBody) *EnableFirewallRuleResponse {
	s.Body = v
	return s
}

type ImportKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the key pair. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key of the key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbO5Govwhb0iHzoMYKkIQxjlHyHH8nxFsW6KF5saxgYhOwdeIpWngpi+/NDWQKvuOnXFFDh/o3eJJkh3rqP+RlMggt4HLQWOd9TS0f4/cgbAzud1caW9PnankCr****
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ImportKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) SetClientToken(v string) *ImportKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

func (s *ImportKeyPairRequest) SetRegionId(v string) *ImportKeyPairRequest {
	s.RegionId = &v
	return s
}

type ImportKeyPairResponseBody struct {
	// The name of the key pair. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with http:// or https://.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) SetKeyPairName(v string) *ImportKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type ImportKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponse) SetHeaders(v map[string]*string) *ImportKeyPairResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyPairResponse) SetStatusCode(v int32) *ImportKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeyPairResponse) SetBody(v *ImportKeyPairResponseBody) *ImportKeyPairResponse {
	s.Body = v
	return s
}

type InstallCloudAssistantRequest struct {
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudAssistantRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantRequest) SetInstanceIds(v []*string) *InstallCloudAssistantRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallCloudAssistantRequest) SetRegionId(v string) *InstallCloudAssistantRequest {
	s.RegionId = &v
	return s
}

type InstallCloudAssistantShrinkRequest struct {
	// The IDs of the simple application servers.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudAssistantShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantShrinkRequest) SetInstanceIdsShrink(v string) *InstallCloudAssistantShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *InstallCloudAssistantShrinkRequest) SetRegionId(v string) *InstallCloudAssistantShrinkRequest {
	s.RegionId = &v
	return s
}

type InstallCloudAssistantResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCloudAssistantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantResponseBody) SetRequestId(v string) *InstallCloudAssistantResponseBody {
	s.RequestId = &v
	return s
}

type InstallCloudAssistantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCloudAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCloudAssistantResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudAssistantResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantResponse) SetHeaders(v map[string]*string) *InstallCloudAssistantResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudAssistantResponse) SetStatusCode(v int32) *InstallCloudAssistantResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudAssistantResponse) SetBody(v *InstallCloudAssistantResponseBody) *InstallCloudAssistantResponse {
	s.Body = v
	return s
}

type InstallCloudMonitorAgentRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly install the CloudMonitor agent. Valid values:
	//
	// 	- true (default value): forcibly installs the CloudMonitor agent.
	//
	// 	- false: does not forcibly install the CloudMonitor agent.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ae7106e68eb4402b0dcbd48a9de****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InstallCloudMonitorAgentRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudMonitorAgentRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorAgentRequest) SetClientToken(v string) *InstallCloudMonitorAgentRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) SetForce(v bool) *InstallCloudMonitorAgentRequest {
	s.Force = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) SetInstanceId(v string) *InstallCloudMonitorAgentRequest {
	s.InstanceId = &v
	return s
}

func (s *InstallCloudMonitorAgentRequest) SetRegionId(v string) *InstallCloudMonitorAgentRequest {
	s.RegionId = &v
	return s
}

type InstallCloudMonitorAgentResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallCloudMonitorAgentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudMonitorAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorAgentResponseBody) SetRequestId(v string) *InstallCloudMonitorAgentResponseBody {
	s.RequestId = &v
	return s
}

type InstallCloudMonitorAgentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallCloudMonitorAgentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallCloudMonitorAgentResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallCloudMonitorAgentResponse) GoString() string {
	return s.String()
}

func (s *InstallCloudMonitorAgentResponse) SetHeaders(v map[string]*string) *InstallCloudMonitorAgentResponse {
	s.Headers = v
	return s
}

func (s *InstallCloudMonitorAgentResponse) SetStatusCode(v int32) *InstallCloudMonitorAgentResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallCloudMonitorAgentResponse) SetBody(v *InstallCloudMonitorAgentResponseBody) *InstallCloudMonitorAgentResponse {
	s.Body = v
	return s
}

type InvokeCommandRequest struct {
	// The command ID. You can call the DescribeCommands operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 50 IDs of simple application servers. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The key-value pairs of custom parameters to specify when the custom parameter feature is enabled.
	//
	// 	- You can specify up to 10 custom parameters. Each key in a Map collection cannot be an empty string and can be up to 64 characters in length.
	//
	// 	- Values in a Map collection can be empty strings. The total length of the custom parameters and the original command cannot exceed 18 KB after they are encoded in Base64.
	//
	// 	- The custom parameter names that you specify for the Parameters parameter must be included in the custom parameter names that you specified when you created the command.
	//
	// 	- You can use empty strings to represent the custom parameters that are not specified. If you want to disable the custom parameter feature, you can leave this parameter empty.
	//
	// example:
	//
	// {"delayed_insert_timeout":"600","max_length_for_sort_data":"2048"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user who runs the command in a simple application server. The username cannot exceed 255 characters in length.
	//
	// 	- For Linux servers, the default value is the root username.
	//
	// 	- For Windows servers, the default value is the system username.
	//
	// You can change the user to run the command only for Linux simple application servers.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s InvokeCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeCommandRequest) GoString() string {
	return s.String()
}

func (s *InvokeCommandRequest) SetCommandId(v string) *InvokeCommandRequest {
	s.CommandId = &v
	return s
}

func (s *InvokeCommandRequest) SetInstanceIds(v string) *InvokeCommandRequest {
	s.InstanceIds = &v
	return s
}

func (s *InvokeCommandRequest) SetParameters(v map[string]interface{}) *InvokeCommandRequest {
	s.Parameters = v
	return s
}

func (s *InvokeCommandRequest) SetRegionId(v string) *InvokeCommandRequest {
	s.RegionId = &v
	return s
}

func (s *InvokeCommandRequest) SetUsername(v string) *InvokeCommandRequest {
	s.Username = &v
	return s
}

type InvokeCommandShrinkRequest struct {
	// The command ID. You can call the DescribeCommands operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 50 IDs of simple application servers. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The key-value pairs of custom parameters to specify when the custom parameter feature is enabled.
	//
	// 	- You can specify up to 10 custom parameters. Each key in a Map collection cannot be an empty string and can be up to 64 characters in length.
	//
	// 	- Values in a Map collection can be empty strings. The total length of the custom parameters and the original command cannot exceed 18 KB after they are encoded in Base64.
	//
	// 	- The custom parameter names that you specify for the Parameters parameter must be included in the custom parameter names that you specified when you created the command.
	//
	// 	- You can use empty strings to represent the custom parameters that are not specified. If you want to disable the custom parameter feature, you can leave this parameter empty.
	//
	// example:
	//
	// {"delayed_insert_timeout":"600","max_length_for_sort_data":"2048"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the user who runs the command in a simple application server. The username cannot exceed 255 characters in length.
	//
	// 	- For Linux servers, the default value is the root username.
	//
	// 	- For Windows servers, the default value is the system username.
	//
	// You can change the user to run the command only for Linux simple application servers.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s InvokeCommandShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvokeCommandShrinkRequest) SetCommandId(v string) *InvokeCommandShrinkRequest {
	s.CommandId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetInstanceIds(v string) *InvokeCommandShrinkRequest {
	s.InstanceIds = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetParametersShrink(v string) *InvokeCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetRegionId(v string) *InvokeCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *InvokeCommandShrinkRequest) SetUsername(v string) *InvokeCommandShrinkRequest {
	s.Username = &v
	return s
}

type InvokeCommandResponseBody struct {
	// The execution ID of the command.
	//
	// example:
	//
	// t-bj02prjhw1n****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InvokeCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeCommandResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeCommandResponseBody) SetInvokeId(v string) *InvokeCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *InvokeCommandResponseBody) SetRequestId(v string) *InvokeCommandResponseBody {
	s.RequestId = &v
	return s
}

type InvokeCommandResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InvokeCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InvokeCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeCommandResponse) GoString() string {
	return s.String()
}

func (s *InvokeCommandResponse) SetHeaders(v map[string]*string) *InvokeCommandResponse {
	s.Headers = v
	return s
}

func (s *InvokeCommandResponse) SetStatusCode(v int32) *InvokeCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *InvokeCommandResponse) SetBody(v *InvokeCommandResponseBody) *InvokeCommandResponse {
	s.Body = v
	return s
}

type ListCustomImageShareAccountsRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-2zehv38jjmwva1ee****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCustomImageShareAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImageShareAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsRequest) SetClientToken(v string) *ListCustomImageShareAccountsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetImageId(v string) *ListCustomImageShareAccountsRequest {
	s.ImageId = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetPageNumber(v int32) *ListCustomImageShareAccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetPageSize(v int32) *ListCustomImageShareAccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomImageShareAccountsRequest) SetRegionId(v string) *ListCustomImageShareAccountsRequest {
	s.RegionId = &v
	return s
}

type ListCustomImageShareAccountsResponseBody struct {
	ImageShareUsers []*ListCustomImageShareAccountsResponseBodyImageShareUsers `json:"ImageShareUsers,omitempty" xml:"ImageShareUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomImageShareAccountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImageShareAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsResponseBody) SetImageShareUsers(v []*ListCustomImageShareAccountsResponseBodyImageShareUsers) *ListCustomImageShareAccountsResponseBody {
	s.ImageShareUsers = v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetPageNumber(v int32) *ListCustomImageShareAccountsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetPageSize(v int32) *ListCustomImageShareAccountsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetRequestId(v string) *ListCustomImageShareAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBody) SetTotalCount(v int32) *ListCustomImageShareAccountsResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomImageShareAccountsResponseBodyImageShareUsers struct {
	SharedTime *string `json:"SharedTime,omitempty" xml:"SharedTime,omitempty"`
	// example:
	//
	// 125111425233****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCustomImageShareAccountsResponseBodyImageShareUsers) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImageShareAccountsResponseBodyImageShareUsers) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) SetSharedTime(v string) *ListCustomImageShareAccountsResponseBodyImageShareUsers {
	s.SharedTime = &v
	return s
}

func (s *ListCustomImageShareAccountsResponseBodyImageShareUsers) SetUserId(v int64) *ListCustomImageShareAccountsResponseBodyImageShareUsers {
	s.UserId = &v
	return s
}

type ListCustomImageShareAccountsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomImageShareAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomImageShareAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImageShareAccountsResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImageShareAccountsResponse) SetHeaders(v map[string]*string) *ListCustomImageShareAccountsResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImageShareAccountsResponse) SetStatusCode(v int32) *ListCustomImageShareAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomImageShareAccountsResponse) SetBody(v *ListCustomImageShareAccountsResponseBody) *ListCustomImageShareAccountsResponse {
	s.Body = v
	return s
}

type ListCustomImagesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data disk snapshot.
	//
	// example:
	//
	// s-acscasca****
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The image IDs of the simple application server. The value can be a JSON array that consists of up to 100 image IDs. Separate multiple image IDs with commas (,).
	//
	// example:
	//
	// ["fe9c66133a9d4688872869726b52****", "794c230fd3e64ea19f83f4d7a0ad****"]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The image names of the simple application servers. The value can be a JSON array that consists of up to 100 image names. Separate multiple image names with commas (,).
	//
	// example:
	//
	// ["test1****", "test2****"]
	ImageNames *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// 	- Maximum value: 100.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers corresponding to the custom images. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7yj2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// False
	Share *bool `json:"Share,omitempty" xml:"Share,omitempty"`
	// The ID of the system disk snapshot.
	//
	// example:
	//
	// s-bp14m09pq8***0g6
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// Tag N of the custom image.
	Tag []*ListCustomImagesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListCustomImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequest) SetClientToken(v string) *ListCustomImagesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListCustomImagesRequest) SetDataSnapshotId(v string) *ListCustomImagesRequest {
	s.DataSnapshotId = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageIds(v string) *ListCustomImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *ListCustomImagesRequest) SetImageNames(v string) *ListCustomImagesRequest {
	s.ImageNames = &v
	return s
}

func (s *ListCustomImagesRequest) SetInstanceId(v string) *ListCustomImagesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCustomImagesRequest) SetPageNumber(v int32) *ListCustomImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImagesRequest) SetPageSize(v int32) *ListCustomImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomImagesRequest) SetRegionId(v string) *ListCustomImagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCustomImagesRequest) SetResourceGroupId(v string) *ListCustomImagesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListCustomImagesRequest) SetShare(v bool) *ListCustomImagesRequest {
	s.Share = &v
	return s
}

func (s *ListCustomImagesRequest) SetSystemSnapshotId(v string) *ListCustomImagesRequest {
	s.SystemSnapshotId = &v
	return s
}

func (s *ListCustomImagesRequest) SetTag(v []*ListCustomImagesRequestTag) *ListCustomImagesRequest {
	s.Tag = v
	return s
}

type ListCustomImagesRequestTag struct {
	// The key of tag N. A tag key can be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. A tag value can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCustomImagesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesRequestTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequestTag) SetKey(v string) *ListCustomImagesRequestTag {
	s.Key = &v
	return s
}

func (s *ListCustomImagesRequestTag) SetValue(v string) *ListCustomImagesRequestTag {
	s.Value = &v
	return s
}

type ListCustomImagesResponseBody struct {
	// The array of queried custom images.
	CustomImages []*ListCustomImagesResponseBodyCustomImages `json:"CustomImages,omitempty" xml:"CustomImages,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBody) SetCustomImages(v []*ListCustomImagesResponseBodyCustomImages) *ListCustomImagesResponseBody {
	s.CustomImages = v
	return s
}

func (s *ListCustomImagesResponseBody) SetPageNumber(v string) *ListCustomImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetPageSize(v string) *ListCustomImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetRequestId(v string) *ListCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetTotalCount(v string) *ListCustomImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListCustomImagesResponseBodyCustomImages struct {
	CreateInstances []*string `json:"CreateInstances,omitempty" xml:"CreateInstances,omitempty" type:"Repeated"`
	// The time when the snapshot was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is in UTC.
	//
	// example:
	//
	// 2022-10-09T02:28:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the data disk snapshot.
	//
	// example:
	//
	// s-bp19rn9u8eqzlfb***
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The name of the data disk snapshot.
	//
	// example:
	//
	// data disk snapshot
	DataSnapshotName *string `json:"DataSnapshotName,omitempty" xml:"DataSnapshotName,omitempty"`
	// The description of the custom image.
	//
	// example:
	//
	// test-custom-image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom image.
	//
	// example:
	//
	// m-bp1e79zktg26n2b***
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Indicates whether the custom image is shared to Elastic Compute Service (ECS).
	//
	// example:
	//
	// false
	InShare *bool `json:"InShare,omitempty" xml:"InShare,omitempty"`
	// example:
	//
	// False
	InShareUser *bool `json:"InShareUser,omitempty" xml:"InShareUser,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 2d06ee0520b44de1ae88d4be****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server.
	//
	// example:
	//
	// swas-asdf23***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the custom image.
	//
	// example:
	//
	// hua
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm2h2lvp3ublq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the custom image.
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the system disk snapshot.
	//
	// example:
	//
	// s-bp1h173hj21puxb***
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// The name of the system disk snapshot.
	//
	// example:
	//
	// system disk snapshot
	SystemSnapshotName *string `json:"SystemSnapshotName,omitempty" xml:"SystemSnapshotName,omitempty"`
	// The tags of the custom image.
	Tags []*ListCustomImagesResponseBodyCustomImagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 180185828710****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListCustomImagesResponseBodyCustomImages) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyCustomImages) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyCustomImages) SetCreateInstances(v []*string) *ListCustomImagesResponseBodyCustomImages {
	s.CreateInstances = v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetCreationTime(v string) *ListCustomImagesResponseBodyCustomImages {
	s.CreationTime = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetDataSnapshotId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.DataSnapshotId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetDataSnapshotName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.DataSnapshotName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetDescription(v string) *ListCustomImagesResponseBodyCustomImages {
	s.Description = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetImageId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.ImageId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInShare(v bool) *ListCustomImagesResponseBodyCustomImages {
	s.InShare = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInShareUser(v bool) *ListCustomImagesResponseBodyCustomImages {
	s.InShareUser = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInstanceId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.InstanceId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetInstanceName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.InstanceName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.Name = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetOsType(v string) *ListCustomImagesResponseBodyCustomImages {
	s.OsType = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetRegionId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.RegionId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetResourceGroupId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetStatus(v string) *ListCustomImagesResponseBodyCustomImages {
	s.Status = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetSystemSnapshotId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.SystemSnapshotId = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetSystemSnapshotName(v string) *ListCustomImagesResponseBodyCustomImages {
	s.SystemSnapshotName = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetTags(v []*ListCustomImagesResponseBodyCustomImagesTags) *ListCustomImagesResponseBodyCustomImages {
	s.Tags = v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImages) SetUserId(v int64) *ListCustomImagesResponseBodyCustomImages {
	s.UserId = &v
	return s
}

type ListCustomImagesResponseBodyCustomImagesTags struct {
	// The tag key of the custom image.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the custom image.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListCustomImagesResponseBodyCustomImagesTags) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyCustomImagesTags) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) SetKey(v string) *ListCustomImagesResponseBodyCustomImagesTags {
	s.Key = &v
	return s
}

func (s *ListCustomImagesResponseBodyCustomImagesTags) SetValue(v string) *ListCustomImagesResponseBodyCustomImagesTags {
	s.Value = &v
	return s
}

type ListCustomImagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponse) SetHeaders(v map[string]*string) *ListCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImagesResponse) SetStatusCode(v int32) *ListCustomImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomImagesResponse) SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse {
	s.Body = v
	return s
}

type ListDisksRequest struct {
	// The IDs of the disks. The value can be a JSON array that consists of up to 100 disk IDs. Separate multiple disk IDs with commas (,).
	//
	// example:
	//
	// ["d-bp14wq0149cpp2x****", "d-bp14wq0149cpp2y****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The disk type. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// By default, system disks and data disks are both queried.
	//
	// example:
	//
	// System
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disks.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags that are added to the disks.
	Tag []*ListDisksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListDisksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDisksRequest) GoString() string {
	return s.String()
}

func (s *ListDisksRequest) SetDiskIds(v string) *ListDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *ListDisksRequest) SetDiskType(v string) *ListDisksRequest {
	s.DiskType = &v
	return s
}

func (s *ListDisksRequest) SetInstanceId(v string) *ListDisksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDisksRequest) SetPageNumber(v int32) *ListDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDisksRequest) SetPageSize(v int32) *ListDisksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDisksRequest) SetRegionId(v string) *ListDisksRequest {
	s.RegionId = &v
	return s
}

func (s *ListDisksRequest) SetResourceGroupId(v string) *ListDisksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDisksRequest) SetTag(v []*ListDisksRequestTag) *ListDisksRequest {
	s.Tag = v
	return s
}

type ListDisksRequestTag struct {
	// The tag key. The tag key can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The tag value can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDisksRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListDisksRequestTag) GoString() string {
	return s.String()
}

func (s *ListDisksRequestTag) SetKey(v string) *ListDisksRequestTag {
	s.Key = &v
	return s
}

func (s *ListDisksRequestTag) SetValue(v string) *ListDisksRequestTag {
	s.Value = &v
	return s
}

type ListDisksResponseBody struct {
	// The queried disks.
	Disks []*ListDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBody) SetDisks(v []*ListDisksResponseBodyDisks) *ListDisksResponseBody {
	s.Disks = v
	return s
}

func (s *ListDisksResponseBody) SetPageNumber(v int32) *ListDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDisksResponseBody) SetPageSize(v int32) *ListDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDisksResponseBody) SetRequestId(v string) *ListDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDisksResponseBody) SetTotalCount(v int32) *ListDisksResponseBody {
	s.TotalCount = &v
	return s
}

type ListDisksResponseBodyDisks struct {
	// The category of the disk. Valid values:
	//
	// 	- ESSD: enhanced SSD (ESSD) of PL0
	//
	// 	- SSD: standard SSD
	//
	// 	- CLOUD_EFFICIENCY: ultra disk
	//
	// example:
	//
	// ESSD
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-08T05:31:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The device name of the disk after the disk is attached to the simple application server.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the disk.
	//
	// example:
	//
	// PrePaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp14wq0149cpp2x****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// SystemDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The disk type. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// System
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the simple application server to which the disk is attached.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server.
	//
	// example:
	//
	// myInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the disk.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the disk belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The size of the disk. Unit: GB.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the disk. Valid values:
	//
	// 	- ReIniting: The disk is being initialized.
	//
	// 	- Creating: The disk is being created.
	//
	// 	- In_use: The disk is in use.
	//
	// 	- Available: The disk can be attached.
	//
	// 	- Attaching: The disk is being attached.
	//
	// 	- Detaching: The disk is being detached.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the disks.
	Tags []*ListDisksResponseBodyDisksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListDisksResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBodyDisks) SetCategory(v string) *ListDisksResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetCreationTime(v string) *ListDisksResponseBodyDisks {
	s.CreationTime = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDevice(v string) *ListDisksResponseBodyDisks {
	s.Device = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskChargeType(v string) *ListDisksResponseBodyDisks {
	s.DiskChargeType = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskId(v string) *ListDisksResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskName(v string) *ListDisksResponseBodyDisks {
	s.DiskName = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetDiskType(v string) *ListDisksResponseBodyDisks {
	s.DiskType = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetInstanceId(v string) *ListDisksResponseBodyDisks {
	s.InstanceId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetInstanceName(v string) *ListDisksResponseBodyDisks {
	s.InstanceName = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetRegionId(v string) *ListDisksResponseBodyDisks {
	s.RegionId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetRemark(v string) *ListDisksResponseBodyDisks {
	s.Remark = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetResourceGroupId(v string) *ListDisksResponseBodyDisks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetSize(v int32) *ListDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetStatus(v string) *ListDisksResponseBodyDisks {
	s.Status = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetTags(v []*ListDisksResponseBodyDisksTags) *ListDisksResponseBodyDisks {
	s.Tags = v
	return s
}

type ListDisksResponseBodyDisksTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDisksResponseBodyDisksTags) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponseBodyDisksTags) GoString() string {
	return s.String()
}

func (s *ListDisksResponseBodyDisksTags) SetKey(v string) *ListDisksResponseBodyDisksTags {
	s.Key = &v
	return s
}

func (s *ListDisksResponseBodyDisksTags) SetValue(v string) *ListDisksResponseBodyDisksTags {
	s.Value = &v
	return s
}

type ListDisksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDisksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDisksResponse) GoString() string {
	return s.String()
}

func (s *ListDisksResponse) SetHeaders(v map[string]*string) *ListDisksResponse {
	s.Headers = v
	return s
}

func (s *ListDisksResponse) SetStatusCode(v int32) *ListDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDisksResponse) SetBody(v *ListDisksResponseBody) *ListDisksResponse {
	s.Body = v
	return s
}

type ListFirewallRulesRequest struct {
	// The ID of the firewall rule.
	//
	// example:
	//
	// 1a16263ab0f541288312a15fa64280de
	FirewallRuleId *string `json:"FirewallRuleId,omitempty" xml:"FirewallRuleId,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags of the firewall rule.
	Tag []*ListFirewallRulesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListFirewallRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesRequest) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesRequest) SetFirewallRuleId(v string) *ListFirewallRulesRequest {
	s.FirewallRuleId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetInstanceId(v string) *ListFirewallRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetPageNumber(v int32) *ListFirewallRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFirewallRulesRequest) SetPageSize(v int32) *ListFirewallRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFirewallRulesRequest) SetRegionId(v string) *ListFirewallRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListFirewallRulesRequest) SetTag(v []*ListFirewallRulesRequestTag) *ListFirewallRulesRequest {
	s.Tag = v
	return s
}

type ListFirewallRulesRequestTag struct {
	// The key of tag N to be added to the firewall rule. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to be added to the firewall rule. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFirewallRulesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesRequestTag) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesRequestTag) SetKey(v string) *ListFirewallRulesRequestTag {
	s.Key = &v
	return s
}

func (s *ListFirewallRulesRequestTag) SetValue(v string) *ListFirewallRulesRequestTag {
	s.Value = &v
	return s
}

type ListFirewallRulesResponseBody struct {
	// The array of firewall rules.
	FirewallRules []*ListFirewallRulesResponseBodyFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFirewallRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBody) SetFirewallRules(v []*ListFirewallRulesResponseBodyFirewallRules) *ListFirewallRulesResponseBody {
	s.FirewallRules = v
	return s
}

func (s *ListFirewallRulesResponseBody) SetPageNumber(v int32) *ListFirewallRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetPageSize(v int32) *ListFirewallRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetRequestId(v string) *ListFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFirewallRulesResponseBody) SetTotalCount(v int32) *ListFirewallRulesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFirewallRulesResponseBodyFirewallRules struct {
	// The firewall policy. Valid values:
	//
	// 	- accept: Access is allowed.
	//
	// 	- drop: Access is refused.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// TEST
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source CIDR block.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The tags of the firewall rule.
	Tags []*ListFirewallRulesResponseBodyFirewallRulesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFirewallRulesResponseBodyFirewallRules) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponseBodyFirewallRules) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetPolicy(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Policy = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetPort(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Port = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRemark(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.Remark = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRuleId(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.RuleId = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetRuleProtocol(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.RuleProtocol = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetSourceCidrIp(v string) *ListFirewallRulesResponseBodyFirewallRules {
	s.SourceCidrIp = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRules) SetTags(v []*ListFirewallRulesResponseBodyFirewallRulesTags) *ListFirewallRulesResponseBodyFirewallRules {
	s.Tags = v
	return s
}

type ListFirewallRulesResponseBodyFirewallRulesTags struct {
	// The key of tag N to be added to the firewall rule. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to be added to the firewall rule. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFirewallRulesResponseBodyFirewallRulesTags) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponseBodyFirewallRulesTags) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) SetKey(v string) *ListFirewallRulesResponseBodyFirewallRulesTags {
	s.Key = &v
	return s
}

func (s *ListFirewallRulesResponseBodyFirewallRulesTags) SetValue(v string) *ListFirewallRulesResponseBodyFirewallRulesTags {
	s.Value = &v
	return s
}

type ListFirewallRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFirewallRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesResponse) GoString() string {
	return s.String()
}

func (s *ListFirewallRulesResponse) SetHeaders(v map[string]*string) *ListFirewallRulesResponse {
	s.Headers = v
	return s
}

func (s *ListFirewallRulesResponse) SetStatusCode(v int32) *ListFirewallRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFirewallRulesResponse) SetBody(v *ListFirewallRulesResponseBody) *ListFirewallRulesResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// The image IDs. The value can be a JSON array that consists of up to 50 image IDs. Format: `["xxx", "yyy", … "zzz"]`. Separate multiple image IDs with commas (,).
	//
	// example:
	//
	// ["fe9c66133a9d4688872869726b52****", "794c230fd3e64ea19f83f4d7a0ad****"]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of the images. Valid values:
	//
	// 	- system: OS images
	//
	// 	- app: application images
	//
	// 	- custom: custom images
	//
	// example:
	//
	// system
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The region ID of the images. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetImageIds(v string) *ListImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *ListImagesRequest) SetImageType(v string) *ListImagesRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

type ListImagesResponseBody struct {
	// The OS type of the image. Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	// The description of the image.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// 794c230fd3e64ea19f83f4d7a0ad****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// WordPress-4.8.1
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the image. Valid values:
	//
	// 	- system
	//
	// 	- app
	//
	// 	- custom
	//
	// example:
	//
	// app
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageName(v string) *ListImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageType(v string) *ListImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetPlatform(v string) *ListImagesResponseBodyImages {
	s.Platform = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstancePlansModificationRequest struct {
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePlansModificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationRequest) SetInstanceId(v string) *ListInstancePlansModificationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePlansModificationRequest) SetRegionId(v string) *ListInstancePlansModificationRequest {
	s.RegionId = &v
	return s
}

type ListInstancePlansModificationResponseBody struct {
	// The operating system types supported by the plan.
	Plans []*ListInstancePlansModificationResponseBodyPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePlansModificationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponseBody) SetPlans(v []*ListInstancePlansModificationResponseBodyPlans) *ListInstancePlansModificationResponseBody {
	s.Plans = v
	return s
}

func (s *ListInstancePlansModificationResponseBody) SetRequestId(v string) *ListInstancePlansModificationResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancePlansModificationResponseBodyPlans struct {
	// The peak bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 3
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The unit of the plan price. Valid values:
	//
	// 	- CNY
	//
	// 	- USD
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The disk size of the simple application server. Unit: GB.
	//
	// example:
	//
	// 40
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The category of the disk. Valid values:
	//
	// 	- SSD: standard SSD
	//
	// 	- ESSD: enhanced SSD
	//
	// example:
	//
	// ESSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The monthly data transfer quota. Unit: GB.
	//
	// example:
	//
	// 400
	Flow *int32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 1
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The price of the plan.
	//
	// example:
	//
	// 60
	OriginPrice *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The ID of the plan.
	//
	// example:
	//
	// swas.s2.c2m1s40b3t04
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The operating system types supported by the plan.
	//
	// example:
	//
	// ["Linux","Windows"]
	SupportPlatform *string `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
}

func (s ListInstancePlansModificationResponseBodyPlans) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetBandwidth(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Bandwidth = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetCore(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Core = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetCurrency(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.Currency = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetDiskSize(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.DiskSize = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetDiskType(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.DiskType = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetFlow(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Flow = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetMemory(v int32) *ListInstancePlansModificationResponseBodyPlans {
	s.Memory = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetOriginPrice(v float64) *ListInstancePlansModificationResponseBodyPlans {
	s.OriginPrice = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetPlanId(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListInstancePlansModificationResponseBodyPlans) SetSupportPlatform(v string) *ListInstancePlansModificationResponseBodyPlans {
	s.SupportPlatform = &v
	return s
}

type ListInstancePlansModificationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePlansModificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePlansModificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePlansModificationResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePlansModificationResponse) SetHeaders(v map[string]*string) *ListInstancePlansModificationResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePlansModificationResponse) SetStatusCode(v int32) *ListInstancePlansModificationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePlansModificationResponse) SetBody(v *ListInstancePlansModificationResponseBody) *ListInstancePlansModificationResponse {
	s.Body = v
	return s
}

type ListInstanceStatusRequest struct {
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusRequest) SetInstanceIds(v string) *ListInstanceStatusRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstanceStatusRequest) SetPageNumber(v int32) *ListInstanceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceStatusRequest) SetPageSize(v int32) *ListInstanceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceStatusRequest) SetRegionId(v string) *ListInstanceStatusRequest {
	s.RegionId = &v
	return s
}

type ListInstanceStatusResponseBody struct {
	// The ID of the simple application server.
	InstanceStatuses []*ListInstanceStatusResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 54
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBody) SetInstanceStatuses(v []*ListInstanceStatusResponseBodyInstanceStatuses) *ListInstanceStatusResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *ListInstanceStatusResponseBody) SetPageNumber(v int32) *ListInstanceStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetPageSize(v int32) *ListInstanceStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetRequestId(v string) *ListInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceStatusResponseBody) SetTotalCount(v int32) *ListInstanceStatusResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceStatusResponseBodyInstanceStatuses struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// a9a6474b935d41bcb531250bb5d****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the simple application server. Valid values:
	//
	// 	- Pending
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Resetting
	//
	// 	- Upgrading
	//
	// 	- Disabled
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceStatusResponseBodyInstanceStatuses) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) SetInstanceId(v string) *ListInstanceStatusResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceStatusResponseBodyInstanceStatuses) SetStatus(v string) *ListInstanceStatusResponseBodyInstanceStatuses {
	s.Status = &v
	return s
}

type ListInstanceStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceStatusResponse) SetHeaders(v map[string]*string) *ListInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceStatusResponse) SetStatusCode(v int32) *ListInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceStatusResponse) SetBody(v *ListInstanceStatusResponseBody) *ListInstanceStatusResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	// The billing method of the simple application servers. Set the value to PrePaid, which indicates the subscription billing method.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// > If you specify both `InstanceIds` and `PublicIpAddresses`, make sure that the specified IDs and the specified public IP addresses belong to the same simple application servers. Otherwise, an empty result is returned.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the simple application servers, which supports fuzzy search using wildcard *.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The public IP addresses of the simple application servers. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,).
	//
	// > If you specify both `InstanceIds` and `PublicIpAddresses`, make sure that the specified IDs and the specified public IP addresses belong to the same simple application servers. Otherwise, an empty result is returned.
	//
	// example:
	//
	// ["42.1.\*\*.**", "42.2.\*\*.**"]
	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty"`
	// The region ID of the simple application servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the simple application servers belong.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the simple application servers. Valid values:
	//
	// 	- Pending
	//
	// 	- Starting
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// 	- Stopped
	//
	// 	- Resetting
	//
	// 	- Upgrading
	//
	// 	- Disabled
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the simple application servers.
	Tag []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetChargeType(v string) *ListInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceIds(v string) *ListInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetPublicIpAddresses(v string) *ListInstancesRequest {
	s.PublicIpAddresses = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

type ListInstancesRequestTag struct {
	// The tag key of the simple application servers. A tag key can be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the simple application servers. A tag value can be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// 01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

type ListInstancesResponseBody struct {
	// Details about the simple application servers.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v []*ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetPageNumber(v int32) *ListInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageSize(v int32) *ListInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	// The status of the server. Valid values:
	//
	// 	- Normal: The server is normal.
	//
	// 	- Expired: The server expires.
	//
	// 	- Overdue: The payment of the server is overdue.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the simple application server.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Indicates whether the simple application server uses a bundle plan.
	//
	// example:
	//
	// false
	Combination *bool `json:"Combination,omitempty" xml:"Combination,omitempty"`
	// The ID of the simple application server that uses a bundle plan.
	//
	// example:
	//
	// com-f6c9a22****45b5b8de68ad608af1ba
	CombinationInstanceId *string `json:"CombinationInstanceId,omitempty" xml:"CombinationInstanceId,omitempty"`
	// The time when the simple application server was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-08T05:31:06Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The DDoS protection status of the server. Valid values:
	//
	// 	- Normal: The DDoS protection status of the server is normal.
	//
	// 	- BlackHole: The server is in blackhole filtering.
	//
	// 	- Defense: The server is being scrubbed.
	//
	// example:
	//
	// Normal
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The reason why the server is disabled. Valid values:
	//
	// 	- FINANCIAL: The server is locked due to overdue payments.
	//
	// 	- SECURITY: The server is locked for security reasons.
	//
	// 	- EXPIRED: The server is expired.
	//
	// example:
	//
	// EXPIRED
	DisableReason *string `json:"DisableReason,omitempty" xml:"DisableReason,omitempty"`
	// The disks that are attached to the simple application server.
	Disks []*ListInstancesResponseBodyInstancesDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The time when the simple application server expires. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-08T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The description of the image.
	Image *ListInstancesResponseBodyInstancesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The ID of the image.
	//
	// example:
	//
	// fe9c66133a9d4688872869726b52****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The private IP address of the simple application server.
	//
	// example:
	//
	// 172.26.XX.XX
	InnerIpAddress *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The ID of the instance plan.
	//
	// example:
	//
	// swas.s2.c2m2s50b4t08
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 42.1.XX.XX
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the server belongs.
	//
	// example:
	//
	// rg-aekz7jmhg5s****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The specifications of the resource.
	ResourceSpec *ListInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// The status of the simple application server. Valid values:
	//
	// 	- Pending: The server is being prepared.
	//
	// 	- Starting: The server is being started.
	//
	// 	- Running: The server is running.
	//
	// 	- Stopping: The server is being stopped.
	//
	// 	- Stopped: The server is stopped.
	//
	// 	- Resetting: The server is being reset.
	//
	// 	- Upgrading: The server is being upgraded.
	//
	// 	- Disabled: The server is not available.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the simple application server.
	Tags []*ListInstancesResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The universally unique identifier (UUID) of the simple application server.
	//
	// example:
	//
	// 41f30524-5df7-49c9-9c6e-32****489001
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetBusinessStatus(v string) *ListInstancesResponseBodyInstances {
	s.BusinessStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetChargeType(v string) *ListInstancesResponseBodyInstances {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCombination(v bool) *ListInstancesResponseBodyInstances {
	s.Combination = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCombinationInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.CombinationInstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetCreationTime(v string) *ListInstancesResponseBodyInstances {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDdosStatus(v string) *ListInstancesResponseBodyInstances {
	s.DdosStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDisableReason(v string) *ListInstancesResponseBodyInstances {
	s.DisableReason = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetDisks(v []*ListInstancesResponseBodyInstancesDisks) *ListInstancesResponseBodyInstances {
	s.Disks = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetExpiredTime(v string) *ListInstancesResponseBodyInstances {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImage(v *ListInstancesResponseBodyInstancesImage) *ListInstancesResponseBodyInstances {
	s.Image = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetImageId(v string) *ListInstancesResponseBodyInstances {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInnerIpAddress(v string) *ListInstancesResponseBodyInstances {
	s.InnerIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceId(v string) *ListInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetInstanceName(v string) *ListInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPlanId(v string) *ListInstancesResponseBodyInstances {
	s.PlanId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetPublicIpAddress(v string) *ListInstancesResponseBodyInstances {
	s.PublicIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetRegionId(v string) *ListInstancesResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceGroupId(v string) *ListInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetResourceSpec(v *ListInstancesResponseBodyInstancesResourceSpec) *ListInstancesResponseBodyInstances {
	s.ResourceSpec = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetTags(v []*ListInstancesResponseBodyInstancesTags) *ListInstancesResponseBodyInstances {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUuid(v string) *ListInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

type ListInstancesResponseBodyInstancesDisks struct {
	// The category of the disk. Valid values:
	//
	// 	- ESSD: ESSD of PL0
	//
	// 	- SSD: standard SSD
	//
	// 	- CLOUD_EFFICIENCY: an ultra disk.
	//
	// example:
	//
	// ESSD
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the simple application server was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-24T02:20:10Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The device name of the disk after the disk is attached to the simple application server.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the disk.
	//
	// example:
	//
	// PrePaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-bp14wq0149cpp2x****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The disk name.
	//
	// example:
	//
	// SystemDisk
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The tags that are added to the disk.
	DiskTags []*ListInstancesResponseBodyInstancesDisksDiskTags `json:"DiskTags,omitempty" xml:"DiskTags,omitempty" type:"Repeated"`
	// The disk type. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// System
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the disk.
	//
	// example:
	//
	// remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the disk belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The disk size. Unit: GB.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the disk. Valid values:
	//
	// 	- ReIniting: The disk is being initialized.
	//
	// 	- Creating: The disk is being created.
	//
	// 	- In_use: The disk is in use.
	//
	// 	- Available: The disk can be attached.
	//
	// 	- Attaching: The disk is being attached.
	//
	// 	- Detaching: The disk is being detached.
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDisks) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDisks) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDisks) SetCategory(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Category = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetCreationTime(v string) *ListInstancesResponseBodyInstancesDisks {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDevice(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Device = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskChargeType(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskId(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskName(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskTags(v []*ListInstancesResponseBodyInstancesDisksDiskTags) *ListInstancesResponseBodyInstancesDisks {
	s.DiskTags = v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetDiskType(v string) *ListInstancesResponseBodyInstancesDisks {
	s.DiskType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetRegionId(v string) *ListInstancesResponseBodyInstancesDisks {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetRemark(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Remark = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetResourceGroupId(v string) *ListInstancesResponseBodyInstancesDisks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetSize(v int32) *ListInstancesResponseBodyInstancesDisks {
	s.Size = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisks) SetStatus(v string) *ListInstancesResponseBodyInstancesDisks {
	s.Status = &v
	return s
}

type ListInstancesResponseBodyInstancesDisksDiskTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesDisksDiskTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesDisksDiskTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) SetKey(v string) *ListInstancesResponseBodyInstancesDisksDiskTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesDisksDiskTags) SetValue(v string) *ListInstancesResponseBodyInstancesDisksDiskTags {
	s.Value = &v
	return s
}

type ListInstancesResponseBodyInstancesImage struct {
	// The image provider.
	//
	// example:
	//
	// https://selfs****e.console.aliyun.com/ticket/createIndex
	ImageContact *string `json:"ImageContact,omitempty" xml:"ImageContact,omitempty"`
	// The URL of the image icon.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O****1vdh9651ReKqWNMI2I_!!6000000002136****-24-24.svg
	ImageIconUrl *string `json:"ImageIconUrl,omitempty" xml:"ImageIconUrl,omitempty"`
	// The name of the image.
	//
	// example:
	//
	// test-custom-1686536882356
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the image. Valid values:
	//
	// 	- system
	//
	// 	- app
	//
	// 	- custom
	//
	// example:
	//
	// system
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image tag.
	//
	// example:
	//
	// V3.5
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The OS.
	//
	// example:
	//
	// windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
}

func (s ListInstancesResponseBodyInstancesImage) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesImage) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageContact(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageContact = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageIconUrl(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageIconUrl = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageName(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageType(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetImageVersion(v string) *ListInstancesResponseBodyInstancesImage {
	s.ImageVersion = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesImage) SetOsType(v string) *ListInstancesResponseBodyInstancesImage {
	s.OsType = &v
	return s
}

type ListInstancesResponseBodyInstancesResourceSpec struct {
	// The bandwidth of the server.
	//
	// example:
	//
	// 5
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs of the simple application server.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The category of the disk. Valid values:
	//
	// 	- ESSD: enhanced SSD (ESSD) of PL0
	//
	// 	- SSD: standard SSD
	//
	// 	- CLOUD_EFFICIENCY: ultra disk
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk size.
	//
	// example:
	//
	// 60
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The amount of the traffic.
	//
	// 	- A value of 0 indicates the traffic amount of a bandwidth-based simple application server.
	//
	// 	- A non-zero value indicates the traffic amount of a data transfer plan-based simple application server.
	//
	// example:
	//
	// 818
	Flow *float64 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size of the server.
	//
	// example:
	//
	// 2
	Memory *float64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListInstancesResponseBodyInstancesResourceSpec) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesResourceSpec) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetBandwidth(v int32) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Bandwidth = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetCpu(v int32) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Cpu = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetDiskCategory(v string) *ListInstancesResponseBodyInstancesResourceSpec {
	s.DiskCategory = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetDiskSize(v int32) *ListInstancesResponseBodyInstancesResourceSpec {
	s.DiskSize = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetFlow(v float64) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Flow = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesResourceSpec) SetMemory(v float64) *ListInstancesResponseBodyInstancesResourceSpec {
	s.Memory = &v
	return s
}

type ListInstancesResponseBodyInstancesTags struct {
	// The tag key of the simple application server.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the simple application server.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesTags) SetKey(v string) *ListInstancesResponseBodyInstancesTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesTags) SetValue(v string) *ListInstancesResponseBodyInstancesTags {
	s.Value = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListInstancesTrafficPackagesRequest struct {
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc******","2ad1ae67295445f598017499dc******"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesTrafficPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesRequest) SetInstanceIds(v string) *ListInstancesTrafficPackagesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesTrafficPackagesRequest) SetRegionId(v string) *ListInstancesTrafficPackagesRequest {
	s.RegionId = &v
	return s
}

type ListInstancesTrafficPackagesResponseBody struct {
	// The data transfers that exceed the quota of the data transfer plan in the current month. Unit: bytes.
	InstanceTrafficPackageUsages []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages `json:"InstanceTrafficPackageUsages,omitempty" xml:"InstanceTrafficPackageUsages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancesTrafficPackagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponseBody) SetInstanceTrafficPackageUsages(v []*ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) *ListInstancesTrafficPackagesResponseBody {
	s.InstanceTrafficPackageUsages = v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBody) SetRequestId(v string) *ListInstancesTrafficPackagesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages struct {
	// The ID of the simple application server.
	//
	// example:
	//
	// ccscqwqwqqqw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The data transfers that exceeds the quota of the data transfer plan in the current month. Unit: Byte.
	//
	// example:
	//
	// 0
	TrafficOverflow *int64 `json:"TrafficOverflow,omitempty" xml:"TrafficOverflow,omitempty"`
	// The unused quota of the data transfer plan in the current month. Unit: Byte.
	//
	// example:
	//
	// 10000
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitempty" xml:"TrafficPackageRemaining,omitempty"`
	// The quota of the data transfer plan in the current month. Unit: Byte.
	//
	// >  TrafficPackageTotal = TrafficUsed + TrafficPackageRemaining
	//
	// example:
	//
	// 20000
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitempty" xml:"TrafficPackageTotal,omitempty"`
	// The used quota of the data transfer plan in the current month. Unit: Byte.
	//
	// example:
	//
	// 10000
	TrafficUsed *int64 `json:"TrafficUsed,omitempty" xml:"TrafficUsed,omitempty"`
}

func (s ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetInstanceId(v string) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficOverflow(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficOverflow = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficPackageRemaining(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficPackageRemaining = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficPackageTotal(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficPackageTotal = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages) SetTrafficUsed(v int64) *ListInstancesTrafficPackagesResponseBodyInstanceTrafficPackageUsages {
	s.TrafficUsed = &v
	return s
}

type ListInstancesTrafficPackagesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesTrafficPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesTrafficPackagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesResponse) SetHeaders(v map[string]*string) *ListInstancesTrafficPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) SetStatusCode(v int32) *ListInstancesTrafficPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesTrafficPackagesResponse) SetBody(v *ListInstancesTrafficPackagesResponseBody) *ListInstancesTrafficPackagesResponse {
	s.Body = v
	return s
}

type ListKeyPairsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the AccessKey pair. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with http:// or https://.
	//
	// example:
	//
	// KeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The page number. Page starts from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *ListKeyPairsRequest) SetClientToken(v string) *ListKeyPairsRequest {
	s.ClientToken = &v
	return s
}

func (s *ListKeyPairsRequest) SetKeyPairName(v string) *ListKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *ListKeyPairsRequest) SetPageNumber(v int32) *ListKeyPairsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeyPairsRequest) SetPageSize(v int32) *ListKeyPairsRequest {
	s.PageSize = &v
	return s
}

func (s *ListKeyPairsRequest) SetRegionId(v string) *ListKeyPairsRequest {
	s.RegionId = &v
	return s
}

type ListKeyPairsResponseBody struct {
	// Details about the queried AccessKey pairs.
	KeyPairs []*ListKeyPairsResponseBodyKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBody) SetKeyPairs(v []*ListKeyPairsResponseBodyKeyPairs) *ListKeyPairsResponseBody {
	s.KeyPairs = v
	return s
}

func (s *ListKeyPairsResponseBody) SetPageNumber(v int32) *ListKeyPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetPageSize(v int32) *ListKeyPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetRequestId(v string) *ListKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetTotalCount(v int32) *ListKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

type ListKeyPairsResponseBodyKeyPairs struct {
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2024-05-06T02:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IDs of simple application servers. A maximum of 50 IDs of simple application servers can be returned.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the AccessKey pair. The name must be 2 to 64 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-). The name must start with a letter and cannot start with http:// or https://.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The content of the public key.
	//
	// example:
	//
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbO5Govwhb0iHzoMYKkIQxjlHyHH8nxFsW6KF5saxgYhOwdeIpWngpi+/NDWQKvuOnXFFDh/o3eJJkh3rqP+RlMggt4HLQWOd9TS0f4/cgbAzud1caW9PnankCr****
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
}

func (s ListKeyPairsResponseBodyKeyPairs) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponseBodyKeyPairs) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetCreationTime(v string) *ListKeyPairsResponseBodyKeyPairs {
	s.CreationTime = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetInstanceIds(v []*string) *ListKeyPairsResponseBodyKeyPairs {
	s.InstanceIds = v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetKeyPairName(v string) *ListKeyPairsResponseBodyKeyPairs {
	s.KeyPairName = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetPublicKey(v string) *ListKeyPairsResponseBodyKeyPairs {
	s.PublicKey = &v
	return s
}

type ListKeyPairsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponse) SetHeaders(v map[string]*string) *ListKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *ListKeyPairsResponse) SetStatusCode(v int32) *ListKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeyPairsResponse) SetBody(v *ListKeyPairsResponseBody) *ListKeyPairsResponse {
	s.Body = v
	return s
}

type ListPlansRequest struct {
	// The region ID of the plans. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPlansRequest) GoString() string {
	return s.String()
}

func (s *ListPlansRequest) SetRegionId(v string) *ListPlansRequest {
	s.RegionId = &v
	return s
}

type ListPlansResponseBody struct {
	// The operating system types supported by the plan.
	Plans []*ListPlansResponseBodyPlans `json:"Plans,omitempty" xml:"Plans,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBody) SetPlans(v []*ListPlansResponseBodyPlans) *ListPlansResponseBody {
	s.Plans = v
	return s
}

func (s *ListPlansResponseBody) SetRequestId(v string) *ListPlansResponseBody {
	s.RequestId = &v
	return s
}

type ListPlansResponseBodyPlans struct {
	// The peak bandwidth. Unit: Mbit/s.
	//
	// example:
	//
	// 3
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The unit of the plan price. Valid values:
	//
	// 	- CNY
	//
	// 	- USD
	//
	// >  CNY is for the China site (aliyun.com). USD is for the international site (alibabacloud.com).
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The size of the disk. Unit: GB.
	//
	// example:
	//
	// 40
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The category of the disk. Valid values:
	//
	// 	- SSD: standard SSDs
	//
	// 	- ESSD: enhanced SSDs
	//
	// example:
	//
	// ESSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The monthly data transfer quota. Unit: GB.
	//
	// example:
	//
	// 400
	Flow *int32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 1
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The monthly price of the plan.
	//
	// example:
	//
	// 60
	OriginPrice *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The ID of the plan.
	//
	// example:
	//
	// swas.s2.c2m1s40b3t04
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The operating system types supported by the plan.
	//
	// example:
	//
	// ["Linux","Windows"]
	SupportPlatform *string `json:"SupportPlatform,omitempty" xml:"SupportPlatform,omitempty"`
}

func (s ListPlansResponseBodyPlans) String() string {
	return tea.Prettify(s)
}

func (s ListPlansResponseBodyPlans) GoString() string {
	return s.String()
}

func (s *ListPlansResponseBodyPlans) SetBandwidth(v int32) *ListPlansResponseBodyPlans {
	s.Bandwidth = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCore(v int32) *ListPlansResponseBodyPlans {
	s.Core = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetCurrency(v string) *ListPlansResponseBodyPlans {
	s.Currency = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetDiskSize(v int32) *ListPlansResponseBodyPlans {
	s.DiskSize = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetDiskType(v string) *ListPlansResponseBodyPlans {
	s.DiskType = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetFlow(v int32) *ListPlansResponseBodyPlans {
	s.Flow = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetMemory(v int32) *ListPlansResponseBodyPlans {
	s.Memory = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetOriginPrice(v float64) *ListPlansResponseBodyPlans {
	s.OriginPrice = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetPlanId(v string) *ListPlansResponseBodyPlans {
	s.PlanId = &v
	return s
}

func (s *ListPlansResponseBodyPlans) SetSupportPlatform(v string) *ListPlansResponseBodyPlans {
	s.SupportPlatform = &v
	return s
}

type ListPlansResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPlansResponse) GoString() string {
	return s.String()
}

func (s *ListPlansResponse) SetHeaders(v map[string]*string) *ListPlansResponse {
	s.Headers = v
	return s
}

func (s *ListPlansResponse) SetStatusCode(v int32) *ListPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPlansResponse) SetBody(v *ListPlansResponseBody) *ListPlansResponse {
	s.Body = v
	return s
}

type ListRegionsRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese
	//
	// 	- **en-US**: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) SetAcceptLanguage(v string) *ListRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type ListRegionsResponseBody struct {
	// The regions.
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// swas.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSnapshotsRequest struct {
	// The disk ID.
	//
	// example:
	//
	// d-bp14wq0149cpp2xy****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the simple application server.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server that corresponds to the snapshots.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple snapshot IDs with commas (,).
	//
	// example:
	//
	// ["s-bp16oazlsold4dks****", "s-bp16oazlsold4abc****"]
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The type of the source disk. Valid values:
	//
	// 	- system: system disk.
	//
	// 	- data: data disk.
	//
	// example:
	//
	// System
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// Tag N that you want to add to the snapshot.
	Tag []*ListSnapshotsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) SetDiskId(v string) *ListSnapshotsRequest {
	s.DiskId = &v
	return s
}

func (s *ListSnapshotsRequest) SetInstanceId(v string) *ListSnapshotsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageNumber(v int32) *ListSnapshotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsRequest) SetPageSize(v int32) *ListSnapshotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsRequest) SetRegionId(v string) *ListSnapshotsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotsRequest) SetResourceGroupId(v string) *ListSnapshotsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSnapshotsRequest) SetSnapshotIds(v string) *ListSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *ListSnapshotsRequest) SetSourceDiskType(v string) *ListSnapshotsRequest {
	s.SourceDiskType = &v
	return s
}

func (s *ListSnapshotsRequest) SetTag(v []*ListSnapshotsRequestTag) *ListSnapshotsRequest {
	s.Tag = v
	return s
}

type ListSnapshotsRequestTag struct {
	// The key of tag N that you want to add to the snapshot. A tag key can be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the snapshot. A tag value can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSnapshotsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsRequestTag) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequestTag) SetKey(v string) *ListSnapshotsRequestTag {
	s.Key = &v
	return s
}

func (s *ListSnapshotsRequestTag) SetValue(v string) *ListSnapshotsRequestTag {
	s.Value = &v
	return s
}

type ListSnapshotsResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the snapshots.
	Snapshots []*ListSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSnapshotsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBody) SetPageNumber(v int32) *ListSnapshotsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetPageSize(v int32) *ListSnapshotsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetRequestId(v string) *ListSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSnapshotsResponseBody) SetSnapshots(v []*ListSnapshotsResponseBodySnapshots) *ListSnapshotsResponseBody {
	s.Snapshots = v
	return s
}

func (s *ListSnapshotsResponseBody) SetTotalCount(v int32) *ListSnapshotsResponseBody {
	s.TotalCount = &v
	return s
}

type ListSnapshotsResponseBodySnapshots struct {
	// The time when the snapshot was created. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-09T07:12:49Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the simple application server.
	//
	// Note: This parameter has a value returned for only system disk snapshots.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The progress of snapshot creation.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the snapshot.
	//
	// example:
	//
	// test-Remark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the snapshot belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time when the last disk rollback was performed.
	//
	// example:
	//
	// 2021-03-09T07:12:49Z
	RollbackTime *string `json:"RollbackTime,omitempty" xml:"RollbackTime,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot.
	//
	// example:
	//
	// test-SnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The ID of the source disk. This parameter has a value even after the source disk is released.
	//
	// example:
	//
	// d-bp14wq0149cpp2xy****
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The type of the source disk. Valid values:
	//
	// 	- system: system disk.
	//
	// 	- data: data disk.
	//
	// example:
	//
	// System
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The status of the snapshot. Valid values:
	//
	// 	- Progressing: The snapshot is being created.
	//
	// 	- Accomplished: The snapshot is created.
	//
	// 	- Failed: The snapshot failed to be created.
	//
	// example:
	//
	// Accomplished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the snapshot.
	Tags []*ListSnapshotsResponseBodySnapshotsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListSnapshotsResponseBodySnapshots) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBodySnapshots) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodySnapshots) SetCreationTime(v string) *ListSnapshotsResponseBodySnapshots {
	s.CreationTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetInstanceId(v string) *ListSnapshotsResponseBodySnapshots {
	s.InstanceId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetProgress(v string) *ListSnapshotsResponseBodySnapshots {
	s.Progress = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRegionId(v string) *ListSnapshotsResponseBodySnapshots {
	s.RegionId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRemark(v string) *ListSnapshotsResponseBodySnapshots {
	s.Remark = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetResourceGroupId(v string) *ListSnapshotsResponseBodySnapshots {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetRollbackTime(v string) *ListSnapshotsResponseBodySnapshots {
	s.RollbackTime = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSnapshotName(v string) *ListSnapshotsResponseBodySnapshots {
	s.SnapshotName = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceDiskId(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceDiskId = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetSourceDiskType(v string) *ListSnapshotsResponseBodySnapshots {
	s.SourceDiskType = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetStatus(v string) *ListSnapshotsResponseBodySnapshots {
	s.Status = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshots) SetTags(v []*ListSnapshotsResponseBodySnapshotsTags) *ListSnapshotsResponseBodySnapshots {
	s.Tags = v
	return s
}

type ListSnapshotsResponseBodySnapshotsTags struct {
	// The tag key of the snapshot.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the snapshot.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSnapshotsResponseBodySnapshotsTags) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponseBodySnapshotsTags) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponseBodySnapshotsTags) SetKey(v string) *ListSnapshotsResponseBodySnapshotsTags {
	s.Key = &v
	return s
}

func (s *ListSnapshotsResponseBodySnapshotsTags) SetValue(v string) *ListSnapshotsResponseBodySnapshotsTags {
	s.Value = &v
	return s
}

type ListSnapshotsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSnapshotsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsResponse) GoString() string {
	return s.String()
}

func (s *ListSnapshotsResponse) SetHeaders(v map[string]*string) *ListSnapshotsResponse {
	s.Headers = v
	return s
}

func (s *ListSnapshotsResponse) SetStatusCode(v int32) *ListSnapshotsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSnapshotsResponse) SetBody(v *ListSnapshotsResponseBody) *ListSnapshotsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nOc1nj4M9UaAZ/I8db***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- instance
	//
	// 	- snapshot
	//
	// 	- customimage
	//
	// 	- command
	//
	// 	- firewallrule
	//
	// 	- disk
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetClientToken(v string) *ListTagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of tag N that you want to add to the simple application server. A tag key can be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the simple application server. A tag value can be up to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nOc1nj4M9UaAZ/I8db***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A collection of resource IDs and tags. The information includes the resource IDs, resource types, and key-value pairs.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The resource ID.
	//
	// example:
	//
	// s-bw526m1vi6x20c6g****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::SWAS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type LoginInstanceRequest struct {
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9ae7106e68eb4402b0dcbd48a9de****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password that corresponds to the username.
	//
	// 	- For a Linux server, you do not need to enter a password.
	//
	// 	- For a Windows server, enter the password that you set. If you have not set a password for the simple application server, set a password. For more information, see [Reset the password](https://help.aliyun.com/document_detail/60055.html).
	//
	// example:
	//
	// Test****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 3389
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username of the simple application server.
	//
	// 	- For a Linux server, you do not need to enter a username.
	//
	// 	- For a Windows server, the default username `administrator` is used.
	//
	// example:
	//
	// administrator
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s LoginInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceRequest) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequest) SetInstanceId(v string) *LoginInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceRequest) SetPassword(v string) *LoginInstanceRequest {
	s.Password = &v
	return s
}

func (s *LoginInstanceRequest) SetPort(v int32) *LoginInstanceRequest {
	s.Port = &v
	return s
}

func (s *LoginInstanceRequest) SetRegionId(v string) *LoginInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *LoginInstanceRequest) SetUsername(v string) *LoginInstanceRequest {
	s.Username = &v
	return s
}

type LoginInstanceResponseBody struct {
	// The URL that you use to log on to the server.
	//
	// example:
	//
	// https://ecs-workbench.aliyun.com/view/instance/single/gbktfz****
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C2DE174B-7196-5778-A00D-6EA2601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LoginInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponseBody) SetRedirectUrl(v string) *LoginInstanceResponseBody {
	s.RedirectUrl = &v
	return s
}

func (s *LoginInstanceResponseBody) SetRequestId(v string) *LoginInstanceResponseBody {
	s.RequestId = &v
	return s
}

type LoginInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LoginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LoginInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s LoginInstanceResponse) GoString() string {
	return s.String()
}

func (s *LoginInstanceResponse) SetHeaders(v map[string]*string) *LoginInstanceResponse {
	s.Headers = v
	return s
}

func (s *LoginInstanceResponse) SetStatusCode(v int32) *LoginInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *LoginInstanceResponse) SetBody(v *LoginInstanceResponseBody) *LoginInstanceResponse {
	s.Body = v
	return s
}

type ModifyDatabaseInstanceDescriptionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Simple Database Service instance.
	//
	// This parameter is required.
	DatabaseInstanceDescription *string `json:"DatabaseInstanceDescription,omitempty" xml:"DatabaseInstanceDescription,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetClientToken(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetDatabaseInstanceDescription(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.DatabaseInstanceDescription = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionRequest) SetRegionId(v string) *ModifyDatabaseInstanceDescriptionRequest {
	s.RegionId = &v
	return s
}

type ModifyDatabaseInstanceDescriptionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseInstanceDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDatabaseInstanceDescriptionResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseInstanceDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetStatusCode(v int32) *ModifyDatabaseInstanceDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseInstanceDescriptionResponse) SetBody(v *ModifyDatabaseInstanceDescriptionResponseBody) *ModifyDatabaseInstanceDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDatabaseInstanceParameterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// Specifies whether to forcibly restart the instance after parameters are modified. Valid values:
	//
	// 	- true: forcibly restarts the instance. If a new parameter value takes effect only after the instance restarts, you must set this parameter to true. Otherwise, the new parameter value cannot take effect.
	//
	// 	- false: does not forcibly restart the instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The JSON strings that consist of instance parameters and the values of the instance parameters. The parameter values are of the string type. Format: {"Parameter name 1":"Parameter value 1","Parameter name 2":"Parameter value 2"...}.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"delayed_insert_timeout":"600","max_length_for_sort_data":"2048"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseInstanceParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterRequest) SetClientToken(v string) *ModifyDatabaseInstanceParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetDatabaseInstanceId(v string) *ModifyDatabaseInstanceParameterRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetForceRestart(v bool) *ModifyDatabaseInstanceParameterRequest {
	s.ForceRestart = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetParameters(v string) *ModifyDatabaseInstanceParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterRequest) SetRegionId(v string) *ModifyDatabaseInstanceParameterRequest {
	s.RegionId = &v
	return s
}

type ModifyDatabaseInstanceParameterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseInstanceParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterResponseBody) SetRequestId(v string) *ModifyDatabaseInstanceParameterResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDatabaseInstanceParameterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseInstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseInstanceParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseInstanceParameterResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseInstanceParameterResponse) SetHeaders(v map[string]*string) *ModifyDatabaseInstanceParameterResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) SetStatusCode(v int32) *ModifyDatabaseInstanceParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseInstanceParameterResponse) SetBody(v *ModifyDatabaseInstanceParameterResponseBody) *ModifyDatabaseInstanceParameterResponse {
	s.Body = v
	return s
}

type ModifyFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range. Valid values: 165535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 10241055.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	//
	// example:
	//
	// custom
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// 	- TCP: the TCP protocol
	//
	// 	- UDP: the UDP protocol
	//
	// 	- TCP+UDP: the TCP and UDP protocols
	//
	// This parameter is required.
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The IP address or CIDR block that is allowed in the firewall rule.
	//
	// example:
	//
	// 10.147.33.**
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s ModifyFirewallRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRuleRequest) SetClientToken(v string) *ModifyFirewallRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetInstanceId(v string) *ModifyFirewallRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetPort(v string) *ModifyFirewallRuleRequest {
	s.Port = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetRegionId(v string) *ModifyFirewallRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetRemark(v string) *ModifyFirewallRuleRequest {
	s.Remark = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetRuleId(v string) *ModifyFirewallRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetRuleProtocol(v string) *ModifyFirewallRuleRequest {
	s.RuleProtocol = &v
	return s
}

func (s *ModifyFirewallRuleRequest) SetSourceCidrIp(v string) *ModifyFirewallRuleRequest {
	s.SourceCidrIp = &v
	return s
}

type ModifyFirewallRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E1FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFirewallRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRuleResponseBody) SetRequestId(v string) *ModifyFirewallRuleResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFirewallRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallRuleResponse) SetHeaders(v map[string]*string) *ModifyFirewallRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallRuleResponse) SetStatusCode(v int32) *ModifyFirewallRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallRuleResponse) SetBody(v *ModifyFirewallRuleResponseBody) *ModifyFirewallRuleResponse {
	s.Body = v
	return s
}

type ModifyFirewallTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the firewall template.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The firewall rule in the template.
	FirewallTemplateRule []*ModifyFirewallTemplateRequestFirewallTemplateRule `json:"FirewallTemplateRule,omitempty" xml:"FirewallTemplateRule,omitempty" type:"Repeated"`
	// The ID of the simple application server to which the firewall template is applied.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the firewall template.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyFirewallTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateRequest) SetClientToken(v string) *ModifyFirewallTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetDescription(v string) *ModifyFirewallTemplateRequest {
	s.Description = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetFirewallTemplateId(v string) *ModifyFirewallTemplateRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetFirewallTemplateRule(v []*ModifyFirewallTemplateRequestFirewallTemplateRule) *ModifyFirewallTemplateRequest {
	s.FirewallTemplateRule = v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetInstanceId(v string) *ModifyFirewallTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetName(v string) *ModifyFirewallTemplateRequest {
	s.Name = &v
	return s
}

func (s *ModifyFirewallTemplateRequest) SetRegionId(v string) *ModifyFirewallTemplateRequest {
	s.RegionId = &v
	return s
}

type ModifyFirewallTemplateRequestFirewallTemplateRule struct {
	// The ID of the firewall rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeea34d9867b4d55a4ff8d5fcfbd****
	FirewallTemplateRuleId *string `json:"FirewallTemplateRuleId,omitempty" xml:"FirewallTemplateRuleId,omitempty"`
	// The port range. Valid values: 1 to 65535. Specify a port range in the format of \\<start port number>/\\<end port number>. Example: `1024/1055`, which indicates that the port range of 1024 to 1055.
	//
	// >  If you set RuleProtocol to ICMP, you must set Port to -1/-1.
	//
	// example:
	//
	// 8080
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall template rule.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol that the rule supports. Valid values:
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- TCP+UDP
	//
	// 	- ICMP
	//
	// example:
	//
	// TCP
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The source address to which you want to grant access permissions. CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
}

func (s ModifyFirewallTemplateRequestFirewallTemplateRule) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallTemplateRequestFirewallTemplateRule) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetFirewallTemplateRuleId(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.FirewallTemplateRuleId = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetPort(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.Port = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetRemark(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.Remark = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetRuleProtocol(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.RuleProtocol = &v
	return s
}

func (s *ModifyFirewallTemplateRequestFirewallTemplateRule) SetSourceCidrIp(v string) *ModifyFirewallTemplateRequestFirewallTemplateRule {
	s.SourceCidrIp = &v
	return s
}

type ModifyFirewallTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyFirewallTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateResponseBody) SetRequestId(v string) *ModifyFirewallTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ModifyFirewallTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirewallTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateResponse) SetHeaders(v map[string]*string) *ModifyFirewallTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallTemplateResponse) SetStatusCode(v int32) *ModifyFirewallTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallTemplateResponse) SetBody(v *ModifyFirewallTemplateResponseBody) *ModifyFirewallTemplateResponse {
	s.Body = v
	return s
}

type ModifyImageShareStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-saacssasc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Valid values:
	//
	// 	- Share
	//
	// 	- UnShare
	//
	// This parameter is required.
	//
	// example:
	//
	// Share
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The region ID of the custom image. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyImageShareStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageShareStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusRequest) SetClientToken(v string) *ModifyImageShareStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetImageId(v string) *ModifyImageShareStatusRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetOperation(v string) *ModifyImageShareStatusRequest {
	s.Operation = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetRegionId(v string) *ModifyImageShareStatusRequest {
	s.RegionId = &v
	return s
}

type ModifyImageShareStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageShareStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageShareStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusResponseBody) SetRequestId(v string) *ModifyImageShareStatusResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageShareStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyImageShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyImageShareStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageShareStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusResponse) SetHeaders(v map[string]*string) *ModifyImageShareStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageShareStatusResponse) SetStatusCode(v int32) *ModifyImageShareStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyImageShareStatusResponse) SetBody(v *ModifyImageShareStatusResponseBody) *ModifyImageShareStatusResponse {
	s.Body = v
	return s
}

type ModifyInstanceVncPasswordRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The existing VNC password.
	//
	// example:
	//
	// ***
	VncPassword *string `json:"VncPassword,omitempty" xml:"VncPassword,omitempty"`
}

func (s ModifyInstanceVncPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceVncPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswordRequest) SetClientToken(v string) *ModifyInstanceVncPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceVncPasswordRequest) SetInstanceId(v string) *ModifyInstanceVncPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVncPasswordRequest) SetRegionId(v string) *ModifyInstanceVncPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceVncPasswordRequest) SetVncPassword(v string) *ModifyInstanceVncPasswordRequest {
	s.VncPassword = &v
	return s
}

type ModifyInstanceVncPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceVncPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceVncPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswordResponseBody) SetRequestId(v string) *ModifyInstanceVncPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceVncPasswordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceVncPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceVncPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceVncPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVncPasswordResponse) SetHeaders(v map[string]*string) *ModifyInstanceVncPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceVncPasswordResponse) SetStatusCode(v int32) *ModifyInstanceVncPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceVncPasswordResponse) SetBody(v *ModifyInstanceVncPasswordResponseBody) *ModifyInstanceVncPasswordResponse {
	s.Body = v
	return s
}

type RebootInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) SetClientToken(v string) *RebootInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetRegionId(v string) *RebootInstanceRequest {
	s.RegionId = &v
	return s
}

type RebootInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponseBody) SetRequestId(v string) *RebootInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RebootInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootInstanceResponse) GoString() string {
	return s.String()
}

func (s *RebootInstanceResponse) SetHeaders(v map[string]*string) *RebootInstanceResponse {
	s.Headers = v
	return s
}

func (s *RebootInstanceResponse) SetStatusCode(v int32) *RebootInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootInstanceResponse) SetBody(v *RebootInstanceResponseBody) *RebootInstanceResponse {
	s.Body = v
	return s
}

type RebootInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly restart the servers. Valid values:
	//
	// 	- true: forcibly restarts the servers. This operation is equivalent to the typical power-off operation. Cache data that is not written to storage devices on the server will be lost.
	//
	// 	- false: normally restarts the instance.
	//
	// Default value: false
	//
	// example:
	//
	// false
	ForceReboot *bool `json:"ForceReboot,omitempty" xml:"ForceReboot,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebootInstancesRequest) SetClientToken(v string) *RebootInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootInstancesRequest) SetForceReboot(v bool) *RebootInstancesRequest {
	s.ForceReboot = &v
	return s
}

func (s *RebootInstancesRequest) SetInstanceIds(v string) *RebootInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *RebootInstancesRequest) SetRegionId(v string) *RebootInstancesRequest {
	s.RegionId = &v
	return s
}

type RebootInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBody) SetRequestId(v string) *RebootInstancesResponseBody {
	s.RequestId = &v
	return s
}

type RebootInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootInstancesResponse) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponse) SetHeaders(v map[string]*string) *RebootInstancesResponse {
	s.Headers = v
	return s
}

func (s *RebootInstancesResponse) SetStatusCode(v int32) *RebootInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootInstancesResponse) SetBody(v *RebootInstancesResponseBody) *RebootInstancesResponse {
	s.Body = v
	return s
}

type ReleasePublicConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleasePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionRequest) SetClientToken(v string) *ReleasePublicConnectionRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleasePublicConnectionRequest) SetDatabaseInstanceId(v string) *ReleasePublicConnectionRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ReleasePublicConnectionRequest) SetRegionId(v string) *ReleasePublicConnectionRequest {
	s.RegionId = &v
	return s
}

type ReleasePublicConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleasePublicConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionResponseBody) SetRequestId(v string) *ReleasePublicConnectionResponseBody {
	s.RequestId = &v
	return s
}

type ReleasePublicConnectionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleasePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleasePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleasePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleasePublicConnectionResponse) SetHeaders(v map[string]*string) *ReleasePublicConnectionResponse {
	s.Headers = v
	return s
}

func (s *ReleasePublicConnectionResponse) SetStatusCode(v int32) *ReleasePublicConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleasePublicConnectionResponse) SetBody(v *ReleasePublicConnectionResponseBody) *ReleasePublicConnectionResponse {
	s.Body = v
	return s
}

type RemoveCustomImageShareAccountRequest struct {
	// This parameter is required.
	Account []*int64 `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// m-2zehv38jjmwva1ee****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveCustomImageShareAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveCustomImageShareAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveCustomImageShareAccountRequest) SetAccount(v []*int64) *RemoveCustomImageShareAccountRequest {
	s.Account = v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) SetClientToken(v string) *RemoveCustomImageShareAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) SetImageId(v string) *RemoveCustomImageShareAccountRequest {
	s.ImageId = &v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) SetRegionId(v string) *RemoveCustomImageShareAccountRequest {
	s.RegionId = &v
	return s
}

type RemoveCustomImageShareAccountResponseBody struct {
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCustomImageShareAccountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveCustomImageShareAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCustomImageShareAccountResponseBody) SetRequestId(v string) *RemoveCustomImageShareAccountResponseBody {
	s.RequestId = &v
	return s
}

type RemoveCustomImageShareAccountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveCustomImageShareAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveCustomImageShareAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveCustomImageShareAccountResponse) GoString() string {
	return s.String()
}

func (s *RemoveCustomImageShareAccountResponse) SetHeaders(v map[string]*string) *RemoveCustomImageShareAccountResponse {
	s.Headers = v
	return s
}

func (s *RemoveCustomImageShareAccountResponse) SetStatusCode(v int32) *RemoveCustomImageShareAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveCustomImageShareAccountResponse) SetBody(v *RemoveCustomImageShareAccountResponseBody) *RemoveCustomImageShareAccountResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The renewal period. Unit: month. Valid values: 1, 3, 6, 12, 24, and 36.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetClientToken(v string) *RenewInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstanceRequest) SetInstanceId(v string) *RenewInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v int32) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetRegionId(v string) *RenewInstanceRequest {
	s.RegionId = &v
	return s
}

type RenewInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetStatusCode(v int32) *RenewInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ResetDatabaseAccountPasswordRequest struct {
	// The password of the database administrator account.
	//
	// This parameter is required.
	//
	// example:
	//
	// Password****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetDatabaseAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDatabaseAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordRequest) SetAccountPassword(v string) *ResetDatabaseAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetClientToken(v string) *ResetDatabaseAccountPasswordRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetDatabaseInstanceId(v string) *ResetDatabaseAccountPasswordRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *ResetDatabaseAccountPasswordRequest) SetRegionId(v string) *ResetDatabaseAccountPasswordRequest {
	s.RegionId = &v
	return s
}

type ResetDatabaseAccountPasswordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDatabaseAccountPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDatabaseAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordResponseBody) SetRequestId(v string) *ResetDatabaseAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetDatabaseAccountPasswordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDatabaseAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDatabaseAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDatabaseAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetDatabaseAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) SetStatusCode(v int32) *ResetDatabaseAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) SetBody(v *ResetDatabaseAccountPasswordResponseBody) *ResetDatabaseAccountPasswordResponse {
	s.Body = v
	return s
}

type ResetDiskRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk to be rolled back.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp14wq0149cpp2xy****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server for which the snapshot is created.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetDiskRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDiskRequest) GoString() string {
	return s.String()
}

func (s *ResetDiskRequest) SetClientToken(v string) *ResetDiskRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetDiskRequest) SetDiskId(v string) *ResetDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ResetDiskRequest) SetRegionId(v string) *ResetDiskRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDiskRequest) SetSnapshotId(v string) *ResetDiskRequest {
	s.SnapshotId = &v
	return s
}

type ResetDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetDiskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDiskResponseBody) SetRequestId(v string) *ResetDiskResponseBody {
	s.RequestId = &v
	return s
}

type ResetDiskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDiskResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDiskResponse) GoString() string {
	return s.String()
}

func (s *ResetDiskResponse) SetHeaders(v map[string]*string) *ResetDiskResponse {
	s.Headers = v
	return s
}

func (s *ResetDiskResponse) SetStatusCode(v int32) *ResetDiskResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDiskResponse) SetBody(v *ResetDiskResponseBody) *ResetDiskResponse {
	s.Body = v
	return s
}

type ResetSystemRequest struct {
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the destination image. If you do not specify this parameter, the current image is reset.
	//
	// example:
	//
	// 794c230fd3e64ea19f83f4d7a0ad****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetSystemRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSystemRequest) GoString() string {
	return s.String()
}

func (s *ResetSystemRequest) SetClientToken(v string) *ResetSystemRequest {
	s.ClientToken = &v
	return s
}

func (s *ResetSystemRequest) SetImageId(v string) *ResetSystemRequest {
	s.ImageId = &v
	return s
}

func (s *ResetSystemRequest) SetInstanceId(v string) *ResetSystemRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetSystemRequest) SetRegionId(v string) *ResetSystemRequest {
	s.RegionId = &v
	return s
}

type ResetSystemResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSystemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSystemResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSystemResponseBody) SetRequestId(v string) *ResetSystemResponseBody {
	s.RequestId = &v
	return s
}

type ResetSystemResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSystemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSystemResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSystemResponse) GoString() string {
	return s.String()
}

func (s *ResetSystemResponse) SetHeaders(v map[string]*string) *ResetSystemResponse {
	s.Headers = v
	return s
}

func (s *ResetSystemResponse) SetStatusCode(v int32) *ResetSystemResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSystemResponse) SetBody(v *ResetSystemResponseBody) *ResetSystemResponse {
	s.Body = v
	return s
}

type RestartDatabaseInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartDatabaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceRequest) SetClientToken(v string) *RestartDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *RestartDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *RestartDatabaseInstanceRequest) SetRegionId(v string) *RestartDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

type RestartDatabaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDatabaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceResponseBody) SetRequestId(v string) *RestartDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartDatabaseInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartDatabaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDatabaseInstanceResponse) SetHeaders(v map[string]*string) *RestartDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartDatabaseInstanceResponse) SetStatusCode(v int32) *RestartDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartDatabaseInstanceResponse) SetBody(v *RestartDatabaseInstanceResponseBody) *RestartDatabaseInstanceResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	// The content of the command. Take note of the following items:
	//
	// 	- If you set `EnableParameter` to true, the custom parameter feature is enabled in the command content and you can configure custom parameters based on the following rules:
	//
	// 	- Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	// 	- The number of custom parameters cannot be greater than 20.
	//
	// 	- A custom parameter name can contain only letters, digits, underscores (_), and hyphens (-). The name is case-insensitive.
	//
	// 	- Each custom parameter name cannot exceed 64 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// ifconfig -s
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Specifies whether to enable the custom parameter feature.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the command.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom parameters in the key-value pair format that are to be passed in when the command includes custom parameters. For example, if the command content is `echo {{name}}`, you can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced with the paired Jack value to generate a new command. As a result, the `echo Jack` command is executed.
	//
	// Number of custom parameters ranges from 0 to 20. Take note of the following items:
	//
	// 	- The key cannot be an empty string. It can be up to 64 characters in length.
	//
	// 	- The value can be an empty string.
	//
	// 	- After custom parameters and original command content are encoded in Base64, the command cannot exceed 16 KB in size.
	//
	// 	- The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is empty by default, which indicates to disable the custom parameter feature.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAIdyvdIqaRY****"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of the command on the server.
	//
	// If a command execution task times out, Command Assistant forcibly terminates the task process. Valid values: 10 to 86400. Unit: seconds. The period of 86400 seconds is equal to 24 hours.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// 	- RunBatScript: batch commands (applicable to Windows servers).
	//
	// 	- RunPowerShellScript: PowerShell commands (applicable to Windows servers).
	//
	// 	- RunShellScript: shell commands (applicable to Linux servers).
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the password to be used to run the command on a Windows server.
	//
	// If you want to use a username other than the default "system" username to run the command on a Windows server, you must specify both the WindowsPasswordName and WorkingUser parameters. To mitigate the risk of password leaks, the password is stored in plaintext in Operation Orchestration Service (OOS) Parameter Store, and only the name of the password is passed in by using WindowsPasswordName.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The execution path of the command. Custom paths are supported. Default execution paths vary based on the operating systems of the servers.
	//
	// 	- For Linux servers, the default path is /root of the root user.
	//
	// 	- For Windows servers, the default path is C:\\Windows\\system32.
	//
	// example:
	//
	// /home/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// A user of the server who runs the command. We recommend that you run the command as a regular user to reduce security risks. Default values:
	//
	// 	- For Linux servers, the default value is root.
	//
	// 	- For Windows servers, the default value is system.
	//
	// example:
	//
	// root
	WorkingUser *string `json:"WorkingUser,omitempty" xml:"WorkingUser,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetEnableParameter(v bool) *RunCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandRequest) SetInstanceId(v string) *RunCommandRequest {
	s.InstanceId = &v
	return s
}

func (s *RunCommandRequest) SetName(v string) *RunCommandRequest {
	s.Name = &v
	return s
}

func (s *RunCommandRequest) SetParameters(v map[string]interface{}) *RunCommandRequest {
	s.Parameters = v
	return s
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int32) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetType(v string) *RunCommandRequest {
	s.Type = &v
	return s
}

func (s *RunCommandRequest) SetWindowsPasswordName(v string) *RunCommandRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandRequest) SetWorkingDir(v string) *RunCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandRequest) SetWorkingUser(v string) *RunCommandRequest {
	s.WorkingUser = &v
	return s
}

type RunCommandShrinkRequest struct {
	// The content of the command. Take note of the following items:
	//
	// 	- If you set `EnableParameter` to true, the custom parameter feature is enabled in the command content and you can configure custom parameters based on the following rules:
	//
	// 	- Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	// 	- The number of custom parameters cannot be greater than 20.
	//
	// 	- A custom parameter name can contain only letters, digits, underscores (_), and hyphens (-). The name is case-insensitive.
	//
	// 	- Each custom parameter name cannot exceed 64 bytes in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// ifconfig -s
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Specifies whether to enable the custom parameter feature.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the command.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom parameters in the key-value pair format that are to be passed in when the command includes custom parameters. For example, if the command content is `echo {{name}}`, you can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced with the paired Jack value to generate a new command. As a result, the `echo Jack` command is executed.
	//
	// Number of custom parameters ranges from 0 to 20. Take note of the following items:
	//
	// 	- The key cannot be an empty string. It can be up to 64 characters in length.
	//
	// 	- The value can be an empty string.
	//
	// 	- After custom parameters and original command content are encoded in Base64, the command cannot exceed 16 KB in size.
	//
	// 	- The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is empty by default, which indicates to disable the custom parameter feature.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAIdyvdIqaRY****"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of the command on the server.
	//
	// If a command execution task times out, Command Assistant forcibly terminates the task process. Valid values: 10 to 86400. Unit: seconds. The period of 86400 seconds is equal to 24 hours.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// 	- RunBatScript: batch commands (applicable to Windows servers).
	//
	// 	- RunPowerShellScript: PowerShell commands (applicable to Windows servers).
	//
	// 	- RunShellScript: shell commands (applicable to Linux servers).
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the password to be used to run the command on a Windows server.
	//
	// If you want to use a username other than the default "system" username to run the command on a Windows server, you must specify both the WindowsPasswordName and WorkingUser parameters. To mitigate the risk of password leaks, the password is stored in plaintext in Operation Orchestration Service (OOS) Parameter Store, and only the name of the password is passed in by using WindowsPasswordName.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The execution path of the command. Custom paths are supported. Default execution paths vary based on the operating systems of the servers.
	//
	// 	- For Linux servers, the default path is /root of the root user.
	//
	// 	- For Windows servers, the default path is C:\\Windows\\system32.
	//
	// example:
	//
	// /home/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// A user of the server who runs the command. We recommend that you run the command as a regular user to reduce security risks. Default values:
	//
	// 	- For Linux servers, the default value is root.
	//
	// 	- For Windows servers, the default value is system.
	//
	// example:
	//
	// root
	WorkingUser *string `json:"WorkingUser,omitempty" xml:"WorkingUser,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetInstanceId(v string) *RunCommandShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetParametersShrink(v string) *RunCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRegionId(v string) *RunCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int32) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetType(v string) *RunCommandShrinkRequest {
	s.Type = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWindowsPasswordName(v string) *RunCommandShrinkRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingDir(v string) *RunCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingUser(v string) *RunCommandShrinkRequest {
	s.WorkingUser = &v
	return s
}

type RunCommandResponseBody struct {
	// The execution ID.
	//
	// example:
	//
	// t-hz02p9545t6****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type StartDatabaseInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDatabaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceRequest) SetClientToken(v string) *StartDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *StartDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *StartDatabaseInstanceRequest) SetRegionId(v string) *StartDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

type StartDatabaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDatabaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceResponseBody) SetRequestId(v string) *StartDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartDatabaseInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDatabaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartDatabaseInstanceResponse) SetHeaders(v map[string]*string) *StartDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartDatabaseInstanceResponse) SetStatusCode(v int32) *StartDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDatabaseInstanceResponse) SetBody(v *StartDatabaseInstanceResponseBody) *StartDatabaseInstanceResponse {
	s.Body = v
	return s
}

type StartInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) SetClientToken(v string) *StartInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) SetRegionId(v string) *StartInstanceRequest {
	s.RegionId = &v
	return s
}

type StartInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartInstanceResponse) SetHeaders(v map[string]*string) *StartInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartInstanceResponse) SetStatusCode(v int32) *StartInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstanceResponse) SetBody(v *StartInstanceResponseBody) *StartInstanceResponse {
	s.Body = v
	return s
}

type StartInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartInstancesRequest) SetClientToken(v string) *StartInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *StartInstancesRequest) SetInstanceIds(v string) *StartInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *StartInstancesRequest) SetRegionId(v string) *StartInstancesRequest {
	s.RegionId = &v
	return s
}

type StartInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBody) SetRequestId(v string) *StartInstancesResponseBody {
	s.RequestId = &v
	return s
}

type StartInstancesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstancesResponse) GoString() string {
	return s.String()
}

func (s *StartInstancesResponse) SetHeaders(v map[string]*string) *StartInstancesResponse {
	s.Headers = v
	return s
}

func (s *StartInstancesResponse) SetStatusCode(v int32) *StartInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstancesResponse) SetBody(v *StartInstancesResponseBody) *StartInstancesResponse {
	s.Body = v
	return s
}

type StartTerminalSessionRequest struct {
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartTerminalSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTerminalSessionRequest) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionRequest) SetInstanceId(v string) *StartTerminalSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetRegionId(v string) *StartTerminalSessionRequest {
	s.RegionId = &v
	return s
}

type StartTerminalSessionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	//
	// example:
	//
	// token-xxxaaz
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The session ID.
	//
	// example:
	//
	// ffb90b6e-b18a-4a33-88cf-86fb88****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The URL of the WebSocket session that is used to connect to the server. The URL contains the session ID (`SessionId`) and the authentication token (`SecurityToken`).
	//
	// example:
	//
	// wss://xxxx
	WebSocketUrl *string `json:"WebSocketUrl,omitempty" xml:"WebSocketUrl,omitempty"`
}

func (s StartTerminalSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTerminalSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionResponseBody) SetRequestId(v string) *StartTerminalSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetSecurityToken(v string) *StartTerminalSessionResponseBody {
	s.SecurityToken = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetSessionId(v string) *StartTerminalSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *StartTerminalSessionResponseBody) SetWebSocketUrl(v string) *StartTerminalSessionResponseBody {
	s.WebSocketUrl = &v
	return s
}

type StartTerminalSessionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTerminalSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTerminalSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTerminalSessionResponse) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionResponse) SetHeaders(v map[string]*string) *StartTerminalSessionResponse {
	s.Headers = v
	return s
}

func (s *StartTerminalSessionResponse) SetStatusCode(v int32) *StartTerminalSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTerminalSessionResponse) SetBody(v *StartTerminalSessionResponseBody) *StartTerminalSessionResponse {
	s.Body = v
	return s
}

type StopDatabaseInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// db-38263fa955774501a2ae1bdaed6f****
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopDatabaseInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDatabaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceRequest) SetClientToken(v string) *StopDatabaseInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDatabaseInstanceRequest) SetDatabaseInstanceId(v string) *StopDatabaseInstanceRequest {
	s.DatabaseInstanceId = &v
	return s
}

func (s *StopDatabaseInstanceRequest) SetRegionId(v string) *StopDatabaseInstanceRequest {
	s.RegionId = &v
	return s
}

type StopDatabaseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDatabaseInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDatabaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceResponseBody) SetRequestId(v string) *StopDatabaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopDatabaseInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDatabaseInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDatabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopDatabaseInstanceResponse) SetHeaders(v map[string]*string) *StopDatabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopDatabaseInstanceResponse) SetStatusCode(v int32) *StopDatabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDatabaseInstanceResponse) SetBody(v *StopDatabaseInstanceResponseBody) *StopDatabaseInstanceResponse {
	s.Body = v
	return s
}

type StopInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) SetClientToken(v string) *StopInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetRegionId(v string) *StopInstanceRequest {
	s.RegionId = &v
	return s
}

type StopInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstanceResponseBody) SetRequestId(v string) *StopInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopInstanceResponse) SetHeaders(v map[string]*string) *StopInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopInstanceResponse) SetStatusCode(v int32) *StopInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstanceResponse) SetBody(v *StopInstanceResponseBody) *StopInstanceResponse {
	s.Body = v
	return s
}

type StopInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly stop the servers.
	//
	// 	- **true**: forcibly stops the servers.
	//
	// 	- **false**: normally stops the servers. This is the default value.
	//
	// example:
	//
	// true
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesRequest) SetClientToken(v string) *StopInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *StopInstancesRequest) SetForceStop(v bool) *StopInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstancesRequest) SetInstanceIds(v string) *StopInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *StopInstancesRequest) SetRegionId(v string) *StopInstancesRequest {
	s.RegionId = &v
	return s
}

type StopInstancesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C2DE174B-7196-5778-A00D-6EA2601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBody) SetRequestId(v string) *StopInstancesResponseBody {
	s.RequestId = &v
	return s
}

type StopInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstancesResponse) GoString() string {
	return s.String()
}

func (s *StopInstancesResponse) SetHeaders(v map[string]*string) *StopInstancesResponse {
	s.Headers = v
	return s
}

func (s *StopInstancesResponse) SetStatusCode(v int32) *StopInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstancesResponse) SetBody(v *StopInstancesResponseBody) *StopInstancesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- instance
	//
	// 	- snapshot
	//
	// 	- customimage
	//
	// 	- command
	//
	// 	- firewallrule
	//
	// 	- disk
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. You can specify up to 20 tags.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetClientToken(v string) *TagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of tag N that you want to add to the simple application server. Valid values of N: 1 to 20.
	//
	// You cannot specify an empty string as a tag key. A tag key can be up to 64 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the simple application server. Valid values of N: 1 to 20.
	//
	// You can specify an empty string as a tag value. A tag value can be up to 64 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags. This parameter takes effect only when TagKey.N is not specified. Valid values: true and false. Default value: false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify up to 50 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid values:
	//
	// 	- instance
	//
	// 	- snapshot
	//
	// 	- customimage
	//
	// 	- command
	//
	// 	- firewallrule
	//
	// 	- disk
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N. Valid values of N: 1 to 20.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetClientToken(v string) *UntagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateCommandAttributeRequest struct {
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The description of the command. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the command. The name supports all character sets and can be up to 128 characters in length.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum timeout period for the command execution on the ECS instance. Unit: seconds. When a command that you created cannot be run, the command execution times out. When the execution times out, the command process is forcefully terminated and the PID of the command is canceled. Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The working directory of the command. The directory can be up to 200 characters in length.
	//
	// example:
	//
	// /home/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s UpdateCommandAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommandAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateCommandAttributeRequest) SetCommandId(v string) *UpdateCommandAttributeRequest {
	s.CommandId = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetDescription(v string) *UpdateCommandAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetName(v string) *UpdateCommandAttributeRequest {
	s.Name = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetRegionId(v string) *UpdateCommandAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetTimeout(v int64) *UpdateCommandAttributeRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateCommandAttributeRequest) SetWorkingDir(v string) *UpdateCommandAttributeRequest {
	s.WorkingDir = &v
	return s
}

type UpdateCommandAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCommandAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommandAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCommandAttributeResponseBody) SetRequestId(v string) *UpdateCommandAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCommandAttributeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCommandAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCommandAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCommandAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCommandAttributeResponse) SetHeaders(v map[string]*string) *UpdateCommandAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCommandAttributeResponse) SetStatusCode(v int32) *UpdateCommandAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCommandAttributeResponse) SetBody(v *UpdateCommandAttributeResponseBody) *UpdateCommandAttributeResponse {
	s.Body = v
	return s
}

type UpdateDiskAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk ID. You can call the ListDisks operation to query the ID of data disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp18kjxg9ebrhsgi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the data disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateDiskAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDiskAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateDiskAttributeRequest) SetClientToken(v string) *UpdateDiskAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateDiskAttributeRequest) SetDiskId(v string) *UpdateDiskAttributeRequest {
	s.DiskId = &v
	return s
}

func (s *UpdateDiskAttributeRequest) SetRegionId(v string) *UpdateDiskAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDiskAttributeRequest) SetRemark(v string) *UpdateDiskAttributeRequest {
	s.Remark = &v
	return s
}

type UpdateDiskAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28D****534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDiskAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDiskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDiskAttributeResponseBody) SetRequestId(v string) *UpdateDiskAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDiskAttributeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDiskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDiskAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDiskAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateDiskAttributeResponse) SetHeaders(v map[string]*string) *UpdateDiskAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateDiskAttributeResponse) SetStatusCode(v int32) *UpdateDiskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDiskAttributeResponse) SetBody(v *UpdateDiskAttributeResponseBody) *UpdateDiskAttributeResponse {
	s.Body = v
	return s
}

type UpdateInstanceAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`. The name can only contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The new password of the simple application server. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The password can contain the following special characters:
	//
	//     ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// >  For security reasons, we recommend that you use HTTPS to send requests if the `Password parameter` is specified.
	//
	// example:
	//
	// Test123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeRequest) SetClientToken(v string) *UpdateInstanceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceId(v string) *UpdateInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceName(v string) *UpdateInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetPassword(v string) *UpdateInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetRegionId(v string) *UpdateInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

type UpdateInstanceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponseBody) SetRequestId(v string) *UpdateInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetStatusCode(v int32) *UpdateInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetBody(v *UpdateInstanceAttributeResponseBody) *UpdateInstanceAttributeResponse {
	s.Body = v
	return s
}

type UpdateSnapshotAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the snapshot of the simple application server.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The snapshot ID. You can call the ListSnapshots operation to query the snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp16oazlsold4dks****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s UpdateSnapshotAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotAttributeRequest) SetClientToken(v string) *UpdateSnapshotAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) SetRegionId(v string) *UpdateSnapshotAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) SetRemark(v string) *UpdateSnapshotAttributeRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSnapshotAttributeRequest) SetSnapshotId(v string) *UpdateSnapshotAttributeRequest {
	s.SnapshotId = &v
	return s
}

type UpdateSnapshotAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSnapshotAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotAttributeResponseBody) SetRequestId(v string) *UpdateSnapshotAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSnapshotAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSnapshotAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSnapshotAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSnapshotAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSnapshotAttributeResponse) SetHeaders(v map[string]*string) *UpdateSnapshotAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSnapshotAttributeResponse) SetStatusCode(v int32) *UpdateSnapshotAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSnapshotAttributeResponse) SetBody(v *UpdateSnapshotAttributeResponseBody) *UpdateSnapshotAttributeResponse {
	s.Body = v
	return s
}

type UpgradeInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new plan. You can call the [ListPlans](https://help.aliyun.com/document_detail/189314.html) operation to query the plans provided by Simple Application Server.
	//
	// This parameter is required.
	//
	// example:
	//
	// swas.s2.c2m2s50b4t08
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceRequest) SetClientToken(v string) *UpgradeInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeInstanceRequest) SetInstanceId(v string) *UpgradeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetPlanId(v string) *UpgradeInstanceRequest {
	s.PlanId = &v
	return s
}

func (s *UpgradeInstanceRequest) SetRegionId(v string) *UpgradeInstanceRequest {
	s.RegionId = &v
	return s
}

type UpgradeInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponseBody) SetRequestId(v string) *UpgradeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeInstanceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceResponse) SetHeaders(v map[string]*string) *UpgradeInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeInstanceResponse) SetStatusCode(v int32) *UpgradeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeInstanceResponse) SetBody(v *UpgradeInstanceResponseBody) *UpgradeInstanceResponse {
	s.Body = v
	return s
}

type UploadInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test_gin
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key.
	//
	// example:
	//
	// ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAID5aQ5bM0Am3mWe+upjSXqisUT4DLR6ExwvA0****	- **@**.com
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UploadInstanceKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadInstanceKeyPairRequest) GoString() string {
	return s.String()
}

func (s *UploadInstanceKeyPairRequest) SetClientToken(v string) *UploadInstanceKeyPairRequest {
	s.ClientToken = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetInstanceId(v string) *UploadInstanceKeyPairRequest {
	s.InstanceId = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetKeyPairName(v string) *UploadInstanceKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetPublicKey(v string) *UploadInstanceKeyPairRequest {
	s.PublicKey = &v
	return s
}

func (s *UploadInstanceKeyPairRequest) SetRegionId(v string) *UploadInstanceKeyPairRequest {
	s.RegionId = &v
	return s
}

type UploadInstanceKeyPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadInstanceKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *UploadInstanceKeyPairResponseBody) SetRequestId(v string) *UploadInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type UploadInstanceKeyPairResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadInstanceKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadInstanceKeyPairResponse) GoString() string {
	return s.String()
}

func (s *UploadInstanceKeyPairResponse) SetHeaders(v map[string]*string) *UploadInstanceKeyPairResponse {
	s.Headers = v
	return s
}

func (s *UploadInstanceKeyPairResponse) SetStatusCode(v int32) *UploadInstanceKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadInstanceKeyPairResponse) SetBody(v *UploadInstanceKeyPairResponseBody) *UploadInstanceKeyPairResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("swas-open"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 共享镜像给用户
//
// @param request - AddCustomImageShareAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddCustomImageShareAccountResponse
func (client *Client) AddCustomImageShareAccountWithOptions(request *AddCustomImageShareAccountRequest, runtime *util.RuntimeOptions) (_result *AddCustomImageShareAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddCustomImageShareAccount"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddCustomImageShareAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 共享镜像给用户
//
// @param request - AddCustomImageShareAccountRequest
//
// @return AddCustomImageShareAccountResponse
func (client *Client) AddCustomImageShareAccount(request *AddCustomImageShareAccountRequest) (_result *AddCustomImageShareAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddCustomImageShareAccountResponse{}
	_body, _err := client.AddCustomImageShareAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint for a Simple Database Service instance.
//
// Description:
//
// By default, no public endpoints are assigned to Simple Database Service instances. If you want to access the databases of a Simple Database Service instance over the Internet by using Simple Container Service or Data Management (DMS), you must apply for a public endpoint for the Simple Database Service instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - AllocatePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AllocatePublicConnectionResponse
func (client *Client) AllocatePublicConnectionWithOptions(request *AllocatePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocatePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AllocatePublicConnection"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AllocatePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Applies for a public endpoint for a Simple Database Service instance.
//
// Description:
//
// By default, no public endpoints are assigned to Simple Database Service instances. If you want to access the databases of a Simple Database Service instance over the Internet by using Simple Container Service or Data Management (DMS), you must apply for a public endpoint for the Simple Database Service instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - AllocatePublicConnectionRequest
//
// @return AllocatePublicConnectionResponse
func (client *Client) AllocatePublicConnection(request *AllocatePublicConnectionRequest) (_result *AllocatePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocatePublicConnectionResponse{}
	_body, _err := client.AllocatePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses a firewall template to apply firewall rules to multiple simple application servers at a time. This improves your efficiency of setting firewall rules.
//
// Description:
//
// If the port range, protocol, and source IP address of a firewall rule in a firewall template are the same as the port range, protocol, and source IP address of an existing rule, the new rule overwrites the existing rule regardless of whether the existing rule is enabled or disabled.
//
// @param request - ApplyFirewallTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApplyFirewallTemplateResponse
func (client *Client) ApplyFirewallTemplateWithOptions(request *ApplyFirewallTemplateRequest, runtime *util.RuntimeOptions) (_result *ApplyFirewallTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ApplyFirewallTemplate"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApplyFirewallTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses a firewall template to apply firewall rules to multiple simple application servers at a time. This improves your efficiency of setting firewall rules.
//
// Description:
//
// If the port range, protocol, and source IP address of a firewall rule in a firewall template are the same as the port range, protocol, and source IP address of an existing rule, the new rule overwrites the existing rule regardless of whether the existing rule is enabled or disabled.
//
// @param request - ApplyFirewallTemplateRequest
//
// @return ApplyFirewallTemplateResponse
func (client *Client) ApplyFirewallTemplate(request *ApplyFirewallTemplateRequest) (_result *ApplyFirewallTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyFirewallTemplateResponse{}
	_body, _err := client.ApplyFirewallTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a key pair to simple application servers.
//
// Description:
//
// You can bind only one key pair to a simple application server in the Simple Application Server console. If a simple application server has a key pair bound, the new key pair overwrites the original key pair.
//
// @param request - AttachKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachKeyPairResponse
func (client *Client) AttachKeyPairWithOptions(request *AttachKeyPairRequest, runtime *util.RuntimeOptions) (_result *AttachKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a key pair to simple application servers.
//
// Description:
//
// You can bind only one key pair to a simple application server in the Simple Application Server console. If a simple application server has a key pair bound, the new key pair overwrites the original key pair.
//
// @param request - AttachKeyPairRequest
//
// @return AttachKeyPairResponse
func (client *Client) AttachKeyPair(request *AttachKeyPairRequest) (_result *AttachKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachKeyPairResponse{}
	_body, _err := client.AttachKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Cloud Assistant command.
//
// @param request - CreateCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCommandResponse
func (client *Client) CreateCommandWithOptions(request *CreateCommandRequest, runtime *util.RuntimeOptions) (_result *CreateCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableParameter)) {
		query["EnableParameter"] = request.EnableParameter
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingDir)) {
		query["WorkingDir"] = request.WorkingDir
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCommand"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Cloud Assistant command.
//
// @param request - CreateCommandRequest
//
// @return CreateCommandResponse
func (client *Client) CreateCommand(request *CreateCommandRequest) (_result *CreateCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCommandResponse{}
	_body, _err := client.CreateCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom image based on a snapshot of a simple application server.
//
// Description:
//
// A custom image is created based on a snapshot of a simple application server. You can use a custom image to create multiple simple application servers that have the same configurations. You can also share custom images to ECS and use the shared images to create ECS instances or replace the OSs of existing ECS instances. For more information about custom images, see [Overview of custom images](https://help.aliyun.com/document_detail/199375.html).
//
// You must create a system disk snapshot of a simple application server before you create a custom image based on the snapshot. For more information, see [CreateSnapshot](https://help.aliyun.com/document_detail/190452.html).
//
// > If you need the data on the data disk of a simple application server when you create a custom image, create a snapshot for the data disk first.
//
// Before you create a custom image, take note of the following items:
//
// 	- The custom image and the corresponding simple application server must reside in the same region.
//
// 	- The maximum number of custom images that can be maintained in an Alibaba Cloud account is triple the number of simple application servers in the account. The value cannot be greater than 15.
//
// 	- You can directly create a custom image only based on the system disk snapshot of a simple application server. If you want a custom image to contain the data on the data disk of the simple application server, you must select a data disk snapshot when you create the custom image.
//
// 	- If a simple application server is released due to expiration or refunds, the custom images that are created based on a snapshot of the server are also released.
//
// 	- If you reset a simple application server by changing the application system or OS of the server or replacing the image of the server, the disk data on the server is cleared. Back up the disk data as needed.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateCustomImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomImageResponse
func (client *Client) CreateCustomImageWithOptions(request *CreateCustomImageRequest, runtime *util.RuntimeOptions) (_result *CreateCustomImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSnapshotId)) {
		query["DataSnapshotId"] = request.DataSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResoureGroupId)) {
		query["ResoureGroupId"] = request.ResoureGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemSnapshotId)) {
		query["SystemSnapshotId"] = request.SystemSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomImage"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom image based on a snapshot of a simple application server.
//
// Description:
//
// A custom image is created based on a snapshot of a simple application server. You can use a custom image to create multiple simple application servers that have the same configurations. You can also share custom images to ECS and use the shared images to create ECS instances or replace the OSs of existing ECS instances. For more information about custom images, see [Overview of custom images](https://help.aliyun.com/document_detail/199375.html).
//
// You must create a system disk snapshot of a simple application server before you create a custom image based on the snapshot. For more information, see [CreateSnapshot](https://help.aliyun.com/document_detail/190452.html).
//
// > If you need the data on the data disk of a simple application server when you create a custom image, create a snapshot for the data disk first.
//
// Before you create a custom image, take note of the following items:
//
// 	- The custom image and the corresponding simple application server must reside in the same region.
//
// 	- The maximum number of custom images that can be maintained in an Alibaba Cloud account is triple the number of simple application servers in the account. The value cannot be greater than 15.
//
// 	- You can directly create a custom image only based on the system disk snapshot of a simple application server. If you want a custom image to contain the data on the data disk of the simple application server, you must select a data disk snapshot when you create the custom image.
//
// 	- If a simple application server is released due to expiration or refunds, the custom images that are created based on a snapshot of the server are also released.
//
// 	- If you reset a simple application server by changing the application system or OS of the server or replacing the image of the server, the disk data on the server is cleared. Back up the disk data as needed.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateCustomImageRequest
//
// @return CreateCustomImageResponse
func (client *Client) CreateCustomImage(request *CreateCustomImageRequest) (_result *CreateCustomImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomImageResponse{}
	_body, _err := client.CreateCustomImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a firewall rule for a simple application server.
//
// Description:
//
// Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
//
// ### QPS limits
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateFirewallRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFirewallRuleResponse
func (client *Client) CreateFirewallRuleWithOptions(request *CreateFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *CreateFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RuleProtocol)) {
		query["RuleProtocol"] = request.RuleProtocol
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a firewall rule for a simple application server.
//
// Description:
//
// Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
//
// ### QPS limits
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateFirewallRuleRequest
//
// @return CreateFirewallRuleResponse
func (client *Client) CreateFirewallRule(request *CreateFirewallRuleRequest) (_result *CreateFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFirewallRuleResponse{}
	_body, _err := client.CreateFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates multiple firewall rules for a simple application server at a time.
//
// Description:
//
// Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
//
// @param tmpReq - CreateFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFirewallRulesResponse
func (client *Client) CreateFirewallRulesWithOptions(tmpReq *CreateFirewallRulesRequest, runtime *util.RuntimeOptions) (_result *CreateFirewallRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateFirewallRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FirewallRules)) {
		request.FirewallRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FirewallRules, tea.String("FirewallRules"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallRulesShrink)) {
		query["FirewallRules"] = request.FirewallRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirewallRules"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates multiple firewall rules for a simple application server at a time.
//
// Description:
//
// Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
//
// @param request - CreateFirewallRulesRequest
//
// @return CreateFirewallRulesResponse
func (client *Client) CreateFirewallRules(request *CreateFirewallRulesRequest) (_result *CreateFirewallRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFirewallRulesResponse{}
	_body, _err := client.CreateFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a firewall template.
//
// Description:
//
// Simple Application Server supports the firewall template feature that provides multiple firewall rules. You can use a template to add a group of firewall rules to one or more simple application servers at a time. This improves the efficiency of setting firewall rules.
//
// @param request - CreateFirewallTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFirewallTemplateResponse
func (client *Client) CreateFirewallTemplateWithOptions(request *CreateFirewallTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateFirewallTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallRule)) {
		query["FirewallRule"] = request.FirewallRule
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirewallTemplate"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirewallTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a firewall template.
//
// Description:
//
// Simple Application Server supports the firewall template feature that provides multiple firewall rules. You can use a template to add a group of firewall rules to one or more simple application servers at a time. This improves the efficiency of setting firewall rules.
//
// @param request - CreateFirewallTemplateRequest
//
// @return CreateFirewallTemplateResponse
func (client *Client) CreateFirewallTemplate(request *CreateFirewallTemplateRequest) (_result *CreateFirewallTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFirewallTemplateResponse{}
	_body, _err := client.CreateFirewallTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds firewall rules to a firewall template based on your business requirements.
//
// Description:
//
// Adding firewall rules to a firewall template does not affect the firewall rules that have been applied to simple application servers..
//
// @param request - CreateFirewallTemplateRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFirewallTemplateRulesResponse
func (client *Client) CreateFirewallTemplateRulesWithOptions(request *CreateFirewallTemplateRulesRequest, runtime *util.RuntimeOptions) (_result *CreateFirewallTemplateRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallRule)) {
		query["FirewallRule"] = request.FirewallRule
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirewallTemplateRules"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirewallTemplateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds firewall rules to a firewall template based on your business requirements.
//
// Description:
//
// Adding firewall rules to a firewall template does not affect the firewall rules that have been applied to simple application servers..
//
// @param request - CreateFirewallTemplateRulesRequest
//
// @return CreateFirewallTemplateRulesResponse
func (client *Client) CreateFirewallTemplateRules(request *CreateFirewallTemplateRulesRequest) (_result *CreateFirewallTemplateRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateFirewallTemplateRulesResponse{}
	_body, _err := client.CreateFirewallTemplateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a key pair for a simple application server.
//
// @param request - CreateInstanceKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceKeyPairResponse
func (client *Client) CreateInstanceKeyPairWithOptions(request *CreateInstanceKeyPairRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a key pair for a simple application server.
//
// @param request - CreateInstanceKeyPairRequest
//
// @return CreateInstanceKeyPairResponse
func (client *Client) CreateInstanceKeyPair(request *CreateInstanceKeyPairRequest) (_result *CreateInstanceKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceKeyPairResponse{}
	_body, _err := client.CreateInstanceKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates subscription simple application servers.
//
// Description:
//
//   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](https://help.aliyun.com/document_detail/58623.html).
//
// 	- A maximum of 20 simple application servers can be maintained in an Alibaba Cloud account.
//
// 	- When you call this operation to create simple application servers, make sure that the balance in your account is sufficient to pay for the servers. If the balance in your account is insufficient, the servers cannot be created.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstancesResponse
func (client *Client) CreateInstancesWithOptions(request *CreateInstancesRequest, runtime *util.RuntimeOptions) (_result *CreateInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		query["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataDiskSize)) {
		query["DataDiskSize"] = request.DataDiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates subscription simple application servers.
//
// Description:
//
//   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](https://help.aliyun.com/document_detail/58623.html).
//
// 	- A maximum of 20 simple application servers can be maintained in an Alibaba Cloud account.
//
// 	- When you call this operation to create simple application servers, make sure that the balance in your account is sufficient to pay for the servers. If the balance in your account is insufficient, the servers cannot be created.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateInstancesRequest
//
// @return CreateInstancesResponse
func (client *Client) CreateInstances(request *CreateInstancesRequest) (_result *CreateInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstancesResponse{}
	_body, _err := client.CreateInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a key pair.
//
// Description:
//
// Alibaba Cloud SSH key pairs offer a secure and efficient logon authentication mechanism, facilitating both verification and encrypted communication within the SSH protocol framework. An SSH key pair is essentially constituted by a public key and a private key. Tailored for Linux-based simple application servers, this security measure enhances security and convenience, effectively addressing your heightened security requirements.
//
// 	- The key pair logon method is only valid for Linux-based simple application servers.
//
// 	- A maximum of 10 key pairs can be created in a region for an Alibaba Cloud account.
//
// 	- Only RSA 2048-bit key pairs can be created in the Simple Application Server console.
//
// @param request - CreateKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPairWithOptions(request *CreateKeyPairRequest, runtime *util.RuntimeOptions) (_result *CreateKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a key pair.
//
// Description:
//
// Alibaba Cloud SSH key pairs offer a secure and efficient logon authentication mechanism, facilitating both verification and encrypted communication within the SSH protocol framework. An SSH key pair is essentially constituted by a public key and a private key. Tailored for Linux-based simple application servers, this security measure enhances security and convenience, effectively addressing your heightened security requirements.
//
// 	- The key pair logon method is only valid for Linux-based simple application servers.
//
// 	- A maximum of 10 key pairs can be created in a region for an Alibaba Cloud account.
//
// 	- Only RSA 2048-bit key pairs can be created in the Simple Application Server console.
//
// @param request - CreateKeyPairRequest
//
// @return CreateKeyPairResponse
func (client *Client) CreateKeyPair(request *CreateKeyPairRequest) (_result *CreateKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyPairResponse{}
	_body, _err := client.CreateKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a snapshot for a disk.
//
// Description:
//
// A snapshot is a point-in-time backup of a disk. Snapshots can be used to back up data, recover data after accidental operations on instances, recover data after network attacks, and create custom images.
//
// > You are not charged for creating snapshots for disks of simple application servers.
//
// ### Precautions
//
// 	- You can create up to three snapshots for disks of each simple application server.
//
// 	- The maximum number of snapshots that can be retained in an Alibaba Cloud account is triple the number of simple application servers that you maintain. The value cannot be greater than 15.
//
// 	- If a simple application server is automatically released due to expiration, the snapshots created for the server are deleted.
//
// 	- If you reset the simple application server after you create a snapshot for a server, the snapshot is retained but cannot be used to roll back the disks of the server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshotWithOptions(request *CreateSnapshotRequest, runtime *util.RuntimeOptions) (_result *CreateSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSnapshot"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a snapshot for a disk.
//
// Description:
//
// A snapshot is a point-in-time backup of a disk. Snapshots can be used to back up data, recover data after accidental operations on instances, recover data after network attacks, and create custom images.
//
// > You are not charged for creating snapshots for disks of simple application servers.
//
// ### Precautions
//
// 	- You can create up to three snapshots for disks of each simple application server.
//
// 	- The maximum number of snapshots that can be retained in an Alibaba Cloud account is triple the number of simple application servers that you maintain. The value cannot be greater than 15.
//
// 	- If a simple application server is automatically released due to expiration, the snapshots created for the server are deleted.
//
// 	- If you reset the simple application server after you create a snapshot for a server, the snapshot is retained but cannot be used to roll back the disks of the server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - CreateSnapshotRequest
//
// @return CreateSnapshotResponse
func (client *Client) CreateSnapshot(request *CreateSnapshotRequest) (_result *CreateSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSnapshotResponse{}
	_body, _err := client.CreateSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a Command Assistant command.
//
// Description:
//
// You cannot delete commands that are being run.
//
// @param request - DeleteCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCommandResponse
func (client *Client) DeleteCommandWithOptions(request *DeleteCommandRequest, runtime *util.RuntimeOptions) (_result *DeleteCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCommand"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a Command Assistant command.
//
// Description:
//
// You cannot delete commands that are being run.
//
// @param request - DeleteCommandRequest
//
// @return DeleteCommandResponse
func (client *Client) DeleteCommand(request *DeleteCommandRequest) (_result *DeleteCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCommandResponse{}
	_body, _err := client.DeleteCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom image.
//
// Description:
//
// You can delete a custom image that you no longer need. After the custom image is deleted, you cannot use the custom image to reset the simple application servers that were created based on the custom image.
//
// > If a custom image is shared to Elastic Compute Service (ECS), you must unshare the image before you can delete it. After you unshare the custom image, you cannot query the custom image by using the ECS console or by calling ECS API operations. If you need to use the custom image in ECS, we recommend that you copy the image before you delete it. For more information, see [Copy a shared image of a simple application server in the ECS console](https://help.aliyun.com/document_detail/199378.html).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DeleteCustomImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomImageResponse
func (client *Client) DeleteCustomImageWithOptions(request *DeleteCustomImageRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomImage"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom image.
//
// Description:
//
// You can delete a custom image that you no longer need. After the custom image is deleted, you cannot use the custom image to reset the simple application servers that were created based on the custom image.
//
// > If a custom image is shared to Elastic Compute Service (ECS), you must unshare the image before you can delete it. After you unshare the custom image, you cannot query the custom image by using the ECS console or by calling ECS API operations. If you need to use the custom image in ECS, we recommend that you copy the image before you delete it. For more information, see [Copy a shared image of a simple application server in the ECS console](https://help.aliyun.com/document_detail/199378.html).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DeleteCustomImageRequest
//
// @return DeleteCustomImageResponse
func (client *Client) DeleteCustomImage(request *DeleteCustomImageRequest) (_result *DeleteCustomImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomImageResponse{}
	_body, _err := client.DeleteCustomImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes custom images. If you no longer require a custom image, you can call this operation to delete the custom image. You can also call this operation to delete multiple custom images at the same time. After a custom image is deleted, you cannot use the custom image to reset the simple application servers that were created based on the custom image.
//
// Description:
//
// If a custom image is shared, you must unshare the image before you can delete it. After a custom image is unshared, you cannot query the custom image by using the Elastic Compute Service (ECS) console or by calling an ECS API operation. If you want to use a custom image to create ECS instances, we recommend that you copy the custom image before you delete it. For more information, see the "Copy custom images" topic.
//
// @param request - DeleteCustomImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomImagesResponse
func (client *Client) DeleteCustomImagesWithOptions(request *DeleteCustomImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageIds)) {
		query["ImageIds"] = request.ImageIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomImages"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCustomImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes custom images. If you no longer require a custom image, you can call this operation to delete the custom image. You can also call this operation to delete multiple custom images at the same time. After a custom image is deleted, you cannot use the custom image to reset the simple application servers that were created based on the custom image.
//
// Description:
//
// If a custom image is shared, you must unshare the image before you can delete it. After a custom image is unshared, you cannot query the custom image by using the Elastic Compute Service (ECS) console or by calling an ECS API operation. If you want to use a custom image to create ECS instances, we recommend that you copy the custom image before you delete it. For more information, see the "Copy custom images" topic.
//
// @param request - DeleteCustomImagesRequest
//
// @return DeleteCustomImagesResponse
func (client *Client) DeleteCustomImages(request *DeleteCustomImagesRequest) (_result *DeleteCustomImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomImagesResponse{}
	_body, _err := client.DeleteCustomImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a firewall rule of a simple application server.
//
// Description:
//
// After a firewall rule is deleted, your business deployed on the simple application server may become inaccessible. Before you delete a firewall rule, make sure that the firewall rule is no longer needed by the simple application server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DeleteFirewallRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallRuleResponse
func (client *Client) DeleteFirewallRuleWithOptions(request *DeleteFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a firewall rule of a simple application server.
//
// Description:
//
// After a firewall rule is deleted, your business deployed on the simple application server may become inaccessible. Before you delete a firewall rule, make sure that the firewall rule is no longer needed by the simple application server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DeleteFirewallRuleRequest
//
// @return DeleteFirewallRuleResponse
func (client *Client) DeleteFirewallRule(request *DeleteFirewallRuleRequest) (_result *DeleteFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFirewallRuleResponse{}
	_body, _err := client.DeleteFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple firewall rules of a simple application server.
//
// Description:
//
// After a firewall rule is deleted, your business deployed on the simple application server may become inaccessible. Before you delete a firewall rule, make sure that the firewall rule is no longer needed by the simple application server.
//
// @param tmpReq - DeleteFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallRulesResponse
func (client *Client) DeleteFirewallRulesWithOptions(tmpReq *DeleteFirewallRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteFirewallRulesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteFirewallRulesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RuleIds)) {
		request.RuleIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RuleIds, tea.String("RuleIds"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleIdsShrink)) {
		query["RuleIds"] = request.RuleIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFirewallRules"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple firewall rules of a simple application server.
//
// Description:
//
// After a firewall rule is deleted, your business deployed on the simple application server may become inaccessible. Before you delete a firewall rule, make sure that the firewall rule is no longer needed by the simple application server.
//
// @param request - DeleteFirewallRulesRequest
//
// @return DeleteFirewallRulesResponse
func (client *Client) DeleteFirewallRules(request *DeleteFirewallRulesRequest) (_result *DeleteFirewallRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFirewallRulesResponse{}
	_body, _err := client.DeleteFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes firewall rules from a firewall template based on your requirements.
//
// Description:
//
// Deletion of firewall rules does not affect the firewall rules that have been applied to simple application servers.
//
// @param request - DeleteFirewallTemplateRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallTemplateRulesResponse
func (client *Client) DeleteFirewallTemplateRulesWithOptions(request *DeleteFirewallTemplateRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteFirewallTemplateRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateRuleId)) {
		query["FirewallTemplateRuleId"] = request.FirewallTemplateRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFirewallTemplateRules"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFirewallTemplateRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes firewall rules from a firewall template based on your requirements.
//
// Description:
//
// Deletion of firewall rules does not affect the firewall rules that have been applied to simple application servers.
//
// @param request - DeleteFirewallTemplateRulesRequest
//
// @return DeleteFirewallTemplateRulesResponse
func (client *Client) DeleteFirewallTemplateRules(request *DeleteFirewallTemplateRulesRequest) (_result *DeleteFirewallTemplateRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFirewallTemplateRulesResponse{}
	_body, _err := client.DeleteFirewallTemplateRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes firewall templates from a simple application server.
//
// Description:
//
// Deleting a firewall template does not affect the firewall rules that have been applied to simple application servers. You can delete firewall templates that you no longer need.
//
// @param request - DeleteFirewallTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFirewallTemplatesResponse
func (client *Client) DeleteFirewallTemplatesWithOptions(request *DeleteFirewallTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteFirewallTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFirewallTemplates"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFirewallTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes firewall templates from a simple application server.
//
// Description:
//
// Deleting a firewall template does not affect the firewall rules that have been applied to simple application servers. You can delete firewall templates that you no longer need.
//
// @param request - DeleteFirewallTemplatesRequest
//
// @return DeleteFirewallTemplatesResponse
func (client *Client) DeleteFirewallTemplates(request *DeleteFirewallTemplatesRequest) (_result *DeleteFirewallTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFirewallTemplatesResponse{}
	_body, _err := client.DeleteFirewallTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the key pair of a simple application server.
//
// @param request - DeleteInstanceKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceKeyPairResponse
func (client *Client) DeleteInstanceKeyPairWithOptions(request *DeleteInstanceKeyPairRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the key pair of a simple application server.
//
// @param request - DeleteInstanceKeyPairRequest
//
// @return DeleteInstanceKeyPairResponse
func (client *Client) DeleteInstanceKeyPair(request *DeleteInstanceKeyPairRequest) (_result *DeleteInstanceKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceKeyPairResponse{}
	_body, _err := client.DeleteInstanceKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the SSH key pairs of simple application servers.
//
// Description:
//
// You must unbind SSH key pairs that you no longer use from simple application servers before you delete the SSH key pairs.
//
// @param request - DeleteKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairsWithOptions(request *DeleteKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairNames)) {
		query["KeyPairNames"] = request.KeyPairNames
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeyPairs"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the SSH key pairs of simple application servers.
//
// Description:
//
// You must unbind SSH key pairs that you no longer use from simple application servers before you delete the SSH key pairs.
//
// @param request - DeleteKeyPairsRequest
//
// @return DeleteKeyPairsResponse
func (client *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (_result *DeleteKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DeleteKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a snapshot of a simple application server.
//
// Description:
//
// You can delete a snapshot if you no longer need it.
//
// > If a custom image was created based on the snapshot, delete the custom image before you delete the snapshot.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DeleteSnapshotRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshotWithOptions(request *DeleteSnapshotRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshot"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a snapshot of a simple application server.
//
// Description:
//
// You can delete a snapshot if you no longer need it.
//
// > If a custom image was created based on the snapshot, delete the custom image before you delete the snapshot.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DeleteSnapshotRequest
//
// @return DeleteSnapshotResponse
func (client *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (_result *DeleteSnapshotResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotResponse{}
	_body, _err := client.DeleteSnapshotWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes snapshots of a simple application server.
//
// @param request - DeleteSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSnapshotsResponse
func (client *Client) DeleteSnapshotsWithOptions(request *DeleteSnapshotsRequest, runtime *util.RuntimeOptions) (_result *DeleteSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSnapshots"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes snapshots of a simple application server.
//
// @param request - DeleteSnapshotsRequest
//
// @return DeleteSnapshotsResponse
func (client *Client) DeleteSnapshots(request *DeleteSnapshotsRequest) (_result *DeleteSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSnapshotsResponse{}
	_body, _err := client.DeleteSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Command Assistant information of simple application servers.
//
// @param tmpReq - DescribeCloudAssistantAttributesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudAssistantAttributesResponse
func (client *Client) DescribeCloudAssistantAttributesWithOptions(tmpReq *DescribeCloudAssistantAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudAssistantAttributesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCloudAssistantAttributesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudAssistantAttributes"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudAssistantAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Command Assistant information of simple application servers.
//
// @param request - DescribeCloudAssistantAttributesRequest
//
// @return DescribeCloudAssistantAttributesResponse
func (client *Client) DescribeCloudAssistantAttributes(request *DescribeCloudAssistantAttributesRequest) (_result *DescribeCloudAssistantAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudAssistantAttributesResponse{}
	_body, _err := client.DescribeCloudAssistantAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether the Cloud Assistant client is installed on simple application servers.
//
// Description:
//
// By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall the client. Otherwise, you cannot run commands on the servers.
//
// @param tmpReq - DescribeCloudAssistantStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudAssistantStatusResponse
func (client *Client) DescribeCloudAssistantStatusWithOptions(tmpReq *DescribeCloudAssistantStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudAssistantStatusResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeCloudAssistantStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudAssistantStatus"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudAssistantStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether the Cloud Assistant client is installed on simple application servers.
//
// Description:
//
// By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall the client. Otherwise, you cannot run commands on the servers.
//
// @param request - DescribeCloudAssistantStatusRequest
//
// @return DescribeCloudAssistantStatusResponse
func (client *Client) DescribeCloudAssistantStatus(request *DescribeCloudAssistantStatusRequest) (_result *DescribeCloudAssistantStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudAssistantStatusResponse{}
	_body, _err := client.DescribeCloudAssistantStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the CloudMonitor agent on simple application servers.
//
// @param request - DescribeCloudMonitorAgentStatusesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCloudMonitorAgentStatusesResponse
func (client *Client) DescribeCloudMonitorAgentStatusesWithOptions(request *DescribeCloudMonitorAgentStatusesRequest, runtime *util.RuntimeOptions) (_result *DescribeCloudMonitorAgentStatusesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCloudMonitorAgentStatuses"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCloudMonitorAgentStatusesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the CloudMonitor agent on simple application servers.
//
// @param request - DescribeCloudMonitorAgentStatusesRequest
//
// @return DescribeCloudMonitorAgentStatusesResponse
func (client *Client) DescribeCloudMonitorAgentStatuses(request *DescribeCloudMonitorAgentStatusesRequest) (_result *DescribeCloudMonitorAgentStatusesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCloudMonitorAgentStatusesResponse{}
	_body, _err := client.DescribeCloudMonitorAgentStatusesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the executions and execution status of a Cloud Assistant command.
//
// @param request - DescribeCommandInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommandInvocationsResponse
func (client *Client) DescribeCommandInvocationsWithOptions(request *DescribeCommandInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeCommandInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.CommandName)) {
		query["CommandName"] = request.CommandName
	}

	if !tea.BoolValue(util.IsUnset(request.CommandType)) {
		query["CommandType"] = request.CommandType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InvocationStatus)) {
		query["InvocationStatus"] = request.InvocationStatus
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		query["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCommandInvocations"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCommandInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the executions and execution status of a Cloud Assistant command.
//
// @param request - DescribeCommandInvocationsRequest
//
// @return DescribeCommandInvocationsResponse
func (client *Client) DescribeCommandInvocations(request *DescribeCommandInvocationsRequest) (_result *DescribeCommandInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCommandInvocationsResponse{}
	_body, _err := client.DescribeCommandInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the commands that you created or the common commands that Alibaba Cloud provides.
//
// @param request - DescribeCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCommandsResponse
func (client *Client) DescribeCommandsWithOptions(request *DescribeCommandsRequest, runtime *util.RuntimeOptions) (_result *DescribeCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCommands"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the commands that you created or the common commands that Alibaba Cloud provides.
//
// @param request - DescribeCommandsRequest
//
// @return DescribeCommandsResponse
func (client *Client) DescribeCommands(request *DescribeCommandsRequest) (_result *DescribeCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCommandsResponse{}
	_body, _err := client.DescribeCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries error logs of databases in a Simple Database Service instance.
//
// Description:
//
// You can call this operation to query the error logs of databases in a Simple Database Service instance and locate faults based on the error logs.
//
// \\### QPS limit You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseErrorLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabaseErrorLogsResponse
func (client *Client) DescribeDatabaseErrorLogsWithOptions(request *DescribeDatabaseErrorLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseErrorLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseErrorLogs"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseErrorLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries error logs of databases in a Simple Database Service instance.
//
// Description:
//
// You can call this operation to query the error logs of databases in a Simple Database Service instance and locate faults based on the error logs.
//
// \\### QPS limit You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseErrorLogsRequest
//
// @return DescribeDatabaseErrorLogsResponse
func (client *Client) DescribeDatabaseErrorLogs(request *DescribeDatabaseErrorLogsRequest) (_result *DescribeDatabaseErrorLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseErrorLogsResponse{}
	_body, _err := client.DescribeDatabaseErrorLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring information about a Simple Database Service instance.
//
// Description:
//
// After you create a Simple Database Service instance, you can query the details about the vCPU, memory, disk size, storage IOPS (input/output operations per second), and total current connection number of the instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseInstanceMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabaseInstanceMetricDataResponse
func (client *Client) DescribeDatabaseInstanceMetricDataWithOptions(request *DescribeDatabaseInstanceMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseInstanceMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseInstanceMetricData"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseInstanceMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring information about a Simple Database Service instance.
//
// Description:
//
// After you create a Simple Database Service instance, you can query the details about the vCPU, memory, disk size, storage IOPS (input/output operations per second), and total current connection number of the instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseInstanceMetricDataRequest
//
// @return DescribeDatabaseInstanceMetricDataResponse
func (client *Client) DescribeDatabaseInstanceMetricData(request *DescribeDatabaseInstanceMetricDataRequest) (_result *DescribeDatabaseInstanceMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseInstanceMetricDataResponse{}
	_body, _err := client.DescribeDatabaseInstanceMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the parameters of a Simple Database Service instance.
//
// Description:
//
// You can call this operation to query the information about parameters of a Simple Database Service instance.
//
// @param request - DescribeDatabaseInstanceParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabaseInstanceParametersResponse
func (client *Client) DescribeDatabaseInstanceParametersWithOptions(request *DescribeDatabaseInstanceParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseInstanceParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseInstanceParameters"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseInstanceParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the parameters of a Simple Database Service instance.
//
// Description:
//
// You can call this operation to query the information about parameters of a Simple Database Service instance.
//
// @param request - DescribeDatabaseInstanceParametersRequest
//
// @return DescribeDatabaseInstanceParametersResponse
func (client *Client) DescribeDatabaseInstanceParameters(request *DescribeDatabaseInstanceParametersRequest) (_result *DescribeDatabaseInstanceParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseInstanceParametersResponse{}
	_body, _err := client.DescribeDatabaseInstanceParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about Simple Database Service instances.
//
// Description:
//
// You can call this operation to query the details of Simple Database Service instances in a region, including the IDs, names, plans, database versions, public endpoint, internal endpoint, creation time, and expiration time of the instances.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabaseInstancesResponse
func (client *Client) DescribeDatabaseInstancesWithOptions(request *DescribeDatabaseInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceIds)) {
		query["DatabaseInstanceIds"] = request.DatabaseInstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about Simple Database Service instances.
//
// Description:
//
// You can call this operation to query the details of Simple Database Service instances in a region, including the IDs, names, plans, database versions, public endpoint, internal endpoint, creation time, and expiration time of the instances.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseInstancesRequest
//
// @return DescribeDatabaseInstancesResponse
func (client *Client) DescribeDatabaseInstances(request *DescribeDatabaseInstancesRequest) (_result *DescribeDatabaseInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseInstancesResponse{}
	_body, _err := client.DescribeDatabaseInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the slow query log details of a Simple Database Service instance.
//
// Description:
//
// You can query the slow query log details of a Simple Database Service instance and locate faults based on the log details.
//
// > Slow query log details are retained for 7 days.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseSlowLogRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDatabaseSlowLogRecordsResponse
func (client *Client) DescribeDatabaseSlowLogRecordsWithOptions(request *DescribeDatabaseSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabaseSlowLogRecords"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabaseSlowLogRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the slow query log details of a Simple Database Service instance.
//
// Description:
//
// You can query the slow query log details of a Simple Database Service instance and locate faults based on the log details.
//
// > Slow query log details are retained for 7 days.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - DescribeDatabaseSlowLogRecordsRequest
//
// @return DescribeDatabaseSlowLogRecordsResponse
func (client *Client) DescribeDatabaseSlowLogRecords(request *DescribeDatabaseSlowLogRecordsRequest) (_result *DescribeDatabaseSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabaseSlowLogRecordsResponse{}
	_body, _err := client.DescribeDatabaseSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of applying a firewall template to simple application servers.
//
// @param request - DescribeFirewallTemplateApplyResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallTemplateApplyResultsResponse
func (client *Client) DescribeFirewallTemplateApplyResultsWithOptions(request *DescribeFirewallTemplateApplyResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeFirewallTemplateApplyResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFirewallTemplateApplyResults"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFirewallTemplateApplyResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of applying a firewall template to simple application servers.
//
// @param request - DescribeFirewallTemplateApplyResultsRequest
//
// @return DescribeFirewallTemplateApplyResultsResponse
func (client *Client) DescribeFirewallTemplateApplyResults(request *DescribeFirewallTemplateApplyResultsRequest) (_result *DescribeFirewallTemplateApplyResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFirewallTemplateApplyResultsResponse{}
	_body, _err := client.DescribeFirewallTemplateApplyResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the result of applying a firewall template rule to simple application servers.
//
// @param request - DescribeFirewallTemplateRulesApplyResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallTemplateRulesApplyResultResponse
func (client *Client) DescribeFirewallTemplateRulesApplyResultWithOptions(request *DescribeFirewallTemplateRulesApplyResultRequest, runtime *util.RuntimeOptions) (_result *DescribeFirewallTemplateRulesApplyResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFirewallTemplateRulesApplyResult"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFirewallTemplateRulesApplyResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the result of applying a firewall template rule to simple application servers.
//
// @param request - DescribeFirewallTemplateRulesApplyResultRequest
//
// @return DescribeFirewallTemplateRulesApplyResultResponse
func (client *Client) DescribeFirewallTemplateRulesApplyResult(request *DescribeFirewallTemplateRulesApplyResultRequest) (_result *DescribeFirewallTemplateRulesApplyResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFirewallTemplateRulesApplyResultResponse{}
	_body, _err := client.DescribeFirewallTemplateRulesApplyResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about firewall templates.
//
// @param request - DescribeFirewallTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirewallTemplatesResponse
func (client *Client) DescribeFirewallTemplatesWithOptions(request *DescribeFirewallTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeFirewallTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFirewallTemplates"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFirewallTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about firewall templates.
//
// @param request - DescribeFirewallTemplatesRequest
//
// @return DescribeFirewallTemplatesResponse
func (client *Client) DescribeFirewallTemplates(request *DescribeFirewallTemplatesRequest) (_result *DescribeFirewallTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFirewallTemplatesResponse{}
	_body, _err := client.DescribeFirewallTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries information about the key pair of a simple application server.
//
// @param request - DescribeInstanceKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceKeyPairResponse
func (client *Client) DescribeInstanceKeyPairWithOptions(request *DescribeInstanceKeyPairRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries information about the key pair of a simple application server.
//
// @param request - DescribeInstanceKeyPairRequest
//
// @return DescribeInstanceKeyPairResponse
func (client *Client) DescribeInstanceKeyPair(request *DescribeInstanceKeyPairRequest) (_result *DescribeInstanceKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceKeyPairResponse{}
	_body, _err := client.DescribeInstanceKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Checks whether a password is set for a simple application server.
//
// @param request - DescribeInstancePasswordsSettingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstancePasswordsSettingResponse
func (client *Client) DescribeInstancePasswordsSettingWithOptions(request *DescribeInstancePasswordsSettingRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancePasswordsSettingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstancePasswordsSetting"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancePasswordsSettingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Checks whether a password is set for a simple application server.
//
// @param request - DescribeInstancePasswordsSettingRequest
//
// @return DescribeInstancePasswordsSettingResponse
func (client *Client) DescribeInstancePasswordsSetting(request *DescribeInstancePasswordsSettingRequest) (_result *DescribeInstancePasswordsSettingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancePasswordsSettingResponse{}
	_body, _err := client.DescribeInstancePasswordsSettingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the VNC connection address of a simple application server.
//
// @param request - DescribeInstanceVncUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInstanceVncUrlResponse
func (client *Client) DescribeInstanceVncUrlWithOptions(request *DescribeInstanceVncUrlRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceVncUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceVncUrl"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceVncUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the VNC connection address of a simple application server.
//
// @param request - DescribeInstanceVncUrlRequest
//
// @return DescribeInstanceVncUrlResponse
func (client *Client) DescribeInstanceVncUrl(request *DescribeInstanceVncUrlRequest) (_result *DescribeInstanceVncUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceVncUrlResponse{}
	_body, _err := client.DescribeInstanceVncUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution result of a command.
//
// Description:
//
//   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the execution result of a command.
//
// 	- You can query the execution results that were generated within the last two weeks. A maximum of 100,000 entries of execution results can be retained.
//
// @param request - DescribeInvocationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationResultResponse
func (client *Client) DescribeInvocationResultWithOptions(request *DescribeInvocationResultRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		query["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocationResult"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution result of a command.
//
// Description:
//
//   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the execution result of a command.
//
// 	- You can query the execution results that were generated within the last two weeks. A maximum of 100,000 entries of execution results can be retained.
//
// @param request - DescribeInvocationResultRequest
//
// @return DescribeInvocationResultResponse
func (client *Client) DescribeInvocationResult(request *DescribeInvocationResultRequest) (_result *DescribeInvocationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationResultResponse{}
	_body, _err := client.DescribeInvocationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about command execution.
//
// Description:
//
//   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the actual execution results.
//
// 	- You can query the execution results that were generated within the last two weeks. Up to 100,000 entries of execution results can be retained.
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeStatus)) {
		query["InvokeStatus"] = request.InvokeStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocations"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about command execution.
//
// Description:
//
//   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the actual execution results.
//
// 	- You can query the execution results that were generated within the last two weeks. Up to 100,000 entries of execution results can be retained.
//
// @param request - DescribeInvocationsRequest
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the monitoring information about vCPUs, memory, disk IOPS, and traffic of a simple application server.
//
// @param request - DescribeMonitorDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMonitorDataResponse
func (client *Client) DescribeMonitorDataWithOptions(request *DescribeMonitorDataRequest, runtime *util.RuntimeOptions) (_result *DescribeMonitorDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Length)) {
		query["Length"] = request.Length
	}

	if !tea.BoolValue(util.IsUnset(request.MetricName)) {
		query["MetricName"] = request.MetricName
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMonitorData"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMonitorDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the monitoring information about vCPUs, memory, disk IOPS, and traffic of a simple application server.
//
// @param request - DescribeMonitorDataRequest
//
// @return DescribeMonitorDataResponse
func (client *Client) DescribeMonitorData(request *DescribeMonitorDataRequest) (_result *DescribeMonitorDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMonitorDataResponse{}
	_body, _err := client.DescribeMonitorDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of the Security Center agent on a simple application server.
//
// @param request - DescribeSecurityAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecurityAgentStatusResponse
func (client *Client) DescribeSecurityAgentStatusWithOptions(request *DescribeSecurityAgentStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityAgentStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityAgentStatus"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of the Security Center agent on a simple application server.
//
// @param request - DescribeSecurityAgentStatusRequest
//
// @return DescribeSecurityAgentStatusResponse
func (client *Client) DescribeSecurityAgentStatus(request *DescribeSecurityAgentStatusRequest) (_result *DescribeSecurityAgentStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityAgentStatusResponse{}
	_body, _err := client.DescribeSecurityAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unbinds key pairs from simple application servers.
//
// Description:
//
// If you want to change the SSH key pairs that are bound to your simple application servers or your end user no longer needs to access a specific simple application server, you can unbind the SSH key pairs from simple application servers to improve the security of the simple application servers or restrict access to the specific simple application server.
//
// @param request - DetachKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachKeyPairResponse
func (client *Client) DetachKeyPairWithOptions(request *DetachKeyPairRequest, runtime *util.RuntimeOptions) (_result *DetachKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unbinds key pairs from simple application servers.
//
// Description:
//
// If you want to change the SSH key pairs that are bound to your simple application servers or your end user no longer needs to access a specific simple application server, you can unbind the SSH key pairs from simple application servers to improve the security of the simple application servers or restrict access to the specific simple application server.
//
// @param request - DetachKeyPairRequest
//
// @return DetachKeyPairResponse
func (client *Client) DetachKeyPair(request *DetachKeyPairRequest) (_result *DetachKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachKeyPairResponse{}
	_body, _err := client.DetachKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disables a firewall rule of a simple application server.
//
// @param request - DisableFirewallRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableFirewallRuleResponse
func (client *Client) DisableFirewallRuleWithOptions(request *DisableFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *DisableFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disables a firewall rule of a simple application server.
//
// @param request - DisableFirewallRuleRequest
//
// @return DisableFirewallRuleResponse
func (client *Client) DisableFirewallRule(request *DisableFirewallRuleRequest) (_result *DisableFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableFirewallRuleResponse{}
	_body, _err := client.DisableFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables a firewall rule for a simple application server.
//
// @param request - EnableFirewallRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableFirewallRuleResponse
func (client *Client) EnableFirewallRuleWithOptions(request *EnableFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *EnableFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables a firewall rule for a simple application server.
//
// @param request - EnableFirewallRuleRequest
//
// @return EnableFirewallRuleResponse
func (client *Client) EnableFirewallRule(request *EnableFirewallRuleRequest) (_result *EnableFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableFirewallRuleResponse{}
	_body, _err := client.EnableFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports an existing key pair to the Simple Application Server console.
//
// Description:
//
// You can call this operation to import an existing key pair to the Simple Application Server console. This way, you can use the key pair to log on to simple application servers. The existing key pair that you want to import must use a supported encryption method. For more information, see [Q2: Which encryption methods must be used by key pairs when I import existing key pairs to the Simple Application Server console?](https://help.aliyun.com/document_detail/59085.html)
//
// @param request - ImportKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPairWithOptions(request *ImportKeyPairRequest, runtime *util.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyBody)) {
		query["PublicKeyBody"] = request.PublicKeyBody
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports an existing key pair to the Simple Application Server console.
//
// Description:
//
// You can call this operation to import an existing key pair to the Simple Application Server console. This way, you can use the key pair to log on to simple application servers. The existing key pair that you want to import must use a supported encryption method. For more information, see [Q2: Which encryption methods must be used by key pairs when I import existing key pairs to the Simple Application Server console?](https://help.aliyun.com/document_detail/59085.html)
//
// @param request - ImportKeyPairRequest
//
// @return ImportKeyPairResponse
func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (_result *ImportKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.ImportKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs the Cloud Assistant client on simple application servers at a time.
//
// Description:
//
// To run commands on your simple application servers, you must install the Cloud Assistant client on your servers. You can call the [DescribeCloudAssistantStatus](https://help.aliyun.com/document_detail/439512.html) operation to check whether the Cloud Assistant client is installed on your simple application servers. If you have not installed the Cloud Assistant client, you can call the InstallCloudAssistant operation to install the client. Then, you can call the [RebootInstance](https://help.aliyun.com/document_detail/190443.html) operation to restart the servers to allow the client to take effect.
//
// @param tmpReq - InstallCloudAssistantRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCloudAssistantResponse
func (client *Client) InstallCloudAssistantWithOptions(tmpReq *InstallCloudAssistantRequest, runtime *util.RuntimeOptions) (_result *InstallCloudAssistantResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InstallCloudAssistantShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceIds)) {
		request.InstanceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceIds, tea.String("InstanceIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIdsShrink)) {
		query["InstanceIds"] = request.InstanceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallCloudAssistant"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCloudAssistantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs the Cloud Assistant client on simple application servers at a time.
//
// Description:
//
// To run commands on your simple application servers, you must install the Cloud Assistant client on your servers. You can call the [DescribeCloudAssistantStatus](https://help.aliyun.com/document_detail/439512.html) operation to check whether the Cloud Assistant client is installed on your simple application servers. If you have not installed the Cloud Assistant client, you can call the InstallCloudAssistant operation to install the client. Then, you can call the [RebootInstance](https://help.aliyun.com/document_detail/190443.html) operation to restart the servers to allow the client to take effect.
//
// @param request - InstallCloudAssistantRequest
//
// @return InstallCloudAssistantResponse
func (client *Client) InstallCloudAssistant(request *InstallCloudAssistantRequest) (_result *InstallCloudAssistantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallCloudAssistantResponse{}
	_body, _err := client.InstallCloudAssistantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Installs the CloudMonitor agent for a simple application server.
//
// @param request - InstallCloudMonitorAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InstallCloudMonitorAgentResponse
func (client *Client) InstallCloudMonitorAgentWithOptions(request *InstallCloudMonitorAgentRequest, runtime *util.RuntimeOptions) (_result *InstallCloudMonitorAgentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallCloudMonitorAgent"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallCloudMonitorAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Installs the CloudMonitor agent for a simple application server.
//
// @param request - InstallCloudMonitorAgentRequest
//
// @return InstallCloudMonitorAgentResponse
func (client *Client) InstallCloudMonitorAgent(request *InstallCloudMonitorAgentRequest) (_result *InstallCloudMonitorAgentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallCloudMonitorAgentResponse{}
	_body, _err := client.InstallCloudMonitorAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs a Command Assistant command for one or more simple application servers.
//
// Description:
//
//   The simple application servers for which you want to call the operation must meet the following conditions. If a simple application server cannot meet the conditions, you must call this operation again.
//
//     	- The simple application servers are in the `Running` state. You can call the [ListInstances](https://help.aliyun.com/document_detail/2361065.html) operation to query the status of simple application servers.
//
//     	- Cloud Assistant Agent is installed on the simple application servers. For more information, see [InstallCloudAssistant](https://help.aliyun.com/document_detail/2361030.html).
//
//     	- If you run a PowerShell command, make sure that the PowerShell module is configured for the simple application servers.
//
// 	- The command may fail to be run due to the abnormal states of simple application servers, network exceptions, or exceptions in Cloud Assistant Agent. If the command fails to be run, no execution information is generated.
//
// 	- If you enable the custom parameter feature when you create a command, you must set the `Parameters` parameter to specify custom parameters when you run the command.
//
// 	- When you call this operation, you can select only one common command or a custom command that you have created.
//
// @param tmpReq - InvokeCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokeCommandResponse
func (client *Client) InvokeCommandWithOptions(tmpReq *InvokeCommandRequest, runtime *util.RuntimeOptions) (_result *InvokeCommandResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InvokeCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InvokeCommand"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InvokeCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs a Command Assistant command for one or more simple application servers.
//
// Description:
//
//   The simple application servers for which you want to call the operation must meet the following conditions. If a simple application server cannot meet the conditions, you must call this operation again.
//
//     	- The simple application servers are in the `Running` state. You can call the [ListInstances](https://help.aliyun.com/document_detail/2361065.html) operation to query the status of simple application servers.
//
//     	- Cloud Assistant Agent is installed on the simple application servers. For more information, see [InstallCloudAssistant](https://help.aliyun.com/document_detail/2361030.html).
//
//     	- If you run a PowerShell command, make sure that the PowerShell module is configured for the simple application servers.
//
// 	- The command may fail to be run due to the abnormal states of simple application servers, network exceptions, or exceptions in Cloud Assistant Agent. If the command fails to be run, no execution information is generated.
//
// 	- If you enable the custom parameter feature when you create a command, you must set the `Parameters` parameter to specify custom parameters when you run the command.
//
// 	- When you call this operation, you can select only one common command or a custom command that you have created.
//
// @param request - InvokeCommandRequest
//
// @return InvokeCommandResponse
func (client *Client) InvokeCommand(request *InvokeCommandRequest) (_result *InvokeCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeCommandResponse{}
	_body, _err := client.InvokeCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询镜像共享给的用户
//
// @param request - ListCustomImageShareAccountsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomImageShareAccountsResponse
func (client *Client) ListCustomImageShareAccountsWithOptions(request *ListCustomImageShareAccountsRequest, runtime *util.RuntimeOptions) (_result *ListCustomImageShareAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomImageShareAccounts"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomImageShareAccountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询镜像共享给的用户
//
// @param request - ListCustomImageShareAccountsRequest
//
// @return ListCustomImageShareAccountsResponse
func (client *Client) ListCustomImageShareAccounts(request *ListCustomImageShareAccountsRequest) (_result *ListCustomImageShareAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomImageShareAccountsResponse{}
	_body, _err := client.ListCustomImageShareAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about custom images in a region.
//
// @param request - ListCustomImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomImagesResponse
func (client *Client) ListCustomImagesWithOptions(request *ListCustomImagesRequest, runtime *util.RuntimeOptions) (_result *ListCustomImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DataSnapshotId)) {
		query["DataSnapshotId"] = request.DataSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageIds)) {
		query["ImageIds"] = request.ImageIds
	}

	if !tea.BoolValue(util.IsUnset(request.ImageNames)) {
		query["ImageNames"] = request.ImageNames
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Share)) {
		query["Share"] = request.Share
	}

	if !tea.BoolValue(util.IsUnset(request.SystemSnapshotId)) {
		query["SystemSnapshotId"] = request.SystemSnapshotId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCustomImages"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about custom images in a region.
//
// @param request - ListCustomImagesRequest
//
// @return ListCustomImagesResponse
func (client *Client) ListCustomImages(request *ListCustomImagesRequest) (_result *ListCustomImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.ListCustomImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about disks in a region.
//
// Description:
//
// You can specify multiple request parameters that you want to query, such as `InstanceId`, `DiskIds`, and `ResourceGroupId`. Specified request parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListDisksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDisksResponse
func (client *Client) ListDisksWithOptions(request *ListDisksRequest, runtime *util.RuntimeOptions) (_result *ListDisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskIds)) {
		query["DiskIds"] = request.DiskIds
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		query["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDisks"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about disks in a region.
//
// Description:
//
// You can specify multiple request parameters that you want to query, such as `InstanceId`, `DiskIds`, and `ResourceGroupId`. Specified request parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListDisksRequest
//
// @return ListDisksResponse
func (client *Client) ListDisks(request *ListDisksRequest) (_result *ListDisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDisksResponse{}
	_body, _err := client.ListDisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the firewall rules of a simple application server.
//
// Description:
//
// You can call the ListFirewallRules operation to query the firewall rule details of a simple application server, including the port range, firewall rule ID, and transport layer protocol.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListFirewallRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFirewallRulesResponse
func (client *Client) ListFirewallRulesWithOptions(request *ListFirewallRulesRequest, runtime *util.RuntimeOptions) (_result *ListFirewallRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FirewallRuleId)) {
		query["FirewallRuleId"] = request.FirewallRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFirewallRules"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFirewallRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the firewall rules of a simple application server.
//
// Description:
//
// You can call the ListFirewallRules operation to query the firewall rule details of a simple application server, including the port range, firewall rule ID, and transport layer protocol.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListFirewallRulesRequest
//
// @return ListFirewallRulesResponse
func (client *Client) ListFirewallRules(request *ListFirewallRulesRequest) (_result *ListFirewallRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFirewallRulesResponse{}
	_body, _err := client.ListFirewallRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about images in a region.
//
// Description:
//
// You can query information about images in a region, including the IDs, names, and types of the images.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageIds)) {
		query["ImageIds"] = request.ImageIds
	}

	if !tea.BoolValue(util.IsUnset(request.ImageType)) {
		query["ImageType"] = request.ImageType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about images in a region.
//
// Description:
//
// You can query information about images in a region, including the IDs, names, and types of the images.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the plans to which you can upgrade your simple application server.
//
// Description:
//
// If the plan of your simple application server does not meet your business requirements, you can call the ListInstancePlansModification operation to obtain a list of plans to which you can upgrade your simple application server. Then, you can call the [UpgradeInstance](https://help.aliyun.com/document_detail/190445.html) operation to upgrade the server.
//
// > We recommend that you create snapshots for the disks of your simple application server to back up data before you upgrade the server. For more information, see [CreateSnapshot](https://help.aliyun.com/document_detail/190452.html).
//
// For the precautions about plan upgrade, see [Upgrade a simple application server](https://help.aliyun.com/document_detail/61433.html).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListInstancePlansModificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePlansModificationResponse
func (client *Client) ListInstancePlansModificationWithOptions(request *ListInstancePlansModificationRequest, runtime *util.RuntimeOptions) (_result *ListInstancePlansModificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancePlansModification"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancePlansModificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the plans to which you can upgrade your simple application server.
//
// Description:
//
// If the plan of your simple application server does not meet your business requirements, you can call the ListInstancePlansModification operation to obtain a list of plans to which you can upgrade your simple application server. Then, you can call the [UpgradeInstance](https://help.aliyun.com/document_detail/190445.html) operation to upgrade the server.
//
// > We recommend that you create snapshots for the disks of your simple application server to back up data before you upgrade the server. For more information, see [CreateSnapshot](https://help.aliyun.com/document_detail/190452.html).
//
// For the precautions about plan upgrade, see [Upgrade a simple application server](https://help.aliyun.com/document_detail/61433.html).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListInstancePlansModificationRequest
//
// @return ListInstancePlansModificationResponse
func (client *Client) ListInstancePlansModification(request *ListInstancePlansModificationRequest) (_result *ListInstancePlansModificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePlansModificationResponse{}
	_body, _err := client.ListInstancePlansModificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of simple application servers.
//
// @param request - ListInstanceStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceStatusResponse
func (client *Client) ListInstanceStatusWithOptions(request *ListInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *ListInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceStatus"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of simple application servers.
//
// @param request - ListInstanceStatusRequest
//
// @return ListInstanceStatusResponse
func (client *Client) ListInstanceStatus(request *ListInstanceStatusRequest) (_result *ListInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceStatusResponse{}
	_body, _err := client.ListInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about simple application servers in a region.
//
// Description:
//
// You can call this operation to query the details of simple application servers in a specified region, including the names, public IP addresses, internal IP addresses, creation time, and expiration time of the servers.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesResponse
func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublicIpAddresses)) {
		query["PublicIpAddresses"] = request.PublicIpAddresses
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about simple application servers in a region.
//
// Description:
//
// You can call this operation to query the details of simple application servers in a specified region, including the names, public IP addresses, internal IP addresses, creation time, and expiration time of the servers.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListInstancesRequest
//
// @return ListInstancesResponse
func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Queries the details about data transfer plans of simple application servers.
//
// Description:
//
// You can query the details of data transfer plans of simple application servers, including the data transfer quota, used amount and unused amount of the data transfer quota, and excess data transfers beyond the quota in the current month.
//
// Simple Application Server provides data transfer quotas in plans. Plan prices include prices of data transfer quotas. You are charged for data transfers that exceed the quotas. Take note of the following items:
//
// 	- Only outbound data transfers of simple application servers over the Internet are calculated. Outbound data transfers include the data transfer quota and the excess data transfers beyond the quota. Inbound data transfers of simple application servers over the Internet are not calculated.
//
// 	- Outbound data transfers from simple application servers to other Alibaba Cloud services over the Internet first consume data transfer quotas. If the quotas are exhausted, you are charged for excess data transfers.
//
// 	- You are not charged for data transfers between simple application servers within the same virtual private cloud (VPC).
//
// For more information, see [Quotas and billing of data transfers](https://help.aliyun.com/document_detail/86281.html).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListInstancesTrafficPackagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancesTrafficPackagesResponse
func (client *Client) ListInstancesTrafficPackagesWithOptions(request *ListInstancesTrafficPackagesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesTrafficPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancesTrafficPackages"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesTrafficPackagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details about data transfer plans of simple application servers.
//
// Description:
//
// You can query the details of data transfer plans of simple application servers, including the data transfer quota, used amount and unused amount of the data transfer quota, and excess data transfers beyond the quota in the current month.
//
// Simple Application Server provides data transfer quotas in plans. Plan prices include prices of data transfer quotas. You are charged for data transfers that exceed the quotas. Take note of the following items:
//
// 	- Only outbound data transfers of simple application servers over the Internet are calculated. Outbound data transfers include the data transfer quota and the excess data transfers beyond the quota. Inbound data transfers of simple application servers over the Internet are not calculated.
//
// 	- Outbound data transfers from simple application servers to other Alibaba Cloud services over the Internet first consume data transfer quotas. If the quotas are exhausted, you are charged for excess data transfers.
//
// 	- You are not charged for data transfers between simple application servers within the same virtual private cloud (VPC).
//
// For more information, see [Quotas and billing of data transfers](https://help.aliyun.com/document_detail/86281.html).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListInstancesTrafficPackagesRequest
//
// @return ListInstancesTrafficPackagesResponse
func (client *Client) ListInstancesTrafficPackages(request *ListInstancesTrafficPackagesRequest) (_result *ListInstancesTrafficPackagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesTrafficPackagesResponse{}
	_body, _err := client.ListInstancesTrafficPackagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs that are bound to simple application servers in a specific region.
//
// @param request - ListKeyPairsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKeyPairsResponse
func (client *Client) ListKeyPairsWithOptions(request *ListKeyPairsRequest, runtime *util.RuntimeOptions) (_result *ListKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKeyPairs"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the AccessKey pairs that are bound to simple application servers in a specific region.
//
// @param request - ListKeyPairsRequest
//
// @return ListKeyPairsResponse
func (client *Client) ListKeyPairs(request *ListKeyPairsRequest) (_result *ListKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeyPairsResponse{}
	_body, _err := client.ListKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all plans provided by Simple Application Server in a region.
//
// Description:
//
// You can query the details of all plans provided by Simple Application Server in a region, including the IDs, prices, disk sizes, and disk categories of the plans.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPlansResponse
func (client *Client) ListPlansWithOptions(request *ListPlansRequest, runtime *util.RuntimeOptions) (_result *ListPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPlans"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all plans provided by Simple Application Server in a region.
//
// Description:
//
// You can query the details of all plans provided by Simple Application Server in a region, including the IDs, prices, disk sizes, and disk categories of the plans.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListPlansRequest
//
// @return ListPlansResponse
func (client *Client) ListPlans(request *ListPlansRequest) (_result *ListPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPlansResponse{}
	_body, _err := client.ListPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all regions in which Simple Application Server is supported.
//
// Description:
//
// The query results include all the Alibaba Cloud regions where Simple Application Server is supported on the international site (alibabacloud.com) and the China site (aliyun.com).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegionsResponse
func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all regions in which Simple Application Server is supported.
//
// Description:
//
// The query results include all the Alibaba Cloud regions where Simple Application Server is supported on the international site (alibabacloud.com) and the China site (aliyun.com).
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListRegionsRequest
//
// @return ListRegionsResponse
func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Queries the information about snapshots that are created for a simple application server.
//
// Description:
//
// You can specify multiple request parameters that you want to query, such as `InstanceId`, `DiskId`, `SnapshotIds`, and `ResourceGroupId`. Specified parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListSnapshotsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshotsWithOptions(request *ListSnapshotsRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDiskType)) {
		query["SourceDiskType"] = request.SourceDiskType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSnapshots"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about snapshots that are created for a simple application server.
//
// Description:
//
// You can specify multiple request parameters that you want to query, such as `InstanceId`, `DiskId`, `SnapshotIds`, and `ResourceGroupId`. Specified parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
//
// ### [](#qps-)QPS limits
//
// You can call this operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you will experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ListSnapshotsRequest
//
// @return ListSnapshotsResponse
func (client *Client) ListSnapshots(request *ListSnapshotsRequest) (_result *ListSnapshotsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSnapshotsResponse{}
	_body, _err := client.ListSnapshotsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 云产品查标签接口
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 云产品查标签接口
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Logs on to a simple application server on Workbench.
//
// Description:
//
// After you create a simple application server, you can log on to the simple application server to build environments and applications on the server.
//
// @param request - LoginInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginInstanceResponse
func (client *Client) LoginInstanceWithOptions(request *LoginInstanceRequest, runtime *util.RuntimeOptions) (_result *LoginInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LoginInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LoginInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Logs on to a simple application server on Workbench.
//
// Description:
//
// After you create a simple application server, you can log on to the simple application server to build environments and applications on the server.
//
// @param request - LoginInstanceRequest
//
// @return LoginInstanceResponse
func (client *Client) LoginInstance(request *LoginInstanceRequest) (_result *LoginInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LoginInstanceResponse{}
	_body, _err := client.LoginInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the description of a Simple Database Service instance.
//
// Description:
//
// You can call this operation to modify the description of a Simple Database Service instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ModifyDatabaseInstanceDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseInstanceDescriptionResponse
func (client *Client) ModifyDatabaseInstanceDescriptionWithOptions(request *ModifyDatabaseInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceDescription)) {
		query["DatabaseInstanceDescription"] = request.DatabaseInstanceDescription
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseInstanceDescription"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseInstanceDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the description of a Simple Database Service instance.
//
// Description:
//
// You can call this operation to modify the description of a Simple Database Service instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ModifyDatabaseInstanceDescriptionRequest
//
// @return ModifyDatabaseInstanceDescriptionResponse
func (client *Client) ModifyDatabaseInstanceDescription(request *ModifyDatabaseInstanceDescriptionRequest) (_result *ModifyDatabaseInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseInstanceDescriptionResponse{}
	_body, _err := client.ModifyDatabaseInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a Simple Database Service instance.
//
// Description:
//
// After you create a Simple Database Service instance, you can view the parameters of the instance or modify the parameters of the instance based on your business requirements.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ModifyDatabaseInstanceParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyDatabaseInstanceParameterResponse
func (client *Client) ModifyDatabaseInstanceParameterWithOptions(request *ModifyDatabaseInstanceParameterRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseInstanceParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ForceRestart)) {
		query["ForceRestart"] = request.ForceRestart
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseInstanceParameter"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseInstanceParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the parameters of a Simple Database Service instance.
//
// Description:
//
// After you create a Simple Database Service instance, you can view the parameters of the instance or modify the parameters of the instance based on your business requirements.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ModifyDatabaseInstanceParameterRequest
//
// @return ModifyDatabaseInstanceParameterResponse
func (client *Client) ModifyDatabaseInstanceParameter(request *ModifyDatabaseInstanceParameterRequest) (_result *ModifyDatabaseInstanceParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseInstanceParameterResponse{}
	_body, _err := client.ModifyDatabaseInstanceParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the firewall rule of a simple application server.
//
// @param request - ModifyFirewallRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirewallRuleResponse
func (client *Client) ModifyFirewallRuleWithOptions(request *ModifyFirewallRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyFirewallRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		query["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleProtocol)) {
		query["RuleProtocol"] = request.RuleProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.SourceCidrIp)) {
		query["SourceCidrIp"] = request.SourceCidrIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFirewallRule"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFirewallRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the firewall rule of a simple application server.
//
// @param request - ModifyFirewallRuleRequest
//
// @return ModifyFirewallRuleResponse
func (client *Client) ModifyFirewallRule(request *ModifyFirewallRuleRequest) (_result *ModifyFirewallRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFirewallRuleResponse{}
	_body, _err := client.ModifyFirewallRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the firewall rule in a firewall template. You can apply the new firewall rule to simple application servers.
//
// Description:
//
// Modifying a firewall template does not affect the firewall rules that have been applied to simple application servers.
//
// @param request - ModifyFirewallTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirewallTemplateResponse
func (client *Client) ModifyFirewallTemplateWithOptions(request *ModifyFirewallTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyFirewallTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateId)) {
		query["FirewallTemplateId"] = request.FirewallTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.FirewallTemplateRule)) {
		query["FirewallTemplateRule"] = request.FirewallTemplateRule
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFirewallTemplate"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFirewallTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the firewall rule in a firewall template. You can apply the new firewall rule to simple application servers.
//
// Description:
//
// Modifying a firewall template does not affect the firewall rules that have been applied to simple application servers.
//
// @param request - ModifyFirewallTemplateRequest
//
// @return ModifyFirewallTemplateResponse
func (client *Client) ModifyFirewallTemplate(request *ModifyFirewallTemplateRequest) (_result *ModifyFirewallTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyFirewallTemplateResponse{}
	_body, _err := client.ModifyFirewallTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Shares or unshares a custom image to Elastic Compute Service (ECS).
//
// Description:
//
// You can share a custom image with ECS. If the configurations of your simple application server cannot meet your business requirements, or you want to use ECS instances to deploy your business, you can share your custom image with ECS to transfer your business from Simple Application Server to ECS.
//
// > The shared image in ECS resides in the same region as the custom image in Simple Application Server.
//
// You can unshare a custom image based on your business requirements or when you want to delete the custom image. Take note of the following items:
//
// 	- After you unshare a custom image, you cannot query or use the custom image in the ECS console or by calling ECS API operations.
//
// 	- After you unshare a custom image, you cannot re-initialize the disks of the ECS instances that were created based on the shared image.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ModifyImageShareStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyImageShareStatusResponse
func (client *Client) ModifyImageShareStatusWithOptions(request *ModifyImageShareStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyImageShareStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyImageShareStatus"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyImageShareStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Shares or unshares a custom image to Elastic Compute Service (ECS).
//
// Description:
//
// You can share a custom image with ECS. If the configurations of your simple application server cannot meet your business requirements, or you want to use ECS instances to deploy your business, you can share your custom image with ECS to transfer your business from Simple Application Server to ECS.
//
// > The shared image in ECS resides in the same region as the custom image in Simple Application Server.
//
// You can unshare a custom image based on your business requirements or when you want to delete the custom image. Take note of the following items:
//
// 	- After you unshare a custom image, you cannot query or use the custom image in the ECS console or by calling ECS API operations.
//
// 	- After you unshare a custom image, you cannot re-initialize the disks of the ECS instances that were created based on the shared image.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ModifyImageShareStatusRequest
//
// @return ModifyImageShareStatusResponse
func (client *Client) ModifyImageShareStatus(request *ModifyImageShareStatusRequest) (_result *ModifyImageShareStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageShareStatusResponse{}
	_body, _err := client.ModifyImageShareStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the VNC password of a simple application server.
//
// @param request - ModifyInstanceVncPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstanceVncPasswordResponse
func (client *Client) ModifyInstanceVncPasswordWithOptions(request *ModifyInstanceVncPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceVncPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VncPassword)) {
		query["VncPassword"] = request.VncPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceVncPassword"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceVncPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Changes the VNC password of a simple application server.
//
// @param request - ModifyInstanceVncPasswordRequest
//
// @return ModifyInstanceVncPasswordResponse
func (client *Client) ModifyInstanceVncPassword(request *ModifyInstanceVncPasswordRequest) (_result *ModifyInstanceVncPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceVncPasswordResponse{}
	_body, _err := client.ModifyInstanceVncPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a simple application server.
//
// Description:
//
//   Only simple application servers that are in the Running state can be restarted.
//
// 	- After you restart a simple application server, it enters the Starting state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - RebootInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootInstanceResponse
func (client *Client) RebootInstanceWithOptions(request *RebootInstanceRequest, runtime *util.RuntimeOptions) (_result *RebootInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a simple application server.
//
// Description:
//
//   Only simple application servers that are in the Running state can be restarted.
//
// 	- After you restart a simple application server, it enters the Starting state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - RebootInstanceRequest
//
// @return RebootInstanceResponse
func (client *Client) RebootInstance(request *RebootInstanceRequest) (_result *RebootInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootInstanceResponse{}
	_body, _err := client.RebootInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts simple application servers.
//
// @param request - RebootInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootInstancesResponse
func (client *Client) RebootInstancesWithOptions(request *RebootInstancesRequest, runtime *util.RuntimeOptions) (_result *RebootInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceReboot)) {
		query["ForceReboot"] = request.ForceReboot
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts simple application servers.
//
// @param request - RebootInstancesRequest
//
// @return RebootInstancesResponse
func (client *Client) RebootInstances(request *RebootInstancesRequest) (_result *RebootInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootInstancesResponse{}
	_body, _err := client.RebootInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of a Simple Database Service instance.
//
// Description:
//
// If you no longer need to use a public endpoint to access a Simple Database Service instance, you can release the public endpoint.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ReleasePublicConnectionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleasePublicConnectionResponse
func (client *Client) ReleasePublicConnectionWithOptions(request *ReleasePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleasePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleasePublicConnection"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleasePublicConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Releases the public endpoint of a Simple Database Service instance.
//
// Description:
//
// If you no longer need to use a public endpoint to access a Simple Database Service instance, you can release the public endpoint.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ReleasePublicConnectionRequest
//
// @return ReleasePublicConnectionResponse
func (client *Client) ReleasePublicConnection(request *ReleasePublicConnectionRequest) (_result *ReleasePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleasePublicConnectionResponse{}
	_body, _err := client.ReleasePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 取消共享镜像到用户
//
// @param request - RemoveCustomImageShareAccountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveCustomImageShareAccountResponse
func (client *Client) RemoveCustomImageShareAccountWithOptions(request *RemoveCustomImageShareAccountRequest, runtime *util.RuntimeOptions) (_result *RemoveCustomImageShareAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Account)) {
		query["Account"] = request.Account
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveCustomImageShareAccount"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveCustomImageShareAccountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 取消共享镜像到用户
//
// @param request - RemoveCustomImageShareAccountRequest
//
// @return RemoveCustomImageShareAccountResponse
func (client *Client) RemoveCustomImageShareAccount(request *RemoveCustomImageShareAccountRequest) (_result *RemoveCustomImageShareAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveCustomImageShareAccountResponse{}
	_body, _err := client.RemoveCustomImageShareAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a simple application server.
//
// Description:
//
//   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](https://help.aliyun.com/document_detail/58623.html).
//
// 	- When you call this operation to renew a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be renewed.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - RenewInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewInstanceResponse
func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Renews a simple application server.
//
// Description:
//
//   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](https://help.aliyun.com/document_detail/58623.html).
//
// 	- When you call this operation to renew a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be renewed.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - RenewInstanceRequest
//
// @return RenewInstanceResponse
func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password of the administrator account of a Simple Database Service instance.
//
// Description:
//
// If the password of your Simple Database Service instance is not strong, you can call this operation to change the password of the administrator account of the instance. To ensure security of the instance, we recommend that you regularly change the password of the instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ResetDatabaseAccountPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDatabaseAccountPasswordResponse
func (client *Client) ResetDatabaseAccountPasswordWithOptions(request *ResetDatabaseAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetDatabaseAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountPassword)) {
		query["AccountPassword"] = request.AccountPassword
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetDatabaseAccountPassword"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetDatabaseAccountPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password of the administrator account of a Simple Database Service instance.
//
// Description:
//
// If the password of your Simple Database Service instance is not strong, you can call this operation to change the password of the administrator account of the instance. To ensure security of the instance, we recommend that you regularly change the password of the instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ResetDatabaseAccountPasswordRequest
//
// @return ResetDatabaseAccountPasswordResponse
func (client *Client) ResetDatabaseAccountPassword(request *ResetDatabaseAccountPasswordRequest) (_result *ResetDatabaseAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDatabaseAccountPasswordResponse{}
	_body, _err := client.ResetDatabaseAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Rolls back a disk based on a snapshot.
//
// Description:
//
//   You can call this operation to roll back a disk only if the associated simple application server is in the Stopped state.
//
// 	- After a disk is rolled back, all data changes that are made from when the snapshot was created to when the disk is rolled back are lost. Back up disk data based on your needs before you roll back the disk.
//
// ### Precautions
//
// After you reset a simple application server, the disk data on the server is deleted. Snapshots created before the resetting operation are retained but cannot be used to roll back the disks of the server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ResetDiskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetDiskResponse
func (client *Client) ResetDiskWithOptions(request *ResetDiskRequest, runtime *util.RuntimeOptions) (_result *ResetDiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetDisk"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetDiskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Rolls back a disk based on a snapshot.
//
// Description:
//
//   You can call this operation to roll back a disk only if the associated simple application server is in the Stopped state.
//
// 	- After a disk is rolled back, all data changes that are made from when the snapshot was created to when the disk is rolled back are lost. Back up disk data based on your needs before you roll back the disk.
//
// ### Precautions
//
// After you reset a simple application server, the disk data on the server is deleted. Snapshots created before the resetting operation are retained but cannot be used to roll back the disks of the server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ResetDiskRequest
//
// @return ResetDiskResponse
func (client *Client) ResetDisk(request *ResetDiskRequest) (_result *ResetDiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDiskResponse{}
	_body, _err := client.ResetDiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets a simple application server.
//
// Description:
//
// You can reset a simple application server to re-install its application system or OS and re-initialize the server. You can reset a simple application server by resetting the current system or replacing the image.
//
// You can use one of the following methods to reset a simple application server:
//
// 	- Reset the current system. You can re-install the operating system without replacing the image.
//
// 	- Replace the image. You can select an Alibaba Cloud image or a custom image that is different from the existing image of the server to reinstall the OS of the server.
//
// ### Precautions
//
// 	- After you reset a simple application server, the disk data on the server is cleared. Back up the data as needed.
//
// 	- After you reset a simple application server, the monitoring operations that are performed on the server may fail. In this case, you can use one of the following methods to install the CloudMonitor agent on the server:
//
//     	- Connect to the server: For more information, see [Manually install the CloudMonitor agent for C++ on an ECS instance](https://help.aliyun.com/document_detail/183482.html).
//
//     	- Use Command Assistant: For more information, see [Use Command Assistant](https://help.aliyun.com/document_detail/438681.html). You can obtain the command that can be used to install CloudMonitor from the "Common commands" section of the [Use Command Assistant](https://help.aliyun.com/document_detail/438681.html) topic.
//
// ### Limits
//
// 	- Snapshots that are created before a server is reset are retained, but the snapshots cannot be used to roll back the disks of the server.
//
// 	- You cannot reset simple application servers that were created based on custom images that contain data of data disks.
//
// 	- Before you reset a simple application server by replacing the existing image with a custom image, take note of the following items:
//
//     	- The custom image must reside in the same region as the current server.
//
//     	- The custom image cannot be created based on the current server. If you want to recover the data on the server, you can use a snapshot of the server to roll back the disks of the server.
//
//     	- If your simple application server resides outside the Chinese mainland, you cannot switch the OS of the server between Windows Server and Linux. You cannot use a Windows Server custom image to reset a Linux simple application server. You also cannot use a Linux custom image to reset a Windows Server simple application server. You can switch the OSs of simple application servers only between Windows Server OSs or between Linux distributions.
//
//     	- The following limits apply to the disks attached to the simple application server:
//
//         	- If the custom image contains a system disk and a data disk but only a system disk is attached to the simple application server and no data disk is attached, you cannot use the custom image to reset the simple application server.
//
//         	- If the system disk size of the custom image is greater than the system disk size of the simple application server, you cannot directly use the custom image to reset the simple application server.
//
//         	- Only if the system disk size of the simple application server is greater than or equal to the system disk size of the custom image, you can use the custom image to reset the simple application server. To increase the system disk size of your simple application server, you can upgrade the server. For more information, see Upgrade a simple application server.
//
//         	- If the data disk size of the custom image is greater than the data disk size of the simple application server, you cannot use the custom image to reset the simple application server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ResetSystemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetSystemResponse
func (client *Client) ResetSystemWithOptions(request *ResetSystemRequest, runtime *util.RuntimeOptions) (_result *ResetSystemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSystem"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSystemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets a simple application server.
//
// Description:
//
// You can reset a simple application server to re-install its application system or OS and re-initialize the server. You can reset a simple application server by resetting the current system or replacing the image.
//
// You can use one of the following methods to reset a simple application server:
//
// 	- Reset the current system. You can re-install the operating system without replacing the image.
//
// 	- Replace the image. You can select an Alibaba Cloud image or a custom image that is different from the existing image of the server to reinstall the OS of the server.
//
// ### Precautions
//
// 	- After you reset a simple application server, the disk data on the server is cleared. Back up the data as needed.
//
// 	- After you reset a simple application server, the monitoring operations that are performed on the server may fail. In this case, you can use one of the following methods to install the CloudMonitor agent on the server:
//
//     	- Connect to the server: For more information, see [Manually install the CloudMonitor agent for C++ on an ECS instance](https://help.aliyun.com/document_detail/183482.html).
//
//     	- Use Command Assistant: For more information, see [Use Command Assistant](https://help.aliyun.com/document_detail/438681.html). You can obtain the command that can be used to install CloudMonitor from the "Common commands" section of the [Use Command Assistant](https://help.aliyun.com/document_detail/438681.html) topic.
//
// ### Limits
//
// 	- Snapshots that are created before a server is reset are retained, but the snapshots cannot be used to roll back the disks of the server.
//
// 	- You cannot reset simple application servers that were created based on custom images that contain data of data disks.
//
// 	- Before you reset a simple application server by replacing the existing image with a custom image, take note of the following items:
//
//     	- The custom image must reside in the same region as the current server.
//
//     	- The custom image cannot be created based on the current server. If you want to recover the data on the server, you can use a snapshot of the server to roll back the disks of the server.
//
//     	- If your simple application server resides outside the Chinese mainland, you cannot switch the OS of the server between Windows Server and Linux. You cannot use a Windows Server custom image to reset a Linux simple application server. You also cannot use a Linux custom image to reset a Windows Server simple application server. You can switch the OSs of simple application servers only between Windows Server OSs or between Linux distributions.
//
//     	- The following limits apply to the disks attached to the simple application server:
//
//         	- If the custom image contains a system disk and a data disk but only a system disk is attached to the simple application server and no data disk is attached, you cannot use the custom image to reset the simple application server.
//
//         	- If the system disk size of the custom image is greater than the system disk size of the simple application server, you cannot directly use the custom image to reset the simple application server.
//
//         	- Only if the system disk size of the simple application server is greater than or equal to the system disk size of the custom image, you can use the custom image to reset the simple application server. To increase the system disk size of your simple application server, you can upgrade the server. For more information, see Upgrade a simple application server.
//
//         	- If the data disk size of the custom image is greater than the data disk size of the simple application server, you cannot use the custom image to reset the simple application server.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - ResetSystemRequest
//
// @return ResetSystemResponse
func (client *Client) ResetSystem(request *ResetSystemRequest) (_result *ResetSystemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSystemResponse{}
	_body, _err := client.ResetSystemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts a Simple Database Service instance.
//
// Description:
//
// You can call this operation to restart a Simple Database Service instance that is in the Running state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - RestartDatabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartDatabaseInstanceResponse
func (client *Client) RestartDatabaseInstanceWithOptions(request *RestartDatabaseInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDatabaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartDatabaseInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartDatabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts a Simple Database Service instance.
//
// Description:
//
// You can call this operation to restart a Simple Database Service instance that is in the Running state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - RestartDatabaseInstanceRequest
//
// @return RestartDatabaseInstanceResponse
func (client *Client) RestartDatabaseInstance(request *RestartDatabaseInstanceRequest) (_result *RestartDatabaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDatabaseInstanceResponse{}
	_body, _err := client.RestartDatabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Runs commands on a simple application server.
//
// Description:
//
// Command Assistant is an automated O\\&M tool for Simple Application Server. You can maintain simple application servers by running shell, PowerShell, and batch commands in the Simple Application Server console without remotely logging on to the servers.
//
// Before you use Command Assistant, take note of the following items:
//
// 	- The simple application server must be in the Running state.
//
// 	- The Cloud Assistant client is installed on the server. By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall it. For more information, see [Install the Cloud Assistant Agent](https://help.aliyun.com/document_detail/64921.html).
//
// @param tmpReq - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithOptions(tmpReq *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		query["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.EnableParameter)) {
		query["EnableParameter"] = request.EnableParameter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.WindowsPasswordName)) {
		query["WindowsPasswordName"] = request.WindowsPasswordName
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingDir)) {
		query["WorkingDir"] = request.WorkingDir
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingUser)) {
		query["WorkingUser"] = request.WorkingUser
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Runs commands on a simple application server.
//
// Description:
//
// Command Assistant is an automated O\\&M tool for Simple Application Server. You can maintain simple application servers by running shell, PowerShell, and batch commands in the Simple Application Server console without remotely logging on to the servers.
//
// Before you use Command Assistant, take note of the following items:
//
// 	- The simple application server must be in the Running state.
//
// 	- The Cloud Assistant client is installed on the server. By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall it. For more information, see [Install the Cloud Assistant Agent](https://help.aliyun.com/document_detail/64921.html).
//
// @param request - RunCommandRequest
//
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a Simple Database Service instance.
//
// Description:
//
// You can call this operation to start a Simple Database Service instance that is in the Stopped state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StartDatabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartDatabaseInstanceResponse
func (client *Client) StartDatabaseInstanceWithOptions(request *StartDatabaseInstanceRequest, runtime *util.RuntimeOptions) (_result *StartDatabaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartDatabaseInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartDatabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a Simple Database Service instance.
//
// Description:
//
// You can call this operation to start a Simple Database Service instance that is in the Stopped state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StartDatabaseInstanceRequest
//
// @return StartDatabaseInstanceResponse
func (client *Client) StartDatabaseInstance(request *StartDatabaseInstanceRequest) (_result *StartDatabaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDatabaseInstanceResponse{}
	_body, _err := client.StartDatabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a simple application server.
//
// Description:
//
// You can call this operation to start a simple application server that is in the Stopped state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StartInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstanceResponse
func (client *Client) StartInstanceWithOptions(request *StartInstanceRequest, runtime *util.RuntimeOptions) (_result *StartInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a simple application server.
//
// Description:
//
// You can call this operation to start a simple application server that is in the Stopped state.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StartInstanceRequest
//
// @return StartInstanceResponse
func (client *Client) StartInstance(request *StartInstanceRequest) (_result *StartInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstanceResponse{}
	_body, _err := client.StartInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts simple application servers.
//
// @param request - StartInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartInstancesResponse
func (client *Client) StartInstancesWithOptions(request *StartInstancesRequest, runtime *util.RuntimeOptions) (_result *StartInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts simple application servers.
//
// @param request - StartInstancesRequest
//
// @return StartInstancesResponse
func (client *Client) StartInstances(request *StartInstancesRequest) (_result *StartInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstancesResponse{}
	_body, _err := client.StartInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a session for a simple application server.
//
// @param request - StartTerminalSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartTerminalSessionResponse
func (client *Client) StartTerminalSessionWithOptions(request *StartTerminalSessionRequest, runtime *util.RuntimeOptions) (_result *StartTerminalSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartTerminalSession"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTerminalSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a session for a simple application server.
//
// @param request - StartTerminalSessionRequest
//
// @return StartTerminalSessionResponse
func (client *Client) StartTerminalSession(request *StartTerminalSessionRequest) (_result *StartTerminalSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartTerminalSessionResponse{}
	_body, _err := client.StartTerminalSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a Simple Database Service instance.
//
// Description:
//
// You can call this operation to stop a Simple Database Service instance that is in the Running state. After the instance is stopped, you cannot log on to or access the instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StopDatabaseInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopDatabaseInstanceResponse
func (client *Client) StopDatabaseInstanceWithOptions(request *StopDatabaseInstanceRequest, runtime *util.RuntimeOptions) (_result *StopDatabaseInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInstanceId)) {
		query["DatabaseInstanceId"] = request.DatabaseInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopDatabaseInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopDatabaseInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a Simple Database Service instance.
//
// Description:
//
// You can call this operation to stop a Simple Database Service instance that is in the Running state. After the instance is stopped, you cannot log on to or access the instance.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StopDatabaseInstanceRequest
//
// @return StopDatabaseInstanceResponse
func (client *Client) StopDatabaseInstance(request *StopDatabaseInstanceRequest) (_result *StopDatabaseInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDatabaseInstanceResponse{}
	_body, _err := client.StopDatabaseInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a simple application server.
//
// Description:
//
// You can stop a simple application server that you do not use for the time being.
//
// >  Stopping a simple application server may interrupt your business. We recommend that you perform the stop operation during off-peak hours.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StopInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstanceResponse
func (client *Client) StopInstanceWithOptions(request *StopInstanceRequest, runtime *util.RuntimeOptions) (_result *StopInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a simple application server.
//
// Description:
//
// You can stop a simple application server that you do not use for the time being.
//
// >  Stopping a simple application server may interrupt your business. We recommend that you perform the stop operation during off-peak hours.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - StopInstanceRequest
//
// @return StopInstanceResponse
func (client *Client) StopInstance(request *StopInstanceRequest) (_result *StopInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Stops simple application servers.
//
// @param request - StopInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInstancesResponse
func (client *Client) StopInstancesWithOptions(request *StopInstancesRequest, runtime *util.RuntimeOptions) (_result *StopInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ForceStop)) {
		query["ForceStop"] = request.ForceStop
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstances"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops simple application servers.
//
// @param request - StopInstancesRequest
//
// @return StopInstancesResponse
func (client *Client) StopInstances(request *StopInstancesRequest) (_result *StopInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstancesResponse{}
	_body, _err := client.StopInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 打标签接口
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// 删标签接口
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删标签接口
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// Modifys the parameter of a command.
//
// @param request - UpdateCommandAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCommandAttributeResponse
func (client *Client) UpdateCommandAttributeWithOptions(request *UpdateCommandAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateCommandAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		query["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		query["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingDir)) {
		query["WorkingDir"] = request.WorkingDir
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCommandAttribute"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCommandAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifys the parameter of a command.
//
// @param request - UpdateCommandAttributeRequest
//
// @return UpdateCommandAttributeResponse
func (client *Client) UpdateCommandAttribute(request *UpdateCommandAttributeRequest) (_result *UpdateCommandAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCommandAttributeResponse{}
	_body, _err := client.UpdateCommandAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the remarks for the data disk that is attached to a simple application server.
//
// @param request - UpdateDiskAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDiskAttributeResponse
func (client *Client) UpdateDiskAttributeWithOptions(request *UpdateDiskAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateDiskAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DiskId)) {
		query["DiskId"] = request.DiskId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDiskAttribute"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDiskAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the remarks for the data disk that is attached to a simple application server.
//
// @param request - UpdateDiskAttributeRequest
//
// @return UpdateDiskAttributeResponse
func (client *Client) UpdateDiskAttribute(request *UpdateDiskAttributeRequest) (_result *UpdateDiskAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDiskAttributeResponse{}
	_body, _err := client.UpdateDiskAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information of a simple application server, including the server name and the password that you use to log on to the server.
//
// Description:
//
// ## Usage notes
//
// After you change the password of a simple application server, you must restart the server by calling the [RebootInstance](https://help.aliyun.com/document_detail/190443.html) operation to allow the new password to take effect.
//
// ### QPS limits
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - UpdateInstanceAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceAttributeResponse
func (client *Client) UpdateInstanceAttributeWithOptions(request *UpdateInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceAttribute"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information of a simple application server, including the server name and the password that you use to log on to the server.
//
// Description:
//
// ## Usage notes
//
// After you change the password of a simple application server, you must restart the server by calling the [RebootInstance](https://help.aliyun.com/document_detail/190443.html) operation to allow the new password to take effect.
//
// ### QPS limits
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - UpdateInstanceAttributeRequest
//
// @return UpdateInstanceAttributeResponse
func (client *Client) UpdateInstanceAttribute(request *UpdateInstanceAttributeRequest) (_result *UpdateInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.UpdateInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the remarks of a snapshot of a simple application server.
//
// @param request - UpdateSnapshotAttributeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSnapshotAttributeResponse
func (client *Client) UpdateSnapshotAttributeWithOptions(request *UpdateSnapshotAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateSnapshotAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Remark)) {
		query["Remark"] = request.Remark
	}

	if !tea.BoolValue(util.IsUnset(request.SnapshotId)) {
		query["SnapshotId"] = request.SnapshotId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSnapshotAttribute"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSnapshotAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the remarks of a snapshot of a simple application server.
//
// @param request - UpdateSnapshotAttributeRequest
//
// @return UpdateSnapshotAttributeResponse
func (client *Client) UpdateSnapshotAttribute(request *UpdateSnapshotAttributeRequest) (_result *UpdateSnapshotAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSnapshotAttributeResponse{}
	_body, _err := client.UpdateSnapshotAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades the plan for a simple application server.
//
// Description:
//
//   The plan of a simple application server cannot be downgraded, but can only be upgraded. For more information about plans, see [Billable items](https://help.aliyun.com/document_detail/58623.html).
//
// 	- When you call this operation to upgrade a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be upgraded.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - UpgradeInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeInstanceResponse
func (client *Client) UpgradeInstanceWithOptions(request *UpgradeInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeInstance"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades the plan for a simple application server.
//
// Description:
//
//   The plan of a simple application server cannot be downgraded, but can only be upgraded. For more information about plans, see [Billable items](https://help.aliyun.com/document_detail/58623.html).
//
// 	- When you call this operation to upgrade a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be upgraded.
//
// ### QPS limit
//
// You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](https://help.aliyun.com/document_detail/347607.html).
//
// @param request - UpgradeInstanceRequest
//
// @return UpgradeInstanceResponse
func (client *Client) UpgradeInstance(request *UpgradeInstanceRequest) (_result *UpgradeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeInstanceResponse{}
	_body, _err := client.UpgradeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a key pair for a simple application server.
//
// @param request - UploadInstanceKeyPairRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadInstanceKeyPairResponse
func (client *Client) UploadInstanceKeyPairWithOptions(request *UploadInstanceKeyPairRequest, runtime *util.RuntimeOptions) (_result *UploadInstanceKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKey)) {
		query["PublicKey"] = request.PublicKey
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadInstanceKeyPair"),
		Version:     tea.String("2020-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadInstanceKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a key pair for a simple application server.
//
// @param request - UploadInstanceKeyPairRequest
//
// @return UploadInstanceKeyPairResponse
func (client *Client) UploadInstanceKeyPair(request *UploadInstanceKeyPairRequest) (_result *UploadInstanceKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UploadInstanceKeyPairResponse{}
	_body, _err := client.UploadInstanceKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
