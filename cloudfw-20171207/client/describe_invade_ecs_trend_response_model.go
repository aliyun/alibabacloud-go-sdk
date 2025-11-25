// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvadeEcsTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvadeEcsTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvadeEcsTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvadeEcsTrendResponseBody) *DescribeInvadeEcsTrendResponse
	GetBody() *DescribeInvadeEcsTrendResponseBody
}

type DescribeInvadeEcsTrendResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvadeEcsTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvadeEcsTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvadeEcsTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvadeEcsTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvadeEcsTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvadeEcsTrendResponse) GetBody() *DescribeInvadeEcsTrendResponseBody {
	return s.Body
}

func (s *DescribeInvadeEcsTrendResponse) SetHeaders(v map[string]*string) *DescribeInvadeEcsTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvadeEcsTrendResponse) SetStatusCode(v int32) *DescribeInvadeEcsTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvadeEcsTrendResponse) SetBody(v *DescribeInvadeEcsTrendResponseBody) *DescribeInvadeEcsTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeInvadeEcsTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
