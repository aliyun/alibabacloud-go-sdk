// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventOperationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityEventOperationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityEventOperationStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityEventOperationStatusResponseBody) *DescribeSecurityEventOperationStatusResponse
	GetBody() *DescribeSecurityEventOperationStatusResponseBody
}

type DescribeSecurityEventOperationStatusResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityEventOperationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityEventOperationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventOperationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventOperationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityEventOperationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityEventOperationStatusResponse) GetBody() *DescribeSecurityEventOperationStatusResponseBody {
	return s.Body
}

func (s *DescribeSecurityEventOperationStatusResponse) SetHeaders(v map[string]*string) *DescribeSecurityEventOperationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetStatusCode(v int32) *DescribeSecurityEventOperationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) SetBody(v *DescribeSecurityEventOperationStatusResponseBody) *DescribeSecurityEventOperationStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityEventOperationStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
