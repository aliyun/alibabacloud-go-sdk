// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableRecoveryTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableRecoveryTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableRecoveryTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableRecoveryTimeResponseBody) *DescribeAvailableRecoveryTimeResponse
	GetBody() *DescribeAvailableRecoveryTimeResponseBody
}

type DescribeAvailableRecoveryTimeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableRecoveryTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableRecoveryTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableRecoveryTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableRecoveryTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableRecoveryTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableRecoveryTimeResponse) GetBody() *DescribeAvailableRecoveryTimeResponseBody {
	return s.Body
}

func (s *DescribeAvailableRecoveryTimeResponse) SetHeaders(v map[string]*string) *DescribeAvailableRecoveryTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponse) SetStatusCode(v int32) *DescribeAvailableRecoveryTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponse) SetBody(v *DescribeAvailableRecoveryTimeResponseBody) *DescribeAvailableRecoveryTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableRecoveryTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
