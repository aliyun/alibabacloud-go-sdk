// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoginBaseConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLoginBaseConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLoginBaseConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLoginBaseConfigResponseBody) *DeleteLoginBaseConfigResponse
	GetBody() *DeleteLoginBaseConfigResponseBody
}

type DeleteLoginBaseConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLoginBaseConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLoginBaseConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoginBaseConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoginBaseConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLoginBaseConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLoginBaseConfigResponse) GetBody() *DeleteLoginBaseConfigResponseBody {
	return s.Body
}

func (s *DeleteLoginBaseConfigResponse) SetHeaders(v map[string]*string) *DeleteLoginBaseConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoginBaseConfigResponse) SetStatusCode(v int32) *DeleteLoginBaseConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoginBaseConfigResponse) SetBody(v *DeleteLoginBaseConfigResponseBody) *DeleteLoginBaseConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLoginBaseConfigResponse) Validate() error {
	return dara.Validate(s)
}
