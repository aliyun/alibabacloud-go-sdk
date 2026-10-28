// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueueMeta interface {
	dara.Model
	String() string
	GoString() string
	SetGmtDequeuedTime(v string) *QueueMeta
	GetGmtDequeuedTime() *string
	SetGmtEnqueuedTime(v string) *QueueMeta
	GetGmtEnqueuedTime() *string
	SetGmtPositionModifiedTime(v string) *QueueMeta
	GetGmtPositionModifiedTime() *string
	SetName(v string) *QueueMeta
	GetName() *string
	SetPosition(v string) *QueueMeta
	GetPosition() *string
	SetQueueStrategy(v string) *QueueMeta
	GetQueueStrategy() *string
	SetQuotaId(v string) *QueueMeta
	GetQuotaId() *string
	SetResource(v *ResourceAmount) *QueueMeta
	GetResource() *ResourceAmount
	SetScheduledResource(v string) *QueueMeta
	GetScheduledResource() *string
	SetStatus(v string) *QueueMeta
	GetStatus() *string
	SetUseOversoldResource(v bool) *QueueMeta
	GetUseOversoldResource() *bool
}

type QueueMeta struct {
	// The time when the task was dequeued.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtDequeuedTime *string `json:"GmtDequeuedTime,omitempty" xml:"GmtDequeuedTime,omitempty"`
	// The time when the task was enqueued.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtEnqueuedTime *string `json:"GmtEnqueuedTime,omitempty" xml:"GmtEnqueuedTime,omitempty"`
	// The time when the position of the task in the queue was changed.
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	GmtPositionModifiedTime *string `json:"GmtPositionModifiedTime,omitempty" xml:"GmtPositionModifiedTime,omitempty"`
	// The name of the task in the queue.
	//
	// example:
	//
	// dlcxxxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The position of the task in the queue. This parameter takes effect only when the task is queued.
	//
	// example:
	//
	// 10
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The dequeue strategy.
	//
	// example:
	//
	// PaiStrategyIntelligent
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// The resource quota queue to which the task belongs.
	//
	// example:
	//
	// quotad2kd8ljpsno
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The resources occupied by the task in the queue.
	Resource *ResourceAmount `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The scheduled resources occupied by the task in the queue.
	ScheduledResource *string `json:"ScheduledResource,omitempty" xml:"ScheduledResource,omitempty"`
	// The status of the task in the queue, such as enqueued or dequeued.
	//
	// example:
	//
	// Enqueued
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Specifies whether the task uses idle resources.
	UseOversoldResource *bool `json:"UseOversoldResource,omitempty" xml:"UseOversoldResource,omitempty"`
}

func (s QueueMeta) String() string {
	return dara.Prettify(s)
}

func (s QueueMeta) GoString() string {
	return s.String()
}

func (s *QueueMeta) GetGmtDequeuedTime() *string {
	return s.GmtDequeuedTime
}

func (s *QueueMeta) GetGmtEnqueuedTime() *string {
	return s.GmtEnqueuedTime
}

func (s *QueueMeta) GetGmtPositionModifiedTime() *string {
	return s.GmtPositionModifiedTime
}

func (s *QueueMeta) GetName() *string {
	return s.Name
}

func (s *QueueMeta) GetPosition() *string {
	return s.Position
}

func (s *QueueMeta) GetQueueStrategy() *string {
	return s.QueueStrategy
}

func (s *QueueMeta) GetQuotaId() *string {
	return s.QuotaId
}

func (s *QueueMeta) GetResource() *ResourceAmount {
	return s.Resource
}

func (s *QueueMeta) GetScheduledResource() *string {
	return s.ScheduledResource
}

func (s *QueueMeta) GetStatus() *string {
	return s.Status
}

func (s *QueueMeta) GetUseOversoldResource() *bool {
	return s.UseOversoldResource
}

func (s *QueueMeta) SetGmtDequeuedTime(v string) *QueueMeta {
	s.GmtDequeuedTime = &v
	return s
}

func (s *QueueMeta) SetGmtEnqueuedTime(v string) *QueueMeta {
	s.GmtEnqueuedTime = &v
	return s
}

func (s *QueueMeta) SetGmtPositionModifiedTime(v string) *QueueMeta {
	s.GmtPositionModifiedTime = &v
	return s
}

func (s *QueueMeta) SetName(v string) *QueueMeta {
	s.Name = &v
	return s
}

func (s *QueueMeta) SetPosition(v string) *QueueMeta {
	s.Position = &v
	return s
}

func (s *QueueMeta) SetQueueStrategy(v string) *QueueMeta {
	s.QueueStrategy = &v
	return s
}

func (s *QueueMeta) SetQuotaId(v string) *QueueMeta {
	s.QuotaId = &v
	return s
}

func (s *QueueMeta) SetResource(v *ResourceAmount) *QueueMeta {
	s.Resource = v
	return s
}

func (s *QueueMeta) SetScheduledResource(v string) *QueueMeta {
	s.ScheduledResource = &v
	return s
}

func (s *QueueMeta) SetStatus(v string) *QueueMeta {
	s.Status = &v
	return s
}

func (s *QueueMeta) SetUseOversoldResource(v bool) *QueueMeta {
	s.UseOversoldResource = &v
	return s
}

func (s *QueueMeta) Validate() error {
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}
