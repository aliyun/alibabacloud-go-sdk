// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingBusinessOptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessMode(v string) *UpdateEventStreamingBusinessOptionRequest
	GetBusinessMode() *string
	SetEventStreamingName(v string) *UpdateEventStreamingBusinessOptionRequest
	GetEventStreamingName() *string
	SetMaxCapacityUnitCount(v int64) *UpdateEventStreamingBusinessOptionRequest
	GetMaxCapacityUnitCount() *int64
	SetMinCapacityUnitCount(v int64) *UpdateEventStreamingBusinessOptionRequest
	GetMinCapacityUnitCount() *int64
}

type UpdateEventStreamingBusinessOptionRequest struct {
	// This parameter is required.
	BusinessMode *string `json:"BusinessMode,omitempty" xml:"BusinessMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rocketmq-sync
	EventStreamingName *string `json:"EventStreamingName,omitempty" xml:"EventStreamingName,omitempty"`
	// example:
	//
	// 2
	MaxCapacityUnitCount *int64 `json:"MaxCapacityUnitCount,omitempty" xml:"MaxCapacityUnitCount,omitempty"`
	// example:
	//
	// 1
	MinCapacityUnitCount *int64 `json:"MinCapacityUnitCount,omitempty" xml:"MinCapacityUnitCount,omitempty"`
}

func (s UpdateEventStreamingBusinessOptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingBusinessOptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingBusinessOptionRequest) GetBusinessMode() *string {
	return s.BusinessMode
}

func (s *UpdateEventStreamingBusinessOptionRequest) GetEventStreamingName() *string {
	return s.EventStreamingName
}

func (s *UpdateEventStreamingBusinessOptionRequest) GetMaxCapacityUnitCount() *int64 {
	return s.MaxCapacityUnitCount
}

func (s *UpdateEventStreamingBusinessOptionRequest) GetMinCapacityUnitCount() *int64 {
	return s.MinCapacityUnitCount
}

func (s *UpdateEventStreamingBusinessOptionRequest) SetBusinessMode(v string) *UpdateEventStreamingBusinessOptionRequest {
	s.BusinessMode = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionRequest) SetEventStreamingName(v string) *UpdateEventStreamingBusinessOptionRequest {
	s.EventStreamingName = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionRequest) SetMaxCapacityUnitCount(v int64) *UpdateEventStreamingBusinessOptionRequest {
	s.MaxCapacityUnitCount = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionRequest) SetMinCapacityUnitCount(v int64) *UpdateEventStreamingBusinessOptionRequest {
	s.MinCapacityUnitCount = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionRequest) Validate() error {
	return dara.Validate(s)
}
