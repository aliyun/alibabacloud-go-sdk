// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSummaryTemplateRequest
	GetInstanceId() *string
	SetTemplateId(v string) *GetSummaryTemplateRequest
	GetTemplateId() *string
}

type GetSummaryTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 43c2671b-8939-4223-86d0-6bd187905cc8_1717664210492
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetSummaryTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetSummaryTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSummaryTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetSummaryTemplateRequest) SetInstanceId(v string) *GetSummaryTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSummaryTemplateRequest) SetTemplateId(v string) *GetSummaryTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetSummaryTemplateRequest) Validate() error {
	return dara.Validate(s)
}
