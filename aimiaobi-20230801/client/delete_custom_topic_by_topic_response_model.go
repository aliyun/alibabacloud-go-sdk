// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTopicByTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomTopicByTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomTopicByTopicResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomTopicByTopicResponseBody) *DeleteCustomTopicByTopicResponse
	GetBody() *DeleteCustomTopicByTopicResponseBody
}

type DeleteCustomTopicByTopicResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomTopicByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomTopicByTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTopicByTopicResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomTopicByTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomTopicByTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomTopicByTopicResponse) GetBody() *DeleteCustomTopicByTopicResponseBody {
	return s.Body
}

func (s *DeleteCustomTopicByTopicResponse) SetHeaders(v map[string]*string) *DeleteCustomTopicByTopicResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomTopicByTopicResponse) SetStatusCode(v int32) *DeleteCustomTopicByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomTopicByTopicResponse) SetBody(v *DeleteCustomTopicByTopicResponseBody) *DeleteCustomTopicByTopicResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomTopicByTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
