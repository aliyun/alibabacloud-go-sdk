// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVectorStoreConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *VectorStoreConfigConfig) *VectorStoreConfig
	GetConfig() *VectorStoreConfigConfig
	SetMysqlConfig(v *VectorStoreConfigMysqlConfig) *VectorStoreConfig
	GetMysqlConfig() *VectorStoreConfigMysqlConfig
	SetProvider(v string) *VectorStoreConfig
	GetProvider() *string
}

type VectorStoreConfig struct {
	Config      *VectorStoreConfigConfig      `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	MysqlConfig *VectorStoreConfigMysqlConfig `json:"mysqlConfig,omitempty" xml:"mysqlConfig,omitempty" type:"Struct"`
	Provider    *string                       `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s VectorStoreConfig) String() string {
	return dara.Prettify(s)
}

func (s VectorStoreConfig) GoString() string {
	return s.String()
}

func (s *VectorStoreConfig) GetConfig() *VectorStoreConfigConfig {
	return s.Config
}

func (s *VectorStoreConfig) GetMysqlConfig() *VectorStoreConfigMysqlConfig {
	return s.MysqlConfig
}

func (s *VectorStoreConfig) GetProvider() *string {
	return s.Provider
}

func (s *VectorStoreConfig) SetConfig(v *VectorStoreConfigConfig) *VectorStoreConfig {
	s.Config = v
	return s
}

func (s *VectorStoreConfig) SetMysqlConfig(v *VectorStoreConfigMysqlConfig) *VectorStoreConfig {
	s.MysqlConfig = v
	return s
}

func (s *VectorStoreConfig) SetProvider(v string) *VectorStoreConfig {
	s.Provider = &v
	return s
}

func (s *VectorStoreConfig) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.MysqlConfig != nil {
		if err := s.MysqlConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VectorStoreConfigConfig struct {
	CollectionName  *string `json:"collectionName,omitempty" xml:"collectionName,omitempty"`
	Endpoint        *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	InstanceName    *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	VectorDimension *int32  `json:"vectorDimension,omitempty" xml:"vectorDimension,omitempty"`
}

func (s VectorStoreConfigConfig) String() string {
	return dara.Prettify(s)
}

func (s VectorStoreConfigConfig) GoString() string {
	return s.String()
}

func (s *VectorStoreConfigConfig) GetCollectionName() *string {
	return s.CollectionName
}

func (s *VectorStoreConfigConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *VectorStoreConfigConfig) GetInstanceName() *string {
	return s.InstanceName
}

func (s *VectorStoreConfigConfig) GetVectorDimension() *int32 {
	return s.VectorDimension
}

func (s *VectorStoreConfigConfig) SetCollectionName(v string) *VectorStoreConfigConfig {
	s.CollectionName = &v
	return s
}

func (s *VectorStoreConfigConfig) SetEndpoint(v string) *VectorStoreConfigConfig {
	s.Endpoint = &v
	return s
}

func (s *VectorStoreConfigConfig) SetInstanceName(v string) *VectorStoreConfigConfig {
	s.InstanceName = &v
	return s
}

func (s *VectorStoreConfigConfig) SetVectorDimension(v int32) *VectorStoreConfigConfig {
	s.VectorDimension = &v
	return s
}

func (s *VectorStoreConfigConfig) Validate() error {
	return dara.Validate(s)
}

type VectorStoreConfigMysqlConfig struct {
	CollectionName  *string `json:"collectionName,omitempty" xml:"collectionName,omitempty"`
	CredentialName  *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	DbName          *string `json:"dbName,omitempty" xml:"dbName,omitempty"`
	Host            *string `json:"host,omitempty" xml:"host,omitempty"`
	Port            *int32  `json:"port,omitempty" xml:"port,omitempty"`
	User            *string `json:"user,omitempty" xml:"user,omitempty"`
	VectorDimension *int32  `json:"vectorDimension,omitempty" xml:"vectorDimension,omitempty"`
}

func (s VectorStoreConfigMysqlConfig) String() string {
	return dara.Prettify(s)
}

func (s VectorStoreConfigMysqlConfig) GoString() string {
	return s.String()
}

func (s *VectorStoreConfigMysqlConfig) GetCollectionName() *string {
	return s.CollectionName
}

func (s *VectorStoreConfigMysqlConfig) GetCredentialName() *string {
	return s.CredentialName
}

func (s *VectorStoreConfigMysqlConfig) GetDbName() *string {
	return s.DbName
}

func (s *VectorStoreConfigMysqlConfig) GetHost() *string {
	return s.Host
}

func (s *VectorStoreConfigMysqlConfig) GetPort() *int32 {
	return s.Port
}

func (s *VectorStoreConfigMysqlConfig) GetUser() *string {
	return s.User
}

func (s *VectorStoreConfigMysqlConfig) GetVectorDimension() *int32 {
	return s.VectorDimension
}

func (s *VectorStoreConfigMysqlConfig) SetCollectionName(v string) *VectorStoreConfigMysqlConfig {
	s.CollectionName = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) SetCredentialName(v string) *VectorStoreConfigMysqlConfig {
	s.CredentialName = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) SetDbName(v string) *VectorStoreConfigMysqlConfig {
	s.DbName = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) SetHost(v string) *VectorStoreConfigMysqlConfig {
	s.Host = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) SetPort(v int32) *VectorStoreConfigMysqlConfig {
	s.Port = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) SetUser(v string) *VectorStoreConfigMysqlConfig {
	s.User = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) SetVectorDimension(v int32) *VectorStoreConfigMysqlConfig {
	s.VectorDimension = &v
	return s
}

func (s *VectorStoreConfigMysqlConfig) Validate() error {
	return dara.Validate(s)
}
