// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPublishStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPublishStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPublishStreamResponse
	GetStatusCode() *int32
	SetBody(v *StartPublishStreamResponseBody) *StartPublishStreamResponse
	GetBody() *StartPublishStreamResponseBody
}

type StartPublishStreamResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPublishStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPublishStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPublishStreamResponse) GoString() string {
	return s.String()
}

func (s *StartPublishStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPublishStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPublishStreamResponse) GetBody() *StartPublishStreamResponseBody {
	return s.Body
}

func (s *StartPublishStreamResponse) SetHeaders(v map[string]*string) *StartPublishStreamResponse {
	s.Headers = v
	return s
}

func (s *StartPublishStreamResponse) SetStatusCode(v int32) *StartPublishStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPublishStreamResponse) SetBody(v *StartPublishStreamResponseBody) *StartPublishStreamResponse {
	s.Body = v
	return s
}

func (s *StartPublishStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
