// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOutboundTaskResponseBody) *DeleteOutboundTaskResponse
	GetBody() *DeleteOutboundTaskResponseBody
}

type DeleteOutboundTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOutboundTaskResponse) GetBody() *DeleteOutboundTaskResponseBody {
	return s.Body
}

func (s *DeleteOutboundTaskResponse) SetHeaders(v map[string]*string) *DeleteOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteOutboundTaskResponse) SetStatusCode(v int32) *DeleteOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOutboundTaskResponse) SetBody(v *DeleteOutboundTaskResponseBody) *DeleteOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
