// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRunResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRunResponseBody) *UpdateRunResponse
	GetBody() *UpdateRunResponseBody
}

type UpdateRunResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRunResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRunResponse) GoString() string {
	return s.String()
}

func (s *UpdateRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRunResponse) GetBody() *UpdateRunResponseBody {
	return s.Body
}

func (s *UpdateRunResponse) SetHeaders(v map[string]*string) *UpdateRunResponse {
	s.Headers = v
	return s
}

func (s *UpdateRunResponse) SetStatusCode(v int32) *UpdateRunResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRunResponse) SetBody(v *UpdateRunResponseBody) *UpdateRunResponse {
	s.Body = v
	return s
}

func (s *UpdateRunResponse) Validate() error {
	return dara.Validate(s)
}
