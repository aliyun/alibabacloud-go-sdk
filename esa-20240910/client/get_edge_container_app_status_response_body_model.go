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
	// The status of the application.
	AppStatus *GetEdgeContainerAppStatusResponseBodyAppStatus `json:"AppStatus,omitempty" xml:"AppStatus,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type GetEdgeContainerAppStatusResponseBodyAppStatus struct {
	// The base version of the application.
	//
	// example:
	//
	// ver-123123123123****
	BaseLineVersion *string `json:"BaseLineVersion,omitempty" xml:"BaseLineVersion,omitempty"`
	// The deployment status of the application.
	//
	// 	- **undeploy**: The application is not deployed.
	//
	// 	- **deploying**: The application is being deployed.
	//
	// 	- **deployed**: The application is deployed.
	//
	// 	- **undeploying**: The deployment is being canceled.
	//
	// example:
	//
	// undeploy
	DeployStatus *string `json:"DeployStatus,omitempty" xml:"DeployStatus,omitempty"`
	// The time when the application was deployed. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-07-25T05:58:05Z
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// The release version of the application.
	//
	// example:
	//
	// ver-123123123123****
	DeployedVersion *string `json:"DeployedVersion,omitempty" xml:"DeployedVersion,omitempty"`
	// The expected release percentage of the application.
	//
	// example:
	//
	// 50%
	ExpectPercentage *int32 `json:"ExpectPercentage,omitempty" xml:"ExpectPercentage,omitempty"`
	// Specifies whether to fully release the version. This parameter takes effect only when PublishType is set to region.
	//
	// example:
	//
	// true
	FullRelease *bool `json:"FullRelease,omitempty" xml:"FullRelease,omitempty"`
	// The environment to which the application was released. Valid values:
	//
	// 	- **prod**: the production environment.
	//
	// 	- **staging**: the staging environment.
	//
	// example:
	//
	// prod
	PublishEnv *string `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
	// The release percentage of the application.
	//
	// example:
	//
	// 50%
	PublishPercentage *int32 `json:"PublishPercentage,omitempty" xml:"PublishPercentage,omitempty"`
	// The release status of the application. Valid values:
	//
	// 	- **publishing**
	//
	// 	- **published**
	//
	// 	- **rollbacking**
	//
	// 	- **rollbacked**
	//
	// example:
	//
	// pubishing
	PublishStatus *string `json:"PublishStatus,omitempty" xml:"PublishStatus,omitempty"`
	// The time when the application was released. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-07-25T05:58:05Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
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
	// The release version of the application.
	//
	// example:
	//
	// ver-123123123123****
	PublishingVersion *string `json:"PublishingVersion,omitempty" xml:"PublishingVersion,omitempty"`
	// The regions to which the version is released.
	Regions *GetEdgeContainerAppStatusResponseBodyAppStatusRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The time when the last rollback was performed.
	//
	// example:
	//
	// 2024-07-25T05:58:05Z
	RollbackTime *string `json:"RollbackTime,omitempty" xml:"RollbackTime,omitempty"`
	// The time when the application deployment was canceled. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-07-25T05:58:05Z
	UnDeployTime *string `json:"UnDeployTime,omitempty" xml:"UnDeployTime,omitempty"`
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
	return dara.Validate(s)
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
