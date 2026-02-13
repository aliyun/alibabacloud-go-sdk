// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AISearchStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AISearchStreamResponse
	GetStatusCode() *int32
	SetId(v string) *AISearchStreamResponse
	GetId() *string
	SetEvent(v string) *AISearchStreamResponse
	GetEvent() *string
	SetBody(v *AISearchStreamResponseBody) *AISearchStreamResponse
	GetBody() *AISearchStreamResponseBody
}

type AISearchStreamResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                     `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                     `json:"event,omitempty" xml:"event,omitempty"`
	Body       *AISearchStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AISearchStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s AISearchStreamResponse) GoString() string {
	return s.String()
}

func (s *AISearchStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AISearchStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AISearchStreamResponse) GetId() *string {
	return s.Id
}

func (s *AISearchStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *AISearchStreamResponse) GetBody() *AISearchStreamResponseBody {
	return s.Body
}

func (s *AISearchStreamResponse) SetHeaders(v map[string]*string) *AISearchStreamResponse {
	s.Headers = v
	return s
}

func (s *AISearchStreamResponse) SetStatusCode(v int32) *AISearchStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *AISearchStreamResponse) SetId(v string) *AISearchStreamResponse {
	s.Id = &v
	return s
}

func (s *AISearchStreamResponse) SetEvent(v string) *AISearchStreamResponse {
	s.Event = &v
	return s
}

func (s *AISearchStreamResponse) SetBody(v *AISearchStreamResponseBody) *AISearchStreamResponse {
	s.Body = v
	return s
}

func (s *AISearchStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
