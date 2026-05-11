// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceRolloutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceRolloutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceRolloutResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceRolloutResponseBody) *DescribeServiceRolloutResponse
	GetBody() *DescribeServiceRolloutResponseBody
}

type DescribeServiceRolloutResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceRolloutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceRolloutResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceRolloutResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceRolloutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceRolloutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceRolloutResponse) GetBody() *DescribeServiceRolloutResponseBody {
	return s.Body
}

func (s *DescribeServiceRolloutResponse) SetHeaders(v map[string]*string) *DescribeServiceRolloutResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceRolloutResponse) SetStatusCode(v int32) *DescribeServiceRolloutResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceRolloutResponse) SetBody(v *DescribeServiceRolloutResponseBody) *DescribeServiceRolloutResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceRolloutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
