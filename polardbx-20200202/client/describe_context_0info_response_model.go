// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0InfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContext0InfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContext0InfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContext0InfoResponseBody) *DescribeContext0InfoResponse
	GetBody() *DescribeContext0InfoResponseBody
}

type DescribeContext0InfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContext0InfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContext0InfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContext0InfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContext0InfoResponse) GetBody() *DescribeContext0InfoResponseBody {
	return s.Body
}

func (s *DescribeContext0InfoResponse) SetHeaders(v map[string]*string) *DescribeContext0InfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeContext0InfoResponse) SetStatusCode(v int32) *DescribeContext0InfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContext0InfoResponse) SetBody(v *DescribeContext0InfoResponseBody) *DescribeContext0InfoResponse {
	s.Body = v
	return s
}

func (s *DescribeContext0InfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
