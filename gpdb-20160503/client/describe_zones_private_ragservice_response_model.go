// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesPrivateRAGServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZonesPrivateRAGServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZonesPrivateRAGServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZonesPrivateRAGServiceResponseBody) *DescribeZonesPrivateRAGServiceResponse
	GetBody() *DescribeZonesPrivateRAGServiceResponseBody
}

type DescribeZonesPrivateRAGServiceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesPrivateRAGServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZonesPrivateRAGServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesPrivateRAGServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesPrivateRAGServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZonesPrivateRAGServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZonesPrivateRAGServiceResponse) GetBody() *DescribeZonesPrivateRAGServiceResponseBody {
	return s.Body
}

func (s *DescribeZonesPrivateRAGServiceResponse) SetHeaders(v map[string]*string) *DescribeZonesPrivateRAGServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesPrivateRAGServiceResponse) SetStatusCode(v int32) *DescribeZonesPrivateRAGServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesPrivateRAGServiceResponse) SetBody(v *DescribeZonesPrivateRAGServiceResponseBody) *DescribeZonesPrivateRAGServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeZonesPrivateRAGServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
