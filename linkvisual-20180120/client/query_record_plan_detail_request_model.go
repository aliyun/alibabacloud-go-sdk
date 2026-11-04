// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlanDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *QueryRecordPlanDetailRequest
	GetPlanId() *string
}

type QueryRecordPlanDetailRequest struct {
	// The ID of the recording plan. You can call [QueryRecordPlans](https://help.aliyun.com/document_detail/107896.html) to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b60bf0dd55944c19d7817442bc9****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s QueryRecordPlanDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlanDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRecordPlanDetailRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *QueryRecordPlanDetailRequest) SetPlanId(v string) *QueryRecordPlanDetailRequest {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlanDetailRequest) Validate() error {
	return dara.Validate(s)
}
