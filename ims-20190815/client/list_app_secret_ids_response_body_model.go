// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppSecretIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecrets(v *ListAppSecretIdsResponseBodyAppSecrets) *ListAppSecretIdsResponseBody
	GetAppSecrets() *ListAppSecretIdsResponseBodyAppSecrets
	SetRequestId(v string) *ListAppSecretIdsResponseBody
	GetRequestId() *string
}

type ListAppSecretIdsResponseBody struct {
	AppSecrets *ListAppSecretIdsResponseBodyAppSecrets `json:"AppSecrets,omitempty" xml:"AppSecrets,omitempty" type:"Struct"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAppSecretIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppSecretIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseBody) GetAppSecrets() *ListAppSecretIdsResponseBodyAppSecrets {
	return s.AppSecrets
}

func (s *ListAppSecretIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppSecretIdsResponseBody) SetAppSecrets(v *ListAppSecretIdsResponseBodyAppSecrets) *ListAppSecretIdsResponseBody {
	s.AppSecrets = v
	return s
}

func (s *ListAppSecretIdsResponseBody) SetRequestId(v string) *ListAppSecretIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppSecretIdsResponseBody) Validate() error {
	if s.AppSecrets != nil {
		if err := s.AppSecrets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppSecretIdsResponseBodyAppSecrets struct {
	AppSecret []*ListAppSecretIdsResponseBodyAppSecretsAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Repeated"`
}

func (s ListAppSecretIdsResponseBodyAppSecrets) String() string {
	return dara.Prettify(s)
}

func (s ListAppSecretIdsResponseBodyAppSecrets) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseBodyAppSecrets) GetAppSecret() []*ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	return s.AppSecret
}

func (s *ListAppSecretIdsResponseBodyAppSecrets) SetAppSecret(v []*ListAppSecretIdsResponseBodyAppSecretsAppSecret) *ListAppSecretIdsResponseBodyAppSecrets {
	s.AppSecret = v
	return s
}

func (s *ListAppSecretIdsResponseBodyAppSecrets) Validate() error {
	if s.AppSecret != nil {
		for _, item := range s.AppSecret {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppSecretIdsResponseBodyAppSecretsAppSecret struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s ListAppSecretIdsResponseBodyAppSecretsAppSecret) String() string {
	return dara.Prettify(s)
}

func (s ListAppSecretIdsResponseBodyAppSecretsAppSecret) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) GetAppId() *string {
	return s.AppId
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) GetAppSecretId() *string {
	return s.AppSecretId
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) SetAppId(v string) *ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	s.AppId = &v
	return s
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) SetAppSecretId(v string) *ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) SetCreateDate(v string) *ListAppSecretIdsResponseBodyAppSecretsAppSecret {
	s.CreateDate = &v
	return s
}

func (s *ListAppSecretIdsResponseBodyAppSecretsAppSecret) Validate() error {
	return dara.Validate(s)
}
