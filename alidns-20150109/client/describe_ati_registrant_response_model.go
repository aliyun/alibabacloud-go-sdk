// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiRegistrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAtiRegistrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAtiRegistrantResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAtiRegistrantResponseBody) *DescribeAtiRegistrantResponse
	GetBody() *DescribeAtiRegistrantResponseBody
}

type DescribeAtiRegistrantResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAtiRegistrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAtiRegistrantResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiRegistrantResponse) GoString() string {
	return s.String()
}

func (s *DescribeAtiRegistrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAtiRegistrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAtiRegistrantResponse) GetBody() *DescribeAtiRegistrantResponseBody {
	return s.Body
}

func (s *DescribeAtiRegistrantResponse) SetHeaders(v map[string]*string) *DescribeAtiRegistrantResponse {
	s.Headers = v
	return s
}

func (s *DescribeAtiRegistrantResponse) SetStatusCode(v int32) *DescribeAtiRegistrantResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAtiRegistrantResponse) SetBody(v *DescribeAtiRegistrantResponseBody) *DescribeAtiRegistrantResponse {
	s.Body = v
	return s
}

func (s *DescribeAtiRegistrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
