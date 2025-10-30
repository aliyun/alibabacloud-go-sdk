// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchToOnlineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchToOnlineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchToOnlineResponse
	GetStatusCode() *int32
	SetBody(v *SwitchToOnlineResponseBody) *SwitchToOnlineResponse
	GetBody() *SwitchToOnlineResponseBody
}

type SwitchToOnlineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchToOnlineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchToOnlineResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchToOnlineResponse) GoString() string {
	return s.String()
}

func (s *SwitchToOnlineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchToOnlineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchToOnlineResponse) GetBody() *SwitchToOnlineResponseBody {
	return s.Body
}

func (s *SwitchToOnlineResponse) SetHeaders(v map[string]*string) *SwitchToOnlineResponse {
	s.Headers = v
	return s
}

func (s *SwitchToOnlineResponse) SetStatusCode(v int32) *SwitchToOnlineResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchToOnlineResponse) SetBody(v *SwitchToOnlineResponseBody) *SwitchToOnlineResponse {
	s.Body = v
	return s
}

func (s *SwitchToOnlineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
