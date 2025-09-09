// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody
	GetArtifacts() []*ListArtifactsResponseBodyArtifacts
	SetMaxResults(v int32) *ListArtifactsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListArtifactsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListArtifactsResponseBody
	GetTotalCount() *int32
}

type ListArtifactsResponseBody struct {
	// The information about deployment packages.
	Artifacts []*ListArtifactsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hRQzE=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46577928-3162-15A6-9084-69820EB9xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBody) GetArtifacts() []*ListArtifactsResponseBodyArtifacts {
	return s.Artifacts
}

func (s *ListArtifactsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListArtifactsResponseBody) SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactsResponseBody) SetMaxResults(v int32) *ListArtifactsResponseBody {
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

func (s *ListArtifactsResponseBody) SetTotalCount(v int32) *ListArtifactsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListArtifactsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListArtifactsResponseBodyArtifacts struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	//
	// example:
	//
	// {\\"CodeRepo\\":{\\"Owner\\":\\"wenle\\",\\"Platform\\":\\"github\\",\\"Branch\\":\\"main\\",\\"RepoName\\":\\"aliyun-computenest/java-springboot-demo\\"}}
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The type of the deployment package.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
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
	MaxVersion *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek25rexxxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the deployment package. Valid values:
	//
	// 	- Created: The deployment package is created.
	//
	// 	- Scanning: The deployment package is being scanned.
	//
	// 	- ScanFailed: The deployment package failed to be scanned.
	//
	// 	- Delivering: The deployment package is being distributed.
	//
	// 	- Available: The deployment package is available.
	//
	// 	- Deleted: The deployment package is deleted.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags []*ListArtifactsResponseBodyArtifactsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListArtifactsResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifacts) GetArtifactBuildProperty() *string {
	return s.ArtifactBuildProperty
}

func (s *ListArtifactsResponseBodyArtifacts) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ListArtifactsResponseBodyArtifacts) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ListArtifactsResponseBodyArtifacts) GetDescription() *string {
	return s.Description
}

func (s *ListArtifactsResponseBodyArtifacts) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListArtifactsResponseBodyArtifacts) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *ListArtifactsResponseBodyArtifacts) GetName() *string {
	return s.Name
}

func (s *ListArtifactsResponseBodyArtifacts) GetPermissionType() *string {
	return s.PermissionType
}

func (s *ListArtifactsResponseBodyArtifacts) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListArtifactsResponseBodyArtifacts) GetStatus() *string {
	return s.Status
}

func (s *ListArtifactsResponseBodyArtifacts) GetTags() []*ListArtifactsResponseBodyArtifactsTags {
	return s.Tags
}

func (s *ListArtifactsResponseBodyArtifacts) SetArtifactBuildProperty(v string) *ListArtifactsResponseBodyArtifacts {
	s.ArtifactBuildProperty = &v
	return s
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

func (s *ListArtifactsResponseBodyArtifacts) SetPermissionType(v string) *ListArtifactsResponseBodyArtifacts {
	s.PermissionType = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetResourceGroupId(v string) *ListArtifactsResponseBodyArtifacts {
	s.ResourceGroupId = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetStatus(v string) *ListArtifactsResponseBodyArtifacts {
	s.Status = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetTags(v []*ListArtifactsResponseBodyArtifactsTags) *ListArtifactsResponseBodyArtifacts {
	s.Tags = v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) Validate() error {
	return dara.Validate(s)
}

type ListArtifactsResponseBodyArtifactsTags struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListArtifactsResponseBodyArtifactsTags) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifactsTags) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifactsTags) GetKey() *string {
	return s.Key
}

func (s *ListArtifactsResponseBodyArtifactsTags) GetValue() *string {
	return s.Value
}

func (s *ListArtifactsResponseBodyArtifactsTags) SetKey(v string) *ListArtifactsResponseBodyArtifactsTags {
	s.Key = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifactsTags) SetValue(v string) *ListArtifactsResponseBodyArtifactsTags {
	s.Value = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifactsTags) Validate() error {
	return dara.Validate(s)
}
