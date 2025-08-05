// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventStreamingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventStreamingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventStreamingResponse
	GetStatusCode() *int32
	SetBody(v *GetEventStreamingResponseBody) *GetEventStreamingResponse
	GetBody() *GetEventStreamingResponseBody
}

type GetEventStreamingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventStreamingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventStreamingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventStreamingResponse) GoString() string {
	return s.String()
}

func (s *GetEventStreamingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventStreamingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventStreamingResponse) GetBody() *GetEventStreamingResponseBody {
	return s.Body
}

func (s *GetEventStreamingResponse) SetHeaders(v map[string]*string) *GetEventStreamingResponse {
	s.Headers = v
	return s
}

func (s *GetEventStreamingResponse) SetStatusCode(v int32) *GetEventStreamingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventStreamingResponse) SetBody(v *GetEventStreamingResponseBody) *GetEventStreamingResponse {
	s.Body = v
	return s
}

func (s *GetEventStreamingResponse) Validate() error {
	return dara.Validate(s)
}
