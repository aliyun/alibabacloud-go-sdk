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
	SetProvider(v string) *VectorStoreConfig
	GetProvider() *string
}

type VectorStoreConfig struct {
	Config   *VectorStoreConfigConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	Provider *string                  `json:"provider,omitempty" xml:"provider,omitempty"`
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

func (s *VectorStoreConfig) GetProvider() *string {
	return s.Provider
}

func (s *VectorStoreConfig) SetConfig(v *VectorStoreConfigConfig) *VectorStoreConfig {
	s.Config = v
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
