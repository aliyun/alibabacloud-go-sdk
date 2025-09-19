// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCycleDays(v []*int32) *GetCheckConfigResponseBody
	GetCycleDays() []*int32
	SetEnableAddCheck(v bool) *GetCheckConfigResponseBody
	GetEnableAddCheck() *bool
	SetEnableAutoCheck(v bool) *GetCheckConfigResponseBody
	GetEnableAutoCheck() *bool
	SetEndTime(v int32) *GetCheckConfigResponseBody
	GetEndTime() *int32
	SetRequestId(v string) *GetCheckConfigResponseBody
	GetRequestId() *string
	SetSelectedChecks(v []*GetCheckConfigResponseBodySelectedChecks) *GetCheckConfigResponseBody
	GetSelectedChecks() []*GetCheckConfigResponseBodySelectedChecks
	SetStandards(v []*GetCheckConfigResponseBodyStandards) *GetCheckConfigResponseBody
	GetStandards() []*GetCheckConfigResponseBodyStandards
	SetStartTime(v int32) *GetCheckConfigResponseBody
	GetStartTime() *int32
}

type GetCheckConfigResponseBody struct {
	// The days in a week on which an automatic check is performed.
	CycleDays []*int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty" type:"Repeated"`
	// Indicates whether the check for new check items in the selected requirement item is enabled by default. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	EnableAddCheck *bool `json:"EnableAddCheck,omitempty" xml:"EnableAddCheck,omitempty"`
	// Indicates whether the automatic check is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableAutoCheck *bool `json:"EnableAutoCheck,omitempty" xml:"EnableAutoCheck,omitempty"`
	// The end time of the check. The value indicates a point in time. The time period that is specified by the start time and end time must be one of the following time periods:
	//
	// 	- **00:00 to 06:00**: If StartTime is set to 00:00, EndTime must be set to 06:00.
	//
	// 	- **06:00 to 12:00**: If StartTime is set to 06:00, EndTime must be set to 12:00.
	//
	// 	- **12:00 to 18:00**: If StartTime is set to 12:00, EndTime must be set to 18:00.
	//
	// 	- **18:00 to 24:00**: If StartTime is set to 18:00, EndTime must be set to 24:00.
	//
	// example:
	//
	// 6
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5A3D5C8F-2A42-5477-BDD8-27E64B5F1739
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The check items selected in the policy.
	SelectedChecks []*GetCheckConfigResponseBodySelectedChecks `json:"SelectedChecks,omitempty" xml:"SelectedChecks,omitempty" type:"Repeated"`
	// The information about the check items.
	Standards []*GetCheckConfigResponseBodyStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
	// The start time of the check. The value indicates a point in time.
	//
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetCheckConfigResponseBody) GetCycleDays() []*int32 {
	return s.CycleDays
}

func (s *GetCheckConfigResponseBody) GetEnableAddCheck() *bool {
	return s.EnableAddCheck
}

func (s *GetCheckConfigResponseBody) GetEnableAutoCheck() *bool {
	return s.EnableAutoCheck
}

func (s *GetCheckConfigResponseBody) GetEndTime() *int32 {
	return s.EndTime
}

func (s *GetCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCheckConfigResponseBody) GetSelectedChecks() []*GetCheckConfigResponseBodySelectedChecks {
	return s.SelectedChecks
}

func (s *GetCheckConfigResponseBody) GetStandards() []*GetCheckConfigResponseBodyStandards {
	return s.Standards
}

func (s *GetCheckConfigResponseBody) GetStartTime() *int32 {
	return s.StartTime
}

func (s *GetCheckConfigResponseBody) SetCycleDays(v []*int32) *GetCheckConfigResponseBody {
	s.CycleDays = v
	return s
}

func (s *GetCheckConfigResponseBody) SetEnableAddCheck(v bool) *GetCheckConfigResponseBody {
	s.EnableAddCheck = &v
	return s
}

func (s *GetCheckConfigResponseBody) SetEnableAutoCheck(v bool) *GetCheckConfigResponseBody {
	s.EnableAutoCheck = &v
	return s
}

func (s *GetCheckConfigResponseBody) SetEndTime(v int32) *GetCheckConfigResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetCheckConfigResponseBody) SetRequestId(v string) *GetCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCheckConfigResponseBody) SetSelectedChecks(v []*GetCheckConfigResponseBodySelectedChecks) *GetCheckConfigResponseBody {
	s.SelectedChecks = v
	return s
}

func (s *GetCheckConfigResponseBody) SetStandards(v []*GetCheckConfigResponseBodyStandards) *GetCheckConfigResponseBody {
	s.Standards = v
	return s
}

func (s *GetCheckConfigResponseBody) SetStartTime(v int32) *GetCheckConfigResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetCheckConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetCheckConfigResponseBodySelectedChecks struct {
	// The ID of the check item.
	//
	// example:
	//
	// 3
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The section ID of the check item.
	//
	// example:
	//
	// 69
	SectionId *int64 `json:"SectionId,omitempty" xml:"SectionId,omitempty"`
}

func (s GetCheckConfigResponseBodySelectedChecks) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConfigResponseBodySelectedChecks) GoString() string {
	return s.String()
}

func (s *GetCheckConfigResponseBodySelectedChecks) GetCheckId() *int64 {
	return s.CheckId
}

func (s *GetCheckConfigResponseBodySelectedChecks) GetSectionId() *int64 {
	return s.SectionId
}

func (s *GetCheckConfigResponseBodySelectedChecks) SetCheckId(v int64) *GetCheckConfigResponseBodySelectedChecks {
	s.CheckId = &v
	return s
}

func (s *GetCheckConfigResponseBodySelectedChecks) SetSectionId(v int64) *GetCheckConfigResponseBodySelectedChecks {
	s.SectionId = &v
	return s
}

func (s *GetCheckConfigResponseBodySelectedChecks) Validate() error {
	return dara.Validate(s)
}

type GetCheckConfigResponseBodyStandards struct {
	// The ID of the check item.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the check item.
	//
	// example:
	//
	// Cloud service configuration management
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- **ON**: The check item is enabled.
	//
	// 	- **OFF**: The check item is disabled.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the check item. Valid values:
	//
	// 	- **RISK**: cloud service configuration management
	//
	// 	- **IDENTITY_PERMISSION**: identity and permission management
	//
	// 	- **COMPLIANCE**: compliance
	//
	// example:
	//
	// RISK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCheckConfigResponseBodyStandards) String() string {
	return dara.Prettify(s)
}

func (s GetCheckConfigResponseBodyStandards) GoString() string {
	return s.String()
}

func (s *GetCheckConfigResponseBodyStandards) GetId() *int64 {
	return s.Id
}

func (s *GetCheckConfigResponseBodyStandards) GetShowName() *string {
	return s.ShowName
}

func (s *GetCheckConfigResponseBodyStandards) GetStatus() *string {
	return s.Status
}

func (s *GetCheckConfigResponseBodyStandards) GetType() *string {
	return s.Type
}

func (s *GetCheckConfigResponseBodyStandards) SetId(v int64) *GetCheckConfigResponseBodyStandards {
	s.Id = &v
	return s
}

func (s *GetCheckConfigResponseBodyStandards) SetShowName(v string) *GetCheckConfigResponseBodyStandards {
	s.ShowName = &v
	return s
}

func (s *GetCheckConfigResponseBodyStandards) SetStatus(v string) *GetCheckConfigResponseBodyStandards {
	s.Status = &v
	return s
}

func (s *GetCheckConfigResponseBodyStandards) SetType(v string) *GetCheckConfigResponseBodyStandards {
	s.Type = &v
	return s
}

func (s *GetCheckConfigResponseBodyStandards) Validate() error {
	return dara.Validate(s)
}
