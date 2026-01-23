// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityClassifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityClassifyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityClassifyResponseBody) *UpdateSecurityClassifyResponse
	GetBody() *UpdateSecurityClassifyResponseBody
}

type UpdateSecurityClassifyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityClassifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityClassifyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityClassifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityClassifyResponse) GetBody() *UpdateSecurityClassifyResponseBody {
	return s.Body
}

func (s *UpdateSecurityClassifyResponse) SetHeaders(v map[string]*string) *UpdateSecurityClassifyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityClassifyResponse) SetStatusCode(v int32) *UpdateSecurityClassifyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityClassifyResponse) SetBody(v *UpdateSecurityClassifyResponseBody) *UpdateSecurityClassifyResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityClassifyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
