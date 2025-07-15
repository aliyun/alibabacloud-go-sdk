// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTicketTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTicketTemplateRequest
	GetInstanceId() *string
	SetTemplateId(v string) *DeleteTicketTemplateRequest
	GetTemplateId() *string
}

type DeleteTicketTemplateRequest struct {
	// example:
	//
	// ef1e71e9-ae9d-487c-96ad-9181d85cf802
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// **43c2671b-8939-4223-***-6bd187905cc8_1717664210492
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTicketTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTicketTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTicketTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTicketTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteTicketTemplateRequest) SetInstanceId(v string) *DeleteTicketTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTicketTemplateRequest) SetTemplateId(v string) *DeleteTicketTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteTicketTemplateRequest) Validate() error {
	return dara.Validate(s)
}
