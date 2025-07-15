// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTopicResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse
	GetBody() *DeleteTopicResponseBody
}

type DeleteTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTopicResponse) GetBody() *DeleteTopicResponseBody {
	return s.Body
}

func (s *DeleteTopicResponse) SetHeaders(v map[string]*string) *DeleteTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteTopicResponse) SetStatusCode(v int32) *DeleteTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTopicResponse) SetBody(v *DeleteTopicResponseBody) *DeleteTopicResponse {
	s.Body = v
	return s
}

func (s *DeleteTopicResponse) Validate() error {
	return dara.Validate(s)
}
