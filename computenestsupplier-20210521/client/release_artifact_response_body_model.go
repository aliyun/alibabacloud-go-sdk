// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseArtifactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ReleaseArtifactResponseBody
	GetArtifactId() *string
	SetArtifactProperty(v string) *ReleaseArtifactResponseBody
	GetArtifactProperty() *string
	SetArtifactType(v string) *ReleaseArtifactResponseBody
	GetArtifactType() *string
	SetArtifactVersion(v string) *ReleaseArtifactResponseBody
	GetArtifactVersion() *string
	SetDescription(v string) *ReleaseArtifactResponseBody
	GetDescription() *string
	SetGmtModified(v string) *ReleaseArtifactResponseBody
	GetGmtModified() *string
	SetRequestId(v string) *ReleaseArtifactResponseBody
	GetRequestId() *string
	SetStatus(v string) *ReleaseArtifactResponseBody
	GetStatus() *string
	SetVersionName(v string) *ReleaseArtifactResponseBody
	GetVersionName() *string
}

type ReleaseArtifactResponseBody struct {
	// The ID of the artifact.
	//
	// example:
	//
	// artifact-9feded91880e4c78xxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The content of the artifact.
	//
	// example:
	//
	// "{\\"Url\\":\\"https://computenest-artifacts-draft-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/130920852836xxxx/cn-hangzhou/service-8072a04e5a134382xxxx/165095355xxxx/changes.txt\\",\\"ConfigurationMetadata\\":\\"{\\\\\\"WorkDir\\\\\\":\\\\\\"/root\\\\\\",\\\\\\"Platform\\\\\\":\\\\\\"Linux\\\\\\",\\\\\\"CommandType\\\\\\":\\\\\\"RunShellScript\\\\\\",\\\\\\"UpgradeScript\\\\\\":\\\\\\"cd /root\\\\\\\\ncp changes.txt cpchanges.txt\\\\\\\\nmv changes.txt mvchangge.txt\\\\\\"}\\"}"
	ArtifactProperty *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The type of the artifact.
	//
	// example:
	//
	// File
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The version of the artifact.
	//
	// example:
	//
	// draft
	ArtifactVersion *string `json:"ArtifactVersion,omitempty" xml:"ArtifactVersion,omitempty"`
	// The description of the artifact.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the artifact was modified.
	//
	// example:
	//
	// 1650954178000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3818BA7D-3F50-1A44-9FF3-04A52A59XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the artifact. Valid values:
	//
	// 	- Created: The artifact is created.
	//
	// 	- Scanning: The artifact is being scanned.
	//
	// 	- ScanFailed: The artifact failed to be scanned.
	//
	// 	- Delivering: The artifact is being distributed.
	//
	// 	- Available: The artifact is available.
	//
	// 	- Deleted: The artifact is deleted.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version name of the artifact.
	//
	// example:
	//
	// V1.0
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ReleaseArtifactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseArtifactResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseArtifactResponseBody) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ReleaseArtifactResponseBody) GetArtifactProperty() *string {
	return s.ArtifactProperty
}

func (s *ReleaseArtifactResponseBody) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *ReleaseArtifactResponseBody) GetArtifactVersion() *string {
	return s.ArtifactVersion
}

func (s *ReleaseArtifactResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ReleaseArtifactResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ReleaseArtifactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseArtifactResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ReleaseArtifactResponseBody) GetVersionName() *string {
	return s.VersionName
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

func (s *ReleaseArtifactResponseBody) Validate() error {
	return dara.Validate(s)
}
