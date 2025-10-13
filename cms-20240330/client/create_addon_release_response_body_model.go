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
	Release *CreateAddonReleaseResponseBodyRelease `json:"release,omitempty" xml:"release,omitempty" type:"Struct"`
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
	// example:
	//
	// cs-gpu
	AddonName *string `json:"addonName,omitempty" xml:"addonName,omitempty"`
	// example:
	//
	// 6
	AlertRuleCount *int64                                             `json:"alertRuleCount,omitempty" xml:"alertRuleCount,omitempty"`
	Conditions     []*CreateAddonReleaseResponseBodyReleaseConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	// example:
	//
	// {"install":{"mode":"auto-install","listenPort":"9400"},"discoverMode":"instances","discover":{"instances":"worker-k8s-for-cs-c126d87c76218487e83ab322017f11b44"},"scrapeInterval":"15","enableSecuritecs-nodeyGroupInjection":"true","metricTags":""}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// 2024-11-05T15:21:30+08:00
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
	// 1654218965xxxxxx
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
	// policy-xxxxxxxxxxx
	ParentAddonReleaseId *string `json:"parentAddonReleaseId,omitempty" xml:"parentAddonReleaseId,omitempty"`
	// example:
	//
	// policy-xxxxxxxxxx
	PolicyId *string `json:"policyId,omitempty" xml:"policyId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 2e898e60-5e6a-46d1-a994-xxxxxxxxxx
	ReleaseId *string `json:"releaseId,omitempty" xml:"releaseId,omitempty"`
	// example:
	//
	// test-gpu-integration-name
	ReleaseName *string `json:"releaseName,omitempty" xml:"releaseName,omitempty"`
	// example:
	//
	// 1
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2024-09-13T02:21:02Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// 165421896xxxxxx
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
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
	// {\\"phase\\": \\"Created\\", \\"executionDetails\\": [], \\"invocations\\": [], \\"latestExecError\\": {\\"message\\": \\"\\", \\"code\\": \\"\\", \\"requestId\\": \\"\\", \\"extraInfo\\": \\"\\", \\"title\\": \\"\\"}}
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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
