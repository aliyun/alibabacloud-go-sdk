// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDictResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDictResponseBody) *DeleteDictResponse
	GetBody() *DeleteDictResponseBody
}

type DeleteDictResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDictResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDictResponse) GoString() string {
	return s.String()
}

func (s *DeleteDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDictResponse) GetBody() *DeleteDictResponseBody {
	return s.Body
}

func (s *DeleteDictResponse) SetHeaders(v map[string]*string) *DeleteDictResponse {
	s.Headers = v
	return s
}

func (s *DeleteDictResponse) SetStatusCode(v int32) *DeleteDictResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDictResponse) SetBody(v *DeleteDictResponseBody) *DeleteDictResponse {
	s.Body = v
	return s
}

func (s *DeleteDictResponse) Validate() error {
	return dara.Validate(s)
}
