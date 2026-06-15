// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHttpApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeployConfigs(v []*HttpApiDeployConfig) *ImportHttpApiRequest
	GetDeployConfigs() []*HttpApiDeployConfig
	SetDescription(v string) *ImportHttpApiRequest
	GetDescription() *string
	SetDryRun(v bool) *ImportHttpApiRequest
	GetDryRun() *bool
	SetGatewayId(v string) *ImportHttpApiRequest
	GetGatewayId() *string
	SetMcpRouteId(v string) *ImportHttpApiRequest
	GetMcpRouteId() *string
	SetName(v string) *ImportHttpApiRequest
	GetName() *string
	SetResourceGroupId(v string) *ImportHttpApiRequest
	GetResourceGroupId() *string
	SetSpecContentBase64(v string) *ImportHttpApiRequest
	GetSpecContentBase64() *string
	SetSpecFileUrl(v string) *ImportHttpApiRequest
	GetSpecFileUrl() *string
	SetSpecOssConfig(v *ImportHttpApiRequestSpecOssConfig) *ImportHttpApiRequest
	GetSpecOssConfig() *ImportHttpApiRequestSpecOssConfig
	SetStrategy(v string) *ImportHttpApiRequest
	GetStrategy() *string
	SetTargetHttpApiId(v string) *ImportHttpApiRequest
	GetTargetHttpApiId() *string
	SetVersionConfig(v *HttpApiVersionConfig) *ImportHttpApiRequest
	GetVersionConfig() *HttpApiVersionConfig
	SetWithGatewayExtension(v bool) *ImportHttpApiRequest
	GetWithGatewayExtension() *bool
}

type ImportHttpApiRequest struct {
	// The API deployment configurations.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The description of the imported API. If this parameter is not specified, the description is extracted from the API definition. Maximum length: 255 bytes.
	//
	// example:
	//
	// 测试专用API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to perform a dry run. If this parameter is enabled, only validation is performed without the actual import.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The MCP route ID.
	//
	// example:
	//
	// xxx
	McpRouteId *string `json:"mcpRouteId,omitempty" xml:"mcpRouteId,omitempty"`
	// The name of the imported API. If this parameter is not specified, the name is extracted from the API definition file. If an API with the same name and version configuration already exists, this import updates the existing API definition based on the strategy parameter.
	//
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The [resource group ID](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-aek23nsa353vmra
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Base64-encoded API definition. OAS 2.0 and OAS 3.0 specifications are supported in YAML or JSON format. This parameter takes priority over the specFileUrl parameter. If the file size exceeds 10 MB, use the specFileUrl parameter instead.
	//
	// example:
	//
	// b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
	SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
	// The download URL of the API definition file. The URL must be accessible over the public network or be an internal network OSS download URL in the same region. The URL must have download permissions. For OSS files that are not publicly readable, refer to References [Download objects using presigned URLs](https://help.aliyun.com/document_detail/39607.html) and provide a URL with download permissions. Only API definition files stored in OSS are supported.
	SpecFileUrl *string `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	// The OSS configuration.
	SpecOssConfig *ImportHttpApiRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// The update strategy to use when the imported API has the same name and version management configuration as an existing API. Valid values:
	//
	// - SpecOnly: uses the imported file as the single source of truth.
	//
	// - SpecFirst: prioritizes the imported file. New operations are added and existing operations are updated. Operations not mentioned in the file remain unchanged.
	//
	// - ExistFirst: prioritizes the existing API. Only new operations are added. Existing operations are not updated.
	//
	// Default value: ExistFirst.
	//
	// example:
	//
	// ExistFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// The ID of the target HTTP API. If this parameter is specified, this import updates the specified API instead of creating a new one or searching for an existing API by name and version management configuration. The target API must be of the REST type.
	//
	// example:
	//
	// api-xxxx
	TargetHttpApiId *string `json:"targetHttpApiId,omitempty" xml:"targetHttpApiId,omitempty"`
	// The API version configuration. If version configuration is enabled and an API with the same version number and name already exists, this import is treated as an update. If version configuration is not enabled and an API with the same name already exists, this import is treated as an update.
	VersionConfig        *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
	WithGatewayExtension *bool                 `json:"withGatewayExtension,omitempty" xml:"withGatewayExtension,omitempty"`
}

func (s ImportHttpApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiRequest) GoString() string {
	return s.String()
}

func (s *ImportHttpApiRequest) GetDeployConfigs() []*HttpApiDeployConfig {
	return s.DeployConfigs
}

func (s *ImportHttpApiRequest) GetDescription() *string {
	return s.Description
}

func (s *ImportHttpApiRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ImportHttpApiRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ImportHttpApiRequest) GetMcpRouteId() *string {
	return s.McpRouteId
}

func (s *ImportHttpApiRequest) GetName() *string {
	return s.Name
}

func (s *ImportHttpApiRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ImportHttpApiRequest) GetSpecContentBase64() *string {
	return s.SpecContentBase64
}

func (s *ImportHttpApiRequest) GetSpecFileUrl() *string {
	return s.SpecFileUrl
}

func (s *ImportHttpApiRequest) GetSpecOssConfig() *ImportHttpApiRequestSpecOssConfig {
	return s.SpecOssConfig
}

func (s *ImportHttpApiRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *ImportHttpApiRequest) GetTargetHttpApiId() *string {
	return s.TargetHttpApiId
}

func (s *ImportHttpApiRequest) GetVersionConfig() *HttpApiVersionConfig {
	return s.VersionConfig
}

func (s *ImportHttpApiRequest) GetWithGatewayExtension() *bool {
	return s.WithGatewayExtension
}

func (s *ImportHttpApiRequest) SetDeployConfigs(v []*HttpApiDeployConfig) *ImportHttpApiRequest {
	s.DeployConfigs = v
	return s
}

func (s *ImportHttpApiRequest) SetDescription(v string) *ImportHttpApiRequest {
	s.Description = &v
	return s
}

func (s *ImportHttpApiRequest) SetDryRun(v bool) *ImportHttpApiRequest {
	s.DryRun = &v
	return s
}

func (s *ImportHttpApiRequest) SetGatewayId(v string) *ImportHttpApiRequest {
	s.GatewayId = &v
	return s
}

func (s *ImportHttpApiRequest) SetMcpRouteId(v string) *ImportHttpApiRequest {
	s.McpRouteId = &v
	return s
}

func (s *ImportHttpApiRequest) SetName(v string) *ImportHttpApiRequest {
	s.Name = &v
	return s
}

func (s *ImportHttpApiRequest) SetResourceGroupId(v string) *ImportHttpApiRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ImportHttpApiRequest) SetSpecContentBase64(v string) *ImportHttpApiRequest {
	s.SpecContentBase64 = &v
	return s
}

func (s *ImportHttpApiRequest) SetSpecFileUrl(v string) *ImportHttpApiRequest {
	s.SpecFileUrl = &v
	return s
}

func (s *ImportHttpApiRequest) SetSpecOssConfig(v *ImportHttpApiRequestSpecOssConfig) *ImportHttpApiRequest {
	s.SpecOssConfig = v
	return s
}

func (s *ImportHttpApiRequest) SetStrategy(v string) *ImportHttpApiRequest {
	s.Strategy = &v
	return s
}

func (s *ImportHttpApiRequest) SetTargetHttpApiId(v string) *ImportHttpApiRequest {
	s.TargetHttpApiId = &v
	return s
}

func (s *ImportHttpApiRequest) SetVersionConfig(v *HttpApiVersionConfig) *ImportHttpApiRequest {
	s.VersionConfig = v
	return s
}

func (s *ImportHttpApiRequest) SetWithGatewayExtension(v bool) *ImportHttpApiRequest {
	s.WithGatewayExtension = &v
	return s
}

func (s *ImportHttpApiRequest) Validate() error {
	if s.DeployConfigs != nil {
		for _, item := range s.DeployConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SpecOssConfig != nil {
		if err := s.SpecOssConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VersionConfig != nil {
		if err := s.VersionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportHttpApiRequestSpecOssConfig struct {
	// The bucket name.
	//
	// example:
	//
	// gms-service-prod
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The full path of the file.
	//
	// example:
	//
	// /test/swagger.json
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ImportHttpApiRequestSpecOssConfig) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiRequestSpecOssConfig) GoString() string {
	return s.String()
}

func (s *ImportHttpApiRequestSpecOssConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *ImportHttpApiRequestSpecOssConfig) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *ImportHttpApiRequestSpecOssConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportHttpApiRequestSpecOssConfig) SetBucketName(v string) *ImportHttpApiRequestSpecOssConfig {
	s.BucketName = &v
	return s
}

func (s *ImportHttpApiRequestSpecOssConfig) SetObjectKey(v string) *ImportHttpApiRequestSpecOssConfig {
	s.ObjectKey = &v
	return s
}

func (s *ImportHttpApiRequestSpecOssConfig) SetRegionId(v string) *ImportHttpApiRequestSpecOssConfig {
	s.RegionId = &v
	return s
}

func (s *ImportHttpApiRequestSpecOssConfig) Validate() error {
	return dara.Validate(s)
}
