// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishEdgeContainerAppVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PublishEdgeContainerAppVersionRequest
	GetAppId() *string
	SetFullRelease(v bool) *PublishEdgeContainerAppVersionRequest
	GetFullRelease() *bool
	SetPercentage(v int32) *PublishEdgeContainerAppVersionRequest
	GetPercentage() *int32
	SetPublishEnv(v string) *PublishEdgeContainerAppVersionRequest
	GetPublishEnv() *string
	SetPublishType(v string) *PublishEdgeContainerAppVersionRequest
	GetPublishType() *string
	SetRegions(v []*string) *PublishEdgeContainerAppVersionRequest
	GetRegions() []*string
	SetRemarks(v string) *PublishEdgeContainerAppVersionRequest
	GetRemarks() *string
	SetStartTime(v string) *PublishEdgeContainerAppVersionRequest
	GetStartTime() *string
	SetVersionId(v string) *PublishEdgeContainerAppVersionRequest
	GetVersionId() *string
}

type PublishEdgeContainerAppVersionRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to fully release the version. This parameter takes effect only when PublishType is set to region.
	//
	// example:
	//
	// true
	FullRelease *bool `json:"FullRelease,omitempty" xml:"FullRelease,omitempty"`
	// The release percentage. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The environment to which you want to release the version. Valid values:
	//
	// 	- prod: the production environment.
	//
	// 	- staging: the staging environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod
	PublishEnv *string `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
	// Specifies how the version is released. Valid values:
	//
	// 	- percentage: releases the version by percentage.
	//
	// 	- region: releases the version by region.
	//
	// If you do not specify this parameter, the version is released by percentage by default.
	//
	// example:
	//
	// percentage
	PublishType *string `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// The regions to which the version is released.
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The remarks. This parameter is empty by default.
	//
	// example:
	//
	// test publish app
	Remarks *string `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	// The time when the application version starts to be released. If you do not specify this parameter, the current time is used by default.
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

func (s PublishEdgeContainerAppVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishEdgeContainerAppVersionRequest) GoString() string {
	return s.String()
}

func (s *PublishEdgeContainerAppVersionRequest) GetAppId() *string {
	return s.AppId
}

func (s *PublishEdgeContainerAppVersionRequest) GetFullRelease() *bool {
	return s.FullRelease
}

func (s *PublishEdgeContainerAppVersionRequest) GetPercentage() *int32 {
	return s.Percentage
}

func (s *PublishEdgeContainerAppVersionRequest) GetPublishEnv() *string {
	return s.PublishEnv
}

func (s *PublishEdgeContainerAppVersionRequest) GetPublishType() *string {
	return s.PublishType
}

func (s *PublishEdgeContainerAppVersionRequest) GetRegions() []*string {
	return s.Regions
}

func (s *PublishEdgeContainerAppVersionRequest) GetRemarks() *string {
	return s.Remarks
}

func (s *PublishEdgeContainerAppVersionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *PublishEdgeContainerAppVersionRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *PublishEdgeContainerAppVersionRequest) SetAppId(v string) *PublishEdgeContainerAppVersionRequest {
	s.AppId = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetFullRelease(v bool) *PublishEdgeContainerAppVersionRequest {
	s.FullRelease = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetPercentage(v int32) *PublishEdgeContainerAppVersionRequest {
	s.Percentage = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetPublishEnv(v string) *PublishEdgeContainerAppVersionRequest {
	s.PublishEnv = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetPublishType(v string) *PublishEdgeContainerAppVersionRequest {
	s.PublishType = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetRegions(v []*string) *PublishEdgeContainerAppVersionRequest {
	s.Regions = v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetRemarks(v string) *PublishEdgeContainerAppVersionRequest {
	s.Remarks = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetStartTime(v string) *PublishEdgeContainerAppVersionRequest {
	s.StartTime = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) SetVersionId(v string) *PublishEdgeContainerAppVersionRequest {
	s.VersionId = &v
	return s
}

func (s *PublishEdgeContainerAppVersionRequest) Validate() error {
	return dara.Validate(s)
}
