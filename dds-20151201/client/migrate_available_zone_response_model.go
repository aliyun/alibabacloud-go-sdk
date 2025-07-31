// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateAvailableZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateAvailableZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateAvailableZoneResponse
	GetStatusCode() *int32
	SetBody(v *MigrateAvailableZoneResponseBody) *MigrateAvailableZoneResponse
	GetBody() *MigrateAvailableZoneResponseBody
}

type MigrateAvailableZoneResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateAvailableZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateAvailableZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateAvailableZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateAvailableZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateAvailableZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateAvailableZoneResponse) GetBody() *MigrateAvailableZoneResponseBody {
	return s.Body
}

func (s *MigrateAvailableZoneResponse) SetHeaders(v map[string]*string) *MigrateAvailableZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateAvailableZoneResponse) SetStatusCode(v int32) *MigrateAvailableZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateAvailableZoneResponse) SetBody(v *MigrateAvailableZoneResponseBody) *MigrateAvailableZoneResponse {
	s.Body = v
	return s
}

func (s *MigrateAvailableZoneResponse) Validate() error {
	return dara.Validate(s)
}
