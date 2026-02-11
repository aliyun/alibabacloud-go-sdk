// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsWithEntityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertsWithEntityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertsWithEntityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertsWithEntityResponseBody) *DescribeAlertsWithEntityResponse
	GetBody() *DescribeAlertsWithEntityResponseBody
}

type DescribeAlertsWithEntityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsWithEntityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsWithEntityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertsWithEntityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertsWithEntityResponse) GetBody() *DescribeAlertsWithEntityResponseBody {
	return s.Body
}

func (s *DescribeAlertsWithEntityResponse) SetHeaders(v map[string]*string) *DescribeAlertsWithEntityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsWithEntityResponse) SetStatusCode(v int32) *DescribeAlertsWithEntityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponse) SetBody(v *DescribeAlertsWithEntityResponseBody) *DescribeAlertsWithEntityResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertsWithEntityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
