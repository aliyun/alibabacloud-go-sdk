// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppKeyResponseBody) *DescribeAppKeyResponse
	GetBody() *DescribeAppKeyResponseBody
}

type DescribeAppKeyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppKeyResponse) GetBody() *DescribeAppKeyResponseBody {
	return s.Body
}

func (s *DescribeAppKeyResponse) SetHeaders(v map[string]*string) *DescribeAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppKeyResponse) SetStatusCode(v int32) *DescribeAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppKeyResponse) SetBody(v *DescribeAppKeyResponseBody) *DescribeAppKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeAppKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
