// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceMirrorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceMirrorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceMirrorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceMirrorResponseBody) *DeleteServiceMirrorResponse
	GetBody() *DeleteServiceMirrorResponseBody
}

type DeleteServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceMirrorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceMirrorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceMirrorResponse) GetBody() *DeleteServiceMirrorResponseBody {
	return s.Body
}

func (s *DeleteServiceMirrorResponse) SetHeaders(v map[string]*string) *DeleteServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceMirrorResponse) SetStatusCode(v int32) *DeleteServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceMirrorResponse) SetBody(v *DeleteServiceMirrorResponseBody) *DeleteServiceMirrorResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceMirrorResponse) Validate() error {
	return dara.Validate(s)
}
