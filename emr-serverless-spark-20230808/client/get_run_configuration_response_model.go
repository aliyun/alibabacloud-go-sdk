// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRunConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRunConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetRunConfigurationResponseBody) *GetRunConfigurationResponse
	GetBody() *GetRunConfigurationResponseBody
}

type GetRunConfigurationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRunConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRunConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRunConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetRunConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRunConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRunConfigurationResponse) GetBody() *GetRunConfigurationResponseBody {
	return s.Body
}

func (s *GetRunConfigurationResponse) SetHeaders(v map[string]*string) *GetRunConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetRunConfigurationResponse) SetStatusCode(v int32) *GetRunConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRunConfigurationResponse) SetBody(v *GetRunConfigurationResponseBody) *GetRunConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetRunConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
