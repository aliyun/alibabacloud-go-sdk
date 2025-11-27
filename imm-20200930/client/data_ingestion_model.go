// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataIngestion interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*DataIngestionActions) *DataIngestion
	GetActions() []*DataIngestionActions
	SetCreateTime(v string) *DataIngestion
	GetCreateTime() *string
	SetError(v string) *DataIngestion
	GetError() *string
	SetId(v string) *DataIngestion
	GetId() *string
	SetInput(v *Input) *DataIngestion
	GetInput() *Input
	SetMarker(v string) *DataIngestion
	GetMarker() *string
	SetNotification(v *DataIngestionNotification) *DataIngestion
	GetNotification() *DataIngestionNotification
	SetPhase(v string) *DataIngestion
	GetPhase() *string
	SetServiceRole(v string) *DataIngestion
	GetServiceRole() *string
	SetState(v string) *DataIngestion
	GetState() *string
	SetStatistic(v *DataIngestionStatistic) *DataIngestion
	GetStatistic() *DataIngestionStatistic
	SetTags(v map[string]interface{}) *DataIngestion
	GetTags() map[string]interface{}
	SetUpdateTime(v string) *DataIngestion
	GetUpdateTime() *string
}

type DataIngestion struct {
	Actions      []*DataIngestionActions    `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	CreateTime   *string                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Error        *string                    `json:"Error,omitempty" xml:"Error,omitempty"`
	Id           *string                    `json:"Id,omitempty" xml:"Id,omitempty"`
	Input        *Input                     `json:"Input,omitempty" xml:"Input,omitempty"`
	Marker       *string                    `json:"Marker,omitempty" xml:"Marker,omitempty"`
	Notification *DataIngestionNotification `json:"Notification,omitempty" xml:"Notification,omitempty" type:"Struct"`
	// example:
	//
	// IncrementalScanning
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// example:
	//
	// AliyunIMMBatchTriggerRole
	ServiceRole *string                 `json:"ServiceRole,omitempty" xml:"ServiceRole,omitempty"`
	State       *string                 `json:"State,omitempty" xml:"State,omitempty"`
	Statistic   *DataIngestionStatistic `json:"Statistic,omitempty" xml:"Statistic,omitempty" type:"Struct"`
	Tags        map[string]interface{}  `json:"Tags,omitempty" xml:"Tags,omitempty"`
	UpdateTime  *string                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DataIngestion) String() string {
	return dara.Prettify(s)
}

func (s DataIngestion) GoString() string {
	return s.String()
}

func (s *DataIngestion) GetActions() []*DataIngestionActions {
	return s.Actions
}

func (s *DataIngestion) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DataIngestion) GetError() *string {
	return s.Error
}

func (s *DataIngestion) GetId() *string {
	return s.Id
}

func (s *DataIngestion) GetInput() *Input {
	return s.Input
}

func (s *DataIngestion) GetMarker() *string {
	return s.Marker
}

func (s *DataIngestion) GetNotification() *DataIngestionNotification {
	return s.Notification
}

func (s *DataIngestion) GetPhase() *string {
	return s.Phase
}

func (s *DataIngestion) GetServiceRole() *string {
	return s.ServiceRole
}

func (s *DataIngestion) GetState() *string {
	return s.State
}

func (s *DataIngestion) GetStatistic() *DataIngestionStatistic {
	return s.Statistic
}

func (s *DataIngestion) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *DataIngestion) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DataIngestion) SetActions(v []*DataIngestionActions) *DataIngestion {
	s.Actions = v
	return s
}

func (s *DataIngestion) SetCreateTime(v string) *DataIngestion {
	s.CreateTime = &v
	return s
}

func (s *DataIngestion) SetError(v string) *DataIngestion {
	s.Error = &v
	return s
}

func (s *DataIngestion) SetId(v string) *DataIngestion {
	s.Id = &v
	return s
}

func (s *DataIngestion) SetInput(v *Input) *DataIngestion {
	s.Input = v
	return s
}

func (s *DataIngestion) SetMarker(v string) *DataIngestion {
	s.Marker = &v
	return s
}

func (s *DataIngestion) SetNotification(v *DataIngestionNotification) *DataIngestion {
	s.Notification = v
	return s
}

func (s *DataIngestion) SetPhase(v string) *DataIngestion {
	s.Phase = &v
	return s
}

func (s *DataIngestion) SetServiceRole(v string) *DataIngestion {
	s.ServiceRole = &v
	return s
}

func (s *DataIngestion) SetState(v string) *DataIngestion {
	s.State = &v
	return s
}

func (s *DataIngestion) SetStatistic(v *DataIngestionStatistic) *DataIngestion {
	s.Statistic = v
	return s
}

func (s *DataIngestion) SetTags(v map[string]interface{}) *DataIngestion {
	s.Tags = v
	return s
}

func (s *DataIngestion) SetUpdateTime(v string) *DataIngestion {
	s.UpdateTime = &v
	return s
}

func (s *DataIngestion) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Statistic != nil {
		if err := s.Statistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataIngestionActions struct {
	FastFailPolicy *FastFailPolicy `json:"FastFailPolicy,omitempty" xml:"FastFailPolicy,omitempty"`
	Name           *string         `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters     []*string       `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
}

func (s DataIngestionActions) String() string {
	return dara.Prettify(s)
}

func (s DataIngestionActions) GoString() string {
	return s.String()
}

func (s *DataIngestionActions) GetFastFailPolicy() *FastFailPolicy {
	return s.FastFailPolicy
}

func (s *DataIngestionActions) GetName() *string {
	return s.Name
}

func (s *DataIngestionActions) GetParameters() []*string {
	return s.Parameters
}

func (s *DataIngestionActions) SetFastFailPolicy(v *FastFailPolicy) *DataIngestionActions {
	s.FastFailPolicy = v
	return s
}

func (s *DataIngestionActions) SetName(v string) *DataIngestionActions {
	s.Name = &v
	return s
}

func (s *DataIngestionActions) SetParameters(v []*string) *DataIngestionActions {
	s.Parameters = v
	return s
}

func (s *DataIngestionActions) Validate() error {
	if s.FastFailPolicy != nil {
		if err := s.FastFailPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataIngestionNotification struct {
	Endpoint *string   `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	MNS      *MNS      `json:"MNS,omitempty" xml:"MNS,omitempty"`
	RocketMQ *RocketMQ `json:"RocketMQ,omitempty" xml:"RocketMQ,omitempty"`
	Topic    *string   `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DataIngestionNotification) String() string {
	return dara.Prettify(s)
}

func (s DataIngestionNotification) GoString() string {
	return s.String()
}

func (s *DataIngestionNotification) GetEndpoint() *string {
	return s.Endpoint
}

func (s *DataIngestionNotification) GetMNS() *MNS {
	return s.MNS
}

func (s *DataIngestionNotification) GetRocketMQ() *RocketMQ {
	return s.RocketMQ
}

func (s *DataIngestionNotification) GetTopic() *string {
	return s.Topic
}

func (s *DataIngestionNotification) SetEndpoint(v string) *DataIngestionNotification {
	s.Endpoint = &v
	return s
}

func (s *DataIngestionNotification) SetMNS(v *MNS) *DataIngestionNotification {
	s.MNS = v
	return s
}

func (s *DataIngestionNotification) SetRocketMQ(v *RocketMQ) *DataIngestionNotification {
	s.RocketMQ = v
	return s
}

func (s *DataIngestionNotification) SetTopic(v string) *DataIngestionNotification {
	s.Topic = &v
	return s
}

func (s *DataIngestionNotification) Validate() error {
	if s.MNS != nil {
		if err := s.MNS.Validate(); err != nil {
			return err
		}
	}
	if s.RocketMQ != nil {
		if err := s.RocketMQ.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataIngestionStatistic struct {
	SkipFiles     *int64 `json:"SkipFiles,omitempty" xml:"SkipFiles,omitempty"`
	SubmitFailure *int64 `json:"SubmitFailure,omitempty" xml:"SubmitFailure,omitempty"`
	SubmitSuccess *int64 `json:"SubmitSuccess,omitempty" xml:"SubmitSuccess,omitempty"`
}

func (s DataIngestionStatistic) String() string {
	return dara.Prettify(s)
}

func (s DataIngestionStatistic) GoString() string {
	return s.String()
}

func (s *DataIngestionStatistic) GetSkipFiles() *int64 {
	return s.SkipFiles
}

func (s *DataIngestionStatistic) GetSubmitFailure() *int64 {
	return s.SubmitFailure
}

func (s *DataIngestionStatistic) GetSubmitSuccess() *int64 {
	return s.SubmitSuccess
}

func (s *DataIngestionStatistic) SetSkipFiles(v int64) *DataIngestionStatistic {
	s.SkipFiles = &v
	return s
}

func (s *DataIngestionStatistic) SetSubmitFailure(v int64) *DataIngestionStatistic {
	s.SubmitFailure = &v
	return s
}

func (s *DataIngestionStatistic) SetSubmitSuccess(v int64) *DataIngestionStatistic {
	s.SubmitSuccess = &v
	return s
}

func (s *DataIngestionStatistic) Validate() error {
	return dara.Validate(s)
}
