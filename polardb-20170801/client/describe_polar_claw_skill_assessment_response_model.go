// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillAssessmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawSkillAssessmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawSkillAssessmentResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawSkillAssessmentResponseBody) *DescribePolarClawSkillAssessmentResponse
	GetBody() *DescribePolarClawSkillAssessmentResponseBody
}

type DescribePolarClawSkillAssessmentResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawSkillAssessmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawSkillAssessmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillAssessmentResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillAssessmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawSkillAssessmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawSkillAssessmentResponse) GetBody() *DescribePolarClawSkillAssessmentResponseBody {
	return s.Body
}

func (s *DescribePolarClawSkillAssessmentResponse) SetHeaders(v map[string]*string) *DescribePolarClawSkillAssessmentResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawSkillAssessmentResponse) SetStatusCode(v int32) *DescribePolarClawSkillAssessmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawSkillAssessmentResponse) SetBody(v *DescribePolarClawSkillAssessmentResponseBody) *DescribePolarClawSkillAssessmentResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawSkillAssessmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
