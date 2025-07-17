// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentAddonsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEnvironmentAddonsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEnvironmentAddonsResponse
	GetStatusCode() *int32
	SetBody(v *ListEnvironmentAddonsResponseBody) *ListEnvironmentAddonsResponse
	GetBody() *ListEnvironmentAddonsResponseBody
}

type ListEnvironmentAddonsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEnvironmentAddonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEnvironmentAddonsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsResponse) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEnvironmentAddonsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEnvironmentAddonsResponse) GetBody() *ListEnvironmentAddonsResponseBody {
	return s.Body
}

func (s *ListEnvironmentAddonsResponse) SetHeaders(v map[string]*string) *ListEnvironmentAddonsResponse {
	s.Headers = v
	return s
}

func (s *ListEnvironmentAddonsResponse) SetStatusCode(v int32) *ListEnvironmentAddonsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEnvironmentAddonsResponse) SetBody(v *ListEnvironmentAddonsResponseBody) *ListEnvironmentAddonsResponse {
	s.Body = v
	return s
}

func (s *ListEnvironmentAddonsResponse) Validate() error {
	return dara.Validate(s)
}
