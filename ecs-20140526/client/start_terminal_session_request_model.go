// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartTerminalSessionRequest
	GetClientToken() *string
	SetCommandLine(v string) *StartTerminalSessionRequest
	GetCommandLine() *string
	SetConnectionType(v string) *StartTerminalSessionRequest
	GetConnectionType() *string
	SetEncryptionOptions(v *StartTerminalSessionRequestEncryptionOptions) *StartTerminalSessionRequest
	GetEncryptionOptions() *StartTerminalSessionRequestEncryptionOptions
	SetInstanceId(v []*string) *StartTerminalSessionRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StartTerminalSessionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartTerminalSessionRequest
	GetOwnerId() *int64
	SetPasswordName(v string) *StartTerminalSessionRequest
	GetPasswordName() *string
	SetPortNumber(v int32) *StartTerminalSessionRequest
	GetPortNumber() *int32
	SetRegionId(v string) *StartTerminalSessionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartTerminalSessionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartTerminalSessionRequest
	GetResourceOwnerId() *int64
	SetTargetServer(v string) *StartTerminalSessionRequest
	GetTargetServer() *string
	SetUsername(v string) *StartTerminalSessionRequest
	GetUsername() *string
}

type StartTerminalSessionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The command to run after the session is initiated. The command can be up to 512 characters in length.
	//
	// > After you specify `CommandLine`, you cannot specify `PortNumber` or `TargetServer`.
	//
	// example:
	//
	// ssh root@192.168.0.246
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The network type of the WebSocket URL required for the remote connection to the instance. Valid values:
	//
	// - Internet: the Internet. This is the default value.
	//
	// - Intranet: the internal network.
	//
	// example:
	//
	// Intranet
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The session encryption configuration.
	EncryptionOptions *StartTerminalSessionRequestEncryptionOptions `json:"EncryptionOptions,omitempty" xml:"EncryptionOptions,omitempty" type:"Struct"`
	// The instance ID list.
	//
	// This parameter is required.
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the password for the user when you use Session Manager on a Windows instance. The name can be up to 255 characters in length.
	//
	// When you want to use Session Manager on a Windows instance as a non-default user (System), specify both Username and this parameter. To reduce the risk of password leaks, store the plaintext password in the parameter repository of operations management and specify only the password name here. For more information, see [Encryption parameters](https://help.aliyun.com/document_detail/186828.html).
	//
	// example:
	//
	// axtSecretPassword
	PasswordName *string `json:"PasswordName,omitempty" xml:"PasswordName,omitempty"`
	// The port number of the ECS instance for data forwarding. After this parameter is specified, Cloud Assistant Agent forwards data to the specified port for port forwarding. For example, SSH uses port 22.
	//
	// Default value: empty, which indicates that no port number is specified for data forwarding.
	//
	// example:
	//
	// 22
	PortNumber *int32 `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The address of the destination server in the VPC that you want to access through the instance.
	//
	// > If this parameter is not empty, `PortNumber` specifies the port number of the destination server in the VPC that you want to access through the managed instance.
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

func (s StartTerminalSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionRequest) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartTerminalSessionRequest) GetCommandLine() *string {
	return s.CommandLine
}

func (s *StartTerminalSessionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *StartTerminalSessionRequest) GetEncryptionOptions() *StartTerminalSessionRequestEncryptionOptions {
	return s.EncryptionOptions
}

func (s *StartTerminalSessionRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StartTerminalSessionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartTerminalSessionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTerminalSessionRequest) GetPasswordName() *string {
	return s.PasswordName
}

func (s *StartTerminalSessionRequest) GetPortNumber() *int32 {
	return s.PortNumber
}

func (s *StartTerminalSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTerminalSessionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartTerminalSessionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartTerminalSessionRequest) GetTargetServer() *string {
	return s.TargetServer
}

func (s *StartTerminalSessionRequest) GetUsername() *string {
	return s.Username
}

func (s *StartTerminalSessionRequest) SetClientToken(v string) *StartTerminalSessionRequest {
	s.ClientToken = &v
	return s
}

func (s *StartTerminalSessionRequest) SetCommandLine(v string) *StartTerminalSessionRequest {
	s.CommandLine = &v
	return s
}

func (s *StartTerminalSessionRequest) SetConnectionType(v string) *StartTerminalSessionRequest {
	s.ConnectionType = &v
	return s
}

func (s *StartTerminalSessionRequest) SetEncryptionOptions(v *StartTerminalSessionRequestEncryptionOptions) *StartTerminalSessionRequest {
	s.EncryptionOptions = v
	return s
}

func (s *StartTerminalSessionRequest) SetInstanceId(v []*string) *StartTerminalSessionRequest {
	s.InstanceId = v
	return s
}

func (s *StartTerminalSessionRequest) SetOwnerAccount(v string) *StartTerminalSessionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartTerminalSessionRequest) SetOwnerId(v int64) *StartTerminalSessionRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetPasswordName(v string) *StartTerminalSessionRequest {
	s.PasswordName = &v
	return s
}

func (s *StartTerminalSessionRequest) SetPortNumber(v int32) *StartTerminalSessionRequest {
	s.PortNumber = &v
	return s
}

func (s *StartTerminalSessionRequest) SetRegionId(v string) *StartTerminalSessionRequest {
	s.RegionId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetResourceOwnerAccount(v string) *StartTerminalSessionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTerminalSessionRequest) SetResourceOwnerId(v int64) *StartTerminalSessionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetTargetServer(v string) *StartTerminalSessionRequest {
	s.TargetServer = &v
	return s
}

func (s *StartTerminalSessionRequest) SetUsername(v string) *StartTerminalSessionRequest {
	s.Username = &v
	return s
}

func (s *StartTerminalSessionRequest) Validate() error {
	if s.EncryptionOptions != nil {
		if err := s.EncryptionOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartTerminalSessionRequestEncryptionOptions struct {
	// Specifies whether to enable end-to-end encryption for the session connection.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The KMS key ID.
	//
	// Note:
	//
	// - Only KMS symmetric keys are supported.
	//
	// - This parameter can be specified only when the encryption mode is set to Kms.
	//
	// example:
	//
	// xxx
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The secret key encryption pattern. Valid values:
	//
	// - Auto: uses an automatically negotiated key to encrypt the session.
	//
	// - Kms: uses a KMS key to encrypt the session.
	//
	// - Default value: Auto.
	//
	// Note:
	//
	// - This parameter can be specified only when session encryption is enabled.
	//
	// example:
	//
	// Auto
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s StartTerminalSessionRequestEncryptionOptions) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionRequestEncryptionOptions) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionRequestEncryptionOptions) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartTerminalSessionRequestEncryptionOptions) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *StartTerminalSessionRequestEncryptionOptions) GetMode() *string {
	return s.Mode
}

func (s *StartTerminalSessionRequestEncryptionOptions) SetEnabled(v bool) *StartTerminalSessionRequestEncryptionOptions {
	s.Enabled = &v
	return s
}

func (s *StartTerminalSessionRequestEncryptionOptions) SetKMSKeyId(v string) *StartTerminalSessionRequestEncryptionOptions {
	s.KMSKeyId = &v
	return s
}

func (s *StartTerminalSessionRequestEncryptionOptions) SetMode(v string) *StartTerminalSessionRequestEncryptionOptions {
	s.Mode = &v
	return s
}

func (s *StartTerminalSessionRequestEncryptionOptions) Validate() error {
	return dara.Validate(s)
}
