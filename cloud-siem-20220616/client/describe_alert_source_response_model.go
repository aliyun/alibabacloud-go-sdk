// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAlertSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAlertSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAlertSourceResponseBody) *DescribeAlertSourceResponse
	GetBody() *DescribeAlertSourceResponseBody
}

type DescribeAlertSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAlertSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAlertSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAlertSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAlertSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAlertSourceResponse) GetBody() *DescribeAlertSourceResponseBody {
	return s.Body
}

func (s *DescribeAlertSourceResponse) SetHeaders(v map[string]*string) *DescribeAlertSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAlertSourceResponse) SetStatusCode(v int32) *DescribeAlertSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAlertSourceResponse) SetBody(v *DescribeAlertSourceResponseBody) *DescribeAlertSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeAlertSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
