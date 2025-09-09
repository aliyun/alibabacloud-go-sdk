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
	// The request ID.
	//
	// example:
	//
	// B1A0198B-F316-1B72-B8DD-28B6F6D6XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the deployment package.
	//
	// example:
	//
	// "/usr/local/share/aliyun-assist/work/script/t-hz04zm90y6og0sg.sh: line 1: pip: command not found"
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// The ID of the region that supports the deployment package.
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
	// The name of the deployment package.
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
