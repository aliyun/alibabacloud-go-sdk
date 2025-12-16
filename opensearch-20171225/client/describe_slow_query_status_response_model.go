// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowQueryStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlowQueryStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlowQueryStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlowQueryStatusResponseBody) *DescribeSlowQueryStatusResponse
	GetBody() *DescribeSlowQueryStatusResponseBody
}

type DescribeSlowQueryStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowQueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlowQueryStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlowQueryStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlowQueryStatusResponse) GetBody() *DescribeSlowQueryStatusResponseBody {
	return s.Body
}

func (s *DescribeSlowQueryStatusResponse) SetHeaders(v map[string]*string) *DescribeSlowQueryStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowQueryStatusResponse) SetStatusCode(v int32) *DescribeSlowQueryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowQueryStatusResponse) SetBody(v *DescribeSlowQueryStatusResponseBody) *DescribeSlowQueryStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSlowQueryStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
