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
	// This parameter is required.
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	FullRelease *bool   `json:"FullRelease,omitempty" xml:"FullRelease,omitempty"`
	Percentage  *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// This parameter is required.
	PublishEnv  *string   `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
	PublishType *string   `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	Regions     []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	Remarks     *string   `json:"Remarks,omitempty" xml:"Remarks,omitempty"`
	StartTime   *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// This parameter is required.
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
