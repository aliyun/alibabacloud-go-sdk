// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigRuntimeChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigRuntimeChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigRuntimeChannelResponse
	GetStatusCode() *int32
	SetBody(v *ConfigRuntimeChannelResponseBody) *ConfigRuntimeChannelResponse
	GetBody() *ConfigRuntimeChannelResponseBody
}

type ConfigRuntimeChannelResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigRuntimeChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigRuntimeChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeChannelResponse) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigRuntimeChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigRuntimeChannelResponse) GetBody() *ConfigRuntimeChannelResponseBody {
	return s.Body
}

func (s *ConfigRuntimeChannelResponse) SetHeaders(v map[string]*string) *ConfigRuntimeChannelResponse {
	s.Headers = v
	return s
}

func (s *ConfigRuntimeChannelResponse) SetStatusCode(v int32) *ConfigRuntimeChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigRuntimeChannelResponse) SetBody(v *ConfigRuntimeChannelResponseBody) *ConfigRuntimeChannelResponse {
	s.Body = v
	return s
}

func (s *ConfigRuntimeChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
