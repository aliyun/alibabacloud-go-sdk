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
	// The information about the artifacts.
	Artifacts []*ListArtifactsResponseBodyArtifacts `json:"Artifacts,omitempty" xml:"Artifacts,omitempty" type:"Repeated"`
	// The number of entries returned per page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set it to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbfVwapgMwCN1wYzPVzLbItEdB0uWSY7AGnM3qCgm/YnjuEfwSnMwiMkcUoI0hR****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B7AE429-B53E-5E73-A5EC-DC91F614F2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that meet the filter criteria.
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
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactsResponseBodyArtifacts struct {
	// The content used to build the artifact. This parameter is used for hosted artifact builds.
	//
	// example:
	//
	// {\\"CodeRepo\\":{\\"Owner\\":\\"wenle\\",\\"Platform\\":\\"github\\",\\"Branch\\":\\"main\\",\\"RepoName\\":\\"heroku/node-js-getting-started\\"}}
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The artifact ID.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The description of the artifact.
	//
	// example:
	//
	// Redhat8_0 image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the artifact was modified.
	//
	// example:
	//
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The latest version.
	//
	// example:
	//
	// 2
	MaxVersion *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The artifact name.
	//
	// example:
	//
	// Redhat8_5 image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The permission field. This parameter is valid for artifacts of the container image, Helm chart, and file types. For other types of artifacts, you can only change the permission from Automatic to Public.
	//
	// Valid values:
	//
	// - Public
	//
	// - Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzmhzoaa****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the artifact.
	//
	// Valid values:
	//
	// - Created: The artifact is created.
	//
	// - Scanning: The artifact is being scanned.
	//
	// - ScanFailed: The artifact failed to be scanned.
	//
	// - Delivering: The artifact is being distributed.
	//
	// - Available: The artifact is available.
	//
	// - Deleted: The artifact is deleted.
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
