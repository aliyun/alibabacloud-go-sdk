// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildProperty(v string) *CreateArtifactResponseBody
	GetArtifactBuildProperty() *string
	SetArtifactBuildType(v string) *CreateArtifactResponseBody
	GetArtifactBuildType() *string
	SetArtifactId(v string) *CreateArtifactResponseBody
	GetArtifactId() *string
	SetArtifactProperty(v string) *CreateArtifactResponseBody
	GetArtifactProperty() *string
	SetArtifactType(v string) *CreateArtifactResponseBody
	GetArtifactType() *string
	SetArtifactVersion(v string) *CreateArtifactResponseBody
	GetArtifactVersion() *string
	SetDescription(v string) *CreateArtifactResponseBody
	GetDescription() *string
	SetGmtModified(v string) *CreateArtifactResponseBody
	GetGmtModified() *string
	SetMaxVersion(v int64) *CreateArtifactResponseBody
	GetMaxVersion() *int64
	SetName(v string) *CreateArtifactResponseBody
	GetName() *string
	SetRequestId(v string) *CreateArtifactResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateArtifactResponseBody
	GetStatus() *string
	SetStatusDetail(v string) *CreateArtifactResponseBody
	GetStatusDetail() *string
	SetSupportRegionIds(v string) *CreateArtifactResponseBody
	GetSupportRegionIds() *string
	SetVersionName(v string) *CreateArtifactResponseBody
	GetVersionName() *string
}

type CreateArtifactResponseBody struct {
	// The content used to build the artifact. This parameter is used for managed artifact builds.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the built artifact.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The artifact ID.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43ae****
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The content of the artifact.
	//
	// example:
	//
	// {\\"CommodityCode\\":\\"cmjj0005****\\",\\"CommodityVersion\\":\\"V2022****\\"}
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The artifact version.
	//
	// example:
	//
	// 2
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
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
	// 2022-11-11T12:00:00Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The latest version of the artifact.
	//
	// example:
	//
	// 1
	MaxVersion *int64 `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	// The artifact name.
	//
	// example:
	//
	// Redhat8_5 image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E73F09DC-6C13-5CB1-A10F-7A4E125ABD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the artifact.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the artifact status.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The IDs of the regions to which the artifact is distributed.
	//
	// example:
	//
	// [
	//
	// 			"cn-beijing",
	//
	// 			"cn-hangzhou",
	//
	// 			"cn-shanghai"
	//
	// 		]
	SupportRegionIds *string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty"`
	// The name of the artifact version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactResponseBody) GetArtifactBuildProperty() *string {
	return s.ArtifactBuildProperty
}

func (s *CreateArtifactResponseBody) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *CreateArtifactResponseBody) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *CreateArtifactResponseBody) GetArtifactProperty() *string {
	return s.ArtifactProperty
}

func (s *CreateArtifactResponseBody) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateArtifactResponseBody) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *CreateArtifactResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateArtifactResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateArtifactResponseBody) GetMaxVersion() *int64 {
	return s.MaxVersion
}

func (s *CreateArtifactResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateArtifactResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateArtifactResponseBody) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *CreateArtifactResponseBody) GetSupportRegionIds() *string {
	return s.SupportRegionIds
}

func (s *CreateArtifactResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateArtifactResponseBody) SetArtifactBuildProperty(v string) *CreateArtifactResponseBody {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *CreateArtifactResponseBody) SetArtifactBuildType(v string) *CreateArtifactResponseBody {
	s.ArtifactBuildType = &v
	return s
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

func (s *CreateArtifactResponseBody) SetStatusDetail(v string) *CreateArtifactResponseBody {
	s.StatusDetail = &v
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

func (s *CreateArtifactResponseBody) Validate() error {
	return dara.Validate(s)
}
