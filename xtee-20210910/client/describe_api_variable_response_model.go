// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiVariableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiVariableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiVariableResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiVariableResponseBody) *DescribeApiVariableResponse
	GetBody() *DescribeApiVariableResponseBody
}

type DescribeApiVariableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiVariableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiVariableResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiVariableResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiVariableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiVariableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiVariableResponse) GetBody() *DescribeApiVariableResponseBody {
	return s.Body
}

func (s *DescribeApiVariableResponse) SetHeaders(v map[string]*string) *DescribeApiVariableResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiVariableResponse) SetStatusCode(v int32) *DescribeApiVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiVariableResponse) SetBody(v *DescribeApiVariableResponseBody) *DescribeApiVariableResponse {
	s.Body = v
	return s
}

func (s *DescribeApiVariableResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
