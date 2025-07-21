// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConsumerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConsumerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConsumerResponseBody) *UpdateConsumerResponse
	GetBody() *UpdateConsumerResponseBody
}

type UpdateConsumerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConsumerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConsumerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConsumerResponse) GetBody() *UpdateConsumerResponseBody {
	return s.Body
}

func (s *UpdateConsumerResponse) SetHeaders(v map[string]*string) *UpdateConsumerResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerResponse) SetStatusCode(v int32) *UpdateConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerResponse) SetBody(v *UpdateConsumerResponseBody) *UpdateConsumerResponse {
	s.Body = v
	return s
}

func (s *UpdateConsumerResponse) Validate() error {
	return dara.Validate(s)
}
