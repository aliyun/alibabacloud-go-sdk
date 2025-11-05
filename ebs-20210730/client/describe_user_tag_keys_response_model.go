// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTagKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserTagKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserTagKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserTagKeysResponseBody) *DescribeUserTagKeysResponse
	GetBody() *DescribeUserTagKeysResponseBody
}

type DescribeUserTagKeysResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserTagKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTagKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserTagKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserTagKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserTagKeysResponse) GetBody() *DescribeUserTagKeysResponseBody {
	return s.Body
}

func (s *DescribeUserTagKeysResponse) SetHeaders(v map[string]*string) *DescribeUserTagKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserTagKeysResponse) SetStatusCode(v int32) *DescribeUserTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserTagKeysResponse) SetBody(v *DescribeUserTagKeysResponseBody) *DescribeUserTagKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeUserTagKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
