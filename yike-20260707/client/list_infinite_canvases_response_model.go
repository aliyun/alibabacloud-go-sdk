// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfiniteCanvasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInfiniteCanvasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInfiniteCanvasesResponse
	GetStatusCode() *int32
	SetBody(v *ListInfiniteCanvasesResponseBody) *ListInfiniteCanvasesResponse
	GetBody() *ListInfiniteCanvasesResponseBody
}

type ListInfiniteCanvasesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInfiniteCanvasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInfiniteCanvasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInfiniteCanvasesResponse) GoString() string {
	return s.String()
}

func (s *ListInfiniteCanvasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInfiniteCanvasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInfiniteCanvasesResponse) GetBody() *ListInfiniteCanvasesResponseBody {
	return s.Body
}

func (s *ListInfiniteCanvasesResponse) SetHeaders(v map[string]*string) *ListInfiniteCanvasesResponse {
	s.Headers = v
	return s
}

func (s *ListInfiniteCanvasesResponse) SetStatusCode(v int32) *ListInfiniteCanvasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInfiniteCanvasesResponse) SetBody(v *ListInfiniteCanvasesResponseBody) *ListInfiniteCanvasesResponse {
	s.Body = v
	return s
}

func (s *ListInfiniteCanvasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
