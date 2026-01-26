// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigAppResponse
	GetStatusCode() *int32
	SetBody(v *ConfigAppResponseBody) *ConfigAppResponse
	GetBody() *ConfigAppResponseBody
}

type ConfigAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigAppResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigAppResponse) GoString() string {
	return s.String()
}

func (s *ConfigAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigAppResponse) GetBody() *ConfigAppResponseBody {
	return s.Body
}

func (s *ConfigAppResponse) SetHeaders(v map[string]*string) *ConfigAppResponse {
	s.Headers = v
	return s
}

func (s *ConfigAppResponse) SetStatusCode(v int32) *ConfigAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigAppResponse) SetBody(v *ConfigAppResponseBody) *ConfigAppResponse {
	s.Body = v
	return s
}

func (s *ConfigAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
