// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAIServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAIServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAIServiceResponseBody) *DescribeAIServiceResponse
	GetBody() *DescribeAIServiceResponseBody
}

type DescribeAIServiceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAIServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAIServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAIServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAIServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAIServiceResponse) GetBody() *DescribeAIServiceResponseBody {
	return s.Body
}

func (s *DescribeAIServiceResponse) SetHeaders(v map[string]*string) *DescribeAIServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAIServiceResponse) SetStatusCode(v int32) *DescribeAIServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAIServiceResponse) SetBody(v *DescribeAIServiceResponseBody) *DescribeAIServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeAIServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
