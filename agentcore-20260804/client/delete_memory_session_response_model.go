// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemorySessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemorySessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemorySessionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemorySessionResponseBody) *DeleteMemorySessionResponse
	GetBody() *DeleteMemorySessionResponseBody
}

type DeleteMemorySessionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemorySessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemorySessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemorySessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemorySessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemorySessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemorySessionResponse) GetBody() *DeleteMemorySessionResponseBody {
	return s.Body
}

func (s *DeleteMemorySessionResponse) SetHeaders(v map[string]*string) *DeleteMemorySessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemorySessionResponse) SetStatusCode(v int32) *DeleteMemorySessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemorySessionResponse) SetBody(v *DeleteMemorySessionResponseBody) *DeleteMemorySessionResponse {
	s.Body = v
	return s
}

func (s *DeleteMemorySessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
