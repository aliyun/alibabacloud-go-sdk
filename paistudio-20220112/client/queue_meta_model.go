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
	// The time the entry was removed from the queue.
	GmtDequeuedTime *string `json:"GmtDequeuedTime,omitempty" xml:"GmtDequeuedTime,omitempty"`
	// The time the entry was added to the queue.
	GmtEnqueuedTime *string `json:"GmtEnqueuedTime,omitempty" xml:"GmtEnqueuedTime,omitempty"`
	// The last time the position of the entry was modified.
	GmtPositionModifiedTime *string `json:"GmtPositionModifiedTime,omitempty" xml:"GmtPositionModifiedTime,omitempty"`
	// The name of the queue entry.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The position of the entry in the queue.
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The queuing strategy for the entry.
	QueueStrategy *string `json:"QueueStrategy,omitempty" xml:"QueueStrategy,omitempty"`
	// The ID of the associated quota.
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The resources that the queue entry requires.
	Resource *ResourceAmount `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The resources scheduled for the entry.
	ScheduledResource *string `json:"ScheduledResource,omitempty" xml:"ScheduledResource,omitempty"`
	// The current status of the entry.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the entry can use oversold resources.
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
