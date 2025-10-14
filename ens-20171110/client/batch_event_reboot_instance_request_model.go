// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRebootInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfos(v []*BatchEventRebootInstanceRequestEventInfos) *BatchEventRebootInstanceRequest
	GetEventInfos() []*BatchEventRebootInstanceRequestEventInfos
}

type BatchEventRebootInstanceRequest struct {
	EventInfos []*BatchEventRebootInstanceRequestEventInfos `json:"EventInfos,omitempty" xml:"EventInfos,omitempty" type:"Repeated"`
}

func (s BatchEventRebootInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *BatchEventRebootInstanceRequest) GetEventInfos() []*BatchEventRebootInstanceRequestEventInfos {
	return s.EventInfos
}

func (s *BatchEventRebootInstanceRequest) SetEventInfos(v []*BatchEventRebootInstanceRequestEventInfos) *BatchEventRebootInstanceRequest {
	s.EventInfos = v
	return s
}

func (s *BatchEventRebootInstanceRequest) Validate() error {
	if s.EventInfos != nil {
		for _, item := range s.EventInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchEventRebootInstanceRequestEventInfos struct {
	// example:
	//
	// e-4452cec5a8f8eb9b2879a054207687d6
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// immediate
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// example:
	//
	// 1742452232000
	PlanTime *int64 `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
	// example:
	//
	// n-54hi3ffi63zrjt4wzx9mepeyh
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s BatchEventRebootInstanceRequestEventInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRebootInstanceRequestEventInfos) GoString() string {
	return s.String()
}

func (s *BatchEventRebootInstanceRequestEventInfos) GetEventId() *string {
	return s.EventId
}

func (s *BatchEventRebootInstanceRequestEventInfos) GetOpsType() *string {
	return s.OpsType
}

func (s *BatchEventRebootInstanceRequestEventInfos) GetPlanTime() *int64 {
	return s.PlanTime
}

func (s *BatchEventRebootInstanceRequestEventInfos) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchEventRebootInstanceRequestEventInfos) SetEventId(v string) *BatchEventRebootInstanceRequestEventInfos {
	s.EventId = &v
	return s
}

func (s *BatchEventRebootInstanceRequestEventInfos) SetOpsType(v string) *BatchEventRebootInstanceRequestEventInfos {
	s.OpsType = &v
	return s
}

func (s *BatchEventRebootInstanceRequestEventInfos) SetPlanTime(v int64) *BatchEventRebootInstanceRequestEventInfos {
	s.PlanTime = &v
	return s
}

func (s *BatchEventRebootInstanceRequestEventInfos) SetResourceId(v string) *BatchEventRebootInstanceRequestEventInfos {
	s.ResourceId = &v
	return s
}

func (s *BatchEventRebootInstanceRequestEventInfos) Validate() error {
	return dara.Validate(s)
}
