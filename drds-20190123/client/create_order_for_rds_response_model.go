// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrderForRdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrderForRdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrderForRdsResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrderForRdsResponseBody) *CreateOrderForRdsResponse
	GetBody() *CreateOrderForRdsResponseBody
}

type CreateOrderForRdsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderForRdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderForRdsResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrderForRdsResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderForRdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrderForRdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrderForRdsResponse) GetBody() *CreateOrderForRdsResponseBody {
	return s.Body
}

func (s *CreateOrderForRdsResponse) SetHeaders(v map[string]*string) *CreateOrderForRdsResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderForRdsResponse) SetStatusCode(v int32) *CreateOrderForRdsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderForRdsResponse) SetBody(v *CreateOrderForRdsResponseBody) *CreateOrderForRdsResponse {
	s.Body = v
	return s
}

func (s *CreateOrderForRdsResponse) Validate() error {
	return dara.Validate(s)
}
