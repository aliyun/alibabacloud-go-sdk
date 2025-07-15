// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableTicketTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableTicketTemplateRequest
	GetInstanceId() *string
	SetTemplateId(v string) *DisableTicketTemplateRequest
	GetTemplateId() *string
}

type DisableTicketTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 43c2671b-8939-4223-86d0-6bd187905cc8_1717664210492
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DisableTicketTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableTicketTemplateRequest) GoString() string {
	return s.String()
}

func (s *DisableTicketTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableTicketTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DisableTicketTemplateRequest) SetInstanceId(v string) *DisableTicketTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableTicketTemplateRequest) SetTemplateId(v string) *DisableTicketTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DisableTicketTemplateRequest) Validate() error {
	return dara.Validate(s)
}
