// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetDetailResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetDetailResponseBody) *ConfigSetDetailResponse
	GetBody() *ConfigSetDetailResponseBody
}

type ConfigSetDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDetailResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetDetailResponse) GetBody() *ConfigSetDetailResponseBody {
	return s.Body
}

func (s *ConfigSetDetailResponse) SetHeaders(v map[string]*string) *ConfigSetDetailResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetDetailResponse) SetStatusCode(v int32) *ConfigSetDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetDetailResponse) SetBody(v *ConfigSetDetailResponseBody) *ConfigSetDetailResponse {
	s.Body = v
	return s
}

func (s *ConfigSetDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
