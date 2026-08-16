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
	// The delivery group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The delivery group name.
	//
	// example:
	//
	// 办公应用
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The network configuration.
	//
	// > To use this parameter, submit a ticket.
	Network *ModifyAppInstanceGroupAttributeRequestNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The resource group object.
	NodePool *ModifyAppInstanceGroupAttributeRequestNodePool `json:"NodePool,omitempty" xml:"NodePool,omitempty" type:"Struct"`
	// Specifies whether only one application can be opened per session.
	//
	// - If enabled, opening multiple applications within the delivery group allocates a separate session for each application, consuming more sessions.
	//
	// example:
	//
	// false
	PerSessionPerApp *bool `json:"PerSessionPerApp,omitempty" xml:"PerSessionPerApp,omitempty"`
	// The AppId of the pre-open application. If the `PreOpenMode` parameter is set to `SINGLE_APP`, the `PreOpenAppId` parameter cannot be an empty string.
	//
	// example:
	//
	// ca-b2ronxxd****
	PreOpenAppId *string `json:"PreOpenAppId,omitempty" xml:"PreOpenAppId,omitempty"`
	// The pre-open mode.
	//
	// example:
	//
	// OFF
	PreOpenMode *string `json:"PreOpenMode,omitempty" xml:"PreOpenMode,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The security policy.
	SecurityPolicy *ModifyAppInstanceGroupAttributeRequestSecurityPolicy `json:"SecurityPolicy,omitempty" xml:"SecurityPolicy,omitempty" type:"Struct"`
	// The session disconnection retention duration, in minutes. After an end user session is disconnected, the session is retained for the duration specified here before being logged off. Set this parameter to `-1` to retain the session indefinitely. Valid values: -1 and 3 to 300. Default value: `15`.
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
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	if s.NodePool != nil {
		if err := s.NodePool.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityPolicy != nil {
		if err := s.SecurityPolicy.Validate(); err != nil {
			return err
		}
	}
	if s.StoragePolicy != nil {
		if err := s.StoragePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppInstanceGroupAttributeRequestNetwork struct {
	// The domain name rule configurations.
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
	if s.DomainRules != nil {
		for _, item := range s.DomainRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAppInstanceGroupAttributeRequestNetworkDomainRules struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The policy value.
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
	// The number of concurrent sessions, which is the number of sessions that can be simultaneously connected to a single resource. If too many sessions are connected simultaneously, the application experience may degrade. The valid value range varies depending on the resource specification. You can call the ListNodeInstanceType operation to obtain the valid value range for each resource specification.
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The resource group ID.
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
	// Specifies whether to reset after unbinding.
	//
	// example:
	//
	// true
	ResetAfterUnbind *bool `json:"ResetAfterUnbind,omitempty" xml:"ResetAfterUnbind,omitempty"`
	// Specifies whether to skip user authorization verification.
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
	// The list of storage types.
	StorageTypeList []*string `json:"StorageTypeList,omitempty" xml:"StorageTypeList,omitempty" type:"Repeated"`
	// The user data roaming configuration.
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
	if s.UserProfile != nil {
		if err := s.UserProfile.Validate(); err != nil {
			return err
		}
	}
	if s.UserProfileFollow != nil {
		if err := s.UserProfileFollow.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAppInstanceGroupAttributeRequestStoragePolicyUserProfile struct {
	// The user data storage system ID (NAS ID).
	//
	// example:
	//
	// 06ae94****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Specifies whether to enable user data roaming.
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
