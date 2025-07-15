// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTopicViewPointByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomTopicViewPointByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomTopicViewPointByIdResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomTopicViewPointByIdResponseBody) *DeleteCustomTopicViewPointByIdResponse
	GetBody() *DeleteCustomTopicViewPointByIdResponseBody
}

type DeleteCustomTopicViewPointByIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTopicViewPointByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTopicViewPointByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTopicViewPointByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicViewPointByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomTopicViewPointByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomTopicViewPointByIdResponse) GetBody() *DeleteCustomTopicViewPointByIdResponseBody {
	return s.Body
}

func (s *DeleteCustomTopicViewPointByIdResponse) SetHeaders(v map[string]*string) *DeleteCustomTopicViewPointByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponse) SetStatusCode(v int32) *DeleteCustomTopicViewPointByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponse) SetBody(v *DeleteCustomTopicViewPointByIdResponseBody) *DeleteCustomTopicViewPointByIdResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomTopicViewPointByIdResponse) Validate() error {
	return dara.Validate(s)
}
