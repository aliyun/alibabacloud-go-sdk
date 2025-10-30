// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceAppKeyResponseBody) *DescribeServiceAppKeyResponse
	GetBody() *DescribeServiceAppKeyResponseBody
}

type DescribeServiceAppKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceAppKeyResponse) GetBody() *DescribeServiceAppKeyResponseBody {
	return s.Body
}

func (s *DescribeServiceAppKeyResponse) SetHeaders(v map[string]*string) *DescribeServiceAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceAppKeyResponse) SetStatusCode(v int32) *DescribeServiceAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceAppKeyResponse) SetBody(v *DescribeServiceAppKeyResponseBody) *DescribeServiceAppKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceAppKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
