// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AllocatePublicConnectionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	PublicConnection *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AllocatePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateCommandRequest struct {
	CommandContent  *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableParameter *bool   `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Timeout         *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkingDir      *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
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

type CreateCommandResponseBody struct {
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data disk snapshot.
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The description of the custom image.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the custom image. The name must be 2 to 128 characters in length, and can contain letters, digits, colons (:), underscores (\_), and hyphens (-). The name must start with a letter or a digit. This parameter is empty by default.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the database. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the system disk snapshot.
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
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

func (s *CreateCustomImageRequest) SetSystemSnapshotId(v string) *CreateCustomImageRequest {
	s.SystemSnapshotId = &v
	return s
}

type CreateCustomImageResponseBody struct {
	// The custom image ID.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The request ID.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range. Valid values: 165535. Specify a port range in the format of \<start port number>/\<end port number>. Example: 1024/1055, which indicates the port range of 10241055.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the simple application server.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// *   TCP: the TCP protocol
	// *   UDP: the UDP protocol
	// *   TCP+UDP: the TCP and UDP protocols
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
	FirewallId *string `json:"FirewallId,omitempty" xml:"FirewallId,omitempty"`
	// The request ID.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The remarks of the firewall rule.
	FirewallRules []*CreateFirewallRulesRequestFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type CreateFirewallRulesRequestFirewallRules struct {
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
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

type CreateFirewallRulesShrinkRequest struct {
	// The client token.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The remarks of the firewall rule.
	FirewallRulesShrink *string `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type CreateFirewallRulesResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFirewallRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirewallRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirewallRulesResponseBody) SetRequestId(v string) *CreateFirewallRulesResponseBody {
	s.RequestId = &v
	return s
}

type CreateFirewallRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the key pair.
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the key pair.
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The private key.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// *   true
	// *   false
	//
	// Default value: false.
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. This parameter is required only when you set `AutoRenew` to true. Unit: months. Valid values: 1, 3, 6, 12, 24, and 36.
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The billing method of the simple application servers. Set the value to PrePaid, which indicates the subscription billing method.
	//
	// Default value: PrePaid.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The size of the data disk that is attached to the server. Unit: GB. Valid values: 0 to 16380. The value must be an integral multiple of 20.
	//
	// *   A value of 0 indicates that no data disk is attached.
	// *   If the disk included in the specified plan is a standard SSD, the data disk must be 20 GB or larger in size.
	//
	// Default value: 0.
	DataDiskSize *int64 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The image ID. You can call the [ListImages](~~189313~~) operation to query the available images in the specified region.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The subscription period of the servers. Unit: months. Valid values: 1, 3, 6, 12, 24, and 36.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The plan ID. You can call the [ListPlans](~~189314~~) operation to query all plans provided by Simple Application Server in the specified region.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateSnapshotRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk ID.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server to which the disk is attached.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot name. The name must be 2 to 50 characters in length. It must start with a letter but cannot start with `http://` or `https://`. The name can only contain letters, digits, colons (:), underscores (\_), periods (.), and hyphens (-).
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
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

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

type CreateSnapshotResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The snapshot ID.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The custom image ID.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID of the custom image. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCustomImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the firewall rule.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteSnapshotRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the snapshot.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSnapshotResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple snapshot IDs with commas (,).
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeCloudAssistantStatusRequest struct {
	// The IDs of the simple application servers.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the Cloud Assistant client is installed on the server.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudAssistantStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	// Indicates whether the Cloud Monitor agent was automatically installed on the simple application server. Valid values:
	//
	// *   true
	// *   false
	InstanceStatusList []*DescribeCloudMonitorAgentStatusesResponseBodyInstanceStatusList `json:"InstanceStatusList,omitempty" xml:"InstanceStatusList,omitempty" type:"Repeated"`
	// The request ID.
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
	AutoInstall *bool   `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCloudMonitorAgentStatusesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CommandId        *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	CommandName      *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	CommandType      *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId         *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	PageNumber       *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	CommandInvocations []*DescribeCommandInvocationsResponseBodyCommandInvocations `json:"CommandInvocations,omitempty" xml:"CommandInvocations,omitempty" type:"Repeated"`
	PageNumber         *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CommandContent     *string                                                                    `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandDescription *string                                                                    `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	CommandId          *string                                                                    `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	CommandName        *string                                                                    `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	CommandType        *string                                                                    `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	CreationTime       *string                                                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	InvocationStatus   *string                                                                    `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	InvokeId           *string                                                                    `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	InvokeInstances    []*DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Repeated"`
	Parameters         *string                                                                    `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Timeout            *int64                                                                     `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Username           *string                                                                    `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkingDir         *string                                                                    `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
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
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorInfo        *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	ExitCode         *int64  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	FinishTime       *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	Output           *string `json:"Output,omitempty" xml:"Output,omitempty"`
	StartTime        *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCommandInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CommandId  *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Provider   *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *DescribeCommandsRequest) SetType(v string) *DescribeCommandsRequest {
	s.Type = &v
	return s
}

type DescribeCommandsResponseBody struct {
	Commands   []*DescribeCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CommandContent       *string                                                     `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandId            *string                                                     `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	CreationTime         *string                                                     `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description          *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableParameter      *bool                                                       `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	Name                 *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	ParameterDefinitions []*DescribeCommandsResponseBodyCommandsParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Repeated"`
	ParameterNames       []*string                                                   `json:"ParameterNames,omitempty" xml:"ParameterNames,omitempty" type:"Repeated"`
	Provider             *string                                                     `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Timeout              *int64                                                      `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Type                 *string                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkingDir           *string                                                     `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
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
	DefaultValue   *string   `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ParameterName  *string   `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	PossibleValues []*string `json:"PossibleValues,omitempty" xml:"PossibleValues,omitempty" type:"Repeated"`
	Required       *bool     `json:"Required,omitempty" xml:"Required,omitempty"`
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

type DescribeCommandsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC. The end time must be later than the start time.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Simple Database Service instance.
	//
	// You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
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
	// The time when the error log entry was generated. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	ErrorLogs []*DescribeDatabaseErrorLogsResponseBodyErrorLogs `json:"ErrorLogs,omitempty" xml:"ErrorLogs,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ErrorInfo  *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseErrorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the metric. Valid values:
	//
	// *   MySQL_MemCpuUsage: The CPU utilization and memory usage of the instance within the entire operating system.
	// *   MySQL_DetailedSpaceUsage: The total space usage, data space, log space, temporary space, and system space of the instance.
	// *   MySQL_Sessions : The total number of active connections.
	// *   MySQL_IOPS: The IOPS of the instance.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mmZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
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
	// *   cpuusage\&memusage
	// *   active_session\&total_session
	// *   ins_size\&data_size\&log_size\&tmp_size\&other_size
	// *   io
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	// The monitoring data.
	MetricData *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	// The name of the metric. Valid values:
	//
	// *   MySQL_MemCpuUsage: The CPU utilization and memory usage of the instance within the entire operating system.
	// *   MySQL_DetailedSpaceUsage: The total space usage, data space, log space, temporary space, and system space of the instance.
	// *   MySQL_Sessions : The total number of active connections.
	// *   MySQL_IOPS: The IOPS of the instance.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unit of the monitoring metric.
	//
	// *   %
	// *   int
	// *   MB
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseInstanceMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine. Valid values:
	//
	// *   5.7: MySQL 5.7.
	// *   8.0: MySQL 8.0.
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The request ID.
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
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ForceModify          *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	ForceRestart         *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
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
	CheckingCode         *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	ForceModify          *string `json:"ForceModify,omitempty" xml:"ForceModify,omitempty"`
	ForceRestart         *string `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
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
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseInstanceParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DatabaseInstanceIds *string `json:"DatabaseInstanceIds,omitempty" xml:"DatabaseInstanceIds,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Simple Database Service instances.
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
	// The name of the super administrator account of the Simple Database Service instance.
	DatabaseInstances []*DescribeDatabaseInstancesResponseBodyDatabaseInstances `json:"DatabaseInstances,omitempty" xml:"DatabaseInstances,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	BusinessStatus          *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType              *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Cpu                     *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreationTime            *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DatabaseInstanceEdition *string `json:"DatabaseInstanceEdition,omitempty" xml:"DatabaseInstanceEdition,omitempty"`
	DatabaseInstanceId      *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	DatabaseInstanceName    *string `json:"DatabaseInstanceName,omitempty" xml:"DatabaseInstanceName,omitempty"`
	DatabaseInstanceStatus  *string `json:"DatabaseInstanceStatus,omitempty" xml:"DatabaseInstanceStatus,omitempty"`
	DatabaseVersion         *string `json:"DatabaseVersion,omitempty" xml:"DatabaseVersion,omitempty"`
	ExpiredTime             *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Memory                  *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PrivateConnection       *string `json:"PrivateConnection,omitempty" xml:"PrivateConnection,omitempty"`
	PublicConnection        *string `json:"PublicConnection,omitempty" xml:"PublicConnection,omitempty"`
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Storage                 *int32  `json:"Storage,omitempty" xml:"Storage,omitempty"`
	SuperAccountName        *string `json:"SuperAccountName,omitempty" xml:"SuperAccountName,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AcsProduct *string `json:"AcsProduct,omitempty" xml:"AcsProduct,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The end of the time range to query. The end time must be later than the start time. The interval between the start time and the end time must be less than 7 days.
	//
	// Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30 to 100.
	//
	// Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The time displayed in the Simple Application Server console is in the format of UTC+8.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDatabaseSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabaseSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseSlowLogRecordsRequest) SetAcsProduct(v string) *DescribeDatabaseSlowLogRecordsRequest {
	s.AcsProduct = &v
	return s
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
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 30 to 100.
	//
	// Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The number of logical reads.
	PhysicalIORead *int64 `json:"PhysicalIORead,omitempty" xml:"PhysicalIORead,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The database name.
	SlowLogs []*DescribeDatabaseSlowLogRecordsResponseBodySlowLogs `json:"SlowLogs,omitempty" xml:"SlowLogs,omitempty" type:"Repeated"`
	// The total number of entries returned.
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
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty"`
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty"`
	LockTimes          *int64  `json:"LockTimes,omitempty" xml:"LockTimes,omitempty"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty"`
	QueryTimeMS        *int64  `json:"QueryTimeMS,omitempty" xml:"QueryTimeMS,omitempty"`
	QueryTimes         *int64  `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabaseSlowLogRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeInstanceKeyPairRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the key pair.
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	InstancePasswordSetting *bool `json:"InstancePasswordSetting,omitempty" xml:"InstancePasswordSetting,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether a VNC connection password is set.
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
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstancePasswordsSettingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The VNC connection address of the server.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceVncUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution ID. You can call the [DescribeInvocations](~~439368~~) operation to query execution IDs.
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The region ID of the simple application server. You can call the [DescribeRegions](~~25609~~) operation to query the most recent region list.
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
	// *   If this parameter is empty, the command is executed normally.
	// *   InstanceNotExists: The specified server does not exist or is released.
	// *   InstanceReleased: The server was released while the command was being executed on the server.
	// *   InstanceNotRunning: The server is not in the Running state while the command is being executed.
	// *   CommandNotApplicable: The command is not applicable to the specified server.
	// *   AccountNotExists: The specified account does not exist.
	// *   DirectoryNotExists: The specified directory does not exist.
	// *   BadCronExpression: The specified cron expression for the execution schedule is invalid.
	// *   ClientNotRunning: The Cloud Assistant client is not running.
	// *   ClientNotResponse: The Cloud Assistant client does not respond.
	// *   ClientIsUpgrading: The Cloud Assistant client is being upgraded.
	// *   ClientNeedUpgrade: The Cloud Assistant client needs to be upgraded.
	// *   DeliveryTimeout: Command sending times out.
	// *   ExecutionTimeout: The execution times out.
	// *   ExecutionException: An exception occurs while the command is being executed.
	// *   ExecutionInterrupted: The execution is interrupted.
	// *   ExitCodeNonzero: The execution is complete, but the exit code is not 0.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the command is not successfully sent or executed. Valid values:
	//
	// *   If this parameter is empty, the command is executed normally.
	// *   the specified instance does not exists: The specified server does not exist or is released.
	// *   the instance has released when create task: The server was released while the command was being executed on the server.
	// *   the instance is not running when create task: The server is not in the Running state while the command is being executed.
	// *   the command is not applicable: The command is not applicable to the specified server.
	// *   the specified account does not exists: The specified account does not exist.
	// *   the specified directory does not exists: The specified directory does not exist.
	// *   the cron job expression is invalid: The specified cron expression is invalid.
	// *   the aliyun service is not running on the instance: The Cloud Assistance client is not running.
	// *   the aliyun service in the instance does not response: The Cloud Assistant client does not respond to your request.
	// *   the aliyun service in the instance is upgrading now: The Cloud Assistant client is being upgraded.
	// *   the aliyun service in the instance need upgrade: The Cloud Assistant client needs to be upgraded.
	// *   the command delivery has been timeout: Command sending times out.
	// *   the command execution has been timeout: The execution times out.
	// *   the command execution got an exception: An exception occurs while the command is being executed.
	// *   the command execution has been interrupted: The execution is interrupted.
	// *   the command execution exit code is not zero: The execution is complete, and the exit code is not 0.
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command.
	//
	// *   For Linux instances, the exit code is the exit code of the shell command.
	// *   For Windows instances, the exit code is the exit code of the batch or PowerShell command.
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the execution ended.
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the execution progress. Valid values:
	//
	// *   Pending: The command is being verified or sent.
	// *   Invalid: The specified command type or parameter is invalid.
	// *   Aborted: The command fails to be sent to the server. To send a command to a server, make sure that the server is in the Running state and the command can be sent within 1 minute.
	// *   Running: The command is being executed on the server.
	// *   Success: The execution is completed, and the exit code is 0.
	// *   Failed: The execution is completed, and the exit code is not 0.
	// *   Error: The execution cannot proceed due to an exception.
	// *   Timeout: The execution times out.
	// *   Cancelled: The execution is canceled, and the command is not executed.
	// *   Stopping: The command in the Running state is being stopped.
	// *   Terminated: The command is terminated while it is being executed.
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The execution ID.
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The status of the execution. Valid values:
	//
	// *   Running
	// *   Finished
	// *   Failed
	// *   Stopped
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The username who executes the command on the simple application server.
	InvokeUser *string `json:"InvokeUser,omitempty" xml:"InvokeUser,omitempty"`
	// The command output.
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the execution started.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInvocationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the command execution. Valid values:
	//
	// *   Running: The command is being executed.
	// *   Finished: The execution is complete.
	// *   Failed: The execution fails.
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The name of the command.
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The type of the command. Valid values:
	//
	// *   RunBatScript: batch command (applicable to Windows instances).
	// *   RunPowerShellScript: PowerShell command (applicable to Windows instances).
	// *   RunShellScript: shell command (applicable to Linux instances).
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The time when the command was created.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The status of the command. Valid values:
	//
	// *   Pending: The command is being verified or sent.
	// *   Invalid: The specified command type or parameter is invalid.
	// *   Aborted: The command failed to be sent. To send a command to an instance, make sure that the instance is in the Running state and the command is sent to the instance within 1 minute.
	// *   Running: The command is being run on the instance.
	// *   Success: The command finishes running, and the exit code is 0.
	// *   Failed: The command finishes running, but the exit code is not 0.
	// *   Error: The running of the command cannot proceed due to an exception.
	// *   Timeout: The running of the command times out.
	// *   Cancelled: The running is canceled, and the command is not run.
	// *   Stopping: The command that is running is being stopped.
	// *   Terminated: The command is terminated while it is being run.
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The ID of the command task.
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The status of the command. Valid values:
	//
	// *   Running: The command is running.
	// *   Finished: The command finishes running.
	// *   Failed: The running of the command failed.
	// *   Stopped: The running is stopped.
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The custom parameters in the command. If no custom parameter exists in the command, the default value is {}.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end of the time range to query. The following formats are supported:
	//
	// *   UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 January 1, 1970.
	// *   Time format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > The interval between the start time and the end time is less than or equal to 31 days.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Valid values: 1 to 1440.
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric name. Valid values:
	//
	// *   MEMORY_ACTUALUSEDSPACE: the memory usage. Unit: bytes.
	// *   DISKUSAGE_USED: the disk usage. Unit: bytes.
	// *   CPU_UTILIZATION: the CPU utilization in percentage.
	// *   VPC_PUBLICIP_INTERNETOUT_RATE: the outbound bandwidth rate of the network. Unit: bits/s.
	// *   VPC_PUBLICIP_INTERNETIN_RATE: the inbound bandwidth rate of the network. Unit: bits/s.
	// *   DISK_READ_IOPS: the read IOPS of the disk. Unit: count/s.
	// *   DISK_WRITE_IOPS: the write IOPS of the disk. Unit: count/s.
	// *   FLOW_USED: the traffic usage. Unit: bytes.
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
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
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The following formats are supported:
	//
	// *   UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 January 1, 1970.
	// *   Time format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > The specified time range includes the end time and excludes the start time. The start time must be earlier than the end time.
	//
	// The interval between the start time and the end time is less than or equal to 31 days.
	//
	// **
	//
	// ****
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
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
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
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMonitorDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	// *   pause: The Security Center agent suspends protection for your server.
	// *   online: The Security Center agent is protecting your server.
	// *   offline: The Security Center agent does not protect your server.
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The request ID.
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSecurityAgentStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DisableFirewallRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule. You can call the ListFirewallRules operation to query the ID of the firewall rule.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The IP address or CIDR block that is allowed in the firewall policy.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type InstallCloudAssistantRequest struct {
	// The IDs of the simple application servers.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallCloudAssistantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly install the CloudMonitor agent. Valid values:
	//
	// *   true (default value): forcibly installs the CloudMonitor agent.
	// *   false: does not forcibly install the CloudMonitor agent.
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallCloudMonitorAgentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CommandId   *string                `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceIds *string                `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Parameters  map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId    *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Username    *string                `json:"Username,omitempty" xml:"Username,omitempty"`
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
	CommandId        *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Username         *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InvokeCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListCustomImagesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data disk snapshot.
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The image IDs of the simple application server. The value can be a JSON array that consists of up to 100 image IDs. Separate multiple image IDs with commas (,).
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The image names of the simple application servers. The value can be a JSON array that consists of up to 100 image names. Separate multiple image names with commas (,).
	ImageNames *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	// The page number. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// *   Maximum value: 100.
	// *   Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers corresponding to the custom images. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the system disk snapshot.
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
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

func (s *ListCustomImagesRequest) SetSystemSnapshotId(v string) *ListCustomImagesRequest {
	s.SystemSnapshotId = &v
	return s
}

type ListCustomImagesResponseBody struct {
	// The queried custom images.
	CustomImages []*ListCustomImagesResponseBodyCustomImages `json:"CustomImages,omitempty" xml:"CustomImages,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	// The time when the snapshot was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the data disk snapshot.
	DataSnapshotId *string `json:"DataSnapshotId,omitempty" xml:"DataSnapshotId,omitempty"`
	// The name of the data disk snapshot.
	DataSnapshotName *string `json:"DataSnapshotName,omitempty" xml:"DataSnapshotName,omitempty"`
	// The description of the custom image.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the custom image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Indicates whether the custom image is shared with Elastic Compute Service (ECS).
	InShare *bool `json:"InShare,omitempty" xml:"InShare,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The name of the custom image.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID of the custom images.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the custom image.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the system disk snapshot.
	SystemSnapshotId *string `json:"SystemSnapshotId,omitempty" xml:"SystemSnapshotId,omitempty"`
	// The name of the system disk snapshot.
	SystemSnapshotName *string `json:"SystemSnapshotName,omitempty" xml:"SystemSnapshotName,omitempty"`
}

func (s ListCustomImagesResponseBodyCustomImages) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyCustomImages) GoString() string {
	return s.String()
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

func (s *ListCustomImagesResponseBodyCustomImages) SetRegionId(v string) *ListCustomImagesResponseBodyCustomImages {
	s.RegionId = &v
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

type ListCustomImagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The type of the disk. Valid values:
	//
	// *   System: system disk.
	// *   Data: data disk.
	//
	// By default, system disks and data disks are both queried.
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disks.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

type ListDisksResponseBody struct {
	// Details about the disks.
	Disks []*ListDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	// *   ESSD: an enhanced SSD (ESSD) at performance level 0 (PL0).
	// *   SSD: a standard SSD.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The device name of the disk on the simple application server.
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the disk.
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The ID of the disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The type of the disk. Valid values:
	//
	// *   System: system disk.
	// *   Data: data disk.
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the simple application server to which the disk is attached.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Name of the simple application server.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the disks.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Description about the disk.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The size of the disk. Unit: GB.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the disk. Valid values:
	//
	// *   ReIniting: The disk is being initialized.
	// *   Creating: The disk is being created.
	// *   In_Use: The disk is being used.
	// *   Available: The disk can be attached.
	// *   Attaching: The disk is being attached.
	// *   Detaching: The disk is being detached.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListDisksResponseBodyDisks) SetSize(v int32) *ListDisksResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *ListDisksResponseBodyDisks) SetStatus(v string) *ListDisksResponseBodyDisks {
	s.Status = &v
	return s
}

type ListDisksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListFirewallRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFirewallRulesRequest) GoString() string {
	return s.String()
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

type ListFirewallRulesResponseBody struct {
	// Details about the firewall rules.
	FirewallRules []*ListFirewallRulesResponseBodyFirewallRules `json:"FirewallRules,omitempty" xml:"FirewallRules,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	// The firewall policy.
	//
	// *   accept: Access is allowed.
	// *   drop: Access is refused.
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The remarks of the firewall rule.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// *   TCP: the TCP protocol.
	// *   UDP: the UDP protocol
	// *   TCP+UDP: the TCP and UDP protocols
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The IP address or CIDR block that is allowed by the firewall rule.
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
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

type ListFirewallRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFirewallRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of the images. Valid values:
	//
	// *   system: OS images
	// *   app: application images
	// *   custom: custom images
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The region ID of the images. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	// *   Linux
	// *   Windows
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The request ID.
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
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the image.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The type of the image. Valid values:
	//
	// *   system
	// *   app
	// *   custom
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The operating system type of the image. Valid values:
	//
	// *   Linux
	// *   Windows
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The unit of the plan price. Valid values:
	//
	// *   CNY
	// *   USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The disk size of the simple application server. Unit: GB.
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The category of the disk. Valid values:
	//
	// *   SSD: standard SSD
	// *   ESSD: enhanced SSD
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The monthly data transfer quota. Unit: GB.
	Flow *int32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size. Unit: GB.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The price of the plan.
	OriginPrice *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The operating system types supported by the plan.
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancePlansModificationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// > If you specify both `InstanceIds` and `PublicIpAddresses`, make sure that the specified IDs and the specified public IP addresses belong to the same simple application servers. Otherwise, an empty result is returned.
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The public IP addresses of the simple application servers. The value can be a JSON array that consists of up to 100 IP addresses. Separate multiple IP addresses with commas (,).
	//
	// > If you specify both `InstanceIds` and `PublicIpAddresses`, make sure that the specified IDs and the specified public IP addresses belong to the same simple application servers. Otherwise, an empty result is returned.
	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty"`
	// The region ID of the simple application servers.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 实例状态，可能值：
	//
	// - Pending：准备中
	// - Starting：启动中
	// - Running：运行中
	// - Stopping：停止中
	// - Stopped：停止
	// - Resetting：重置中
	// - Upgrading：升级中
	// - Disabled：不可用
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

type ListInstancesResponseBody struct {
	// Details about the simple application servers.
	Instances []*ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	// *   Normal: The server is normal.
	// *   Expired: The server expires.
	// *   Overdue: The payment of the server is overdue.
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method of the simple application server.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Indicates whether the plan is a bundle plan.
	Combination *bool `json:"Combination,omitempty" xml:"Combination,omitempty"`
	// The ID of the bundle plan.
	CombinationInstanceId *string `json:"CombinationInstanceId,omitempty" xml:"CombinationInstanceId,omitempty"`
	// The time when the simple application server was created. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The DDoS protection status of the server. Valid values:
	//
	// *   Normal: The DDoS protection status of the server is normal.
	// *   BlackHole: The server is in blackhole filtering.
	// *   Defense: The server is being scrubbed.
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The reason why the server is disabled. Valid values:
	//
	// *   FINANCIAL: The server is locked due to overdue payments.
	// *   SECURITY: The server is locked due to security reasons.
	// *   EXPIRED: The server has expired.
	DisableReason *string `json:"DisableReason,omitempty" xml:"DisableReason,omitempty"`
	// The time when the server expires. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The description of the image.
	Image *ListInstancesResponseBodyInstancesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The ID of an image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address of the simple application server.
	InnerIpAddress *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The plan ID.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The public IP address of the server.
	PublicIpAddress *string `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	// The region ID of the servers.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The specifications of the resource.
	ResourceSpec *ListInstancesResponseBodyInstancesResourceSpec `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty" type:"Struct"`
	// The status of the simple application server. Valid values:
	//
	// *   Pending
	// *   Starting
	// *   Running
	// *   Stopping
	// *   Stopped
	// *   Resetting
	// *   Upgrading
	// *   Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The universally unique identifier (UUID) of the server.
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

func (s *ListInstancesResponseBodyInstances) SetResourceSpec(v *ListInstancesResponseBodyInstancesResourceSpec) *ListInstancesResponseBodyInstances {
	s.ResourceSpec = v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetStatus(v string) *ListInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstances) SetUuid(v string) *ListInstancesResponseBodyInstances {
	s.Uuid = &v
	return s
}

type ListInstancesResponseBodyInstancesImage struct {
	// The image provider.
	ImageContact *string `json:"ImageContact,omitempty" xml:"ImageContact,omitempty"`
	// The URL of the image icon.
	ImageIconUrl *string `json:"ImageIconUrl,omitempty" xml:"ImageIconUrl,omitempty"`
	// The image name.
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image type. Valid values:
	//
	// *   system
	// *   app
	// *   custom
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The image tag.
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The OS.
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
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The category of the disk. Valid values:
	//
	// *   ESSD: an enhanced SSD (ESSD) at performance level 0 (PL0).
	// *   SSD: a standard SSD.
	// *   CLOUD_EFFICIENCY: an ultra disk.
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk size.
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The amount of the traffic.
	//
	// *   A value of 0 indicates that the server is a bandwidth-based server.
	// *   A value of none-zero indicates that the server is a data transfer plan-based server.
	Flow *float64 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size.
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

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AcsProduct *string `json:"AcsProduct,omitempty" xml:"AcsProduct,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesTrafficPackagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesTrafficPackagesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesTrafficPackagesRequest) SetAcsProduct(v string) *ListInstancesTrafficPackagesRequest {
	s.AcsProduct = &v
	return s
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The data transfers that exceeds the quota of the data transfer plan in the current month. Unit: Byte.
	TrafficOverflow *int64 `json:"TrafficOverflow,omitempty" xml:"TrafficOverflow,omitempty"`
	// The unused quota of the data transfer plan in the current month. Unit: Byte.
	TrafficPackageRemaining *int64 `json:"TrafficPackageRemaining,omitempty" xml:"TrafficPackageRemaining,omitempty"`
	// The quota of the data transfer plan in the current month. Unit: Byte.
	//
	// >  TrafficPackageTotal = TrafficUsed + TrafficPackageRemaining
	TrafficPackageTotal *int64 `json:"TrafficPackageTotal,omitempty" xml:"TrafficPackageTotal,omitempty"`
	// The used quota of the data transfer plan in the current month. Unit: Byte.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesTrafficPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListPlansRequest struct {
	// The region ID of the plans. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The number of vCPUs.
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The unit of the plan price. Valid values:
	//
	// *   CNY
	// *   USD
	//
	// >  CNY is for the China site (aliyun.com). USD is for the international site (alibabacloud.com).
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The size of the disk. Unit: GB.
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The category of the disk. Valid values:
	//
	// *   SSD: standard SSDs
	// *   ESSD: enhanced SSDs
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The monthly data transfer quota. Unit: GB.
	Flow *int32 `json:"Flow,omitempty" xml:"Flow,omitempty"`
	// The memory size. Unit: GB.
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The monthly price of the plan.
	OriginPrice *float64 `json:"OriginPrice,omitempty" xml:"OriginPrice,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The operating system types supported by the plan.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListRegionsResponseBody struct {
	// The region ID.
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AcsProduct *string `json:"AcsProduct,omitempty" xml:"AcsProduct,omitempty"`
	// The disk ID.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server that corresponds to the snapshots.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot IDs. The value can be a JSON array that consists of up to 100 snapshot IDs. Separate multiple snapshot IDs with commas (,).
	SnapshotIds *string `json:"SnapshotIds,omitempty" xml:"SnapshotIds,omitempty"`
	// The type of the source disk. Valid values:
	//
	// *   system: system disk.
	// *   data: data disk.
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
}

func (s ListSnapshotsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSnapshotsRequest) GoString() string {
	return s.String()
}

func (s *ListSnapshotsRequest) SetAcsProduct(v string) *ListSnapshotsRequest {
	s.AcsProduct = &v
	return s
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

func (s *ListSnapshotsRequest) SetSnapshotIds(v string) *ListSnapshotsRequest {
	s.SnapshotIds = &v
	return s
}

func (s *ListSnapshotsRequest) SetSourceDiskType(v string) *ListSnapshotsRequest {
	s.SourceDiskType = &v
	return s
}

type ListSnapshotsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the snapshots.
	Snapshots []*ListSnapshotsResponseBodySnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The total number of entries returned.
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
	// The time when the snapshot was created. The time follows the [ISO 8601](~~25696~~) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the simple application server.
	//
	// Note: This parameter has a value for system disk snapshots. This parameter is left empty for data disk snapshots.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The progress of snapshot creation.
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The region ID of the snapshots.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the snapshot.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The time when the last disk rollback was performed.
	RollbackTime *string `json:"RollbackTime,omitempty" xml:"RollbackTime,omitempty"`
	// The snapshot ID.
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot.
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// The ID of the source disk based on which the snapshot is created. This parameter has a value even if the source disk is released.
	SourceDiskId *string `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	// The type of the source disk. Valid values:
	//
	// *   system: system disk.
	// *   data: data disk.
	SourceDiskType *string `json:"SourceDiskType,omitempty" xml:"SourceDiskType,omitempty"`
	// The status of the snapshot. Valid values:
	//
	// *   Progressing: The snapshot is being created.
	// *   Accomplished: The snapshot is created.
	// *   Failed: The snapshot failed to be created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

type ListSnapshotsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSnapshotsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type LoginInstanceRequest struct {
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password that corresponds to the username.
	//
	// *   For a Linux server, you do not need to enter a password.
	// *   For a Windows server, enter the password that you set. If you have not set a password for the simple application server, set a password. For more information, see [Reset the password](~~60055~~).
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The username of the simple application server.
	//
	// *   For a Linux server, you do not need to enter a username.
	// *   For a Windows server, the default username is `administrator`.
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
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The request ID.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *LoginInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the Simple Database Service instance.
	DatabaseInstanceDescription *string `json:"DatabaseInstanceDescription,omitempty" xml:"DatabaseInstanceDescription,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseInstanceDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// Specifies whether to forcibly restart the instance after parameters are modified. Valid values:
	//
	// *   true: forcibly restarts the instance. If a new parameter value takes effect only after the instance restarts, you must set this parameter to true. Otherwise, the new parameter value cannot take effect.
	// *   false: does not forcibly restart the instance.
	//
	// Default value: false.
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// The JSON strings that consist of instance parameters and the values of the instance parameters. The parameter values are of the string type. Format: {"Parameter name 1":"Parameter value 1","Parameter name 2":"Parameter value 2"...}.
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseInstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range. Valid values: 165535. Specify a port range in the format of \<start port number>/\<end port number>. Example: `1024/1055`, which indicates that the port range of 10241055.
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the firewall rule.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the firewall rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The transport layer protocol. Valid values:
	//
	// *   TCP: the TCP protocol
	// *   UDP: the UDP protocol
	// *   TCP+UDP: the TCP and UDP protocols
	RuleProtocol *string `json:"RuleProtocol,omitempty" xml:"RuleProtocol,omitempty"`
	// The IP address or CIDR block that is allowed in the firewall rule.
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFirewallRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyImageShareStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image ID.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Valid values:
	//
	// *   Share
	// *   UnShare
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The region ID of the custom image. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyImageShareStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The existing VNC password.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceVncPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebootInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly restart the servers. Valid values:
	//
	// *   true: forcibly restarts the servers. This operation is equivalent to the typical power-off operation. Cache data that is not written to storage devices on the server will be lost.
	// *   false: normally restarts the instance.
	//
	// Default value: false
	ForceReboot *bool `json:"ForceReboot,omitempty" xml:"ForceReboot,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebootInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleasePublicConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RenewInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The renewal period. Unit: month. Valid values: 1, 3, 6, 12, 24, and 36.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the server.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetDatabaseAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the disk to be rolled back.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server for which the snapshot is created.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The snapshot ID.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetDiskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the image that is used to replace the image of the simple application server. If you do not specify this parameter, the current image of the simple application server is replaced by default.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetSystemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// *   If you set `EnableParameter` to true, the custom parameter feature is enabled in the command content and you can configure custom parameters based on the following rules:
	// *   Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	// *   The number of custom parameters cannot be greater than 20.
	// *   A custom parameter name can contain only letters, digits, underscores (\_), and hyphens (-). The name is case-insensitive.
	// *   Each custom parameter name cannot exceed 64 bytes in length.
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Specifies whether to enable the custom parameter feature.
	//
	// Default value: false.
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the command.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom parameters in the key-value pair format that are to be passed in when the command includes custom parameters. For example, if the command content is `echo {{name}}`, you can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced with the paired Jack value to generate a new command. As a result, the `echo Jack` command is executed.
	//
	// Number of custom parameters ranges from 0 to 20. Take note of the following items:
	//
	// *   The key cannot be an empty string. It can be up to 64 characters in length.
	// *   The value can be an empty string.
	// *   After custom parameters and original command content are encoded in Base64, the command cannot exceed 16 KB in size.
	// *   The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is empty by default, which indicates to disable the custom parameter feature.
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of the command on the server.
	//
	// If a command execution task times out, Command Assistant forcibly terminates the task process. Valid values: 10 to 86400. Unit: seconds. The period of 86400 seconds is equal to 24 hours.
	//
	// Default value: 60.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// *   RunBatScript: batch commands (applicable to Windows servers).
	// *   RunPowerShellScript: PowerShell commands (applicable to Windows servers).
	// *   RunShellScript: shell commands (applicable to Linux servers).
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the password to be used to run the command on a Windows server.
	//
	// If you want to use a username other than the default "system" username to run the command on a Windows server, you must specify both the WindowsPasswordName and WorkingUser parameters. To mitigate the risk of password leaks, the password is stored in plaintext in Operation Orchestration Service (OOS) Parameter Store, and only the name of the password is passed in by using WindowsPasswordName.
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The execution path of the command. You can specify a value for the parameter. Default execution paths vary based on the operating systems of the servers.
	//
	// *   For Linux servers, the default execution path is the /home directory of the root user.
	// *   For Windows servers, the default execution path is C:\Windows\system32.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// A user of the server who runs the command. We recommend that you run the command as a regular user to reduce security risks. Default values:
	//
	// *   For Linux servers, the default value is root.
	// *   For Windows servers, the default value is system.
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
	// *   If you set `EnableParameter` to true, the custom parameter feature is enabled in the command content and you can configure custom parameters based on the following rules:
	// *   Define custom parameters in the {{}} format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	// *   The number of custom parameters cannot be greater than 20.
	// *   A custom parameter name can contain only letters, digits, underscores (\_), and hyphens (-). The name is case-insensitive.
	// *   Each custom parameter name cannot exceed 64 bytes in length.
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Specifies whether to enable the custom parameter feature.
	//
	// Default value: false.
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the command.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom parameters in the key-value pair format that are to be passed in when the command includes custom parameters. For example, if the command content is `echo {{name}}`, you can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced with the paired Jack value to generate a new command. As a result, the `echo Jack` command is executed.
	//
	// Number of custom parameters ranges from 0 to 20. Take note of the following items:
	//
	// *   The key cannot be an empty string. It can be up to 64 characters in length.
	// *   The value can be an empty string.
	// *   After custom parameters and original command content are encoded in Base64, the command cannot exceed 16 KB in size.
	// *   The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is empty by default, which indicates to disable the custom parameter feature.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The timeout period of the command on the server.
	//
	// If a command execution task times out, Command Assistant forcibly terminates the task process. Valid values: 10 to 86400. Unit: seconds. The period of 86400 seconds is equal to 24 hours.
	//
	// Default value: 60.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The language type of the command. Valid values:
	//
	// *   RunBatScript: batch commands (applicable to Windows servers).
	// *   RunPowerShellScript: PowerShell commands (applicable to Windows servers).
	// *   RunShellScript: shell commands (applicable to Linux servers).
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the password to be used to run the command on a Windows server.
	//
	// If you want to use a username other than the default "system" username to run the command on a Windows server, you must specify both the WindowsPasswordName and WorkingUser parameters. To mitigate the risk of password leaks, the password is stored in plaintext in Operation Orchestration Service (OOS) Parameter Store, and only the name of the password is passed in by using WindowsPasswordName.
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The execution path of the command. You can specify a value for the parameter. Default execution paths vary based on the operating systems of the servers.
	//
	// *   For Linux servers, the default execution path is the /home directory of the root user.
	// *   For Windows servers, the default execution path is C:\Windows\system32.
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// A user of the server who runs the command. We recommend that you run the command as a regular user to reduce security risks. Default values:
	//
	// *   For Linux servers, the default value is root.
	// *   For Windows servers, the default value is system.
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
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The request ID.
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the server.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security token included in the WebSocket request header. The system uses this token to authenticate the request.
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The session ID.
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The URL of the WebSocket session that is used to connect to the server. The URL contains the session ID (`SessionId`) and the authentication token (`SecurityToken`).
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartTerminalSessionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Simple Database Service instance.
	DatabaseInstanceId *string `json:"DatabaseInstanceId,omitempty" xml:"DatabaseInstanceId,omitempty"`
	// The region ID of the Simple Database Service instance. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopDatabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the server.
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly stop the servers.
	//
	// *   **true**: forcibly stops the servers.
	// *   **false**: normally stops the servers. This is the default value.
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateCommandAttributeRequest struct {
	CommandId   *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Timeout     *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkingDir  *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCommandAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The disk ID. You can call the ListDisks operation to query the ID of data disk.
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the data disk.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDiskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the simple application server. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`. The name can only contain letters, digits, colons (:), underscores (\_), periods (.), and hyphens (-).
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The new password of the simple application server. The password must be 8 to 30 characters in length. It must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Only the following special characters are supported:
	//
	// `()~!@#$%^&*-+=|{}[]:;<>,.?/`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID of the simple application server.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken** can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remarks of the snapshot of the simple application server.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The snapshot ID. You can call the ListSnapshots operation to query the snapshot ID.
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSnapshotAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new plan. You can call the [ListPlans](~~189314~~) operation to query the plans provided by Simple Application Server.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The region ID of the simple application server.
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpgradeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](~~25693~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the simple application server.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the key pair.
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The public key.
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](~~189315~~) operation to query the most recent region list.
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadInstanceKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

/**
 * By default, no public endpoints are assigned to Simple Database Service instances. If you want to access the databases of a Simple Database Service instance over the Internet by using Simple Container Service or Data Management (DMS), you must apply for a public endpoint for the Simple Database Service instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request AllocatePublicConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return AllocatePublicConnectionResponse
 */
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

/**
 * By default, no public endpoints are assigned to Simple Database Service instances. If you want to access the databases of a Simple Database Service instance over the Internet by using Simple Container Service or Data Management (DMS), you must apply for a public endpoint for the Simple Database Service instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request AllocatePublicConnectionRequest
 * @return AllocatePublicConnectionResponse
 */
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

/**
 * A custom image is created based on a snapshot of a simple application server. You can use a custom image to create multiple simple application servers that have the same configurations. You can also share custom images to ECS and use the shared images to create ECS instances or replace the OSs of existing ECS instances. For more information about custom images, see [Overview of custom images](~~199375~~).
 * You must create a system disk snapshot of a simple application server before you create a custom image based on the snapshot. For more information, see [CreateSnapshot](~~190452~~).
 * > If you need the data on the data disk of a simple application server when you create a custom image, create a snapshot for the data disk first.
 * Before you create a custom image, take note of the following items:
 * *   The custom image and the corresponding simple application server must reside in the same region.
 * *   The maximum number of custom images that can be maintained in an Alibaba Cloud account is triple the number of simple application servers in the account. The value cannot be greater than 15.
 * *   You can directly create a custom image only based on the system disk snapshot of a simple application server. If you want a custom image to contain the data on the data disk of the simple application server, you must select a data disk snapshot when you create the custom image.
 * *   If a simple application server is released due to expiration or refunds, the custom images that are created based on a snapshot of the server are also released.
 * *   If you reset a simple application server by changing the application system or OS of the server or replacing the image of the server, the disk data on the server is cleared. Back up the disk data as needed.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateCustomImageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateCustomImageResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.SystemSnapshotId)) {
		query["SystemSnapshotId"] = request.SystemSnapshotId
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

/**
 * A custom image is created based on a snapshot of a simple application server. You can use a custom image to create multiple simple application servers that have the same configurations. You can also share custom images to ECS and use the shared images to create ECS instances or replace the OSs of existing ECS instances. For more information about custom images, see [Overview of custom images](~~199375~~).
 * You must create a system disk snapshot of a simple application server before you create a custom image based on the snapshot. For more information, see [CreateSnapshot](~~190452~~).
 * > If you need the data on the data disk of a simple application server when you create a custom image, create a snapshot for the data disk first.
 * Before you create a custom image, take note of the following items:
 * *   The custom image and the corresponding simple application server must reside in the same region.
 * *   The maximum number of custom images that can be maintained in an Alibaba Cloud account is triple the number of simple application servers in the account. The value cannot be greater than 15.
 * *   You can directly create a custom image only based on the system disk snapshot of a simple application server. If you want a custom image to contain the data on the data disk of the simple application server, you must select a data disk snapshot when you create the custom image.
 * *   If a simple application server is released due to expiration or refunds, the custom images that are created based on a snapshot of the server are also released.
 * *   If you reset a simple application server by changing the application system or OS of the server or replacing the image of the server, the disk data on the server is cleared. Back up the disk data as needed.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateCustomImageRequest
 * @return CreateCustomImageResponse
 */
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

/**
 * Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
 * ### QPS limits
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateFirewallRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFirewallRuleResponse
 */
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

/**
 * Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
 * ### QPS limits
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateFirewallRuleRequest
 * @return CreateFirewallRuleResponse
 */
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

/**
 * Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
 *
 * @param tmpReq CreateFirewallRulesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFirewallRulesResponse
 */
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

/**
 * Firewalls serve to control network access to simple application servers and isolate security domains in the cloud. By default, SSH port 22, HTTP port 80, and HTTPS port 443 are enabled for simple application servers. Other ports are disabled. You can add firewall rules to enable more ports.
 *
 * @param request CreateFirewallRulesRequest
 * @return CreateFirewallRulesResponse
 */
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

/**
 * *   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](~~58623~~).
 * *   A maximum of 20 simple application servers can be maintained in an Alibaba Cloud account.
 * *   When you call this operation to create simple application servers, make sure that the balance in your account is sufficient to pay for the servers. If the balance in your account is insufficient, the servers cannot be created.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateInstancesResponse
 */
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

/**
 * *   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](~~58623~~).
 * *   A maximum of 20 simple application servers can be maintained in an Alibaba Cloud account.
 * *   When you call this operation to create simple application servers, make sure that the balance in your account is sufficient to pay for the servers. If the balance in your account is insufficient, the servers cannot be created.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateInstancesRequest
 * @return CreateInstancesResponse
 */
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

/**
 * A snapshot is a point-in-time backup of a disk. Snapshots can be used to back up data, recover data after accidental operations on instances, recover data after network attacks, and create custom images.
 * > You are not charged for creating snapshots for disks of simple application servers.
 * ### Precautions
 * *   You can create up to three snapshots for disks of each simple application server.
 * *   The maximum number of snapshots that can be retained in an Alibaba Cloud account is triple the number of simple application servers that you maintain. The value cannot be greater than 15.
 * *   If a simple application server is automatically released due to expiration, the snapshots created for the server are deleted.
 * *   If you reset the simple application server after you create a snapshot for a server, the snapshot is retained but cannot be used to roll back the disks of the server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateSnapshotRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateSnapshotResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.SnapshotName)) {
		query["SnapshotName"] = request.SnapshotName
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

/**
 * A snapshot is a point-in-time backup of a disk. Snapshots can be used to back up data, recover data after accidental operations on instances, recover data after network attacks, and create custom images.
 * > You are not charged for creating snapshots for disks of simple application servers.
 * ### Precautions
 * *   You can create up to three snapshots for disks of each simple application server.
 * *   The maximum number of snapshots that can be retained in an Alibaba Cloud account is triple the number of simple application servers that you maintain. The value cannot be greater than 15.
 * *   If a simple application server is automatically released due to expiration, the snapshots created for the server are deleted.
 * *   If you reset the simple application server after you create a snapshot for a server, the snapshot is retained but cannot be used to roll back the disks of the server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request CreateSnapshotRequest
 * @return CreateSnapshotResponse
 */
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

/**
 * You can delete a custom image that you no longer need. After the custom image is deleted, you cannot use the custom image to reset the simple application servers that were created based on the custom image.
 * > If a custom image is shared to Elastic Compute Service (ECS), you must unshare the image before you can delete it. After you unshare the custom image, you cannot query the custom image by using the ECS console or by calling ECS API operations. If you need to use the custom image in ECS, we recommend that you copy the image before you delete it. For more information, see [Copy a shared image of a simple application server in the ECS console](~~199378~~).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DeleteCustomImageRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteCustomImageResponse
 */
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

/**
 * You can delete a custom image that you no longer need. After the custom image is deleted, you cannot use the custom image to reset the simple application servers that were created based on the custom image.
 * > If a custom image is shared to Elastic Compute Service (ECS), you must unshare the image before you can delete it. After you unshare the custom image, you cannot query the custom image by using the ECS console or by calling ECS API operations. If you need to use the custom image in ECS, we recommend that you copy the image before you delete it. For more information, see [Copy a shared image of a simple application server in the ECS console](~~199378~~).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DeleteCustomImageRequest
 * @return DeleteCustomImageResponse
 */
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

/**
 * After a firewall rule is deleted, your business deployed on the simple application server may become inaccessible. Before you delete a firewall rule, make sure that the firewall rule is no longer needed by the simple application server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DeleteFirewallRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteFirewallRuleResponse
 */
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

/**
 * After a firewall rule is deleted, your business deployed on the simple application server may become inaccessible. Before you delete a firewall rule, make sure that the firewall rule is no longer needed by the simple application server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DeleteFirewallRuleRequest
 * @return DeleteFirewallRuleResponse
 */
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

/**
 * You can delete a snapshot if you no longer need it.
 * > If a custom image was created based on the snapshot, delete the custom image before you delete the snapshot.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DeleteSnapshotRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteSnapshotResponse
 */
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

/**
 * You can delete a snapshot if you no longer need it.
 * > If a custom image was created based on the snapshot, delete the custom image before you delete the snapshot.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DeleteSnapshotRequest
 * @return DeleteSnapshotResponse
 */
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

/**
 * By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall the client. Otherwise, you cannot run commands on the servers.
 *
 * @param tmpReq DescribeCloudAssistantStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeCloudAssistantStatusResponse
 */
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

/**
 * By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall the client. Otherwise, you cannot run commands on the servers.
 *
 * @param request DescribeCloudAssistantStatusRequest
 * @return DescribeCloudAssistantStatusResponse
 */
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

/**
 * You can call this operation to query the error logs of databases in a Simple Database Service instance and locate faults based on the error logs.
 * \\### QPS limit You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseErrorLogsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDatabaseErrorLogsResponse
 */
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

/**
 * You can call this operation to query the error logs of databases in a Simple Database Service instance and locate faults based on the error logs.
 * \\### QPS limit You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseErrorLogsRequest
 * @return DescribeDatabaseErrorLogsResponse
 */
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

/**
 * After you create a Simple Database Service instance, you can query the details about the vCPU, memory, disk size, storage IOPS (input/output operations per second), and total current connection number of the instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseInstanceMetricDataRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDatabaseInstanceMetricDataResponse
 */
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

/**
 * After you create a Simple Database Service instance, you can query the details about the vCPU, memory, disk size, storage IOPS (input/output operations per second), and total current connection number of the instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseInstanceMetricDataRequest
 * @return DescribeDatabaseInstanceMetricDataResponse
 */
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

/**
 * You can call this operation to query the information about parameters of a Simple Database Service instance.
 *
 * @param request DescribeDatabaseInstanceParametersRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDatabaseInstanceParametersResponse
 */
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

/**
 * You can call this operation to query the information about parameters of a Simple Database Service instance.
 *
 * @param request DescribeDatabaseInstanceParametersRequest
 * @return DescribeDatabaseInstanceParametersResponse
 */
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

/**
 * You can call this operation to query the details of Simple Database Service instances in a region, including the IDs, names, plans, database versions, public endpoint, internal endpoint, creation time, and expiration time of the instances.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDatabaseInstancesResponse
 */
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

/**
 * You can call this operation to query the details of Simple Database Service instances in a region, including the IDs, names, plans, database versions, public endpoint, internal endpoint, creation time, and expiration time of the instances.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseInstancesRequest
 * @return DescribeDatabaseInstancesResponse
 */
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

/**
 * You can query the slow query log details of a Simple Database Service instance and locate faults based on the log details.
 * > Slow query log details are retained for 7 days.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseSlowLogRecordsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDatabaseSlowLogRecordsResponse
 */
func (client *Client) DescribeDatabaseSlowLogRecordsWithOptions(request *DescribeDatabaseSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabaseSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsProduct)) {
		query["AcsProduct"] = request.AcsProduct
	}

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

/**
 * You can query the slow query log details of a Simple Database Service instance and locate faults based on the log details.
 * > Slow query log details are retained for 7 days.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request DescribeDatabaseSlowLogRecordsRequest
 * @return DescribeDatabaseSlowLogRecordsResponse
 */
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

/**
 * *   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the execution result of a command.
 * *   You can query the execution results that were generated within the last two weeks. A maximum of 100,000 entries of execution results can be retained.
 *
 * @param request DescribeInvocationResultRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeInvocationResultResponse
 */
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

/**
 * *   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the execution result of a command.
 * *   You can query the execution results that were generated within the last two weeks. A maximum of 100,000 entries of execution results can be retained.
 *
 * @param request DescribeInvocationResultRequest
 * @return DescribeInvocationResultResponse
 */
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

/**
 * *   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the actual execution results.
 * *   You can query the execution results that were generated within the last two weeks. Up to 100,000 entries of execution results can be retained.
 *
 * @param request DescribeInvocationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeInvocationsResponse
 */
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

/**
 * *   After you execute a command, the command may not succeed or return the expected results. You can call this operation to query the actual execution results.
 * *   You can query the execution results that were generated within the last two weeks. Up to 100,000 entries of execution results can be retained.
 *
 * @param request DescribeInvocationsRequest
 * @return DescribeInvocationsResponse
 */
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

/**
 * To run commands on your simple application servers, you must install the Cloud Assistant client on your servers. You can call the [DescribeCloudAssistantStatus](~~439512~~) operation to check whether the Cloud Assistant client is installed on your simple application servers. If you have not installed the Cloud Assistant client, you can call the InstallCloudAssistant operation to install the client. Then, you can call the [RebootInstance](~~190443~~) operation to restart the servers to allow the client to take effect.
 *
 * @param tmpReq InstallCloudAssistantRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return InstallCloudAssistantResponse
 */
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

/**
 * To run commands on your simple application servers, you must install the Cloud Assistant client on your servers. You can call the [DescribeCloudAssistantStatus](~~439512~~) operation to check whether the Cloud Assistant client is installed on your simple application servers. If you have not installed the Cloud Assistant client, you can call the InstallCloudAssistant operation to install the client. Then, you can call the [RebootInstance](~~190443~~) operation to restart the servers to allow the client to take effect.
 *
 * @param request InstallCloudAssistantRequest
 * @return InstallCloudAssistantResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SystemSnapshotId)) {
		query["SystemSnapshotId"] = request.SystemSnapshotId
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

/**
 * You can specify multiple request parameters such as `InstanceId` and `DiskIds`. Specified parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](~~347607~~).
 *
 * @param request ListDisksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListDisksResponse
 */
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

/**
 * You can specify multiple request parameters such as `InstanceId` and `DiskIds`. Specified parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limit](~~347607~~).
 *
 * @param request ListDisksRequest
 * @return ListDisksResponse
 */
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

/**
 * You can call the ListFirewallRules operation to query the firewall rule details of a simple application server, including the port range, firewall rule ID, and transport layer protocol.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListFirewallRulesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListFirewallRulesResponse
 */
func (client *Client) ListFirewallRulesWithOptions(request *ListFirewallRulesRequest, runtime *util.RuntimeOptions) (_result *ListFirewallRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
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

/**
 * You can call the ListFirewallRules operation to query the firewall rule details of a simple application server, including the port range, firewall rule ID, and transport layer protocol.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListFirewallRulesRequest
 * @return ListFirewallRulesResponse
 */
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

/**
 * You can query information about images in a region, including the IDs, names, and types of the images.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListImagesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListImagesResponse
 */
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

/**
 * You can query information about images in a region, including the IDs, names, and types of the images.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListImagesRequest
 * @return ListImagesResponse
 */
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

/**
 * If the plan of your simple application server does not meet your business requirements, you can call the ListInstancePlansModification operation to obtain a list of plans to which you can upgrade your simple application server. Then, you can call the [UpgradeInstance](~~190445~~) operation to upgrade the server.
 * > We recommend that you create snapshots for the disks of your simple application server to back up data before you upgrade the server. For more information, see [CreateSnapshot](~~190452~~).
 * For the precautions about plan upgrade, see [Upgrade a simple application server](~~61433~~).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListInstancePlansModificationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListInstancePlansModificationResponse
 */
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

/**
 * If the plan of your simple application server does not meet your business requirements, you can call the ListInstancePlansModification operation to obtain a list of plans to which you can upgrade your simple application server. Then, you can call the [UpgradeInstance](~~190445~~) operation to upgrade the server.
 * > We recommend that you create snapshots for the disks of your simple application server to back up data before you upgrade the server. For more information, see [CreateSnapshot](~~190452~~).
 * For the precautions about plan upgrade, see [Upgrade a simple application server](~~61433~~).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListInstancePlansModificationRequest
 * @return ListInstancePlansModificationResponse
 */
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

/**
 * You can call this operation to query the details of simple application servers in a specified region, including the names, public IP addresses, internal IP addresses, creation time, and expiration time of the servers.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListInstancesResponse
 */
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
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

/**
 * You can call this operation to query the details of simple application servers in a specified region, including the names, public IP addresses, internal IP addresses, creation time, and expiration time of the servers.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListInstancesRequest
 * @return ListInstancesResponse
 */
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

/**
 * You can query the details of data transfer plans of simple application servers, including the data transfer quota, used amount and unused amount of the data transfer quota, and excess data transfers beyond the quota in the current month.
 * Simple Application Server provides data transfer quotas in plans. Plan prices include prices of data transfer quotas. You are charged for data transfers that exceed the quotas. Take note of the following items:
 * *   Only outbound data transfers of simple application servers over the Internet are calculated. Outbound data transfers include the data transfer quota and the excess data transfers beyond the quota. Inbound data transfers of simple application servers over the Internet are not calculated.
 * *   Outbound data transfers from simple application servers to other Alibaba Cloud services over the Internet first consume data transfer quotas. If the quotas are exhausted, you are charged for excess data transfers.
 * *   You are not charged for data transfers between simple application servers within the same virtual private cloud (VPC).
 * For more information, see [Quotas and billing of data transfers](~~86281~~).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListInstancesTrafficPackagesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListInstancesTrafficPackagesResponse
 */
func (client *Client) ListInstancesTrafficPackagesWithOptions(request *ListInstancesTrafficPackagesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesTrafficPackagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsProduct)) {
		query["AcsProduct"] = request.AcsProduct
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

/**
 * You can query the details of data transfer plans of simple application servers, including the data transfer quota, used amount and unused amount of the data transfer quota, and excess data transfers beyond the quota in the current month.
 * Simple Application Server provides data transfer quotas in plans. Plan prices include prices of data transfer quotas. You are charged for data transfers that exceed the quotas. Take note of the following items:
 * *   Only outbound data transfers of simple application servers over the Internet are calculated. Outbound data transfers include the data transfer quota and the excess data transfers beyond the quota. Inbound data transfers of simple application servers over the Internet are not calculated.
 * *   Outbound data transfers from simple application servers to other Alibaba Cloud services over the Internet first consume data transfer quotas. If the quotas are exhausted, you are charged for excess data transfers.
 * *   You are not charged for data transfers between simple application servers within the same virtual private cloud (VPC).
 * For more information, see [Quotas and billing of data transfers](~~86281~~).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListInstancesTrafficPackagesRequest
 * @return ListInstancesTrafficPackagesResponse
 */
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

/**
 * You can query the details of all plans provided by Simple Application Server in a region, including the IDs, prices, disk sizes, and disk categories of the plans.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListPlansRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListPlansResponse
 */
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

/**
 * You can query the details of all plans provided by Simple Application Server in a region, including the IDs, prices, disk sizes, and disk categories of the plans.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListPlansRequest
 * @return ListPlansResponse
 */
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

/**
 * The query results include all the Alibaba Cloud regions where Simple Application Server is supported on the international site (alibabacloud.com) and the China site (aliyun.com).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListRegionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRegionsResponse
 */
func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
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

/**
 * The query results include all the Alibaba Cloud regions where Simple Application Server is supported on the international site (alibabacloud.com) and the China site (aliyun.com).
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @return ListRegionsResponse
 */
func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can specify multiple request parameters such as `InstanceId`, `DiskId`, and `SnapshotIds` to query snapshots. Specified parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListSnapshotsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListSnapshotsResponse
 */
func (client *Client) ListSnapshotsWithOptions(request *ListSnapshotsRequest, runtime *util.RuntimeOptions) (_result *ListSnapshotsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcsProduct)) {
		query["AcsProduct"] = request.AcsProduct
	}

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

	if !tea.BoolValue(util.IsUnset(request.SnapshotIds)) {
		query["SnapshotIds"] = request.SnapshotIds
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDiskType)) {
		query["SourceDiskType"] = request.SourceDiskType
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

/**
 * You can specify multiple request parameters such as `InstanceId`, `DiskId`, and `SnapshotIds` to query snapshots. Specified parameters have logical AND relations. Only the specified parameters are included in the filter conditions.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ListSnapshotsRequest
 * @return ListSnapshotsResponse
 */
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

/**
 * ##
 * After you create a simple application server, you can log on to the simple application server to build environments and applications on the server.
 *
 * @param request LoginInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return LoginInstanceResponse
 */
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

/**
 * ##
 * After you create a simple application server, you can log on to the simple application server to build environments and applications on the server.
 *
 * @param request LoginInstanceRequest
 * @return LoginInstanceResponse
 */
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

/**
 * You can call this operation to modify the description of a Simple Database Service instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ModifyDatabaseInstanceDescriptionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDatabaseInstanceDescriptionResponse
 */
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

/**
 * You can call this operation to modify the description of a Simple Database Service instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ModifyDatabaseInstanceDescriptionRequest
 * @return ModifyDatabaseInstanceDescriptionResponse
 */
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

/**
 * After you create a Simple Database Service instance, you can view the parameters of the instance or modify the parameters of the instance based on your business requirements.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ModifyDatabaseInstanceParameterRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyDatabaseInstanceParameterResponse
 */
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

/**
 * After you create a Simple Database Service instance, you can view the parameters of the instance or modify the parameters of the instance based on your business requirements.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ModifyDatabaseInstanceParameterRequest
 * @return ModifyDatabaseInstanceParameterResponse
 */
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

/**
 * You can share a custom image with ECS. If the configurations of your simple application server cannot meet your business requirements, or you want to use ECS instances to deploy your business, you can share your custom image with ECS to transfer your business from Simple Application Server to ECS.
 * > The shared image in ECS resides in the same region as the custom image in Simple Application Server.
 * You can unshare a custom image based on your business requirements or when you want to delete the custom image. Take note of the following items:
 * *   After you unshare a custom image, you cannot query or use the custom image in the ECS console or by calling ECS API operations.
 * *   After you unshare a custom image, you cannot re-initialize the disks of the ECS instances that were created based on the shared image.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ModifyImageShareStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyImageShareStatusResponse
 */
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

/**
 * You can share a custom image with ECS. If the configurations of your simple application server cannot meet your business requirements, or you want to use ECS instances to deploy your business, you can share your custom image with ECS to transfer your business from Simple Application Server to ECS.
 * > The shared image in ECS resides in the same region as the custom image in Simple Application Server.
 * You can unshare a custom image based on your business requirements or when you want to delete the custom image. Take note of the following items:
 * *   After you unshare a custom image, you cannot query or use the custom image in the ECS console or by calling ECS API operations.
 * *   After you unshare a custom image, you cannot re-initialize the disks of the ECS instances that were created based on the shared image.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ModifyImageShareStatusRequest
 * @return ModifyImageShareStatusResponse
 */
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

/**
 * *   Only simple application servers that are in the Running state can be restarted.
 * *   After you restart a simple application server, it enters the Starting state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request RebootInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RebootInstanceResponse
 */
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

/**
 * *   Only simple application servers that are in the Running state can be restarted.
 * *   After you restart a simple application server, it enters the Starting state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request RebootInstanceRequest
 * @return RebootInstanceResponse
 */
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

/**
 * If you no longer need to use a public endpoint to access a Simple Database Service instance, you can release the public endpoint.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ReleasePublicConnectionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ReleasePublicConnectionResponse
 */
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

/**
 * If you no longer need to use a public endpoint to access a Simple Database Service instance, you can release the public endpoint.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ReleasePublicConnectionRequest
 * @return ReleasePublicConnectionResponse
 */
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

/**
 * *   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](~~58623~~).
 * *   When you call this operation to renew a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be renewed.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request RenewInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RenewInstanceResponse
 */
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

/**
 * *   Before you call this operation, we recommend that you understand the billing of Simple Application Server. For more information, see [Billable items](~~58623~~).
 * *   When you call this operation to renew a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be renewed.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request RenewInstanceRequest
 * @return RenewInstanceResponse
 */
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

/**
 * If the password of your Simple Database Service instance is not strong, you can call this operation to change the password of the administrator account of the instance. To ensure security of the instance, we recommend that you regularly change the password of the instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ResetDatabaseAccountPasswordRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetDatabaseAccountPasswordResponse
 */
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

/**
 * If the password of your Simple Database Service instance is not strong, you can call this operation to change the password of the administrator account of the instance. To ensure security of the instance, we recommend that you regularly change the password of the instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ResetDatabaseAccountPasswordRequest
 * @return ResetDatabaseAccountPasswordResponse
 */
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

/**
 * *   You can call this operation to roll back a disk only if the associated simple application server is in the Stopped state.
 * *   After a disk is rolled back, all data changes that are made from when the snapshot was created to when the disk is rolled back are lost. Back up disk data based on your needs before you roll back the disk.
 * ### Precautions
 * After you reset a simple application server, the disk data on the server is deleted. Snapshots created before the resetting operation are retained but cannot be used to roll back the disks of the server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ResetDiskRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetDiskResponse
 */
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

/**
 * *   You can call this operation to roll back a disk only if the associated simple application server is in the Stopped state.
 * *   After a disk is rolled back, all data changes that are made from when the snapshot was created to when the disk is rolled back are lost. Back up disk data based on your needs before you roll back the disk.
 * ### Precautions
 * After you reset a simple application server, the disk data on the server is deleted. Snapshots created before the resetting operation are retained but cannot be used to roll back the disks of the server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ResetDiskRequest
 * @return ResetDiskResponse
 */
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

/**
 * You can reset a simple application server to re-install its application system or OS and re-initialize the server. You can reset a simple application server by resetting the current system or replacing the image.
 * You can use one of the following methods to reset a simple application server:
 * *   Reset the current system. You can re-install the operating system without replacing the image.
 * *   Replace the image. You can select an Alibaba Cloud image or a custom image that is different from the existing image of the server to reinstall the OS of the server.
 * ### Precautions
 * *   After you reset a simple application server, the disk data on the server is cleared. Back up the data as needed.
 * *   After you reset a simple application server, the monitoring operations that are performed on the server may fail. In this case, you can use one of the following methods to install the CloudMonitor agent on the server:
 *     *   Connect to the server: For more information, see [Manually install the CloudMonitor agent for C++ on an ECS instance](~~183482~~).
 *     *   Use Command Assistant: For more information, see [Use Command Assistant](~~438681~~). You can obtain the command that can be used to install CloudMonitor from the "Common commands" section of the [Use Command Assistant](~~438681~~) topic.
 * ### Limits
 * *   Snapshots that are created before a server is reset are retained, but the snapshots cannot be used to roll back the disks of the server.
 * *   You cannot reset simple application servers that were created based on custom images that contain data of data disks.
 * *   Before you reset a simple application server by replacing the existing image with a custom image, take note of the following items:
 *     *   The custom image must reside in the same region as the current server.
 *     *   The custom image cannot be created based on the current server. If you want to recover the data on the server, you can use a snapshot of the server to roll back the disks of the server.
 *     *   If your simple application server resides outside the Chinese mainland, you cannot switch the OS of the server between Windows Server and Linux. You cannot use a Windows Server custom image to reset a Linux simple application server. You also cannot use a Linux custom image to reset a Windows Server simple application server. You can switch the OSs of simple application servers only between Windows Server OSs or between Linux distributions.
 *     *   The following limits apply to the disks attached to the simple application server:
 *         *   If the custom image contains a system disk and a data disk but only a system disk is attached to the simple application server and no data disk is attached, you cannot use the custom image to reset the simple application server.
 *         *   If the system disk size of the custom image is greater than the system disk size of the simple application server, you cannot directly use the custom image to reset the simple application server.
 *         *   Only if the system disk size of the simple application server is greater than or equal to the system disk size of the custom image, you can use the custom image to reset the simple application server. To increase the system disk size of your simple application server, you can upgrade the server. For more information, see Upgrade a simple application server.
 *         *   If the data disk size of the custom image is greater than the data disk size of the simple application server, you cannot use the custom image to reset the simple application server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ResetSystemRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ResetSystemResponse
 */
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

/**
 * You can reset a simple application server to re-install its application system or OS and re-initialize the server. You can reset a simple application server by resetting the current system or replacing the image.
 * You can use one of the following methods to reset a simple application server:
 * *   Reset the current system. You can re-install the operating system without replacing the image.
 * *   Replace the image. You can select an Alibaba Cloud image or a custom image that is different from the existing image of the server to reinstall the OS of the server.
 * ### Precautions
 * *   After you reset a simple application server, the disk data on the server is cleared. Back up the data as needed.
 * *   After you reset a simple application server, the monitoring operations that are performed on the server may fail. In this case, you can use one of the following methods to install the CloudMonitor agent on the server:
 *     *   Connect to the server: For more information, see [Manually install the CloudMonitor agent for C++ on an ECS instance](~~183482~~).
 *     *   Use Command Assistant: For more information, see [Use Command Assistant](~~438681~~). You can obtain the command that can be used to install CloudMonitor from the "Common commands" section of the [Use Command Assistant](~~438681~~) topic.
 * ### Limits
 * *   Snapshots that are created before a server is reset are retained, but the snapshots cannot be used to roll back the disks of the server.
 * *   You cannot reset simple application servers that were created based on custom images that contain data of data disks.
 * *   Before you reset a simple application server by replacing the existing image with a custom image, take note of the following items:
 *     *   The custom image must reside in the same region as the current server.
 *     *   The custom image cannot be created based on the current server. If you want to recover the data on the server, you can use a snapshot of the server to roll back the disks of the server.
 *     *   If your simple application server resides outside the Chinese mainland, you cannot switch the OS of the server between Windows Server and Linux. You cannot use a Windows Server custom image to reset a Linux simple application server. You also cannot use a Linux custom image to reset a Windows Server simple application server. You can switch the OSs of simple application servers only between Windows Server OSs or between Linux distributions.
 *     *   The following limits apply to the disks attached to the simple application server:
 *         *   If the custom image contains a system disk and a data disk but only a system disk is attached to the simple application server and no data disk is attached, you cannot use the custom image to reset the simple application server.
 *         *   If the system disk size of the custom image is greater than the system disk size of the simple application server, you cannot directly use the custom image to reset the simple application server.
 *         *   Only if the system disk size of the simple application server is greater than or equal to the system disk size of the custom image, you can use the custom image to reset the simple application server. To increase the system disk size of your simple application server, you can upgrade the server. For more information, see Upgrade a simple application server.
 *         *   If the data disk size of the custom image is greater than the data disk size of the simple application server, you cannot use the custom image to reset the simple application server.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request ResetSystemRequest
 * @return ResetSystemResponse
 */
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

/**
 * You can call this operation to restart a Simple Database Service instance that is in the Running state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request RestartDatabaseInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RestartDatabaseInstanceResponse
 */
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

/**
 * You can call this operation to restart a Simple Database Service instance that is in the Running state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request RestartDatabaseInstanceRequest
 * @return RestartDatabaseInstanceResponse
 */
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

/**
 * Command Assistant is an automated O\\&M tool for Simple Application Server. You can maintain simple application servers by running shell, PowerShell, and batch commands in the Simple Application Server console without remotely logging on to the servers.
 * Before you use Command Assistant, take note of the following items:
 * *   The simple application server must be in the Running state.
 * *   The Cloud Assistant client is installed on the server. By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall it. For more information, see [Install the Cloud Assistant Agent](~~64921~~).
 *
 * @param tmpReq RunCommandRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RunCommandResponse
 */
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

/**
 * Command Assistant is an automated O\\&M tool for Simple Application Server. You can maintain simple application servers by running shell, PowerShell, and batch commands in the Simple Application Server console without remotely logging on to the servers.
 * Before you use Command Assistant, take note of the following items:
 * *   The simple application server must be in the Running state.
 * *   The Cloud Assistant client is installed on the server. By default, the Cloud Assistant client is installed on simple application servers. If you have manually uninstalled the client, you must reinstall it. For more information, see [Install the Cloud Assistant Agent](~~64921~~).
 *
 * @param request RunCommandRequest
 * @return RunCommandResponse
 */
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

/**
 * You can call this operation to start a Simple Database Service instance that is in the Stopped state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StartDatabaseInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartDatabaseInstanceResponse
 */
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

/**
 * You can call this operation to start a Simple Database Service instance that is in the Stopped state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StartDatabaseInstanceRequest
 * @return StartDatabaseInstanceResponse
 */
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

/**
 * You can call this operation to start a simple application server that is in the Stopped state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StartInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartInstanceResponse
 */
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

/**
 * You can call this operation to start a simple application server that is in the Stopped state.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StartInstanceRequest
 * @return StartInstanceResponse
 */
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

/**
 * You can call this operation to stop a Simple Database Service instance that is in the Running state. After the instance is stopped, you cannot log on to or access the instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StopDatabaseInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopDatabaseInstanceResponse
 */
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

/**
 * You can call this operation to stop a Simple Database Service instance that is in the Running state. After the instance is stopped, you cannot log on to or access the instance.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StopDatabaseInstanceRequest
 * @return StopDatabaseInstanceResponse
 */
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

/**
 * You can stop a simple application server that you do not use for the time being.
 * >  Stopping a simple application server may interrupt your business. We recommend that you perform the stop operation during off-peak hours.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StopInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopInstanceResponse
 */
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

/**
 * You can stop a simple application server that you do not use for the time being.
 * >  Stopping a simple application server may interrupt your business. We recommend that you perform the stop operation during off-peak hours.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request StopInstanceRequest
 * @return StopInstanceResponse
 */
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

/**
 * ## Usage notes
 * After you change the password of a simple application server, you must restart the server by calling the [RebootInstance](~~190443~~) operation to allow the new password to take effect.
 * ### QPS limits
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request UpdateInstanceAttributeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateInstanceAttributeResponse
 */
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

/**
 * ## Usage notes
 * After you change the password of a simple application server, you must restart the server by calling the [RebootInstance](~~190443~~) operation to allow the new password to take effect.
 * ### QPS limits
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request UpdateInstanceAttributeRequest
 * @return UpdateInstanceAttributeResponse
 */
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

/**
 * *   The plan of a simple application server cannot be downgraded, but can only be upgraded. For more information about plans, see [Billable items](~~58623~~).
 * *   When you call this operation to upgrade a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be upgraded.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request UpgradeInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpgradeInstanceResponse
 */
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

/**
 * *   The plan of a simple application server cannot be downgraded, but can only be upgraded. For more information about plans, see [Billable items](~~58623~~).
 * *   When you call this operation to upgrade a server, make sure that the balance in your account is sufficient. If the balance in your account is insufficient, the server cannot be upgraded.
 * ### QPS limit
 * You can call this API operation up to 10 times per minute per account. Requests that exceed this limit are dropped and you may experience service interruptions. We recommend that you take note of this limit when you call this operation. For more information, see [QPS limits](~~347607~~).
 *
 * @param request UpgradeInstanceRequest
 * @return UpgradeInstanceResponse
 */
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
