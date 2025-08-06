// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEditingProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEditingProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEditingProjectResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEditingProjectResponseBody) *DeleteEditingProjectResponse
	GetBody() *DeleteEditingProjectResponseBody
}

type DeleteEditingProjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEditingProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEditingProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEditingProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteEditingProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEditingProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEditingProjectResponse) GetBody() *DeleteEditingProjectResponseBody {
	return s.Body
}

func (s *DeleteEditingProjectResponse) SetHeaders(v map[string]*string) *DeleteEditingProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteEditingProjectResponse) SetStatusCode(v int32) *DeleteEditingProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEditingProjectResponse) SetBody(v *DeleteEditingProjectResponseBody) *DeleteEditingProjectResponse {
	s.Body = v
	return s
}

func (s *DeleteEditingProjectResponse) Validate() error {
	return dara.Validate(s)
}
