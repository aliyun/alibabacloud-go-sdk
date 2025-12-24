// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTerminalSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTerminalSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTerminalSettingsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTerminalSettingsResponseBody) *ModifyTerminalSettingsResponse
	GetBody() *ModifyTerminalSettingsResponseBody
}

type ModifyTerminalSettingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTerminalSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTerminalSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTerminalSettingsResponse) GoString() string {
	return s.String()
}

func (s *ModifyTerminalSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTerminalSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTerminalSettingsResponse) GetBody() *ModifyTerminalSettingsResponseBody {
	return s.Body
}

func (s *ModifyTerminalSettingsResponse) SetHeaders(v map[string]*string) *ModifyTerminalSettingsResponse {
	s.Headers = v
	return s
}

func (s *ModifyTerminalSettingsResponse) SetStatusCode(v int32) *ModifyTerminalSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTerminalSettingsResponse) SetBody(v *ModifyTerminalSettingsResponseBody) *ModifyTerminalSettingsResponse {
	s.Body = v
	return s
}

func (s *ModifyTerminalSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
