// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceToTargetZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchInstanceToTargetZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchInstanceToTargetZoneResponse
	GetStatusCode() *int32
	SetBody(v *SwitchInstanceToTargetZoneResponseBody) *SwitchInstanceToTargetZoneResponse
	GetBody() *SwitchInstanceToTargetZoneResponseBody
}

type SwitchInstanceToTargetZoneResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchInstanceToTargetZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchInstanceToTargetZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceToTargetZoneResponse) GoString() string {
	return s.String()
}

func (s *SwitchInstanceToTargetZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchInstanceToTargetZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchInstanceToTargetZoneResponse) GetBody() *SwitchInstanceToTargetZoneResponseBody {
	return s.Body
}

func (s *SwitchInstanceToTargetZoneResponse) SetHeaders(v map[string]*string) *SwitchInstanceToTargetZoneResponse {
	s.Headers = v
	return s
}

func (s *SwitchInstanceToTargetZoneResponse) SetStatusCode(v int32) *SwitchInstanceToTargetZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchInstanceToTargetZoneResponse) SetBody(v *SwitchInstanceToTargetZoneResponseBody) *SwitchInstanceToTargetZoneResponse {
	s.Body = v
	return s
}

func (s *SwitchInstanceToTargetZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
