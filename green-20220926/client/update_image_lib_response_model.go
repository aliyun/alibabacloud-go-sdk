// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateImageLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateImageLibResponse
	GetStatusCode() *int32
	SetBody(v *UpdateImageLibResponseBody) *UpdateImageLibResponse
	GetBody() *UpdateImageLibResponseBody
}

type UpdateImageLibResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateImageLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateImageLibResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageLibResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateImageLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateImageLibResponse) GetBody() *UpdateImageLibResponseBody {
	return s.Body
}

func (s *UpdateImageLibResponse) SetHeaders(v map[string]*string) *UpdateImageLibResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageLibResponse) SetStatusCode(v int32) *UpdateImageLibResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageLibResponse) SetBody(v *UpdateImageLibResponseBody) *UpdateImageLibResponse {
	s.Body = v
	return s
}

func (s *UpdateImageLibResponse) Validate() error {
	return dara.Validate(s)
}
