// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTopicResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTopicResponseBody) *UpdateTopicResponse
	GetBody() *UpdateTopicResponseBody
}

type UpdateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTopicResponse) GoString() string {
	return s.String()
}

func (s *UpdateTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTopicResponse) GetBody() *UpdateTopicResponseBody {
	return s.Body
}

func (s *UpdateTopicResponse) SetHeaders(v map[string]*string) *UpdateTopicResponse {
	s.Headers = v
	return s
}

func (s *UpdateTopicResponse) SetStatusCode(v int32) *UpdateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTopicResponse) SetBody(v *UpdateTopicResponseBody) *UpdateTopicResponse {
	s.Body = v
	return s
}

func (s *UpdateTopicResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
