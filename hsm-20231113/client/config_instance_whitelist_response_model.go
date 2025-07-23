// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigInstanceWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigInstanceWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ConfigInstanceWhitelistResponseBody) *ConfigInstanceWhitelistResponse
	GetBody() *ConfigInstanceWhitelistResponseBody
}

type ConfigInstanceWhitelistResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigInstanceWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigInstanceWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigInstanceWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigInstanceWhitelistResponse) GetBody() *ConfigInstanceWhitelistResponseBody {
	return s.Body
}

func (s *ConfigInstanceWhitelistResponse) SetHeaders(v map[string]*string) *ConfigInstanceWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ConfigInstanceWhitelistResponse) SetStatusCode(v int32) *ConfigInstanceWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigInstanceWhitelistResponse) SetBody(v *ConfigInstanceWhitelistResponseBody) *ConfigInstanceWhitelistResponse {
	s.Body = v
	return s
}

func (s *ConfigInstanceWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
