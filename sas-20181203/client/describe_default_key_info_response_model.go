// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefaultKeyInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefaultKeyInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefaultKeyInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefaultKeyInfoResponseBody) *DescribeDefaultKeyInfoResponse
	GetBody() *DescribeDefaultKeyInfoResponseBody
}

type DescribeDefaultKeyInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefaultKeyInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefaultKeyInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefaultKeyInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefaultKeyInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefaultKeyInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefaultKeyInfoResponse) GetBody() *DescribeDefaultKeyInfoResponseBody {
	return s.Body
}

func (s *DescribeDefaultKeyInfoResponse) SetHeaders(v map[string]*string) *DescribeDefaultKeyInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefaultKeyInfoResponse) SetStatusCode(v int32) *DescribeDefaultKeyInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefaultKeyInfoResponse) SetBody(v *DescribeDefaultKeyInfoResponseBody) *DescribeDefaultKeyInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDefaultKeyInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
