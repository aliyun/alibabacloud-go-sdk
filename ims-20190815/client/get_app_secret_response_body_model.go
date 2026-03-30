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
	// The details of the application secret.
	AppSecret *GetAppSecretResponseBodyAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EE46FC3C-3BDE-4771-B531-27B7B6EB533D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The content of the application secret.
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
