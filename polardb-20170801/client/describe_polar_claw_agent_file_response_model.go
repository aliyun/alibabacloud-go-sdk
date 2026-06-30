// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawAgentFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawAgentFileResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawAgentFileResponseBody) *DescribePolarClawAgentFileResponse
	GetBody() *DescribePolarClawAgentFileResponseBody
}

type DescribePolarClawAgentFileResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawAgentFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawAgentFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentFileResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawAgentFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawAgentFileResponse) GetBody() *DescribePolarClawAgentFileResponseBody {
	return s.Body
}

func (s *DescribePolarClawAgentFileResponse) SetHeaders(v map[string]*string) *DescribePolarClawAgentFileResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawAgentFileResponse) SetStatusCode(v int32) *DescribePolarClawAgentFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawAgentFileResponse) SetBody(v *DescribePolarClawAgentFileResponseBody) *DescribePolarClawAgentFileResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawAgentFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
