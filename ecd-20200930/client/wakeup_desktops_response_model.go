// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeupDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WakeupDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WakeupDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *WakeupDesktopsResponseBody) *WakeupDesktopsResponse
	GetBody() *WakeupDesktopsResponseBody
}

type WakeupDesktopsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WakeupDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WakeupDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s WakeupDesktopsResponse) GoString() string {
	return s.String()
}

func (s *WakeupDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WakeupDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WakeupDesktopsResponse) GetBody() *WakeupDesktopsResponseBody {
	return s.Body
}

func (s *WakeupDesktopsResponse) SetHeaders(v map[string]*string) *WakeupDesktopsResponse {
	s.Headers = v
	return s
}

func (s *WakeupDesktopsResponse) SetStatusCode(v int32) *WakeupDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *WakeupDesktopsResponse) SetBody(v *WakeupDesktopsResponseBody) *WakeupDesktopsResponse {
	s.Body = v
	return s
}

func (s *WakeupDesktopsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
