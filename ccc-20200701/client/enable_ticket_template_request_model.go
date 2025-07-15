// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableTicketTemplateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *EnableTicketTemplateRequest
  GetInstanceId() *string 
  SetTemplateId(v string) *EnableTicketTemplateRequest
  GetTemplateId() *string 
}

type EnableTicketTemplateRequest struct {
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

func (s EnableTicketTemplateRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableTicketTemplateRequest) GoString() string {
  return s.String()
}

func (s *EnableTicketTemplateRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EnableTicketTemplateRequest) GetTemplateId() *string  {
  return s.TemplateId
}

func (s *EnableTicketTemplateRequest) SetInstanceId(v string) *EnableTicketTemplateRequest {
  s.InstanceId = &v
  return s
}

func (s *EnableTicketTemplateRequest) SetTemplateId(v string) *EnableTicketTemplateRequest {
  s.TemplateId = &v
  return s
}

func (s *EnableTicketTemplateRequest) Validate() error {
  return dara.Validate(s)
}

