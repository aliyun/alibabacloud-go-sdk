// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaslUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateSaslUserRequest
	GetInstanceId() *string
	SetMechanism(v string) *CreateSaslUserRequest
	GetMechanism() *string
	SetPassword(v string) *CreateSaslUserRequest
	GetPassword() *string
	SetRegionId(v string) *CreateSaslUserRequest
	GetRegionId() *string
	SetType(v string) *CreateSaslUserRequest
	GetType() *string
	SetUsername(v string) *CreateSaslUserRequest
	GetUsername() *string
}

type CreateSaslUserRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The encryption mechanism. Valid values:
	//
	// - SCRAM-SHA-512 (selected by default)
	//
	// - SCRAM-SHA-256
	//
	// >This parameter is supported only for Serverless instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// The password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type. Valid values:
	//
	// - plain: a simple username and password verification mechanism. ApsaraMQ for Kafka has optimized the PLAIN mechanism to support dynamically adding SASL users without restarting the instance.
	//
	// - scram: a username and password verification mechanism that provides higher security than PLAIN. ApsaraMQ for Kafka uses SCRAM-SHA-256.
	//
	// - LDAP: applicable only to adding users for Confluent instances.
	//
	// Default value: plain.
	//
	// example:
	//
	// plain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateSaslUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSaslUserRequest) GoString() string {
	return s.String()
}

func (s *CreateSaslUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSaslUserRequest) GetMechanism() *string {
	return s.Mechanism
}

func (s *CreateSaslUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateSaslUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSaslUserRequest) GetType() *string {
	return s.Type
}

func (s *CreateSaslUserRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateSaslUserRequest) SetInstanceId(v string) *CreateSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSaslUserRequest) SetMechanism(v string) *CreateSaslUserRequest {
	s.Mechanism = &v
	return s
}

func (s *CreateSaslUserRequest) SetPassword(v string) *CreateSaslUserRequest {
	s.Password = &v
	return s
}

func (s *CreateSaslUserRequest) SetRegionId(v string) *CreateSaslUserRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSaslUserRequest) SetType(v string) *CreateSaslUserRequest {
	s.Type = &v
	return s
}

func (s *CreateSaslUserRequest) SetUsername(v string) *CreateSaslUserRequest {
	s.Username = &v
	return s
}

func (s *CreateSaslUserRequest) Validate() error {
	return dara.Validate(s)
}
