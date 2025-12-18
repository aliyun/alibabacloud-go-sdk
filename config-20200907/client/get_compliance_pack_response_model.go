// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *GetCompliancePackResponseBody) *GetCompliancePackResponse
	GetBody() *GetCompliancePackResponseBody
}

type GetCompliancePackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCompliancePackResponse) GetBody() *GetCompliancePackResponseBody {
	return s.Body
}

func (s *GetCompliancePackResponse) SetHeaders(v map[string]*string) *GetCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *GetCompliancePackResponse) SetStatusCode(v int32) *GetCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompliancePackResponse) SetBody(v *GetCompliancePackResponseBody) *GetCompliancePackResponse {
	s.Body = v
	return s
}

func (s *GetCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
