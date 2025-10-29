// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTurnOnZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TurnOnZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TurnOnZoneResponse
	GetStatusCode() *int32
	SetBody(v *TurnOnZoneResponseBody) *TurnOnZoneResponse
	GetBody() *TurnOnZoneResponseBody
}

type TurnOnZoneResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TurnOnZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TurnOnZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s TurnOnZoneResponse) GoString() string {
	return s.String()
}

func (s *TurnOnZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TurnOnZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TurnOnZoneResponse) GetBody() *TurnOnZoneResponseBody {
	return s.Body
}

func (s *TurnOnZoneResponse) SetHeaders(v map[string]*string) *TurnOnZoneResponse {
	s.Headers = v
	return s
}

func (s *TurnOnZoneResponse) SetStatusCode(v int32) *TurnOnZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *TurnOnZoneResponse) SetBody(v *TurnOnZoneResponseBody) *TurnOnZoneResponse {
	s.Body = v
	return s
}

func (s *TurnOnZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
