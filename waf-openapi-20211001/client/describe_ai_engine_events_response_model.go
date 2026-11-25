// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAiEngineEventsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAiEngineEventsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAiEngineEventsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAiEngineEventsResponseBody) *DescribeAiEngineEventsResponse
	GetBody() *DescribeAiEngineEventsResponseBody
}

type DescribeAiEngineEventsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAiEngineEventsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAiEngineEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAiEngineEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAiEngineEventsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAiEngineEventsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAiEngineEventsResponse) GetBody() *DescribeAiEngineEventsResponseBody {
	return s.Body
}

func (s *DescribeAiEngineEventsResponse) SetHeaders(v map[string]*string) *DescribeAiEngineEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAiEngineEventsResponse) SetStatusCode(v int32) *DescribeAiEngineEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAiEngineEventsResponse) SetBody(v *DescribeAiEngineEventsResponseBody) *DescribeAiEngineEventsResponse {
	s.Body = v
	return s
}

func (s *DescribeAiEngineEventsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
