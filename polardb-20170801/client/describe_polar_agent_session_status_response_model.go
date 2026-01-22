// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentSessionStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarAgentSessionStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarAgentSessionStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarAgentSessionStatusResponseBody) *DescribePolarAgentSessionStatusResponse
	GetBody() *DescribePolarAgentSessionStatusResponseBody
}

type DescribePolarAgentSessionStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarAgentSessionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarAgentSessionStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentSessionStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentSessionStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarAgentSessionStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarAgentSessionStatusResponse) GetBody() *DescribePolarAgentSessionStatusResponseBody {
	return s.Body
}

func (s *DescribePolarAgentSessionStatusResponse) SetHeaders(v map[string]*string) *DescribePolarAgentSessionStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarAgentSessionStatusResponse) SetStatusCode(v int32) *DescribePolarAgentSessionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarAgentSessionStatusResponse) SetBody(v *DescribePolarAgentSessionStatusResponseBody) *DescribePolarAgentSessionStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePolarAgentSessionStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
