// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageUserMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveMessageUserMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveMessageUserMessageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveMessageUserMessageResponseBody) *DeleteLiveMessageUserMessageResponse
	GetBody() *DeleteLiveMessageUserMessageResponseBody
}

type DeleteLiveMessageUserMessageResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveMessageUserMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveMessageUserMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageUserMessageResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageUserMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveMessageUserMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveMessageUserMessageResponse) GetBody() *DeleteLiveMessageUserMessageResponseBody {
	return s.Body
}

func (s *DeleteLiveMessageUserMessageResponse) SetHeaders(v map[string]*string) *DeleteLiveMessageUserMessageResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveMessageUserMessageResponse) SetStatusCode(v int32) *DeleteLiveMessageUserMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveMessageUserMessageResponse) SetBody(v *DeleteLiveMessageUserMessageResponseBody) *DeleteLiveMessageUserMessageResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveMessageUserMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
