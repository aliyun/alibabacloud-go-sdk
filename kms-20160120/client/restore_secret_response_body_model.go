// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreSecretResponseBody
	GetRequestId() *string
	SetSecretName(v string) *RestoreSecretResponseBody
	GetSecretName() *string
}

type RestoreSecretResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// e4885adf-548f-4ca5-8075-f540bbd3a55f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// secret001
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s RestoreSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreSecretResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreSecretResponseBody) GetSecretName() *string {
	return s.SecretName
}

func (s *RestoreSecretResponseBody) SetRequestId(v string) *RestoreSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreSecretResponseBody) SetSecretName(v string) *RestoreSecretResponseBody {
	s.SecretName = &v
	return s
}

func (s *RestoreSecretResponseBody) Validate() error {
	return dara.Validate(s)
}
