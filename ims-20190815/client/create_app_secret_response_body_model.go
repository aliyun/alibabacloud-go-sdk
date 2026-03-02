// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecret(v *CreateAppSecretResponseBodyAppSecret) *CreateAppSecretResponseBody
	GetAppSecret() *CreateAppSecretResponseBodyAppSecret
	SetRequestId(v string) *CreateAppSecretResponseBody
	GetRequestId() *string
}

type CreateAppSecretResponseBody struct {
	AppSecret *CreateAppSecretResponseBodyAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponseBody) GetAppSecret() *CreateAppSecretResponseBodyAppSecret {
	return s.AppSecret
}

func (s *CreateAppSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppSecretResponseBody) SetAppSecret(v *CreateAppSecretResponseBodyAppSecret) *CreateAppSecretResponseBody {
	s.AppSecret = v
	return s
}

func (s *CreateAppSecretResponseBody) SetRequestId(v string) *CreateAppSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppSecretResponseBody) Validate() error {
	if s.AppSecret != nil {
		if err := s.AppSecret.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppSecretResponseBodyAppSecret struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecretId    *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty"`
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s CreateAppSecretResponseBodyAppSecret) String() string {
	return dara.Prettify(s)
}

func (s CreateAppSecretResponseBodyAppSecret) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponseBodyAppSecret) GetAppId() *string {
	return s.AppId
}

func (s *CreateAppSecretResponseBodyAppSecret) GetAppSecretId() *string {
	return s.AppSecretId
}

func (s *CreateAppSecretResponseBodyAppSecret) GetAppSecretValue() *string {
	return s.AppSecretValue
}

func (s *CreateAppSecretResponseBodyAppSecret) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateAppSecretResponseBodyAppSecret) SetAppId(v string) *CreateAppSecretResponseBodyAppSecret {
	s.AppId = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) SetAppSecretId(v string) *CreateAppSecretResponseBodyAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) SetAppSecretValue(v string) *CreateAppSecretResponseBodyAppSecret {
	s.AppSecretValue = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) SetCreateDate(v string) *CreateAppSecretResponseBodyAppSecret {
	s.CreateDate = &v
	return s
}

func (s *CreateAppSecretResponseBodyAppSecret) Validate() error {
	return dara.Validate(s)
}
