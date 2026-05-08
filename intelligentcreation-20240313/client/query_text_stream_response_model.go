// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTextStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryTextStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryTextStreamResponse
	GetStatusCode() *int32
	SetId(v string) *QueryTextStreamResponse
	GetId() *string
	SetEvent(v string) *QueryTextStreamResponse
	GetEvent() *string
	SetBody(v *QueryTextStreamResponseBody) *QueryTextStreamResponse
	GetBody() *QueryTextStreamResponseBody
}

type QueryTextStreamResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                      `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                      `json:"event,omitempty" xml:"event,omitempty"`
	Body       *QueryTextStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryTextStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryTextStreamResponse) GoString() string {
	return s.String()
}

func (s *QueryTextStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryTextStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryTextStreamResponse) GetId() *string {
	return s.Id
}

func (s *QueryTextStreamResponse) GetEvent() *string {
	return s.Event
}

func (s *QueryTextStreamResponse) GetBody() *QueryTextStreamResponseBody {
	return s.Body
}

func (s *QueryTextStreamResponse) SetHeaders(v map[string]*string) *QueryTextStreamResponse {
	s.Headers = v
	return s
}

func (s *QueryTextStreamResponse) SetStatusCode(v int32) *QueryTextStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryTextStreamResponse) SetId(v string) *QueryTextStreamResponse {
	s.Id = &v
	return s
}

func (s *QueryTextStreamResponse) SetEvent(v string) *QueryTextStreamResponse {
	s.Event = &v
	return s
}

func (s *QueryTextStreamResponse) SetBody(v *QueryTextStreamResponseBody) *QueryTextStreamResponse {
	s.Body = v
	return s
}

func (s *QueryTextStreamResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
