// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRelease(v *CreateAddonReleaseResponseBodyRelease) *CreateAddonReleaseResponseBody
	GetRelease() *CreateAddonReleaseResponseBodyRelease
	SetRequestId(v string) *CreateAddonReleaseResponseBody
	GetRequestId() *string
}

type CreateAddonReleaseResponseBody struct {
	// The information about the integrated add-on.
	Release *CreateAddonReleaseResponseBodyRelease `json:"release,omitempty" xml:"release,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAddonReleaseResponseBody) GetRelease() *CreateAddonReleaseResponseBodyRelease {
	return s.Release
}

func (s *CreateAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAddonReleaseResponseBody) SetRelease(v *CreateAddonReleaseResponseBodyRelease) *CreateAddonReleaseResponseBody {
	s.Release = v
	return s
}

func (s *CreateAddonReleaseResponseBody) SetRequestId(v string) *CreateAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAddonReleaseResponseBody) Validate() error {
	if s.Release != nil {
		if err := s.Release.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAddonReleaseResponseBodyRelease struct {
	// The name of the integrated add-on.
	//
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// The number of alert rules.
	//
	// example:
	//
	// 6
	AlertRuleCount *int64 `json:"alertRuleCount,omitempty" xml:"alertRuleCount,omitempty"`
	// The information about the add-on installation phases.
	Conditions []*CreateAddonReleaseResponseBodyReleaseConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// The configuration of the add-on.
	//
	// example:
	//
	// {"install":{"mode":"auto-install","listenPort":"9400"},"discoverMode":"instances","discover":{"instances":"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44"},"scrapeInterval":"15","enableSecuritecs-nodeyGroupInjection":"true","metricTags":""}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// The time when the add-on was integrated.
	//
	// example:
	//
	// 2024-11-05T15:21:30+08:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The number of dashboards.
	//
	// example:
	//
	// 3
	DashboardCount *int64 `json:"dashboardCount,omitempty" xml:"dashboardCount,omitempty"`
	// The details of the entity.
	EntityRules *EntityGroupBase `json:"entityRules,omitempty" xml:"entityRules,omitempty"`
	// The environment type.
	//
	// example:
	//
	// CS
	EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// policy-xxxxxxxxxxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The number of exporters.
	//
	// example:
	//
	// 2
	ExporterCount *int64 `json:"exporterCount,omitempty" xml:"exporterCount,omitempty"`
	// Indicates whether a configuration is available.
	//
	// example:
	//
	// true
	HaveConfig *bool `json:"haveConfig,omitempty" xml:"haveConfig,omitempty"`
	// The ID of the user who installed the add-on.
	//
	// example:
	//
	// 1654218965xxxxxx
	InstallUserId *string `json:"installUserId,omitempty" xml:"installUserId,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Indicates whether the add-on is a managed add-on.
	//
	// example:
	//
	// true
	Managed *bool `json:"managed,omitempty" xml:"managed,omitempty"`
	// The ID of the parent AddonRelease.
	//
	// example:
	//
	// policy-xxxxxxxxxxx
	ParentAddonReleaseId *string `json:"parentAddonReleaseId,omitempty" xml:"parentAddonReleaseId,omitempty"`
	// The ID of the environment policy.
	//
	// example:
	//
	// policy-xxxxxxxxxx
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the release.
	//
	// example:
	//
	// 2e898e60-5e6a-46d1-a994-xxxxxxxxxx
	ReleaseId *string `json:"releaseId,omitempty" xml:"releaseId,omitempty"`
	// The name of the release.
	//
	// example:
	//
	// test-gpu-integration-name
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// The scenario of the add-on.
	//
	// example:
	//
	// 1
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The status of the add-on.
	//
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the add-on was last updated.
	//
	// example:
	//
	// 2024-09-13T02:21:02Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 165421896xxxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The version of the add-on.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The workspace.
	//
	// example:
	//
	// default
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateAddonReleaseResponseBodyRelease) String() string {
	return dara.Prettify(s)
}

func (s CreateAddonReleaseResponseBodyRelease) GoString() string {
	return s.String()
}

func (s *CreateAddonReleaseResponseBodyRelease) GetAddonName() *string {
	return s.AddonName
}

func (s *CreateAddonReleaseResponseBodyRelease) GetAlertRuleCount() *int64 {
	return s.AlertRuleCount
}

func (s *CreateAddonReleaseResponseBodyRelease) GetConditions() []*CreateAddonReleaseResponseBodyReleaseConditions {
	return s.Conditions
}

func (s *CreateAddonReleaseResponseBodyRelease) GetConfig() *string {
	return s.Config
}

func (s *CreateAddonReleaseResponseBodyRelease) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAddonReleaseResponseBodyRelease) GetDashboardCount() *int64 {
	return s.DashboardCount
}

func (s *CreateAddonReleaseResponseBodyRelease) GetEntityRules() *EntityGroupBase {
	return s.EntityRules
}

func (s *CreateAddonReleaseResponseBodyRelease) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateAddonReleaseResponseBodyRelease) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetExporterCount() *int64 {
	return s.ExporterCount
}

func (s *CreateAddonReleaseResponseBodyRelease) GetHaveConfig() *bool {
	return s.HaveConfig
}

func (s *CreateAddonReleaseResponseBodyRelease) GetInstallUserId() *string {
	return s.InstallUserId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetLanguage() *string {
	return s.Language
}

func (s *CreateAddonReleaseResponseBodyRelease) GetManaged() *bool {
	return s.Managed
}

func (s *CreateAddonReleaseResponseBodyRelease) GetParentAddonReleaseId() *string {
	return s.ParentAddonReleaseId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetReleaseId() *string {
	return s.ReleaseId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *CreateAddonReleaseResponseBodyRelease) GetScene() *string {
	return s.Scene
}

func (s *CreateAddonReleaseResponseBodyRelease) GetStatus() *string {
	return s.Status
}

func (s *CreateAddonReleaseResponseBodyRelease) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateAddonReleaseResponseBodyRelease) GetUserId() *string {
	return s.UserId
}

func (s *CreateAddonReleaseResponseBodyRelease) GetVersion() *string {
	return s.Version
}

func (s *CreateAddonReleaseResponseBodyRelease) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateAddonReleaseResponseBodyRelease) SetAddonName(v string) *CreateAddonReleaseResponseBodyRelease {
	s.AddonName = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetAlertRuleCount(v int64) *CreateAddonReleaseResponseBodyRelease {
	s.AlertRuleCount = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetConditions(v []*CreateAddonReleaseResponseBodyReleaseConditions) *CreateAddonReleaseResponseBodyRelease {
	s.Conditions = v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetConfig(v string) *CreateAddonReleaseResponseBodyRelease {
	s.Config = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetCreateTime(v string) *CreateAddonReleaseResponseBodyRelease {
	s.CreateTime = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetDashboardCount(v int64) *CreateAddonReleaseResponseBodyRelease {
	s.DashboardCount = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetEntityRules(v *EntityGroupBase) *CreateAddonReleaseResponseBodyRelease {
	s.EntityRules = v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetEnvType(v string) *CreateAddonReleaseResponseBodyRelease {
	s.EnvType = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetEnvironmentId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.EnvironmentId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetExporterCount(v int64) *CreateAddonReleaseResponseBodyRelease {
	s.ExporterCount = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetHaveConfig(v bool) *CreateAddonReleaseResponseBodyRelease {
	s.HaveConfig = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetInstallUserId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.InstallUserId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetLanguage(v string) *CreateAddonReleaseResponseBodyRelease {
	s.Language = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetManaged(v bool) *CreateAddonReleaseResponseBodyRelease {
	s.Managed = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetParentAddonReleaseId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.ParentAddonReleaseId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetPolicyId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.PolicyId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetRegionId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.RegionId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetReleaseId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.ReleaseId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetReleaseName(v string) *CreateAddonReleaseResponseBodyRelease {
	s.ReleaseName = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetScene(v string) *CreateAddonReleaseResponseBodyRelease {
	s.Scene = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetStatus(v string) *CreateAddonReleaseResponseBodyRelease {
	s.Status = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetUpdateTime(v string) *CreateAddonReleaseResponseBodyRelease {
	s.UpdateTime = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetUserId(v string) *CreateAddonReleaseResponseBodyRelease {
	s.UserId = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetVersion(v string) *CreateAddonReleaseResponseBodyRelease {
	s.Version = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) SetWorkspace(v string) *CreateAddonReleaseResponseBodyRelease {
	s.Workspace = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyRelease) Validate() error {
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
	return nil
}

type CreateAddonReleaseResponseBodyReleaseConditions struct {
	// The time when the phase first transitioned.
	//
	// example:
	//
	// 2024-11-04T16:10:22+08:00
	FirstTransitionTime *string `json:"firstTransitionTime,omitempty" xml:"firstTransitionTime,omitempty"`
	// The time when the phase last transitioned.
	//
	// example:
	//
	// 2024-11-04T16:10:22+08:00
	LastTransitionTime *string `json:"lastTransitionTime,omitempty" xml:"lastTransitionTime,omitempty"`
	// The details of the phase.
	//
	// example:
	//
	// The addon loaded successfully
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The status of the phase.
	//
	// example:
	//
	// {\\"phase\\": \\"Created\\", \\"executionDetails\\": [], \\"invocations\\": [], \\"latestExecError\\": {\\"message\\": \\"\\", \\"code\\": \\"\\", \\"requestId\\": \\"\\", \\"extraInfo\\": \\"\\", \\"title\\": \\"\\"}}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the phase.
	//
	// example:
	//
	// Loaded
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAddonReleaseResponseBodyReleaseConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateAddonReleaseResponseBodyReleaseConditions) GoString() string {
	return s.String()
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) GetMessage() *string {
	return s.Message
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) GetStatus() *string {
	return s.Status
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) GetType() *string {
	return s.Type
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) SetFirstTransitionTime(v string) *CreateAddonReleaseResponseBodyReleaseConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) SetLastTransitionTime(v string) *CreateAddonReleaseResponseBodyReleaseConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) SetMessage(v string) *CreateAddonReleaseResponseBodyReleaseConditions {
	s.Message = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) SetStatus(v string) *CreateAddonReleaseResponseBodyReleaseConditions {
	s.Status = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) SetType(v string) *CreateAddonReleaseResponseBodyReleaseConditions {
	s.Type = &v
	return s
}

func (s *CreateAddonReleaseResponseBodyReleaseConditions) Validate() error {
	return dara.Validate(s)
}
