// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPeekMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PeekMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PeekMessageResponse
	GetStatusCode() *int32
	SetBody(v *PeekMessageResponseBody) *PeekMessageResponse
	GetBody() *PeekMessageResponseBody
}

type PeekMessageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PeekMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PeekMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s PeekMessageResponse) GoString() string {
	return s.String()
}

func (s *PeekMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PeekMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PeekMessageResponse) GetBody() *PeekMessageResponseBody {
	return s.Body
}

func (s *PeekMessageResponse) SetHeaders(v map[string]*string) *PeekMessageResponse {
	s.Headers = v
	return s
}

func (s *PeekMessageResponse) SetStatusCode(v int32) *PeekMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *PeekMessageResponse) SetBody(v *PeekMessageResponseBody) *PeekMessageResponse {
	s.Body = v
	return s
}

func (s *PeekMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
