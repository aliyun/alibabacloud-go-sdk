// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConsumerGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConsumerGroupResponseBody) *DeleteConsumerGroupResponse
	GetBody() *DeleteConsumerGroupResponseBody
}

type DeleteConsumerGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConsumerGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConsumerGroupResponse) GetBody() *DeleteConsumerGroupResponseBody {
	return s.Body
}

func (s *DeleteConsumerGroupResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupResponse) SetStatusCode(v int32) *DeleteConsumerGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerGroupResponse) SetBody(v *DeleteConsumerGroupResponseBody) *DeleteConsumerGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteConsumerGroupResponse) Validate() error {
	return dara.Validate(s)
}
