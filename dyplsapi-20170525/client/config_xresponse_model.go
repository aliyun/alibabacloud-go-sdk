// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigXResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigXResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigXResponse
	GetStatusCode() *int32
	SetBody(v *ConfigXResponseBody) *ConfigXResponse
	GetBody() *ConfigXResponseBody
}

type ConfigXResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigXResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigXResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigXResponse) GoString() string {
	return s.String()
}

func (s *ConfigXResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigXResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigXResponse) GetBody() *ConfigXResponseBody {
	return s.Body
}

func (s *ConfigXResponse) SetHeaders(v map[string]*string) *ConfigXResponse {
	s.Headers = v
	return s
}

func (s *ConfigXResponse) SetStatusCode(v int32) *ConfigXResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigXResponse) SetBody(v *ConfigXResponseBody) *ConfigXResponse {
	s.Body = v
	return s
}

func (s *ConfigXResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
