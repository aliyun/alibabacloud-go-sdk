// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetAccessConfigurationResponseBody) *GetAccessConfigurationResponse
	GetBody() *GetAccessConfigurationResponseBody
}

type GetAccessConfigurationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccessConfigurationResponse) GetBody() *GetAccessConfigurationResponseBody {
	return s.Body
}

func (s *GetAccessConfigurationResponse) SetHeaders(v map[string]*string) *GetAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetAccessConfigurationResponse) SetStatusCode(v int32) *GetAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccessConfigurationResponse) SetBody(v *GetAccessConfigurationResponseBody) *GetAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
