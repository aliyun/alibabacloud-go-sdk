// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaSettingsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKibanaSettingsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKibanaSettingsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKibanaSettingsResponseBody) *UpdateKibanaSettingsResponse
	GetBody() *UpdateKibanaSettingsResponseBody
}

type UpdateKibanaSettingsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKibanaSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKibanaSettingsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSettingsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKibanaSettingsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKibanaSettingsResponse) GetBody() *UpdateKibanaSettingsResponseBody {
	return s.Body
}

func (s *UpdateKibanaSettingsResponse) SetHeaders(v map[string]*string) *UpdateKibanaSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateKibanaSettingsResponse) SetStatusCode(v int32) *UpdateKibanaSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKibanaSettingsResponse) SetBody(v *UpdateKibanaSettingsResponseBody) *UpdateKibanaSettingsResponse {
	s.Body = v
	return s
}

func (s *UpdateKibanaSettingsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
