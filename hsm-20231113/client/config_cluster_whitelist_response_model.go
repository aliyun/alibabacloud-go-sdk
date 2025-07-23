// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigClusterWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigClusterWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *ConfigClusterWhitelistResponseBody) *ConfigClusterWhitelistResponse
	GetBody() *ConfigClusterWhitelistResponseBody
}

type ConfigClusterWhitelistResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigClusterWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigClusterWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ConfigClusterWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigClusterWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigClusterWhitelistResponse) GetBody() *ConfigClusterWhitelistResponseBody {
	return s.Body
}

func (s *ConfigClusterWhitelistResponse) SetHeaders(v map[string]*string) *ConfigClusterWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ConfigClusterWhitelistResponse) SetStatusCode(v int32) *ConfigClusterWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigClusterWhitelistResponse) SetBody(v *ConfigClusterWhitelistResponseBody) *ConfigClusterWhitelistResponse {
	s.Body = v
	return s
}

func (s *ConfigClusterWhitelistResponse) Validate() error {
	return dara.Validate(s)
}
