// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpsApplicationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpsApplicationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpsApplicationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpsApplicationConfigurationResponseBody) *GetHttpsApplicationConfigurationResponse
	GetBody() *GetHttpsApplicationConfigurationResponseBody
}

type GetHttpsApplicationConfigurationResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpsApplicationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpsApplicationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpsApplicationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetHttpsApplicationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpsApplicationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpsApplicationConfigurationResponse) GetBody() *GetHttpsApplicationConfigurationResponseBody {
	return s.Body
}

func (s *GetHttpsApplicationConfigurationResponse) SetHeaders(v map[string]*string) *GetHttpsApplicationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetHttpsApplicationConfigurationResponse) SetStatusCode(v int32) *GetHttpsApplicationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpsApplicationConfigurationResponse) SetBody(v *GetHttpsApplicationConfigurationResponseBody) *GetHttpsApplicationConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetHttpsApplicationConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
