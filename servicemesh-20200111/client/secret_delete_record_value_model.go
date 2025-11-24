// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSecretDeleteRecordValue interface {
	dara.Model
	String() string
	GoString() string
	SetState(v string) *SecretDeleteRecordValue
	GetState() *string
	SetClusterId(v string) *SecretDeleteRecordValue
	GetClusterId() *string
	SetMessage(v string) *SecretDeleteRecordValue
	GetMessage() *string
}

type SecretDeleteRecordValue struct {
	// The result of deleting the secret. Valid values:
	//
	// 	- `success`: The secret was deleted.
	//
	// 	- `fail`: The secret failed to be deleted.
	//
	// example:
	//
	// success
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The error message returned when exceptions occur. Otherwise, an empty value is returned.
	//
	// example:
	//
	// timeout
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SecretDeleteRecordValue) String() string {
	return dara.Prettify(s)
}

func (s SecretDeleteRecordValue) GoString() string {
	return s.String()
}

func (s *SecretDeleteRecordValue) GetState() *string {
	return s.State
}

func (s *SecretDeleteRecordValue) GetClusterId() *string {
	return s.ClusterId
}

func (s *SecretDeleteRecordValue) GetMessage() *string {
	return s.Message
}

func (s *SecretDeleteRecordValue) SetState(v string) *SecretDeleteRecordValue {
	s.State = &v
	return s
}

func (s *SecretDeleteRecordValue) SetClusterId(v string) *SecretDeleteRecordValue {
	s.ClusterId = &v
	return s
}

func (s *SecretDeleteRecordValue) SetMessage(v string) *SecretDeleteRecordValue {
	s.Message = &v
	return s
}

func (s *SecretDeleteRecordValue) Validate() error {
	return dara.Validate(s)
}
