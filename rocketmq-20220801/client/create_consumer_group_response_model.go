// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConsumerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConsumerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConsumerGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateConsumerGroupResponseBody) *CreateConsumerGroupResponse
	GetBody() *CreateConsumerGroupResponseBody
}

type CreateConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConsumerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConsumerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConsumerGroupResponse) GetBody() *CreateConsumerGroupResponseBody {
	return s.Body
}

func (s *CreateConsumerGroupResponse) SetHeaders(v map[string]*string) *CreateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerGroupResponse) SetStatusCode(v int32) *CreateConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConsumerGroupResponse) SetBody(v *CreateConsumerGroupResponseBody) *CreateConsumerGroupResponse {
	s.Body = v
	return s
}

func (s *CreateConsumerGroupResponse) Validate() error {
	return dara.Validate(s)
}
