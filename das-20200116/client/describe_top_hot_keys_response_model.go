// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTopHotKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTopHotKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTopHotKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTopHotKeysResponseBody) *DescribeTopHotKeysResponse
	GetBody() *DescribeTopHotKeysResponseBody
}

type DescribeTopHotKeysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTopHotKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTopHotKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTopHotKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopHotKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTopHotKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTopHotKeysResponse) GetBody() *DescribeTopHotKeysResponseBody {
	return s.Body
}

func (s *DescribeTopHotKeysResponse) SetHeaders(v map[string]*string) *DescribeTopHotKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopHotKeysResponse) SetStatusCode(v int32) *DescribeTopHotKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopHotKeysResponse) SetBody(v *DescribeTopHotKeysResponseBody) *DescribeTopHotKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeTopHotKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
