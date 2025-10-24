// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceInstanceCertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceInstanceCertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceInstanceCertsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceInstanceCertsResponseBody) *DescribeResourceInstanceCertsResponse
	GetBody() *DescribeResourceInstanceCertsResponseBody
}

type DescribeResourceInstanceCertsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceInstanceCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceInstanceCertsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceInstanceCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceInstanceCertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceInstanceCertsResponse) GetBody() *DescribeResourceInstanceCertsResponseBody {
	return s.Body
}

func (s *DescribeResourceInstanceCertsResponse) SetHeaders(v map[string]*string) *DescribeResourceInstanceCertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceInstanceCertsResponse) SetStatusCode(v int32) *DescribeResourceInstanceCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceInstanceCertsResponse) SetBody(v *DescribeResourceInstanceCertsResponseBody) *DescribeResourceInstanceCertsResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceInstanceCertsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
