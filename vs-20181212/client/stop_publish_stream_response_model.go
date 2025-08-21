// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPublishStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPublishStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPublishStreamResponse
	GetStatusCode() *int32
	SetBody(v *StopPublishStreamResponseBody) *StopPublishStreamResponse
	GetBody() *StopPublishStreamResponseBody
}

type StopPublishStreamResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPublishStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPublishStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPublishStreamResponse) GoString() string {
	return s.String()
}

func (s *StopPublishStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPublishStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPublishStreamResponse) GetBody() *StopPublishStreamResponseBody {
	return s.Body
}

func (s *StopPublishStreamResponse) SetHeaders(v map[string]*string) *StopPublishStreamResponse {
	s.Headers = v
	return s
}

func (s *StopPublishStreamResponse) SetStatusCode(v int32) *StopPublishStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPublishStreamResponse) SetBody(v *StopPublishStreamResponseBody) *StopPublishStreamResponse {
	s.Body = v
	return s
}

func (s *StopPublishStreamResponse) Validate() error {
	return dara.Validate(s)
}
