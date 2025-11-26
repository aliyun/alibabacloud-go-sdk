// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPrometheusUserSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPrometheusUserSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPrometheusUserSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetPrometheusUserSettingResponseBody) *GetPrometheusUserSettingResponse
	GetBody() *GetPrometheusUserSettingResponseBody
}

type GetPrometheusUserSettingResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrometheusUserSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrometheusUserSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPrometheusUserSettingResponse) GoString() string {
	return s.String()
}

func (s *GetPrometheusUserSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPrometheusUserSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPrometheusUserSettingResponse) GetBody() *GetPrometheusUserSettingResponseBody {
	return s.Body
}

func (s *GetPrometheusUserSettingResponse) SetHeaders(v map[string]*string) *GetPrometheusUserSettingResponse {
	s.Headers = v
	return s
}

func (s *GetPrometheusUserSettingResponse) SetStatusCode(v int32) *GetPrometheusUserSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrometheusUserSettingResponse) SetBody(v *GetPrometheusUserSettingResponseBody) *GetPrometheusUserSettingResponse {
	s.Body = v
	return s
}

func (s *GetPrometheusUserSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
