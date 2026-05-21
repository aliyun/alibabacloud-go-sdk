// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateInstanceZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateInstanceZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateInstanceZoneResponse
	GetStatusCode() *int32
	SetBody(v *MigrateInstanceZoneResponseBody) *MigrateInstanceZoneResponse
	GetBody() *MigrateInstanceZoneResponseBody
}

type MigrateInstanceZoneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateInstanceZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateInstanceZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateInstanceZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateInstanceZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateInstanceZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateInstanceZoneResponse) GetBody() *MigrateInstanceZoneResponseBody {
	return s.Body
}

func (s *MigrateInstanceZoneResponse) SetHeaders(v map[string]*string) *MigrateInstanceZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateInstanceZoneResponse) SetStatusCode(v int32) *MigrateInstanceZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateInstanceZoneResponse) SetBody(v *MigrateInstanceZoneResponseBody) *MigrateInstanceZoneResponse {
	s.Body = v
	return s
}

func (s *MigrateInstanceZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
