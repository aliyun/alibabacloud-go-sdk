// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventResponseBody) *DeleteEventResponse
	GetBody() *DeleteEventResponseBody
}

type DeleteEventResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventResponse) GetBody() *DeleteEventResponseBody {
	return s.Body
}

func (s *DeleteEventResponse) SetHeaders(v map[string]*string) *DeleteEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventResponse) SetStatusCode(v int32) *DeleteEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventResponse) SetBody(v *DeleteEventResponseBody) *DeleteEventResponse {
	s.Body = v
	return s
}

func (s *DeleteEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
