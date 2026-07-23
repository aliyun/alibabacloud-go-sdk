// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineBuildConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineBuildConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineBuildConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineBuildConfigurationResponseBody) *GetRoutineBuildConfigurationResponse
	GetBody() *GetRoutineBuildConfigurationResponseBody
}

type GetRoutineBuildConfigurationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineBuildConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineBuildConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineBuildConfigurationResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineBuildConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineBuildConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineBuildConfigurationResponse) GetBody() *GetRoutineBuildConfigurationResponseBody {
	return s.Body
}

func (s *GetRoutineBuildConfigurationResponse) SetHeaders(v map[string]*string) *GetRoutineBuildConfigurationResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineBuildConfigurationResponse) SetStatusCode(v int32) *GetRoutineBuildConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineBuildConfigurationResponse) SetBody(v *GetRoutineBuildConfigurationResponseBody) *GetRoutineBuildConfigurationResponse {
	s.Body = v
	return s
}

func (s *GetRoutineBuildConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
