// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemoryResponseBody) *DeleteMemoryResponse
	GetBody() *DeleteMemoryResponseBody
}

type DeleteMemoryResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemoryResponse) GetBody() *DeleteMemoryResponseBody {
	return s.Body
}

func (s *DeleteMemoryResponse) SetHeaders(v map[string]*string) *DeleteMemoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryResponse) SetStatusCode(v int32) *DeleteMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryResponse) SetBody(v *DeleteMemoryResponseBody) *DeleteMemoryResponse {
	s.Body = v
	return s
}

func (s *DeleteMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
