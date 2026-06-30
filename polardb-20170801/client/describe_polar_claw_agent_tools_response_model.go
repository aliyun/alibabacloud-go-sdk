// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawAgentToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePolarClawAgentToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePolarClawAgentToolsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePolarClawAgentToolsResponseBody) *DescribePolarClawAgentToolsResponse
	GetBody() *DescribePolarClawAgentToolsResponseBody
}

type DescribePolarClawAgentToolsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePolarClawAgentToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePolarClawAgentToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawAgentToolsResponse) GoString() string {
	return s.String()
}

func (s *DescribePolarClawAgentToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePolarClawAgentToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePolarClawAgentToolsResponse) GetBody() *DescribePolarClawAgentToolsResponseBody {
	return s.Body
}

func (s *DescribePolarClawAgentToolsResponse) SetHeaders(v map[string]*string) *DescribePolarClawAgentToolsResponse {
	s.Headers = v
	return s
}

func (s *DescribePolarClawAgentToolsResponse) SetStatusCode(v int32) *DescribePolarClawAgentToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePolarClawAgentToolsResponse) SetBody(v *DescribePolarClawAgentToolsResponseBody) *DescribePolarClawAgentToolsResponse {
	s.Body = v
	return s
}

func (s *DescribePolarClawAgentToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
