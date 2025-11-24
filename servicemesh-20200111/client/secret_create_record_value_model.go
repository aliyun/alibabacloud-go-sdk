// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecretCreateRecordValue interface {
	dara.Model
	String() string
	GoString() string
	SetState(v string) *SecretCreateRecordValue
	GetState() *string
	SetClusterId(v string) *SecretCreateRecordValue
	GetClusterId() *string
	SetMessage(v string) *SecretCreateRecordValue
	GetMessage() *string
}

type SecretCreateRecordValue struct {
	// The result of creating the secret. Valid values:
	//
	// 	- `success`: The secret was created.
	//
	// 	- `fail`: The secret failed to be created.
	//
	// example:
	//
	// success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the cluster on the data plane.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8xe10d****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The error message returned when exceptions occur. Otherwise, an empty value is returned.
	//
	// example:
	//
	// timeout
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SecretCreateRecordValue) String() string {
	return dara.Prettify(s)
}

func (s SecretCreateRecordValue) GoString() string {
	return s.String()
}

func (s *SecretCreateRecordValue) GetState() *string {
	return s.State
}

func (s *SecretCreateRecordValue) GetClusterId() *string {
	return s.ClusterId
}

func (s *SecretCreateRecordValue) GetMessage() *string {
	return s.Message
}

func (s *SecretCreateRecordValue) SetState(v string) *SecretCreateRecordValue {
	s.State = &v
	return s
}

func (s *SecretCreateRecordValue) SetClusterId(v string) *SecretCreateRecordValue {
	s.ClusterId = &v
	return s
}

func (s *SecretCreateRecordValue) SetMessage(v string) *SecretCreateRecordValue {
	s.Message = &v
	return s
}

func (s *SecretCreateRecordValue) Validate() error {
	return dara.Validate(s)
}
