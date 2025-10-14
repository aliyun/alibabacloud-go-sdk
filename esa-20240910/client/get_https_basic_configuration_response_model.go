// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpsBasicConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpsBasicConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpsBasicConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpsBasicConfigurationResponseBody) *GetHttpsBasicConfigurationResponse
	GetBody() *GetHttpsBasicConfigurationResponseBody
}

type GetHttpsBasicConfigurationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpsBasicConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpsBasicConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpsBasicConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetHttpsBasicConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpsBasicConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpsBasicConfigurationResponse) GetBody() *GetHttpsBasicConfigurationResponseBody {
	return s.Body
}

func (s *GetHttpsBasicConfigurationResponse) SetHeaders(v map[string]*string) *GetHttpsBasicConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetHttpsBasicConfigurationResponse) SetStatusCode(v int32) *GetHttpsBasicConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponse) SetBody(v *GetHttpsBasicConfigurationResponseBody) *GetHttpsBasicConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetHttpsBasicConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
