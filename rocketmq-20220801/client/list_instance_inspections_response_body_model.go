// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceInspectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInstanceInspectionsResponseBodyData) *ListInstanceInspectionsResponseBody
	GetData() []*ListInstanceInspectionsResponseBodyData
	SetMaxResults(v int32) *ListInstanceInspectionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstanceInspectionsResponseBody
	GetNextToken() *string
}

type ListInstanceInspectionsResponseBody struct {
	// example:
	//
	// []
	Data []*ListInstanceInspectionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListInstanceInspectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBody) GetData() []*ListInstanceInspectionsResponseBodyData {
	return s.Data
}

func (s *ListInstanceInspectionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceInspectionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceInspectionsResponseBody) SetData(v []*ListInstanceInspectionsResponseBodyData) *ListInstanceInspectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListInstanceInspectionsResponseBody) SetMaxResults(v int32) *ListInstanceInspectionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceInspectionsResponseBody) SetNextToken(v string) *ListInstanceInspectionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceInspectionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceInspectionsResponseBodyData struct {
	// example:
	//
	// {}
	Config *ListInstanceInspectionsResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// true
	Configured *bool `json:"configured,omitempty" xml:"configured,omitempty"`
	// example:
	//
	// rmq-cn-example
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// []
	Items []*ListInstanceInspectionsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-09-09T03:00:00+08:00
	NextRunAt *string `json:"nextRunAt,omitempty" xml:"nextRunAt,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
}

func (s ListInstanceInspectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBodyData) GetConfig() *ListInstanceInspectionsResponseBodyDataConfig {
	return s.Config
}

func (s *ListInstanceInspectionsResponseBodyData) GetConfigured() *bool {
	return s.Configured
}

func (s *ListInstanceInspectionsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceInspectionsResponseBodyData) GetItems() []*ListInstanceInspectionsResponseBodyDataItems {
	return s.Items
}

func (s *ListInstanceInspectionsResponseBodyData) GetNextRunAt() *string {
	return s.NextRunAt
}

func (s *ListInstanceInspectionsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstanceInspectionsResponseBodyData) GetRevision() *int64 {
	return s.Revision
}

func (s *ListInstanceInspectionsResponseBodyData) SetConfig(v *ListInstanceInspectionsResponseBodyDataConfig) *ListInstanceInspectionsResponseBodyData {
	s.Config = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) SetConfigured(v bool) *ListInstanceInspectionsResponseBodyData {
	s.Configured = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) SetInstanceId(v string) *ListInstanceInspectionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) SetItems(v []*ListInstanceInspectionsResponseBodyDataItems) *ListInstanceInspectionsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) SetNextRunAt(v string) *ListInstanceInspectionsResponseBodyData {
	s.NextRunAt = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) SetRegionId(v string) *ListInstanceInspectionsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) SetRevision(v int64) *ListInstanceInspectionsResponseBodyData {
	s.Revision = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyData) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceInspectionsResponseBodyDataConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*ListInstanceInspectionsResponseBodyDataConfigItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *ListInstanceInspectionsResponseBodyDataConfigResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *ListInstanceInspectionsResponseBodyDataConfigSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
}

func (s ListInstanceInspectionsResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) GetItems() []*ListInstanceInspectionsResponseBodyDataConfigItems {
	return s.Items
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) GetResourceScope() *ListInstanceInspectionsResponseBodyDataConfigResourceScope {
	return s.ResourceScope
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) GetSchedule() *ListInstanceInspectionsResponseBodyDataConfigSchedule {
	return s.Schedule
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) SetItems(v []*ListInstanceInspectionsResponseBodyDataConfigItems) *ListInstanceInspectionsResponseBodyDataConfig {
	s.Items = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) SetResourceScope(v *ListInstanceInspectionsResponseBodyDataConfigResourceScope) *ListInstanceInspectionsResponseBodyDataConfig {
	s.ResourceScope = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) SetSchedule(v *ListInstanceInspectionsResponseBodyDataConfigSchedule) *ListInstanceInspectionsResponseBodyDataConfig {
	s.Schedule = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfig) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceScope != nil {
		if err := s.ResourceScope.Validate(); err != nil {
			return err
		}
	}
	if s.Schedule != nil {
		if err := s.Schedule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceInspectionsResponseBodyDataConfigItems struct {
	// This parameter is required.
	//
	// example:
	//
	// {"lookback":"1h","consumerDelayThresholdSeconds":600}
	Parameters interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s ListInstanceInspectionsResponseBodyDataConfigItems) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBodyDataConfigItems) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBodyDataConfigItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *ListInstanceInspectionsResponseBodyDataConfigItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListInstanceInspectionsResponseBodyDataConfigItems) SetParameters(v interface{}) *ListInstanceInspectionsResponseBodyDataConfigItems {
	s.Parameters = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigItems) SetTemplateCode(v string) *ListInstanceInspectionsResponseBodyDataConfigItems {
	s.TemplateCode = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigItems) Validate() error {
	return dara.Validate(s)
}

type ListInstanceInspectionsResponseBodyDataConfigResourceScope struct {
	// example:
	//
	// ["orders-g1"]
	ConsumerGroups []*string `json:"consumerGroups,omitempty" xml:"consumerGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// All
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// ["orders"]
	Topics []*string `json:"topics,omitempty" xml:"topics,omitempty" type:"Repeated"`
}

func (s ListInstanceInspectionsResponseBodyDataConfigResourceScope) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBodyDataConfigResourceScope) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) GetMode() *string {
	return s.Mode
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) SetConsumerGroups(v []*string) *ListInstanceInspectionsResponseBodyDataConfigResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) SetMode(v string) *ListInstanceInspectionsResponseBodyDataConfigResourceScope {
	s.Mode = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) SetTopics(v []*string) *ListInstanceInspectionsResponseBodyDataConfigResourceScope {
	s.Topics = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigResourceScope) Validate() error {
	return dara.Validate(s)
}

type ListInstanceInspectionsResponseBodyDataConfigSchedule struct {
	// example:
	//
	// 0 3 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s ListInstanceInspectionsResponseBodyDataConfigSchedule) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBodyDataConfigSchedule) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBodyDataConfigSchedule) GetCron() *string {
	return s.Cron
}

func (s *ListInstanceInspectionsResponseBodyDataConfigSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListInstanceInspectionsResponseBodyDataConfigSchedule) SetCron(v string) *ListInstanceInspectionsResponseBodyDataConfigSchedule {
	s.Cron = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigSchedule) SetEnabled(v bool) *ListInstanceInspectionsResponseBodyDataConfigSchedule {
	s.Enabled = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataConfigSchedule) Validate() error {
	return dara.Validate(s)
}

type ListInstanceInspectionsResponseBodyDataItems struct {
	// example:
	//
	// 检查消费延迟和消费健康度
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 消费健康巡检
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 3
	ResourceCount *int64 `json:"resourceCount,omitempty" xml:"resourceCount,omitempty"`
	// example:
	//
	// ["ConsumerGroup"]
	ResourceTypes []*string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
	// example:
	//
	// rmq.consumer-health
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
}

func (s ListInstanceInspectionsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceInspectionsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListInstanceInspectionsResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListInstanceInspectionsResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListInstanceInspectionsResponseBodyDataItems) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *ListInstanceInspectionsResponseBodyDataItems) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *ListInstanceInspectionsResponseBodyDataItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListInstanceInspectionsResponseBodyDataItems) SetDescription(v string) *ListInstanceInspectionsResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataItems) SetName(v string) *ListInstanceInspectionsResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataItems) SetResourceCount(v int64) *ListInstanceInspectionsResponseBodyDataItems {
	s.ResourceCount = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataItems) SetResourceTypes(v []*string) *ListInstanceInspectionsResponseBodyDataItems {
	s.ResourceTypes = v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataItems) SetTemplateCode(v string) *ListInstanceInspectionsResponseBodyDataItems {
	s.TemplateCode = &v
	return s
}

func (s *ListInstanceInspectionsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
