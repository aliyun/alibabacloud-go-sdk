// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetParseSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetParseSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetParseSettingsResponseBody) *GetParseSettingsResponse
	GetBody() *GetParseSettingsResponseBody
}

type GetParseSettingsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParseSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParseSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetParseSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetParseSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetParseSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetParseSettingsResponse) GetBody() *GetParseSettingsResponseBody {
	return s.Body
}

func (s *GetParseSettingsResponse) SetHeaders(v map[string]*string) *GetParseSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetParseSettingsResponse) SetStatusCode(v int32) *GetParseSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParseSettingsResponse) SetBody(v *GetParseSettingsResponseBody) *GetParseSettingsResponse {
	s.Body = v
	return s
}

func (s *GetParseSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
