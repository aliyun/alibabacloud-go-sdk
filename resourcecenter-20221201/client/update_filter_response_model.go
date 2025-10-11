// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFilterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFilterResponseBody) *UpdateFilterResponse
	GetBody() *UpdateFilterResponseBody
}

type UpdateFilterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFilterResponse) GoString() string {
	return s.String()
}

func (s *UpdateFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFilterResponse) GetBody() *UpdateFilterResponseBody {
	return s.Body
}

func (s *UpdateFilterResponse) SetHeaders(v map[string]*string) *UpdateFilterResponse {
	s.Headers = v
	return s
}

func (s *UpdateFilterResponse) SetStatusCode(v int32) *UpdateFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFilterResponse) SetBody(v *UpdateFilterResponseBody) *UpdateFilterResponse {
	s.Body = v
	return s
}

func (s *UpdateFilterResponse) Validate() error {
	return dara.Validate(s)
}
