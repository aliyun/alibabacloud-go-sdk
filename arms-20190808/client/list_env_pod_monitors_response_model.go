// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvPodMonitorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvPodMonitorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvPodMonitorsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvPodMonitorsResponseBody) *ListEnvPodMonitorsResponse
	GetBody() *ListEnvPodMonitorsResponseBody
}

type ListEnvPodMonitorsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvPodMonitorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvPodMonitorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvPodMonitorsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvPodMonitorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvPodMonitorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvPodMonitorsResponse) GetBody() *ListEnvPodMonitorsResponseBody {
	return s.Body
}

func (s *ListEnvPodMonitorsResponse) SetHeaders(v map[string]*string) *ListEnvPodMonitorsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvPodMonitorsResponse) SetStatusCode(v int32) *ListEnvPodMonitorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvPodMonitorsResponse) SetBody(v *ListEnvPodMonitorsResponseBody) *ListEnvPodMonitorsResponse {
	s.Body = v
	return s
}

func (s *ListEnvPodMonitorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
