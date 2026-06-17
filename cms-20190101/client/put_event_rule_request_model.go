// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEventRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PutEventRuleRequest
	GetDescription() *string
	SetEventPattern(v []*PutEventRuleRequestEventPattern) *PutEventRuleRequest
	GetEventPattern() []*PutEventRuleRequestEventPattern
	SetEventType(v string) *PutEventRuleRequest
	GetEventType() *string
	SetGroupId(v string) *PutEventRuleRequest
	GetGroupId() *string
	SetRegionId(v string) *PutEventRuleRequest
	GetRegionId() *string
	SetRuleName(v string) *PutEventRuleRequest
	GetRuleName() *string
	SetSilenceTime(v int64) *PutEventRuleRequest
	GetSilenceTime() *int64
	SetState(v string) *PutEventRuleRequest
	GetState() *string
}

type PutEventRuleRequest struct {
	// The description of the Event-triggered Alert Rule.
	//
	// example:
	//
	// Event alert test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The pattern of the Event-triggered Alert Rule.
	//
	// This parameter is required.
	EventPattern []*PutEventRuleRequestEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
	// The type of the Event-triggered Alert Rule. Valid values:
	//
	// - SYSTEM: system event.
	//
	// - CUSTOM: custom event.
	//
	// example:
	//
	// SYSTEM
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the application group to which the Event-triggered Alert Rule belongs.
	//
	// example:
	//
	// 7378****
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the Event-triggered Alert Rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// myRuleName
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period. Unit: seconds.
	//
	// example:
	//
	// 86400
	SilenceTime *int64 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The status of the Event-triggered Alert Rule. Valid values:
	//
	// - ENABLED: enabled.
	//
	// - DISABLED: disabled.
	//
	// example:
	//
	// ENABLED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s PutEventRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleRequest) GoString() string {
	return s.String()
}

func (s *PutEventRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *PutEventRuleRequest) GetEventPattern() []*PutEventRuleRequestEventPattern {
	return s.EventPattern
}

func (s *PutEventRuleRequest) GetEventType() *string {
	return s.EventType
}

func (s *PutEventRuleRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *PutEventRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutEventRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PutEventRuleRequest) GetSilenceTime() *int64 {
	return s.SilenceTime
}

func (s *PutEventRuleRequest) GetState() *string {
	return s.State
}

func (s *PutEventRuleRequest) SetDescription(v string) *PutEventRuleRequest {
	s.Description = &v
	return s
}

func (s *PutEventRuleRequest) SetEventPattern(v []*PutEventRuleRequestEventPattern) *PutEventRuleRequest {
	s.EventPattern = v
	return s
}

func (s *PutEventRuleRequest) SetEventType(v string) *PutEventRuleRequest {
	s.EventType = &v
	return s
}

func (s *PutEventRuleRequest) SetGroupId(v string) *PutEventRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutEventRuleRequest) SetRegionId(v string) *PutEventRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PutEventRuleRequest) SetRuleName(v string) *PutEventRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PutEventRuleRequest) SetSilenceTime(v int64) *PutEventRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *PutEventRuleRequest) SetState(v string) *PutEventRuleRequest {
	s.State = &v
	return s
}

func (s *PutEventRuleRequest) Validate() error {
	if s.EventPattern != nil {
		for _, item := range s.EventPattern {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutEventRuleRequestEventPattern struct {
	// The keyword for event filtering. When the event content contains this keyword, an alert is automatically triggered.
	//
	// example:
	//
	// Stopping
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// The type of the Event-triggered Alert Rule. Valid values of N: 1 to 50. Valid values:
	//
	// - StatusNotification: fault notification.
	//
	// - Exception: exception.
	//
	// - Maintenance: O&M.
	//
	// - \\*: unlimited.
	//
	// example:
	//
	// Exception
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
	// The level of the Event-triggered Alert Rule. Valid values of N: 1 to 50. Valid values:
	//
	// - CRITICAL: critical.
	//
	// - WARN: warning.
	//
	// - INFO: information.
	//
	// - \\*: all levels.
	//
	// example:
	//
	// CRITICAL
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
	// The name of the Event-triggered Alert Rule. Valid values of N: 1 to 50.
	//
	// example:
	//
	// Agent_Status_Stopped
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
	// The Alibaba Cloud service type. Valid values of N: 1 to 50.
	//
	// >For information about the Alibaba Cloud services supported by Event-triggered Alert Rules, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The SQL filter for events. When the event content meets the SQL condition, an alert is automatically triggered.
	//
	// > The syntax of the SQL event filter is consistent with the query syntax of Simple Log Service (SLS).
	//
	// example:
	//
	// 192.168.XX.XX and Executed
	SQLFilter *string `json:"SQLFilter,omitempty" xml:"SQLFilter,omitempty"`
	// The status of the Event-triggered Alert Rule. Valid values of N: 1 to 50.
	//
	// example:
	//
	// Failed
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s PutEventRuleRequestEventPattern) String() string {
	return dara.Prettify(s)
}

func (s PutEventRuleRequestEventPattern) GoString() string {
	return s.String()
}

func (s *PutEventRuleRequestEventPattern) GetCustomFilters() *string {
	return s.CustomFilters
}

func (s *PutEventRuleRequestEventPattern) GetEventTypeList() []*string {
	return s.EventTypeList
}

func (s *PutEventRuleRequestEventPattern) GetLevelList() []*string {
	return s.LevelList
}

func (s *PutEventRuleRequestEventPattern) GetNameList() []*string {
	return s.NameList
}

func (s *PutEventRuleRequestEventPattern) GetProduct() *string {
	return s.Product
}

func (s *PutEventRuleRequestEventPattern) GetSQLFilter() *string {
	return s.SQLFilter
}

func (s *PutEventRuleRequestEventPattern) GetStatusList() []*string {
	return s.StatusList
}

func (s *PutEventRuleRequestEventPattern) SetCustomFilters(v string) *PutEventRuleRequestEventPattern {
	s.CustomFilters = &v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetEventTypeList(v []*string) *PutEventRuleRequestEventPattern {
	s.EventTypeList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetLevelList(v []*string) *PutEventRuleRequestEventPattern {
	s.LevelList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetNameList(v []*string) *PutEventRuleRequestEventPattern {
	s.NameList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetProduct(v string) *PutEventRuleRequestEventPattern {
	s.Product = &v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetSQLFilter(v string) *PutEventRuleRequestEventPattern {
	s.SQLFilter = &v
	return s
}

func (s *PutEventRuleRequestEventPattern) SetStatusList(v []*string) *PutEventRuleRequestEventPattern {
	s.StatusList = v
	return s
}

func (s *PutEventRuleRequestEventPattern) Validate() error {
	return dara.Validate(s)
}
