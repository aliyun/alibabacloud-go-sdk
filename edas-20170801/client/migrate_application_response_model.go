// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *MigrateApplicationResponseBody) *MigrateApplicationResponse
	GetBody() *MigrateApplicationResponseBody
}

type MigrateApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateApplicationResponse) GoString() string {
	return s.String()
}

func (s *MigrateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateApplicationResponse) GetBody() *MigrateApplicationResponseBody {
	return s.Body
}

func (s *MigrateApplicationResponse) SetHeaders(v map[string]*string) *MigrateApplicationResponse {
	s.Headers = v
	return s
}

func (s *MigrateApplicationResponse) SetStatusCode(v int32) *MigrateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateApplicationResponse) SetBody(v *MigrateApplicationResponseBody) *MigrateApplicationResponse {
	s.Body = v
	return s
}

func (s *MigrateApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
