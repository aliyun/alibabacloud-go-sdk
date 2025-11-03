// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventSourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventSourceResponseBody) *DeleteEventSourceResponse
	GetBody() *DeleteEventSourceResponseBody
}

type DeleteEventSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventSourceResponse) GetBody() *DeleteEventSourceResponseBody {
	return s.Body
}

func (s *DeleteEventSourceResponse) SetHeaders(v map[string]*string) *DeleteEventSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventSourceResponse) SetStatusCode(v int32) *DeleteEventSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventSourceResponse) SetBody(v *DeleteEventSourceResponseBody) *DeleteEventSourceResponse {
	s.Body = v
	return s
}

func (s *DeleteEventSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
