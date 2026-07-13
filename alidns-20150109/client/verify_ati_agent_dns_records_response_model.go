// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyAtiAgentDnsRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyAtiAgentDnsRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyAtiAgentDnsRecordsResponse
	GetStatusCode() *int32
	SetBody(v *VerifyAtiAgentDnsRecordsResponseBody) *VerifyAtiAgentDnsRecordsResponse
	GetBody() *VerifyAtiAgentDnsRecordsResponseBody
}

type VerifyAtiAgentDnsRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyAtiAgentDnsRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyAtiAgentDnsRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyAtiAgentDnsRecordsResponse) GoString() string {
	return s.String()
}

func (s *VerifyAtiAgentDnsRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyAtiAgentDnsRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyAtiAgentDnsRecordsResponse) GetBody() *VerifyAtiAgentDnsRecordsResponseBody {
	return s.Body
}

func (s *VerifyAtiAgentDnsRecordsResponse) SetHeaders(v map[string]*string) *VerifyAtiAgentDnsRecordsResponse {
	s.Headers = v
	return s
}

func (s *VerifyAtiAgentDnsRecordsResponse) SetStatusCode(v int32) *VerifyAtiAgentDnsRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyAtiAgentDnsRecordsResponse) SetBody(v *VerifyAtiAgentDnsRecordsResponseBody) *VerifyAtiAgentDnsRecordsResponse {
	s.Body = v
	return s
}

func (s *VerifyAtiAgentDnsRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
