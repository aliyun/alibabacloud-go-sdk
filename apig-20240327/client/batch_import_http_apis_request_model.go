// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchImportHttpApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUpdate(v bool) *BatchImportHttpApisRequest
	GetAllowUpdate() *bool
	SetApiType(v string) *BatchImportHttpApisRequest
	GetApiType() *string
	SetDryRun(v bool) *BatchImportHttpApisRequest
	GetDryRun() *bool
	SetGatewayId(v string) *BatchImportHttpApisRequest
	GetGatewayId() *string
	SetResourceGroupId(v string) *BatchImportHttpApisRequest
	GetResourceGroupId() *string
	SetSpecFileUrl(v string) *BatchImportHttpApisRequest
	GetSpecFileUrl() *string
	SetSpecOssConfig(v *BatchImportHttpApisRequestSpecOssConfig) *BatchImportHttpApisRequest
	GetSpecOssConfig() *BatchImportHttpApisRequestSpecOssConfig
	SetStrategy(v string) *BatchImportHttpApisRequest
	GetStrategy() *string
	SetWithGatewayExtension(v bool) *BatchImportHttpApisRequest
	GetWithGatewayExtension() *bool
}

type BatchImportHttpApisRequest struct {
	AllowUpdate *bool `json:"allowUpdate,omitempty" xml:"allowUpdate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Http
	ApiType *string `json:"apiType,omitempty" xml:"apiType,omitempty"`
	DryRun  *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// rg-xxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com/my-bucket/imports/batch.zip
	SpecFileUrl   *string                                  `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	SpecOssConfig *BatchImportHttpApisRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// example:
	//
	// ExistFirst
	Strategy             *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	WithGatewayExtension *bool   `json:"withGatewayExtension,omitempty" xml:"withGatewayExtension,omitempty"`
}

func (s BatchImportHttpApisRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchImportHttpApisRequest) GoString() string {
	return s.String()
}

func (s *BatchImportHttpApisRequest) GetAllowUpdate() *bool {
	return s.AllowUpdate
}

func (s *BatchImportHttpApisRequest) GetApiType() *string {
	return s.ApiType
}

func (s *BatchImportHttpApisRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *BatchImportHttpApisRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *BatchImportHttpApisRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *BatchImportHttpApisRequest) GetSpecFileUrl() *string {
	return s.SpecFileUrl
}

func (s *BatchImportHttpApisRequest) GetSpecOssConfig() *BatchImportHttpApisRequestSpecOssConfig {
	return s.SpecOssConfig
}

func (s *BatchImportHttpApisRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *BatchImportHttpApisRequest) GetWithGatewayExtension() *bool {
	return s.WithGatewayExtension
}

func (s *BatchImportHttpApisRequest) SetAllowUpdate(v bool) *BatchImportHttpApisRequest {
	s.AllowUpdate = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetApiType(v string) *BatchImportHttpApisRequest {
	s.ApiType = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetDryRun(v bool) *BatchImportHttpApisRequest {
	s.DryRun = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetGatewayId(v string) *BatchImportHttpApisRequest {
	s.GatewayId = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetResourceGroupId(v string) *BatchImportHttpApisRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetSpecFileUrl(v string) *BatchImportHttpApisRequest {
	s.SpecFileUrl = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetSpecOssConfig(v *BatchImportHttpApisRequestSpecOssConfig) *BatchImportHttpApisRequest {
	s.SpecOssConfig = v
	return s
}

func (s *BatchImportHttpApisRequest) SetStrategy(v string) *BatchImportHttpApisRequest {
	s.Strategy = &v
	return s
}

func (s *BatchImportHttpApisRequest) SetWithGatewayExtension(v bool) *BatchImportHttpApisRequest {
	s.WithGatewayExtension = &v
	return s
}

func (s *BatchImportHttpApisRequest) Validate() error {
	if s.SpecOssConfig != nil {
		if err := s.SpecOssConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchImportHttpApisRequestSpecOssConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// my-bucket
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// imports/batch.zip
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s BatchImportHttpApisRequestSpecOssConfig) String() string {
	return dara.Prettify(s)
}

func (s BatchImportHttpApisRequestSpecOssConfig) GoString() string {
	return s.String()
}

func (s *BatchImportHttpApisRequestSpecOssConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *BatchImportHttpApisRequestSpecOssConfig) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *BatchImportHttpApisRequestSpecOssConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *BatchImportHttpApisRequestSpecOssConfig) SetBucketName(v string) *BatchImportHttpApisRequestSpecOssConfig {
	s.BucketName = &v
	return s
}

func (s *BatchImportHttpApisRequestSpecOssConfig) SetObjectKey(v string) *BatchImportHttpApisRequestSpecOssConfig {
	s.ObjectKey = &v
	return s
}

func (s *BatchImportHttpApisRequestSpecOssConfig) SetRegionId(v string) *BatchImportHttpApisRequestSpecOssConfig {
	s.RegionId = &v
	return s
}

func (s *BatchImportHttpApisRequestSpecOssConfig) Validate() error {
	return dara.Validate(s)
}
