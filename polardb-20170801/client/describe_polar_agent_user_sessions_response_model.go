// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentUserSessionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarAgentUserSessionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarAgentUserSessionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarAgentUserSessionsResponseBody) *DescribePolarAgentUserSessionsResponse
	GetBody() *DescribePolarAgentUserSessionsResponseBody
}

type DescribePolarAgentUserSessionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarAgentUserSessionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarAgentUserSessionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentUserSessionsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentUserSessionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarAgentUserSessionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarAgentUserSessionsResponse) GetBody() *DescribePolarAgentUserSessionsResponseBody {
	return s.Body
}

func (s *DescribePolarAgentUserSessionsResponse) SetHeaders(v map[string]*string) *DescribePolarAgentUserSessionsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarAgentUserSessionsResponse) SetStatusCode(v int32) *DescribePolarAgentUserSessionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarAgentUserSessionsResponse) SetBody(v *DescribePolarAgentUserSessionsResponseBody) *DescribePolarAgentUserSessionsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarAgentUserSessionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
