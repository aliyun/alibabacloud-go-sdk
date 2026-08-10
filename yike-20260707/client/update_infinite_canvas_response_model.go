// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInfiniteCanvasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInfiniteCanvasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInfiniteCanvasResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInfiniteCanvasResponseBody) *UpdateInfiniteCanvasResponse
	GetBody() *UpdateInfiniteCanvasResponseBody
}

type UpdateInfiniteCanvasResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInfiniteCanvasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInfiniteCanvasResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInfiniteCanvasResponse) GoString() string {
	return s.String()
}

func (s *UpdateInfiniteCanvasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInfiniteCanvasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInfiniteCanvasResponse) GetBody() *UpdateInfiniteCanvasResponseBody {
	return s.Body
}

func (s *UpdateInfiniteCanvasResponse) SetHeaders(v map[string]*string) *UpdateInfiniteCanvasResponse {
	s.Headers = v
	return s
}

func (s *UpdateInfiniteCanvasResponse) SetStatusCode(v int32) *UpdateInfiniteCanvasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInfiniteCanvasResponse) SetBody(v *UpdateInfiniteCanvasResponseBody) *UpdateInfiniteCanvasResponse {
	s.Body = v
	return s
}

func (s *UpdateInfiniteCanvasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
