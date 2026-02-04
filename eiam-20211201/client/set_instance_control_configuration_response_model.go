// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceControlConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetInstanceControlConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetInstanceControlConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *SetInstanceControlConfigurationResponseBody) *SetInstanceControlConfigurationResponse
	GetBody() *SetInstanceControlConfigurationResponseBody
}

type SetInstanceControlConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstanceControlConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstanceControlConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceControlConfigurationResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceControlConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetInstanceControlConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetInstanceControlConfigurationResponse) GetBody() *SetInstanceControlConfigurationResponseBody {
	return s.Body
}

func (s *SetInstanceControlConfigurationResponse) SetHeaders(v map[string]*string) *SetInstanceControlConfigurationResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceControlConfigurationResponse) SetStatusCode(v int32) *SetInstanceControlConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceControlConfigurationResponse) SetBody(v *SetInstanceControlConfigurationResponseBody) *SetInstanceControlConfigurationResponse {
	s.Body = v
	return s
}

func (s *SetInstanceControlConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
