// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAuthorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccountAuthorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccountAuthorityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccountAuthorityResponseBody) *DescribeAccountAuthorityResponse
	GetBody() *DescribeAccountAuthorityResponseBody
}

type DescribeAccountAuthorityResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccountAuthorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccountAuthorityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAuthorityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountAuthorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccountAuthorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccountAuthorityResponse) GetBody() *DescribeAccountAuthorityResponseBody {
	return s.Body
}

func (s *DescribeAccountAuthorityResponse) SetHeaders(v map[string]*string) *DescribeAccountAuthorityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccountAuthorityResponse) SetStatusCode(v int32) *DescribeAccountAuthorityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccountAuthorityResponse) SetBody(v *DescribeAccountAuthorityResponseBody) *DescribeAccountAuthorityResponse {
	s.Body = v
	return s
}

func (s *DescribeAccountAuthorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
