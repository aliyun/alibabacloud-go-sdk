// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigSetDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigSetDeleteResponse
	GetStatusCode() *int32
	SetBody(v *ConfigSetDeleteResponseBody) *ConfigSetDeleteResponse
	GetBody() *ConfigSetDeleteResponseBody
}

type ConfigSetDeleteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigSetDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigSetDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDeleteResponse) GoString() string {
	return s.String()
}

func (s *ConfigSetDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigSetDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigSetDeleteResponse) GetBody() *ConfigSetDeleteResponseBody {
	return s.Body
}

func (s *ConfigSetDeleteResponse) SetHeaders(v map[string]*string) *ConfigSetDeleteResponse {
	s.Headers = v
	return s
}

func (s *ConfigSetDeleteResponse) SetStatusCode(v int32) *ConfigSetDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigSetDeleteResponse) SetBody(v *ConfigSetDeleteResponseBody) *ConfigSetDeleteResponse {
	s.Body = v
	return s
}

func (s *ConfigSetDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
