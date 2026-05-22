// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppStatus(v *GetEdgeContainerAppStatusResponseBodyAppStatus) *GetEdgeContainerAppStatusResponseBody
	GetAppStatus() *GetEdgeContainerAppStatusResponseBodyAppStatus
	SetRequestId(v string) *GetEdgeContainerAppStatusResponseBody
	GetRequestId() *string
}

type GetEdgeContainerAppStatusResponseBody struct {
	AppStatus *GetEdgeContainerAppStatusResponseBodyAppStatus `json:"AppStatus,omitempty" xml:"AppStatus,omitempty" type:"Struct"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeContainerAppStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppStatusResponseBody) GetAppStatus() *GetEdgeContainerAppStatusResponseBodyAppStatus {
	return s.AppStatus
}

func (s *GetEdgeContainerAppStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppStatusResponseBody) SetAppStatus(v *GetEdgeContainerAppStatusResponseBodyAppStatus) *GetEdgeContainerAppStatusResponseBody {
	s.AppStatus = v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBody) SetRequestId(v string) *GetEdgeContainerAppStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBody) Validate() error {
	if s.AppStatus != nil {
		if err := s.AppStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeContainerAppStatusResponseBodyAppStatus struct {
	BaseLineVersion   *string                                                `json:"BaseLineVersion,omitempty" xml:"BaseLineVersion,omitempty"`
	DeployStatus      *string                                                `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	DeployTime        *string                                                `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	DeployedVersion   *string                                                `json:"DeployedVersion,omitempty" xml:"DeployedVersion,omitempty"`
	ExpectPercentage  *int32                                                 `json:"ExpectPercentage,omitempty" xml:"ExpectPercentage,omitempty"`
	FullRelease       *bool                                                  `json:"FullRelease,omitempty" xml:"FullRelease,omitempty"`
	PublishEnv        *string                                                `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
	PublishPercentage *int32                                                 `json:"PublishPercentage,omitempty" xml:"PublishPercentage,omitempty"`
	PublishStatus     *string                                                `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	PublishTime       *string                                                `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	PublishType       *string                                                `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	PublishingVersion *string                                                `json:"PublishingVersion,omitempty" xml:"PublishingVersion,omitempty"`
	Regions           *GetEdgeContainerAppStatusResponseBodyAppStatusRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RollbackTime      *string                                                `json:"RollbackTime,omitempty" xml:"RollbackTime,omitempty"`
	UnDeployTime      *string                                                `json:"UnDeployTime,omitempty" xml:"UnDeployTime,omitempty"`
}

func (s GetEdgeContainerAppStatusResponseBodyAppStatus) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppStatusResponseBodyAppStatus) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetBaseLineVersion() *string {
	return s.BaseLineVersion
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetDeployTime() *string {
	return s.DeployTime
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetDeployedVersion() *string {
	return s.DeployedVersion
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetExpectPercentage() *int32 {
	return s.ExpectPercentage
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetFullRelease() *bool {
	return s.FullRelease
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetPublishEnv() *string {
	return s.PublishEnv
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetPublishPercentage() *int32 {
	return s.PublishPercentage
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetPublishStatus() *string {
	return s.PublishStatus
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetPublishTime() *string {
	return s.PublishTime
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetPublishType() *string {
	return s.PublishType
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetPublishingVersion() *string {
	return s.PublishingVersion
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetRegions() *GetEdgeContainerAppStatusResponseBodyAppStatusRegions {
	return s.Regions
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetRollbackTime() *string {
	return s.RollbackTime
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) GetUnDeployTime() *string {
	return s.UnDeployTime
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetBaseLineVersion(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.BaseLineVersion = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetDeployStatus(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.DeployStatus = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetDeployTime(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.DeployTime = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetDeployedVersion(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.DeployedVersion = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetExpectPercentage(v int32) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.ExpectPercentage = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetFullRelease(v bool) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.FullRelease = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetPublishEnv(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.PublishEnv = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetPublishPercentage(v int32) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.PublishPercentage = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetPublishStatus(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.PublishStatus = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetPublishTime(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.PublishTime = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetPublishType(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.PublishType = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetPublishingVersion(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.PublishingVersion = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetRegions(v *GetEdgeContainerAppStatusResponseBodyAppStatusRegions) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.Regions = v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetRollbackTime(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.RollbackTime = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) SetUnDeployTime(v string) *GetEdgeContainerAppStatusResponseBodyAppStatus {
	s.UnDeployTime = &v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatus) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEdgeContainerAppStatusResponseBodyAppStatusRegions struct {
	Region []*string `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s GetEdgeContainerAppStatusResponseBodyAppStatusRegions) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppStatusResponseBodyAppStatusRegions) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatusRegions) GetRegion() []*string {
	return s.Region
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatusRegions) SetRegion(v []*string) *GetEdgeContainerAppStatusResponseBodyAppStatusRegions {
	s.Region = v
	return s
}

func (s *GetEdgeContainerAppStatusResponseBodyAppStatusRegions) Validate() error {
	return dara.Validate(s)
}
