// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCAvailableResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCAvailableResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCAvailableResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCAvailableResourceResponseBody) *DescribeRCAvailableResourceResponse
	GetBody() *DescribeRCAvailableResourceResponseBody
}

type DescribeRCAvailableResourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCAvailableResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCAvailableResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCAvailableResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCAvailableResourceResponse) GetBody() *DescribeRCAvailableResourceResponseBody {
	return s.Body
}

func (s *DescribeRCAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeRCAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCAvailableResourceResponse) SetStatusCode(v int32) *DescribeRCAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCAvailableResourceResponse) SetBody(v *DescribeRCAvailableResourceResponseBody) *DescribeRCAvailableResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeRCAvailableResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
