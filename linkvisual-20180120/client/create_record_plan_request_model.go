// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateRecordPlanRequest
	GetName() *string
	SetTemplateId(v string) *CreateRecordPlanRequest
	GetTemplateId() *string
}

type CreateRecordPlanRequest struct {
	// The name of the recording plan.
	//
	// The name must be less than 48 characters in length and can contain letters, digits, and the following special characters: hyphens (-), underscores (_), at signs (@), periods (.), and colons (:).
	//
	// This parameter is required.
	//
	// example:
	//
	// plan1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the time template for the recording plan. You can call the [CreateTimeTemplate](https://help.aliyun.com/document_detail/107888.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 36f06322442e4e3f85a5ae8862c9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateRecordPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordPlanRequest) GetName() *string {
	return s.Name
}

func (s *CreateRecordPlanRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateRecordPlanRequest) SetName(v string) *CreateRecordPlanRequest {
	s.Name = &v
	return s
}

func (s *CreateRecordPlanRequest) SetTemplateId(v string) *CreateRecordPlanRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateRecordPlanRequest) Validate() error {
	return dara.Validate(s)
}
