// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateToOtherZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateToOtherZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateToOtherZoneResponse
	GetStatusCode() *int32
	SetBody(v *MigrateToOtherZoneResponseBody) *MigrateToOtherZoneResponse
	GetBody() *MigrateToOtherZoneResponseBody
}

type MigrateToOtherZoneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateToOtherZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateToOtherZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateToOtherZoneResponse) GoString() string {
	return s.String()
}

func (s *MigrateToOtherZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateToOtherZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateToOtherZoneResponse) GetBody() *MigrateToOtherZoneResponseBody {
	return s.Body
}

func (s *MigrateToOtherZoneResponse) SetHeaders(v map[string]*string) *MigrateToOtherZoneResponse {
	s.Headers = v
	return s
}

func (s *MigrateToOtherZoneResponse) SetStatusCode(v int32) *MigrateToOtherZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateToOtherZoneResponse) SetBody(v *MigrateToOtherZoneResponseBody) *MigrateToOtherZoneResponse {
	s.Body = v
	return s
}

func (s *MigrateToOtherZoneResponse) Validate() error {
	return dara.Validate(s)
}
