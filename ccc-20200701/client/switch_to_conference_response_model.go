// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchToConferenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchToConferenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchToConferenceResponse
	GetStatusCode() *int32
	SetBody(v *SwitchToConferenceResponseBody) *SwitchToConferenceResponse
	GetBody() *SwitchToConferenceResponseBody
}

type SwitchToConferenceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchToConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchToConferenceResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchToConferenceResponse) GoString() string {
	return s.String()
}

func (s *SwitchToConferenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchToConferenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchToConferenceResponse) GetBody() *SwitchToConferenceResponseBody {
	return s.Body
}

func (s *SwitchToConferenceResponse) SetHeaders(v map[string]*string) *SwitchToConferenceResponse {
	s.Headers = v
	return s
}

func (s *SwitchToConferenceResponse) SetStatusCode(v int32) *SwitchToConferenceResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchToConferenceResponse) SetBody(v *SwitchToConferenceResponseBody) *SwitchToConferenceResponse {
	s.Body = v
	return s
}

func (s *SwitchToConferenceResponse) Validate() error {
	return dara.Validate(s)
}
