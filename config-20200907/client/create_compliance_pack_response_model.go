// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *CreateCompliancePackResponseBody) *CreateCompliancePackResponse
	GetBody() *CreateCompliancePackResponseBody
}

type CreateCompliancePackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCompliancePackResponse) GetBody() *CreateCompliancePackResponseBody {
	return s.Body
}

func (s *CreateCompliancePackResponse) SetHeaders(v map[string]*string) *CreateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *CreateCompliancePackResponse) SetStatusCode(v int32) *CreateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCompliancePackResponse) SetBody(v *CreateCompliancePackResponseBody) *CreateCompliancePackResponse {
	s.Body = v
	return s
}

func (s *CreateCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
