// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNamespacesResponseBody) *DescribeNamespacesResponse
	GetBody() *DescribeNamespacesResponseBody
}

type DescribeNamespacesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNamespacesResponse) GetBody() *DescribeNamespacesResponseBody {
	return s.Body
}

func (s *DescribeNamespacesResponse) SetHeaders(v map[string]*string) *DescribeNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeNamespacesResponse) SetStatusCode(v int32) *DescribeNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNamespacesResponse) SetBody(v *DescribeNamespacesResponseBody) *DescribeNamespacesResponse {
	s.Body = v
	return s
}

func (s *DescribeNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
