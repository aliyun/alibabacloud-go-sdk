// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWebAuthnConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAaguids(v []*SetWebAuthnConfigurationRequestAaguids) *SetWebAuthnConfigurationRequest
	GetAaguids() []*SetWebAuthnConfigurationRequestAaguids
	SetEnableAaguidVerification(v bool) *SetWebAuthnConfigurationRequest
	GetEnableAaguidVerification() *bool
	SetEnableMetadataServiceVerification(v bool) *SetWebAuthnConfigurationRequest
	GetEnableMetadataServiceVerification() *bool
	SetEnableUserSelfRegistration(v bool) *SetWebAuthnConfigurationRequest
	GetEnableUserSelfRegistration() *bool
	SetInstanceId(v string) *SetWebAuthnConfigurationRequest
	GetInstanceId() *string
}

type SetWebAuthnConfigurationRequest struct {
	// AAGUID及其名称列表
	Aaguids []*SetWebAuthnConfigurationRequestAaguids `json:"Aaguids,omitempty" xml:"Aaguids,omitempty" type:"Repeated"`
	// 是否开启AAGUID校验
	//
	// example:
	//
	// true
	EnableAaguidVerification *bool `json:"EnableAaguidVerification,omitempty" xml:"EnableAaguidVerification,omitempty"`
	// 是否开启WebAuthn认证器MDS校验
	//
	// example:
	//
	// true
	EnableMetadataServiceVerification *bool `json:"EnableMetadataServiceVerification,omitempty" xml:"EnableMetadataServiceVerification,omitempty"`
	// 是否允许用户自注册WebAuthn认证器
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableUserSelfRegistration *bool `json:"EnableUserSelfRegistration,omitempty" xml:"EnableUserSelfRegistration,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetWebAuthnConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWebAuthnConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetWebAuthnConfigurationRequest) GetAaguids() []*SetWebAuthnConfigurationRequestAaguids {
	return s.Aaguids
}

func (s *SetWebAuthnConfigurationRequest) GetEnableAaguidVerification() *bool {
	return s.EnableAaguidVerification
}

func (s *SetWebAuthnConfigurationRequest) GetEnableMetadataServiceVerification() *bool {
	return s.EnableMetadataServiceVerification
}

func (s *SetWebAuthnConfigurationRequest) GetEnableUserSelfRegistration() *bool {
	return s.EnableUserSelfRegistration
}

func (s *SetWebAuthnConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetWebAuthnConfigurationRequest) SetAaguids(v []*SetWebAuthnConfigurationRequestAaguids) *SetWebAuthnConfigurationRequest {
	s.Aaguids = v
	return s
}

func (s *SetWebAuthnConfigurationRequest) SetEnableAaguidVerification(v bool) *SetWebAuthnConfigurationRequest {
	s.EnableAaguidVerification = &v
	return s
}

func (s *SetWebAuthnConfigurationRequest) SetEnableMetadataServiceVerification(v bool) *SetWebAuthnConfigurationRequest {
	s.EnableMetadataServiceVerification = &v
	return s
}

func (s *SetWebAuthnConfigurationRequest) SetEnableUserSelfRegistration(v bool) *SetWebAuthnConfigurationRequest {
	s.EnableUserSelfRegistration = &v
	return s
}

func (s *SetWebAuthnConfigurationRequest) SetInstanceId(v string) *SetWebAuthnConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetWebAuthnConfigurationRequest) Validate() error {
	if s.Aaguids != nil {
		for _, item := range s.Aaguids {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetWebAuthnConfigurationRequestAaguids struct {
	// AAGUID
	//
	// example:
	//
	// adce0002-35bc-c60a-648b-0b25f1f05503
	Aaguid *string `json:"Aaguid,omitempty" xml:"Aaguid,omitempty"`
	// AAGUID名称
	//
	// example:
	//
	// Touch ID
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetWebAuthnConfigurationRequestAaguids) String() string {
	return dara.Prettify(s)
}

func (s SetWebAuthnConfigurationRequestAaguids) GoString() string {
	return s.String()
}

func (s *SetWebAuthnConfigurationRequestAaguids) GetAaguid() *string {
	return s.Aaguid
}

func (s *SetWebAuthnConfigurationRequestAaguids) GetName() *string {
	return s.Name
}

func (s *SetWebAuthnConfigurationRequestAaguids) SetAaguid(v string) *SetWebAuthnConfigurationRequestAaguids {
	s.Aaguid = &v
	return s
}

func (s *SetWebAuthnConfigurationRequestAaguids) SetName(v string) *SetWebAuthnConfigurationRequestAaguids {
	s.Name = &v
	return s
}

func (s *SetWebAuthnConfigurationRequestAaguids) Validate() error {
	return dara.Validate(s)
}
