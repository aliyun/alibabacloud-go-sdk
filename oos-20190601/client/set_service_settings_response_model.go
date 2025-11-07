// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetServiceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetServiceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetServiceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *SetServiceSettingsResponseBody) *SetServiceSettingsResponse
	GetBody() *SetServiceSettingsResponseBody
}

type SetServiceSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetServiceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetServiceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetServiceSettingsResponse) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetServiceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetServiceSettingsResponse) GetBody() *SetServiceSettingsResponseBody {
	return s.Body
}

func (s *SetServiceSettingsResponse) SetHeaders(v map[string]*string) *SetServiceSettingsResponse {
	s.Headers = v
	return s
}

func (s *SetServiceSettingsResponse) SetStatusCode(v int32) *SetServiceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetServiceSettingsResponse) SetBody(v *SetServiceSettingsResponseBody) *SetServiceSettingsResponse {
	s.Body = v
	return s
}

func (s *SetServiceSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
