// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAddonReleaseResponseBody
	GetCode() *string
	SetData(v *DescribeAddonReleaseResponseBodyData) *DescribeAddonReleaseResponseBody
	GetData() *DescribeAddonReleaseResponseBodyData
	SetMessage(v string) *DescribeAddonReleaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAddonReleaseResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeAddonReleaseResponseBody
	GetSuccess() *string
}

type DescribeAddonReleaseResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The release information.
	Data *DescribeAddonReleaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of the request.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAddonReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonReleaseResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAddonReleaseResponseBody) GetData() *DescribeAddonReleaseResponseBodyData {
	return s.Data
}

func (s *DescribeAddonReleaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAddonReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAddonReleaseResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeAddonReleaseResponseBody) SetCode(v string) *DescribeAddonReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAddonReleaseResponseBody) SetData(v *DescribeAddonReleaseResponseBodyData) *DescribeAddonReleaseResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAddonReleaseResponseBody) SetMessage(v string) *DescribeAddonReleaseResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAddonReleaseResponseBody) SetRequestId(v string) *DescribeAddonReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddonReleaseResponseBody) SetSuccess(v string) *DescribeAddonReleaseResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAddonReleaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAddonReleaseResponseBodyData struct {
	// The configuration information of the add-on release.
	//
	// example:
	//
	// {"host":"mysql-service.default","port":3306,"username":"root","password":"roots"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The release information.
	Release *DescribeAddonReleaseResponseBodyDataRelease `json:"Release,omitempty" xml:"Release,omitempty" type:"Struct"`
}

func (s DescribeAddonReleaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonReleaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAddonReleaseResponseBodyData) GetConfig() *string {
	return s.Config
}

func (s *DescribeAddonReleaseResponseBodyData) GetRelease() *DescribeAddonReleaseResponseBodyDataRelease {
	return s.Release
}

func (s *DescribeAddonReleaseResponseBodyData) SetConfig(v string) *DescribeAddonReleaseResponseBodyData {
	s.Config = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyData) SetRelease(v *DescribeAddonReleaseResponseBodyDataRelease) *DescribeAddonReleaseResponseBodyData {
	s.Release = v
	return s
}

func (s *DescribeAddonReleaseResponseBodyData) Validate() error {
	if s.Release != nil {
		if err := s.Release.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAddonReleaseResponseBodyDataRelease struct {
	// The name of the add-on.
	//
	// example:
	//
	// mysql
	AddonName *string `json:"AddonName,omitempty" xml:"AddonName,omitempty"`
	// The number of alert rules.
	//
	// example:
	//
	// 1
	AlertRuleCount *int64 `json:"AlertRuleCount,omitempty" xml:"AlertRuleCount,omitempty"`
	// The installation phase.
	Conditions []*DescribeAddonReleaseResponseBodyDataReleaseConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The time when the add-on was created.
	//
	// example:
	//
	// 2023-09-22T16:56:29+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of dashboards.
	//
	// example:
	//
	// 1
	DashboardCount *int64 `json:"DashboardCount,omitempty" xml:"DashboardCount,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-xxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The number of exporters.
	//
	// example:
	//
	// 2
	ExporterCount *int64 `json:"ExporterCount,omitempty" xml:"ExporterCount,omitempty"`
	// Indicates whether the configuration is available.
	//
	// example:
	//
	// true
	HaveConfig *bool `json:"HaveConfig,omitempty" xml:"HaveConfig,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 23810923891
	InstallUserId *string `json:"InstallUserId,omitempty" xml:"InstallUserId,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Indicates whether the component is fully managed.
	Managed *bool `json:"Managed,omitempty" xml:"Managed,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The release ID after installation.
	//
	// example:
	//
	// be29c093-3fd6-4fb6-9430-797030cc183a
	ReleaseId *string `json:"ReleaseId,omitempty" xml:"ReleaseId,omitempty"`
	// The name of the release.
	//
	// example:
	//
	// mysql-1695372983039
	ReleaseName *string `json:"ReleaseName,omitempty" xml:"ReleaseName,omitempty"`
	// The scenario.
	//
	// example:
	//
	// database
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// The status.
	//
	// example:
	//
	// 0
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the add-on was updated.
	//
	// example:
	//
	// 2023-09-22T16:56:29+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 13818734031
	UserID *string `json:"UserID,omitempty" xml:"UserID,omitempty"`
	// The version of the add-on.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAddonReleaseResponseBodyDataRelease) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonReleaseResponseBodyDataRelease) GoString() string {
	return s.String()
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetAddonName() *string {
	return s.AddonName
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetAlertRuleCount() *int64 {
	return s.AlertRuleCount
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetConditions() []*DescribeAddonReleaseResponseBodyDataReleaseConditions {
	return s.Conditions
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetDashboardCount() *int64 {
	return s.DashboardCount
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetExporterCount() *int64 {
	return s.ExporterCount
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetHaveConfig() *bool {
	return s.HaveConfig
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetInstallUserId() *string {
	return s.InstallUserId
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetLanguage() *string {
	return s.Language
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetManaged() *bool {
	return s.Managed
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetReleaseId() *string {
	return s.ReleaseId
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetScene() *string {
	return s.Scene
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetStatus() *string {
	return s.Status
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetUserID() *string {
	return s.UserID
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) GetVersion() *string {
	return s.Version
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetAddonName(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.AddonName = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetAlertRuleCount(v int64) *DescribeAddonReleaseResponseBodyDataRelease {
	s.AlertRuleCount = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetConditions(v []*DescribeAddonReleaseResponseBodyDataReleaseConditions) *DescribeAddonReleaseResponseBodyDataRelease {
	s.Conditions = v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetCreateTime(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.CreateTime = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetDashboardCount(v int64) *DescribeAddonReleaseResponseBodyDataRelease {
	s.DashboardCount = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetEnvironmentId(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetExporterCount(v int64) *DescribeAddonReleaseResponseBodyDataRelease {
	s.ExporterCount = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetHaveConfig(v bool) *DescribeAddonReleaseResponseBodyDataRelease {
	s.HaveConfig = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetInstallUserId(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.InstallUserId = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetLanguage(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.Language = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetManaged(v bool) *DescribeAddonReleaseResponseBodyDataRelease {
	s.Managed = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetRegionId(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.RegionId = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetReleaseId(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.ReleaseId = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetReleaseName(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.ReleaseName = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetScene(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.Scene = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetStatus(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.Status = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetUpdateTime(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.UpdateTime = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetUserID(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.UserID = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) SetVersion(v string) *DescribeAddonReleaseResponseBodyDataRelease {
	s.Version = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataRelease) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAddonReleaseResponseBodyDataReleaseConditions struct {
	// The first transition time.
	//
	// example:
	//
	// 2018-01-31T14:32:19Z
	FirstTransitionTime *string `json:"FirstTransitionTime,omitempty" xml:"FirstTransitionTime,omitempty"`
	// The last transition time.
	//
	// example:
	//
	// 2018-01-31T14:32:19Z
	LastTransitionTime *string `json:"LastTransitionTime,omitempty" xml:"LastTransitionTime,omitempty"`
	// The returned message.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// xxxx
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the phase.
	//
	// example:
	//
	// True
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the phase.
	//
	// example:
	//
	// Loaded
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAddonReleaseResponseBodyDataReleaseConditions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonReleaseResponseBodyDataReleaseConditions) GoString() string {
	return s.String()
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) GetMessage() *string {
	return s.Message
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) GetReason() *string {
	return s.Reason
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) GetStatus() *string {
	return s.Status
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) GetType() *string {
	return s.Type
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) SetFirstTransitionTime(v string) *DescribeAddonReleaseResponseBodyDataReleaseConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) SetLastTransitionTime(v string) *DescribeAddonReleaseResponseBodyDataReleaseConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) SetMessage(v string) *DescribeAddonReleaseResponseBodyDataReleaseConditions {
	s.Message = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) SetReason(v string) *DescribeAddonReleaseResponseBodyDataReleaseConditions {
	s.Reason = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) SetStatus(v string) *DescribeAddonReleaseResponseBodyDataReleaseConditions {
	s.Status = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) SetType(v string) *DescribeAddonReleaseResponseBodyDataReleaseConditions {
	s.Type = &v
	return s
}

func (s *DescribeAddonReleaseResponseBodyDataReleaseConditions) Validate() error {
	return dara.Validate(s)
}
