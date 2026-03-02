// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppSecret(v *GetAppSecretResponseBodyAppSecret) *GetAppSecretResponseBody
	GetAppSecret() *GetAppSecretResponseBodyAppSecret
	SetRequestId(v string) *GetAppSecretResponseBody
	GetRequestId() *string
}

type GetAppSecretResponseBody struct {
	AppSecret *GetAppSecretResponseBodyAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Struct"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAppSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseBody) GetAppSecret() *GetAppSecretResponseBodyAppSecret {
	return s.AppSecret
}

func (s *GetAppSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppSecretResponseBody) SetAppSecret(v *GetAppSecretResponseBodyAppSecret) *GetAppSecretResponseBody {
	s.AppSecret = v
	return s
}

func (s *GetAppSecretResponseBody) SetRequestId(v string) *GetAppSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppSecretResponseBody) Validate() error {
	if s.AppSecret != nil {
		if err := s.AppSecret.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppSecretResponseBodyAppSecret struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecretId    *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty"`
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
}

func (s GetAppSecretResponseBodyAppSecret) String() string {
	return dara.Prettify(s)
}

func (s GetAppSecretResponseBodyAppSecret) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseBodyAppSecret) GetAppId() *string {
	return s.AppId
}

func (s *GetAppSecretResponseBodyAppSecret) GetAppSecretId() *string {
	return s.AppSecretId
}

func (s *GetAppSecretResponseBodyAppSecret) GetAppSecretValue() *string {
	return s.AppSecretValue
}

func (s *GetAppSecretResponseBodyAppSecret) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetAppSecretResponseBodyAppSecret) SetAppId(v string) *GetAppSecretResponseBodyAppSecret {
	s.AppId = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) SetAppSecretId(v string) *GetAppSecretResponseBodyAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) SetAppSecretValue(v string) *GetAppSecretResponseBodyAppSecret {
	s.AppSecretValue = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) SetCreateDate(v string) *GetAppSecretResponseBodyAppSecret {
	s.CreateDate = &v
	return s
}

func (s *GetAppSecretResponseBodyAppSecret) Validate() error {
	return dara.Validate(s)
}
