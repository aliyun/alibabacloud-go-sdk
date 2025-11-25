// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConfiguredDestinationIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConfiguredDestinationIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConfiguredDestinationIPResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConfiguredDestinationIPResponseBody) *DescribeConfiguredDestinationIPResponse
	GetBody() *DescribeConfiguredDestinationIPResponseBody
}

type DescribeConfiguredDestinationIPResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConfiguredDestinationIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConfiguredDestinationIPResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConfiguredDestinationIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfiguredDestinationIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConfiguredDestinationIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConfiguredDestinationIPResponse) GetBody() *DescribeConfiguredDestinationIPResponseBody {
	return s.Body
}

func (s *DescribeConfiguredDestinationIPResponse) SetHeaders(v map[string]*string) *DescribeConfiguredDestinationIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfiguredDestinationIPResponse) SetStatusCode(v int32) *DescribeConfiguredDestinationIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfiguredDestinationIPResponse) SetBody(v *DescribeConfiguredDestinationIPResponseBody) *DescribeConfiguredDestinationIPResponse {
	s.Body = v
	return s
}

func (s *DescribeConfiguredDestinationIPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
