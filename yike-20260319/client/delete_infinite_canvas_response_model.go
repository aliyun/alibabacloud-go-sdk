// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInfiniteCanvasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInfiniteCanvasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInfiniteCanvasResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInfiniteCanvasResponseBody) *DeleteInfiniteCanvasResponse
	GetBody() *DeleteInfiniteCanvasResponseBody
}

type DeleteInfiniteCanvasResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInfiniteCanvasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInfiniteCanvasResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInfiniteCanvasResponse) GoString() string {
	return s.String()
}

func (s *DeleteInfiniteCanvasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInfiniteCanvasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInfiniteCanvasResponse) GetBody() *DeleteInfiniteCanvasResponseBody {
	return s.Body
}

func (s *DeleteInfiniteCanvasResponse) SetHeaders(v map[string]*string) *DeleteInfiniteCanvasResponse {
	s.Headers = v
	return s
}

func (s *DeleteInfiniteCanvasResponse) SetStatusCode(v int32) *DeleteInfiniteCanvasResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInfiniteCanvasResponse) SetBody(v *DeleteInfiniteCanvasResponseBody) *DeleteInfiniteCanvasResponse {
	s.Body = v
	return s
}

func (s *DeleteInfiniteCanvasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
