// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateRecordPlanRequest
	GetName() *string
	SetPlanId(v string) *UpdateRecordPlanRequest
	GetPlanId() *string
	SetTemplateId(v string) *UpdateRecordPlanRequest
	GetTemplateId() *string
}

type UpdateRecordPlanRequest struct {
	// The name of the recording plan.
	//
	// The name must be less than 48 characters in length and can contain letters, digits, and the following special characters: hyphens (-), underscores (_), at signs (@), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the recording plan. You can call [QueryRecordPlans](https://help.aliyun.com/document_detail/107896.html) to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6c5a397bc17c4b48ace86f0f0c30****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the time template to be associated with the recording plan. You can call [QueryTimeTemplate](https://help.aliyun.com/document_detail/107890.html) to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36f06322442e4e3f85a5ae8862c9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateRecordPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordPlanRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRecordPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateRecordPlanRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateRecordPlanRequest) SetName(v string) *UpdateRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetPlanId(v string) *UpdateRecordPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateRecordPlanRequest) SetTemplateId(v string) *UpdateRecordPlanRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateRecordPlanRequest) Validate() error {
	return dara.Validate(s)
}
