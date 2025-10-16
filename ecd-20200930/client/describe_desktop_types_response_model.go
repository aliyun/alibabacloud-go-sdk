// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopTypesResponseBody) *DescribeDesktopTypesResponse
	GetBody() *DescribeDesktopTypesResponseBody
}

type DescribeDesktopTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopTypesResponse) GetBody() *DescribeDesktopTypesResponseBody {
	return s.Body
}

func (s *DescribeDesktopTypesResponse) SetHeaders(v map[string]*string) *DescribeDesktopTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopTypesResponse) SetStatusCode(v int32) *DescribeDesktopTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopTypesResponse) SetBody(v *DescribeDesktopTypesResponseBody) *DescribeDesktopTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
