// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoThrottleRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetAutoThrottleRulesResponseBody
	GetCode() *int64
	SetData(v *GetAutoThrottleRulesResponseBodyData) *GetAutoThrottleRulesResponseBody
	GetData() *GetAutoThrottleRulesResponseBodyData
	SetMessage(v string) *GetAutoThrottleRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoThrottleRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoThrottleRulesResponseBody
	GetSuccess() *bool
}

type GetAutoThrottleRulesResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetAutoThrottleRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7172BECE-588A-5961-8126-C216E16B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoThrottleRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoThrottleRulesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoThrottleRulesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetAutoThrottleRulesResponseBody) GetData() *GetAutoThrottleRulesResponseBodyData {
	return s.Data
}

func (s *GetAutoThrottleRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoThrottleRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoThrottleRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoThrottleRulesResponseBody) SetCode(v int64) *GetAutoThrottleRulesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBody) SetData(v *GetAutoThrottleRulesResponseBodyData) *GetAutoThrottleRulesResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoThrottleRulesResponseBody) SetMessage(v string) *GetAutoThrottleRulesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBody) SetRequestId(v string) *GetAutoThrottleRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBody) SetSuccess(v bool) *GetAutoThrottleRulesResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoThrottleRulesResponseBodyData struct {
	// The number of database instances for which the automatic SQL throttling feature is currently enabled.
	//
	// example:
	//
	// 1
	EnableAutoThrottleCount *int64 `json:"EnableAutoThrottleCount,omitempty" xml:"EnableAutoThrottleCount,omitempty"`
	// The database instances for which the automatic SQL throttling feature is currently enabled.
	EnableAutoThrottleList []*GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList `json:"EnableAutoThrottleList,omitempty" xml:"EnableAutoThrottleList,omitempty" type:"Repeated"`
	// The number of database instances that do not exist or for which the automatic SQL throttling feature has never been enabled.
	//
	// >  If a database instance does not exist, the instance has been released or the specified instance ID is invalid.
	//
	// example:
	//
	// 1
	NeverEnableAutoThrottleOrReleasedInstanceCount *int64 `json:"NeverEnableAutoThrottleOrReleasedInstanceCount,omitempty" xml:"NeverEnableAutoThrottleOrReleasedInstanceCount,omitempty"`
	// The number of database instances that do not exist or for which the automatic SQL throttling feature has never been enabled.
	//
	// >  If a database instance does not exist, the instance has been released or the specified instance ID is invalid.
	NeverEnableAutoThrottleOrReleasedInstanceIdList []*string `json:"NeverEnableAutoThrottleOrReleasedInstanceIdList,omitempty" xml:"NeverEnableAutoThrottleOrReleasedInstanceIdList,omitempty" type:"Repeated"`
	// The number of databases for which the automatic SQL throttling feature has been enabled.
	//
	// example:
	//
	// 3
	TotalAutoThrottleRulesCount *int64 `json:"TotalAutoThrottleRulesCount,omitempty" xml:"TotalAutoThrottleRulesCount,omitempty"`
	// The number of database instances for which the automatic SQL throttling feature was once enabled but is currently disabled.
	//
	// example:
	//
	// 1
	TurnOffAutoThrottleCount *int64 `json:"TurnOffAutoThrottleCount,omitempty" xml:"TurnOffAutoThrottleCount,omitempty"`
	// The database instances for which the automatic SQL throttling feature was once enabled but is currently disabled.
	TurnOffAutoThrottleList []*GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList `json:"TurnOffAutoThrottleList,omitempty" xml:"TurnOffAutoThrottleList,omitempty" type:"Repeated"`
}

func (s GetAutoThrottleRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoThrottleRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoThrottleRulesResponseBodyData) GetEnableAutoThrottleCount() *int64 {
	return s.EnableAutoThrottleCount
}

func (s *GetAutoThrottleRulesResponseBodyData) GetEnableAutoThrottleList() []*GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	return s.EnableAutoThrottleList
}

func (s *GetAutoThrottleRulesResponseBodyData) GetNeverEnableAutoThrottleOrReleasedInstanceCount() *int64 {
	return s.NeverEnableAutoThrottleOrReleasedInstanceCount
}

func (s *GetAutoThrottleRulesResponseBodyData) GetNeverEnableAutoThrottleOrReleasedInstanceIdList() []*string {
	return s.NeverEnableAutoThrottleOrReleasedInstanceIdList
}

func (s *GetAutoThrottleRulesResponseBodyData) GetTotalAutoThrottleRulesCount() *int64 {
	return s.TotalAutoThrottleRulesCount
}

func (s *GetAutoThrottleRulesResponseBodyData) GetTurnOffAutoThrottleCount() *int64 {
	return s.TurnOffAutoThrottleCount
}

func (s *GetAutoThrottleRulesResponseBodyData) GetTurnOffAutoThrottleList() []*GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	return s.TurnOffAutoThrottleList
}

func (s *GetAutoThrottleRulesResponseBodyData) SetEnableAutoThrottleCount(v int64) *GetAutoThrottleRulesResponseBodyData {
	s.EnableAutoThrottleCount = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) SetEnableAutoThrottleList(v []*GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) *GetAutoThrottleRulesResponseBodyData {
	s.EnableAutoThrottleList = v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) SetNeverEnableAutoThrottleOrReleasedInstanceCount(v int64) *GetAutoThrottleRulesResponseBodyData {
	s.NeverEnableAutoThrottleOrReleasedInstanceCount = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) SetNeverEnableAutoThrottleOrReleasedInstanceIdList(v []*string) *GetAutoThrottleRulesResponseBodyData {
	s.NeverEnableAutoThrottleOrReleasedInstanceIdList = v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) SetTotalAutoThrottleRulesCount(v int64) *GetAutoThrottleRulesResponseBodyData {
	s.TotalAutoThrottleRulesCount = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) SetTurnOffAutoThrottleCount(v int64) *GetAutoThrottleRulesResponseBodyData {
	s.TurnOffAutoThrottleCount = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) SetTurnOffAutoThrottleList(v []*GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) *GetAutoThrottleRulesResponseBodyData {
	s.TurnOffAutoThrottleList = v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyData) Validate() error {
	if s.EnableAutoThrottleList != nil {
		for _, item := range s.EnableAutoThrottleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TurnOffAutoThrottleList != nil {
		for _, item := range s.TurnOffAutoThrottleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList struct {
	// The maximum period of time during which an exception occurs when automatic SQL throttling is triggered. Unit: minutes.
	//
	// example:
	//
	// 2
	AbnormalDuration *float64 `json:"AbnormalDuration,omitempty" xml:"AbnormalDuration,omitempty"`
	// The maximum number of active sessions.
	//
	// example:
	//
	// 32
	ActiveSessions *int64 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The end time of the throttling window. The value of this parameter is in UTC.
	//
	// example:
	//
	// 23:59Z
	AllowThrottleEndTime *string `json:"AllowThrottleEndTime,omitempty" xml:"AllowThrottleEndTime,omitempty"`
	// The start time of the throttling window. The value of this parameter is in UTC.
	//
	// example:
	//
	// 00:00Z
	AllowThrottleStartTime *string `json:"AllowThrottleStartTime,omitempty" xml:"AllowThrottleStartTime,omitempty"`
	// Indicates whether abnormal SQL statements in execution are terminated at a time. Valid values:
	//
	// > Abnormal SQL statements use the same template as the SQL statements that need to be throttled.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoKillSession *bool `json:"AutoKillSession,omitempty" xml:"AutoKillSession,omitempty"`
	// The logical relationship between the CPU utilization threshold and the maximum number of active sessions. Valid values:
	//
	// 	- **AND**
	//
	// 	- **OR**
	//
	// example:
	//
	// AND
	CpuSessionRelation *string `json:"CpuSessionRelation,omitempty" xml:"CpuSessionRelation,omitempty"`
	// The CPU utilization threshold.
	//
	// example:
	//
	// 70
	CpuUsage *float64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum throttling duration. Unit: minutes.
	//
	// example:
	//
	// 10
	MaxThrottleTime *float64 `json:"MaxThrottleTime,omitempty" xml:"MaxThrottleTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// example:
	//
	// 140692647406****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Indicates whether the automatic SQL throttling feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GoString() string {
	return s.String()
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetAbnormalDuration() *float64 {
	return s.AbnormalDuration
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetActiveSessions() *int64 {
	return s.ActiveSessions
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetAllowThrottleEndTime() *string {
	return s.AllowThrottleEndTime
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetAllowThrottleStartTime() *string {
	return s.AllowThrottleStartTime
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetAutoKillSession() *bool {
	return s.AutoKillSession
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetCpuSessionRelation() *string {
	return s.CpuSessionRelation
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetCpuUsage() *float64 {
	return s.CpuUsage
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetMaxThrottleTime() *float64 {
	return s.MaxThrottleTime
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetUserId() *string {
	return s.UserId
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) GetVisible() *bool {
	return s.Visible
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetAbnormalDuration(v float64) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.AbnormalDuration = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetActiveSessions(v int64) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.ActiveSessions = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetAllowThrottleEndTime(v string) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.AllowThrottleEndTime = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetAllowThrottleStartTime(v string) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.AllowThrottleStartTime = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetAutoKillSession(v bool) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.AutoKillSession = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetCpuSessionRelation(v string) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.CpuSessionRelation = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetCpuUsage(v float64) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.CpuUsage = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetInstanceId(v string) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetMaxThrottleTime(v float64) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.MaxThrottleTime = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetUserId(v string) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.UserId = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) SetVisible(v bool) *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList {
	s.Visible = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataEnableAutoThrottleList) Validate() error {
	return dara.Validate(s)
}

type GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList struct {
	// The maximum period of time during which the automatic SQL throttling feature is triggered. Unit: minutes.
	//
	// example:
	//
	// 2
	AbnormalDuration *float64 `json:"AbnormalDuration,omitempty" xml:"AbnormalDuration,omitempty"`
	// The maximum number of active sessions.
	//
	// example:
	//
	// 64
	ActiveSessions *int64 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The end time of the throttling window. The value of this parameter is in UTC.
	//
	// example:
	//
	// 23:59Z
	AllowThrottleEndTime *string `json:"AllowThrottleEndTime,omitempty" xml:"AllowThrottleEndTime,omitempty"`
	// The start time of the throttling window. The value of this parameter is in UTC.
	//
	// example:
	//
	// 00:00Z
	AllowThrottleStartTime *string `json:"AllowThrottleStartTime,omitempty" xml:"AllowThrottleStartTime,omitempty"`
	// Indicates whether abnormal SQL statements in execution are terminated at a time. Valid values:
	//
	// > Abnormal SQL statements use the same template as the SQL statements that need to be throttled.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoKillSession *bool `json:"AutoKillSession,omitempty" xml:"AutoKillSession,omitempty"`
	// The logical relationship between the CPU utilization threshold and the maximum number of active sessions. Valid values:
	//
	// 	- **AND**
	//
	// 	- **OR**
	//
	// example:
	//
	// OR
	CpuSessionRelation *string `json:"CpuSessionRelation,omitempty" xml:"CpuSessionRelation,omitempty"`
	// The CPU utilization threshold.
	//
	// example:
	//
	// 80
	CpuUsage *float64 `json:"CpuUsage,omitempty" xml:"CpuUsage,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze9xrhze0709****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum throttling duration. Unit: minutes.
	//
	// example:
	//
	// 10
	MaxThrottleTime *float64 `json:"MaxThrottleTime,omitempty" xml:"MaxThrottleTime,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// example:
	//
	// 140692647406****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// Indicates whether the automatic SQL throttling feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) String() string {
	return dara.Prettify(s)
}

func (s GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GoString() string {
	return s.String()
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetAbnormalDuration() *float64 {
	return s.AbnormalDuration
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetActiveSessions() *int64 {
	return s.ActiveSessions
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetAllowThrottleEndTime() *string {
	return s.AllowThrottleEndTime
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetAllowThrottleStartTime() *string {
	return s.AllowThrottleStartTime
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetAutoKillSession() *bool {
	return s.AutoKillSession
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetCpuSessionRelation() *string {
	return s.CpuSessionRelation
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetCpuUsage() *float64 {
	return s.CpuUsage
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetMaxThrottleTime() *float64 {
	return s.MaxThrottleTime
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetUserId() *string {
	return s.UserId
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) GetVisible() *bool {
	return s.Visible
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetAbnormalDuration(v float64) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.AbnormalDuration = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetActiveSessions(v int64) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.ActiveSessions = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetAllowThrottleEndTime(v string) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.AllowThrottleEndTime = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetAllowThrottleStartTime(v string) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.AllowThrottleStartTime = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetAutoKillSession(v bool) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.AutoKillSession = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetCpuSessionRelation(v string) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.CpuSessionRelation = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetCpuUsage(v float64) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.CpuUsage = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetInstanceId(v string) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetMaxThrottleTime(v float64) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.MaxThrottleTime = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetUserId(v string) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.UserId = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) SetVisible(v bool) *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList {
	s.Visible = &v
	return s
}

func (s *GetAutoThrottleRulesResponseBodyDataTurnOffAutoThrottleList) Validate() error {
	return dara.Validate(s)
}
