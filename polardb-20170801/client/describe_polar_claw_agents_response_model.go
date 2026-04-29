// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawAgentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawAgentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawAgentsResponseBody) *DescribePolarClawAgentsResponse
	GetBody() *DescribePolarClawAgentsResponseBody
}

type DescribePolarClawAgentsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawAgentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawAgentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawAgentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawAgentsResponse) GetBody() *DescribePolarClawAgentsResponseBody {
	return s.Body
}

func (s *DescribePolarClawAgentsResponse) SetHeaders(v map[string]*string) *DescribePolarClawAgentsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawAgentsResponse) SetStatusCode(v int32) *DescribePolarClawAgentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawAgentsResponse) SetBody(v *DescribePolarClawAgentsResponseBody) *DescribePolarClawAgentsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawAgentsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
