// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetListResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetListResponseBody) *ConfigSetListResponse
	GetBody() *ConfigSetListResponseBody
}

type ConfigSetListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetListResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetListResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetListResponse) GetBody() *ConfigSetListResponseBody {
	return s.Body
}

func (s *ConfigSetListResponse) SetHeaders(v map[string]*string) *ConfigSetListResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetListResponse) SetStatusCode(v int32) *ConfigSetListResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetListResponse) SetBody(v *ConfigSetListResponseBody) *ConfigSetListResponse {
	s.Body = v
	return s
}

func (s *ConfigSetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
