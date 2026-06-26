// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientSecrets(v []*ListClientSecretsResponseBodyClientSecrets) *ListClientSecretsResponseBody
	GetClientSecrets() []*ListClientSecretsResponseBodyClientSecrets
	SetRequestId(v string) *ListClientSecretsResponseBody
	GetRequestId() *string
}

type ListClientSecretsResponseBody struct {
	ClientSecrets []*ListClientSecretsResponseBodyClientSecrets `json:"ClientSecrets,omitempty" xml:"ClientSecrets,omitempty" type:"Repeated"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientSecretsResponseBody) GetClientSecrets() []*ListClientSecretsResponseBodyClientSecrets {
	return s.ClientSecrets
}

func (s *ListClientSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientSecretsResponseBody) SetClientSecrets(v []*ListClientSecretsResponseBodyClientSecrets) *ListClientSecretsResponseBody {
	s.ClientSecrets = v
	return s
}

func (s *ListClientSecretsResponseBody) SetRequestId(v string) *ListClientSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientSecretsResponseBody) Validate() error {
	if s.ClientSecrets != nil {
		for _, item := range s.ClientSecrets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClientSecretsResponseBodyClientSecrets struct {
	ClientId       *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	ClientName     *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	ClientSecretId *string `json:"ClientSecretId,omitempty" xml:"ClientSecretId,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UserPoolName   *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s ListClientSecretsResponseBodyClientSecrets) String() string {
	return dara.Prettify(s)
}

func (s ListClientSecretsResponseBodyClientSecrets) GoString() string {
	return s.String()
}

func (s *ListClientSecretsResponseBodyClientSecrets) GetClientId() *string {
	return s.ClientId
}

func (s *ListClientSecretsResponseBodyClientSecrets) GetClientName() *string {
	return s.ClientName
}

func (s *ListClientSecretsResponseBodyClientSecrets) GetClientSecretId() *string {
	return s.ClientSecretId
}

func (s *ListClientSecretsResponseBodyClientSecrets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClientSecretsResponseBodyClientSecrets) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *ListClientSecretsResponseBodyClientSecrets) SetClientId(v string) *ListClientSecretsResponseBodyClientSecrets {
	s.ClientId = &v
	return s
}

func (s *ListClientSecretsResponseBodyClientSecrets) SetClientName(v string) *ListClientSecretsResponseBodyClientSecrets {
	s.ClientName = &v
	return s
}

func (s *ListClientSecretsResponseBodyClientSecrets) SetClientSecretId(v string) *ListClientSecretsResponseBodyClientSecrets {
	s.ClientSecretId = &v
	return s
}

func (s *ListClientSecretsResponseBodyClientSecrets) SetCreateTime(v string) *ListClientSecretsResponseBodyClientSecrets {
	s.CreateTime = &v
	return s
}

func (s *ListClientSecretsResponseBodyClientSecrets) SetUserPoolName(v string) *ListClientSecretsResponseBodyClientSecrets {
	s.UserPoolName = &v
	return s
}

func (s *ListClientSecretsResponseBodyClientSecrets) Validate() error {
	return dara.Validate(s)
}
