// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDocResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDocResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDocResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDocResponseBody) *UpdateDocResponse
	GetBody() *UpdateDocResponseBody
}

type UpdateDocResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDocResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDocResponse) GoString() string {
	return s.String()
}

func (s *UpdateDocResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDocResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDocResponse) GetBody() *UpdateDocResponseBody {
	return s.Body
}

func (s *UpdateDocResponse) SetHeaders(v map[string]*string) *UpdateDocResponse {
	s.Headers = v
	return s
}

func (s *UpdateDocResponse) SetStatusCode(v int32) *UpdateDocResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDocResponse) SetBody(v *UpdateDocResponseBody) *UpdateDocResponse {
	s.Body = v
	return s
}

func (s *UpdateDocResponse) Validate() error {
	return dara.Validate(s)
}
