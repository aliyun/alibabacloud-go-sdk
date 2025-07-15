// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerOffsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConsumerOffsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConsumerOffsetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConsumerOffsetResponseBody) *UpdateConsumerOffsetResponse
	GetBody() *UpdateConsumerOffsetResponseBody
}

type UpdateConsumerOffsetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConsumerOffsetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConsumerOffsetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerOffsetResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerOffsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConsumerOffsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConsumerOffsetResponse) GetBody() *UpdateConsumerOffsetResponseBody {
	return s.Body
}

func (s *UpdateConsumerOffsetResponse) SetHeaders(v map[string]*string) *UpdateConsumerOffsetResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerOffsetResponse) SetStatusCode(v int32) *UpdateConsumerOffsetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerOffsetResponse) SetBody(v *UpdateConsumerOffsetResponseBody) *UpdateConsumerOffsetResponse {
	s.Body = v
	return s
}

func (s *UpdateConsumerOffsetResponse) Validate() error {
	return dara.Validate(s)
}
