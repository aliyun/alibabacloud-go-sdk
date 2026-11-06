// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawSkillsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawSkillsResponseBody) *DescribePolarClawSkillsResponse
	GetBody() *DescribePolarClawSkillsResponseBody
}

type DescribePolarClawSkillsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawSkillsResponse) GetBody() *DescribePolarClawSkillsResponseBody {
	return s.Body
}

func (s *DescribePolarClawSkillsResponse) SetHeaders(v map[string]*string) *DescribePolarClawSkillsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawSkillsResponse) SetStatusCode(v int32) *DescribePolarClawSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawSkillsResponse) SetBody(v *DescribePolarClawSkillsResponseBody) *DescribePolarClawSkillsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
