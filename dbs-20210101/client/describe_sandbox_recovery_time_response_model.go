// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxRecoveryTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSandboxRecoveryTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSandboxRecoveryTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSandboxRecoveryTimeResponseBody) *DescribeSandboxRecoveryTimeResponse
	GetBody() *DescribeSandboxRecoveryTimeResponseBody
}

type DescribeSandboxRecoveryTimeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSandboxRecoveryTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSandboxRecoveryTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxRecoveryTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSandboxRecoveryTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSandboxRecoveryTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSandboxRecoveryTimeResponse) GetBody() *DescribeSandboxRecoveryTimeResponseBody {
	return s.Body
}

func (s *DescribeSandboxRecoveryTimeResponse) SetHeaders(v map[string]*string) *DescribeSandboxRecoveryTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) SetStatusCode(v int32) *DescribeSandboxRecoveryTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) SetBody(v *DescribeSandboxRecoveryTimeResponseBody) *DescribeSandboxRecoveryTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeSandboxRecoveryTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
