// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CreateArtifactRequest struct {
	ArtifactId       *string                                `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *CreateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	ArtifactType     *string                                `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Description      *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string                                `json:"Name,omitempty" xml:"Name,omitempty"`
	SupportRegionIds []*string                              `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	VersionName      *string                                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequest) SetArtifactId(v string) *CreateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactRequest) SetArtifactProperty(v *CreateArtifactRequestArtifactProperty) *CreateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *CreateArtifactRequest) SetArtifactType(v string) *CreateArtifactRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactRequest) SetDescription(v string) *CreateArtifactRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactRequest) SetName(v string) *CreateArtifactRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactRequest) SetSupportRegionIds(v []*string) *CreateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactRequest) SetVersionName(v string) *CreateArtifactRequest {
	s.VersionName = &v
	return s
}

type CreateArtifactRequestArtifactProperty struct {
	CommodityCode      *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityVersion   *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	FileScriptMetadata *string `json:"FileScriptMetadata,omitempty" xml:"FileScriptMetadata,omitempty"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScriptMetadata     *string `json:"ScriptMetadata,omitempty" xml:"ScriptMetadata,omitempty"`
	Url                *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateArtifactRequestArtifactProperty) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityCode(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetFileScriptMetadata(v string) *CreateArtifactRequestArtifactProperty {
	s.FileScriptMetadata = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetImageId(v string) *CreateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRegionId(v string) *CreateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetScriptMetadata(v string) *CreateArtifactRequestArtifactProperty {
	s.ScriptMetadata = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetUrl(v string) *CreateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

type CreateArtifactShrinkRequest struct {
	ArtifactId             *string   `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactPropertyShrink *string   `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	ArtifactType           *string   `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Description            *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	SupportRegionIds       []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	VersionName            *string   `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequest) SetArtifactId(v string) *CreateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *CreateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactType(v string) *CreateArtifactShrinkRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetDescription(v string) *CreateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetName(v string) *CreateArtifactShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *CreateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetVersionName(v string) *CreateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

type CreateArtifactResponseBody struct {
	ArtifactId       *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	ArtifactType     *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	ArtifactVersion  *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MaxVersion       *int64  `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	VersionName      *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponseBody) SetArtifactId(v string) *CreateArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactProperty(v string) *CreateArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactType(v string) *CreateArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactVersion(v string) *CreateArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *CreateArtifactResponseBody) SetDescription(v string) *CreateArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *CreateArtifactResponseBody) SetGmtModified(v string) *CreateArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateArtifactResponseBody) SetMaxVersion(v int64) *CreateArtifactResponseBody {
	s.MaxVersion = &v
	return s
}

func (s *CreateArtifactResponseBody) SetName(v string) *CreateArtifactResponseBody {
	s.Name = &v
	return s
}

func (s *CreateArtifactResponseBody) SetRequestId(v string) *CreateArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactResponseBody) SetStatus(v string) *CreateArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *CreateArtifactResponseBody) SetSupportRegionIds(v string) *CreateArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *CreateArtifactResponseBody) SetVersionName(v string) *CreateArtifactResponseBody {
	s.VersionName = &v
	return s
}

type CreateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponse) SetHeaders(v map[string]*string) *CreateArtifactResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactResponse) SetStatusCode(v int32) *CreateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactResponse) SetBody(v *CreateArtifactResponseBody) *CreateArtifactResponse {
	s.Body = v
	return s
}

type DeleteArtifactRequest struct {
	ArtifactId      *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s DeleteArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactRequest) SetArtifactId(v string) *DeleteArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *DeleteArtifactRequest) SetArtifactVersion(v string) *DeleteArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

type DeleteArtifactResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponseBody) SetRequestId(v string) *DeleteArtifactResponseBody {
	s.RequestId = &v
	return s
}

type DeleteArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactResponse) SetHeaders(v map[string]*string) *DeleteArtifactResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactResponse) SetStatusCode(v int32) *DeleteArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactResponse) SetBody(v *DeleteArtifactResponseBody) *DeleteArtifactResponse {
	s.Body = v
	return s
}

type GetArtifactRequest struct {
	ArtifactId      *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
}

func (s GetArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRequest) SetArtifactId(v string) *GetArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactRequest) SetArtifactVersion(v string) *GetArtifactRequest {
	s.ArtifactVersion = &v
	return s
}

type GetArtifactResponseBody struct {
	ArtifactId       *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	ArtifactType     *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	ArtifactVersion  *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MaxVersion       *int64  `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Progress         *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	VersionName      *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBody) SetArtifactId(v string) *GetArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactProperty(v string) *GetArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactType(v string) *GetArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactVersion(v string) *GetArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *GetArtifactResponseBody) SetDescription(v string) *GetArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *GetArtifactResponseBody) SetGmtModified(v string) *GetArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetArtifactResponseBody) SetMaxVersion(v int64) *GetArtifactResponseBody {
	s.MaxVersion = &v
	return s
}

func (s *GetArtifactResponseBody) SetName(v string) *GetArtifactResponseBody {
	s.Name = &v
	return s
}

func (s *GetArtifactResponseBody) SetProgress(v string) *GetArtifactResponseBody {
	s.Progress = &v
	return s
}

func (s *GetArtifactResponseBody) SetRequestId(v string) *GetArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactResponseBody) SetStatus(v string) *GetArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *GetArtifactResponseBody) SetSupportRegionIds(v string) *GetArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *GetArtifactResponseBody) SetVersionName(v string) *GetArtifactResponseBody {
	s.VersionName = &v
	return s
}

type GetArtifactResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactResponse) SetHeaders(v map[string]*string) *GetArtifactResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactResponse) SetStatusCode(v int32) *GetArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactResponse) SetBody(v *GetArtifactResponseBody) *GetArtifactResponse {
	s.Body = v
	return s
}

type GetArtifactRepositoryCredentialsRequest struct {
	ArtifactType   *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsRequest) SetArtifactType(v string) *GetArtifactRepositoryCredentialsRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsRequest) SetDeployRegionId(v string) *GetArtifactRepositoryCredentialsRequest {
	s.DeployRegionId = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBody struct {
	AvailableResources []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources `json:"AvailableResources,omitempty" xml:"AvailableResources,omitempty" type:"Repeated"`
	Credentials        *GetArtifactRepositoryCredentialsResponseBodyCredentials          `json:"Credentials,omitempty" xml:"Credentials,omitempty" type:"Struct"`
	ExpireDate         *string                                                           `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	RequestId          *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetAvailableResources(v []*GetArtifactRepositoryCredentialsResponseBodyAvailableResources) *GetArtifactRepositoryCredentialsResponseBody {
	s.AvailableResources = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetCredentials(v *GetArtifactRepositoryCredentialsResponseBodyCredentials) *GetArtifactRepositoryCredentialsResponseBody {
	s.Credentials = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetExpireDate(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.ExpireDate = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBody) SetRequestId(v string) *GetArtifactRepositoryCredentialsResponseBody {
	s.RequestId = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBodyAvailableResources struct {
	Path           *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyAvailableResources) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetPath(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.Path = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRegionId(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RegionId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyAvailableResources) SetRepositoryName(v string) *GetArtifactRepositoryCredentialsResponseBodyAvailableResources {
	s.RepositoryName = &v
	return s
}

type GetArtifactRepositoryCredentialsResponseBodyCredentials struct {
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Password        *string `json:"Password,omitempty" xml:"Password,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Username        *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponseBodyCredentials) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeyId(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeyId = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetAccessKeySecret(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.AccessKeySecret = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetPassword(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Password = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetSecurityToken(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.SecurityToken = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponseBodyCredentials) SetUsername(v string) *GetArtifactRepositoryCredentialsResponseBodyCredentials {
	s.Username = &v
	return s
}

type GetArtifactRepositoryCredentialsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetArtifactRepositoryCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetArtifactRepositoryCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactRepositoryCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactRepositoryCredentialsResponse) SetHeaders(v map[string]*string) *GetArtifactRepositoryCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetStatusCode(v int32) *GetArtifactRepositoryCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactRepositoryCredentialsResponse) SetBody(v *GetArtifactRepositoryCredentialsResponseBody) *GetArtifactRepositoryCredentialsResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	CreateTime                *string                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableInstanceOps         *bool                                        `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	EndTime                   *string                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsOperated                *bool                                        `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	LicenseMetadata           *string                                      `json:"LicenseMetadata,omitempty" xml:"LicenseMetadata,omitempty"`
	Name                      *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkConfig             *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	OperatedServiceInstanceId *string                                      `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	OperationEndTime          *string                                      `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OperationStartTime        *string                                      `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	Outputs                   *string                                      `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	Parameters                *string                                      `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PayType                   *string                                      `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Progress                  *int64                                       `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RequestId                 *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources                 *string                                      `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service                   *GetServiceInstanceResponseBodyService       `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	ServiceInstanceId         *string                                      `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	ServiceType               *string                                      `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string                                      `json:"Source,omitempty" xml:"Source,omitempty"`
	Status                    *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDetail              *string                                      `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	SupplierUid               *int64                                       `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	Tags                      []*GetServiceInstanceResponseBodyTags        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName              *string                                      `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UpdateTime                *string                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                    *int64                                       `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseMetadata(v string) *GetServiceInstanceResponseBody {
	s.LicenseMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	EndpointId                   *string                                                                    `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointServiceId            *string                                                                    `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	PrivateVpcConnections        []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections        `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	EndpointId        *string                                                                              `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointServiceId *string                                                                              `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
	PrivateZoneName   *string                                                                              `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs struct {
	EndpointIps           []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	IngressEndpointStatus *string   `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	NetworkServiceStatus  *string   `json:"NetworkServiceStatus,omitempty" xml:"NetworkServiceStatus,omitempty"`
	SecurityGroups        []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	VSwitches             []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VpcId                 *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetEndpointIps(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.EndpointIps = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetIngressEndpointStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.IngressEndpointStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetNetworkServiceStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.NetworkServiceStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetSecurityGroups(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.SecurityGroups = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVSwitches(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VSwitches = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVpcId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VpcId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	EndpointId        *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointServiceId *string `json:"EndpointServiceId,omitempty" xml:"EndpointServiceId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointServiceId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointServiceId = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	DeployMetadata    *string                                              `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	DeployType        *string                                              `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	PublishTime       *string                                              `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	ServiceDocUrl     *string                                              `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	ServiceId         *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos      []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	ServiceProductUrl *string                                              `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	ServiceType       *string                                              `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status            *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName      *string                                              `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl       *string                                              `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	Version           *string                                              `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName       *string                                              `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceInstanceResponseBodyTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type ListArtifactVersionsRequest struct {
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	MaxResult  *string `json:"MaxResult,omitempty" xml:"MaxResult,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListArtifactVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsRequest) SetArtifactId(v string) *ListArtifactVersionsRequest {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetMaxResult(v string) *ListArtifactVersionsRequest {
	s.MaxResult = &v
	return s
}

func (s *ListArtifactVersionsRequest) SetNextToken(v string) *ListArtifactVersionsRequest {
	s.NextToken = &v
	return s
}

type ListArtifactVersionsResponseBody struct {
	Artifacts  []*ListArtifactVersionsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	MaxResults *string                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBody) SetArtifacts(v []*ListArtifactVersionsResponseBodyArtifacts) *ListArtifactVersionsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetMaxResults(v string) *ListArtifactVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetNextToken(v string) *ListArtifactVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetRequestId(v string) *ListArtifactVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactVersionsResponseBody) SetTotalCount(v string) *ListArtifactVersionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactVersionsResponseBodyArtifacts struct {
	ArtifactId          *string            `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty    *string            `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	ArtifactType        *string            `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	ArtifactVersion     *string            `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	GmtCreate           *string            `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string            `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageDelivery       map[string]*string `json:"ImageDelivery,omitempty" xml:"ImageDelivery,omitempty"`
	Progress            *string            `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ResultFile          *string            `json:"ResultFile,omitempty" xml:"ResultFile,omitempty"`
	SecurityAuditResult *string            `json:"SecurityAuditResult,omitempty" xml:"SecurityAuditResult,omitempty"`
	Status              *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportRegionIds    *string            `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	VersionName         *string            `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListArtifactVersionsResponseBodyArtifacts) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactProperty(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactProperty = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetArtifactVersion(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ArtifactVersion = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtCreate(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtCreate = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetImageDelivery(v map[string]*string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ImageDelivery = v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetProgress(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Progress = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetResultFile(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.ResultFile = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSecurityAuditResult(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SecurityAuditResult = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetStatus(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetSupportRegionIds(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.SupportRegionIds = &v
	return s
}

func (s *ListArtifactVersionsResponseBodyArtifacts) SetVersionName(v string) *ListArtifactVersionsResponseBodyArtifacts {
	s.VersionName = &v
	return s
}

type ListArtifactVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListArtifactVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListArtifactVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactVersionsResponse) SetHeaders(v map[string]*string) *ListArtifactVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactVersionsResponse) SetStatusCode(v int32) *ListArtifactVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactVersionsResponse) SetBody(v *ListArtifactVersionsResponseBody) *ListArtifactVersionsResponse {
	s.Body = v
	return s
}

type ListArtifactsRequest struct {
	Filter     []*ListArtifactsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	MaxResults *string                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListArtifactsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequest) SetFilter(v []*ListArtifactsRequestFilter) *ListArtifactsRequest {
	s.Filter = v
	return s
}

func (s *ListArtifactsRequest) SetMaxResults(v string) *ListArtifactsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsRequest) SetNextToken(v string) *ListArtifactsRequest {
	s.NextToken = &v
	return s
}

type ListArtifactsRequestFilter struct {
	Name   *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListArtifactsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListArtifactsRequestFilter) SetName(v string) *ListArtifactsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListArtifactsRequestFilter) SetValues(v []*string) *ListArtifactsRequestFilter {
	s.Values = v
	return s
}

type ListArtifactsResponseBody struct {
	Artifacts  []*ListArtifactsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	MaxResults *string                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBody) SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactsResponseBody) SetMaxResults(v string) *ListArtifactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsResponseBody) SetNextToken(v string) *ListArtifactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsResponseBody) SetRequestId(v string) *ListArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactsResponseBody) SetTotalCount(v string) *ListArtifactsResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactsResponseBodyArtifacts struct {
	ArtifactId   *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	MaxVersion   *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListArtifactsResponseBodyArtifacts) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactId(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactId = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactType(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetDescription(v string) *ListArtifactsResponseBodyArtifacts {
	s.Description = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetGmtModified(v string) *ListArtifactsResponseBodyArtifacts {
	s.GmtModified = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetMaxVersion(v string) *ListArtifactsResponseBodyArtifacts {
	s.MaxVersion = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetName(v string) *ListArtifactsResponseBodyArtifacts {
	s.Name = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetStatus(v string) *ListArtifactsResponseBodyArtifacts {
	s.Status = &v
	return s
}

type ListArtifactsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListArtifactsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListArtifactsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactsResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponse) SetHeaders(v map[string]*string) *ListArtifactsResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactsResponse) SetStatusCode(v int32) *ListArtifactsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactsResponse) SetBody(v *ListArtifactsResponseBody) *ListArtifactsResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	Filter     []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	MaxResults *string                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tag        []*ListServiceInstancesRequestTag    `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v string) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	MaxResults       *string                                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	TotalCount       *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v string) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	CreateTime                *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableInstanceOps         *bool                                                    `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	EndTime                   *string                                                  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsOperated                *bool                                                    `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	Name                      *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OperatedServiceInstanceId *string                                                  `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	OperationEndTime          *string                                                  `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	OperationStartTime        *string                                                  `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	Parameters                *string                                                  `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PayType                   *string                                                  `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Progress                  *int64                                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Service                   *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	ServiceInstanceId         *string                                                  `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	ServiceType               *string                                                  `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Source                    *string                                                  `json:"Source,omitempty" xml:"Source,omitempty"`
	Status                    *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDetail              *string                                                  `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Tags                      []*ListServiceInstancesResponseBodyServiceInstancesTags  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName              *string                                                  `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	UpdateTime                *string                                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId                    *int64                                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetIsOperated(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.IsOperated = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUserId(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.UserId = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	DeployMetadata *string                                                                `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	DeployType     *string                                                                `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	PublishTime    *string                                                                `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	ServiceId      *string                                                                `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos   []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	ServiceType    *string                                                                `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	Status         *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName   *string                                                                `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	SupplierUrl    *string                                                                `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	Version        *string                                                                `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName    *string                                                                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployMetadata(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployMetadata = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	Image            *string `json:"Image,omitempty" xml:"Image,omitempty"`
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListServiceUsagesRequest struct {
	Filter     []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	MaxResults *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListServiceUsagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequest) SetFilter(v []*ListServiceUsagesRequestFilter) *ListServiceUsagesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceUsagesRequest) SetMaxResults(v int32) *ListServiceUsagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesRequest) SetNextToken(v string) *ListServiceUsagesRequest {
	s.NextToken = &v
	return s
}

type ListServiceUsagesRequestFilter struct {
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceUsagesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequestFilter) SetName(v string) *ListServiceUsagesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceUsagesRequestFilter) SetValue(v []*string) *ListServiceUsagesRequestFilter {
	s.Value = v
	return s
}

type ListServiceUsagesResponseBody struct {
	MaxResults    *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken     *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceUsages []*ListServiceUsagesResponseBodyServiceUsages `json:"ServiceUsages,omitempty" xml:"ServiceUsages,omitempty" type:"Repeated"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceUsagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBody) SetMaxResults(v int32) *ListServiceUsagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetNextToken(v string) *ListServiceUsagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetRequestId(v string) *ListServiceUsagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceUsagesResponseBody) SetServiceUsages(v []*ListServiceUsagesResponseBodyServiceUsages) *ListServiceUsagesResponseBody {
	s.ServiceUsages = v
	return s
}

func (s *ListServiceUsagesResponseBody) SetTotalCount(v int32) *ListServiceUsagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceUsagesResponseBodyServiceUsages struct {
	Comments        *string                                                    `json:"Comments,omitempty" xml:"Comments,omitempty"`
	CreateTime      *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceId       *string                                                    `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName     *string                                                    `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status          *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName    *string                                                    `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	UpdateTime      *string                                                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserAliUid      *int64                                                     `json:"UserAliUid,omitempty" xml:"UserAliUid,omitempty"`
	UserInformation *ListServiceUsagesResponseBodyServiceUsagesUserInformation `json:"UserInformation,omitempty" xml:"UserInformation,omitempty" type:"Struct"`
}

func (s ListServiceUsagesResponseBodyServiceUsages) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBodyServiceUsages) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetComments(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Comments = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetCreateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.CreateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceId(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceId = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetServiceName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.ServiceName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetStatus(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.Status = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetSupplierName(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.SupplierName = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUpdateTime(v string) *ListServiceUsagesResponseBodyServiceUsages {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserAliUid(v int64) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserAliUid = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsages) SetUserInformation(v *ListServiceUsagesResponseBodyServiceUsagesUserInformation) *ListServiceUsagesResponseBodyServiceUsages {
	s.UserInformation = v
	return s
}

type ListServiceUsagesResponseBodyServiceUsagesUserInformation struct {
	Company      *string `json:"Company,omitempty" xml:"Company,omitempty"`
	Country      *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EmailAddress *string `json:"EmailAddress,omitempty" xml:"EmailAddress,omitempty"`
	Industry     *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Source       *string `json:"Source,omitempty" xml:"Source,omitempty"`
	Telephone    *string `json:"Telephone,omitempty" xml:"Telephone,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListServiceUsagesResponseBodyServiceUsagesUserInformation) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponseBodyServiceUsagesUserInformation) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetCompany(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Company = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetCountry(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Country = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetEmailAddress(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.EmailAddress = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetIndustry(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Industry = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetName(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Name = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetSource(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Source = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetTelephone(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Telephone = &v
	return s
}

func (s *ListServiceUsagesResponseBodyServiceUsagesUserInformation) SetTitle(v string) *ListServiceUsagesResponseBodyServiceUsagesUserInformation {
	s.Title = &v
	return s
}

type ListServiceUsagesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceUsagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceUsagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceUsagesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesResponse) SetHeaders(v map[string]*string) *ListServiceUsagesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceUsagesResponse) SetStatusCode(v int32) *ListServiceUsagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceUsagesResponse) SetBody(v *ListServiceUsagesResponseBody) *ListServiceUsagesResponse {
	s.Body = v
	return s
}

type ReleaseArtifactRequest struct {
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
}

func (s ReleaseArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactRequest) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactRequest) SetArtifactId(v string) *ReleaseArtifactRequest {
	s.ArtifactId = &v
	return s
}

type ReleaseArtifactResponseBody struct {
	ArtifactId       *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	ArtifactType     *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	ArtifactVersion  *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VersionName      *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ReleaseArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponseBody) SetArtifactId(v string) *ReleaseArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactProperty(v string) *ReleaseArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactType(v string) *ReleaseArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetArtifactVersion(v string) *ReleaseArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetDescription(v string) *ReleaseArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetGmtModified(v string) *ReleaseArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetRequestId(v string) *ReleaseArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetStatus(v string) *ReleaseArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *ReleaseArtifactResponseBody) SetVersionName(v string) *ReleaseArtifactResponseBody {
	s.VersionName = &v
	return s
}

type ReleaseArtifactResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseArtifactResponse) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponse) SetHeaders(v map[string]*string) *ReleaseArtifactResponse {
	s.Headers = v
	return s
}

func (s *ReleaseArtifactResponse) SetStatusCode(v int32) *ReleaseArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseArtifactResponse) SetBody(v *ReleaseArtifactResponseBody) *ReleaseArtifactResponse {
	s.Body = v
	return s
}

type UpdateArtifactRequest struct {
	ArtifactId       *string                                `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *UpdateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	Description      *string                                `json:"Description,omitempty" xml:"Description,omitempty"`
	SupportRegionIds []*string                              `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	VersionName      *string                                `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequest) SetArtifactId(v string) *UpdateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactRequest) SetArtifactProperty(v *UpdateArtifactRequestArtifactProperty) *UpdateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *UpdateArtifactRequest) SetDescription(v string) *UpdateArtifactRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactRequest) SetSupportRegionIds(v []*string) *UpdateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactRequest) SetVersionName(v string) *UpdateArtifactRequest {
	s.VersionName = &v
	return s
}

type UpdateArtifactRequestArtifactProperty struct {
	CommodityCode      *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityVersion   *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	FileScriptMetadata *string `json:"FileScriptMetadata,omitempty" xml:"FileScriptMetadata,omitempty"`
	ImageId            *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ScriptMetadata     *string `json:"ScriptMetadata,omitempty" xml:"ScriptMetadata,omitempty"`
	Url                *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateArtifactRequestArtifactProperty) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityCode(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetFileScriptMetadata(v string) *UpdateArtifactRequestArtifactProperty {
	s.FileScriptMetadata = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetImageId(v string) *UpdateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRegionId(v string) *UpdateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetScriptMetadata(v string) *UpdateArtifactRequestArtifactProperty {
	s.ScriptMetadata = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetUrl(v string) *UpdateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

type UpdateArtifactShrinkRequest struct {
	ArtifactId             *string   `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactPropertyShrink *string   `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	Description            *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	SupportRegionIds       []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	VersionName            *string   `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactShrinkRequest) SetArtifactId(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetDescription(v string) *UpdateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *UpdateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetVersionName(v string) *UpdateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

type UpdateArtifactResponseBody struct {
	ArtifactId       *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	ArtifactType     *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	ArtifactVersion  *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	VersionName      *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponseBody) SetArtifactId(v string) *UpdateArtifactResponseBody {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactProperty(v string) *UpdateArtifactResponseBody {
	s.ArtifactProperty = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactType(v string) *UpdateArtifactResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactVersion(v string) *UpdateArtifactResponseBody {
	s.ArtifactVersion = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetDescription(v string) *UpdateArtifactResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetGmtModified(v string) *UpdateArtifactResponseBody {
	s.GmtModified = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetRequestId(v string) *UpdateArtifactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetStatus(v string) *UpdateArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetSupportRegionIds(v string) *UpdateArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetVersionName(v string) *UpdateArtifactResponseBody {
	s.VersionName = &v
	return s
}

type UpdateArtifactResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateArtifactResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateArtifactResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponse) SetHeaders(v map[string]*string) *UpdateArtifactResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactResponse) SetStatusCode(v int32) *UpdateArtifactResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactResponse) SetBody(v *UpdateArtifactResponseBody) *UpdateArtifactResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenestsupplier"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateArtifactWithOptions(tmpReq *CreateArtifactRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactProperty)) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, tea.String("ArtifactProperty"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactPropertyShrink)) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRegionIds)) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateArtifact(request *CreateArtifactRequest) (_result *CreateArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArtifactResponse{}
	_body, _err := client.CreateArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteArtifactWithOptions(request *DeleteArtifactRequest, runtime *util.RuntimeOptions) (_result *DeleteArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteArtifact(request *DeleteArtifactRequest) (_result *DeleteArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteArtifactResponse{}
	_body, _err := client.DeleteArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetArtifactWithOptions(request *GetArtifactRequest, runtime *util.RuntimeOptions) (_result *GetArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactVersion)) {
		query["ArtifactVersion"] = request.ArtifactVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetArtifact(request *GetArtifactRequest) (_result *GetArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactResponse{}
	_body, _err := client.GetArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetArtifactRepositoryCredentialsWithOptions(request *GetArtifactRepositoryCredentialsRequest, runtime *util.RuntimeOptions) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactRepositoryCredentials"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetArtifactRepositoryCredentials(request *GetArtifactRepositoryCredentialsRequest) (_result *GetArtifactRepositoryCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactRepositoryCredentialsResponse{}
	_body, _err := client.GetArtifactRepositoryCredentialsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstance"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceInstance(request *GetServiceInstanceRequest) (_result *GetServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.GetServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListArtifactVersionsWithOptions(request *ListArtifactVersionsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResult)) {
		query["MaxResult"] = request.MaxResult
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactVersions"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListArtifactVersions(request *ListArtifactVersionsRequest) (_result *ListArtifactVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactVersionsResponse{}
	_body, _err := client.ListArtifactVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListArtifactsWithOptions(request *ListArtifactsRequest, runtime *util.RuntimeOptions) (_result *ListArtifactsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifacts"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListArtifacts(request *ListArtifactsRequest) (_result *ListArtifactsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactsResponse{}
	_body, _err := client.ListArtifactsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceInstances(request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceUsagesWithOptions(request *ListServiceUsagesRequest, runtime *util.RuntimeOptions) (_result *ListServiceUsagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceUsages"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceUsages(request *ListServiceUsagesRequest) (_result *ListServiceUsagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceUsagesResponse{}
	_body, _err := client.ListServiceUsagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseArtifactWithOptions(request *ReleaseArtifactRequest, runtime *util.RuntimeOptions) (_result *ReleaseArtifactResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseArtifact(request *ReleaseArtifactRequest) (_result *ReleaseArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseArtifactResponse{}
	_body, _err := client.ReleaseArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateArtifactWithOptions(tmpReq *UpdateArtifactRequest, runtime *util.RuntimeOptions) (_result *UpdateArtifactResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateArtifactShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ArtifactProperty)) {
		request.ArtifactPropertyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ArtifactProperty, tea.String("ArtifactProperty"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactId)) {
		query["ArtifactId"] = request.ArtifactId
	}

	if !tea.BoolValue(util.IsUnset(request.ArtifactPropertyShrink)) {
		query["ArtifactProperty"] = request.ArtifactPropertyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRegionIds)) {
		query["SupportRegionIds"] = request.SupportRegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateArtifact"),
		Version:     tea.String("2021-05-21"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateArtifact(request *UpdateArtifactRequest) (_result *UpdateArtifactResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateArtifactResponse{}
	_body, _err := client.UpdateArtifactWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
