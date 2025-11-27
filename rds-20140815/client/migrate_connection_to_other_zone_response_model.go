// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateConnectionToOtherZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateConnectionToOtherZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateConnectionToOtherZoneResponse
	GetStatusCode() *int32
	SetBody(v *MigrateConnectionToOtherZoneResponseBody) *MigrateConnectionToOtherZoneResponse
	GetBody() *MigrateConnectionToOtherZoneResponseBody
}

type MigrateConnectionToOtherZoneResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateConnectionToOtherZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateConnectionToOtherZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateConnectionToOtherZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateConnectionToOtherZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateConnectionToOtherZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateConnectionToOtherZoneResponse) GetBody() *MigrateConnectionToOtherZoneResponseBody {
	return s.Body
}

func (s *MigrateConnectionToOtherZoneResponse) SetHeaders(v map[string]*string) *MigrateConnectionToOtherZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateConnectionToOtherZoneResponse) SetStatusCode(v int32) *MigrateConnectionToOtherZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateConnectionToOtherZoneResponse) SetBody(v *MigrateConnectionToOtherZoneResponseBody) *MigrateConnectionToOtherZoneResponse {
	s.Body = v
	return s
}

func (s *MigrateConnectionToOtherZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
