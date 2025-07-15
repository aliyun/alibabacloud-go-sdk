// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTicketTemplateRequest
	GetInstanceId() *string
	SetTemplateId(v string) *GetTicketTemplateRequest
	GetTemplateId() *string
	SetTemplateVersion(v string) *GetTicketTemplateRequest
	GetTemplateVersion() *string
}

type GetTicketTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 4ca2e2-c8d19b82c-d7ce393ac8197d3ab
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 1703517780627
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTicketTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTicketTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTicketTemplateRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTicketTemplateRequest) SetInstanceId(v string) *GetTicketTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTicketTemplateRequest) SetTemplateId(v string) *GetTicketTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTicketTemplateRequest) SetTemplateVersion(v string) *GetTicketTemplateRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTicketTemplateRequest) Validate() error {
	return dara.Validate(s)
}
