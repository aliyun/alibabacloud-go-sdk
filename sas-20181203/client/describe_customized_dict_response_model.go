// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizedDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizedDictResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizedDictResponseBody) *DescribeCustomizedDictResponse
	GetBody() *DescribeCustomizedDictResponseBody
}

type DescribeCustomizedDictResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizedDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizedDictResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedDictResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizedDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizedDictResponse) GetBody() *DescribeCustomizedDictResponseBody {
	return s.Body
}

func (s *DescribeCustomizedDictResponse) SetHeaders(v map[string]*string) *DescribeCustomizedDictResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizedDictResponse) SetStatusCode(v int32) *DescribeCustomizedDictResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizedDictResponse) SetBody(v *DescribeCustomizedDictResponseBody) *DescribeCustomizedDictResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizedDictResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
