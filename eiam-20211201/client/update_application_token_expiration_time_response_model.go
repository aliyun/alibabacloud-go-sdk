// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTokenExpirationTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationTokenExpirationTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationTokenExpirationTimeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationTokenExpirationTimeResponseBody) *UpdateApplicationTokenExpirationTimeResponse
	GetBody() *UpdateApplicationTokenExpirationTimeResponseBody
}

type UpdateApplicationTokenExpirationTimeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationTokenExpirationTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationTokenExpirationTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTokenExpirationTimeResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTokenExpirationTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationTokenExpirationTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationTokenExpirationTimeResponse) GetBody() *UpdateApplicationTokenExpirationTimeResponseBody {
	return s.Body
}

func (s *UpdateApplicationTokenExpirationTimeResponse) SetHeaders(v map[string]*string) *UpdateApplicationTokenExpirationTimeResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeResponse) SetStatusCode(v int32) *UpdateApplicationTokenExpirationTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeResponse) SetBody(v *UpdateApplicationTokenExpirationTimeResponseBody) *UpdateApplicationTokenExpirationTimeResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationTokenExpirationTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
