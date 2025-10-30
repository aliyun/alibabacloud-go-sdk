// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListSecretsResponseBody
	GetCount() *int64
	SetMessage(v string) *ListSecretsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecretsResponseBody
	GetRequestId() *string
	SetSecrets(v *ListSecretsResponseBodySecrets) *ListSecretsResponseBody
	GetSecrets() *ListSecretsResponseBodySecrets
	SetStatus(v string) *ListSecretsResponseBody
	GetStatus() *string
}

type ListSecretsResponseBody struct {
	// The number of access credentials.
	//
	// example:
	//
	// 4
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried access credentials.
	Secrets *ListSecretsResponseBodySecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Struct"`
	// The status of the operation. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListSecretsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretsResponseBody) GetSecrets() *ListSecretsResponseBodySecrets {
	return s.Secrets
}

func (s *ListSecretsResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListSecretsResponseBody) SetCount(v int64) *ListSecretsResponseBody {
	s.Count = &v
	return s
}

func (s *ListSecretsResponseBody) SetMessage(v string) *ListSecretsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecretsResponseBody) SetRequestId(v string) *ListSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretsResponseBody) SetSecrets(v *ListSecretsResponseBodySecrets) *ListSecretsResponseBody {
	s.Secrets = v
	return s
}

func (s *ListSecretsResponseBody) SetStatus(v string) *ListSecretsResponseBody {
	s.Status = &v
	return s
}

func (s *ListSecretsResponseBody) Validate() error {
	if s.Secrets != nil {
		if err := s.Secrets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecretsResponseBodySecrets struct {
	Secrets []*ListSecretsResponseBodySecretsSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s ListSecretsResponseBodySecrets) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodySecrets) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecrets) GetSecrets() []*ListSecretsResponseBodySecretsSecrets {
	return s.Secrets
}

func (s *ListSecretsResponseBodySecrets) SetSecrets(v []*ListSecretsResponseBodySecretsSecrets) *ListSecretsResponseBodySecrets {
	s.Secrets = v
	return s
}

func (s *ListSecretsResponseBodySecrets) Validate() error {
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

type ListSecretsResponseBodySecretsSecrets struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1033***
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gp-bp14****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The description of the access credential.
	//
	// example:
	//
	// test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the access credential for the created Data API account. Format: `acs:gpdb:{{region}}:{{accountId}}:secret/{{secretName}}-{{32 digits random string}`.
	//
	// example:
	//
	// acs:gpdb:cn-beijing:1033**:secret/testsecret-eG2AQGRIwQ0zFp4VA7mYL3uiCXTfDQbQ
	SecretArn *string `json:"SecretArn,omitempty" xml:"SecretArn,omitempty"`
	// The name of the access credential.
	//
	// example:
	//
	// testsecret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The name of the database account.
	//
	// example:
	//
	// testacc
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListSecretsResponseBodySecretsSecrets) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodySecretsSecrets) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodySecretsSecrets) GetAccountId() *string {
	return s.AccountId
}

func (s *ListSecretsResponseBodySecretsSecrets) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ListSecretsResponseBodySecretsSecrets) GetDescription() *string {
	return s.Description
}

func (s *ListSecretsResponseBodySecretsSecrets) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSecretsResponseBodySecretsSecrets) GetSecretArn() *string {
	return s.SecretArn
}

func (s *ListSecretsResponseBodySecretsSecrets) GetSecretName() *string {
	return s.SecretName
}

func (s *ListSecretsResponseBodySecretsSecrets) GetUsername() *string {
	return s.Username
}

func (s *ListSecretsResponseBodySecretsSecrets) SetAccountId(v string) *ListSecretsResponseBodySecretsSecrets {
	s.AccountId = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) SetDBInstanceId(v string) *ListSecretsResponseBodySecretsSecrets {
	s.DBInstanceId = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) SetDescription(v string) *ListSecretsResponseBodySecretsSecrets {
	s.Description = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) SetRegionId(v string) *ListSecretsResponseBodySecretsSecrets {
	s.RegionId = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) SetSecretArn(v string) *ListSecretsResponseBodySecretsSecrets {
	s.SecretArn = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) SetSecretName(v string) *ListSecretsResponseBodySecretsSecrets {
	s.SecretName = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) SetUsername(v string) *ListSecretsResponseBodySecretsSecrets {
	s.Username = &v
	return s
}

func (s *ListSecretsResponseBodySecretsSecrets) Validate() error {
	return dara.Validate(s)
}
