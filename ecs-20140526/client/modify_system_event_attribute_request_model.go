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
	// This parameter is required.
	//
	// example:
	//
	// e-2zeielxl1qzq8slb****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-30T00:00:00Z
	NotBefore *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
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
