// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEventRecordPlanDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *QueryEventRecordPlanDetailRequest
	GetPlanId() *string
}

type QueryEventRecordPlanDetailRequest struct {
	// The ID of the intelligent alert event recording plan. You can call [QueryEventRecordPlans](https://help.aliyun.com/document_detail/107907.html) to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryEventRecordPlanDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEventRecordPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryEventRecordPlanDetailRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *QueryEventRecordPlanDetailRequest) SetPlanId(v string) *QueryEventRecordPlanDetailRequest {
	s.PlanId = &v
	return s
}

func (s *QueryEventRecordPlanDetailRequest) Validate() error {
	return dara.Validate(s)
}
