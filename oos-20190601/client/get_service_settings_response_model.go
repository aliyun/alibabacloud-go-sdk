// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceSettingsResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceSettingsResponseBody) *GetServiceSettingsResponse
	GetBody() *GetServiceSettingsResponseBody
}

type GetServiceSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceSettingsResponse) GetBody() *GetServiceSettingsResponseBody {
	return s.Body
}

func (s *GetServiceSettingsResponse) SetHeaders(v map[string]*string) *GetServiceSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceSettingsResponse) SetStatusCode(v int32) *GetServiceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceSettingsResponse) SetBody(v *GetServiceSettingsResponseBody) *GetServiceSettingsResponse {
	s.Body = v
	return s
}

func (s *GetServiceSettingsResponse) Validate() error {
	return dara.Validate(s)
}
