// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecordPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *DeleteRecordPlanRequest
	GetPlanId() *string
}

type DeleteRecordPlanRequest struct {
	// The ID of the recording plan. You can call the [QueryRecordPlans](https://help.aliyun.com/document_detail/107896.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6c5a397bc17c4b48ace86f0f0c30****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteRecordPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *DeleteRecordPlanRequest) SetPlanId(v string) *DeleteRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *DeleteRecordPlanRequest) Validate() error {
	return dara.Validate(s)
}
