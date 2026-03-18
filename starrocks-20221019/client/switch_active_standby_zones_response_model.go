// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchActiveStandbyZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchActiveStandbyZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchActiveStandbyZonesResponse
	GetStatusCode() *int32
	SetBody(v *SwitchActiveStandbyZonesResponseBody) *SwitchActiveStandbyZonesResponse
	GetBody() *SwitchActiveStandbyZonesResponseBody
}

type SwitchActiveStandbyZonesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchActiveStandbyZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchActiveStandbyZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchActiveStandbyZonesResponse) GoString() string {
	return s.String()
}

func (s *SwitchActiveStandbyZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchActiveStandbyZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchActiveStandbyZonesResponse) GetBody() *SwitchActiveStandbyZonesResponseBody {
	return s.Body
}

func (s *SwitchActiveStandbyZonesResponse) SetHeaders(v map[string]*string) *SwitchActiveStandbyZonesResponse {
	s.Headers = v
	return s
}

func (s *SwitchActiveStandbyZonesResponse) SetStatusCode(v int32) *SwitchActiveStandbyZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchActiveStandbyZonesResponse) SetBody(v *SwitchActiveStandbyZonesResponseBody) *SwitchActiveStandbyZonesResponse {
	s.Body = v
	return s
}

func (s *SwitchActiveStandbyZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
