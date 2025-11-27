// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocalAvailableRecoveryTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLocalAvailableRecoveryTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLocalAvailableRecoveryTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLocalAvailableRecoveryTimeResponseBody) *DescribeLocalAvailableRecoveryTimeResponse
	GetBody() *DescribeLocalAvailableRecoveryTimeResponseBody
}

type DescribeLocalAvailableRecoveryTimeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLocalAvailableRecoveryTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLocalAvailableRecoveryTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocalAvailableRecoveryTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) GetBody() *DescribeLocalAvailableRecoveryTimeResponseBody {
	return s.Body
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) SetHeaders(v map[string]*string) *DescribeLocalAvailableRecoveryTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) SetStatusCode(v int32) *DescribeLocalAvailableRecoveryTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) SetBody(v *DescribeLocalAvailableRecoveryTimeResponseBody) *DescribeLocalAvailableRecoveryTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeLocalAvailableRecoveryTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
