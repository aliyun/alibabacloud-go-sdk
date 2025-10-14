// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpsBasicConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpsBasicConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpsBasicConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpsBasicConfigurationResponseBody) *DeleteHttpsBasicConfigurationResponse
	GetBody() *DeleteHttpsBasicConfigurationResponseBody
}

type DeleteHttpsBasicConfigurationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpsBasicConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpsBasicConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpsBasicConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpsBasicConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpsBasicConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpsBasicConfigurationResponse) GetBody() *DeleteHttpsBasicConfigurationResponseBody {
	return s.Body
}

func (s *DeleteHttpsBasicConfigurationResponse) SetHeaders(v map[string]*string) *DeleteHttpsBasicConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpsBasicConfigurationResponse) SetStatusCode(v int32) *DeleteHttpsBasicConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpsBasicConfigurationResponse) SetBody(v *DeleteHttpsBasicConfigurationResponseBody) *DeleteHttpsBasicConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpsBasicConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
