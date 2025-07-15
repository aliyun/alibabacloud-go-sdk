// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *CreateMessageAppResponseBody) *CreateMessageAppResponse
	GetBody() *CreateMessageAppResponseBody
}

type CreateMessageAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageAppResponse) GoString() string {
	return s.String()
}

func (s *CreateMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMessageAppResponse) GetBody() *CreateMessageAppResponseBody {
	return s.Body
}

func (s *CreateMessageAppResponse) SetHeaders(v map[string]*string) *CreateMessageAppResponse {
	s.Headers = v
	return s
}

func (s *CreateMessageAppResponse) SetStatusCode(v int32) *CreateMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMessageAppResponse) SetBody(v *CreateMessageAppResponseBody) *CreateMessageAppResponse {
	s.Body = v
	return s
}

func (s *CreateMessageAppResponse) Validate() error {
	return dara.Validate(s)
}
