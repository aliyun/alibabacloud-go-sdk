// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePartnerConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePartnerConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePartnerConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribePartnerConfigResponseBody) *DescribePartnerConfigResponse
	GetBody() *DescribePartnerConfigResponseBody
}

type DescribePartnerConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePartnerConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePartnerConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePartnerConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribePartnerConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePartnerConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePartnerConfigResponse) GetBody() *DescribePartnerConfigResponseBody {
	return s.Body
}

func (s *DescribePartnerConfigResponse) SetHeaders(v map[string]*string) *DescribePartnerConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribePartnerConfigResponse) SetStatusCode(v int32) *DescribePartnerConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePartnerConfigResponse) SetBody(v *DescribePartnerConfigResponseBody) *DescribePartnerConfigResponse {
	s.Body = v
	return s
}

func (s *DescribePartnerConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
