// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCubeBySqlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCubeBySqlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCubeBySqlResponse
	GetStatusCode() *int32
	SetBody(v *CreateCubeBySqlResponseBody) *CreateCubeBySqlResponse
	GetBody() *CreateCubeBySqlResponseBody
}

type CreateCubeBySqlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCubeBySqlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCubeBySqlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCubeBySqlResponse) GoString() string {
	return s.String()
}

func (s *CreateCubeBySqlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCubeBySqlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCubeBySqlResponse) GetBody() *CreateCubeBySqlResponseBody {
	return s.Body
}

func (s *CreateCubeBySqlResponse) SetHeaders(v map[string]*string) *CreateCubeBySqlResponse {
	s.Headers = v
	return s
}

func (s *CreateCubeBySqlResponse) SetStatusCode(v int32) *CreateCubeBySqlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCubeBySqlResponse) SetBody(v *CreateCubeBySqlResponseBody) *CreateCubeBySqlResponse {
	s.Body = v
	return s
}

func (s *CreateCubeBySqlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
