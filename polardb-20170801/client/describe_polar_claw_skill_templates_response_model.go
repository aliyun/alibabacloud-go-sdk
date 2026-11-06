// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawSkillTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawSkillTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawSkillTemplatesResponseBody) *DescribePolarClawSkillTemplatesResponse
	GetBody() *DescribePolarClawSkillTemplatesResponseBody
}

type DescribePolarClawSkillTemplatesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawSkillTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawSkillTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawSkillTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawSkillTemplatesResponse) GetBody() *DescribePolarClawSkillTemplatesResponseBody {
	return s.Body
}

func (s *DescribePolarClawSkillTemplatesResponse) SetHeaders(v map[string]*string) *DescribePolarClawSkillTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponse) SetStatusCode(v int32) *DescribePolarClawSkillTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponse) SetBody(v *DescribePolarClawSkillTemplatesResponseBody) *DescribePolarClawSkillTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawSkillTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
