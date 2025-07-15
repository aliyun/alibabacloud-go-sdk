// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStateConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateStateConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateStateConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateStateConfigurationResponseBody) *CreateStateConfigurationResponse
	GetBody() *CreateStateConfigurationResponseBody
}

type CreateStateConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStateConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateStateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateStateConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateStateConfigurationResponse) GetBody() *CreateStateConfigurationResponseBody {
	return s.Body
}

func (s *CreateStateConfigurationResponse) SetHeaders(v map[string]*string) *CreateStateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateStateConfigurationResponse) SetStatusCode(v int32) *CreateStateConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStateConfigurationResponse) SetBody(v *CreateStateConfigurationResponseBody) *CreateStateConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateStateConfigurationResponse) Validate() error {
	return dara.Validate(s)
}
