// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScriptsResponse
	GetStatusCode() *int32
	SetBody(v *ListScriptsResponseBody) *ListScriptsResponse
	GetBody() *ListScriptsResponseBody
}

type ListScriptsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScriptsResponse) GoString() string {
	return s.String()
}

func (s *ListScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScriptsResponse) GetBody() *ListScriptsResponseBody {
	return s.Body
}

func (s *ListScriptsResponse) SetHeaders(v map[string]*string) *ListScriptsResponse {
	s.Headers = v
	return s
}

func (s *ListScriptsResponse) SetStatusCode(v int32) *ListScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScriptsResponse) SetBody(v *ListScriptsResponseBody) *ListScriptsResponse {
	s.Body = v
	return s
}

func (s *ListScriptsResponse) Validate() error {
	return dara.Validate(s)
}
