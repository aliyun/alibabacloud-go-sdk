// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMemoryNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMemoryNodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMemoryNodeResponseBody) *DeleteMemoryNodeResponse
	GetBody() *DeleteMemoryNodeResponseBody
}

type DeleteMemoryNodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMemoryNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMemoryNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteMemoryNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMemoryNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMemoryNodeResponse) GetBody() *DeleteMemoryNodeResponseBody {
	return s.Body
}

func (s *DeleteMemoryNodeResponse) SetHeaders(v map[string]*string) *DeleteMemoryNodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteMemoryNodeResponse) SetStatusCode(v int32) *DeleteMemoryNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMemoryNodeResponse) SetBody(v *DeleteMemoryNodeResponseBody) *DeleteMemoryNodeResponse {
	s.Body = v
	return s
}

func (s *DeleteMemoryNodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
