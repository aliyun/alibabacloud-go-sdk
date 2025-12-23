// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceConfigurationResponseBody) *GetResourceConfigurationResponse
	GetBody() *GetResourceConfigurationResponseBody
}

type GetResourceConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceConfigurationResponse) GetBody() *GetResourceConfigurationResponseBody {
	return s.Body
}

func (s *GetResourceConfigurationResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationResponse) SetStatusCode(v int32) *GetResourceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationResponse) SetBody(v *GetResourceConfigurationResponseBody) *GetResourceConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetResourceConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
