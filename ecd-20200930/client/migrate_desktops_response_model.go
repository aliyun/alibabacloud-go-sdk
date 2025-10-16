// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *MigrateDesktopsResponseBody) *MigrateDesktopsResponse
	GetBody() *MigrateDesktopsResponseBody
}

type MigrateDesktopsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateDesktopsResponse) GoString() string {
	return s.String()
}

func (s *MigrateDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateDesktopsResponse) GetBody() *MigrateDesktopsResponseBody {
	return s.Body
}

func (s *MigrateDesktopsResponse) SetHeaders(v map[string]*string) *MigrateDesktopsResponse {
	s.Headers = v
	return s
}

func (s *MigrateDesktopsResponse) SetStatusCode(v int32) *MigrateDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateDesktopsResponse) SetBody(v *MigrateDesktopsResponseBody) *MigrateDesktopsResponse {
	s.Body = v
	return s
}

func (s *MigrateDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
