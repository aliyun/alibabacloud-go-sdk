// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSubscribeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEventSubscribeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEventSubscribeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEventSubscribeResponseBody) *DeleteEventSubscribeResponse
	GetBody() *DeleteEventSubscribeResponseBody
}

type DeleteEventSubscribeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventSubscribeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSubscribeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEventSubscribeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEventSubscribeResponse) GetBody() *DeleteEventSubscribeResponseBody {
	return s.Body
}

func (s *DeleteEventSubscribeResponse) SetHeaders(v map[string]*string) *DeleteEventSubscribeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventSubscribeResponse) SetStatusCode(v int32) *DeleteEventSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventSubscribeResponse) SetBody(v *DeleteEventSubscribeResponseBody) *DeleteEventSubscribeResponse {
	s.Body = v
	return s
}

func (s *DeleteEventSubscribeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
