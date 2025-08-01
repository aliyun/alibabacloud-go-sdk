// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageQueueRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMessageQueueRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMessageQueueRouteResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMessageQueueRouteResponseBody) *UpdateMessageQueueRouteResponse
	GetBody() *UpdateMessageQueueRouteResponseBody
}

type UpdateMessageQueueRouteResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageQueueRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageQueueRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageQueueRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageQueueRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMessageQueueRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageQueueRouteResponse) GetBody() *UpdateMessageQueueRouteResponseBody {
	return s.Body
}

func (s *UpdateMessageQueueRouteResponse) SetHeaders(v map[string]*string) *UpdateMessageQueueRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageQueueRouteResponse) SetStatusCode(v int32) *UpdateMessageQueueRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageQueueRouteResponse) SetBody(v *UpdateMessageQueueRouteResponseBody) *UpdateMessageQueueRouteResponse {
	s.Body = v
	return s
}

func (s *UpdateMessageQueueRouteResponse) Validate() error {
	return dara.Validate(s)
}
