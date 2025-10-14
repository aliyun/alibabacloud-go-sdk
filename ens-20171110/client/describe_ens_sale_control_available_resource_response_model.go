// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlAvailableResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnsSaleControlAvailableResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnsSaleControlAvailableResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnsSaleControlAvailableResourceResponseBody) *DescribeEnsSaleControlAvailableResourceResponse
	GetBody() *DescribeEnsSaleControlAvailableResourceResponseBody
}

type DescribeEnsSaleControlAvailableResourceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnsSaleControlAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) GetBody() *DescribeEnsSaleControlAvailableResourceResponseBody {
	return s.Body
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeEnsSaleControlAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) SetStatusCode(v int32) *DescribeEnsSaleControlAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) SetBody(v *DescribeEnsSaleControlAvailableResourceResponseBody) *DescribeEnsSaleControlAvailableResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
