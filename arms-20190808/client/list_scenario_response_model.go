// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScenarioResponse
	GetStatusCode() *int32
	SetBody(v *ListScenarioResponseBody) *ListScenarioResponse
	GetBody() *ListScenarioResponseBody
}

type ListScenarioResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScenarioResponse) GoString() string {
	return s.String()
}

func (s *ListScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScenarioResponse) GetBody() *ListScenarioResponseBody {
	return s.Body
}

func (s *ListScenarioResponse) SetHeaders(v map[string]*string) *ListScenarioResponse {
	s.Headers = v
	return s
}

func (s *ListScenarioResponse) SetStatusCode(v int32) *ListScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScenarioResponse) SetBody(v *ListScenarioResponseBody) *ListScenarioResponse {
	s.Body = v
	return s
}

func (s *ListScenarioResponse) Validate() error {
	return dara.Validate(s)
}
