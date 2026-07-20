// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishEdgeContainerAppVersionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetAppId() *string
	SetFullRelease(v bool) *PublishEdgeContainerAppVersionShrinkRequest
	GetFullRelease() *bool
	SetPercentage(v int32) *PublishEdgeContainerAppVersionShrinkRequest
	GetPercentage() *int32
	SetPublishEnv(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetPublishEnv() *string
	SetPublishType(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetPublishType() *string
	SetRegionsShrink(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetRegionsShrink() *string
	SetRemarks(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetRemarks() *string
	SetStartTime(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetStartTime() *string
	SetVersionId(v string) *PublishEdgeContainerAppVersionShrinkRequest
	GetVersionId() *string
}

type PublishEdgeContainerAppVersionShrinkRequest struct {
	// The application ID.
	//
	// > 1) Obtain the AppId by calling CreateEdgeContainerApp. 2) Obtain the VersionId by calling CreateEdgeContainerAppVersion (which requires the AppId). 3) The complete call chain is CreateEdgeContainerApp → CreateEdgeContainerAppVersion → PublishEdgeContainerAppVersion.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to perform a full release. This parameter takes effect only when PublishType is set to region.
	//
	// example:
	//
	// true
	FullRelease *bool `json:"FullRelease,omitempty" xml:"FullRelease,omitempty"`
	// The publishing percentage. Valid values: **1 to 100**. Default value: **100**.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The publishing environment. Valid values:
	//
	// - **prod**: production environment.
	//
	// - **staging**: staging environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod
	PublishEnv *string `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
	// The publishing type. Valid values:
	//
	// - **percentage**: Publish by percentage.
	//
	// - **region**: Publish by region.
	//
	// If this parameter is not specified, percentage-based publishing is used by default.
	//
	// example:
	//
	// percentage
	PublishType *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// The list of publishing regions.
	RegionsShrink *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The remarks. Default value: empty.
	//
	// example:
	//
	// test publish app
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The time when the publishing starts. If this parameter is not specified, the current time is used by default.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2023-06-05T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ver-87962637161651****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s PublishEdgeContainerAppVersionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishEdgeContainerAppVersionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetFullRelease() *bool {
	return s.FullRelease
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetPercentage() *int32 {
	return s.Percentage
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetPublishEnv() *string {
	return s.PublishEnv
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetPublishType() *string {
	return s.PublishType
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetRegionsShrink() *string {
	return s.RegionsShrink
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetAppId(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.AppId = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetFullRelease(v bool) *PublishEdgeContainerAppVersionShrinkRequest {
	s.FullRelease = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetPercentage(v int32) *PublishEdgeContainerAppVersionShrinkRequest {
	s.Percentage = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetPublishEnv(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.PublishEnv = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetPublishType(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.PublishType = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetRegionsShrink(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.RegionsShrink = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetRemarks(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.Remarks = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetStartTime(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) SetVersionId(v string) *PublishEdgeContainerAppVersionShrinkRequest {
	s.VersionId = &v
	return s
}

func (s *PublishEdgeContainerAppVersionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
