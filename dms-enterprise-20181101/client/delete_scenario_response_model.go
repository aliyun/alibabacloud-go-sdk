// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScenarioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScenarioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScenarioResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScenarioResponseBody) *DeleteScenarioResponse
	GetBody() *DeleteScenarioResponseBody
}

type DeleteScenarioResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScenarioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScenarioResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScenarioResponse) GoString() string {
	return s.String()
}

func (s *DeleteScenarioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScenarioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScenarioResponse) GetBody() *DeleteScenarioResponseBody {
	return s.Body
}

func (s *DeleteScenarioResponse) SetHeaders(v map[string]*string) *DeleteScenarioResponse {
	s.Headers = v
	return s
}

func (s *DeleteScenarioResponse) SetStatusCode(v int32) *DeleteScenarioResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScenarioResponse) SetBody(v *DeleteScenarioResponseBody) *DeleteScenarioResponse {
	s.Body = v
	return s
}

func (s *DeleteScenarioResponse) Validate() error {
	return dara.Validate(s)
}
