// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonReleasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAddonReleasesResponseBody
	GetCode() *int32
	SetData(v *ListAddonReleasesResponseBodyData) *ListAddonReleasesResponseBody
	GetData() *ListAddonReleasesResponseBodyData
	SetMessage(v string) *ListAddonReleasesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAddonReleasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAddonReleasesResponseBody
	GetSuccess() *bool
}

type ListAddonReleasesResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result returned.
	Data *ListAddonReleasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9C9DA3D-10FE-472E-9EEF-2D0A3E41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAddonReleasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAddonReleasesResponseBody) GetData() *ListAddonReleasesResponseBodyData {
	return s.Data
}

func (s *ListAddonReleasesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAddonReleasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonReleasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAddonReleasesResponseBody) SetCode(v int32) *ListAddonReleasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAddonReleasesResponseBody) SetData(v *ListAddonReleasesResponseBodyData) *ListAddonReleasesResponseBody {
	s.Data = v
	return s
}

func (s *ListAddonReleasesResponseBody) SetMessage(v string) *ListAddonReleasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAddonReleasesResponseBody) SetRequestId(v string) *ListAddonReleasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonReleasesResponseBody) SetSuccess(v bool) *ListAddonReleasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAddonReleasesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAddonReleasesResponseBodyData struct {
	// The queried add-ons.
	Releases []*ListAddonReleasesResponseBodyDataReleases `json:"Releases,omitempty" xml:"Releases,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 12
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAddonReleasesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBodyData) GetReleases() []*ListAddonReleasesResponseBodyDataReleases {
	return s.Releases
}

func (s *ListAddonReleasesResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListAddonReleasesResponseBodyData) SetReleases(v []*ListAddonReleasesResponseBodyDataReleases) *ListAddonReleasesResponseBodyData {
	s.Releases = v
	return s
}

func (s *ListAddonReleasesResponseBodyData) SetTotal(v int64) *ListAddonReleasesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListAddonReleasesResponseBodyData) Validate() error {
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

type ListAddonReleasesResponseBodyDataReleases struct {
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
	Conditions []*ListAddonReleasesResponseBodyDataReleasesConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The configuration information of the add-on release.
	//
	// example:
	//
	// {"port":"9379"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
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
	// 1268790592387
	InstallUserId *string `json:"InstallUserId,omitempty" xml:"InstallUserId,omitempty"`
	// The language.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// Indicates whether the component is fully managed.
	//
	// example:
	//
	// false
	Managed *bool `json:"Managed,omitempty" xml:"Managed,omitempty"`
	// The latest version.
	//
	// example:
	//
	// 0.0.4
	NextVersion *string `json:"NextVersion,omitempty" xml:"NextVersion,omitempty"`
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
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version of the add-on.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAddonReleasesResponseBodyDataReleases) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBodyDataReleases) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetAddonName() *string {
	return s.AddonName
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetAlertRuleCount() *int64 {
	return s.AlertRuleCount
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetConditions() []*ListAddonReleasesResponseBodyDataReleasesConditions {
	return s.Conditions
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetConfig() *string {
	return s.Config
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetDashboardCount() *int64 {
	return s.DashboardCount
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetExporterCount() *int64 {
	return s.ExporterCount
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetHaveConfig() *bool {
	return s.HaveConfig
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetInstallUserId() *string {
	return s.InstallUserId
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetLanguage() *string {
	return s.Language
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetManaged() *bool {
	return s.Managed
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetNextVersion() *string {
	return s.NextVersion
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetReleaseId() *string {
	return s.ReleaseId
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetScene() *string {
	return s.Scene
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetStatus() *string {
	return s.Status
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetUserId() *string {
	return s.UserId
}

func (s *ListAddonReleasesResponseBodyDataReleases) GetVersion() *string {
	return s.Version
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetAddonName(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.AddonName = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetAlertRuleCount(v int64) *ListAddonReleasesResponseBodyDataReleases {
	s.AlertRuleCount = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetConditions(v []*ListAddonReleasesResponseBodyDataReleasesConditions) *ListAddonReleasesResponseBodyDataReleases {
	s.Conditions = v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetConfig(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.Config = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetCreateTime(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.CreateTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetDashboardCount(v int64) *ListAddonReleasesResponseBodyDataReleases {
	s.DashboardCount = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetEnvironmentId(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.EnvironmentId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetExporterCount(v int64) *ListAddonReleasesResponseBodyDataReleases {
	s.ExporterCount = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetHaveConfig(v bool) *ListAddonReleasesResponseBodyDataReleases {
	s.HaveConfig = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetInstallUserId(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.InstallUserId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetLanguage(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.Language = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetManaged(v bool) *ListAddonReleasesResponseBodyDataReleases {
	s.Managed = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetNextVersion(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.NextVersion = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetRegionId(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.RegionId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetReleaseId(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.ReleaseId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetReleaseName(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.ReleaseName = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetScene(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.Scene = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetStatus(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.Status = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetUpdateTime(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.UpdateTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetUserId(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.UserId = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) SetVersion(v string) *ListAddonReleasesResponseBodyDataReleases {
	s.Version = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleases) Validate() error {
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

type ListAddonReleasesResponseBodyDataReleasesConditions struct {
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
	// The detailed information.
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

func (s ListAddonReleasesResponseBodyDataReleasesConditions) String() string {
	return dara.Prettify(s)
}

func (s ListAddonReleasesResponseBodyDataReleasesConditions) GoString() string {
	return s.String()
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) GetMessage() *string {
	return s.Message
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) GetReason() *string {
	return s.Reason
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) GetStatus() *string {
	return s.Status
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) GetType() *string {
	return s.Type
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) SetFirstTransitionTime(v string) *ListAddonReleasesResponseBodyDataReleasesConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) SetLastTransitionTime(v string) *ListAddonReleasesResponseBodyDataReleasesConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) SetMessage(v string) *ListAddonReleasesResponseBodyDataReleasesConditions {
	s.Message = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) SetReason(v string) *ListAddonReleasesResponseBodyDataReleasesConditions {
	s.Reason = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) SetStatus(v string) *ListAddonReleasesResponseBodyDataReleasesConditions {
	s.Status = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) SetType(v string) *ListAddonReleasesResponseBodyDataReleasesConditions {
	s.Type = &v
	return s
}

func (s *ListAddonReleasesResponseBodyDataReleasesConditions) Validate() error {
	return dara.Validate(s)
}
