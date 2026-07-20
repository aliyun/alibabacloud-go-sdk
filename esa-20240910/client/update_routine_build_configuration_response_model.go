// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineBuildConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRoutineBuildConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRoutineBuildConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRoutineBuildConfigurationResponseBody) *UpdateRoutineBuildConfigurationResponse
	GetBody() *UpdateRoutineBuildConfigurationResponseBody
}

type UpdateRoutineBuildConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRoutineBuildConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRoutineBuildConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineBuildConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoutineBuildConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRoutineBuildConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRoutineBuildConfigurationResponse) GetBody() *UpdateRoutineBuildConfigurationResponseBody {
	return s.Body
}

func (s *UpdateRoutineBuildConfigurationResponse) SetHeaders(v map[string]*string) *UpdateRoutineBuildConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoutineBuildConfigurationResponse) SetStatusCode(v int32) *UpdateRoutineBuildConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoutineBuildConfigurationResponse) SetBody(v *UpdateRoutineBuildConfigurationResponseBody) *UpdateRoutineBuildConfigurationResponse {
	s.Body = v
	return s
}

func (s *UpdateRoutineBuildConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
