// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteQueueResponse
	GetStatusCode() *int32
	SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse
	GetBody() *DeleteQueueResponseBody
}

type DeleteQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteQueueResponse) GetBody() *DeleteQueueResponseBody {
	return s.Body
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetStatusCode(v int32) *DeleteQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

func (s *DeleteQueueResponse) Validate() error {
	return dara.Validate(s)
}
