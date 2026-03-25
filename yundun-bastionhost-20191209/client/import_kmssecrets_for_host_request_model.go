// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKMSSecretsForHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostId(v int32) *ImportKMSSecretsForHostRequest
	GetHostId() *int32
	SetInstanceId(v string) *ImportKMSSecretsForHostRequest
	GetInstanceId() *string
	SetRegionId(v string) *ImportKMSSecretsForHostRequest
	GetRegionId() *string
	SetSecrets(v []*ImportKMSSecretsForHostRequestSecrets) *ImportKMSSecretsForHostRequest
	GetSecrets() []*ImportKMSSecretsForHostRequestSecrets
}

type ImportKMSSecretsForHostRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *int32 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Secrets  []*ImportKMSSecretsForHostRequestSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s ImportKMSSecretsForHostRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKMSSecretsForHostRequest) GoString() string {
	return s.String()
}

func (s *ImportKMSSecretsForHostRequest) GetHostId() *int32 {
	return s.HostId
}

func (s *ImportKMSSecretsForHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportKMSSecretsForHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportKMSSecretsForHostRequest) GetSecrets() []*ImportKMSSecretsForHostRequestSecrets {
	return s.Secrets
}

func (s *ImportKMSSecretsForHostRequest) SetHostId(v int32) *ImportKMSSecretsForHostRequest {
	s.HostId = &v
	return s
}

func (s *ImportKMSSecretsForHostRequest) SetInstanceId(v string) *ImportKMSSecretsForHostRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportKMSSecretsForHostRequest) SetRegionId(v string) *ImportKMSSecretsForHostRequest {
	s.RegionId = &v
	return s
}

func (s *ImportKMSSecretsForHostRequest) SetSecrets(v []*ImportKMSSecretsForHostRequestSecrets) *ImportKMSSecretsForHostRequest {
	s.Secrets = v
	return s
}

func (s *ImportKMSSecretsForHostRequest) Validate() error {
	if s.Secrets != nil {
		for _, item := range s.Secrets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportKMSSecretsForHostRequestSecrets struct {
	// example:
	//
	// secret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// example:
	//
	// ECS
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
}

func (s ImportKMSSecretsForHostRequestSecrets) String() string {
	return dara.Prettify(s)
}

func (s ImportKMSSecretsForHostRequestSecrets) GoString() string {
	return s.String()
}

func (s *ImportKMSSecretsForHostRequestSecrets) GetSecretName() *string {
	return s.SecretName
}

func (s *ImportKMSSecretsForHostRequestSecrets) GetSecretType() *string {
	return s.SecretType
}

func (s *ImportKMSSecretsForHostRequestSecrets) SetSecretName(v string) *ImportKMSSecretsForHostRequestSecrets {
	s.SecretName = &v
	return s
}

func (s *ImportKMSSecretsForHostRequestSecrets) SetSecretType(v string) *ImportKMSSecretsForHostRequestSecrets {
	s.SecretType = &v
	return s
}

func (s *ImportKMSSecretsForHostRequestSecrets) Validate() error {
	return dara.Validate(s)
}
