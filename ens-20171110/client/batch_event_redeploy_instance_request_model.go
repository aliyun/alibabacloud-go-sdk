// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEventRedeployInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventInfos(v []*BatchEventRedeployInstanceRequestEventInfos) *BatchEventRedeployInstanceRequest
	GetEventInfos() []*BatchEventRedeployInstanceRequestEventInfos
}

type BatchEventRedeployInstanceRequest struct {
	EventInfos []*BatchEventRedeployInstanceRequestEventInfos `json:"EventInfos,omitempty" xml:"EventInfos,omitempty" type:"Repeated"`
}

func (s BatchEventRedeployInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRedeployInstanceRequest) GoString() string {
	return s.String()
}

func (s *BatchEventRedeployInstanceRequest) GetEventInfos() []*BatchEventRedeployInstanceRequestEventInfos {
	return s.EventInfos
}

func (s *BatchEventRedeployInstanceRequest) SetEventInfos(v []*BatchEventRedeployInstanceRequestEventInfos) *BatchEventRedeployInstanceRequest {
	s.EventInfos = v
	return s
}

func (s *BatchEventRedeployInstanceRequest) Validate() error {
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

type BatchEventRedeployInstanceRequestEventInfos struct {
	// example:
	//
	// e-d71ff150945b9c02eb6ebc0016328468
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// example:
	//
	// 1742452232000
	OpsType *string `json:"OpsType,omitempty" xml:"OpsType,omitempty"`
	// example:
	//
	// immediate
	PlanTime *int64 `json:"PlanTime,omitempty" xml:"PlanTime,omitempty"`
	// example:
	//
	// i-55qi8m11rr53c4i964md8a00l
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s BatchEventRedeployInstanceRequestEventInfos) String() string {
	return dara.Prettify(s)
}

func (s BatchEventRedeployInstanceRequestEventInfos) GoString() string {
	return s.String()
}

func (s *BatchEventRedeployInstanceRequestEventInfos) GetEventId() *string {
	return s.EventId
}

func (s *BatchEventRedeployInstanceRequestEventInfos) GetOpsType() *string {
	return s.OpsType
}

func (s *BatchEventRedeployInstanceRequestEventInfos) GetPlanTime() *int64 {
	return s.PlanTime
}

func (s *BatchEventRedeployInstanceRequestEventInfos) GetResourceId() *string {
	return s.ResourceId
}

func (s *BatchEventRedeployInstanceRequestEventInfos) SetEventId(v string) *BatchEventRedeployInstanceRequestEventInfos {
	s.EventId = &v
	return s
}

func (s *BatchEventRedeployInstanceRequestEventInfos) SetOpsType(v string) *BatchEventRedeployInstanceRequestEventInfos {
	s.OpsType = &v
	return s
}

func (s *BatchEventRedeployInstanceRequestEventInfos) SetPlanTime(v int64) *BatchEventRedeployInstanceRequestEventInfos {
	s.PlanTime = &v
	return s
}

func (s *BatchEventRedeployInstanceRequestEventInfos) SetResourceId(v string) *BatchEventRedeployInstanceRequestEventInfos {
	s.ResourceId = &v
	return s
}

func (s *BatchEventRedeployInstanceRequestEventInfos) Validate() error {
	return dara.Validate(s)
}
