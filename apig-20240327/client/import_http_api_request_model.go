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
}

type ImportHttpApiRequest struct {
	// The API deployment configurations.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The description of the API to import. If omitted, the description is taken from the API definition. The maximum length is 255 bytes.
	//
	// example:
	//
	// 测试专用API
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Indicates whether to perform a dry run. If `true`, the system validates the request but does not import the API.
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
	// The name of the API to import. If omitted, the name is taken from the API definition file. If an API with the same name and versioning configuration already exists, the import acts as an update based on the specified `strategy`.
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
	// The Base64-encoded API definition. It supports OpenAPI Specification (OAS) 2.0 and 3.0 and can be in either YAML or JSON format. This parameter takes precedence over `specFileUrl`. If the file size is larger than 10 MB, use the `specFileUrl` parameter.
	//
	// example:
	//
	// b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
	SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
	// The URL of the API definition file stored in OSS. The URL must be accessible from the public network or be an internal OSS endpoint in the same region. For OSS objects that are not publicly readable, use a presigned URL. For details, see [Download a file by using a presigned URL](https://help.aliyun.com/document_detail/39607.html).
	SpecFileUrl *string `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	// Configuration for fetching the API definition from an OSS bucket.
	SpecOssConfig *ImportHttpApiRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// The update strategy to apply when an API with the same name and versioning configuration already exists.
	//
	// - `SpecOnly`: Overwrites the existing API completely with the imported definition.
	//
	// - `SpecFirst`: Updates existing APIs and creates new ones based on the imported definition. Existing APIs not included in the import file are unaffected.
	//
	// - `ExistFirst`: Creates new APIs from the imported definition but does not modify any existing APIs. This is the default strategy.
	//
	// example:
	//
	// ExistFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// If you specify this parameter, the import updates the specified API instead of creating a new one or searching for an existing API by name and versioning configuration. The target API must be an HTTP API.
	//
	// example:
	//
	// api-xxxx
	TargetHttpApiId *string `json:"targetHttpApiId,omitempty" xml:"targetHttpApiId,omitempty"`
	// The versioning configuration for the API. If an existing API matches the specified name (and version, if enabled), this import updates that API.
	VersionConfig *HttpApiVersionConfig `json:"versionConfig,omitempty" xml:"versionConfig,omitempty"`
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
	// The name of the OSS bucket that contains the API definition file.
	//
	// example:
	//
	// gms-service-prod
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The object key (file path) of the API definition file within the bucket.
	//
	// example:
	//
	// /test/swagger.json
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// The ID of the region where the OSS bucket is located.
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
