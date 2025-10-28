// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContainerConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContainerConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContainerConfigurationResponseBody) *UpdateContainerConfigurationResponse
	GetBody() *UpdateContainerConfigurationResponseBody
}

type UpdateContainerConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContainerConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContainerConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateContainerConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContainerConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContainerConfigurationResponse) GetBody() *UpdateContainerConfigurationResponseBody {
	return s.Body
}

func (s *UpdateContainerConfigurationResponse) SetHeaders(v map[string]*string) *UpdateContainerConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateContainerConfigurationResponse) SetStatusCode(v int32) *UpdateContainerConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContainerConfigurationResponse) SetBody(v *UpdateContainerConfigurationResponseBody) *UpdateContainerConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateContainerConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
