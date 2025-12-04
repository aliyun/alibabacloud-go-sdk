// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchVccConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchVccConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchVccConnectionResponse
	GetStatusCode() *int32
	SetBody(v *SwitchVccConnectionResponseBody) *SwitchVccConnectionResponse
	GetBody() *SwitchVccConnectionResponseBody
}

type SwitchVccConnectionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchVccConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchVccConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchVccConnectionResponse) GoString() string {
	return s.String()
}

func (s *SwitchVccConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchVccConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchVccConnectionResponse) GetBody() *SwitchVccConnectionResponseBody {
	return s.Body
}

func (s *SwitchVccConnectionResponse) SetHeaders(v map[string]*string) *SwitchVccConnectionResponse {
	s.Headers = v
	return s
}

func (s *SwitchVccConnectionResponse) SetStatusCode(v int32) *SwitchVccConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchVccConnectionResponse) SetBody(v *SwitchVccConnectionResponseBody) *SwitchVccConnectionResponse {
	s.Body = v
	return s
}

func (s *SwitchVccConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
