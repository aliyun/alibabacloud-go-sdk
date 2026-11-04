// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventRecordPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *DeleteEventRecordPlanRequest
	GetPlanId() *string
}

type DeleteEventRecordPlanRequest struct {
	// The ID of the intelligent alert event recording plan. You can call the [QueryEventRecordPlans](https://help.aliyun.com/document_detail/107907.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e34a1989aa54debac22358aca7f****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteEventRecordPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRecordPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteEventRecordPlanRequest) SetPlanId(v string) *DeleteEventRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteEventRecordPlanRequest) Validate() error {
	return dara.Validate(s)
}
