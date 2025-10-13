// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonReleasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReleases(v []*ListAddonReleasesResponseBodyReleases) *ListAddonReleasesResponseBody
	GetReleases() []*ListAddonReleasesResponseBodyReleases
	SetRequestId(v string) *ListAddonReleasesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAddonReleasesResponseBody
	GetTotal() *int64
}

type ListAddonReleasesResponseBody struct {
	Releases []*ListAddonReleasesResponseBodyReleases `json:"releases,omitempty" xml:"releases,omitempty" type:"Repeated"`
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAddonReleasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBody) GetReleases() []*ListAddonReleasesResponseBodyReleases {
	return s.Releases
}

func (s *ListAddonReleasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonReleasesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAddonReleasesResponseBody) SetReleases(v []*ListAddonReleasesResponseBodyReleases) *ListAddonReleasesResponseBody {
	s.Releases = v
	return s
}

func (s *ListAddonReleasesResponseBody) SetRequestId(v string) *ListAddonReleasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonReleasesResponseBody) SetTotal(v int64) *ListAddonReleasesResponseBody {
	s.Total = &v
	return s
}

func (s *ListAddonReleasesResponseBody) Validate() error {
	if s.Releases != nil {
		for _, item := range s.Releases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAddonReleasesResponseBodyReleases struct {
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// example:
	//
	// 6
	AlertRuleCount *int64                                             `json:"alertRuleCount,omitempty" xml:"alertRuleCount,omitempty"`
	ApiVersion     *string                                            `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	Conditions     []*ListAddonReleasesResponseBodyReleasesConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// example:
	//
	// {"install":{"mode":"auto-install","listenPort":"9400"},"discoverMode":"instances","discover":{"instances":"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44"},"scrapeInterval":"15","enableSecuritecs-nodeyGroupInjection":"true","metricTags":""}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// 2024-11-04T16:10:12+08:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3
	DashboardCount *int64           `json:"dashboardCount,omitempty" xml:"dashboardCount,omitempty"`
	EntityRules    *EntityGroupBase `json:"entityRules,omitempty" xml:"entityRules,omitempty"`
	// example:
	//
	// CS
	EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
	// example:
	//
	// policy-xxxxxxxxxxxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// 2
	ExporterCount *int64 `json:"exporterCount,omitempty" xml:"exporterCount,omitempty"`
	// example:
	//
	// true
	HaveConfig *bool `json:"haveConfig,omitempty" xml:"haveConfig,omitempty"`
	// example:
	//
	// 175xxxxxxxxx
	InstallUserId *string `json:"installUserId,omitempty" xml:"installUserId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// true
	Managed     *bool   `json:"managed,omitempty" xml:"managed,omitempty"`
	NextVersion *string `json:"nextVersion,omitempty" xml:"nextVersion,omitempty"`
	// example:
	//
	// policy-xxxxxxxxxxxxxx
	ParentAddonReleaseId *string `json:"parentAddonReleaseId,omitempty" xml:"parentAddonReleaseId,omitempty"`
	// example:
	//
	// policy-xxxxxxxxxxxxxx
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Release ID。
	//
	// example:
	//
	// 7339d808-66f9-4d40-83fa-xxxxxxxxxxx
	ReleaseId *string `json:"releaseId,omitempty" xml:"releaseId,omitempty"`
	// example:
	//
	// test-gpu-integration-name
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// example:
	//
	// container
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// running
	Status          *string                                               `json:"status,omitempty" xml:"status,omitempty"`
	SubAddonRelease *ListAddonReleasesResponseBodyReleasesSubAddonRelease `json:"subAddonRelease,omitempty" xml:"subAddonRelease,omitempty" type:"Struct"`
	// example:
	//
	// 2024-11-04T16:10:23+08:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 175xxxxxxxxxxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 0.0.2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// default
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListAddonReleasesResponseBodyReleases) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBodyReleases) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBodyReleases) GetAddonName() *string {
	return s.AddonName
}

func (s *ListAddonReleasesResponseBodyReleases) GetAlertRuleCount() *int64 {
	return s.AlertRuleCount
}

func (s *ListAddonReleasesResponseBodyReleases) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListAddonReleasesResponseBodyReleases) GetConditions() []*ListAddonReleasesResponseBodyReleasesConditions {
	return s.Conditions
}

func (s *ListAddonReleasesResponseBodyReleases) GetConfig() *string {
	return s.Config
}

func (s *ListAddonReleasesResponseBodyReleases) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAddonReleasesResponseBodyReleases) GetDashboardCount() *int64 {
	return s.DashboardCount
}

func (s *ListAddonReleasesResponseBodyReleases) GetEntityRules() *EntityGroupBase {
	return s.EntityRules
}

func (s *ListAddonReleasesResponseBodyReleases) GetEnvType() *string {
	return s.EnvType
}

func (s *ListAddonReleasesResponseBodyReleases) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListAddonReleasesResponseBodyReleases) GetExporterCount() *int64 {
	return s.ExporterCount
}

func (s *ListAddonReleasesResponseBodyReleases) GetHaveConfig() *bool {
	return s.HaveConfig
}

func (s *ListAddonReleasesResponseBodyReleases) GetInstallUserId() *string {
	return s.InstallUserId
}

func (s *ListAddonReleasesResponseBodyReleases) GetLanguage() *string {
	return s.Language
}

func (s *ListAddonReleasesResponseBodyReleases) GetManaged() *bool {
	return s.Managed
}

func (s *ListAddonReleasesResponseBodyReleases) GetNextVersion() *string {
	return s.NextVersion
}

func (s *ListAddonReleasesResponseBodyReleases) GetParentAddonReleaseId() *string {
	return s.ParentAddonReleaseId
}

func (s *ListAddonReleasesResponseBodyReleases) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListAddonReleasesResponseBodyReleases) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAddonReleasesResponseBodyReleases) GetReleaseId() *string {
	return s.ReleaseId
}

func (s *ListAddonReleasesResponseBodyReleases) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *ListAddonReleasesResponseBodyReleases) GetScene() *string {
	return s.Scene
}

func (s *ListAddonReleasesResponseBodyReleases) GetStatus() *string {
	return s.Status
}

func (s *ListAddonReleasesResponseBodyReleases) GetSubAddonRelease() *ListAddonReleasesResponseBodyReleasesSubAddonRelease {
	return s.SubAddonRelease
}

func (s *ListAddonReleasesResponseBodyReleases) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAddonReleasesResponseBodyReleases) GetUserId() *string {
	return s.UserId
}

func (s *ListAddonReleasesResponseBodyReleases) GetVersion() *string {
	return s.Version
}

func (s *ListAddonReleasesResponseBodyReleases) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListAddonReleasesResponseBodyReleases) SetAddonName(v string) *ListAddonReleasesResponseBodyReleases {
	s.AddonName = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetAlertRuleCount(v int64) *ListAddonReleasesResponseBodyReleases {
	s.AlertRuleCount = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetApiVersion(v string) *ListAddonReleasesResponseBodyReleases {
	s.ApiVersion = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetConditions(v []*ListAddonReleasesResponseBodyReleasesConditions) *ListAddonReleasesResponseBodyReleases {
	s.Conditions = v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetConfig(v string) *ListAddonReleasesResponseBodyReleases {
	s.Config = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetCreateTime(v string) *ListAddonReleasesResponseBodyReleases {
	s.CreateTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetDashboardCount(v int64) *ListAddonReleasesResponseBodyReleases {
	s.DashboardCount = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetEntityRules(v *EntityGroupBase) *ListAddonReleasesResponseBodyReleases {
	s.EntityRules = v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetEnvType(v string) *ListAddonReleasesResponseBodyReleases {
	s.EnvType = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetEnvironmentId(v string) *ListAddonReleasesResponseBodyReleases {
	s.EnvironmentId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetExporterCount(v int64) *ListAddonReleasesResponseBodyReleases {
	s.ExporterCount = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetHaveConfig(v bool) *ListAddonReleasesResponseBodyReleases {
	s.HaveConfig = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetInstallUserId(v string) *ListAddonReleasesResponseBodyReleases {
	s.InstallUserId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetLanguage(v string) *ListAddonReleasesResponseBodyReleases {
	s.Language = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetManaged(v bool) *ListAddonReleasesResponseBodyReleases {
	s.Managed = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetNextVersion(v string) *ListAddonReleasesResponseBodyReleases {
	s.NextVersion = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetParentAddonReleaseId(v string) *ListAddonReleasesResponseBodyReleases {
	s.ParentAddonReleaseId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetPolicyId(v string) *ListAddonReleasesResponseBodyReleases {
	s.PolicyId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetRegionId(v string) *ListAddonReleasesResponseBodyReleases {
	s.RegionId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetReleaseId(v string) *ListAddonReleasesResponseBodyReleases {
	s.ReleaseId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetReleaseName(v string) *ListAddonReleasesResponseBodyReleases {
	s.ReleaseName = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetScene(v string) *ListAddonReleasesResponseBodyReleases {
	s.Scene = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetStatus(v string) *ListAddonReleasesResponseBodyReleases {
	s.Status = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetSubAddonRelease(v *ListAddonReleasesResponseBodyReleasesSubAddonRelease) *ListAddonReleasesResponseBodyReleases {
	s.SubAddonRelease = v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetUpdateTime(v string) *ListAddonReleasesResponseBodyReleases {
	s.UpdateTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetUserId(v string) *ListAddonReleasesResponseBodyReleases {
	s.UserId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetVersion(v string) *ListAddonReleasesResponseBodyReleases {
	s.Version = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) SetWorkspace(v string) *ListAddonReleasesResponseBodyReleases {
	s.Workspace = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleases) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EntityRules != nil {
		if err := s.EntityRules.Validate(); err != nil {
			return err
		}
	}
	if s.SubAddonRelease != nil {
		if err := s.SubAddonRelease.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAddonReleasesResponseBodyReleasesConditions struct {
	// example:
	//
	// 2024-11-04T16:10:22+08:00
	FirstTransitionTime *string `json:"firstTransitionTime,omitempty" xml:"firstTransitionTime,omitempty"`
	// example:
	//
	// 2024-11-04T16:10:22+08:00
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
	// example:
	//
	// The addon loaded successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// True
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// Loaded
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAddonReleasesResponseBodyReleasesConditions) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBodyReleasesConditions) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) GetMessage() *string {
	return s.Message
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) GetStatus() *string {
	return s.Status
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) GetType() *string {
	return s.Type
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) SetFirstTransitionTime(v string) *ListAddonReleasesResponseBodyReleasesConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) SetLastTransitionTime(v string) *ListAddonReleasesResponseBodyReleasesConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) SetMessage(v string) *ListAddonReleasesResponseBodyReleasesConditions {
	s.Message = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) SetStatus(v string) *ListAddonReleasesResponseBodyReleasesConditions {
	s.Status = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) SetType(v string) *ListAddonReleasesResponseBodyReleasesConditions {
	s.Type = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesConditions) Validate() error {
	return dara.Validate(s)
}

type ListAddonReleasesResponseBodyReleasesSubAddonRelease struct {
	Ready *int32 `json:"ready,omitempty" xml:"ready,omitempty"`
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAddonReleasesResponseBodyReleasesSubAddonRelease) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBodyReleasesSubAddonRelease) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBodyReleasesSubAddonRelease) GetReady() *int32 {
	return s.Ready
}

func (s *ListAddonReleasesResponseBodyReleasesSubAddonRelease) GetTotal() *int32 {
	return s.Total
}

func (s *ListAddonReleasesResponseBodyReleasesSubAddonRelease) SetReady(v int32) *ListAddonReleasesResponseBodyReleasesSubAddonRelease {
	s.Ready = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesSubAddonRelease) SetTotal(v int32) *ListAddonReleasesResponseBodyReleasesSubAddonRelease {
	s.Total = &v
	return s
}

func (s *ListAddonReleasesResponseBodyReleasesSubAddonRelease) Validate() error {
	return dara.Validate(s)
}
