// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateInstanceInspectionRequest
	GetClientToken() *string
	SetItems(v []*UpdateInstanceInspectionRequestItems) *UpdateInstanceInspectionRequest
	GetItems() []*UpdateInstanceInspectionRequestItems
	SetResourceScope(v *UpdateInstanceInspectionRequestResourceScope) *UpdateInstanceInspectionRequest
	GetResourceScope() *UpdateInstanceInspectionRequestResourceScope
	SetRevision(v int64) *UpdateInstanceInspectionRequest
	GetRevision() *int64
	SetSchedule(v *UpdateInstanceInspectionRequestSchedule) *UpdateInstanceInspectionRequest
	GetSchedule() *UpdateInstanceInspectionRequestSchedule
	SetRegionId(v string) *UpdateInstanceInspectionRequest
	GetRegionId() *string
}

type UpdateInstanceInspectionRequest struct {
	// 同一实例配置生命周期内有效 24 小时；同 token 同请求返回首次成功结果，不同请求返回冲突。新请求仍须携带当前 revision。
	//
	// This parameter is required.
	//
	// example:
	//
	// update-inspection-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*UpdateInstanceInspectionRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// {"mode":"All"}
	ResourceScope *UpdateInstanceInspectionRequestResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *UpdateInstanceInspectionRequestSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s UpdateInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateInstanceInspectionRequest) GetItems() []*UpdateInstanceInspectionRequestItems {
	return s.Items
}

func (s *UpdateInstanceInspectionRequest) GetResourceScope() *UpdateInstanceInspectionRequestResourceScope {
	return s.ResourceScope
}

func (s *UpdateInstanceInspectionRequest) GetRevision() *int64 {
	return s.Revision
}

func (s *UpdateInstanceInspectionRequest) GetSchedule() *UpdateInstanceInspectionRequestSchedule {
	return s.Schedule
}

func (s *UpdateInstanceInspectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateInstanceInspectionRequest) SetClientToken(v string) *UpdateInstanceInspectionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateInstanceInspectionRequest) SetItems(v []*UpdateInstanceInspectionRequestItems) *UpdateInstanceInspectionRequest {
	s.Items = v
	return s
}

func (s *UpdateInstanceInspectionRequest) SetResourceScope(v *UpdateInstanceInspectionRequestResourceScope) *UpdateInstanceInspectionRequest {
	s.ResourceScope = v
	return s
}

func (s *UpdateInstanceInspectionRequest) SetRevision(v int64) *UpdateInstanceInspectionRequest {
	s.Revision = &v
	return s
}

func (s *UpdateInstanceInspectionRequest) SetSchedule(v *UpdateInstanceInspectionRequestSchedule) *UpdateInstanceInspectionRequest {
	s.Schedule = v
	return s
}

func (s *UpdateInstanceInspectionRequest) SetRegionId(v string) *UpdateInstanceInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceInspectionRequest) Validate() error {
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

type UpdateInstanceInspectionRequestItems struct {
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

func (s UpdateInstanceInspectionRequestItems) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionRequestItems) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionRequestItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *UpdateInstanceInspectionRequestItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateInstanceInspectionRequestItems) SetParameters(v interface{}) *UpdateInstanceInspectionRequestItems {
	s.Parameters = v
	return s
}

func (s *UpdateInstanceInspectionRequestItems) SetTemplateCode(v string) *UpdateInstanceInspectionRequestItems {
	s.TemplateCode = &v
	return s
}

func (s *UpdateInstanceInspectionRequestItems) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionRequestResourceScope struct {
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

func (s UpdateInstanceInspectionRequestResourceScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionRequestResourceScope) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionRequestResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *UpdateInstanceInspectionRequestResourceScope) GetMode() *string {
	return s.Mode
}

func (s *UpdateInstanceInspectionRequestResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *UpdateInstanceInspectionRequestResourceScope) SetConsumerGroups(v []*string) *UpdateInstanceInspectionRequestResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *UpdateInstanceInspectionRequestResourceScope) SetMode(v string) *UpdateInstanceInspectionRequestResourceScope {
	s.Mode = &v
	return s
}

func (s *UpdateInstanceInspectionRequestResourceScope) SetTopics(v []*string) *UpdateInstanceInspectionRequestResourceScope {
	s.Topics = v
	return s
}

func (s *UpdateInstanceInspectionRequestResourceScope) Validate() error {
	return dara.Validate(s)
}

type UpdateInstanceInspectionRequestSchedule struct {
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

func (s UpdateInstanceInspectionRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceInspectionRequestSchedule) GoString() string {
	return s.String()
}

func (s *UpdateInstanceInspectionRequestSchedule) GetCron() *string {
	return s.Cron
}

func (s *UpdateInstanceInspectionRequestSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateInstanceInspectionRequestSchedule) SetCron(v string) *UpdateInstanceInspectionRequestSchedule {
	s.Cron = &v
	return s
}

func (s *UpdateInstanceInspectionRequestSchedule) SetEnabled(v bool) *UpdateInstanceInspectionRequestSchedule {
	s.Enabled = &v
	return s
}

func (s *UpdateInstanceInspectionRequestSchedule) Validate() error {
	return dara.Validate(s)
}
