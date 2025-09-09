// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildProperty(v string) *GetArtifactResponseBody
	GetArtifactBuildProperty() *string
	SetArtifactBuildType(v string) *GetArtifactResponseBody
	GetArtifactBuildType() *string
	SetArtifactId(v string) *GetArtifactResponseBody
	GetArtifactId() *string
	SetArtifactProperty(v string) *GetArtifactResponseBody
	GetArtifactProperty() *string
	SetArtifactType(v string) *GetArtifactResponseBody
	GetArtifactType() *string
	SetArtifactVersion(v string) *GetArtifactResponseBody
	GetArtifactVersion() *string
	SetDescription(v string) *GetArtifactResponseBody
	GetDescription() *string
	SetGmtModified(v string) *GetArtifactResponseBody
	GetGmtModified() *string
	SetMaxVersion(v int64) *GetArtifactResponseBody
	GetMaxVersion() *int64
	SetName(v string) *GetArtifactResponseBody
	GetName() *string
	SetPermissionType(v string) *GetArtifactResponseBody
	GetPermissionType() *string
	SetProgress(v string) *GetArtifactResponseBody
	GetProgress() *string
	SetRequestId(v string) *GetArtifactResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetArtifactResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetArtifactResponseBody
	GetStatus() *string
	SetStatusDetail(v string) *GetArtifactResponseBody
	GetStatusDetail() *string
	SetSupportRegionIds(v string) *GetArtifactResponseBody
	GetSupportRegionIds() *string
	SetTags(v []*GetArtifactResponseBodyTags) *GetArtifactResponseBody
	GetTags() []*GetArtifactResponseBodyTags
	SetVersionName(v string) *GetArtifactResponseBody
	GetVersionName() *string
}

type GetArtifactResponseBody struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the deployment package to be built.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005xxxx\\",\\"CommodityVersion\\":\\"V2022xxxx\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the deployment package was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The latest version of the deployment package.
	//
	// example:
	//
	// 2
	MaxVersion *int64 `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The name of the deployment package.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact They can only change from Automatic to Public. Options:
	//
	// - Public
	//
	// - Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The distribution progress of the deployment package.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
	//
	// example:
	//
	// ["cn-hangzhou","cn-beijing"]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The tags of the deployment package.
	Tags []*GetArtifactResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBody) GetArtifactBuildProperty() *string {
	return s.ArtifactBuildProperty
}

func (s *GetArtifactResponseBody) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *GetArtifactResponseBody) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *GetArtifactResponseBody) GetArtifactProperty() *string {
	return s.ArtifactProperty
}

func (s *GetArtifactResponseBody) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetArtifactResponseBody) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *GetArtifactResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetArtifactResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetArtifactResponseBody) GetMaxVersion() *int64 {
	return s.MaxVersion
}

func (s *GetArtifactResponseBody) GetName() *string {
	return s.Name
}

func (s *GetArtifactResponseBody) GetPermissionType() *string {
	return s.PermissionType
}

func (s *GetArtifactResponseBody) GetProgress() *string {
	return s.Progress
}

func (s *GetArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetArtifactResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetArtifactResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetArtifactResponseBody) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *GetArtifactResponseBody) GetSupportRegionIds() *string {
	return s.SupportRegionIds
}

func (s *GetArtifactResponseBody) GetTags() []*GetArtifactResponseBodyTags {
	return s.Tags
}

func (s *GetArtifactResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *GetArtifactResponseBody) SetArtifactBuildProperty(v string) *GetArtifactResponseBody {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *GetArtifactResponseBody) SetArtifactBuildType(v string) *GetArtifactResponseBody {
	s.ArtifactBuildType = &v
	return s
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

func (s *GetArtifactResponseBody) SetPermissionType(v string) *GetArtifactResponseBody {
	s.PermissionType = &v
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

func (s *GetArtifactResponseBody) SetResourceGroupId(v string) *GetArtifactResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetArtifactResponseBody) SetStatus(v string) *GetArtifactResponseBody {
	s.Status = &v
	return s
}

func (s *GetArtifactResponseBody) SetStatusDetail(v string) *GetArtifactResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetArtifactResponseBody) SetSupportRegionIds(v string) *GetArtifactResponseBody {
	s.SupportRegionIds = &v
	return s
}

func (s *GetArtifactResponseBody) SetTags(v []*GetArtifactResponseBodyTags) *GetArtifactResponseBody {
	s.Tags = v
	return s
}

func (s *GetArtifactResponseBody) SetVersionName(v string) *GetArtifactResponseBody {
	s.VersionName = &v
	return s
}

func (s *GetArtifactResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetArtifactResponseBodyTags struct {
	// The tag key of the deployment package.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the deployment package.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetArtifactResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetArtifactResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetArtifactResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetArtifactResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetArtifactResponseBodyTags) SetKey(v string) *GetArtifactResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetArtifactResponseBodyTags) SetValue(v string) *GetArtifactResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetArtifactResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
