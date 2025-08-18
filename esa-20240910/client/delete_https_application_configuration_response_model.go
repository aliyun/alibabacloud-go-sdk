// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpsApplicationConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHttpsApplicationConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHttpsApplicationConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHttpsApplicationConfigurationResponseBody) *DeleteHttpsApplicationConfigurationResponse
	GetBody() *DeleteHttpsApplicationConfigurationResponseBody
}

type DeleteHttpsApplicationConfigurationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHttpsApplicationConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHttpsApplicationConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpsApplicationConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteHttpsApplicationConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHttpsApplicationConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHttpsApplicationConfigurationResponse) GetBody() *DeleteHttpsApplicationConfigurationResponseBody {
	return s.Body
}

func (s *DeleteHttpsApplicationConfigurationResponse) SetHeaders(v map[string]*string) *DeleteHttpsApplicationConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteHttpsApplicationConfigurationResponse) SetStatusCode(v int32) *DeleteHttpsApplicationConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHttpsApplicationConfigurationResponse) SetBody(v *DeleteHttpsApplicationConfigurationResponseBody) *DeleteHttpsApplicationConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteHttpsApplicationConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
