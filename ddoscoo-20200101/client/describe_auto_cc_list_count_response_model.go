// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoCcListCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoCcListCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoCcListCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoCcListCountResponseBody) *DescribeAutoCcListCountResponse
	GetBody() *DescribeAutoCcListCountResponseBody
}

type DescribeAutoCcListCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoCcListCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoCcListCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoCcListCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoCcListCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoCcListCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoCcListCountResponse) GetBody() *DescribeAutoCcListCountResponseBody {
	return s.Body
}

func (s *DescribeAutoCcListCountResponse) SetHeaders(v map[string]*string) *DescribeAutoCcListCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoCcListCountResponse) SetStatusCode(v int32) *DescribeAutoCcListCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoCcListCountResponse) SetBody(v *DescribeAutoCcListCountResponseBody) *DescribeAutoCcListCountResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoCcListCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
