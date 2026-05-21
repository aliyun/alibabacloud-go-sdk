// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorName(v string) *AddConnectorRequest
	GetConnectorName() *string
	SetConnectorType(v string) *AddConnectorRequest
	GetConnectorType() *string
	SetDescription(v string) *AddConnectorRequest
	GetDescription() *string
	SetFileConnectorConfig(v *AddConnectorRequestFileConnectorConfig) *AddConnectorRequest
	GetFileConnectorConfig() *AddConnectorRequestFileConnectorConfig
}

type AddConnectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-connector
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FILE
	ConnectorType *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	// This parameter is required.
	Description         *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	FileConnectorConfig *AddConnectorRequestFileConnectorConfig `json:"FileConnectorConfig,omitempty" xml:"FileConnectorConfig,omitempty" type:"Struct"`
}

func (s AddConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s AddConnectorRequest) GoString() string {
	return s.String()
}

func (s *AddConnectorRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *AddConnectorRequest) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *AddConnectorRequest) GetDescription() *string {
	return s.Description
}

func (s *AddConnectorRequest) GetFileConnectorConfig() *AddConnectorRequestFileConnectorConfig {
	return s.FileConnectorConfig
}

func (s *AddConnectorRequest) SetConnectorName(v string) *AddConnectorRequest {
	s.ConnectorName = &v
	return s
}

func (s *AddConnectorRequest) SetConnectorType(v string) *AddConnectorRequest {
	s.ConnectorType = &v
	return s
}

func (s *AddConnectorRequest) SetDescription(v string) *AddConnectorRequest {
	s.Description = &v
	return s
}

func (s *AddConnectorRequest) SetFileConnectorConfig(v *AddConnectorRequestFileConnectorConfig) *AddConnectorRequest {
	s.FileConnectorConfig = v
	return s
}

func (s *AddConnectorRequest) Validate() error {
	if s.FileConnectorConfig != nil {
		if err := s.FileConnectorConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddConnectorRequestFileConnectorConfig struct {
	// example:
	//
	// zyb-docker-registry-jn
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// OSS_CUSTOM
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s AddConnectorRequestFileConnectorConfig) String() string {
	return dara.Prettify(s)
}

func (s AddConnectorRequestFileConnectorConfig) GoString() string {
	return s.String()
}

func (s *AddConnectorRequestFileConnectorConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *AddConnectorRequestFileConnectorConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *AddConnectorRequestFileConnectorConfig) GetStorageType() *string {
	return s.StorageType
}

func (s *AddConnectorRequestFileConnectorConfig) SetBucketName(v string) *AddConnectorRequestFileConnectorConfig {
	s.BucketName = &v
	return s
}

func (s *AddConnectorRequestFileConnectorConfig) SetRegionId(v string) *AddConnectorRequestFileConnectorConfig {
	s.RegionId = &v
	return s
}

func (s *AddConnectorRequestFileConnectorConfig) SetStorageType(v string) *AddConnectorRequestFileConnectorConfig {
	s.StorageType = &v
	return s
}

func (s *AddConnectorRequestFileConnectorConfig) Validate() error {
	return dara.Validate(s)
}
