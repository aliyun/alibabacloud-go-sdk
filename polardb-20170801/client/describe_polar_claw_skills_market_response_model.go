// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawSkillsMarketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawSkillsMarketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawSkillsMarketResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawSkillsMarketResponseBody) *DescribePolarClawSkillsMarketResponse
	GetBody() *DescribePolarClawSkillsMarketResponseBody
}

type DescribePolarClawSkillsMarketResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawSkillsMarketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawSkillsMarketResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawSkillsMarketResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawSkillsMarketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawSkillsMarketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawSkillsMarketResponse) GetBody() *DescribePolarClawSkillsMarketResponseBody {
	return s.Body
}

func (s *DescribePolarClawSkillsMarketResponse) SetHeaders(v map[string]*string) *DescribePolarClawSkillsMarketResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawSkillsMarketResponse) SetStatusCode(v int32) *DescribePolarClawSkillsMarketResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawSkillsMarketResponse) SetBody(v *DescribePolarClawSkillsMarketResponseBody) *DescribePolarClawSkillsMarketResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawSkillsMarketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
