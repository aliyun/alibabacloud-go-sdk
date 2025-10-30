// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAuditDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAuditDetailsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAuditDetailsResponseBody) *DescribeAuditDetailsResponse
	GetBody() *DescribeAuditDetailsResponseBody
}

type DescribeAuditDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAuditDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAuditDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuditDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAuditDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAuditDetailsResponse) GetBody() *DescribeAuditDetailsResponseBody {
	return s.Body
}

func (s *DescribeAuditDetailsResponse) SetHeaders(v map[string]*string) *DescribeAuditDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuditDetailsResponse) SetStatusCode(v int32) *DescribeAuditDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAuditDetailsResponse) SetBody(v *DescribeAuditDetailsResponseBody) *DescribeAuditDetailsResponse {
	s.Body = v
	return s
}

func (s *DescribeAuditDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
