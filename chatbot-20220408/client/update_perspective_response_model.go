// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePerspectiveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePerspectiveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePerspectiveResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePerspectiveResponseBody) *UpdatePerspectiveResponse
	GetBody() *UpdatePerspectiveResponseBody
}

type UpdatePerspectiveResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePerspectiveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePerspectiveResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePerspectiveResponse) GoString() string {
	return s.String()
}

func (s *UpdatePerspectiveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePerspectiveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePerspectiveResponse) GetBody() *UpdatePerspectiveResponseBody {
	return s.Body
}

func (s *UpdatePerspectiveResponse) SetHeaders(v map[string]*string) *UpdatePerspectiveResponse {
	s.Headers = v
	return s
}

func (s *UpdatePerspectiveResponse) SetStatusCode(v int32) *UpdatePerspectiveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePerspectiveResponse) SetBody(v *UpdatePerspectiveResponseBody) *UpdatePerspectiveResponse {
	s.Body = v
	return s
}

func (s *UpdatePerspectiveResponse) Validate() error {
	return dara.Validate(s)
}
