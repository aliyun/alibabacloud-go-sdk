// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJvmConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetJvmConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetJvmConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetJvmConfigurationResponseBody) *GetJvmConfigurationResponse
	GetBody() *GetJvmConfigurationResponseBody
}

type GetJvmConfigurationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetJvmConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetJvmConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetJvmConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetJvmConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetJvmConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetJvmConfigurationResponse) GetBody() *GetJvmConfigurationResponseBody {
	return s.Body
}

func (s *GetJvmConfigurationResponse) SetHeaders(v map[string]*string) *GetJvmConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetJvmConfigurationResponse) SetStatusCode(v int32) *GetJvmConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJvmConfigurationResponse) SetBody(v *GetJvmConfigurationResponseBody) *GetJvmConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetJvmConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
