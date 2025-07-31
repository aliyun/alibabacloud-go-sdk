// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUpDownStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceUpDownStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceUpDownStreamResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceUpDownStreamResponseBody) *GetInstanceUpDownStreamResponse
	GetBody() *GetInstanceUpDownStreamResponseBody
}

type GetInstanceUpDownStreamResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceUpDownStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceUpDownStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUpDownStreamResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceUpDownStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceUpDownStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceUpDownStreamResponse) GetBody() *GetInstanceUpDownStreamResponseBody {
	return s.Body
}

func (s *GetInstanceUpDownStreamResponse) SetHeaders(v map[string]*string) *GetInstanceUpDownStreamResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceUpDownStreamResponse) SetStatusCode(v int32) *GetInstanceUpDownStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceUpDownStreamResponse) SetBody(v *GetInstanceUpDownStreamResponseBody) *GetInstanceUpDownStreamResponse {
	s.Body = v
	return s
}

func (s *GetInstanceUpDownStreamResponse) Validate() error {
	return dara.Validate(s)
}
