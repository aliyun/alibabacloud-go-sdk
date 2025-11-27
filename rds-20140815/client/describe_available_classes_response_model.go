// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableClassesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableClassesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableClassesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableClassesResponseBody) *DescribeAvailableClassesResponse
	GetBody() *DescribeAvailableClassesResponseBody
}

type DescribeAvailableClassesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableClassesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableClassesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableClassesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableClassesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableClassesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableClassesResponse) GetBody() *DescribeAvailableClassesResponseBody {
	return s.Body
}

func (s *DescribeAvailableClassesResponse) SetHeaders(v map[string]*string) *DescribeAvailableClassesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableClassesResponse) SetStatusCode(v int32) *DescribeAvailableClassesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableClassesResponse) SetBody(v *DescribeAvailableClassesResponseBody) *DescribeAvailableClassesResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableClassesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
