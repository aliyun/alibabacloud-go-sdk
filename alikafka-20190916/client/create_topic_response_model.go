// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTopicResponse
	GetStatusCode() *int32
	SetBody(v *CreateTopicResponseBody) *CreateTopicResponse
	GetBody() *CreateTopicResponseBody
}

type CreateTopicResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponse) GoString() string {
	return s.String()
}

func (s *CreateTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTopicResponse) GetBody() *CreateTopicResponseBody {
	return s.Body
}

func (s *CreateTopicResponse) SetHeaders(v map[string]*string) *CreateTopicResponse {
	s.Headers = v
	return s
}

func (s *CreateTopicResponse) SetStatusCode(v int32) *CreateTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTopicResponse) SetBody(v *CreateTopicResponseBody) *CreateTopicResponse {
	s.Body = v
	return s
}

func (s *CreateTopicResponse) Validate() error {
	return dara.Validate(s)
}
