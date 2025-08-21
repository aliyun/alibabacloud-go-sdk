// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLogstashSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLogstashSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLogstashSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLogstashSettingsResponseBody) *UpdateLogstashSettingsResponse
	GetBody() *UpdateLogstashSettingsResponseBody
}

type UpdateLogstashSettingsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLogstashSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLogstashSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLogstashSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateLogstashSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLogstashSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLogstashSettingsResponse) GetBody() *UpdateLogstashSettingsResponseBody {
	return s.Body
}

func (s *UpdateLogstashSettingsResponse) SetHeaders(v map[string]*string) *UpdateLogstashSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateLogstashSettingsResponse) SetStatusCode(v int32) *UpdateLogstashSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLogstashSettingsResponse) SetBody(v *UpdateLogstashSettingsResponseBody) *UpdateLogstashSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateLogstashSettingsResponse) Validate() error {
	return dara.Validate(s)
}
