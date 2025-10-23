// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCreateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetCreateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetCreateResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetCreateResponseBody) *ConfigSetCreateResponse
	GetBody() *ConfigSetCreateResponseBody
}

type ConfigSetCreateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetCreateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetCreateResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCreateResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetCreateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetCreateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetCreateResponse) GetBody() *ConfigSetCreateResponseBody {
	return s.Body
}

func (s *ConfigSetCreateResponse) SetHeaders(v map[string]*string) *ConfigSetCreateResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetCreateResponse) SetStatusCode(v int32) *ConfigSetCreateResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetCreateResponse) SetBody(v *ConfigSetCreateResponseBody) *ConfigSetCreateResponse {
	s.Body = v
	return s
}

func (s *ConfigSetCreateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
