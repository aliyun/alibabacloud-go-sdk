// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRotateSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArn(v string) *RotateSecretResponseBody
	GetArn() *string
	SetRequestId(v string) *RotateSecretResponseBody
	GetRequestId() *string
	SetSecretName(v string) *RotateSecretResponseBody
	GetSecretName() *string
	SetVersionId(v string) *RotateSecretResponseBody
	GetVersionId() *string
}

type RotateSecretResponseBody struct {
	// The Alibaba Cloud Resource Name (ARN) of the secret.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:secret/RdsSecret/Mysql5.4/MyCred
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 10257c86-269d-43aa-aaf3-90ed4144bb7c
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// RdsSecret/Mysql5.4/MyCred
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The version number of the secret after the secret is rotated.
	//
	// example:
	//
	// 000000123
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s RotateSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RotateSecretResponseBody) GoString() string {
	return s.String()
}

func (s *RotateSecretResponseBody) GetArn() *string {
	return s.Arn
}

func (s *RotateSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RotateSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *RotateSecretResponseBody) GetVersionId() *string {
	return s.VersionId
}

func (s *RotateSecretResponseBody) SetArn(v string) *RotateSecretResponseBody {
	s.Arn = &v
	return s
}

func (s *RotateSecretResponseBody) SetRequestId(v string) *RotateSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *RotateSecretResponseBody) SetSecretName(v string) *RotateSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *RotateSecretResponseBody) SetVersionId(v string) *RotateSecretResponseBody {
	s.VersionId = &v
	return s
}

func (s *RotateSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
