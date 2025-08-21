// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliwsDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAliwsDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAliwsDictResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAliwsDictResponseBody) *UpdateAliwsDictResponse
	GetBody() *UpdateAliwsDictResponseBody
}

type UpdateAliwsDictResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAliwsDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliwsDictResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliwsDictResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliwsDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAliwsDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAliwsDictResponse) GetBody() *UpdateAliwsDictResponseBody {
	return s.Body
}

func (s *UpdateAliwsDictResponse) SetHeaders(v map[string]*string) *UpdateAliwsDictResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliwsDictResponse) SetStatusCode(v int32) *UpdateAliwsDictResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAliwsDictResponse) SetBody(v *UpdateAliwsDictResponseBody) *UpdateAliwsDictResponse {
	s.Body = v
	return s
}

func (s *UpdateAliwsDictResponse) Validate() error {
	return dara.Validate(s)
}
