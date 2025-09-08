// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertsCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertsCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertsCountResponseBody) *DescribeAlertsCountResponse
	GetBody() *DescribeAlertsCountResponseBody
}

type DescribeAlertsCountResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertsCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertsCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertsCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertsCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertsCountResponse) GetBody() *DescribeAlertsCountResponseBody {
	return s.Body
}

func (s *DescribeAlertsCountResponse) SetHeaders(v map[string]*string) *DescribeAlertsCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertsCountResponse) SetStatusCode(v int32) *DescribeAlertsCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertsCountResponse) SetBody(v *DescribeAlertsCountResponseBody) *DescribeAlertsCountResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertsCountResponse) Validate() error {
	return dara.Validate(s)
}
