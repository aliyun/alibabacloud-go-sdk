// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawSkillTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawSkillTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawSkillTemplateResponseBody) *DescribePolarClawSkillTemplateResponse
	GetBody() *DescribePolarClawSkillTemplateResponseBody
}

type DescribePolarClawSkillTemplateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawSkillTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawSkillTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawSkillTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawSkillTemplateResponse) GetBody() *DescribePolarClawSkillTemplateResponseBody {
	return s.Body
}

func (s *DescribePolarClawSkillTemplateResponse) SetHeaders(v map[string]*string) *DescribePolarClawSkillTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawSkillTemplateResponse) SetStatusCode(v int32) *DescribePolarClawSkillTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawSkillTemplateResponse) SetBody(v *DescribePolarClawSkillTemplateResponseBody) *DescribePolarClawSkillTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawSkillTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
