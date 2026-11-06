// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigNetStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigNetStatusResponse
	GetStatusCode() *int32
	SetBody(v *ConfigNetStatusResponseBody) *ConfigNetStatusResponse
	GetBody() *ConfigNetStatusResponseBody
}

type ConfigNetStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigNetStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigNetStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetStatusResponse) GoString() string {
	return s.String()
}

func (s *ConfigNetStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigNetStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigNetStatusResponse) GetBody() *ConfigNetStatusResponseBody {
	return s.Body
}

func (s *ConfigNetStatusResponse) SetHeaders(v map[string]*string) *ConfigNetStatusResponse {
	s.Headers = v
	return s
}

func (s *ConfigNetStatusResponse) SetStatusCode(v int32) *ConfigNetStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigNetStatusResponse) SetBody(v *ConfigNetStatusResponseBody) *ConfigNetStatusResponse {
	s.Body = v
	return s
}

func (s *ConfigNetStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
