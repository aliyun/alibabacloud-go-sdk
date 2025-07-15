// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMixStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMixStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMixStreamResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMixStreamResponseBody) *UpdateMixStreamResponse
	GetBody() *UpdateMixStreamResponseBody
}

type UpdateMixStreamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMixStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMixStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMixStreamResponse) GoString() string {
	return s.String()
}

func (s *UpdateMixStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMixStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMixStreamResponse) GetBody() *UpdateMixStreamResponseBody {
	return s.Body
}

func (s *UpdateMixStreamResponse) SetHeaders(v map[string]*string) *UpdateMixStreamResponse {
	s.Headers = v
	return s
}

func (s *UpdateMixStreamResponse) SetStatusCode(v int32) *UpdateMixStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMixStreamResponse) SetBody(v *UpdateMixStreamResponseBody) *UpdateMixStreamResponse {
	s.Body = v
	return s
}

func (s *UpdateMixStreamResponse) Validate() error {
	return dara.Validate(s)
}
