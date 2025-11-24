// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGuestClusterNamespacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGuestClusterNamespacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGuestClusterNamespacesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGuestClusterNamespacesResponseBody) *DescribeGuestClusterNamespacesResponse
	GetBody() *DescribeGuestClusterNamespacesResponseBody
}

type DescribeGuestClusterNamespacesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGuestClusterNamespacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGuestClusterNamespacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGuestClusterNamespacesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGuestClusterNamespacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGuestClusterNamespacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGuestClusterNamespacesResponse) GetBody() *DescribeGuestClusterNamespacesResponseBody {
	return s.Body
}

func (s *DescribeGuestClusterNamespacesResponse) SetHeaders(v map[string]*string) *DescribeGuestClusterNamespacesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponse) SetStatusCode(v int32) *DescribeGuestClusterNamespacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGuestClusterNamespacesResponse) SetBody(v *DescribeGuestClusterNamespacesResponseBody) *DescribeGuestClusterNamespacesResponse {
	s.Body = v
	return s
}

func (s *DescribeGuestClusterNamespacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
