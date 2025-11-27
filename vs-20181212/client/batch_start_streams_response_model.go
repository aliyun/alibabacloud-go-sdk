// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartStreamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchStartStreamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchStartStreamsResponse
	GetStatusCode() *int32
	SetBody(v *BatchStartStreamsResponseBody) *BatchStartStreamsResponse
	GetBody() *BatchStartStreamsResponseBody
}

type BatchStartStreamsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchStartStreamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchStartStreamsResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchStartStreamsResponse) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchStartStreamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchStartStreamsResponse) GetBody() *BatchStartStreamsResponseBody {
	return s.Body
}

func (s *BatchStartStreamsResponse) SetHeaders(v map[string]*string) *BatchStartStreamsResponse {
	s.Headers = v
	return s
}

func (s *BatchStartStreamsResponse) SetStatusCode(v int32) *BatchStartStreamsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchStartStreamsResponse) SetBody(v *BatchStartStreamsResponseBody) *BatchStartStreamsResponse {
	s.Body = v
	return s
}

func (s *BatchStartStreamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
