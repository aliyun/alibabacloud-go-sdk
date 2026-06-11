// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiPolicyRedisConfig interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseNumber(v int32) *AiPolicyRedisConfig
	GetDatabaseNumber() *int32
	SetHost(v string) *AiPolicyRedisConfig
	GetHost() *string
	SetPassword(v string) *AiPolicyRedisConfig
	GetPassword() *string
	SetPort(v int32) *AiPolicyRedisConfig
	GetPort() *int32
	SetTimeout(v int32) *AiPolicyRedisConfig
	GetTimeout() *int32
	SetUsername(v string) *AiPolicyRedisConfig
	GetUsername() *string
}

type AiPolicyRedisConfig struct {
	// The zero-based index of the Redis database to use. The default value is 0.
	DatabaseNumber *int32 `json:"databaseNumber,omitempty" xml:"databaseNumber,omitempty"`
	// The hostname of the Redis instance.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// The password for Redis authentication.
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The port number of the Redis instance.
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The connection timeout in milliseconds.
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The username for Redis authentication.
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s AiPolicyRedisConfig) String() string {
	return dara.Prettify(s)
}

func (s AiPolicyRedisConfig) GoString() string {
	return s.String()
}

func (s *AiPolicyRedisConfig) GetDatabaseNumber() *int32 {
	return s.DatabaseNumber
}

func (s *AiPolicyRedisConfig) GetHost() *string {
	return s.Host
}

func (s *AiPolicyRedisConfig) GetPassword() *string {
	return s.Password
}

func (s *AiPolicyRedisConfig) GetPort() *int32 {
	return s.Port
}

func (s *AiPolicyRedisConfig) GetTimeout() *int32 {
	return s.Timeout
}

func (s *AiPolicyRedisConfig) GetUsername() *string {
	return s.Username
}

func (s *AiPolicyRedisConfig) SetDatabaseNumber(v int32) *AiPolicyRedisConfig {
	s.DatabaseNumber = &v
	return s
}

func (s *AiPolicyRedisConfig) SetHost(v string) *AiPolicyRedisConfig {
	s.Host = &v
	return s
}

func (s *AiPolicyRedisConfig) SetPassword(v string) *AiPolicyRedisConfig {
	s.Password = &v
	return s
}

func (s *AiPolicyRedisConfig) SetPort(v int32) *AiPolicyRedisConfig {
	s.Port = &v
	return s
}

func (s *AiPolicyRedisConfig) SetTimeout(v int32) *AiPolicyRedisConfig {
	s.Timeout = &v
	return s
}

func (s *AiPolicyRedisConfig) SetUsername(v string) *AiPolicyRedisConfig {
	s.Username = &v
	return s
}

func (s *AiPolicyRedisConfig) Validate() error {
	return dara.Validate(s)
}
