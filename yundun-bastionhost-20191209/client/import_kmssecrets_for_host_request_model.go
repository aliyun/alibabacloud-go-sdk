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
	// The ID of the host to import the KMS secrets to.
	//
	// > Only ECS hosts can import KMS secrets. You can call the [ListHosts](https://help.aliyun.com/document_detail/200665.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	HostId *int32 `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about region IDs and names, see [Regions and availability zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The KMS secrets to import.
	Secrets []*ImportKMSSecretsForHostRequestSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
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
	// The name of the KMS secret.
	//
	// example:
	//
	// secret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the KMS secret. Valid values:
	//
	// - **ECS**: an ECS credential.
	//
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
