// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContactResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContactResponseBody) *DeleteContactResponse
	GetBody() *DeleteContactResponseBody
}

type DeleteContactResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContactResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactResponse) GoString() string {
	return s.String()
}

func (s *DeleteContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContactResponse) GetBody() *DeleteContactResponseBody {
	return s.Body
}

func (s *DeleteContactResponse) SetHeaders(v map[string]*string) *DeleteContactResponse {
	s.Headers = v
	return s
}

func (s *DeleteContactResponse) SetStatusCode(v int32) *DeleteContactResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContactResponse) SetBody(v *DeleteContactResponseBody) *DeleteContactResponse {
	s.Body = v
	return s
}

func (s *DeleteContactResponse) Validate() error {
	return dara.Validate(s)
}
