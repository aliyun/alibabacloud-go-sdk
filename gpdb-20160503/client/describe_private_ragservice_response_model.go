// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateRAGServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrivateRAGServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrivateRAGServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrivateRAGServiceResponseBody) *DescribePrivateRAGServiceResponse
	GetBody() *DescribePrivateRAGServiceResponseBody
}

type DescribePrivateRAGServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrivateRAGServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrivateRAGServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateRAGServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribePrivateRAGServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrivateRAGServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrivateRAGServiceResponse) GetBody() *DescribePrivateRAGServiceResponseBody {
	return s.Body
}

func (s *DescribePrivateRAGServiceResponse) SetHeaders(v map[string]*string) *DescribePrivateRAGServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribePrivateRAGServiceResponse) SetStatusCode(v int32) *DescribePrivateRAGServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrivateRAGServiceResponse) SetBody(v *DescribePrivateRAGServiceResponseBody) *DescribePrivateRAGServiceResponse {
	s.Body = v
	return s
}

func (s *DescribePrivateRAGServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
