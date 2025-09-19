// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityStatInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttackEvent(v *DescribeSecurityStatInfoResponseBodyAttackEvent) *DescribeSecurityStatInfoResponseBody
	GetAttackEvent() *DescribeSecurityStatInfoResponseBodyAttackEvent
	SetHealthCheck(v *DescribeSecurityStatInfoResponseBodyHealthCheck) *DescribeSecurityStatInfoResponseBody
	GetHealthCheck() *DescribeSecurityStatInfoResponseBodyHealthCheck
	SetRequestId(v string) *DescribeSecurityStatInfoResponseBody
	GetRequestId() *string
	SetSecurityEvent(v *DescribeSecurityStatInfoResponseBodySecurityEvent) *DescribeSecurityStatInfoResponseBody
	GetSecurityEvent() *DescribeSecurityStatInfoResponseBodySecurityEvent
	SetSuccess(v bool) *DescribeSecurityStatInfoResponseBody
	GetSuccess() *bool
	SetVulnerability(v *DescribeSecurityStatInfoResponseBodyVulnerability) *DescribeSecurityStatInfoResponseBody
	GetVulnerability() *DescribeSecurityStatInfoResponseBodyVulnerability
}

type DescribeSecurityStatInfoResponseBody struct {
	// The detailed statistics of attacks.
	AttackEvent *DescribeSecurityStatInfoResponseBodyAttackEvent `json:"AttackEvent,omitempty" xml:"AttackEvent,omitempty" type:"Struct"`
	// The detailed statistics of baseline risk items.
	HealthCheck *DescribeSecurityStatInfoResponseBodyHealthCheck `json:"HealthCheck,omitempty" xml:"HealthCheck,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A3E61730-85E2-4789-8017-B9B1B70F0568
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed statistics of unhandled alerts.
	SecurityEvent *DescribeSecurityStatInfoResponseBodySecurityEvent `json:"SecurityEvent,omitempty" xml:"SecurityEvent,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The detailed statistics of unfixed vulnerabilities.
	Vulnerability *DescribeSecurityStatInfoResponseBodyVulnerability `json:"Vulnerability,omitempty" xml:"Vulnerability,omitempty" type:"Struct"`
}

func (s DescribeSecurityStatInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBody) GetAttackEvent() *DescribeSecurityStatInfoResponseBodyAttackEvent {
	return s.AttackEvent
}

func (s *DescribeSecurityStatInfoResponseBody) GetHealthCheck() *DescribeSecurityStatInfoResponseBodyHealthCheck {
	return s.HealthCheck
}

func (s *DescribeSecurityStatInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityStatInfoResponseBody) GetSecurityEvent() *DescribeSecurityStatInfoResponseBodySecurityEvent {
	return s.SecurityEvent
}

func (s *DescribeSecurityStatInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSecurityStatInfoResponseBody) GetVulnerability() *DescribeSecurityStatInfoResponseBodyVulnerability {
	return s.Vulnerability
}

func (s *DescribeSecurityStatInfoResponseBody) SetAttackEvent(v *DescribeSecurityStatInfoResponseBodyAttackEvent) *DescribeSecurityStatInfoResponseBody {
	s.AttackEvent = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetHealthCheck(v *DescribeSecurityStatInfoResponseBodyHealthCheck) *DescribeSecurityStatInfoResponseBody {
	s.HealthCheck = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetRequestId(v string) *DescribeSecurityStatInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetSecurityEvent(v *DescribeSecurityStatInfoResponseBodySecurityEvent) *DescribeSecurityStatInfoResponseBody {
	s.SecurityEvent = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetSuccess(v bool) *DescribeSecurityStatInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) SetVulnerability(v *DescribeSecurityStatInfoResponseBodyVulnerability) *DescribeSecurityStatInfoResponseBody {
	s.Vulnerability = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityStatInfoResponseBodyAttackEvent struct {
	// The points in time when the number of attacks is collected in the trend chart.
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// The total number of attacks on the current day.
	//
	// example:
	//
	// 1096
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The numbers of attacks at all points in time.
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodyAttackEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodyAttackEvent) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) GetDateArray() []*string {
	return s.DateArray
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) GetValueArray() []*string {
	return s.ValueArray
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodyAttackEvent {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodyAttackEvent {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodyAttackEvent {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyAttackEvent) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityStatInfoResponseBodyHealthCheck struct {
	// The points in time when data of baseline risk items is collected in the trend chart.
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// The number of baseline risk items that have the high-risk level on the current day.
	//
	// example:
	//
	// 10
	HighCount *int32 `json:"HighCount,omitempty" xml:"HighCount,omitempty"`
	// The numbers of baseline risk items that have the high-risk level at all points in time.
	HighList []*string `json:"HighList,omitempty" xml:"HighList,omitempty" type:"Repeated"`
	// The risk levels of baseline risk items.
	LevelsOn []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// The number of baseline risk items that have the low-risk level on the current day.
	//
	// example:
	//
	// 0
	LowCount *int32 `json:"LowCount,omitempty" xml:"LowCount,omitempty"`
	// The numbers of baseline risk items that have the low-risk level at all points in time.
	LowList []*string `json:"LowList,omitempty" xml:"LowList,omitempty" type:"Repeated"`
	// The number of baseline risk items that have the medium-risk level on the current day.
	//
	// example:
	//
	// 21
	MediumCount *int32 `json:"MediumCount,omitempty" xml:"MediumCount,omitempty"`
	// The numbers of baseline risk items that have the medium-risk level at all points in time.
	MediumList []*string `json:"MediumList,omitempty" xml:"MediumList,omitempty" type:"Repeated"`
	// The time periods during which data of baseline risk items is collected.
	TimeArray []*string `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Repeated"`
	// The total number of baseline risk items on the current day.
	//
	// example:
	//
	// 32
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of baseline risk items at all points in time.
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodyHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodyHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetDateArray() []*string {
	return s.DateArray
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetHighCount() *int32 {
	return s.HighCount
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetHighList() []*string {
	return s.HighList
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetLevelsOn() []*string {
	return s.LevelsOn
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetLowCount() *int32 {
	return s.LowCount
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetLowList() []*string {
	return s.LowList
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetMediumCount() *int32 {
	return s.MediumCount
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetMediumList() []*string {
	return s.MediumList
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetTimeArray() []*string {
	return s.TimeArray
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) GetValueArray() []*string {
	return s.ValueArray
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetHighCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.HighCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetHighList(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.HighList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetLevelsOn(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.LevelsOn = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetLowCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.LowCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetLowList(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.LowList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetMediumCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.MediumCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetMediumList(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.MediumList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetTimeArray(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.TimeArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodyHealthCheck {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyHealthCheck) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityStatInfoResponseBodySecurityEvent struct {
	// The points in time when data of unhandled alerts is collected in the trend chart.
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// The risk levels of unhandled alerts.
	LevelsOn []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// The number of **remind*	- alerts on the current day.
	//
	// example:
	//
	// 0
	RemindCount *int32 `json:"RemindCount,omitempty" xml:"RemindCount,omitempty"`
	// The numbers of remind alerts at all points in time.
	RemindList []*string `json:"RemindList,omitempty" xml:"RemindList,omitempty" type:"Repeated"`
	// The number of **serious*	- alerts on the current day.
	//
	// example:
	//
	// 404
	SeriousCount *int32 `json:"SeriousCount,omitempty" xml:"SeriousCount,omitempty"`
	// The numbers of serious alerts at all points in time.
	SeriousList []*string `json:"SeriousList,omitempty" xml:"SeriousList,omitempty" type:"Repeated"`
	// The number of **suspicious*	- alerts on the current day.
	//
	// example:
	//
	// 148
	SuspiciousCount *int32 `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	// The numbers of suspicious alerts at all points in time.
	SuspiciousList []*string `json:"SuspiciousList,omitempty" xml:"SuspiciousList,omitempty" type:"Repeated"`
	// The time periods during which data of the same alert is collected.
	TimeArray []*string `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Repeated"`
	// The total number of unhandled alerts on the current day.
	//
	// example:
	//
	// 552
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The numbers of unhandled alerts at all points in time.
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodySecurityEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodySecurityEvent) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetDateArray() []*string {
	return s.DateArray
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetLevelsOn() []*string {
	return s.LevelsOn
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetRemindCount() *int32 {
	return s.RemindCount
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetRemindList() []*string {
	return s.RemindList
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetSeriousCount() *int32 {
	return s.SeriousCount
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetSeriousList() []*string {
	return s.SeriousList
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetSuspiciousCount() *int32 {
	return s.SuspiciousCount
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetSuspiciousList() []*string {
	return s.SuspiciousList
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetTimeArray() []*string {
	return s.TimeArray
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) GetValueArray() []*string {
	return s.ValueArray
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetLevelsOn(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.LevelsOn = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetRemindCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.RemindCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetRemindList(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.RemindList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSeriousCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SeriousCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSeriousList(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SeriousList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSuspiciousCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SuspiciousCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetSuspiciousList(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.SuspiciousList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetTimeArray(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.TimeArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodySecurityEvent {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodySecurityEvent) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityStatInfoResponseBodyVulnerability struct {
	// The number of **high-risk*	- unfixed vulnerabilities on the current day.
	//
	// example:
	//
	// 109
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// The numbers of high-risk unfixed vulnerabilities at all points in time.
	AsapList []*string `json:"AsapList,omitempty" xml:"AsapList,omitempty" type:"Repeated"`
	// The points in time when data of unfixed vulnerabilities is collected in the trend chart.
	DateArray []*string `json:"DateArray,omitempty" xml:"DateArray,omitempty" type:"Repeated"`
	// The number of **medium-risk*	- unfixed vulnerabilities on the current day.
	//
	// example:
	//
	// 275
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// The numbers of medium-risk unfixed vulnerabilities at all points in time.
	LaterList []*string `json:"LaterList,omitempty" xml:"LaterList,omitempty" type:"Repeated"`
	// The risk levels of unfixed vulnerabilities.
	LevelsOn []*string `json:"LevelsOn,omitempty" xml:"LevelsOn,omitempty" type:"Repeated"`
	// The number of **low-risk*	- unfixed vulnerabilities on the current day.
	//
	// example:
	//
	// 0
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// The numbers of low-risk unfixed vulnerabilities at all points in time.
	NntfList []*string `json:"NntfList,omitempty" xml:"NntfList,omitempty" type:"Repeated"`
	// The time periods during which data of unfixed vulnerabilities is collected.
	TimeArray []*string `json:"TimeArray,omitempty" xml:"TimeArray,omitempty" type:"Repeated"`
	// The total number of unfixed vulnerabilities on the current day.
	//
	// example:
	//
	// 384
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The numbers of unfixed vulnerabilities at all points in time.
	ValueArray []*string `json:"ValueArray,omitempty" xml:"ValueArray,omitempty" type:"Repeated"`
}

func (s DescribeSecurityStatInfoResponseBodyVulnerability) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityStatInfoResponseBodyVulnerability) GoString() string {
	return s.String()
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetAsapCount() *int32 {
	return s.AsapCount
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetAsapList() []*string {
	return s.AsapList
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetDateArray() []*string {
	return s.DateArray
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetLaterCount() *int32 {
	return s.LaterCount
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetLaterList() []*string {
	return s.LaterList
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetLevelsOn() []*string {
	return s.LevelsOn
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetNntfCount() *int32 {
	return s.NntfCount
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetNntfList() []*string {
	return s.NntfList
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetTimeArray() []*string {
	return s.TimeArray
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) GetValueArray() []*string {
	return s.ValueArray
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetAsapCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.AsapCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetAsapList(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.AsapList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetDateArray(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.DateArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetLaterCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.LaterCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetLaterList(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.LaterList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetLevelsOn(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.LevelsOn = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetNntfCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.NntfCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetNntfList(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.NntfList = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetTimeArray(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.TimeArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetTotalCount(v int32) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) SetValueArray(v []*string) *DescribeSecurityStatInfoResponseBodyVulnerability {
	s.ValueArray = v
	return s
}

func (s *DescribeSecurityStatInfoResponseBodyVulnerability) Validate() error {
	return dara.Validate(s)
}
