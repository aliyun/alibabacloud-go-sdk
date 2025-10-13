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
	// The deployment configuration.
	DeployConfigs []*HttpApiDeployConfig `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	// The API description, which cannot exceed 255 bytes in length. If you do not specify a description, a description is extracted from the definition file.
	//
	// example:
	//
	// API for testing
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to perform a dry run. If this parameter is set to true, a dry run is performed without importing the file.
	//
	// example:
	//
	// false
	DryRun    *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The MCP route ID.
	McpRouteId *string `json:"mcpRouteId,omitempty" xml:"mcpRouteId,omitempty"`
	// The API name. If you do not specify a name, a name is extracted from the definition file. If a name and a versioning configuration already exist, the existing API definition is updated based on the strategy field.
	//
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// [The resource group ID](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfm3q4zjh7fkki
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The Base64-encoded API definition. OAS 2.0 and OAS 3.0 specifications are supported. YAML and JSON formats are supported. This parameter precedes over the specFileUrl parameter. However, if the file size exceeds 10 MB, use the specFileUrl parameter to pass the definition.
	//
	// example:
	//
	// b3BlbmFwaTogMy4wLjAKaW5mbzoKICAgIHRpdGxlOiBkZW1vCiAgICBkZXNjcmlwdGlvbjogdGhpc2lzZGVtbwogICAgdmVyc2lvbjogIiIKcGF0aHM6CiAgICAvdXNlci97dXNlcklkfToKICAgICAgICBnZXQ6CiAgICAgICAgICAgIHN1bW1hcnk6IOiOt+WPlueUqOaIt+S/oeaBrwogICAgICAgICAgICBkZXNjcmlwdGlvbjog6I635Y+W55So5oi35L+h5oGvCiAgICAgICAgICAgIG9wZXJhdGlvbklkOiBHZXRVc2VySW5mbwogICAgICAgICAgICByZXNwb25zZXM6CiAgICAgICAgICAgICAgICAiMjAwIjoKICAgICAgICAgICAgICAgICAgICBkZXNjcmlwdGlvbjog5oiQ5YqfCiAgICAgICAgICAgICAgICAgICAgY29udGVudDoKICAgICAgICAgICAgICAgICAgICAgICAgYXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04OgogICAgICAgICAgICAgICAgICAgICAgICAgICAgc2NoZW1hOiBudWxsCnNlcnZlcnM6CiAgICAtIHVybDogaHR0cDovL2FwaS5leGFtcGxlLmNvbS92MQo=
	SpecContentBase64 *string `json:"specContentBase64,omitempty" xml:"specContentBase64,omitempty"`
	// The download URL of the API definition file. You can download the file over the Internet or by using an Object Storage Service (OSS) internal download URL that belongs to the current region. You must obtain the required permissions to download the file. For OSS URLs that are not publicly readable, refer to [Download objects using presigned URLs](https://help.aliyun.com/document_detail/39607.html) to specify URLs that provide download permissions. Currently, only OSS URLs are supported.
	//
	// example:
	//
	// https://my-bucket.oss-cn-hangzhou.aliyuncs.com/my-api/api.yaml
	SpecFileUrl *string `json:"specFileUrl,omitempty" xml:"specFileUrl,omitempty"`
	// The OSS information.
	SpecOssConfig *ImportHttpApiRequestSpecOssConfig `json:"specOssConfig,omitempty" xml:"specOssConfig,omitempty" type:"Struct"`
	// The update policy when the API to be imported has the same version and name as an existing one. Valid values:
	//
	// 	- SpectOnly: All configurations in the file take effect.
	//
	// 	- SpecFirst: The file takes precedence. New APIs are created and existing ones are updated. APIs not included in the file remain unchanged.
	//
	// 	- ExistFirst (default): The existing APIs take precedence. New APIs are created but existing ones remain unchanged. If this parameter is not specified, the ExistFirst policy takes effect.
	//
	// example:
	//
	// ExistFirst
	Strategy *string `json:"strategy,omitempty" xml:"strategy,omitempty"`
	// The API to be updated. If this parameter is specified, this import updates only the specified API. New APIs are not created and unspecified existing APIs are not updated. Only REST APIs can be specified.
	//
	// example:
	//
	// api-xxxx
	TargetHttpApiId *string `json:"targetHttpApiId,omitempty" xml:"targetHttpApiId,omitempty"`
	// The API versioning configuration. If versioning is enabled for an API and the version and name of an API to be imported are the same as those of the existing API, the existing API is updated by this import. If versioning is not enabled for an API and the name of an API to be imported are the same as that of the existing API, the existing API is updated by this import.
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
	return dara.Validate(s)
}

type ImportHttpApiRequestSpecOssConfig struct {
	// The bucket name.
	//
	// example:
	//
	// api-1
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
