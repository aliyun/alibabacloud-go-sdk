// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCubeBySqlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCubeBySqlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCubeBySqlResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCubeBySqlResponseBody) *UpdateCubeBySqlResponse
	GetBody() *UpdateCubeBySqlResponseBody
}

type UpdateCubeBySqlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCubeBySqlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCubeBySqlResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCubeBySqlResponse) GoString() string {
	return s.String()
}

func (s *UpdateCubeBySqlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCubeBySqlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCubeBySqlResponse) GetBody() *UpdateCubeBySqlResponseBody {
	return s.Body
}

func (s *UpdateCubeBySqlResponse) SetHeaders(v map[string]*string) *UpdateCubeBySqlResponse {
	s.Headers = v
	return s
}

func (s *UpdateCubeBySqlResponse) SetStatusCode(v int32) *UpdateCubeBySqlResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCubeBySqlResponse) SetBody(v *UpdateCubeBySqlResponseBody) *UpdateCubeBySqlResponse {
	s.Body = v
	return s
}

func (s *UpdateCubeBySqlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
