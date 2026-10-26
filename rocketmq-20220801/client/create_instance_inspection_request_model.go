// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInspectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateInstanceInspectionRequest
	GetClientToken() *string
	SetItems(v []*CreateInstanceInspectionRequestItems) *CreateInstanceInspectionRequest
	GetItems() []*CreateInstanceInspectionRequestItems
	SetResourceScope(v *CreateInstanceInspectionRequestResourceScope) *CreateInstanceInspectionRequest
	GetResourceScope() *CreateInstanceInspectionRequestResourceScope
	SetSchedule(v *CreateInstanceInspectionRequestSchedule) *CreateInstanceInspectionRequest
	GetSchedule() *CreateInstanceInspectionRequestSchedule
	SetRegionId(v string) *CreateInstanceInspectionRequest
	GetRegionId() *string
}

type CreateInstanceInspectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// inspection-request-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"templateCode":"rmq.consumer-health","parameters":{"lookback":"1h"}},{"templateCode":"rmq.traffic-capacity","parameters":{"lookback":"6h"}}]
	Items []*CreateInstanceInspectionRequestItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// {"mode":"All"}
	ResourceScope *CreateInstanceInspectionRequestResourceScope `json:"resourceScope,omitempty" xml:"resourceScope,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// {"enabled":true,"cron":"0 3 	- 	- *"}
	Schedule *CreateInstanceInspectionRequestSchedule `json:"schedule,omitempty" xml:"schedule,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateInstanceInspectionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceInspectionRequest) GetItems() []*CreateInstanceInspectionRequestItems {
	return s.Items
}

func (s *CreateInstanceInspectionRequest) GetResourceScope() *CreateInstanceInspectionRequestResourceScope {
	return s.ResourceScope
}

func (s *CreateInstanceInspectionRequest) GetSchedule() *CreateInstanceInspectionRequestSchedule {
	return s.Schedule
}

func (s *CreateInstanceInspectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceInspectionRequest) SetClientToken(v string) *CreateInstanceInspectionRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceInspectionRequest) SetItems(v []*CreateInstanceInspectionRequestItems) *CreateInstanceInspectionRequest {
	s.Items = v
	return s
}

func (s *CreateInstanceInspectionRequest) SetResourceScope(v *CreateInstanceInspectionRequestResourceScope) *CreateInstanceInspectionRequest {
	s.ResourceScope = v
	return s
}

func (s *CreateInstanceInspectionRequest) SetSchedule(v *CreateInstanceInspectionRequestSchedule) *CreateInstanceInspectionRequest {
	s.Schedule = v
	return s
}

func (s *CreateInstanceInspectionRequest) SetRegionId(v string) *CreateInstanceInspectionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceInspectionRequest) Validate() error {
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

type CreateInstanceInspectionRequestItems struct {
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

func (s CreateInstanceInspectionRequestItems) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionRequestItems) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionRequestItems) GetParameters() interface{} {
	return s.Parameters
}

func (s *CreateInstanceInspectionRequestItems) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateInstanceInspectionRequestItems) SetParameters(v interface{}) *CreateInstanceInspectionRequestItems {
	s.Parameters = v
	return s
}

func (s *CreateInstanceInspectionRequestItems) SetTemplateCode(v string) *CreateInstanceInspectionRequestItems {
	s.TemplateCode = &v
	return s
}

func (s *CreateInstanceInspectionRequestItems) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionRequestResourceScope struct {
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

func (s CreateInstanceInspectionRequestResourceScope) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionRequestResourceScope) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionRequestResourceScope) GetConsumerGroups() []*string {
	return s.ConsumerGroups
}

func (s *CreateInstanceInspectionRequestResourceScope) GetMode() *string {
	return s.Mode
}

func (s *CreateInstanceInspectionRequestResourceScope) GetTopics() []*string {
	return s.Topics
}

func (s *CreateInstanceInspectionRequestResourceScope) SetConsumerGroups(v []*string) *CreateInstanceInspectionRequestResourceScope {
	s.ConsumerGroups = v
	return s
}

func (s *CreateInstanceInspectionRequestResourceScope) SetMode(v string) *CreateInstanceInspectionRequestResourceScope {
	s.Mode = &v
	return s
}

func (s *CreateInstanceInspectionRequestResourceScope) SetTopics(v []*string) *CreateInstanceInspectionRequestResourceScope {
	s.Topics = v
	return s
}

func (s *CreateInstanceInspectionRequestResourceScope) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceInspectionRequestSchedule struct {
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

func (s CreateInstanceInspectionRequestSchedule) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInspectionRequestSchedule) GoString() string {
	return s.String()
}

func (s *CreateInstanceInspectionRequestSchedule) GetCron() *string {
	return s.Cron
}

func (s *CreateInstanceInspectionRequestSchedule) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateInstanceInspectionRequestSchedule) SetCron(v string) *CreateInstanceInspectionRequestSchedule {
	s.Cron = &v
	return s
}

func (s *CreateInstanceInspectionRequestSchedule) SetEnabled(v bool) *CreateInstanceInspectionRequestSchedule {
	s.Enabled = &v
	return s
}

func (s *CreateInstanceInspectionRequestSchedule) Validate() error {
	return dara.Validate(s)
}
