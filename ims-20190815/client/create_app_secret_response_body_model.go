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
	// The details of the application secret.
	AppSecret *CreateAppSecretResponseBodyAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EE46FC3C-3BDE-4771-B531-27B7B6EB533D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the application.
	//
	// example:
	//
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the application secret.
	//
	// example:
	//
	// 2efd5004-005c-4f05-83c6-5b1dd176****
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty"`
	// The content of the application secret. This value can be used as the client secret for open authorization (OAuth).
	//
	// example:
	//
	// ai78ZmmxnlUG1jXlBZRDFKos9DIjY4m17Q7dCpMwn1rqXsTGb1X1XmrmveMp****
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-26T02:52:31Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
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
