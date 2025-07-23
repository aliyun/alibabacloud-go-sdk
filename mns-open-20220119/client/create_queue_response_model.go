// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateQueueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateQueueResponse
	GetStatusCode() *int32
	SetBody(v *CreateQueueResponseBody) *CreateQueueResponse
	GetBody() *CreateQueueResponseBody
}

type CreateQueueResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateQueueResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueResponse) GoString() string {
	return s.String()
}

func (s *CreateQueueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateQueueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateQueueResponse) GetBody() *CreateQueueResponseBody {
	return s.Body
}

func (s *CreateQueueResponse) SetHeaders(v map[string]*string) *CreateQueueResponse {
	s.Headers = v
	return s
}

func (s *CreateQueueResponse) SetStatusCode(v int32) *CreateQueueResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueueResponse) SetBody(v *CreateQueueResponseBody) *CreateQueueResponse {
	s.Body = v
	return s
}

func (s *CreateQueueResponse) Validate() error {
	return dara.Validate(s)
}
