// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySystemEventAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *ModifySystemEventAttributeRequest
	GetEventId() *string
	SetInstanceId(v string) *ModifySystemEventAttributeRequest
	GetInstanceId() *string
	SetNotBefore(v string) *ModifySystemEventAttributeRequest
	GetNotBefore() *string
	SetRegionId(v string) *ModifySystemEventAttributeRequest
	GetRegionId() *string
}

type ModifySystemEventAttributeRequest struct {
	// The event ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e-2zeielxl1qzq8slb****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new planned execution time of the system event. Specify the time in the [ISO 8601](https://www.alibabacloud.com/help/en/ecs/developer-reference/iso-8601-time-format) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-30T00:00:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifySystemEventAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySystemEventAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySystemEventAttributeRequest) GetEventId() *string {
	return s.EventId
}

func (s *ModifySystemEventAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifySystemEventAttributeRequest) GetNotBefore() *string {
	return s.NotBefore
}

func (s *ModifySystemEventAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySystemEventAttributeRequest) SetEventId(v string) *ModifySystemEventAttributeRequest {
	s.EventId = &v
	return s
}

func (s *ModifySystemEventAttributeRequest) SetInstanceId(v string) *ModifySystemEventAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySystemEventAttributeRequest) SetNotBefore(v string) *ModifySystemEventAttributeRequest {
	s.NotBefore = &v
	return s
}

func (s *ModifySystemEventAttributeRequest) SetRegionId(v string) *ModifySystemEventAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySystemEventAttributeRequest) Validate() error {
	return dara.Validate(s)
}
