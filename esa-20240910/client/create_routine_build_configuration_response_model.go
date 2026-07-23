// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineBuildConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineBuildConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineBuildConfigurationResponseBody) *CreateRoutineBuildConfigurationResponse
	GetBody() *CreateRoutineBuildConfigurationResponseBody
}

type CreateRoutineBuildConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineBuildConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineBuildConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineBuildConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineBuildConfigurationResponse) GetBody() *CreateRoutineBuildConfigurationResponseBody {
	return s.Body
}

func (s *CreateRoutineBuildConfigurationResponse) SetHeaders(v map[string]*string) *CreateRoutineBuildConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineBuildConfigurationResponse) SetStatusCode(v int32) *CreateRoutineBuildConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineBuildConfigurationResponse) SetBody(v *CreateRoutineBuildConfigurationResponseBody) *CreateRoutineBuildConfigurationResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineBuildConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
