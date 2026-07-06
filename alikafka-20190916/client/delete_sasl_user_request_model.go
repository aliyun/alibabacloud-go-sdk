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
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Encryption method. Valid values:
	//
	// - SCRAM-SHA-512 (selected by default)
	//
	// - SCRAM-SHA-256
	//
	// > This parameter is only supported for Serverless instances.
	//
	// example:
	//
	// SCRAM-SHA-256
	Mechanism *string `json:"Mechanism,omitempty" xml:"Mechanism,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Type. Valid values:
	//
	// - **plain**: A simple username and password verification mechanism. MSMQ optimizes the PLAIN mechanism to support adding SASL users dynamically without restarting the instance.
	//
	// - **scram**: A username and password verification mechanism with higher security than PLAIN. MSMQ uses SCRAM-SHA-256.
	//
	// - **LDAP**: Only applicable for deleting Confluent instance users.
	//
	// Default value: **plain**.
	//
	// example:
	//
	// scram
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Username.
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
