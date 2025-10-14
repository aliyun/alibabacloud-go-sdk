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
	// The description of the event-triggered alert rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	EventPattern []*PutEventRuleRequestEventPattern `json:"EventPattern,omitempty" xml:"EventPattern,omitempty" type:"Repeated"`
	// The type of the event-triggered alert rule. Valid values:
	//
	// 	- SYSTEM: system event-triggered alert rule
	//
	// 	- CUSTOM: custom event-triggered alert rule
	//
	// example:
	//
	// SYSTEM
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the application group to which the event-triggered alert rule belongs.
	//
	// example:
	//
	// 7378****
	GroupId  *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the event-triggered alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// myRuleName
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The mute period during which new alerts are not sent even if the trigger conditions are met. Unit: seconds.
	//
	// example:
	//
	// 86400
	SilenceTime *int64 `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	// The status of the event-triggered alert rule. Valid values:
	//
	// 	- ENABLED: enabled
	//
	// 	- DISABLED: disabled
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
	// The keyword that is used to filter events. If the content of an event contains the specified keyword, an alert is automatically triggered.
	//
	// example:
	//
	// Stopping
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// example:
	//
	// Exception
	EventTypeList []*string `json:"EventTypeList,omitempty" xml:"EventTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// CRITICAL
	LevelList []*string `json:"LevelList,omitempty" xml:"LevelList,omitempty" type:"Repeated"`
	// example:
	//
	// Agent_Status_Stopped
	NameList []*string `json:"NameList,omitempty" xml:"NameList,omitempty" type:"Repeated"`
	// The type of the cloud service. Valid values of N: 1 to 50.
	//
	// >  You can call the DescribeSystemEventMetaList operation to query the cloud services that support event-triggered alerts. For more information, see [DescribeSystemEventMetaList](https://help.aliyun.com/document_detail/114972.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The SQL condition that is used to filter events. If the content of an event meets the specified SQL condition, an alert is automatically triggered.
	//
	// >  The syntax of SQL event filtering is consistent with the query syntax of Log Service.
	//
	// example:
	//
	// 192.168.XX.XX and Executed
	SQLFilter *string `json:"SQLFilter,omitempty" xml:"SQLFilter,omitempty"`
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
