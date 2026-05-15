// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAiOutboundTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAiOutboundTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAiOutboundTaskResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAiOutboundTaskResponseBody) *UpdateAiOutboundTaskResponse
	GetBody() *UpdateAiOutboundTaskResponseBody
}

type UpdateAiOutboundTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAiOutboundTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAiOutboundTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAiOutboundTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateAiOutboundTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAiOutboundTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAiOutboundTaskResponse) GetBody() *UpdateAiOutboundTaskResponseBody {
	return s.Body
}

func (s *UpdateAiOutboundTaskResponse) SetHeaders(v map[string]*string) *UpdateAiOutboundTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateAiOutboundTaskResponse) SetStatusCode(v int32) *UpdateAiOutboundTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAiOutboundTaskResponse) SetBody(v *UpdateAiOutboundTaskResponseBody) *UpdateAiOutboundTaskResponse {
	s.Body = v
	return s
}

func (s *UpdateAiOutboundTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
