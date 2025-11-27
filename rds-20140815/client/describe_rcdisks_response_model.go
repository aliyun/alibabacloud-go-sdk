// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCDisksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCDisksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCDisksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCDisksResponseBody) *DescribeRCDisksResponse
	GetBody() *DescribeRCDisksResponseBody
}

type DescribeRCDisksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCDisksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCDisksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCDisksResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCDisksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCDisksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCDisksResponse) GetBody() *DescribeRCDisksResponseBody {
	return s.Body
}

func (s *DescribeRCDisksResponse) SetHeaders(v map[string]*string) *DescribeRCDisksResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCDisksResponse) SetStatusCode(v int32) *DescribeRCDisksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCDisksResponse) SetBody(v *DescribeRCDisksResponseBody) *DescribeRCDisksResponse {
	s.Body = v
	return s
}

func (s *DescribeRCDisksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
