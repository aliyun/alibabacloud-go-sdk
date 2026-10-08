// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenGovernanceServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenGovernanceServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenGovernanceServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenGovernanceServiceResponseBody) *OpenGovernanceServiceResponse
	GetBody() *OpenGovernanceServiceResponseBody
}

type OpenGovernanceServiceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenGovernanceServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenGovernanceServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenGovernanceServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenGovernanceServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenGovernanceServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenGovernanceServiceResponse) GetBody() *OpenGovernanceServiceResponseBody {
	return s.Body
}

func (s *OpenGovernanceServiceResponse) SetHeaders(v map[string]*string) *OpenGovernanceServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenGovernanceServiceResponse) SetStatusCode(v int32) *OpenGovernanceServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenGovernanceServiceResponse) SetBody(v *OpenGovernanceServiceResponseBody) *OpenGovernanceServiceResponse {
	s.Body = v
	return s
}

func (s *OpenGovernanceServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
