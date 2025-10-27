// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCheckWarningCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCheckWarningCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCheckWarningCountResponseBody) *DescribeCheckWarningCountResponse
	GetBody() *DescribeCheckWarningCountResponseBody
}

type DescribeCheckWarningCountResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckWarningCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckWarningCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCheckWarningCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCheckWarningCountResponse) GetBody() *DescribeCheckWarningCountResponseBody {
	return s.Body
}

func (s *DescribeCheckWarningCountResponse) SetHeaders(v map[string]*string) *DescribeCheckWarningCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckWarningCountResponse) SetStatusCode(v int32) *DescribeCheckWarningCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckWarningCountResponse) SetBody(v *DescribeCheckWarningCountResponseBody) *DescribeCheckWarningCountResponse {
	s.Body = v
	return s
}

func (s *DescribeCheckWarningCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
