// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootECSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *RebootECSRequest
	GetEventId() *string
	SetInstanceId(v string) *RebootECSRequest
	GetInstanceId() *string
	SetRebootTime(v int64) *RebootECSRequest
	GetRebootTime() *int64
}

type RebootECSRequest struct {
	// example:
	//
	// 8c96a3fc8a0d4a48b5db5fdb9742fbbc
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1772076782
	RebootTime *int64 `json:"RebootTime,omitempty" xml:"RebootTime,omitempty"`
}

func (s RebootECSRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootECSRequest) GoString() string {
	return s.String()
}

func (s *RebootECSRequest) GetEventId() *string {
	return s.EventId
}

func (s *RebootECSRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootECSRequest) GetRebootTime() *int64 {
	return s.RebootTime
}

func (s *RebootECSRequest) SetEventId(v string) *RebootECSRequest {
	s.EventId = &v
	return s
}

func (s *RebootECSRequest) SetInstanceId(v string) *RebootECSRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootECSRequest) SetRebootTime(v int64) *RebootECSRequest {
	s.RebootTime = &v
	return s
}

func (s *RebootECSRequest) Validate() error {
	return dara.Validate(s)
}
