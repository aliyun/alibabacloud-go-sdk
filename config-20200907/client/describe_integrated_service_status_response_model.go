// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntegratedServiceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIntegratedServiceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIntegratedServiceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIntegratedServiceStatusResponseBody) *DescribeIntegratedServiceStatusResponse
	GetBody() *DescribeIntegratedServiceStatusResponseBody
}

type DescribeIntegratedServiceStatusResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntegratedServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntegratedServiceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntegratedServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntegratedServiceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIntegratedServiceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIntegratedServiceStatusResponse) GetBody() *DescribeIntegratedServiceStatusResponseBody {
	return s.Body
}

func (s *DescribeIntegratedServiceStatusResponse) SetHeaders(v map[string]*string) *DescribeIntegratedServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntegratedServiceStatusResponse) SetStatusCode(v int32) *DescribeIntegratedServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntegratedServiceStatusResponse) SetBody(v *DescribeIntegratedServiceStatusResponseBody) *DescribeIntegratedServiceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeIntegratedServiceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
