// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySAGAdminPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySAGAdminPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySAGAdminPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifySAGAdminPasswordResponseBody) *ModifySAGAdminPasswordResponse
	GetBody() *ModifySAGAdminPasswordResponseBody
}

type ModifySAGAdminPasswordResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySAGAdminPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySAGAdminPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySAGAdminPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifySAGAdminPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySAGAdminPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySAGAdminPasswordResponse) GetBody() *ModifySAGAdminPasswordResponseBody {
	return s.Body
}

func (s *ModifySAGAdminPasswordResponse) SetHeaders(v map[string]*string) *ModifySAGAdminPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifySAGAdminPasswordResponse) SetStatusCode(v int32) *ModifySAGAdminPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySAGAdminPasswordResponse) SetBody(v *ModifySAGAdminPasswordResponseBody) *ModifySAGAdminPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifySAGAdminPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
