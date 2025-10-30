// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditConfigResponseBody) *DescribeAuditConfigResponse
	GetBody() *DescribeAuditConfigResponseBody
}

type DescribeAuditConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditConfigResponse) GetBody() *DescribeAuditConfigResponseBody {
	return s.Body
}

func (s *DescribeAuditConfigResponse) SetHeaders(v map[string]*string) *DescribeAuditConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditConfigResponse) SetStatusCode(v int32) *DescribeAuditConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditConfigResponse) SetBody(v *DescribeAuditConfigResponseBody) *DescribeAuditConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
