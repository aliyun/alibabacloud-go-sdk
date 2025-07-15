// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateApplicationResponseBody) *UpdateApplicationResponse
	GetBody() *UpdateApplicationResponseBody
}

type UpdateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateApplicationResponse) GetBody() *UpdateApplicationResponseBody {
	return s.Body
}

func (s *UpdateApplicationResponse) SetHeaders(v map[string]*string) *UpdateApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationResponse) SetStatusCode(v int32) *UpdateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationResponse) SetBody(v *UpdateApplicationResponseBody) *UpdateApplicationResponse {
	s.Body = v
	return s
}

func (s *UpdateApplicationResponse) Validate() error {
	return dara.Validate(s)
}
