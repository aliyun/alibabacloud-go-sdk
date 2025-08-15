// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConsumerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConsumerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConsumerGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConsumerGroupResponseBody) *UpdateConsumerGroupResponse
	GetBody() *UpdateConsumerGroupResponseBody
}

type UpdateConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConsumerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateConsumerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConsumerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConsumerGroupResponse) GetBody() *UpdateConsumerGroupResponseBody {
	return s.Body
}

func (s *UpdateConsumerGroupResponse) SetHeaders(v map[string]*string) *UpdateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateConsumerGroupResponse) SetStatusCode(v int32) *UpdateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConsumerGroupResponse) SetBody(v *UpdateConsumerGroupResponseBody) *UpdateConsumerGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateConsumerGroupResponse) Validate() error {
	return dara.Validate(s)
}
