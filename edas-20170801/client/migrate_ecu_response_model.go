// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateEcuResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateEcuResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateEcuResponse
	GetStatusCode() *int32
	SetBody(v *MigrateEcuResponseBody) *MigrateEcuResponse
	GetBody() *MigrateEcuResponseBody
}

type MigrateEcuResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateEcuResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateEcuResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateEcuResponse) GoString() string {
	return s.String()
}

func (s *MigrateEcuResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateEcuResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateEcuResponse) GetBody() *MigrateEcuResponseBody {
	return s.Body
}

func (s *MigrateEcuResponse) SetHeaders(v map[string]*string) *MigrateEcuResponse {
	s.Headers = v
	return s
}

func (s *MigrateEcuResponse) SetStatusCode(v int32) *MigrateEcuResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateEcuResponse) SetBody(v *MigrateEcuResponseBody) *MigrateEcuResponse {
	s.Body = v
	return s
}

func (s *MigrateEcuResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
