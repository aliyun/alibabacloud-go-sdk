// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInfiniteCanvasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInfiniteCanvasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInfiniteCanvasResponse
	GetStatusCode() *int32
	SetBody(v *CreateInfiniteCanvasResponseBody) *CreateInfiniteCanvasResponse
	GetBody() *CreateInfiniteCanvasResponseBody
}

type CreateInfiniteCanvasResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInfiniteCanvasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInfiniteCanvasResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInfiniteCanvasResponse) GoString() string {
	return s.String()
}

func (s *CreateInfiniteCanvasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInfiniteCanvasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInfiniteCanvasResponse) GetBody() *CreateInfiniteCanvasResponseBody {
	return s.Body
}

func (s *CreateInfiniteCanvasResponse) SetHeaders(v map[string]*string) *CreateInfiniteCanvasResponse {
	s.Headers = v
	return s
}

func (s *CreateInfiniteCanvasResponse) SetStatusCode(v int32) *CreateInfiniteCanvasResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInfiniteCanvasResponse) SetBody(v *CreateInfiniteCanvasResponseBody) *CreateInfiniteCanvasResponse {
	s.Body = v
	return s
}

func (s *CreateInfiniteCanvasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
