// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKeyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKeyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKeyVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKeyVersionResponseBody) *DescribeKeyVersionResponse
	GetBody() *DescribeKeyVersionResponseBody
}

type DescribeKeyVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKeyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKeyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKeyVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKeyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKeyVersionResponse) GetBody() *DescribeKeyVersionResponseBody {
	return s.Body
}

func (s *DescribeKeyVersionResponse) SetHeaders(v map[string]*string) *DescribeKeyVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeKeyVersionResponse) SetStatusCode(v int32) *DescribeKeyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKeyVersionResponse) SetBody(v *DescribeKeyVersionResponseBody) *DescribeKeyVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeKeyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
