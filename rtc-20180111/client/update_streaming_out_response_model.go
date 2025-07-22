// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamingOutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStreamingOutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStreamingOutResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStreamingOutResponseBody) *UpdateStreamingOutResponse
	GetBody() *UpdateStreamingOutResponseBody
}

type UpdateStreamingOutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStreamingOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStreamingOutResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutResponse) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStreamingOutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStreamingOutResponse) GetBody() *UpdateStreamingOutResponseBody {
	return s.Body
}

func (s *UpdateStreamingOutResponse) SetHeaders(v map[string]*string) *UpdateStreamingOutResponse {
	s.Headers = v
	return s
}

func (s *UpdateStreamingOutResponse) SetStatusCode(v int32) *UpdateStreamingOutResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStreamingOutResponse) SetBody(v *UpdateStreamingOutResponseBody) *UpdateStreamingOutResponse {
	s.Body = v
	return s
}

func (s *UpdateStreamingOutResponse) Validate() error {
	return dara.Validate(s)
}
