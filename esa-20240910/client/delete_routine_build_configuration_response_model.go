// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineBuildConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRoutineBuildConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRoutineBuildConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRoutineBuildConfigurationResponseBody) *DeleteRoutineBuildConfigurationResponse
	GetBody() *DeleteRoutineBuildConfigurationResponseBody
}

type DeleteRoutineBuildConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRoutineBuildConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRoutineBuildConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineBuildConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoutineBuildConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRoutineBuildConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRoutineBuildConfigurationResponse) GetBody() *DeleteRoutineBuildConfigurationResponseBody {
	return s.Body
}

func (s *DeleteRoutineBuildConfigurationResponse) SetHeaders(v map[string]*string) *DeleteRoutineBuildConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoutineBuildConfigurationResponse) SetStatusCode(v int32) *DeleteRoutineBuildConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoutineBuildConfigurationResponse) SetBody(v *DeleteRoutineBuildConfigurationResponseBody) *DeleteRoutineBuildConfigurationResponse {
	s.Body = v
	return s
}

func (s *DeleteRoutineBuildConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
