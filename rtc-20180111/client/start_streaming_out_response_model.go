// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamingOutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartStreamingOutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartStreamingOutResponse
	GetStatusCode() *int32
	SetBody(v *StartStreamingOutResponseBody) *StartStreamingOutResponse
	GetBody() *StartStreamingOutResponseBody
}

type StartStreamingOutResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartStreamingOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartStreamingOutResponse) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutResponse) GoString() string {
	return s.String()
}

func (s *StartStreamingOutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartStreamingOutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartStreamingOutResponse) GetBody() *StartStreamingOutResponseBody {
	return s.Body
}

func (s *StartStreamingOutResponse) SetHeaders(v map[string]*string) *StartStreamingOutResponse {
	s.Headers = v
	return s
}

func (s *StartStreamingOutResponse) SetStatusCode(v int32) *StartStreamingOutResponse {
	s.StatusCode = &v
	return s
}

func (s *StartStreamingOutResponse) SetBody(v *StartStreamingOutResponseBody) *StartStreamingOutResponse {
	s.Body = v
	return s
}

func (s *StartStreamingOutResponse) Validate() error {
	return dara.Validate(s)
}
