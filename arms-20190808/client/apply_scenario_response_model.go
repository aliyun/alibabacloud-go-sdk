// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyScenarioResponse
	GetStatusCode() *int32
	SetBody(v *ApplyScenarioResponseBody) *ApplyScenarioResponse
	GetBody() *ApplyScenarioResponseBody
}

type ApplyScenarioResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyScenarioResponse) GoString() string {
	return s.String()
}

func (s *ApplyScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyScenarioResponse) GetBody() *ApplyScenarioResponseBody {
	return s.Body
}

func (s *ApplyScenarioResponse) SetHeaders(v map[string]*string) *ApplyScenarioResponse {
	s.Headers = v
	return s
}

func (s *ApplyScenarioResponse) SetStatusCode(v int32) *ApplyScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyScenarioResponse) SetBody(v *ApplyScenarioResponseBody) *ApplyScenarioResponse {
	s.Body = v
	return s
}

func (s *ApplyScenarioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
