// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlannedDeleteTime(v string) *DeleteSecretResponseBody
	GetPlannedDeleteTime() *string
	SetRequestId(v string) *DeleteSecretResponseBody
	GetRequestId() *string
	SetSecretName(v string) *DeleteSecretResponseBody
	GetSecretName() *string
}

type DeleteSecretResponseBody struct {
	// The time when the secret is scheduled to be deleted.
	//
	// example:
	//
	// 2022-09-15T07:02:14Z
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 38bbed2a-15e0-45ad-98d4-816ad2ccf4ea
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s DeleteSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponseBody) GetPlannedDeleteTime() *string {
	return s.PlannedDeleteTime
}

func (s *DeleteSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *DeleteSecretResponseBody) SetPlannedDeleteTime(v string) *DeleteSecretResponseBody {
	s.PlannedDeleteTime = &v
	return s
}

func (s *DeleteSecretResponseBody) SetRequestId(v string) *DeleteSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponseBody) SetSecretName(v string) *DeleteSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
