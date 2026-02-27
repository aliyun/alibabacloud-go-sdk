// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceLoginInfo(v *LoginInstanceRequestInstanceLoginInfo) *LoginInstanceRequest
	GetInstanceLoginInfo() *LoginInstanceRequestInstanceLoginInfo
	SetPartnerInfo(v *LoginInstanceRequestPartnerInfo) *LoginInstanceRequest
	GetPartnerInfo() *LoginInstanceRequestPartnerInfo
	SetUserAccount(v *LoginInstanceRequestUserAccount) *LoginInstanceRequest
	GetUserAccount() *LoginInstanceRequestUserAccount
}

type LoginInstanceRequest struct {
	InstanceLoginInfo *LoginInstanceRequestInstanceLoginInfo `json:"InstanceLoginInfo,omitempty" xml:"InstanceLoginInfo,omitempty" type:"Struct"`
	PartnerInfo       *LoginInstanceRequestPartnerInfo       `json:"PartnerInfo,omitempty" xml:"PartnerInfo,omitempty" type:"Struct"`
	UserAccount       *LoginInstanceRequestUserAccount       `json:"UserAccount,omitempty" xml:"UserAccount,omitempty" type:"Struct"`
}

func (s LoginInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequest) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequest) GetInstanceLoginInfo() *LoginInstanceRequestInstanceLoginInfo {
	return s.InstanceLoginInfo
}

func (s *LoginInstanceRequest) GetPartnerInfo() *LoginInstanceRequestPartnerInfo {
	return s.PartnerInfo
}

func (s *LoginInstanceRequest) GetUserAccount() *LoginInstanceRequestUserAccount {
	return s.UserAccount
}

func (s *LoginInstanceRequest) SetInstanceLoginInfo(v *LoginInstanceRequestInstanceLoginInfo) *LoginInstanceRequest {
	s.InstanceLoginInfo = v
	return s
}

func (s *LoginInstanceRequest) SetPartnerInfo(v *LoginInstanceRequestPartnerInfo) *LoginInstanceRequest {
	s.PartnerInfo = v
	return s
}

func (s *LoginInstanceRequest) SetUserAccount(v *LoginInstanceRequestUserAccount) *LoginInstanceRequest {
	s.UserAccount = v
	return s
}

func (s *LoginInstanceRequest) Validate() error {
	if s.InstanceLoginInfo != nil {
		if err := s.InstanceLoginInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PartnerInfo != nil {
		if err := s.PartnerInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UserAccount != nil {
		if err := s.UserAccount.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceRequestInstanceLoginInfo struct {
	// example:
	//
	// password/certificate
	AuthenticationType *string `json:"AuthenticationType,omitempty" xml:"AuthenticationType,omitempty"`
	// example:
	//
	// ----begin----
	//
	// ----end----
	Certificate         *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	CredentialToken     *string `json:"CredentialToken,omitempty" xml:"CredentialToken,omitempty"`
	DockerContainerName *string `json:"DockerContainerName,omitempty" xml:"DockerContainerName,omitempty"`
	DockerExec          *string `json:"DockerExec,omitempty" xml:"DockerExec,omitempty"`
	// example:
	//
	// 123
	DurationSeconds   *int64                                                  `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	EncryptionOptions *LoginInstanceRequestInstanceLoginInfoEncryptionOptions `json:"EncryptionOptions,omitempty" xml:"EncryptionOptions,omitempty" type:"Struct"`
	// example:
	//
	// 2022-11-30 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 127.0.0.1
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// i-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ecs/eci/ack
	InstanceType              *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	LoginByInstanceCredential *bool   `json:"LoginByInstanceCredential,omitempty" xml:"LoginByInstanceCredential,omitempty"`
	LoginByInstanceShortcut   *bool   `json:"LoginByInstanceShortcut,omitempty" xml:"LoginByInstanceShortcut,omitempty"`
	// example:
	//
	// vpc/classic
	NetworkAccessMode *string                                       `json:"NetworkAccessMode,omitempty" xml:"NetworkAccessMode,omitempty"`
	Options           *LoginInstanceRequestInstanceLoginInfoOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// example:
	//
	// xxxx
	PassPhrase *string `json:"PassPhrase,omitempty" xml:"PassPhrase,omitempty"`
	// example:
	//
	// xxxxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 22/3389
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// ssh/rdp/ack
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// cn-hangzhou/cn-beijing
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ShortcutToken   *string `json:"ShortcutToken,omitempty" xml:"ShortcutToken,omitempty"`
	// example:
	//
	// root/Administrator
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// vpc-abc
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfo) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfo) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetCertificate() *string {
	return s.Certificate
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetCredentialToken() *string {
	return s.CredentialToken
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetDockerContainerName() *string {
	return s.DockerContainerName
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetDockerExec() *string {
	return s.DockerExec
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetEncryptionOptions() *LoginInstanceRequestInstanceLoginInfoEncryptionOptions {
	return s.EncryptionOptions
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetHost() *string {
	return s.Host
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetInstanceType() *string {
	return s.InstanceType
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetLoginByInstanceCredential() *bool {
	return s.LoginByInstanceCredential
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetLoginByInstanceShortcut() *bool {
	return s.LoginByInstanceShortcut
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetNetworkAccessMode() *string {
	return s.NetworkAccessMode
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetOptions() *LoginInstanceRequestInstanceLoginInfoOptions {
	return s.Options
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetPassPhrase() *string {
	return s.PassPhrase
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetPassword() *string {
	return s.Password
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetPort() *int32 {
	return s.Port
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetShortcutToken() *string {
	return s.ShortcutToken
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetUsername() *string {
	return s.Username
}

func (s *LoginInstanceRequestInstanceLoginInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetAuthenticationType(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.AuthenticationType = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetCertificate(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Certificate = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetCredentialToken(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.CredentialToken = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetDockerContainerName(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.DockerContainerName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetDockerExec(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.DockerExec = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetDurationSeconds(v int64) *LoginInstanceRequestInstanceLoginInfo {
	s.DurationSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetEncryptionOptions(v *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) *LoginInstanceRequestInstanceLoginInfo {
	s.EncryptionOptions = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetExpireTime(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.ExpireTime = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetHost(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Host = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetInstanceId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.InstanceId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetInstanceType(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.InstanceType = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetLoginByInstanceCredential(v bool) *LoginInstanceRequestInstanceLoginInfo {
	s.LoginByInstanceCredential = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetLoginByInstanceShortcut(v bool) *LoginInstanceRequestInstanceLoginInfo {
	s.LoginByInstanceShortcut = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetNetworkAccessMode(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.NetworkAccessMode = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetOptions(v *LoginInstanceRequestInstanceLoginInfoOptions) *LoginInstanceRequestInstanceLoginInfo {
	s.Options = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetPassPhrase(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.PassPhrase = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetPassword(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Password = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetPort(v int32) *LoginInstanceRequestInstanceLoginInfo {
	s.Port = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetProtocol(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Protocol = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetRegionId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.RegionId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetResourceGroupId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetShortcutToken(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.ShortcutToken = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetUsername(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.Username = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) SetVpcId(v string) *LoginInstanceRequestInstanceLoginInfo {
	s.VpcId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfo) Validate() error {
	if s.EncryptionOptions != nil {
		if err := s.EncryptionOptions.Validate(); err != nil {
			return err
		}
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceRequestInstanceLoginInfoEncryptionOptions struct {
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfoEncryptionOptions) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfoEncryptionOptions) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) GetEnabled() *bool {
	return s.Enabled
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) GetMode() *string {
	return s.Mode
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) SetEnabled(v bool) *LoginInstanceRequestInstanceLoginInfoEncryptionOptions {
	s.Enabled = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) SetKMSKeyId(v string) *LoginInstanceRequestInstanceLoginInfoEncryptionOptions {
	s.KMSKeyId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) SetMode(v string) *LoginInstanceRequestInstanceLoginInfoEncryptionOptions {
	s.Mode = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoEncryptionOptions) Validate() error {
	return dara.Validate(s)
}

type LoginInstanceRequestInstanceLoginInfoOptions struct {
	AudioMuteSeconds *int32                                                     `json:"AudioMuteSeconds,omitempty" xml:"AudioMuteSeconds,omitempty"`
	ContainerInfo    *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo `json:"ContainerInfo,omitempty" xml:"ContainerInfo,omitempty" type:"Struct"`
	FixedHeight      *int32                                                     `json:"FixedHeight,omitempty" xml:"FixedHeight,omitempty"`
	FixedWidth       *int32                                                     `json:"FixedWidth,omitempty" xml:"FixedWidth,omitempty"`
	// example:
	//
	// abc
	NotificationEventTypes *string `json:"NotificationEventTypes,omitempty" xml:"NotificationEventTypes,omitempty"`
	// example:
	//
	// abc
	NotificationRecipientUrl *string `json:"NotificationRecipientUrl,omitempty" xml:"NotificationRecipientUrl,omitempty"`
	// example:
	//
	// 10
	NotificationRetryIntervalSeconds *int32 `json:"NotificationRetryIntervalSeconds,omitempty" xml:"NotificationRetryIntervalSeconds,omitempty"`
	// example:
	//
	// 3
	NotificationRetryLimit  *int32 `json:"NotificationRetryLimit,omitempty" xml:"NotificationRetryLimit,omitempty"`
	OperationDisableSeconds *int32 `json:"OperationDisableSeconds,omitempty" xml:"OperationDisableSeconds,omitempty"`
	// example:
	//
	// abc
	SessionControl     *string `json:"SessionControl,omitempty" xml:"SessionControl,omitempty"`
	VideoFreezeSeconds *int32  `json:"VideoFreezeSeconds,omitempty" xml:"VideoFreezeSeconds,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfoOptions) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfoOptions) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetAudioMuteSeconds() *int32 {
	return s.AudioMuteSeconds
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetContainerInfo() *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	return s.ContainerInfo
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetFixedHeight() *int32 {
	return s.FixedHeight
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetFixedWidth() *int32 {
	return s.FixedWidth
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetNotificationEventTypes() *string {
	return s.NotificationEventTypes
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetNotificationRecipientUrl() *string {
	return s.NotificationRecipientUrl
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetNotificationRetryIntervalSeconds() *int32 {
	return s.NotificationRetryIntervalSeconds
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetNotificationRetryLimit() *int32 {
	return s.NotificationRetryLimit
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetOperationDisableSeconds() *int32 {
	return s.OperationDisableSeconds
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetSessionControl() *string {
	return s.SessionControl
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) GetVideoFreezeSeconds() *int32 {
	return s.VideoFreezeSeconds
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetAudioMuteSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.AudioMuteSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetContainerInfo(v *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.ContainerInfo = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetFixedHeight(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.FixedHeight = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetFixedWidth(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.FixedWidth = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationEventTypes(v string) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationEventTypes = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationRecipientUrl(v string) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationRecipientUrl = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationRetryIntervalSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationRetryIntervalSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetNotificationRetryLimit(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.NotificationRetryLimit = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetOperationDisableSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.OperationDisableSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetSessionControl(v string) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.SessionControl = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) SetVideoFreezeSeconds(v int32) *LoginInstanceRequestInstanceLoginInfoOptions {
	s.VideoFreezeSeconds = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptions) Validate() error {
	if s.ContainerInfo != nil {
		if err := s.ContainerInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo struct {
	// example:
	//
	// abcdef
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// abc
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// abc
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// example:
	//
	// abc
	Deployment *string `json:"Deployment,omitempty" xml:"Deployment,omitempty"`
	// example:
	//
	// abc
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// {"abc":"def"}
	Headers map[string]interface{} `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// example:
	//
	// abc
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// abc
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetClusterName() *string {
	return s.ClusterName
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetContainerName() *string {
	return s.ContainerName
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetDeployment() *string {
	return s.Deployment
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetHeaders() map[string]interface{} {
	return s.Headers
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) GetPodName() *string {
	return s.PodName
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetClusterId(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.ClusterId = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetClusterName(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.ClusterName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetContainerName(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.ContainerName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetDeployment(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Deployment = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetEndpoint(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Endpoint = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetHeaders(v map[string]interface{}) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Headers = v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetNamespace(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.Namespace = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) SetPodName(v string) *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo {
	s.PodName = &v
	return s
}

func (s *LoginInstanceRequestInstanceLoginInfoOptionsContainerInfo) Validate() error {
	return dara.Validate(s)
}

type LoginInstanceRequestPartnerInfo struct {
	// example:
	//
	// abc
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// example:
	//
	// abc
	PartnerName *string `json:"PartnerName,omitempty" xml:"PartnerName,omitempty"`
}

func (s LoginInstanceRequestPartnerInfo) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestPartnerInfo) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestPartnerInfo) GetPartnerId() *string {
	return s.PartnerId
}

func (s *LoginInstanceRequestPartnerInfo) GetPartnerName() *string {
	return s.PartnerName
}

func (s *LoginInstanceRequestPartnerInfo) SetPartnerId(v string) *LoginInstanceRequestPartnerInfo {
	s.PartnerId = &v
	return s
}

func (s *LoginInstanceRequestPartnerInfo) SetPartnerName(v string) *LoginInstanceRequestPartnerInfo {
	s.PartnerName = &v
	return s
}

func (s *LoginInstanceRequestPartnerInfo) Validate() error {
	return dara.Validate(s)
}

type LoginInstanceRequestUserAccount struct {
	// example:
	//
	// 1234
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// aas
	AccountPlatform *string `json:"AccountPlatform,omitempty" xml:"AccountPlatform,omitempty"`
	// example:
	//
	// 2/3/4
	AccountStructure *string `json:"AccountStructure,omitempty" xml:"AccountStructure,omitempty"`
	// example:
	//
	// 100
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	// example:
	//
	// 123abc
	EmpId *string `json:"EmpId,omitempty" xml:"EmpId,omitempty"`
	// example:
	//
	// 2022-11-30 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// abc
	LoginName *string                                 `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	Options   *LoginInstanceRequestUserAccountOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
	// example:
	//
	// 1234
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s LoginInstanceRequestUserAccount) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestUserAccount) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestUserAccount) GetAccountId() *int64 {
	return s.AccountId
}

func (s *LoginInstanceRequestUserAccount) GetAccountPlatform() *string {
	return s.AccountPlatform
}

func (s *LoginInstanceRequestUserAccount) GetAccountStructure() *string {
	return s.AccountStructure
}

func (s *LoginInstanceRequestUserAccount) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *LoginInstanceRequestUserAccount) GetEmpId() *string {
	return s.EmpId
}

func (s *LoginInstanceRequestUserAccount) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *LoginInstanceRequestUserAccount) GetLoginName() *string {
	return s.LoginName
}

func (s *LoginInstanceRequestUserAccount) GetOptions() *LoginInstanceRequestUserAccountOptions {
	return s.Options
}

func (s *LoginInstanceRequestUserAccount) GetParentId() *int64 {
	return s.ParentId
}

func (s *LoginInstanceRequestUserAccount) SetAccountId(v int64) *LoginInstanceRequestUserAccount {
	s.AccountId = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetAccountPlatform(v string) *LoginInstanceRequestUserAccount {
	s.AccountPlatform = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetAccountStructure(v string) *LoginInstanceRequestUserAccount {
	s.AccountStructure = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetDurationSeconds(v int64) *LoginInstanceRequestUserAccount {
	s.DurationSeconds = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetEmpId(v string) *LoginInstanceRequestUserAccount {
	s.EmpId = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetExpireTime(v string) *LoginInstanceRequestUserAccount {
	s.ExpireTime = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetLoginName(v string) *LoginInstanceRequestUserAccount {
	s.LoginName = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetOptions(v *LoginInstanceRequestUserAccountOptions) *LoginInstanceRequestUserAccount {
	s.Options = v
	return s
}

func (s *LoginInstanceRequestUserAccount) SetParentId(v int64) *LoginInstanceRequestUserAccount {
	s.ParentId = &v
	return s
}

func (s *LoginInstanceRequestUserAccount) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LoginInstanceRequestUserAccountOptions struct {
	// example:
	//
	// 3
	LoginLimit *int64 `json:"LoginLimit,omitempty" xml:"LoginLimit,omitempty"`
}

func (s LoginInstanceRequestUserAccountOptions) String() string {
	return dara.Prettify(s)
}

func (s LoginInstanceRequestUserAccountOptions) GoString() string {
	return s.String()
}

func (s *LoginInstanceRequestUserAccountOptions) GetLoginLimit() *int64 {
	return s.LoginLimit
}

func (s *LoginInstanceRequestUserAccountOptions) SetLoginLimit(v int64) *LoginInstanceRequestUserAccountOptions {
	s.LoginLimit = &v
	return s
}

func (s *LoginInstanceRequestUserAccountOptions) Validate() error {
	return dara.Validate(s)
}
