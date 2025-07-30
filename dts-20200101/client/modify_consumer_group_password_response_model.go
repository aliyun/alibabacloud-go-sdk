// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConsumerGroupPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyConsumerGroupPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyConsumerGroupPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ModifyConsumerGroupPasswordResponseBody) *ModifyConsumerGroupPasswordResponse
	GetBody() *ModifyConsumerGroupPasswordResponseBody
}

type ModifyConsumerGroupPasswordResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyConsumerGroupPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyConsumerGroupPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyConsumerGroupPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyConsumerGroupPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyConsumerGroupPasswordResponse) GetBody() *ModifyConsumerGroupPasswordResponseBody {
	return s.Body
}

func (s *ModifyConsumerGroupPasswordResponse) SetHeaders(v map[string]*string) *ModifyConsumerGroupPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumerGroupPasswordResponse) SetStatusCode(v int32) *ModifyConsumerGroupPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponse) SetBody(v *ModifyConsumerGroupPasswordResponseBody) *ModifyConsumerGroupPasswordResponse {
	s.Body = v
	return s
}

func (s *ModifyConsumerGroupPasswordResponse) Validate() error {
	return dara.Validate(s)
}
