// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAccessConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAccessConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAccessConfigurationResponseBody) *DeleteAccessConfigurationResponse
	GetBody() *DeleteAccessConfigurationResponseBody
}

type DeleteAccessConfigurationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAccessConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAccessConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAccessConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAccessConfigurationResponse) GetBody() *DeleteAccessConfigurationResponseBody {
	return s.Body
}

func (s *DeleteAccessConfigurationResponse) SetHeaders(v map[string]*string) *DeleteAccessConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteAccessConfigurationResponse) SetStatusCode(v int32) *DeleteAccessConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAccessConfigurationResponse) SetBody(v *DeleteAccessConfigurationResponseBody) *DeleteAccessConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteAccessConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
