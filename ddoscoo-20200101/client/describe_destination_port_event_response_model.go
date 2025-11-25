// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDestinationPortEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDestinationPortEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDestinationPortEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDestinationPortEventResponseBody) *DescribeDestinationPortEventResponse
	GetBody() *DescribeDestinationPortEventResponseBody
}

type DescribeDestinationPortEventResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDestinationPortEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDestinationPortEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDestinationPortEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeDestinationPortEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDestinationPortEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDestinationPortEventResponse) GetBody() *DescribeDestinationPortEventResponseBody {
	return s.Body
}

func (s *DescribeDestinationPortEventResponse) SetHeaders(v map[string]*string) *DescribeDestinationPortEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeDestinationPortEventResponse) SetStatusCode(v int32) *DescribeDestinationPortEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDestinationPortEventResponse) SetBody(v *DescribeDestinationPortEventResponseBody) *DescribeDestinationPortEventResponse {
	s.Body = v
	return s
}

func (s *DescribeDestinationPortEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
