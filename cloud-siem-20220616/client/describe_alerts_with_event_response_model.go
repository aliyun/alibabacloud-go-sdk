// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsWithEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertsWithEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertsWithEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertsWithEventResponseBody) *DescribeAlertsWithEventResponse
	GetBody() *DescribeAlertsWithEventResponseBody
}

type DescribeAlertsWithEventResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsWithEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsWithEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertsWithEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertsWithEventResponse) GetBody() *DescribeAlertsWithEventResponseBody {
	return s.Body
}

func (s *DescribeAlertsWithEventResponse) SetHeaders(v map[string]*string) *DescribeAlertsWithEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsWithEventResponse) SetStatusCode(v int32) *DescribeAlertsWithEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponse) SetBody(v *DescribeAlertsWithEventResponseBody) *DescribeAlertsWithEventResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertsWithEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
