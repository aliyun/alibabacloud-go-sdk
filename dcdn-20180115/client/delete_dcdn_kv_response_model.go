// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDcdnKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDcdnKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDcdnKvResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDcdnKvResponseBody) *DeleteDcdnKvResponse
	GetBody() *DeleteDcdnKvResponseBody
}

type DeleteDcdnKvResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDcdnKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDcdnKvResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDcdnKvResponse) GoString() string {
	return s.String()
}

func (s *DeleteDcdnKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDcdnKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDcdnKvResponse) GetBody() *DeleteDcdnKvResponseBody {
	return s.Body
}

func (s *DeleteDcdnKvResponse) SetHeaders(v map[string]*string) *DeleteDcdnKvResponse {
	s.Headers = v
	return s
}

func (s *DeleteDcdnKvResponse) SetStatusCode(v int32) *DeleteDcdnKvResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDcdnKvResponse) SetBody(v *DeleteDcdnKvResponseBody) *DeleteDcdnKvResponse {
	s.Body = v
	return s
}

func (s *DeleteDcdnKvResponse) Validate() error {
	return dara.Validate(s)
}
