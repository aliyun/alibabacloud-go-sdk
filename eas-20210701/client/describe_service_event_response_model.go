// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceEventResponseBody) *DescribeServiceEventResponse
	GetBody() *DescribeServiceEventResponseBody
}

type DescribeServiceEventResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceEventResponse) GetBody() *DescribeServiceEventResponseBody {
	return s.Body
}

func (s *DescribeServiceEventResponse) SetHeaders(v map[string]*string) *DescribeServiceEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceEventResponse) SetStatusCode(v int32) *DescribeServiceEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceEventResponse) SetBody(v *DescribeServiceEventResponseBody) *DescribeServiceEventResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
