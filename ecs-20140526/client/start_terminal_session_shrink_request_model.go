// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartTerminalSessionShrinkRequest
	GetClientToken() *string
	SetCommandLine(v string) *StartTerminalSessionShrinkRequest
	GetCommandLine() *string
	SetConnectionType(v string) *StartTerminalSessionShrinkRequest
	GetConnectionType() *string
	SetEncryptionOptionsShrink(v string) *StartTerminalSessionShrinkRequest
	GetEncryptionOptionsShrink() *string
	SetInstanceId(v []*string) *StartTerminalSessionShrinkRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StartTerminalSessionShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartTerminalSessionShrinkRequest
	GetOwnerId() *int64
	SetPasswordName(v string) *StartTerminalSessionShrinkRequest
	GetPasswordName() *string
	SetPortNumber(v int32) *StartTerminalSessionShrinkRequest
	GetPortNumber() *int32
	SetRegionId(v string) *StartTerminalSessionShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartTerminalSessionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartTerminalSessionShrinkRequest
	GetResourceOwnerId() *int64
	SetTargetServer(v string) *StartTerminalSessionShrinkRequest
	GetTargetServer() *string
	SetUsername(v string) *StartTerminalSessionShrinkRequest
	GetUsername() *string
}

type StartTerminalSessionShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command to be executed after the session is initiated. The command cannot exceed 512 characters in length.
	//
	// > After CommandLine is specified, you cannot specify PortNumber or TargetServer.
	//
	// example:
	//
	// ssh root@192.168.0.246
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The network type of the WebSocket URL required for a remote connection to the instance. Valid values:
	//
	// - Internet: the Internet. This is the default value.
	//
	// - Intranet: the internal network.
	//
	// example:
	//
	// Intranet
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The session encryption configurations.
	EncryptionOptionsShrink *string `json:"EncryptionOptions,omitempty" xml:"EncryptionOptions,omitempty"`
	// The list of instance IDs.
	//
	// This parameter is required.
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the user password used when Session Manager is used on a Windows instance. The name cannot exceed 255 characters in length.
	//
	// If you want to use Session Manager on a Windows instance as a non-default user (System), you must specify both the Username parameter and this parameter. To reduce the risk of password leakage, the plaintext password must be hosted in the operations management parameter repository, and only the password name is passed here. For more information, see [Encrypted parameters](https://help.aliyun.com/document_detail/186828.html).
	//
	// example:
	//
	// axtSecretPassword
	PasswordName *string `json:"PasswordName,omitempty" xml:"PasswordName,omitempty"`
	// The port number of the ECS instance, which is used for data forwarding. After this parameter is specified, data forwarding of Cloud Assistant Agent is directed to the specified port number for port forwarding. For example, SSH uses port 22.
	//
	// Default value: null, which indicates that the port number for data forwarding is not specified.
	//
	// example:
	//
	// 22
	PortNumber *int32 `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The destination service address in the VPC that is accessed through the instance.
	//
	// > When this parameter is not empty, PortNumber specifies the port number of the destination service in the VPC that is accessed through the managed instance.
	//
	// example:
	//
	// 192.168.0.246
	TargetServer *string `json:"TargetServer,omitempty" xml:"TargetServer,omitempty"`
	// The username used for the connection.
	//
	// example:
	//
	// testUser
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s StartTerminalSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartTerminalSessionShrinkRequest) GetCommandLine() *string {
	return s.CommandLine
}

func (s *StartTerminalSessionShrinkRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *StartTerminalSessionShrinkRequest) GetEncryptionOptionsShrink() *string {
	return s.EncryptionOptionsShrink
}

func (s *StartTerminalSessionShrinkRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StartTerminalSessionShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartTerminalSessionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTerminalSessionShrinkRequest) GetPasswordName() *string {
	return s.PasswordName
}

func (s *StartTerminalSessionShrinkRequest) GetPortNumber() *int32 {
	return s.PortNumber
}

func (s *StartTerminalSessionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTerminalSessionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartTerminalSessionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartTerminalSessionShrinkRequest) GetTargetServer() *string {
	return s.TargetServer
}

func (s *StartTerminalSessionShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *StartTerminalSessionShrinkRequest) SetClientToken(v string) *StartTerminalSessionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetCommandLine(v string) *StartTerminalSessionShrinkRequest {
	s.CommandLine = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetConnectionType(v string) *StartTerminalSessionShrinkRequest {
	s.ConnectionType = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetEncryptionOptionsShrink(v string) *StartTerminalSessionShrinkRequest {
	s.EncryptionOptionsShrink = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetInstanceId(v []*string) *StartTerminalSessionShrinkRequest {
	s.InstanceId = v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetOwnerAccount(v string) *StartTerminalSessionShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetOwnerId(v int64) *StartTerminalSessionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetPasswordName(v string) *StartTerminalSessionShrinkRequest {
	s.PasswordName = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetPortNumber(v int32) *StartTerminalSessionShrinkRequest {
	s.PortNumber = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetRegionId(v string) *StartTerminalSessionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetResourceOwnerAccount(v string) *StartTerminalSessionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetResourceOwnerId(v int64) *StartTerminalSessionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetTargetServer(v string) *StartTerminalSessionShrinkRequest {
	s.TargetServer = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetUsername(v string) *StartTerminalSessionShrinkRequest {
	s.Username = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
