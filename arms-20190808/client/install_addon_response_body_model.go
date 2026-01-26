// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAddonResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InstallAddonResponseBody
	GetCode() *int32
	SetData(v *InstallAddonResponseBodyData) *InstallAddonResponseBody
	GetData() *InstallAddonResponseBodyData
	SetMessage(v string) *InstallAddonResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallAddonResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallAddonResponseBody
	GetSuccess() *bool
}

type InstallAddonResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *InstallAddonResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallAddonResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonResponseBody) GoString() string {
	return s.String()
}

func (s *InstallAddonResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstallAddonResponseBody) GetData() *InstallAddonResponseBodyData {
	return s.Data
}

func (s *InstallAddonResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallAddonResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallAddonResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallAddonResponseBody) SetCode(v int32) *InstallAddonResponseBody {
	s.Code = &v
	return s
}

func (s *InstallAddonResponseBody) SetData(v *InstallAddonResponseBodyData) *InstallAddonResponseBody {
	s.Data = v
	return s
}

func (s *InstallAddonResponseBody) SetMessage(v string) *InstallAddonResponseBody {
	s.Message = &v
	return s
}

func (s *InstallAddonResponseBody) SetRequestId(v string) *InstallAddonResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallAddonResponseBody) SetSuccess(v bool) *InstallAddonResponseBody {
	s.Success = &v
	return s
}

func (s *InstallAddonResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallAddonResponseBodyData struct {
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
	Conditions []*InstallAddonResponseBodyDataConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
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
	// 123456
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
	// true
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
	// 111
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version of the add-on.
	//
	// example:
	//
	// 0.0.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s InstallAddonResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonResponseBodyData) GoString() string {
	return s.String()
}

func (s *InstallAddonResponseBodyData) GetAddonName() *string {
	return s.AddonName
}

func (s *InstallAddonResponseBodyData) GetAlertRuleCount() *int64 {
	return s.AlertRuleCount
}

func (s *InstallAddonResponseBodyData) GetConditions() []*InstallAddonResponseBodyDataConditions {
	return s.Conditions
}

func (s *InstallAddonResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *InstallAddonResponseBodyData) GetDashboardCount() *int64 {
	return s.DashboardCount
}

func (s *InstallAddonResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *InstallAddonResponseBodyData) GetExporterCount() *int64 {
	return s.ExporterCount
}

func (s *InstallAddonResponseBodyData) GetHaveConfig() *bool {
	return s.HaveConfig
}

func (s *InstallAddonResponseBodyData) GetInstallUserId() *string {
	return s.InstallUserId
}

func (s *InstallAddonResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *InstallAddonResponseBodyData) GetManaged() *bool {
	return s.Managed
}

func (s *InstallAddonResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallAddonResponseBodyData) GetReleaseId() *string {
	return s.ReleaseId
}

func (s *InstallAddonResponseBodyData) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *InstallAddonResponseBodyData) GetScene() *string {
	return s.Scene
}

func (s *InstallAddonResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *InstallAddonResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *InstallAddonResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *InstallAddonResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *InstallAddonResponseBodyData) SetAddonName(v string) *InstallAddonResponseBodyData {
	s.AddonName = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetAlertRuleCount(v int64) *InstallAddonResponseBodyData {
	s.AlertRuleCount = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetConditions(v []*InstallAddonResponseBodyDataConditions) *InstallAddonResponseBodyData {
	s.Conditions = v
	return s
}

func (s *InstallAddonResponseBodyData) SetCreateTime(v string) *InstallAddonResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetDashboardCount(v int64) *InstallAddonResponseBodyData {
	s.DashboardCount = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetEnvironmentId(v string) *InstallAddonResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetExporterCount(v int64) *InstallAddonResponseBodyData {
	s.ExporterCount = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetHaveConfig(v bool) *InstallAddonResponseBodyData {
	s.HaveConfig = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetInstallUserId(v string) *InstallAddonResponseBodyData {
	s.InstallUserId = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetLanguage(v string) *InstallAddonResponseBodyData {
	s.Language = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetManaged(v bool) *InstallAddonResponseBodyData {
	s.Managed = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetRegionId(v string) *InstallAddonResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetReleaseId(v string) *InstallAddonResponseBodyData {
	s.ReleaseId = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetReleaseName(v string) *InstallAddonResponseBodyData {
	s.ReleaseName = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetScene(v string) *InstallAddonResponseBodyData {
	s.Scene = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetStatus(v string) *InstallAddonResponseBodyData {
	s.Status = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetUpdateTime(v string) *InstallAddonResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetUserId(v string) *InstallAddonResponseBodyData {
	s.UserId = &v
	return s
}

func (s *InstallAddonResponseBodyData) SetVersion(v string) *InstallAddonResponseBodyData {
	s.Version = &v
	return s
}

func (s *InstallAddonResponseBodyData) Validate() error {
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

type InstallAddonResponseBodyDataConditions struct {
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

func (s InstallAddonResponseBodyDataConditions) String() string {
	return dara.Prettify(s)
}

func (s InstallAddonResponseBodyDataConditions) GoString() string {
	return s.String()
}

func (s *InstallAddonResponseBodyDataConditions) GetFirstTransitionTime() *string {
	return s.FirstTransitionTime
}

func (s *InstallAddonResponseBodyDataConditions) GetLastTransitionTime() *string {
	return s.LastTransitionTime
}

func (s *InstallAddonResponseBodyDataConditions) GetMessage() *string {
	return s.Message
}

func (s *InstallAddonResponseBodyDataConditions) GetReason() *string {
	return s.Reason
}

func (s *InstallAddonResponseBodyDataConditions) GetStatus() *string {
	return s.Status
}

func (s *InstallAddonResponseBodyDataConditions) GetType() *string {
	return s.Type
}

func (s *InstallAddonResponseBodyDataConditions) SetFirstTransitionTime(v string) *InstallAddonResponseBodyDataConditions {
	s.FirstTransitionTime = &v
	return s
}

func (s *InstallAddonResponseBodyDataConditions) SetLastTransitionTime(v string) *InstallAddonResponseBodyDataConditions {
	s.LastTransitionTime = &v
	return s
}

func (s *InstallAddonResponseBodyDataConditions) SetMessage(v string) *InstallAddonResponseBodyDataConditions {
	s.Message = &v
	return s
}

func (s *InstallAddonResponseBodyDataConditions) SetReason(v string) *InstallAddonResponseBodyDataConditions {
	s.Reason = &v
	return s
}

func (s *InstallAddonResponseBodyDataConditions) SetStatus(v string) *InstallAddonResponseBodyDataConditions {
	s.Status = &v
	return s
}

func (s *InstallAddonResponseBodyDataConditions) SetType(v string) *InstallAddonResponseBodyDataConditions {
	s.Type = &v
	return s
}

func (s *InstallAddonResponseBodyDataConditions) Validate() error {
	return dara.Validate(s)
}
