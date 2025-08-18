// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeRequest
	GetAppInstanceGroupName() *string
	SetNetwork(v *ModifyAppInstanceGroupAttributeRequestNetwork) *ModifyAppInstanceGroupAttributeRequest
	GetNetwork() *ModifyAppInstanceGroupAttributeRequestNetwork
	SetNodePool(v *ModifyAppInstanceGroupAttributeRequestNodePool) *ModifyAppInstanceGroupAttributeRequest
	GetNodePool() *ModifyAppInstanceGroupAttributeRequestNodePool
	SetPerSessionPerApp(v bool) *ModifyAppInstanceGroupAttributeRequest
	GetPerSessionPerApp() *bool
	SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeRequest
	GetPreOpenAppId() *string
	SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeRequest
	GetPreOpenMode() *string
	SetProductType(v string) *ModifyAppInstanceGroupAttributeRequest
	GetProductType() *string
	SetSecurityPolicy(v *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) *ModifyAppInstanceGroupAttributeRequest
	GetSecurityPolicy() *ModifyAppInstanceGroupAttributeRequestSecurityPolicy
	SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeRequest
	GetSessionTimeout() *int32
	SetStoragePolicy(v *ModifyAppInstanceGroupAttributeRequestStoragePolicy) *ModifyAppInstanceGroupAttributeRequest
	GetStoragePolicy() *ModifyAppInstanceGroupAttributeRequestStoragePolicy
}

type ModifyAppInstanceGroupAttributeRequest struct {
	// The ID of the delivery group.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery group.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The network settings.
	//
	// >  If you want to use this parameter, submit a ticket.
	Network *ModifyAppInstanceGroupAttributeRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The information about the resource group.
	NodePool *ModifyAppInstanceGroupAttributeRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// Specifies whether only one application can be opened in a session.
	//
	// 	- After you enable this feature, the system assigns a session to each application if you open multiple applications in a delivery group. This consumes a larger number of sessions.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// The application ID of the pre-open application. If you set `PreOpenMode` to `SINGLE_APP`, you cannot leave this parameter empty.``
	//
	// example:
	//
	// ca-b2ronxxd****
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// The pre-open mode.
	//
	// Valid values:
	//
	// 	- SINGLE_APP: enables the pre-open mode for a single application.
	//
	// 	- OFF: disables the pre-open mode. This is the default value.
	//
	// example:
	//
	// OFF
	PreOpenMode *string `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The security policy.
	SecurityPolicy *ModifyAppInstanceGroupAttributeRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The duration for which sessions are retained after disconnection. Unit: minutes. After an end user disconnects from a session, the session is closed only after the specified duration elapses. If you want to permanently retain sessions, set this parameter to `-1`. Valid values:-1 and 3 to 300. Default value: `15`.
	//
	// example:
	//
	// 15
	SessionTimeout *int32 `json:"SessionTimeout,omitempty" xml:"SessionTimeout,omitempty"`
	// The storage policy.
	StoragePolicy *ModifyAppInstanceGroupAttributeRequestStoragePolicy `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty" type:"Struct"`
}

func (s ModifyAppInstanceGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetNetwork() *ModifyAppInstanceGroupAttributeRequestNetwork {
	return s.Network
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetNodePool() *ModifyAppInstanceGroupAttributeRequestNodePool {
	return s.NodePool
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetPerSessionPerApp() *bool {
	return s.PerSessionPerApp
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetPreOpenAppId() *string {
	return s.PreOpenAppId
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetPreOpenMode() *string {
	return s.PreOpenMode
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetSecurityPolicy() *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	return s.SecurityPolicy
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetSessionTimeout() *int32 {
	return s.SessionTimeout
}

func (s *ModifyAppInstanceGroupAttributeRequest) GetStoragePolicy() *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	return s.StoragePolicy
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetAppInstanceGroupName(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNetwork(v *ModifyAppInstanceGroupAttributeRequestNetwork) *ModifyAppInstanceGroupAttributeRequest {
	s.Network = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetNodePool(v *ModifyAppInstanceGroupAttributeRequestNodePool) *ModifyAppInstanceGroupAttributeRequest {
	s.NodePool = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPerSessionPerApp(v bool) *ModifyAppInstanceGroupAttributeRequest {
	s.PerSessionPerApp = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPreOpenAppId(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.PreOpenAppId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetPreOpenMode(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.PreOpenMode = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetProductType(v string) *ModifyAppInstanceGroupAttributeRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSecurityPolicy(v *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) *ModifyAppInstanceGroupAttributeRequest {
	s.SecurityPolicy = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetSessionTimeout(v int32) *ModifyAppInstanceGroupAttributeRequest {
	s.SessionTimeout = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) SetStoragePolicy(v *ModifyAppInstanceGroupAttributeRequestStoragePolicy) *ModifyAppInstanceGroupAttributeRequest {
	s.StoragePolicy = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestNetwork struct {
	// The domain name rules.
	DomainRules []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules `json:"DomainRules,omitempty" xml:"DomainRules,omitempty" type:"Repeated"`
}

func (s ModifyAppInstanceGroupAttributeRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNetwork) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNetwork) GetDomainRules() []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	return s.DomainRules
}

func (s *ModifyAppInstanceGroupAttributeRequestNetwork) SetDomainRules(v []*ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) *ModifyAppInstanceGroupAttributeRequestNetwork {
	s.DomainRules = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNetwork) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestNetworkDomainRules struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The policy used for the domain name.
	//
	// Valid values:
	//
	// 	- allow
	//
	// 	- block
	//
	// example:
	//
	// block
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) GetDomain() *string {
	return s.Domain
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) GetPolicy() *string {
	return s.Policy
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) SetDomain(v string) *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	s.Domain = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) SetPolicy(v string) *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules {
	s.Policy = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNetworkDomainRules) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestNodePool struct {
	// The maximum number of sessions to which a resource can connect at the same time. If a resource connects to a large number of sessions at the same time, user experience can be compromised. The value range varies based on the resource type. The following items describe the value ranges of different resource types:
	//
	// 	- appstreaming.general.4c8g: 1 to 2
	//
	// 	- appstreaming.general.8c16g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c16g.4g: 1 to 4
	//
	// 	- appstreaming.vgpu.8c31g.16g: 1 to 4
	//
	// 	- appstreaming.vgpu.14c93g.12g: 1 to 6
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-ew7va2g1wl3vm****
	NodePoolId *string `json:"NodePoolId,omitempty" xml:"NodePoolId,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestNodePool) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) GetNodeCapacity() *int32 {
	return s.NodeCapacity
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) GetNodePoolId() *string {
	return s.NodePoolId
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodeCapacity(v int32) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) SetNodePoolId(v string) *ModifyAppInstanceGroupAttributeRequestNodePool {
	s.NodePoolId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestNodePool) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestSecurityPolicy struct {
	// Specifies whether to reset after unbinding from a delivery group.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ResetAfterUnbind *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	// Specifies whether to skip user permission verification.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false: This is the default value.
	//
	// example:
	//
	// false
	SkipUserAuthCheck *bool `json:"SkipUserAuthCheck,omitempty" xml:"SkipUserAuthCheck,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestSecurityPolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestSecurityPolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) GetResetAfterUnbind() *bool {
	return s.ResetAfterUnbind
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) GetSkipUserAuthCheck() *bool {
	return s.SkipUserAuthCheck
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) SetResetAfterUnbind(v bool) *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	s.ResetAfterUnbind = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) SetSkipUserAuthCheck(v bool) *ModifyAppInstanceGroupAttributeRequestSecurityPolicy {
	s.SkipUserAuthCheck = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestSecurityPolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicy struct {
	// The storage types.
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
	// The configurations of user data roaming.
	UserProfile       *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile       `json:"UserProfile,omitempty" xml:"UserProfile,omitempty" type:"Struct"`
	UserProfileFollow *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow `json:"UserProfileFollow,omitempty" xml:"UserProfileFollow,omitempty" type:"Struct"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicy) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicy) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) GetStorageTypeList() []*string {
	return s.StorageTypeList
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) GetUserProfile() *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile {
	return s.UserProfile
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) GetUserProfileFollow() *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow {
	return s.UserProfileFollow
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetStorageTypeList(v []*string) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.StorageTypeList = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetUserProfile(v *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.UserProfile = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) SetUserProfileFollow(v *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) *ModifyAppInstanceGroupAttributeRequestStoragePolicy {
	s.UserProfileFollow = v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicy) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile struct {
	// The ID of the File Storage NAS (NAS) file system used to store user data.
	//
	// example:
	//
	// 06ae94****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Specifies whether user data roaming is enabled.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	UserProfileSwitch *bool `json:"UserProfileSwitch,omitempty" xml:"UserProfileSwitch,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) GetUserProfileSwitch() *bool {
	return s.UserProfileSwitch
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) SetFileSystemId(v string) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) SetUserProfileSwitch(v bool) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile {
	s.UserProfileSwitch = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile) Validate() error {
	return dara.Validate(s)
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow struct {
	FileSystemId        *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	ProfileFollowSwitch *bool   `json:"ProfileFollowSwitch,omitempty" xml:"ProfileFollowSwitch,omitempty"`
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) GetProfileFollowSwitch() *bool {
	return s.ProfileFollowSwitch
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) SetFileSystemId(v string) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) SetProfileFollowSwitch(v bool) *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow {
	s.ProfileFollowSwitch = &v
	return s
}

func (s *ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfileFollow) Validate() error {
	return dara.Validate(s)
}
