// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCostRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCostRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCostRulesResponseBody) *DescribeCostRulesResponse
	GetBody() *DescribeCostRulesResponseBody
}

type DescribeCostRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCostRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCostRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCostRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCostRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCostRulesResponse) GetBody() *DescribeCostRulesResponseBody {
	return s.Body
}

func (s *DescribeCostRulesResponse) SetHeaders(v map[string]*string) *DescribeCostRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCostRulesResponse) SetStatusCode(v int32) *DescribeCostRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCostRulesResponse) SetBody(v *DescribeCostRulesResponseBody) *DescribeCostRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeCostRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
