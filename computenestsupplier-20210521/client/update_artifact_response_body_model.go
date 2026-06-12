// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildProperty(v string) *UpdateArtifactResponseBody
	GetArtifactBuildProperty() *string
	SetArtifactBuildType(v string) *UpdateArtifactResponseBody
	GetArtifactBuildType() *string
	SetArtifactId(v string) *UpdateArtifactResponseBody
	GetArtifactId() *string
	SetArtifactProperty(v string) *UpdateArtifactResponseBody
	GetArtifactProperty() *string
	SetArtifactType(v string) *UpdateArtifactResponseBody
	GetArtifactType() *string
	SetArtifactVersion(v string) *UpdateArtifactResponseBody
	GetArtifactVersion() *string
	SetDescription(v string) *UpdateArtifactResponseBody
	GetDescription() *string
	SetGmtModified(v string) *UpdateArtifactResponseBody
	GetGmtModified() *string
	SetRequestId(v string) *UpdateArtifactResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateArtifactResponseBody
	GetStatus() *string
	SetStatusDetail(v string) *UpdateArtifactResponseBody
	GetStatusDetail() *string
	SetSupportRegionIds(v string) *UpdateArtifactResponseBody
	GetSupportRegionIds() *string
	SetVersionName(v string) *UpdateArtifactResponseBody
	GetVersionName() *string
}

type UpdateArtifactResponseBody struct {
	// The properties for building the artifact. This is used for managed artifact builds.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
	ArtifactBuildProperty *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The build type of the artifact.
	//
	// example:
	//
	// "{\\"RegionId\\":\\"xxx\\", \\"SourceImageId\\":\\"xxx\\", \\"\\":\\"xxx\\", \\"CommandType\\":\\"xxx\\", \\"CommandContent\\":\\"xxx\\"}"
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
	// 2022-10-20T02:19:55Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9F350409-2ACC-5B61-ACD9-3C8995792F8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details of the artifact status.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The IDs of the regions to which the artifact can be distributed.
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
	// The version name.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactResponseBody) GetArtifactBuildProperty() *string {
	return s.ArtifactBuildProperty
}

func (s *UpdateArtifactResponseBody) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *UpdateArtifactResponseBody) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *UpdateArtifactResponseBody) GetArtifactProperty() *string {
	return s.ArtifactProperty
}

func (s *UpdateArtifactResponseBody) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UpdateArtifactResponseBody) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *UpdateArtifactResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateArtifactResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateArtifactResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateArtifactResponseBody) GetStatusDetail() *string {
	return s.StatusDetail
}

func (s *UpdateArtifactResponseBody) GetSupportRegionIds() *string {
	return s.SupportRegionIds
}

func (s *UpdateArtifactResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *UpdateArtifactResponseBody) SetArtifactBuildProperty(v string) *UpdateArtifactResponseBody {
	s.ArtifactBuildProperty = &v
	return s
}

func (s *UpdateArtifactResponseBody) SetArtifactBuildType(v string) *UpdateArtifactResponseBody {
	s.ArtifactBuildType = &v
	return s
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

func (s *UpdateArtifactResponseBody) SetStatusDetail(v string) *UpdateArtifactResponseBody {
	s.StatusDetail = &v
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

func (s *UpdateArtifactResponseBody) Validate() error {
	return dara.Validate(s)
}
