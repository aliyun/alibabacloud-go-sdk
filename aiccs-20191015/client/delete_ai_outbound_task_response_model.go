// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAiOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAiOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAiOutboundTaskResponseBody) *DeleteAiOutboundTaskResponse
	GetBody() *DeleteAiOutboundTaskResponseBody
}

type DeleteAiOutboundTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAiOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAiOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteAiOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAiOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAiOutboundTaskResponse) GetBody() *DeleteAiOutboundTaskResponseBody {
	return s.Body
}

func (s *DeleteAiOutboundTaskResponse) SetHeaders(v map[string]*string) *DeleteAiOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteAiOutboundTaskResponse) SetStatusCode(v int32) *DeleteAiOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAiOutboundTaskResponse) SetBody(v *DeleteAiOutboundTaskResponseBody) *DeleteAiOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *DeleteAiOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
