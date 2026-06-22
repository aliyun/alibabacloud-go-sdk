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
	// The periodic check schedule.
	CycleDays []*int32 `json:"CycleDays,omitempty" xml:"CycleDays,omitempty" type:"Repeated"`
	// Indicates whether new check items added to the selected standards are checked by default. Valid values:
	//
	// - **true:*	- Enabled.
	//
	// - **false:*	- Disabled.
	//
	// example:
	//
	// false
	EnableAddCheck *bool `json:"EnableAddCheck,omitempty" xml:"EnableAddCheck,omitempty"`
	// Indicates whether automatic periodic checks are enabled. Valid values:
	//
	// - **true:*	- Enabled.
	//
	// - **false:*	- Disabled.
	//
	// example:
	//
	// true
	EnableAutoCheck *bool `json:"EnableAutoCheck,omitempty" xml:"EnableAutoCheck,omitempty"`
	// The end time of the check period, indicating the hour of the day. The start time and end time must fall within one of the following time ranges:
	//
	// - **0~6**: If the start time is 0, the end time must be set to 6.
	//
	// - **6~12**: If the start time is 6, the end time must be set to 12.
	//
	// - **12~18**: If the start time is 12, the end time must be set to 18.
	//
	// - **18~24**: If the start time is 18, the end time must be set to 24.
	//
	// example:
	//
	// 6
	EndTime *int32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// 5A3D5C8F-2A42-5477-BDD8-27E64B5F1739
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The check items selected in the policy.
	SelectedChecks []*GetCheckConfigResponseBodySelectedChecks `json:"SelectedChecks,omitempty" xml:"SelectedChecks,omitempty" type:"Repeated"`
	// The list of check item information.
	Standards []*GetCheckConfigResponseBodyStandards `json:"Standards,omitempty" xml:"Standards,omitempty" type:"Repeated"`
	// The start time of the check period, indicating the hour of the day.
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
	if s.SelectedChecks != nil {
		for _, item := range s.SelectedChecks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Standards != nil {
		for _, item := range s.Standards {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// 云产品配置管理
	ShowName *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	// The enabling status of the check item. Valid values:
	//
	// - **ON:*	- Enabled.
	//
	// - **OFF:*	- Shutdown.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the check item. Valid values:
	//
	// - **RISK:*	- cloud service configuration management
	//
	// - **IDENTITY_PERMISSION:*	- permission management
	//
	// - **COMPLIANCE:*	- compliance.
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
