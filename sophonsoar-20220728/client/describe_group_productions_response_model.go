// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupProductionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupProductionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupProductionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupProductionsResponseBody) *DescribeGroupProductionsResponse
	GetBody() *DescribeGroupProductionsResponseBody
}

type DescribeGroupProductionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupProductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupProductionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupProductionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupProductionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupProductionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupProductionsResponse) GetBody() *DescribeGroupProductionsResponseBody {
	return s.Body
}

func (s *DescribeGroupProductionsResponse) SetHeaders(v map[string]*string) *DescribeGroupProductionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupProductionsResponse) SetStatusCode(v int32) *DescribeGroupProductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupProductionsResponse) SetBody(v *DescribeGroupProductionsResponseBody) *DescribeGroupProductionsResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupProductionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
