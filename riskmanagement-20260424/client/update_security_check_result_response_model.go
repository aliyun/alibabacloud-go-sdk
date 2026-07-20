// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityCheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityCheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityCheckResultResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityCheckResultResponseBody) *UpdateSecurityCheckResultResponse
	GetBody() *UpdateSecurityCheckResultResponseBody
}

type UpdateSecurityCheckResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityCheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityCheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityCheckResultResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityCheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityCheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityCheckResultResponse) GetBody() *UpdateSecurityCheckResultResponseBody {
	return s.Body
}

func (s *UpdateSecurityCheckResultResponse) SetHeaders(v map[string]*string) *UpdateSecurityCheckResultResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityCheckResultResponse) SetStatusCode(v int32) *UpdateSecurityCheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityCheckResultResponse) SetBody(v *UpdateSecurityCheckResultResponseBody) *UpdateSecurityCheckResultResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityCheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
