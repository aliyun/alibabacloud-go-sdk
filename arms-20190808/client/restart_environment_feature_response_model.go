// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartEnvironmentFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartEnvironmentFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartEnvironmentFeatureResponse
	GetStatusCode() *int32
	SetBody(v *RestartEnvironmentFeatureResponseBody) *RestartEnvironmentFeatureResponse
	GetBody() *RestartEnvironmentFeatureResponseBody
}

type RestartEnvironmentFeatureResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartEnvironmentFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartEnvironmentFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartEnvironmentFeatureResponse) GoString() string {
	return s.String()
}

func (s *RestartEnvironmentFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartEnvironmentFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartEnvironmentFeatureResponse) GetBody() *RestartEnvironmentFeatureResponseBody {
	return s.Body
}

func (s *RestartEnvironmentFeatureResponse) SetHeaders(v map[string]*string) *RestartEnvironmentFeatureResponse {
	s.Headers = v
	return s
}

func (s *RestartEnvironmentFeatureResponse) SetStatusCode(v int32) *RestartEnvironmentFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartEnvironmentFeatureResponse) SetBody(v *RestartEnvironmentFeatureResponseBody) *RestartEnvironmentFeatureResponse {
	s.Body = v
	return s
}

func (s *RestartEnvironmentFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
