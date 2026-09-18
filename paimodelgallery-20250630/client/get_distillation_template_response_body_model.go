// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDistillationTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistillationTemplate(v *DistillationTemplate) *GetDistillationTemplateResponseBody
	GetDistillationTemplate() *DistillationTemplate
	SetRequestId(v string) *GetDistillationTemplateResponseBody
	GetRequestId() *string
}

type GetDistillationTemplateResponseBody struct {
	// The distillation template details.
	DistillationTemplate *DistillationTemplate `json:"DistillationTemplate,omitempty" xml:"DistillationTemplate,omitempty"`
	// **Request ID**
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDistillationTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDistillationTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetDistillationTemplateResponseBody) GetDistillationTemplate() *DistillationTemplate {
	return s.DistillationTemplate
}

func (s *GetDistillationTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDistillationTemplateResponseBody) SetDistillationTemplate(v *DistillationTemplate) *GetDistillationTemplateResponseBody {
	s.DistillationTemplate = v
	return s
}

func (s *GetDistillationTemplateResponseBody) SetRequestId(v string) *GetDistillationTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDistillationTemplateResponseBody) Validate() error {
	if s.DistillationTemplate != nil {
		if err := s.DistillationTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}
