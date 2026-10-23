// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReInitDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoStartInstance(v bool) *ReInitDiskRequest
	GetAutoStartInstance() *bool
	SetDiskId(v string) *ReInitDiskRequest
	GetDiskId() *string
	SetKeyPairName(v string) *ReInitDiskRequest
	GetKeyPairName() *string
	SetOwnerAccount(v string) *ReInitDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReInitDiskRequest
	GetOwnerId() *int64
	SetPassword(v string) *ReInitDiskRequest
	GetPassword() *string
	SetResourceOwnerAccount(v string) *ReInitDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReInitDiskRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *ReInitDiskRequest
	GetSecurityEnhancementStrategy() *string
}

type ReInitDiskRequest struct {
	// Specifies whether to automatically start the instance after the cloud disk is reinitialized. Valid values:
	//
	// - true: Yes. The instance is automatically started.
	//
	// - false: No.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoStartInstance *bool `json:"AutoStartInstance,omitempty" xml:"AutoStartInstance,omitempty"`
	// The ID of the cloud disk to be reinitialized. The instance to which the cloud disk is attached must have been started at least once and must currently be in the Stopped state.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4ph****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the key pair.
	//
	// > This parameter is applicable only to Linux instances. When you reinitialize the system cloud disk, you can bind an SSH key pair to the ECS instance as the logon credential. After an SSH key pair is used, the password logon method is disabled.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName  *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to reset the logon password of the ECS instance when you reinitialize the system cloud disk. The password must be 8 to 30 characters in length and must contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	// ```
	//
	// ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// ```
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If you specify the `Password` parameter, send the request over HTTPS to prevent password leaks.
	//
	// example:
	//
	// EcsV587!
	Password             *string `json:"Password,omitempty" xml:"Password,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to use Security Center for free after the cloud disk is reinitialized when the specified cloud disk is a system cloud disk. Valid values:
	//
	//
	//
	// - Active: Uses Security Center. This value is supported only for public images.
	//
	// - Deactive: Does not use Security Center. This value is supported for all images.
	//
	// Default value: Deactive.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
}

func (s ReInitDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s ReInitDiskRequest) GoString() string {
	return s.String()
}

func (s *ReInitDiskRequest) GetAutoStartInstance() *bool {
	return s.AutoStartInstance
}

func (s *ReInitDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ReInitDiskRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ReInitDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReInitDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReInitDiskRequest) GetPassword() *string {
	return s.Password
}

func (s *ReInitDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReInitDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReInitDiskRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *ReInitDiskRequest) SetAutoStartInstance(v bool) *ReInitDiskRequest {
	s.AutoStartInstance = &v
	return s
}

func (s *ReInitDiskRequest) SetDiskId(v string) *ReInitDiskRequest {
	s.DiskId = &v
	return s
}

func (s *ReInitDiskRequest) SetKeyPairName(v string) *ReInitDiskRequest {
	s.KeyPairName = &v
	return s
}

func (s *ReInitDiskRequest) SetOwnerAccount(v string) *ReInitDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReInitDiskRequest) SetOwnerId(v int64) *ReInitDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *ReInitDiskRequest) SetPassword(v string) *ReInitDiskRequest {
	s.Password = &v
	return s
}

func (s *ReInitDiskRequest) SetResourceOwnerAccount(v string) *ReInitDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReInitDiskRequest) SetResourceOwnerId(v int64) *ReInitDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReInitDiskRequest) SetSecurityEnhancementStrategy(v string) *ReInitDiskRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *ReInitDiskRequest) Validate() error {
	return dara.Validate(s)
}
