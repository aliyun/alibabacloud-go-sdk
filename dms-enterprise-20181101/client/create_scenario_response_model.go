// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScenarioResponse
	GetStatusCode() *int32
	SetBody(v *CreateScenarioResponseBody) *CreateScenarioResponse
	GetBody() *CreateScenarioResponseBody
}

type CreateScenarioResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScenarioResponse) GoString() string {
	return s.String()
}

func (s *CreateScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScenarioResponse) GetBody() *CreateScenarioResponseBody {
	return s.Body
}

func (s *CreateScenarioResponse) SetHeaders(v map[string]*string) *CreateScenarioResponse {
	s.Headers = v
	return s
}

func (s *CreateScenarioResponse) SetStatusCode(v int32) *CreateScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScenarioResponse) SetBody(v *CreateScenarioResponseBody) *CreateScenarioResponse {
	s.Body = v
	return s
}

func (s *CreateScenarioResponse) Validate() error {
	return dara.Validate(s)
}
