// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTopicResponse
	GetStatusCode() *int32
	SetBody(v *ListTopicResponseBody) *ListTopicResponse
	GetBody() *ListTopicResponseBody
}

type ListTopicResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTopicResponse) GoString() string {
	return s.String()
}

func (s *ListTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTopicResponse) GetBody() *ListTopicResponseBody {
	return s.Body
}

func (s *ListTopicResponse) SetHeaders(v map[string]*string) *ListTopicResponse {
	s.Headers = v
	return s
}

func (s *ListTopicResponse) SetStatusCode(v int32) *ListTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicResponse) SetBody(v *ListTopicResponseBody) *ListTopicResponse {
	s.Body = v
	return s
}

func (s *ListTopicResponse) Validate() error {
	return dara.Validate(s)
}
