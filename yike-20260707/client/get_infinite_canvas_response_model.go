// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInfiniteCanvasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInfiniteCanvasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInfiniteCanvasResponse
	GetStatusCode() *int32
	SetBody(v *GetInfiniteCanvasResponseBody) *GetInfiniteCanvasResponse
	GetBody() *GetInfiniteCanvasResponseBody
}

type GetInfiniteCanvasResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInfiniteCanvasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInfiniteCanvasResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInfiniteCanvasResponse) GoString() string {
	return s.String()
}

func (s *GetInfiniteCanvasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInfiniteCanvasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInfiniteCanvasResponse) GetBody() *GetInfiniteCanvasResponseBody {
	return s.Body
}

func (s *GetInfiniteCanvasResponse) SetHeaders(v map[string]*string) *GetInfiniteCanvasResponse {
	s.Headers = v
	return s
}

func (s *GetInfiniteCanvasResponse) SetStatusCode(v int32) *GetInfiniteCanvasResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInfiniteCanvasResponse) SetBody(v *GetInfiniteCanvasResponseBody) *GetInfiniteCanvasResponse {
	s.Body = v
	return s
}

func (s *GetInfiniteCanvasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
