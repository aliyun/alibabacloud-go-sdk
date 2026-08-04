// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSaasServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSaasServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSaasServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *CheckSaasServiceVersionResponseBody) *CheckSaasServiceVersionResponse
	GetBody() *CheckSaasServiceVersionResponseBody
}

type CheckSaasServiceVersionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSaasServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSaasServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSaasServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *CheckSaasServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSaasServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSaasServiceVersionResponse) GetBody() *CheckSaasServiceVersionResponseBody {
	return s.Body
}

func (s *CheckSaasServiceVersionResponse) SetHeaders(v map[string]*string) *CheckSaasServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *CheckSaasServiceVersionResponse) SetStatusCode(v int32) *CheckSaasServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSaasServiceVersionResponse) SetBody(v *CheckSaasServiceVersionResponseBody) *CheckSaasServiceVersionResponse {
	s.Body = v
	return s
}

func (s *CheckSaasServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
