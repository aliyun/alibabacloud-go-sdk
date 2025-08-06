// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaStorageClassResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaStorageClassResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaStorageClassResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaStorageClassResponseBody) *UpdateMediaStorageClassResponse
	GetBody() *UpdateMediaStorageClassResponseBody
}

type UpdateMediaStorageClassResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaStorageClassResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaStorageClassResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaStorageClassResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaStorageClassResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaStorageClassResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaStorageClassResponse) GetBody() *UpdateMediaStorageClassResponseBody {
	return s.Body
}

func (s *UpdateMediaStorageClassResponse) SetHeaders(v map[string]*string) *UpdateMediaStorageClassResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaStorageClassResponse) SetStatusCode(v int32) *UpdateMediaStorageClassResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaStorageClassResponse) SetBody(v *UpdateMediaStorageClassResponseBody) *UpdateMediaStorageClassResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaStorageClassResponse) Validate() error {
	return dara.Validate(s)
}
