// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchGlobalBroadcastTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SwitchGlobalBroadcastTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SwitchGlobalBroadcastTypeResponse
	GetStatusCode() *int32
	SetBody(v *SwitchGlobalBroadcastTypeResponseBody) *SwitchGlobalBroadcastTypeResponse
	GetBody() *SwitchGlobalBroadcastTypeResponseBody
}

type SwitchGlobalBroadcastTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchGlobalBroadcastTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchGlobalBroadcastTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s SwitchGlobalBroadcastTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchGlobalBroadcastTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SwitchGlobalBroadcastTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SwitchGlobalBroadcastTypeResponse) GetBody() *SwitchGlobalBroadcastTypeResponseBody {
	return s.Body
}

func (s *SwitchGlobalBroadcastTypeResponse) SetHeaders(v map[string]*string) *SwitchGlobalBroadcastTypeResponse {
	s.Headers = v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) SetStatusCode(v int32) *SwitchGlobalBroadcastTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) SetBody(v *SwitchGlobalBroadcastTypeResponseBody) *SwitchGlobalBroadcastTypeResponse {
	s.Body = v
	return s
}

func (s *SwitchGlobalBroadcastTypeResponse) Validate() error {
	return dara.Validate(s)
}
