// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAlertCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserAlertCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserAlertCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserAlertCountResponseBody) *DescribeUserAlertCountResponse
	GetBody() *DescribeUserAlertCountResponseBody
}

type DescribeUserAlertCountResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAlertCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserAlertCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlertCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAlertCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserAlertCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserAlertCountResponse) GetBody() *DescribeUserAlertCountResponseBody {
	return s.Body
}

func (s *DescribeUserAlertCountResponse) SetHeaders(v map[string]*string) *DescribeUserAlertCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAlertCountResponse) SetStatusCode(v int32) *DescribeUserAlertCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAlertCountResponse) SetBody(v *DescribeUserAlertCountResponseBody) *DescribeUserAlertCountResponse {
	s.Body = v
	return s
}

func (s *DescribeUserAlertCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
