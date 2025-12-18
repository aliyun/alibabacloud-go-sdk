// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCompliancePackResponseBody) *UpdateCompliancePackResponse
	GetBody() *UpdateCompliancePackResponseBody
}

type UpdateCompliancePackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCompliancePackResponse) GetBody() *UpdateCompliancePackResponseBody {
	return s.Body
}

func (s *UpdateCompliancePackResponse) SetHeaders(v map[string]*string) *UpdateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *UpdateCompliancePackResponse) SetStatusCode(v int32) *UpdateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCompliancePackResponse) SetBody(v *UpdateCompliancePackResponseBody) *UpdateCompliancePackResponse {
	s.Body = v
	return s
}

func (s *UpdateCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
