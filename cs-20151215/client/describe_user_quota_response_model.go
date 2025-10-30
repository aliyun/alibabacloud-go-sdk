// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserQuotaResponseBody) *DescribeUserQuotaResponse
	GetBody() *DescribeUserQuotaResponseBody
}

type DescribeUserQuotaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserQuotaResponse) GetBody() *DescribeUserQuotaResponseBody {
	return s.Body
}

func (s *DescribeUserQuotaResponse) SetHeaders(v map[string]*string) *DescribeUserQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserQuotaResponse) SetStatusCode(v int32) *DescribeUserQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserQuotaResponse) SetBody(v *DescribeUserQuotaResponseBody) *DescribeUserQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeUserQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
