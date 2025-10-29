// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOffZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TurnOffZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TurnOffZoneResponse
	GetStatusCode() *int32
	SetBody(v *TurnOffZoneResponseBody) *TurnOffZoneResponse
	GetBody() *TurnOffZoneResponseBody
}

type TurnOffZoneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TurnOffZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TurnOffZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s TurnOffZoneResponse) GoString() string {
	return s.String()
}

func (s *TurnOffZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TurnOffZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TurnOffZoneResponse) GetBody() *TurnOffZoneResponseBody {
	return s.Body
}

func (s *TurnOffZoneResponse) SetHeaders(v map[string]*string) *TurnOffZoneResponse {
	s.Headers = v
	return s
}

func (s *TurnOffZoneResponse) SetStatusCode(v int32) *TurnOffZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *TurnOffZoneResponse) SetBody(v *TurnOffZoneResponseBody) *TurnOffZoneResponse {
	s.Body = v
	return s
}

func (s *TurnOffZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
