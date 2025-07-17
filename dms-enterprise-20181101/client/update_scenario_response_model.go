// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScenarioResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScenarioResponseBody) *UpdateScenarioResponse
	GetBody() *UpdateScenarioResponseBody
}

type UpdateScenarioResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScenarioResponse) GoString() string {
	return s.String()
}

func (s *UpdateScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScenarioResponse) GetBody() *UpdateScenarioResponseBody {
	return s.Body
}

func (s *UpdateScenarioResponse) SetHeaders(v map[string]*string) *UpdateScenarioResponse {
	s.Headers = v
	return s
}

func (s *UpdateScenarioResponse) SetStatusCode(v int32) *UpdateScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScenarioResponse) SetBody(v *UpdateScenarioResponseBody) *UpdateScenarioResponse {
	s.Body = v
	return s
}

func (s *UpdateScenarioResponse) Validate() error {
	return dara.Validate(s)
}
