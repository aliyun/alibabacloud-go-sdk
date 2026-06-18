// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeQuotaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeQuotaResponseBody) *DescribeQuotaResponse
	GetBody() *DescribeQuotaResponseBody
}

type DescribeQuotaResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeQuotaResponse) GetBody() *DescribeQuotaResponseBody {
	return s.Body
}

func (s *DescribeQuotaResponse) SetHeaders(v map[string]*string) *DescribeQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeQuotaResponse) SetStatusCode(v int32) *DescribeQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQuotaResponse) SetBody(v *DescribeQuotaResponseBody) *DescribeQuotaResponse {
	s.Body = v
	return s
}

func (s *DescribeQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
