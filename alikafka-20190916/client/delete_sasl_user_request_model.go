// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSaslUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSaslUserRequest
	GetInstanceId() *string
	SetMechanism(v string) *DeleteSaslUserRequest
	GetMechanism() *string
	SetRegionId(v string) *DeleteSaslUserRequest
	GetRegionId() *string
	SetType(v string) *DeleteSaslUserRequest
	GetType() *string
	SetUsername(v string) *DeleteSaslUserRequest
	GetUsername() *string
}

type DeleteSaslUserRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The encryption method. Valid values:
	//
	// 	- SCRAM-SHA-512. This is the default value.
	//
	// 	- SCRAM-SHA-256
	//
	// >  This parameter is available only for serverless ApsaraMQ for Kafka instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the Simple Authentication and Security Layer (SASL) user. Valid values:
	//
	// 	- **plain**: a simple mechanism that uses usernames and passwords to verify user identities. ApsaraMQ for Kafka provides an improved PLAIN mechanism that allows you to dynamically add SASL users without the need to restart an instance.
	//
	// 	- **SCRAM**: a mechanism that uses usernames and passwords to verify user identities. Compared with the PLAIN mechanism, this mechanism provides better security protection. ApsaraMQ for Kafka uses the SCRAM-SHA-256 algorithm.
	//
	// 	- **LDAP**: This value is available only for the SASL users of ApsaraMQ for Confluent instances.
	//
	// Default value: **plain**.
	//
	// example:
	//
	// scram
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The name of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test***
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DeleteSaslUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSaslUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteSaslUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSaslUserRequest) GetMechanism() *string {
	return s.Mechanism
}

func (s *DeleteSaslUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSaslUserRequest) GetType() *string {
	return s.Type
}

func (s *DeleteSaslUserRequest) GetUsername() *string {
	return s.Username
}

func (s *DeleteSaslUserRequest) SetInstanceId(v string) *DeleteSaslUserRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSaslUserRequest) SetMechanism(v string) *DeleteSaslUserRequest {
	s.Mechanism = &v
	return s
}

func (s *DeleteSaslUserRequest) SetRegionId(v string) *DeleteSaslUserRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSaslUserRequest) SetType(v string) *DeleteSaslUserRequest {
	s.Type = &v
	return s
}

func (s *DeleteSaslUserRequest) SetUsername(v string) *DeleteSaslUserRequest {
	s.Username = &v
	return s
}

func (s *DeleteSaslUserRequest) Validate() error {
	return dara.Validate(s)
}
