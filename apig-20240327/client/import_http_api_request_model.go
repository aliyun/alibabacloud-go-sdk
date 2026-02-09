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
	// The API deployment configuration.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The imported API description (255-byte limit). If not specified, a description is extracted from the API definition file. A maximum of 255 bytes is supported.
	//
	// example:
	//
	// API for testing
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to perform a precheck. If set to true, a check is performed without actual import.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// Gateway ID.
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
	// The imported API name. If not specified, a name is extracted from the API definition file. If the API name and versioning configuration already exist, this import will update the existing API definition based on the strategy field.
	//
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The [resource group ID](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfm3q4zjh7fkki
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Base64-encoded API definition (supports OAS 2.0/OAS 3.0 in YAML/JSON). This parameter has higher priority than the specFileUrl parameter. However, if the file size exceeds 10 MB, use the specFileUrl parameter to pass the definition.
	//
	// example:
	//
	// b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
	SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
	// The download URL of the API definition file. Must be either a publicly accessible Object Storage Service (OSS) URL or an OSS intranet endpoint within the same region. Requires download permissions. For OSS URLs that are not publicly readable, refer to [https://www.alibabacloud.com/help/en/oss/user-guide/how-to-obtain-the-url-of-a-single-object-or-the-urls-of-multiple-objects](https://help.aliyun.com/document_detail/39607.html) and use URLs with download permissions. Currently, only OSS URLs are supported.
	//
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/my-api/api.yaml
	SpecFileUrl *string `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	// The OSS configuration details.
	SpecOssConfig *ImportHttpApiRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// The conflict resolution strategy when the API to be imported has the same name and version as an existing one. Valid values:
	//
	// 	- SpecOnly: full override.
	//
	// 	- SpecFirst: Merge with priority on the newly imported file. New APIs are created and existing ones are updated. APIs not included in the file remain unchanged.
	//
	// 	- ExistFirst (default): Merge with priority on existing APIs. New APIs are created but existing ones remain unchanged. If this parameter is not specified, the ExistFirst policy takes effect.
	//
	// example:
	//
	// ExistFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// The target REST API ID for direct updates. If specified, the import operation will directly update the designated API instead of creating new APIs or updating existing APIs based on the name and version. Only REST APIs can be specified.
	//
	// example:
	//
	// api-xxxx
	TargetHttpApiId *string `json:"targetHttpApiId,omitempty" xml:"targetHttpApiId,omitempty"`
	// The API versioning configuration. If versioning is enabled, an imported API that matches both the version number and the API name of an existing API will update that API. If versioning is disabled, an imported API that matches the API name of an existing API will update it.
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
	// The OSS bucket name.
	//
	// example:
	//
	// api-1
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The full file path in OSS.
	//
	// example:
	//
	// /test/swagger.json
	ObjectKey *string `json:"objectKey,omitempty" xml:"objectKey,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
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
