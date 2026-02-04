// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceControlConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceControlConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceControlConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceControlConfigurationResponseBody) *GetInstanceControlConfigurationResponse
	GetBody() *GetInstanceControlConfigurationResponseBody
}

type GetInstanceControlConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceControlConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceControlConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceControlConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceControlConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceControlConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceControlConfigurationResponse) GetBody() *GetInstanceControlConfigurationResponseBody {
	return s.Body
}

func (s *GetInstanceControlConfigurationResponse) SetHeaders(v map[string]*string) *GetInstanceControlConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceControlConfigurationResponse) SetStatusCode(v int32) *GetInstanceControlConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceControlConfigurationResponse) SetBody(v *GetInstanceControlConfigurationResponseBody) *GetInstanceControlConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetInstanceControlConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
