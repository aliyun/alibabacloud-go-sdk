// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAScriptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAScriptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAScriptsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAScriptsResponseBody) *UpdateAScriptsResponse
	GetBody() *UpdateAScriptsResponseBody
}

type UpdateAScriptsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAScriptsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAScriptsResponse) GoString() string {
	return s.String()
}

func (s *UpdateAScriptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAScriptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAScriptsResponse) GetBody() *UpdateAScriptsResponseBody {
	return s.Body
}

func (s *UpdateAScriptsResponse) SetHeaders(v map[string]*string) *UpdateAScriptsResponse {
	s.Headers = v
	return s
}

func (s *UpdateAScriptsResponse) SetStatusCode(v int32) *UpdateAScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAScriptsResponse) SetBody(v *UpdateAScriptsResponseBody) *UpdateAScriptsResponse {
	s.Body = v
	return s
}

func (s *UpdateAScriptsResponse) Validate() error {
	return dara.Validate(s)
}
