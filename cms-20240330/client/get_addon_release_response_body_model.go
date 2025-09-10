// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetAddonReleaseResponseBody
	GetConfig() *string
	SetRelease(v *GetAddonReleaseResponseBodyRelease) *GetAddonReleaseResponseBody
	GetRelease() *GetAddonReleaseResponseBodyRelease
	SetRequestId(v string) *GetAddonReleaseResponseBody
	GetRequestId() *string
}

type GetAddonReleaseResponseBody struct {
	// example:
	//
	// {"install":{"mode":"auto-install","listenPort":"9400"},"discoverMode":"instances","discover":{"instances":"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44"},"scrapeInterval":"15","enableSecuritecs-nodeyGroupInjection":"true","metricTags":""}
	Config  *string                             `json:"config,omitempty" xml:"config,omitempty"`
	Release *GetAddonReleaseResponseBodyRelease `json:"release,omitempty" xml:"release,omitempty" type:"Struct"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-A01D6CC3F4B8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetAddonReleaseResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetAddonReleaseResponseBody) GetRelease() *GetAddonReleaseResponseBodyRelease {
	return s.Release
}

func (s *GetAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAddonReleaseResponseBody) SetConfig(v string) *GetAddonReleaseResponseBody {
	s.Config = &v
	return s
}

func (s *GetAddonReleaseResponseBody) SetRelease(v *GetAddonReleaseResponseBodyRelease) *GetAddonReleaseResponseBody {
	s.Release = v
	return s
}

func (s *GetAddonReleaseResponseBody) SetRequestId(v string) *GetAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAddonReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAddonReleaseResponseBodyRelease struct {
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// example:
	//
	// 6
	AlertRuleCount *int64                                          `json:"alertRuleCount,omitempty" xml:"alertRuleCount,omitempty"`
	Conditions     []*GetAddonReleaseResponseBodyReleaseConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
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
	// policy-xxxxxxxxxxx
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
	// 1707xxxxxxxxxxxx
	InstallUserId *string `json:"installUserId,omitempty" xml:"installUserId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// true
	Managed *bool `json:"managed,omitempty" xml:"managed,omitempty"`
	// example:
	//
	// policy-xxxxxxxxxxxxx
	ParentAddonReleaseId *string `json:"parentAddonReleaseId,omitempty" xml:"parentAddonReleaseId,omitempty"`
	// example:
	//
	// policy-xxxxxxxxxxxxx
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
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2024-11-04T16:10:12+08:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 1707xxxxxxxxxxxx
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

func (s GetAddonReleaseResponseBodyRelease) String() string {
	return dara.Prettify(s)
}

func (s GetAddonReleaseResponseBodyRelease) GoString() string {
	return s.String()
}

func (s *GetAddonReleaseResponseBodyRelease) GetAddonName() *string {
	return s.AddonName
}

func (s *GetAddonReleaseResponseBodyRelease) GetAlertRuleCount() *int64 {
	return s.AlertRuleCount
}

func (s *GetAddonReleaseResponseBodyRelease) GetConditions() []*GetAddonReleaseResponseBodyReleaseConditions {
	return s.Conditions
}

func (s *GetAddonReleaseResponseBodyRelease) GetConfig() *string {
	return s.Config
}

func (s *GetAddonReleaseResponseBodyRelease) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAddonReleaseResponseBodyRelease) GetDashboardCount() *int64 {
	return s.DashboardCount
}

func (s *GetAddonReleaseResponseBodyRelease) GetEntityRules() *EntityGroupBase {
	return s.EntityRules
}

func (s *GetAddonReleaseResponseBodyRelease) GetEnvType() *string {
	return s.EnvType
}

func (s *GetAddonReleaseResponseBodyRelease) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetAddonReleaseResponseBodyRelease) GetExporterCount() *int64 {
	return s.ExporterCount
}

func (s *GetAddonReleaseResponseBodyRelease) GetHaveConfig() *bool {
	return s.HaveConfig
}

func (s *GetAddonReleaseResponseBodyRelease) GetInstallUserId() *string {
	return s.InstallUserId
}

func (s *GetAddonReleaseResponseBodyRelease) GetLanguage() *string {
	return s.Language
}

func (s *GetAddonReleaseResponseBodyRelease) GetManaged() *bool {
	return s.Managed
}

func (s *GetAddonReleaseResponseBodyRelease) GetParentAddonReleaseId() *string {
	return s.ParentAddonReleaseId
}

func (s *GetAddonReleaseResponseBodyRelease) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetAddonReleaseResponseBodyRelease) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAddonReleaseResponseBodyRelease) GetReleaseId() *string {
	return s.ReleaseId
}

func (s *GetAddonReleaseResponseBodyRelease) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *GetAddonReleaseResponseBodyRelease) GetScene() *string {
	return s.Scene
}

func (s *GetAddonReleaseResponseBodyRelease) GetStatus() *string {
	return s.Status
}

func (s *GetAddonReleaseResponseBodyRelease) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetAddonReleaseResponseBodyRelease) GetUserId() *string {
	return s.UserId
}

func (s *GetAddonReleaseResponseBodyRelease) GetVersion() *string {
	return s.Version
}

func (s *GetAddonReleaseResponseBodyRelease) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetAddonReleaseResponseBodyRelease) SetAddonName(v string) *GetAddonReleaseResponseBodyRelease {
	s.AddonName = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetAlertRuleCount(v int64) *GetAddonReleaseResponseBodyRelease {
	s.AlertRuleCount = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetConditions(v []*GetAddonReleaseResponseBodyReleaseConditions) *GetAddonReleaseResponseBodyRelease {
	s.Conditions = v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetConfig(v string) *GetAddonReleaseResponseBodyRelease {
	s.Config = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetCreateTime(v string) *GetAddonReleaseResponseBodyRelease {
	s.CreateTime = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetDashboardCount(v int64) *GetAddonReleaseResponseBodyRelease {
	s.DashboardCount = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetEntityRules(v *EntityGroupBase) *GetAddonReleaseResponseBodyRelease {
	s.EntityRules = v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetEnvType(v string) *GetAddonReleaseResponseBodyRelease {
	s.EnvType = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetEnvironmentId(v string) *GetAddonReleaseResponseBodyRelease {
	s.EnvironmentId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetExporterCount(v int64) *GetAddonReleaseResponseBodyRelease {
	s.ExporterCount = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetHaveConfig(v bool) *GetAddonReleaseResponseBodyRelease {
	s.HaveConfig = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetInstallUserId(v string) *GetAddonReleaseResponseBodyRelease {
	s.InstallUserId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetLanguage(v string) *GetAddonReleaseResponseBodyRelease {
	s.Language = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetManaged(v bool) *GetAddonReleaseResponseBodyRelease {
	s.Managed = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetParentAddonReleaseId(v string) *GetAddonReleaseResponseBodyRelease {
	s.ParentAddonReleaseId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetPolicyId(v string) *GetAddonReleaseResponseBodyRelease {
	s.PolicyId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetRegionId(v string) *GetAddonReleaseResponseBodyRelease {
	s.RegionId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetReleaseId(v string) *GetAddonReleaseResponseBodyRelease {
	s.ReleaseId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetReleaseName(v string) *GetAddonReleaseResponseBodyRelease {
	s.ReleaseName = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetScene(v string) *GetAddonReleaseResponseBodyRelease {
	s.Scene = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetStatus(v string) *GetAddonReleaseResponseBodyRelease {
	s.Status = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetUpdateTime(v string) *GetAddonReleaseResponseBodyRelease {
	s.UpdateTime = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetUserId(v string) *GetAddonReleaseResponseBodyRelease {
	s.UserId = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetVersion(v string) *GetAddonReleaseResponseBodyRelease {
	s.Version = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) SetWorkspace(v string) *GetAddonReleaseResponseBodyRelease {
	s.Workspace = &v
	return s
}

func (s *GetAddonReleaseResponseBodyRelease) Validate() error {
	return dara.Validate(s)
}

type GetAddonReleaseResponseBodyReleaseConditions struct {
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

func (s GetAddonReleaseResponseBodyReleaseConditions) String() string {
	return dara.Prettify(s)
}

func (s GetAddonReleaseResponseBodyReleaseConditions) GoString() string {
	return s.String()
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) GetMessage() *string {
	return s.Message
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) GetStatus() *string {
	return s.Status
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) GetType() *string {
	return s.Type
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) SetFirstTransitionTime(v string) *GetAddonReleaseResponseBodyReleaseConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) SetLastTransitionTime(v string) *GetAddonReleaseResponseBodyReleaseConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) SetMessage(v string) *GetAddonReleaseResponseBodyReleaseConditions {
	s.Message = &v
	return s
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) SetStatus(v string) *GetAddonReleaseResponseBodyReleaseConditions {
	s.Status = &v
	return s
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) SetType(v string) *GetAddonReleaseResponseBodyReleaseConditions {
	s.Type = &v
	return s
}

func (s *GetAddonReleaseResponseBodyReleaseConditions) Validate() error {
	return dara.Validate(s)
}
