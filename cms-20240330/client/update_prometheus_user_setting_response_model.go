// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrometheusUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrometheusUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrometheusUserSettingResponseBody) *UpdatePrometheusUserSettingResponse
	GetBody() *UpdatePrometheusUserSettingResponseBody
}

type UpdatePrometheusUserSettingResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrometheusUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrometheusUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusUserSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrometheusUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrometheusUserSettingResponse) GetBody() *UpdatePrometheusUserSettingResponseBody {
	return s.Body
}

func (s *UpdatePrometheusUserSettingResponse) SetHeaders(v map[string]*string) *UpdatePrometheusUserSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrometheusUserSettingResponse) SetStatusCode(v int32) *UpdatePrometheusUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrometheusUserSettingResponse) SetBody(v *UpdatePrometheusUserSettingResponseBody) *UpdatePrometheusUserSettingResponse {
	s.Body = v
	return s
}

func (s *UpdatePrometheusUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
