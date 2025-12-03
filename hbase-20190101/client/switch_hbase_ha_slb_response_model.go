// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchHbaseHaSlbResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchHbaseHaSlbResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchHbaseHaSlbResponse
	GetStatusCode() *int32
	SetBody(v *SwitchHbaseHaSlbResponseBody) *SwitchHbaseHaSlbResponse
	GetBody() *SwitchHbaseHaSlbResponseBody
}

type SwitchHbaseHaSlbResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchHbaseHaSlbResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchHbaseHaSlbResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchHbaseHaSlbResponse) GoString() string {
	return s.String()
}

func (s *SwitchHbaseHaSlbResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchHbaseHaSlbResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchHbaseHaSlbResponse) GetBody() *SwitchHbaseHaSlbResponseBody {
	return s.Body
}

func (s *SwitchHbaseHaSlbResponse) SetHeaders(v map[string]*string) *SwitchHbaseHaSlbResponse {
	s.Headers = v
	return s
}

func (s *SwitchHbaseHaSlbResponse) SetStatusCode(v int32) *SwitchHbaseHaSlbResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchHbaseHaSlbResponse) SetBody(v *SwitchHbaseHaSlbResponseBody) *SwitchHbaseHaSlbResponse {
	s.Body = v
	return s
}

func (s *SwitchHbaseHaSlbResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
