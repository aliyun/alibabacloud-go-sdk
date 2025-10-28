// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHotKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHotKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHotKeysResponseBody) *DescribeHotKeysResponse
	GetBody() *DescribeHotKeysResponseBody
}

type DescribeHotKeysResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHotKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHotKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHotKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHotKeysResponse) GetBody() *DescribeHotKeysResponseBody {
	return s.Body
}

func (s *DescribeHotKeysResponse) SetHeaders(v map[string]*string) *DescribeHotKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotKeysResponse) SetStatusCode(v int32) *DescribeHotKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHotKeysResponse) SetBody(v *DescribeHotKeysResponseBody) *DescribeHotKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeHotKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
