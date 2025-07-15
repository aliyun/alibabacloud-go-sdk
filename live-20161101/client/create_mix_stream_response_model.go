// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMixStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMixStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMixStreamResponse
	GetStatusCode() *int32
	SetBody(v *CreateMixStreamResponseBody) *CreateMixStreamResponse
	GetBody() *CreateMixStreamResponseBody
}

type CreateMixStreamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMixStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMixStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMixStreamResponse) GoString() string {
	return s.String()
}

func (s *CreateMixStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMixStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMixStreamResponse) GetBody() *CreateMixStreamResponseBody {
	return s.Body
}

func (s *CreateMixStreamResponse) SetHeaders(v map[string]*string) *CreateMixStreamResponse {
	s.Headers = v
	return s
}

func (s *CreateMixStreamResponse) SetStatusCode(v int32) *CreateMixStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMixStreamResponse) SetBody(v *CreateMixStreamResponseBody) *CreateMixStreamResponse {
	s.Body = v
	return s
}

func (s *CreateMixStreamResponse) Validate() error {
	return dara.Validate(s)
}
